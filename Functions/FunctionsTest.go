package main

import "fmt"

func main() {
	// simpleFunc()
	// paramFunc(10, 15)
	// res := returnFunc(10, 20)
	// fmt.Println("Res:", res)

	// //return all the values
	// sum, product := multipleReturnFunc(3, 4)
	// fmt.Println(sum, product)

	// //return single value with blank identifier
	// _, product := multipleReturnFunc(3, 4)
	// fmt.Println(product)

	// //methods
	// p1 := Person{name: "Alice", age: 25}
	// p1.greet() //calling the method

	// //interface
	// var s Speaker
	// s = Human{name: "Alice"}
	// fmt.Println(s.Speak())
	deferExample()

}

// A function in Go is a named block of code that performs a specific task and can be called multiple times(Reusable code)
func simpleFunc() {
	fmt.Println("Hello , Bhuvana")
}

func paramFunc(a int, b int) {
	fmt.Println("Sum:", a+b)
}

func returnFunc(a int, b int) int {
	return a + b
}

func multipleReturnFunc(a int, b int) (int, int) {
	return a + b, a * b
}

//Methods-A method is a function that has a receiver , allowing it to be associated with a type

//struct
type Person struct {
	name string
	age  int
}

//method
func (p Person) greet() {
	fmt.Println("Hello my name is ", p.name)
}

//interfaces-An interface in Go is a type that specifies method signatures ,
//  and type implementing those methods implicitly satisfies the interface

type Speaker interface {
	Speak() string
}
type Human struct {
	name string
}

func (h Human) Speak() string {
	return "Hello , my name is " + h.name
}

//Defer-defer is a keyword in Go used to delay the execution of a function until the surrounding function returns.(LIFO)

func deferExample() {
	defer fmt.Println("First")
	defer fmt.Println("Second")
	fmt.Println("Hello")

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
}
