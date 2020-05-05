package main

import (
	"fmt"
	"net/http"

	"github.com/mateuspmello/stocks/controller"
	"github.com/mateuspmello/stocks/repo"
)

func main() {

	err := repo.AbreConexaoComBancoDeDadosSQL()
	if err != nil {
		fmt.Println("Erro ao abrir a conexão com o banco de dados: ", err.Error())
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá mundo")
	})
	http.HandleFunc("/funcao", controller.Funcao)
	http.HandleFunc("/ticker/nome", controller.NomeTicker)
	fmt.Println("Servidor inicializado")
	http.ListenAndServe(":8181", nil)

	// cliente := &http.Client{}

	// stock := model.Stock{}
	// stock.Ticker = "VVAR3.SA"
	// stock.Price = 4.42

	//iGET
	// request, err = http.NewRequest("GET", "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=VVAR3.SA&apikey=IHFYQDZP1LSOW6N0", nil)
	// if err != nil {
	// 	fmt.Println("erro na requisiçãi ao site; Erro:", err.Error())
	// 	return
	// }
	// // request.SetBasicAuth("mateus", "mateus")
	// resposta, err = cliente.Do(request)
	// defer resposta.Body.Close()

	// if resposta.StatusCode == 200 {
	// 	corpo, err := ioutil.ReadAll(resposta.Body)
	// 	if err != nil {
	// 		fmt.Println("erro ao ler o conteudo do site; Erro:", err.Error())
	// 		return
	// 	}
	// 	fmt.Println(string(corpo))
	// }
	//fGET
}
