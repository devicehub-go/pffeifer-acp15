package pffeiferacp15

import (
	"github.com/devicehub-go/pffeifer-acp15/protocol"
	"github.com/devicehub-go/unicomm"
)

/*
Creates a new Pfeiffer ACP 15 instance that allows to communicate
with the device through the selected protocol

Parameters:
  - options: communication options to connect to the
    device, including the protocol and the respective options

Note: delimiter is not required
*/
func New(options unicomm.Options) *protocol.ACP15 {
	options.Delimiter = "\r"
	return &protocol.ACP15{
		Communication: unicomm.New(options),
	}
}
