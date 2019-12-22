package main

func Column(line string, params []string) CommandResult {
	return CommandResult{
		StopProcessing: false,
		Result:         params[0] + line,
	}
}
