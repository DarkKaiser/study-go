package main

import (
	"github.com/darkkaiser/study-go/webserve11_TodoList/app"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

func main() {
	m := app.MakeHandler()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
