package flags

import "flag"

//Config struct keeps track of cli flags
type Config struct {
	SchemaPath string
	OutputDir  string
	Host       string
	ServerPort int
	Module     string
}

//InitFlags sets cli flags
func InitFlags() Config {
	var config Config

	flag.StringVar(&config.SchemaPath, "sqlSchema", "./schema.sql", "Path to sql schema file. Needs to be .sql")
	flag.StringVar(&config.OutputDir, "output", "./generated", "Path to output generated files")
	flag.StringVar(&config.Host, "host", "localhost", "Database host. Default: localhost")
	flag.IntVar(&config.ServerPort, "server_port", 8888, "If api set to true, provide server port")
	flag.StringVar(&config.Module, "module", "", "Go package module. Example: github.com/hyperxpizza")

	flag.Parse()

	return config
}
