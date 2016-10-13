package ic

import (
	"image"
	_"image/jpeg"
	_ "image/png"
	"os"
	"fmt"
)

func Resize(ifile, ofile string, width, height int) error {
	reader, err := os.Open(ifile)
	if err != nil {
		return err
	}
	defer reader.Close()
	m, _, err := image.Decode(reader)
	if err != nil {
		return err
	}
	bounds := m.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := m.At(x,y).RGBA()
			fmt.Println(r)
			fmt.Println(g)
			fmt.Println(b)
			fmt.Println(a)
		}
	}
	fmt.Println(ofile)
	fmt.Println(width)
	fmt.Println(height)
	return nil
}
