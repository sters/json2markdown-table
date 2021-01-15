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
			wantResult: `|foo|bar|
|-|-|
|1|2|
|3|4|
`,
			wantErr: false,
		},
		{
			name: "mixed",
			r:    strings.NewReader(`[{"foo": "foo", "bar": "bar"}, {"baz": "baz", "boo": "boo"}]`),
			wantResult: `|foo|bar|baz|boo|
|-|-|-|-|
|foo|bar|||
|||baz|boo|
`,
			wantErr: false,
		},
		{
			name:       "invalid json",
			r:          strings.NewReader(`[`),
			wantResult: "",
			wantErr:    true,
		},
		{
			name:       "invalid not slice map",
			r:          strings.NewReader(`{}`),
			wantResult: "",
			wantErr:    true,
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
