package database

import (
	"context"
	"fmt"
	"github.com/alecthomas/log4go"
	"github.com/jackc/pgx"
	"sync"
)

var (
	conn *pgx.Conn

	log *log4go.Logger

	sqlSelect = "SELECT rquid, message FROM public.notification_data nd WHERE need_notification is true AND sent IS NOT true"

	sqlChangeSent = "UPDATE notification_data SET sent = true WHERE rquid = $1"

	sqlInsertMessages = "INSERT INTO public.notification_data (rquid, message, need_notification)" +
		"VALUES ($1, $2, $3)"

	mutex = &sync.Mutex{}
)

// initialize DAO
func InitDao(connection *pgx.Conn, logger *log4go.Logger) {
	conn = connection
	log = logger
}

// select messages which need to send
func SelectMessages() map[string]string {
	mutex.Lock()
	log.Info("Start searching for notification messages")
	result := make(map[string]string)
	rows, err := conn.Query(context.Background(), sqlSelect)
	mutex.Unlock()
	if err != nil {
		_ = log.Error(fmt.Sprintf("Error while executing query: %v\n", err))
	}
	defer rows.Close()
	for rows.Next() {
		var rquid, message string
		err := rows.Scan(&rquid, &message)
		if err != nil {
			_ = log.Error(fmt.Sprintf("Error while scanning row: %v\n", err))
		} else {
			result[rquid] = message
		}
	}
	log.Info(fmt.Sprintf("Found %d messages that need notification: %v\n", len(result), result))
	return result
}

// change status of sent message
func ChangeStatusSent(rquid string) bool {
	mutex.Lock()
	log.Info(fmt.Sprintf("Going to set sent=true to rquid = %s", rquid))
	_, err := conn.Exec(context.Background(), sqlChangeSent, rquid)
	mutex.Unlock()
	if err != nil {
		_ = log.Error(fmt.Sprintf("Error while changing 'sent' for rquid=%s\n%v\n", rquid, err))
		return false
	} else {
		log.Info(fmt.Sprintf("Successfuly set 'sent'=true for rquid=%s", rquid))
		return true
	}
}

// insert new message
func InsertMessage(rquid, message string, needNotification bool) {
	mutex.Lock()
	_, err := conn.Exec(context.Background(), sqlInsertMessages, rquid, message, needNotification)
	mutex.Unlock()
	if err != nil {
		_ = log.Error(fmt.Sprintf("Error while inserting new message: %s\n", err))
	} else {
		log.Info(fmt.Sprintf("Inserted new message rquid=%s", rquid))
	}
}
