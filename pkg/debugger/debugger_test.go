package debugger

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func Test_Add(t *testing.T) {
	data, err := hex.DecodeString("6001600101")
	if err != nil {
		t.Error(err)
	}
	debugger := Debugger{
		Bytecode: data,
	}
	debugger.StepDebugger()
	debugger.StepDebugger()
	debugger.StepDebugger()
	should := [1]byte{0x02}
	if should[0] != debugger.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", debugger.Stack[0])
		t.Fail()
	}
}

func Test_AddOverflow(t *testing.T) {
	/* testcase
	x := 0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff + 0x01
	expected
	x == 0x00
	*/
	data, err := hex.DecodeString("7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600101")
	if err != nil {
		t.Error(err)
	}
	debugger := Debugger{
		Bytecode: data,
	}
	debugger.StepDebugger()
	debugger.StepDebugger()
	debugger.StepDebugger()
	should := [1]byte{0x00}
	if should[0] != debugger.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", debugger.Stack[0])
		t.Fail()
	}
}

func Test_AddBigNumber(t *testing.T) {
	/* test:
	x := 0x11111111111111111111111111111111111111111111111111111111111111ff + 0x01
	should be
	x == 0x1111111111111111111111111111111111111111111111111111111111111200
	*/
	data, err := hex.DecodeString("7f11111111111111111111111111111111111111111111111111111111111111ff600101")
	if err != nil {
		t.Error(err)
	}
	debugger := Debugger{
		Bytecode: data,
	}
	debugger.StepDebugger()
	debugger.StepDebugger()
	debugger.StepDebugger()
	should := [32]byte{0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x12, 0x00}
	if should[0] != debugger.Stack[0][0] || should[31] != debugger.Stack[0][31] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", debugger.Stack[0])
		t.Fail()
	}
}

func Test_Mul(t *testing.T) {
	/* testcase
	x := 0x02 * 0x02
	expected
	x == 0x04
	*/
	data, err := hex.DecodeString("6002600202")
	if err != nil {
		t.Error(err)
	}
	debugger := Debugger{
		Bytecode: data,
	}
	debugger.StepDebugger()
	debugger.StepDebugger()
	debugger.StepDebugger()
	should := [1]byte{0x04}
	if should[0] != debugger.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", debugger.Stack[0])
		t.Fail()
	}
}

func Test_MulOverflow(t *testing.T) {
	/* testcase
	x := 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff * 0x02
	expected
	x == 0x04
	*/
	data, err := hex.DecodeString("7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600202")
	if err != nil {
		t.Error(err)
	}
	debugger := Debugger{
		Bytecode: data,
	}
	debugger.StepDebugger()
	debugger.StepDebugger()
	debugger.StepDebugger()
	should := [32]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}
	if should[0] != debugger.Stack[0][0] || should[31] != debugger.Stack[0][31] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", debugger.Stack[0])
		t.Fail()
	}
}

func Test_Sub(t *testing.T) {
	/* testcase
	x := 0x06 - 0x02
	expected
	x == 0x04
	*/
	data, err := hex.DecodeString("6002600603")
	if err != nil {
		t.Error(err)
	}
	debugger := Debugger{
		Bytecode: data,
	}
	debugger.StepDebugger()
	debugger.StepDebugger()
	debugger.StepDebugger()
	should := [1]byte{0x04}
	if should[0] != debugger.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", debugger.Stack[0])
		t.Fail()
	}
}

func Test_SubUnderFlow(t *testing.T) {
	/* testcase
	x := 0x00 - 0x02
	expected
	x == 0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe
	*/
	data, err := hex.DecodeString("6002600003")
	if err != nil {
		t.Error(err)
	}
	debugger := Debugger{
		Bytecode: data,
	}
	debugger.StepDebugger()
	debugger.StepDebugger()
	debugger.StepDebugger()
	should := [32]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}
	if should[0] != debugger.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", debugger.Stack[0])
		t.Fail()
	}
}

