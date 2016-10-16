package main

import (
	"github.com/coldmine/ic"
	"flag"
	"log"
	"os"
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
		os.Exit(0)
	}

	//예외처리 넣을 것.

	//프로세싱.
	err := ic.Resize(Input, Output, Width, Height)
	if err != nil {
		log.Fatal(err)
	}
}
