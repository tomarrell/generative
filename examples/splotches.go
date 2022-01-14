package examples

import (
	"image"
	"math/rand"

	"github.com/fogleman/gg"
	"github.com/tomarrell/graphics/colors"
	"github.com/tomarrell/graphics/frame"
)

func Splotches(w, h float64) image.Image {
	c := gg.NewContext(int(w), int(h))

	c.DrawImage(frame.Artboard(w, h), 0, 0)

	for i := 0; i < 10; i++ {
		x, y, r := rand.Intn(int(w)), rand.Intn(int(h)), rand.Intn(30)
		c.SetColor(colors.Red)
		c.DrawCircle(float64(x), float64(y), float64(r))
		c.Fill()
	}

	return c.Image()
}
