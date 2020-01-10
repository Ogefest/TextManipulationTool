package main

import (
	"flag"
	"os"
)

func GetRuntimeOptions() ApplicationOptions {
	// options := flag.String("async", "in", "out")

	async := flag.Bool("async", false, "a bool")
	file_in := flag.String("in", "", "a string")
	file_out := flag.String("out", "", "a string")

	flag.Parse()

	result := ApplicationOptions{
		AsyncThreads: 1,
		FileIn:       os.Stdin,
		FileOut:      os.Stdout,
	}

	if *async == true {
		result.AsyncThreads = 10
	}
	if *file_in != "" {
		f, err := os.Open(*file_in)
		if err != nil {
			panic(err)
		}
		result.FileIn = f
	}
	if *file_out != "" {
		f, err := os.Create(*file_out)
		if err != nil {
			panic(err)
		}
		result.FileOut = f
	}

	return result

}
