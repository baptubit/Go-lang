package main

import "fmt"

func main() {
	fmt.Println("Welcome to the world of GO-lang")
	fmt.Println("I am in main")
	foo()
	fmt.Println("Lets do some iterative stuff")
	FUNC2()
	DeclareVariable()
	for i := 0; i < 10; i++ {
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

func DeclareVariable(){
	x:=45 //first time to declare variable we should use ":="
	fmt.Println("Value of initilized Variable=",x)
	x=90
	fmt.Println("Changed value of x  is=",x)
}

//Control flow:
//1.Sequence
//2.loop;iterative
//3.Conditional
