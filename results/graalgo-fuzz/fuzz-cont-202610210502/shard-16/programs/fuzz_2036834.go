package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 uint; F2 int }
type GenS5 struct { F0 uint16; F1 bool; F2 uint8 }
func genFn0(p0 int8, p1 uint8, p2 uint16) uint8 {
    var v5 uint64 = 4294967295
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 float64 = 0.1
        if false {
            var v8 uint16 = p2
            fmt.Println(p1)
            for fi9 := 0; fi9 < 9; fi9++ {
                switch (fi9 * fi9) {
                case -2147483648:
                    var v10 int8 = p0
                    _ = v10
                case -9223372036854775808:
                    p1 = p1
                    if true {
                        v11 := []int{fi9, fi9, -9223372036854775808}
                        _ = v11
                    } else {
                        fmt.Println(v5)
                        var v12 float64 = 1.5
                        _ = v12
                    }
                case -661:
                    v8 -= (p2 ^ (v8 % (((v8 + p2)) | 1)))
                }
                fmt.Println(p2, p0, fi9)
                switch "hello" {
                case "a":
                    switch fi9 {
                    case -905:
                        v13 := []int{fi9, fi9, int(p0)}
                        _ = v13
                    case 983:
                        if (!(!(!false))) {
                            v8 += 42
                            fmt.Println(fi6, fi9)
                            var v14 float32 = (-2.25 - float32(1.5))
                            _ = v14
                        }
                    case 1:
                        fmt.Println(p1, fi9)
                    }
                }
                _ = fi9
            }
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (p1 << ((0) & 63))
}
func genFn1(p0 int32, p1 int) int64 {
    for fi15 := 0; fi15 < 15; fi15++ {
        v16 := []int{(fi15 % ((fi15) | 1))}
        if false {
            switch "foo bar" {
            case "hello":
                if (193 != 0) {
                    var v17 float32 = 1e10
                    var v18 string = "hello"
                    _ = v17
                    _ = v18
                } else {
                    fmt.Println(p1)
                }
            case "\n\t":
                v19 := GenS1{F0: 2, F1: 152}
                _ = v19
            default:
                v20 := GenS2{F0: 313, F1: 42}
                _ = v20
            }
            _ = len(v16)
            v21 := GenS1{F0: 0, F1: genFn0(1, 128, 255)}
            _ = v21
        }
        var v22 uint8 = 0
        _ = fi15
        _ = v16
        _ = v22
    }
    var v23 int32 = (^(p0))
    _ = p0
    _ = p1
    _ = v23
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 926
    v3 := []int{}
    var v4 uint8 = 2
    v24 := GenS3{F0: 32767}
    var v25 int64 = 42
    var v26 float32 = 0.0
    fmt.Println(v26, v25)
    for fi27 := 0; fi27 < 15; fi27++ {
        v28 := GenS1{F0: 677, F1: v4}
        fmt.Println(fi27, v25)
        _ = fi27
        _ = v28
    }
    switch ("" + "") {
    case "":
        var v29 uint16 = 726
        _ = v29
    case "\n\t":
        for fi30 := 0; fi30 < 6; fi30++ {
            var v31 int32 = 0
            v32 := GenS1{F0: 2, F1: v4}
            _ = fi30
            _ = v31
            _ = v32
        }
    default:
        for fi33 := 0; fi33 < 1; fi33++ {
            if (true || ((false || false) || true)) {
                var v34 float32 = v26
                if ((fi33 > fi33) || false) {
                    v35 := GenS3{F0: -771}
                    fmt.Println(v25, v26, v34)
                    _ = v35
                }
                var v36 int16 = 42
                _ = v34
                _ = v36
            }
            if (0 < 42) {
                switch "foo bar" {
                case "\n\t":
                    var v37 float32 = ((v26 / 0.1) + (v26 * (v26 * 3.14159)))
                    _ = v37
                case "":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                switch -862 {
                case 42:
                    switch "a" {
                    case "":
                        v1["\n\t"] = ((v26 - (v26 + ((1e10 + float32(1.5)) - v26))) >= v26)
                    case "a":
                        switch "hello" {
                        case "a":
                            v3 = append(v3, (^(fi33)))
                        case "foo bar":
                            var v38 uint = 42
                            _ = v38
                        default:
                            fmt.Println(v4)
                        }
                    }
                default:
                    v25 = -9223372036854775808
                }
                if (!false) {
                    delete(v1, "hello")
                    v26 = 0.1
                } else {
                    fmt.Println(v26)
                    switch (fi33 ^ int(v4)) {
                    case 127:
                        var v39 int16 = 1000
                        _ = v39
                    default:
                        var v40 int64 = (v25 ^ (v25 >> ((uint(v4)) & 63)))
                        _ = v40
                    }
                }
            } else {
                if (true || (false || (true || false))) {
                    fmt.Println(v2, v25, fi33)
                    v4 -= v4
                } else {
                    v26 = (1.5 / v26)
                    var v41 float32 = ((v26 * float32(-2.25)) + v26)
                    _ = v41
                }
                fmt.Println(v26, v26, v26)
            }
            v4 = (v4 | (v4 + (^(v4))))
            _ = fi33
        }
    }
    for fi42 := 0; fi42 < 8; fi42++ {
        if false {
            fmt.Println(v25)
            if false {
                var v43 int32 = 770
                _ = v43
            } else {
                if false {
                    v26 -= v26
                    if true {
                        if true {
                            v4 -= (v4 | v4)
                            switch -1 {
                            case 0:
                                for fi44 := 0; fi44 < 9; fi44++ {
                                    var v45 float64 = 1.5
                                    v46 := GenS2{F0: 276, F1: (^(v25))}
                                    _ = fi44
                                    _ = v45
                                    _ = v46
                                }
                            case 42:
                                if (42 == -176) {
                                    v4 = v4
                                    for fi47 := 0; fi47 < 14; fi47++ {
                                        var v48 uint8 = v4
                                        fmt.Println(fi42)
                                        if (3.14159 <= float64(0.1)) {
                                            if true {
                                                v48 += genFn0(78, v4, 32767)
                                                v49 := GenS3{F0: 1000}
                                                v4 -= 42
                                                _ = v49
                                            } else {
                                                v50 := GenS4{F0: (^(v2)), F1: 255, F2: (fi42 % ((fi47) | 1))}
                                                v51 := v1[""]
                                                var v52 uint64 = 42
                                                _ = v50
                                                _ = v51
                                                _ = v52
                                            }
                                        } else {
                                            fmt.Println(fi47, fi42, v2)
                                        }
                                        v2 = (v2 >> ((uint(v4)) & 63))
                                        _ = fi47
                                        _ = v48
                                    }
                                }
                            case 4294967295:
                                if true {
                                    for fi53 := 0; fi53 < 12; fi53++ {
                                        var v54 uint = 443
                                        switch "a" {
                                        case "foo bar":
                                            var v55 float64 = 1.5
                                            _ = v55
                                        case "\n\t":
                                            v56 := GenS2{F0: 0, F1: (v25 / ((1) | 1))}
                                            _ = v56
                                        }
                                        _ = fi53
                                        _ = v54
                                    }
                                } else {
                                    if true {
                                        if false {
                                            var v57 uint = 575
                                            _ = v57
                                        } else {
                                            if (978 <= 42) {
                                                v3 = append(v3, (fi42 % (((^(fi42))) | 1)))
                                                var v58 uint8 = genFn0(50, 97, 32767)
                                                _ = v58
                                            } else {
                                                if (32767 <= 65535) {
                                                    for fi59 := 0; fi59 < 7; fi59++ {
                                                        v25 = (v25 | (^(v25)))
                                                        _ = fi59
                                                    }
                                                    var v60 uint = 795
                                                    _ = v60
                                                } else {
                                                    v2 = v2
                                                    for fi61 := 0; fi61 < 14; fi61++ {
                                                        v2 = v2
                                                        fmt.Println(v2, v4, fi42)
                                                        _ = fi61
                                                    }
                                                    switch "hello" {
                                                    case "hello":
                                                        switch 127 {
                                                        case -1:
                                                            if (v2 != 709) {
                                                                var v62 string = "\n\t"
                                                                _ = v62
                                                            } else {
                                                                v2 = v2
                                                            }
                                                        case 637:
                                                            var v63 uint16 = 349
                                                            _ = v63
                                                        case 1:
                                                            var v64 int64 = (v25 >> ((4294967295) & 63))
                                                            _ = v64
                                                        }
                                                    case "foo bar":
                                                        v4 = genFn0(57, (v4 ^ (v4 + v4)), 0)
                                                    case "":
                                                        switch (fi42 ^ -469) {
                                                        case 127:
                                                            if true {
                                                                var v65 string = ""
                                                                if len(v3) > 0 {
                                                                    v3[0] = (fi42 * fi42)
                                                                }
                                                                if ((v26 + 0.1) < v26) {
                                                                    fmt.Println(v26, fi42)
                                                                    var v66 bool = (!(v2 < (v2 + 2147483648)))
                                                                    var v67 float32 = v26
                                                                    _ = v66
                                                                    _ = v67
                                                                }
                                                                _ = v65
                                                            }
                                                        }
                                                    }
                                                }
                                                var v68 uint32 = (v2 / ((0) | 1))
                                                _ = v68
                                            }
                                        }
                                        v25 = genFn1(1, fi42)
                                        for fi69 := 0; fi69 < 8; fi69++ {
                                            v1[(("foo bar" + ("foo bar" + ("\n\t" + "\n\t"))) + "")] = true
                                            if true {
                                                var v70 bool = false
                                                if (true && (v25 > v25)) {
                                                    if v70 {
                                                        v70 = v70
                                                        v2 -= v2
                                                    }
                                                    var v71 float32 = (float32(-2.25) / 1.0)
                                                    _ = v71
                                                } else {
                                                    v70 = (4294967295 >= fi69)
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                                _ = v70
                                            }
                                            _ = fi69
                                        }
                                    } else {
                                        _ = len(v1)
                                        var v72 uint16 = 65535
                                        v72 = v72
                                        _ = v72
                                    }
                                }
                            }
                            var v73 float64 = (0.1 + float64(float32(float64(float32((float64(v26) - (float64((float32(0.1) + v26)) * 3.14159)))))))
                            _ = v73
                        } else {
                            if true {
                                switch fi42 {
                                case -255:
                                    fmt.Println(v2, v4)
                                default:
                                    for fi74 := 0; fi74 < 8; fi74++ {
                                        switch (^(fi74)) {
                                        case 265:
                                            var v75 int64 = (^(v25))
                                            _ = v75
                                        case 2147483647:
                                            v2 = v2
                                        default:
                                            var v76 string = "a"
                                            _ = v76
                                        }
                                        _ = fi74
                                    }
                                }
                            }
                            fmt.Println(v25, v26, v2)
                        }
                    }
                    for fi77 := 0; fi77 < 14; fi77++ {
                        v4 = 128
                        switch ("foo bar" + ("foo bar" + "\n\t")) {
                        case "":
                            v25 = int64(v2)
                        }
                        if false {
                            switch "foo bar" {
                            case "foo bar":
                                fmt.Println(fi42)
                            case "hello":
                                switch "" {
                                case "\n\t":
                                    fmt.Println(v26)
                                case "foo bar":
                                    fmt.Println(v25, v2, v25)
                                case "":
                                    if len(v3) > 0 {
                                        v3[0] = 32768
                                    }
                                default:
                                    var v78 uint16 = 660
                                    _ = v78
                                }
                            case "a":
                                v79 := GenS3{F0: int16(v4)}
                                _ = v79
                            }
                            v26 = v26
                            var v80 uint = 255
                            _ = v80
                        } else {
                            v26 = (v26 * 1e10)
                            if (-375 >= 715) {
                                fmt.Println(fi77, fi77)
                                var v81 int = fi77
                                var v82 int32 = 0
                                _ = v81
                                _ = v82
                            }
                            switch fi77 {
                            case -2147483648:
                                v83 := GenS5{F0: 2, F1: (fi42 > -9223372036854775808), F2: genFn0(-128, v4, 32767)}
                                _ = v83
                            case 9223372036854775807:
                                var v84 int16 = 1
                                _ = v84
                            default:
                                v85 := GenS5{F0: 42, F1: false, F2: v4}
                                _ = v85
                            }
                        }
                        _ = fi77
                    }
                }
                if true {
                    var v86 int8 = -1
                    var v87 int = (fi42 ^ int(v86))
                    fmt.Println(v26)
                    _ = v86
                    _ = v87
                } else {
                    if true {
                        fmt.Println(v26, v4)
                        v88 := GenS4{F0: (v2 << ((18446744073709551615) & 63)), F1: 61, F2: (fi42 * -746)}
                        var v89 int64 = v25
                        _ = v88
                        _ = v89
                    }
                    if false {
                        var v90 bool = true
                        delete(v1, "")
                        v91 := GenS2{F0: 1, F1: (-(v25))}
                        _ = v90
                        _ = v91
                    }
                    v26 = v26
                }
            }
            v92 := GenS5{F0: 778, F1: false, F2: uint8(v2)}
            _ = v92
        } else {
            var v93 uint32 = (v2 * (v2 & v2))
            switch ("" + "\n\t") {
            case "a":
                switch "\n\t" {
                case "hello":
                    v94 := GenS2{F0: 4294967295, F1: (-(v25))}
                    _ = v94
                }
            case "\n\t":
                v95 := GenS5{F0: 0, F1: (-2.25 <= (0.0 / (((0.1 + 3.14159) / 1.0) / -2.25))), F2: v4}
                _ = v95
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = v93
        }
        _ = fi42
    }
    var v96 uint64 = 0
    if ((!true) && true) {
        v26 -= 0.0
    }
    v25 = v25
    if (42 >= 42) {
        v97 := GenS4{F0: (v2 ^ (v2 % ((v2) | 1))), F1: 18446744073709551615, F2: int(v25)}
        var v98 bool = false
        _ = v97
        _ = v98
    }
    v4 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v96
}
