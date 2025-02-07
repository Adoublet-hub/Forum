package module

import (
	"fmt"
	"html/template"
	"net/http"
)

// Load user page
func Index(w http.ResponseWriter, r *http.Request) {
	cookieUser, cookieToken, _ := Checklog(w, r)
	database.User.Username = cookieUser
	database.User.Token = cookieToken
	if r.Method == "GET" {
		t, err := template.ParseFiles("./templates/blog.html")
		if err != nil {
			fmt.Println(err)
			handle500(w, err)
			return
		}
		err = t.Execute(w, database)
		if err != nil {
			fmt.Println(err)
			handle500(w, err)
			return
		}
	}
}
