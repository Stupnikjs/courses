package mailer

import (
	"log"
	"net/smtp"
)

func Mail() {
	// Define the email sender and recipient
	from := "stupnikjs@gmail.com"
	password := "Phoenix190129023903." // Use application-specific password if required
	to := []string{"n.boudier.ph@gmail.com"}

	// Define the SMTP server and port
	smtpHost := "smtp.gmail.com" // e.g., smtp.gmail.com for Gmail
	smtpPort := "465"

	// Set up the message content
	subject := "Subject: Test Email from Go\n"
	body := "This is a test email sent from Go using the net/smtp package."
	message := []byte(subject + "\n" + body)

	// Authentication for the SMTP server
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send the email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}

	log.Println("Email sent successfully!")
}
