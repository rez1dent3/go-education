package hw05patterns_test

import (
	"fmt"
	"testing"

	hw05patterns "github.com/rez1dent3/go-education/hw05_patterns"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MinSuite struct {
	suite.Suite
}

func TestMinSuite(t *testing.T) {
	suite.Run(t, new(MinSuite))
}

func (s *MaxSuite) TestMin() {
	s.T().Parallel()

	tests := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{2, 5, 3, 1, 4},
			expected: 1,
		},
		{
			input:    []int{-10, 10},
			expected: -10,
		},
		{
			input:    []int{0, -10, 10},
			expected: -10,
		},
		{
			input:    []int{},
			expected: 0,
		},
		{
			input:    []int{2},
			expected: 2,
		},
	}

	for _, tt := range tests {
		s.Run(fmt.Sprintf("Min(%v)", tt.input), func() {
			assert.Equal(s.T(), tt.expected, hw05patterns.Min(tt.input))
		})
	}
}
