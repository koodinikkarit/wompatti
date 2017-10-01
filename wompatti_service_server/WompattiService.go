package WompattiServiceServer

import (
	"github.com/koodinikkarit/wompatti/context"
	"github.com/koodinikkarit/wompatti/pinger"
)

type Server struct {
	newContext func() *WompattiContext.Context
	pinger     *WompattiPinger.Pinger
}
