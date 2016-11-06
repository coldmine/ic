package main

import (
	"flag"
	"image"
	"image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/coldmine/ic"
)

func main() {
	var Input string
	var Output string
	var Width int
	var Height int
	//var Gamma float64
	flag.StringVar(&Input, "i", "", "`Input` image path")
	flag.StringVar(&Input, "input", "", "`Input` image path")
	flag.StringVar(&Output, "o", "", "`Output` image path")
	flag.StringVar(&Output, "output", "", "`Output` image path")
	flag.IntVar(&Width, "w", 0, "Output `Width` pixel")
	flag.IntVar(&Width, "width", 0, "Output `Width` pixel")
	flag.IntVar(&Height, "h", 0, "Output `Height` pixel")
	flag.IntVar(&Height, "height", 0, "Output `Height` pixel")
	//flag.Float64Var(&Gamma, "g", 1.0, "`Gamma` value")
	//flag.Float64Var(&Gamma, "gamma", 1.0, "`Gamma` value")
	flag.Parse()

	if (Width == 0 && Height == 0) || Input == "" || Output == "" {
		flag.PrintDefaults()
		os.Exit(1)
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

	//예외처리 넣을 것.

	// 파일 읽기
	in, err := os.Open(Input)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()
	src, _, err := image.Decode(in)
	if err != nil {
		log.Fatal(err)
	}

	dst := ic.Resize(src, width, height)

	// 이미지 쓰기
	// TODO: 여러 파일포맷 지원
	out, err := os.Create(Output)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	err = jpeg.Encode(out, dst, nil)
	if err != nil {
		log.Fatal(err)
	}
}
