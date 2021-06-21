package main

import (
	"bufio"
	"flag"
	"log"
	"os"
)

func main() {
	var outputFlag string

	flag.StringVar(&outputFlag, "output", "", "the file to output to")
	flag.Parse()

	if outputFlag == "" {
		flag.Usage()
		return
	}

	f, err := os.Create(outputFlag)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := Render{
		w: bufio.NewWriter(f),
	}

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
