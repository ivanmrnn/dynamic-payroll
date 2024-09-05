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

	// If the`active` of the company being saved is true, set all other companies as inactive.
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

	//makes sure that `active` is set to false when it is set to false  
	if !c.Active {
		c.Active = false
	}
	
	uadmin.Save(c)
}