package main

import (
	_ "github.com/coldmine/ic"
	"fmt"
	"flag"
)

func main() {
	var Input string
	var Output string
	var Width int
	var Height int
	var Gamma float64
	flag.StringVar(&Input, "i", "", "`Input` image path")
	flag.StringVar(&Input, "input", "", "`Input` image path")
	flag.StringVar(&Output, "o", "", "`Output` image path")
	flag.StringVar(&Output, "output", "", "`Output` image path")
	flag.IntVar(&Width, "w", 0, "Output `Width` pixel")
	flag.IntVar(&Width, "width", 0, "Output `Width` pixel")
	flag.IntVar(&Height, "h", 0, "Output `Height` pixel")
	flag.IntVar(&Height, "height", 0, "Output `Height` pixel")
	flag.Float64Var(&Gamma, "g", 1.0, "`Gamma` value")
	flag.Float64Var(&Gamma, "gamma", 1.0, "`Gamma` value")
	flag.Parse()
	fmt.Println(Input)
	fmt.Println(Output)
	fmt.Println(Width)
	fmt.Println(Height)
	fmt.Println(Gamma)
	flag.PrintDefaults()
	// 2. 이미지를 컨버팅하도록 명령어를 처리한다.
}
