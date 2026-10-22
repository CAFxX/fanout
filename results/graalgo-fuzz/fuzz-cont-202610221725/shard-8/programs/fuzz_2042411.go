package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 bool; F2 string }
func genFn0(p0 uint, p1 bool) int {
    fmt.Println(p1, p1)
    v5 := []int{}
    if p1 {
        var v6 int16 = 1
        v7 := GenS1{F0: 2, F1: 128, F2: 274}
        switch "foo bar" {
        case "foo bar":
            var v8 int = -9223372036854775808
            _ = v8
        default:
            for fi9 := 0; fi9 < 13; fi9++ {
                fmt.Println(p0, v6)
                if (-2.25 >= float64((float32((0.0 * float64(0.0))) + 3.14159))) {
                    for fi10 := 0; fi10 < 1; fi10++ {
                        if (895 < 0) {
                            v5 = append(v5, fi10)
                            switch "foo bar" {
                            case "\n\t":
                                v5 = append(v5, -2147483648)
                            }
                            if len(v5) > 0 {
                                v5[0] = 9223372036854775807
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = len(v5)
                        } else {
                            fmt.Println(p0, p0, v6)
                        }
                        _ = fi10
                    }
                }
                var v11 uint64 = 2
                _ = fi9
                _ = v11
            }
        }
        _ = v6
        _ = v7
    } else {
        v12 := GenS1{F0: 0, F1: 2, F2: 0}
        _ = v12
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = v5
    return -9223372036854775808
}
func genFn1(p0 uint, p1 string) uint {
    switch p1 {
    case "a":
        for fi13 := 0; fi13 < 12; fi13++ {
            v14 := []int{fi13}
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            _ = fi13
            _ = v14
        }
    case "":
        p0 -= (p0 / ((18446744073709551615) | 1))
    case "hello":
        v15 := GenS1{F0: 0, F1: 128, F2: (p0 & p0)}
        _ = v15
    default:
        var v16 int16 = 42
        _ = v16
    }
    if (1 < -939) {
        if false {
            v17 := GenS1{F0: 4294967295, F1: 197, F2: p0}
            _ = v17
        }
    } else {
        v18 := []int{}
        _ = v18
    }
    var v19 uint = (p0 >> ((65535) & 63))
    v20 := []int{774}
    v20 = append(v20, 128)
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    return 4294967295
}

func main() {
    var v1 int = 4294967295
    var v2 uint16 = 255
    var v3 int = (-(v1))
    v4 := []uint32{4294967295, 2147483648, 42, 65535}
    if (0 >= 1) {
        var v21 int64 = -1
        if (32767 < 1000) {
            v22 := GenS1{F0: 724, F1: 2, F2: uint(v2)}
            _ = v22
        } else {
            v1 = 32767
        }
        v1 = v1
        _ = v21
    } else {
        v23 := GenS1{F0: 663, F1: 1, F2: 2}
        _ = v23
    }
    var v24 int8 = -89
    switch genFn0(1, (!(false && (false && ((4294967295 > 2) && (2147483647 >= 852)))))) {
    case 9223372036854775807:
        v3 += (v3 - (v3 - v3))
    }
    v25 := GenS1{F0: uint32(v2), F1: 128, F2: 42}
    v24 = v24
    v26 := GenS2{F0: int64(v3)}
    var v27 bool = true
    var v28 float32 = 1.5
    if v27 {
        v29 := GenS3{F0: false, F1: v27, F2: ("" + "hello")}
        v4 = append(v4, 0)
        _ = v29
    } else {
        v24 -= (v24 >> ((genFn1(genFn1(18446744073709551615, ("foo bar" + "")), "\n\t")) & 63))
    }
    v1 -= genFn0(genFn1(18446744073709551615, "foo bar"), v27)
    v30 := GenS2{F0: 1}
    v31 := GenS2{F0: -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
