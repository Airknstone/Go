// //pass_fail reposts whether a grade is passing or failing.
// // The go fmt command automatically reformats source files to use Go standard formatting. You should run go fmt on any code that you plan to share with others.

// // The go build command compiles Go source code into a binary format that computers can execute.

// // The go run command compiles and runs a program without saving an executable file in the current directory.
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	fmt.Print("Enter a grade: ")
// 	reader := bufio.NewReader(os.Stdin)
// 	input, err := reader.ReadString('\n')
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	input = strings.TrimSpace(input)
// 	grade, err := strconv.ParseFloat(input, 64)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var status string
// 	if grade >= 60 {
// 		status = "passing"
// 	} else {
// 		status = "failing"
// 	}

// 	fmt.Println("A grade of", grade, "is", status)
// }
