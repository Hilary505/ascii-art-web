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

// ProcessInput accepts the contents of the ASCII art file and the input string,
// and processes the input to display the corresponding ASCII art
func ProcessInput(contents []string, input string) (strArt string) {
	count := 0
	// Replace newline and tab characters with their respective escape sequences
	strInput := strings.ReplaceAll(input, "\n", "\\n")
	strInput = strings.ReplaceAll(strInput, "\\t", "    ")
	newInput := strings.Split(strInput, "\\n")

	for _, arg := range newInput {
		if arg == "" {
			count++
			if count < len(newInput) {
				strArt += "\n"
			}
			continue
		}
		for i := 1; i <= 8; i++ {
			for _, ch := range arg {
				if ch > 126 {
					fmt.Println("The text contains an unprintable character", ch)
					os.Exit(0)
				}

				index := int(ch-32)*9 + i
				// check if the index of contents are within range of bannerFile
				if index >= 0 && index < len(contents) {
					strArt += (contents[index])
				}
			}
			strArt += "\n"
		}
	}

	return strArt
}

func FindFile(input, font string) (string, int) {
	var filename string
	switch font {
	case "shadow":
		filename = "shadow.txt"
	case "standard":
		filename = "standard.txt"
	case "thinkertoy":
		filename = "thinkertoy.txt"
	default:
		return "", 500
	}
	return filename, 200
}
