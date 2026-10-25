package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint; F1 uint64; F2 uint64 }
func genFn0(p0 uint16, p1 bool, p2 int32) int32 {
    for fi4 := 0; fi4 < 17; fi4++ {
        fmt.Println(p0, p1)
        v5 := []int{140, fi4}
        p0 += p0
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := []int{(fi6 - (^(fi6)))}
        for fi8 := 0; fi8 < 14; fi8++ {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            v9 := GenS1{F0: 683, F1: 2, F2: 2}
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    var v10 uint8 = 34
    switch 88 {
    case 1:
        var v11 int8 = -128
        _ = v11
    case -2147483648:
        var v12 uint16 = p0
        _ = v12
    default:
        for fi13 := 0; fi13 < 16; fi13++ {
            fmt.Println(p2, p2)
            _ = fi13
        }
    }
    fmt.Println(p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return (p2 * 42)
}
func genFn1(p0 int8) int64 {
    if true {
        switch -272 {
        case 128:
            if false {
                for fi14 := 0; fi14 < 12; fi14++ {
                    v15 := []int{32768, fi14}
                    fmt.Println(fi14)
                    _ = fi14
                    _ = v15
                }
                var v16 uint16 = 255
                v17 := GenS1{F0: 0, F1: 1, F2: 2}
                _ = v16
                _ = v17
            } else {
                if true {
                    p0 -= p0
                }
            }
        case 2147483647:
            if (false || (0 <= 1)) {
                v18 := []int{-739, 2147483647, -9223372036854775808}
                _ = v18
            }
        }
        p0 = (p0 | 42)
        var v19 float64 = (3.14159 * 3.14159)
        _ = v19
    }
    v20 := []int{-9223372036854775808, 2147483647, 9223372036854775807}
    p0 = (^(p0))
    p0 += (p0 << ((130) & 63))
    _ = p0
    _ = v20
    return 9223372036854775807
}

func main() {
    var v1 int8 = -128
    var v2 float32 = 1e10
    var v3 uint8 = 38
    var v21 uint8 = 42
    for fi22 := 0; fi22 < 20; fi22++ {
        var v23 uint16 = 2
        _ = fi22
        _ = v23
    }
    var v24 int16 = 0
    v25 := []int{32768, 4294967295}
    v25 = append(v25, 9223372036854775807)
    v26 := GenS1{F0: 2, F1: 965, F2: 115}
    v27 := GenS2{F0: 18446744073709551615, F1: 18446744073709551615, F2: 313}
    fmt.Println(v1, v2)
    var v28 uint64 = 0
    fmt.Println(v28)
    fmt.Println(v21, v24, v2)
    v25 = append(v25, 32767)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
