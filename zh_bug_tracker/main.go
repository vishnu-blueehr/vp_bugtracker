package main

import (
	_ "database/sql"
	_ "fmt"
	_ "html/template"
	_ "io"
	"net/http"

	_ "tawesoft.co.uk/go/dialog"
	// this package is used for dialog

	_ "zh_bug_tracker/initial"

	bug "zh_bug_tracker/bug"

	project "zh_bug_tracker/project"

	mates "zh_bug_tracker/project_mates"

	config "zh_bug_tracker/config"

	um "zh_bug_tracker/user_managment"
)

func main() {
	file := http.FileServer(http.Dir("src"))
	http.Handle("/src/", http.StripPrefix("/src", file))

	//Handle function for  "Project"...
	http.HandleFunc("/project_list/", project.Project_retrive_data)
	http.HandleFunc("/new_project", project.Project_new_db)
	http.HandleFunc("/create_project", project.Project_create)
	http.HandleFunc("/updateproject/", project.Project_updation_data)
	http.HandleFunc("/edit_project", project.Project_edit_db)
	http.HandleFunc("/delete_project/", project.Project_delete_db)
	http.HandleFunc("/request_listing/", project.Project_release_listing)
	http.HandleFunc("/releaseproject/", project.Project_release_data_user)
	http.HandleFunc("/project_release_form", project.Project_release_data_db)
	http.HandleFunc("/delete_project_release/", project.Project_release_delete_db)
	http.HandleFunc("/accept_project_release/", project.Project_release_accept_db)
	http.HandleFunc("/decline_project_release/", project.Project_release_decline_db)

	//Handle function for  "Project Mates"...
	http.HandleFunc("/add_pro_member/", mates.Mate_create)
	http.HandleFunc("/new_project_mate", mates.Mate_new_db)
	http.HandleFunc("/update_mates/", mates.Mates_updation_data)
	http.HandleFunc("/delete_project_mate/", mates.Mate_delete_db)

	//Handle function for  "Bug"...
	http.HandleFunc("/create_bug/", bug.Bug_create)
	http.HandleFunc("/bug_list/", bug.Bug_list)
	http.HandleFunc("/new_bug", bug.Bug_new_db)
	http.HandleFunc("/update_bug/", bug.Bug_update)
	http.HandleFunc("/edit_bug", bug.Bug_edit)
	http.HandleFunc("/delete_bug/", bug.Bug_delete)

	//Handle function for  "Config"...
	http.HandleFunc("/", config.Config_home)
	http.HandleFunc("/login", config.Config_user_login)
	http.HandleFunc("/logout", config.Config_user_logout)
	http.HandleFunc("/user_login", config.Config_user_login_validation)
	http.HandleFunc("/terms_condition", config.Config_terms_condition)
	http.HandleFunc("/forget_pass", config.Config_forget_pass)
	http.HandleFunc("/create_user", config.Config_admin_create_user)
	http.HandleFunc("/user_home", config.Config_user_home)
	http.HandleFunc("/home_redirect", config.Config_home_redirect)
	http.HandleFunc("/admin_home", config.Config_admin_home)
	http.HandleFunc("/signup", config.Config_user_signup)
	http.HandleFunc("/admin_user_registration", config.Config_admin_user_registration_db)
	http.HandleFunc("/signup_user_registration", config.Config_signup_user_registration_db)

	//Handle function for  "User Managment"...
	http.HandleFunc("/user_managment", um.Um_user_managment)
	http.HandleFunc("/update_emp_data", um.Um_update_data_emp)
	http.HandleFunc("/profile_manange/", um.Um_profile_manange)
	http.HandleFunc("/update_emp/", um.Um_update_employee)
	http.HandleFunc("/delete_emp/", um.Um_delete_employee)
	http.ListenAndServe(":8080", nil)

	// http.HandleFunc("/create_user", config.Config_create_user)
	// http.HandleFunc("/create_emp_data", create_emp_data)
	// http.HandleFunc("/project_listing.html", project_listing)
	// http.HandleFunc("/ulogin", user)

}

// type Emp_data struct {
// 	Emp_id string
// }

// type project_data struct {
// 	P_id      string
// 	P_name    string
// 	P_key     string
// 	P_type    string
// 	P_lead    string
// 	P_empname string
// 	P_empid   string
// 	PM_id     string
// }

// type bug_data struct {
// 	B_id   string
// 	B_name string
// 	P_name string
// 	B_sum  string
// 	B_lead string
// 	B_sta  string
// 	B_ass  string
// }

// type empl_data struct {
// 	E_id      string
// 	E_empid   string
// 	E_empname string
// 	E_empmail string
// 	E_emppass string
// }

// type Data struct {
// 	Id      string
// 	Name    string
// 	Key     string
// 	Type    string
// 	Lead    string
// 	Empname string
// 	Empid   string
// }
