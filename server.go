package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	index, _ := os.Open("index.html")
	data, _ := ioutil.ReadAll(index)
	w.Write([]byte(data))
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	x, y := queries["x"][0], queries["y"][0]
	xf, _ := strconv.ParseFloat(x, 64)
	yf, _ := strconv.ParseFloat(y, 64)
	ans := strconv.FormatFloat(xf+yf, 'f', -1, 64)
	printer := x + " + " + y + " = " + ans + "\nThanks for using our calculator!"
	w.Write([]byte(printer))
}

func SubHandler(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	x, y := queries["x"][0], queries["y"][0]
	xf, _ := strconv.ParseFloat(x, 64)
	yf, _ := strconv.ParseFloat(y, 64)
	ans := strconv.FormatFloat(xf-yf, 'f', -1, 64)
	printer := x + " - " + y + " = " + ans + "\nThanks for using our calculator!"
	w.Write([]byte(printer))
}

func MultHandler(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	x, y := queries["x"][0], queries["y"][0]
	xf, _ := strconv.ParseFloat(x, 64)
	yf, _ := strconv.ParseFloat(y, 64)
	ans := strconv.FormatFloat((xf * yf), 'f', -1, 64)
	printer := x + " * " + y + " = " + ans + "\nThanks for using our calculator!"
	w.Write([]byte(printer))
}

func DivHandler(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	x, y := queries["x"][0], queries["y"][0]
	xf, _ := strconv.ParseFloat(x, 64)
	yf, _ := strconv.ParseFloat(y, 64)
	ans := strconv.FormatFloat((xf / yf), 'f', -1, 64)
	printer := x + " / " + y + " = " + ans + "\nThanks for using our calculator!"
	w.Write([]byte(printer))
}

func main() {
	r := mux.NewRouter()
	fmt.Println("Listening on port 8000\n")
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/add", AddHandler)
	r.HandleFunc("/sub", SubHandler)
	r.HandleFunc("/mult", MultHandler)
	r.HandleFunc("/div", DivHandler)
	//Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
