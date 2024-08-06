package ascii_art

import (
	"fmt"
	"os"
	"strings"
)

// GetFile reads from the file specified by filename and returns its contents
func GetFile(filename string) ([]string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("An error", err)
		os.Exit(1)
	}

	if len(file) == 0 {
		fmt.Println("Error: The banner file is empty")
		os.Exit(1)
	}

	myfile := string(file)
	var contents []string

	// Different line splitting logic based on the file type
	if filename == "thinkertoy.txt" {
		contents = strings.Split(myfile, "\r\n")
	} else {
		contents = strings.Split(myfile, "\n")
	}

	return contents, nil
}
