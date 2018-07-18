# lFlag

Support for flags of maps, slices and arrays.

## Usage
### Maps
```go
graphOfFriends := lFlag.NewMapStringString()

flag.Var(&graphOfFriends, "m", "graphOfFriends)
flag.Parse()

fmt.Println(graphOfFriends)
```

-------------------------------------------
```shell
go build example.go
./example.go -m person1:person2 -m person2:person3 -m someoneElse:person1

{"person1":"person2","person2":"person3","someoneElse":"person1"}
```
------------------------------------------

Escaped semicolons will be ignored and the leading backslash removed!

For example:
-m aPunctuationMark:\\\\:

->

{"aPunctuationMark":":"}
