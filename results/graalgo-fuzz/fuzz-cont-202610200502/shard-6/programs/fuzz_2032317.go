package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32; F1 int16 }
func genFn0(p0 int32, p1 uint, p2 int64) int {
    for fi6 := 0; fi6 < 13; fi6++ {
        p1 = (p1 ^ p1)
        _ = fi6
    }
    var v7 float32 = (((1.5 * float32(0.1)) / ((3.14159 - 1e10) / (-2.25 / 1.0))) + 0.1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    v1 := []uint32{}
    var v2 int8 = 0
    var v3 uint64 = 117
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]int)
    var v8 string = "a"
    v2 = v2
    fmt.Println(v2, v2, v3)
    var v9 bool = (-1 < 908)
    if (-1 != 732) {
        var v10 int = -1
        v2 += 1
        _ = v10
    } else {
        switch v8 {
        case "\n\t":
            var v11 float64 = (3.14159 - 1e10)
            _ = v11
        case "a":
            v4[2147483647] = 587
        default:
            var v12 uint64 = (v3 >> ((42) & 63))
            _ = v12
        }
    }
    v2 -= 42
    var v13 float64 = 0.0
    for fi14 := 0; fi14 < 1; fi14++ {
        v3 = 28
        for fi15 := 0; fi15 < 14; fi15++ {
            v2 += v2
            v9 = v9
            _ = fi15
        }
        v9 = (42 != 0)
        _ = fi14
    }
    if v9 {
        v13 = (v13 + (v13 + v13))
        var v16 int16 = 42
        if v9 {
            if (!(0.1 <= ((-2.25 / float32(v13)) - (1.5 + 1e10)))) {
                v17 := GenS1{F0: 227, F1: 1}
                v8 = (v8 + "foo bar")
                _ = v17
            }
        } else {
            fmt.Println(v9, v13, v16)
        }
        _ = v16
    }
    var v18 float32 = 3.14159
    _ = len(v1)
    v9 = true
    var v19 bool = (v9 || (821 < 0))
    for fi20 := 0; fi20 < 7; fi20++ {
        if v9 {
            v8 = v8
            if false {
                var v21 uint8 = 239
                v13 += (0.1 + v13)
                _ = v21
            } else {
                var v22 uint16 = 255
                v13 += (0.0 / v13)
                if v19 {
                    for fi23 := 0; fi23 < 1; fi23++ {
                        var v24 int = (^(fi20))
                        var v25 float32 = v18
                        if ((!v9) || false) {
                            var v26 float64 = v13
                            var v27 int32 = 633
                            fmt.Println(v2, v13)
                            _ = len(v4)
                            _ = v26
                            _ = v27
                        } else {
                            if (2 >= 0) {
                                v28 := GenS2{F0: 1, F1: 938}
                                _ = v28
                            }
                            v29 := GenS2{F0: 4294967295, F1: 2}
                            v19 = (!(v3 >= (v3 % ((v3) | 1))))
                            _ = v29
                        }
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    switch fi20 {
                    case 42:
                        v30 := GenS3{F0: 0}
                        _ = v30
                    case -9223372036854775808:
                        v31 := GenS4{F0: 1, F1: 18446744073709551615}
                        _ = v31
                    case 1:
                        v32 := GenS5{F0: 1}
                        _ = v32
                    default:
                        fmt.Println(fi20)
                    }
                    _ = len(v1)
                }
                _ = v22
            }
            if ((^(v3)) > (v3 | (v3 << ((65535) & 63)))) {
                v1 = append(v1, 770)
                if ((v2 / ((v2) | 1)) != v2) {
                    var v33 float64 = -2.25
                    v9 = (2 > 132)
                    _ = v33
                } else {
                    var v34 bool = v19
                    _ = v34
                }
            }
        } else {
            fmt.Println(v3)
            v9 = false
            v35 := GenS4{F0: (v3 % (((^(v3))) | 1)), F1: 1}
            _ = v35
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi20
    }
    v36 := GenS3{F0: 42}
    switch "hello" {
    case "":
        var v37 string = v8
        _ = v37
    case "\n\t":
        var v38 int8 = v2
        _ = v38
    }
    switch -164 {
    case 128:
        fmt.Println(v13, v13)
    default:
        v8 = v8
    }
    v39 := GenS6{F0: -1, F1: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v36
    _ = v39
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
