package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func directionKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"Сервис",
				CallbackDirectionService,
			),
			tgbotapi.NewInlineKeyboardButtonData(
				"Детейлинг",
				CallbackDirectionDetailing,
			),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"Мойка",
				CallbackDirectionCarWash,
			),
			tgbotapi.NewInlineKeyboardButtonData(
				"Кузовной ремонт",
				CallbackDirectionBodyWork,
			),
		),
	)
}

func addressKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"Поэтический",
				CallbackAddressPoeticheskiy,
			),
			tgbotapi.NewInlineKeyboardButtonData(
				"Сикейроса",
				CallbackAddressSikeirosa,
			),
		),
	)
}
