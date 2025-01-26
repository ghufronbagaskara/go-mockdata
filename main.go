package main

import (
	"flag"
	"fmt"
	"os"
)


func main() {
	var help bool
	var inputPath, outputPath string

	flag.BoolVar(&help, "h", false, "Show how to use the program")
	flag.BoolVar(&help, "help", false, "Show how to use the program")

	flag.StringVar(&inputPath, "i", "", "JSON file to read from as an input")
	flag.StringVar(&inputPath, "input", "", "JSON file to read from as an input")
	flag.StringVar(&outputPath, "o", "", "JSON file to read from as an output")
	flag.StringVar(&outputPath, "output", "", "JSON file to read from as an output")

	flag.Parse()

	if help {
		fmt.Println("mockdata -i <input file> -o <output file>")
		os.Exit(0)
	}

	if inputPath == "" {
		fmt.Println("Please provide an input file")
		os.Exit(1)
	}

	if outputPath == "" {
		fmt.Println("Please provide an output file")
		os.Exit(1)
	}
}