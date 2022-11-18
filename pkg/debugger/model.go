package debugger

import "github.com/r4gn4r-l0/evm-knife/pkg/evm"

type Debugger struct {
	Tx evm.Tx
}

func New() Debugger {
	// TOOD: no 0x00 address
	// TODO: remove hardcoded value
	address := "0x00"
	tx := evm.Tx{
		Origin: address,
		Value:  []byte{0x01},
	}
	return Debugger{
		Tx: tx,
	}
}
