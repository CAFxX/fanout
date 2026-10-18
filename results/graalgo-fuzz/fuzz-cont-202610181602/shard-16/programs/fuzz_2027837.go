package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }

func main() {
    v1 := []uint32{}
    var v2 int32 = 1
    var v3 uint64 = 42
    var v4 uint64 = v3
    fmt.Println(v2)
    v2 += (v2 + (-(v2)))
    if false {
        v2 = (v2 >> ((18446744073709551615) & 63))
    }
    for fi5 := 0; fi5 < 1; fi5++ {
        v2 += v2
        var v6 string = "a"
        _ = fi5
        _ = v6
    }
    var v7 bool = true
    for fi8 := 0; fi8 < 9; fi8++ {
        v3 -= 0
        _ = fi8
    }
    fmt.Println(v2, v2, v3)
    _ = len(v1)
    var v9 uint = 42
    if (true || (((!true) || v7) && true)) {
        var v10 uint32 = 117
        var v11 int = 32768
        _ = v10
        _ = v11
    } else {
        v9 = (v9 >> (((v9 | uint(v3))) & 63))
        var v12 int = -986
        _ = v12
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        v7 = (!(!false))
        _ = fi13
    }
    v14 := GenS1{F0: 1, F1: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
