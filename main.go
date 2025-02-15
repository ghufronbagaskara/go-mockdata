package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ghufronbagaskara/go-mockdata/data"
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

	var mapping map[string]string
	err := readInput(inputPath, &mapping); if err != nil {
		fmt.Printf("Read input error: %s \n", err)
		os.Exit(0)
	}


	if err := valydateType(mapping); err != nil{
		fmt.Printf("Error validating data type: %s \n", err)
		os.Exit(0)
	}



	result, err := generateOutput(mapping)
	if err != nil {
		fmt.Printf("Error generating output data: %s \n", err)
		os.Exit(0)
	}

	if err := writeOutput(outputPath, result); err != nil {
		fmt.Printf("Error writing result: %s \n", err)
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

func readInput(path string, mapping *map[string]string) error {
	if path == "" {
		return errors.New("Invalid path")
	}

	if mapping == nil {
		return errors.New("Invalid mapping")
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	fileByte, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if len(fileByte) == 0 {
		return errors.New("Empty input")
	}
	
	if err := json.Unmarshal(fileByte, &mapping); err != nil {
		return err
	}

	
	return nil
}	

func valydateType(mapping map[string]string) error {	
	for _ , value := range mapping {
		if !data.Supported[value] {
			return errors.New("Unsupported data type")
		}
	}

	return nil
}

func generateOutput(mapping map[string]string) (map[string]any, error) {
	result := make(map[string]any)

	for key, dataType := range mapping {
		result[key] = data.Generate(dataType)
	}

	return result, nil
}

func writeOutput(path string, result map[string]any) error {
	if path == "" {
		return errors.New("Invalid path")
	}


	flags := os.O_RDWR | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile(path, flags, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	
	resultByte, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return err
	}

	if _, err := file.Write(resultByte); err != nil {
		return err
	}

	
	return nil
}