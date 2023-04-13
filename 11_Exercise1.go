package main

import "fmt"

type Player struct {
	health int
}

func (player *Player) takeDamageFromExplosion(dmg int) {
	fmt.Println("Player is taking damage from explosion")
	player.health -= dmg
}

func takeDamageFromExplosion(player *Player, dmg int) {
	fmt.Println("Player is taking damage from explosion")
	player.health -= dmg
}

func main() {

	player := &Player{100}
	player2 := &Player{120}

	fmt.Printf("Before Explosion Player 1 %+v\n", player)
	fmt.Printf("Before Explosion Player 2 %+v\n", player)
	player.takeDamageFromExplosion(50)
	takeDamageFromExplosion(player2, 20)
	fmt.Printf("After explosion Player 1%+v\n", player)
	fmt.Printf("After explosion Player 2%+v\n", player)

}
