package api

import (
	"app/models"
	"fmt"
	"net/http"
	"strings"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	firstname := r.Header.Get("firstname")
	lastname := r.Header.Get("lastname")

	if firstname == " " {
		fmt.Fprint(w, "Please input firstname")
		return
	}
	if lastname == " " {
		fmt.Fprint(w, "Please input lastname")
		return
	}
	firstnam := strings.Trim(firstname, "!*#&^%@~:;/][{}")
	lastnam := strings.Trim(lastname, "!*#&^%@~:;/][{}")

	senn := models.User{
		Firstname: firstname,
		Lastname:  lastname,
	}
	fmt.Fprint(w, "user:"+firstnam+" "+lastnam+" "+"has been created")
	fmt.Println(senn)

}
