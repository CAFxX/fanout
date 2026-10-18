package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 int16 }

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{981, 1}
    var v3 string = "hello"
    v4 := make(map[string]int)
    var v5 int32 = 0
    var v6 int16 = 0
    var v7 int64 = -9223372036854775808
    if (false && (0 >= 42)) {
        v2 = append(v2, 1)
        fmt.Println(v3, v3, v6)
        fmt.Println(v7)
    }
    _ = len(v4)
    if false {
        for fi8 := 0; fi8 < 17; fi8++ {
            if false {
                delete(v4, "hello")
            }
            for fi9 := 0; fi9 < 9; fi9++ {
                v4[v3] = (fi9 & (-(fi8)))
                if false {
                    v2 = append(v2, 0)
                    v3 = v3
                } else {
                    var v10 uint16 = 255
                    _ = v10
                }
                v7 = (v7 + (v7 % (((v7 % ((1) | 1))) | 1)))
                _ = fi9
            }
            _ = len(v1)
            _ = fi8
        }
        var v11 uint32 = 42
        v12 := GenS1{F0: (v3 + v3)}
        _ = v11
        _ = v12
    }
    var v13 float64 = float64(-2.25)
    _ = len(v2)
    if (2 == 1) {
        switch -287 {
        case 9223372036854775807:
            v14 := GenS2{F0: -128}
            _ = v14
        case 111:
            if (18446744073709551615 <= 2) {
                switch 128 {
                case -9223372036854775808:
                    for fi15 := 0; fi15 < 9; fi15++ {
                        var v16 uint32 = 2147483648
                        v17 := GenS3{F0: (v3 + v3), F1: (v6 / (((v6 | (v6 / ((32767) | 1)))) | 1))}
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                case 1:
                    fmt.Println(v6, v13)
                }
                v2 = append(v2, 0)
            } else {
                v5 = v5
                _ = len(v1)
                var v18 string = "a"
                _ = v18
            }
        case 128:
            var v19 int16 = v6
            _ = v19
        default:
            switch ("" + v3) {
            case "foo bar":
                v2 = append(v2, 42)
            }
        }
        fmt.Println(v6)
        for fi20 := 0; fi20 < 8; fi20++ {
            var v21 float32 = 3.14159
            if (65535 == 2147483648) {
                var v22 int16 = (^(v6))
                var v23 int8 = 38
                _ = v22
                _ = v23
            } else {
                var v24 float64 = v13
                _ = v24
            }
            _ = fi20
            _ = v21
        }
    } else {
        var v25 bool = false
        var v26 uint64 = 395
        for fi27 := 0; fi27 < 16; fi27++ {
            delete(v1, fi27)
            v28 := GenS3{F0: v3, F1: 32767}
            switch 1 {
            case -604:
                var v29 int32 = 1
                _ = v29
            case -9223372036854775808:
                v30 := GenS2{F0: 88}
                _ = v30
            case 32768:
                v31 := GenS3{F0: ("a" + ""), F1: v6}
                _ = v31
            }
            _ = fi27
            _ = v28
        }
        _ = v25
        _ = v26
    }
    var v32 uint32 = 42
    for fi33 := 0; fi33 < 14; fi33++ {
        var v34 uint8 = 95
        _ = fi33
        _ = v34
    }
    if (2 != 0) {
        var v35 float32 = -2.25
        _ = v35
    } else {
        v5 += 42
        var v36 uint8 = 1
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
