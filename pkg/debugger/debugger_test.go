package debugger

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/r4gn4r-l0/evm-knife/pkg/evm"
)

func Test_Add(t *testing.T) {
	data, err := hex.DecodeString("6001600101")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)

	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [1]byte{0x02}
	if should[0] != contract.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
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
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [1]byte{0x00}
	if should[0] != contract.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
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
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [32]byte{0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x12, 0x00}
	if should[0] != contract.Stack[0][0] || should[31] != contract.Stack[0][31] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
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
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [1]byte{0x04}
	if should[0] != contract.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
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
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [32]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}
	if should[0] != contract.Stack[0][0] || should[31] != contract.Stack[0][31] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
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
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [1]byte{0x04}
	if should[0] != contract.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
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
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [32]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}
	if should[0] != contract.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
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
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [1]byte{0x03}
	if should[0] != contract.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
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
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [1]byte{0x00}
	if should[0] != contract.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
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
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [1]byte{0x02}
	if should[0] != contract.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
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
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [1]byte{0x01}
	if should[0] != contract.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
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
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [1]byte{0x02}
	if should[0] != contract.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
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
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [1]byte{0x01}
	if should[0] != contract.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
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
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [1]byte{0x02}
	if should[0] != contract.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
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
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [32]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}
	if should[0] != contract.Stack[0][0] || should[31] != contract.Stack[0][31] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
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
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [1]byte{0x04}
	if should[0] != contract.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
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
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [1]byte{0x04}
	if should[0] != contract.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
		t.Fail()
	}
}

func Test_SIGNEXTEND(t *testing.T) {
	data, err := hex.DecodeString("600060ff0b")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [32]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	if should[0] != contract.Stack[0][0] || should[31] != contract.Stack[0][31] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
		t.Fail()
	}
}

func Test_LT(t *testing.T) {
	data, err := hex.DecodeString("600a600910")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [1]byte{0x01}
	if should[0] != contract.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
		t.Fail()
	}
}

func Test_GT(t *testing.T) {
	data, err := hex.DecodeString("6009600a11")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [1]byte{0x01}
	if should[0] != contract.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
		t.Fail()
	}
}

func Test_AND(t *testing.T) {
	data, err := hex.DecodeString("600f600f16")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [1]byte{0x0f}
	if should[0] != contract.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
		t.Fail()
	}
}

func Test_AND2(t *testing.T) {
	data, err := hex.DecodeString("60ff600016")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [1]byte{0x00}
	if should[0] != contract.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
		t.Fail()
	}
}

func Test_OR(t *testing.T) {
	data, err := hex.DecodeString("600f60f017")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [1]byte{0xff}
	if should[0] != contract.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
		t.Fail()
	}
}

func Test_XOR(t *testing.T) {
	data, err := hex.DecodeString("600f60f018")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [1]byte{0xff}
	if should[0] != contract.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
		t.Fail()
	}
}

func Test_XOR2(t *testing.T) {
	data, err := hex.DecodeString("60ff60ff18")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [1]byte{0x00}
	if should[0] != contract.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
		t.Fail()
	}
}

func Test_NOT(t *testing.T) {
	data, err := hex.DecodeString("600019")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [32]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	if should[0] != contract.Stack[0][0] || should[31] != contract.Stack[0][31] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
		t.Fail()
	}
}

func Test_BYTE(t *testing.T) {
	data, err := hex.DecodeString("60ff601f1a")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [1]byte{0xff}
	if should[0] != contract.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
		t.Fail()
	}
}

func Test_BYTE2(t *testing.T) {
	data, err := hex.DecodeString("61ff00601e1a")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [1]byte{0xff}
	if should[0] != contract.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
		t.Fail()
	}
}

func Test_SHL(t *testing.T) {
	data, err := hex.DecodeString("600160011b")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [1]byte{0x02}
	if should[0] != contract.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
		t.Fail()
	}
}

