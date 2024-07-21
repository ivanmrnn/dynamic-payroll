package models

import "github.com/uadmin/uadmin"

type Responsibility struct {
	uadmin.Model
	Name      string
	Description string
	ButtonName string
}
