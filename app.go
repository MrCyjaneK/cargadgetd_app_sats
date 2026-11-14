package main

import (
	"image"
	"sync"
	"time"
)

// satNote is one lock or drop from the daemon broadcast.
type satNote struct {
	Gain                       bool
	Kicker, Title, Sub, Detail string
	Flag                       string
}

// Shell paints the satellite toasts. Empty pixels stay clear.
type Shell struct {
	mu       sync.Mutex
	uiScale  float32
	satStack []satCard
}

func NewShell() *Shell { return &Shell{uiScale: 1} }

func (s *Shell) GPSSat(gain bool, kicker, title, sub, detail, flag string) {
	s.queueSatNotes([]satNote{{
		Gain: gain, Kicker: kicker, Title: title, Sub: sub, Detail: detail, Flag: flag,
	}})
}

func (s *Shell) Draw(dst *image.RGBA, now time.Time) {
	if dst == nil {
		return
	}
	sc := float32(dst.Bounds().Dy()) / 480
	if sc < 0.25 {
		sc = 0.25
	}
	s.mu.Lock()
	s.uiScale = sc
	s.mu.Unlock()
	for i := range dst.Pix {
		dst.Pix[i] = 0
	}
	s.paintSat(dst, now)
}

var satsApp = NewShell()
