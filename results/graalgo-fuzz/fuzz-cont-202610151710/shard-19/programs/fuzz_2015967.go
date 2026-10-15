package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS2 struct { F0 int64; F1 int64 }
func genFn0(p0 uint64, p1 int8, p2 int8) bool {
    fmt.Println(p0, p2)
    fmt.Println(p0, p2)
    if false {
        for fi6 := 0; fi6 < 14; fi6++ {
            var v7 uint = 0
            _ = fi6
            _ = v7
        }
        fmt.Println(p2, p0)
    } else {
        p0 += (p0 + (p0 - p0))
        p2 -= p2
        v8 := GenS1{F0: 2, F1: 938, F2: 42}
        _ = v8
    }
    p0 -= (p0 & p0)
    _ = p0
    _ = p1
    _ = p2
    return true
}

func main() {
    var v1 int = -1
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]string)
    var v4 int64 = 1
    v5 := make(map[int]int)
    v9 := GenS2{F0: 1, F1: (v4 ^ (v4 & (v4 | v4)))}
    if (!((0.0 != (0.1 / ((1.5 / 3.14159) + (-2.25 + (((1.5 + 0.1) * 1.5) / -2.25))))) && true)) {
        var v10 int32 = 1
        _ = v10
    } else {
        for fi11 := 0; fi11 < 7; fi11++ {
            fmt.Println(fi11, fi11)
            _ = fi11
        }
        for fi12 := 0; fi12 < 17; fi12++ {
            fmt.Println(v1, v1)
            _ = fi12
        }
    }
    switch v1 {
    case 9223372036854775807:
        v1 += (v1 - v1)
    default:
        for fi13 := 0; fi13 < 9; fi13++ {
            fmt.Println(v4, fi13)
            v1 += 654
            _ = fi13
        }
    }
    v4 = (v4 | (v4 >> ((42) & 63)))
    var v14 float32 = 0.1
    v15 := GenS2{F0: 9223372036854775807, F1: v4}
    v16 := []int{v1}
    v4 -= 620
    for fi17 := 0; fi17 < 13; fi17++ {
        if (2 < 42) {
            var v18 uint = 1
            _ = v18
        }
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
        _ = fi17
    }
    v14 = -2.25
    v1 += v1
    switch -880 {
    case -366:
        var v19 float64 = 1e10
        _ = v19
    case 32767:
        v16 = append(v16, 42)
    case 2147483647:
        var v20 int64 = -1
        _ = v20
    default:
        if true {
            _ = len(v2)
            var v21 int32 = -2147483648
            var v22 int = (v1 >> ((42) & 63))
            _ = v21
            _ = v22
        } else {
            var v23 uint32 = 836
            var v24 uint64 = 366
            _ = v23
            _ = v24
        }
    }
    var v25 uint64 = 18446744073709551615
    fmt.Println(v1)
    _ = len(v3)
    var v26 uint8 = 126
    if false {
        var v27 int8 = 114
        var v28 int32 = -1
        fmt.Println(v28, v1, v4)
        _ = v27
        _ = v28
    } else {
        v14 = (3.14159 / v14)
        _ = len(v2)
    }
    _ = len(v16)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
