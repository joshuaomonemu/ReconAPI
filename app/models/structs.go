package models

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Userid    string `json:"userid"`
}
