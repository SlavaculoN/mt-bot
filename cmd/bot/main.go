package main

import (
	"log"

	"mt-bot/internal/config"
	"mt-bot/internal/service"
	"mt-bot/internal/storage"
	"mt-bot/internal/telegram"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	sessionStorage := storage.NewMemoryStorage()

	requestService := service.NewRequestService(
		sessionStorage,
	)

	bot, err := telegram.NewBot(
		cfg.BotToken,
		requestService,
		cfg.WorkChatID,
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := bot.Run(); err != nil {
		log.Fatal(err)
	}
}
