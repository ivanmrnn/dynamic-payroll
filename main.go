package main

import (
	"net/http"

	"github.com/ivanmrnn/dynamic-payroll/models"
	"github.com/ivanmrnn/dynamic-payroll/views"
	"github.com/uadmin/uadmin"
)

func main() {
	uadmin.Register(
		models.Company{},
		models.Employee{},
		models.Role{},
		models.Position{},
		models.Department{},
		models.Team{},
		models.Responsibility{},
	)

	uadmin.RegisterInlines(
		models.Role{},
		map[string]string{
			"Employee": "RoleID",
		},
	)

	uadmin.RootURL = "/admin/"
	uadmin.SiteName = "Dynamic Payroll"

	http.HandleFunc("/logout/", uadmin.Handler(views.LogoutHandler))
	http.HandleFunc("/", uadmin.Handler(views.MainHandler))

	uadmin.StartServer()
}
