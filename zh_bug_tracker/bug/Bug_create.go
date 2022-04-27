package bug

import (

	// this package is used for dialog

	"zh_bug_tracker/initial"
	// this package is used for DB initial

	_ "database/sql"
	// this package is for DB query Execution

	"fmt"
	"net/http"
)

type Data struct {
	Id      string
	Name    string
	Key     string
	Type    string
	Lead    string
	Empname string
	Empid   string
}

func Bug_create(w http.ResponseWriter, r *http.Request) {
	var (
		query      string
		Sl_project []Data
	)
	emp_id := r.FormValue("emp")

	fmt.Print("emp_id :", emp_id)

	//db, err := initial.Connect_db()
	//initial.CheckError(err)

	if emp_id == "1" {
		query = fmt.Sprintf("select pr.p_id ,p.p_name,pr.emp_id,e.emp_name from project p inner join project_mates pr on p.p_id = pr.p_id inner join employee e on e.emp_id=pr.emp_id order by p.p_name")
		fmt.Print("query :", query)

	} else {
		query = fmt.Sprintf("select pr.p_id ,p.p_name,pr.emp_id,e.emp_name from project p inner join project_mates pr on p.p_id = pr.p_id and  pr.p_id in (select pr.p_id from project_mates pr where pr.emp_id='%s') inner join employee e on e.emp_id=pr.emp_id;", emp_id)
		fmt.Print("query :", query)
	}

	rows, err := initial.Db.Query(query)
	initial.CheckError(err)

	fmt.Print("query values:", rows)
	Project_data := Data{}
	for rows.Next() {
		err := rows.Scan(&Project_data.Id, &Project_data.Name, &Project_data.Empid, &Project_data.Empname)
		initial.CheckError(err)

		Sl_project = append(Sl_project, Project_data)
	}
	fmt.Println("slice of data :", Sl_project)

	defer rows.Close()

	initial.Tpl.ExecuteTemplate(w, "bug_create.html", Sl_project)
	// query := `select pr.p_id ,p.p_name,pr.emp_id,e.emp_name from project p inner join project_mates pr on p.p_id = pr.p_id and  pr.p_id in (select pr.p_id from project_mates pr where pr.emp_id=$1) inner join employee e on e.emp_id=pr.emp_id;`
	// rows, err := db.Query("select pr.p_id,p.p_name,e.emp_id,e.emp_name from project_mates pr inner join employee e on pr.emp_id='$1' inner join project p on p.p_id =pr.p_id;")

}