func Test_SHL2(t *testing.T) {
	data, err := hex.DecodeString("7fFF0000000000000000000000000000000000000000000000000000000000000060041b")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [32]byte{0xf0}
	if should[0] != contract.Stack[0][0] || should[31] != contract.Stack[0][31] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
		t.Fail()
	}
}

func Test_SHR(t *testing.T) {
	data, err := hex.DecodeString("600260011c")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [1]byte{0x01}
	if should[0] != contract.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
		t.Fail()
	}
}

func Test_SHR2(t *testing.T) {
	data, err := hex.DecodeString("60ff60041c")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [1]byte{0x0f}
	if should[0] != contract.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
		t.Fail()
	}
}

func Test_SAR(t *testing.T) {
	data, err := hex.DecodeString("600260011d")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [1]byte{0x01}
	if should[0] != contract.Stack[0][0] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
		t.Fail()
	}
}

func Test_SAR2(t *testing.T) {
	data, err := hex.DecodeString("7fFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF060041d")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should := [32]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	if should[0] != contract.Stack[0][0] || should[31] != contract.Stack[0][31] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
		t.Fail()
	}
}

func Test_SHA3(t *testing.T) {
	data, err := hex.DecodeString("63ffffffff6000526004600020")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	should, err := hex.DecodeString("29045A592007D0C246EF02C2223570DA9522D0CF0F73282C79A1BC8F0BB2C238")
	if err != nil {
		t.Error(err)
	}
	if should[0] != contract.Stack[0][0] || should[31] != contract.Stack[0][31] {
		fmt.Printf("expected: %x\n", should)
		fmt.Printf("is: %x\n", contract.Stack[0])
		t.Fail()
	}
}

func Test_Address(t *testing.T) {
	data, err := hex.DecodeString("30")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)

	if err != nil {
		t.Error(err)
	}
	if !strings.EqualFold(hex.EncodeToString(contract.Stack[0]), contract.Address[2:]) {
		fmt.Println("expected: ", contract.Address)
		fmt.Println("is: ", hex.EncodeToString(contract.Stack[0]))
		t.Fail()
	}
}

func Test_Balance(t *testing.T) {
	data, err := hex.DecodeString("600031")
	if err != nil {
		t.Error(err)
	}
	evm.GetEVM().AddressBalanceMap = make(map[string][]byte)
	evm.GetEVM().AddressBalanceMap["0x00"] = []byte{0x01}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	if contract.Stack[0][0] != 0x01 {
		fmt.Println("expected: 1")
		fmt.Println("is:  + ", contract.Stack[0][0])
		t.Fail()
	}
}

func Test_Origin(t *testing.T) {
	data, err := hex.DecodeString("32")
	if err != nil {
		t.Error(err)
	}
	evm.GetEVM().AddressBalanceMap = make(map[string][]byte)
	evm.GetEVM().AddressBalanceMap["0x00"] = []byte{0x01}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	addr, _ := hex.DecodeString(debugger.Tx.Origin[2:])
	if contract.Stack[0][0] != addr[0] {
		fmt.Printf("expected: 0x%x\n", addr)
		fmt.Printf("is: 0x%x\n", contract.Stack[0])
		t.Fail()
	}
}

func Test_CALLER(t *testing.T) {
	fmt.Println("Not yet implemented")
	t.Fail()
}

func Test_CALLVALUE(t *testing.T) {
	data, err := hex.DecodeString("34")
	if err != nil {
		t.Error(err)
	}
	evm.GetEVM().AddressBalanceMap = make(map[string][]byte)
	evm.GetEVM().AddressBalanceMap["0x00"] = []byte{0x01}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	if contract.Stack[0][0] != 0x01 {
		fmt.Printf("expected: 0x%x\n", []byte{0x01})
		fmt.Printf("is: 0x%x\n", contract.Stack[0])
		t.Fail()
	}
}

