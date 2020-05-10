package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/rds/rdsutils"
	//"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	dbEndpoint = "jsondatabase.cm1zpldcgrzz.us-east-1.rds.amazonaws.com"
	awsRegion  = "us-east-1f"
	dbUser     = "postgres"
	dbName     = "jsondatabase"
	password   = "ufkS3GfMnclQUCyKVbAN" //aye fam delete this
)

// func init() {
// 	// loads values from .env into the system
// 	if err := godotenv.Load(); err != nil {
// 		log.Print("No .env file found")
// 	}
// }

func main() {
	awsCreds := credentials.NewEnvCredentials()

	_, err := awsCreds.Get()
	if err != nil {
		panic(err.Error())
	}

	authToken, err := rdsutils.BuildAuthToken(dbEndpoint, awsRegion, dbUser, awsCreds)
	if err != nil {
		log.Println("Token Failed to be created", err)
		return
	}
	//%s:%s@tcp(%s)/%s?tls=true
	//postgres://%s:%s@%s/%s?sslmode=verify-full
	var dnsStr = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=verify-full", dbUser, url.PathEscape(authToken), dbEndpoint, dbName)

	db, err := sql.Open("postgres", dnsStr)

	if err != nil {
		log.Println("Damn Again", err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("Cannot ping db: %s\n", err)
		return
	}

	db.Close()
}
