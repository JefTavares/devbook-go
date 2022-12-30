package main

import (
	"api/src/config"
	"api/src/logs"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

//Gera a secret_key
// func init() {
// 	chave := make([]byte, 64)
// 	if _, err := rand.Read(chave); err != nil {
// 		log.Fatal(err)
// 	}
// 	stringBase64 := base64.StdEncoding.EncodeToString(chave)
// 	fmt.Println(stringBase64)
// }

func main() {

	//logs.Write("olá mundo!", "./ola.log")
	logs.Write("olá mundo!", "")

	config.Carregar()
	r := router.Gerar()

	fmt.Printf("Escutando na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
