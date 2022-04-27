package bug

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

func Bug_update(w http.ResponseWriter, r *http.Request) {
	var bug_id = r.FormValue("b_id")
	fmt.Println("p_id:", bug_id)

	//db, er := initial.Connect_db()
	//initial.CheckError(er)

	query := fmt.Sprintf(`SELECT bug_id,p_id,assigne,bug_name,summary,reporter,status from bug where bug_id='%s'`, bug_id)
	// query := fmt.Sprintf(`SELECT p_id,p_name,p_key,type,lead FROM project  where p_id='$1'`,
	fmt.Println("query update:", query)

	rows, err := initial.Db.Query(query)
	initial.CheckError(err)
	fmt.Println("rows update:", rows)

	var slc_project []bug_data
	b_data := bug_data{}

	for rows.Next() {
		err = rows.Scan(&b_data.B_id, &b_data.P_name, &b_data.B_ass, &b_data.B_name, &b_data.B_sum, &b_data.B_lead, &b_data.B_sta)
		initial.CheckError(err)

		slc_project = append(slc_project, b_data)
	}
	fmt.Println("slice of data :", slc_project)

	initial.Tpl.ExecuteTemplate(w, "update_bug.html", slc_project)
}
