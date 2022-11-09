package opcode

import "encoding/hex"

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
	for _, line := range o.Output {
		output += line.Cmd + "\t" + hex.EncodeToString(line.Input) + "\n"
	}
	return output
}
