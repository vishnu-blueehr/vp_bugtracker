package project

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

func Project_delete_db(w http.ResponseWriter, r *http.Request) {
	var project_id = r.FormValue("dlproject")

	emp_data := initial.Emp_data{
		Emp_id: initial.Employee_id,
	}
	fmt.Println("\n initial.Employee_id:", initial.Employee_id)
	fmt.Println("\n Emp data:", emp_data)
	fmt.Println("\n logger(the file Project_delete_db-go func is executing..)")

	//db, err := initial.Connect_db()
	//initial.CheckError(err)
	fmt.Printf("\nLogger('Successfully connected to database!')\n")

	query1 := `DELETE FROM project_mates WHERE p_id = $1`
	query2 := `DELETE FROM bug WHERE p_id = $1`
	query3 := `DELETE FROM project WHERE p_id = $1`

	_, err := initial.Db.Exec(query1, project_id)
	initial.CheckError(err)
	_, err = initial.Db.Exec(query2, project_id)
	initial.CheckError(err)
	_, err = initial.Db.Exec(query3, project_id)
	initial.CheckError(err)

	dialog.Alert("Succesfully Deleted the Project : %s", project_id)
	fmt.Printf("\nLogger('Successfully deleted  data from database!')\n")

	switch initial.Employee_id {
	case "1":
		initial.Tpl.ExecuteTemplate(w, "admin_home.html", emp_data)
	default:
		initial.Tpl.ExecuteTemplate(w, "user_home.html", emp_data)
	}

}
