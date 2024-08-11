package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func LeavesLogsHandler(w http.ResponseWriter, r *http.Request, session *uadmin.Session, bc BaseContext) map[string]interface{} {
	return map[string]interface{}{
		"Title": "Leaves Logs",
		// Add any timesheet-specific data here
	}
}
