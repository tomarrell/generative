package examples

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/fogleman/gg"
	"github.com/tomarrell/graphics/frame"
)

const inset = 4

type point struct {
	x, y float64
}

func Shapes(width, height float64) image.Image {
	c := gg.NewContext(int(width), int(height))

	const count = 10
	w := width / count

	c.DrawImage(frame.Artboard(width, height), 0, 0)

	for i := 0; i < count; i++ {
		c.Push()
		for j := 0; j < count; j++ {
			c.Push()
			c.Rotate(gg.Radians(float64(rand.Intn(360))))
			c.DrawImage(triangle(w), 0, 0)
			c.Pop()

			c.Push()
			if rand.Intn(2) > 0 {
				c.Scale(1.2, 1.2)
				c.DrawImage(square(w), 0, 0)
			}
			c.Pop()

			c.Translate(w, 0)
		}
		c.Pop()
		c.Translate(0, w)
	}

	return c.Image()
}

func triangle(w float64) image.Image {
	c := gg.NewContext(int(w), int(w))

	points := []point{
		{w / 2, 0 + inset},
		{0 + inset, w - inset},
		{w - inset, w - inset},
		{w / 2, 0 + inset},
	}

	c.SetColor(color.Black)
	c.SetLineWidth(2)

	for _, p := range points {
		c.LineTo(p.x, p.y)
	}

	c.Stroke()

	return c.Image()
}

func square(w float64) image.Image {
	c := gg.NewContext(int(w), int(w))

	points := []point{
		{0 + inset, 0 + inset},
		{w - inset, 0 + inset},
		{w - inset, w - inset},
		{0 + inset, w - inset},
		{0 + inset, 0 + inset},
	}

	c.SetColor(color.Black)
	c.SetLineWidth(2)

	for _, p := range points {
		c.LineTo(p.x, p.y)
	}

	c.Stroke()

	return c.Image()
}
