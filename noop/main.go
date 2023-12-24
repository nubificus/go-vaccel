package main

/*

#cgo pkg-config: vaccel
#include <vaccel.h>
#include <stdlib.h>
#include <stdio.h>

*/
import (
	"github.com/nubificus/go-vaccel/vaccel"
	"os"
	"fmt"
	"C"
)

func main() {

	/* Session */
	var session vaccel.Session
	err := vaccel.SessionInit(&session, 0)

	if err != 0 {
		fmt.Println("error initializing session")
		os.Exit(int(err))
	}

	/* Run the operation */
	err = vaccel.NoOp(&session)

	if err != 0 {
		fmt.Println("An error occured while running the operation")
		os.Exit(err)
	}

	/* Free Session */
	if vaccel.SessionFree(&session) != 0 {
		fmt.Println("An error occured while freeing the session")
	}

}
