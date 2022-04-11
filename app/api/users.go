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
	}
	if lastname == " " {
		fmt.Fprint(w, "Please input lastname")
	}
	firstname = strings.Trim(firstname, "!*#&^%@~:;/][{}")
	lastname = strings.Trim(lastname, "!*#&^%@~:;/][{}")

	senn := models.User{
		Firstname: firstname,
		Lastname:  lastname,
	}
	fmt.Fprint(w, "user:"+firstname+" "+lastname+"has been created")

}
