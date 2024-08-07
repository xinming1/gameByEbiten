package game

import (
	"game_by_ebiten/config"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Game struct {
	bgImg *ebiten.Image
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
	bgImg, _, err := ebitenutil.NewImageFromFile(config.Cfg.BgImg)
	if err != nil {
		log.Fatal(err)
	}
	game.bgImg = bgImg

	return game
}

func (g *Game) Update() error {
	//g.dog.Update()
	g.sword.Update()
	//g.goblin.Update()
	g.goblinManager.Update(g.sword.x, g.sword.y)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	i := config.Cfg.ScreenWidth / g.bgImg.Bounds().Dx()
	j := config.Cfg.ScreenHeight / g.bgImg.Bounds().Dy()
	op.GeoM.Scale(float64(i), float64(j))
	screen.DrawImage(g.bgImg, op)
	g.goblinManager.Draw(screen)
	g.sword.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.Cfg.ScreenWidth, config.Cfg.ScreenHeight
}
