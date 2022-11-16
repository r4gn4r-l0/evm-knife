package debugger

import (
	"errors"

	"github.com/holiman/uint256"
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
		a := new(uint256.Int).SetBytes(o.stackPop())
		b := new(uint256.Int).SetBytes(o.stackPop())
		x := a.Add(a, b)
		// correctOverflow(x)
		o.stackPush(x.Bytes())
	case code == 0x02: // MUL (multiplication)
		a := new(uint256.Int).SetBytes(o.stackPop())
		b := new(uint256.Int).SetBytes(o.stackPop())
		x := a.Mul(a, b)
		// correctOverflow(x)
		o.stackPush(x.Bytes())
	case code == 0x03: // SUB
		a := new(uint256.Int).SetBytes(o.stackPop())
		b := new(uint256.Int).SetBytes(o.stackPop())
		x := a.Sub(a, b)
		// correctUnderflow(x)
		o.stackPush(x.Bytes())
	case code == 0x04: // DIV
		a := new(uint256.Int).SetBytes(o.stackPop())
		b := new(uint256.Int).SetBytes(o.stackPop())
		x := new(uint256.Int).Div(a, b)
		// correctUnderflow(x)
		o.stackPush(x.Bytes())
	case code == 0x05: // SDIV (signed div)
		a := new(uint256.Int).SetBytes(o.stackPop())
		b := new(uint256.Int).SetBytes(o.stackPop())
		x := new(uint256.Int).SDiv(a, b)
		o.stackPush(x.Bytes())
	case code == 0x06: // MOD
		a := new(uint256.Int).SetBytes(o.stackPop())
		b := new(uint256.Int).SetBytes(o.stackPop())
		x := new(uint256.Int).Mod(a, b)
		o.stackPush(x.Bytes())
	case code == 0x07: // SMOD
		a := new(uint256.Int).SetBytes(o.stackPop())
		b := new(uint256.Int).SetBytes(o.stackPop())
		x := new(uint256.Int).SMod(a, b)
		o.stackPush(x.Bytes())
	case code == 0x08: // ADDMOD
		a := new(uint256.Int).SetBytes(o.stackPop())
		b := new(uint256.Int).SetBytes(o.stackPop())
		c := new(uint256.Int).SetBytes(o.stackPop())
		x := new(uint256.Int).AddMod(a, b, c)
		o.stackPush(x.Bytes())
	case code == 0x09: // MULMOD
		a := new(uint256.Int).SetBytes(o.stackPop())
		b := new(uint256.Int).SetBytes(o.stackPop())
		c := new(uint256.Int).SetBytes(o.stackPop())
		x := new(uint256.Int).MulMod(a, b, c)
		o.stackPush(x.Bytes())
	case code == 0x0a: // EXP
		a := new(uint256.Int).SetBytes(o.stackPop())
		b := new(uint256.Int).SetBytes(o.stackPop())
		x := new(uint256.Int).Exp(a, b)
		o.stackPush(x.Bytes())
	case code == 0x0b: // SIGNEXTEND
		a := new(uint256.Int).SetBytes(o.stackPop())
		b := new(uint256.Int).SetBytes(o.stackPop())
		x := new(uint256.Int).ExtendSign(a, b)
		o.stackPush(x.Bytes())
	case code == 0x10: // LT
		a := new(uint256.Int).SetBytes(o.stackPop())
		b := new(uint256.Int).SetBytes(o.stackPop())
		x := a.Lt(b)
		if x {
			o.stackPush([]byte{0x01})
		} else {
			o.stackPush([]byte{0x00})
		}
	case code == 0x11: // GT
		a := new(uint256.Int).SetBytes(o.stackPop())
		b := new(uint256.Int).SetBytes(o.stackPop())
		x := a.Gt(b)
		if x {
			o.stackPush([]byte{0x01})
		} else {
			o.stackPush([]byte{0x00})
		}
	case code == 0x12: // SLT
		a := new(uint256.Int).SetBytes(o.stackPop())
		b := new(uint256.Int).SetBytes(o.stackPop())
		x := a.Slt(b)
		if x {
			o.stackPush([]byte{0x01})
		} else {
			o.stackPush([]byte{0x00})
		}
	case code == 0x13: // SGT
		a := new(uint256.Int).SetBytes(o.stackPop())
		b := new(uint256.Int).SetBytes(o.stackPop())
		x := a.Sgt(b)
		if x {
			o.stackPush([]byte{0x01})
		} else {
			o.stackPush([]byte{0x00})
		}
	case code == 0x14: // EQ
		a := new(uint256.Int).SetBytes(o.stackPop())
		b := new(uint256.Int).SetBytes(o.stackPop())
		x := a.Eq(b)
		if x {
			o.stackPush([]byte{0x01})
		} else {
			o.stackPush([]byte{0x00})
		}
	case code == 0x15: // ISZERO
		a := new(uint256.Int).SetBytes(o.stackPop())
		b := new(uint256.Int).SetBytes([]byte{0x00})
		x := a.Eq(b)
		if x {
			o.stackPush([]byte{0x01})
		} else {
			o.stackPush([]byte{0x00})
		}
	case code == 0x16: // AND
		a := new(uint256.Int).SetBytes(o.stackPop())
		b := new(uint256.Int).SetBytes(o.stackPop())
		x := new(uint256.Int).And(a, b)
		o.stackPush(x.Bytes())
	case code == 0x17: // OR
		a := new(uint256.Int).SetBytes(o.stackPop())
		b := new(uint256.Int).SetBytes(o.stackPop())
		x := new(uint256.Int).Or(a, b)
		o.stackPush(x.Bytes())
	case code == 0x18: // XOR
		a := new(uint256.Int).SetBytes(o.stackPop())
		b := new(uint256.Int).SetBytes(o.stackPop())
		x := new(uint256.Int).Xor(a, b)
		o.stackPush(x.Bytes())
	case code == 0x19: // NOT
		a := new(uint256.Int).SetBytes(o.stackPop())
		x := new(uint256.Int).Not(a)
		o.stackPush(x.Bytes())
	case code == 0x1a: // BYTE
		a := new(uint256.Int).SetBytes(o.stackPop())
		b := new(uint256.Int).SetBytes(o.stackPop())
		x := b.Byte(a)
		o.stackPush(x.Bytes())
	case code == 0x1b: // SHL
		a := new(uint256.Int).SetBytes(o.stackPop())
		b := new(uint256.Int).SetBytes(o.stackPop())
		if a.LtUint64(256) {
			b = b.Lsh(b, uint(a.Uint64()))
		} else {
			b = b.Clear()
		}
		o.stackPush(b.Bytes())
	case code == 0x1c: // SHR
		a := new(uint256.Int).SetBytes(o.stackPop())
		b := new(uint256.Int).SetBytes(o.stackPop())
		if a.LtUint64(256) {
			b = b.Rsh(b, uint(a.Uint64()))
		} else {
			b = b.Clear()
		}
		o.stackPush(b.Bytes())
	case code == 0x1d: // SAR
		a := new(uint256.Int).SetBytes(o.stackPop())
		b := new(uint256.Int).SetBytes(o.stackPop())
		if a.GtUint64(256) {
			if b.Sign() >= 0 {
				b = b.Clear()
			} else {
				// Max negative shift: all bits set
				b = b.SetAllOne()
			}
		} else {
			n := uint(a.Uint64())
			b = b.SRsh(b, n)
		}
		o.stackPush(b.Bytes())
	case code >= 0x60 && code <= 0x7f: // PUSHx
		to := int16(code) - 0x5e
		value := o.Bytecode[(o.ProgramCounter + 0x01):(o.ProgramCounter + to)]
		o.ProgramCounter = o.ProgramCounter + to
		o.stackPush(value)
	case code == 0x52: // MSTORE
		byteArr := o.stackPop()
		offset := new(uint256.Int).SetBytes(byteArr)
		value := o.stackPop()
		byteStart := int64(offset.Uint64()) + int64(0x20) - 1
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
	if len(value) == 0 {
		value = []byte{0x00}
	}
	o.Stack = append(o.Stack, value)
}

func (o *Debugger) stackPop() []byte {
	top := len(o.Stack) - 1
	value := o.Stack[top]
	o.Stack = o.Stack[:top] // remove from stack
	return value
}
