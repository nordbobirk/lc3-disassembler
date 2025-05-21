package cli

import (
	"fmt"
	"lc3-disassembler/internal"
	"strings"
)

// Run is the entry point for the CLI.
func Run(args []string) {
	if len(args) <= 1 {
		Help()
		return
	}

	switch strings.ToLower(args[1]) {
	case "disassemble":
		internal.Disassemble()
	case "help":
		Help()
	default:
		fmt.Println("Unknown command:", args[1])
		Help()
	}
}
