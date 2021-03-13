package main

import (
	"fmt"
	"log"

	"github.com/hyperxpizza/crud-generator/flags"
	"github.com/hyperxpizza/crud-generator/generator/general"
	"github.com/hyperxpizza/crud-generator/generator/sqldb"
	"github.com/hyperxpizza/crud-generator/parser"
)

var config flags.Config

func init() {
	config = flags.InitFlags()
}

func main() {
	//Prepare new directory and set up boilerplate

	err := general.SetUpBoilerplate(config.OutputDir, config.Module)
	if err != nil {
		log.Fatalf("Setting up boilerplate failed: %v", err)
	}

	sqlData, err := parser.LoadSQLfile(config.SchemaPath)
	if err != nil {
		log.Fatal(err)
	}

	err = sqldb.SetUpDatabaseConnectionTemplate(config.OutputDir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sqlData)

}
