package ic

import (
	"image"
	"image/jpeg"
	_ "image/png"
	"os"
	"golang.org/x/image/draw"
)

func Resize(ifile, ofile string, width, height int) error {
	reader, err := os.Open(ifile)
	if err != nil {
		return err
	}
	defer reader.Close()
	src, _, err := image.Decode(reader)
	if err != nil {
		return err
	}

	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	// TODO: Support multiple kernels.
	draw.BiLinear.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Src, nil)

	// TODO: Support multiple file extensions.
	out, err := os.Create(ofile)
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
