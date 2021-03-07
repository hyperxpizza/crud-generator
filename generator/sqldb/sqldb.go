package sqldb

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/hyperxpizza/crud-generator/templates"
)

//LoadSQLfile reads provided sql file
func LoadSQLfile(path string) (string, error) {
	log.Println("Loading sql schema...")

	if path == "" {
		return "", fmt.Errorf("Invalid filepath")
	}

	if filepath.Ext(path) != ".sql" {
		return "", fmt.Errorf("Invalid file extension")
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("reading sql schema failed: %v\n", err)
		return "", err
	}

	log.Println("Sql schema loaded successfully!")

	return string(data), nil
}

//SetUpDatabaseConnectionTemplate creates database.go file responsible for postgres connection
func SetUpDatabaseConnectionTemplate(dir, module string) error {
	log.Println()

	file, err := os.Create(fmt.Sprintf("%s/database/database.go", dir))
	if err != nil {
		log.Fatalf("os.Create")
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

	return nil
}
