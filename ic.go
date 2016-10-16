package ic

import (
	"image"
	"image/jpeg"
	_ "image/png"
	"os"

	"golang.org/x/image/draw"
)

func Resize(srcpath, dstpath string, width, height int) error {
	reader, err := os.Open(srcpath)
	if err != nil {
		return err
	}
	defer reader.Close()
	src, _, err := image.Decode(reader)
	if err != nil {
		return err
	}

	// 빈 width, height 값 자동으로 계산
	if width == 0 || height == 0 {
		s := src.Bounds().Max
		if width == 0 {
			width = int(float32(height) / float32(s.Y) * float32(s.X))
		} else {
			height = int(float32(width) / float32(s.X) * float32(s.Y))
		}
	}

	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	// TODO: Support multiple kernels(Kernel is an interpolator that blends source pixels weighted by a symmetric).
	draw.BiLinear.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Src, nil)

	// TODO: Support multiple file extensions.
	out, err := os.Create(dstpath)
	if err != nil {
		return err
	}
	defer out.Close()
	err = jpeg.Encode(out, dst, nil)
	if err != nil {
		return err
	}
	return nil
}
