package ascii_art  

import (
	"fmt"
	"os" 
	"strings"
)
// function that maps the inputstring to the banner
func ProcessInput(contents []string, input string) (strArt string) {
	count := 0
	newInput := strings.Split(input, "\n")

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
