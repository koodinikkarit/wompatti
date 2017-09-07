package WompattiServiceServer

import (
	"golang.org/x/net/context"

	WompattiService "github.com/koodinikkarit/wompatti/wompatti_service"
)

func (s *Server) CreateTelnetInterface(
	ctx context.Context,
	in *WompattiService.CreateTelnetInterfaceRequest,
) (*WompattiService.CreateTelnetInterfaceResponse, error) {
	res := &WompattiService.CreateTelnetInterfaceResponse{}
	return res, nil
}

func (s *Server) EditTelnetInterface(
	ctx context.Context,
	in *WompattiService.EditTelnetInterfaceRequest,
) (*WompattiService.EditTelnetInterfaceResponse, error) {
	res := &WompattiService.EditTelnetInterfaceResponse{}
	return res, nil
}

func (s *Server) RemoveTelnetInterface(
	ctx context.Context,
	in *WompattiService.RemoveTelnetInterfaceRequest,
) (*WompattiService.RemoveTelnetInterfaceResponse, error) {
	res := &WompattiService.RemoveTelnetInterfaceResponse{}
	return res, nil
}

func (s *Server) CreateDevice(
	ctxt context.Context,
	in *WompattiService.CreateDeviceRequest,
) (*WompattiService.CreateDeviceResponse, error) {
	res := &WompattiService.CreateDeviceResponse{}
	return res, nil
}

func (s *Server) EditDevice(
	ctx context.Context,
	in *WompattiService.EditDeviceRequest,
) (*WompattiService.EditDeviceResponse, error) {
	res := &WompattiService.EditDeviceResponse{}
	return res, nil
}

func (s *Server) RemoveDevice(
	ctx context.Context,
	in *WompattiService.RemoveDeviceRequest,
) (*WompattiService.RemoveDeviceResponse, error) {
	res := &WompattiService.RemoveDeviceResponse{}
	return res, nil
}

func (s *Server) CreateDeviceType(
	ctx context.Context,
	in *WompattiService.CreateDeviceTypeRequest,
) (*WompattiService.CreateDeviceTypeResponse, error) {
	res := &WompattiService.CreateDeviceTypeResponse{}
	return res, nil
}

func (s *Server) EditDeviceType(
	ctx context.Context,
	in *WompattiService.EditDeviceTypeRequest,
) (*WompattiService.EditDeviceTypeResponse, error) {
	res := &WompattiService.EditDeviceTypeResponse{}
	return res, nil
}

func (s *Server) RemoveDeviceType(
	ctx context.Context,
	in *WompattiService.RemoveDeviceTypeRequest,
) (*WompattiService.RemoveDeviceTypeResponse, error) {
	res := &WompattiService.RemoveDeviceTypeResponse{}
	return res, nil
}

func (s *Server) CreateCommand(
	ctx context.Context,
	in *WompattiService.CreateCommandRequest,
) (*WompattiService.CreateCommandResponse, error) {
	res := &WompattiService.CreateCommandResponse{}
	return res, nil
}

func (s *Server) EditCommand(
	ctx context.Context,
	in *WompattiService.EditCommandRequest,
) (*WompattiService.EditCommandReponse, error) {
	res := &WompattiService.EditCommandReponse{}
	return res, nil
}

func (s *Server) RemoveCommand(
	ctx context.Context,
	in *WompattiService.RemoveCommandRequest,
) (*WompattiService.RemoveCommandResponse, error) {
	res := &WompattiService.RemoveCommandResponse{}
	return res, nil
}

func (s *Server) CreateKeijo(
	ctx context.Context,
	in *WompattiService.CreateKeijoRequest,
) (*WompattiService.CreateKeijoResponse, error) {
	res := &WompattiService.CreateKeijoResponse{}
	return res, nil
}

func (s *Server) EditKeijo(
	ctx context.Context,
	in *WompattiService.EditKeijoRequest,
) (*WompattiService.EditKeijoResponse, error) {
	res := &WompattiService.EditKeijoResponse{}
	return res, nil
}

func (s *Server) RemoveKeijo(
	ctx context.Context,
	in *WompattiService.RemoveKeijoRequest,
) (*WompattiService.RemoveKeijoResponse, error) {
	res := &WompattiService.RemoveKeijoResponse{}
	return res, nil
}

func (s *Server) CreateSeveri(
	ctx context.Context,
	in *WompattiService.CreateSeveriRequest,
) (*WompattiService.CreateSeveriResponse, error) {
	res := &WompattiService.CreateSeveriResponse{}
	return res, nil
}

func (s *Server) EditSeveri(
	ctx context.Context,
	in *WompattiService.EditSeveriRequest,
) (*WompattiService.EditSeveriResponse, error) {
	res := &WompattiService.EditSeveriResponse{}
	return res, nil
}

func (s *Server) RemoveSeveri(
	ctx context.Context,
	in *WompattiService.RemoveSeveriRequest,
) (*WompattiService.RemoveSeveriResponse, error) {
	res := &WompattiService.RemoveSeveriResponse{}
	return res, nil
}

func (s *Server) CreateWolInterface(
	ctx context.Context,
	in *WompattiService.CreateWolInterfaceRequest,
) (*WompattiService.CreateWolInterfaceResponse, error) {
	res := &WompattiService.CreateWolInterfaceResponse{}
	context := s.newContext()
	wolInterface := context.CreateWolInterface(
		in.EthernetInterfaceId,
		in.Mac,
	)
	context.Commit()
	res.WolInterface = NewWolInterface(wolInterface)
	return res, nil
}

func (s *Server) EditWolInterface(
	ctx context.Context,
	in *WompattiService.EditWolInterfaceRequest,
) (*WompattiService.EditWolInterfaceResponse, error) {
	res := &WompattiService.EditWolInterfaceResponse{}
	context := s.newContext()
	wolInterface := context.EditWolInterface(
		in.WolInterfaceId,
		in.EthernetInterfaceId,
		in.Mac,
	)
	context.Commit()
	res.WolInterface = NewWolInterface(wolInterface)
	return res, nil
}

func (s *Server) RemoveWolInterface(
	ctx context.Context,
	in *WompattiService.RemoveWolInterfaceRequest,
) (*WompattiService.RemoveWolInterfaceResponse, error) {
	res := &WompattiService.RemoveWolInterfaceResponse{}
	context := s.newContext()
	res.Success = context.RemoveWolInterface(in.WolInterfaceId)
	context.Commit()
	return res, nil
}

func (s *Server) WakeupWolInterface(
	ctx context.Context,
	in *WompattiService.WakeupWolInterfaceRequest,
) (*WompattiService.WakeupWolInterfaceResponse, error) {
	res := &WompattiService.WakeupWolInterfaceResponse{}
	return res, nil
}
