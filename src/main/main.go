package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"gopkg.in/macaron.v1"
	"github.com/go-macaron/binding"
)

type Config struct {
	Http  httpConfig `toml:"http"`
	DB    dbConfig `toml:"database"`
	Smtp  smtpConfig `toml:"smtp"`
	Email emailConfig `toml:"email"`
}

type httpConfig struct {
	Port int
  Redirect string
}

type dbConfig struct {
	ConnectionString string
}

type emailConfig struct {
	From     string
	To       []string
	Subject  string
	Template string
}

type smtpConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Ssl	bool
}

var conf Config

func handleForm(ctx *macaron.Context, form HttpForm) {
	if ba, err := BuildBusinessActivity(form); err != nil {
		log.Println(err)
	} else {
		if issueId, err := AddBusinessActivity(ba); err != nil {
			log.Println("Error adding business activity")
			log.Println(err)
		} else {
      id := *issueId
      log.Printf("Business activity %s has just added", id)
			if err := SendBusinessActivity(id, ba); err != nil {
				log.Println("Unable to send email")
				log.Println(err)
			} else {
        ctx.Redirect(conf.Http.Redirect + "?id=" + id)
        return
      }
		}
	}

  ctx.Redirect(conf.Http.Redirect)
}

func main() {

	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		log.Println("Error parsing config file")
		log.Panic(err)
	}

	// init mailer
	if err := InitMailer(conf.Smtp, conf.Email); err != nil {
		log.Println("Can't initialize mailer")
		log.Panic(err)
	}

	// init database connection
	InitDB(conf.DB.ConnectionString)
	defer CloseDB()

	m := macaron.Classic()
	m.Post("/send-form", binding.MultipartForm(HttpForm{}), handleForm)
	m.Run(conf.Http.Port)
}