package main

import (
	"context"
	"crypto/rand"
	"crypto/tls"
	host "github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/routing"
	discovery "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/p2p/net/connmgr"
	libp2ptls "github.com/libp2p/go-libp2p/p2p/security/tls"
	"github.com/libp2p/go-libp2p/p2p/transport/tcp"
	"github.com/libp2p/go-yamux/v4"
	"github.com/multiformats/go-multiaddr"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

const service = "LaSalleCollege"
//
The code represents a package that sets up a peer-to-peer (P2P) network using the libp2p library.
It contains a P2P struct, which holds the context, host, Kademlia Distributed Hash Table (KadDHT), routing discovery,
and publisher-subscriber (PubSub) components.

The newP2P() function initializes and binds these components together.
It first sets up the host node and the DHT using SetupHost(), then binds the three components with the bootstrapDHT()
function. The routing discovery and PubSub handlers are then activated.
EnterClientIntoChatRoom() is used to advertise the service to the network,
which allows a client to enter the chat room. Finally, AnnounceConnect() generates an ID for the client and
makes an announcement on the network for it to connect with other peers.

The code uses logrus for logging and is well-documented, with each function's purpose being described in detail.
The SetupHost(), bootstrapDHT(), and setupPubSub() functions provide detailed explanations and error handling.

The newP2P() function initializes the required P2P components and binds them
together to create a functional P2P network.
The EnterClientIntoChatRoom() function is used to let a client come in and start chatting,
while the AnnounceConnect() function allows the announcement of the client to connect with other peers.

type P2P struct {
	Ctx       context.Context
	Host      host.Host
	KadDHT    *dht.IpfsDHT
	Discovery *discovery.RoutingDiscovery
	PubSub    *pubsub.Pubsub
}

func newP2P() *P2P {
	ctx := context.Background() // To InitializeComponents

	// Set up the host node, in fact is a host and one dht

	nodehost, kaddht := SetupHost(ctx)
	logrus.Debugln("We just created the P2P host and the DHT")

	//Binding the three adove-mentioned components together

	bootstrapDHT(ctx, nodehost, kaddht)
	logrus.Debugln("Bootstraping (binding parameters) is done!")

	//Activate routing discovery

	routingDiscovery := discovery.NewRoutingDiscovery(kaddht)
	logrus.Debudln("Routing discovery service is now activate!")

	// Activating the pubsub

	pubsubhandler := setupPubSub(ctx, nodehost, routingDiscovery)
	logrus.Debugln("PubSub handler (publisher handler) is now activate")

	return &P2P{
		Ctx:       ctx,
		Host:      nodehost,
		KadDHT:    kaddht,
		Discovery: routingDiscovery,
		PubSub:    pubsubhandler,
	}

}

func (p2p *P2P) EnterClientIntoChatRoom() {
	// Call the discovery function to let a client come in

	ttl, err := p2p.Discovery.Advertise(p2p.Ctx, service)
	logrus.Debugln("Discovery function verified client to come in")
	time.Sleep(time.Second * 10)
	logrus.Debugf("Time out for title: %s", ttl)

	peerChan, err := p2p.Discovery.FindPeers(p2p.Ctx, service)

	if err != nil {
		logrus.WithFields(
			logrus.Fields{
				"Wow Error": err.Error(),
			},
		).Fatalln("A Fatal error happened")
	}

	logrus.Traceln("Print the trace of the new client into the history of chatroom")
	go handlePeerDiscovery(p2p.Host, peerChan)
	logrus.Traceln("Voila, start talking")
}

func (p2p *P2P) AnnounceConnect() {
	cidvalue := generateCID(service)
	logrus.Debugln("An ID has been generated for the client")
	err := p2p.KadDHT.Provide(p2p.Ctx, cidvalue, true)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Error": err.Error(),
		}).Fatalln("Failed to anounce service CID")
	}
	peerChan := p2p.KadDHT.FindProvidersAsync(p2p.Ctx, cidvalue, 0)
	go hanlePeerDiscovery(p2p.Host, peerChan)
	logrus.Debugln("Peer connection is being handled!")
}

//
//The code is written in Golang and sets up a libp2p host and a Kademlia DHT, which are used to enable peer-to-peer
// communication and content discovery in decentralized applications.
//The setupHost function generates a private key, sets up a TLS transport with a TCP transport, and configures various
//options for the libp2p host, including the use of the YAMUX multiplexing protocol, connection management,
//NAT port mapping, and relaying. The setupKadDHT function sets up a Kademlia DHT using the given libp2p
//host and bootstrap peers, and returns a pointer to the DHT.
//
//In general, a DHT is a distributed hash table that enables content discovery by storing key-value pairs across a
//network of nodes, while libp2p is a modular networking stack for building peer-to-peer applications.
//Together, a DHT and a libp2p host enable decentralized applications to find and communicate
//with each other in a peer-to-peer network, without relying on centralized servers.
//This makes them useful for building decentralized applications that require secure and scalable communication,
//such as blockchain-based systems, file-sharing networks, and messaging apps.

