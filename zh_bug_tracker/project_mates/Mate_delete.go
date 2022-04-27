package mates

import (
	"tawesoft.co.uk/go/dialog"
	// this package is used for dialog

	"zh_bug_tracker/initial"
	// this package is used for DB initial

	_ "database/sql"
	// this package is for DB query Execution

	"fmt"
	"net/http"
)

func Mate_delete_db(w http.ResponseWriter, r *http.Request) {

	fmt.Println("\n initial.Employee_id:", initial.Employee_id)
	fmt.Println("\n logger(the file Mate_new_db-go func is executing..)")

	emp_data := initial.Emp_data{
		Emp_id: initial.Employee_id,
	}
	fmt.Println("\n emp data:", emp_data)

	var project_mate_id = r.FormValue("dlpromate")
	//db, err := initial.Connect_db()
	//initial.CheckError(err)
	fmt.Printf("\nLogger('Successfully connected to database!')\n")

	query1 := `DELETE FROM project_mates WHERE id = $1`

	_, err := initial.Db.Exec(query1, project_mate_id)
	initial.CheckError(err)

	dialog.Alert("Succesfully Deleted the Project Member")
	fmt.Printf("\nLogger('Successfully deleted  data from database!')\n")

	switch initial.Employee_id {
	case "1":
		initial.Tpl.ExecuteTemplate(w, "admin_home.html", emp_data)
	default:
		initial.Tpl.ExecuteTemplate(w, "user_home.html", emp_data)
	}
}
