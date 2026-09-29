package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken   string
	WorkChatID int64
}

func Load() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("load env: %w", &err)
	}

	token := strings.TrimSpace(os.Getenv("BOT_TOKEN"))
	if token == "" {
		return Config{}, fmt.Errorf("BOT_TOKEN is not set")
	}

	workChatIDValue := strings.TrimSpace(os.Getenv("WORK_CHAT_ID"))
	if workChatIDValue == "" {
		return Config{}, fmt.Errorf("WORK_CHAT_ID is not set")
	}

	workChatID, err := strconv.ParseInt(workChatIDValue, 10, 64)
	if err != nil {
		return Config{}, fmt.Errorf("parse WORK_CHAT_ID: %w", err)
	}

	return Config{
		BotToken:   token,
		WorkChatID: workChatID,
	}, nil
}
