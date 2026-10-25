package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int16 }
func genFn0() int {
    var v6 int8 = -33
    v7 := []int{422}
    _ = v6
    _ = v7
    return 144
}
func genFn1(p0 int32) uint {
    fmt.Println(p0)
    p0 = 739
    _ = p0
    return 582
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = 913
    var v3 int16 = -32768
    v4 := make(map[int]uint32)
    var v5 uint64 = 310
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 float32 = 1e10
        fmt.Println(fi8)
        v10 := GenS1{F0: 42, F1: (v2 / ((410) | 1)), F2: (v3 ^ 410)}
        _ = fi8
        _ = v10
        _ = v9
    }
    v3 = (v3 ^ (v3 >> ((18446744073709551615) & 63)))
    var v11 uint32 = 2147483648
    if false {
        fmt.Println(v11)
    }
    v12 := []int{2147483647}
    for fi13 := 0; fi13 < 12; fi13++ {
        switch genFn0() {
        case 4294967295:
            delete(v1, (^(v11)))
        default:
            v11 = v11
        }
        if (4294967295 >= (^(v5))) {
            v14 := GenS1{F0: 42, F1: v2, F2: v3}
            _ = v14
        }
        var v15 uint = genFn1(-946)
        _ = fi13
        _ = v15
    }
    v2 += (v2 << ((2) & 63))
    v2 = v2
    for fi16 := 0; fi16 < 18; fi16++ {
        v5 -= v5
        _ = fi16
    }
    if (-37 != 1) {
        var v17 int8 = -56
        _ = v17
    } else {
        v5 = (v5 / ((1) | 1))
    }
    var v18 int64 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
