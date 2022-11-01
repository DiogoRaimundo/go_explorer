package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var argsWithoutProg = os.Args[1:]

func main() {
	funcToRun := All_Options

	if len(argsWithoutProg) == 0 {
		funcToRun = getOptionToRunFromInput()
	} else if len(argsWithoutProg) > 0 {
		funcToRun = getOptionToRunFromArgs()
	} else {
		printErrorAndExit(("Value can't be less than 0"))
	}

	runFunc(funcToRun)
}

func getOptionToRunFromInput() int {
	fmt.Println("-----========== GO EXPLORER ==========-----")

	maxOption := len(options)
	maxOptionNChars := len(fmt.Sprint(maxOption)) + 1

	for idx, options := range options {
		optionNChars := len(fmt.Sprint(idx))
		optionSpacesToRender := strings.Repeat(" ", maxOptionNChars-optionNChars)

		fmt.Printf("[%d]%s%s\n", idx, optionSpacesToRender, options.Name)
	}

	fmt.Printf("[%d] %s\n", maxOption, "Run All")

	return readOptionFromInput()
}

func readOptionFromInput() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nSelect an option: ")

	optionAsText, _ := reader.ReadString('\n')

	if len(optionAsText) == 2 {
		return All_Options
	}

	optionAsText = optionAsText[:len(optionAsText)-2]

	return parseIntFromString(optionAsText)
}

func getOptionToRunFromArgs() int {
	return parseIntFromString(argsWithoutProg[0])
}

func parseIntFromString(stringValue string) int {
	funcToRun, err := strconv.Atoi(stringValue)
	if err != nil {
		printErrorAndExit(fmt.Sprintf("Option must be an number. \"%s\" received.", stringValue))
	}

	return funcToRun
}

func runFunc(funcToRun int) {
	if funcToRun < 0 {
		printErrorAndExit(("\"funcToRun\" can't be less than 0"))
	}

	if funcToRun < len(options) {
		function := options[funcToRun]
		announceAndRun(function.Name, function.Run)
		return
	}

	for _, function := range options {
		announceAndRun(function.Name, function.Run)
	}
}

func printErrorAndExit(errorToPrint string) {
	fmt.Printf("[ERROR] %s\n", errorToPrint)
	os.Exit(1)
}

func announceAndRun(name string, run func()) {
	announce := "---===== Running " + name + " =====---"
	fmt.Println()
	fmt.Println(announce)

	run()

	fmt.Println("---" + strings.Repeat("=", len(announce)-6) + "---")
	fmt.Println()
}

func printUnableToCompile() {
	fmt.Println("[ERROR] Unable to compile this function")
}
