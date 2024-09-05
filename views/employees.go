package views

import (
	"net/http"

	"github.com/ivanmrnn/dynamic-payroll/models"
	"github.com/uadmin/uadmin"
)

func EmployeesHandler(w http.ResponseWriter, r *http.Request, session *uadmin.Session, bc BaseContext) map[string]interface{} {
	employees := []models.Employee{}
	uadmin.All(&employees) // Fetch all employees

	return map[string]interface{}{
		"Title":      "Employees",
		"ActivePage": "employees",
	}
}
