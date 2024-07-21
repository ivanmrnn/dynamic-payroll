package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type Employee struct {
	uadmin.Model
	UserID uint
	User   uadmin.User

	IdEmployee string
	FirstName  string
	MiddleName string
	LastName   string
	SuffixName string

	Password string `uadmin:"password"`

	DateOfBirth time.Time
	Email       string

	SSS        string
	Tin        string
	PhilHealth string
	PagIbig    string

	Role             Role
	RoleID           uint

	Position         []Position       `gorm:"many2many:employee_positions;"`
	Department       []Department     `gorm:"many2many:employee_departments;"`
	Team             []Team           `gorm:"many2many:employee_teams;"`
	Responsibilities []Responsibility `gorm:"many2many:employee_responsibilities;"`
}

func (e *Employee) Save() {
	user := uadmin.User{}
	user.Username = e.IdEmployee
	user.FirstName = e.FirstName
	user.LastName = e.LastName
	user.Password = "password"
	user.Email = e.Email

	user.Active = true

	user.Save()
	uadmin.Save(e)
}
