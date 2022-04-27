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
	// this package is used for Encrtypting password
	// "golang.org/x/crypto/bcrypt"
)

func Config_user_login(w http.ResponseWriter, r *http.Request) {
	initial.Tpl.ExecuteTemplate(w, "login.html", nil)
	// io.WriteString(w, "Hello fcc ")
}

func Config_user_login_validation(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "POST" {
	// http.Redirect(w, r, "/", http.StatusSeeOther)
	// return
	// }

	// To open up a database connectivity at every login
	initial.Connect_db()

	var (
		user_valid      = false
		emp_id, emp_psw string
	)
	e_id := r.FormValue("emp-id")
	initial.Employee_id = r.FormValue("emp-id")

	fmt.Println("\n initial.Employee_id:", initial.Employee_id)

	e_psw := r.FormValue("psw")

	emp_data := initial.Emp_data{
		Emp_id: initial.Employee_id,
	}

	fmt.Println("id:", e_id, "pass", e_psw)

	//db, err := initial.Connect_db()
	//initial.CheckError(err)

	query := fmt.Sprintf(`SELECT emp_id , password FROM login where emp_id='%s'`, e_id)
	fmt.Println("\n select query:", query)

	rows, err := initial.Db.Query(query)
	initial.CheckError(err)

	for rows.Next() {
		err = rows.Scan(&emp_id, &emp_psw)
		match := CheckPasswordHash(e_psw, emp_psw)
		fmt.Println("is match:   ", match)

		// if emp_id == e_id && emp_psw == e_psw {
		if emp_id == e_id && match == true {

			fmt.Print("login success...")
			dialog.Alert("✓ Successfully Logged in as %v user", emp_id)
			user_valid = true

			if emp_id == "1" {
				// initial.Employee_id = e_id
				initial.Tpl.ExecuteTemplate(w, "admin_home.html", emp_data)

			} else {
				// initial.Employee_id = e_id
				initial.Tpl.ExecuteTemplate(w, "user_home.html", emp_data)
				// disp_data(emp_data)
			}

		}
	}
	if user_valid == false {
		fmt.Print("login failed...Invalid Accredentials")
		dialog.Alert("User Login failed,Invalid Accredentials")
		initial.Tpl.ExecuteTemplate(w, "login.html", nil)

	}

}
