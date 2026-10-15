package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint16; F1 uint8 }

func main() {
    v1 := []uint32{4294967295}
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]bool)
    var v4 int = 1
    v5 := make(map[uint32]bool)
    v6 := []int{v4, 601, 42}
    v4 = v4
    fmt.Println(v4)
    for fi7 := 0; fi7 < 5; fi7++ {
        fmt.Println(v4, fi7)
        _ = fi7
    }
    if (1e10 > -2.25) {
        switch v4 {
        case 766:
            var v8 uint16 = 42
            _ = v8
        default:
            v9 := GenS1{F0: 551, F1: -4}
            _ = v9
        }
    }
    switch (v4 * (v4 % (((-(v4))) | 1))) {
    case 9223372036854775807:
        var v10 int8 = 69
        _ = v10
    case 2147483647:
        v11 := GenS2{F0: 652, F1: 80}
        _ = v11
    default:
        if true {
            var v12 int32 = 42
            var v13 int16 = 32767
            _ = v12
            _ = v13
        }
    }
    v4 -= v4
    var v14 int8 = -1
    var v15 int16 = 42
    var v16 string = ("a" + "foo bar")
    fmt.Println(v14, v4, v14)
    v14 = v14
    switch v4 {
    case -1:
        v16 = v16
    case 0:
        fmt.Println(v16)
    case -284:
        _ = len(v1)
    }
    var v17 uint = 18446744073709551615
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
