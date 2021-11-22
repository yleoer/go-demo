package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"negative + negative",
			args{-1, -1},
			-2,
		},
		{
			"negative + positive",
			args{-1, 1},
			0,
		},
		{
			"positive + positive",
			args{1, 1},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, Add(tt.args.x, tt.args.y), tt.want, tt.name)
		})
	}
}
