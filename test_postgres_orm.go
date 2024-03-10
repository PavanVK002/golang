package main

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

type User struct {
    ID   uint
    Name string
}

func main() {
    // Connect to the database
    dsn := "postgres://postgres:1&Onlypavan@localhost/bseSample?sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    // Auto Migrate the User struct
    err = db.AutoMigrate(&User{})
    if err != nil {
        log.Fatal(err)
    }

    // Create a new user
    user := User{Name: "John"}
    db.Create(&user)

    // Read the user
    var retrievedUser User
    db.First(&retrievedUser, user.ID)

    log.Println("Retrieved User:", retrievedUser)
}
