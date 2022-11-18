package debugger

type Debugger struct {
	Address string
}

func New() Debugger {
	//TOOD: no 0x00 address
	address := "0x00"
	return Debugger{
		Address: address,
	}
}
