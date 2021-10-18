package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/go-session/session"
	"context"
)


func HomePage(page http.ResponseWriter, req *http.Request) {
	temp, err := template.ParseFiles("templates/html/HomePage.html")

	if err != nil {
		fmt.Fprintf(page, err.Error())
	}
		temp.ExecuteTemplate(page, "home_page", nil)

	login := req.FormValue("login")
	password := req.FormValue("password")

	if login == ""{
		fmt.Fprintf(page, "Login cant be nil")
	}else{
		if password == ""{
			fmt.Fprintf(page, "Password cant be nil")
		}else{
			querry := fmt.Sprintf("https://vavylon.herokuapp.com/api?title=reg&token=cardinal&login=%s&password=%s", login, password)
			resp, err := http.Get(querry)
			if err != nil {
				print(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				print(err)
			}
			//fmt.Print(string(body))
			var res map[string]interface{}
			json.Unmarshal([]byte(body), &res)
			if res["Status"] == "OK"{
				log.Println("Nice request")
				store, err := session.Start(context.Background(), page, req)
				if err != nil {
					fmt.Fprint(page, err)
					return
				}

				store.Set("active_login", login)
				err = store.Save()
				if err != nil {
					fmt.Fprint(page, err)
					return
				}


				fmt.Fprintf(page, "Вы были успешно авторизированы. И помните, сайт в бета тестировании...")
			}

			if res["Code"] == "202 Incorrect Token Cardinal"{
				log.Println("Bad request. 202 Incorrect Token Cardinal")
				fmt.Fprintf(page, "202 Incorrect Token Cardinal. И помните, сайт в бета тестировании...")
			}

			if res["Code"] == "??? Empty url login"{
				log.Println("Bad request. ??? Empty url login")
				fmt.Fprintf(page, "??? Empty url login. И помните, сайт в бета тестировании...")
			}

			if res["Code"] == "??? Empty url password"{
				log.Println("Bad request. ??? Empty url password")
				fmt.Fprintf(page, "??? Empty url password. И помните, сайт в бета тестировании...")
			}

			if res["Code"] == "??? Empty url password"{
				log.Println("Bad request. ??? Empty url password")
				fmt.Fprintf(page, "??? Empty url password. И помните, сайт в бета тестировании...")
			}


			if res["Code"] == "304 User not found"{
				log.Println("Bad request. 304 User not found")
				fmt.Fprintf(page, "304 User not found. И помните, сайт в бета тестировании...")
			}

			if res["Code"] == "302 Bad password"{
				log.Println("Bad request. 302 Bad password")
				fmt.Fprintf(page, "302 Bad password. И помните, сайт в бета тестировании...")
			}

		}
	}

}
