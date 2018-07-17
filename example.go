package main

import (
	mFlag "./maps"

	"flag"
	"fmt"
)

func main() {
	graphOfFriends := mFlag.NewMapStringString()
	flag.Var(&graphOfFriends, "m", "graph of friends")
	flag.Parse()

	fmt.Println(graphOfFriends.String())
}
