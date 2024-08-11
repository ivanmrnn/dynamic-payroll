package models

import "github.com/uadmin/uadmin"

type MenuName struct {
	uadmin.Model
	Name     string
	DisplayName string
	MenuIcon string
}

func (m *MenuName) Save() {
	uadmin.Save(m)
}
