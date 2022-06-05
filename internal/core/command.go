package core

type Command interface {
	Run() error
}
