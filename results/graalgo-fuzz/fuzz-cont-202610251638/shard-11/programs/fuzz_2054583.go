package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 uint16; F1 uint8 }
func genFn0(p0 uint64, p1 uint16, p2 int) uint16 {
    if false {
        v7 := GenS1{F0: 65535, F1: 753}
        var v8 int16 = 32767
        var v9 float32 = (0.1 * (3.14159 * 0.1))
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 635, F1: -9223372036854775808}
    v11 := []int{p2, int(p0), -226}
    if false {
        p1 -= 684
        var v12 float64 = (-2.25 * -2.25)
        fmt.Println(p2, v12, p0)
        _ = v12
    } else {
        if true {
            v13 := GenS1{F0: 2, F1: 42}
            if len(v11) > 0 {
                v11[0] = -716
            }
            _ = v13
        }
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return p1
}

func main() {
    var v1 float64 = -2.25
    v2 := []string{"hello", ("\n\t" + ("a" + (("\n\t" + "\n\t") + ("hello" + "hello")))), ("hello" + ("a" + ""))}
    var v3 uint32 = 2147483647
    var v4 int64 = -9223372036854775808
    var v5 int16 = -899
    var v6 uint32 = (v3 & (^(v3)))
    var v14 int = 9223372036854775807
    v6 = 859
    if (v14 < 32) {
        v5 = v5
    }
    switch ("hello" + "a") {
    case "a":
        v15 := GenS2{F0: 429, F1: 930, F2: (v3 & 926)}
        _ = v15
    case "hello":
        if (true || (42 <= 697)) {
            if true {
                _ = len(v2)
                v16 := GenS1{F0: 854, F1: v4}
                _ = v16
            } else {
                v5 = 436
            }
            for fi17 := 0; fi17 < 19; fi17++ {
                switch "\n\t" {
                case "a":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case "":
                    _ = len(v2)
                default:
                    if ((!(!((1 >= 0) || false))) || true) {
                        if (0 >= 2147483647) {
                            v1 = float64(-2.25)
                            switch v14 {
                            case 127:
                                v18 := GenS1{F0: 18446744073709551615, F1: 0}
                                _ = v18
                            case 32767:
                                v19 := GenS3{F0: 42, F1: (236 < 1)}
                                _ = v19
                            default:
                                var v20 bool = (v6 < v3)
                                _ = v20
                            }
                        }
                    } else {
                        v21 := GenS1{F0: 894, F1: (v4 * v4)}
                        if true {
                            if false {
                                fmt.Println(fi17, v1)
                                var v22 string = "hello"
                                if false {
                                    fmt.Println(v1)
                                    for fi23 := 0; fi23 < 9; fi23++ {
                                        var v24 int64 = -1
                                        if len(v2) > 0 {
                                            v2[0] = ("a" + v22)
                                        }
                                        _ = fi23
                                        _ = v24
                                    }
                                    v6 -= (v6 | 65535)
                                }
                                _ = v22
                            }
                            v25 := GenS2{F0: 42, F1: genFn0(1, genFn0(515, 552, v14), 645), F2: 4294967295}
                            if (!false) {
                                for fi26 := 0; fi26 < 8; fi26++ {
                                    if (-1 != 2147483647) {
                                        var v27 int = 553
                                        v28 := GenS4{F0: 1, F1: 255}
                                        _ = v27
                                        _ = v28
                                    }
                                    v3 = v6
                                    v14 = fi17
                                    _ = fi26
                                }
                                fmt.Println(v3)
                            }
                            _ = v25
                        }
                        _ = v21
                    }
                }
                v1 = float64(1.5)
                var v29 int8 = -31
                _ = fi17
                _ = v29
            }
            fmt.Println(v1, v4)
        }
    default:
        v5 = v5
    }
    v2 = append(v2, "hello")
    switch (-(v14)) {
    case 4294967295:
        var v30 int8 = 1
        _ = v30
    default:
        var v31 int64 = v4
        _ = v31
    }
    fmt.Println(v4, v14)
    var v32 float64 = (v1 * (v1 - 1.5))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
