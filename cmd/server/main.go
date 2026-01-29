package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/luanaands/multithreading-golang/configs"
	_ "github.com/luanaands/multithreading-golang/docs"
	"github.com/luanaands/multithreading-golang/internal/infra/service"
	"github.com/luanaands/multithreading-golang/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Desafio CEP API - golang
// @version 1.0
// @description API para consulta de CEP em paralelo (BrasilAPI + ViaCEP)
// @termsOfService http://swagger.io/terms/

// @contact.name Luana Andrade
// @contact.email luanaands@gmail.com

// @host localhost:8000
// @basePath /
func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("BrasilApHost", configs.ApiHost))
	r.Use(middleware.WithValue("ViaCepHost", configs.OtherApiHost))

	var cepService service.CepInterface = service.NewCepService()
	handler := handlers.NewCepHandler(cepService)

	r.Get("/cep", handler.GetCep)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)
}
