package main

import (
	"strconv"
	"strings"
)

func Replace(input *ParamDefinition, params []string) CommandResult {
	var from = params[0]
	var to = params[1]

	return CommandResult{
		StopProcessing: false,
		Result:         strings.Replace(input.line, from, to, -1),
	}
}

func Filter(input *ParamDefinition, params []string) CommandResult {
	c := strings.Contains(input.line, params[0])

	return CommandResult{
		StopProcessing: !c,
		Result:         input.line,
	}
}

func Lower(input *ParamDefinition, params []string) CommandResult {
	return CommandResult{
		StopProcessing: false,
		Result:         strings.ToLower(input.line),
	}
}

func Upper(input *ParamDefinition, params []string) CommandResult {
	return CommandResult{
		StopProcessing: false,
		Result:         strings.ToUpper(input.line),
	}
}

func Title(input *ParamDefinition, params []string) CommandResult {
	return CommandResult{
		StopProcessing: false,
		Result:         strings.ToTitle(input.line),
	}
}

func Trim(input *ParamDefinition, params []string) CommandResult {
	return CommandResult{
		StopProcessing: false,
		Result:         strings.TrimSpace(input.line),
	}
}

func Duplicate(input *ParamDefinition, params []string) CommandResult {
	return CommandResult{
		StopProcessing: false,
		Result:         input.line + "\n" + input.line,
	}
}

func Prefix(input *ParamDefinition, params []string) CommandResult {
	return CommandResult{
		StopProcessing: false,
		Result:         params[0] + input.line,
	}
}

func Suffix(input *ParamDefinition, params []string) CommandResult {
	return CommandResult{
		StopProcessing: false,
		Result:         params[0] + input.line,
	}
}

func Cut(input *ParamDefinition, params []string) CommandResult {

	lines := strings.Split(params[0], ":")

	var from int
	var to int

	if lines[1] == "" {
		to = len(input.line)
	} else {
		to, _ = strconv.Atoi(lines[1])
	}

	if lines[0] == "" {
		from = 0
	} else {
		from, _ = strconv.Atoi(lines[0])
	}
	runes := []rune(input.line)

	return CommandResult{
		StopProcessing: false,
		Result:         string(runes[from:to]),
	}
}
