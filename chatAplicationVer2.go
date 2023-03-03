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

type P2P struct {
	Ctx       context.Context
	Host      host.Host
	KadDHT    *dht.IpfsDHT
	Discovery *discovery.RoutingDiscovery
	PubSub    *pubsub.Pubsub
}

func newP2P() *P2P {
	ctx := context.Background() // To InitializeComponents

	nodehost, kaddht := SetupHost(ctx)
	logrus.Debugln("We just created the P2P host and the DHT")

	bootstrapDHT(ctx, nodehost, kaddht)
	logrus.Debugln("Bootstraping (binding parameters) is done!")

	routingDiscovery := discovery.NewRoutingDiscovery(kaddht)
	logrus.Debudln("Routing discovery service is now activate!")

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
