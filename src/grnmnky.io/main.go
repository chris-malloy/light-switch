package light_switch

// client
func main() {

	// receiver
	lamp := new(Light)

	// concrete command
	switchOn := &SwitchOnCommand{Light: lamp}
	switchOff := &SwitchOffCommand{Light: lamp}

	// invoker
	mySwitch := new(Switch)
	mySwitch.StoreAndExecute(switchOn)
	mySwitch.StoreAndExecute(switchOff)
}
