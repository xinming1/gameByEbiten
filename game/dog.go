package game

import (
	"game_by_ebiten/config"
	"game_by_ebiten/definition"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Dog struct {
	direction, lastDirection int
	image                    *ebiten.Image
	width, height            float64
	x, y                     float64
	speed                    float64
	imgOptions               *ebiten.DrawImageOptions
}

func NewDog(dogConfig *config.DogConfig) *Dog {
	img, _, err := ebitenutil.NewImageFromFile(dogConfig.Img)
	if err != nil {
		log.Fatal(err)
	}
	dog := &Dog{
		direction:     definition.Left,
		lastDirection: definition.Left,
		image:         img,
		width:         float64(img.Bounds().Dx()),
		height:        float64(img.Bounds().Dy()),
		speed:         dogConfig.DefaultSpeed,
	}
	dog.x = (float64(config.Cfg.ScreenWidth) - dog.width) / 2
	dog.y = (float64(config.Cfg.ScreenHeight) - dog.height) / 2
	dog.imgOptions = &ebiten.DrawImageOptions{}
	dog.imgOptions.GeoM.Translate(dog.x, dog.y)
	return dog
}

func (d *Dog) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	if d.direction != definition.Left {
		op.GeoM.Scale(-1, 1)
		d.lastDirection = d.direction
	}
	op.GeoM.Translate(d.x, d.y)
	screen.DrawImage(d.image, op)

	//d.imgOptions.GeoM.Translate(d.x, d.y)
	//screen.DrawImage(d.image, d.imgOptions)
}

func (d *Dog) Update() {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyLeft):
		d.x -= d.speed
		d.direction = definition.Left
	case ebiten.IsKeyPressed(ebiten.KeyRight):
		d.x += d.speed
		d.direction = definition.Right
	case ebiten.IsKeyPressed(ebiten.KeyUp):
		d.y -= d.speed
	case ebiten.IsKeyPressed(ebiten.KeyDown):
		d.y += d.speed

	}
	//d.imgOptions.GeoM.Translate(d.x, d.y)

}
