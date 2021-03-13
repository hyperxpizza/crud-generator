package sqldb

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hyperxpizza/crud-generator/templates"
)

//SetUpDatabaseConnectionTemplate creates database.go file responsible for postgres connection
func SetUpDatabaseConnectionTemplate(dir string) error {
	log.Println("Setting up database.go...")

	file, err := os.Create(fmt.Sprintf("%s/database/database.go", dir))
	if err != nil {
		log.Fatalf("os.Create database.go failed: %v\n", err)
		return err
	}

	defer file.Close()

	err = templates.DatabaseConnectionTemplate.Execute(file, struct {
		Timestamp time.Time
	}{
		Timestamp: time.Now(),
	})
	if err != nil {
		log.Fatalf("databaseConnectionTemplate.Execute failed: %v\n", err)
		return err
	}

	log.Println("Created database.go")

	defer file.Close()

	return nil
}

func parseSQL(data string) error {
	return nil
}
