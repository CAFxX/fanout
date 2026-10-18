package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 uint; F1 string; F2 uint }
type GenS4 struct { F0 int32; F1 int; F2 int16 }
type GenS5 struct { F0 int; F1 int16; F2 uint32 }
type GenS6 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS7 struct { F0 int8; F1 int; F2 int32 }
type GenS8 struct { F0 uint16; F1 int8; F2 int32 }
type GenS9 struct { F0 int64; F1 uint; F2 bool }
type GenS10 struct { F0 int64; F1 int64 }
type GenS11 struct { F0 uint64; F1 bool; F2 int16 }
type GenS12 struct { F0 uint8; F1 string }
func genFn0(p0 int16) uint {
    var v7 uint32 = 2
    v8 := []int{-2147483648, 704, -389}
    for fi9 := 0; fi9 < 12; fi9++ {
        p0 += p0
        v10 := GenS1{F0: "hello"}
        if (!((!false) || (-85 <= -1))) {
            switch "" {
            case "":
                v7 -= (v7 - (v7 ^ v7))
            case "\n\t":
                p0 = (p0 / (((p0 >> ((65535) & 63))) | 1))
            case "hello":
                for fi11 := 0; fi11 < 1; fi11++ {
                    fmt.Println(fi11)
                    _ = fi11
                }
            default:
                v7 = 42
            }
        }
        _ = fi9
        _ = v10
    }
    v8 = append(v8, 32767)
    fmt.Println(v7)
    _ = p0
    _ = v7
    _ = v8
    return 68
}
func genFn1(p0 int32) int8 {
    v12 := GenS2{F0: 0, F1: 1}
    p0 = (-(p0))
    v13 := GenS2{F0: 42, F1: 1}
    for fi14 := 0; fi14 < 3; fi14++ {
        for fi15 := 0; fi15 < 9; fi15++ {
            v16 := []int{fi15, (fi14 | fi14), (fi14 ^ (fi15 >> ((745) & 63)))}
            if (1 != 1) {
                if (991 >= 2) {
                    var v17 string = ""
                    switch (fi14 - (fi14 & fi14)) {
                    case 42:
                        switch 2 {
                        case 32767:
                            v16 = append(v16, (fi14 / ((9223372036854775807) | 1)))
                        case -862:
                            switch (fi14 - (fi15 % ((fi14) | 1))) {
                            case 2147483647:
                                switch "\n\t" {
                                case "":
                                    _ = len(v16)
                                case "hello":
                                    switch ("a" + v17) {
                                    case "hello":
                                        v17 = ((("a" + "\n\t") + v17) + "\n\t")
                                        fmt.Println(fi15, v17, p0)
                                        if (42 == 191) {
                                            if len(v16) >= 2 {
                                                v16 = v16[1:len(v16)]
                                            }
                                            v17 = v17
                                            var v18 uint8 = 192
                                            _ = v18
                                        }
                                        fmt.Println(fi15)
                                    case "":
                                        v16 = append(v16, 42)
                                    default:
                                        v19 := GenS3{F0: 598, F1: ((v17 + "") + (v17 + "a")), F2: 2}
                                        _ = v19
                                    }
                                case "\n\t":
                                    if true {
                                        v20 := GenS2{F0: 531, F1: 570}
                                        _ = v20
                                    } else {
                                        if len(v16) >= 2 {
                                            v16 = v16[1:len(v16)]
                                        }
                                        switch v17 {
                                        case "a":
                                            p0 = -883
                                        }
                                        var v21 uint16 = 255
                                        _ = v21
                                    }
                                default:
                                    switch (^(fi14)) {
                                    case 0:
                                        var v22 uint32 = 0
                                        _ = v22
                                    case 1:
                                        p0 = p0
                                    }
                                }
                            case 534:
                                v23 := GenS3{F0: 1, F1: (v17 + v17), F2: 142}
                                _ = v23
                            case 4294967295:
                                var v24 int8 = 127
                                _ = v24
                            default:
                                p0 = p0
                            }
                        case 32768:
                            var v25 int16 = -1
                            switch v17 {
                            case "foo bar":
                                switch fi15 {
                                case 32767:
                                    var v26 int = 9223372036854775807
                                    _ = v26
                                }
                            default:
                                fmt.Println(v17)
                            }
                            _ = v25
                        }
                        if false {
                            p0 = (p0 * -2147483648)
                            v27 := GenS2{F0: 2, F1: 983}
                            _ = v27
                        } else {
                            _ = len(v16)
                        }
                        p0 = p0
                    default:
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                    }
                    v17 = "hello"
                    _ = v17
                }
                if (true && (1e10 < float32(1e10))) {
                    if (-2.25 < float32((0.0 * float64(3.14159)))) {
                        switch "foo bar" {
                        case "hello":
                            v28 := GenS4{F0: (p0 / ((p0) | 1)), F1: (fi15 ^ fi14), F2: 456}
                            _ = v28
                        case "a":
                            fmt.Println(fi14, fi15, fi14)
                        case "foo bar":
                            var v29 int32 = p0
                            var v30 uint64 = 81
                            if true {
                                fmt.Println(p0, v30, v29)
                            } else {
                                v30 += v30
                            }
                            _ = v29
                            _ = v30
                        default:
                            switch "hello" {
                            case "foo bar":
                                p0 = int32(fi14)
                            case "a":
                                var v31 int16 = 32767
                                _ = v31
                            }
                        }
                        var v32 float64 = 0.0
                        var v33 float32 = -2.25
                        _ = v32
                        _ = v33
                    } else {
                        p0 = (p0 | p0)
                        v34 := GenS3{F0: genFn0(42), F1: "foo bar", F2: 42}
                        _ = v34
                    }
                    switch (fi14 - fi15) {
                    case 7:
                        if ((p0 | -1) == p0) {
                            p0 -= p0
                            switch (-(fi15)) {
                            case 424:
                                if true {
                                    v35 := GenS5{F0: (^(fi14)), F1: 642, F2: 42}
                                    var v36 string = ("a" + "hello")
                                    _ = v35
                                    _ = v36
                                }
                            case 4294967295:
                                if (0 < -9223372036854775808) {
                                    var v37 uint64 = 18446744073709551615
                                    p0 = -1
                                    _ = v37
                                }
                            }
                        } else {
                            if ((int16(fi15) > -773) || ((!(!true)) && (42 >= 80))) {
                                p0 = (p0 & 1)
                                switch ("a" + "a") {
                                case "hello":
                                    if len(v16) >= 2 {
                                        v16 = v16[1:len(v16)]
                                    }
                                case "":
                                    v38 := GenS2{F0: 4294967295, F1: 2}
                                    _ = v38
                                default:
                                    p0 = p0
                                }
                                var v39 uint64 = 2
                                _ = v39
                            } else {
                                switch "" {
                                case "foo bar":
                                    if false {
                                        switch fi14 {
                                        case 851:
                                            v16 = append(v16, (fi14 | (fi14 | fi15)))
                                        case -2147483648:
                                            var v40 int16 = -869
                                            _ = v40
                                        default:
                                            fmt.Println(p0, fi14)
                                        }
                                    }
                                    p0 -= (p0 << ((65535) & 63))
                                    var v41 uint32 = 0
                                    if len(v16) >= 2 {
                                        v16 = v16[1:len(v16)]
                                    }
                                    switch (fi14 / ((fi14) | 1)) {
                                    case 162:
                                        v42 := GenS5{F0: fi15, F1: -32768, F2: (v41 & v41)}
                                        _ = v42
                                    case 128:
                                        v43 := GenS4{F0: -2147483648, F1: fi15, F2: 42}
                                        _ = v43
                                    case 42:
                                        v41 += v41
                                    }
                                    if (!(1 != 42)) {
                                        var v44 int32 = (p0 << ((951) & 63))
                                        var v45 int64 = -1
                                        v16 = append(v16, (fi14 << ((2) & 63)))
                                        _ = v44
                                        _ = v45
                                    } else {
                                        switch (fi15 - fi14) {
                                        case -1:
                                            v41 -= (v41 | v41)
                                        }
                                        p0 = p0
                                        v46 := GenS1{F0: (("a" + "hello") + ("a" + "hello"))}
                                        v47 := GenS6{F0: 42, F1: uint(v41), F2: 215}
                                        _ = v46
                                        _ = v47
                                    }
                                    var v48 bool = ((1 != 0) || (92 == 132))
                                    v49 := GenS5{F0: (fi15 >> ((371) & 63)), F1: 1000, F2: (v41 + (^(v41)))}
                                    if (v48 && ((1.5 <= 1e10) && (1 != 42))) {
                                        v50 := GenS4{F0: p0, F1: -2147483648, F2: 32767}
                                        switch "" {
                                        case "hello":
                                            v51 := GenS7{F0: 42, F1: fi14, F2: p0}
                                            _ = v51
                                        case "a":
                                            var v52 float64 = 1e10
                                            _ = v52
                                        }
                                        if (!true) {
                                            v53 := GenS8{F0: 856, F1: -33, F2: (p0 << ((18446744073709551615) & 63))}
                                            var v54 uint = genFn0(1000)
                                            var v55 uint = (v54 / ((v54) | 1))
                                            _ = v53
                                            _ = v54
                                            _ = v55
                                        } else {
                                            var v56 int32 = p0
                                            v48 = false
                                            _ = v56
                                        }
                                        _ = v50
                                    }
                                    fmt.Println(p0, v41, fi15)
                                    _ = v41
                                    _ = v48
                                    _ = v49
                                case "a":
                                    var v57 float32 = 3.14159
                                    _ = v57
                                case "hello":
                                    fmt.Println(fi14)
                                }
                                var v58 int16 = 32767
                                var v59 int64 = int64(p0)
                                _ = v58
                                _ = v59
                            }
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                            fmt.Println(fi14)
                        }
                    case 41:
                        v60 := GenS9{F0: 9223372036854775807, F1: 42, F2: false}
                        _ = v60
                    case 42:
                        if true {
                            if (((fi14 / ((782) | 1)) > fi14) && false) {
                                v61 := GenS9{F0: 0, F1: 0, F2: true}
                                switch fi15 {
                                case 265:
                                    var v62 int8 = 100
                                    _ = v62
                                case 248:
                                    fmt.Println(fi15)
                                case -649:
                                    switch "hello" {
                                    case "foo bar":
                                        fmt.Println(fi15, fi14, fi14)
                                    case "a":
                                        if ((float64(0.0) != 0.1) && false) {
                                            if len(v16) > 0 {
                                                v16[0] = 1
                                            }
                                            p0 = (p0 % ((p0) | 1))
                                        } else {
                                            if len(v16) > 0 {
                                                v16[0] = fi15
                                            }
                                            _ = len(v16)
                                            v63 := GenS10{F0: 42, F1: -821}
                                            _ = v63
                                        }
                                    case "hello":
                                        switch "a" {
                                        case "\n\t":
                                            if len(v16) >= 2 {
                                                v16 = v16[1:len(v16)]
                                            }
                                        }
                                    }
                                default:
                                    var v64 float32 = -2.25
                                    _ = v64
                                }
                                fmt.Println(p0)
                                _ = v61
                            } else {
                                v65 := GenS6{F0: 0, F1: 4294967295, F2: 1}
                                var v66 uint16 = 0
                                if (!(0 == genFn0(720))) {
                                    v66 = 255
                                    v66 = 341
                                } else {
                                    if (!(!(2 != 18446744073709551615))) {
                                        v16 = append(v16, int(v66))
                                    }
                                    v66 = (v66 & v66)
                                }
                                _ = v65
                                _ = v66
                            }
                            p0 = 2147483647
                        }
                        var v67 int = int(p0)
                        switch fi15 {
                        case -174:
                            p0 = p0
                        default:
                            p0 += p0
                        }
                        _ = v67
                    }
                }
            } else {
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                _ = len(v16)
            }
            switch "foo bar" {
            case "\n\t":
                var v68 bool = (!(((3.14159 + float32(3.14159)) < 0.0) && (405 >= 4294967295)))
                _ = v68
            }
            var v69 int64 = -9223372036854775808
            v70 := GenS6{F0: 94, F1: 843, F2: 18446744073709551615}
            _ = fi15
            _ = v16
            _ = v69
            _ = v70
        }
        _ = fi14
    }
    v71 := []int{}
    _ = p0
    _ = v12
    _ = v13
    _ = v71
    return -128
}

