package opcode

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/r4gn4r-l0/evm-knife/pkg/opcode"
	"github.com/spf13/cobra"
)

var OpcodeCmd = &cobra.Command{
	Use: "opcode",
	Run: func(cmd *cobra.Command, args []string) {
		fi, err := os.Stdin.Stat()
		if err != nil {
			panic(err)
		}
		var hexbytecode []byte
		if fi.Mode()&os.ModeNamedPipe == 0 {
			fmt.Println("no pipe :(")
		} else {
			reader := bufio.NewReader(os.Stdin)
			input, _, err := reader.ReadLine()
			if err != nil && err == io.EOF {
				fmt.Println("ERROR... TODO")
			}
			hexbytecode, _ = hex.DecodeString(strings.Replace(string(input), "0x", "", 1))
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
