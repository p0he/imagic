package main

import (
	"flag"
	"fmt"
	"os"

	"imagic/src/cmd"
	"imagic/src/converter"
	"imagic/src/file"
	"imagic/src/validator"
)

func main() {
	fileOperator := file.NewFileOperator()
	convOperator := converter.NewConverterOperator()
	validateOperator := validator.NewValidatorOperator()
	cmdOP := cmd.NewCMDOperator(fileOperator, convOperator, validateOperator)

	flag.Parse()
	args := flag.Args()

	if err := cmdOP.Run(args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}
