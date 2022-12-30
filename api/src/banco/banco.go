package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver
)

// Abre a conexão com o banco de dados e a retorna
func Conectar() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StringDeConexaoBanco)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
