package models

type User struct {
	ID             uint
	PassportSerie  uint
	PassportNumber uint
	Surname        string
	Name           string
	Patronymic     string
	Address        string
}
