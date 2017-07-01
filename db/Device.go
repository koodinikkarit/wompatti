package wompatti

// Device
type Device struct {
	ID                uint32
	DeviceInfoID      uint32
	Name              string
	DeviceTypeID      uint32
	SerialInterfaceID uint32
	TelnetInterfaceID uint32
	CecInterfaceID    uint32
}
