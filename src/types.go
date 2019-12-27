package main

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

// type CommandResult struct {
// 	StopProcessing bool
// 	// Result         string
// }
