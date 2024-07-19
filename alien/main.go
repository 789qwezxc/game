package main

import (
	"alien/model"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	err := ebiten.RunGame(model.NewGame())
	if err != nil {
		log.Fatal("%v", err)
	}
}