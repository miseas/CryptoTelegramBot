package repo

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

const db_name string = "sqlite-database.db"

func Start_db() {
	// os.Remove("sqlite-database.db")
	_, err := os.Stat(db_name)
	if os.IsNotExist(err) {
		log.Println("Creating db " + db_name)
		file, err := os.Create(db_name) // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		log.Println(db_name + " created")
	}
}

func Open_db_Connect() (db *sql.DB) {
	sqliteDatabase, _ := sql.Open("sqlite3", "./"+db_name)
	log.Println("Connecting to db...")
	return sqliteDatabase
}

func Close_db_Connect(db *sql.DB) {
	log.Println("Closing connection to db...")
	db.Close()
}
