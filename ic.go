package ic

import (
	"image"
	"image/jpeg"
	_ "image/png"
	"os"
	"golang.org/x/image/draw"
)

func Resize(srcpath, dstpath string, width, height int) error {
	// TODO: 만약 width값이 0 이라면 height값을 기준으로 autoWidth값을 가지고 온다.
	// TODO: 만약 height값이 0 이라면 width값을 기준으로 autoHeight값을 가지고 온다.
	reader, err := os.Open(srcpath)
	if err != nil {
		return err
	}
	defer reader.Close()
	src, _, err := image.Decode(reader)
	if err != nil {
		return err
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
