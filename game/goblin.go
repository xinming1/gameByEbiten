package game

import (
	"fmt"
	"game_by_ebiten/config"
	"game_by_ebiten/definition"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Goblin struct {
	x, y      float64
	w, h      float64
	speed     float64
	runImage  []*ebiten.Image
	idx       int
	direction int
}

func NewGoblin(c *config.GoblinConfig) *Goblin {

	var imgList = make([]*ebiten.Image, 0)
	for i := 0; i < c.ImgNum; i++ {
		tmpImg := fmt.Sprintf(c.Img, i)
		img, _, err := ebitenutil.NewImageFromFile(tmpImg)
		if err != nil {
			log.Fatal(err)
		}

		imgList = append(imgList, img)
	}

	goblin := &Goblin{
		w:        float64(imgList[0].Bounds().Dx()),
		h:        float64(imgList[0].Bounds().Dy()),
		runImage: imgList,
		idx:      0,
		speed:    c.Speed,
	}
	goblin.x = (float64(config.Cfg.ScreenWidth) - goblin.w) / 2
	goblin.y = (float64(config.Cfg.ScreenWidth) - goblin.h) / 2
	return goblin
}

func (goblin *Goblin) Update() {
	goblin.idx++
	if goblin.idx == config.Cfg.GoblinConfig.ImgNum-1 {
		goblin.idx = 0
	}
	goblin.Run()

}
func (goblin *Goblin) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	if goblin.direction != definition.Left {
		op.GeoM.Translate(goblin.x+(goblin.w)/2, 0)
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(-(goblin.x + (goblin.w)/2), 0)
	}

	op.GeoM.Translate(goblin.x, goblin.y)

	screen.DrawImage(goblin.runImage[goblin.idx], op)
}

func (goblin *Goblin) Run() {
	x, y := ebiten.CursorPosition()
	if goblin.x < float64(x) {
		goblin.x += goblin.speed
		goblin.direction = definition.Left
	} else {
		goblin.x -= goblin.speed
		goblin.direction = definition.Right
	}
	if goblin.y < float64(y) {
		goblin.y += goblin.speed
	} else {
		goblin.y -= goblin.speed
	}

}
