package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint64 }
type GenS2 struct { F0 int; F1 uint16; F2 bool }
type GenS3 struct { F0 int64; F1 bool }

func main() {
    v1 := []int{-447, 128}
    var v2 uint8 = 188
    var v3 int16 = -1
    v2 -= v2
    var v4 string = "hello"
    var v5 uint16 = 32767
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := GenS1{F0: true, F1: 0, F2: 0}
        _ = fi6
        _ = v7
    }
    v2 = v2
    if (-1 <= -1) {
        switch 4294967295 {
        case 2147483647:
            var v8 int8 = 100
            _ = v8
        }
    } else {
        var v9 int32 = int32(v5)
        _ = v9
    }
    switch v4 {
    case "hello":
        fmt.Println(v3)
    }
    if true {
        if ((float32((1.5 / (-2.25 * 0.1))) / ((1e10 + (3.14159 - float32(float64(3.14159)))) / (3.14159 / 1.0))) != (0.0 / (0.1 * -2.25))) {
            if true {
                if (42 > 507) {
                    fmt.Println(v3, v5)
                } else {
                    fmt.Println(v3)
                    v1 = append(v1, -9223372036854775808)
                    v10 := GenS1{F0: true, F1: -1, F2: uint64(v5)}
                    _ = v10
                }
                if len(v1) > 0 {
                    v1[0] = -1
                }
                for fi11 := 0; fi11 < 9; fi11++ {
                    fmt.Println(v4, v5)
                    var v12 uint = 65535
                    _ = fi11
                    _ = v12
                }
            }
        }
        switch -1 {
        case -995:
            v13 := GenS2{F0: -265, F1: v5, F2: false}
            _ = v13
        case -474:
            switch 270 {
            case 32768:
                v14 := GenS2{F0: -278, F1: v5, F2: false}
                _ = v14
            case 9223372036854775807:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case 127:
                if (((4294967295 > 1) && (v5 <= (^(v5)))) || false) {
                    v5 = v5
                    v2 = v2
                } else {
                    v15 := GenS3{F0: 942, F1: (false || true)}
                    _ = v15
                }
            default:
                if len(v1) > 0 {
                    v1[0] = -709
                }
            }
        case 2147483647:
            var v16 int32 = 42
            _ = v16
        default:
            fmt.Println(v4, v3)
        }
        fmt.Println(v2, v3, v5)
    }
    switch ((v4 + ("hello" + (v4 + (v4 + v4)))) + v4) {
    case "a":
        v3 = -492
    default:
        if (777 != 18) {
            for fi17 := 0; fi17 < 15; fi17++ {
                v3 += (v3 * (-(v3)))
                _ = fi17
            }
            v1 = append(v1, -188)
        }
    }
    var v18 uint16 = 967
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