func Test_Div(t *testing.T) {
	/* testcase
	x := 0x06 / 0x02
	expected
	x == 0x03
	*/
	data, err := hex.DecodeString("6002600604")
	if err != nil {
		t.Error(err)
	}
	debugger := Debugger{
		Bytecode: data,
	}
	debugger.StepDebugger()
	debugger.StepDebugger()
	debugger.StepDebugger()
	should := [1]byte{0x03}
	if should[0] != debugger.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", debugger.Stack[0])
		t.Fail()
	}
}

func Test_DivFloatingPoint(t *testing.T) {
	/* testcase
	x := 0x01 / 0x02
	expected
	x == 0x00
	*/
	data, err := hex.DecodeString("6002600104")
	if err != nil {
		t.Error(err)
	}
	debugger := Debugger{
		Bytecode: data,
	}
	debugger.StepDebugger()
	debugger.StepDebugger()
	debugger.StepDebugger()
	should := [1]byte{0x00}
	if should[0] != debugger.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", debugger.Stack[0])
		t.Fail()
	}
}

func Test_DivFloatingPoint2(t *testing.T) {
	/* testcase
	x := 0x05 / 0x02
	expected
	x == 0x02
	*/
	data, err := hex.DecodeString("6002600504")
	if err != nil {
		t.Error(err)
	}
	debugger := Debugger{
		Bytecode: data,
	}
	debugger.StepDebugger()
	debugger.StepDebugger()
	debugger.StepDebugger()
	should := [1]byte{0x02}
	if should[0] != debugger.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", debugger.Stack[0])
		t.Fail()
	}
}

func Test_SDiv(t *testing.T) {
	/* testcase
	x := 0x10 / 0x10
	expected
	x == 0x01
	*/
	data, err := hex.DecodeString("600a600a05")
	if err != nil {
		t.Error(err)
	}
	debugger := Debugger{
		Bytecode: data,
	}
	debugger.StepDebugger()
	debugger.StepDebugger()
	debugger.StepDebugger()
	should := [1]byte{0x01}
	if should[0] != debugger.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", debugger.Stack[0])
		t.Fail()
	}
}

func Test_SDivOverflow(t *testing.T) {
	/* testcase
	x := 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFE / 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF
	expected
	x == 0x02
	*/
	data, err := hex.DecodeString("7fFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF7fFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFE05")
	if err != nil {
		t.Error(err)
	}
	debugger := Debugger{
		Bytecode: data,
	}
	debugger.StepDebugger()
	debugger.StepDebugger()
	debugger.StepDebugger()
	should := [1]byte{0x02}
	if should[0] != debugger.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", debugger.Stack[0])
		t.Fail()
	}
}

func Test_Mod(t *testing.T) {
	/* testcase
	x := 0x0a % 0x03
	expected
	x == 0x01
	*/
	data, err := hex.DecodeString("6003600a06")
	if err != nil {
		t.Error(err)
	}
	debugger := Debugger{
		Bytecode: data,
	}
	debugger.StepDebugger()
	debugger.StepDebugger()
	debugger.StepDebugger()
	should := [1]byte{0x01}
	if should[0] != debugger.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", debugger.Stack[0])
		t.Fail()
	}
}

func Test_Mod2(t *testing.T) {
	/* testcase
	x := 0x10 % 0x05
	expected
	x == 0x02
	*/
	data, err := hex.DecodeString("6005601106")
	if err != nil {
		t.Error(err)
	}
	debugger := Debugger{
		Bytecode: data,
	}
	debugger.StepDebugger()
	debugger.StepDebugger()
	debugger.StepDebugger()
	should := [1]byte{0x02}
	if should[0] != debugger.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", debugger.Stack[0])
		t.Fail()
	}
}

