package repo

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//Db armaze conexão com bd
var Db *sqlx.DB

//AbreConexaoComBancoDeDadosSQL abre a conexão com o banco de dados
func AbreConexaoComBancoDeDadosSQL() (err error) {
	err = nil
	Db, err := sqlx.Open("mysql", "root@tcp(127.0.0.1)/portifolio")
	if err != nil {
		return
	}
	err = Db.Ping()
	if err != nil {
		return
	}
	return
}
