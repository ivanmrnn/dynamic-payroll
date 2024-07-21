package models

import "github.com/uadmin/uadmin"

type Department struct {
	uadmin.Model
	Name  string
	Teams []Team `gorm:"many2many:department_teams;"`
}
