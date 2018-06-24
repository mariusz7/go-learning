package mystringbuilder


import (
	"testing"
	"strings"
)

func doSingleTest(caption string, appended []string, expected string,
	t *testing.T) MyStringBuilder {

	var sb MyStringBuilder	

	for _, str := range appended {
		sb.Append(str)
	}

	var concStr string = sb.String()

	if concStr != expected {
		t.Errorf("%v - Got actual: %v, but expected: %v\n",
			caption, concStr, expected)
	}

	return sb
}

func doSingleTestOnSB(caption string, appended []string, expected string,
	sb MyStringBuilder, t *testing.T) MyStringBuilder {

	for _, str := range appended {
		sb.Append(str)
	}

	var concStr string = sb.String()

	if concStr != expected {
		t.Errorf("%v - Got actual: %v, but expected: %v\n",
			caption, concStr, expected)
	}
	return sb
}

func TestMyStringBuilder(t *testing.T) {

	appended := []string{"", "", "", "", ""}
	expected := ""
	doSingleTest("Empty strings", appended, expected, t)

	appended = []string{"a", "b", "", "c", ""}
	expected = "abc"
	doSingleTest("Mixup empty strings and single letters", appended, expected, t)

	appended = []string{"aaaaaaaaaa", "b", "ccccccccccccccccccc",
		"dddddddddddddd", "111111111111111"}
	expected = "aaaaaaaaaabcccccccccccccccccccdddddddddddddd111111111111111"
	doSingleTest("Non empty various size strings", appended, expected, t)

	appended = []string{"aaaaaaaaaa", "b", "ccccccccccccccccccc",
		"dddddddddddddd", "111111111111111"}
	expected = "aaaaaaaaaabcccccccccccccccccccdddddddddddddd111111111111111"
	var sb MyStringBuilder = doSingleTest(
		"Non empty various size strings 2", appended, expected, t)

	appended = []string{"", "", "", "", "z"}
	expected = "aaaaaaaaaabcccccccccccccccccccdddddddddddddd111111111111111z"
	doSingleTestOnSB("Appended something to already used string buffer",
		appended, expected, sb, t)
}


func TestMyStringBuilderInALoop1(t *testing.T) {

	var mySB MyStringBuilder
	var stdSB strings.Builder

	for i := 0; i < 1000; i++ {
		mySB.Append(string(i))
		stdSB.WriteString(string(i))

		myStr := mySB.String()
		stdStr := stdSB.String()

		if myStr != stdStr {
			t.Fatalf("In loop %v actual: %v, but expected: %v\n",
				i, myStr, stdStr)
		}
	}
}

func TestMyStringBuilderInALoop2(t *testing.T) {

	var mySB MyStringBuilder
	var stdSB strings.Builder

	for i := 0; i < 1000; i++ {
		mySB.Append(string(i))
		stdSB.WriteString(string(i))
	}

	myStr := mySB.String()
	stdStr := stdSB.String()

	if myStr != stdStr {
		t.Fatalf("Test failed; actual: %v, but expected: %v\n",
			myStr, stdStr)
	}
}

