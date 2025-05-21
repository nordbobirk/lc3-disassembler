package internal

import (
	"fmt"
	"log"
	"math/big"
	"regexp"
	"strings"
)

// Disassemble disassembles the LC-3 binary instructions in the input file and writes the output to the output file.
func Disassemble() {
	fmt.Println("Disassembling...")
	input := ReadInputFile()

	instructions := strings.Split(input, "\n")

	output := ""

	for i, instruction := range instructions {
		if i == len(instructions)-1 {
			output += disassembleInstruction(instruction) // last instruction, dont trim newline
			continue
		}
		output += disassembleInstruction(instruction[:len(instruction)-1]) + "\n" // trim newline
	}

	WriteOutputFile(output)
}

func disassembleInstruction(instruction string) string {
	validateInstruction(instruction)

	opcode := getOpcode(instruction)

	return opcode + func() string {
		if opcode == "BR" {
			return ""
		}
		return " "
	}() + getOperands(instruction[4:], opcode)
}

func validateInstruction(instruction string) {
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
	case "0011":
		return "ST"
	case "1011":
		return "STI"
	case "0111":
		return "STR"
	case "1111":
		return "TRAP"
	case "1000":
		return "RTI"
	default:
		log.Fatal("Invalid opcode: ", instruction[0:4])
		return ""
	}
}

func getOperands(operands string, opcode string) string {
	switch opcode {
	case "ADD":
		return getAddAndOperands(operands)
	case "AND":
		return getAddAndOperands(operands)
	case "TRAP":
		return getTrapOperands(operands)
	case "ST":
		return getStStiLeaLdiLdOperands(operands)
	case "STI":
		return getStStiLeaLdiLdOperands(operands)
	case "LEA":
		return getStStiLeaLdiLdOperands(operands)
	case "LDI":
		return getStStiLeaLdiLdOperands(operands)
	case "LD":
		return getStStiLeaLdiLdOperands(operands)
	case "BR":
		return getBrOperands(operands)
	case "JMP":
		return getJmpJsrrOperands(operands)
	case "RET":
		return getRetOperands()
	case "JSRR":
		return getJmpJsrrOperands(operands)
	case "JSR":
		return getJsrOperands(operands)
	case "STR":
		return getStrLdrOperands(operands)
	case "LDR":
		return getStrLdrOperands(operands)
	case "RTI":
		return getRtiOperands()
	default:
		log.Fatal("Invalid opcode: ", opcode)
		return ""
	}
}

func getAddAndOperands(operands string) string {
	if operands[6:7] == "0" {
		// xxxx DR SR 000 SR2
		return getRegister(operands[0:3]) + ", " + getRegister(operands[3:6]) + ", " + getRegister(operands[9:])
	} else {
		// xxxx DR SR 1 imm5
		return getRegister(operands[0:3]) + ", " + getRegister(operands[3:6]) + ", " + getImmediate(operands[7:])
	}
}

func getTrapOperands(operands string) string {
	// 1111 0000 trapvect8
	return getHex(operands[4:])
}

func getStStiLeaLdiLdOperands(operands string) string {
	// xxxx SR PCoffset9
	return getRegister(operands[:3]) + ", " + getImmediate(operands[3:])
}

func getBrOperands(operands string) string {
	// xxxx nzp PCoffset9
	nzp := ""
	if operands[0:1] == "1" {
		nzp += "n"
	}
	if operands[1:2] == "1" {
		nzp += "z"
	}
	if operands[2:3] == "1" {
		nzp += "p"
	}
	return nzp + " " + getImmediate(operands[3:])
}

func getJmpJsrrOperands(operands string) string {
	// xxxx 000 baseR 000000
	return getRegister(operands[2:5])
}

func getRetOperands() string {
	return "" // no operands for opcode RET
}

func getJsrOperands(operands string) string {
	// xxxx 1 Pcoffset11
	return getImmediate(operands[1:])
}

func getStrLdrOperands(operands string) string {
	// xxxx SR/DR baseR offset6
	return getRegister(operands[:3]) + ", " + getRegister(operands[3:6]) + ", " + getImmediate(operands[6:])
}

func getRtiOperands() string {
	return "" // no operands for opcode RTI
}

func getRegister(register string) string {
	switch register {
	case "000":
		return "r0"
	case "001":
		return "r1"
	case "010":
		return "r2"
	case "011":
		return "r3"
	case "100":
		return "r4"
	case "101":
		return "r5"
	case "110":
		return "r6"
	case "111":
		return "r7"
	default:
		log.Fatal("Invalid register: ", register)
		return ""
	}
}

func getImmediate(immediate string) string {
	n := len(immediate)
	if n == 0 {
		return "0"
	}

	value := new(big.Int)
	value.SetString(immediate, 2)

	if immediate[0] == '1' {
		twoPowerN := new(big.Int).Lsh(big.NewInt(1), uint(n))
		value.Sub(value, twoPowerN)
	}

	return "#" + value.String()
}

func getHex(binStr string) string {
	n := len(binStr)
	if n == 0 {
		return "0x0"
	}

	value := new(big.Int)
	value.SetString(binStr, 2)

	if binStr[0] == '1' {
		twoPowerN := new(big.Int).Lsh(big.NewInt(1), uint(n))
		value.Sub(value, twoPowerN)
	}

	hexStr := fmt.Sprintf("0x%x", value)

	return hexStr
}
