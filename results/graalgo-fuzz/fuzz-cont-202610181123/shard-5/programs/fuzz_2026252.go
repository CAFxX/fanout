package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 bool; F2 int32 }
type GenS3 struct { F0 uint8; F1 uint8 }
func genFn0(p0 int16, p1 int16, p2 int32) uint8 {
    if (true || false) {
        var v5 uint16 = 2
        _ = v5
    }
    p2 += 42
    p1 -= p1
    fmt.Println(p2, p2)
    if (1 == 65) {
        v6 := GenS1{F0: 2}
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return 141
}

func main() {
    v1 := []string{"a"}
    var v2 int32 = -2147483648
    var v3 int8 = 100
    var v4 uint8 = 206
    switch -9223372036854775808 {
    case 127:
        v7 := GenS1{F0: 18446744073709551615}
        _ = v7
    default:
        if (42 >= 65535) {
            var v8 float32 = (1e10 / (0.0 + 1.5))
            for fi9 := 0; fi9 < 1; fi9++ {
                v8 -= (v8 * v8)
                if (1.5 >= 0.1) {
                    v10 := GenS2{F0: v2, F1: true, F2: v2}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v10
                } else {
                    var v11 float64 = 0.0
                    v2 += v2
                    _ = v11
                }
                _ = fi9
            }
            _ = v8
        }
    }
    v4 = (v4 - (v4 & (v4 % ((127) | 1))))
    var v12 uint16 = 0
    v1 = append(v1, "foo bar")
    v12 += v12
    v13 := GenS2{F0: int32(v4), F1: true, F2: v2}
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 float32 = -2.25
        _ = fi14
        _ = v15
    }
    v2 = (v2 + int32(v4))
    var v16 float32 = -2.25
    if true {
        var v17 uint64 = 2
        _ = v17
    } else {
        var v18 int64 = 1
        for fi19 := 0; fi19 < 3; fi19++ {
            fmt.Println(v4, v3, v4)
            switch fi19 {
            case 32767:
                v20 := GenS1{F0: 619}
                _ = v20
            case 32768:
                var v21 uint64 = 2
                _ = v21
            default:
                if false {
                    var v22 uint = 231
                    for fi23 := 0; fi23 < 9; fi23++ {
                        v18 -= v18
                        v24 := GenS2{F0: -1, F1: (1.5 != 3.14159), F2: (v2 | int32(v4))}
                        if len(v1) > 0 {
                            v1[0] = "hello"
                        }
                        var v25 float64 = -2.25
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    _ = v22
                } else {
                    v26 := GenS3{F0: (^(v4)), F1: v4}
                    _ = v26
                }
            }
            _ = fi19
        }
        _ = v18
    }
    fmt.Println(v2, v12, v4)
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    var v27 int32 = (v2 * (v2 & (v2 ^ v2)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
