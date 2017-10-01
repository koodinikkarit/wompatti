package WompattiServiceServer

import (
	"github.com/koodinikkarit/wompatti/models"
	WompattiService "github.com/koodinikkarit/wompatti/wompatti_service"
)

func NewComputer(in *WompattiModels.Computer) *WompattiService.Computer {
	return &WompattiService.Computer{
		Id:             in.ID,
		Name:           in.Name,
		Ip:             in.IP,
		Mac:            in.Mac,
		WolInterfaceId: in.WolInterfaceID,
	}
}

func NewWolInterface(in *WompattiModels.WolInterface) *WompattiService.WolInterface {
	return &WompattiService.WolInterface{
		Id:   in.ID,
		Ip:   in.IP,
		Port: in.Port,
	}
}
