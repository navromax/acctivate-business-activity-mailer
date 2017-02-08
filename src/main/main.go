package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"gopkg.in/macaron.v1"
	"github.com/go-macaron/binding"
  "flag"
)

type Config struct {
	Http  httpConfig  `toml:"http"`
	DB    dbConfig    `toml:"database"`
	Smtp  smtpConfig  `toml:"smtp"`
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

var configFileName = flag.String("config", "config.toml", "specifies config file to use")

var conf Config

func handleForm(ctx *macaron.Context, form HttpForm) {
	if ba, err := BuildBusinessActivity(form); err != nil {
		log.Println(err)
	} else {
		if issueId, err := AddBusinessActivity(ba); err != nil {
			log.Println("Error adding business activity. EMails won't be sent")
			log.Println(err)
		} else {
      id := *issueId
      log.Printf("Business activity %s has just added. Sending ...", id)
			if err := SendBusinessActivity(id, ba); err != nil {
				log.Println("Unable to send email")
				log.Println(err)
			} else {
        log.Printf("Business activity %s successfully sent", id)
        ctx.Redirect(conf.Http.Redirect + "?id=" + id)
        return
      }
		}
	}

  ctx.Redirect(conf.Http.Redirect)
}

func parseConfig() {
  if _, err := toml.DecodeFile(*configFileName, &conf); err != nil {
    log.Printf("Error parsing config file. %s", *configFileName)
    log.Panic(err)
  }
}

func main() {
  flag.Parse()
  parseConfig()

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