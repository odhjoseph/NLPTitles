package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

func main() {
	config := getEnv()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])
	db, err := sql.Open("postgres", psqlInfo)

	defer db.Close()

	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println"Connection... on ", dbhost)

}

func 

func getEnv() map[string]string {
	config := make(map[string]string)
	host, err := os.LookupEnv(dbhost)
	if !err {
		panic("DBHOST required")
	}
	port, err := os.LookupEnv(dbport)
	if !err {
		panic("DBPORT required")
	}
	user, err := os.LookupEnv(dbuser)
	if !err {
		panic("DBUSER required")
	}
	password, err := os.LookupEnv(dbpass)
	if !err {
		panic("DBPASS required")
	}
	name, err := os.LookupEnv(dbname)
	if !err {
        panic("DBNAME required")
        
    }

	config[dbhost] = host
	config[dbport] = port
	config[dbuser] = user
	config[dbpass] = password
	config[dbname] = name
	return config
}
