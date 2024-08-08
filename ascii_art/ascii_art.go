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
		return nil, err
	}

	if len(file) == 0 {
		err := fmt.Errorf("file is empty")
		return nil, err
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
