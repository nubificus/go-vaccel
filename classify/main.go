package main

import (
	"fmt"
	"os"

	"github.com/nubificus/go-vaccel/vaccel"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./classify <filename>")
		return
	}

	/* Get the filename from command line argument */
	filePath := os.Args[1]

	/* Session */
	var session vaccel.Session

	err := vaccel.SessionInit(&session, 0)

	if err != 0 {
		fmt.Println("error initializing session")
		os.Exit(err)
	}

	var outText string

	/* Run the Operation providing the path */
	outText, err = vaccel.ImageClassificationFromFile(&session, filePath)

	if err != 0 {
		fmt.Println("Image Classification failed")
		os.Exit(err)
	}

	fmt.Println("Output(1): ", outText)

	/* Or by providing the bytes */
	imageBytes, e := os.ReadFile(filePath)
	if e != nil {
		fmt.Printf("Error reading file: %s\n", e)
		os.Exit(1)
	}

	outText, err = vaccel.ImageClassification(&session, imageBytes)

	if err != 0 {
		fmt.Println("Image Classification failed")
		os.Exit(err)
	}

	fmt.Println("Output(2): ", outText)

	/* Free Session */
	if vaccel.SessionFree(&session) != 0 {
		fmt.Println("An error occurred while freeing the session")
	}
}
