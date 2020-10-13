package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type Employee struct{
	Firstname string
	Lastname string
	Country string
	City     sql.NullString
	TelePhoneCode int `db:"telcode"`
}

func db_connect() *sqlx.DB{
	drivername := "mysql"
	username := "perennial"
	pass := "perennial"
	db_name := "perennial"
	db, err := sqlx.Connect(drivername, username+":"+pass+"@(127.0.0.1:3307)/"+db_name)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Create(){
	db:= db_connect()
	defer db.Close()
	db.Exec("DROP TABLE IF EXISTS employee")
	query := `CREATE TABLE IF NOT EXISTS employee(
                firstname TEXT NOT NULL,
                lastname TEXT NOT NULL,
                country TEXT ,
                city TEXT NULL,
                telcode INTEGER);`
	if _,err:= db.Exec(query); err!= nil{
		log.Fatalln(err)
	}

}

func Insert(){
	db:= db_connect()
	defer db.Close()
	CityTel := `INSERT INTO employee(firstname,lastname,country,telcode) VALUES(?,?,?,?)`
	CountryCity := `INSERT INTO employee(firstname,lastname,country,city,telcode) VALUES(?,?,?,?,?)`

	db.MustExec(CityTel,"steve","jobs","USA",85)
	db.MustExec(CountryCity,"smith","neymar","south Africa","johanesburg",28)
	db.MustExec(CityTel,"sandip","torane","India",91)
}

func Read(){
	db := db_connect()
	defer db.Close()
	fmt.Println("----for all row id''s----")
	rows, err := db.Queryx(`SELECT * FROM employee`)
	for rows.Next(){
		var user Employee
		err =rows.StructScan(&user)
		fmt.Println(user.Firstname,user.Lastname,user.Country,user.City,user.TelePhoneCode)
	}
	if err!=nil{
		log.Fatal(err)
	}
	defer rows.Close()

}

func Update(){
	db := db_connect()
	defer db.Close()
	firstname := "smith"
	lastname := "neymar"
	country_to_change :="UK"
	changes,err := db.Preparex("UPDATE employee SET country =? WHERE (firstname = ? AND lastname=?)")
	if err != nil {
		panic(err.Error())
	}
	changes.MustExec(country_to_change,firstname,lastname)
	fmt.Println("UPDATE:",country_to_change,"on name =",firstname,lastname)
}

func Delete(){
	db := db_connect()
	defer  db.Close()
	del, err := db.Preparex("DELETE FROM employee WHERE (firstname = ? AND lastname=?)")
	if err != nil {
		panic(err.Error())
	}
	del.MustExec("steve","jobs")
	log.Println("DELETE")
}

func get_select(){
	fmt.Println("..In get select ..")
	db:= db_connect()
	p := Employee{}
	pp := []Employee{}

	// this will pull the first place directly into p
	err := db.Get(&p, "SELECT * FROM employee LIMIT 1")
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(p)

	// this will pull places with telcode > 10 into the slice pp
	err = db.Select(&pp, "SELECT * FROM employee WHERE telcode > ?", 10)
	fmt.Println(pp)
	// they work with regular types as well
	var id int
	err = db.Get(&id, "SELECT count(*) FROM employee")
	fmt.Println(id)
	// fetch at most 10 place names
	var names []string
	err = db.Select(&names, "SELECT firstname FROM employee LIMIT 10")
	fmt.Println("names:",names)
}
func main(){
	Create()
	Insert()
	Read()
	Update()
	Delete()
	fmt.Println("after deleting ")
	Read()

	get_select()

}
