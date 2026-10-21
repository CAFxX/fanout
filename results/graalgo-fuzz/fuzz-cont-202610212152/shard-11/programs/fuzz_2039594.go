package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 float64 = (0.1 - 1.5)
    var v2 int8 = 127
    var v3 int8 = -1
    if ((!false) || (!("hello" != (("hello" + "") + (("" + ("" + "hello")) + "foo bar"))))) {
        var v4 uint32 = 0
        v1 += 3.14159
        _ = v4
    } else {
        if true {
            var v5 uint64 = 2
            _ = v5
        }
    }
    if (1e10 == (1e10 / 1.0)) {
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := []int{(fi6 + (fi6 - (fi6 ^ (fi6 >> ((18446744073709551615) & 63))))), -779}
            _ = fi6
            _ = v7
        }
        fmt.Println(v1, v3, v3)
    } else {
        if (v1 != float64((3.14159 + 1e10))) {
            v3 -= (v2 | (v3 ^ v2))
        }
        v1 = (v1 - v1)
        fmt.Println(v3, v2, v2)
    }
    v8 := []int{362}
    v2 += v3
    if ((v3 | (v2 + v2)) >= (-(v2))) {
        for fi9 := 0; fi9 < 6; fi9++ {
            if false {
                var v10 uint16 = 2
                v3 += v3
                switch ("foo bar" + ("" + ("" + "foo bar"))) {
                case "":
                    var v11 uint = 4294967295
                    _ = v11
                case "hello":
                    if ((v10 | (^(v10))) < v10) {
                        switch (fi9 | 128) {
                        case 32767:
                            switch (((("" + "a") + ("a" + "hello")) + "foo bar") + "") {
                            case "":
                                var v12 int32 = -453
                                _ = v12
                            case "hello":
                                var v13 int8 = v3
                                _ = v13
                            case "a":
                                v14 := GenS1{F0: v10, F1: 42, F2: 812}
                                _ = v14
                            }
                        }
                        v2 = v3
                        if ((true || (65535 == (v10 / ((v10) | 1)))) && ((!false) || false)) {
                            v15 := GenS1{F0: v10, F1: 1, F2: 65535}
                            _ = v15
                        }
                    } else {
                        v8 = append(v8, -9223372036854775808)
                        v16 := GenS2{F0: 614}
                        var v17 uint16 = (v10 + v10)
                        _ = v16
                        _ = v17
                    }
                default:
                    fmt.Println(v2, v1)
                }
                _ = v10
            }
            fmt.Println(v2, v1)
            for fi18 := 0; fi18 < 12; fi18++ {
                fmt.Println(fi9, v1)
                var v19 float64 = (v1 - 0.0)
                v20 := GenS3{F0: 42, F1: 0}
                if ((true && (true || ((v2 ^ (v2 & v2)) >= -56))) && (42 < 128)) {
                    var v21 float64 = (v19 / -2.25)
                    if true {
                        switch 4294967295 {
                        case 32767:
                            switch "a" {
                            case "":
                                var v22 uint8 = 128
                                _ = v22
                            case "foo bar":
                                var v23 uint64 = 0
                                _ = v23
                            default:
                                var v24 bool = false
                                _ = v24
                            }
                        case -1:
                            v25 := GenS4{F0: 128}
                            _ = v25
                        default:
                            v21 += ((((3.14159 * v19) * (1.5 / v21)) / (v19 + v21)) + (v19 / (v19 + ((v1 / 1.0) - v21))))
                        }
                    } else {
                        if (!false) {
                            switch fi18 {
                            case 32768:
                                var v26 string = "\n\t"
                                _ = v26
                            case -1:
                                var v27 float64 = float64(0.0)
                                _ = v27
                            default:
                                v2 += 35
                            }
                            v19 = (3.14159 + float64(((((-2.25 + 0.1) / float32(float64(((1.5 - (3.14159 * 3.14159)) * 0.1)))) * 1.5) / 0.1)))
                        } else {
                            v3 = 0
                        }
                        switch "hello" {
                        case "foo bar":
                            v28 := GenS4{F0: 1}
                            _ = v28
                        case "\n\t":
                            if true {
                                fmt.Println(v2)
                                fmt.Println(v19)
                                v29 := GenS3{F0: 178, F1: 1}
                                _ = v29
                            }
                        case "":
                            v2 += 42
                        default:
                            if ((^(fi9)) > 128) {
                                v21 -= 3.14159
                            } else {
                                v19 = 1e10
                                var v30 string = "\n\t"
                                v21 -= 3.14159
                                _ = v30
                            }
                        }
                    }
                    _ = v21
                }
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = fi9
        }
    } else {
        var v31 int32 = 1
        _ = v31
    }
    fmt.Println(v1)
    v1 += (((float64(3.14159) - v1) / (3.14159 * (1.5 / 0.1))) / (3.14159 - (v1 * ((v1 / 1e10) + v1))))
    if false {
        v1 = (v1 - ((1e10 / (((v1 * float64((((1.5 + 3.14159) - 0.1) * float32(v1)))) * v1) - v1)) + v1))
    }
    if len(v8) > 0 {
        v8[0] = 32767
    }
    v2 = v3
    fmt.Println(v3, v1, v1)
    var v32 int32 = 2147483647
    for fi33 := 0; fi33 < 6; fi33++ {
        for fi34 := 0; fi34 < 20; fi34++ {
            v1 = v1
            _ = fi34
        }
        switch -769 {
        case 127:
            v3 += (^(v3))
        case 83:
            _ = len(v8)
        default:
            v2 += v3
        }
        _ = fi33
    }
    if len(v8) > 0 {
        v8[0] = 558
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v8
}
