package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"

    "jaytaylor.com/html2text"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: go run main.go <url> <output_file>")
        os.Exit(1)
    }

    url := os.Args[1]
    outputFile := os.Args[2]

    if err := convertHTMLToText(url, outputFile); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
}

func convertHTMLToText(url, outputFile string) error {
    // Download HTML content
    resp, err := http.Get(url)
    if err != nil {
        return fmt.Errorf("failed to download HTML: %v", err)
    }
    defer resp.Body.Close()

    // Read HTML content
    html, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return fmt.Errorf("failed to read HTML content: %v", err)
    }

    // Convert HTML to text
    text, err := html2text.FromString(string(html), html2text.Options{PrettyTables: false})
    if err != nil {
        return fmt.Errorf("failed to convert HTML to text: %v", err)
    }

    // Write text to output file
    err = ioutil.WriteFile(outputFile, []byte(text), 0644)
    if err != nil {
        return fmt.Errorf("failed to write text to file: %v", err)
    }

    fmt.Printf("HTML converted to text and saved to: %s\n", outputFile)
    return nil
}
