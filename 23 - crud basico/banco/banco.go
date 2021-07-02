package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Drover de conexao com o MySQL
)

// Conectar abre a conexao com o banco de dados
func Conectar() (*sql.DB, error) {
	urlConexao := "erbalo:erbalo123@/learning-go?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", urlConexao)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
