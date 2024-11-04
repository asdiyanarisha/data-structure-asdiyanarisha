package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type TableTest struct {
	s        string
	goal     string
	expected bool
	nameTest string
}

func rotateString2(s string, goal string) bool {
	sSplit := strings.Split(s, "")
	goalSplit := strings.Split(goal, "")

	if len(sSplit) != len(goalSplit) {
		return false
	}

	fmt.Println(sSplit, goalSplit)
	var sIdx int

	// find sIdx

	for i := 0; i < len(goalSplit); i++ {
		if goalSplit[0] == sSplit[i] {
			sIdx = i
			break
		}
	}
	fmt.Println("begin sIdx", sIdx)

	var result bool
	for i := 0; i < len(goalSplit); i++ {

		if goalSplit[i] == sSplit[sIdx] {
			result = true
			if sIdx == len(sSplit)-1 {
				sIdx = 0
				continue
			}

			sIdx++
		} else {

			fmt.Println("sIdx", sIdx)
			result = false
			break
		}
	}

	return result
}

func rotateString(s string, goal string) bool {
	goalSplit := strings.Split(goal, "")

	if len(strings.Split(s, "")) != len(goalSplit) {
		return false
	}

	sSplit := strings.Split(s+s, "")

	// find sIdx
	for i := 0; i < len(sSplit); i++ {

		var start int
		startS := i
		run := true

		for start < len(goalSplit) {
			if startS > len(sSplit)-1 {
				startS = 0
			}

			if goalSplit[start] == sSplit[startS] {
				start++
				startS++
			} else {
				run = false
				break
			}
		}

		if !run {
			continue
		}

		return run
	}

	return false
}

func TestRotateString(t *testing.T) {
	tableTests := []TableTest{
		{
			s:        "abcde",
			goal:     "cdeab",
			expected: true,
			nameTest: "Test Case 1",
		},
		{
			s:        "abcde",
			goal:     "abced",
			expected: false,
			nameTest: "Test Case 2",
		},
		{
			s:        "defdefdefabcabc",
			goal:     "defdefabcabcdef",
			expected: true,
			nameTest: "Test Case 3",
		},
		{
			s:        "xjab",
			goal:     "abcd",
			expected: false,
			nameTest: "Test Case 4",
		},
	}

	for _, testCase := range tableTests {
		t.Run(testCase.nameTest, func(t *testing.T) {
			got := rotateString(testCase.s, testCase.goal)
			assert.Equal(t, testCase.expected, got)
		})
	}

}
