package service

import (
	"fmt"
	"github.com/anjush-bhargavan/go_trade_notification/pkg/config"
	"log"

	// "strings"

	"gopkg.in/gomail.v2"
)

type Messages struct {
	Username string
	Email    string
	Messages string
	Subject  string
}

// SendEmail function send the generated email
func SendEmail(cnfg *config.Config, details Messages) error {
	sender := cnfg.Email
	password := cnfg.Password

	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", details.Email)
	fmt.Println(details.Email)
	m.SetHeader("Subject", details.Subject)
	m.SetBody("text/plain", details.Messages)

	d := gomail.NewDialer("smtp.gmail.com", 587, sender, password)
	if err := d.DialAndSend(m); err != nil {
		log.Printf("Could not send mail %v", err)
		return err
	} else {
		log.Printf("Email Sent Successfully")
	}
	return nil
}
