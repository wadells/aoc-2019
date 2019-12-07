package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var EOP = errors.New("End of Program")

var OPCODES = getOpcodes()

type opcode func(int, []int) error

func opcode1(idx int, program []int) error {
	/*
		Opcode 1 adds together numbers read from two positions and stores the
		result in a third position. The three integers immediately after
		the opcode tell you these three positions - the first two indicate
		the positions from which you should read the input values, and the
		third indicates the position at which the output should be stored.
	*/
	arg_ptr1 := program[idx+1]
	arg_ptr2 := program[idx+2]
	sum_ptr := program[idx+3]
	sum := program[arg_ptr1] + program[arg_ptr2]
	program[sum_ptr] = sum
	return nil
}

func opcode2(idx int, program []int) error {
	/*
		Opcode 2 works exactly like opcode 1, except it multiplies the two inputs
		instead of adding them. Again, the three integers after the opcode
		indicate where the inputs and outputs are, not their values.
	*/
	arg_ptr1 := program[idx+1]
	arg_ptr2 := program[idx+2]
	sum_ptr := program[idx+3]
	product := program[arg_ptr1] * program[arg_ptr2]
	program[sum_ptr] = product
	return nil
}

func opcode99(idx int, program []int) error {
	return EOP
}

func opcodeUnknown(idx int, program []int) error {
	codepoint := program[idx]
	return errors.New(fmt.Sprintf("Invalid opcode %d at index %d", codepoint, idx))
}

func execute(instructions []opcode, program []int) {
	n := len(program)
	for i := 0; i < n; i += 4 {
		instruction := instructions[program[i]]
		err := instruction(i, program)
		if err != nil {
			if err == EOP {
				break
			} else {
				panic(err)
			}
		}
	}
}

func readIntcode(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	// we know there is only a single line, so no loop necessary
	line, err := reader.Read()
	if err != nil {
		panic(err)
	}
	program := make([]int, len(line))
	for i, str := range line {
		val, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		program[i] = val
	}

	return program
}

func getOpcodes() []opcode {
	instructions := make([]opcode, 100)
	for i := range instructions {
		instructions[i] = opcodeUnknown
	}
	instructions[1] = opcode1
	instructions[2] = opcode2
	instructions[99] = opcode99
	return instructions
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		panic("please provide the input filename as the first and only argument")
	}
	filename := args[0]
	program := readIntcode(filename)
	// per the instructions: replace position 1 with the value 12 and replace
	// position 2 with the value 2.
	program[1] = 12
	program[2] = 2

	execute(OPCODES, program)

	fmt.Printf("Value at position 0 after program halts: %d\n", program[0])
}
