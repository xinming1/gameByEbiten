package game

import (
	"game_by_ebiten/config"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	dog   *Dog
	sword *Sword
	//goblin        *Goblin
	goblinManager *GoblinManager
}

func NewGame() *Game {
	game := &Game{
		//dog:   NewDog(config.Cfg.Dog),
		sword: NewSword(config.Cfg.SwordConfig),
	}
	game.goblinManager = NewGoblinManager(game)
	return game
}

func (g *Game) Update() error {
	//g.dog.Update()
	g.sword.Update()
	//g.goblin.Update()
	g.goblinManager.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//g.dog.Draw(screen)
	g.sword.Draw(screen)
	//g.goblin.Draw(screen)
	g.goblinManager.Draw(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.Cfg.ScreenWidth, config.Cfg.ScreenHeight
}
