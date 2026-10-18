package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 string; F2 string }
type GenS4 struct { F0 int64; F1 uint8; F2 uint16 }
type GenS5 struct { F0 int64; F1 uint32; F2 uint32 }

func main() {
    var v1 string = ""
    v2 := make(map[string]bool)
    var v3 float32 = -2.25
    v4 := []uint32{1, 2147483648, 2147483648, 2}
    var v5 uint64 = 4294967295
    v1 = ((v1 + "a") + (v1 + ""))
    if (!true) {
        v6 := GenS1{F0: 0, F1: 1}
        _ = v6
    }
    switch ("hello" + v1) {
    case "foo bar":
        var v7 int = -2147483648
        _ = v7
    case "":
        v3 -= (v3 - v3)
    case "a":
        for fi8 := 0; fi8 < 19; fi8++ {
            if ((fi8 / (((fi8 ^ fi8)) | 1)) >= fi8) {
                switch (((v1 + (v1 + v1)) + v1) + v1) {
                case "foo bar":
                    if (v3 >= (0.1 * v3)) {
                        v9 := GenS1{F0: -32768, F1: 42}
                        _ = v9
                    }
                default:
                    v10 := GenS2{F0: 921}
                    _ = v10
                }
                if (65535 < 255) {
                    v5 += v5
                    var v11 int16 = 1
                    fmt.Println(v11, v1)
                    _ = v11
                } else {
                    for fi12 := 0; fi12 < 19; fi12++ {
                        v13 := GenS3{F0: 2147483648, F1: ("a" + v1), F2: ((v1 + v1) + "a")}
                        fmt.Println(v5)
                        var v14 uint8 = uint8(v5)
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    v5 += 29
                    v2[v1] = true
                }
                v3 = 1e10
            }
            _ = fi8
        }
    default:
        var v15 bool = true
        _ = v15
    }
    var v16 int = 904
    switch (v16 % ((-772) | 1)) {
    case 32768:
        var v17 uint16 = 0
        _ = v17
    case 32767:
        v1 = (((v1 + (((("\n\t" + "") + (v1 + "hello")) + v1) + (v1 + "hello"))) + v1) + v1)
    case -9223372036854775808:
        if (587 > 132) {
            v18 := GenS4{F0: 573, F1: 128, F2: 286}
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v18
        } else {
            v4 = append(v4, 462)
        }
    default:
        if ((v3 + (v3 + (((float32(0.0) * 3.14159) - float32(1.5)) + (v3 / -2.25)))) != (0.1 / (-2.25 / v3))) {
            if (!(!(true && false))) {
                v19 := GenS5{F0: 1, F1: 848, F2: 673}
                var v20 string = ((v1 + (v1 + v1)) + v1)
                _ = v19
                _ = v20
            }
            var v21 int32 = 897
            if true {
                var v22 string = (v1 + ("" + v1))
                var v23 int16 = -383
                _ = len(v2)
                _ = v22
                _ = v23
            }
            _ = v21
        } else {
            if false {
                v2[(v1 + ("a" + ("\n\t" + v1)))] = true
                fmt.Println(v1, v3)
            }
            v4 = append(v4, 0)
        }
    }
    var v24 bool = true
    var v25 uint64 = v5
    var v26 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
