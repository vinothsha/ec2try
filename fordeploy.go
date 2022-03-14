package main

import (
	"fmt"
	"net/http"
)

func Hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "hello AWS ec2 instance ")
}

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}
