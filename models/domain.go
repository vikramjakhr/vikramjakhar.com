package models

type User struct {
	Id                string
	FirstName         string
	LastName          string
	UserName          string
	Email             string
	Password          string
	ImageUrl          interface{}
	Mobile            interface{}
	Address           interface{}
	IsBlocked         interface{}
	IsAdmin           interface{}
	IsAccountVerified interface{}
	DateCreated       interface{}
	LastUpdated       interface{}
}
