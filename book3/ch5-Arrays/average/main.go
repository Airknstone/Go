package main

import (
	"datafile"
	"fmt"
	"log"
)

func main() {

	numbers, err := datafile.GetFloats("/home/nodejs/golang/book3/ch5-Arrays/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64

	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)

}
