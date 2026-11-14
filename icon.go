package main

import (
	"bytes"
	"embed"
	"image"
	"image/draw"
	"image/png"
)

//go:embed icon.png
var iconFS embed.FS

var iconPix []byte
var iconW, iconH int

func loadIcon() {
	if iconPix != nil {
		return
	}
	raw, err := iconFS.ReadFile("icon.png")
	if err != nil {
		return
	}
	img, err := png.Decode(bytes.NewReader(raw))
	if err != nil {
		return
	}
	b := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, b.Min, draw.Src)
	iconPix = append([]byte(nil), rgba.Pix...)
	iconW, iconH = rgba.Bounds().Dx(), rgba.Bounds().Dy()
}
