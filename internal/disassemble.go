package internal

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

// Disassemble disassembles the LC-3 binary instructions in the input file and writes the output to the output file.
func Disassemble() {
	fmt.Println("Disassembling...")
	input := ReadInputFile()

	instructions := strings.Split(input, "\n")

	for i, instruction := range instructions {
		if i == len(instructions)-1 {
			decompileInstruction(instruction) // last instruction, dont trim newline
			continue
		}
		decompileInstruction(instruction[:len(instruction)-1]) // trim newline
	}

	WriteOutputFile(input)
}

func decompileInstruction(instruction string) {
	if len(instruction) != InstructionLength {
		log.Fatal("Invalid instruction length: ", len(instruction), " for instruction: ", instruction)
	}

	validInstruction, err := regexp.MatchString("^[0-1]{16}$", instruction)

	if err != nil {
		log.Fatal(err)
	}

	if !validInstruction {
		log.Fatal("Invalid instruction: ", instruction)
	}

	fmt.Println(instruction, getOpcode(instruction))
}

func getOpcode(instruction string) string {
	switch instruction[0:4] {
	case "0001":
		return "ADD"
	case "0101":
		return "AND"
	case "0000":
		return "BR"
	case "1100":
		if instruction[7:10] == "111" {
			return "RET"
		}
		return "JMP"
	case "0100":
		if instruction[4:5] == "0" {
			return "JSRR"
		}
		return "JSR"
	case "0010":
		return "LD"
	case "1010":
		return "LDI"
	case "0110":
		return "LDR"
	case "1110":
		return "LEA"
	case "1001":
		return "NOT"
	default:
		log.Fatal("Invalid opcode: ", instruction[0:4])
		return ""
	}
}
