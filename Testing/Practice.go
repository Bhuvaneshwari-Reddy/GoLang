package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("Result: ")
	printHello()
	add(5, 6)
	practice()
	stringPractice()
	booleanPractice()
	smallValuePractice()
	floatPractice()
	string1()
	Testing()
	Target()
	Variables()
	string2()
	sum()
	dataTypes()
	constants()
	isStatement()
	switchStatement()
	forLoops()

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type:%T \n", anotherVariable)

	//no var style
	numberOfUsers := 3000
	fmt.Printf("variable is of type:%T \n", numberOfUsers)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter the rating for our pizza:")

	//comma ok syntax
	// input, _ := reader.ReadString('\n')
	// fmt.Println("Thanks for reading, ", input)

	fmt.Println(time.Now())

}

func printHello() {
	fmt.Println("Hello world")
}

func add(a int, b int) {
	fmt.Println(a + b)
}

func practice() {
	fmt.Println("Hello, World!")
	fmt.Print("Hello Bhuvana")

	for i := 5; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Print("Hellooooo")
	for i := 1; i < 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func stringPractice() {
	var username string = "Bhuvana"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)
}

func booleanPractice() {
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)
}

func smallValuePractice() {
	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("variable is of type:%T \n", smallValue)
}

func floatPractice() {
	var smallFloat float32 = 255.4555555555
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type:%T \n", smallFloat)
}

func string1() {
	var student1 string
	student1 = "John"
	fmt.Println(student1)
}

func string2() {
	var student2 string
	student2 = "John"
	fmt.Println(student2)
}

func Testing() {
	var a, b, c, d int = 1, 3, 5, 7
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func Target() {
	var (
		a int
		b int    = 1
		c string = "hello"
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func Variables() {
	var a, b = 6, "Hello"
	c, d := 7, "World!"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func sum() {
	n := 5
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println("Sum of first", n, "number is :", sum)
}

func dataTypes() {
	var age int = 25
	var pi float64 = 3.1414
	var isActive bool = true
	var name string = "Bhuvana"
	fmt.Println(age, pi, isActive, name)
}

func constants() {
	var x int = 10      //normal declaration of an variable
	y := 20             //shorthand declaration (only inside the functions)
	var name = "Golang" //type inferred automatically
	fmt.Println(x, y, name)
}

func isStatement() {
	num := 10
	if num > 0 {
		fmt.Println("Positive")
	} else if num < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}
}

func switchStatement() {
	day := 5

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Other day")
	}

}

func forLoops() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}
