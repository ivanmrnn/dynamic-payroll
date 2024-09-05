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
		models.Department{},
		models.Employee{},
		models.MenuName{},
		models.Position{},
		models.Responsibility{},
		models.Role{},
		models.RoleResponsibility{},
		models.Team{},
	)

	

	uadmin.RootURL = "/admin/"
	uadmin.SiteName = "Dynamic Payroll"

	createAdminRoleAndUser()

	http.HandleFunc("/logout/", uadmin.Handler(views.LogoutHandler))
	http.HandleFunc("/", uadmin.Handler(views.MainHandler))

	uadmin.StartServer()
}

// createAdminRoleAndUser checks if an admin role and user already exist. If they don't, it creates an entry for the admin at the role and employees database.
func createAdminRoleAndUser() {
    // Create admin role
    var adminRole models.Role
    uadmin.Get(&adminRole, "name = ?", "Admin")
    if adminRole.ID == 0 {
        adminRole = models.Role{
            Name:  "Admin",
            Level: 4,
        }
        uadmin.Save(&adminRole)
        
        // Set responsibilities for the admin role
        models.SetResponsibilitiesForLevel(&adminRole)
    }

    // Create admin user
    var adminUser models.Employee
    uadmin.Get(&adminUser, "id_employee = ?", "admin")
    if adminUser.ID == 0 {
        adminUser = models.Employee{
            IdEmployee: "admin",
            FirstName:  "System",
            LastName:   "Admin",
            Password:   "admin",
            Active:     true,
        }

        // Assign the admin role to the user
        adminUser.Role = append(adminUser.Role, adminRole) 
        adminUser.Save()
    }
}
