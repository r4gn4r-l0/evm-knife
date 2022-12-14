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

func (o *Debugger) StepDebugger(contract *evm.Contract, ctx *evm.Context) (bool, error) {
	if len(contract.Bytecode) > 0 {
		code := contract.Bytecode[ctx.ProgramCounter]
		finished, err := contract.ExecuteCode(code, &o.Tx, ctx)
		return finished, err
	} else {
		return true, errors.New("no bytecode given")
	}
}
