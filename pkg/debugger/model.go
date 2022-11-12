package debugger

type Debugger struct {
	Bytecode       []byte
	ProgramCounter int16    // evm smart contract has max. size of 24kb => 24.576b
	Stack          [][]byte // reverse stack => len(Stack)-1 is top
	Memory         []byte
	Storage        [][]byte
}
