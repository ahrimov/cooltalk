package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

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

func (mainDB *MainDB) GetUserByID(id string) (*User, error) {
	var user User
	db := mainDB.db

	row := db.QueryRow("SELECT * FROM users WHERE id = $1", id)

	if err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt); err != nil {
		return nil, err
	}

	return &user, nil
}

func (mainDB *MainDB) SuggestUsersByUsername(username string) ([]User, error) {
	var users []User
	db := mainDB.db

	rows, err := db.Query("SELECT id, username FROM users WHERE username ILIKE $1 || '%'", username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			return nil, fmt.Errorf("fail in finding user: %v", err)
		}
		users = append(users, user)
	}

	return users, nil
}

func (mainDB *MainDB) AddNewUser(user User) (int64, error) {
	db := mainDB.db
	var id int64

	result := db.QueryRow("INSERT INTO users (username, password, email) VALUES($1, $2, $3) RETURNING id", user.Username, user.Password, user.Email)
	if err := result.Scan(&id); err != nil {
		return 0, fmt.Errorf("AddNewUser error: %v", err)
	}
	return id, nil

}

func (mainDB *MainDB) DeleteUser(id string) (int64, error) {
	db := mainDB.db
	var deletedId int64

	result := db.QueryRow("DELETE FROM users WHERE id = $1 RETURNING id", id)
	if err := result.Scan(&deletedId); err != nil {
		return 0, fmt.Errorf("DeleteUser: %v", err)
	}
	return deletedId, nil
}

func (mainDB *MainDB) UpdateUser(id string, updatedData map[string]interface{}) (*User, error) {
	db := mainDB.db
	var user User

	params := make([]interface{}, 0, len(updatedData))
	paramsCount := 1
	query := "UPDATE users SET "
	for key, value := range updatedData {
		queryPart := fmt.Sprintf("%s = $%d,", key, paramsCount)
		query += queryPart
		params = append(params, value)
		paramsCount++
	}

	query = strings.TrimSuffix(query, ",") + " WHERE id = " + id + " RETURNING id, username, email"

	result := db.QueryRow(query, params...)
	if err := result.Scan(&user.ID, &user.Username, &user.Email); err != nil {
		return nil, fmt.Errorf("update user error: %v", err)
	}

	return &user, nil
}
