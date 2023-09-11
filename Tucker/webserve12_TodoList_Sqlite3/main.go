package main

import (
	"github.com/darkkaiser/study-go/webserve12_TodoList_Sqlite3/app"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

func main() {
	m := app.MakeHandler("./test.db")
	defer m.Close()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
