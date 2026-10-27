package protocol

import "fmt"

type PumpStatus struct {
	RotationSpeed   uint16
	MotorPower      uint16
	Temperature     uint16
	OperatingHours  uint32
	IsRunning       bool
	SetNominalSpeed bool
	SetTargetSpeed  bool
	IsStandBy       bool
	PowerFailure    bool
	OverTemperature bool
	OverCurrent     bool
}

// Gets the identification of the connected device
func (a *ACP15) GetIDN(address uint8) (string, error) {
	response, err := a.Request(ControlMessage{
		Address:   address,
		Order:     "IDN",
		Parameter: "",
	})
	if err != nil {
		return "", err
	}
	return response.Value, nil
}

// Gets the pump status
func (a *ACP15) GetStatus(address uint8) (PumpStatus, error) {
	response, err := a.Request(ControlMessage{
		Address:   address,
		Order:     "STA",
		Parameter: "",
	})
	if err != nil {
		return PumpStatus{}, err
	}
	var status PumpStatus
	fmt.Sscanf(response.Value, ",%d,%d,%d,%d,%t,%t,%t,%t,%t,%t",
		&status.RotationSpeed,
		&status.MotorPower,
		&status.Temperature,
		&status.OperatingHours,
		&status.IsRunning,
		&status.SetNominalSpeed,
		&status.SetTargetSpeed,
		&status.IsStandBy,
		&status.PowerFailure,
		&status.OverTemperature,
	)
	return status, nil
}
