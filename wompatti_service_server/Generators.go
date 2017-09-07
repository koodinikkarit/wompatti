package WompattiServiceServer

import (
	"github.com/koodinikkarit/wompatti/models"
	WompattiService "github.com/koodinikkarit/wompatti/wompatti_service"
)

func NewWolInterface(in *WompattiModels.WolInterface) *WompattiService.WolInterface {
	return &WompattiService.WolInterface{
		Id:                  in.ID,
		EthernetInterfaceId: in.EthernetInterfaceID,
		Mac:                 in.Mac,
	}
}
