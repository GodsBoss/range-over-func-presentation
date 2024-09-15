package main

import (
	"fmt"
	"io"
	"iter"
	"net/http"
	"slices"
	"time"

	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/BooleanCat/go-functional/v2/it/filter"
	"github.com/GodsBoss/g/seq/iterate"
	"github.com/GodsBoss/g/seq/throttle"
)

func main() {
	rawURLs := []string{
		"https://foo.example.com/",
		"https://example.com/",
		"https://example.org/",
		"https://baz.example.org",
	}

	client := &http.Client{
		Timeout: time.Second,
	}

	throttling, err := throttle.IterationsPerInterval{
		Maximum:   1,
		Timeframe: time.Second,
	}.Strategy()
	if err != nil {
		panic(err)
	}

	iterateRawURLs := slices.Values(rawURLs)
	logRawURLs := logValues(iterateRawURLs)
	errs := iterate.Map(checkAvailability(client))(logRawURLs)
	throttled := throttle.Iteration[error](throttling)(errs)
	onlyErrs := it.Filter(throttled, filter.Not(isNilError))

	for err := range onlyErrs {
		fmt.Println(err)
	}
}

func checkAvailability(client *http.Client) func(url string) error {
	return func(url string) error {
		response, err := client.Get(url)
		if err != nil {
			return fmt.Errorf("could not reach %s: %w", url, err)
		}
		io.Copy(io.Discard, response.Body)
		response.Body.Close()
		return nil
	}
}

// isNilError is a predicate to check whether an error is nil.
func isNilError(err error) bool {
	return err == nil
}

// logValues converts a one-valued iterator to a one-valued iterator of the same type that logs its values.
func logValues[T any](iterator iter.Seq[T]) iter.Seq[T] {
	return func(yield func(value T) bool) {
		iterator(
			func(value T) bool {
				fmt.Printf("%s: %v\n", time.Now(), value)
				return yield(value)
			},
		)
	}
}
