package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 string }
func genFn0(p0 uint16) uint32 {
    if (2 <= 132) {
        v6 := GenS1{F0: 42, F1: 1, F2: "hello"}
        var v7 uint16 = p0
        _ = v6
        _ = v7
    } else {
        fmt.Println(p0)
    }
    var v8 uint64 = 608
    fmt.Println(p0)
    v8 = (v8 & v8)
    v9 := GenS1{F0: -9223372036854775808, F1: v8, F2: ""}
    _ = p0
    _ = v8
    _ = v9
    return 427
}
func genFn1(p0 string, p1 int8) int8 {
    p0 = p0
    v10 := GenS1{F0: -1, F1: 2, F2: "foo bar"}
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := []int{127, (-(fi11))}
        for fi13 := 0; fi13 < 8; fi13++ {
            if (!(2 < 1)) {
                v12 = append(v12, -1)
                var v14 int32 = 0
                fmt.Println(fi13)
                _ = v14
            }
            p1 = (p1 * 11)
            _ = fi13
        }
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    return (p1 | (p1 << ((1) & 63)))
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]uint32)
    var v3 int16 = -737
    var v4 uint64 = 1
    v5 := make(map[int]uint32)
    fmt.Println(v4, v4)
    if (!true) {
        if (genFn1(("\n\t" + ""), 1) != -116) {
            var v15 int32 = 0
            _ = v15
        } else {
            var v16 int8 = -1
            if true {
                v1[46] = 0
                v4 = v4
                var v17 int16 = -32768
                _ = v17
            }
            _ = v16
        }
        var v18 uint64 = v4
        _ = v18
    }
    if (!false) {
        var v19 uint = 241
        _ = v19
    } else {
        if true {
            fmt.Println(v4, v3)
        }
    }
    var v20 uint8 = 0
    v21 := []int{127, 1}
    v21 = append(v21, -2147483648)
    for fi22 := 0; fi22 < 12; fi22++ {
        v4 = 42
        _ = fi22
    }
    var v23 int32 = -1
    v20 = v20
    var v24 uint8 = v20
    var v25 uint32 = 583
    delete(v2, 84)
    var v26 uint64 = uint64(v24)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
