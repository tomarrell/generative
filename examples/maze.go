package examples

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/fogleman/gg"
	"github.com/tomarrell/graphics/frame"
)

var points = []struct{ x1, y1, x2, y2 float64 }{
	{0, 0, 1, 1},
	{1, 0, 0, 1},
}

func Maze(width, height float64) image.Image {
	c := gg.NewContext(int(width), int(height))
	c.DrawImage(frame.Artboard(width, height), 0, 0)

	const count = 20

	w := width / count

	c.SetColor(color.Black)
	for i := 0; i < count; i++ {
		c.Push()
		for j := 0; j < count; j++ {
			c.Push()
			if rand.Intn(2) > 0 {
				c.LineTo(w*points[0].x1, w*points[0].y1)
				c.LineTo(w*points[0].x2, w*points[0].y2)
			} else {
				c.LineTo(w*points[1].x1, w*points[1].y1)
				c.LineTo(w*points[1].x2, w*points[1].y2)
			}
			c.Stroke()
			c.Pop()
			c.Translate(w, 0)
		}
		c.Pop()
		c.Translate(0, w)
	}

	return c.Image()
}
