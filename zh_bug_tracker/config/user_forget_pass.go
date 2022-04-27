package config

import (
	"tawesoft.co.uk/go/dialog"
	// this package is used for dialog

	"zh_bug_tracker/initial"
	// this package is used for DB initial

	_ "database/sql"
	// this package is for DB query Execution

	_ "fmt"
	"net/http"
)

func Config_forget_pass(w http.ResponseWriter, r *http.Request) {

	dialog.Alert("Please contact System Admin at systemadmin@blueehr.com for Accredentials")
	initial.Tpl.ExecuteTemplate(w, "home.html", nil)
}
