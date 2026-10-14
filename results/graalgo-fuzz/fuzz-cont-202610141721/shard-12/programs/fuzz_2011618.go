package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint }
func genFn0(p0 int32) uint64 {
    if true {
        if (-2.25 > -2.25) {
            for fi7 := 0; fi7 < 20; fi7++ {
                v8 := GenS1{F0: (-(fi7)), F1: -128}
                var v9 float64 = 1e10
                _ = fi7
                _ = v8
                _ = v9
            }
        }
        var v10 float32 = float32(1e10)
        var v11 uint32 = 992
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: -2147483648, F1: 87}
    p0 -= (^(p0))
    for fi13 := 0; fi13 < 6; fi13++ {
        var v14 uint8 = 127
        v14 = v14
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = v12
    return 0
}
func genFn1(p0 uint16) int64 {
    p0 += 2
    var v15 int16 = 1000
    _ = p0
    _ = v15
    return -9223372036854775808
}

func main() {
    var v1 float32 = 0.1
    v2 := make(map[int]int)
    var v3 uint32 = 285
    var v4 int8 = 0
    var v5 float32 = v1
    var v6 float32 = v1
    switch ("" + ("hello" + ("\n\t" + "foo bar"))) {
    case "foo bar":
        v3 -= 2147483648
    }
    v16 := []int{464}
    v17 := GenS2{F0: 97, F1: 456, F2: 2}
    v6 = float32((1.5 + (-2.25 - -2.25)))
    v6 = ((3.14159 / (1e10 + (0.0 / v6))) / (0.1 / (v5 * v6)))
    var v18 uint16 = 42
    delete(v2, -9223372036854775808)
    var v19 int64 = 0
    v18 -= v18
    if false {
        fmt.Println(v5, v4)
        v3 += v3
        for fi20 := 0; fi20 < 20; fi20++ {
            var v21 bool = (!(!false))
            switch fi20 {
            case 9223372036854775807:
                var v22 int64 = -9223372036854775808
                _ = v22
            case -9223372036854775808:
                if v21 {
                    v19 = (v19 + (^(v19)))
                }
            case 338:
                fmt.Println(v1, v4, v1)
            default:
                if false {
                    v18 = 65535
                    if (!true) {
                        var v23 bool = (93 > 652)
                        v6 += ((v6 * ((v5 + v6) + 1.5)) + 1e10)
                        fmt.Println(v21, v19)
                        _ = v23
                    } else {
                        fmt.Println(v18, v4, v21)
                        if true {
                            for fi24 := 0; fi24 < 13; fi24++ {
                                v25 := GenS1{F0: (fi20 % ((fi20) | 1)), F1: (v4 | int8(v18))}
                                switch 2147483647 {
                                case -979:
                                    v18 = v18
                                case 128:
                                    v3 += v3
                                default:
                                    switch ("" + "foo bar") {
                                    case "foo bar":
                                        v21 = v21
                                    default:
                                        var v26 int64 = (v19 % ((int64(v3)) | 1))
                                        _ = v26
                                    }
                                }
                                switch 42 {
                                case -360:
                                    v1 = (-2.25 / 0.1)
                                case 32768:
                                    var v27 uint8 = 255
                                    _ = v27
                                case -1:
                                    var v28 float32 = float32(0.1)
                                    _ = v28
                                }
                                _ = fi24
                                _ = v25
                            }
                        }
                        var v29 uint8 = uint8(v18)
                        _ = v29
                    }
                } else {
                    v16 = append(v16, fi20)
                    switch fi20 {
                    case 1:
                        v30 := GenS3{F0: 0}
                        _ = v30
                    default:
                        v3 += 65535
                    }
                }
            }
            _ = fi20
            _ = v21
        }
    } else {
        var v31 int8 = v4
        _ = v31
    }
    switch "foo bar" {
    case "":
        fmt.Println(v5, v5, v6)
    case "foo bar":
        for fi32 := 0; fi32 < 9; fi32++ {
            var v33 uint8 = 0
            v34 := GenS3{F0: 65535}
            _ = fi32
            _ = v33
            _ = v34
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
