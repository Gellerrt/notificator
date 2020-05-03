package job

import (
	"math/rand"
	"notificator/notificator/internal/database"
	"time"
)

var letters = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// generate new message
func AddMessage() {
	log.Info("Trying to insert new message")
	rand.Seed(time.Now().UnixNano())
	rquid := randomSequence(15)
	message := randomSequence(30)
	send := randBool()
	database.InsertMessage(rquid, message, send)
}

// generate random sequence
func randomSequence(n int) string {
	result := make([]rune, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// generate random bool
func randBool() bool {
	return rand.Float32() < 0.5
}
