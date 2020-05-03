package job

import (
	"fmt"
	"github.com/alecthomas/log4go"
	"net/http"
	"notificator/notificator/internal/database"
	"strings"
)

var (
	log  *log4go.Logger
	url  string
	port string
)

// do notification job, find and send messages and change theis statuses
func DoJob() {
	notifications := database.SelectMessages()
	log.Info(fmt.Sprintf("Trying to send %d notifications", len(notifications)))
	count := 0
	for rquid, message := range notifications {
		count += sendNotificationAndChangeStatus(rquid, message)
	}
	log.Info(fmt.Sprintf("Sent and changed status for %d notifications, see ya", count))
}

// send message
func sendNotificationAndChangeStatus(rquid, message string) int {
	count := 0
	reader := strings.NewReader(fmt.Sprintf("rquid=%s\nmessage=%s", rquid, message))
	address := fmt.Sprintf("%s:%s", url, port)
	log.Info(fmt.Sprintf("Trying to send message to address=%s", address))
	_, err := http.Post(address, "text/html; charset=utf-8", reader)
	if err != nil {
		_ = log.Error(fmt.Sprintf("Error while sending POST request: %v\n", err))
	} else {
		count = changeStatus(rquid)
	}
	return count
}

// change status of message
func changeStatus(rquid string) int {
	count := 0
	sent := database.ChangeStatusSent(rquid)
	if sent {
		log.Info(fmt.Sprintf("Successfuly sent notification rquid = %s", rquid))
		count = 1
	} else {
		_ = log.Error(fmt.Sprintf("Cannot change 'sent' for rquid = %s", rquid))
	}
	return count
}

// initialize job
func InitJob(logger *log4go.Logger, urlToSend string, portToSend string) {
	log = logger
	url = urlToSend
	port = portToSend
}
