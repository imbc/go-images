package main

import (
	"image/color"
)

func (c Canvas) DrawLine(color color.RGBA, from Vector, to Vector) {
	delta := to.Sub(from)
	length := delta.Length()
	x_step, y_step := delta.X/length, delta.Y/length
	limit := int(length + 0.5)
	for i := 0; i < limit; i++ {
		x := from.X + float64(i)*x_step
		y := from.Y + float64(i)*y_step
		c.Set(int(x), int(y), color)
	}
}
