package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type info struct{
	id int
	firstname string
	lastname string
	email string
}

func db_connect() *sql.DB{
	drivername := "mysql"
	username := "perennial"
	pass := "perennial"
	db_name := "perennial"
	db, err := sql.Open(drivername, username+":"+pass+"@(127.0.0.1:3307)/"+db_name)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

func create(){
	db:= db_connect()
	defer db.Close()
	db.Exec("DROP TABLE IF EXISTS info")
	query := `CREATE TABLE IF NOT EXISTS info (
                id INT AUTO_INCREMENT,
                firstname TEXT NOT NULL,
                lastname TEXT NOT NULL,
                email  varchar(50) UNIQUE NOT NULL,
                PRIMARY KEY (id)
            );`
	if _,err:= db.Exec(query); err!= nil{
		log.Fatalln(err)
	}

}
func insert(firstname,lastname,email string){
	db:= db_connect()
	defer db.Close()
	result, err := db.Exec(`INSERT INTO info (firstname, lastname, email) VALUES (?, ?, ?)`, firstname, lastname, email)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	fmt.Println(id)
}

func read(){
	db := db_connect()
	defer db.Close()
	var(
		id int
		firstname string
		lastname  string
		email string
	)
	fmt.Println("for single id")
	query := "SELECT id, firstname, lastname, email FROM info WHERE id = ?"
	if err := db.QueryRow(query, 1).Scan(&id, &firstname, &lastname, &email); err != nil {
		log.Fatal(err)
	}

	fmt.Println(id, firstname, lastname, email)

	fmt.Println("for all row id''s")
	rows, err := db.Query(`SELECT id, firstname, lastname, email FROM info`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []info
	for rows.Next() {
		var u info

		err := rows.Scan(&u.id, &u.firstname, &u.lastname, &u.email)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _,user := range users{
		fmt.Println(user.id,user.firstname,user.lastname,user.email)
	}
}

func update(){
	db := db_connect()
	defer db.Close()
	firstname := "dhiraj"
	lastname := "gurav"
	email := "dhirajgurav@gmail.com"
	id_to_udpate :=2
	changes,err := db.Prepare("UPDATE info SET firstname = ?,lastname = ?,email = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	changes.Exec(firstname,lastname,email,id_to_udpate)
    fmt.Println("UPDATE:",firstname,lastname,email,"on id =",id_to_udpate)
}

func delete(id_to_delete int){
	db := db_connect()
	defer  db.Close()
	del, err := db.Prepare("DELETE FROM info WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	del.Exec(id_to_delete)
	log.Println("DELETE")
}
func main(){
	create()
	firstname:= "sandip"
	lastname := "torane"
	email := "sandiptorane@gmail.com"
	insert(firstname,lastname,email)
	insert("pankaj","torane","pankajtorane@gmail.com")
	read()
	update()
	id_to_delete :=2
	delete(id_to_delete)
	fmt.Println("after deleting ",id_to_delete)
	read()

}