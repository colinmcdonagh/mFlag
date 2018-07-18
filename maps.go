package mFlag

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var /* const */ colonRegexp = regexp.MustCompile("[^\\\\]:")

func split(flagString string) (key string, val string, err error) {
	matchIndices := colonRegexp.FindAllIndex([]byte(flagString), -1)

	if len(matchIndices) != 1 {
		err = errors.New("match indices not of length 1!")
		return
	}

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
	key = strings.Replace(key, "\\:", ":", -1)
	val = strings.Replace(val, "\\:", ":", -1)

	(*m)[key] = val
	return nil
}

type MapStringInt map[string]int

func NewMapStringInt() MapStringInt {
	return make(map[string]int)
}

func (m *MapStringInt) String() string {
	bytes, _ := json.Marshal(*m) // probably shouldn't ignore this error.
	return string(bytes)
}

func (m *MapStringInt) Set(flagString string) error {
	key, val, err := split(flagString)
	if err != nil {
		return err
	}
	key = strings.Replace(key, "\\:", ":", -1)
	var iVal int
	iVal, err = strconv.Atoi(val)
	if err != nil {
		return err
	}

	(*m)[key] = iVal
	return nil
}

type MapIntString map[int]string

func NewMapIntString() MapIntString {
	return make(map[int]string)
}

func (m *MapIntString) String() string {
	bytes, _ := json.Marshal(*m) // probably shouldn't ignore this error.
	return string(bytes)
}

func (m *MapIntString) Set(flagString string) error {
	key, val, err := split(flagString)
	if err != nil {
		return err
	}
	var iKey int
	iKey, err = strconv.Atoi(key)
	if err != nil {
		return err
	}
	val = strings.Replace(val, "\\:", ":", -1)

	(*m)[iKey] = val
	return nil
}

type MapIntInt map[int]int

func NewMapIntInt() MapIntInt {
	return make(map[int]int)
}

func (m *MapIntInt) String() string {
	bytes, _ := json.Marshal(*m) // probably shouldn't ignore this error.
	return string(bytes)
}

func (m *MapIntInt) Set(flagString string) error {
	key, val, err := split(flagString)
	if err != nil {
		return err
	}
	var (
		iKey, iVal int
	)
	iKey, err = strconv.Atoi(key)
	if err != nil {
		return err
	}
	iVal, err = strconv.Atoi(val)
	if err != nil {
		return err
	}

	(*m)[iKey] = iVal
	return nil
}
