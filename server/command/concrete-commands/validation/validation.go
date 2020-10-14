package validation

import "concurrent-tcp-server/server/command/receiver"

type Validate struct {
	Device receiver.Device
}

func (v *Validate) Execute(input *string) {
	v.Device.Validate(input)
}
