package game

import (
	"game_by_ebiten/config"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	dog    *Dog
	sword  *Sword
	goblin *Goblin
}

func NewGame() *Game {
	return &Game{
		//dog:   NewDog(config.Cfg.Dog),
		sword:  NewSword(config.Cfg.SwordConfig),
		goblin: NewGoblin(config.Cfg.GoblinConfig),
	}
}

func (g *Game) Update() error {
	//g.dog.Update()
	g.sword.Update()
	g.goblin.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//g.dog.Draw(screen)
	g.sword.Draw(screen)
	g.goblin.Draw(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.Cfg.ScreenWidth, config.Cfg.ScreenHeight
}
