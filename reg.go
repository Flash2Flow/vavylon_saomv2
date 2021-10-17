package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func reg(page http.ResponseWriter, req *http.Request) {
	login := req.FormValue("login")
	password := req.FormValue("password")
	email := req.FormValue("email")

	if login == ""{
		fmt.Fprintf(page, "Login cant be nil")
	}else{
		if password == ""{
			fmt.Fprintf(page, "Password cant be nil")
		}else{
			if email == ""{
				fmt.Fprintf(page, "Email cant be nil")
			}else{
				querry := fmt.Sprintf("https://vavylon.herokuapp.com/api?title=reg&token=cardinal&login=%s&password=%s&email=%s", login, password, email)
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
					fmt.Fprintf(page, "Вы были успешно зарегистрированы, теперь авторизируйтесь. И помните, сайт в бета тестировании...")
				}

				if res["Code"] == "999 Unknown error ( бтв такой пользователь уже есть )"{
					log.Println("Bad request Already have this login/email")
					fmt.Fprintf(page, "Пользователь с таким логином или мылом уже есть... Напоминаю, всё ещё бета тест...")
				}

			}
		}
	}
}