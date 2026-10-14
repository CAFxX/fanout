package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
func genFn0(p0 int, p1 int) int32 {
    for fi4 := 0; fi4 < 8; fi4++ {
        var v5 int32 = 42
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: 65535, F1: -2147483648}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v6
    return -2147483648
}
func genFn1() int64 {
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 uint64 = 18446744073709551615
        _ = fi7
        _ = v8
    }
    var v9 int16 = 1
    v9 += 802
    var v10 float64 = -2.25
    _ = v10
    _ = v9
    return 1
}

func main() {
    var v1 int16 = 32767
    var v2 int32 = 1
    var v3 uint32 = 0
    v3 -= v3
    var v11 uint8 = 128
    if (v11 > 67) {
        v12 := GenS1{F0: 626, F1: 9223372036854775807}
        for fi13 := 0; fi13 < 2; fi13++ {
            if true {
                v14 := GenS1{F0: 0, F1: (^(fi13))}
                _ = v14
            }
            v1 = v1
            fmt.Println(v3)
            _ = fi13
        }
        v2 += v2
        _ = v12
    }
    switch 128 {
    case 4294967295:
        for fi15 := 0; fi15 < 11; fi15++ {
            var v16 uint64 = 42
            var v17 uint16 = 629
            if true {
                v18 := GenS1{F0: 255, F1: 2147483647}
                v19 := GenS1{F0: 65535, F1: (fi15 - (fi15 << ((2) & 63)))}
                v11 -= 42
                _ = v18
                _ = v19
            }
            _ = fi15
            _ = v16
            _ = v17
        }
    case 191:
        v20 := []int{-9223372036854775808, 824}
        _ = v20
    case -1:
        v1 -= v1
    default:
        var v21 bool = (-641 == -520)
        _ = v21
    }
    v2 = -221
    for fi22 := 0; fi22 < 10; fi22++ {
        v23 := []int{(fi22 & fi22), int(v1), (fi22 % (((fi22 * fi22)) | 1))}
        var v24 bool = (false || (!false))
        var v25 uint16 = 1
        _ = fi22
        _ = v23
        _ = v24
        _ = v25
    }
    v1 += 872
    var v26 float32 = ((-2.25 * 0.1) + ((-2.25 / (0.1 / 0.1)) + 3.14159))
    v1 = 695
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v26
    _ = v3
}
