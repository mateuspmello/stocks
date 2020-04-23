package controller

import (
	"fmt"
	"net/http"
)

//Funcao manipulador
func Funcao(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Olá mundo da funcao")
}
