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
	Dog          *DogConfig `json:"dog"`
}

type DogConfig struct {
	Img          string  `json:"img"`
	DefaultSpeed float64 `json:"defaultSpeed"`
}

var Cfg *Config

func LoadConfig(configPath string) {
	f, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("os.Open failed: %v\n", err)
	}

	err = json.NewDecoder(f).Decode(&Cfg)
	if err != nil {
		log.Fatalf("json.Decode failed: %v\n", err)
	}
}
