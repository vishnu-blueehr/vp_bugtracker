package config

import (

	// this package is used for dialog

	"zh_bug_tracker/initial"
	// this package is used for DB initial

	_ "database/sql"
	// this package is for DB query Execution

	_ "fmt"
	"net/http"
)

func Config_home(w http.ResponseWriter, r *http.Request) {
	initial.Tpl.ExecuteTemplate(w, "home.html", nil)
	// io.WriteString(w, "Hello fcc ")
}

func Config_home_redirect(w http.ResponseWriter, r *http.Request) {
	// emp_id := r.FormValue("emp")

	emp_data := initial.Emp_data{
		Emp_id: initial.Employee_id,
	}

	switch initial.Employee_id {
	case "1":
		initial.Tpl.ExecuteTemplate(w, "admin_home.html", emp_data)
	default:
		initial.Tpl.ExecuteTemplate(w, "user_home.html", emp_data)
	}

}
