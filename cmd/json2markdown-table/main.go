package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/morikuni/failure"
)

func main() {
	if err := do(os.Stdin, os.Stdout); err != nil {
		panic(err)
	}
}

type dataInner map[string]interface{}
type data []dataInner

func do(r io.Reader, w io.Writer) error {
	var read data

	decoder := json.NewDecoder(r)
	decoder.UseNumber()
	if err := decoder.Decode(&read); err != nil {
		return failure.Wrap(err)
	}

	keys := make([]string, 0, len(read[0]))
	added := map[string]struct{}{}
	for _, inner := range read {
		for k := range inner {
			if _, ok := added[k]; ok {
				continue
			}
			keys = append(keys, k)
			added[k] = struct{}{}
		}
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	for _, k := range keys {
		fmt.Fprintf(w, `|%s`, k)
	}
	fmt.Fprintf(w, "|\n")

	for range keys {
		fmt.Fprintf(w, `|-`)
	}
	fmt.Fprintf(w, "|\n")

	for _, inner := range read {
		for _, k := range keys {
			v, ok := inner[k]
			if !ok {
				v = ""
			}
			fmt.Fprintf(w, `|%s`, v)
		}
		fmt.Fprintf(w, "|\n")
	}

	return nil
}
