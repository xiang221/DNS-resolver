package dnsmessage

import (
	"bytes"
	"encoding/binary"
)

type Header struct {
	ID      uint16
	Flags   uint16
	QDCount uint16
	ANCount uint16
	NSCount uint16
	ARCount uint16
}

func (h *Header) ToBytes() []byte {
	buf := new(bytes.Buffer)

	binary.Write(buf, binary.BigEndian, h.ID)
	binary.Write(buf, binary.BigEndian, h.Flags)
	binary.Write(buf, binary.BigEndian, h.QDCount)
	binary.Write(buf, binary.BigEndian, h.ANCount)
	binary.Write(buf, binary.BigEndian, h.NSCount)
	binary.Write(buf, binary.BigEndian, h.ARCount)

	return buf.Bytes()
}

type HeaderFlag struct {
	OR     bool
	Opcode uint8
	AA     bool
	TC     bool
	RD     bool
	RA     bool
	Z      uint8
	RCode  uint8
}

func (hf *HeaderFlag) GenerateFlag() uint16 {
	qr := uint16(boolToInt(hf.OR))
	opcode := uint16(hf.Opcode)
	aa := uint16(boolToInt(hf.AA))
	tc := uint16(boolToInt(hf.TC))
	rd := uint16(boolToInt(hf.RD))
	ra := uint16(boolToInt(hf.RA))
	z := uint16(hf.Z)
	rcode := uint16(hf.RCode)

	return uint16(qr<<15 | opcode<<11 | aa<<10 | tc<<9 | rd<<8 | ra<<7 | z<<4 | rcode)
}

func boolToInt(boolVal bool) uint16 {
	if boolVal {
		return 1
	}
	return 0
}
