package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type GimulType int

// Type aliasing

const (
	GimulTypeNone GimulType = -1 + iota
	GimulTypeGreenWang
	GimulTypeGreenJa
	GimulTypeGreenJang
	GimulTypeGreenSang
	GimulTypeRedWang
	GimulTypeRedJa
	GimulTypeRedJang
	GimulTypeRedSang
	GimulTypeMax
)

//when you no need to value == only need to tell apart
//You can use iota grammer.
//iota in const start at 0 -> next const : 1 -> next next const : 2 -> next next next const: 3

var (
	board     [4][3]GimulType
	bgimg     *ebiten.Image
	gimulImgs [GimulTypeMax]*ebiten.Image
)

func update(screen *ebiten.Image) error {
	screen.DrawImage(bgimg, nil)

	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			switch board[i][j] {
			case GimulTypeGreenWang:
				//Draw GreenWang
				screen.DrawImage(gimulImgs[GimulTypeGreenWang], nil)
			case GimulTypeGreenJa:
				//Draw GreenJa
				screen.DrawImage(gimulImgs[GimulTypeGreenJa], nil)
			case GimulTypeGreenJang:
				//Draw GreenJang
			case GimulTypeGreenSang:
				//Draw GreenSang
			case GimulTypeRedWang:
				//Draw RedWang
			case GimulTypeRedJa:
				//Draw RedJa
			case GimulTypeRedJang:
				//Draw RedJang
			case GimulTypeRedSang:
				//Draw RedSang
			}
		}
	}
	return nil
}

func main() {
	var err error
	bgimg, _, err = ebitenutil.NewImageFromFile("./images/bgimg.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	gimulImgs[GimulTypeGreenWang], _, err = ebitenutil.NewImageFromFile("./images/green_wang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	gimulImgs[GimulTypeGreenJa], _, err = ebitenutil.NewImageFromFile("./images/green_ja.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	gimulImgs[GimulTypeGreenJang], _, err = ebitenutil.NewImageFromFile("./images/green_jang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	gimulImgs[GimulTypeGreenSang], _, err = ebitenutil.NewImageFromFile("./images/green_sang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	// Initialize board

	//green
	board[0][0] = GimulTypeGreenSang
	board[0][1] = GimulTypeGreenWang
	board[0][2] = GimulTypeGreenJang
	board[1][1] = GimulTypeGreenJa

	//red
	board[2][1] = GimulTypeRedJa
	board[3][0] = GimulTypeGreenSang
	board[3][1] = GimulTypeGreenWang
	board[3][2] = GimulTypeGreenJang

	err = ebiten.Run(update, 500, 400, 1.0, "12 chess")

	if err != nil {
		log.Fatalf("Ebiten run error: %v", err)
	}
}
