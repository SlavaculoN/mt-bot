package domain

type Direction string

const (
	DirectionService   Direction = "service"
	DirectionDetailing Direction = "detailing"
	DirectionCarWash   Direction = "carwash"
	DirectionBodyWork  Direction = "bodywork"
)

type Address string

const (
	AddressPoeticheskiy Address = "poeticheskiy"
	AddressSikeirosa    Address = "sikeirosa"
)

type Request struct {
	Direction   Direction
	Address     Address
	Name        string
	PhoneNumber string
	Car         string
	Problem     string
}

type UserState string

const (
	StateNone             UserState = ""
	StateWaitAddress      UserState = "wait_address"
	StateWaitName         UserState = "wait_name"
	StateWaitPhoneNumber  UserState = "wait_phone_number"
	StateWaitCar          UserState = "wait_car"
	StateWaitProblem      UserState = "wait_problem"
	StateWaitConfirmation UserState = "wait_confirmation"
)

type UserSession struct {
	State   UserState
	Request Request
}
