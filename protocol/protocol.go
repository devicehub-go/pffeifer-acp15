package protocol

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/devicehub-go/unicomm"
)

type ACP15 struct {
	Communication unicomm.Unicomm
}

type ControlMessage struct {
	Address   uint8
	Order     string
	Parameter string
}

type ResponseMessage struct {
	Address uint8
	Value   string
	Error   error
}

// Establishes a connection with the device
func (a *ACP15) Connect() error {
	return a.Communication.Connect()
}

// Closes the connection with the device
func (a *ACP15) Disconnect() error {
	return a.Communication.Disconnect()
}

// Checks if the device is connected
func (a *ACP15) IsConnected() bool {
	return a.Communication.IsConnected()
}

// Writes a control message to the device
func (a *ACP15) WriteControlMessage(me ControlMessage) error {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("%02X", me.Address))
	sb.WriteString(me.Order)
	sb.WriteString(me.Parameter)
	sb.WriteString("\r")

	return a.Communication.Write([]byte(sb.String()))
}

// Returns the error code as a program error
func getErrorDescription(code string) error {
	switch code {
	case "ERR0":
		return fmt.Errorf("setting fault")
	case "ERR1":
		return fmt.Errorf("context fault")
	case "ERR2":
		return fmt.Errorf("parameter fault")
	case "ERR3":
		return fmt.Errorf("order fault")
	default:
		return fmt.Errorf("unknown error")
	}
}

// Reads the response from the device
func (a *ACP15) ReadResponse() (ResponseMessage, error) {
	response, err := a.Communication.ReadUntil("\r")
	if err != nil {
		return ResponseMessage{}, err
	}

	expected := regexp.MustCompile(`^\#(\d{3})(.*)\r$`)
	matches := expected.FindStringSubmatch(string(response))
	if len(matches) != 3 {
		return ResponseMessage{}, fmt.Errorf("invalid response, got: %s", &response)
	}

	addr, err := strconv.Atoi(matches[1])
	if err != nil {
		return ResponseMessage{}, fmt.Errorf("invalid address in response: %v", err)
	}

	responseMsg := ResponseMessage{
		Address: uint8(addr),
		Value:   strings.TrimSuffix(matches[2], "\r"),
	}
	if strings.Contains(matches[2], "ERR") {
		responseMsg.Error = getErrorDescription(matches[2])
	}
	return responseMsg, nil
}

// Sends a control message and reat the response
func (a *ACP15) Request(message ControlMessage) (ResponseMessage, error) {
	if err := a.WriteControlMessage(message); err != nil {
		return ResponseMessage{}, fmt.Errorf("error on write control message: %w", err)
	}
	response, err := a.ReadResponse()
	if err != nil {
		return ResponseMessage{}, fmt.Errorf("error read response: %w", err)
	}
	if response.Error != nil {
		return response, response.Error
	}
	if response.Address != message.Address {
		return ResponseMessage{}, fmt.Errorf(
			"response address mismatch: expected %02X, got %02X",
			message.Address, response.Address,
		)
	}
	return response, nil
}
