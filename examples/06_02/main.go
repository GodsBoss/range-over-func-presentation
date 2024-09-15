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
		callYield := true
		flagset.VisitAll(
			func(f *flag.Flag) {
				if callYield {
					callYield = yield(f)
				}
			},
		)
	}
}
