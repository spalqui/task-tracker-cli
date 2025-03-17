package commands

type Executor interface {
	SetCommand(fn CommandFunc, params ...string)
	Execute() error
}

type CommandFunc func(params ...string) error

type commander struct {
	fn     CommandFunc
	params []string
}

func NewCommander() Executor {
	return &commander{}
}

func (c *commander) SetCommand(fn CommandFunc, params ...string) {
	c.fn = fn
	c.params = params
}

func (c *commander) Execute() error {
	return c.fn(c.params...)
}
