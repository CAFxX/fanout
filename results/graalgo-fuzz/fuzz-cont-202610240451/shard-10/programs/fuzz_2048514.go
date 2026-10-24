package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 int; F2 string }
type GenS4 struct { F0 int32; F1 int; F2 uint8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64; F1 bool; F2 bool }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int8; F1 uint8 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int8; F1 int32; F2 int32 }
type GenS11 struct { F0 bool }
func genFn0(p0 uint32, p1 uint8, p2 uint8) uint16 {
    if (1 >= uint(p0)) {
        var v4 int16 = 1
        switch 0 {
        case 1:
            var v5 int8 = 100
            _ = v5
        case 128:
            var v6 uint16 = 65535
            _ = v6
        default:
            p0 = p0
        }
        _ = v4
    }
    v7 := []int{-2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    var v1 uint8 = 1
    var v2 int64 = 1
    v3 := make(map[uint32]string)
    fmt.Println(v2, v2)
    for fi8 := 0; fi8 < 18; fi8++ {
        if true {
            for fi9 := 0; fi9 < 17; fi9++ {
                v10 := GenS1{F0: -989}
                _ = fi9
                _ = v10
            }
            v11 := []int{fi8}
            _ = v11
        } else {
            switch fi8 {
            case 32767:
                fmt.Println(v1)
            }
        }
        if (0 <= 42) {
            var v12 uint64 = 4294967295
            var v13 int32 = -1
            v14 := GenS2{F0: -1}
            _ = v12
            _ = v13
            _ = v14
        } else {
            if (float64((1e10 * 0.1)) >= ((3.14159 * 0.1) * -2.25)) {
                v15 := []int{(fi8 << ((933) & 63)), (fi8 / (((fi8 >> ((42) & 63))) | 1)), (fi8 * int(v1))}
                switch (-(fi8)) {
                case 42:
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                case -839:
                    var v16 uint32 = 443
                    _ = v16
                case 2147483647:
                    if len(v15) > 0 {
                        v15[0] = (-(fi8))
                    }
                }
                _ = v15
            } else {
                v17 := v3[865]
                _ = v17
            }
        }
        _ = fi8
    }
    if ("" >= ("\n\t" + "\n\t")) {
        for fi18 := 0; fi18 < 8; fi18++ {
            v1 += (v1 >> ((1) & 63))
            var v19 uint32 = 65535
            for fi20 := 0; fi20 < 14; fi20++ {
                delete(v3, (v19 % ((v19) | 1)))
                fmt.Println(v2, fi20, fi18)
                var v21 int = (fi18 & 125)
                var v22 uint8 = (v1 ^ (v1 >> ((1) & 63)))
                var v23 uint32 = 2147483647
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
            _ = fi18
            _ = v19
        }
    }
    if (((158 < 65535) || ((127 > -192) && true)) || (42 > 125)) {
        v1 = v1
        var v24 float32 = 0.0
        _ = v24
    } else {
        var v25 uint32 = 1
        if (1.5 >= (0.1 * (-2.25 - 0.0))) {
            if (false && ((v2 & v2) > v2)) {
                fmt.Println(v1, v2, v2)
                v26 := GenS2{F0: -1}
                if false {
                    fmt.Println(v2, v2)
                    switch -1 {
                    case -2147483648:
                        v25 = 360
                    }
                } else {
                    v27 := []int{-150, 757}
                    _ = v27
                }
                _ = v26
            }
            var v28 int8 = 127
            _ = v28
        }
        _ = v25
    }
    switch "" {
    case "hello":
        v29 := GenS1{F0: 886}
        _ = v29
    }
    v1 = v1
    for fi30 := 0; fi30 < 12; fi30++ {
        var v31 uint32 = 60
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 20; fi32++ {
        if (4294967295 > 4294967295) {
            v33 := GenS3{F0: 1, F1: 128, F2: (("hello" + "foo bar") + "hello")}
            _ = v33
        }
        var v34 float64 = 1.5
        v34 += ((v34 + ((v34 / (v34 - ((v34 + 3.14159) * -2.25))) + (v34 - v34))) - v34)
        _ = fi32
        _ = v34
    }
    if (65535 != 65535) {
        v35 := v3[992]
        if (2 != 42) {
            v2 = (-(v2))
            var v36 uint64 = uint64(v1)
            _ = v36
        }
        var v37 uint32 = 1
        _ = v35
        _ = v37
    } else {
        for fi38 := 0; fi38 < 17; fi38++ {
            if true {
                for fi39 := 0; fi39 < 14; fi39++ {
                    v1 -= (^(v1))
                    _ = len(v3)
                    v1 -= v1
                    var v40 int8 = -94
                    var v41 uint = 18446744073709551615
                    switch fi39 {
                    case 2147483647:
                        fmt.Println(fi38)
                    case 241:
                        if true {
                            v40 += v40
                            v42 := GenS4{F0: 1, F1: fi38, F2: 127}
                            _ = v42
                        }
                    default:
                        switch "" {
                        case "foo bar":
                            v43 := GenS2{F0: -10}
                            _ = v43
                        case "hello":
                            fmt.Println(fi38)
                        case "a":
                            var v44 int8 = 42
                            _ = v44
                        }
                    }
                    fmt.Println(v41, v2, v41)
                    _ = fi39
                    _ = v40
                    _ = v41
                }
                v2 = v2
            } else {
                v45 := []int{(^(fi38)), fi38, (fi38 ^ int(v2)), -9223372036854775808}
                var v46 int16 = 285
                _ = v45
                _ = v46
            }
            v1 -= v1
            _ = fi38
        }
    }
    v47 := []int{}
    v1 -= v1
    v2 = -1
    if false {
        switch -898 {
        case -156:
            switch 256 {
            case 0:
                fmt.Println(v2, v2)
            }
        }
        v1 = v1
        if true {
            v1 += 255
            v2 += (v2 ^ -1)
        }
    } else {
        if false {
            for fi48 := 0; fi48 < 19; fi48++ {
                for fi49 := 0; fi49 < 19; fi49++ {
                    if (-1 <= 127) {
                        fmt.Println(fi48, fi49)
                        switch "foo bar" {
                        case "foo bar":
                            var v50 int = fi49
                            _ = v50
                        case "a":
                            var v51 uint16 = 2
                            _ = v51
                        case "\n\t":
                            switch "foo bar" {
                            case "\n\t":
                                v52 := GenS5{F0: (fi49 | fi48)}
                                _ = v52
                            case "foo bar":
                                v2 = (v2 | (v2 * (v2 >> ((18446744073709551615) & 63))))
                            case "a":
                                if true {
                                    v1 += (v1 % ((2) | 1))
                                    var v53 uint16 = 255
                                    if len(v47) > 0 {
                                        v47[0] = fi48
                                    }
                                    _ = v53
                                }
                            default:
                                v1 = v1
                            }
                        default:
                            var v54 uint64 = uint64(v1)
                            _ = v54
                        }
                    } else {
                        v2 = v2
                        switch (fi49 | fi49) {
                        case 608:
                            switch 127 {
                            case 2147483647:
                                var v55 uint16 = 255
                                _ = v55
                            case 128:
                                v56 := GenS3{F0: 0, F1: (fi48 % ((fi48) | 1)), F2: "\n\t"}
                                v1 = v1
                                v1 = (v1 % ((v1) | 1))
                                _ = v56
                            }
                        case -2147483648:
                            if false {
                                if false {
                                    switch ((((((((("\n\t" + ("" + "\n\t")) + "hello") + "hello") + "foo bar") + "a") + "") + "a") + "\n\t") + "a") {
                                    case "\n\t":
                                        fmt.Println(v2, fi48)
                                        var v57 uint8 = (^(v1))
                                        _ = v57
                                    case "":
                                        v3[2] = "hello"
                                    case "foo bar":
                                        if len(v47) > 0 {
                                            v47[0] = 687
                                        }
                                    }
                                    v58 := GenS6{F0: 18446744073709551615, F1: false, F2: (((-2.25 / -2.25) / -2.25) == 0.1)}
                                    v2 += (v2 | (v2 & v2))
                                    _ = v58
                                } else {
                                    var v59 uint8 = (v1 & v1)
                                    var v60 int32 = -699
                                    fmt.Println(v59, fi48, v60)
                                    var v61 uint32 = 65535
                                    switch (fi49 - 4294967295) {
                                    case 32767:
                                        v1 = v1
                                    default:
                                        if false {
                                            fmt.Println(fi49, v59)
                                        } else {
                                            var v62 float64 = 3.14159
                                            switch fi49 {
                                            case 644:
                                                var v63 uint8 = (^(v1))
                                                _ = v63
                                            default:
                                                v64 := GenS7{F0: 42}
                                                _ = v64
                                            }
                                            var v65 bool = (!false)
                                            _ = v62
                                            _ = v65
                                        }
                                    }
                                    _ = v59
                                    _ = v60
                                    _ = v61
                                }
                                switch (-(fi49)) {
                                case 42:
                                    v66 := GenS4{F0: 790, F1: fi48, F2: (v1 & v1)}
                                    _ = v66
                                case 878:
                                    var v67 uint = 4294967295
                                    switch 32768 {
                                    case 2147483647:
                                        v1 = (v1 % ((128) | 1))
                                    default:
                                        v1 += v1
                                    }
                                    switch (("a" + "a") + ("\n\t" + "hello")) {
                                    case "hello":
                                        v68 := GenS8{F0: -1, F1: (v1 << ((4294967295) & 63))}
                                        _ = v68
                                    case "foo bar":
                                        v69 := GenS2{F0: 42}
                                        _ = v69
                                    default:
                                        var v70 uint8 = v1
                                        _ = v70
                                    }
                                    _ = v67
                                default:
                                    var v71 bool = (true && true)
                                    _ = v71
                                }
                                if true {
                                    fmt.Println(fi48, v2)
                                    fmt.Println(v2, v2, v2)
                                    if false {
                                        fmt.Println(fi49, fi48)
                                    }
                                } else {
                                    var v72 int64 = (v2 - v2)
                                    _ = v72
                                }
                                if (-9223372036854775808 > fi49) {
                                    var v73 int8 = 100
                                    _ = v73
                                } else {
                                    var v74 uint64 = 605
                                    v75 := GenS7{F0: uint64(v1)}
                                    _ = len(v3)
                                    _ = v74
                                    _ = v75
                                }
                                v76 := GenS9{F0: v2}
                                var v77 float64 = 0.0
                                v78 := GenS4{F0: 0, F1: fi48, F2: (^(v1))}
                                _ = len(v47)
                                var v79 float64 = 0.1
                                _ = len(v47)
                                if true {
                                    v79 = ((float64((-2.25 * (1e10 * -2.25))) * (float64((0.0 * 0.0)) + 3.14159)) + v77)
                                } else {
                                    var v80 int = -201
                                    v2 = (v2 & (v2 >> ((255) & 63)))
                                    if true {
                                        v81 := GenS2{F0: 0}
                                        var v82 uint8 = v1
                                        if true {
                                            var v83 uint32 = 65535
                                            fmt.Println(v77, fi48)
                                            _ = v83
                                        }
                                        _ = v81
                                        _ = v82
                                    } else {
                                        switch "" {
                                        case "foo bar":
                                            switch "\n\t" {
                                            case "foo bar":
                                                if true {
                                                    if len(v47) >= 2 {
                                                        v47 = v47[1:len(v47)]
                                                    }
                                                    fmt.Println(fi49)
                                                    var v84 uint32 = 301
                                                    _ = v84
                                                } else {
                                                    var v85 int = (fi49 - 486)
                                                    v85 = (fi49 ^ -422)
                                                    if len(v47) > 0 {
                                                        v47[0] = (-(fi48))
                                                    }
                                                    _ = v85
                                                }
                                            case "":
                                                var v86 float32 = 3.14159
                                                _ = v86
                                            }
                                        case "hello":
                                            _ = len(v47)
                                        case "":
                                            v79 += 0.1
                                        }
                                        v79 -= v79
                                        v3[65535] = ("" + "hello")
                                    }
                                    v79 = float64(-2.25)
                                    _ = v80
                                }
                                v79 = float64((float32(v77) + 1e10))
                                _ = v76
                                _ = v77
                                _ = v78
                                _ = v79
                            }
                        default:
                            switch (fi48 & fi49) {
                            case -812:
                                v2 = (v2 - v2)
                            default:
                                var v87 int16 = -970
                                v1 = v1
                                _ = v87
                            }
                        }
                        v3[2] = ""
                    }
                    v2 = v2
                    _ = fi49
                }
                for fi88 := 0; fi88 < 7; fi88++ {
                    fmt.Println(fi48)
                    delete(v3, 2147483648)
                    if ((!true) && true) {
                        if true {
                            if true {
                                switch "\n\t" {
                                case "":
                                    var v89 uint16 = 65535
                                    _ = v89
                                case "hello":
                                    switch "hello" {
                                    case "\n\t":
                                        v1 = v1
                                    case "a":
                                        switch "foo bar" {
                                        case "a":
                                            if true {
                                                switch ("a" + "hello") {
                                                case "hello":
                                                    if false {
                                                        var v90 uint8 = v1
                                                        v91 := GenS8{F0: -1, F1: (v1 * (^(v1)))}
                                                        v1 = (v1 * 42)
                                                        _ = v90
                                                        _ = v91
                                                    } else {
                                                        if true {
                                                            var v92 bool = ((1e10 * (-2.25 + 0.1)) != float64(float32(1e10)))
                                                            v1 += v1
                                                            var v93 uint64 = 357
                                                            var v94 uint = 4294967295
                                                            _ = v92
                                                            _ = v93
                                                            _ = v94
                                                        } else {
                                                            v2 = v2
                                                        }
                                                        v2 += (v2 / (((v2 << ((2) & 63))) | 1))
                                                        var v95 uint16 = 665
                                                        _ = v95
                                                    }
                                                case "\n\t":
                                                    v1 += (v1 / ((1) | 1))
                                                }
                                            }
                                        }
                                    default:
                                        switch "foo bar" {
                                        case "hello":
                                            var v96 int8 = 99
                                            _ = v96
                                        case "\n\t":
                                            var v97 int64 = v2
                                            _ = v97
                                        case "a":
                                            v1 -= v1
                                        default:
                                            switch "hello" {
                                            case "a":
                                                v98 := GenS9{F0: 0}
                                                _ = v98
                                            case "\n\t":
                                                v2 -= 9223372036854775807
                                                v47 = append(v47, 9223372036854775807)
                                                v2 = (v2 * v2)
                                            case "foo bar":
                                                v2 -= v2
                                            default:
                                                fmt.Println(v2, v2)
                                            }
                                        }
                                    }
                                default:
                                    v99 := GenS1{F0: 500}
                                    _ = v99
                                }
                                var v100 uint32 = 42
                                _ = v100
                            } else {
                                var v101 float32 = (1.5 + (float32(-2.25) / 0.1))
                                v2 = 42
                                var v102 int8 = 127
                                _ = v101
                                _ = v102
                            }
                            v103 := GenS9{F0: v2}
                            _ = len(v47)
                            _ = v103
                        }
                        var v104 int8 = -75
                        _ = v104
                    }
                    _ = fi88
                }
                if false {
                    var v105 string = ("" + (("hello" + "hello") + (("a" + "") + "")))
                    v105 = ""
                    _ = v105
                }
                _ = fi48
            }
            var v106 uint = 366
            v47 = append(v47, 32767)
            _ = v106
        } else {
            for fi107 := 0; fi107 < 4; fi107++ {
                v108 := GenS10{F0: 1, F1: 2147483647, F2: -1}
                _ = fi107
                _ = v108
            }
            if (74 <= 1) {
                v2 += 42
                v109 := GenS11{F0: (-425 >= 0)}
                _ = v109
            }
        }
        for fi110 := 0; fi110 < 9; fi110++ {
            var v111 uint32 = 1
            _ = fi110
            _ = v111
        }
    }
    var v112 int32 = -87
    var v113 uint = 1
    v114 := GenS9{F0: 0}
    v1 -= uint8(v113)
    fmt.Println("v1:", v1)
    fmt.Println("v112:", v112)
    fmt.Println("v113:", v113)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v47:", len(v47), v47)
    _ = v1
    _ = v112
    _ = v113
    _ = v114
    _ = v2
    _ = v3
    _ = v47
}
