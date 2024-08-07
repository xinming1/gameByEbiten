package config

import (
	"encoding/json"
	"image/color"
	"log"
	"os"
)

type Config struct {
	ScreenWidth  int           `json:"screenWidth"`
	ScreenHeight int           `json:"screenHeight"`
	Title        string        `json:"title"`
	BgColor      color.RGBA    `json:"bgColor"`
	BgImg        string        `json:"bgImg"`
	Dog          *DogConfig    `json:"dogConfig"`
	SwordConfig  *SwordConfig  `json:"swordConfig"`
	GoblinConfig *GoblinConfig `json:"goblinConfig"`
}

type DogConfig struct {
	Img          string  `json:"img"`
	DefaultSpeed float64 `json:"defaultSpeed"`
}

type SwordConfig struct {
	Img string `json:"img"`
}

type GoblinConfig struct {
	RunImg  ImgConfig `json:"runImg"`
	DiedImg ImgConfig `json:"diedImg"`
	Speed   float64   `json:"speed"`
}

type ImgConfig struct {
	Img    string `json:"imgTemplate"`
	ImgNum int    `json:"imgNum"`
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
