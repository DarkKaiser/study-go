package main

import (
	"html/template"
	"os"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) IsOld() bool {
	return u.Age > 30
}

func main() {
	user := User{
		Name:  "tucker",
		Email: "tucker@naver.com",
		Age:   43,
	}
	user2 := User{
		Name:  "darkkaiser",
		Email: "darkkaiser@naver.com",
		Age:   22,
	}
	users := []User{user, user2}

	tmpl, err := template.New("Tmpl1").ParseFiles("templates/tmpl1.tmpl", "templates/tmpl2.tmpl")
	if err != nil {
		panic(err)
	}

	//tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", user)
	//tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", user2)
	tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", users)
}
