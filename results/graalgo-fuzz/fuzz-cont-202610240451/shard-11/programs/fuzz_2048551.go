package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 int64; F2 uint16 }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 int64) bool {
    var v7 float32 = 0.0
    v8 := []int{-976, 127, 845}
    v7 = 0.1
    _ = p0
    _ = v7
    _ = v8
    return false
}

func main() {
    var v1 int8 = 114
    var v2 float64 = -2.25
    var v3 int64 = -9223372036854775808
    var v4 uint16 = 65535
    v5 := []int{4294967295, 42, 32767, 127}
    var v6 uint16 = v4
    if len(v5) > 0 {
        v5[0] = 4294967295
    }
    v9 := GenS1{F0: 65535, F1: 42}
    v2 += v2
    switch "a" {
    case "foo bar":
        var v10 string = ("hello" + "")
        _ = v10
    default:
        var v11 int8 = (v1 & (v1 + v1))
        _ = v11
    }
    var v12 float32 = (0.0 * (0.1 * ((0.0 * 0.1) / 1.0)))
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 uint16 = (v4 - (v4 ^ (v6 - v6)))
        _ = fi13
        _ = v14
    }
    switch -256 {
    case 128:
        var v15 float64 = (3.14159 * v2)
        _ = v15
    case 9223372036854775807:
        var v16 int32 = 728
        _ = v16
    case 426:
        v17 := GenS2{F0: 1}
        _ = v17
    default:
        fmt.Println(v3, v12)
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        v12 -= 3.14159
        _ = fi18
    }
    _ = len(v5)
    if true {
        switch -201 {
        case 1:
            switch 32767 {
            case 1:
                for fi19 := 0; fi19 < 9; fi19++ {
                    fmt.Println(v4, v4)
                    _ = fi19
                }
            default:
                var v20 uint64 = 42
                _ = v20
            }
        case 32767:
            fmt.Println(v2, v1)
        case -1:
            switch "foo bar" {
            case "hello":
                var v21 uint64 = 42
                _ = v21
            case "foo bar":
                fmt.Println(v12, v6, v3)
            default:
                var v22 int16 = -599
                _ = v22
            }
        default:
            if (true && (v6 <= 0)) {
                v4 = v6
            }
        }
        for fi23 := 0; fi23 < 12; fi23++ {
            for fi24 := 0; fi24 < 4; fi24++ {
                if true {
                    if true {
                        v25 := GenS3{F0: 2147483647, F1: v3, F2: v6}
                        var v26 uint16 = v4
                        _ = v25
                        _ = v26
                    }
                    var v27 int32 = 578
                    if (true || ((v6 * v4) >= (v6 + (^(v6))))) {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        var v28 int16 = 1000
                        v2 -= ((v2 - ((float64(float32(v2)) - v2) + v2)) - 1.5)
                        v28 += 439
                        _ = v28
                    }
                    _ = v27
                } else {
                    v29 := GenS3{F0: 42, F1: (v3 * (-(v3))), F2: (v6 % ((v4) | 1))}
                    _ = v29
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v5 = append(v5, fi24)
                _ = fi24
            }
            v30 := GenS4{F0: v1, F1: (v1 / ((v1) | 1))}
            if (-948 == 1) {
                fmt.Println(v3, fi23, v1)
                v31 := GenS1{F0: 0, F1: 2147483647}
                fmt.Println(v3, v12)
                _ = v31
            }
            _ = fi23
            _ = v30
        }
    }
    fmt.Println(v3, v4, v4)
    for fi32 := 0; fi32 < 6; fi32++ {
        fmt.Println(v4, v2, v12)
        _ = fi32
    }
    var v33 uint8 = 150
    if false {
        if true {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            switch 787 {
            case 32767:
                if (true || (2 != 490)) {
                    var v34 float32 = 1e10
                    _ = v34
                }
            case 42:
                v35 := GenS5{F0: 0}
                _ = v35
            case -20:
                _ = len(v5)
            }
        }
    }
    v6 = 636
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v12 += (float32(((((v2 * v2) + v2) * float64(v12)) + 3.14159)) - float32(v2))
    for fi36 := 0; fi36 < 13; fi36++ {
        var v37 int16 = -275
        fmt.Println(v37, v6, v12)
        var v38 int32 = 1
        _ = fi36
        _ = v37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
