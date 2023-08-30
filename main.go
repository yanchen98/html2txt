package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"jaytaylor.com/html2text"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <url_or_file> <output_file>")
		os.Exit(1)
	}

	source := os.Args[1]
	outputFile := os.Args[2]

	var err error
	var html []byte

	if strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://") {
		html, err = downloadHTML(source)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	} else {
		html, err = ioutil.ReadFile(source)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}

	if err := convertHTMLToText(html, outputFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func downloadHTML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to download HTML: %v", err)
	}
	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read HTML content: %v", err)
	}

	return html, nil
}

func convertHTMLToText(html []byte, outputFile string) error {
	text, err := html2text.FromString(string(html), html2text.Options{PrettyTables: false})
	if err != nil {
		return fmt.Errorf("failed to convert HTML to text: %v", err)
	}

	err = ioutil.WriteFile(outputFile, []byte(text), 0644)
	if err != nil {
		return fmt.Errorf("failed to write text to file: %v", err)
	}

	fmt.Printf("HTML converted to text and saved to: %s\n", outputFile)
	return nil
}
