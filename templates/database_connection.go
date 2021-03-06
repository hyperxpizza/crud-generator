package templates

import "text/template"

//DatabaseConnectionTemplate is a template for database.go file which is responslible for postgres connection
var DatabaseConnectionTemplate = template.Must(template.New("database.go").Parse(`
// This file was generated by robots at {{ .Timestamp }}
// Feel free to be a part of this project: github.com/hyperxpizza/crud-generator

package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	//Postgresql driver
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(user, password, dbname, host, port string) error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

	database, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("sql.Open failed: %v\n", err)
		return err
	}

	err = database.Ping()
	if err != nil {
		log.Fatalf("database.Ping failed: %v\n", err)
		return err
	}

	db = database
	log.Println("[+] Connected to the database")

	return nil
}`))
