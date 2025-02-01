package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kelseyhightower/envconfig"
)

type Appconfig struct {
	AccountURL string `json:"ACCOUNT_SERVICE_URL"`
	CatalogURL string `json:"CATALOG_SERVICE_URL"`
	OrderURL   string `json:"ORDER_SERVICE_URL"`
}

func main() {
	var cfg Appconfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	s, err := NewGraphQLServer(cfg.AccountURL, cfg.CatalogURL, cfg.OrderURL)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/graphql", handler.New(s.ToExecutableSchema()))
	http.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))

	log.Println("connect to http://localhost:8080/playground for GraphQL playground")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
