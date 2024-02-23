package types

type (
	Device struct {
		id      string
		name    string
		address string
		banned  bool
	}

	DeviceArguments struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Address string `json:"address"`
	}
)

func (d *Device) ID() string {
	return d.id
}

func (d *Device) Name() string {
	return d.name
}

func (d *Device) Address() string {
	return d.address
}

func (d *Device) SetID(id string) error {
	d.id = id
	return nil
}

func (d *Device) SetName(name string) error {
	d.name = name
	return nil
}

func (d *Device) SetAddress(address string) error {
	d.address = address
	return nil
}

func (d *Device) Banned(banned bool) error {
	d.banned = banned
	return nil
}
func (d *Device) IsBanned() bool {
	return d.banned
}
