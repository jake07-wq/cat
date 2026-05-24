package main

import (
	"io"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			return
		}
		return
	}

	hasError := false

	for _, filename := range args {
		file, err := os.Open(filename)
		if err != nil {
			os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
			hasError = true
			continue
		}
		_, err = io.Copy(os.Stdout, file)
		file.Close()
		if err != nil {
			os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
			hasError = true
		}
	}
	if hasError {
		os.Exit(1)
	}
}
