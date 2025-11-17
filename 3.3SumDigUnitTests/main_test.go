package main

import "testing"

func TestSumDigits(t *testing.T) {
	tests := []struct {
		numbers int
		want    int
	}{
		{-3, 3},
		{0, 0},
		{99999, 45},
		{-2147483648, 45},
	}
	for _, tests := range tests {
		got := SumDigits(tests.numbers)
		t.Logf("Calling SumDigits(%v), want %d  result %d\n", tests.numbers, tests.want, got)
		// expected := tests.want
		if got != tests.want {
			t.Errorf("Negative Test Error: Sumdigits(%d) expected %d, but got %d", tests.numbers, tests.want, got)
		}
	}
}

// numbers := 1
// want := 1
// for Result := SumDidits(numbers)

// t.Logf("Calling SumDigits(%d), result %d\n", numbers, want),

// if Result != want {
// t.Errorf("Incorrect result, Want %d, got %d", want, Result)
// }{
// 	{
// 		numbers: 1,
// 		want:    1,
// 	},
// {
// 	numbers: int(-123),
// 	want:    6,
// },
// {
// 	numbers: int(-125473),
// 	want:    22,
// },
// {
// 	numbers: int(99999),
// 	want:    45,
// },
// {0, 0},
// {-123, 6},
// {-125473, 22},
// {999999, 54},
// {100000001, 2},
// 	}

// }

// for testCase := range testTable {
// 	result := SumDidits(testCase.numbers)
// 	// t.Logf("Calling SumDigits(%v), result %d\n", testCase.numbers, result)

// 	if result != testCase.want {
// 		t.Errorf("Incorrect result, Want %d, got %d", testCase.want, result)

// testname := fmt.Sprintf("%d", tt.a)
// t.Run(testname, func(t *testing.T)) {
// 	ans := SumDigits(tt.a)
// 	if ans != tt.want{
// 		t.Errorf("got %d, want %d", ans, tt.want)