func setupHost(ctx context.Context) (host.Host, *dht.IpfsDHT) {

	key, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, rand.Reader)
	identity := libp2ptls.Identity(key)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Error": err.Error(),
		}).Fatalln("Failed to generate an identity for the client!")
	}
	tlstransport, err := tls.New(key)
	security := libp2p.Security(tls.ID, tlstransport)
	transport := libp2ptls.Transport(tcp.NewTCPTransport)
	logrus.Debugln("TLS and security and transport are set successfully")
	muladdr, err := multiaddr.NewMultiaddr("/ip4/0.0.0.0./tcp/0")
	listen := libp2p.ListenAddrs(muladdr)
	muxer := libp2p.Muxer("yamux/1.0.0", yamux.DefaultTransport)
	conn := libp2p.ConnectionManager(connmgr.NewConnManager(100, 400, time.Minute))
	nat := libp2p.NATPortMap()
	relay := libp2p.EnableRelay() // Repeat something you have
	var kaddht *dht.IpfsDHT
	routing := libp2p.Routing(func(h host.Host) (routing.PeerRouting,error)){
		kaddht = setupKadDHT(ctx,h)
		return kaddht, err
	})
	opts := libp2p.ChainOptions(identity, listen, security, transport, muxer, conn, nat, routing., relay)
	libhost, err := libp2p.New(ctx,opts)
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err.Error(),
		}).Fataln("Failed out create the P2P host")
	}
}

func setupKadDHT(ctx context.Context, nodehost host.Host) *dht.IpfsDHT{
	dhtmode := dht.Mode(dht.ModeServer)
	bootstrapPeers := dht.GetDefaultBootstrapPeerAddrInfos()
	dhtpeers := dht.BootstrapPeers(bootstrapPeers...)
	logrus.Degugln("KadDHT is up and renning!")
	kaddht, err := dht.New(ctx, nodehost, dhtmode, dhtpeers)
	if err != nil{
		logrus.WithFields(logrus.Fields{
			"Error": err.Error(),
		}).Fatalln("Failed to create Kademlia DHT")
	}
	return kaddht
}
//
//The two functions build on the earlier functions to add more functionality to the P2P network.
//The setupPubSub function sets up a pubsub mechanism for communication between nodes in the network,
//using the GossipSub implementation of pubsub. The pubsub.WithDiscovery option is used to enable
//discovery of peers in the network.
//
//The bootstrapDHT function bootstraps the DHT and connects to a set of default bootstrap peers.
//It first calls the kaddht.Bootstrap method to get the list of known bootstrap peers for the DHT.
//Then, it loops through the list of default bootstrap peers, and connects to each peer using the nodehost.Connect method.
//It tracks the number of connected bootstrap peers and logs the result.
//
//This bootstrap process is important because it helps to discover and connect to other peers in the network,
//which enables the DHT to function properly.
//
//Both functions are essential for building a functional decentralized application
//that requires peer-to-peer communication and discovery of other nodes in the network.
//The setupPubSub function allows for message broadcasting and subscription among peers in the network,
//while the bootstrapDHT function ensures that the DHT is properly set up and connected to other peers in
//the network for efficient content discovery.

func setupPubSub(ctx context.Context, nodehost host.Host, routingdiscovery *discovery.RotingDiscovery) *pubsub.PubSub{

	pubsubhandler, err := pubsub.NewGossipSub(ctx, nodehost, pubsub.WithDiscovery(routingdiscovery))
	if err != nil{
		logrus.WithFields(logrus, Fields{
			"error": err.Error(),
			"type": "Gossipsub",
		}).Fatalln("PubSub Handler Creation Failed")
	}
}


func bootstrapDHT(ctx context.Context, nodehost host.Host, kaddht *dht.IpfsDHT){

	//Agregation
	if err := kaddht.Bootstrap(ctx); err != nil{
		logrus.WithFields(logrus.Fields{
			"Error": err.Error(),
		}).Fatalln("Failed to Bootstrap kaddht")
	}
	var wg sync.WaitGroup
	var connectedBootPeers int
	var totalBootPeers int
	for _, peeraddr := range dht.DefaultBootstrapPeers{
		peerinfo, _ := peer.AddrInfoFromP2pAddr(peeraddr)
		wg.Add(1)
		go func(){
			defer wg.Done()
			if err := nodehost.Connect(ctx, *peerinfo); err != nil{
				totalBootPeers++
			}else{
				connectedBootPeers++
				totalBootPeers++
			}
		}()
	}
	wg.Wait()
	logrus.Debugf("Connected to %d out of %d Bootstrat Peers"),
	connectedBootPeers, totalBootPeers)
}





func main() {

}
