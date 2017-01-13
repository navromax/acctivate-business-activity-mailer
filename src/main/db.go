package main

import (
	_ "github.com/denisenkom/go-mssqldb"
	"database/sql"
	"log"
)

var db *sql.DB

func AddBusinessActivity(ba *BusinessActivity) (*string, error) {
  var issueId string
  if err := db.QueryRow(`
		EXECUTE spAddBusinessActivity
			@Type = ?1,
			@Code = ?2,
			@AssignedTo = ?3,
			@Contact = ?4,
			@Email = ?5,
			@AddressName = ?6,
			@OpenedBy = ?7,
			@Description = ?8,
			@Discussion = ?9,
			@Reference = ?10,
			@Reference2 = ?11`,
    ba.ActivityType,
    ba.ActivityCode,
    ba.AssignedTo,
    ba.Contact,
    ba.Email,
    ba.AddressName,
    ba.OpenedBy,
    ba.Description,
    ba.Discussion,
    ba.Reference,
    ba.Reference2).Scan(&issueId); err != nil {
    return nil, err
  } else {
    return &issueId, nil
  }
}


func InitDB(connString string) {
	var err error
	db, err = sql.Open("mssql", connString)



	if err != nil {
    log.Println("Error connecting to database")
    log.Printf("Database connection string: %s", connString)
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
    log.Println("Login failed")
    log.Printf("Database connection string: %s", connString)
		log.Panic(err)
	}
}

func CloseDB() {
	db.Close()
}