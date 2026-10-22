package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 bool }

func main() {
    v1 := make(map[int]string)
    var v2 uint = 255
    v3 := []uint32{}
    var v4 uint = (v2 | v2)
    v5 := GenS1{F0: 0, F1: v4}
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := GenS2{F0: true}
        _ = fi6
        _ = v7
    }
    v4 = (v4 << ((v4) & 63))
    var v8 string = "hello"
    if false {
        var v9 uint32 = 4294967295
        _ = len(v1)
        _ = v9
    }
    switch v8 {
    case "\n\t":
        if len(v3) > 0 {
            v3[0] = 339
        }
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 int8 = 0
        var v12 uint64 = 2
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println(v2, v4)
    _ = len(v3)
    fmt.Println(v8, v8, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
