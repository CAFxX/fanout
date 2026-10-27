package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 string; F2 int16 }
func genFn0() uint {
    var v6 uint = 18446744073709551615
    for fi7 := 0; fi7 < 7; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            switch "\n\t" {
            case "\n\t":
                if (true && true) {
                    v9 := GenS1{F0: fi8, F1: -325, F2: (v6 % ((v6) | 1))}
                    _ = v9
                }
            }
            v6 += (^(v6))
            if true {
                if (42 > 1) {
                    if (!(-795 >= 32767)) {
                        var v10 uint8 = 32
                        _ = v10
                    }
                    var v11 int = (fi7 & (fi8 << (((v6 / (((v6 & v6)) | 1))) & 63)))
                    var v12 uint64 = 1
                    _ = v11
                    _ = v12
                }
                fmt.Println(fi8, fi8, v6)
                v6 = v6
            } else {
                switch fi8 {
                case 2147483647:
                    if (3.14159 <= (float64((float32(0.0) * (3.14159 - (0.0 * ((((0.0 * 1.5) * 1e10) / float32(float64(1.5))) * (((1.5 - (3.14159 / -2.25)) + -2.25) - ((1e10 - -2.25) - -2.25))))))) - 3.14159)) {
                        v6 += (v6 & (v6 / ((0) | 1)))
                        var v13 float64 = 0.0
                        _ = v13
                    }
                case -1:
                    fmt.Println(v6, v6)
                case 42:
                    fmt.Println(v6, fi7, v6)
                default:
                    if true {
                        v14 := []int{(fi7 % (((fi8 >> ((658) & 63))) | 1)), fi8}
                        var v15 bool = true
                        v15 = (1.5 > (float64(1.5) / (((1e10 / 1.0) / ((1.5 / -2.25) - (float64(3.14159) + (1e10 * (1e10 - (3.14159 + 3.14159)))))) - 0.0)))
                        _ = v14
                        _ = v15
                    }
                }
            }
            _ = fi8
        }
        v6 = 18446744073709551615
        if (false && true) {
            fmt.Println(fi7)
        } else {
            var v16 float64 = ((-2.25 / (3.14159 / 1.0)) + -2.25)
            var v17 uint8 = 2
            _ = v16
            _ = v17
        }
        _ = fi7
    }
    fmt.Println(v6)
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 string = "a"
        for fi20 := 0; fi20 < 15; fi20++ {
            v6 = (^(v6))
            _ = fi20
        }
        fmt.Println(v6, v6)
        _ = fi18
        _ = v19
    }
    _ = v6
    return (v6 / ((v6) | 1))
}

func main() {
    var v1 int8 = 0
    var v2 int32 = 2147483647
    v3 := []int{1}
    var v4 int8 = (^(v1))
    var v5 float32 = -2.25
    if true {
        var v21 float64 = 0.1
        _ = v21
    } else {
        var v22 uint16 = 1
        for fi23 := 0; fi23 < 19; fi23++ {
            switch fi23 {
            case 361:
                v1 = (v1 | v1)
            case 808:
                var v24 uint8 = 162
                _ = v24
            default:
                if (!(true || true)) {
                    fmt.Println(v1, v4)
                    v1 = -32
                } else {
                    if false {
                        var v25 int16 = -418
                        v26 := GenS2{F0: 42, F1: v2}
                        _ = v25
                        _ = v26
                    } else {
                        if false {
                            fmt.Println(v2, v22, v5)
                            var v27 uint16 = 0
                            fmt.Println(v4, v22)
                            _ = v27
                        }
                        switch 415 {
                        case 128:
                            var v28 uint32 = 4294967295
                            _ = v28
                        }
                        var v29 bool = true
                        _ = v29
                    }
                    if (215 < -1) {
                        var v30 int8 = -1
                        _ = v30
                    }
                }
            }
            switch "hello" {
            case "hello":
                if false {
                    v31 := GenS3{F0: ("hello" + ("\n\t" + ("a" + ("hello" + "hello"))))}
                    v1 += v1
                    if (32767 < 0) {
                        fmt.Println(v1, v22, v5)
                        for fi32 := 0; fi32 < 8; fi32++ {
                            var v33 int32 = v2
                            _ = fi32
                            _ = v33
                        }
                    } else {
                        for fi34 := 0; fi34 < 4; fi34++ {
                            if ("" < "hello") {
                                v4 = (v1 ^ -14)
                            }
                            _ = fi34
                        }
                        var v35 float64 = (1e10 - -2.25)
                        _ = v35
                    }
                    _ = v31
                } else {
                    fmt.Println(v2, v1)
                    v4 -= v1
                    if len(v3) > 0 {
                        v3[0] = fi23
                    }
                }
            }
            _ = fi23
        }
        var v36 int = -1
        _ = v22
        _ = v36
    }
    var v37 uint = 0
    if true {
        switch "" {
        case "\n\t":
            var v38 string = (("hello" + "a") + "hello")
            _ = v38
        case "foo bar":
            v39 := GenS4{F0: 9223372036854775807}
            _ = v39
        case "":
            v40 := GenS5{F0: (v2 << (((v37 * v37)) & 63)), F1: "a", F2: -716}
            _ = v40
        default:
            v2 = v2
        }
        v1 = (v4 ^ 0)
    }
    fmt.Println(v37, v37, v5)
    v41 := GenS1{F0: 202, F1: 42, F2: v37}
    v42 := GenS3{F0: "a"}
    v4 -= (v1 | (v4 % ((v1) | 1)))
    v5 -= (float32(1e10) * v5)
    var v43 int8 = (v1 - v1)
    for fi44 := 0; fi44 < 16; fi44++ {
        v3 = append(v3, (fi44 * (fi44 * fi44)))
        var v45 int16 = int16(v2)
        v3 = append(v3, fi44)
        _ = fi44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
