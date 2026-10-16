package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 bool; F2 uint32 }
func genFn0(p0 uint32, p1 int) int8 {
    if true {
        p1 = p1
    }
    p1 -= p1
    v4 := []int{}
    var v5 int16 = 0
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -67
}
func genFn1(p0 uint, p1 uint8) int32 {
    var v6 bool = true
    var v7 int8 = 127
    if (344 >= -228) {
        if v6 {
            p0 += (p0 / (((^(p0))) | 1))
            v8 := GenS1{F0: 1, F1: -1, F2: 18446744073709551615}
            _ = v8
        } else {
            v9 := GenS2{F0: -112}
            _ = v9
        }
        v7 += -109
        p1 = 255
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -2147483648
}

func main() {
    var v1 int8 = 1
    var v2 uint32 = 666
    v3 := make(map[uint32]string)
    v10 := GenS3{F0: v1, F1: ((2147483647 <= -2147483648) && true), F2: (^(v2))}
    var v11 int = 352
    switch v11 {
    case 909:
        var v12 int = 477
        _ = v12
    case 270:
        for fi13 := 0; fi13 < 18; fi13++ {
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 int8 = genFn0((v2 + 2147483647), fi14)
                _ = fi14
                _ = v15
            }
            v16 := v3[v2]
            v17 := v3[2147483648]
            _ = fi13
            _ = v16
            _ = v17
        }
    case 127:
        v18 := GenS2{F0: -215}
        _ = v18
    default:
        var v19 int8 = (v1 * 70)
        _ = v19
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        var v21 uint8 = 127
        var v22 int8 = 100
        _ = fi20
        _ = v21
        _ = v22
    }
    if (0 != -83) {
        var v23 uint8 = 255
        _ = v23
    }
    v2 = v2
    v2 -= 2147483647
    if false {
        v1 -= 55
        switch "a" {
        case "\n\t":
            for fi24 := 0; fi24 < 10; fi24++ {
                if ((1 == 83) && true) {
                    for fi25 := 0; fi25 < 1; fi25++ {
                        fmt.Println(v1, fi24)
                        _ = fi25
                    }
                    var v26 float64 = float64(((3.14159 / (-2.25 / 1.0)) / 3.14159))
                    _ = v26
                }
                v2 = v2
                v27 := GenS3{F0: 0, F1: (((1e10 + 1e10) + 0.0) == (1.5 / -2.25)), F2: v2}
                _ = fi24
                _ = v27
            }
        }
        for fi28 := 0; fi28 < 7; fi28++ {
            var v29 string = "\n\t"
            _ = fi28
            _ = v29
        }
    } else {
        fmt.Println(v2, v2, v11)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
}
