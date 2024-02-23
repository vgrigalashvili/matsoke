package usecases

import "github.com/vgrigalashvili/matsoke/internal/domain/types"

func CreateDevice(a *types.DeviceArguments) *types.Device {
	device := &types.Device{}

	device.SetID(a.ID)
	device.SetName(a.Name)
	device.SetAddress(a.Address)
	device.Banned(false)

	return device
}
