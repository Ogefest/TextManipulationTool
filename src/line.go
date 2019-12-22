package main

import (
	"strconv"
	"strings"
)

func Replace(line string, params []string) CommandResult {
	var from = params[0]
	var to = params[1]

	return CommandResult{
		StopProcessing: false,
		Result:         strings.Replace(line, from, to, -1),
	}
}

func Filter(line string, params []string) CommandResult {
	c := strings.Contains(line, params[0])

	return CommandResult{
		StopProcessing: !c,
		Result:         line,
	}
}

func Lower(line string, params []string) CommandResult {
	return CommandResult{
		StopProcessing: false,
		Result:         strings.ToLower(line),
	}
}

func Upper(line string, params []string) CommandResult {
	return CommandResult{
		StopProcessing: false,
		Result:         strings.ToUpper(line),
	}
}

func Title(line string, params []string) CommandResult {
	return CommandResult{
		StopProcessing: false,
		Result:         strings.ToTitle(line),
	}
}

func Trim(line string, params []string) CommandResult {
	return CommandResult{
		StopProcessing: false,
		Result:         strings.TrimSpace(line),
	}
}

func Duplicate(line string, params []string) CommandResult {
	return CommandResult{
		StopProcessing: false,
		Result:         line + "\n" + line,
	}
}

func Prefix(line string, params []string) CommandResult {
	return CommandResult{
		StopProcessing: false,
		Result:         params[0] + line,
	}
}

func Suffix(line string, params []string) CommandResult {
	return CommandResult{
		StopProcessing: false,
		Result:         params[0] + line,
	}
}

func Cut(line string, params []string) CommandResult {

	lines := strings.Split(params[0], ":")

	var from int
	var to int

	if lines[1] == "" {
		to = len(line)
	} else {
		to, _ = strconv.Atoi(lines[1])
	}

	if lines[0] == "" {
		from = 0
	} else {
		from, _ = strconv.Atoi(lines[0])
	}
	runes := []rune(line)

	return CommandResult{
		StopProcessing: false,
		Result:         string(runes[from:to]),
	}
}
