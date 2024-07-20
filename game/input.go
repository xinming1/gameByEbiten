package game

import "github.com/hajimehoshi/ebiten/v2"

type Input struct {
	msg string
}

func (i *Input) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		i.msg = "left pressed"
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		i.msg = "right pressed"
	} else if ebiten.IsKeyPressed(ebiten.KeySpace) {
		i.msg = "space pressed"
	}
}
