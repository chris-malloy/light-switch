package light_switch

import "fmt"

func Hello() string {
	return "Hello, world."
}

// command
// knows about `receiver` and invokes a method of the `receiver`
type Command interface {
	execute()
}

type SwitchOnCommand struct {
	Light *Light
}

func (cmd *SwitchOnCommand) execute() {
	cmd.Light.TurnOn()
}

type SwitchOffCommand struct {
	Light *Light
}

func (cmd *SwitchOffCommand) execute() {
	cmd.Light.TurnOff()
}

// receiver
// this does work when `execute()` is called
type Light struct{}

func (l *Light) TurnOn() {
	fmt.Print("the light is off")
}

func (l *Light) TurnOff() {
	fmt.Print("the light is off")
}

// invoker
// knows how to execute a command, and optionally does bookkeeping
type Switch struct {
	commands []Command
}
func (sw *Switch) StoreAndExecute(cmd Command) {
	sw.commands = append(sw.commands, cmd)
	cmd.execute()
}

// client
// decides which receiver object it assigns to the command objects
