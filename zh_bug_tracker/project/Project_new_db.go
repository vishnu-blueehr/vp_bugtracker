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

func Project_new_db(w http.ResponseWriter, r *http.Request) {
	var emp_id string
	project_registered := true
	p_name := r.FormValue("p_name")
	p_key := r.FormValue("p_key")
	p_type := r.FormValue("p_type")
	p_id := r.FormValue("p_id")
	p_lead := r.FormValue("p_lead")

	fmt.Println("\n initial.Employee_id:", initial.Employee_id)
	fmt.Println("\n logger(the file Project_new_db-go func is executing..)")

	emp_data := initial.Emp_data{
		Emp_id: initial.Employee_id,
	}

	//db, err := initial.Connect_db()
	//initial.CheckError(err)

	rows, err := initial.Db.Query("SELECT p_id FROM project")
	initial.CheckError(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&emp_id)
		initial.CheckError(err)
		if emp_id == p_id {
			dialog.Alert("Project Registration failed,Project_id already exist")
			// tpl.ExecuteTemplate(w, "home.html", nil)
			project_registered = false
			fmt.Println("\n Logger(Project Registration failed satying in same page)")
			initial.Tpl.ExecuteTemplate(w, "reg_project.html", nil)
			break
		}
	}

	if project_registered == true {
		insertstmt_to_project := fmt.Sprintf(`insert into "project" values(DEFAULT, '%s', '%s', '%s', '%s', '%s')`, p_name, p_key, p_type, p_lead, p_id)
		insertstmt_to_project_mates := fmt.Sprintf(`insert into "project_mates" values(DEFAULT, '%s', '%s')`, p_id, p_lead)

		_, err = initial.Db.Exec(insertstmt_to_project)
		_, err = initial.Db.Exec(insertstmt_to_project_mates)

		initial.CheckError(err)
		dialog.Alert("Succesfully Registerd the Project : %s", p_id)

		switch initial.Employee_id {
		case "1":
			initial.Tpl.ExecuteTemplate(w, "admin_home.html", emp_data)
		default:
			initial.Tpl.ExecuteTemplate(w, "user_home.html", emp_data)
		}

	}

}
