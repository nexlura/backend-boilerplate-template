package utilities

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/backend-boilerplate-template/utilities/startup"
)

const (
	ErrParseJSON         = "cannot parse JSON"
	ErrInternalServer    = "internal Server Error"
	ErrProcessingRequest = "error processing request"
)

func ReadQuery(filePath string) (string, error) {
	query, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading SQL file '%s': %v", filePath, err)
		return "", errors.New("error processing request")
	}
	return string(query), nil
}

func ExecuteQuery(db *sql.DB, query string, args ...interface{}) error {
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Printf("Error preparing query: %v", err)
		return errors.New("error processing request")
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return errors.New("error processing request")
	}

	return nil
}

func ReadAndExecuteQuery(db startup.Database, filePath string, args ...interface{}) error {
	query, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading SQL file '%s': %v", filePath, err)
		return errors.New("error processing request")
	}

	_, err = db.Conn.Exec(context.Background(), string(query), args...)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return errors.New("error processing request")
	}

	return nil
}
