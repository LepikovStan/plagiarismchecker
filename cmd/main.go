package main

import (
	"database/sql"
	"fmt"
	"log"

	//"gorm.io/driver/postgres"
	//"gorm.io/driver/postgres"
	//"gorm.io/driver/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	//_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	_ "github.com/lib/pq"
)

type Task struct {
	ID              string
	State           string
	OriginalArticle string
	ErrorMessage    string
}

func main() {
	//dbConn, err := sql.Open("postgres", fmt.Sprintf(`postgres://pl-checker-psql:vD-/ER_"C0daIF6A@35.202.217.250/plagiarism_checker?sslmode=disable`)) //"postgres://postgres:vD-/ER_\"C0daIF6A@35.238.150.37:5432/plagiarism_checker?sslmode=disable")
	dbConn, err := sql.Open("postgres", fmt.Sprintf(`user=postgres password=1234 host=35.202.217.250 port=5432 dbname=plagiarism_checker sslmode=disable`)) //"postgres://postgres:vD-/ER_\"C0daIF6A@35.238.150.37:5432/plagiarism_checker?sslmode=disable")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("dbConn -->", dbConn)

	if err := dbConn.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("done")

	// db, err := gorm.Open(postgres.New(postgres.Config{
	// 	DriverName: "cloudsqlpostgres",
	// 	DSN:        "host=plagiarism-checker-377309:us-central1:pl-checker-psql user=puser dbname=plagiarism_checker password=12345 sslmode=disable",
	// }))
	dsn := "user=pl-checker-psql  password=1234 host=35.202.217.250 port=5432 dbname=plagiarism_checker sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	result := db.Select("State", "OriginalArticle").Create(&Task{State: "created", OriginalArticle: "test"})
	fmt.Println(result.RowsAffected)
}
