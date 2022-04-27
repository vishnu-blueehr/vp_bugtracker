package mates

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

func Mates_updation_data(w http.ResponseWriter, r *http.Request) {
	var project_id = r.FormValue("pid")

	fmt.Println("p_id:", project_id)

	//db, er := initial.Connect_db()
	//initial.CheckError(er)

	query := fmt.Sprintf(`select pm.id,p.p_id,p.p_name,pm.emp_id,e.emp_name from project  p inner join project_mates pm on p.p_id=pm.p_id inner join employee e on e.emp_id=pm.emp_id where p.p_id='%s'`, project_id)
	fmt.Println("query update:", query)

	rows, err := initial.Db.Query(query)
	initial.CheckError(err)
	fmt.Println("rows update:", rows)

	var slc_project []project_data
	p_data := project_data{}

	for rows.Next() {
		err = rows.Scan(&p_data.PM_id, &p_data.P_id, &p_data.P_name, &p_data.P_empid, &p_data.P_empname)
		initial.CheckError(err)

		slc_project = append(slc_project, p_data)
	}
	fmt.Println("slice of data :", slc_project)

	initial.Tpl.ExecuteTemplate(w, "update_pro_mates.html", slc_project)

}
