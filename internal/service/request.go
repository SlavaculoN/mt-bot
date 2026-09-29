package service

import (
	"mt-bot/internal/domain"
)

type SessionStorage interface {
	Get(chatID int64) (*domain.UserSession, bool)
	Set(chatID int64, session *domain.UserSession)
	Delete(chatID int64)
}

type RequestService struct {
	storage SessionStorage
}

type HandleResult struct {
	Text    string
	Action  Action
	Request *domain.Request
}

type Action string

const (
	ActionNone          Action = ""
	ActionChooseAddress Action = "choose_address"
	ActionSendRequest   Action = "send_request"
	ActionShowKeyboard  Action = "show_phone_keyboard"
)

func NewRequestService(storage SessionStorage) *RequestService {
	return &RequestService{
		storage: storage,
	}
}

func (r *RequestService) StartRequest(
	chatID int64,
	direction domain.Direction,
) HandleResult {
	session := &domain.UserSession{
		State: domain.StateWaitAddress,
		Request: domain.Request{
			Direction: direction,
		},
	}

	r.storage.Set(chatID, session)

	return HandleResult{
		Text:   "Вы выбрали направление.\nВыберете адрес:",
		Action: ActionChooseAddress,
	}
}

func (r *RequestService) ConfirmRequest(chatID int64) (*domain.Request, bool) {
	session, ok := r.storage.Get(chatID)
	if !ok {
		return nil, false
	}

	if session.State != domain.StateWaitConfirmation {
		return nil, false
	}

	request := session.Request

	return &request, true
}

func (r *RequestService) FinishRequest(chatID int64) {
	r.storage.Delete(chatID)
}

func (r *RequestService) CancelRequest(chatID int64) bool {
	_, ok := r.storage.Get(chatID)
	if !ok {
		return false
	}

	r.storage.Delete(chatID)

	return true
}

func (r *RequestService) SetAddress(
	chatID int64,
	address domain.Address,
) HandleResult {
	session, ok := r.storage.Get(chatID)

	if !ok {
		return HandleResult{
			Text: "Активная заявка не найдена. Нажмите /start.",
		}
	}

	if session.State != domain.StateWaitAddress {
		return HandleResult{
			Text: "Сейчас выбор адреса не доступен.",
		}
	}

	if !isAddressAllowed(session.Request.Direction, address) {
		return HandleResult{
			Text: "Выбранный адрес недоступен для этого направления.",
		}
	}

	session.Request.Address = address
	session.State = domain.StateWaitName

	r.storage.Set(chatID, session)

	return HandleResult{
		Text: "Адрес записал.\nВведите Ваше имя:",
	}
}

func isAddressAllowed(direction domain.Direction, address domain.Address) bool {

	switch direction {
	case domain.DirectionService:
		return address == domain.AddressPoeticheskiy || address == domain.AddressSikeirosa
	case domain.DirectionDetailing, domain.DirectionBodyWork, domain.DirectionCarWash:
		return address == domain.AddressPoeticheskiy
	default:
		return false
	}

}

func (r *RequestService) SetName(chatID int64, name string) bool {
	session, ok := r.storage.Get(chatID)
	if !ok {
		return false
	}

	if session.State != domain.StateWaitName {
		return false
	}

	session.Request.Name = name
	session.State = domain.StateWaitPhoneNumber

	r.storage.Set(chatID, session)

	return true
}

func (r *RequestService) SetPhoneNumber(
	chatID int64,
	phone string,
) HandleResult {
	session, ok := r.storage.Get(chatID)
	if !ok {
		return HandleResult{
			Text: "Активная заявка не найдена. Нажмите /start.",
		}
	}

	if session.State != domain.StateWaitPhoneNumber {
		return HandleResult{
			Text: "Сейчас номер теелфона не запрашивается.",
		}
	}

	session.Request.PhoneNumber = phone
	session.State = domain.StateWaitCar

	r.storage.Set(chatID, session)

	return HandleResult{
		Text: "Номер телефона записал.\nВведите марку автомобиля и модель:",
	}
}

func (r *RequestService) SetCar(chatID int64, car string) bool {
	session, ok := r.storage.Get(chatID)
	if !ok {
		return false
	}

	if session.State != domain.StateWaitCar {
		return false
	}

	session.Request.Car = car
	session.State = domain.StateWaitProblem

	r.storage.Set(chatID, session)

	return true
}

func (r *RequestService) SetProblem(chatID int64, problem string) (*domain.Request, bool) {
	session, ok := r.storage.Get(chatID)
	if !ok {
		return nil, false
	}

	if session.State != domain.StateWaitProblem {
		return nil, false
	}

	session.Request.Problem = problem

	request := session.Request

	r.storage.Delete(chatID)

	return &request, true
}

func (r *RequestService) HandleText(
	chatID int64,
	text string,
) HandleResult {
	session, ok := r.storage.Get(chatID)
	if !ok {
		return HandleResult{
			Text: "У Вас нет активной заявки. Нажмите /start.",
		}
	}

	switch session.State {
	case domain.StateWaitName:
		session.Request.Name = text
		session.State = domain.StateWaitPhoneNumber

		r.storage.Set(chatID, session)

		return HandleResult{
			Text:   "Имя записал.\nВведите номер телефона:",
			Action: ActionShowKeyboard,
		}

	case domain.StateWaitPhoneNumber:
		session.Request.PhoneNumber = text
		session.State = domain.StateWaitCar

		r.storage.Set(chatID, session)

		return HandleResult{
			Text: "Телефон записал.\nВведите марку автомобиля и модель:",
		}

	case domain.StateWaitCar:
		session.Request.Car = text
		session.State = domain.StateWaitProblem

		r.storage.Set(chatID, session)

		return HandleResult{
			Text: "Автомобиль записал.\nОпишите проблему:",
		}

	case domain.StateWaitProblem:
		session.Request.Problem = text
		session.State = domain.StateWaitConfirmation

		r.storage.Set(chatID, session)

		request := session.Request

		return HandleResult{
			Request: &request,
		}

	default:
		return HandleResult{
			Text: "Неизвестное состояние заявки",
		}
	}
}
