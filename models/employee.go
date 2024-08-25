package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type Employee struct {
	uadmin.Model
	Name    string `uadmin:"read_only"`
	UserID  uint   `uadmin:"read_only"`

	IdEmployee string
	FirstName  string
	MiddleName string
	LastName   string
	SuffixName string

	Password string `uadmin:"password;list_exclude"`

	Photo string `uadmin:"image;list_exclude"`

	DateOfBirth time.Time
	Email       string

	SSS        string
	Tin        string
	PhilHealth string
	PagIbig    string

	Role       []Role `gorm:"many2many:employee_roles;"`

	Position         []Position       `gorm:"many2many:employee_positions;"`
	Department       []Department     `gorm:"many2many:employee_departments;"`
	Team             []Team           `gorm:"many2many:employee_teams;"`

	Active bool `uadmin:"list_exclude"`
}

// Save method for the Employee model
// Save method for the Employee model
func (e *Employee) Save() {
	e.Name = e.FirstName + " " + e.LastName

	// Deactivate any previous Employee record with the same IdEmployee
	previousEmployees := []Employee{}
	uadmin.Filter(&previousEmployees, "id_employee = ? AND active = ?", e.IdEmployee, true)
	for _, emp := range previousEmployees {
		emp.Active = false
		uadmin.Save(&emp)
	}

	var user uadmin.User
	uadmin.Get(&user, "username = ?", e.IdEmployee)

	if user.ID == 0 {
		// User does not exist, create a new one
		user = uadmin.User{
			Username:  e.IdEmployee,
			FirstName: e.FirstName,
			LastName:  e.LastName,
			Password:  e.Password,
			Email:     e.Email,
			Active:    true,
			Photo: e.Photo,
		}
	} else {
		// User exists, update their details
		user.FirstName = e.FirstName
		user.LastName = e.LastName
		user.Password = e.Password
		user.Email = e.Email
		user.Active = true
		user.Photo = e.Photo
	}

	// Set admin privileges for the system admin
	if e.IdEmployee == "admin" {
		user.Admin = true
	}

	user.Save()

	// Set the UserID in Employee to the ID of the user
	e.UserID = user.ID

	// Mark the current Employee record as active
	e.Active = true

	uadmin.Save(e)
}

