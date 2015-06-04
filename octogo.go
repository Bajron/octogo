package main

import (
	"flag"
	"fmt"
	"github.com/Bajron/octogo/octogo"
	"log"
	"path/filepath"
)

var inFile, outFile string

const (
	DEFAULT_INPUT  = "in.(png|jpg|bmp|gif)"
	DEFAULT_OUTPUT = "out.(png|jpg|bmp|gif)"
)

func init() {
	const (
		INPUT_USAGE = "specifies a file to read (jpg,png,gif or bmp)"
		OUPUT_USAGE = "specifies output filename"
	)

	flag.StringVar(&inFile, "input", DEFAULT_INPUT, INPUT_USAGE)
	flag.StringVar(&inFile, "f", DEFAULT_INPUT, INPUT_USAGE+" (shorthand for --input)")

	flag.StringVar(&outFile, "output", DEFAULT_OUTPUT, OUPUT_USAGE)
	flag.StringVar(&outFile, "o", DEFAULT_OUTPUT, OUPUT_USAGE+" (shorthand for --output)")
}

func main() {
	flag.Parse()

	if flag.NArg() > 0 && inFile == DEFAULT_INPUT {
		inFile = flag.Arg(0)
	}
	if flag.NArg() > 1 && outFile == DEFAULT_OUTPUT {
		outFile = flag.Arg(1)
	}

	if inFile == DEFAULT_INPUT {
		fmt.Printf("You need to provide input file!\n")
		flag.Usage()
		return
	}

	if outFile == DEFAULT_OUTPUT {
		outFile = "out.png"
	}

	ext := filepath.Ext(outFile)
	if ext == "" {
		outFile += ".png"
	}

	log.Printf("Processing %s -> %s!\n", inFile, outFile)
	octogo.Process(inFile, outFile, octogo.Copy)
}
