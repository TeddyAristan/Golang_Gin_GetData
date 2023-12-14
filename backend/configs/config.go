package configs

import (
    "database/sql"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

func MyPort() (string,error){
	port := os.Getenv("PORT")
	if port == "" {
        port = "8080"
    }
	return ":" + port, nil
}

func Connectdb()(*sql.DB, error){
	db, errdb := sql.Open("mysql", "root:1234@tcp(localhost:3306)/dbgomycms")
	if errdb!= nil {
        return nil, errdb
    }
	err := db.Ping()
	return db,err
}