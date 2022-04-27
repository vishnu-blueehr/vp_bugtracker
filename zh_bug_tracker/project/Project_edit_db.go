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

func Project_edit_db(w http.ResponseWriter, r *http.Request) {

	// project_updated := true
	r.ParseForm()
	p_name := r.FormValue("p_name")
	p_key := r.FormValue("p_key")
	p_type := r.FormValue("p_type")
	p_id := r.FormValue("p_id")
	p_lead := r.FormValue("p_lead")

	fmt.Println("employee id:", initial.Employee_id)

	emp_data := initial.Emp_data{
		Emp_id: initial.Employee_id,
	}
	fmt.Println("emp_data:", emp_data)

	fmt.Println("values:", p_name, p_key, p_type, p_id, p_lead)
	//db, err := initial.Connect_db()
	//initial.CheckError(err)

	// updatestmt_to_project := `update project set p_name='$1',p_key='$2', type='$3',lead=$4 where p_id='$5'`
	updatestmt_to_project := fmt.Sprintf(`update project set p_name='%s',p_key='%s',type='%s',lead='%s' where p_id='%s'`, p_name, p_key, p_type, p_lead, p_id)
	fmt.Println("\n\n query:", updatestmt_to_project)
	rows, err := initial.Db.Query(updatestmt_to_project)

	fmt.Println("\n\n query after:", rows)

	initial.CheckError(err)
	dialog.Alert("Succesfully Updated the Project : %s", p_id)

	switch initial.Employee_id {
	case "1":
		initial.Tpl.ExecuteTemplate(w, "admin_home.html", emp_data)
	default:
		initial.Tpl.ExecuteTemplate(w, "user_home.html", emp_data)
	}

}
