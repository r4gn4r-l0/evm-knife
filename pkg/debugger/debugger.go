package debugger

import (
	"errors"

	"github.com/r4gn4r-l0/evm-knife/pkg/evm"
)

func (o *Debugger) DeployContract(contract []byte) *evm.Contract {
	contractObj := evm.NewContract(contract)
	evm.GetEVM().AddContract(contractObj.Address, &contractObj)
	return &contractObj
}

func (o *Debugger) StepDebugger(contract *evm.Contract) (bool, error) {
	if len(contract.Bytecode) > 0 {
		code := contract.Bytecode[contract.ProgramCounter]
		finished := contract.ExecuteCode(code, &o.Tx)
		return finished, nil
	} else {
		return true, errors.New("no bytecode given")
	}
}
