# Acctivate! form sender

[![Build Status](https://travis-ci.org/andyglow/acctivate-business-activity-mailer.svg?branch=master)](https://travis-ci.org/andyglow/acctivate-business-activity-mailer)

This simple program do the simple things.
1. It hosts an html file with the form you want to handle
2. It is able to process submission if the form
3. The form will be passed to Acctivate! `spAddBusinessActivity` function
4. The notification will be sent via email to pre-configured recipients 

## Installation
### Linux
1. Copy `acctivate-business-activity-mailer` to your app folder (eg. `/opt/acctivate/apps/form-mailer`)
2. Put `config.toml` and `email.tpl` to the same dir.
3. Make necessary changes to `config.toml`

### Windows
1. Copy `acctivate-business-activity-mailer.exe` (`i386` / `x64`) to your app folder (eg. `C:\Acctuate\Apps\Form-Mailer`)
2. Put `config.toml` and `email.tpl` to the same dir.
3. Make necessary changes to `config.toml`

## Configuration
### Http
You can define port the server will be listening to
```toml
[http]
Port = 8080
```

### Database
You should specify database connection string
```toml
[database]
ConnectionString = "server=localhost;user id=sa;password=sa;database=ACCTivate$DB"
```

### SMTP
To be able to send notifications you should configure smtp server connection settings
```toml
[smtp]
Host = "localhost"
Port = 25
User = ""
Password = ""
Ssl = false
```

### Email
And the last. It is possible to specify email attributes like `From`, list of recipients and `Subject`.
```toml
[email]
From = "info@site.com"
To = [ "admin@site.com", "service@site.com" ]
Subject = "Notification. Activity ID:"
Template = "email.tpl"
```
You also need to specify template used to generate email body.

## Email template
Email body is a simple text, so emails send using `text/plain` content type.
Writing a template you should know that it is built with [[https://golang.org/pkg/text/template/]].
And following fields are presented and could be used within template:
- 	`ActivityType`
-  	`ActivityCode`
-  	`ActivityDescription`
-  	`AssignedTo`
-  	`Reference`
-  	`Reference2`
-  	`OpenedBy`
-  	`Contact`
-  	`Email`
-  	`AddressName`
-  	`Discussion`
-  	`Description`
-  	`AttachmentPath`
