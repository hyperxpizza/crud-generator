package flags

import "flag"

//Config struct keeps track of cli flags
type Config struct {
	SchemaPath string
	OutputDir  string
	Host       string
	Port       int
	API        bool
	ServerPort int
	Module     string
}

//InitFlags sets cli flags
func InitFlags() Config {
	var config Config

	flag.StringVar(&config.SchemaPath, "sqlSchema", "", "Path to sql schema file. Needs to be .sql")
	flag.StringVar(&config.OutputDir, "output", "", "Path to output generated files")
	flag.StringVar(&config.Host, "host", "localhost", "Database host. Default: localhost")
	flag.IntVar(&config.Port, "port", 5432, "Database port. Default: 5432")
	flag.BoolVar(&config.API, "api", false, "If set to true, generates gin rest api.")
	flag.IntVar(&config.ServerPort, "server_port", 8888, "If api set to true, provide server port")
	flag.StringVar(&config.Module, "module", "", "Go package module. Example: github.com/hyperxpizza")

	flag.Parse()

	return config
}
