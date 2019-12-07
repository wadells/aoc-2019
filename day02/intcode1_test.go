package main

import (
	"reflect"
	"testing"
)

func TestAddOpcode(t *testing.T) {
	program := []int{1, 0, 0, 0, 99}
	execute(OPCODES, program)
	expected := []int{2, 0, 0, 0, 99}
	if !reflect.DeepEqual(program, expected) {
		t.Errorf("Got %v\tExpected %v", program, expected)
	}
}

func TestMultiplyOpcode(t *testing.T) {
	program := []int{2, 3, 0, 3, 99}
	execute(OPCODES, program)
	expected := []int{2, 3, 0, 6, 99}
	if !reflect.DeepEqual(program, expected) {
		t.Errorf("Got %v\tExpected %v", program, expected)
	}
}

func TestSaveAfter99(t *testing.T) {
	program := []int{2, 4, 4, 5, 99, 0}
	execute(OPCODES, program)
	expected := []int{2, 4, 4, 5, 99, 9801}
	if !reflect.DeepEqual(program, expected) {
		t.Errorf("Got %v\tExpected %v", program, expected)
	}
}

func TestChangeFutureOpcode(t *testing.T) {
	program := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	execute(OPCODES, program)
	expected := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
	if !reflect.DeepEqual(program, expected) {
		t.Errorf("Got %v\tExpected %v", program, expected)
	}
}
