package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int, p1 int32) int32 {
    var v6 uint8 = 0
    v7 := GenS1{F0: 1, F1: 2}
    v8 := GenS1{F0: 114, F1: 429}
    p1 = p1
    v6 = v6
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return (p1 | (p1 / ((42) | 1)))
}

func main() {
    v1 := []uint32{}
    v2 := []int{-406, -9223372036854775808, 9223372036854775807}
    var v3 uint16 = 545
    v4 := make(map[uint32]bool)
    var v5 int32 = 2147483647
    fmt.Println(v3, v5)
    v5 -= (v5 >> ((18446744073709551615) & 63))
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 uint16 = v3
        var v11 int16 = 188
        _ = fi9
        _ = v10
        _ = v11
    }
    v5 += v5
    for fi12 := 0; fi12 < 1; fi12++ {
        switch "hello" {
        case "foo bar":
            v13 := GenS1{F0: 100, F1: 59}
            _ = v13
        default:
            var v14 float32 = 0.1
            _ = v14
        }
        v15 := GenS2{F0: 128}
        _ = fi12
        _ = v15
    }
    v16 := GenS1{F0: 127, F1: 42}
    var v17 bool = (2 == 975)
    v5 = v5
    var v18 uint32 = 402
    for fi19 := 0; fi19 < 16; fi19++ {
        v5 += v5
        _ = fi19
    }
    v20 := GenS1{F0: 57, F1: 42}
    v3 = v3
    _ = len(v2)
    if ("" >= "") {
        fmt.Println(v18, v5, v3)
    }
    v17 = v17
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
