package models

import "github.com/uadmin/uadmin"

type Role struct {
	uadmin.Model
	Name           string
	Employee       []Employee       `gorm:"many2many:employee_roles;"`
	Responsibility []Responsibility `gorm:"many2many:role_responsibility;"`
}
