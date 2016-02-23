package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/sum/", sum)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	f, _ := ioutil.ReadFile("./index.html")
	w.Write(f)
}

func sum(w http.ResponseWriter, r *http.Request) {
	x, _ := strconv.Atoi(r.URL.Query().Get("x"))
	y, _ := strconv.Atoi(r.URL.Query().Get("y"))
	fmt.Fprint(w, x+y)
}
