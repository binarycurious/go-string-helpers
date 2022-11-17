package stringhelpers

import "testing"

type args struct {
	s *string
}

func makeMyStr(s string) args {
	var a args
	a.s = &s
	return a
}

func makeNilArgs() args {
	var a args
	a.s = nil
	return a
}

func TestIsEmptyOrWhitespace(t *testing.T) {

	tests := []struct {
		name string
		a    args
		want bool
	}{
		{
			name: "Nil string is empty",
			a:    makeNilArgs(),
			want: true,
		},
		{
			name: "Empty string is empty",
			a:    makeMyStr(""),
			want: true,
		},
		{
			name: "Whitespace string is Notempty",
			a:    makeMyStr("  		"),
			want: true,
		},
		{
			name: "More Whitespace string is empty",
			a:    makeMyStr("  		\n\t"),
			want: true,
		},
		{
			name: "Whitespace and char string is Not empty",
			a:    makeMyStr("  		\nc"),
			want: false,
		},
		{
			name: "Single char Not empty",
			a:    makeMyStr("c"),
			want: false,
		},
		{
			name: "Sentence Not empty",
			a:    makeMyStr("This is a normal sentence, with punctuation"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmptyOrWhitespace(tt.a.s); got != tt.want {
				t.Errorf("IsEmptyOrWhitespace() = %v, want %v", got, tt.want)
			}
		})
	}
}
