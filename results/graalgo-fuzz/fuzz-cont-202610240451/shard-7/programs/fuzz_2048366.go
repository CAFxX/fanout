package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int8; F2 uint64 }
func genFn0(p0 uint64, p1 uint16) int32 {
    var v5 string = (("foo bar" + "a") + ("foo bar" + "foo bar"))
    v5 = "a"
    v5 = v5
    _ = p0
    _ = p1
    _ = v5
    return 0
}

func main() {
    var v1 bool = false
    v2 := []string{}
    var v3 int16 = 42
    v4 := make(map[string]uint32)
    v1 = (v1 || v1)
    v6 := GenS1{F0: "\n\t"}
    v7 := GenS2{F0: 2147483647, F1: 100, F2: 1}
    v8 := GenS1{F0: (("hello" + "\n\t") + "")}
    v1 = (275 != 831)
    v3 = (v3 | 334)
    _ = len(v4)
    var v9 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
