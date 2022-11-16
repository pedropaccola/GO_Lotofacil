package main

import (
	"math/rand"
	"time"

	"github.com/pedropaccola/go-lotofacil/game"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	g := game.NewGame()
	g.Start()
}
