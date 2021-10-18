package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-session/session"
	"io/ioutil"
	"log"
	"net/http"
)

func auth(page http.ResponseWriter, req *http.Request) {
					store, err := session.Start(context.Background(), page, req)
				if err != nil {
					fmt.Fprint(page, err)
					return
				}
	login := req.FormValue("login")
	password := req.FormValue("password")

	if login == ""{
		fmt.Fprintf(page, "Login cant be nil")
	}else{
		if password == ""{
			fmt.Fprintf(page, "Password cant be nil")
		}else{
			querry := fmt.Sprintf("https://vavylon.herokuapp.com/api?title=auth&token=cardinal&login=%s&password=%s", login, password)
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

				store.Set("active_login", login)
				err = store.Save()
				if err != nil {
					fmt.Fprint(page, err)
					return
				}
				http.Redirect(page, req, "/", 302)
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
