package mFlag

import (
	"testing"
)

func Test_MapStringString(t *testing.T) {

	mSS := NewMapStringString()
	flagString := "person1:person2"
	mSS.Set(flagString)
	flagString = "person2:per\\:son3"
	mSS.Set(flagString)

	mapString := "{\"person1\":\"person2\",\"person2\":\"per:son3\"}"
	if mSS.String() != mapString {
		t.Errorf("mSS.String == %s\nnot %s", mSS.String(), mapString)
	}
}
