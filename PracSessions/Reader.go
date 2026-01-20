package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza:")

	//comma ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for reading, ", input)

}
