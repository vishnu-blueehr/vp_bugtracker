package bug

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

func Bug_delete(w http.ResponseWriter, r *http.Request) {

	var bug_id = r.FormValue("b_id")
	//db, err := initial.Connect_db()
	//initial.CheckError(err)

	fmt.Println("\n initial.Employee_id:", initial.Employee_id)
	fmt.Println("\n logger(the file Bug_delete-go func is executing..)")

	emp_data := initial.Emp_data{
		Emp_id: initial.Employee_id,
	}

	fmt.Printf("\nLogger('Successfully connected to database!')\n")

	query1 := `DELETE FROM bug WHERE bug_id = $1`

	_, err := initial.Db.Exec(query1, bug_id)
	initial.CheckError(err)

	dialog.Alert("Succesfully Deleted the Bug : %s", bug_id)
	fmt.Printf("\nLogger('Successfully deleted  data from database!')\n")

	switch initial.Employee_id {
	case "1":
		initial.Tpl.ExecuteTemplate(w, "admin_home.html", emp_data)
	default:
		initial.Tpl.ExecuteTemplate(w, "user_home.html", emp_data)
	}
}
