package awesomeProject

import "testing"

var formatTestParams = []struct {
	in  int
	out string
}{
	{0, "zero"},
	{1, "one"},
	{2, "two"},
	{3, "three"},
	{4, "four"},
	{5, "five"},
	{6, "six"},
	{7, "seven"},
	{8, "eight"},
	{9, "nine"},
	{10, "ten"},
	{11, "eleven"},
	{12, "twelve"},
	{13, "thirteen"},
	{14, "fourteen"},
	{15, "fifteen"},
	{16, "sixteen"},
	{17, "seventeen"},
	{18, "eighteen"},
	{19, "nineteen"},
	{20, "twenty"},
	{21, "twenty one"},
	{30, "thirty"},
	{40, "forty"},
	{121, "one hundred twenty one"},
	{122, "one hundred twenty two"},
	{1122, "one thousand one hundred twenty two"},
}

func Test_GivenInput_ReturnsOutput(t *testing.T) {
	for _, params := range formatTestParams {
		result, err := format(params.in)
		if err != nil {
			t.Errorf("Expected no error but got %v", err)
		}
		if result != params.out {
			t.Errorf("Expected %v but got %v", params.out, result)
		}
	}
}

func Test_GivenNegativeNumber_Errors(t *testing.T) {
	_, err := format(-1)
	if err == nil {
		t.Error("Expected error but didn't get one")
	}
}