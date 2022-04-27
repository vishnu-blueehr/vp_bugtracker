package usermanagment

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

func Um_profile_manange(w http.ResponseWriter, r *http.Request) {
	emp_id := r.FormValue("emp")
	fmt.Printf(emp_id)
	//db, er := initial.Connect_db()
	//initial.CheckError(er)

	query := fmt.Sprintf(`select e.id,e.emp_id,e.emp_name,e.emp_mail ,l.password from employee e inner join login l on e.emp_id=l.emp_id  and e.emp_id='%s'`, emp_id)
	fmt.Println("query update:", query)

	rows, err := initial.Db.Query(query)
	initial.CheckError(err)
	defer rows.Close()
	fmt.Println("rows update:", rows)

	var slc_project []empl_data
	e_data := empl_data{}

	for rows.Next() {
		err := rows.Scan(&e_data.E_id, &e_data.E_empid, &e_data.E_empname, &e_data.E_empmail, &e_data.E_emppass)
		initial.CheckError(err)

		slc_project = append(slc_project, e_data)
	}

	fmt.Println("slice:", slc_project)
	initial.Tpl.ExecuteTemplate(w, "update_user.html", slc_project)

}
