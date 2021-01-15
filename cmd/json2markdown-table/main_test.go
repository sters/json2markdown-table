package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func Test_do(t *testing.T) {
	tests := []struct {
		name       string
		r          io.Reader
		wantResult string
		wantErr    bool
	}{
		{
			name: "simple",
			r:    strings.NewReader(`[{"foo": 1, "bar": 2}, {"foo": 3, "bar": 4}]`),
			wantResult: `|bar|foo|
|-|-|
|2|1|
|4|3|
`,
			wantErr: false,
		},
		{
			name: "mixed",
			r:    strings.NewReader(`[{"foo": 1, "bar": 2}, {"baz": 3, "boo": 4}]`),
			wantResult: `|bar|baz|boo|foo|
|-|-|-|-|
|2|||1|
||3|4||
`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer([]byte{})
			if err := do(tt.r, buf); (err != nil) != tt.wantErr {
				t.Errorf("do() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := buf.String()
			if tt.wantResult != got {
				t.Errorf("do() = %v, wantResult = %v", got, tt.wantResult)
			}
		})
	}
}
