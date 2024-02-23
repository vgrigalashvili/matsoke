package usecases

import "github.com/vgrigalashvili/matsoke/internal/domain/types"

func UpdateDevice(device *types.Device, id, name, address string) *types.Device {

	if id != "" {
		device.SetID(id)
	}

	if name != "" {
		device.SetName(name)
	}

	if address != "" {
		device.SetAddress(address)
	}

	return device
}
