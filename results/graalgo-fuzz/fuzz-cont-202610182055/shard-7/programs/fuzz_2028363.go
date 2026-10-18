package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int64 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 int; F1 int64; F2 bool }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS8 struct { F0 int; F1 uint16; F2 uint64 }
type GenS9 struct { F0 uint16; F1 int16 }
func genFn0(p0 uint16, p1 bool) uint {
    v5 := []int{-1, 820}
    p0 = (p0 + 0)
    _ = p0
    _ = p1
    _ = v5
    return 65535
}
func genFn1(p0 uint16, p1 int8) uint {
    v6 := GenS1{F0: p1, F1: (!(9223372036854775807 >= 4294967295)), F2: -1}
    if (!true) {
        fmt.Println(p0)
        p1 -= 1
        switch "hello" {
        case "a":
            v7 := GenS1{F0: 1, F1: ((p1 % (((p1 - 100)) | 1)) < p1), F2: 9223372036854775807}
            _ = v7
        case "hello":
            for fi8 := 0; fi8 < 11; fi8++ {
                if (!true) {
                    var v9 float32 = -2.25
                    var v10 int64 = -472
                    var v11 uint32 = 0
                    _ = v10
                    _ = v11
                    _ = v9
                }
                _ = fi8
            }
        case "\n\t":
            for fi12 := 0; fi12 < 1; fi12++ {
                v13 := GenS1{F0: (p1 & 127), F1: true, F2: 24}
                switch 365 {
                case 354:
                    p0 -= (p0 / (((p0 >> ((2) & 63))) | 1))
                default:
                    p1 = p1
                }
                _ = fi12
                _ = v13
            }
        }
    } else {
        fmt.Println(p1)
        if false {
            p0 = 32767
            p0 -= (^(p0))
            v14 := []int{}
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return uint(p0)
}

func main() {
    v1 := []string{"foo bar", "\n\t", "foo bar", "foo bar"}
    v2 := []string{""}
    var v3 uint8 = 1
    var v4 int32 = 164
    if false {
        if (298 != 0) {
            v4 = v4
            for fi15 := 0; fi15 < 15; fi15++ {
                fmt.Println(v4, fi15, v4)
                v3 = 1
                _ = fi15
            }
            v16 := GenS2{F0: 127, F1: -680}
            _ = v16
        } else {
            v17 := GenS3{F0: int(v4), F1: -1, F2: (!(!false))}
            v18 := GenS4{F0: (!(-949 == -9223372036854775808)), F1: 624}
            _ = v17
            _ = v18
        }
        fmt.Println(v3, v3)
        var v19 float64 = 0.0
        _ = v19
    } else {
        for fi20 := 0; fi20 < 9; fi20++ {
            v21 := GenS2{F0: 77, F1: (fi20 * (^(fi20)))}
            v22 := GenS5{F0: (!("hello" != (("a" + "\n\t") + ("foo bar" + (("\n\t" + "\n\t") + "\n\t"))))), F1: 1}
            v23 := GenS1{F0: 42, F1: (32767 <= 136), F2: -9223372036854775808}
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
    }
    v24 := GenS4{F0: (-2.25 == 1e10), F1: 1}
    v4 = -2147483648
    v25 := GenS4{F0: (4294967295 >= 306), F1: 42}
    if false {
        for fi26 := 0; fi26 < 6; fi26++ {
            if (false || true) {
                v3 -= v3
                for fi27 := 0; fi27 < 14; fi27++ {
                    switch -9223372036854775808 {
                    case 2147483647:
                        if true {
                            if ((0 == 18446744073709551615) || (!false)) {
                                if len(v1) > 0 {
                                    v1[0] = "\n\t"
                                }
                                var v28 int8 = 42
                                if len(v1) > 0 {
                                    v1[0] = "a"
                                }
                                _ = v28
                            } else {
                                switch ("\n\t" + (((("a" + "\n\t") + "") + "a") + ("foo bar" + "a"))) {
                                case "":
                                    switch 32767 {
                                    case -2147483648:
                                        _ = len(v1)
                                    default:
                                        v3 += v3
                                    }
                                case "\n\t":
                                    fmt.Println(fi27, fi26, fi26)
                                case "hello":
                                    fmt.Println(fi27)
                                default:
                                    var v29 int16 = 1
                                    _ = v29
                                }
                            }
                        }
                    }
                    fmt.Println(fi27)
                    var v30 float32 = 3.14159
                    _ = fi27
                    _ = v30
                }
                v1 = append(v1, "\n\t")
            } else {
                var v31 int32 = -1
                _ = v31
            }
            _ = fi26
        }
        var v32 int64 = 975
        for fi33 := 0; fi33 < 1; fi33++ {
            if (("" + "a") < "\n\t") {
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
                fmt.Println(v3, fi33, v32)
            } else {
                v1 = append(v1, ("\n\t" + "\n\t"))
                for fi34 := 0; fi34 < 8; fi34++ {
                    if true {
                        v35 := GenS2{F0: -128, F1: (^(fi33))}
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        fmt.Println(fi34, fi34, fi33)
                        _ = v35
                    } else {
                        v1 = append(v1, "\n\t")
                    }
                    _ = fi34
                }
            }
            _ = fi33
        }
        _ = v32
    } else {
        v3 = 127
        var v36 uint = 65535
        _ = v36
    }
    var v37 uint32 = 1
    var v38 uint16 = 32767
    v38 -= 2
    for fi39 := 0; fi39 < 5; fi39++ {
        if true {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v40 uint32 = v37
            for fi41 := 0; fi41 < 14; fi41++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi41
            }
            _ = v40
        }
        for fi42 := 0; fi42 < 19; fi42++ {
            v38 = v38
            _ = fi42
        }
        if false {
            v3 = 128
            for fi43 := 0; fi43 < 15; fi43++ {
                v37 = (v37 & (v37 >> ((18446744073709551615) & 63)))
                switch "foo bar" {
                case "\n\t":
                    var v44 int64 = 42
                    _ = v44
                case "hello":
                    v37 = v37
                case "foo bar":
                    if len(v1) > 0 {
                        v1[0] = ""
                    }
                default:
                    if len(v2) > 0 {
                        v2[0] = ("\n\t" + "")
                    }
                }
                if true {
                    v45 := GenS2{F0: 0, F1: fi43}
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(fi39)
                    case "a":
                        if (false && (!(false && false))) {
                            if false {
                                if true {
                                    v46 := GenS6{F0: v38}
                                    v4 = 42
                                    _ = v46
                                }
                                var v47 uint64 = 18446744073709551615
                                _ = v47
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            fmt.Println(v3, v38)
                        }
                    case "\n\t":
                        v4 = v4
                    default:
                        fmt.Println(v38, v38, v3)
                    }
                    var v48 string = "foo bar"
                    _ = v45
                    _ = v48
                } else {
                    if (128 == (^(v3))) {
                        switch "foo bar" {
                        case "foo bar":
                            var v49 int16 = 0
                            _ = v49
                        case "\n\t":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        var v50 int64 = -1
                        var v51 uint = 189
                        _ = v50
                        _ = v51
                    } else {
                        if ((float64(float32(float64(float32(3.14159)))) <= (float64(1e10) * 1e10)) || false) {
                            if false {
                                switch fi39 {
                                case -9223372036854775808:
                                    v37 -= (v37 + v37)
                                case -1:
                                    if len(v1) > 0 {
                                        v1[0] = ""
                                    }
                                case -2147483648:
                                    switch "" {
                                    case "foo bar":
                                        var v52 float32 = 1.5
                                        var v53 int32 = v4
                                        _ = v52
                                        _ = v53
                                    case "a":
                                        var v54 string = "\n\t"
                                        _ = v54
                                    case "":
                                        fmt.Println(fi39, fi39, v4)
                                    default:
                                        if false {
                                            if (255 > v3) {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v55 int = fi39
                                                    _ = v55
                                                case "hello":
                                                    v56 := GenS5{F0: (!false), F1: 0}
                                                    _ = v56
                                                case "\n\t":
                                                    if (false || true) {
                                                        switch -940 {
                                                        case 32767:
                                                            if (false || ((v37 << ((4294967295) & 63)) < (^(v37)))) {
                                                                v3 -= (v3 / ((v3) | 1))
                                                                v4 += (v4 + (v4 << ((genFn1((v38 >> ((1) & 63)), -112)) & 63)))
                                                            } else {
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                v38 -= (^(v38))
                                                                v3 += (v3 | (v3 >> ((4294967295) & 63)))
                                                            }
                                                        case 128:
                                                            var v57 float64 = 1e10
                                                            _ = v57
                                                        }
                                                        v58 := GenS7{F0: v37, F1: v37, F2: 844}
                                                        var v59 int64 = 1
                                                        _ = v58
                                                        _ = v59
                                                    }
                                                }
                                                v37 = uint32(v3)
                                            } else {
                                                var v60 int = (fi39 - (fi43 ^ (fi43 << ((2) & 63))))
                                                _ = v60
                                            }
                                        }
                                    }
                                }
                            } else {
                                v3 += 255
                                if (1 < (v37 ^ (v37 * v37))) {
                                    v1 = append(v1, (("\n\t" + "hello") + ""))
                                    v4 = 1
                                    v37 = (v37 % ((v37) | 1))
                                } else {
                                    fmt.Println(v4, v4)
                                    v4 -= 42
                                    if ((v38 % (((v38 ^ 1)) | 1)) <= (v38 % (((v38 * v38)) | 1))) {
                                        v61 := GenS7{F0: (v37 % ((0) | 1)), F1: (v37 | (v37 ^ (v37 % (((v37 ^ (v37 >> ((373) & 63)))) | 1)))), F2: 2}
                                        _ = v61
                                    } else {
                                        var v62 string = "a"
                                        _ = v62
                                    }
                                }
                            }
                            switch ("a" + "foo bar") {
                            case "":
                                var v63 uint16 = 252
                                _ = v63
                            case "\n\t":
                                v3 = 42
                            case "hello":
                                fmt.Println(v38, v37, v4)
                                if (true && (1.5 <= ((3.14159 * 0.0) / float32(0.1)))) {
                                    v38 += (v38 >> ((4294967295) & 63))
                                    v3 -= v3
                                } else {
                                    v1 = append(v1, "hello")
                                    fmt.Println(fi39, v37)
                                    switch "foo bar" {
                                    case "hello":
                                        if false {
                                            switch ("" + "hello") {
                                            case "\n\t":
                                                v37 = (v37 * v37)
                                            case "a":
                                                var v64 int32 = (v4 << ((803) & 63))
                                                _ = v64
                                            }
                                            if true {
                                                var v65 uint64 = 1
                                                _ = len(v1)
                                                v37 -= (v37 >> ((0) & 63))
                                                _ = v65
                                            }
                                            fmt.Println(v38)
                                            v3 = (v3 & v3)
                                            var v66 int16 = 794
                                            v1 = append(v1, "hello")
                                            _ = v66
                                        }
                                    case "":
                                        v3 -= 2
                                    case "foo bar":
                                        var v67 float64 = (float64(-2.25) - 0.1)
                                        _ = v67
                                    }
                                }
                            }
                            if (!true) {
                                v37 += (v37 * v37)
                            } else {
                                v1 = append(v1, "hello")
                                if false {
                                    fmt.Println(v3, v38, v3)
                                }
                            }
                        } else {
                            if (725 == 42) {
                                v37 = v37
                            } else {
                                v3 = 2
                                v68 := GenS6{F0: 1000}
                                v4 -= 2147483647
                                _ = v68
                            }
                            if (((v37 ^ v37) <= (^(v37))) || true) {
                                v4 = v4
                            }
                            var v69 string = ("foo bar" + ("" + "foo bar"))
                            v3 -= v3
                            if len(v1) > 0 {
                                v1[0] = "\n\t"
                            }
                            _ = v69
                        }
                        v70 := GenS8{F0: (fi39 + -2147483648), F1: 42, F2: 969}
                        switch "a" {
                        case "\n\t":
                            v38 = 32767
                        case "foo bar":
                            if (1.5 > float64(-2.25)) {
                                fmt.Println(v37)
                                var v71 uint = 1
                                _ = v71
                            } else {
                                v3 = v3
                            }
                        case "hello":
                            switch "hello" {
                            case "":
                                var v72 float32 = 0.0
                                _ = v72
                            case "a":
                                v38 = v38
                            case "\n\t":
                                var v73 uint64 = 1
                                _ = v73
                            }
                        default:
                            var v74 uint = 606
                            _ = v74
                        }
                        _ = v70
                    }
                }
                _ = fi43
            }
            var v75 uint8 = 2
            _ = v75
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = fi39
    }
    if false {
        v3 = v3
    }
    var v76 int8 = 0
    v3 = v3
    fmt.Println(v38)
    var v77 uint64 = 376
    var v78 float32 = 1e10
    for fi79 := 0; fi79 < 17; fi79++ {
        v80 := GenS9{F0: 42, F1: 409}
        for fi81 := 0; fi81 < 12; fi81++ {
            var v82 int = fi79
            _ = fi81
            _ = v82
        }
        v38 = (v38 & (v38 << ((883) & 63)))
        _ = fi79
        _ = v80
    }
    var v83 bool = (float64(0.1) <= 0.0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v76
    _ = v77
    _ = v78
    _ = v83
}
