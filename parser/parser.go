package parser

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/hyperxpizza/crud-generator/helpers"
)

var fieldTypes = []string{
	"varchar",
	"integer",
	"text",
	"boolean",
}

type Data struct {
	Name   string
	Fields []Field
}

type Field struct {
	Name string
	Type string
	Null bool
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

func Parse(sql string) ([]*Data, error) {
	var data []*Data

	a1 := strings.Split(sql, ";")
	for _, v := range a1 {
		var d Data

		//Clear string of empty lines
		v1 := strings.Replace(v, "\n\n", "", -1)
		//Get Name
		d.Name = helpers.Between(v1, "create table ", " (")
		if d.Name == "" {
			continue
		}

		d.Fields = getFields(v1, d.Name)

		data = append(data, &d)
	}

	return data, nil
}

func getFields(s, name string) []Field {
	var fields []Field
	s = strings.ToLower(s)

	d := strings.Replace(s, fmt.Sprintf("create table %s (", name), "", -1)
	arr := strings.Split(d, ",")

	for _, v := range arr {
		var field Field
		field.Null = false

		//Get name
		field.Name = helpers.Between(v, "", " ")
		//Get type
		t := helpers.CheckIfContains(v, fieldTypes)
		field.Type = getType(t)

		//check if not null
		if strings.Contains(v, "not null") {
			field.Null = true
		}

		fields = append(fields, field)
	}

	return fields
}

func getType(s string) string {
	switch s {
	case "varchar":
		return "string"
	case "text":
		return "string"
	case "integer":
		return "int"
	case "boolean":
		return "bool"
	}
}
