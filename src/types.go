package main

type InternalFunction func(*ParamDefinition, []string) CommandResult

type ParamDefinition struct {
	line             string
	processingString string
	columnSeparator  string
}

type CommandDefinition struct {
	Command        string
	NumberOfParams int
	Function       InternalFunction
}

type CommandResult struct {
	StopProcessing bool
	Result         string
}
