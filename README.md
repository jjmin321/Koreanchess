장기판 만들기 (사진)
장기 말들 만들기 (사진)
board 변수에 이차원 배열 사용해서 말들 배치 컴퓨터에게 알려주기 
좌표평면에 말들 그림을 그려서 넣어야함.
```go
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
```