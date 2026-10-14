package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint64; F1 int8; F2 int16 }
type GenS4 struct { F0 string; F1 int32; F2 int16 }
func genFn0(p0 uint8, p1 int32) bool {
    var v6 uint = 560
    var v7 bool = ((29 < 42) && (2147483647 <= 3))
    v7 = ((2 <= 0) && (9223372036854775807 > 1))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return ((p0 | (^(p0))) == (^(p0)))
}

func main() {
    v1 := []uint32{65535}
    var v2 uint16 = 270
    var v3 int16 = -691
    var v4 float64 = 3.14159
    var v5 uint16 = (^(v2))
    v1 = append(v1, 631)
    v8 := GenS1{F0: 322}
    var v9 uint8 = 0
    if false {
        var v10 float32 = float32(float64(0.1))
        _ = v10
    }
    v11 := GenS2{F0: 2, F1: -1}
    v12 := GenS3{F0: 920, F1: -1, F2: (v3 << ((42) & 63))}
    var v13 float64 = float64(float32(((v4 - ((v4 * ((float64(((3.14159 / -2.25) * (((-2.25 - 3.14159) * -2.25) - float32((1e10 / ((float64(0.1) + 0.0) + v4)))))) - float64((0.1 * 1e10))) * v4)) + 0.0)) / v4)))
    var v14 uint8 = (^(v9))
    switch 127 {
    case 42:
        v13 += ((((((v13 / (v13 * float64(float32(v4)))) * v13) / v4) * v13) * 3.14159) - (v4 / 0.1))
    default:
        v9 = v14
    }
    v9 = v9
    fmt.Println(v13)
    for fi15 := 0; fi15 < 8; fi15++ {
        if (!(110 != 42)) {
            var v16 string = "a"
            if (!(false && (!true))) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                if true {
                    v4 -= (0.1 / ((((3.14159 - float64(0.0)) / (v4 - ((v4 / v4) / v13))) + (v4 / (v13 * (v4 + (-2.25 * (v4 / v4)))))) + v4))
                    if false {
                        fmt.Println(v2, v9, v9)
                        var v17 uint = 4294967295
                        fmt.Println(v13)
                        _ = v17
                    } else {
                        if false {
                            if false {
                                switch "foo bar" {
                                case "a":
                                    fmt.Println(v16)
                                case "hello":
                                    v3 = (v3 << ((2) & 63))
                                case "\n\t":
                                    v18 := GenS4{F0: v16, F1: 42, F2: v3}
                                    _ = v18
                                default:
                                    var v19 uint64 = 29
                                    _ = v19
                                }
                                var v20 float64 = v4
                                _ = v20
                            }
                        }
                        var v21 float64 = v13
                        _ = v21
                    }
                }
            }
            _ = v16
        }
        _ = fi15
    }
    v22 := GenS1{F0: 255}
    for fi23 := 0; fi23 < 8; fi23++ {
        _ = len(v1)
        _ = fi23
    }
    switch "\n\t" {
    case "\n\t":
        var v24 int32 = -1
        _ = v24
    case "":
        switch 2147483647 {
        case 127:
            var v25 int = -257
            _ = v25
        case 0:
            var v26 int8 = int8(v14)
            _ = v26
        }
    }
    switch "foo bar" {
    case "\n\t":
        if true {
            var v27 bool = ((3.14159 - 1e10) > 1e10)
            switch -139 {
            case -9223372036854775808:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = v27
        }
    default:
        var v28 int64 = 42
        _ = v28
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
