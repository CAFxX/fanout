package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 int }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]uint32)
    var v3 string = (("foo bar" + ("a" + (("a" + "hello") + ("hello" + "foo bar")))) + "")
    v4 := make(map[uint32]bool)
    var v5 bool = (1 <= 779)
    v5 = (32768 != -239)
    v5 = (!(!(4294967295 <= 1)))
    var v6 int64 = 42
    v7 := GenS1{F0: 112}
    var v8 int32 = -60
    v9 := GenS2{F0: -751, F1: 32768}
    var v10 uint64 = 807
    fmt.Println(v10, v3, v6)
    v8 -= -1
    v10 -= v10
    if v5 {
        fmt.Println(v3)
    } else {
        var v11 int8 = 100
        _ = v11
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
