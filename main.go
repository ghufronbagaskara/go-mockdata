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

	if help || inputPath == "" || outputPath == ""{
		printUsage()
		os.Exit(0)
	}
}

func printUsage() {
	fmt.Println("Usage : mockdata [-i | --input] <input file> [-o | --output] <output file>")
	fmt.Println("-i --input: JSON input file as a template")
	fmt.Println("-o --output: JSON output file")
}