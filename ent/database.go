package ent

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Database struct {
	*Client
}

// singleton
var database = newDatabase()

func newDatabase() *Database {
	entOptions := []Option{}

	env := os.Getenv("DEBUG")

	if env == "ON" {
		entOptions = append(entOptions, Debug())
	}

	dbConf := newDbConfig()

	datasource := fmt.Sprintf("host=%s port=%s user=%s dbname= %s password=%s sslmode=disable", dbConf.host, dbConf.port, dbConf.user, dbConf.name, dbConf.pass)

	client, err := Open(dbConf.driver, datasource, entOptions...)
	if err != nil {
		log.Fatal(err)
	}

	return &Database{client}
}

func DatabaseConnect() *Database {
	return database
}

type dbConfig struct {
	host   string
	port   string
	user   string
	name   string
	pass   string
	driver string
}

func newDbConfig() dbConfig {
	conf := new(dbConfig)

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	conf.host = host

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}
	conf.port = port

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "template"
	}
	conf.user = user

	name := os.Getenv("DB_NAME")
	if name == "" {
		name = "template"
	}
	conf.name = name

	pass := os.Getenv("DB_PASS")
	if pass == "" {
		pass = "template"
	}
	conf.pass = pass

	driver := os.Getenv("DB_DRIVER")
	if driver == "" {
		driver = "postgres"
	}
	conf.driver = driver

	return *conf
}
