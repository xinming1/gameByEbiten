package game

import (
	"game_by_ebiten/config"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"math"
)

type Sword struct {
	width, height float64
	x, y          float64
	image         *ebiten.Image
	theta         float64
}

func NewSword(swordConfig *config.SwordConfig) *Sword {
	img, _, err := ebitenutil.NewImageFromFile(swordConfig.Img)
	if err != nil {
		log.Fatal(err)
	}
	sword := &Sword{
		width:  float64(img.Bounds().Dx()),
		height: float64(img.Bounds().Dy()),
		image:  img,
	}
	sword.x = (float64(config.Cfg.ScreenWidth) - sword.width) / 2
	sword.y = (float64(config.Cfg.ScreenHeight) - sword.height) / 2
	//go sword.revolve()
	return sword
}

func (sword *Sword) revolve() {

	sword.theta += math.Pi / 100
	if sword.theta > 2*math.Pi {
		sword.theta = 0.0
	}

}

func (sword *Sword) Update() {
	sword.revolve()
	x, y := ebiten.CursorPosition()
	sword.x = float64(x) - sword.width/2
	sword.y = float64(y) - sword.height
}
func (sword *Sword) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	//op.GeoM.Scale(0.7, 0.7)
	op.GeoM.Translate(-sword.width/2.0, -sword.height)
	op.GeoM.Rotate(sword.theta)
	op.GeoM.Translate(sword.width/2.0, sword.height)

	op.GeoM.Translate(sword.x, sword.y)

	screen.DrawImage(sword.image, op)
}
