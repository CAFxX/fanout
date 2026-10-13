package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint; F1 bool; F2 int64 }
func genFn0(p0 uint32, p1 int32, p2 uint64) int16 {
    v4 := []int{127}
    var v5 uint64 = p2
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 1000
}
func genFn1(p0 uint32) uint16 {
    var v6 float64 = 3.14159
    v7 := []int{-1, 9223372036854775807}
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 bool = false
        _ = fi8
        _ = v9
    }
    if (!false) {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        v6 += 0.0
        if true {
            if (!true) {
                v6 = v6
                fmt.Println(v6, v6)
                if (!(!(!((65535 > 18446744073709551615) || (false && true))))) {
                    p0 = (p0 % ((p0) | 1))
                } else {
                    if len(v7) > 0 {
                        v7[0] = 2147483647
                    }
                    if true {
                        _ = len(v7)
                        for fi10 := 0; fi10 < 14; fi10++ {
                            for fi11 := 0; fi11 < 2; fi11++ {
                                fmt.Println(fi11, fi11, p0)
                                v6 += (v6 / v6)
                                _ = fi11
                            }
                            v12 := GenS1{F0: -128}
                            _ = fi10
                            _ = v12
                        }
                    }
                    if (625 != -558) {
                        v6 -= 1e10
                        v6 = float64(((0.0 + float32((v6 + v6))) * 3.14159))
                        if false {
                            if (1e10 == ((((-2.25 * (0.0 + (1e10 * 1.5))) * 0.0) + 0.1) / 1.0)) {
                                switch 42 {
                                case -908:
                                    var v13 uint64 = 2
                                    _ = v13
                                case 4294967295:
                                    _ = len(v7)
                                case -961:
                                    p0 = 2147483647
                                }
                                if (!true) {
                                    var v14 uint32 = p0
                                    if len(v7) > 0 {
                                        v7[0] = 42
                                    }
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                    _ = v14
                                } else {
                                    _ = len(v7)
                                    var v15 float64 = v6
                                    _ = v15
                                }
                            }
                        }
                    }
                }
            }
        } else {
            var v16 float64 = 0.1
            v16 = (-2.25 + v16)
            _ = v16
        }
    }
    _ = p0
    _ = v6
    _ = v7
    return 880
}

func main() {
    var v1 int16 = 0
    var v2 uint16 = 2
    var v3 uint64 = 745
    v3 = 2
    v17 := []int{}
    if len(v17) > 0 {
        v17[0] = 127
    }
    fmt.Println(v2, v3, v2)
    var v18 int = 9223372036854775807
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    for fi19 := 0; fi19 < 15; fi19++ {
        var v20 int = -1
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        if false {
            fmt.Println(v3)
            v18 = -104
        }
        var v22 int64 = -9223372036854775808
        var v23 uint64 = (v3 >> ((206) & 63))
        _ = fi21
        _ = v22
        _ = v23
    }
    v17 = append(v17, v18)
    if false {
        if (false && (true || true)) {
            v18 = 290
            v18 = (v18 | 128)
        }
        _ = len(v17)
        fmt.Println(v3)
    }
    switch 697 {
    case 42:
        fmt.Println(v18)
    }
    var v24 uint32 = 2147483648
    v25 := GenS1{F0: 0}
    for fi26 := 0; fi26 < 12; fi26++ {
        v27 := GenS2{F0: (v2 ^ v2), F1: (1 != 798)}
        _ = fi26
        _ = v27
    }
    var v28 string = ("\n\t" + "hello")
    if true {
        switch v18 {
        case 780:
            v29 := GenS1{F0: 18}
            _ = v29
        case 94:
            var v30 int32 = -459
            _ = v30
        case 1:
            var v31 bool = true
            _ = v31
        default:
            v32 := GenS3{F0: 511, F1: false, F2: int64(v3)}
            _ = v32
        }
    }
    var v33 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v33
}
