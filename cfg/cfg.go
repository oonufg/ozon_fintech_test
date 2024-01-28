package cfg

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Cfg struct {
	DATABASE           string
	DATABASE_USER      string
	DATABASE_PASSWORD  string
	DATABASE_HOST      string
	DATABASE_PORT      string
	GRPC_PORT          string
	GRPC_ADDR          string
	PERSISTANCE_MODE   string
	HTTP_GATEWAY_ADDRE string
	HTTP_GATEWAY_PORT  string
}

func LoadCFG() *Cfg {
	workingDirrectory, _ := os.Getwd()
	error := godotenv.Load()
	if error != nil {
		log.Fatalf("Failed to read env file - \n Dirrectory %s", workingDirrectory)
	}

	return &Cfg{
		DATABASE:           pullEnvVariable("PG_DB_NAME"),
		DATABASE_USER:      pullEnvVariable("PG_USER"),
		DATABASE_PASSWORD:  pullEnvVariable("PG_PASSWORD"),
		DATABASE_HOST:      pullEnvVariable("PG_HOST"),
		DATABASE_PORT:      pullEnvVariable("PG_PORT"),
		GRPC_PORT:          pullEnvVariable("GRPC_PORT"),
		GRPC_ADDR:          pullEnvVariable("GRPC_ADDR"),
		HTTP_GATEWAY_ADDRE: pullEnvVariable("HTTP_GATEWAY_ADDRE"),
		HTTP_GATEWAY_PORT:  pullEnvVariable("HTTP_GATEWAY_PORT"),
		PERSISTANCE_MODE:   pullPersistanceType(),
	}
}

func pullEnvVariable(key string) string {
	value, isExist := os.LookupEnv(key)
	if !isExist {
		log.Fatalf("Environment variable %s not found in env", key)
	}
	return value
}

func pullPersistanceType() string {
	if len(os.Args) < 2 {
		return ""
	} else {
		return os.Args[1]
	}
}
