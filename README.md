# KoreanChess

### What I learned and other thing, etc...
- Code refactoring
- how to make 2D game with golang
- Importance of const
- Type aliasing
# 2D 게임을 구현시킬 수 있게 해주는 golang 패키지 다운받기 
- github.com/hajimehoshi/ebiten
- github.com/hajimehoshi/ebiten/ebitenutil

# 함수에서 사용할 상수, 변수 선언
```go
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
- board 이차원 배열에 이니셜을 나눠주어 구분시킨다.
```go
//green
	board[0][0] = GimulTypeGreenSang 
	board[0][1] = GimulTypeGreenWang 
	board[0][2] = GimulTypeGreenJang 
	board[1][1] = GimulTypeGreenJa   

	//red
	board[2][1] = GimulTypeRedJa    
	board[3][0] = GimulTypeRedSang   
	board[3][1] = GimulTypeRedWang 
	board[3][2] = GimulTypeRedJang 
```
- 이 코드가 어렵게 보일수도 있겠지만 사실상 이 코드와 같은 의미이다.

```go
    board[0][0] = 3 //board[0][0]을 3이라고 부르겠다.
	board[0][1] = 0 //board[0][1]을 0이라고 부르겠다.
	board[0][2] = 1 //board[0][2]를 1이라고 부르겠다.
	board[1][1] = 2 //board[1][1]를 2라고 부르겠다.

	//red
	board[2][1] = 7 //board[2][1]을 7이라고 부르겠다.
	board[3][0] = 4 //board[3][0]를 4라고 부르겠다.
	board[3][1] = 5 //board[3][1]를 5라고 부르겠다.
	board[3][2] = 6 //board[3][2]를 6이라고 부르겠다.
```
# 좌표평면에 말들 그림을 그려서 넣기

- 장기판에 장기말들의 배치만 그림을 넣으니 오류가 생겼다. 장기말들이 없어야할 곳도 생기는 것이였다.
- 그래서 board의 모든 값을 GimulTypeNone으로 먼저 초기화시키고 switch case문을 사용했다.

```go
// Initialize board
	for i := 0; i < BoardWidth; i++ {
		for j := 0; j < BoardHeight; j++ {
			board[i][j] = GimulTypeNone
		}
	}
``` 
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
