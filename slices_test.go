package mFlag

import (
	"testing"
)

func Test_SliceString(t *testing.T) {

	ss := NewSliceString()
	flagString := "a"
	ss.Set(flagString)
	flagString = "b"
	ss.Set(flagString)
	flagString = "c"
	ss.Set(flagString)

	sliceString := "[\"a\",\"b\",\"c\"]"
	if ss.String() != sliceString {
		t.Errorf("ss.String() == %s\nnot %s\n", ss.String(), sliceString)
	}
}

func Test_SliceInt(t *testing.T) {

	si := NewSliceInt()
	flagString := "0"
	si.Set(flagString)
	flagString = "1"
	si.Set(flagString)
	flagString = "2"
	si.Set(flagString)

	sliceString := "[0,1,2]"
	if si.String() != sliceString {
		t.Errorf("si.String() == %s\nnot %s\n", si.String(), sliceString)
	}
}
