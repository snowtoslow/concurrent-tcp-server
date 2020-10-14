package receivedata

import "concurrent-tcp-server/server/command/receiver"

type ReceiveData struct {
	Device receiver.Device
}

func (p *ReceiveData) Execute(input *string) {
	p.Device.PrintResponse(input)
}
