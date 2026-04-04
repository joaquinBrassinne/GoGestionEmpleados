package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var templates = template.Must(template.ParseGlob("templates/*"))

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/edit", Edit)

	log.Println("Server running...")

	http.ListenAndServe(":8080", nil)
}

func DBConnect() (c *sql.DB) {
	Driver := "mysql"
	User := getEnv("DB_USER", "root")
	Pass := getEnv("DB_PASSWORD", "root")
	Host := getEnv("DB_HOST", "localhost")
	Name := getEnv("DB_NAME", "golang_crud")

	c, err := sql.Open(Driver, User+":"+Pass+"@tcp("+Host+":3306)/"+Name)

	if err != nil {
		log.Fatal(err)
	}
	return c
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

type Employe struct {
	Id    int
	Name  string
	Email string
}

func Home(w http.ResponseWriter, r *http.Request) {
	connect := DBConnect()
	register, err := connect.Query("SELECT * FROM employees")

	if err != nil {
		log.Fatal(err)
	}

	employees := Employe{}
	employeesList := []Employe{}

	for register.Next() {
		var id int
		var name string
		var email string
		err = register.Scan(&id, &name, &email)

		if err != nil {
			log.Fatal(err)
		}

		employees.Id = id
		employees.Name = name
		employees.Email = email

		employeesList = append(employeesList, employees)
	}
	log.Println(employeesList)

	defer connect.Close()

	templates.ExecuteTemplate(w, "start-tmpl.html", employeesList)
}

func Create(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "create-tmpl.html", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		email := r.FormValue("email")

		connect := DBConnect()
		insert, err := connect.Prepare("INSERT INTO employees(name,email) VALUES (?,?)")

		if err != nil {
			log.Fatal(err)
		}

		insert.Exec(name, email)

		defer connect.Close()
		http.Redirect(w, r, "/", 301)
	}

}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	connect := DBConnect()
	delete, err := connect.Prepare("DELETE FROM employees WHERE id = ?")

	if err != nil {
		log.Fatal(err)
	}

	delete.Exec(id)

	defer connect.Close()
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	connect := DBConnect()
	register, err := connect.Query("SELECT * FROM employees WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
	}

	employees := Employe{}
	for register.Next() {
		var id int
		var name string
		var email string
		err = register.Scan(&id, &name, &email)

		if err != nil {
			log.Fatal(err)
		}

		employees.Id = id
		employees.Name = name
		employees.Email = email
	}

	defer connect.Close()
	templates.ExecuteTemplate(w, "edit-tmpl.html", employees)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		email := r.FormValue("email")

		connect := DBConnect()
		update, err := connect.Prepare("UPDATE employees SET name = ?, email = ? WHERE id = ?")

		if err != nil {
			log.Fatal(err)
		}

		update.Exec(name, email, id)

		defer connect.Close()
		http.Redirect(w, r, "/", 301)
	}
}
