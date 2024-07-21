package views

import (
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
	// Base context for all templates
	baseContext := map[string]interface{}{
		"SessionUser":      bc.SessionUser,
		"SessionUserPhoto": bc.SessionUserPhoto,
		"CompanyExists":    bc.CompanyExists,
		"CompanyName":      bc.CompanyName,
		"CompanyLogo":      bc.CompanyLogo,
		"CompanyPrimary":   bc.CompanyPrimary,
		"CompanySecondary": bc.CompanySecondary,
	}

	// Get page-specific context
	var pageContext map[string]interface{}
	switch page {
	case "dashboard":
		pageContext = DashboardHandler(w, r, session, bc)
	default:
		page = "dashboard"
		pageContext = DashboardHandler(w, r, session, bc)
	}

	// Merge page-specific context with base context
	for key, value := range pageContext {
		baseContext[key] = value
	}

	// Render the page with the combined context
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
