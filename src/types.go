package main

type InternalFunction func(string, []string) CommandResult

type CommandDefinition struct {
	Command        string
	NumberOfParams int
	Function       InternalFunction
}

type CommandResult struct {
	StopProcessing bool
	Result         string
}
