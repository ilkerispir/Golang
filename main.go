package main

import ("fmt" 
		"math/rand")

func foo() {
	fmt.Println("The square root of 1-100", rand.Intn(100))
}

func main()  {
	foo()
}