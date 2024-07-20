package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

func LoginHandler(w http.ResponseWriter, r *http.Request, baseContext BaseContext) {
	type Context struct {
		BaseContext
		Err       string
		ErrExists bool
		Password  string
	}

	c := Context{
		BaseContext: baseContext,
	}

	// If the request method is POST
	if r.Method == "POST" {
		//Login request from the user.
		username := r.PostFormValue("username")
		username = strings.TrimSpace(strings.ToLower(username))
		password := r.PostFormValue("password")

		// Login using username and password.
		session, _ := uadmin.Login(r, username, password)

		// Check whether the session returned is nil or the user is not active.
		if session == nil || !session.User.Active {
			// Assign the login validation here that will be used for UI displaying.
			c.ErrExists = true
			c.Err = "Invalid username or password"
		} else {
			// Create a session cookie
			cookie, _ := r.Cookie("session")
			if cookie == nil {
				cookie = &http.Cookie{}
			}
			cookie.Name = "session"
			cookie.Value = session.Key
			cookie.Path = "/"
			cookie.SameSite = http.SameSiteStrictMode
			http.SetCookie(w, cookie)

			// If valid, proceed to dashboard
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			return
		}
	}

	// Render the login filepath and pass the context data object to the HTML file.
	uadmin.RenderHTML(w, r, "templates/login.html", c)
}
