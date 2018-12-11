package main

import (
	"fmt"
	_"github.com/lib/pq"
	"database/sql"
)

//connectionString := "User ID=sa;Password=Password;Host=gopoc.c9lfhf7ayih8.ap-south-1.rds.amazonaws.com;Port=5432;Database=smarthome;Pooling=true;Min Pool Size=0;Max Pool Size=100;Connection Lifetime=0;"

const (
	host     = "gopoc.c9lfhf7ayih8.ap-south-1.rds.amazonaws.com"
	port     = 5432
	user     = "sa"
	password = "Password"
	dbname   = "smarthome"

  )


func main(){
	ConnectRDS();
}


func SaveEvent( key string,  value string){

}

func ConnectRDS( ){
	
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
	
	

	db,err := sql.Open("postgres",psqlInfo)
	
	if err != nil{
	fmt.Println("error : ",err)
		panic(err)
	}

	fmt.Println("rds connected")


	query := "CREATE TABLE Event(event_id VARCHAR (50)  PRIMARY KEY,value STRING	);"
	 _,er := db.Exec(query)

	 if er != nil{
		fmt.Println("error : ",er)
			panic(er)
		}
	fmt.Println("Event Table created")
		
}