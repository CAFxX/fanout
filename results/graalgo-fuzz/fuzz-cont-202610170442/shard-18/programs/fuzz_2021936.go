package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 585
    var v3 int32 = -1
    v4 := []int{42}
    v5 := make(map[uint32]uint32)
    var v6 uint32 = 0
    v7 := GenS1{F0: (-(v3)), F1: 42}
    if (v2 == v2) {
        delete(v1, 9223372036854775807)
        v1[454] = "hello"
    }
    var v8 uint64 = v2
    for fi9 := 0; fi9 < 4; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            fmt.Println(v2)
            var v11 uint = 0
            _ = fi10
            _ = v11
        }
        var v12 int16 = 1000
        for fi13 := 0; fi13 < 9; fi13++ {
            v12 = v12
            _ = fi13
        }
        _ = fi9
        _ = v12
    }
    v4 = append(v4, 32767)
    if false {
        var v14 int = 9223372036854775807
        var v15 int16 = 1000
        _ = v14
        _ = v15
    }
    v1[-1] = "a"
    v8 = v8
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
