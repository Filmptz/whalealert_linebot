package svc

import (
	"github.com/joho/godotenv"
    "log"
    "os"
)

func DotEnv(key string)(string,error){
	err := godotenv.Load()
  	if err != nil {
    	log.Fatal("Error loading .env file")
  	}

  return os.Getenv(key),nil
}