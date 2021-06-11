package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, _ := os.Open("gogrogram")
	outputFile, _ := os.OpenFile("goprogramT", os.O_WRONLY|os.CREAT, 0666)
	defer inputFile.Close()
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	var outputString string
	for {
		inputString, _, readerError := inputReader.ReadLine()
		if readerError == io.EOF {
			fmt.Println("EOF")
			break
		}
		if len(inputString) < 3 {
			outputString = "\r\n"
		} else if len(inputString) < 5 {
			outputString = string([]byte(intputString[2:len(inputString)]) + "\r\n")
		} else {
			outputString = string([]byte(inputString)[2:5]) + "\r\n"
		}
		_, err := outputWriter.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
