package hw05patterns_test

import (
	"fmt"
	"testing"

	hw05patterns "github.com/rez1dent3/go-education/hw05_patterns"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UniqSuite struct {
	suite.Suite
}

func TestUniqSuite(t *testing.T) {
	suite.Run(t, new(UniqSuite))
}

func (s *MaxSuite) TestUniq() {
	s.T().Parallel()

	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{2, 5, 3, 1, 4},
			expected: []int{2, 5, 3, 1, 4},
		},
		{
			input:    []int{-10, -10, 10},
			expected: []int{-10, 10},
		},
		{
			input:    []int{10, 0, -10, 10},
			expected: []int{10, 0, -10},
		},
		{
			input:    []int{},
			expected: []int{},
		},
		{
			input:    []int{2},
			expected: []int{2},
		},
	}

	for _, tt := range tests {
		s.Run(fmt.Sprintf("Uniq(%v)", tt.input), func() {
			assert.Equal(s.T(), tt.expected, hw05patterns.Uniq(tt.input))
		})
	}
}
