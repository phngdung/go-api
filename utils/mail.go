package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/smtp"
	"os"
)

type MailServerInfo struct {
	MailFrom string
	MailPass string
	MailHost string
	MailPort string
}

var (
	mail *MailServerInfo
)

func GetMailOption() *MailServerInfo {
	if mail != nil {
		return mail
	}
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&mail)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return mail
}
func Send(toMail string, subject string, body string) (bool, string) {

	mailConfig := GetMailOption()

	from := mailConfig.MailFrom
	password := mailConfig.MailPass

	to := []string{
		toMail,
	}

	smtpHost := mailConfig.MailHost
	smtpPort := mailConfig.MailPort

	message := []byte("From: Demo<" + from + ">\r\n" +
		"Subject:" + subject + "\r\n" +
		"\r\n" +
		body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return false, err.Error()
	}
	fmt.Println("Email Sent Successfully!")
	return true, "OK"
}
