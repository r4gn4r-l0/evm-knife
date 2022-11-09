package get

import (
	"github.com/r4gn4r-l0/evm-knife/cmd/get/opcode"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use: "get",
}

func init() {
	GetCmd.AddCommand(opcode.OpcodeCmd)
}
