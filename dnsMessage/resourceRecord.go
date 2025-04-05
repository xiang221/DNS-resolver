package dnsmessage

type ResourceRecord struct {
	Name        string
	Type        uint16
	Class       uint16
	TTL         uint32
	RDLength    uint16
	RData       []byte
	RDataParsed string
}
