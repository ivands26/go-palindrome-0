package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountPalindrome(t *testing.T) {
	type TestCase struct {
		Param1   int
		Param2   int
		Expected int
		Message  string
	}

	arrTestCase := []TestCase{
		{1, 10, 9, "Wrong Asnwer"},
		{100, 150, 5, "Wrong Asnwer"},
		{99, 100, 1, "Wrong Asnwer"},
		{21, 31, 1, "Wrong Asnwer"},
	}

	for i := 0; i < len(arrTestCase); i++ {
		t.Run(fmt.Sprint("Case", i+1), func(t *testing.T) {
			assert.Equal(t, arrTestCase[i].Expected, CountPalindrome(arrTestCase[i].Param1, arrTestCase[i].Param2), arrTestCase[i].Message)
		})
	}
}
