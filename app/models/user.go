package models

type User struct {
	ID    int    `jsonapi:"primary,users"`
	Name  string `jsonapi:"attr,name"`
	Email string `jsonapi:"attr,email"`
}
