package avatar

import (
	"image"
	"image/color"
)

type Circle struct {
	p image.Point
	r int
}

func (c *Circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *Circle) Bounds() image.Rectangle {
	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

func (c *Circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.p.X), float64(y-c.p.Y), float64(c.r)
	if xx*xx+yy*yy < rr*rr {
		return color.RGBA{0, 0, 0, 255}
	}
	return color.Alpha{0}
}
