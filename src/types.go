package main

import "os"

type InternalFunction func(*ParamDefinition, []string)

type ParamDefinition struct {
	line                  string
	temporaryLine         string
	columnSeparator       string
	stopProcessing        bool
	skipCurrentLine       bool
	isColumnProcessing    bool
	columnProcessingIndex int
	callError             error
}

type CommandDefinition struct {
	Command        string
	NumberOfParams int
	Function       InternalFunction
}

type ApplicationOptions struct {
	AsyncThreads int
	FileIn       *os.File
	FileOut      *os.File
	Json         string
}
