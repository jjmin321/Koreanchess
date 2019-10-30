# 2D 게임을 구현시킬 수 있게 해주는 golang 패키지 다운받기 
- github.com/hajimehoshi/ebiten
- github.com/hajimehoshi/ebiten/ebitenutil

# 함수에서 사용할 상수, 변수 선언
```go
type GimulType int

// Type aliasing

const (
	GimulTypeNone GimulType = -1 + iota		
	GimulTypeGreenWang //0
	GimulTypeGreenJa //1
	GimulTypeGreenJang //2
	GimulTypeGreenSang //3
	GimulTypeRedWang //4
	GimulTypeRedJa //5
	GimulTypeRedJang //6
	GimulTypeRedSang //7
	GimulTypeMax //8
)

//when you no need to value == only need to tell apart
//You can use iota grammer.
//iota in const start at 0 -> next const : 1 -> next next const : 2 -> next next next const: 3

var (
	board     [4][3]GimulType //크기가 [4][3]인 GimulType 타입 board 배열을 선언한다.
	bgimg     *ebiten.Image //*ebiten.Image 타입 bgimg를 선언한다. 
	gimulImgs [GimulTypeMax]*ebiten.Image //크기가 GimulTypeMax인 *ebiten.Image 타입 gimulImgs 배열을 선언한다.
)
```
# 장기 메인 만들기 
```go
err = ebiten.Run(update, 500, 400, 1.0, "12 chess")

	if err != nil {
		log.Fatalf("Ebiten run error: %v", err)
	}
```
### 장기판 만들기 (사진)
```go
var err error
	bgimg, _, err = ebitenutil.NewImageFromFile("./images/bgimg.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
```
### 장기 말들 만들기 (사진)
```go
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
```
### board 변수에 이차원 배열 사용해서 말들 배치 컴퓨터에게 알려주기 
```go
// Initialize board

	//green
	board[0][0] = GimulTypeGreenSang //3
	board[0][1] = GimulTypeGreenWang //0
	board[0][2] = GimulTypeGreenJang //1
	board[1][1] = GimulTypeGreenJa //2

	//red
	board[2][1] = GimulTypeRedJa //7
	board[3][0] = GimulTypeRedSang //4
	board[3][1] = GimulTypeRedWang //5
	board[3][2] = GimulTypeRedJang //6
```
### 좌표평면에 말들 그림을 그려서 넣는거
```go
func update(screen *ebiten.Image) error {
	screen.DrawImage(bgimg, nil)

	S := 20
	W := 116

	T := 23
	H := 116
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {

			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(S+W*i), float64(T+H*j)) //x좌표, y좌표
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
```