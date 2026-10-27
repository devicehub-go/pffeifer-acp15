package protocol

import "fmt"

// Switches from set point speed to the nominal speed (100 Hz)
func (a *ACP15) SetNominalSpeed(address uint8) error {
	_, err := a.Request(ControlMessage{
		Address:   address,
		Order:     "NSP",
		Parameter: "",
	})
	return err
}

// Sets the target speed of the pump in Hz
func (a *ACP15) SetTargetSpeed(address uint8, speed uint16) error {
	if speed < 35 || speed > 100 {
		return fmt.Errorf("speed must be between 35-100 Hz, got %d", speed)
	}
	if err := a.SetStandBySpeed(address); err != nil {
		return err
	}
	_, err := a.Request(ControlMessage{
		Address:   address,
		Order:     "TSP",
		Parameter: fmt.Sprintf("%04d", speed),
	})
	return err
}

// Stand-by speed is reset at 35 Hz
func (a *ACP15) SetStandBySpeed(address uint8) error {
	_, err := a.Request(ControlMessage{
		Address:   address,
		Order:     "SBY",
		Parameter: "",
	})
	return err
}

// Turns the pump on
func (a *ACP15) TurnOn(address uint8) error {
	_, err := a.Request(ControlMessage{
		Address:   address,
		Order:     "ACP",
		Parameter: "ON",
	})
	return err
}

// Turns the pump off
func (a *ACP15) TurnOff(address uint8) error {
	_, err := a.Request(ControlMessage{
		Address:   address,
		Order:     "ACP",
		Parameter: "OFF",
	})
	return err
}
