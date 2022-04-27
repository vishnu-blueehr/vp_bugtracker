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

type empl_data struct {
	E_id      string
	E_empid   string
	E_empname string
	E_empmail string
	E_emppass string
}

func Um_user_managment(w http.ResponseWriter, r *http.Request) {

	//db, er := initial.Connect_db()
	//initial.CheckError(er)

	query := fmt.Sprintf(`select e.id,e.emp_id,e.emp_name,e.emp_mail from employee e order by e.emp_id asc`)
	// query := fmt.Sprintf(`select e.id,e.emp_id,e.emp_name,e.emp_mail ,l.password from employee e inner join login l on e.emp_id=l.emp_id order by e.emp_id asc`)
	fmt.Println("query update:", query)

	rows, err := initial.Db.Query(query)
	initial.CheckError(err)
	defer rows.Close()
	fmt.Println("rows update:", rows)

	var slc_project []empl_data
	e_data := empl_data{}

	for rows.Next() {
		err := rows.Scan(&e_data.E_id, &e_data.E_empid, &e_data.E_empname, &e_data.E_empmail)
		// err := rows.Scan(&e_data.E_id, &e_data.E_empid, &e_data.E_empname, &e_data.E_empmail, &e_data.E_emppass)
		initial.CheckError(err)

		slc_project = append(slc_project, e_data)
	}
	fmt.Println("slice:", slc_project)
	initial.Tpl.ExecuteTemplate(w, "user_managment.html", slc_project)

}
