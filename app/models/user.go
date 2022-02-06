package models

type User struct {
	ID    int    `jsonapi:"primary,users"`
	Name  string `jsonapi:"name"`
	Email string `jsonapi:"email"`
}
