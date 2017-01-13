package main

import (
	"gopkg.in/gomail.v2"
	"text/template"
	"bytes"
	"net/textproto"
)

var tpl *template.Template
var dialer *gomail.Dialer
var emailAttributes *emailConfig

func InitMailer(smtp smtpConfig, attrs emailConfig) error {
	var err error
	tpl, err = template.New(attrs.Template).ParseFiles(attrs.Template)
	if err != nil {
		return err
	}

	if textproto.TrimString(smtp.User) == "" {
		dialer = &gomail.Dialer{Host: smtp.Host, Port: smtp.Port}
	} else {
		dialer = gomail.NewDialer(smtp.Host, smtp.Port, smtp.User, smtp.Password)
	}
	dialer.SSL = smtp.Ssl
	emailAttributes = &attrs
	return nil
}

func SendBusinessActivity(issueId string, ba *BusinessActivity) error {
	buffer := bytes.NewBufferString("")
	if err := tpl.Execute(buffer, *ba); err != nil {
		return err
	}

	body := buffer.String()

	m := gomail.NewMessage()
	m.SetHeader("From", emailAttributes.From)
	m.SetHeader("To", emailAttributes.To...)
	m.SetHeader("Subject", emailAttributes.Subject + issueId)
	m.SetBody("text/plain", body)
	if ba.AttachmentPath != nil {
		m.Attach(*ba.AttachmentPath)
	}

	return dialer.DialAndSend(m)
}