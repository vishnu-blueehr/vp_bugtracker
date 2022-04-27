package project

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

func Project_release_accept_db(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\n Logger(on page Project_release_accept_db-go function)")
	fmt.Println("\n initial.Employee_id:", initial.Employee_id)

	var pm_id string
	rel_id := r.FormValue("relproject")
	fmt.Printf(rel_id)

	emp_data := initial.Emp_data{
		Emp_id: initial.Employee_id,
	}

	//db, er := initial.Connect_db()
	//initial.CheckError(er)

	query1 := fmt.Sprintf("select pm_id from user_project_release where id='%s'", rel_id)
	rows, err := initial.Db.Query(query1)
	fmt.Println("query:", rows)
	initial.CheckError(err)

	for rows.Next() {
		err := rows.Scan(&pm_id)
		initial.CheckError(err)
	}

	query := fmt.Sprintf("update user_project_release set status = 'Release granted' where id='%s'", rel_id)
	rows, err = initial.Db.Query(query)
	fmt.Println("query:", rows)
	initial.CheckError(err)

	fmt.Println("query:", query)
	initial.CheckError(err)
	dialog.Alert("Succesfully Updated the Release request : %s", rel_id)

	query2 := fmt.Sprintf(`DELETE FROM user_project_release  WHERE id = '%s'`, rel_id)
	rows, err = initial.Db.Query(query)
	_, err = initial.Db.Exec(query2)
	initial.CheckError(err)

	query3 := fmt.Sprintf(`DELETE FROM project_mates WHERE id = '%s'`, pm_id)
	rows, err = initial.Db.Query(query)
	_, err = initial.Db.Exec(query3)
	initial.CheckError(err)

	dialog.Alert("Succesfully Deleted the Project Member")
	fmt.Printf("\nLogger('Successfully deleted  data from database!')\n")

	switch initial.Employee_id {
	case "1":
		initial.Tpl.ExecuteTemplate(w, "admin_home.html", emp_data)
	default:
		initial.Tpl.ExecuteTemplate(w, "user_home.html", emp_data)
	}
}

func Project_release_decline_db(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\n Logger(on page Project_release_decline_db-go function)")
	fmt.Println("\n initial.Employee_id:", initial.Employee_id)

	rel_id := r.FormValue("relproject")
	fmt.Printf(rel_id)

	emp_data := initial.Emp_data{
		Emp_id: initial.Employee_id,
	}

	//db, er := initial.Connect_db()
	//initial.CheckError(er)

	query := fmt.Sprintf("update user_project_release set status = 'Release Declined' where id='%s'", rel_id)
	rows, err := initial.Db.Query(query)
	fmt.Println("query:", rows)

	fmt.Println("query:", query)
	initial.CheckError(err)
	dialog.Alert("Succesfully Updated the Release request : %s", rel_id)

	switch initial.Employee_id {
	case "1":
		initial.Tpl.ExecuteTemplate(w, "admin_home.html", emp_data)
	default:
		initial.Tpl.ExecuteTemplate(w, "user_home.html", emp_data)
	}
}

func Project_release_delete_db(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\n Logger(on page Project_retrive_data-go function)")
	fmt.Println("\n initial.Employee_id:", initial.Employee_id)

	rel_id := r.FormValue("relproject")
	fmt.Printf(rel_id)

	emp_data := initial.Emp_data{
		Emp_id: initial.Employee_id,
	}

	//db, er := initial.Connect_db()
	//initial.CheckError(er)

	query := fmt.Sprintf("DELETE FROM user_project_release WHERE id ='%s'", rel_id)
	rows, err := initial.Db.Query(query)
	fmt.Println("query:", rows)

	fmt.Println("query:", query)
	initial.CheckError(err)
	dialog.Alert("Succesfully Deleted the Release request : %s", rel_id)

	switch initial.Employee_id {
	case "1":
		initial.Tpl.ExecuteTemplate(w, "admin_home.html", emp_data)
	default:
		initial.Tpl.ExecuteTemplate(w, "user_home.html", emp_data)
	}

}
