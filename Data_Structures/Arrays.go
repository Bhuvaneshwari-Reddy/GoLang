package main

import (
	"fmt"
)

func main() {
	// array()
	// slice()
	// maps()
	// structs()
	// pointers()
	// pointers1()
	Human()

	bhuvana := User{"Bhuvana", "bhuvana@gmail.com", true, 26}
	fmt.Println(bhuvana)
	bhuvana.GetStatus()
	deferPrac()

}

func array() {
	//An array is a fixed-size, ordered collection of elements of the same type.
	var numbers [3]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30

	// var numbers = [3]int{10, 20, 30}

	fmt.Println("Array:", numbers)
	fmt.Println("Length:", len(numbers))

	for i, val := range numbers {
		fmt.Printf("Index %d->%d\n", i, val)
	}

}

func slice() {
	//A slice is a dynamic, flexible view into an array.Unlike arrays, slices can grow or shrink.(they don't have fixed length)

	nums := []int{10, 20, 30}
	fmt.Println("Slice", nums)

	//Append elements
	nums = append(nums, 40, 50)
	fmt.Println("After append:", nums)

	//Slice from slice
	part := nums[1:4]
	fmt.Println("Sub slice:", part)

	//Iterate
	for i, val := range nums {
		fmt.Printf("Index %d->%d\n", i, val)
	}

}

func maps() {
	//A map is a built-in data structure that stores key-value pairs, like dictionaries in Python or objects in JavaScript.
	fruits := map[string]int{
		"apple":  5,
		"banana": 10,
		"mango":  7,
	}

	fmt.Println("Map:", fruits)
	fmt.Println("apple count:", fruits["apple"])

	//add new items
	fruits["grapes"] = 6
	fmt.Println("Map after adding:", fruits)

	//delete item
	delete(fruits, "banana")
	fmt.Println("Map after deleting:", fruits)

	//check if key exists
	val, exits := fruits["banana"]
	if exits {
		fmt.Println("Banana count:", val)
	} else {
		fmt.Println("Banana not found")
	}

	//iterate over map
	for key, value := range fruits {
		fmt.Printf("%s->%d\n", key, value)
	}

}

// A struct is a custom data type that groups related data fields together.It’s like a class without methods (you can attach methods later).
type Person struct {
	name string
	age  int
}

func structs() {

	//create a struct (similar to object creation)
	p1 := Person{name: "Alice", age: 25}
	fmt.Println("Person:", p1)
	fmt.Printf("Details:%+v\n", p1)

	//Access fields
	fmt.Println("Name", p1.name)
	fmt.Println("Age", p1.age)

	//Modify field
	p1.age = 26
	fmt.Println("Updated age", p1.age)

	//Anonymous struct
	car := struct {
		brand string
		year  int
	}{"Tesla", 2024}
	fmt.Println("Car:", car)

}

// * is the symbol used for pointers and reference means &
// Pointer is a variable that is used to store the memory address of another variable.
/*
* Operator also termed as the dereferencing operator used to declare pointer variable and access the value stored in the address.
& operator termed as address operator used to returns the address of a variable or to access the address of a variable to a pointer.
*/
func pointers() {
	fmt.Println("Welcome to a class on pointers")
	// var ptr *int
	// fmt.Println("value of the pointer is ", ptr)
	myNumber := 20
	var ptr = &myNumber
	fmt.Println("value of the actual pointer is ", ptr)
	fmt.Println("value of the actual pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is :", myNumber)
}

func pointers1() {
	myName := "Bhuvana"
	var ptr = &myName
	fmt.Println("Value of pointer", ptr)
	fmt.Println("Value of pointer", *ptr)
}

//methods

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User active: ", u.Status)

}

// defer-delay the execution of the function until the nearby function returns  (lifo)
func deferPrac() {
	defer fmt.Println("Hello")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("World")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)

	}
}

type Persons struct {
	name string
	age  int
}

func Human() {
	p := Persons{name: "Bhuvana", age: 26}
	fmt.Printf("Persons", p)
}
