package dnsmessage

type DNSMessage struct {
	Header        Header
	Question      []Question
	Answers       []ResourceRecord
	AuthorityRRs  []ResourceRecord
	AdditionalRRs []ResourceRecord
}

func NewDNSMessage(header Header, questions []Question, records ...[]ResourceRecord) *DNSMessage {
	answers := make([]ResourceRecord, 0)
	authorityRRs := make([]ResourceRecord, 0)
	additionalRRs := make([]ResourceRecord, 0)

	if len(records) > 0 {
		answers = records[0]
	}

	if len(records) > 1 {
		authorityRRs = records[1]
	}

	if len(records) > 2 {
		additionalRRs = records[2]
	}

	return &DNSMessage{
		Header:        header,
		Question:      questions,
		Answers:       answers,
		AuthorityRRs:  authorityRRs,
		AdditionalRRs: additionalRRs,
	}
}
