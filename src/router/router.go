package router

import (
	"github.com/gorilla/mux"
	"api/src/router/rotas"
)

// Vai retornar um router com as rotas 
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
