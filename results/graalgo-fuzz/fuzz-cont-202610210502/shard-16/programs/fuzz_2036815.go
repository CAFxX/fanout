package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint; F1 int16; F2 int32 }
type GenS3 struct { F0 string; F1 string; F2 uint32 }

func main() {
    var v1 float32 = 0.0
    var v2 float64 = ((0.0 * 0.1) - 1.5)
    var v3 uint16 = 886
    if false {
        var v4 int16 = -958
        v4 = (^(v4))
        _ = v4
    }
    v5 := []int{}
    v6 := GenS1{F0: 1, F1: 590}
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 uint8 = 128
        var v9 int16 = 0
        _ = fi7
        _ = v8
        _ = v9
    }
    switch 0 {
    case -79:
        var v10 int16 = 0
        _ = v10
    case -2147483648:
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 int8 = 0
            for fi13 := 0; fi13 < 9; fi13++ {
                fmt.Println(fi13, v3, fi13)
                _ = fi13
            }
            v2 += ((v2 / float64(v1)) + 1e10)
            _ = fi11
            _ = v12
        }
    case -247:
        if len(v5) > 0 {
            v5[0] = -681
        }
    default:
        v1 = ((v1 + 1e10) + v1)
    }
    _ = len(v5)
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 int64 = 553
        _ = fi14
        _ = v15
    }
    if (42 < -350) {
        for fi16 := 0; fi16 < 7; fi16++ {
            fmt.Println(fi16)
            _ = fi16
        }
        if (!(42 != 691)) {
            v1 -= v1
            if len(v5) > 0 {
                v5[0] = 530
            }
            v3 = v3
        } else {
            for fi17 := 0; fi17 < 7; fi17++ {
                if true {
                    v3 += v3
                    v18 := GenS2{F0: uint(v3), F1: -366, F2: 0}
                    _ = v18
                }
                _ = fi17
            }
            _ = len(v5)
            fmt.Println(v3)
        }
    } else {
        var v19 bool = (972 > 282)
        var v20 int64 = 735
        v3 -= 255
        _ = v19
        _ = v20
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if len(v5) > 0 {
        v5[0] = -9223372036854775808
    }
    var v21 string = "a"
    switch -517 {
    case 4294967295:
        var v22 string = v21
        _ = v22
    default:
        fmt.Println(v1)
    }
    switch -1 {
    case 2147483647:
        var v23 uint64 = 1
        _ = v23
    case -1:
        fmt.Println(v1, v3, v2)
    case 9223372036854775807:
        var v24 uint64 = 4294967295
        _ = v24
    }
    v25 := GenS3{F0: v21, F1: v21, F2: 2147483647}
    v1 = (0.0 * v1)
    fmt.Println(v1)
    v26 := GenS2{F0: 2, F1: 0, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v5
    _ = v6
}
