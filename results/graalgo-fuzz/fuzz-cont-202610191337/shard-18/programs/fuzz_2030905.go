package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32 }
func genFn0(p0 bool, p1 string, p2 int8) uint64 {
    var v6 uint32 = 65535
    var v7 float64 = ((0.0 / (1.5 - (1e10 / (0.1 / 1e10)))) + float64((-2.25 / -2.25)))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    var v1 float32 = (3.14159 / 1.0)
    var v2 int64 = 901
    v3 := []string{"hello", (("hello" + ((((("" + "a") + "hello") + "hello") + "") + "\n\t")) + ""), "foo bar"}
    v4 := make(map[int]string)
    v5 := []uint32{65535, 2147483647, 2147483648}
    if false {
        v2 -= (^(v2))
    }
    if false {
        if false {
            var v8 int32 = 42
            v9 := GenS1{F0: 65535, F1: 65535, F2: -128}
            _ = v8
            _ = v9
        } else {
            var v10 uint64 = 937
            var v11 float64 = float64(v1)
            fmt.Println(v11, v1, v10)
            _ = v10
            _ = v11
        }
        _ = len(v3)
    }
    var v12 int64 = 1
    v13 := GenS1{F0: 2, F1: 0, F2: 100}
    v12 += (-(v12))
    v14 := GenS2{F0: (!(false || (0.0 >= 1.5))), F1: 0}
    for fi15 := 0; fi15 < 14; fi15++ {
        switch "hello" {
        case "a":
            if (!true) {
                v16 := GenS3{F0: 0}
                var v17 int64 = -881
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = v16
                _ = v17
            } else {
                switch fi15 {
                case 9223372036854775807:
                    v12 = v2
                case 407:
                    v12 = v2
                default:
                    var v18 uint = 1
                    _ = v18
                }
                v1 -= (-2.25 * float32(((0.1 / 1.0) * 1e10)))
                if ((v12 & 744) <= 65) {
                    switch "" {
                    case "":
                        v19 := GenS3{F0: 479}
                        _ = v19
                    case "foo bar":
                        var v20 uint8 = 128
                        _ = v20
                    case "a":
                        var v21 int64 = (v12 ^ -695)
                        _ = v21
                    }
                    var v22 int = fi15
                    fmt.Println(v2, v12, v22)
                    _ = v22
                }
            }
        default:
            v12 -= (v2 >> ((393) & 63))
        }
        var v23 uint16 = 707
        _ = fi15
        _ = v23
    }
    v1 += (v1 + (1e10 * v1))
    v5 = append(v5, 1)
    if (0 < 42) {
        if (((((("a" + (("\n\t" + "\n\t") + "")) + "hello") + ("foo bar" + ((("foo bar" + "") + "") + "a"))) + "foo bar") + "") == (("a" + "") + ("" + "\n\t"))) {
            var v24 uint16 = 255
            var v25 int32 = 802
            if ((-(v2)) < (v2 / (((v2 + int64(v25))) | 1))) {
                if false {
                    v12 = -121
                }
            } else {
                v12 -= v2
                var v26 float64 = (float64(0.1) / 1.5)
                _ = v26
            }
            _ = v24
            _ = v25
        }
        v1 = v1
    } else {
        if (!false) {
            fmt.Println(v12, v2, v12)
            if (-80 != 42) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v12 = v2
                v27 := GenS4{F0: 849}
                _ = v27
            }
        }
        v1 = v1
    }
    v28 := GenS2{F0: false, F1: 486}
    v29 := GenS1{F0: 0, F1: 121, F2: -66}
    var v30 int = 42
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        var v31 uint8 = 127
        _ = v31
    case "a":
        v5 = append(v5, 1)
    default:
        v12 += (-(v2))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
