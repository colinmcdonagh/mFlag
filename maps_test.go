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
		t.Errorf("mSS.String() == %s\nnot %s", mSS.String(), mapString)
	}
}

func Test_MapStringInt(t *testing.T) {

	mSI := NewMapStringInt()
	flagString := "person1:1"
	mSI.Set(flagString)
	flagString = "person2:2"
	mSI.Set(flagString)

	mapString := "{\"person1\":1,\"person2\":2}"
	if mSI.String() != mapString {
		t.Errorf("mSI.String() == %s\nnot %s", mSI.String(), mapString)
	}
}

func Test_MapIntString(t *testing.T) {

	mIS := NewMapIntString()
	flagString := "123:person1"
	mIS.Set(flagString)
	flagString = "234:person2"
	mIS.Set(flagString)
	flagString = "345:person3"
	mIS.Set(flagString)

	mapString := "{\"123\":\"person1\",\"234\":\"person2\",\"345\":\"person3\"}"
	// json marshals a map int index like a string in double quotes.
	if mIS.String() != mapString {
		t.Errorf("mIS.String() == %s\nnot %s", mIS.String(), mapString)
	}
}

func Test_MapIntInt(t *testing.T) {

	mII := NewMapIntInt()
	flagString := "1:10"
	mII.Set(flagString)
	flagString = "2:20"
	mII.Set(flagString)
	flagString = "3:30"
	mII.Set(flagString)

	mapString := "{\"1\":10,\"2\":20,\"3\":30}"
	// json marshals a map int index like a string in double quotes.
	if mII.String() != mapString {
		t.Errorf("mII.String() == %s\nnot %s", mII.String(), mapString)
	}
}
