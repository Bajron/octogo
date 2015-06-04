package main

import (
	"flag"
	"fmt"
	"github.com/Bajron/octogo/octogo"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var inFile, outFile, mode string

var (
	kDefaultInput  = "in.(png|jpg|gif)"
	kDefaultOutput = "out.(png|jpg|bmp|gif)"
	kDefaultMode   = "(copy|gray|...)"
)

func init() {
	const (
		INPUT_USAGE = "a file to read"
		OUPUT_USAGE = "output filename"
		MODE_USAGE  = "processing function"
	)

	kDefaultOutput = "out.(" + strings.Join(octogo.GetEncoders(), "|") + ")"
	kDefaultMode = "(" + strings.Join(octogo.GetModes(), "|") + ")"

	flag.StringVar(&inFile, "input", kDefaultInput, INPUT_USAGE)
	flag.StringVar(&inFile, "f", kDefaultInput, INPUT_USAGE+" (shorthand for --input)")

	flag.StringVar(&outFile, "output", kDefaultOutput, OUPUT_USAGE)
	flag.StringVar(&outFile, "o", kDefaultOutput, OUPUT_USAGE+" (shorthand for --output)")

	flag.StringVar(&mode, "mode", kDefaultMode, MODE_USAGE)
	flag.StringVar(&mode, "m", kDefaultMode, MODE_USAGE+" (shorthand for --mode)")
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			"Usage:\n"+
				"\t%s [options] [input file] [output file]\n\n",
			os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() > 0 && inFile == kDefaultInput {
		inFile = flag.Arg(0)
	}
	if flag.NArg() > 1 && outFile == kDefaultOutput {
		outFile = flag.Arg(1)
	}

	if inFile == kDefaultInput {
		fmt.Printf("You need to provide input file!\n")
		flag.Usage()
		return
	}

	if outFile == kDefaultOutput {
		outFile = "out.png"
	}

	if mode == kDefaultMode {
		mode = "copy"
	}

	ext := filepath.Ext(outFile)
	if ext == "" {
		outFile += ".png"
	}

	log.Printf("Processing %s -[%s]-> %s!\n", inFile, mode, outFile)
	octogo.Process(inFile, outFile, octogo.GetProcessingFunction(mode))
}
