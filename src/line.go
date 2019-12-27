package main

import (
	"strconv"
	"strings"
)

func Replace(input *ParamDefinition, params []string) {
	var from = params[0]
	var to = params[1]

	input.line = strings.Replace(input.line, from, to, -1)

}

func Filter(input *ParamDefinition, params []string) {
	c := strings.Contains(input.line, params[0])

	input.stopProcessing = !c
}

func FIf(input *ParamDefinition, params []string) {
	c := strings.Contains(input.line, params[0])

	if !c {
		input.skipCurrentLine = true
	}

}

func FIfNot(input *ParamDefinition, params []string) {
	c := strings.Contains(input.line, params[0])

	if c {
		input.skipCurrentLine = true
	}

}

func Lower(input *ParamDefinition, params []string) {

	input.line = strings.ToLower(input.line)

}

func Upper(input *ParamDefinition, params []string) {

	input.line = strings.ToUpper(input.line)

}

func Title(input *ParamDefinition, params []string) {

	input.line = strings.ToTitle(input.line)

}

func Trim(input *ParamDefinition, params []string) {

	input.line = strings.TrimSpace(input.line)

}

func Duplicate(input *ParamDefinition, params []string) {

	input.line = input.line + "\n" + input.line

}

func Prefix(input *ParamDefinition, params []string) {

	input.line = params[0] + input.line

}

func Suffix(input *ParamDefinition, params []string) {

	input.line = input.line + params[0]

}

func Cut(input *ParamDefinition, params []string) {

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

	input.line = string(runes[from:to])

}
