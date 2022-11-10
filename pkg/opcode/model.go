package opcode

import (
	"encoding/hex"
	"fmt"
)

type Opcode struct {
	HexBytecode []byte
	Output      []*Command
}

type Command struct {
	Cmd    string
	Input  []byte
	MinGas int
}

func (o *Opcode) ToString() string {
	var output string = ""
	for i, line := range o.Output {
		output += fmt.Sprintf("0x%x", i) + "\t" + line.Cmd + "\t\t" + hex.EncodeToString(line.Input) + "\n"
	}
	return output
}
