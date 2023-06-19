package utils

import (
	"crypto/tls"
	"fmt"
	"os"
	"strconv"

	gomail "github.com/go-gomail/gomail"
)

var (
	port int
)

type Mail struct {
	smtpHost    string
	port        string
	senderEmail string
	password    string
}

type msgData struct {
	Username string
	Token    string
}

var VerifyMessage = `

	<h2 style="text-align: center"> Auth </h2><br>
	
	<h2 style="text-align:center"> Verify Your Email </h2>
	
	<p> Hey %s click the link below to finish verifying your email address </p><br>
	
	<p style="text-align:center; height:30px; width:200px; background-color:blue; border-radius: 5px; padding:10px;"><a href="http://127.0.0.1:8081/api/auth/verify/%s/%s" style="text-decoration:none; color:white; text-align:center;"> Confirm email </a></p>
	<br>
`

func NewMailStruct() Mail {
	return Mail{
		smtpHost:    os.Getenv("SMTP_HOST"),
		port:        os.Getenv("SMTP_PORT"),
		senderEmail: os.Getenv("SENDER_EMAIL"),
		password:    os.Getenv("SENDER_PASSWORD"),
	}
}

func NewMessage(receiver, subject, message, username, token string) (Mail, *gomail.Message) {
	ms := NewMailStruct()

	mail := gomail.NewMessage()
	mail.SetHeader("From", ms.senderEmail)
	mail.SetHeader("To", receiver)
	mail.SetHeader("Subject", subject)

	msg := fmt.Sprintf(message, username, username, token)

	mail.SetBody("text/html", msg)

	return ms, mail

}

func NewMail(receiver, subject, message, username, token string) error {

	ms, msg := NewMessage(receiver, subject, message, username, token)

	port, err := strconv.Atoi(ms.port)
	if err != nil {
		return err
	}

	dialer := gomail.NewDialer(ms.smtpHost, port, ms.senderEmail, ms.password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err = dialer.DialAndSend(msg)
	if err != nil {
		return err
	}

	return nil
}
