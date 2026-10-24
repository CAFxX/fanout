package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int8 }
type GenS2 struct { F0 int; F1 uint64; F2 int64 }
type GenS3 struct { F0 string; F1 uint64 }

func main() {
    v1 := []int{32767}
    v2 := make(map[int]string)
    v3 := make(map[int]uint32)
    var v4 float64 = 0.0
    v4 = 1e10
    v4 += v4
    fmt.Println(v4)
    var v5 int8 = -128
    v4 += (v4 * (v4 * v4))
    v4 += -2.25
    delete(v3, 127)
    if (false && false) {
        var v6 int64 = 1
        if true {
            var v7 bool = (1 > 87)
            if len(v1) > 0 {
                v1[0] = 128
            }
            var v8 string = ("a" + "")
            _ = v7
            _ = v8
        }
        var v9 uint64 = 297
        _ = v6
        _ = v9
    } else {
        v10 := GenS1{F0: 42, F1: (624 < 340), F2: v5}
        v11 := GenS2{F0: 2147483647, F1: 458, F2: -68}
        _ = v10
        _ = v11
    }
    var v12 float64 = (v4 + float64(0.1))
    var v13 int8 = v5
    v2[734] = ""
    v14 := GenS2{F0: -9223372036854775808, F1: 2, F2: 597}
    var v15 int8 = v5
    v16 := GenS2{F0: -518, F1: 4294967295, F2: 1}
    v4 -= 0.0
    if (true || false) {
        if false {
            v17 := GenS1{F0: 0, F1: true, F2: 0}
            _ = v17
        } else {
            v18 := GenS3{F0: ("hello" + "a"), F1: 0}
            v15 = (-(v13))
            v3[32768] = 2
            _ = v18
        }
        if (false && (((true || true) || true) || false)) {
            if (!true) {
                var v19 uint = 1
                v4 -= ((((v12 / ((v12 - 1e10) + -2.25)) + ((1.5 / v4) / v4)) / v4) * v12)
                _ = v19
            }
        } else {
            for fi20 := 0; fi20 < 5; fi20++ {
                if (!true) {
                    if len(v1) > 0 {
                        v1[0] = 293
                    }
                    var v21 int = (fi20 ^ (fi20 << ((42) & 63)))
                    fmt.Println(v13, v13)
                    _ = v21
                } else {
                    for fi22 := 0; fi22 < 3; fi22++ {
                        fmt.Println(v15, v13, v4)
                        if (0 < 1) {
                            if (true && (false && false)) {
                                if (false || (9223372036854775807 < -1)) {
                                    v12 -= -2.25
                                    if false {
                                        v12 += float64(((float32(0.0) * 1e10) - ((1e10 + 1e10) + ((float32(float64(0.1)) / -2.25) / 1.0))))
                                        var v23 float32 = ((float32((1e10 + float64((1.5 + float32(0.1))))) * ((((-2.25 + 3.14159) / 0.1) + 1e10) * 1e10)) * (((1.5 + -2.25) + 3.14159) * 0.0))
                                        _ = v23
                                    } else {
                                        switch "hello" {
                                        case "\n\t":
                                            var v24 float64 = (1e10 + ((v12 + v4) * (v4 - 0.1)))
                                            _ = v24
                                        case "hello":
                                            v5 += v15
                                        default:
                                            switch (fi20 * fi20) {
                                            case -205:
                                                var v25 int16 = 42
                                                _ = v25
                                            case 32768:
                                                if (v4 != 1.5) {
                                                    v12 = (((1e10 + v4) + v4) / v4)
                                                }
                                            }
                                        }
                                        v4 = float64(float32((v12 - (v4 / (1e10 / v12)))))
                                    }
                                }
                                var v26 int64 = -1
                                v4 = ((v12 - (-2.25 * (v4 * 1.5))) - (v12 * -2.25))
                                _ = v26
                            } else {
                                var v27 float64 = -2.25
                                v15 = (v15 | v13)
                                fmt.Println(v13, v4)
                                _ = v27
                            }
                        } else {
                            switch "hello" {
                            case "a":
                                v13 -= (v5 << ((2) & 63))
                            case "hello":
                                _ = len(v3)
                            case "\n\t":
                                switch fi22 {
                                case 4294967295:
                                    var v28 uint16 = 505
                                    _ = v28
                                }
                            default:
                                _ = len(v3)
                            }
                            v13 = (v15 % (((v15 % ((v5) | 1))) | 1))
                        }
                        _ = fi22
                    }
                    v5 = (v15 | (v13 & v15))
                }
                _ = fi20
            }
            v29 := GenS1{F0: -168, F1: (true && (42 >= 1)), F2: (^(v5))}
            switch -2147483648 {
            case 1:
                if false {
                    var v30 uint = 0
                    fmt.Println(v30, v5, v15)
                    _ = v30
                }
            default:
                var v31 uint16 = 255
                _ = v31
            }
            _ = v29
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
