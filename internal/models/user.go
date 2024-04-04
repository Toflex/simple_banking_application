package models

type User struct {
	Base

	FirstName string
	LastName  string
	Email     string
	Account   Account
}
