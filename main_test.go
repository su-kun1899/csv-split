package main

import (
	"flag"
	"os"
	"testing"
)

func Test_run_flagVar(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want int
	}{
		{
			name: "No args",
			args: []string{},
			want: ExitCodeError,
		},
		{
			name: "record_count 50",
			args: []string{"-record_count", "50", "input.csv"},
			want: ExitCodeOK,
		},
		{
			name: "multiple args",
			args: []string{"input1.csv", "input2.csv"},
			want: ExitCodeError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			commandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			if got := run(tt.args); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
