package config

import (
	"encoding/json"
	"image/color"
	"log"
	"os"
)

type Config struct {
	ScreenWidth  int        `json:"screenWidth"`
	ScreenHeight int        `json:"screenHeight"`
	Title        string     `json:"title"`
	BgColor      color.RGBA `json:"bgColor"`
}

var Cfg *Config

func loadConfig() {
	f, err := os.Open("./config.json")
	if err != nil {
		log.Fatalf("os.Open failed: %v\n", err)
	}

	err = json.NewDecoder(f).Decode(&Cfg)
	if err != nil {
		log.Fatalf("json.Decode failed: %v\n", err)
	}
}
