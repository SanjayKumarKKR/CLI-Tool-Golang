package mailUtils

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendMail(from string, password string, to []string, subject string, message string) {
	// Configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	fmt.Println("from : ", from)
	fmt.Println("to : ", to)
	fmt.Println("subject : ", subject)
	fmt.Println("message : ", message)
	var msg string = "To: " + to[0] + "\r\n" + "Subject: " + subject + "\r\n" + message + "\r\n"

	//authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send actual message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}
