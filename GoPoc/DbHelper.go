package main

import (
	"fmt"
	"database/sql"
	_"github.com/lib/pq"
)

//connectionString := "User ID=sa;Password=Password;Host=gopoc.c9lfhf7ayih8.ap-south-1.rds.amazonaws.com;Port=5432;Database=smarthome;Pooling=true;Min Pool Size=0;Max Pool Size=100;Connection Lifetime=0;"

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "postgres"
  )



func SaveEvent( key ,  value string){

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db,err := sql.Open("postgres",psqlInfo)
	
	if err != nil{
	fmt.Println("error : ",err)
		panic(err)
	}

	fmt.Println("rds connected")


	sqlStatement := `INSERT INTO events (event_id, value) VALUES ($1,$2)`

	_,err = db.Query(sqlStatement,key,value)
	if err != nil {
		panic(err)
	}

	 if err != nil{
		fmt.Println("error : ",err)
			panic(err)
		}
	//fmt.Println("Event Table created")
	fmt.Println("Query executed successfully")

		
}