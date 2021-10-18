package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"github.com/go-session/session"
)


func HomePage(page http.ResponseWriter, req *http.Request) {
	temp, err := template.ParseFiles("templates/html/HomePage.html")

	if err != nil {
		fmt.Fprintf(page, err.Error())
	}

	temp.ExecuteTemplate(page, "home_page", nil)

	store, err := session.Start(context.Background(), page, req)
	if err != nil {
		fmt.Fprint(page, err)
		return
	}



}
