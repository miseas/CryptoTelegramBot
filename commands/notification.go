package commands

import (
	"cryptoTelegramBot/repo"
	"cryptoTelegramBot/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func CreateNotification(user_id string, sym string, condition string, compare_value string) (string, error) {
	sym = strings.ToUpper(sym)
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

func GetAllNotifications() []repo.CryptoNotification {
	notifications := repo.GetAllNotifications()
	return notifications
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
		alert_msg = fmt.Sprintf("🔊Current value of %s is higher than %s (Current: %s)",
			notification.Symbol, fmt.Sprintf("%.2f", notification.CompareValue), p.Last)
	} else if notification.CompareCondition == "<" && last < notification.CompareValue {
		alert_msg = fmt.Sprintf("🔊Current value of %s is lower than %s (Current: %s)",
			notification.Symbol, fmt.Sprintf("%.2f", notification.CompareValue), p.Last)
	}
	return alert_msg, err
}

func DeleteNotification(user_id string, sym string, condition string, compare_value string) string {
	sym = strings.ToUpper(sym)
	f_number, _ := strconv.ParseFloat(compare_value, 32)
	affected_rows, _ := repo.DeleteNotification(user_id, sym, condition, f_number)
	return fmt.Sprintf("Notifications removed: %d", affected_rows)
}

func UpdateNotificationCounter(notification repo.CryptoNotification) bool {
	deleted := false
	notification.Counter = notification.Counter - 1
	if notification.Counter <= 0 {
		repo.DeleteNotification(notification.UserId, notification.Symbol, notification.CompareCondition, notification.CompareValue)
		log.Println("Notification deleted after counter set to 0...")
		deleted = true
	} else {
		repo.UpdateNotification(notification)
	}

	return deleted
}
