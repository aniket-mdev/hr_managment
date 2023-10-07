package main

import (
	"database/sql"
	"log"

	"github.com/aniket-mdev/hr_managment/apis"
	"github.com/aniket-mdev/hr_managment/apis/router"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

var (
	db_driver = "postgres"
	db_source = "postgresql://postgres:root@localhost:5432/postgres?sslmode=disable"
	address   = "0.0.0.0:8080"
)

func main() {

	db, err := sql.Open(db_driver, db_source)

	if err != nil {
		log.Fatal(err)
	}

	store := apis.NewStore(db)

	router_gin := gin.Default()

	router.UserRouter(router_gin, store)

	router_gin.Run(address)
}
