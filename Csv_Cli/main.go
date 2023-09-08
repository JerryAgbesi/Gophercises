package main

import (
	"flag"
	"os"
	"fmt"
)

func main(){
	csvFilename := flag.String("csv","questions.csv","Use a csv file with content formatted as 'question,answer'")
	flag.Parse()
	
	file,err := os.Open(*csvFilename)

	if err != nil {
		fmt.Printf("Error reading from file: %s",*csvFilename)
	}
	os.Exit(1)

	_ = file


}