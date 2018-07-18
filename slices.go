package mFlag

import (
	"encoding/json"
	"strconv"
)

type SliceString []string

func NewSliceString() SliceString {
	return make([]string, 0)
}

func (s *SliceString) String() string {
	bytes, _ := json.Marshal(*s) // probably shouldn't ignore this error.
	return string(bytes)
}

func (s *SliceString) Set(flagString string) error {
	*s = append(*s, flagString)
	return nil
}

type SliceInt []int

func NewSliceInt() SliceInt {
	return make([]int, 0)
}

func (s *SliceInt) String() string {
	bytes, _ := json.Marshal(*s) // probably shouldn't ignore this error.
	return string(bytes)
}

func (s *SliceInt) Set(flagString string) error {
	i, err := strconv.Atoi(flagString)
	if err != nil {
		return err
	}
	*s = append(*s, i)
	return nil
}
