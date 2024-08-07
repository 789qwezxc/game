package model

import (
	"alien/config"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "golang.org/x/image/bmp"
)

type Ship struct {
	GameObj
	image *ebiten.Image
}

func NewShip(screenWidth, screenHeight int) *Ship {
	image, _, err := ebitenutil.NewImageFromFile("./ship.bmp")
	if err != nil {
		log.Fatalf("%v", err)
		fmt.Println("解析图片失败")
	}
	width, height := image.Bounds().Dx(), image.Bounds().Dy()
	s := &Ship{
		image: image,
	}
	s.GameObj.width = width
	s.GameObj.height = height
	s.GameObj.x = screenWidth / 2
	s.GameObj.y = screenHeight - height
	return s
}

func (ship *Ship) Draw(screen *ebiten.Image, cfg *config.Config) {
	// draw by self
	op := &ebiten.DrawImageOptions{}
	//init ship at the screen center
	op.GeoM.Translate(float64(ship.X()), float64(ship.Y()))
	screen.DrawImage(ship.image, op)
}