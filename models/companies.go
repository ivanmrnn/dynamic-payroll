package models

import (
	"github.com/uadmin/uadmin"
)

type Companies struct {
	uadmin.Model
	Name string
	Logo string `uadmin:"image"`
	Primary string
	Secondary string
	Active bool
}

func (c *Companies) Save(){
	if c.Active {
		companies := []Companies{}
		uadmin.Filter(&companies, "active = ?", true)
		for _, company := range companies {
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