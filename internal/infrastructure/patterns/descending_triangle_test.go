package patterns

import "testing"

func TestDescendingTrianglePattern_Generate(t *testing.T) {
	tests := []struct {
		name  string
		size  int
		down  int
		right int
		want  string
	}{
		{
			name:  "3 size with 5 repeated horizontally",
			size:  3,
			down:  1,
			right: 5,
			want:  "***************\n** ** ** ** **\n*  *  *  *  *",
		},
		{
			name:  "invalid size",
			size:  0,
			down:  1,
			right: 5,
			want:  "",
		},
		{
			name:  "invalid down",
			size:  3,
			down:  0,
			right: 5,
			want:  "",
		},
		{
			name:  "invalid right",
			size:  3,
			down:  1,
			right: 0,
			want:  "",
		},
	}

	pattern := DescendingTrianglePattern{}

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
