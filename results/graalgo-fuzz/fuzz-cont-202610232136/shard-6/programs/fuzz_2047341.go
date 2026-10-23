package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 int64; F2 int }
func genFn0(p0 uint64, p1 uint64, p2 int64) int8 {
    switch 0 {
    case -2147483648:
        var v5 int64 = p2
        _ = v5
    case -21:
        var v6 int64 = -9223372036854775808
        _ = v6
    }
    v7 := GenS1{F0: -2147483648}
    v8 := []int{-673}
    var v9 uint8 = 128
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return -1
}
func genFn1(p0 uint8, p1 string) int {
    p0 -= p0
    var v10 uint16 = 42
    var v11 uint = 948
    if (v11 == v11) {
        p0 = p0
        var v12 uint8 = p0
        v13 := []int{-1, 32767, -236, 4294967295}
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return -73
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]string)
    var v3 uint32 = 4294967295
    var v4 uint32 = (v3 | 1)
    switch 127 {
    case 768:
        v14 := GenS2{F0: -1, F1: 1, F2: -1}
        _ = v14
    default:
        if (float64(1.5) < 3.14159) {
            fmt.Println(v3)
            var v15 uint64 = 598
            _ = v15
        }
    }
    v4 = (v3 - 163)
    v16 := []int{}
    if true {
        fmt.Println(v3, v3)
    } else {
        if (false || (1 > 325)) {
            if (328 >= 313) {
                v17 := v2[1]
                v3 = v4
                v18 := v1[9223372036854775807]
                _ = v17
                _ = v18
            }
            v16 = append(v16, 424)
        } else {
            fmt.Println(v3, v4)
            switch "hello" {
            case "\n\t":
                var v19 int64 = -9223372036854775808
                _ = v19
            case "":
                var v20 uint16 = 57
                _ = v20
            }
        }
        for fi21 := 0; fi21 < 1; fi21++ {
            _ = len(v16)
            v3 += 642
            var v22 int32 = 0
            _ = fi21
            _ = v22
        }
    }
    v4 = v3
    v23 := GenS2{F0: -12, F1: 42, F2: genFn1(128, "foo bar")}
    v24 := GenS1{F0: 32768}
    switch "\n\t" {
    case "a":
        switch "\n\t" {
        case "foo bar":
            if (true && false) {
                for fi25 := 0; fi25 < 20; fi25++ {
                    for fi26 := 0; fi26 < 14; fi26++ {
                        fmt.Println(v4, v3)
                        v3 = (v3 * v4)
                        _ = fi26
                    }
                    _ = fi25
                }
            } else {
                if ((true && ((0 > 127) || true)) && (float32(1e10) >= 1e10)) {
                    _ = len(v1)
                }
            }
        case "\n\t":
            v4 -= (^(v3))
        case "hello":
            delete(v1, -385)
        default:
            for fi27 := 0; fi27 < 6; fi27++ {
                v28 := v2[127]
                for fi29 := 0; fi29 < 3; fi29++ {
                    var v30 int16 = 1
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                    v28 = v28
                    _ = fi29
                    _ = v30
                }
                fmt.Println(v4, fi27, v3)
                _ = fi27
                _ = v28
            }
        }
    default:
        for fi31 := 0; fi31 < 3; fi31++ {
            var v32 uint16 = 1
            var v33 int32 = 394
            _ = fi31
            _ = v32
            _ = v33
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
