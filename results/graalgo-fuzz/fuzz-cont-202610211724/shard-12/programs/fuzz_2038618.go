package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint16; F1 uint32; F2 int64 }
func genFn0(p0 uint32, p1 int32, p2 int16) int8 {
    var v5 uint16 = 32767
    v6 := GenS1{F0: 82, F1: 32768}
    if (!((2 < 728) || (!true))) {
        p0 += p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -97
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = float64(-2.25)
    v3 := []uint32{1}
    v4 := []uint32{65535, 2, 1, 2147483648}
    v2 = (v2 / float64(-2.25))
    var v7 int64 = 0
    v8 := GenS1{F0: 42, F1: 127}
    switch 578 {
    case -315:
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 float64 = v2
            v11 := GenS1{F0: 128, F1: fi9}
            _ = fi9
            _ = v10
            _ = v11
        }
    case -913:
        var v12 float64 = ((((float64(1.5) / (0.0 + 3.14159)) + 1.5) * v2) - v2)
        _ = v12
    default:
        v13 := v1[-2147483648]
        _ = v13
    }
    var v14 int = 2147483647
    var v15 float32 = -2.25
    if (18446744073709551615 < 2) {
        v14 -= v14
        switch ("" + "hello") {
        case "hello":
            var v16 int64 = 0
            _ = v16
        case "foo bar":
            switch 947 {
            case 32768:
                var v17 int8 = -1
                _ = v17
            case 9223372036854775807:
                v18 := GenS1{F0: 213, F1: v14}
                _ = v18
            default:
                var v19 int8 = -95
                _ = v19
            }
        }
    }
    var v20 uint32 = 2147483648
    if false {
        if (v2 == (float64((((0.1 * float32(-2.25)) + v15) + ((v15 + v15) - (v15 - 0.0)))) * v2)) {
            _ = len(v3)
            var v21 float64 = 0.1
            _ = v21
        } else {
            for fi22 := 0; fi22 < 9; fi22++ {
                for fi23 := 0; fi23 < 15; fi23++ {
                    var v24 int8 = -11
                    var v25 float32 = (3.14159 - float32((1.5 + v2)))
                    if ((42 <= 42) && true) {
                        switch -2147483648 {
                        case -156:
                            var v26 int = v14
                            _ = v26
                        case 0:
                            switch "\n\t" {
                            case "":
                                fmt.Println(v15, v24)
                            case "a":
                                if false {
                                    v1[(fi22 & (fi23 / ((v14) | 1)))] = 233
                                    fmt.Println(v20, fi22)
                                    v27 := GenS1{F0: 128, F1: (fi23 | 32767)}
                                    _ = v27
                                } else {
                                    var v28 float32 = ((((v25 / v25) + v25) - 0.0) + v15)
                                    _ = v28
                                }
                            case "\n\t":
                                var v29 uint = 4294967295
                                _ = v29
                            }
                        default:
                            var v30 int16 = -32768
                            _ = v30
                        }
                        switch -9223372036854775808 {
                        case 32767:
                            v20 = 2147483647
                        case 251:
                            v31 := GenS1{F0: 1, F1: 42}
                            _ = v31
                        default:
                            var v32 int8 = int8(v20)
                            _ = v32
                        }
                        v33 := GenS2{F0: 32767, F1: (v20 - (^(v20))), F2: (v7 & (v7 + (v7 >> ((18446744073709551615) & 63))))}
                        switch v14 {
                        case -9223372036854775808:
                            fmt.Println(v2, v2)
                        default:
                            v15 += 3.14159
                        }
                        switch ("a" + "foo bar") {
                        case "":
                            var v34 int = (fi23 % ((v14) | 1))
                            _ = v34
                        case "\n\t":
                            var v35 bool = (0 == 231)
                            _ = v35
                        case "foo bar":
                            var v36 uint16 = 0
                            _ = v36
                        }
                        _ = v33
                    } else {
                        v14 -= v14
                    }
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                _ = len(v1)
                if false {
                    var v37 uint8 = 2
                    fmt.Println(v14, v7, fi22)
                    _ = v37
                }
                _ = fi22
            }
            v20 += 0
        }
    }
    v1[-9223372036854775808] = v20
    fmt.Println(v14)
    fmt.Println(v14)
    v20 = (v20 ^ (v20 ^ (v20 - (v20 + v20))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
