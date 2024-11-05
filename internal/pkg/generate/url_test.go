// Package generate contains help funcs to generate some stuff
package generate

import "testing"

func TestString(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "generate empty string",
			args: args{length: 0},
		},
		{
			name: "generate short string",
			args: args{length: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.length); len(got) != tt.args.length {
				t.Errorf("String() = %v, want %v", got, tt.args.length)
			}
		})
	}
}
