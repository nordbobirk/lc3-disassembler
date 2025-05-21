package cli

import (
	"fmt"
)

func Help() {
	fmt.Println("lc3-disassembler")
	fmt.Println("Disassemble LC-3 binaries.")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("	disassemble - Disassemble the LC-3 binary instructions in the input file and write the output to the output file.")
	fmt.Println("	help      - Display this help message.")
}
