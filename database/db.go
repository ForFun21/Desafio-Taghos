package database

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("pgx", "postgres://postgres:210420@localhost:5432/meu_banco")
	if err != nil {
		log.Fatal("Erro ao se conectar ao banco de dados:", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("Erro ao verificar a conexão com o banco de dados:", err)
	}
}
