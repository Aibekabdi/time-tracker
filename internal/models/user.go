package models

type User struct {
	PassportSerie  uint
	PassportNumber uint
	Surname        string
	Name           string
	Patronymic     string
	Address        string
}
