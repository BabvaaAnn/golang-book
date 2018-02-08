package main

import "fmt"

func main(){
	// Long Declaration
	var x string = "Hello, World"
	fmt.Println(x)

	var y string
	y = "Hello, World"
	fmt.Println(y)

	// Shot Declaration
	// Type Inference
	z := "Hello, World" // ":" --> declare, "=" --> assign value
	fmt.Println(z)
	fmt.Printf("Type: %T\n",z)

	zz := 1234
	fmt.Println(zz)
	fmt.Printf("Type: %T\n",zz)

	// const xx string = "Hello, World"
	// xx = "Other String"
	// fmt.Println(xx)

	var (
		a = 5
		b = 10
		c = 15
	)
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)

	v1, v2 := "first", "sec"
	v1, v2 = v2, v1
	fmt.Println(v1)
	fmt.Println(v2)
}