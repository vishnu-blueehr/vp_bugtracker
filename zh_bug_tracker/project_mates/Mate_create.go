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

func Mate_create(w http.ResponseWriter, r *http.Request) {
	var project_id = r.FormValue("pid")

	fmt.Println("p_id:", project_id)

	initial.Tpl.ExecuteTemplate(w, "new_pro_member.html", project_id)

}
