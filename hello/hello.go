package main

import (
	"bytes"
	"net/http"
	"os"
	"time"
)

var name = []byte{}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if string(name) == "" {
			html, _ := os.ReadFile("public/name.html")
			w.Write(html)
			return
		}
		html, _ := os.ReadFile("public/index.html")
		named := bytes.ReplaceAll(html, []byte("{name}"), name)
		w.Write(bytes.ReplaceAll(named, []byte("{time}"), []byte(time.Now().GoString())))
	})

	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		r.ParseForm()
		name = []byte(r.PostForm.Get("name"))
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.ListenAndServe(":8080", nil)
}
