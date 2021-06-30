package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ola mundo!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar pagina de usuarios!"))
}

func main() {
	// HTTP e um protocolo de comunicacao - base da comunicacao web

	// Cliente (faz requisicao) - servidor (processa requisicao e envia resposta)

	// Request - response

	// Rotas
	// URI - identificador do recurso
	// Metodo - Accion do recurso, GET, POST, PUT, DELETE

	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
