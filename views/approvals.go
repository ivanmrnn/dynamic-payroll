package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func ApprovalsHandler(w http.ResponseWriter, r *http.Request, session *uadmin.Session, bc BaseContext) map[string]interface{} {
	return map[string]interface{}{
		"Title": "Approvals",
		"ActivePage": "approvals",
	}
}
