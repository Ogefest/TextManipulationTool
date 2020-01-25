package main

import (
	"flag"
	"io/ioutil"
	"os"
)

func GetRuntimeOptions() ApplicationOptions {
	// options := flag.String("async", "in", "out")

	async := flag.Bool("async", false, "a bool")
	file_in := flag.String("in", "", "a string")
	file_out := flag.String("out", "", "a string")
	json_file := flag.String("json", "", "a string")

	flag.Parse()

	result := ApplicationOptions{
		AsyncThreads: 1,
		FileIn:       os.Stdin,
		FileOut:      os.Stdout,
		Json:         "",
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
	if *json_file != "" {
		json_content, err := ioutil.ReadFile(*json_file)
		if err != nil {
			panic(err)
		}
		result.Json = string(json_content)
	}

	return result

}
