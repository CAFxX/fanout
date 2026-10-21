package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := []int{}
    var v2 int16 = 42
    var v3 uint64 = 42
    var v4 int8 = 1
    var v5 uint16 = 32767
    v6 := make(map[int]uint32)
    var v7 int32 = 42
    _ = len(v6)
    switch 914 {
    case 526:
        var v8 int = -2147483648
        _ = v8
    case 44:
        fmt.Println(v5, v5)
    case 535:
        delete(v6, 0)
    default:
        for fi9 := 0; fi9 < 15; fi9++ {
            v2 = (-(v2))
            v3 = (v3 >> ((340) & 63))
            _ = fi9
        }
    }
    v10 := GenS1{F0: v2}
    for fi11 := 0; fi11 < 12; fi11++ {
        fmt.Println(fi11, v2, v5)
        _ = fi11
    }
    if true {
        var v12 uint = 0
        _ = v12
    }
    v13 := v6[-9223372036854775808]
    v7 = v7
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
