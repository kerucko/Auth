package models

type User struct {
	Id           int
	Email        string
	PasswordHash []byte
}

type App struct {
	Id     int
	Name   string
	Secret string
}