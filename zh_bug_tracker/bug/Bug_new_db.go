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

func Bug_new_db(w http.ResponseWriter, r *http.Request) {

	var b_name_db string

	fmt.Println("employee id:", initial.Employee_id)

	emp_data := initial.Emp_data{
		Emp_id: initial.Employee_id,
	}
	fmt.Println("emp_data:", emp_data)

	bug_registered := true
	p_id := r.FormValue("project_id")
	emp_id := r.FormValue("emp_id")
	bug_name := r.FormValue("bug_name")
	bug_summary := r.FormValue("summary")
	reporter_id := r.FormValue("reporter")
	bug_status := r.FormValue("status")

	fmt.Println("/nproject_id:", p_id, "/n emp_id: ", emp_id, "/n bug_name: ", bug_name, "/n summary: ", bug_summary, "/n reporter: ", reporter_id, "/n status: ", bug_status)

	// emp_data := emp_data{
	// 	Emp_id:  "emp_id",
	// 	Emp_psw: "123456789",
	// }

	//db, err := initial.Connect_db()
	//initial.CheckError(err)

	rows, err := initial.Db.Query("SELECT bug_name FROM bug")
	initial.CheckError(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&b_name_db)
		initial.CheckError(err)
		if b_name_db == bug_name {
			dialog.Alert("BUG Registration failed,BUG already exist")
			// tpl.ExecuteTemplate(w, "home.html", nil)
			bug_registered = false
			fmt.Println("\n Logger(BUG Registration failed satying in same page)")
			switch initial.Employee_id {
			case "1":
				initial.Tpl.ExecuteTemplate(w, "admin_home.html", emp_data)
			default:
				initial.Tpl.ExecuteTemplate(w, "user_home.html", emp_data)
			}
			// initial.Tpl.ExecuteTemplate(w, "bug_create.html", nil)
			break
		}
	}

	if bug_registered == true {
		insertstmt_to_bug := fmt.Sprintf(`insert into "bug" values(DEFAULT, '%s', '%s', '%s', '%s', '%s','%s')`, p_id, bug_summary, reporter_id, emp_id, bug_status, bug_name)

		fmt.Println("/n query:", insertstmt_to_bug)

		_, err = initial.Db.Exec(insertstmt_to_bug)

		initial.CheckError(err)
		dialog.Alert("Succesfully Registerd the BUG : %s", bug_name)

		switch initial.Employee_id {
		case "1":
			initial.Tpl.ExecuteTemplate(w, "admin_home.html", emp_data)
		default:
			initial.Tpl.ExecuteTemplate(w, "user_home.html", emp_data)
		}
	}

}
