package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string
	AddressID   uint
	Address     Address
	PhoneNumber string
	Email       string
}

type Address struct {
	ID         uint
	StreetNo   int
	StreetName string
	City       string
	Country    string
}