func Test_CALLDATA(t *testing.T) {
	data, err := hex.DecodeString("601f35")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	if contract.Stack[0][0] != 0xff { // hardcoded CALLDATA in debugger.New() function
		fmt.Printf("expected: 0x%x\n", []byte{0xff})
		fmt.Printf("is: 0x%x\n", contract.Stack[0])
		t.Fail()
	}
}

func Test_CALLDATASIZE(t *testing.T) {
	data, err := hex.DecodeString("36")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	if contract.Stack[0][0] != 0x20 { // hardcoded CALLDATA in debugger.New() function
		fmt.Printf("expected: 0x%x\n", []byte{0x20})
		fmt.Printf("is: 0x%x\n", contract.Stack[0])
		t.Fail()
	}
}

func Test_CALLDATACOPY(t *testing.T) {
	data, err := hex.DecodeString("60206000600037")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	if contract.Memory[0] != 0xff || contract.Memory[31] != 0xff { // hardcoded CALLDATA in debugger.New() function
		fmt.Printf("expected: 0x%x\n", []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
		fmt.Printf("is: 0x%x\n", contract.Memory)
		t.Fail()
	}
}

func Test_CALLDATACOPY2(t *testing.T) {
	data, err := hex.DecodeString("6008601f600037")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	if contract.Memory[0] != 0xff ||
		contract.Memory[1] != 0x00 ||
		contract.Memory[7] != 0x00 ||
		contract.Memory[31] != 0x00 { // hardcoded CALLDATA in debugger.New() function
		fmt.Printf("expected: 0x%x\n", []byte{0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
		fmt.Printf("is: 0x%x\n", contract.Memory)
		t.Fail()
	}
}

func Test_CALLDATACOPY3(t *testing.T) {
	data, err := hex.DecodeString("60206000600137")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	if contract.Memory[0] != 0x00 ||
		contract.Memory[1] != 0xff ||
		contract.Memory[7] != 0xff ||
		contract.Memory[31] != 0xff ||
		contract.Memory[63] != 0x00 { // hardcoded CALLDATA in debugger.New() function
		fmt.Printf("expected: 0x%x\n", []byte{0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
		fmt.Printf("is: 0x%x\n", contract.Memory)
		t.Fail()
	}
}

func Test_CODESIZE(t *testing.T) {
	data, err := hex.DecodeString("7c00000000000000000000000000000000000000000000000000000000005038")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	if contract.Stack[0][0] != 0x20 {
		fmt.Printf("expected: 0x%x\n", []byte{0x20})
		fmt.Printf("is: 0x%x\n", contract.Memory)
		t.Fail()
	}
}

func Test_CODECOPY(t *testing.T) {
	data, err := hex.DecodeString("7DFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF5060206000600039")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	if contract.Memory[0] != 0x7d ||
		contract.Memory[1] != 0xff ||
		contract.Memory[30] != 0xff ||
		contract.Memory[31] != 0x50 {
		fmt.Printf("expected: 0x%x\n", []byte{0x7d, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x50})
		fmt.Printf("is: 0x%x\n", contract.Memory)
		t.Fail()
	}
}

func Test_PushAndMstore(t *testing.T) {
	data, err := hex.DecodeString("62424240600252")
	if err != nil {
		t.Error(err)
	}
	debugger := New()
	contract := debugger.DeployContract(data)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	debugger.StepDebugger(contract)
	expected := [64]byte{
		byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x42), byte(0x42), byte(0x40), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00), byte(0x00),
	}
	if contract.Memory[0] != expected[0] ||
		// offset 02 => array 31+2; opcode 62 (PUSH3) => PUSH 3 bytes
		contract.Memory[31] != expected[31] || // value = 0x00
		contract.Memory[32] != expected[32] || // value = 0x42
		contract.Memory[33] != expected[33] || // value = 0x42
		contract.Memory[34] != expected[34] || // value = 0x40
		contract.Memory[35] != expected[35] { // value = 0x00
		fmt.Printf("expected: 0x%x\n", expected)
		fmt.Printf("is: 0x%x\n", contract.Memory)
		t.Fail()
	}
}
