package telegram

import (
	"log"
	"mt-bot/internal/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	CallbackDirectionService   = "direction:service"
	CallbackDirectionDetailing = "direction:detailing"
	CallbackDirectionCarWash   = "direction:carwash"
	CallbackDirectionBodyWork  = "direction:bodywork"

	CallbackAddressPoeticheskiy = "address:poeticheskiy"
	CallbackAddressSikeirosa    = "address:sikeirosa"
)

const (
	CallbackRequestConfirm = "request:confirm"
	CallbackRequestCancel  = "request:cancel"
)

func parseDirection(data string) (domain.Direction, bool) {
	switch data {
	case CallbackDirectionService:
		return domain.DirectionService, true

	case CallbackDirectionDetailing:
		return domain.DirectionDetailing, true

	case CallbackDirectionCarWash:
		return domain.DirectionCarWash, true

	case CallbackDirectionBodyWork:
		return domain.DirectionBodyWork, true

	default:
		return "", false
	}
}

func parseAddress(data string) (domain.Address, bool) {
	switch data {
	case CallbackAddressPoeticheskiy:
		return domain.AddressPoeticheskiy, true

	case CallbackAddressSikeirosa:
		return domain.AddressSikeirosa, true

	default:
		return "", false
	}
}

func answerCallback(bot *tgbotapi.BotAPI, callbackID string) {
	_, err := bot.Request(
		tgbotapi.NewCallback(callbackID, ""),
	)

	if err != nil {
		log.Println(err)
	}
}
