package initial

import (
	"database/sql"
	"fmt"

	"html/template"

	_ "github.com/lib/pq"
)

var (
	Employee_id string
	Tpl         *template.Template
	Db          *sql.DB
	Err         error
)

type Emp_data struct {
	Emp_id string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456789"
	dbname   = "bugtracker_db"
)

func init() {
	Tpl = template.Must(template.ParseGlob("templates/*.html"))

	// psqlconn := fmt.Sprintf("host=%s port =%d user =%s password =%s dbname =%s sslmode =disable", host, port, user, password, dbname)

	// Db, Err = sql.Open("postgres", psqlconn)
	// CheckError(Err)

	// Err = Db.Ping()
	// CheckError(Err)

	// defer Db.Close()

}

func Connect_db() {
	psqlconn := fmt.Sprintf("host=%s port =%d user =%s password =%s dbname =%s sslmode =disable", host, port, user, password, dbname)

	Db, Err = sql.Open("postgres", psqlconn)
	CheckError(Err)

	Err = Db.Ping()
	CheckError(Err)

}

// func Connect_db() (*sql.DB, error) {
// 	psqlconn := fmt.Sprintf("host=%s port =%d user =%s password =%s dbname =%s sslmode =disable", host, port, user, password, dbname)

// 	db, err := sql.Open("postgres", psqlconn)
// 	CheckError(err)

// 	// defer db.Close()

// 	err = db.Ping()
// 	CheckError(err)

// 	// insertstmt := fmt.Sprintf(`insert into "login"("id", "name","password") values(DEFAULT, '%s', '%s')`, name, pass)

// 	// _, e := initial.Db.Exec(insertstmt)
// 	// CheckError(e)
// 	return db, nil

// 	//  insertdystmt := 'insert into "login" ("id","name", "password") values (DEFAULT,'kiran','1234')'

// }
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
