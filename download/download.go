package main

import (
	"io"
	"net/http"
	"os"
)

// type greeting string

// func (g greeting) Greet() {
// 	fmt.Println("Hello Universe")
// }

// // exported
// var Greeter greeting

type urlString string

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func (u urlString) DownloadFile(url string) {
	// url := string(u) //"https://golangcode.com/images/avatar.jpg"
	filepath := "avatar.jpg"
	println("Download at____%s", filepath)
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		println("error 1:-->", err, 1)
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		println("error 2:-->", err, url)
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		println("error 3:-->%s", err)
	}

	println("---->>>Done Download url:%s", url)
}

var DownloadIDE urlString
