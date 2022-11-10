package opcode

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/r4gn4r-l0/evm-knife/pkg/opcode"
	"github.com/spf13/cobra"
)

var filename string
var OpcodeCmd = &cobra.Command{
	Use: "opcode",
	Run: func(cmd *cobra.Command, args []string) {
		fi, err := os.Stdin.Stat()
		if err != nil {
			panic(err)
		}
		var hexbytecode []byte
		var reader *bufio.Reader
		if len(filename) > 0 {
			f, err := os.Open(filename)
			if err != nil {
				panic(err)
			}
			reader = bufio.NewReader(f)
		} else if fi.Mode()&os.ModeNamedPipe != 0 {
			reader = bufio.NewReader(os.Stdin)
		} else {
			cmd.Help()
			panic("no bytecode given")
		}
		input, _, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}
		hexbytecode, err = hex.DecodeString(strings.Replace(string(input), "0x", "", 1))
		if err != nil {
			panic(err)
		}
		var opcode = opcode.Opcode{
			HexBytecode: hexbytecode,
		}
		err = opcode.Generate()
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		fmt.Println(opcode.ToString())
	},
}

func init() {
	OpcodeCmd.Flags().StringVar(&filename, "file", "", "Full path to the hex binary file")
}
