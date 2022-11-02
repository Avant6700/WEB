package main

import (
	"awesomeProject1/internal/app/config"
	"awesomeProject1/internal/pkg/app"
	"context"
	log "github.com/sirupsen/logrus"
	"os"
)

// @title Comics-shop
// @version 1.0
// @description Store with actual comics

// @contact.name Alexey Fotchin
// @contact.url https://vk.com/id250446192
// @contact.email fotchin02@mail.ru

// @host 127.0.0.1:8080
// @schemes http https
// @BasePath /

func main() {
	log.Println("app start")

	ctx := context.Background()

	_, err := config.NewConfig(ctx) //Config пока не используется
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("can't init config")
		os.Exit(2)
	}

	application, err := app.New(ctx)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("can`t create app")
		os.Exit(2)
	}

	err = application.Run(ctx)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("can`t run app")
		os.Exit(2)
	}

	log.Println("app terminated")
}
