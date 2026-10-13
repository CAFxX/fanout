package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint; F1 bool; F2 int16 }
type GenS4 struct { F0 uint32; F1 int64; F2 int32 }
func genFn0(p0 int) int16 {
    fmt.Println(p0)
    if true {
        v6 := []int{}
        p0 = 2147483647
        _ = v6
    } else {
        if false {
            switch p0 {
            case 4294967295:
                if (128 <= 91) {
                    v7 := []int{(p0 + p0), (p0 - 9223372036854775807)}
                    _ = v7
                }
            case -499:
                switch "\n\t" {
                case "foo bar":
                    for fi8 := 0; fi8 < 8; fi8++ {
                        if (-761 <= 255) {
                            var v9 uint64 = 18446744073709551615
                            var v10 float32 = (-2.25 + 3.14159)
                            v11 := GenS1{F0: (v9 * (v9 % (((v9 / ((v9) | 1))) | 1)))}
                            _ = v10
                            _ = v11
                            _ = v9
                        } else {
                            p0 += (p0 & fi8)
                            var v12 uint = 0
                            _ = v12
                        }
                        p0 += (fi8 ^ 1)
                        _ = fi8
                    }
                case "a":
                    p0 = (p0 | p0)
                default:
                    switch "foo bar" {
                    case "\n\t":
                        var v13 int8 = 115
                        _ = v13
                    }
                }
            case -9223372036854775808:
                var v14 int64 = 1
                _ = v14
            default:
                if false {
                    v15 := []int{p0, -905, p0, (p0 - (p0 * 4294967295))}
                    _ = v15
                } else {
                    var v16 int8 = -128
                    v16 -= int8(p0)
                    p0 = p0
                    _ = v16
                }
            }
            v17 := []int{669}
            _ = v17
        }
        if true {
            switch (p0 * 9223372036854775807) {
            case 1:
                var v18 int16 = 42
                _ = v18
            default:
                p0 -= (p0 & (p0 + (^(p0))))
            }
            fmt.Println(p0)
        }
    }
    v19 := GenS1{F0: 1}
    _ = p0
    _ = v19
    return 42
}

func main() {
    v1 := []uint32{421, 2147483647, 0, 4294967295}
    var v2 uint64 = 2
    var v3 float64 = 1e10
    var v4 int = 0
    var v5 uint = 481
    v4 = (^(v4))
    var v20 float64 = v3
    if ((0.1 > 0.1) && ((65535 <= 519) && ("" <= "\n\t"))) {
        for fi21 := 0; fi21 < 9; fi21++ {
            for fi22 := 0; fi22 < 9; fi22++ {
                var v23 uint64 = v2
                switch -66 {
                case 42:
                    v24 := GenS1{F0: 950}
                    _ = v24
                }
                v25 := GenS2{F0: 32767, F1: (v4 * fi22)}
                _ = fi22
                _ = v23
                _ = v25
            }
            if len(v1) > 0 {
                v1[0] = 65535
            }
            _ = fi21
        }
        var v26 string = ""
        _ = v26
    } else {
        for fi27 := 0; fi27 < 20; fi27++ {
            v2 = (v2 / ((v2) | 1))
            _ = fi27
        }
        switch (v4 - 4294967295) {
        case 2147483647:
            v1 = append(v1, 42)
        default:
            for fi28 := 0; fi28 < 5; fi28++ {
                _ = len(v1)
                _ = len(v1)
                _ = len(v1)
                _ = fi28
            }
        }
    }
    if false {
        v29 := GenS1{F0: (v2 >> ((v5) & 63))}
        v5 -= v5
        v30 := GenS3{F0: v5, F1: true, F2: -32768}
        _ = v29
        _ = v30
    } else {
        var v31 uint8 = 2
        var v32 uint64 = (v2 + v2)
        v20 = (float64((((-2.25 * -2.25) * ((3.14159 * (0.1 * 1.5)) - 1e10)) - (float32((float64(float32(0.0)) / (float64(3.14159) + v3))) * 1e10))) * v3)
        _ = v31
        _ = v32
    }
    fmt.Println(v20, v2, v20)
    v3 += (v3 * float64(1.5))
    var v33 int8 = -128
    if len(v1) > 0 {
        v1[0] = 0
    }
    v1 = append(v1, 4294967295)
    v33 += (-(v33))
    var v34 int8 = v33
    v35 := GenS4{F0: 876, F1: 0, F2: 1}
    v1 = append(v1, uint32(v2))
    for fi36 := 0; fi36 < 6; fi36++ {
        _ = len(v1)
        v37 := GenS1{F0: v2}
        v34 = 40
        _ = fi36
        _ = v37
    }
    var v38 int32 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v38
    _ = v4
    _ = v5
}
