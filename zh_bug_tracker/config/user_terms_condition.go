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

func Config_terms_condition(w http.ResponseWriter, r *http.Request) {
	initial.Tpl.ExecuteTemplate(w, "terms&condition.html", nil)
	// io.WriteString(w, "Hello fcc ")
}
