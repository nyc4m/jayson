package main

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	_, err := JaysonFromFile("sample/test.json")
	if err != nil {
		t.Fail()
	}
}

func TestReadAddress(t *testing.T) {
	file, _ := JaysonFromFile("sample/test.json")
	if file.address != "jeuxvideo.com" {
		t.Fail()
	}
}
