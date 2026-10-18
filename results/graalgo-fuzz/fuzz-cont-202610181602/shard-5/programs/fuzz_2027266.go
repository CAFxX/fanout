package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 int64 }
func genFn0(p0 int64, p1 uint16, p2 int) string {
    v6 := []int{(p2 ^ p2), 127, (p2 ^ 0)}
    for fi7 := 0; fi7 < 11; fi7++ {
        switch ("\n\t" + "hello") {
        case "\n\t":
            p0 = p0
        }
        p1 += (^(p1))
        _ = fi7
    }
    p0 -= (p0 % ((-10) | 1))
    if (469 != 851) {
        p1 = (p1 / (((p1 | (^(p1)))) | 1))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return ""
}
func genFn1(p0 bool, p1 uint8) int64 {
    switch 42 {
    case -963:
        v8 := GenS1{F0: -614, F1: -965}
        _ = v8
    case 127:
        p0 = false
    default:
        fmt.Println(p1, p1)
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 bool = true
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    return -1
}

func main() {
    var v1 uint64 = 42
    var v2 float32 = 1e10
    v3 := []string{}
    var v4 uint64 = 876
    v5 := make(map[uint32]uint32)
    v11 := GenS2{F0: 42, F1: 42}
    v12 := v5[2147483647]
    v3 = append(v3, "a")
    var v13 float32 = 3.14159
    v14 := GenS3{F0: 18446744073709551615}
    for fi15 := 0; fi15 < 18; fi15++ {
        v16 := GenS3{F0: 540}
        _ = fi15
        _ = v16
    }
    var v17 uint16 = 1
    if true {
        v18 := GenS4{F0: 1, F1: 9223372036854775807}
        v4 = 293
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
