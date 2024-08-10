package views

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/ivanmrnn/dynamic-payroll/models"
	"github.com/uadmin/uadmin"
)

type BaseContext struct {
	SessionUser      string
	SessionUserPhoto string `uadmin:"image"`
	CompanyExists    bool
	CompanyName      string
	CompanyLogo      string `uadmin:"image"`
	CompanyPrimary   string
	CompanySecondary string
	MenuItems        []MenuItem
	IdEmployee string
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")
	page := strings.TrimSuffix(r.URL.Path, "/")

	session := uadmin.IsAuthenticated(r)

	if session != nil {
		// If user is authenticated and trying to access login or root, redirect to dashboard
		if page == "" || page == "login" {
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			return
		}
	} else {
		// If user is not authenticated and trying to access any page other than login, redirect to login
		if page != "login" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
	}

	// Create a base context
	bc := BaseContext{}
	if session != nil {
		bc.SessionUser = session.User.FirstName + " " + session.User.LastName
		if session.User.Photo != "" {
			bc.SessionUserPhoto = session.User.Photo
		}

		var employee models.Employee
        err := uadmin.Get(&employee, "user_id = ?", session.UserID)
        if err == nil {
            bc.IdEmployee = employee.IdEmployee
        }
		
		
		// Generate menu items based on user roles and responsibilities
		bc.MenuItems = GeneratePayrollMenu(session.UserID)

		UpdateResponsibilitiesMenuNames()
		CheckResponsibilitiesAndMenuNames()
	}

	activeCompany := getActiveCompany()

	// Populate the context with active company details
	if activeCompany != nil {
		bc.CompanyExists = true
		bc.CompanyLogo = activeCompany.Logo
		bc.CompanyName = activeCompany.Name
		bc.CompanyPrimary = activeCompany.Primary
		bc.CompanySecondary = activeCompany.Secondary
	}

	// Determine the page to render
	switch page {
	case "login":
		if session != nil {
			// If already authenticated, redirect to dashboard
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			return
		}
		LoginHandler(w, r, bc)
	default:
		renderPage(w, r, page, session, bc)
	}
}

func renderPage(w http.ResponseWriter, r *http.Request, page string, session *uadmin.Session, bc BaseContext) {
	baseContext := map[string]interface{}{
		"SessionUser":      bc.SessionUser,
		"SessionUserPhoto": bc.SessionUserPhoto,
		"CompanyExists":    bc.CompanyExists,
		"CompanyName":      bc.CompanyName,
		"CompanyLogo":      bc.CompanyLogo,
		"CompanyPrimary":   bc.CompanyPrimary,
		"CompanySecondary": bc.CompanySecondary,
		"MenuItems":        bc.MenuItems,
		"IdEmployee":       bc.IdEmployee,
	}

	var pageContext map[string]interface{}
	switch page {
	case "dashboard":
		pageContext = DashboardHandler(w, r, session, bc)
	default:
		page = "dashboard"
		pageContext = DashboardHandler(w, r, session, bc)
	}

	for key, value := range pageContext {
		baseContext[key] = value
	}

	renderTemplates(w, r, page, baseContext)
}

func renderTemplates(w http.ResponseWriter, r *http.Request, page string, context map[string]interface{}) {
	// List of templates to render
	templateList := []string{"./templates/base.html"}
	pagePath := "./templates/" + page + ".html"
	templateList = append(templateList, pagePath)

	// Render the templates with the provided context
	uadmin.RenderMultiHTML(w, r, templateList, context)
}

func getActiveCompany() *models.Company {
	companies := []models.Company{}

	err := uadmin.Filter(&companies, "active = ?", true)
	if err != nil {
		return nil
	}

	if len(companies) > 0 {
		return &companies[0]
	}

	return nil
}

type MenuItem struct {
	MenuName string
	MenuIcon template.HTML // Change to template.HTML
}

// Add this at the package level
var menuOrder = []string{
    "Dashboard",
    "Timesheet",
    "Payroll",
    "Leaves",
	"Approvals",
    "Department",
    "Employees",
    "Roles",
}

func GeneratePayrollMenu(userID uint) []MenuItem {
    var employee models.Employee
    err := uadmin.Get(&employee, "user_id = ?", userID)
    if err != nil || employee.ID == 0 {
        return []MenuItem{}
    }

    menuItems := make(map[string]MenuItem)

    for _, role := range employee.Role {
        var roleResponsibilities []models.RoleResponsibility
        uadmin.Filter(&roleResponsibilities, "role_id = ?", role.ID)

        for _, roleResp := range roleResponsibilities {
            var responsibility models.Responsibility
            err := uadmin.Get(&responsibility, "id = ?", roleResp.ResponsibilityID)
            if err != nil {
                continue
            }

            var menuName models.MenuName
            err = uadmin.Get(&menuName, "id = ?", responsibility.MenuNameID)
            if err != nil {
                continue
            }

            menuIcon := template.HTML(UnescapeSVG(menuName.MenuIcon))

            menuItems[menuName.Name] = MenuItem{
                MenuName: menuName.Name,
                MenuIcon: menuIcon,
            }
        }
    }

    // Sort menu items according to the predefined order
    result := make([]MenuItem, 0, len(menuItems))
    for _, menuName := range menuOrder {
        if item, exists := menuItems[menuName]; exists {
            result = append(result, item)
        }
    }

    return result
}
func CheckResponsibilitiesAndMenuNames() {
	var responsibilities []models.Responsibility
	uadmin.All(&responsibilities)

	for _, resp := range responsibilities {
		var menuName models.MenuName
		err := uadmin.Get(&menuName, resp.MenuNameID)
		if err != nil {
			continue
		}
	}
}

func UpdateResponsibilitiesMenuNames() {
	var responsibilities []models.Responsibility
	uadmin.All(&responsibilities)

	for _, resp := range responsibilities {
		var menuName models.MenuName
		err := uadmin.Get(&menuName, resp.MenuNameID)
		if err == nil && menuName.ID != 0 {
			resp.MenuName = menuName
			uadmin.Save(&resp)
		}
	}
}

func UnescapeSVG(svg string) string {
	svg = strings.ReplaceAll(svg, `\u003c`, `<`)
	svg = strings.ReplaceAll(svg, `\u003e`, `>`)
	svg = strings.ReplaceAll(svg, `&amp;`, `&`)
	svg = strings.ReplaceAll(svg, `&quot;`, `"`)
	svg = strings.ReplaceAll(svg, `&#x27;`, `'`)
	return svg
}
