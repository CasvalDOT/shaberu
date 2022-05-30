package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"shaberu/dialects"
	"shaberu/formatter"
)

func getExtensionFromPath(filePath *string) string {
	extension := filepath.Ext(*filePath)

	// Return the extension ignoring the "." char at
	// starting string
	return extension[1:]
}

func main() {
	// Parse the flags provided
	// using the CLI
	separator := flag.String("s", ";", "The character separator")
	inputPath := flag.String("i", "", "The path of input file")
	outputPath := flag.String("o", "output.json", "The path of the output file")
	formatterName := flag.String("f", "", "The formatter to use. For a list of available formatters read the docs.")

	flag.Parse()

	// Check if a source path is set
	// if this parameter is not set
	// the script cannot do so much
	if *inputPath == "" {
		fmt.Println("Missing input file")
		os.Exit(1)
	}

	// Take the source type using the extension
	// if this argument is not explicited set
	readerType := getExtensionFromPath(inputPath)

	// Take the output type using the extension
	// if this argument is not expplicited set
	writerType := getExtensionFromPath(outputPath)

	// Create the reader
	reader := dialects.New(readerType)
	if reader == nil {
		fmt.Println("Input type not supported")
		os.Exit(1)
	}

	// Create the writer
	writer := dialects.New(writerType)
	if writer == nil {
		fmt.Println("Output type not supported")
		os.Exit(1)
	}

	// Read the input file
	sourceMap, err := reader.Read(inputPath, separator)
	if err != nil {
		fmt.Println("Reading error:", err.Error())
		os.Exit(1)
	}

	// Write in the output file
	err = writer.Write(outputPath, sourceMap)
	if err != nil {
		fmt.Println("Writing error:", err.Error())
		os.Exit(1)
	}

	// Apply a formatter if provided
	if *formatterName != "" {
		formatterInstance := formatter.New(*formatterName)
		formatterInstance.Format(outputPath)
	}

	// Finish
	fmt.Println("Convert successfully")
}
