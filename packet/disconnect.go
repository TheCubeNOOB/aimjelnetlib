package packet

import (
	"github.com/thecubenoob/aimjelnetlib/chat"
)

type DisconnectLogin struct {
	Reason chat.Message
}

func (l DisconnectLogin) ID() int32 {
	return 0x00
}

func (l *DisconnectLogin) Decode(r *Reader) error {
	return NotImplemented
}

func (l DisconnectLogin) Encode(w Writer) error {
	return w.String(l.Reason.String())
}

type DisconnectPlay struct {
	DisconnectLogin
}

func (p DisconnectPlay) ID() int32 {
	return 0x1a
}
