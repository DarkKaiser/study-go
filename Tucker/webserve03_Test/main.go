package main

import (
	"github.com/darkkaiser/study-go/webserve03_Test/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
