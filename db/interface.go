package wompatti

type Interface struct {
	ID      uint32
	Name    string
	Mac     []byte
	Index   uint32
	Mtu     uint32
	Flags   uint32
	ArttuID uint32
}
