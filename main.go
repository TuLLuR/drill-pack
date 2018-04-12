package main

import (
	"fmt"
	"net/http"
	"time"

	"./src/lissajous"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func lissajousCurve(w http.ResponseWriter, r *http.Request) {
	lissajous.Lissajous(w)
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/lissajous", lissajousCurve)
	fmt.Println("It Works!")
	http.ListenAndServe(":8080", nil)
}
