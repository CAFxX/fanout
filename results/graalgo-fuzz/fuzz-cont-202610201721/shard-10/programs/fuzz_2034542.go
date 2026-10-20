package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 int8; F2 string }

func main() {
    var v1 uint16 = 255
    v2 := make(map[int]bool)
    var v3 int8 = -1
    var v4 float64 = 1.5
    v5 := make(map[uint32]uint32)
    var v6 uint16 = 2
    v4 = (v4 * v4)
    v7 := GenS1{F0: (v6 | 255)}
    for fi8 := 0; fi8 < 8; fi8++ {
        for fi9 := 0; fi9 < 14; fi9++ {
            fmt.Println(fi8, fi9)
            var v10 float32 = 0.1
            v10 = (v10 / float32(float64(0.1)))
            v4 = 0.1
            var v11 float64 = -2.25
            switch ("hello" + (("a" + "a") + "a")) {
            case "\n\t":
                v12 := GenS1{F0: 2}
                _ = v12
            default:
                v13 := []int{9223372036854775807, 658}
                if true {
                    v10 = 0.1
                    fmt.Println(v1, v6)
                    var v14 int32 = 0
                    _ = v14
                } else {
                    var v15 uint32 = 2147483647
                    _ = v15
                }
                fmt.Println(fi9)
                _ = v13
            }
            switch fi8 {
            case -9223372036854775808:
                fmt.Println(v3, v11)
            default:
                v16 := v2[32768]
                _ = v16
            }
            if false {
                v4 += v11
                v17 := GenS1{F0: v6}
                _ = v17
            } else {
                if (v10 >= 1.5) {
                    var v18 int16 = 384
                    v3 = v3
                    _ = v18
                } else {
                    switch "" {
                    case "hello":
                        v4 += v11
                    case "\n\t":
                        v6 += v1
                    }
                }
            }
            var v19 uint64 = 1
            if (4294967295 >= 255) {
                switch ("a" + (("" + "") + "a")) {
                case "":
                    v20 := []int{9223372036854775807}
                    _ = v20
                case "a":
                    var v21 uint8 = 90
                    _ = v21
                default:
                    var v22 uint8 = uint8(v6)
                    _ = v22
                }
                v1 = v6
                switch (fi9 / (((fi8 % (((fi9 & (fi9 | fi8))) | 1))) | 1)) {
                case -9223372036854775808:
                    switch 0 {
                    case -909:
                        var v23 int = fi8
                        _ = v23
                    default:
                        v24 := GenS1{F0: (v6 >> ((18446744073709551615) & 63))}
                        fmt.Println(fi8, v19)
                        if (!false) {
                            fmt.Println(fi9, fi9, v10)
                        }
                        _ = v24
                    }
                case 128:
                    var v25 int8 = -128
                    _ = v25
                case 961:
                    var v26 uint32 = 0
                    _ = v26
                }
                var v27 float32 = v10
                v4 = v11
                v28 := GenS2{F0: 255, F1: v3, F2: "foo bar"}
                _ = v27
                _ = v28
            }
            if (true && true) {
                v1 += 42
                fmt.Println(v4)
            } else {
                var v29 uint16 = v1
                var v30 uint64 = (^(v19))
                v31 := GenS1{F0: (v6 ^ 266)}
                _ = v29
                _ = v30
                _ = v31
            }
            if (!(!(0 != 518))) {
                switch "foo bar" {
                case "":
                    v19 -= 0
                case "hello":
                    v32 := v2[(fi8 / (((fi9 - (fi9 ^ fi8))) | 1))]
                    var v33 uint16 = 42
                    _ = v32
                    _ = v33
                }
                fmt.Println(v4, v11, v6)
            }
            _ = fi9
            _ = v10
            _ = v11
            _ = v19
        }
        var v34 uint64 = 990
        _ = fi8
        _ = v34
    }
    for fi35 := 0; fi35 < 9; fi35++ {
        v36 := GenS1{F0: (^(v1))}
        if (!(342 <= 561)) {
            v4 = (float64(1.5) - 0.0)
        }
        var v37 float32 = 3.14159
        _ = fi35
        _ = v36
        _ = v37
    }
    fmt.Println(v1, v1, v4)
    v1 = 255
    var v38 bool = ((1e10 + (-2.25 * 0.0)) > (v4 / v4))
    fmt.Println(v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
