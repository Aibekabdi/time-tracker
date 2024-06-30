package models

type User struct {
	ID             uint
	PassportSerie  string
	PassportNumber string
	Surname        string
	Name           string
	Patronymic     string
	Address        string
}
