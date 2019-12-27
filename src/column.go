package main

import (
	"errors"
	"strconv"
	"strings"
)

func Column(input *ParamDefinition, params []string) {
	cols := strings.Split(input.line, input.columnSeparator)
	colNumber, _ := strconv.Atoi(params[0])
	colNumber--

	input.line = cols[colNumber]
}

func ColumnSeparator(input *ParamDefinition, params []string) {
	input.columnSeparator = params[0]

}

func ColumnRemove(input *ParamDefinition, params []string) {

	cols := strings.Split(input.line, input.columnSeparator)

	columnToRemove, _ := strconv.Atoi(params[0])
	tmp := append(cols[:columnToRemove-1], cols[columnToRemove:]...)

	input.line = strings.Join(tmp, input.columnSeparator)

}

func ColumnAdd(input *ParamDefinition, params []string) {

	cols := strings.Split(input.line, input.columnSeparator)

	columnToAdd, _ := strconv.Atoi(params[0])
	columnToAddValue := params[1]

	p1 := cols[:columnToAdd]
	p2 := cols[columnToAdd:]
	copy(p2[1:], p2)
	p2[0] = columnToAddValue

	result := append(p1, p2...)

	input.line = strings.Join(result, input.columnSeparator)

}

func ColumnSelect(input *ParamDefinition, params []string) {

	if input.isColumnProcessing {
		ColumnDeselect(input, params)
	}

	cols := strings.Split(input.line, input.columnSeparator)
	colNumber, _ := strconv.Atoi(params[0])
	colNumber--

	input.temporaryLine = input.line
	input.line = cols[colNumber]
	input.isColumnProcessing = true
	input.columnProcessingIndex = colNumber

}

func ColumnDeselect(input *ParamDefinition, params []string) {

	cols := strings.Split(input.temporaryLine, input.columnSeparator)
	cols[input.columnProcessingIndex] = input.line

	input.line = strings.Join(cols, input.columnSeparator)
	input.isColumnProcessing = false
	input.columnProcessingIndex = 0

}

func ColumnOrder(input *ParamDefinition, params []string) {

	cols := strings.Split(input.line, input.columnSeparator)
	orderCols := strings.Split(params[0], ",")

	result := make([]string, len(cols))
	if len(orderCols) != len(result) {
		input.callError = errors.New("Number of column order must match number of columns in line")
		return
	}

	for index, val := range orderCols {
		intval, _ := strconv.Atoi(val)
		intval--

		result[index] = cols[intval]
	}

	input.line = strings.Join(result, input.columnSeparator)
}
