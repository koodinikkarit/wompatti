package WompattiServiceServer

import (
	"github.com/koodinikkarit/wompatti/context"
)

type Server struct {
	newContext func() *WompattiContext.Context
}
