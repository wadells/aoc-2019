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

func execute(instructions []opcode, memory []int) {
	n := len(memory)
	for i := 0; i < n; i += 4 {
		instruction := instructions[memory[i]]
		err := instruction(i, memory)
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
	// retain the original, do not corrupt
	program := readIntcode(filename)
	// a temporary buffer for running intcode programs
	memory := make([]int, len(program))

	verb := 0
	noun := 0
	found := false
	for noun = 0; !found && noun < 100; noun++ {
		for verb = 0; !found && verb < 100; verb++ {
			copy(memory, program)
			memory[1] = noun
			memory[2] = verb
			execute(OPCODES, memory)
			output := memory[0]
			if output == 19690720 {
				found = true
			}
		}
	}

	if !found {
		panic("no solution found")
	}

	// decrement noun & verb by 1, as they'll both increment by one after
	// the solution is found -- just before the loop condition fails
	noun--
	verb--

	fmt.Printf("100 * noun + verb =  %d\n", 100*noun+verb)
}
