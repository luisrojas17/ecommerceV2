package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/luisrojas17/ecommerceV2/accounts"

	"github.com/tinrab/retry"
)

type Config struct {
	DatabseURL string `envconfig:"DATABASE_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var r accounts.Respository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {

		r, err = accounts.NewPostgresRespository(cfg.DatabseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()
	log.Println("Listening on port 8080...")

	s := accounts.NewService(r)

	log.Fatal(accounts.ListenGRPC(s, 8080))

}
