package receiver

type Device interface {
	Validate(*string)
	PrintResponse(*string)
}
