package main

import (
	"io"
	"net/http"
)

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func bar(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "bar ran")
}

func main() {
	http.HandleFunc("/", foo)     //GET http://localhost:8080/
	http.HandleFunc("/dog/", bar) //GET: http://localhost:8080/dog/
	http.ListenAndServe(":8080", nil)
}
