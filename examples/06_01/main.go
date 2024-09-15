package main

import (
	"flag"
	"fmt"
)

func main() {
	fs := flag.NewFlagSet("app", flag.ContinueOnError)

	_ = fs.Bool("isset", true, "")
	_ = fs.Bool("isnotset", true, "")

	if err := fs.Parse([]string{"-isset"}); err != nil {
		panic(err)
	}

	for f := range VisitAllFlags(fs) {
		fmt.Println(f.Name)
		break
	}
}

func VisitAllFlags(flagset *flag.FlagSet) func(yield func(*flag.Flag) bool) {
	return func(yield func(*flag.Flag) bool) {
		flagset.VisitAll(
			func(f *flag.Flag) {
				// If yield returned false and we call it again, a panic occurs:
				// panic: runtime error: range function continued iteration after function for loop body returned false
				yield(f)
			},
		)
	}
}
