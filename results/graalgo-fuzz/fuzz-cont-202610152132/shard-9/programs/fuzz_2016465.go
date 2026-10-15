package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int64; F1 uint; F2 bool }
func genFn0(p0 int64, p1 uint16) uint8 {
    v7 := GenS1{F0: 128, F1: 826}
    for fi8 := 0; fi8 < 8; fi8++ {
        p0 -= p0
        for fi9 := 0; fi9 < 8; fi9++ {
            if true {
                v10 := []int{fi8, fi9}
                _ = v10
            }
            _ = fi9
        }
        _ = fi8
    }
    v11 := []int{-723, -67}
    switch 4294967295 {
    case 32767:
        switch -206 {
        case 946:
            var v12 int = 32768
            _ = v12
        case -1:
            p1 = (p1 >> ((18446744073709551615) & 63))
        case 1:
            _ = len(v11)
        default:
            p1 = (p1 - (p1 >> ((42) & 63)))
        }
    case 42:
        _ = len(v11)
    default:
        for fi13 := 0; fi13 < 14; fi13++ {
            p1 = p1
            fmt.Println(fi13)
            _ = fi13
        }
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    return 128
}

func main() {
    var v1 string = "foo bar"
    var v2 float64 = 0.1
    var v3 uint32 = 879
    var v4 int = 4294967295
    var v5 int32 = -2147483648
    var v6 float64 = float64(0.1)
    var v14 float64 = v6
    v15 := GenS2{F0: 4294967295, F1: 1}
    if (!(true && (-9223372036854775808 < v4))) {
        var v16 bool = ((true && false) || ((1000 > int16(v4)) || false))
        if (v16 && (1.5 != -2.25)) {
            fmt.Println(v4)
        }
        switch -31 {
        case 128:
            if false {
                for fi17 := 0; fi17 < 7; fi17++ {
                    switch (v4 * (fi17 >> ((65535) & 63))) {
                    case -1:
                        v18 := GenS3{F0: -1, F1: 4294967295, F2: v16}
                        _ = v18
                    case 128:
                        if (155 >= 42) {
                            var v19 uint64 = 18446744073709551615
                            var v20 float64 = 0.0
                            var v21 float32 = (-2.25 / (1e10 + (((float32(-2.25) + (1e10 * float32(v14))) - 0.1) - 3.14159)))
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                    case 1:
                        var v22 uint = 201
                        _ = v22
                    }
                    v23 := GenS2{F0: 451, F1: 113}
                    _ = fi17
                    _ = v23
                }
                for fi24 := 0; fi24 < 18; fi24++ {
                    v25 := GenS3{F0: -1, F1: 535, F2: (false || v16)}
                    if (v16 && ((((true && (42 <= 42)) || false) || v16) || v16)) {
                        if false {
                            for fi26 := 0; fi26 < 19; fi26++ {
                                v27 := []int{(fi26 * v4), (fi24 - fi26), (v4 / ((v4) | 1)), fi26}
                                var v28 uint16 = 0
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                        } else {
                            fmt.Println(v3, v5, v16)
                            if v16 {
                                var v29 int8 = 1
                                _ = v29
                            }
                        }
                    }
                    fmt.Println(v2, v5)
                    _ = fi24
                    _ = v25
                }
                v1 = v1
            } else {
                v16 = v16
                v30 := GenS1{F0: v4, F1: v3}
                v2 -= ((float64(0.0) * v6) / v2)
                _ = v30
            }
        default:
            v31 := []int{v4, -2147483648, (v4 | v4), (v4 * (-(v4)))}
            _ = v31
        }
        _ = v16
    }
    var v32 uint16 = 376
    v1 = "a"
    v5 = (v5 * 61)
    fmt.Println(v5, v1, v14)
    v2 += (v6 * v2)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
