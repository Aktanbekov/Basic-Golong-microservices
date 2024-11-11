package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/handler"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	AccountURL	string `envconfig:"ACCOUNT_SERVICES_URL"`
	CatalogURL	string `envconfig:"CATALOG_SERVICES_URL"`
	OrderURL	string `envconfig:"ORDER_SERVICES_URL"`
}


func main() {
	var cfg AppConfig
	err :=envconfig.Process("", &cfg)
	if err != nil{
		log.Fatal(err)
	}

	s, err := NewGraphQLServer(cfg.AccountURL, cfg.CatalogURL, cfg.OrderURL)
	if err != nil {
		log.Fatal(err)
	} 

	http.Handle("/graphql", handler.GraphQL(s.ToExecutableSchema()))
	http.Handle("/playground", playground.Handler("aktai", "/graphql"))

	log.Fatal(http.ListenAndServe(":8080", nil))
}