package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("messsages.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	var currentLine string

	for {
		buf := make([]byte, 8)
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}

		chunk := string(buf[:n])
		parts := strings.Split(chunk, "\n")

		// Print all complete lines
		for _, part := range parts[:len(parts)-1] {
			fmt.Printf("read: %s\n", currentLine+part)
			currentLine = ""
		}

		// Save the last part (could be incomplete)
		currentLine += parts[len(parts)-1]
	}

	// Print any remaining text after EOF
	if currentLine != "" {
		fmt.Printf("read: %s\n", currentLine)
	}
}
