package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint16 }
func genFn0(p0 bool, p1 uint64) int16 {
    p1 -= (p1 >> ((2) & 63))
    p1 = p1
    fmt.Println(p0, p1)
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return -32768
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 1
    var v3 string = ("a" + "")
    v4 := []uint32{0}
    switch -1 {
    case 628:
        fmt.Println(v2)
    default:
        var v5 int32 = 2147483647
        _ = v5
    }
    if true {
        _ = len(v4)
    } else {
        v6 := GenS1{F0: 42, F1: 0, F2: 65535}
        var v7 int8 = 1
        var v8 int32 = 1
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 uint8 = 128
    v2 = v2
    v3 = v3
    var v10 uint16 = 65535
    var v11 uint = 588
    fmt.Println(v9)
    for fi12 := 0; fi12 < 4; fi12++ {
        fmt.Println(v2)
        _ = fi12
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
