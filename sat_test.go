package main

import (
	"image"
	"image/color"
	"testing"
	"time"
)

func TestSatToastPaintsAtTheOrigin(t *testing.T) {
	sh := NewShell()
	sh.GPSSat(true, "SAT LOCK", "E26", "Galileo · European Union", "EL 37    AZ 67    SNR 34", "eu")
	dst := image.NewRGBA(image.Rect(0, 0, 800, 480))
	now := time.Date(2026, 10, 27, 21, 48, 0, 0, time.UTC)
	sh.Draw(dst, now)
	box := satToastBox(dst.Bounds(), sh.uiScale, satNote{
		Kicker: "SAT LOCK", Title: "E26", Sub: "Galileo · European Union", Detail: "EL 37    AZ 67    SNR 34",
	}, 0)
	n := 0
	eu := 0
	for y := box.Min.Y; y < box.Max.Y; y++ {
		for x := box.Min.X; x < box.Max.X; x++ {
			p := dst.RGBAAt(x, y)
			if p.R == colSatBG.R && p.G == colSatBG.G && p.B == colSatBG.B && p.A == colSatBG.A {
				n++
			}
			if p == (color.RGBA{R: 0, G: 51, B: 153, A: 255}) {
				eu++
			}
		}
	}
	if n < 200 {
		t.Fatalf("card pixels %d", n)
	}
	if eu < 20 {
		t.Fatal("flag missing")
	}
	if dst.RGBAAt(dst.Bounds().Dx()-4, dst.Bounds().Dy()-4).A != 0 {
		t.Fatal("toast painted the opposite corner")
	}
	sh.Draw(dst, now.Add(satHold))
	if dst.RGBAAt(2, 2).A != 0 {
		t.Fatal("toast stuck")
	}
}
