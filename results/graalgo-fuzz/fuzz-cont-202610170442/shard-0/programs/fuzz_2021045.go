package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 uint }

func main() {
    var v1 float64 = -2.25
    v2 := make(map[string]string)
    v3 := []uint32{0, 0, 4294967295}
    v4 := make(map[int]string)
    v5 := []string{"foo bar", "hello"}
    v6 := []string{"foo bar"}
    var v7 bool = true
    if (!v7) {
        v1 += v1
        if len(v5) > 0 {
            v5[0] = ""
        }
        switch "a" {
        case "a":
            for fi8 := 0; fi8 < 15; fi8++ {
                for fi9 := 0; fi9 < 1; fi9++ {
                    if v7 {
                        v10 := v4[0]
                        _ = v10
                    }
                    switch fi9 {
                    case 32768:
                        v11 := GenS1{F0: false, F1: 47}
                        _ = v11
                    }
                    v12 := v4[(fi9 * (fi9 | (fi8 * 303)))]
                    _ = fi9
                    _ = v12
                }
                _ = fi8
            }
        case "foo bar":
            if len(v5) > 0 {
                v5[0] = ""
            }
        case "":
            v13 := GenS2{F0: 4294967295}
            _ = v13
        default:
            for fi14 := 0; fi14 < 15; fi14++ {
                for fi15 := 0; fi15 < 2; fi15++ {
                    if v7 {
                        v16 := GenS1{F0: (-128 == 0), F1: 0}
                        if len(v6) > 0 {
                            v6[0] = "a"
                        }
                        _ = v16
                    }
                    switch "hello" {
                    case "a":
                        _ = len(v3)
                        var v17 int16 = -32768
                        _ = v17
                    default:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    fmt.Println(v1)
                    _ = fi15
                }
                v1 = (float64((1e10 / ((1.5 - float32(v1)) + float32(v1)))) + float64(float32(v1)))
                v1 = -2.25
                _ = fi14
            }
        }
    }
    if (100 >= -1) {
        if (42 != 753) {
            v7 = (955 == 657)
        } else {
            fmt.Println(v1)
            var v18 float32 = 1e10
            delete(v4, 9223372036854775807)
            _ = v18
        }
        var v19 uint8 = 2
        _ = v19
    } else {
        _ = len(v4)
        fmt.Println(v1)
        var v20 float32 = ((-2.25 + -2.25) + 1.5)
        _ = v20
    }
    v3 = append(v3, 0)
    if (!(true || (!v7))) {
        v1 = (1.5 / float64(3.14159))
        v7 = (!((0 > 0) && true))
        v7 = v7
    }
    if ((!false) || v7) {
        fmt.Println(v1, v7)
    }
    _ = len(v5)
    if true {
        v1 = 0.1
        if len(v3) > 0 {
            v3[0] = 723
        }
    } else {
        switch -9223372036854775808 {
        case 0:
            v21 := GenS2{F0: 4294967295}
            _ = v21
        case 127:
            var v22 uint64 = 4294967295
            _ = v22
        case 2147483647:
            v23 := GenS3{F0: 1, F1: 206}
            _ = v23
        default:
            v24 := v4[127]
            _ = v24
        }
        var v25 int64 = -562
        fmt.Println(v7, v7, v25)
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
