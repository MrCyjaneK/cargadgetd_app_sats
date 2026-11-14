package main

import (
	"image"
	"image/color"
)

func fillCircle(dst *image.RGBA, cx, cy, r int, c color.Color) {
	if r < 1 || dst == nil {
		return
	}
	rr := r * r
	for y := cy - r; y <= cy+r; y++ {
		for x := cx - r; x <= cx+r; x++ {
			dx, dy := x-cx, y-cy
			if dx*dx+dy*dy <= rr && image.Pt(x, y).In(dst.Bounds()) {
				dst.Set(x, y, c)
			}
		}
	}
}
