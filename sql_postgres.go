package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/lib/pq"
// )

// type User struct {
// 	ID   int
// 	Name string
// 	Age  int
// }

// func main() {
// 	db, err := sql.Open("postgres", "postgres://postgres:1&Onlypavan@localhost/bseSample?sslmode=disable")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
	
// 	// Create Table
// 	createTable := `
// 	CREATE TABLE IF NOT EXISTS users (
// 		id INTEGER PRIMARY KEY,
// 		name TEXT,
// 		age INTEGER
// 	);
// 	`
// 	_, err = db.Exec(createTable)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Create
// 	user := User{ID: 1, Name: "Ricky Pointing", Age: 52}
// 	err = createUser(db, user)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("User1 created successfully")
// 	user2 := User{ID: 2, Name: "John Cenna", Age: 54}
// 	err = createUser(db, user2)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("User2 created successfully")

// 	// Read
// 	users, err := getUsers(db)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Users:")
// 	for _, u := range users {
// 		fmt.Printf("%d: %s, %d\n", u.ID, u.Name, u.Age)
// 	}

// 	// Update
// 	userToUpdate := User{ID: 2, Name: "John Cenna", Age: 56}
// 	err = updateUser(db, userToUpdate)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("User2 updated successfully")

// 	// Delete
// 	err = deleteUser(db, 2)
// 	err = deleteUser(db, 1)
	
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("User1 deleted successfully")
// }

// func createUser(db *sql.DB, user User) error {
// 	_, err := db.Exec("INSERT INTO users (id,name, age) VALUES ($1, $2, $3)", user.ID, user.Name, user.Age)
// 	return err
// }

// func getUsers(db *sql.DB) ([]User, error) {
// 	rows, err := db.Query("SELECT id, name, age FROM users")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var users []User
// 	for rows.Next() {
// 		var user User
// 		err := rows.Scan(&user.ID, &user.Name, &user.Age)
// 		if err != nil {
// 			return nil, err
// 		}
// 		users = append(users, user)
// 	}
// 	return users, nil
// }

// func updateUser(db *sql.DB, user User) error {
// 	_, err := db.Exec("UPDATE users SET name=$1, age=$2 WHERE id=$3", user.Name, user.Age, user.ID)
// 	return err
// }

// func deleteUser(db *sql.DB, id int) error {
// 	_, err := db.Exec("DELETE FROM users WHERE id=$1", id)
// 	return err
// }
import (
    "fmt"
    "database/sql"
    "encoding/json"
    // "fmt"
    "log"

    _ "github.com/lib/pq"
)

type EODData struct {
    IndicesWatchName string  `json:"IndicesWatchName"`
    Curvalue         float64 `json:"Curvalue"`
    PrevDayClose     float64 `json:"PrevDayClose"`
    CHNG             float64 `json:"CHNG"`
    CHNGPER          float64 `json:"CHNGPER"`
    DT_TM            string  `json:"DT_TM"`
}

func main() {
    // JSON data
    jsonData := `{"RealTime":[],"ASON":[],"EOD":[{"IndicesWatchName":"S&P BSE SENSEX 1-Month Real Vol ","Curvalue":15.16416067646739,"PrevDayClose":10.23648900702580,"CHNG":4.92767166944159,"CHNGPER":48.14,"DT_TM":"2024-03-01T00:00:00","IndexSrNo":1.0,"rn":1},{"IndicesWatchName":"S&P BSE SENSEX 2-Month Real Vol ","Curvalue":12.07862932775694,"PrevDayClose":11.02405730019034,"CHNG":1.05457202756660,"CHNGPER":9.57,"DT_TM":"2024-03-01T00:00:00","IndexSrNo":2.0,"rn":1},{"IndicesWatchName":"S&P BSE SENSEX 3-Month Real Vol ","Curvalue":12.07862932775694,"PrevDayClose":11.02405730019034,"CHNG":1.05457202756660,"CHNGPER":9.57,"DT_TM":"2024-03-01T00:00:00","IndexSrNo":3.0,"rn":1}]}`
    fmt.Println(jsonData)
    // Parse JSON data
    var jsonDataMap map[string][]EODData
    err := json.Unmarshal([]byte(jsonData), &jsonDataMap)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(jsonDataMap)
    // Open connection to PostgreSQL
    db, err := sql.Open("postgres", "postgres://postgres:1&Onlypavan@localhost/bseSample?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Create EOD data table if it does not exist
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS eod_data (
        "Index" TEXT,
        "Current Value" FLOAT,
        "Prev. Close" FLOAT,
        "Ch (pts)" FLOAT,
        "Ch (%)" FLOAT,
        "Date" TIMESTAMP
    )`)
    if err != nil {
        log.Fatal(err)
    }

    // Insert EOD data into PostgreSQL
    for _, eod := range jsonDataMap["EOD"] {
        _, err := db.Exec("INSERT INTO eod_data (\"Index\", \"Current Value\", \"Prev. Close\", \"Ch (pts)\", \"Ch (%)\", \"Date\") VALUES ($1, $2, $3, $4, $5, $6)",
            eod.IndicesWatchName, eod.Curvalue, eod.PrevDayClose, eod.CHNG, eod.CHNGPER, eod.DT_TM)
        if err != nil {
            log.Fatal(err)
        }
    }
}
