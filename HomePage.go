package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"github.com/go-session/session"
)


func HomePage(page http.ResponseWriter, req *http.Request) {
		store, err := session.Start(context.Background(), page, req)
	if err != nil {
		fmt.Fprint(page, err)
		return
	}


	active, ok := store.Get("active_login")
	if ok {
		fmt.Fprintf(page, "Nice session %s", active)
	}else{
		fmt.Fprintf(page, "Error session" )
	}

	temp, err := template.ParseFiles("templates/html/HomePage.html")

	if err != nil {
		fmt.Fprintf(page, err.Error())
	}

	temp.ExecuteTemplate(page, "home_page", nil)



}
