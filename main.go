package main

import (
	a "ascii/ascii_art"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type ExecOutput struct {
	In  string
	Out string
}

// ValidAscii checks if the input string contains only ASCII characters.
func ValidAscii(s string) bool {
	for _, i := range []byte(s) {
		if i > 127 {
			return false
		}
	}
	return true
}

// internalServerError handles internal server errors.
func internalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	t, err := template.ParseFiles("error/500.html")
	if err != nil {
		log.Printf("Error parsing 500.html template: %v", err)
		return
	}
	if err := t.Execute(w, nil); err != nil {
		log.Printf("Error executing 500.html template: %v", err)
	}
}

// Handler handles the HTTP requests.
// Handler handles the HTTP requests.
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		switch r.Method {
		case "GET":
			t, err := template.ParseFiles("index.html")
			if err != nil {
				internalServerError(w)
				return
			}
			t.Execute(w, nil)
		case "POST":
			if err := r.ParseForm(); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				log.Printf("Error parsing form: %v", err)
				return
			}
			input := r.Form.Get("input")
			font := r.Form.Get("font")
			if !ValidAscii(input) {
				w.WriteHeader(http.StatusBadRequest)
				t, err := template.ParseFiles("error/400.html")
				if err != nil {
					internalServerError(w)
					return
				}
				t.Execute(w, nil)
			} else {
				file, status := a.FindFile(input, font)
				if status == 500 {
					internalServerError(w)
					return
				}
				contents, err := a.GetFile(file)
				if err != nil {
					internalServerError(w)
					return
				}
				output := a.ProcessInput(contents, input)
				log.Printf("method: %v / font: %v / input: %v / statuscode: %v\n", r.Method, font, input, status)
				ex := ExecOutput{
					In:  input,
					Out: output,
				}
				t, err := template.ParseFiles("index.html")
				if err != nil {
					internalServerError(w)
					return
				}
				t.Execute(w, ex)
			}
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		t, err := template.ParseFiles("error/404.html")
		if err != nil {
			internalServerError(w)
			return
		}
		t.Execute(w, nil)
	}
}

// main starts the HTTP server.
func main() {
	log.Println("Server is starting...")
	http.HandleFunc("/", Handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	fmt.Println("Server up at port 8080\nhttp status :", http.StatusOK)
	//Openbrowser("http.localhost:8080")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
