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

func Project_release_data_user(w http.ResponseWriter, r *http.Request) {

	fmt.Println("\n initial.Employee_id:", initial.Employee_id)
	fmt.Println("\n logger(the file Project_release_data_user-go func is executing..)")

	p_id := r.FormValue("relproject")
	emp_id := initial.Employee_id

	//db, er := initial.Connect_db()
	//initial.CheckError(er)

	query := fmt.Sprintf(`select pm.id,p.p_id,p_name,p_key,type,pm.emp_id ,lead,emp.emp_name FROM project p inner join project_mates pm on p.p_id=pm.p_id left join employee emp on emp.emp_id=lead where  pm.emp_id='%s'  and pm.p_id='%s'`, emp_id, p_id)
	fmt.Println("query update:", query)

	rows, err := initial.Db.Query(query)
	initial.CheckError(err)
	fmt.Println("rows update:", rows)

	var slc_project []project_data
	p_data := project_data{}

	for rows.Next() {
		err = rows.Scan(&p_data.PM_id, &p_data.P_id, &p_data.P_name, &p_data.P_key, &p_data.P_type, &p_data.P_empid, &p_data.P_lead, &p_data.P_empname)
		initial.CheckError(err)

		slc_project = append(slc_project, p_data)
	}
	fmt.Println("slice of data :", slc_project)

	initial.Tpl.ExecuteTemplate(w, "Project_user_release_form.html", slc_project)

}

func Project_release_data_db(w http.ResponseWriter, r *http.Request) {

	fmt.Println("\n initial.Employee_id:", initial.Employee_id)
	fmt.Println("\n logger(the file Project_release_data_db-go func is executing..)")

	emp_data := initial.Emp_data{
		Emp_id: initial.Employee_id,
	}
	fmt.Println("emp_data:", emp_data)

	pm_id := r.FormValue("pm_id")
	p_empid := r.FormValue("emp_id")
	p_name := r.FormValue("p_name")
	p_key := r.FormValue("p_key")
	p_type := r.FormValue("p_type")
	p_id := r.FormValue("p_id")
	p_lead := r.FormValue("p_empname")
	reason := r.FormValue("reason")
	status := r.FormValue("status")

	var (
		pr_id, em_id, prm_id string
		release_registered   = true
	)
	fmt.Printf("\n values:\n pm_id:%s ,emp_id:%s , pr_name:%s , pr_id:%s , reason :%s,status:%s,%s,%s,%s", pm_id, p_empid, p_name, p_id, reason, status, p_key, p_type, p_lead)

	//db, err := initial.Connect_db()
	//initial.CheckError(err)

	query := fmt.Sprintf(`SELECT p_id,emp_id,pm_id FROM user_project_release`)
	rows, err := initial.Db.Query(query)
	initial.CheckError(err)
	defer rows.Close()

	if rows.Next() {
		for rows.Next() {
			err := rows.Scan(&pr_id, &em_id, &prm_id)
			initial.CheckError(err)
			if p_id == pr_id && em_id == p_empid && prm_id == pm_id {
				dialog.Alert("Release Request Registration failed,Request already exist,Please wait for request reply")
				// tpl.ExecuteTemplate(w, "home.html", nil)
				release_registered = false
				fmt.Println("\n Logger(Release request Registration failed satying in same page)")
				initial.Tpl.ExecuteTemplate(w, "reg_project.html", nil)
				break
			}
		}
	}

	if release_registered == true {
		insert_query_to_release := fmt.Sprintf(`insert into user_project_release values(default,'%s','%s','%s','%s','%s')`, p_id, p_empid, pm_id, status, reason)
		fmt.Println("\n query:", query)

		_, err = initial.Db.Exec(insert_query_to_release)
		initial.CheckError(err)
		dialog.Alert("Succesfully Registerd the Request for Project release ,please wait for release confirmation from higher authority")

		switch initial.Employee_id {
		case "1":
			initial.Tpl.ExecuteTemplate(w, "admin_home.html", emp_data)
		default:
			initial.Tpl.ExecuteTemplate(w, "user_home.html", emp_data)
		}
	}

}
