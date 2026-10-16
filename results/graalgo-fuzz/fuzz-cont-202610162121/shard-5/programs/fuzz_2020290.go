package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint64, p1 uint) int8 {
    var v4 string = "\n\t"
    var v5 float64 = 0.1
    v6 := []int{}
    v4 = (((((v4 + v4) + v4) + v4) + v4) + (v4 + (v4 + "")))
    switch 32767 {
    case 2147483647:
        v7 := GenS1{F0: 0}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return -128
}

func main() {
    var v1 int = -9223372036854775808
    var v2 int = v1
    var v3 int8 = 0
    if (950 >= 4294967295) {
        v8 := GenS1{F0: -99}
        var v9 int32 = 2147483647
        _ = v8
        _ = v9
    }
    switch v2 {
    case 958:
        v10 := []int{}
        _ = v10
    }
    v11 := []int{}
    v12 := GenS2{F0: 991}
    var v13 uint = 42
    if (false && true) {
        _ = len(v11)
        for fi14 := 0; fi14 < 1; fi14++ {
            v15 := GenS2{F0: 2147483648}
            for fi16 := 0; fi16 < 12; fi16++ {
                var v17 uint64 = 80
                v17 -= v17
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                _ = fi16
                _ = v17
            }
            _ = fi14
            _ = v15
        }
        if true {
            v18 := GenS1{F0: -1}
            v19 := GenS3{F0: (v3 | genFn0(4294967295, (v13 / (((^(v13))) | 1)))), F1: 32767, F2: 127}
            v2 = v1
            _ = v18
            _ = v19
        }
    } else {
        v1 = v2
        v2 += 4294967295
        v3 -= (v3 / ((v3) | 1))
    }
    if true {
        fmt.Println(v3, v1, v2)
        v3 -= v3
        v1 += -1
    } else {
        var v20 uint32 = 4294967295
        _ = v20
    }
    v2 = (v1 & v2)
    v13 -= v13
    switch 128 {
    case 42:
        v21 := GenS1{F0: -9223372036854775808}
        _ = v21
    case 2147483647:
        var v22 int16 = 0
        _ = v22
    case 32767:
        v23 := GenS4{F0: false}
        _ = v23
    default:
        switch "" {
        case "\n\t":
            v3 += (v3 ^ genFn0(850, v13))
        case "":
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
        default:
            if true {
                if true {
                    var v24 int = v1
                    v1 += v24
                    v2 = (-(v2))
                    _ = v24
                }
                v3 += v3
                var v25 uint16 = 511
                _ = v25
            }
        }
    }
    v26 := GenS3{F0: genFn0(933, (^(v13))), F1: 432, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v26
    _ = v3
}
