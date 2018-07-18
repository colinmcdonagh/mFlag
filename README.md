# mFlag

Support for map and slice flags (containing 'string's and 'int's).

## Map Usage 
map\_example.go:
```go
graphOfFriends := lFlag.NewMapStringString() 

flag.Var(&graphOfFriends, "m", "a map containing a directed graph of who considers who a friend")
flag.Parse()

graphOfFriends["some1else"] = "some2else"

fmt.Println(graphOfFriends) // graphOfFriends.String()
```
-------------------------------------------
shell:
```shell
go run map_example.go -m person1:person2 -m person2:person3 

{"person1":"person2","person2":"person3","some1else":"some2else"}
```
## Slice Usage
slice\_example.go:
```go
list := lFlag.NewSliceInt()

flag.Var(&list, "l", "a list of numbers")
flag.Parse()

l = append(l, 99)

fmt.Println(list) // list.String()
```
-------------------------------------------
shell:
```shell
go run slice_example.go -l 0 -l 1 -l 2

[0,1,2,99]
```

## 'New' Functions
Each 'New' function, such as NewMapStringString or NewSliceInt, returns a specific type; map[string]string and []int in this case.
A generic approach of using any map[Type]Type or []Type isn't used since the map and slice datatypes cannot be extended directly. Therefore, to extend maps, one approach would be to extend specific types, such as map[string]string, or extend a map[interface{}]interface{}.
The difficulty with the latter is that another argument for the type of map would have to be supplied on the command line, this flag would then
be parsed before parsing the map/slice flags, and then the map[interface{}]interface{} or []interface{} reference supplied to flag.Var
would then need to be cast by the user to the type input on the command line.

Therefore, the former approach of defining a 'New' function for each possible type is taken. The type combinations are as follows:
* map[string]string - NewMapStringString()
* map[string]int - NewMapStringInt()
* map[int]string - NewMapIntString()
* map[int]int - NewMapIntInt()
* []string - NewSliceString()
* []int - NewSliceInt()

## Encoded Semicolons

Escaped semicolons will be ignored and the leading backslash removed when used in a map type! (but not with slices)

For example:
-m aPunctuationMark:\\\\: --> {"aPunctuationMark":":"}
