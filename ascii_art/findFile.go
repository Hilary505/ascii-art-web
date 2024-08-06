package ascii_art 

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
