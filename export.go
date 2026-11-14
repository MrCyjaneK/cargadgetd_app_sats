package main

/*
#cgo CFLAGS: -I${SRCDIR}/include
#include "cargadgetd/cg_app.h"
#include <stdint.h>
#include <stdlib.h>
#include <string.h>
*/
import "C"

import (
	"image"
	"time"
	"unsafe"
)

var (
	satsName = C.CString("Satellites")
	iconC    *C.uint8_t
)

func ensureIconC() {
	if iconC != nil {
		return
	}
	loadIcon()
	if len(iconPix) == 0 {
		return
	}
	n := C.size_t(len(iconPix))
	iconC = (*C.uint8_t)(C.malloc(n))
	if iconC == nil {
		return
	}
	C.memcpy(unsafe.Pointer(iconC), unsafe.Pointer(&iconPix[0]), n)
}

//export cargadgetd_app_sats_icon
func cargadgetd_app_sats_icon(w, h *C.int) *C.uint8_t {
	ensureIconC()
	if w != nil {
		*w = C.int(iconW)
	}
	if h != nil {
		*h = C.int(iconH)
	}
	return iconC
}

//export cargadgetd_app_sats_init
func cargadgetd_app_sats_init(*C.struct_cargadgetd_host) {}

//export cargadgetd_app_sats_draw
func cargadgetd_app_sats_draw(rgba *C.uint32_t, w, h, stride C.int) {
	if rgba == nil || w < 1 || h < 1 || stride < w*4 {
		return
	}
	pix := unsafe.Slice((*uint8)(unsafe.Pointer(rgba)), int(h)*int(stride))
	dst := &image.RGBA{Pix: pix, Stride: int(stride), Rect: image.Rect(0, 0, int(w), int(h))}
	satsApp.Draw(dst, time.Now())
}

//export cargadgetd_app_sats_touch
func cargadgetd_app_sats_touch(C.int, C.int, C.int) {}

//export cargadgetd_app_sats_pause
func cargadgetd_app_sats_pause() {}

//export cargadgetd_app_sats_resume
func cargadgetd_app_sats_resume() {}

//export cargadgetd_app_sats_stop
func cargadgetd_app_sats_stop() {}

//export cargadgetd_app_sats_name
func cargadgetd_app_sats_name() *C.char { return satsName }

//export cargadgetd_app_sats_gps_sat
func cargadgetd_app_sats_gps_sat(gain C.int, kicker, title, sub, detail, flag *C.char) {
	satsApp.GPSSat(gain != 0, C.GoString(kicker), C.GoString(title), C.GoString(sub), C.GoString(detail), C.GoString(flag))
}

func main() {}
