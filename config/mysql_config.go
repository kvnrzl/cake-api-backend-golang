package config

import "os"

var MYSQL_HOST = os.Getenv("MYSQL_HOST")
var MYSQL_PORT = os.Getenv("MYSQL_PORT")
var MYSQL_USER = os.Getenv("MYSQL_USER")
var MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")
var MYSQL_DATABASE = os.Getenv("MYSQL_DATABASE")

// var MYSQL_HOST = "localhost"
// var MYSQL_PORT = "3306"
// var MYSQL_USER = "root"
// var MYSQL_PASSWORD = ""
// var MYSQL_DATABASE = "technical_test_privy_db"
