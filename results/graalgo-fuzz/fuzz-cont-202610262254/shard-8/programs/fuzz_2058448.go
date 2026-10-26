package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 bool; F2 int8 }
type GenS3 struct { F0 int; F1 int32; F2 uint }

func main() {
    var v1 int32 = 1
    v2 := []string{"hello"}
    v3 := make(map[int]uint32)
    var v4 uint16 = 89
    v4 = (v4 ^ (v4 << ((2) & 63)))
    v1 = (v1 << ((255) & 63))
    v4 -= v4
    var v5 float32 = -2.25
    if false {
        var v6 float32 = v5
        for fi7 := 0; fi7 < 13; fi7++ {
            v1 = v1
            var v8 bool = true
            _ = fi7
            _ = v8
        }
        switch -168 {
        case 32767:
            for fi9 := 0; fi9 < 2; fi9++ {
                if (true || true) {
                    v10 := GenS1{F0: 328}
                    switch fi9 {
                    case -9223372036854775808:
                        var v11 int8 = 109
                        _ = v11
                    case 9223372036854775807:
                        v12 := GenS1{F0: 18446744073709551615}
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 10; fi13++ {
                        switch fi13 {
                        case -9223372036854775808:
                            if (!(v1 < 0)) {
                                if ((^(fi9)) < (fi13 - fi9)) {
                                    delete(v3, fi13)
                                } else {
                                    v2 = append(v2, ("foo bar" + ("foo bar" + "\n\t")))
                                    var v14 bool = true
                                    _ = v14
                                }
                                v1 += 1
                            }
                        default:
                            v3[fi13] = 0
                        }
                        v1 = v1
                        _ = fi13
                    }
                    _ = v10
                } else {
                    v3[0] = 65535
                    v2 = append(v2, "")
                }
                v15 := GenS2{F0: 2, F1: (!true), F2: -1}
                _ = fi9
                _ = v15
            }
        case 9223372036854775807:
            v1 += v1
        case -1:
            var v16 uint = 4294967295
            _ = v16
        default:
            v6 -= float32(1.5)
        }
        _ = v6
    } else {
        if (!(3.14159 >= 0.0)) {
            var v17 bool = true
            var v18 string = ""
            if (-1 != 42) {
                var v19 bool = v17
                _ = v19
            } else {
                for fi20 := 0; fi20 < 13; fi20++ {
                    if (v17 || v17) {
                        v5 -= v5
                    } else {
                        if len(v2) > 0 {
                            v2[0] = v18
                        }
                    }
                    for fi21 := 0; fi21 < 16; fi21++ {
                        var v22 int64 = -9223372036854775808
                        if (-2.25 < ((((0.1 / (((v5 - v5) + 1e10) * float32(((0.0 + (((0.0 / (((0.0 * -2.25) * (float64(v5) / ((1e10 - 0.0) / (-2.25 - 1.5)))) * float64(v5))) + (0.1 + ((((float64(v5) - (1.5 / 1.0)) * (3.14159 - 3.14159)) / 1.5) + 0.0))) / 1.5)) + -2.25)))) + v5) * (v5 / (v5 * (v5 - 0.0)))) * v5)) {
                            v23 := GenS2{F0: 1, F1: v17, F2: 100}
                            _ = v23
                        }
                        var v24 uint = 83
                        var v25 string = "foo bar"
                        _ = fi21
                        _ = v22
                        _ = v24
                        _ = v25
                    }
                    var v26 int16 = -32768
                    _ = fi20
                    _ = v26
                }
            }
            _ = v17
            _ = v18
        } else {
            v4 = (v4 / ((46) | 1))
        }
        for fi27 := 0; fi27 < 17; fi27++ {
            var v28 int8 = 42
            v28 += v28
            _ = fi27
            _ = v28
        }
        fmt.Println(v4, v1)
    }
    for fi29 := 0; fi29 < 6; fi29++ {
        fmt.Println(v1)
        switch fi29 {
        case 146:
            v4 = 42
        default:
            v1 = 0
        }
        v5 = ((0.0 / ((v5 - v5) / v5)) / 1.5)
        _ = fi29
    }
    v30 := GenS2{F0: 484, F1: (42 > 128), F2: -128}
    for fi31 := 0; fi31 < 3; fi31++ {
        v1 = int32(v4)
        v1 += 2147483647
        v5 = v5
        _ = fi31
    }
    var v32 uint32 = 2
    for fi33 := 0; fi33 < 10; fi33++ {
        v5 -= v5
        fmt.Println(fi33)
        _ = fi33
    }
    v4 = (^(v4))
    var v34 uint32 = (v32 >> ((18446744073709551615) & 63))
    v1 = -531
    if true {
        v35 := GenS3{F0: 127, F1: v1, F2: uint(v34)}
        switch "" {
        case "\n\t":
            var v36 int16 = 1000
            _ = v36
        case "foo bar":
            var v37 float64 = -2.25
            _ = v37
        }
        var v38 int32 = v1
        _ = v35
        _ = v38
    } else {
        var v39 string = "\n\t"
        v2 = append(v2, "hello")
        _ = v39
    }
    if (!((1.5 * ((((1.5 * float64(v5)) / float64(3.14159)) / 1.0) * 0.0)) < 0.1)) {
        var v40 uint8 = 63
        _ = v40
    } else {
        v41 := GenS1{F0: 65535}
        if (true || true) {
            for fi42 := 0; fi42 < 13; fi42++ {
                v2 = append(v2, "foo bar")
                _ = fi42
            }
        }
        if (v5 <= 0.0) {
            switch 127 {
            case -387:
                _ = len(v2)
            }
            if ((v34 << ((255) & 63)) >= 2147483648) {
                _ = len(v3)
                var v43 bool = false
                _ = v43
            } else {
                v5 = v5
            }
        }
        _ = v41
    }
    v44 := GenS1{F0: 2}
    if len(v2) > 0 {
        v2[0] = ("" + "\n\t")
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v34
    _ = v4
    _ = v44
    _ = v5
}
