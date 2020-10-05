package utils

import "github.com/mailgun/mailgun-go"

func SendEmailToMe(subject string, body string) error {
	config := GetConfig().Mailgun

	mg := mailgun.NewMailgun(config.Domain, config.ApiKey)

	message := mg.NewMessage(config.Sender, subject, body, config.Reciver)

	_, _, err := mg.Send(message)
	return err
}
