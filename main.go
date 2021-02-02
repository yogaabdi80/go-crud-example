package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yogaabdi80/go-crud-example/api"
	"github.com/yogaabdi80/go-crud-example/config"
)


func main() {

	r := api.Router()
    fmt.Println("Server dijalankan pada port 8080...")

	db := config.CreateConnection()
	defer db.Close()

	log.Fatal(http.ListenAndServe(":8080", r))
}

