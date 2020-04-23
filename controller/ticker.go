package controller

import (
	"fmt"
	"net/http"

	"github.com/mateuspmello/stocks/model"
)

//NomeTicker nome do ticker
func NomeTicker(w http.ResponseWriter, r *http.Request) {
	nome := "VVAR3.SA"
	ticker := model.Stock{}

	ticker.Ticker = nome
	if err := Modelos.ExecuteTemplate(w, "ticker.html", ticker); err != nil {
		http.Error(w, "Erro ao fazer a renderização da pagina", http.StatusInternalServerError)
		fmt.Println("[Ticker] Erro na execucao do modelo: ", err.Error())
	}
}
