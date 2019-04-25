package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	var path string
	flag.StringVar(&path, "path", "", "Path of the file to analyse")
	flag.Parse()
	if path == "" {
		panic("No args given")
	}

	jay, err := JaysonFromFile(path)
	if err != nil {
		panic(err)
	}
	resp, err := jay.Query()
	if err != nil {
		fmt.Println("Error while querying")
	}

	allBytes, _ := ioutil.ReadAll(resp.Body)
	content := string(allBytes)
	fmt.Println(content)
}
