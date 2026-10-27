package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
func genFn0(p0 uint) int8 {
    var v4 int64 = -9223372036854775808
    var v5 uint8 = 255
    _ = p0
    _ = v4
    _ = v5
    return -128
}

func main() {
    var v1 uint64 = 2
    v2 := []string{"foo bar", ("foo bar" + "hello")}
    v3 := make(map[int]uint32)
    var v6 int64 = 42
    v1 = 667
    var v7 int32 = 1
    if len(v2) > 0 {
        v2[0] = ("hello" + "hello")
    }
    v8 := GenS1{F0: v7, F1: 297}
    v9 := GenS1{F0: -1, F1: 255}
    v7 = v7
    fmt.Println(v6)
    var v10 uint64 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
