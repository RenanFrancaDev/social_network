package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConectionDB = ""
	Port              = 0
	SecretKey         []byte
)

func HandleConfig() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatalf("[config] [msg: Error to load .env file %v]", err)
	}

	//ParseInt string -> int
	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		println("[config] [msg: error with Port]")
	}

	StringConectionDB = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
