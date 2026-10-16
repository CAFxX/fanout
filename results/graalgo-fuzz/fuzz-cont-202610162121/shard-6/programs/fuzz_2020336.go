package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint16 }

func main() {
    var v1 uint = 346
    var v2 int32 = -2147483648
    var v3 int = 128
    var v4 uint16 = 255
    v5 := make(map[uint32]uint32)
    v6 := make(map[uint32]uint32)
    fmt.Println(v4, v1)
    switch "foo bar" {
    case "hello":
        var v7 uint8 = 38
        _ = v7
    case "":
        fmt.Println(v4, v4, v2)
    default:
        switch -2147483648 {
        case 77:
            var v8 uint64 = 83
            _ = v8
        case 2147483647:
            v2 -= (v2 | 640)
        }
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 int = (v3 << ((v1) & 63))
        var v11 int16 = -1
        for fi12 := 0; fi12 < 5; fi12++ {
            fmt.Println(v1)
            v4 -= v4
            _ = fi12
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    v13 := GenS1{F0: -49, F1: v3, F2: v4}
    for fi14 := 0; fi14 < 9; fi14++ {
        if (!true) {
            _ = len(v6)
            fmt.Println(v4, fi14, fi14)
        } else {
            var v15 uint32 = 2
            _ = v15
        }
        var v16 uint16 = 256
        _ = fi14
        _ = v16
    }
    if (false || false) {
        v1 += (^(v1))
        v17 := v6[158]
        var v18 uint8 = 127
        _ = v17
        _ = v18
    } else {
        _ = len(v6)
        v2 -= 491
        for fi19 := 0; fi19 < 6; fi19++ {
            var v20 uint8 = 24
            var v21 int32 = (-(v2))
            var v22 uint = 4294967295
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
    }
    var v23 float64 = 3.14159
    v2 -= (v2 >> ((42) & 63))
    v24 := []int{1, v3, 642, v3}
    v25 := GenS1{F0: 599, F1: v3, F2: 798}
    if len(v24) >= 2 {
        v24 = v24[1:len(v24)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
