package dnsmessage

import (
	"bytes"
	"encoding/binary"
	"strings"
)

type Question struct {
	Name   string
	QName  string
	QType  uint16
	QClass uint16
}

func encodeName(name string) string {
	domainParts := strings.Split(name, ".")
	qname := ""
	for _, part := range domainParts {
		newDomainPart := string(byte(len(part))) + part
		qname += newDomainPart
	}

	return qname + "\x00"
}

func (q *Question) ToBytes() []byte {
	buf := new(bytes.Buffer)
	buf.Write([]byte(q.QName))
	binary.Write(buf, binary.BigEndian, q.QType)
	binary.Write(buf, binary.BigEndian, q.QClass)
	return buf.Bytes()
}
