package main

import (
    // "context"
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
func main() {
    // Database connection string
    // connString := "Server=(localdb)\\Local;Database=LearningDB;Integrated Security=true;"
	
	var (
		server    string = "localhost" // for example
		user1      string = "pavan.vasant"    // Database user
		password  string = "PV&97jh@@X"   // User Password
		port      int    = 1433        // Database port
	)
	
	
	connString := fmt.Sprintf("server=%s;user id=%s;Database=LearningDB;password=%s;port=%d", server, user1, password, port)
	
    // Create a new SQL Server database connection
    db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create Table
	createTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		name TEXT,
		age INTEGER
	);
	`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	user := User{ID: 1, Name: "Ricky Pointing", Age: 52}
	err = createUser(db, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User1 created successfully")
	user2 := User{ID: 2, Name: "John Cenna", Age: 54}
	err = createUser(db, user2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User2 created successfully")

	// Read
	users, err := getUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Users:")
	for _, u := range users {
		fmt.Printf("%d: %s, %d\n", u.ID, u.Name, u.Age)
	}

	// Update
	userToUpdate := User{ID: 2, Name: "John Cenna", Age: 56}
	err = updateUser(db, userToUpdate)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User2 updated successfully")

	// Delete
	err = deleteUser(db, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User1 deleted successfully")
}

func createUser(db *sql.DB, user User) error {
	_, err := db.Exec("INSERT INTO users (id,name, age) VALUES ($1, $2, $3)", user.ID, user.Name, user.Age)
	return err
}

func getUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func updateUser(db *sql.DB, user User) error {
	_, err := db.Exec("UPDATE users SET name=$1, age=$2 WHERE id=$3", user.Name, user.Age, user.ID)
	return err
}

func deleteUser(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}
