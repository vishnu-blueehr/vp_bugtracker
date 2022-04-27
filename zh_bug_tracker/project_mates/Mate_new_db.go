package mates

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

func Mate_new_db(w http.ResponseWriter, r *http.Request) {
	var emp_id string

	fmt.Println("\n initial.Employee_id:", initial.Employee_id)
	fmt.Println("\n logger(the file Mate_new_db-go func is executing..)")

	emp_data := initial.Emp_data{
		Emp_id: initial.Employee_id,
	}
	fmt.Println("\n emp data:", emp_data)

	member_added := true
	p_id := r.FormValue("project_id")
	P_empid := r.FormValue("emp_id")
	//db, err := initial.Connect_db()
	//initial.CheckError(err)

	query := fmt.Sprintf(`SELECT emp_id FROM project_mates where p_id ='%s'`, p_id)
	fmt.Println("query:", query)
	rows, err := initial.Db.Query(query)
	initial.CheckError(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&emp_id)
		initial.CheckError(err)
		if emp_id == P_empid {
			dialog.Alert("Adding new Member to Project failed,Member already exist")
			initial.Tpl.ExecuteTemplate(w, "new_pro_member.html", p_id)
			member_added = false
			break
		}
	}

	if member_added == true {
		insertstmt_to_project_mate := fmt.Sprintf(`insert into "project_mates" values(DEFAULT, '%s', '%s')`, p_id, P_empid)

		_, err = initial.Db.Exec(insertstmt_to_project_mate)

		initial.CheckError(err)
		dialog.Alert("Succesfully Registerd employee :%s to Project : %s", P_empid, p_id)

		switch initial.Employee_id {
		case "1":
			initial.Tpl.ExecuteTemplate(w, "admin_home.html", emp_data)
		default:
			initial.Tpl.ExecuteTemplate(w, "user_home.html", emp_data)
		}
	}

}