func Test_SMod(t *testing.T) {
	/* testcase
	x := 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF8 % 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFD
	expected
	x == 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFE
	*/
	data, err := hex.DecodeString("7fFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFD7fFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF807")
	if err != nil {
		t.Error(err)
	}
	debugger := Debugger{
		Bytecode: data,
	}
	debugger.StepDebugger()
	debugger.StepDebugger()
	debugger.StepDebugger()
	should := [32]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}
	if should[0] != debugger.Stack[0][0] || should[31] != debugger.Stack[0][31] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", debugger.Stack[0])
		t.Fail()
	}
}

func Test_AddMod(t *testing.T) {
	/* testcase
	x := (0x0a + 0x0a) % 0x08
	expected
	x == 0x4
	*/
	data, err := hex.DecodeString("6008600a600a08")
	if err != nil {
		t.Error(err)
	}
	debugger := Debugger{
		Bytecode: data,
	}
	debugger.StepDebugger()
	debugger.StepDebugger()
	debugger.StepDebugger()
	debugger.StepDebugger()
	should := [1]byte{0x04}
	if should[0] != debugger.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", debugger.Stack[0])
		t.Fail()
	}
}

func Test_MulMod(t *testing.T) {
	/* testcase
	x := (0x0a * 0x0a) % 0x08
	expected
	x == 0x4
	*/
	data, err := hex.DecodeString("6008600a600a09")
	if err != nil {
		t.Error(err)
	}
	debugger := Debugger{
		Bytecode: data,
	}
	debugger.StepDebugger()
	debugger.StepDebugger()
	debugger.StepDebugger()
	debugger.StepDebugger()
	should := [1]byte{0x04}
	if should[0] != debugger.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", debugger.Stack[0])
		t.Fail()
	}
}

func Test_SIGNEXTEND(t *testing.T) {
	data, err := hex.DecodeString("600060ff0b")
	if err != nil {
		t.Error(err)
	}
	debugger := Debugger{
		Bytecode: data,
	}
	debugger.StepDebugger()
	debugger.StepDebugger()
	debugger.StepDebugger()
	should := [32]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	if should[0] != debugger.Stack[0][0] || should[31] != debugger.Stack[0][31] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", debugger.Stack[0])
		t.Fail()
	}
}

func Test_LT(t *testing.T) {
	data, err := hex.DecodeString("600a600910")
	if err != nil {
		t.Error(err)
	}
	debugger := Debugger{
		Bytecode: data,
	}
	debugger.StepDebugger()
	debugger.StepDebugger()
	debugger.StepDebugger()
	should := [1]byte{0x01}
	if should[0] != debugger.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", debugger.Stack[0])
		t.Fail()
	}
}

func Test_GT(t *testing.T) {
	data, err := hex.DecodeString("6009600a11")
	if err != nil {
		t.Error(err)
	}
	debugger := Debugger{
		Bytecode: data,
	}
	debugger.StepDebugger()
	debugger.StepDebugger()
	debugger.StepDebugger()
	should := [1]byte{0x01}
	if should[0] != debugger.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", debugger.Stack[0])
		t.Fail()
	}
}

func Test_PushAndMstore(t *testing.T) {
	data, err := hex.DecodeString("62424240600252")
	if err != nil {
		t.Error(err)
	}
	debugger := Debugger{
		Bytecode: data,
	}
	debugger.StepDebugger()
	debugger.StepDebugger()
	debugger.StepDebugger()
	expected := [64]byte{
		byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x42), byte(0x42), byte(0x40), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00),
	}
	if debugger.Memory[0] != expected[0] ||
		// offset 02 => array 31+2; opcode 62 (PUSH3) => PUSH 3 bytes
		debugger.Memory[31] != expected[31] || // value = 0x00
		debugger.Memory[32] != expected[32] || // value = 0x42
		debugger.Memory[33] != expected[33] || // value = 0x42
		debugger.Memory[34] != expected[34] || // value = 0x40
		debugger.Memory[35] != expected[35] { // value = 0x00
		fmt.Printf("expected: 0x%x\n", expected)
		fmt.Printf("is: 0x%x\n", debugger.Memory)
		t.Fail()
	}
}
