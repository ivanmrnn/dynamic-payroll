package models

import (
	"github.com/uadmin/uadmin"
)

type Company struct {
	uadmin.Model
	Name string
	Logo string `uadmin:"image"`
	Primary string
	Secondary string
	Active bool
}

func (c *Company) Save(){
	if c.Active {
		company := []Company{}
		uadmin.Filter(&company, "active = ?", true)
		for _, company := range company {
			if company.ID != c.ID {
				company.Active = false
				uadmin.Save(&company)
			}
		}
	}

	if !c.Active {
		c.Active = false
	}
	
	uadmin.Save(c)
}