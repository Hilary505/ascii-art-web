package handlers


import (
	a "ascii/ascii_art"
     "net/http"
	 "html/template"
	 "log"
)

type ExecOutput struct {
	In  string
	Out string
	Font  string
}
func ValidAscii(s string) bool {
	for _, i := range []byte(s) {
		if i > 127 {
			return false
		}
	}
	return true
}

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
			t.Execute(w, ExecOutput{
                In:   "",
                Font: "standard", // default font
                Out:  "",
            })
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
					StatusBadRequest(w)
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
				
				ex := ExecOutput{
					In:  input,
					Out: output,
					Font: font,
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
			StatusNotFound(w)
			return
		}
		t.Execute(w, nil)
	}
}


func internalServerError(w http.ResponseWriter) {
    w.WriteHeader(http.StatusInternalServerError)
    t, err := template.ParseFiles("errors/500.html")
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, nil)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}

func  StatusNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
    t, err := template.ParseFiles("errors/404.html")
    if err != nil {
        internalServerError(w)
        return
    }
    err = t.Execute(w, nil)
    if err != nil {
        internalServerError(w)
    }
}

func  StatusBadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
    t, err := template.ParseFiles("errors/400.html")
    if err != nil {
        internalServerError(w)
        return
    }
    err = t.Execute(w, nil)
    if err != nil {
        internalServerError(w)
    }
}