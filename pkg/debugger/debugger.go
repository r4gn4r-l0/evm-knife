package debugger

import (
	"errors"
	"math/big"
)

func (o *Debugger) StepDebugger() (bool, error) {
	if len(o.Bytecode) > 0 {
		code := o.Bytecode[o.ProgramCounter]
		finished := o.executeCode(code)
		return finished, nil
	} else {
		return true, errors.New("no bytecode given")
	}
}

func (o *Debugger) executeCode(code byte) bool {
	switch {
	case code == 0x00: // STOP
		return true
	case code == 0x01: // ADD
		a := new(big.Int).SetBytes(o.stackPop())
		b := new(big.Int).SetBytes(o.stackPop())
		x := a.Add(a, b)
		o.stackPush(x.Bytes())
	case code >= 0x60 && code <= 0x7f: // PUSHx
		to := int16(code) - 0x5e
		value := o.Bytecode[(o.ProgramCounter + 0x01):(o.ProgramCounter + to)]
		o.ProgramCounter = o.ProgramCounter + to
		o.stackPush(value)
	case code == 0x52: // MSTORE
		byteArr := o.stackPop()
		offset := new(big.Int).SetBytes(byteArr)
		value := o.stackPop()
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
			o.Memory[(byteStart - int64(len(value)-1) + int64(i))] = val
		}
	}
	return false
}

func (o *Debugger) stackPush(value []byte) {
	o.Stack = append(o.Stack, value)
}

func (o *Debugger) stackPop() []byte {
	top := len(o.Stack) - 1
	value := o.Stack[top]
	o.Stack = o.Stack[:top] // remove from stack
	return value
}
