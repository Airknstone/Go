//pass_fail #2 reposts whether a grade is passing or failing.
// The go fmt command automatically reformats source files to use Go standard formatting. You should run go fmt on any code that you plan to share with others.

// The go build command compiles Go source code into a binary format that computers can execute.

// The go run command compiles and runs a program without saving an executable file in the current directory.
package passFail

import (
	"fmt"
	"keyboard" // imports keyboard module we created to handle floats
	"log"
)

func PassFail() {
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "Failing"
	}
	//Expected output if >= 60 passing else failing
	fmt.Println("A grade of", grade, "is", status)

}
