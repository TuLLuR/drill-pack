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
	t.Execute(w, &page{Title: "Just Page", Msg: "Hello, World!"})
}

func lissajousCurve(w http.ResponseWriter, r *http.Request) {
	lissajous.Lissajous(w)
}

func lissIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	t, _ := template.ParseFiles("html/lissajous.html")
	t.Execute(w, &page{Title: "Lissajous", Msg: "GIF Image"})
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/lissajous", lissajousCurve)
	http.HandleFunc("/index", index)
	http.HandleFunc("/lissIndex", lissIndex)
	fmt.Println("It Works!")
	http.ListenAndServe(":8080", nil)
}
