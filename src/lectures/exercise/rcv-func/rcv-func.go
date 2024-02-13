//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	name              string
	health, maxHealth uint
	energy, maxEnergy uint
}

func (player *Player) reduceHealth(amount uint) {
	if amount >= player.health {
		player.health = 0
	} else {

		player.health -= amount
	}
}
func (player *Player) healPlayer(amount uint) {
	if amount >= player.maxHealth {
		player.health = player.maxHealth
	} else {
		player.health += amount
	}
}

func (player *Player) reduceEnergy(amount uint) {
	if amount >= player.energy {
		player.energy = 0
	} else {

		player.energy -= amount
	}
}
func (player *Player) energizePlayer(amount uint) {
	if amount >= player.maxHealth {
		player.energy = player.maxEnergy
	} else {
		player.energy += amount
	}
}

func main() {
	wizard := Player{name: "Gandalf", health: 1000, maxHealth: 1000, energy: 500, maxEnergy: 500}
	fmt.Println("Player created: ", wizard)

	wizard.reduceHealth(50)
	fmt.Println("Player attacked, the attack has a strength of 50, health is:", wizard.health, wizard)

	wizard.reduceHealth(5550)
	fmt.Println("Player attacked, the attack has a strength of 5550, health is:", wizard.health, wizard)

	wizard.healPlayer(250)
	fmt.Println("Player healed with 250, health is:", wizard.health, wizard)

	wizard.reduceEnergy(100)
	fmt.Println("Player used 100 points of energy, energy is", wizard.energy)

	wizard.energizePlayer(10)
	fmt.Println("Player received 10 points of energy, energy is", wizard.energy)

	fmt.Println("Overflow cases ///////")
	wizard.reduceEnergy(500000)
	fmt.Println("Player used 500000 points of energy, energy is", wizard.energy)

	wizard.energizePlayer(111110)
	fmt.Println("Player received 111110 points of energy, energy is", wizard.energy)

}
