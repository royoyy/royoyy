package main

import (
	"./stack"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("INPUT:\nnumber1 ENTER number2 ENTER operator ENTER\nenter \"q\" to quit")
	var st stack.Stack
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		input = input[:len(input)-1]
		switch input {
		case "q":
			return
		case "+":
			q, _ := st.Pop()
			p, _ := st.Pop()
			fmt.Printf("%d %s %d = %d\n", p, input, q, p.(int)+q.(int))
		case "-":
			q, _ := st.Pop()
			p, _ := st.Pop()
			fmt.Printf("%d %s %d = %d\n", p, input, q, p.(int)-q.(int))
		default:
			i, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println(err)
				return
			}
			st.Push(i)
		}
	}
}
