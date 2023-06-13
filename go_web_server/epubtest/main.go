package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	// Open the EPUB file
	epubFile, err := zip.OpenReader("./englishlanguage00kell_englishlanguage00kell.epub")
	if err != nil {
		log.Fatal(err)
	}
	defer epubFile.Close()

	// Iterate over the files in the EPUB archive
	for _, file := range epubFile.File {
		// Check if the file is an HTML file (content file)
		if filepath.Ext(file.Name) == ".html" {
			// Open the content file
			content, err := file.Open()
			if err != nil {
				log.Fatal(err)
			}
			defer content.Close()

			// Read the content of the HTML file
			htmlContent, err := ioutil.ReadAll(content)
			if err != nil {
				log.Fatal(err)
			}

			// Print the content of the HTML file
			fmt.Println(string(htmlContent))
		}
	}
}
