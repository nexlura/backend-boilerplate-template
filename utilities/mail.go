package utilities

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
)

// smtpServer data to smtp server
type smtpServer struct {
	host string
	port string
}

// Address URI to smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

// Request struct
type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

type GmailStruct struct {
	RecipientEmail string
	RecipientName  string
	LinkToken      string
}

func GmailSendResetPasswordEmail(req GmailStruct) error {
	username := os.Getenv("GMAIL_USERNAME")
	password := os.Getenv("GMAIL_PASSWORD")
	//fmt.Println(username, password)

	from := "info@propati.xyz"
	to := []string{req.RecipientEmail}

	// smtp server configuration.
	gmailSmtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}

	//Authentication.
	auth := smtp.PlainAuth("", username, password, gmailSmtpServer.host)

	// set up the email template
	t, err := template.ParseFiles("public/templates/resetPasswordEmail.html")

	if err != nil {
		log.Fatal(err)
		println("Template Error::: ", err)
		return err
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Reset Password! \n%s\n\n", mimeHeaders)))
	t.Execute(&body, struct {
		Name string
		URL  string
	}{
		Name: req.RecipientName,
		URL:  fmt.Sprintf("https://app.propati.xyz/auth/password/reset?token=%s", req.LinkToken),
	})

	// Sending email.
	err = smtp.SendMail(gmailSmtpServer.Address(), auth, from, to, body.Bytes())

	fmt.Println(err)
	if err != nil {
		//log.Fatal(err)
		fmt.Println("Email err::: ", err)
		return err
	}

	fmt.Println("Reset email sent successfully")

	return nil
}
