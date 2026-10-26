package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 int32 = 659
    v2 := make(map[int]int)
    var v3 int32 = v1
    for fi4 := 0; fi4 < 3; fi4++ {
        v5 := GenS1{F0: -850, F1: true}
        _ = fi4
        _ = v5
    }
    _ = len(v2)
    var v6 uint8 = 228
    v7 := GenS2{F0: 1}
    if true {
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 string = (("foo bar" + "") + "")
            fmt.Println(v9)
            _ = fi8
            _ = v9
        }
        var v10 float64 = -2.25
        _ = v10
    }
    v1 -= (v1 << ((0) & 63))
    var v11 uint32 = 0
    delete(v2, -1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
