package main

import (
	"os"

	"github.com/igorlage/curso/fullcycle/desafio1/application/grpc"
	"github.com/igorlage/curso/fullcycle/desafio1/infrastructure/db"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
