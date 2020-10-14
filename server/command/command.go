package command

type Command interface {
	Execute(input *string)
}
