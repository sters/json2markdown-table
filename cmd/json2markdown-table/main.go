package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	mapslice "github.com/mickep76/mapslice-json"
	"github.com/morikuni/failure"
	"github.com/spf13/cast"
)

func main() {
	if err := do(os.Stdin, os.Stdout); err != nil {
		panic(err)
	}
}

func do(r io.Reader, w io.Writer) error {
	data := []mapslice.MapSlice{}
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return failure.Wrap(err)
	}

	var keys []interface{}
	added := map[interface{}]struct{}{}
	for _, inner := range data {
		for _, item := range inner {
			if _, ok := added[item.Key]; ok {
				continue
			}
			keys = append(keys, item.Key)
			added[item.Key] = struct{}{}
		}
	}

	for _, k := range keys {
		fmt.Fprintf(w, `|%s`, k)
	}
	fmt.Fprintf(w, "|\n")

	for range keys {
		fmt.Fprintf(w, `|-`)
	}
	fmt.Fprintf(w, "|\n")

	for _, inner := range data {
		for _, k := range keys {
			var found interface{} = ""
			for _, item := range inner {
				if item.Key == k {
					found = item.Value
					break
				}
			}

			fmt.Fprintf(w, "|%s", cast.ToString(found))
		}
		fmt.Fprintf(w, "|\n")
	}

	return nil
}
