package main

import (
	"image"
	"image/png"
	"io"
)

const (
	imageWidth  = 1280
	imageHeight = 720
)

type Render struct {
	w io.Writer
}

func (r *Render) Run() error {
	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	return png.Encode(r.w, img)
}
