package hiddenWord

import "testing"

type args struct {
	numberChain int
}

type testArgs struct {
	name string
	args args
	want string
}

func TestHiddenWord(t *testing.T) {
	tests := []testArgs{
		{name: "Test one number become a letter", args: args{1}, want: "b"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HiddenWord(tt.args.numberChain); got != tt.want {
				t.Errorf("HiddenWord(%v) = %v, want %v", tt.args.numberChain, got, tt.want)
			}
		})
	}
}
