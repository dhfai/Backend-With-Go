package main

import (
	"database/sql"
	"github.com/ddhaifullah/go-auth-api/cmd/api"
	"github.com/ddhaifullah/go-auth-api/config"
	"github.com/ddhaifullah/go-auth-api/db"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := db.NewMySqlStorage(mysql.Config{
		User:   config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr:   config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		//Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPISertver(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected successfully!")
}
