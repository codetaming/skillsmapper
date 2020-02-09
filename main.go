package main

import (
	"fmt"
	pipeline "github.com/codetaming/skillsmapper/internal/pipeline"
	"io/ioutil"
	"os"
)

func main() {
	raw, err := ioutil.ReadFile("./examples/checkDetails.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	pipeline.Check(raw, "")
}
