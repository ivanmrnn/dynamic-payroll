package models

import "github.com/uadmin/uadmin"

type Responsibility struct {
    uadmin.Model
    Name       string
    MenuName   MenuName `uadmin:"fk:MenuNameID"`
    MenuNameID uint
}

func (r *Responsibility) Save() {
    uadmin.Save(r)
}
