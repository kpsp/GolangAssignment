package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func downloadFile(URL string) (bool, error) {
	tokens := strings.Split(URL, "/")
	fileName := tokens[len(tokens)-1]
	fmt.Println("Downloading", URL, "to", fileName)

	output, err := os.Create(fileName)
	if err != nil {
		return false, err
	}
	defer output.Close()

	response, err := http.Get(URL)
	if err != nil {
		return false, err
	} else {
		defer response.Body.Close()
		_, err = io.Copy(output, response.Body)
		if err != nil {
			return false, err
		} else {
			fmt.Println("Downloaded", fileName)
		}
	}
	return true, nil
}

func downloadMultipleFiles(urls []string) ([]bool, error) {
	done := make(chan bool, len(urls))
	errch := make(chan error, len(urls))
	for _, URL := range urls {
		go func(URL string) {
			b, err := downloadFile(URL)
			done <- b
			errch <- err
		}(URL)
	}
	var errStr string
	var tests []bool
	for i := 0; i < len(urls); i++ {
		test := <-done
		tests = append(tests, test)
		if err := <-errch; err != nil {
			errStr = errStr + " " + err.Error()
		}
	}
	var err error
	if errStr != "" {
		err = errors.New(errStr)
	}
	return tests, err
}

func main() {
	var urls []string
	urls = []string{
		"http://212.183.159.230/5MB.zip",
		"https://www.africau.edu/images/default/sample.pdf",
		"https://instagram.com/favicon.ico",
		"https://media.geeksforgeeks.org/wp-content/uploads/gfg-40.png",
	}
	_, err := downloadMultipleFiles(urls)
	//fmt.Println(d)
	fmt.Println(err)
}
