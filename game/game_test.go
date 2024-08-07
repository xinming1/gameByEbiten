package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"testing"
)

func TestGeoMApply(t *testing.T) {
	trans := ebiten.GeoM{}
	trans.Translate(1, 2)

	scale := ebiten.GeoM{}
	scale.Scale(1.5, 2.5)

	cpx := ebiten.GeoM{}
	cpx.Rotate(math.Pi)
	cpx.Scale(1.5, 2.5)
	cpx.Translate(-2, -3)

	cases := []struct {
		GeoM ebiten.GeoM
		InX  float64
		InY  float64
		OutX float64
		OutY float64
	}{
		{
			GeoM: ebiten.GeoM{},
			InX:  3.14159,
			InY:  2.81828,
			OutX: 3.14159,
			OutY: 2.81828,
		},
		{
			GeoM: trans,
			InX:  3.14159,
			InY:  2.81828,
			OutX: 4.14159,
			OutY: 4.81828,
		},
		{
			GeoM: scale,
			InX:  3.14159,
			InY:  2.81828,
			OutX: 4.71239,
			OutY: 7.04570,
		},
		{
			GeoM: cpx,
			InX:  3.14159,
			InY:  2.81828,
			OutX: -6.71239,
			OutY: -10.04570,
		},
	}

	const delta = 0.00001

	for _, c := range cases {
		rx, ry := c.GeoM.Apply(c.InX, c.InY)
		t.Logf("%s.Apply(%f, %f) = (%f, %f), want (%f, %f)", c.GeoM.String(), c.InX, c.InY, rx, ry, c.OutX, c.OutY)
		if math.Abs(rx-c.OutX) > delta || math.Abs(ry-c.OutY) > delta {
			t.Errorf("%s.Apply(%f, %f) = (%f, %f), want (%f, %f)", c.GeoM.String(), c.InX, c.InY, rx, ry, c.OutX, c.OutY)
		}
	}
}
