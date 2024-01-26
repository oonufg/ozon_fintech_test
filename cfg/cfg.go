package cfg

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Cfg struct {
	DATABASE          string
	DATABASE_USER     string
	DATABASE_PASSWORD string
	DATABASE_HOST     string
	DATABASE_PORT     string

	PERSISTANCE_MODE string
}

func LoadCFG() *Cfg {
	workingDirrectory, _ := os.Getwd()
	error := godotenv.Load()
	if error != nil {
		log.Fatalf("Failed to read env file - \n Dirrectory %s", workingDirrectory)
	}

	return &Cfg{
		DATABASE:          pullEnvVariable("PG_DB_NAME"),
		DATABASE_USER:     pullEnvVariable("PG_USER"),
		DATABASE_PASSWORD: pullEnvVariable("PG_PASSWORD"),
		DATABASE_HOST:     pullEnvVariable("PG_HOST"),
		DATABASE_PORT:     pullEnvVariable("PG_PORT"),
		PERSISTANCE_MODE:  os.Args[0],
	}
}

func pullEnvVariable(key string) string {
	value, isExist := os.LookupEnv(key)
	if !isExist {
		log.Fatalf("Environment variable %s not found in env", key)
	}
	return value
}
