package routes

import (
	"app/api"
	"app/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var r = mux.NewRouter()

func Start() {
	Routes()
	log.Fatal(http.ListenAndServe(":2020", r))
}
func Routes() {
	r.HandleFunc("/home", handlers.Appenduser)
	// Api for User
	// r.HandleFunc("/api/user/get", api.GetUser)
	r.HandleFunc("/api/user/create", api.CreateUser)
	// r.HandleFunc("/api/user/update", api.UpdateUser)
	// r.HandleFunc("/api/user/delete", api.DeleteUser)
	// r.HandleFunc("/api/user/addfriend", api.Newfriend)
}
