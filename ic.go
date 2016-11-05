package ic

import (
	"image"

	"golang.org/x/image/draw"
)

func Resize(src image.Image, width, height int) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	// TODO: Support multiple kernels(Kernel is an interpolator that blends source pixels weighted by a symmetric).
	draw.BiLinear.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Src, nil)
	return dst
}
