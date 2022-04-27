package config

import (
	"fmt"

	"tawesoft.co.uk/go/dialog"
	// this package is used for dialog

	"zh_bug_tracker/initial"
	// this package is used for DB initial

	_ "database/sql"
	// this package is for DB query Execution

	_ "fmt"
	"net/http"
)

func Config_user_logout(w http.ResponseWriter, r *http.Request) {
	initial.Employee_id = "0"

	// To close the database after every logout
	defer initial.Db.Close()
	fmt.Println("The database connenctivity closed...!")

	dialog.Alert(" Successfully Logged Out.")
	initial.Tpl.ExecuteTemplate(w, "home.html", nil)
}
