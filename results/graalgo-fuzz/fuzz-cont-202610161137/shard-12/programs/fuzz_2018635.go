package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 bool }
type GenS2 struct { F0 string; F1 int16; F2 uint32 }
type GenS3 struct { F0 uint16; F1 int64; F2 string }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]uint32)
    var v3 string = ""
    var v4 float64 = (0.1 + 0.0)
    v5 := make(map[uint32]int)
    var v6 int16 = -561
    v7 := GenS1{F0: -2147483648, F1: 0, F2: (!(!(-9223372036854775808 <= 735)))}
    v3 = v3
    fmt.Println(v3)
    var v8 int16 = (-(v6))
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := GenS2{F0: v3, F1: v8, F2: 2}
        for fi11 := 0; fi11 < 19; fi11++ {
            fmt.Println(v8, v3)
            _ = fi11
        }
        _ = fi9
        _ = v10
    }
    if (-1 == 0) {
        fmt.Println(v4)
    }
    v12 := GenS1{F0: 127, F1: 65535, F2: false}
    switch v3 {
    case "foo bar":
        v3 = v3
    case "a":
        v13 := []int{543, 563, -9223372036854775808}
        _ = v13
    default:
        var v14 string = "a"
        _ = v14
    }
    switch 2147483647 {
    case 9223372036854775807:
        fmt.Println(v6, v6, v6)
    case 0:
        var v15 int16 = -32768
        _ = v15
    case 212:
        v16 := GenS3{F0: 257, F1: 1, F2: ""}
        _ = v16
    default:
        v8 -= (-(v6))
    }
    v17 := v5[42]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
