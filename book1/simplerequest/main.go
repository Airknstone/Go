package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	//Create the variables for the respone and error
	var r *http.Response
	var err error

	//Request index.html from example.com
	r, err = http.Get("http://www.example.com/index.html")

	//If ther eis a problem stop the program and pring the error
	if err != nil {
		panic(err)
	}

	//Check the status code returned by the server
	if r.StatusCode == 200 {
		//the request was sucessful!
		var webPageContent []byte

		//We know the size of the response is 1270 from previous example
		var bodyLength int = 1270

		//Initialize the byte array to the size of the data
		webPageContent = make([]byte, bodyLength)

		//read the data from the server
		r.Body.Read(webPageContent)

		// Open a writable file on your computer (create if it does not  exist)
		var out *os.File
		out, err = os.OpenFile("index.html", os.O_CREATE|os.O_WRONLY, 0664)

		if err != nil {
			panic(err)
		}

		// Write the contents to a file
		out.Write(webPageContent)
		out.Close()
	} else {
		log.Fatal("Failed to retrieve the webpage. Received status code",
			r.Status)
	}
}
