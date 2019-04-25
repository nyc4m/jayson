package main

import (
	_ "bufio"
	"bytes"
	_ "fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//JaysonRequest represents a json request with
//its associated address
type JaysonRequest struct {
	address string
	content []byte
}

func parseTitle(tokens []string) (title string, updatedTokens []string) {
	title = tokens[0]
	updatedTokens = tokens[1:]
	title = strings.TrimLeft(title, "//")
	return
}

func parseContent(tokens []string) (content []byte) {
	content = []byte(strings.Join(tokens, "\n"))
	return
}

func parseFile(binaryContent []byte) (string, []byte) {
	strContent := string(binaryContent)
	tokens := strings.Split(strContent, "\n")
	address, tokens := parseTitle(tokens)
	content := parseContent(tokens)
	return address, content
}

func JaysonFromFile(path string) (*JaysonRequest, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	rawContent, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	address, content := parseFile(rawContent)

	return &JaysonRequest{address: address, content: content}, nil
}

func (j *JaysonRequest) Query() (*http.Response, error) {
	return http.Post(j.address, "application/json", bytes.NewBuffer(j.content))
}
