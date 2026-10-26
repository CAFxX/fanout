package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int }
func genFn0(p0 uint8, p1 int8, p2 int8) uint64 {
    var v4 uint = 4294967295
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 int8 = (p1 & (p2 >> (((v4 + v4)) & 63)))
        switch "" {
        case "foo bar":
            var v7 uint32 = 128
            _ = v7
        case "a":
            var v8 float64 = ((0.1 / 1.5) * (float64(0.0) + 0.1))
            _ = v8
        default:
            v9 := []int{}
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    p2 = -1
    if (true && false) {
        v10 := []int{-9223372036854775808, 1}
        _ = v10
    } else {
        if ((65535 == uint32(p0)) && (2147483647 <= 42)) {
            var v11 float64 = -2.25
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}

func main() {
    var v1 uint8 = 127
    v2 := []string{"a"}
    v3 := []uint32{}
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := GenS1{F0: -1}
        _ = fi12
        _ = v13
    }
    v14 := GenS2{F0: 0, F1: (v1 / (((^(v1))) | 1))}
    var v15 uint8 = v1
    if len(v2) > 0 {
        v2[0] = (("foo bar" + "\n\t") + "\n\t")
    }
    v15 += 128
    for fi16 := 0; fi16 < 13; fi16++ {
        v17 := GenS2{F0: 1, F1: 161}
        _ = fi16
        _ = v17
    }
    if false {
        v2 = append(v2, "\n\t")
        switch 773 {
        case 774:
            var v18 string = "hello"
            _ = v18
        case 2147483647:
            v19 := GenS3{F0: -923}
            _ = v19
        default:
            fmt.Println(v15)
        }
    }
    if len(v3) > 0 {
        v3[0] = 910
    }
    v20 := GenS2{F0: -9223372036854775808, F1: v15}
    var v21 int8 = 100
    var v22 float64 = float64(0.0)
    _ = len(v3)
    v15 = (v15 + v15)
    fmt.Println(v22, v1)
    for fi23 := 0; fi23 < 19; fi23++ {
        _ = len(v3)
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
