package repo

import (
	// "fmt"
	// "math"
	// "strconv"

	// tb "gopkg.in/tucnak/telebot.v2"

	// "cryptoTelegramBot/utils"
	// "cryptoTelegramBot/utils"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

type CryptoNotification struct {
	Id               int
	UserId           string
	Symbol           string
	CompareCondition string
	CompareValue     float64
	CurrentValue     float64
	UpdatedAt        string
}

func Create_tables() {
	log.Println("Creating tables.db...")
	db := Open_db_Connect()
	createNotificationTableSQL := `
						CREATE TABLE IF NOT EXISTS crypto_notifications (
								id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
								user_id text NOT NULL,
								symbol text NOT NULL,
								compare_condition text NOT NULL,
								compare_value DECIMAL NOT NULL,
								current_value DECIMAL,
								updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
								created_at DATETIME DEFAULT CURRENT_TIMESTAMP
								);
								`
	statement, err := db.Prepare(createNotificationTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("tables created!")
	Close_db_Connect(db)
}

func InsertNotification(user_id string, symbol string, compare_condition string, compare_value float64) {
	log.Println("Inserting new notification record...")
	db := Open_db_Connect()
	insertSQL := `INSERT INTO crypto_notifications(user_id, symbol, compare_condition, compare_value) VALUES (?,?,?,?);`
	statement, err := db.Prepare(insertSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	res, err := statement.Exec(user_id, symbol, compare_condition, compare_value)
	log.Println(res)
	if err != nil {
		log.Fatalln(err.Error())
	}
	Close_db_Connect(db)
}

func GetAllNotifications() []CryptoNotification {
	log.Println("Get all notifications...")
	db := Open_db_Connect()
	row, err := db.Query("SELECT id, user_id, symbol, compare_condition, compare_value, updated_at FROM crypto_notifications ORDER BY user_id")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	var notifications []CryptoNotification
	for row.Next() { // Iterate and fetch the records from result cursor
		notification := CryptoNotification{}
		row.Scan(&notification.Id, &notification.UserId, &notification.Symbol,
			&notification.CompareCondition, &notification.CompareValue, &notification.UpdatedAt)
		notifications = append(notifications, notification)
	}
	Close_db_Connect(db)

	return notifications
}

func GetUserNotifications(user_id string) []CryptoNotification {
	log.Println("Get user notifications...")
	db := Open_db_Connect()
	row, err := db.Query(fmt.Sprintf(`SELECT id, user_id, symbol, compare_condition, compare_value, updated_at 
							FROM crypto_notifications
							WHERE user_id = %s`, user_id))
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	var notifications []CryptoNotification
	for row.Next() { // Iterate and fetch the records from result cursor
		notification := CryptoNotification{}
		row.Scan(&notification.Id, &notification.UserId, &notification.Symbol,
			&notification.CompareCondition, &notification.CompareValue, &notification.UpdatedAt)
		notifications = append(notifications, notification)
	}
	Close_db_Connect(db)

	return notifications
}

func DeleteNotification(user_id string, symbol string, compare_condition string, compare_value float64) (int64, error) {
	log.Println("Delete notification record...")
	db := Open_db_Connect()
	deleteSQL := `Delete from crypto_notifications where user_id = ? and symbol = ? and compare_condition = ? and compare_value = ?;`
	statement, err := db.Prepare(deleteSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	res, err := statement.Exec(user_id, symbol, compare_condition, compare_value)
	affected_rows, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err.Error())
	}
	Close_db_Connect(db)
	return affected_rows, err
}
