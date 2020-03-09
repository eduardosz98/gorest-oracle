package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/godror/godror"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func main() {
	var errDotenv error
	errDotenv = godotenv.Load()
	if errDotenv != nil {
		log.Fatalf("Error getting env, %v", errDotenv)
	} else {
		fmt.Println("We are getting the env values")
	}

	var connString string
	connString = fmt.Sprintf("%s/%s@%s", os.Getenv("ORACLE_USER"), os.Getenv("ORACLE_PASS"), os.Getenv("ORACLE_STRING"))
	fmt.Println(connString)
	db, err := sql.Open("godror", connString)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select descricao from fsenet_produtos")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var thedate string
	for rows.Next() {

		rows.Scan(&thedate)
	}
	fmt.Printf("The date is: %s\n", thedate)
}
