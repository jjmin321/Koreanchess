장기판 만들기 (사진)
장기 말들 만들기 (사진)
board 변수에 이차원 배열 사용해서 말들 배치 컴퓨터에게 알려주기 
# 좌표평면에 말들 그림을 그려서 넣는거
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
				screen.DrawImage(gimulImgs[GimulTypeGreenJang], nil)
			case GimulTypeGreenSang:
				//Draw GreenSang
				screen.DrawImage(gimulImgs[GimulTypeGreenSang], nil)
			case GimulTypeRedWang:
				//Draw RedWang
				screen.DrawImage(gimulImgs[GimulTypeRedWang], nil)
			case GimulTypeRedJa:
				//Draw RedJa
				screen.DrawImage(gimulImgs[GimulTypeRedJa], nil)
			case GimulTypeRedJang:
				//Draw RedJang
				screen.DrawImage(gimulImgs[GimulTypeRedJang], nil)
			case GimulTypeRedSang:
				//Draw RedSang
				screen.DrawImage(gimulImgs[GimulTypeRedSang], nil)
			}
		}
	}
	return nil
}
```