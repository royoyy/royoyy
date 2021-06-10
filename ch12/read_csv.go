package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	inputFile, inputError := os.Open("./ch12/data.csv")
	if inputError != nil {
		fmt.Println(inputError)
		return
	}
	defer inputFile.Close()

	inputReader := csv.NewReader(inputFile)
	recond, err := inputReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(recond)
}
