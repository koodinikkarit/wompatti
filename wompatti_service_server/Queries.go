package WompattiServiceServer

import (
	"golang.org/x/net/context"

	"github.com/koodinikkarit/wompatti/models"
	WompattiService "github.com/koodinikkarit/wompatti/wompatti_service"
)

func (s *Server) FetchComputers(
	ctx context.Context,
	in *WompattiService.FetchComputersRequest,
) (*WompattiService.FetchComputersResponse, error) {
	res := &WompattiService.FetchComputersResponse{}
	c := s.newContext()
	computers := []WompattiModels.Computer{}
	c.GetDb().Find(&computers)
	for i := 0; i < len(computers); i++ {
		res.Computers = append(
			res.Computers,
			NewComputer(&computers[i]),
		)
	}
	return res, nil
}

func (s *Server) FetchComputerById(
	ctx context.Context,
	in *WompattiService.FetchComputerByIdRequest,
) (*WompattiService.FetchComputerByIdResponse, error) {
	res := &WompattiService.FetchComputerByIdResponse{}
	c := s.newContext()
	computers := []WompattiModels.Computer{}
	c.GetDb().Where("id in (?)", in.ComputerIds).
		Find(&computers)

	for i := 0; i < len(in.ComputerIds); i++ {
		found := false
		for j := 0; j < len(computers); j++ {
			if in.ComputerIds[i] == computers[j].ID {
				found = true
				res.Computers = append(
					res.Computers,
					NewComputer(&computers[j]),
				)
			}
		}
		if found == false {
			res.Computers = append(
				res.Computers,
				&WompattiService.Computer{
					Id: 0,
				},
			)
		}
	}

	return res, nil
}

func (s *Server) FetchTelnetInterfaces(
	ctx context.Context,
	in *WompattiService.FetchTelnetInterfacesRequest,
) (*WompattiService.FetchTelnetInterfacesResponse, error) {
	res := &WompattiService.FetchTelnetInterfacesResponse{}

	return res, nil
}

func (s *Server) FetchTelnetInterfaceById(
	ctx context.Context,
	in *WompattiService.FetchTelnetInterfaceByIdRequest,
) (*WompattiService.FetchTelnetInterfaceByIdResponse, error) {
	res := &WompattiService.FetchTelnetInterfaceByIdResponse{}
	return res, nil
}

func (s *Server) FetchDevices(
	ctx context.Context,
	in *WompattiService.FetchDevicesRequest,
) (*WompattiService.FetchDevicesResponse, error) {
	res := &WompattiService.FetchDevicesResponse{}
	return res, nil
}

func (s *Server) FetchDeviceById(
	ctx context.Context,
	in *WompattiService.FetchDeviceByIdRequest,
) (*WompattiService.FetchDeviceByIdResponse, error) {
	res := &WompattiService.FetchDeviceByIdResponse{}
	return res, nil
}

func (s *Server) FetchDeviceTypes(
	ctx context.Context,
	in *WompattiService.FetchDeviceTypesRequest,
) (*WompattiService.FetchDeviceTypesResponse, error) {
	res := &WompattiService.FetchDeviceTypesResponse{}
	return res, nil
}

func (s *Server) FetchDeviceTypeById(
	ctx context.Context,
	in *WompattiService.FetchDeviceTypeByIdRequest,
) (*WompattiService.FetchDeviceTypeByIdResponse, error) {
	res := &WompattiService.FetchDeviceTypeByIdResponse{}
	return res, nil
}

func (s *Server) FetchCommands(
	ctx context.Context,
	in *WompattiService.FetchCommandsRequest,
) (*WompattiService.FetchCommandsResponse, error) {
	res := &WompattiService.FetchCommandsResponse{}
	return res, nil
}

func (s *Server) FetchCommandsByDeviceTypeId(
	ctx context.Context,
	in *WompattiService.FetchCommandsByDeviceTypeIdRequest,
) (*WompattiService.FetchCommandsByDeviceTypeIdResponse, error) {
	res := &WompattiService.FetchCommandsByDeviceTypeIdResponse{}
	return res, nil
}

func (s *Server) FetchCommandById(
	ctx context.Context,
	in *WompattiService.FetchCommandByIdRequest,
) (*WompattiService.FetchCommandByIdResponse, error) {
	res := &WompattiService.FetchCommandByIdResponse{}
	return res, nil
}

func (s *Server) FetchKeijos(
	ctx context.Context,
	in *WompattiService.FetchKeijosRequest,
) (*WompattiService.FetchKeijosResponse, error) {
	res := &WompattiService.FetchKeijosResponse{}
	return res, nil
}

func (s *Server) FetchKeijoById(
	ctx context.Context,
	in *WompattiService.FetchKeijoByIdRequest,
) (*WompattiService.FetchKeijoByIdResponse, error) {
	res := &WompattiService.FetchKeijoByIdResponse{}
	return res, nil
}

func (s *Server) FetchSeveris(
	ctx context.Context,
	in *WompattiService.FetchSeverisRequest,
) (*WompattiService.FetchSeverisResponse, error) {
	res := &WompattiService.FetchSeverisResponse{}
	return res, nil
}

func (s *Server) FetchSeveriById(
	ctx context.Context,
	in *WompattiService.FetchSeveriByIdRequest,
) (*WompattiService.FetchSeveriByIdResponse, error) {
	res := &WompattiService.FetchSeveriByIdResponse{}
	return res, nil
}

func (s *Server) FetchWolInterfaces(
	ctx context.Context,
	in *WompattiService.FetchWolInterfacesRequest,
) (*WompattiService.FetchWolInterfacesResponse, error) {
	res := &WompattiService.FetchWolInterfacesResponse{}
	c := s.newContext()
	wompattiDb := c.GetDb()
	wolInterfaces := []WompattiModels.WolInterface{}
	wompattiDb.Find(&wolInterfaces)
	for i := 0; i < len(wolInterfaces); i++ {
		res.WolInterfaces = append(
			res.WolInterfaces,
			NewWolInterface(&wolInterfaces[i]),
		)
	}
	return res, nil
}

func (s *Server) FetchWolInterfaceById(
	ctx context.Context,
	in *WompattiService.FetchWolInterfaceByIdRequest,
) (*WompattiService.FetchWolInterfaceByIdResponse, error) {
	res := &WompattiService.FetchWolInterfaceByIdResponse{}
	context := s.newContext()

	wolInterfaces := []WompattiModels.WolInterface{}
	context.GetDb().Where("id in (?)", in.WolInterfaceIds).
		Find(&wolInterfaces)

	for i := 0; i < len(in.WolInterfaceIds); i++ {
		found := false
		for j := 0; j < len(wolInterfaces); j++ {
			if in.WolInterfaceIds[i] == wolInterfaces[j].ID {
				found = true
				res.WolInterfaces = append(
					res.WolInterfaces,
					NewWolInterface(&wolInterfaces[j]),
				)
			}
		}
		if found == false {
			res.WolInterfaces = append(
				res.WolInterfaces,
				&WompattiService.WolInterface{
					Id: 0,
				},
			)
		}
	}

	return res, nil
}

func (s *Server) Ping(
	ctx context.Context,
	in *WompattiService.PingRequest,
) (
	*WompattiService.PingResponse,
	error,
) {
	res := <-s.pinger.Ping(in.IpAddress)
	return &res, nil
}
