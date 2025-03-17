package commands

type Executor interface {
	SetCommand(fn CommandFunc)
	Execute() error
}

type CommandFunc func() error

type commander struct {
	fn CommandFunc
}

func NewCommander() Executor {
	return &commander{}
}

func (c *commander) SetCommand(fn CommandFunc) {
	c.fn = fn
}

func (c *commander) Execute() error {
	return c.fn()
}
