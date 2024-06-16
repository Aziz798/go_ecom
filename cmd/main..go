package main

import (
	"database/sql"
	"log"

	"github.com/Aziz798/ecom/cmd/api"
	"github.com/Aziz798/ecom/config"
	"github.com/Aziz798/ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMysqlStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal("Error in db %w", err)
	}
	intiStorage(db)
	server := api.NewAPIServer(config.Envs.Port, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func intiStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB successfully connected")
}
