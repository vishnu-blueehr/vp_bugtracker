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

func Project_updation_data(w http.ResponseWriter, r *http.Request) {
	var project_id = r.FormValue("upproject")
	fmt.Println("p_id:", project_id)

	//db, er := initial.Connect_db()
	//initial.CheckError(er)

	query := fmt.Sprintf(`SELECT p_id,p_name,p_key,type,lead FROM project  where p_id='%s'`, project_id)
	// query := fmt.Sprintf(`SELECT p_id,p_name,p_key,type,lead FROM project  where p_id='$1'`,
	fmt.Println("query update:", query)

	rows, err := initial.Db.Query(query)
	initial.CheckError(err)
	fmt.Println("rows update:", rows)

	var slc_project []project_data
	p_data := project_data{}

	for rows.Next() {
		err = rows.Scan(&p_data.P_id, &p_data.P_name, &p_data.P_key, &p_data.P_type, &p_data.P_lead)
		initial.CheckError(err)

		slc_project = append(slc_project, p_data)
	}
	fmt.Println("slice of data :", slc_project)

	initial.Tpl.ExecuteTemplate(w, "update_project.html", slc_project)
}
