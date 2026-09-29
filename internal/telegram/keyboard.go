package telegram

import (
	"mt-bot/internal/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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

func addressKeyboard(direction domain.Direction) tgbotapi.InlineKeyboardMarkup {
	if direction == domain.DirectionService {
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

	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"Поэтический",
				CallbackAddressPoeticheskiy,
			),
		),
	)
}

func phoneKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButtonContact("📱 Поделиться номером"),
		),
	)
}

func confirmationKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"✅ Отправить",
				"request:confirm",
			),
			tgbotapi.NewInlineKeyboardButtonData(
				"❌ Отменить",
				"request:cancel",
			),
		),
	)
}
