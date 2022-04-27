package project

import (
	_ "tawesoft.co.uk/go/dialog"
	// this package is used for dialog

	"zh_bug_tracker/initial"
	// this package is used for DB initial

	_ "database/sql"
	// this package is for DB query Execution

	"fmt"
	"net/http"
)

type project_data struct {
	P_id      string
	P_name    string
	P_key     string
	P_type    string
	P_lead    string
	P_empname string
	P_empid   string
	PM_id     string
}

func Project_retrive_data(w http.ResponseWriter, r *http.Request) {

	// r.ParseForm()
	fmt.Println("\n Logger(on page Project_retrive_data-go function)")
	fmt.Println("\n initial.Employee_id:", initial.Employee_id)
	user := initial.Employee_id

	emp_id := r.FormValue("emp")
	fmt.Printf(emp_id)
	//db, er := initial.Connect_db()
	//initial.CheckError(er)
	var (
		query string
		// rows  *sql.Rows
		err error
		// e_id  int
	)

	//read data from the table
	// if emp_id == "1" {
	switch emp_id {
	case "1":
		query = fmt.Sprintf("select p_id,p_name,p_key,type,emp_name FROM project p left join employee e on p.lead = e.emp_id ")
		// query = fmt.Sprintf("SELECT p_id,p_name,p_key,type,emp_name FROM project p inner join employee e on p.lead = e.emp_id")
		// rows, err = db.Query(query)

	default:
		query = fmt.Sprintf("SELECT p.p_id,p_name,p_key,type,e.emp_name as lead FROM project p left join employee e on p.lead = e.emp_id inner join project_mates pm on p.p_id = pm.p_id and pm.emp_id = '%s'", emp_id)
		// query = fmt.Sprintf("SELECT p_id,p_name,p_key,type,emp_name FROM project p inner join employee e on e.emp_id = '%s'", emp_id)
		// query = `SELECT p_id,p_name,p_key,type,emp_name FROM project p inner join employee e on e.emp_id = '$1' `
		// rows, err = db.Query(query)
	}
	rows, err := initial.Db.Query(query)
	fmt.Println("emp_id:", emp_id)
	fmt.Println("query:", query)

	// } else {
	// 	query = `SELECT p.p_id,p_name,p_key,type,e.emp_name as lead FROM project p inner join employee e on p.lead = e.emp_id inner join project_mates pm on p.p_id = pm.p_id and pm.emp_id =1234`
	// 	// query ="SELECT p.p_id,p_name,p_key,type,e.emp_name as lead FROM project p inner join employee e on p.lead = e.emp_id inner join project_mates pm on p.p_id = pm.p_id and pm.emp_id =1234 "
	// }
	initial.CheckError(err)

	var slc_project []project_data
	p_data := project_data{}

	for rows.Next() {
		err = rows.Scan(&p_data.P_id, &p_data.P_name, &p_data.P_key, &p_data.P_type, &p_data.P_lead)
		initial.CheckError(err)

		slc_project = append(slc_project, p_data)
	}

	fmt.Print(slc_project)

	switch user {
	case "1":
		initial.Tpl.ExecuteTemplate(w, "admin_project_listing.html", slc_project)
	default:
		initial.Tpl.ExecuteTemplate(w, "user_project_listing.html", slc_project)
	}

}
