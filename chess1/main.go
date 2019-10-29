package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	bgimg *ebiten.Image
)

func main() {
	var err error
	bgimg, _, err = ebitenutil.NewImageFromFile("./main.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	err = ebiten.Run(func(*ebiten.Image) error {
		return nil
	}, 300, 400, 1.0, "12 chess")

	if err != nil {
		log.Fatalf("Ebiten run error: %v", err)
	}
}
