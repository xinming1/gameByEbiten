package game

import (
	"game_by_ebiten/config"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	dog *Dog
}

func NewGame() *Game {
	return &Game{
		dog: NewDog(config.Cfg.Dog),
	}
}

func (g *Game) Update() error {
	g.dog.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.dog.Draw(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.Cfg.ScreenWidth, config.Cfg.ScreenHeight
}
