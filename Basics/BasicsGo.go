package main

import "fmt"

func main() {
	// variables()
	// stringPractice()
	// pointers()
	pointers1()
	// structs()
	// structures()
	// arrays()
	// slices()
	// maps()
	// mapLiteral()
}

//A variable is a named memory location used to store a value of a specific data type.
func variables() {

	var a = "initail"
	fmt.Println(a)

	var b, c = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	// f := "apple" //shorthand declaration
	var f string = "apple"
	fmt.Println(f)

}

func stringPractice() {
	var username string = "Bhuvana"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)
}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

//A pointer stores the memory address of another variable, allowing direct access and modification of the original value.
func pointers() {
	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)
	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", i)
}

func pointers1() {
	x := 10
	p := &x
	fmt.Println("x =", x)
	fmt.Println("address of x =", p)
	fmt.Println("value at pointer =", *p)
}

//A struct is a user-defined data type that groups related variables of different types into a single unit.
type Person struct {
	name string
	age  int
}

func structs() {
	var u Person
	u.name = "Reddy"
	u.age = 26
	fmt.Println(u)
}

type Address struct {
	city  string
	state string
}
type Employee struct {
	name    string
	address Address
}

func printEmp(u Employee) {
	fmt.Println("Name: ", u.name)
	fmt.Println("City: ", u.address.city)
	fmt.Println("State: ", u.address.state)
}

func structures() {
	u := Employee{
		name: "Bhuvana",
		address: Address{
			city:  "Bangalore",
			state: "Karnataka",
		},
	}
	printEmp(u)
	// fmt.Println(u)
}

//An array is a fixed-size collection of elements of the same data type.
func arrays() {
	//inline initialization
	nums := [3]int{10, 20, 30}
	// var nums [3]int
	// nums[0] = 10
	// nums[1] = 20
	// nums[2] = 30
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

}

//An slice is a dynamically sized , reference type that represents a portion of an array and provides flexible access to the sequence of elements
func slices() {
	nums := []int{2, 4, 6}
	nums = append(nums, 8)
	for _, num := range nums {
		fmt.Println(num)
	}

}

//maps-Maps are Go’s built-in associative data type used to store the data in key and value pair(unordered), key-unique
//syntax :map[keyType]valueType

//range-range iterates over elements in a variety of data structures.
func maps() {
	//using make
	ages := make(map[string]int)
	ages["Alice"] = 25
	ages["bob"] = 22
	for name, age := range ages {
		fmt.Println(name, age)
	}
}

func mapLiteral() {
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	for name, age := range ages {
		fmt.Println(name, age)
	}
	delete(ages, "Bob")
}
