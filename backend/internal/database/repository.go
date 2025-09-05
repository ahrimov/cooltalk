package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type MainDB struct {
	db *sql.DB
}

func OpenDatabase() *MainDB {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBNAME"))

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Cannot connect to database: %s", psqlInfo)
	}

	if pingErr := db.Ping(); pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Printf("Connected to %s:%s/%s", os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBNAME"))
	return &MainDB{db: db}
}

func (mainDB *MainDB) CloseDatabase() {
	mainDB.db.Close()
}

func (mainDB *MainDB) GetAllUsers() ([]User, error) {
	var users []User
	db := mainDB.db

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("get all users error: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt); err != nil {
			return nil, fmt.Errorf("get all user error: %v", err)
		}
		users = append(users, user)
	}

	return users, nil
}
