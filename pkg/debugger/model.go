package debugger

import "github.com/r4gn4r-l0/evm-knife/pkg/evm"

type Debugger struct {
	evm evm.EVM
}

func New() Debugger {
	// TODO load evm from cache (e.g. file)
	evmObj := evm.EVM{}
	return Debugger{
		evm: evmObj,
	}
}
