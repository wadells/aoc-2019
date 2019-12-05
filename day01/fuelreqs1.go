package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func moduleFuelCost(module int) int {
	return module/3 - 2
}

func shipFuelCost(modules []int) int {
	sum := 0
	for _, module := range modules {
		sum += moduleFuelCost(module)
	}
	return sum
}

func readModules(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var modules []int
	for scanner.Scan() {
		line := scanner.Text()
		module, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		modules = append(modules, module)
	}
	return modules
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		panic("please provide the input filename as the first and only argument")
	}
	filename := args[0]
	modules := readModules(filename)
	cost := shipFuelCost(modules)
	fmt.Printf("The total fuel cost is %d.\n", cost)
}
