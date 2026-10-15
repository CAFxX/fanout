package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
func genFn0(p0 uint32, p1 uint, p2 int64) uint64 {
    v6 := GenS1{F0: 42, F1: false}
    var v7 int64 = 0
    v8 := []int{128, 2147483647}
    var v9 bool = true
    switch 20 {
    case 0:
        if (!v9) {
            p2 += p2
            v9 = (float32(0.0) > 1.5)
        } else {
            var v10 int8 = 100
            switch 32767 {
            case -402:
                var v11 bool = (42 <= 3)
                _ = v11
            case 127:
                var v12 uint = p1
                _ = v12
            case 630:
                if false {
                    var v13 uint64 = 486
                    _ = len(v8)
                    if len(v8) > 0 {
                        v8[0] = int(p1)
                    }
                    _ = v13
                }
            }
            _ = v10
        }
    case -9223372036854775808:
        var v14 uint64 = 18446744073709551615
        _ = v14
    case -383:
        switch -382 {
        case -9223372036854775808:
            v15 := GenS1{F0: 0, F1: (0.0 == 1e10)}
            _ = v15
        case 130:
            var v16 int8 = -128
            _ = v16
        case 80:
            switch "foo bar" {
            case "a":
                fmt.Println(p1, p0)
            default:
                if v9 {
                    switch "a" {
                    case "hello":
                        _ = len(v8)
                    case "\n\t":
                        if len(v8) > 0 {
                            v8[0] = 2147483647
                        }
                    case "":
                        var v17 int8 = 1
                        _ = v17
                    default:
                        switch ("" + "hello") {
                        case "":
                            var v18 uint32 = p0
                            _ = v18
                        default:
                            if (0 > 127) {
                                var v19 int = 32768
                                _ = v19
                            }
                        }
                    }
                    fmt.Println(p2, v7, p1)
                    v8 = append(v8, 863)
                }
            }
        }
    default:
        v7 -= (v7 << (((p1 >> (((p1 | p1)) & 63))) & 63))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    var v1 float64 = -2.25
    var v2 int = 760
    var v3 int = (-(v2))
    var v4 uint32 = 2147483648
    var v5 uint8 = 2
    fmt.Println(v1, v2)
    v20 := GenS1{F0: 0, F1: ((v3 / ((int(v5)) | 1)) >= (v3 & (-(v2))))}
    v5 -= (v5 / ((1) | 1))
    var v21 float64 = (float64(1e10) - v1)
    fmt.Println(v5)
    v22 := []int{891, (v3 * v2), v2}
    if true {
        _ = len(v22)
    } else {
        if (0 < genFn0((v4 % (((v4 & 550)) | 1)), 255, -991)) {
            for fi23 := 0; fi23 < 12; fi23++ {
                var v24 uint32 = (v4 >> ((0) & 63))
                v4 = (v4 | (v4 | v4))
                switch "hello" {
                case "\n\t":
                    if len(v22) > 0 {
                        v22[0] = fi23
                    }
                case "a":
                    var v25 float32 = float32((v1 - (3.14159 / v21)))
                    _ = v25
                default:
                    for fi26 := 0; fi26 < 4; fi26++ {
                        var v27 uint32 = v24
                        if ("foo bar" < "a") {
                            var v28 int = v3
                            var v29 int16 = 32767
                            var v30 int32 = -2147483648
                            var v31 string = ("" + "\n\t")
                            _ = len(v22)
                            fmt.Println(v21, fi26, v2)
                            _ = v28
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                        _ = fi26
                        _ = v27
                    }
                }
                _ = fi23
                _ = v24
            }
        } else {
            if false {
                for fi32 := 0; fi32 < 7; fi32++ {
                    v22 = append(v22, fi32)
                    _ = fi32
                }
            } else {
                for fi33 := 0; fi33 < 17; fi33++ {
                    var v34 uint64 = 2
                    _ = len(v22)
                    if true {
                        _ = len(v22)
                        var v35 string = ("\n\t" + ("foo bar" + "foo bar"))
                        fmt.Println(v2, v3, v1)
                        _ = v35
                    }
                    _ = fi33
                    _ = v34
                }
                v36 := GenS1{F0: 1, F1: false}
                _ = v36
            }
            if false {
                v22 = append(v22, v2)
            }
        }
    }
    switch v2 {
    case -708:
        var v37 uint = 0
        _ = v37
    }
    for fi38 := 0; fi38 < 14; fi38++ {
        var v39 uint32 = 2147483647
        var v40 float64 = v21
        _ = fi38
        _ = v39
        _ = v40
    }
    v2 -= (v3 / ((v3) | 1))
    var v41 uint = 2
    v4 += v4
    if false {
        var v42 float64 = v1
        var v43 string = "foo bar"
        var v44 uint16 = 1
        _ = v42
        _ = v43
        _ = v44
    }
    v21 -= v1
    v5 -= 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v41
    _ = v5
}
