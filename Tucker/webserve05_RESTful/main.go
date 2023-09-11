package main

import (
	"github.com/darkkaiser/study-go/webserve05_RESTful/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
