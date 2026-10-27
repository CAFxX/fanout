package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 int64; F2 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32; F1 uint }
func genFn0() uint8 {
    if true {
        var v6 uint8 = 2
        if (!false) {
            if ((42 <= 2) && false) {
                fmt.Println(v6)
            } else {
                v7 := []int{}
                for fi8 := 0; fi8 < 13; fi8++ {
                    fmt.Println(fi8)
                    v9 := GenS1{F0: 42, F1: 18446744073709551615}
                    v10 := GenS1{F0: 211, F1: 4294967295}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                var v11 uint = 42
                _ = v11
                _ = v7
            }
        }
        _ = v6
    } else {
        var v12 int32 = 42
        _ = v12
    }
    if (!true) {
        var v13 int16 = 1
        var v14 uint8 = 11
        var v15 float64 = (0.1 - (((float64((0.0 / (1.5 + (-2.25 + 1e10)))) / 1.5) * 0.0) + 0.1))
        _ = v13
        _ = v14
        _ = v15
    }
    switch "foo bar" {
    case "":
        var v16 uint8 = 0
        _ = v16
    case "hello":
        var v17 int8 = -32
        _ = v17
    }
    return 1
}
func genFn1(p0 int32, p1 string) bool {
    v18 := []int{9223372036854775807, -901, -213, 9223372036854775807}
    var v19 uint64 = 18446744073709551615
    fmt.Println(v19, p1, p0)
    v20 := GenS1{F0: 1, F1: 369}
    var v21 uint32 = 0
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    _ = v20
    _ = v21
    return true
}

func main() {
    var v1 string = "hello"
    var v2 uint16 = 325
    v3 := []int{670}
    var v4 uint16 = (v2 & 0)
    v5 := []int{-2147483648, 2147483647}
    if true {
        v4 = v4
    }
    fmt.Println(v4, v1, v1)
    var v22 float32 = -2.25
    v22 = v22
    if (-256 != -794) {
        var v23 int16 = 1
        for fi24 := 0; fi24 < 12; fi24++ {
            v25 := GenS2{F0: 1000}
            v22 = v22
            v4 -= (v4 << ((1) & 63))
            _ = fi24
            _ = v25
        }
        _ = v23
    }
    var v26 float64 = -2.25
    switch "a" {
    case "":
        var v27 string = "hello"
        _ = v27
    default:
        switch ("foo bar" + "hello") {
        case "foo bar":
            var v28 string = "\n\t"
            _ = v28
        case "a":
            v29 := GenS3{F0: -978, F1: -9223372036854775808, F2: 1}
            _ = v29
        case "hello":
            switch v1 {
            case "hello":
                if len(v5) > 0 {
                    v5[0] = -51
                }
            case "foo bar":
                v26 += ((1e10 - ((((v26 - v26) + float64(v22)) + 3.14159) - v26)) * v26)
            }
        default:
            for fi30 := 0; fi30 < 8; fi30++ {
                var v31 uint8 = 119
                switch (v1 + v1) {
                case "\n\t":
                    if true {
                        v4 = (v2 ^ (v2 / ((664) | 1)))
                        switch (fi30 & -1) {
                        case 42:
                            _ = len(v5)
                        case 868:
                            switch (fi30 / ((fi30) | 1)) {
                            case 4294967295:
                                v31 = v31
                            case 32767:
                                for fi32 := 0; fi32 < 20; fi32++ {
                                    switch "hello" {
                                    case "":
                                        if len(v5) > 0 {
                                            v5[0] = fi30
                                        }
                                        fmt.Println(v1)
                                    case "hello":
                                        var v33 int64 = -436
                                        _ = v33
                                    }
                                    _ = fi32
                                }
                            case 9223372036854775807:
                                v5 = append(v5, fi30)
                            }
                        case 1:
                            if true {
                                fmt.Println(v22, v22)
                            }
                        default:
                            v34 := GenS4{F0: v2}
                            _ = v34
                        }
                    } else {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                case "a":
                    var v35 int16 = 1
                    _ = v35
                default:
                    if true {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if true {
                            for fi36 := 0; fi36 < 13; fi36++ {
                                fmt.Println(fi30)
                                v22 += (v22 * v22)
                                _ = fi36
                            }
                            var v37 int32 = -776
                            for fi38 := 0; fi38 < 16; fi38++ {
                                switch v1 {
                                case "\n\t":
                                    if (!false) {
                                        v37 += -2147483648
                                        if (0 < 152) {
                                            if false {
                                                v37 = -2147483648
                                                v39 := GenS5{F0: 2, F1: 2}
                                                _ = v39
                                            }
                                            var v40 float32 = (v22 + (v22 - 1e10))
                                            v1 = (v1 + v1)
                                            _ = v40
                                        }
                                    }
                                }
                                v1 = v1
                                _ = fi38
                            }
                            _ = v37
                        }
                    }
                }
                fmt.Println(fi30)
                _ = fi30
                _ = v31
            }
        }
    }
    v41 := GenS2{F0: 0}
    for fi42 := 0; fi42 < 4; fi42++ {
        for fi43 := 0; fi43 < 2; fi43++ {
            fmt.Println(v1, v26, v2)
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v2 -= (^(v4))
            _ = fi43
        }
        for fi44 := 0; fi44 < 12; fi44++ {
            if (-63 >= 2147483647) {
                v1 = ("" + v1)
                switch v1 {
                case "foo bar":
                    v45 := GenS2{F0: 9223372036854775807}
                    _ = v45
                case "":
                    if len(v5) > 0 {
                        v5[0] = fi42
                    }
                }
                var v46 int8 = -128
                _ = v46
            }
            v2 -= v2
            switch (fi44 / ((32768) | 1)) {
            case -1:
                var v47 int16 = int16(v2)
                _ = v47
            case 4294967295:
                if false {
                    var v48 float32 = float32((v26 - float64((1.5 + (v22 / v22)))))
                    var v49 uint64 = 18446744073709551615
                    if ((v22 / v22) < v22) {
                        v5 = append(v5, -1)
                        switch ("" + (((v1 + "\n\t") + ((v1 + "\n\t") + "")) + "foo bar")) {
                        case "a":
                            var v50 float32 = float32(((v26 * 1e10) / float64(3.14159)))
                            _ = v50
                        case "":
                            if true {
                                switch "foo bar" {
                                case "\n\t":
                                    fmt.Println(v2)
                                }
                                fmt.Println(v49, fi42)
                            }
                        default:
                            switch (fi42 >> ((65535) & 63)) {
                            case 32767:
                                v51 := GenS2{F0: -1}
                                _ = v51
                            case 270:
                                fmt.Println(fi42, v49, fi44)
                            }
                        }
                        v4 += (^(v2))
                        var v52 uint8 = 128
                        v3 = append(v3, (fi42 & 2147483647))
                        _ = v52
                    }
                    v22 -= v22
                    _ = v48
                    _ = v49
                }
            case -2147483648:
                var v53 int32 = 945
                _ = v53
            }
            v2 += 248
            fmt.Println(v26, fi44, v1)
            _ = fi44
        }
        _ = fi42
    }
    v2 += 2
    v22 += v22
    if true {
        v54 := GenS1{F0: -852, F1: 627}
        _ = v54
    }
    v2 = (v4 * (^(v2)))
    var v55 string = v1
    fmt.Println(v22, v55)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v55
}
