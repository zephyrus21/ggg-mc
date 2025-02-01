package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
	"github.com/zephyrus21/ggg-mc/account"
)

type Config struct {
	DBURL string `envconfig:"DB_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var r account.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = account.NewPostgresRepository(cfg.DBURL)
		if err != nil {
			log.Println("failed to connect to db:", err)
		}

		return
	})

	defer r.Close()
	log.Println("Listening on :8080")

	s := account.NewService(r)
	log.Fatal(account.ServeGRPC(s, 8080))
}
