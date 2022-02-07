package models

type User struct {
	ID    int    `jsonapi:"primary,users"`
	Name  string `jsonapi:"attr,name"`
	Password  string `jsonapi:"attr,password"`
	Email string `jsonapi:"attr,email"`
}
