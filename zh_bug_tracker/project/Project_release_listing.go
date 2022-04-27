package project

import (
	_ "tawesoft.co.uk/go/dialog"
	// this package is used for dialog

	"zh_bug_tracker/initial"
	// this package is used for DB initial

	"database/sql"
	// this package is for DB query Execution

	"fmt"
	"net/http"
)

func Project_release_listing(w http.ResponseWriter, r *http.Request) {
	emp_id := r.FormValue("emp")
	fmt.Printf(emp_id)
	//db, er := initial.Connect_db()
	//initial.CheckError(er)
	var (
		query string
		rows  *sql.Rows
		err   error
		// e_id  int
	)

	//read data from the table
	// if emp_id == "1" {
	switch emp_id {
	case "1":
		query = fmt.Sprintf("select r.id,p.p_id,p.p_name,r.emp_id,emp.emp_name,r.status,r.reason  from user_project_release r left join project p on p.p_id=r.p_id left join employee emp on emp.emp_id=r.emp_id")
		// rows, err = db.Query(query)

	default:
		query = fmt.Sprintf("select r.id,p.p_id,p.p_name,r.emp_id,emp.emp_name,r.status,r.reason  from user_project_release r left join project p on p.p_id=r.p_id left join employee emp on emp.emp_id=r.emp_id where r.emp_id='%s'", emp_id)
	}
	rows, err = initial.Db.Query(query)
	fmt.Println("emp_id:", emp_id)
	fmt.Println("query:", query)

	initial.CheckError(err)

	var slc_project []project_data
	p_data := project_data{}

	for rows.Next() {
		err = rows.Scan(&p_data.PM_id, &p_data.P_id, &p_data.P_name, &p_data.P_empid, &p_data.P_empname, &p_data.P_key, &p_data.P_type)
		initial.CheckError(err)

		slc_project = append(slc_project, p_data)
	}

	fmt.Print(slc_project)

	switch emp_id {
	case "1":
		initial.Tpl.ExecuteTemplate(w, "admin_release_listing.html", slc_project)
	default:
		initial.Tpl.ExecuteTemplate(w, "user_release_listing.html", slc_project)
	}

}
