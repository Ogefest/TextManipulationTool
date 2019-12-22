package main

import (
	"bufio"
	"io"
	"os"
	"sync"
)

func FileProcess(filein *os.File, fileout *os.File) {

	result := make(chan string, 10)
	jobs := make(chan string, 10)

	wg := new(sync.WaitGroup)

	for w := 0; w < 10; w++ {
		wg.Add(1)
		go lineProceed(jobs, result, wg)
	}

	go func() {
		r := bufio.NewReader(filein)
		for {
			// line, err := r.ReadString(10) // 0x0A separator = newline
			line, _, err := r.ReadLine()
			if err == io.EOF {
				break
			} else if err != nil {
				// return err // if you return error
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

	w := bufio.NewWriter(fileout)
	for str := range result {
		w.WriteString(str + "\n")
		w.Flush()
	}

}

func lineProceed(jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {

	defer wg.Done()

	var allCommands map[string]CommandDefinition = GetDefinitions()

	for lineToProceed := range jobs {

		sendResult := true

		for i := 0; i < len(os.Args); i++ {
			part := os.Args[i]

			if cmd, contains := allCommands[part]; contains {

				paramsToCall := os.Args[i+1 : i+cmd.NumberOfParams+1]

				i += cmd.NumberOfParams

				callResult := cmd.Function(lineToProceed, paramsToCall)
				if callResult.StopProcessing {
					sendResult = false
					break
				}
				lineToProceed = callResult.Result

			}
		}
		if sendResult {
			results <- lineToProceed
		}

	}

}
