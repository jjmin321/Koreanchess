package main //40분

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type GimulType int

// Type aliasing

const (
	GimulTypeNone      GimulType = -1 + iota
	GimulTypeGreenWang           //0
	GimulTypeGreenJa             //1
	GimulTypeGreenJang           //2
	GimulTypeGreenSang           //3
	GimulTypeRedWang             //4
	GimulTypeRedJa               //5
	GimulTypeRedJang             //6
	GimulTypeRedSang             //7
	GimulTypeMax                 //8
)

const (
	GimulStartX = 20  //starting coordinate X
	GimulStartY = 23  //starting coordinate Y
	GridWidth   = 116 //한 칸의 가로
	GridHeight  = 116 //한 칸의 세로
	BoardWidth  = 4   // Constant declaration : easy to change value.
	BoardHeight = 3   // Constant declaration : easy to change value.
)

//when you no need to value == only need to tell apart
//You can use iota grammer.
//iota in const start at 0 -> next const : 1 -> next next const : 2 -> next next next const: 3

var (
	board     [BoardWidth][BoardHeight]GimulType //크기가 [4][3]인 GimulType 타입 board 배열을 선언한다.
	bgimg     *ebiten.Image                      //*ebiten.Image 타입 bgimg를 선언한다.
	gimulImgs [GimulTypeMax]*ebiten.Image        //크기가 GimulTypeMax인 *ebiten.Image 타입 gimulImgs 배열을 선언한다.
)

func update(screen *ebiten.Image) error {
	screen.DrawImage(bgimg, nil)

	for i := 0; i < BoardWidth; i++ {
		for j := 0; j < BoardHeight; j++ {

			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(GimulStartX+GridWidth*i), float64(GimulStartY+GridHeight*j)) //x좌표, y좌표
			switch board[i][j] {
			case GimulTypeGreenWang:
				//Draw GreenWang
				screen.DrawImage(gimulImgs[GimulTypeGreenWang], opts)
			case GimulTypeGreenJa:
				//Draw GreenJa
				screen.DrawImage(gimulImgs[GimulTypeGreenJa], opts)
			case GimulTypeGreenJang:
				//Draw GreenJang
				screen.DrawImage(gimulImgs[GimulTypeGreenJang], opts)
			case GimulTypeGreenSang:
				//Draw GreenSang
				screen.DrawImage(gimulImgs[GimulTypeGreenSang], opts)
			case GimulTypeRedWang:
				//Draw RedWang
				screen.DrawImage(gimulImgs[GimulTypeRedWang], opts)
			case GimulTypeRedJa:
				//Draw RedJa
				screen.DrawImage(gimulImgs[GimulTypeRedJa], opts)
			case GimulTypeRedJang:
				//Draw RedJang
				screen.DrawImage(gimulImgs[GimulTypeRedJang], opts)
			case GimulTypeRedSang:
				//Draw RedSang
				screen.DrawImage(gimulImgs[GimulTypeRedSang], opts)
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
	gimulImgs[GimulTypeRedWang], _, err = ebitenutil.NewImageFromFile("./images/red_wang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	gimulImgs[GimulTypeRedJa], _, err = ebitenutil.NewImageFromFile("./images/red_ja.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	gimulImgs[GimulTypeRedJang], _, err = ebitenutil.NewImageFromFile("./images/red_jang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	gimulImgs[GimulTypeRedSang], _, err = ebitenutil.NewImageFromFile("./images/red_sang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	// Initialize board

	//green
	board[0][0] = GimulTypeGreenSang //board[0][0]을 3이라고 부르겠다.
	board[0][1] = GimulTypeGreenWang //board[0][1]을 0이라고 부르겠다.
	board[0][2] = GimulTypeGreenJang //board[0][2]를 1이라고 부르겠다.
	board[1][1] = GimulTypeGreenJa   //board[1][1]를 2라고 부르겠다.

	//red
	board[2][1] = GimulTypeRedJa   //board[2][1]을 7이라고 부르겠다.
	board[3][0] = GimulTypeRedSang //board[3][0]를 4라고 부르겠다.
	board[3][1] = GimulTypeRedWang //board[3][1]를 5라고 부르겠다.
	board[3][2] = GimulTypeRedJang //board[3][2]를 6이라고 부르겠다.

	err = ebiten.Run(update, 500, 400, 1.0, "12 chess")

	if err != nil {
		log.Fatalf("Ebiten run error: %v", err)
	}
}
