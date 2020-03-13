package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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

type Article struct {
	Hyperlink string
}

func main() {
	// var paths = []string{"/Users/josephodhiambo/Python/NLPTitles/scripts/jsonFeeds/",
	// 	"/Users/josephodhiambo/Python/NLPTitles/scripts/titleDisplay/"}
	var articles map[string]interface{}

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

	log.Println("Connection... on")
	//sqlState := `CREATE TABLE feeds(hyperlink TEXT NOT NULL, titles TEXT NOT NULL);`
	// sqlState := `
	// INSERT INTO feeds (hyperlink, titles)
	// VALUES('tester.com', 'pleaseWork');`

	jsonFile, err := os.Open("/Users/josephodhiambo/go/godev/NLPTitles/scripts/titleDisplay/search2020-03-10.json")
	if err != nil {
		log.Println("This shouldn't be possible, unless empty directory", err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &articles)

	for k := range articles {
		if articles[k] != nil {
			str := fmt.Sprintf("%v", articles[k])
			sqlState := `
			INSERT INTO feeds(hyperlink, titles)
			VALUES('` + k + `', '` + str + `');`
			_, err = db.Exec(sqlState)
			if err != nil {
				log.Print("Couldn't write this aritcle ", k, err)
				continue
			}
		}

	}

}

func writeJSONtoSQL(db *sql.DB, file string) bool {

	//	sqlState := `CREATE TABLE feeds(id SERIAL PRIMARY KEY, hyperlink TEXT NOT NULL, titles TEXT NOT NULL);`
	if isAlreadyInDatabase(file) {
		return false //fix implementaiton to handle error handling
	}
	var paths = []string{"/Users/josephodhiambo/Python/NLPTitles/scripts/jsonFeeds/",
		"/Users/josephodhiambo/Python/NLPTitles/scripts/titleDisplay/"}

	var articles map[string]interface{}

	jsonFile, err := os.Open(paths[1] + file)
	if err != nil {
		log.Println("This shouldn't be possible, unless empty directory", err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &articles)

	sqlstate := `
	INSERT INTO rssInfo (hyperlink, titles)
	VALUES('tester.com2', 'pleaseWork2');`
	_, err = db.Exec(sqlstate)

	return true //doesn't work
}

func isAlreadyInDatabase(fileName string) bool {
	file, err := os.Open("written.csv")

	if err != nil {
		panic("Failed to find CSV file")
	}

	r := csv.NewReader(file)

	for {
		readFile, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		return contains(readFile, fileName)
	}

	return true

}

func contains(allFiles []string, targetFile string) bool {
	for _, fileName := range allFiles {
		if fileName == targetFile {
			return true
		}
	}
	return false
}

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
