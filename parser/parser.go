package parser

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

type Query struct {
	Type      string
	TableName string
	Fields    map[string]interface{}
}

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

func Parse(sql string) (*Query, error) {
	return nil, nil
}
