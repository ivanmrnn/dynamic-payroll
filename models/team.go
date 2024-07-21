package models

import "github.com/uadmin/uadmin"

type Team struct {
	uadmin.Model
	Name      string
	Employee []Employee `gorm:"many2many:employee_roles;"`
}
