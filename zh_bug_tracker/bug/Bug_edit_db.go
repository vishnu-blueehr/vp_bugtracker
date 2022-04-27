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

func Bug_edit(w http.ResponseWriter, r *http.Request) {

	fmt.Println("\n Logger(on page Bug_edit-go function)")

	// project_updated := true
	r.ParseForm()
	// bug_registered := true
	bug_id := r.FormValue("b_id")
	p_id := r.FormValue("project_id")
	emp_id := r.FormValue("emp_id")
	bug_name := r.FormValue("bug_name")
	bug_summary := r.FormValue("summary")
	reporter_id := r.FormValue("reporter")
	bug_status := r.FormValue("status")

	emp_data := initial.Emp_data{
		Emp_id: initial.Employee_id,
	}
	fmt.Println("employee id:", initial.Employee_id)
	fmt.Println("emp_data:", emp_data)

	fmt.Println("value of bug_id:", bug_id)
	fmt.Println("values:", bug_id, p_id, emp_id, bug_name, bug_summary, reporter_id, bug_status, bug_id)
	//db, err := initial.Connect_db()
	//initial.CheckError(err)

	// rows, err := db.Query("SELECT p_id FROM project")
	// initial.CheckError(err)
	// defer rows.Close()

	// updatestmt_to_project := `update project set p_name='$1',p_key='$2', type='$3',lead=$4 where p_id='$5'`
	updatestmt_to_bug := fmt.Sprintf(`update bug set p_id='%s',summary='%s',reporter='%s',assigne='%s',status='%s', bug_name='%s' where bug_id='%s'`, p_id, bug_summary, reporter_id, emp_id, bug_status, bug_name, bug_id)
	fmt.Println("\n\n query:", updatestmt_to_bug)
	// _, err = initial.Db.Exec(updatestmt_to_project)

	// rows, err := db.Query(updatestmt_to_project, p_name, p_key, p_type, p_lead, p_id)
	rows, err := initial.Db.Query(updatestmt_to_bug)

	fmt.Println("\n\n query after:", rows)

	initial.CheckError(err)
	dialog.Alert("Succesfully Updated the BUG : %s", bug_id)

	switch initial.Employee_id {
	case "1":
		initial.Tpl.ExecuteTemplate(w, "admin_home.html", emp_data)
	default:
		initial.Tpl.ExecuteTemplate(w, "user_home.html", emp_data)
	}

}
