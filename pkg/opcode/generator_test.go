package opcode

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func Test_opcode_generator(t *testing.T) {

	/*
		[00]	PUSH1	42
		[02]	PUSH1	00
		[04]	MSTORE
		[05]	PUSH1	20
		[07]	PUSH1	00
		[09]	RETURN
	*/

	data, err := hex.DecodeString("604260005260206000F3")
	if err != nil {
		panic(err)
	}

	var opcode = Opcode{
		HexBytecode: data,
	}
	err = opcode.Generate()
	if err != nil {
		t.Fatal(err.Error())
	}

	value0, _ := hex.DecodeString("42")
	value1, _ := hex.DecodeString("00")
	value3, _ := hex.DecodeString("20")
	value4, _ := hex.DecodeString("00")

	var expectedOutput = [6]Command{
		{Cmd: "PUSH1", Input: value0},
		{Cmd: "PUSH1", Input: value1},
		{Cmd: "MSTORE", Input: nil},
		{Cmd: "PUSH1", Input: value3},
		{Cmd: "PUSH1", Input: value4},
		{Cmd: "RETURN", Input: nil},
	}

	if len(opcode.Output) != len(expectedOutput) {
		fmt.Println("Expected:")
		fmt.Println(expectedOutput)
		fmt.Println("But is:")
		fmt.Println(opcode.Output)
		t.Fatal("opcode does not match")
	}

	for i := 0; i < len(opcode.Output); i++ {
		if opcode.Output[i].Cmd != expectedOutput[i].Cmd {
			t.Fatal("opcode does not match")
		}
		if hex.EncodeToString(opcode.Output[i].Input) != hex.EncodeToString(expectedOutput[i].Input) {
			t.Fatal("opcode does not match")
		}
	}

}
