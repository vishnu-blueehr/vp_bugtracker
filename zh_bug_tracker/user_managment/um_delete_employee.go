package usermanagment

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

func Um_delete_employee(w http.ResponseWriter, r *http.Request) {
	var emp_id = r.FormValue("dlemp")

	emp_data := initial.Emp_data{
		Emp_id: initial.Employee_id,
	}
	fmt.Println("employee id:", initial.Employee_id)
	fmt.Println("emp_data:", emp_data)

	//db, err := initial.Connect_db()
	//initial.CheckError(err)

	fmt.Printf("\nLogger('Successfully connected to database!')\n")

	query1 := `DELETE FROM project_mates WHERE emp_id = $1`
	query2 := `DELETE FROM employee WHERE emp_id = $1`
	// query3 := `DELETE FROM bug WHERE p_id = $1`

	_, err := initial.Db.Exec(query1, emp_id)
	initial.CheckError(err)
	_, err = initial.Db.Exec(query2, emp_id)
	initial.CheckError(err)
	// _, err = initial.Db.Exec(query3, project_id)
	// initial.CheckError(err)

	dialog.Alert("Succesfully Deleted the Employee : %s", emp_id)
	fmt.Printf("\nLogger('Successfully deleted  data from database!')\n")
	initial.Tpl.ExecuteTemplate(w, "admin_home.html", nil)

}
