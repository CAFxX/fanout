package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint8 }
type GenS2 struct { F0 bool; F1 int64; F2 uint64 }
func genFn0(p0 int8) string {
    if (-1 >= 1) {
        v7 := []int{}
        v8 := GenS1{F0: false, F1: -2147483648, F2: 223}
        _ = v7
        _ = v8
    }
    v9 := []int{127}
    _ = p0
    _ = v9
    return "foo bar"
}
func genFn1(p0 int64, p1 uint16) int8 {
    switch 32767 {
    case -1:
        var v10 uint8 = 2
        _ = v10
    case 35:
        if ((!false) && true) {
            p0 = p0
        } else {
            var v11 int32 = 1
            _ = v11
        }
    case -333:
        if false {
            fmt.Println(p0, p0)
        }
    default:
        var v12 uint16 = p1
        _ = v12
    }
    fmt.Println(p0, p0)
    p0 += p0
    if false {
        fmt.Println(p1)
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 float64 = (0.1 * 0.1)
        p0 += (p0 / (((p0 ^ 1)) | 1))
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    return 127
}

func main() {
    var v1 int8 = 42
    v2 := make(map[uint32]uint32)
    var v3 int64 = 1
    v4 := make(map[int]bool)
    v5 := make(map[uint32]uint32)
    v6 := make(map[int]uint32)
    var v15 uint = 903
    for fi16 := 0; fi16 < 16; fi16++ {
        if (!((1 == 2) || (!(genFn1(v3, 2) != (v1 * v1))))) {
            v3 -= v3
            fmt.Println(v1)
            v15 = (v15 & v15)
        }
        _ = len(v2)
        _ = fi16
    }
    var v17 float64 = 3.14159
    if (127 <= 128) {
        v17 = v17
    }
    v1 = v1
    if (4294967295 > 42) {
        v18 := GenS2{F0: (-9223372036854775808 != (v3 & (v3 / ((v3) | 1)))), F1: -9223372036854775808, F2: 18446744073709551615}
        _ = v18
    } else {
        var v19 int64 = v3
        _ = v19
    }
    v20 := GenS2{F0: (true && (0.1 <= 0.0)), F1: (v3 & v3), F2: 4294967295}
    v17 = (-2.25 * v17)
    fmt.Println(v1, v15)
    v3 = -200
    var v21 uint8 = 2
    for fi22 := 0; fi22 < 17; fi22++ {
        if (!(1 != 42)) {
            v23 := []int{2147483647, (fi22 << ((v15) & 63)), (fi22 ^ (fi22 << (((^(v15))) & 63)))}
            fmt.Println(v17)
            _ = v23
        }
        _ = fi22
    }
    var v24 uint8 = (v21 ^ (v21 + (v21 | 151)))
    delete(v6, 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
