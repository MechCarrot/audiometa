package interfaces

type Command interface {
	ParseFlag([]string) error
	Run() error
	Name() string
}
