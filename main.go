package main

import (
	"log"

	"github.com/hyperxpizza/crud-generator/flags"
	"github.com/hyperxpizza/crud-generator/generator/general"
	"github.com/hyperxpizza/crud-generator/generator/sqldb"
)

var config flags.Config

func init() {
	config = flags.InitFlags()
}

func main() {
	//Prepare new directory and set up boilerplate
	err := general.SetUpBoilerplate(config.OutputDir, config.Module)

	sqlData, err := sqldb.LoadSQLfile(config.SchemaPath)
	if err != nil {
		log.Fatal(err)
	}

}
