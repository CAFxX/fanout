package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 int64; F1 uint64 }
func genFn0() uint8 {
    var v6 int = -196
    var v7 float32 = 1e10
    v8 := []int{(v6 - 128), (v6 << ((445) & 63)), v6}
    fmt.Println(v7, v7)
    for fi9 := 0; fi9 < 8; fi9++ {
        v8 = append(v8, v6)
        v10 := GenS1{F0: 2, F1: 1}
        _ = fi9
        _ = v10
    }
    _ = v6
    _ = v7
    _ = v8
    return 128
}
func genFn1(p0 int64, p1 int32, p2 int32) int64 {
    for fi11 := 0; fi11 < 18; fi11++ {
        v12 := GenS1{F0: 42, F1: p0}
        v13 := GenS1{F0: 896, F1: 1}
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println(p1, p1, p1)
    for fi14 := 0; fi14 < 12; fi14++ {
        if (false || (true || false)) {
            var v15 bool = true
            p0 += -9223372036854775808
            for fi16 := 0; fi16 < 16; fi16++ {
                p1 += p1
                _ = fi16
            }
            _ = v15
        }
        _ = fi14
    }
    switch 2147483647 {
    case 128:
        v17 := GenS1{F0: 18446744073709551615, F1: (p0 ^ -9223372036854775808)}
        _ = v17
    case -1:
        for fi18 := 0; fi18 < 4; fi18++ {
            p0 = (p0 | p0)
            if (false && (!false)) {
                p2 = (p2 | p1)
            } else {
                fmt.Println(p0)
                p1 = -268
            }
            var v19 int = 42
            _ = fi18
            _ = v19
        }
    }
    var v20 uint = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return p0
}

func main() {
    var v1 int = -9223372036854775808
    var v2 int64 = 42
    var v3 uint8 = 255
    v4 := make(map[int]string)
    v5 := []int{1, v1}
    var v21 uint = 476
    fmt.Println(v2, v3, v2)
    v21 += (v21 >> ((1) & 63))
    if len(v5) > 0 {
        v5[0] = v1
    }
    v22 := GenS2{F0: (v2 | 9223372036854775807), F1: 4294967295}
    for fi23 := 0; fi23 < 17; fi23++ {
        v4[fi23] = ("hello" + ("" + "hello"))
        v21 = 65535
        _ = fi23
    }
    v5 = append(v5, int(v3))
    v2 = v2
    v3 = (v3 >> (((v21 & 1)) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
