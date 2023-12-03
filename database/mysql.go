package database

// func init digunakan seperti __construct() kalau di php

var connection string

func init() {
	connection = "MySQL"
}

func GetConnection() string {
	return connection
}
