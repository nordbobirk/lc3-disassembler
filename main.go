package main

import (
	"lc3-disassembler/cli"
	"lc3-disassembler/internal"
	"log"
	"os"
)

func main() {
	initDataDir()
	initInputFile()
	cli.Run(os.Args)
}

func initDataDir() {
	if !dataDirExists() {
		err := os.Mkdir(internal.DataDir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func dataDirExists() bool {
	_, err := os.Stat(internal.DataDir)
	return err == nil || !os.IsNotExist(err)
}

func initInputFile() {
	if !inputFileExists() {
		err := os.WriteFile(internal.InputFileName, []byte(""), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func inputFileExists() bool {
	_, err := os.Stat(internal.InputFileName)
	return err == nil || !os.IsNotExist(err)
}
