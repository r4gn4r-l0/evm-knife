package debugger

import (
	"errors"
	"math/big"
)

func (o *Debugger) StepDebugger() error {
	if len(o.Bytecode) > 0 {
		code := o.Bytecode[o.ProgramCounter]
		o.executeCode(code)
	} else {
		return errors.New("no bytecode given")
	}
	return nil
}

func (o *Debugger) executeCode(code byte) {
	switch {
	case code >= 0x60 && code < 0x7f: // PUSHx
		to := int16(code) - 0x5e
		value := o.Bytecode[(o.ProgramCounter + 0x01):(o.ProgramCounter + to)]
		o.ProgramCounter = o.ProgramCounter + to
		o.push(value)
	case code == 0x52: // MSTORE
		byteArr := o.pop()
		offset := new(big.Int).SetBytes(byteArr)
		value := o.pop()
		// offset.Add(offset, big.NewInt(1))
		byteStart := offset.Int64() + int64(0x20) - 1
		words := byteStart / 0x20
		if byteStart%0x20 > 0 {
			words += 1
		}
		memSize := words * 0x20
		if int64(len(o.Memory)) < memSize {
			appendSize := memSize - int64(len(o.Memory))
			appendArray := make([]byte, appendSize)
			o.Memory = append(o.Memory, appendArray...)
		}
		for i, val := range value {
			o.Memory[(byteStart - int64(i))] = val
		}
	}

}

func (o *Debugger) push(value []byte) {
	o.Stack = append(o.Stack, value)
}

func (o *Debugger) pop() []byte {
	top := len(o.Stack) - 1
	value := o.Stack[top]
	o.Stack = o.Stack[:top] // remove from stack
	return value
}
