package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

// Page holds all the information we need to generate a new
// HTML page from a text file on the filesystem.
type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content      string
}

func main() {
	filePath := "filePath"
	// Making a flag
	filePtr := flag.String("file", " ", "File to read.")
	flag.Parse()
	// flag function
	htmlPath := func() string {
		return strings.Replace(*filePtr, ".txt", ".html", -1)
	}
	// read the file
	fileContents, err := ioutil.ReadFile(*filePtr)
	// creating a new Page object
	page := Page{
		TextFilePath: filePath,
		TextFileName: "first",
		HTMLPagePath: htmlPath(),
		Content:      string(fileContents),
	}
	// create the template t
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	// Create a new, blank HTML file.
	newFile, err := os.Create(htmlPath())
	if err != nil {
		panic(err)
	}
	// inject the newly created page into the new htmlfile named new.html
	t.Execute(newFile, page)
}
