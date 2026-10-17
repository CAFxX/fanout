package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 string }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 int64 }
func genFn0(p0 uint, p1 uint16) int64 {
    var v7 uint16 = 2
    v8 := []int{-844, -791, 9223372036854775807, -1}
    switch "foo bar" {
    case "hello":
        var v9 float64 = -2.25
        _ = v9
    case "a":
        if false {
            if false {
                fmt.Println(p1, p1)
                var v10 uint32 = 1
                if len(v8) > 0 {
                    v8[0] = 296
                }
                _ = v10
            } else {
                fmt.Println(v7)
                if len(v8) > 0 {
                    v8[0] = 127
                }
                if ((!((false && false) || true)) && true) {
                    for fi11 := 0; fi11 < 13; fi11++ {
                        v7 = p1
                        for fi12 := 0; fi12 < 19; fi12++ {
                            p1 = v7
                            if true {
                                switch "a" {
                                case "":
                                    switch (fi11 - fi12) {
                                    case 415:
                                        _ = len(v8)
                                        if len(v8) > 0 {
                                            v8[0] = fi12
                                        }
                                        if (-9223372036854775808 == -1) {
                                            var v13 float32 = 0.0
                                            v14 := GenS1{F0: 784, F1: 1, F2: "\n\t"}
                                            switch fi11 {
                                            case 2147483647:
                                                v13 += 3.14159
                                            case -1:
                                                var v15 uint32 = 2
                                                _ = v15
                                            }
                                            _ = v13
                                            _ = v14
                                        } else {
                                            v7 += (^(v7))
                                        }
                                        p1 += (v7 / (((v7 | (p1 % (((^(v7))) | 1)))) | 1))
                                        v16 := GenS2{F0: "", F1: 1}
                                        _ = v16
                                    case 32768:
                                        var v17 float64 = (1e10 - ((1e10 - 1.5) * 1e10))
                                        _ = v17
                                    default:
                                        fmt.Println(fi11, p0, fi12)
                                    }
                                case "hello":
                                    switch ("" + ("\n\t" + "a")) {
                                    case "a":
                                        switch "\n\t" {
                                        case "hello":
                                            var v18 float32 = 1.5
                                            _ = v18
                                        case "\n\t":
                                            var v19 int64 = 0
                                            _ = v19
                                        default:
                                            if len(v8) > 0 {
                                                v8[0] = 32767
                                            }
                                        }
                                    default:
                                        switch (fi11 + fi11) {
                                        case 42:
                                            fmt.Println(p1, p1)
                                            var v20 int64 = 42
                                            _ = v20
                                        case 4294967295:
                                            var v21 int = fi12
                                            _ = v21
                                        }
                                    }
                                    var v22 float64 = 3.14159
                                    v7 += (p1 - v7)
                                    _ = v22
                                case "\n\t":
                                    var v23 uint64 = uint64(p0)
                                    _ = v23
                                }
                                if true {
                                    var v24 int = -2147483648
                                    if len(v8) > 0 {
                                        v8[0] = (fi11 ^ 9223372036854775807)
                                    }
                                    v25 := GenS3{F0: p1, F1: (0 == (fi12 - (fi11 & (v24 - v24)))), F2: 255}
                                    _ = v24
                                    _ = v25
                                } else {
                                    v7 += (p1 * p1)
                                }
                                if (887 <= -9223372036854775808) {
                                    var v26 int32 = -370
                                    switch (-(fi11)) {
                                    case -1:
                                        v8 = append(v8, (fi11 % (((^(fi11))) | 1)))
                                    case 0:
                                        var v27 string = "\n\t"
                                        _ = v27
                                    case 168:
                                        switch "foo bar" {
                                        case "\n\t":
                                            fmt.Println(v7, p1)
                                        case "hello":
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                        }
                                    }
                                    _ = v26
                                } else {
                                    switch fi11 {
                                    case 42:
                                        v28 := GenS3{F0: (v7 - 255), F1: false, F2: 1}
                                        _ = v28
                                    }
                                    var v29 uint = (p0 & p0)
                                    _ = v29
                                }
                            }
                            _ = fi12
                        }
                        _ = fi11
                    }
                }
            }
        }
    default:
        p0 = uint(p1)
    }
    v8 = append(v8, 0)
    p0 -= (p0 - (^(p0)))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 bool = true
    v2 := []int{}
    var v3 uint = 4294967295
    var v4 uint8 = 2
    var v5 uint32 = 0
    var v6 int8 = 119
    fmt.Println(v6, v3, v3)
    v2 = append(v2, -1)
    if (1.5 >= ((float32(((0.0 - (float64(1e10) - (1e10 - float64(0.0)))) + (float64(-2.25) * float64((1.5 - (((0.1 + 0.0) - 3.14159) / 1.0)))))) * 3.14159) - -2.25)) {
        v30 := GenS2{F0: "", F1: -839}
        _ = v30
    }
    fmt.Println(v6, v4)
    var v31 float64 = float64(0.0)
    var v32 uint16 = 42
    for fi33 := 0; fi33 < 19; fi33++ {
        v4 = (v4 & (v4 - v4))
        if (((0.0 * 0.1) - 1.5) == 1.5) {
            v34 := GenS3{F0: uint16(v3), F1: ((v4 - (v4 << (((v3 | (v3 ^ (v3 << ((v3) & 63))))) & 63))) == (v4 >> ((uint(v4)) & 63))), F2: 1}
            var v35 int64 = genFn0((^(v3)), (^(v32)))
            _ = v34
            _ = v35
        }
        for fi36 := 0; fi36 < 12; fi36++ {
            var v37 int32 = 513
            switch (fi36 - fi33) {
            case 1:
                if len(v2) > 0 {
                    v2[0] = fi33
                }
            }
            if false {
                v32 = 85
                var v38 uint32 = (v5 << (((v3 / (((v3 + 18446744073709551615)) | 1))) & 63))
                _ = v38
            }
            _ = fi36
            _ = v37
        }
        _ = fi33
    }
    for fi39 := 0; fi39 < 14; fi39++ {
        for fi40 := 0; fi40 < 15; fi40++ {
            var v41 uint8 = (^(v4))
            _ = fi40
            _ = v41
        }
        for fi42 := 0; fi42 < 15; fi42++ {
            var v43 string = "hello"
            switch (fi39 % ((4294967295) | 1)) {
            case 187:
                var v44 float64 = v31
                _ = v44
            case -361:
                v45 := GenS4{F0: v6, F1: fi42}
                _ = v45
            }
            _ = fi42
            _ = v43
        }
        _ = fi39
    }
    var v46 int = -9223372036854775808
    v4 += v4
    switch (("\n\t" + (((("foo bar" + (("foo bar" + ("a" + ("\n\t" + "\n\t"))) + "a")) + "\n\t") + "foo bar") + (("foo bar" + "hello") + ""))) + "foo bar") {
    case "a":
        v47 := GenS5{F0: (v46 * 32767), F1: genFn0(541, uint16(v3))}
        _ = v47
    case "\n\t":
        switch v46 {
        case 710:
            v2 = append(v2, (v46 | -1))
        case 0:
            for fi48 := 0; fi48 < 10; fi48++ {
                fmt.Println(v32, v6, v31)
                v49 := GenS6{F0: int64(v5)}
                _ = fi48
                _ = v49
            }
        default:
            v4 = v4
        }
    default:
        if true {
            v50 := GenS5{F0: v46, F1: 9223372036854775807}
            v31 -= v31
            _ = v50
        }
    }
    v6 = v6
    for fi51 := 0; fi51 < 18; fi51++ {
        var v52 uint16 = 32767
        _ = len(v2)
        _ = fi51
        _ = v52
    }
    var v53 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v46
    _ = v5
    _ = v53
    _ = v6
}
