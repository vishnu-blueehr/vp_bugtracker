package usermanagment

import (
	"tawesoft.co.uk/go/dialog"

	// this package is used for dialog

	config "zh_bug_tracker/config"
	"zh_bug_tracker/initial"

	// this package is used for DB initial

	_ "database/sql"
	// this package is for DB query Execution

	"fmt"
	"net/http"
)

func Um_update_employee(w http.ResponseWriter, r *http.Request) {
	var emp_id = r.FormValue("uppemp")

	fmt.Println("emp_id:", emp_id)

	//db, er := initial.Connect_db()
	//initial.CheckError(er)

	query := fmt.Sprintf(`select e.id,e.emp_id,e.emp_name,e.emp_mail from employee e where e.id='%s'`, emp_id)
	// query := fmt.Sprintf(`select e.id,e.emp_id,e.emp_name,e.emp_mail ,l.password from employee e inner join login l on e.emp_id=l.emp_id  where e.id='%s'`, emp_id)
	fmt.Println("query update:", query)

	rows, err := initial.Db.Query(query)
	initial.CheckError(err)
	fmt.Println("rows update:", rows)

	var slc_project []empl_data
	e_data := empl_data{}

	for rows.Next() {
		err := rows.Scan(&e_data.E_id, &e_data.E_empid, &e_data.E_empname, &e_data.E_empmail)
		initial.CheckError(err)

		slc_project = append(slc_project, e_data)
	}
	fmt.Println("slice of data :", slc_project)
	initial.Tpl.ExecuteTemplate(w, "update_user.html", slc_project)

}

func Um_update_data_emp(w http.ResponseWriter, r *http.Request) {

	fmt.Println("\n logger(the file Um_update_data_emp-go func is executing..)")

	r.ParseForm()
	// project_updated := true
	e_id := r.FormValue("table_id")
	e_empid := r.FormValue("emp_id")
	e_name := r.FormValue("emp_name")
	e_mail := r.FormValue("emp_email")
	e_pass := r.FormValue("emp_psw")

	fmt.Println("\n initial.Employee_id:", initial.Employee_id)

	emp_data := initial.Emp_data{
		Emp_id: initial.Employee_id,
	}

	fmt.Println("values:", e_id, e_empid, e_name, e_mail, e_pass)
	//db, err := initial.Connect_db()
	//initial.CheckError(err)

	query1_updatestmt_to_emp_table := fmt.Sprintf(`update employee set emp_name='%s',emp_mail='%s' where emp_id='%s'`, e_name, e_mail, e_empid)
	fmt.Println("\n\n query:", query1_updatestmt_to_emp_table)

	rows, err := initial.Db.Query(query1_updatestmt_to_emp_table)
	fmt.Println("\n\n query after:", rows)
	initial.CheckError(err)

	if e_pass != "" {

		fmt.Println("Password:", e_pass)
		hashed_pass, _ := config.HashPassword(e_pass)
		// hashed_pass, _ := config.HashPassword(e_pass)
		fmt.Println("Hash:    ", hashed_pass)

		query2_updatestmt_to_login_table := fmt.Sprintf(`update login set password='%s' where emp_id='%s'`, hashed_pass, e_empid)
		fmt.Println("\n\n query:", query2_updatestmt_to_login_table)

		rows, err = initial.Db.Query(query2_updatestmt_to_login_table)
		fmt.Println("\n\n query after:", rows)
		initial.CheckError(err)
		fmt.Println("successfully updated user password...")
	}
	dialog.Alert("Succesfully Updated the Employee : %s", e_empid)

	switch initial.Employee_id {
	case "1":
		initial.Tpl.ExecuteTemplate(w, "admin_home.html", emp_data)
	default:
		initial.Tpl.ExecuteTemplate(w, "user_home.html", emp_data)
	}

}
