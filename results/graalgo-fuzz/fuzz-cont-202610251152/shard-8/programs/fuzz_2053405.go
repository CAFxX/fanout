package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS4 struct { F0 int32 }
func genFn0(p0 bool, p1 uint16, p2 uint32) uint8 {
    if p0 {
        v7 := GenS1{F0: p1, F1: 2, F2: 1}
        p0 = (4294967295 >= 39)
        _ = v7
    }
    p0 = p0
    v8 := GenS1{F0: (p1 % ((p1) | 1)), F1: p1, F2: 18446744073709551615}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 42
}
func genFn1() int16 {
    var v9 int32 = 583
    if (-912 > 32768) {
        v9 += -361
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        v9 = (v9 * (v9 ^ (v9 & v9)))
        _ = fi10
    }
    _ = v9
    return -32768
}

func main() {
    v1 := []uint32{412, 445}
    var v2 int64 = 42
    var v3 uint = 4294967295
    var v4 uint32 = 2147483648
    v5 := make(map[string]bool)
    v6 := []uint32{(v4 >> ((2) & 63)), (v4 / ((42) | 1)), (v4 + 824)}
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 uint16 = 32767
        v2 = -9223372036854775808
        for fi13 := 0; fi13 < 20; fi13++ {
            var v14 bool = true
            switch (fi11 % ((975) | 1)) {
            case 2147483647:
                var v15 int16 = 813
                v16 := GenS1{F0: (v12 | (v12 % ((881) | 1))), F1: (v12 % ((v12) | 1)), F2: 2}
                _ = len(v5)
                _ = v15
                _ = v16
            case -871:
                v2 += v2
            case 111:
                v17 := GenS1{F0: 1, F1: (v12 + v12), F2: 42}
                _ = v17
            }
            _ = fi13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    v3 -= (^(v3))
    v2 -= (v2 - v2)
    v3 += (v3 + (^(v3)))
    fmt.Println(v3, v4)
    fmt.Println(v3, v2, v4)
    if (false || true) {
        fmt.Println(v2, v3, v2)
        v3 = 73
        var v18 uint = (v3 & 0)
        _ = v18
    }
    v19 := v5["hello"]
    for fi20 := 0; fi20 < 13; fi20++ {
        var v21 float64 = 1.5
        fmt.Println(v4, v3)
        var v22 int64 = v2
        _ = fi20
        _ = v21
        _ = v22
    }
    if v19 {
        v23 := GenS1{F0: 42, F1: 0, F2: 98}
        v24 := GenS2{F0: 1, F1: 2}
        _ = v23
        _ = v24
    } else {
        var v25 uint16 = 199
        _ = v25
    }
    if false {
        if (1e10 >= 3.14159) {
            var v26 int = -100
            _ = v26
        }
    }
    v4 -= (v4 / ((v4) | 1))
    for fi27 := 0; fi27 < 17; fi27++ {
        var v28 int = (fi27 & (fi27 + -2147483648))
        switch (fi27 - -295) {
        case -142:
            v29 := GenS3{F0: v19, F1: (v4 % (((^(v4))) | 1)), F2: uint16(v3)}
            _ = v29
        case 4294967295:
            switch 4294967295 {
            case -323:
                fmt.Println(v3, v28, v2)
            case 32768:
                if (-1 > 26) {
                    var v30 int32 = 42
                    v31 := GenS4{F0: -2147483648}
                    _ = v30
                    _ = v31
                } else {
                    switch 32767 {
                    case -1:
                        fmt.Println(fi27, v3, v2)
                    case 42:
                        var v32 uint16 = 0
                        _ = v32
                    }
                }
            default:
                v4 -= v4
            }
        }
        _ = len(v6)
        _ = fi27
        _ = v28
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
