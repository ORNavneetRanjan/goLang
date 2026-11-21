package main

import (
	"fmt"
	"strconv"
)

type Game interface {
	GetName() string
}

type Games struct {
	ID   int
	Name string
}

type BattleGround struct {
	ID int
	Name string
	Players int
}

func (g Games) GetName() string {
	return g.Name
}

func (g BattleGround) GetName() string {
	return g.Name + " Players: " + strconv.Itoa(g.Players)
}

func displayGames(p Game) {
	fmt.Println(p.GetName())
}

func main2() {
	gta := Games {
		ID: 1,
		Name: "GTA",
	}

	pubg := BattleGround {
		ID: 2, 
		Name: "PUBG",
		Players: 64,
	}
	displayGames(pubg)
	displayGames(gta) 
}
