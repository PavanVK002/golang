package main

import (
    "context"
    "database/sql"
    "fmt"
    "log"
    // "time"

    _ "github.com/denisenkom/go-mssqldb"
)
type User struct {
	ID   int
	Name string
	Age  int
}
var db *sql.DB
func main() {
    // Database connection string
    //connString := "Server=(localdb)\\Local;Database=LearningDB;Integrated Security=true;"
	
	var (
		server    string = "localhost" // for example
		user      string = "pavanvk"    // Database user
		password  string = "1&Onlypavan"   // User Password
		port      int    = 1433        // Database port
	)
	
	
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=LearningDB", server, user, password, port)
	db, err := sql.Open("sqlserver", connString)
	
	// Test if the connection is OK or not
	if err != nil {
		panic("Cannot connect to database")
	} else {
		fmt.Println("Connected!")
		fmt.Println(db)
	
	  }

	  defer db.Close()
	 selectversion()
	//   row := db.QueryRowContext(context.Background(), "SELECT FirstName, LastName FROM Students WHERE StudentID = ?", "15")
	//   // extract data
	//   var firstname, lastname string
	//   err = row.Scan(&firstname, &lastname)
	//   if err != nil {
	// 	  log.Fatalf("impossible to scan rows of query: %s", err)
	//   }
	//   log.Printf("student number 15 is %s %s", firstname, lastname)
	// // Don't forget to close the connection to your database
	
}

func selectversion(){
	ctx := context.Background()

	err := db.PingContext(ctx)
	if err != nil{
		log.Fatal("Error pinging database : " +err.Error())

	}
	var result string

	err =db.QueryRowContext(ctx,"SELECT @@version").Scan(&result)
	if err !=nil{
		log.Fatal("Scan failed ",err.Error())
	}
	fmt.Println(result)
}


