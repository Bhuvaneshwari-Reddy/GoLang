package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"sync"
	"time"
)

//libraries
//[fmt,error,flag,os,string,strc lonv]

func main() {
	// bufio1()
	// str()
	// syncs()
	// maths()
	timeFn()

}

func bufio1() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text:")

	text, _ := reader.ReadString('\n')
	fmt.Println("You entered:", text)

}

func str() {
	str := "Hello world"
	length := len(str)
	fmt.Println("Length of the string:", length)
	strUpper := strings.ToUpper(str)
	fmt.Println("Uppercase:", strUpper)
	strLower := strings.ToLower(str)
	fmt.Println("LowerCase", strLower)
	fmt.Println(strings.Contains("hello world", "world"))
}

// sync.WaitGroup is used to make the main goroutine WAIT until other goroutines finish their work.
func syncs() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hello from goroutine ")
	}()

	wg.Wait()
	fmt.Println("Main function finished")

}

func maths() {
	fmt.Println(math.Sqrt(16))
	fmt.Println(math.Floor(4.3))
	fmt.Println(math.Ceil(4.3))
}

func timeFn() {
	fmt.Println(time.Now())
	fmt.Println(time.Sleep)

}
