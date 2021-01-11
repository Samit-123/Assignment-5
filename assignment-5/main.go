package main

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Location   string `json:"location"`
	Department string `json:"department"`
	Income     int    `json:"Income"`
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "Employee_db"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	//new template engine

	router.GET("/", func(ctx *gin.Context) {
		//render only file, must full name with extension
		db := dbConn()
		sd, err := db.Query("SELECT * FROM Person ORDER BY id DESC")
		if err != nil {
			panic(err.Error())
		}
		emp := Person{}
		res := []Person{}
		for sd.Next() {
			var id, income int
			var name, location, department string
			err = sd.Scan(&id, &name, &location, &department, &income)
			if err != nil {
				panic(err.Error())
			}
			emp.Id = id
			emp.Name = name
			emp.Location = location
			emp.Department = department
			emp.Income = income
			res = append(res, emp)
		}
		//var a = "hello words"
		ctx.HTML(http.StatusOK, "index.html", gin.H{"title": "index file title!!", "a": res})
	})

	router.GET("/about", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "about.html", gin.H{"title": "index file title!!"})
	})

	router.GET("/contact", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "contact.html", gin.H{"title": "index file title!!"})
	})

	router.GET("/addnew", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "addnew.html", gin.H{"title": "index file title!!"})
	})

	router.GET("/submit", func(ctx *gin.Context) {
		//render only file, must full name with extension
		var name, location, department string
		var income int

		name = ctx.Request.FormValue("name")
		location = ctx.Request.FormValue("location")
		department = ctx.Request.FormValue("department")

		sal := ctx.Request.FormValue("income")
		income, _ = strconv.Atoi(sal)
		db := dbConn()
		insForm, err := db.Prepare("INSERT INTO Person(name, location, department, income) VALUES(?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, location, department, income)

		sd, err := db.Query("SELECT * FROM Person ORDER BY id DESC")
		if err != nil {
			panic(err.Error())
		}
		emp := Person{}
		res := []Person{}
		for sd.Next() {
			var id, income int
			var name, location, department string
			err = sd.Scan(&id, &name, &location, &department, &income)
			if err != nil {
				panic(err.Error())
			}
			emp.Id = id
			emp.Name = name
			emp.Location = location
			emp.Department = department
			emp.Income = income
			res = append(res, emp)
		}

		ctx.HTML(http.StatusOK, "index.html", gin.H{"title": "index file title!!", "a": res})
	})

	router.Run(":9090")
}
