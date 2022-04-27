package config

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

func Config_user_signup(w http.ResponseWriter, r *http.Request) {
	initial.Tpl.ExecuteTemplate(w, "signup.html", nil)
	// io.WriteString(w, "Hello fcc ")
}

func Config_admin_create_user(w http.ResponseWriter, r *http.Request) {
	initial.Tpl.ExecuteTemplate(w, "create_user.html", nil)
}

func Config_admin_user_registration_db(w http.ResponseWriter, r *http.Request) {

	fmt.Println("\n Logger(on page Config_admin_user_registration_db-go function)")

	var (
		user_registered = true
		emp_id          string
	)
	// if r.Method != "POST" {
	// 	http.Redirect(w, r, "/", http.StatusSeeOther)
	// 	return
	// }

	emp_data := initial.Emp_data{
		Emp_id: initial.Employee_id,
	}
	fmt.Println("employee id:", initial.Employee_id)
	fmt.Println("emp_data:", emp_data)

	e_name := r.FormValue("emp_name")
	e_email := r.FormValue("emp_email")
	e_id := r.FormValue("emp_id")
	e_psw := r.FormValue("emp_psw")
	// ename := r.FormValue("emp_email")

	// initial.
	//db, err := initial.Connect_db()
	//initial.CheckError(err)

	rows, err := initial.Db.Query("SELECT emp_id FROM employee")
	initial.CheckError(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&emp_id)
		initial.CheckError(err)
		if emp_id == e_id {
			dialog.Alert("User Registration failed,User already exist")
			initial.Tpl.ExecuteTemplate(w, "create_user.html", nil)
			user_registered = false
			break
		}
	}

	if user_registered == true {

		fmt.Println("Password:", e_psw)
		hashed_pass, _ := HashPassword(e_psw)
		fmt.Println("Hash:    ", hashed_pass)

		insertstmt_to_login := fmt.Sprintf(`insert into "login"("id", "emp_id","password") values(DEFAULT, '%s', '%s')`, e_id, hashed_pass)
		insertstmt_to_emp := fmt.Sprintf(`insert into "employee"("id", "emp_id","emp_name","emp_mail") values(DEFAULT, '%s', '%s', '%s')`, e_id, e_name, e_email)

		_, err = initial.Db.Exec(insertstmt_to_login)
		_, err = initial.Db.Exec(insertstmt_to_emp)

		initial.CheckError(err)
		dialog.Alert("Succesfully Registerd the User : %s", e_id)

		switch initial.Employee_id {
		case "1":
			initial.Tpl.ExecuteTemplate(w, "admin_home.html", emp_data)
		default:
			initial.Tpl.ExecuteTemplate(w, "user_home.html", emp_data)
		}

	}
}

func Config_signup_user_registration_db(w http.ResponseWriter, r *http.Request) {

	fmt.Println("\n Logger(on page Config_signup_user_registration_db-go function)")

	var (
		user_registered = true
		emp_id          string
	)
	// if r.Method != "POST" {
	// 	http.Redirect(w, r, "/", http.StatusSeeOther)
	// 	return
	// }

	e_name := r.FormValue("emp_name")
	e_email := r.FormValue("emp_email")
	e_id := r.FormValue("emp_id")
	e_psw := r.FormValue("emp_psw")
	// ename := r.FormValue("emp_email")

	// initial.
	//db, err := initial.Connect_db()
	//initial.CheckError(err)

	hashed_pass, _ := HashPassword(e_psw)
	fmt.Println("Password:", e_psw)
	fmt.Println("Hash:    ", hashed_pass)

	query := fmt.Sprintf(`SELECT emp_id FROM employee where emp_id='%s'`, e_id)
	fmt.Println("query:   ", query)

	rows, err := initial.Db.Query(query)
	initial.CheckError(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&emp_id)
		initial.CheckError(err)
		if emp_id == e_id {
			dialog.Alert("User Registration failed,User already exist")
			initial.Tpl.ExecuteTemplate(w, "signup.html", nil)
			user_registered = false
			break
		}
	}

	if user_registered == true {
		insertstmt_to_login := fmt.Sprintf(`insert into "login"("id", "emp_id","password") values(DEFAULT, '%s', '%s')`, e_id, hashed_pass)
		_, err = initial.Db.Exec(insertstmt_to_login)
		// insertstmt_to_login := fmt.Sprintf(`insert into "login"("id", "emp_id","password") values(DEFAULT, '%s', '%s')`, e_id, e_psw)
		insertstmt_to_emp := fmt.Sprintf(`insert into "employee"("id", "emp_id","emp_name","emp_mail") values(DEFAULT, '%s', '%s', '%s')`, e_id, e_name, e_email)

		_, err = initial.Db.Exec(insertstmt_to_emp)

		initial.CheckError(err)
		dialog.Alert("Succesfully Registerd the User : %s", e_id)
		initial.Tpl.ExecuteTemplate(w, "home.html", nil)

	}
}
