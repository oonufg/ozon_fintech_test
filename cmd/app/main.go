package main

import (
	_ "database/sql"
	"fmt"
	"ozon_fintech_test/cfg"
	_ "ozon_fintech_test/internal/persistance"

	_ "github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg := cfg.LoadCFG()
	fmt.Println(cfg)
}
