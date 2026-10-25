package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 int8; F2 string }
type GenS4 struct { F0 string; F1 int64; F2 int64 }

func main() {
    v1 := []uint32{2147483648, 2147483647}
    var v2 string = "\n\t"
    var v3 int64 = 9223372036854775807
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]bool)
    var v6 uint16 = 866
    for fi7 := 0; fi7 < 13; fi7++ {
        fmt.Println(v6, v6, fi7)
        switch 9223372036854775807 {
        case 4294967295:
            v3 = 78
        case -259:
            fmt.Println(v6)
        default:
            if false {
                v3 = v3
                _ = len(v1)
                v1 = append(v1, 1)
            } else {
                for fi8 := 0; fi8 < 19; fi8++ {
                    v2 = "\n\t"
                    _ = fi8
                }
                if (!true) {
                    if false {
                        v3 -= (v3 % ((143) | 1))
                        var v9 int = (fi7 * fi7)
                        _ = v9
                    } else {
                        delete(v5, 2147483648)
                        v10 := GenS1{F0: -773, F1: -93}
                        _ = v10
                    }
                    var v11 int8 = 39
                    for fi12 := 0; fi12 < 12; fi12++ {
                        fmt.Println(v11, v2)
                        var v13 int = (fi7 % ((fi12) | 1))
                        switch "a" {
                        case "":
                            v6 -= (v6 - 0)
                        case "foo bar":
                            v14 := GenS2{F0: 39}
                            _ = v14
                        default:
                            delete(v5, 1)
                        }
                        _ = fi12
                        _ = v13
                    }
                    _ = v11
                }
            }
        }
        v15 := GenS3{F0: 963, F1: 42, F2: "a"}
        _ = fi7
        _ = v15
    }
    var v16 uint16 = v6
    _ = len(v5)
    var v17 float32 = 1.5
    v18 := GenS2{F0: 951}
    _ = len(v1)
    switch ((v2 + "") + "hello") {
    case "":
        switch v2 {
        case "hello":
            fmt.Println(v6, v16)
        case "\n\t":
            if (0.1 == (3.14159 * 1.5)) {
                if (1e10 == 1e10) {
                    if (1000 > 1000) {
                        fmt.Println(v2)
                    } else {
                        fmt.Println(v2, v6)
                        delete(v5, 4294967295)
                        fmt.Println(v17, v3)
                    }
                    var v19 uint16 = v16
                    v19 = v19
                    _ = v19
                }
            }
        default:
            if (!(true && true)) {
                for fi20 := 0; fi20 < 20; fi20++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v2 = (v2 + v2)
                    _ = fi20
                }
                v1 = append(v1, 0)
                var v21 int8 = 1
                _ = v21
            }
        }
    case "a":
        if true {
            v5[1] = false
            v22 := GenS1{F0: 42, F1: -50}
            _ = v22
        } else {
            var v23 uint32 = 998
            _ = v23
        }
    case "hello":
        v6 = v6
    default:
        switch -979 {
        case 128:
            var v24 int8 = 1
            _ = v24
        case -186:
            if (128 >= -2147483648) {
                v25 := GenS2{F0: 255}
                _ = v25
            }
        case 652:
            v1 = append(v1, 817)
        default:
            if (false && false) {
                var v26 float32 = (float32(0.0) * v17)
                _ = v26
            }
        }
    }
    switch "foo bar" {
    case "a":
        var v27 uint64 = 42
        _ = v27
    case "hello":
        v1 = append(v1, 266)
    case "\n\t":
        v17 = v17
    default:
        for fi28 := 0; fi28 < 17; fi28++ {
            if len(v1) > 0 {
                v1[0] = 2
            }
            fmt.Println(v6, v17, v17)
            v2 = v2
            _ = fi28
        }
    }
    if (1000 == -32768) {
        var v29 int64 = -9223372036854775808
        v30 := GenS4{F0: (v2 + ("\n\t" + v2)), F1: (v29 - (-(v29))), F2: int64(v16)}
        v6 += (v16 * 65535)
        _ = v29
        _ = v30
    }
    fmt.Println(v3, v6)
    var v31 uint16 = (^(v16))
    if len(v1) > 0 {
        v1[0] = 0
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
