package bug

import (

	// this package is used for dialog

	"zh_bug_tracker/initial"
	// this package is used for DB initial

	"database/sql"
	_ "database/sql"

	// this package is for DB query Execution

	"fmt"
	"net/http"
)

type bug_data struct {
	B_id   string
	B_name string
	P_name string
	B_sum  string
	B_lead string
	B_sta  string
	B_ass  string
}

func Bug_list(w http.ResponseWriter, r *http.Request) {

	// r.ParseForm()
	emp_id := r.FormValue("emp")
	fmt.Printf(emp_id)
	//db, er := initial.Connect_db()
	//initial.CheckError(er)
	var (
		bug_query string
		rows      *sql.Rows
		err       error
	)

	switch emp_id {
	case "1":
		bug_query = fmt.Sprintf(`select bug_id,bug_name,p.p_name,summary,e.emp_name,status,1 from bug b inner join project p on p.p_id=b.p_id inner join employee e on  e.emp_id=b.reporter`)
		rows, err = initial.Db.Query(bug_query)

	default:
		bug_query := fmt.Sprintf(`select bug_id,bug_name,p.p_name,summary,e.emp_name,status,b.assigne from bug b inner join project p on p.p_id=b.p_id inner join employee e on  e.emp_id=b.reporter and ( assigne='%s' or reporter='%s')`, emp_id, emp_id)
		// 	query = `SELECT p_id,p_name,p_key,type,emp_name FROM project p inner join employee e on e.emp_id = '$1' `
		rows, err = initial.Db.Query(bug_query)

	}
	fmt.Println("query:", bug_query)

	fmt.Println("emp_id:", emp_id)

	// } else {
	// 	query = `SELECT p.p_id,p_name,p_key,type,e.emp_name as lead FROM project p inner join employee e on p.lead = e.emp_id inner join project_mates pm on p.p_id = pm.p_id and pm.emp_id =1234`
	// 	// query ="SELECT p.p_id,p_name,p_key,type,e.emp_name as lead FROM project p inner join employee e on p.lead = e.emp_id inner join project_mates pm on p.p_id = pm.p_id and pm.emp_id =1234 "
	// }
	initial.CheckError(err)

	var slc_project []bug_data
	b_data := bug_data{}

	for rows.Next() {
		err = rows.Scan(&b_data.B_id, &b_data.B_name, &b_data.P_name, &b_data.B_sum, &b_data.B_lead, &b_data.B_sta, &b_data.B_ass)
		initial.CheckError(err)

		slc_project = append(slc_project, b_data)
	}
	fmt.Print(slc_project)

	initial.Tpl.ExecuteTemplate(w, "bug_listing.html", slc_project)
}
