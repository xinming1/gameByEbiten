package main

import (
	"game_by_ebiten/config"
	"game_by_ebiten/game"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	config.LoadConfig("./config/config.json")
	ebiten.SetWindowSize(config.Cfg.ScreenWidth, config.Cfg.ScreenHeight)
	ebiten.SetWindowTitle("")
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}
}
