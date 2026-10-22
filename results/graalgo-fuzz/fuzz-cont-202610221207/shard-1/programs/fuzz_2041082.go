package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 int16; F2 uint64 }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 int8 }
func genFn0() string {
    switch "\n\t" {
    case "foo bar":
        var v7 uint8 = 0
        _ = v7
    }
    var v8 int64 = 9223372036854775807
    for fi9 := 0; fi9 < 18; fi9++ {
        v8 += (v8 % ((v8) | 1))
        _ = fi9
    }
    if (float64(0.0) != 1.5) {
        v10 := GenS1{F0: "foo bar", F1: 1, F2: -1}
        var v11 uint16 = 255
        _ = v10
        _ = v11
    }
    _ = v8
    return "\n\t"
}
func genFn1(p0 uint32, p1 int16) int8 {
    if true {
        switch 531 {
        case -9223372036854775808:
            switch ("hello" + "foo bar") {
            case "":
                p0 = p0
            case "a":
                var v12 float32 = (float32(-2.25) - (0.0 - 0.1))
                _ = v12
            default:
                if (p1 >= (p1 >> ((18446744073709551615) & 63))) {
                    p1 += (-(p1))
                    var v13 uint64 = 1
                    _ = v13
                } else {
                    v14 := []int{-256, 1, 127, -2147483648}
                    p0 -= (p0 % (((^(p0))) | 1))
                    _ = v14
                }
            }
        case 32767:
            if false {
                for fi15 := 0; fi15 < 11; fi15++ {
                    for fi16 := 0; fi16 < 18; fi16++ {
                        var v17 uint64 = 1
                        switch fi15 {
                        case 128:
                            p1 = p1
                        case -9223372036854775808:
                            v18 := GenS1{F0: "\n\t", F1: (v17 ^ (v17 / ((v17) | 1))), F2: 1}
                            _ = v18
                        }
                        _ = fi16
                        _ = v17
                    }
                    p1 -= 27
                    p1 += -337
                    _ = fi15
                }
                var v19 float32 = (1.5 * 3.14159)
                v20 := GenS2{F0: 32768}
                _ = v19
                _ = v20
            } else {
                var v21 int32 = 42
                _ = v21
            }
        default:
            fmt.Println(p0, p1)
        }
    } else {
        p1 = p1
    }
    switch -530 {
    case -9223372036854775808:
        switch 1 {
        case -9223372036854775808:
            fmt.Println(p1)
        }
    case 0:
        if (!true) {
            p1 -= (p1 | (p1 | (p1 - p1)))
            v22 := []int{}
            _ = v22
        }
    case 9223372036854775807:
        v23 := []int{481, -300, 1, int(p0)}
        _ = v23
    default:
        fmt.Println(p1)
    }
    v24 := GenS3{F0: 18446744073709551615, F1: 1, F2: 443}
    var v25 int64 = 9223372036854775807
    v26 := GenS4{F0: 847, F1: 1}
    _ = p0
    _ = p1
    _ = v24
    _ = v25
    _ = v26
    return 100
}

func main() {
    var v1 float64 = 0.0
    var v2 uint32 = 2147483648
    v3 := make(map[int]int)
    var v4 uint64 = 4294967295
    var v5 uint = 255
    var v6 int64 = -86
    v3[42] = -691
    switch "\n\t" {
    case "a":
        var v27 uint16 = 1
        _ = v27
    case "\n\t":
        if (128 <= 0) {
            v4 = 472
        }
    default:
        switch 1 {
        case 9223372036854775807:
            v28 := GenS5{F0: 2147483647, F1: (-843 > 0)}
            _ = v28
        }
    }
    v29 := GenS4{F0: 1, F1: uint64(v5)}
    v4 += (v4 >> ((v5) & 63))
    v30 := GenS6{F0: 100}
    v4 += v4
    fmt.Println(v4, v2, v6)
    var v31 uint8 = 30
    var v32 string = "hello"
    var v33 uint8 = v31
    var v34 uint16 = 1
    fmt.Println(v4, v4, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
