package WompattiServiceServer

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/koodinikkarit/wompatti/models"
	"golang.org/x/net/context"

	WompattiService "github.com/koodinikkarit/wompatti/wompatti_service"
)

func (s *Server) CreateComputer(
	ctx context.Context,
	in *WompattiService.CreateComputerRequest,
) (*WompattiService.CreateComputerResponse, error) {
	res := &WompattiService.CreateComputerResponse{}
	context := s.newContext()
	res.Computer = NewComputer(context.CreateComputer(in.Name))
	//context.Commit()
	return res, nil
}

func (s *Server) UpdateComputer(
	ctx context.Context,
	in *WompattiService.UpdateComputerRequest,
) (*WompattiService.UpdateComputerResponse, error) {
	res := &WompattiService.UpdateComputerResponse{}
	context := s.newContext()
	res.Computer = NewComputer(context.UpdateComputer(
		in.ComputerId,
		in.Name,
		in.WolInterfaceId,
		in.Ip,
		in.Mac,
	))
	//context.Commit()
	return res, nil
}

func (s *Server) RemoveComputer(
	ctx context.Context,
	in *WompattiService.RemoveComputerRequest,
) (*WompattiService.RemoveComputerResponse, error) {
	res := &WompattiService.RemoveComputerResponse{}
	c := s.newContext()
	res.Success = c.RemoveComputer(in.ComputerId)
	c.Commit()
	return res, nil
}

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
		in.Ip,
		in.Port,
	)
	//context.Commit()
	res.WolInterface = NewWolInterface(wolInterface)
	return res, nil
}

func (s *Server) UpdateWolInterface(
	ctx context.Context,
	in *WompattiService.UpdateWolInterfaceRequest,
) (*WompattiService.UpdateWolInterfaceResponse, error) {
	res := &WompattiService.UpdateWolInterfaceResponse{}
	context := s.newContext()
	wolInterface := context.UpdateWolInterface(
		in.WolInterfaceId,
		in.Ip,
		in.Port,
	)
	//context.Commit()
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
	//context.Commit()
	return res, nil
}

func (s *Server) Wakeup(
	ctx context.Context,
	in *WompattiService.WakeupRequest,
) (*WompattiService.WakeupResponse, error) {
	res := &WompattiService.WakeupResponse{}

	context := s.newContext()

	var computer WompattiModels.Computer
	context.GetDb().First(&computer, in.ComputerId)
	if computer.ID > 0 {
		var wolInterface WompattiModels.WolInterface
		context.GetDb().First(&wolInterface, computer.WolInterfaceID)
		if wolInterface.ID > 0 {
			conn, err := net.Dial("udp", wolInterface.IP+":"+strconv.Itoa(int(wolInterface.Port)))
			if err != nil {
				return nil, nil
			}
			defer conn.Close()

			var mac []byte

			macParts := strings.Split(computer.Mac, ":")
			for i := 0; i < len(macParts); i++ {
				x, _ := strconv.ParseInt(macParts[i], 16, 0)
				fmt.Println(x)
				mac = append(mac, byte(x))
			}

			var buff []byte
			buff = append(buff, 255)
			buff = append(buff, 255)
			buff = append(buff, 255)
			buff = append(buff, 255)
			buff = append(buff, 255)
			buff = append(buff, 255)

			for i := 0; i < 16; i++ {
				for j := 0; j < len(mac); j++ {
					//buff[6+(i*6)+j] = mac[j]
					buff = append(buff, mac[j])
				}
			}

			conn.Write(buff)
			res.Success = true
		}
	}

	res.Success = false

	// context := s.newContext()
	// var wolInterface WompattiModels.WolInterface
	// context.GetDb().First(&wolInterface, in.WolInterfaceId)
	// if wolInterface.ID > 0 {
	// 	wol.MagicWake(wolInterface.Mac, "255.255.255.255")
	// 	res.Success = true
	// } else {
	// 	res.Success = false
	// }

	return res, nil
}
