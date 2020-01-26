package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"
)

func FileProcess() {

	options := GetRuntimeOptions()

	result := make(chan string, 10)
	jobs := make(chan string, 10)

	wg := new(sync.WaitGroup)

	for w := 0; w < options.AsyncThreads; w++ {
		wg.Add(1)
		go lineProceed(jobs, result, wg)
	}

	go func() {
		r := bufio.NewReader(options.FileIn)
		for {
			// line, err := r.ReadString(10) // 0x0A separator = newline
			line, _, err := r.ReadLine()
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("ERR: %s", err)
			} else {
				jobs <- string(line)
			}
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(result)
	}()

	w := bufio.NewWriter(options.FileOut)
	for str := range result {
		w.WriteString(str + "\n")
		w.Flush()
	}

}

func lineProceed(jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {

	defer wg.Done()

	var allCommands map[string]CommandDefinition = GetDefinitions()

	for lineToProceed := range jobs {

		proceedParam := ParamDefinition{
			line:                  lineToProceed,
			temporaryLine:         "",
			columnSeparator:       " ",
			stopProcessing:        false,
			isColumnProcessing:    false,
			skipCurrentLine:       false,
			columnProcessingIndex: 0,
		}

		sendResult := true

		for i := 0; i < len(os.Args); i++ {
			part := os.Args[i]

			if cmd, contains := allCommands[part]; contains {

				paramsToCall := os.Args[i+1 : i+cmd.NumberOfParams+1]

				i += cmd.NumberOfParams
				cmd.Function(&proceedParam, paramsToCall)
				if proceedParam.stopProcessing {
					sendResult = false
					break
				}
				if proceedParam.skipCurrentLine {
					break
				}
				if proceedParam.callError != nil {
					fmt.Printf("ERR: %s", proceedParam.callError)
					proceedParam.callError = nil
				}
			}
		}

		if proceedParam.isColumnProcessing {
			ColumnDeselect(&proceedParam, []string{})
		}

		if sendResult {
			results <- proceedParam.line
		}

	}

}
