package opcode

import (
	"errors"
	"fmt"
)

func (o *Opcode) Generate() error {
	if len(o.HexBytecode) <= 0 {
		return fmt.Errorf("bytecode is not set")
	}
	currentIndex := 0
	for i, b := range o.HexBytecode {
		if i >= currentIndex {
			currentIndex = i
			command, err := getCommand(b, &currentIndex, o.HexBytecode)
			if err != nil {
				return err
			}
			o.Output = append(o.Output, &command)
		}
	}
	return nil
}

func getCommand(b byte, index *int, bytecode []byte) (Command, error) {
	var command = new(Command)
	var err error = nil
	switch b {
	case 0x00:
		command.Cmd = "STOP"
		command.MinGas = 0
	case 0x01:
		command.Cmd = "ADD"
		command.MinGas = 3
	case 0x02:
		command.Cmd = "MUL"
		command.MinGas = 5
	case 0x03:
		command.Cmd = "SUB"
		command.MinGas = 3
	case 0x04:
		command.Cmd = "DIV"
		command.MinGas = 5
	case 0x05:
		command.Cmd = "SDIV"
		command.MinGas = 5
	case 0x06:
		command.Cmd = "MOD"
		command.MinGas = 5
	case 0x07:
		command.Cmd = "SMOD"
		command.MinGas = 5
	case 0x08:
		command.Cmd = "ADDMOD"
		command.MinGas = 8
	case 0x09:
		command.Cmd = "MULMOD"
		command.MinGas = 8
	case 0x0a:
		command.Cmd = "EXP"
		command.MinGas = 10
	case 0x0b:
		command.Cmd = "SIGNEXTEND"
		command.MinGas = 5
	case 0x10:
		command.Cmd = "LT"
		command.MinGas = 3
	case 0x11:
		command.Cmd = "GT"
		command.MinGas = 3
	case 0x12:
		command.Cmd = "SLT"
		command.MinGas = 3
	case 0x13:
		command.Cmd = "SGT"
		command.MinGas = 3
	case 0x14:
		command.Cmd = "EQ"
		command.MinGas = 3
	case 0x15:
		command.Cmd = "ISZERO"
		command.MinGas = 3
	case 0x16:
		command.Cmd = "AND"
		command.MinGas = 3
	case 0x17:
		command.Cmd = "OR"
		command.MinGas = 3
	case 0x18:
		command.Cmd = "XOR"
		command.MinGas = 3
	case 0x19:
		command.Cmd = "NOT"
		command.MinGas = 3
	case 0x1a:
		command.Cmd = "BYTE"
		command.MinGas = 3
	case 0x1b:
		command.Cmd = "SHL"
		command.MinGas = 3
	case 0x1c:
		command.Cmd = "SHR"
		command.MinGas = 3
	case 0x1d:
		command.Cmd = "SAR"
		command.MinGas = 3
	case 0x20:
		command.Cmd = "SHA3"
		command.MinGas = 30
	case 0x30:
		command.Cmd = "ADDRESS"
		command.MinGas = 2
	case 0x31:
		command.Cmd = "BALANCE"
		command.MinGas = 100
	case 0x32:
		command.Cmd = "ORIGIN"
		command.MinGas = 2
	case 0x33:
		command.Cmd = "CALLER"
		command.MinGas = 2
	case 0x34:
		command.Cmd = "CALLVALUE"
		command.MinGas = 2
	case 0x35:
		command.Cmd = "CALLDATALOAD"
		command.MinGas = 3
	case 0x36:
		command.Cmd = "CALLDATASIZE"
		command.MinGas = 3
	case 0x37:
		command.Cmd = "CALLDATACOPY"
		command.MinGas = 3
	case 0x38:
		command.Cmd = "CODESIZE"
		command.MinGas = 2
	case 0x39:
		command.Cmd = "CODECOPY"
		command.MinGas = 3
	case 0x3a:
		command.Cmd = "GASPRICE"
		command.MinGas = 2
	case 0x3b:
		command.Cmd = "EXTCODESIZE"
		command.MinGas = 100
	case 0x3c:
		command.Cmd = "EXTCODECOPY"
		command.MinGas = 100
	case 0x3d:
		command.Cmd = "RETURNDATASIZE"
		command.MinGas = 2
	case 0x3e:
		command.Cmd = "RETURNDATACOPY"
		command.MinGas = 3
	case 0x3f:
		command.Cmd = "EXTCODEHASH"
		command.MinGas = 100
	case 0x40:
		command.Cmd = "BLOCKHASH"
		command.MinGas = 20
	case 0x41:
		command.Cmd = "COINBASE"
		command.MinGas = 2
	case 0x42:
		command.Cmd = "TIMESTAMP"
		command.MinGas = 2
	case 0x43:
		command.Cmd = "NUMBER"
		command.MinGas = 2
	case 0x44:
		command.Cmd = "PREVRANDAO"
		command.MinGas = 2
	case 0x45:
		command.Cmd = "GASLIMIT"
		command.MinGas = 2
	case 0x46:
		command.Cmd = "CHAINID"
		command.MinGas = 2
	case 0x47:
		command.Cmd = "SELFBALANCE"
		command.MinGas = 5
	case 0x48:
		command.Cmd = "BASEFEE"
		command.MinGas = 2
	case 0x50:
		command.Cmd = "POP"
		command.MinGas = 2
	case 0x51:
		command.Cmd = "MLOAD"
		command.MinGas = 2
	case 0x52:
		command.Cmd = "MSTORE"
		command.MinGas = 3
	case 0x53:
		command.Cmd = "MSTORE8"
		command.MinGas = 3
	case 0x54:
		command.Cmd = "SLOAD"
		command.MinGas = 100
	case 0x55:
		command.Cmd = "SSTORE"
		command.MinGas = 100
	case 0x56:
		command.Cmd = "JUMP"
		command.MinGas = 8
	case 0x57:
		command.Cmd = "JUMPI"
		command.MinGas = 10
	case 0x58:
		command.Cmd = "PC"
		command.MinGas = 2
	case 0x59:
		command.Cmd = "MSIZE"
		command.MinGas = 2
	case 0x5a:
		command.Cmd = "GAS"
		command.MinGas = 2
	case 0x5B:
		command.Cmd = "JUMPDEST"
		command.MinGas = 1
	case 0x60:
		command.Cmd = "PUSH1"
		command.Input = getInputValue(index, 1, bytecode)
		command.MinGas = 3
	case 0x61:
		command.Cmd = "PUSH2"
		command.Input = getInputValue(index, 2, bytecode)
		command.MinGas = 3
	case 0x62:
		command.Cmd = "PUSH3"
		command.Input = getInputValue(index, 3, bytecode)
		command.MinGas = 3
	case 0x63:
		command.Cmd = "PUSH4"
		command.Input = getInputValue(index, 4, bytecode)
		command.MinGas = 3
	case 0x64:
		command.Cmd = "PUSH5"
		command.Input = getInputValue(index, 5, bytecode)
		command.MinGas = 3
	case 0x65:
		command.Cmd = "PUSH6"
		command.Input = getInputValue(index, 6, bytecode)
		command.MinGas = 3
	case 0x66:
		command.Cmd = "PUSH7"
		command.Input = getInputValue(index, 7, bytecode)
		command.MinGas = 3
	case 0x67:
		command.Cmd = "PUSH8"
		command.Input = getInputValue(index, 8, bytecode)
		command.MinGas = 3
	case 0x68:
		command.Cmd = "PUSH9"
		command.Input = getInputValue(index, 9, bytecode)
		command.MinGas = 3
	case 0x69:
		command.Cmd = "PUSH10"
		command.Input = getInputValue(index, 10, bytecode)
		command.MinGas = 3
	case 0x6a:
		command.Cmd = "PUSH11"
		command.Input = getInputValue(index, 11, bytecode)
		command.MinGas = 3
	case 0x6b:
		command.Cmd = "PUSH12"
		command.Input = getInputValue(index, 12, bytecode)
		command.MinGas = 3
	case 0x6c:
		command.Cmd = "PUSH13"
		command.Input = getInputValue(index, 13, bytecode)
		command.MinGas = 3
	case 0x6d:
		command.Cmd = "PUSH14"
		command.Input = getInputValue(index, 14, bytecode)
		command.MinGas = 3
	case 0x6e:
		command.Cmd = "PUSH15"
		command.Input = getInputValue(index, 15, bytecode)
		command.MinGas = 3
	case 0x6f:
		command.Cmd = "PUSH16"
		command.Input = getInputValue(index, 16, bytecode)
		command.MinGas = 3
	case 0x70:
		command.Cmd = "PUSH17"
		command.Input = getInputValue(index, 17, bytecode)
		command.MinGas = 3
	case 0x71:
		command.Cmd = "PUSH18"
		command.Input = getInputValue(index, 18, bytecode)
		command.MinGas = 3
	case 0x72:
		command.Cmd = "PUSH19"
		command.Input = getInputValue(index, 19, bytecode)
		command.MinGas = 3
	case 0x73:
		command.Cmd = "PUSH20"
		command.Input = getInputValue(index, 20, bytecode)
		command.MinGas = 3
	case 0x74:
		command.Cmd = "PUSH21"
		command.Input = getInputValue(index, 21, bytecode)
		command.MinGas = 3
	case 0x75:
		command.Cmd = "PUSH22"
		command.Input = getInputValue(index, 22, bytecode)
		command.MinGas = 3
	case 0x76:
		command.Cmd = "PUSH23"
		command.Input = getInputValue(index, 23, bytecode)
		command.MinGas = 3
	case 0x77:
		command.Cmd = "PUSH24"
		command.Input = getInputValue(index, 24, bytecode)
		command.MinGas = 3
	case 0x78:
		command.Cmd = "PUSH25"
		command.Input = getInputValue(index, 25, bytecode)
		command.MinGas = 3
	case 0x79:
		command.Cmd = "PUSH26"
		command.Input = getInputValue(index, 26, bytecode)
		command.MinGas = 3
	case 0x7a:
		command.Cmd = "PUSH27"
		command.Input = getInputValue(index, 27, bytecode)
		command.MinGas = 3
	case 0x7b:
		command.Cmd = "PUSH28"
		command.Input = getInputValue(index, 28, bytecode)
		command.MinGas = 3
	case 0x7c:
		command.Cmd = "PUSH29"
		command.Input = getInputValue(index, 29, bytecode)
		command.MinGas = 3
	case 0x7d:
		command.Cmd = "PUSH30"
		command.Input = getInputValue(index, 30, bytecode)
		command.MinGas = 3
	case 0x7e:
		command.Cmd = "PUSH31"
		command.Input = getInputValue(index, 31, bytecode)
		command.MinGas = 3
	case 0x7f:
		command.Cmd = "PUSH32"
		command.Input = getInputValue(index, 32, bytecode)
		command.MinGas = 3
	case 0x80:
		command.Cmd = "DUP1"
		command.MinGas = 3
	case 0x81:
		command.Cmd = "DUP2"
		command.MinGas = 3
	case 0x82:
		command.Cmd = "DUP3"
		command.MinGas = 3
	case 0x83:
		command.Cmd = "DUP4"
		command.MinGas = 3
	case 0x84:
		command.Cmd = "DUP5"
		command.MinGas = 3
	case 0x85:
		command.Cmd = "DUP6"
		command.MinGas = 3
	case 0x86:
		command.Cmd = "DUP7"
		command.MinGas = 3
	case 0x87:
		command.Cmd = "DUP8"
		command.MinGas = 3
	case 0x88:
		command.Cmd = "DUP9"
		command.MinGas = 3
	case 0x89:
		command.Cmd = "DUP10"
		command.MinGas = 3
	case 0x8a:
		command.Cmd = "DUP11"
		command.MinGas = 3
	case 0x8b:
		command.Cmd = "DUP12"
		command.MinGas = 3
	case 0x8c:
		command.Cmd = "DUP13"
		command.MinGas = 3
	case 0x8d:
		command.Cmd = "DUP14"
		command.MinGas = 3
	case 0x8e:
		command.Cmd = "DUP15"
		command.MinGas = 3
	case 0x8f:
		command.Cmd = "DUP16"
		command.MinGas = 3
	case 0x90:
		command.Cmd = "SWAP1"
		command.MinGas = 3
	case 0x91:
		command.Cmd = "SWAP2"
		command.MinGas = 3
	case 0x92:
		command.Cmd = "SWAP3"
		command.MinGas = 3
	case 0x93:
		command.Cmd = "SWAP4"
		command.MinGas = 3
	case 0x94:
		command.Cmd = "SWAP5"
		command.MinGas = 3
	case 0x95:
		command.Cmd = "SWAP6"
		command.MinGas = 3
	case 0x96:
		command.Cmd = "SWAP7"
		command.MinGas = 3
	case 0x97:
		command.Cmd = "SWAP8"
		command.MinGas = 3
	case 0x98:
		command.Cmd = "SWAP9"
		command.MinGas = 3
	case 0x99:
		command.Cmd = "SWAP10"
		command.MinGas = 3
	case 0x9a:
		command.Cmd = "SWAP11"
		command.MinGas = 3
	case 0x9b:
		command.Cmd = "SWAP12"
		command.MinGas = 3
	case 0x9c:
		command.Cmd = "SWAP13"
		command.MinGas = 3
	case 0x9d:
		command.Cmd = "SWAP14"
		command.MinGas = 3
	case 0x9e:
		command.Cmd = "SWAP15"
		command.MinGas = 3
	case 0x9f:
		command.Cmd = "SWAP16"
		command.MinGas = 3
	case 0xa0:
		command.Cmd = "LOG0"
		command.MinGas = 375
	case 0xa1:
		command.Cmd = "LOG1"
		command.MinGas = 750
	case 0xa2:
		command.Cmd = "LOG2"
		command.MinGas = 1125
	case 0xa3:
		command.Cmd = "LOG3"
		command.MinGas = 1500
	case 0xa4:
		command.Cmd = "LOG4"
		command.MinGas = 1875
	case 0xf0:
		command.Cmd = "CREATE"
		command.MinGas = 32000
	case 0xf1:
		command.Cmd = "CALL"
		command.MinGas = 100
	case 0xf2:
		command.Cmd = "CALLCODE"
		command.MinGas = 100
	case 0xf3:
		command.Cmd = "RETURN"
		command.MinGas = 0
	case 0xf4:
		command.Cmd = "DELEGATECALL"
		command.MinGas = 100
	case 0xf5:
		command.Cmd = "CREATE2"
		command.MinGas = 32000
	case 0xfa:
		command.Cmd = "STATICCALL"
		command.MinGas = 100
	case 0xfd:
		command.Cmd = "REVERT"
		command.MinGas = 0
	case 0xfe:
		command.Cmd = "INVALID"
		command.MinGas = 0
	case 0xff:
		command.Cmd = "SELFDESTRUCT"
		command.MinGas = 5000
	default:
		err = errors.New("invalid byte")
	}
	return *command, err
}

func getInputValue(index *int, bytes int, bytecode []byte) []byte {
	var response = make([]byte, bytes)
	for i := 1; i <= bytes; i++ {
		response[i-1] = bytecode[*index+i]
	}
	*index += bytes + 1
	return response
}
