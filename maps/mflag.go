package mFlag

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"
)

var /* const */ colonRegexp = regexp.MustCompile("[^\\\\]:")

func split(flagString string) (key string, val string, err error) {
	matchIndices := colonRegexp.FindAllIndex([]byte(flagString), -1)

	if len(matchIndices) != 1 {
		err = errors.New("match indices not of length 1!")
		return
	}

	flagString = strings.Replace(flagString, "\\:", ":", -1)

	key = flagString[0 : matchIndices[0][1]-1]
	val = flagString[matchIndices[0][1]:]
	return
}

type MapStringString map[string]string

func NewMapStringString() MapStringString {
	return make(map[string]string)
}

func (m *MapStringString) String() string {
	bytes, _ := json.Marshal(*m) // probably shouldn't ignore this error?
	return string(bytes)
}

func (m *MapStringString) Set(flagString string) error {
	key, val, err := split(flagString)
	if err != nil {
		return err
	}
	(*m)[key] = val
	return nil
}

/* TODO: all other type combinations...
string:int
string:bool
int:string
int:int
int:bool
*/
