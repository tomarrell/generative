package frame

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"
)

var (
	paddingColor = color.RGBA{245, 245, 245, 255}
)

func Frame(width, height, borderWidth, shadowWidth float64, frameColor color.Color) image.Image {
	c := gg.NewContext(int(width+shadowWidth), int(height+shadowWidth))
	c.InvertY()
	c.RotateAbout(gg.Radians(180), float64(width+shadowWidth)/2, float64(height+shadowWidth)/2)

	c.Push()
	runs := 20
	for i := runs; i > 0; i-- {
		c.SetColor(color.RGBA{0, 0, 0, uint8(runs - i)})
		c.DrawRectangle(float64(i*1), float64(i*1), width, height)
		c.Fill()
	}
	c.Pop()

	c.SetColor(frameColor)
	c.DrawRectangle(0, 0, width, height)
	c.Fill()

	c.SetColor(paddingColor)
	c.DrawRectangle(borderWidth, borderWidth, width-2*borderWidth, height-2*borderWidth)
	c.Fill()

	return c.Image()
}

func Artboard(w, h float64) image.Image {
	c := gg.NewContext(int(w), int(h))
	c.SetColor(color.RGBA{0, 0, 0, 10})
	c.DrawRectangle(0, 0, w, h)
	c.Fill()

	c.Translate(0, 2)
	c.SetColor(color.White)
	c.DrawRectangle(0, 0, w-2, h)
	c.Fill()

	return c.Image()
}
