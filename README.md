# mFlag

Support for flags of maps and slices.

## Map Usage 
map\_example.go:
```go
graphOfFriends := lFlag.NewMapStringString() 

flag.Var(&graphOfFriends, "m", "graphOfFriends)
flag.Parse()

fmt.Println(graphOfFriends) // graphOfFriends.String()
```
-------------------------------------------
shell:
```shell
go run map_example.go -m person1:person2 -m person2:person3 -m someoneElse:person1

{"person1":"person2","person2":"person3","someoneElse":"person1"}
```
## Slice Usage
slice\_example.go:
```go
```
-------------------------------------------
shell:
```shell
```

## New Functions
Each 'New' function, such as NewMapStringString or NewSliceInt, returns a specific type; map[string]string and []int in this case.
A generic approach of using any map[Type]Type or []Type isn't used since the map and slice datatypes cannot be extended directly. Therefore, to extend maps, one approach would be to extend specific types, such as map[string]string, or extend a map[interface{}]interface{}.
The difficulty with the latter is that another argument for the type of map would have to be supplied on the command line, this flag would then
be parsed before parsing the map/slice flags, and then the map[interface{}]interface{} or []interface{} reference supplied to flag.Var
would then need to be cast by the user to the type input on the command line.

Therefore, the former approach of defining a 'New' function for each possible type is taken. The type combinations are as follows:
* map[string]string
* map[string]int
* map[int]string
* map[int]int
* []string
* []int

## Encoded Semicolons

Escaped semicolons will be ignored and the leading backslash removed when used in a map type! (but not with slices)

For example:
-m aPunctuationMark:\\\\: --> {"aPunctuationMark":":"}
