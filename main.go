package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
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

	if err := validateInput(inputPath); err != nil {
		fmt.Printf("Invalid input: %s \n", err)
		os.Exit(0)
	}
	
	if err := validateOutput(outputPath); err != nil {
		fmt.Printf("Invalid output: %s \n", err)
		os.Exit(0)
	}

}



func printUsage() {
	fmt.Println("Usage : mockdata [-i | --input] <input file> [-o | --output] <output file>")
	fmt.Println("-i --input: JSON input file as a template")
	fmt.Println("-o --output: JSON output file")
}

func validateInput(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err){
		return err
	}
	
	return nil 
}

func validateOutput(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err){
		return nil
	}
	
	confirmOverwrite()
	return nil
}

func confirmOverwrite() {
	fmt.Println("Output file already exists. Do you want to overwrite it? (y/n)") 
	
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = strings.ToLower(strings.TrimSpace(response))

	if response != "y" && response != "yes" && response != "ya" {
		fmt.Println("Exiting program")
		os.Exit(0)
	}
}