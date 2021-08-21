package ws

type Frame struct {
	IsFragment bool // if the
	Opcode     byte
	Reserved   byte
	IsMasked   bool
	Length     uint64
	Payload    []byte
}

// IsControl checks if the frame is a control frame identified by opcodes where the most significant bit of the opcode is 1
func (f *Frame) IsControl() bool {
	return f.Opcode&0x08 == 0x08
}

func (f *Frame) HasReservedOpcode() bool {
	return f.Opcode > 10 || (f.Opcode >= 3 && f.Opcode <= 7)
}