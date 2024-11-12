package main

import (
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/siddheshRajendraNimbalkar/API-JWT-MYSQL/cmd/api"
	"github.com/siddheshRajendraNimbalkar/API-JWT-MYSQL/config"
	"github.com/siddheshRajendraNimbalkar/API-JWT-MYSQL/db"
)

func main() {
	fmt.Println("Server is now running")
	//db connection
	mydb, err := db.NewMYSqlStorage(mysql.Config{
		User:                 config.InitConfig.DBUser,
		Passwd:               config.InitConfig.Passwd,
		Addr:                 config.InitConfig.DBAddress,
		DBName:               config.InitConfig.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	err = db.InitDataBase(mydb)

	if err != nil {
		log.Fatal("[ERROR IN PACKAGE DB]:", err)
	}

	server := api.NewAPIServer(":8080", mydb)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
