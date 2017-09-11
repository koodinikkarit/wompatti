package WompattiServiceServer

import (
	"github.com/koodinikkarit/wompatti/models"
	WompattiService "github.com/koodinikkarit/wompatti/wompatti_service"
)

func NewComputer(in *WompattiModels.Computer) *WompattiService.Computer {
	return &WompattiService.Computer{
		Id:   in.ID,
		Name: in.Name,
	}
}

func NewWolInterface(in *WompattiModels.WolInterface) *WompattiService.WolInterface {
	return &WompattiService.WolInterface{
		Id:                  in.ID,
		EthernetInterfaceId: in.EthernetInterfaceID,
		Mac:                 in.Mac,
	}
}
