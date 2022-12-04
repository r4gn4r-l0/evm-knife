package evm

func (ctx *Context) execute(contract *Contract, tx *Tx) bool {
	if len(contract.Bytecode) > 0 {
		finished := false
		var err error
		for !finished && int(ctx.ProgramCounter) != len(contract.Bytecode) {
			code := contract.Bytecode[ctx.ProgramCounter]
			finished, err = contract.ExecuteCode(code, tx, ctx)
			if err != nil {
				return false
			}
		}
		return true
	}
	return false
}