func main() {
    var v1 int = 1
    var v2 int8 = -1
    var v3 uint32 = 381
    var v4 uint16 = 841
    var v5 int16 = 32767
    v6 := []int{}
    var v72 uint64 = 2
    var v73 float64 = (-2.25 / (1.5 / (1.5 - -2.25)))
    var v74 int16 = v5
    _ = len(v6)
    fmt.Println(v74, v4)
    for fi75 := 0; fi75 < 6; fi75++ {
        v4 += 699
        switch "a" {
        case "hello":
            if (42 <= -1) {
                if true {
                    var v76 bool = (false && (v4 <= 336))
                    _ = v76
                }
                var v77 float32 = -2.25
                _ = v77
            }
        default:
            _ = len(v6)
        }
        _ = fi75
    }
    v78 := GenS11{F0: (v72 & (v72 % ((v72) | 1))), F1: (true || ((0 == -9223372036854775808) && (!(-602 != 42)))), F2: v74}
    v3 -= (v3 % ((829) | 1))
    for fi79 := 0; fi79 < 20; fi79++ {
        var v80 int32 = -2147483648
        switch 128 {
        case 594:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
        _ = fi79
        _ = v80
    }
    var v81 int32 = 972
    v74 = v74
    v82 := GenS8{F0: v4, F1: (v2 | v2), F2: v81}
    v3 = 582
    for fi83 := 0; fi83 < 13; fi83++ {
        for fi84 := 0; fi84 < 19; fi84++ {
            fmt.Println(fi84, fi83)
            _ = fi84
        }
        switch "\n\t" {
        case "":
            fmt.Println(v1)
        case "foo bar":
            var v85 int16 = -32768
            _ = v85
        case "\n\t":
            switch "\n\t" {
            case "":
                if (!false) {
                    v81 = v81
                }
            case "hello":
                var v86 uint64 = 0
                _ = v86
            case "a":
                for fi87 := 0; fi87 < 17; fi87++ {
                    if ("" == "") {
                        _ = len(v6)
                        var v88 bool = (false || false)
                        var v89 int32 = (v81 * v81)
                        _ = v88
                        _ = v89
                    }
                    _ = fi87
                }
            }
        }
        _ = fi83
    }
    switch v1 {
    case 586:
        v5 += -792
    case 559:
        v90 := GenS12{F0: 32, F1: "foo bar"}
        _ = v90
    case 1:
        var v91 bool = ((18446744073709551615 == 255) && false)
        _ = v91
    default:
        v74 -= 1
    }
    var v92 uint16 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    fmt.Println("v81:", v81)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v72
    _ = v73
    _ = v74
    _ = v78
    _ = v81
    _ = v82
    _ = v92
}
