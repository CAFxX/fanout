package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 uint8; F2 int64 }
type GenS4 struct { F0 string; F1 int; F2 string }
func genFn0(p0 int16, p1 uint8) int8 {
    p1 = (^(p1))
    var v4 bool = (321 != -145)
    _ = p0
    _ = p1
    _ = v4
    return 25
}

func main() {
    var v1 float64 = 1e10
    var v2 uint16 = 41
    var v3 uint16 = (v2 >> ((2) & 63))
    fmt.Println(v2, v1, v2)
    v2 += v2
    v5 := []int{9223372036854775807}
    v3 -= v3
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := GenS1{F0: ((4294967295 > 0) && false)}
        _ = fi6
        _ = v7
    }
    v3 -= (v3 * 1)
    for fi8 := 0; fi8 < 14; fi8++ {
        switch "foo bar" {
        case "foo bar":
            var v9 float64 = 0.0
            _ = v9
        case "a":
            var v10 float64 = (1e10 / -2.25)
            _ = v10
        case "hello":
            v2 += 0
        default:
            v1 += v1
        }
        for fi11 := 0; fi11 < 2; fi11++ {
            v3 = 32767
            _ = fi11
        }
        _ = fi8
    }
    v3 += 1
    var v12 uint8 = 83
    var v13 uint32 = uint32(v3)
    if (42 == -538) {
        switch 42 {
        case 805:
            var v14 int = 2147483647
            _ = v14
        case -826:
            v15 := GenS2{F0: 9223372036854775807}
            _ = v15
        }
    }
    fmt.Println(v13)
    var v16 int8 = 0
    switch ("" + ("" + "")) {
    case "a":
        v16 = (v16 | v16)
    default:
        var v17 int = 42
        _ = v17
    }
    if len(v5) > 0 {
        v5[0] = 924
    }
    if (-1 <= 1) {
        v13 += v13
        var v18 uint8 = (v12 / ((1) | 1))
        if len(v5) > 0 {
            v5[0] = 991
        }
        _ = v18
    }
    v19 := GenS3{F0: true, F1: v12, F2: 42}
    v20 := GenS4{F0: "foo bar", F1: 136, F2: ("hello" + "foo bar")}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v5
}
