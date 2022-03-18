package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func process(w http.ResponseWriter, r *http.Request) {
	value, _ := parser(r.FormValue("value"))
	t, _ := template.ParseFiles("./template/index.tmpl")
	t.Execute(w, value)
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(wrapper{http.Dir("./static")})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fs))
	mux.HandleFunc("/", process)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}

type wrapper struct {
	fs http.FileSystem
}

//wrapping around FileServer to deny access to internal files
//method Open() gets callled each time http.FileServer receives a request
func (w wrapper) Open(path string) (http.File, error) {
	f, err := w.fs.Open(path)

	if err != nil {
		return nil, err
	}

	s, err := f.Stat()

	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := w.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}
	return f, err
}
