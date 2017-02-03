package main

import (
	_ "github.com/denisenkom/go-mssqldb"
	"database/sql"
	"log"
  "fmt"
)

var db *sql.DB

func AddBusinessActivity(ba *BusinessActivity) (*string, error) {
  if rows, err := db.Query(`
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
    ba.Reference2); err != nil {
    return nil, err
  } else {
    defer rows.Close()

    if rows.Next() {
      var issueId string
      if err := GetColumn(rows, "IssueID", &issueId); err != nil {
        return nil, err
      } else {
        return &issueId, nil
      }
    } else {
      return nil, fmt.Errorf("Result set hasn't been recieved")
    }
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