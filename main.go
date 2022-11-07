package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)
var (
	PostgresUser = "postgres"
	PostgresPassword = "1234"
	PostgresHost = "localhost"
	PostgresPort = 5432
	PostgresDatabase = "postgres"
)

func main (){
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		PostgresHost,
		PostgresPort,
		PostgresUser,
		PostgresPassword,
		PostgresDatabase,
	)
	db, err := sql.Open("postgres",connStr)
	if err != nil{
		log.Fatalf("failed to connection database: %v",err)
	}

	dbManager := NewDBManager(db)

	var newCar Car

	newCar.Brand = "Ford"
	newCar.Color = "green"
	newCar.Year = 2010

	// car, err := dbManager.CreateCar(&newCar)
	// if err != nil {
	// 	log.Fatalf("failed to update car: %v",err)
	// }
	// fmt.Println(car)

	fmt.Println(dbManager.GetCar(3))

////  =========================================
//                   |
//                   |
//                   |
//                   |
//                   |
//                 \ | /
//                 	\ /


//    delete da update qilyapti lekin hatolik beryapti



	car, err := dbManager.DeleteCar(7)
	if err != nil {
		log.Fatalf("failed to delete car: %v", err)
	}

	fmt.Println(car)
}