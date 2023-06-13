package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello world!")
	x := complex(2.5, 3.1)
	y := complex(10.2, 2)
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(real(x))
	fmt.Println(imag(x))
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("printing", j, "from inside of ananonymous function")
		}(i)
	}
}
