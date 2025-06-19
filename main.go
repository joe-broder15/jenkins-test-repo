package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter text now  : ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		fmt.Println("You entered:", input)
	}
}
