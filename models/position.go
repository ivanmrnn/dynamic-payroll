package models

import "github.com/uadmin/uadmin"

type Position struct {
	uadmin.Model
	Name      string
	Employee []Employee `gorm:"many2many:employee_roles;"`
}
