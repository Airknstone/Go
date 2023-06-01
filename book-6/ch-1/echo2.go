// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for l, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
		fmt.Println(l)
	}
	fmt.Println(s)
	fmt.Println(os.Args[0])
}
