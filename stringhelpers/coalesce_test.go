package stringhelpers

import "testing"

func newStrSlice(ss ...string) []string {
	return ss
}

func newStrPtr(s string) *string {
	return &s
}

func TestCoalesces(t *testing.T) {
	type args struct {
		ss []string
	}

	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "No strings should return empty string",
			args: args{ss: make([]string, 0)},
			want: newStrPtr(""),
		},
		{
			name: "All empty should return empty",
			args: args{ss: make([]string, 10)},
			want: newStrPtr(""),
		},
		{
			name: "Should return elem3",
			args: args{ss: []string{"", "", "elem3"}},
			want: newStrPtr("elem3"),
		},
		{
			name: "Should return elem1",
			args: args{ss: []string{"elem1", "", "elem3"}},
			want: newStrPtr("elem1"),
		},
		{
			name: "Should return elem2",
			args: args{ss: []string{"", "elem2", "elem3"}},
			want: newStrPtr("elem2"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Coalesces(tt.args.ss...); *got != *tt.want {
				t.Errorf("Coalesces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoalesceWhitespace(t *testing.T) {
	type args struct {
		ss []string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "No strings should return empty string",
			args: args{ss: make([]string, 0)},
			want: newStrPtr(""),
		},
		{
			name: "All empty should return empty",
			args: args{ss: make([]string, 10)},
			want: newStrPtr(""),
		},
		{
			name: "All whitespace should return empty",
			args: args{ss: []string{" ", "   ", "\n\t"}},
			want: newStrPtr(""),
		},
		{
			name: "Should return elem3",
			args: args{ss: []string{"", "", "elem3"}},
			want: newStrPtr("elem3"),
		},
		{
			name: "Should return elem1",
			args: args{ss: []string{"elem1", "", "elem3"}},
			want: newStrPtr("elem1"),
		},
		{
			name: "Should return elem2",
			args: args{ss: []string{"", "elem2", "elem3"}},
			want: newStrPtr("elem2"),
		},
		{
			name: "Should return elem3",
			args: args{ss: []string{"\n", "\n\n\n   ", "\nelem3"}},
			want: newStrPtr("\nelem3"),
		},
		{
			name: "Should return elem1",
			args: args{ss: []string{"elem1", "\n", "elem3"}},
			want: newStrPtr("elem1"),
		},
		{
			name: "Should return elem2",
			args: args{ss: []string{"", "elem2\n  ", "elem3"}},
			want: newStrPtr("elem2\n  "),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CoalesceWhitespace(tt.args.ss...); *got != *tt.want {
				t.Errorf("CoalesceWhitespace() = %v, want %v", got, tt.want)
			}
		})
	}
}
