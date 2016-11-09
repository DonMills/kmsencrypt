package filefuncs

import (
	"fmt"
	"testing"
)

func TestConcatSplit(t *testing.T) {
	testdata := []byte("This is the test data")
	testkey := []byte("This is a test key")
	testiv := []byte("This is a test iv")
	fmt.Printf("Test Data: %v\nTest Key: %v\nTest IV: %v\n", testdata, testkey, testiv)
	concatfile := CreateEncFile(testdata, testiv, testkey)
	fmt.Printf("Concatenated Data: %v\n", concatfile)
	splitdata, splitiv, splitkey := SplitEncFile(concatfile)
	fmt.Printf("Split Data: %v\nSplit Key: %v\nSplit IV: %v\n", splitdata, splitkey, splitiv)
	if string(testdata) != string(splitdata) {
		t.Error("Concat/split failed - bad data! Error in CreateEncFile/SplitEncFile functions")
	} else if string(testkey) != string(splitkey) {
		t.Error("Concat/split failed - bad key! Error in CreateEncFile/SplitEncFile functions")
	} else if string(testiv) != string(splitiv) {
		t.Error("Concat/split failed - iv! Error in CreateEncFile/SplitEncFile functions")
	} else {
		fmt.Printf("All test data matches\n")
	}
}
