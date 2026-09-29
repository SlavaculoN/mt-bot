package telegram

import (
	"fmt"
	"log"

	"mt-bot/internal/domain"
	"mt-bot/internal/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api            *tgbotapi.BotAPI
	requestService *service.RequestService
	workChatID     int64
}

func NewBot(
	token string,
	requestService *service.RequestService,
	workChatID int64,
) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	log.Println("Бот запущен:", api.Self.UserName)

	return &Bot{
		api:            api,
		requestService: requestService,
		workChatID:     workChatID,
	}, nil
}

func (b *Bot) Run() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.api.GetUpdatesChan(u)

	for update := range updates {
		b.handleUpdate(update)
	}

	return nil
}

func (b *Bot) handleUpdate(update tgbotapi.Update) {
	if update.Message != nil {
		b.handleMessage(update)
		return
	}

	if update.CallbackQuery != nil {
		b.handleCallback(update)
		return
	}
}

func (b *Bot) handleMessage(update tgbotapi.Update) {
	message := update.Message

	if message.Text == "" {
		return
	}

	if message.IsCommand() {
		b.handleCommand(message)
		return
	}

	chatID := message.Chat.ID

	result := b.requestService.HandleText(
		chatID,
		message.Text,
	)

	if result.Text != "" {
		if err := b.sendMessage(chatID, result.Text); err != nil {
			log.Println("send message:", err)
		}
	}

	if result.Request != nil {
		requestText := formatRequest(*result.Request)

		if err := b.sendMessage(b.workChatID, requestText); err != nil {
			log.Println("send request to work chat:", err)
		}
	}
}

func (b *Bot) handleCommand(message *tgbotapi.Message) {
	chatID := message.Chat.ID

	switch message.Command() {
	case "start":
		err := b.sendMessageWithKeyboard(
			chatID,
			"Привет! Выберите направление:",
			directionKeyboard(),
		)
		if err != nil {
			log.Println("send /start:", err)
		}

	case "chatid":
		err := b.sendMessage(
			chatID,
			fmt.Sprintf("Chat ID: %d", chatID),
		)
		if err != nil {
			log.Println("send /chatid:", err)
		}

	case "help":
		err := b.sendMessage(
			chatID,
			"Доступные команды:\n"+
				"/start — начать оформление заявки\n"+
				"/chatid — узнать ID текущего чата\n"+
				"/help — список команд",
		)
		if err != nil {
			log.Println("send /help:", err)
		}
	}
}

func (b *Bot) handleCallback(update tgbotapi.Update) {
	callback := update.CallbackQuery

	if callback == nil || callback.Message == nil {
		return
	}

	defer b.answerCallback(callback.ID)

	chatID := callback.Message.Chat.ID

	if direction, ok := parseDirection(callback.Data); ok {
		result := b.requestService.StartRequest(
			chatID,
			direction,
		)

		if err := b.sendMessageWithKeyboard(
			chatID,
			result.Text,
			addressKeyboard(),
		); err != nil {
			log.Println("send address keyboard:", err)
		}

		return
	}

	if address, ok := parseAddress(callback.Data); ok {
		result := b.requestService.SetAddress(
			chatID,
			address,
		)

		if err := b.sendMessage(
			chatID,
			result.Text,
		); err != nil {
			log.Println("send address result:", err)
		}
	}
}

func (b *Bot) sendMessage(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)

	_, err := b.api.Send(msg)

	return err
}

func (b *Bot) sendMessageWithKeyboard(
	chatID int64,
	text string,
	keyboard tgbotapi.InlineKeyboardMarkup,
) error {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = keyboard

	_, err := b.api.Send(msg)

	return err
}

func (b *Bot) answerCallback(callbackID string) {
	_, err := b.api.Request(
		tgbotapi.NewCallback(callbackID, ""),
	)
	if err != nil {
		log.Println("answer callback:", err)
	}
}

func directionName(direction domain.Direction) string {
	switch direction {
	case domain.DirectionService:
		return "Сервис"
	case domain.DirectionDetailing:
		return "Детейлинг"
	case domain.DirectionCarWash:
		return "Мойка"
	case domain.DirectionBodyWork:
		return "Кузовной ремонт"
	default:
		return "Неизвестно"
	}
}

func addressName(address domain.Address) string {
	switch address {
	case domain.AddressPoeticheskiy:
		return "Поэтический"
	case domain.AddressSikeirosa:
		return "Сикейроса"
	default:
		return "Неизвестно"
	}
}

func formatRequest(req domain.Request) string {
	return fmt.Sprintf(
		"Новая заявка\n\n"+
			"Направление: %s\n"+
			"Адрес: %s\n"+
			"Имя: %s\n"+
			"Телефон: %s\n"+
			"Автомобиль: %s\n"+
			"Проблема: %s",
		directionName(req.Direction),
		addressName(req.Address),
		req.Name,
		req.PhoneNumber,
		req.Car,
		req.Problem,
	)
}
