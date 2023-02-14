package main

import (
	"log"

	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/app"
	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/config"
)

func main() {
	cfg, err := config.NewConfig("config.yml", "env.example")
	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
