package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetDefaultList() {
	db, err := DbConnect()
	if err != nil {
		return
	}

	rows, err := db.Query("CALL get_default_list()")
	if err != nil {
		fmt.Printf("Failed to execute DB procedure! -> Error : %s\n", err.Error())
		return
	}
	defer rows.Close()

	db.Close()
}

func DbConnect() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUsername, DBPassword, DBHost, DBPort, DBName)
	Conn, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Printf("Failed to connect to database! -> \nError : %s\nDSN : %s\n", err.Error(), dsn)
		return nil, err
	}

	err = Conn.Ping()
	if err != nil {
		fmt.Printf("DB Connection timed out or invalid connection! -> \nError : %s\n", err.Error())
		return nil, err
	}

	fmt.Println("Connected to the database!")
	return Conn, nil
}
