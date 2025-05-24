package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/handler"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	AccountURL string `envconfig:"ACCOUNT_SERVICE_URL"`
	ProductURL string `envconfig:"PRODUCT_SERVICE_URL"`
	OrderURL   string `envconfig:"ORDER_SERVICE_URL"`
}

func main() {
	var cfg AppConfig

	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	// To call graph.go script
	s, err := NewGraphqlServer(cfg.AccountURL, cfg.ProductURL, cfg.OrderURL)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/graphql", handler.GraphQL(s.ToExecutableSchema()))
	http.Handle("/playground", playground.Handler("luisrojas17", "/grapql"))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
