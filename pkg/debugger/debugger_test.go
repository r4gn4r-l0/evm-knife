package debugger

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func Test_PushAndMstore(t *testing.T) {
	data, err := hex.DecodeString("62424240600252")
	if err != nil {
		t.Error(err)
	}
	debugger := Debugger{
		Bytecode: data,
	}
	debugger.StepDebugger()
	fmt.Println("\nNext step:")
	fmt.Print("Stack:\t")
	fmt.Println(debugger.Stack)
	fmt.Print("Memory:\t")
	fmt.Printf("0x%x\n", debugger.Memory)

	debugger.StepDebugger()
	fmt.Println("\nNext step:")
	fmt.Print("Stack:\t")
	fmt.Println(debugger.Stack)
	fmt.Print("Memory:\t")
	fmt.Printf("0x%x\n", debugger.Memory)

	debugger.StepDebugger()
	fmt.Println("\nNext step:")
	fmt.Print("Stack:\t")
	fmt.Println(debugger.Stack)
	fmt.Print("Memory:\t")
	fmt.Printf("0x%x\n", debugger.Memory)
	expected := [64]byte{
		byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x42), byte(0x42), byte(0x40), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00),
	}
	fmt.Println("\n\n\n==============\nexpected:")
	fmt.Printf("0x%x\n", expected)
	fmt.Println("is:")
	fmt.Printf("0x%x\n", debugger.Memory)
	if debugger.Memory[0] != expected[0] ||
		// offset 02 => array 31+2; opcode 62 (PUSH3) => PUSH 3 bytes
		debugger.Memory[31] != expected[31] || // value = 0x00
		debugger.Memory[32] != expected[32] || // value = 0x42
		debugger.Memory[33] != expected[33] || // value = 0x42
		debugger.Memory[34] != expected[34] || // value = 0x40
		debugger.Memory[35] != expected[35] { // value = 0x00
		t.Fail()
	}
}
