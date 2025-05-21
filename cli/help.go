package cli

import (
	"fmt"
)

func Help() {
	fmt.Println("lc3-decompiler")
	fmt.Println("Decompile LC-3 binary files to LC-3 assembly.")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("	decompile - Decompile the LC-3 binary instructions in the input file and write the output to the output file.")
	fmt.Println("	help      - Display this help message.")
}
