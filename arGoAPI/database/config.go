package database

import ("fmt")

//config to maintain DB configuration properties
type Config struct {
	ServerName string
	User string
	Password string
	DB string
}

var GetConnectionString = func(config Config) string {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true",
						config.User, config.Password, config.ServerName, config.DB)

	return connectionString					
}