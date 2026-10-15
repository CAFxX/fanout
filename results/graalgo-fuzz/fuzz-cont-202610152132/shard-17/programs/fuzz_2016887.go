package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]uint32)
    var v3 int64 = -9223372036854775808
    v3 = v3
    v3 = (v3 >> ((42) & 63))
    if (false || (!(1 >= 42))) {
        var v4 uint = 255
        _ = v4
    } else {
        fmt.Println(v3)
        var v5 int64 = 0
        _ = v5
    }
    switch 128 {
    case 398:
        fmt.Println(v3)
    case -621:
        v6 := GenS1{F0: 42, F1: 65, F2: -409}
        _ = v6
    }
    var v7 int64 = v3
    var v8 bool = false
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 uint8 = 127
        var v11 int64 = v3
        _ = fi9
        _ = v10
        _ = v11
    }
    v7 = (v7 + 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
