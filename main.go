package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"./src/lissajous"
)

type page struct {
	Title string
	Msg   string
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	t, _ := template.ParseFiles("html/index.html")
	t.Execute(w, &page{Title: "Title", Msg: "Message"})
}

func lissajousCurve(w http.ResponseWriter, r *http.Request) {
	lissajous.Lissajous(w)
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/lissajous", lissajousCurve)
	http.HandleFunc("/index", index)
	fmt.Println("It Works!")
	http.ListenAndServe(":8080", nil)
}
