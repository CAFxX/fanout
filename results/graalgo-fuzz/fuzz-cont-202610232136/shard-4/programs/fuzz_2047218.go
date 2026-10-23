package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 uint; F1 uint64; F2 int8 }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int8; F1 int8; F2 int8 }
type GenS7 struct { F0 uint64; F1 uint }
type GenS8 struct { F0 int8 }
func genFn0(p0 int64, p1 string) int32 {
    v4 := []int{32768, 127}
    p1 = p1
    p0 -= p0
    switch -2147483648 {
    case 4294967295:
        var v5 int8 = 42
        _ = v5
    case -881:
        for fi6 := 0; fi6 < 8; fi6++ {
            switch 9223372036854775807 {
            case 42:
                for fi7 := 0; fi7 < 9; fi7++ {
                    v8 := GenS1{F0: 0, F1: 1}
                    p1 = p1
                    if (false && true) {
                        v9 := GenS2{F0: 271, F1: -58}
                        _ = v9
                    } else {
                        p1 = p1
                    }
                    _ = fi7
                    _ = v8
                }
            case -9223372036854775808:
                if (0 != 2) {
                    p1 = p1
                }
            }
            _ = fi6
        }
    default:
        v10 := GenS3{F0: 42, F1: 97, F2: 100}
        _ = v10
    }
    if (1 == p0) {
        switch -2147483648 {
        case -441:
            v11 := GenS3{F0: 0, F1: 1, F2: -1}
            _ = v11
        }
        p0 += (p0 << ((578) & 63))
    }
    _ = p0
    _ = p1
    _ = v4
    return 1
}
func genFn1(p0 uint8, p1 uint32, p2 uint) uint16 {
    v12 := []int{128, -595, -1, -503}
    v13 := GenS4{F0: p2, F1: p2}
    p1 -= (p1 % (((^(p1))) | 1))
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = 9223372036854775807
    var v3 int64 = 313
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 int16 = 1000
        var v16 uint16 = 249
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := []int{(v2 & v2), (-(v2)), 674}
    v3 += v3
    var v18 float32 = (-2.25 - float32(float64(3.14159)))
    switch v2 {
    case 127:
        var v19 uint32 = 2
        _ = v19
    }
    for fi20 := 0; fi20 < 11; fi20++ {
        var v21 bool = true
        if v21 {
            v22 := GenS5{F0: 928}
            if v21 {
                if v21 {
                    if (!(v21 || v21)) {
                        var v23 int16 = 0
                        for fi24 := 0; fi24 < 17; fi24++ {
                            if (false && v21) {
                                v3 -= v3
                            }
                            v17 = append(v17, fi20)
                            _ = fi24
                        }
                        _ = v23
                    }
                    v18 += float32(3.14159)
                }
            } else {
                v18 -= v18
                if v21 {
                    v3 -= (v3 - 1)
                    v25 := GenS4{F0: 296, F1: 187}
                    fmt.Println(v3)
                    _ = v25
                }
            }
            v26 := GenS6{F0: -128, F1: 42, F2: -50}
            _ = v22
            _ = v26
        }
        var v27 int = 42
        _ = fi20
        _ = v21
        _ = v27
    }
    switch 9223372036854775807 {
    case 801:
        if false {
            v3 -= v3
            v28 := v1[938]
            _ = v28
        }
    }
    v29 := GenS7{F0: 4294967295, F1: 2}
    var v30 uint16 = 1
    var v31 int32 = 0
    var v32 uint8 = 10
    v33 := GenS8{F0: int8(v3)}
    v3 -= v3
    v18 = 0.0
    var v34 uint = 1
    fmt.Println(v2, v34, v34)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
}
