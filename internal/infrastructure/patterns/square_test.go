package patterns

import "testing"

func TestSquarePattern_Generate(t *testing.T) {
	tests := []struct {
		name  string
		size  int
		down  int
		right int
		want  string
	}{
		{
			name:  "3x3 square",
			size:  3,
			down:  3,
			right: 1,
			want:  "***\n***\n***",
		},
		{
			name:  "invalid size",
			size:  0,
			down:  3,
			right: 1,
			want:  "",
		},
		{
			name:  "invalid down",
			size:  3,
			down:  0,
			right: 1,
			want:  "",
		},
		{
			name:  "invalid right",
			size:  3,
			down:  3,
			right: 0,
			want:  "",
		},
	}

	pattern := &SquarePattern{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pattern.Generate(
				tt.size,
				tt.down,
				tt.right,
			)

			if got != tt.want {
				t.Errorf(
					"Generate() = %q, want %q",
					got,
					tt.want,
				)
			}
		})
	}
}
