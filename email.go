package main

import (
	"bytes"
	"log/slog"
	"text/template"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func SendEmail(subject, htmlBody string) error {
	// Create a new AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ca-central-1"), // Specify your AWS region
	})
	if err != nil {
		panic(err)
	}

	// Create an SES client
	svc := ses.New(sess)

	// Construct the email input parameters
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{aws.String("oren.mazor@protonmail.com"), aws.String("LeeEva@gmail.com")},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Data: aws.String(htmlBody),
				},
			},
			Subject: &ses.Content{
				Data: aws.String(subject),
			},
		},
		Source: aws.String("costco-monitor@orenmazor.com"), // Must be a verified email address in your SES account
	}

	// Send the email
	foo, err := svc.SendEmail(input)
	if err != nil {
		panic(err)
	}

	slog.Info(foo.String())

	return nil
}

func GenerateEmailHTML(results map[string][]CostcoResult) string {
	templateFile := "templates/email.tpl"
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, results)

	return buf.String()
}
