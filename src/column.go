package main

import (
	"strconv"
	"strings"
)

func Column(input *ParamDefinition, params []string) CommandResult {
	return CommandResult{
		StopProcessing: false,
		Result:         params[0] + input.line,
	}
}

func ColumnSeparator(input *ParamDefinition, params []string) CommandResult {
	input.columnSeparator = params[0]
	return CommandResult{
		StopProcessing: false,
		Result:         input.line,
	}
}

func ColumnRemove(input *ParamDefinition, params []string) CommandResult {

	cols := strings.Split(input.line, input.columnSeparator)

	columnToRemove, _ := strconv.Atoi(params[0])
	tmp := append(cols[:columnToRemove-1], cols[columnToRemove:]...)

	return CommandResult{
		StopProcessing: false,
		Result:         strings.Join(tmp, input.columnSeparator),
	}
}

func ColumnAdd(input *ParamDefinition, params []string) CommandResult {

	cols := strings.Split(input.line, input.columnSeparator)

	columnToAdd, _ := strconv.Atoi(params[0])
	columnToAddValue := params[1]

	p1 := cols[:columnToAdd]
	p2 := cols[columnToAdd:]
	copy(p2[1:], p2)
	p2[0] = columnToAddValue

	result := append(p1, p2...)

	return CommandResult{
		StopProcessing: false,
		Result:         strings.Join(result, input.columnSeparator),
	}
}

// func ColumnSelect(input *ParamDefinition, params []string) CommandResult {

// 	cols := strings.Split(input.line, input.columnSeparator)
// 	colNumber, _ := strconv.Atoi(params[0])
// 	input.line = input.processingString
// 	input.processingString = cols[colNumber]

// 	return CommandResult{
// 		StopProcessing: false,
// 		Result:         input.processingString,
// 	}
// }

// func ColumnDeselect(input *ParamDefinition, params []string) CommandResult {

// 	cols := strings.Split(input.line, input.columnSeparator)
// 	colNumber, _ := strconv.Atoi(params[0])
// 	input.line = input.processingString
// 	input.processingString = cols[colNumber]

// 	return CommandResult{
// 		StopProcessing: false,
// 		Result:         input.processingString,
// 	}
// }
