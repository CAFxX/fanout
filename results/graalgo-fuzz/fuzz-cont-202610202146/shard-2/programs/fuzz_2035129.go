package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 bool }

func main() {
    var v1 int8 = 127
    var v2 uint8 = 127
    var v3 int64 = -1
    v4 := make(map[int]bool)
    var v5 uint64 = 1
    var v6 int64 = -9223372036854775808
    fmt.Println(v5)
    if true {
        if false {
            var v7 uint64 = v5
            for fi8 := 0; fi8 < 3; fi8++ {
                for fi9 := 0; fi9 < 13; fi9++ {
                    v10 := []int{fi9, int(v2), 128, fi9}
                    var v11 int8 = (v1 % ((-113) | 1))
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                _ = fi8
            }
            _ = v7
        } else {
            v12 := GenS1{F0: (v1 ^ v1), F1: 42, F2: v5}
            fmt.Println(v2)
            var v13 bool = true
            _ = v12
            _ = v13
        }
        var v14 int16 = 32767
        _ = v14
    }
    if false {
        v15 := GenS2{F0: (v6 ^ (^(v3)))}
        for fi16 := 0; fi16 < 12; fi16++ {
            v1 = 100
            _ = fi16
        }
        _ = v15
    } else {
        var v17 int32 = 1
        v1 -= (v1 & (v1 >> ((42) & 63)))
        switch 127 {
        case 32767:
            fmt.Println(v1, v6)
        case -842:
            v18 := []int{}
            _ = v18
        case 127:
            v3 -= (v3 & v6)
        default:
            if ("foo bar" < ((("a" + ("a" + "foo bar")) + "a") + "foo bar")) {
                if true {
                    switch 265 {
                    case 835:
                        _ = len(v4)
                    default:
                        v6 = -650
                    }
                    for fi19 := 0; fi19 < 18; fi19++ {
                        var v20 int = (^(fi19))
                        v21 := GenS3{F0: v1, F1: true}
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    v22 := GenS1{F0: (-(v1)), F1: 1000, F2: (v5 << ((255) & 63))}
                    _ = v22
                } else {
                    if (true && ((((1e10 / -2.25) * 1e10) < 0.1) && false)) {
                        v6 = v6
                        v4[int(v2)] = (v6 == v3)
                    }
                }
                var v23 int16 = -205
                var v24 uint64 = (v5 << ((18446744073709551615) & 63))
                _ = v23
                _ = v24
            }
        }
        _ = v17
    }
    v25 := GenS1{F0: -128, F1: -1, F2: 1}
    v1 = v1
    v5 -= v5
    if true {
        v3 = (v6 ^ v6)
        switch "\n\t" {
        case "\n\t":
            v26 := []int{-1, 127}
            _ = v26
        default:
            var v27 uint16 = 230
            _ = v27
        }
    }
    var v28 uint32 = 319
    var v29 int = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
