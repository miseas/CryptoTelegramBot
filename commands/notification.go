package commands

import (
	// "fmt"
	// "math"
	// "strconv"

	// "cryptoTelegramBot/repo"
	"cryptoTelegramBot/repo"
	"cryptoTelegramBot/utils"
	"fmt"
	"log"
	"strconv"
)

func CreateNotification(user_id string, sym string, condition string, compare_value string) (string, error) {
	p, err := utils.GetApiCall(sym)
	if p.Symbol == "" {
		return "", err
	}
	f_number, err := strconv.ParseFloat(compare_value, 32)
	log.Printf(fmt.Sprintf("%.2f", f_number))
	repo.DeleteNotification(user_id, sym, condition, f_number)
	repo.InsertNotification(user_id, sym, condition, f_number)
	return p.Symbol, err
}

func GetNotificationsByUser(user_id string) string {
	var msg_notifications string
	notifications := repo.GetUserNotifications(user_id)
	if len(notifications) == 0 {
		return "Notifications not found"
	}

	for _, notif := range notifications {
		msg_notifications += fmt.Sprintf("\n----\nSymbol: %s\nCondition: %s\nCompare Value: %f\n----", notif.Symbol, notif.CompareCondition, notif.CompareValue)
	}
	return msg_notifications
}

func CheckNotification(notification repo.CryptoNotification) (string, error) {
	p, err := utils.GetApiCall(notification.Symbol)
	if p.Symbol == "" {
		return "", err
	}
	last, err := strconv.ParseFloat(p.Last, 32)
	var alert_msg string
	if notification.CompareCondition == ">" && last > notification.CompareValue {
		alert_msg = fmt.Sprintf("Current value of %s is higher than %s (Current: %s)",
			notification.Symbol, fmt.Sprintf("%.2f", notification.CompareValue), p.Last)
	} else if notification.CompareCondition == "<" && last < notification.CompareValue {
		alert_msg = fmt.Sprintf("Current value of %s is lower than %s (Current: %s)",
			notification.Symbol, fmt.Sprintf("%.2f", notification.CompareValue), p.Last)
	}
	return alert_msg, err
}

func DeleteNotification(user_id string, sym string, condition string, compare_value string) string {
	f_number, _ := strconv.ParseFloat(compare_value, 32)
	affected_rows, _ := repo.DeleteNotification(user_id, sym, condition, f_number)
	return fmt.Sprintf("Notifications removed: %d", affected_rows)
}
