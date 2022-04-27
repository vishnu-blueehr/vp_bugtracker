package config

import (
	_ "tawesoft.co.uk/go/dialog"
	// this package is used for dialog

	"zh_bug_tracker/initial"
	// this package is used for DB initial

	_ "database/sql"
	// this package is for DB query Execution

	_ "fmt"
	"net/http"
)

func Config_user_home(w http.ResponseWriter, r *http.Request) {
	// tpl.ExecuteTemplate(w, "admin_home.html", nil)
	initial.Tpl.ExecuteTemplate(w, "user_home.html", nil)
	// io.WriteString(w, "Hello fcc ")
}
