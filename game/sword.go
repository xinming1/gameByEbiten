package game

import (
	"game_by_ebiten/config"
	"game_by_ebiten/logger"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"math"
)

type Sword struct {
	imgW, imgH     float64
	x, y           float64
	image          *ebiten.Image
	theta          float64
	drawOp         *ebiten.DrawImageOptions
	killNum        int
	revolveSpeed   float64
	sizeMultiplier float64
}

func NewSword(swordConfig *config.SwordConfig) *Sword {
	img, _, err := ebitenutil.NewImageFromFile(swordConfig.Img)
	if err != nil {
		log.Fatal(err)
	}
	sword := &Sword{
		imgW:           float64(img.Bounds().Dx()),
		imgH:           float64(img.Bounds().Dy()),
		image:          img,
		drawOp:         &ebiten.DrawImageOptions{},
		revolveSpeed:   1.0,
		sizeMultiplier: 1.0,
	}
	sword.x = (float64(config.Cfg.ScreenWidth) - sword.imgW) / 2
	sword.y = (float64(config.Cfg.ScreenHeight) - sword.imgH) / 2
	//go sword.revolve()
	return sword
}

func (sword *Sword) revolve() {

	sword.theta += math.Pi / 100
	if sword.theta > 2*math.Pi {
		sword.theta = 0.0
	}

}

func (sword *Sword) upgrade() {
	if sword.revolveSpeed == 10 {
		return
	}
	if sword.killNum > 3 {
		sword.revolveSpeed += 0.3
		sword.sizeMultiplier += 0.3
		dx := float64(sword.image.Bounds().Dx())
		sword.imgW = sword.sizeMultiplier * dx
		sword.imgH = sword.sizeMultiplier * dx
		sword.killNum = 0
	}
}

func (sword *Sword) Update() {
	sword.revolve()
	sword.upgrade()
	x, y := ebiten.CursorPosition()
	sword.x = float64(x) - sword.imgW/2
	sword.y = float64(y) - sword.imgH
}
func (sword *Sword) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-sword.imgW/2.0, -sword.imgH)
	op.GeoM.Rotate(sword.theta)
	op.GeoM.Translate(sword.imgW/2.0, sword.imgH)
	op.GeoM.Scale(sword.sizeMultiplier, sword.sizeMultiplier)

	logger.Debug("sword x:%f,y:%f", sword.x, sword.y)
	op.GeoM.Translate(sword.x, sword.y)
	sword.drawOp = op
	screen.DrawImage(sword.image, sword.drawOp)
}
