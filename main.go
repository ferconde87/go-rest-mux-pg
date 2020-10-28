package main

import "os"

//environment variables
// var APP_DB_USERNAME = "postgres"
// var APP_DB_PASSWORD = ""
// var APP_DB_NAME = "postgres"

func main() {
	a := App{}

	os.Setenv("APP_DB_USERNAME", "postgres")
	os.Setenv("APP_DB_PASSWORD", "")
	os.Setenv("APP_DB_NAME", "postgres")


	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8010")
}