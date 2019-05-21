package main

import "fmt"

func main() {
	fmt.Println("Welcome to the world of GO-lang")
	fmt.Println("I am in main")
	foo()
	fmt.Println("Lets do some iterative stuff")
    FUNC2()
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println("The current loop=number is =", i)
		}

	}
}

func foo() {
	fmt.Println("I am in foo")
}

func FUNC2(){
	n,err:=fmt.Println("Welcome to the world of GO-lang")
	fmt.Println("Number of bytes printed=",n)
	fmt.Println(err)	
}

//Control flow:
//1.Sequence
//2.loop;iterative
//3.Conditional
