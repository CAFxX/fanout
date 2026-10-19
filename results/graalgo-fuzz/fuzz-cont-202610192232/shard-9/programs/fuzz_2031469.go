package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 int16 }

func main() {
    var v1 uint = 4294967295
    v2 := []string{("foo bar" + "\n\t")}
    var v3 int8 = 42
    var v4 float32 = 1e10
    var v5 uint16 = 426
    v6 := make(map[uint32]string)
    v1 += (v1 ^ v1)
    if false {
        var v7 int = 32768
        fmt.Println(v5, v7, v3)
        _ = v7
    } else {
        switch 4294967295 {
        case 127:
            v4 = (((3.14159 / -2.25) + (float32(float64(v4)) / 1.0)) - (-2.25 + float32(0.0)))
        case -2147483648:
            for fi8 := 0; fi8 < 3; fi8++ {
                v9 := GenS1{F0: (^(v5))}
                var v10 uint8 = 86
                _ = fi8
                _ = v10
                _ = v9
            }
        case 32767:
            fmt.Println(v1, v5, v4)
        default:
            if ((float64(-2.25) == (0.0 + 3.14159)) && false) {
                var v11 float32 = v4
                if true {
                    v3 -= (v3 >> ((uint(v5)) & 63))
                    var v12 uint = (^(v1))
                    _ = v12
                } else {
                    v13 := GenS2{F0: "a"}
                    _ = v13
                }
                _ = v11
            } else {
                if (1 >= 128) {
                    var v14 string = "a"
                    switch 728 {
                    case 32767:
                        v15 := v6[1]
                        _ = v15
                    case -9223372036854775808:
                        var v16 int16 = 901
                        _ = v16
                    default:
                        fmt.Println(v3)
                    }
                    var v17 uint = v1
                    _ = v14
                    _ = v17
                } else {
                    var v18 int = -891
                    v19 := GenS3{F0: -406}
                    _ = v18
                    _ = v19
                }
                fmt.Println(v1, v3)
                fmt.Println(v1)
            }
        }
    }
    fmt.Println(v3)
    v20 := GenS4{F0: (v1 * 365), F1: 1}
    var v21 string = "a"
    v22 := GenS5{F0: 1000}
    if (((!true) || true) || true) {
        var v23 int8 = (v3 >> ((v1) & 63))
        var v24 string = ("foo bar" + "a")
        v4 -= float32(((1.5 * 0.1) - (3.14159 + 3.14159)))
        _ = v23
        _ = v24
    } else {
        delete(v6, 65535)
        fmt.Println(v5)
        var v25 uint32 = 2147483648
        _ = v25
    }
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
