package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS3 struct { F0 uint32; F1 string; F2 uint64 }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16; F1 int64; F2 int32 }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 uint32; F1 int8; F2 int64 }
type GenS9 struct { F0 int16; F1 int16; F2 int16 }
type GenS10 struct { F0 bool; F1 int32 }
type GenS11 struct { F0 int16; F1 int; F2 int8 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 bool; F1 string }
func genFn0(p0 int32) uint64 {
    var v4 int32 = p0
    if false {
        if (0 == (v4 ^ v4)) {
            v4 = v4
            v4 = (-(v4))
        } else {
            p0 = (v4 * p0)
            p0 += p0
        }
    }
    var v5 int16 = 104
    _ = p0
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 int64, p1 string, p2 bool) int8 {
    switch p1 {
    case "\n\t":
        switch ("a" + p1) {
        case "a":
            p1 = p1
        default:
            var v6 uint8 = 127
            _ = v6
        }
    }
    for fi7 := 0; fi7 < 12; fi7++ {
        p2 = false
        if p2 {
            switch (^(fi7)) {
            case 127:
                fmt.Println(p2, p2)
            case 32767:
                var v8 int8 = 0
                _ = v8
            default:
                v9 := GenS1{F0: 0}
                _ = v9
            }
            v10 := []int{}
            _ = v10
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    return -128
}

func main() {
    var v1 uint32 = 2147483647
    v2 := []uint32{}
    v3 := make(map[int]uint32)
    for fi11 := 0; fi11 < 8; fi11++ {
        v1 = v1
        switch "hello" {
        case "a":
            for fi12 := 0; fi12 < 12; fi12++ {
                v13 := GenS1{F0: 2}
                _ = fi12
                _ = v13
            }
        case "\n\t":
            fmt.Println(fi11, fi11)
        default:
            delete(v3, (^(fi11)))
        }
        _ = fi11
    }
    if true {
        for fi14 := 0; fi14 < 13; fi14++ {
            v1 += 462
            _ = fi14
        }
        v1 += v1
    }
    var v15 int = -2147483648
    v1 += (v1 >> ((4294967295) & 63))
    v15 = (^(v15))
    v16 := GenS1{F0: 2}
    for fi17 := 0; fi17 < 2; fi17++ {
        v1 += (v1 & 478)
        v18 := GenS2{F0: 185, F1: (^(v1)), F2: -32768}
        if false {
            delete(v3, 0)
        }
        _ = fi17
        _ = v18
    }
    if true {
        if true {
            if false {
                var v19 string = "\n\t"
                _ = v19
            } else {
                var v20 string = "a"
                var v21 int32 = -2147483648
                var v22 bool = false
                _ = v20
                _ = v21
                _ = v22
            }
        }
    } else {
        switch (((("foo bar" + "a") + "foo bar") + ("" + "")) + "hello") {
        case "foo bar":
            for fi23 := 0; fi23 < 3; fi23++ {
                switch (fi23 ^ fi23) {
                case -127:
                    fmt.Println(fi23, v1, v1)
                case 4294967295:
                    switch "\n\t" {
                    case "foo bar":
                        if (!false) {
                            for fi24 := 0; fi24 < 7; fi24++ {
                                v3[v15] = 2147483647
                                _ = fi24
                            }
                            v25 := v3[v15]
                            _ = v25
                        }
                    }
                case 42:
                    for fi26 := 0; fi26 < 20; fi26++ {
                        v1 += v1
                        _ = fi26
                    }
                }
                if true {
                    for fi27 := 0; fi27 < 6; fi27++ {
                        v28 := GenS1{F0: 42}
                        v1 -= (v1 >> ((0) & 63))
                        var v29 bool = (!false)
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                    var v30 int = (v15 - fi23)
                    _ = v30
                }
                _ = fi23
            }
        }
        v1 = 399
    }
    if (-2147483648 == 1) {
        v31 := GenS2{F0: 127, F1: (^(v1)), F2: 591}
        if true {
            v15 += (v15 << ((2) & 63))
            v15 = v15
        } else {
            switch (-(v15)) {
            case -1:
                v15 += v15
            case 42:
                v1 = 4294967295
            case 128:
                v1 -= (v1 - v1)
            }
            v32 := GenS3{F0: (v1 + (v1 % ((v1) | 1))), F1: "\n\t", F2: 18446744073709551615}
            fmt.Println(v15, v1)
            _ = v32
        }
        v33 := GenS3{F0: 0, F1: ("a" + ""), F2: 1}
        _ = v31
        _ = v33
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi34 := 0; fi34 < 5; fi34++ {
        v15 -= (v15 << ((42) & 63))
        if true {
            v15 = 9223372036854775807
            switch ("hello" + "") {
            case "\n\t":
                for fi35 := 0; fi35 < 1; fi35++ {
                    if (false || false) {
                        v1 -= 772
                        v15 = v15
                        v36 := GenS2{F0: uint8(v1), F1: v1, F2: 1000}
                        v37 := GenS4{F0: 1, F1: true}
                        _ = v36
                        _ = v37
                    } else {
                        var v38 string = "\n\t"
                        v15 -= v15
                        _ = v38
                    }
                    _ = fi35
                }
            default:
                v39 := GenS1{F0: 2}
                _ = v39
            }
            switch (v15 ^ fi34) {
            case 32767:
                var v40 uint32 = (^(v1))
                _ = v40
            default:
                switch v15 {
                case 9223372036854775807:
                    if (!true) {
                        fmt.Println(v15, v15, v1)
                        var v41 int16 = 52
                        _ = v41
                    }
                case -728:
                    if (fi34 < -702) {
                        switch 9223372036854775807 {
                        case -136:
                            if (!false) {
                                v42 := GenS4{F0: 180, F1: (751 != 65535)}
                                if ((128 <= 32767) || false) {
                                    if (18446744073709551615 != 255) {
                                        switch "hello" {
                                        case "hello":
                                            if (255 == 232) {
                                                v15 = 127
                                                if true {
                                                    switch "" {
                                                    case "":
                                                        switch 240 {
                                                        case 1:
                                                            var v43 int = (v15 << ((255) & 63))
                                                            _ = v43
                                                        default:
                                                            v44 := GenS4{F0: 65535, F1: (0 >= -1)}
                                                            _ = v44
                                                        }
                                                    case "foo bar":
                                                        v1 = (v1 - 2147483647)
                                                    }
                                                    v15 -= 5
                                                }
                                            } else {
                                                var v45 float32 = (float32((3.14159 * (0.0 * 0.0))) - 0.1)
                                                switch ("foo bar" + "foo bar") {
                                                case "hello":
                                                    for fi46 := 0; fi46 < 4; fi46++ {
                                                        var v47 int32 = 314
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            var v48 uint64 = 233
                                                            _ = v48
                                                        case "":
                                                            switch fi34 {
                                                            case -1:
                                                                v1 = v1
                                                            case -96:
                                                                fmt.Println(v15)
                                                            case -29:
                                                                if (!(v45 == (v45 / v45))) {
                                                                    fmt.Println(v15, v45, v45)
                                                                    if (!false) {
                                                                        switch "foo bar" {
                                                                        case "foo bar":
                                                                            switch "" {
                                                                            case "\n\t":
                                                                                var v49 uint16 = 255
                                                                                _ = v49
                                                                            case "foo bar":
                                                                                if false {
                                                                                    var v50 int8 = 1
                                                                                    if false {
                                                                                        var v51 uint8 = 2
                                                                                        switch (fi34 + (fi46 | int(v51))) {
                                                                                        case 128:
                                                                                            v1 = (v1 / ((v1) | 1))
                                                                                        case -538:
                                                                                            if (false && ("foo bar" > "hello")) {
                                                                                                fmt.Println(v1, v1)
                                                                                            } else {
                                                                                                v1 += v1
                                                                                                var v52 uint = 981
                                                                                                v1 -= (v1 / (((v1 - (v1 / ((v1) | 1)))) | 1))
                                                                                                if len(v2) >= 2 {
                                                                                                    v2 = v2[1:len(v2)]
                                                                                                }
                                                                                                var v53 int8 = (v50 & (v50 >> (((^(v52))) & 63)))
                                                                                                var v54 int8 = v50
                                                                                                _ = v52
                                                                                                _ = v53
                                                                                                _ = v54
                                                                                            }
                                                                                        default:
                                                                                            v1 -= 1
                                                                                        }
                                                                                        if true {
                                                                                            if false {
                                                                                                v55 := GenS5{F0: (fi34 + -330)}
                                                                                                v50 += (v50 * v50)
                                                                                                v3[fi46] = v1
                                                                                                if (!false) {
                                                                                                    v51 -= 2
                                                                                                } else {
                                                                                                    if len(v2) > 0 {
                                                                                                        v2[0] = v1
                                                                                                    }
                                                                                                    fmt.Println(v50, v15, fi46)
                                                                                                    fmt.Println(v51, v51)
                                                                                                }
                                                                                                fmt.Println(v45, v15, v51)
                                                                                                _ = v55
                                                                                            } else {
                                                                                                v56 := GenS6{F0: 684, F1: 9223372036854775807, F2: (v47 / ((v47) | 1))}
                                                                                                if (!false) {
                                                                                                    if ((2 > 1) && false) {
                                                                                                        var v57 float64 = -2.25
                                                                                                        _ = v57
                                                                                                    }
                                                                                                    v47 -= (-(v47))
                                                                                                }
                                                                                                if (("\n\t" + "") < "a") {
                                                                                                    switch "a" {
                                                                                                    case "a":
                                                                                                        var v58 string = (("hello" + "") + (("" + (("a" + (("foo bar" + ("hello" + "\n\t")) + "a")) + "\n\t")) + "a"))
                                                                                                        switch (((v58 + "") + "\n\t") + v58) {
                                                                                                        case "foo bar":
                                                                                                            if false {
                                                                                                                switch fi46 {
                                                                                                                case -1:
                                                                                                                    var v59 int8 = 0
                                                                                                                    _ = v59
                                                                                                                case 128:
                                                                                                                    switch "foo bar" {
                                                                                                                    case "a":
                                                                                                                        var v60 int16 = int16(v1)
                                                                                                                        _ = v60
                                                                                                                    case "":
                                                                                                                        _ = len(v2)
                                                                                                                    case "\n\t":
                                                                                                                        var v61 uint = 18446744073709551615
                                                                                                                        _ = v61
                                                                                                                    }
                                                                                                                case -783:
                                                                                                                    switch (v58 + v58) {
                                                                                                                    case "":
                                                                                                                        if len(v2) >= 2 {
                                                                                                                            v2 = v2[1:len(v2)]
                                                                                                                        }
                                                                                                                    case "foo bar":
                                                                                                                        v50 = v50
                                                                                                                    case "a":
                                                                                                                        v62 := GenS6{F0: 0, F1: 614, F2: v47}
                                                                                                                        _ = v62
                                                                                                                    default:
                                                                                                                        var v63 int8 = (v50 / (((v50 / (((v50 * v50)) | 1))) | 1))
                                                                                                                        _ = v63
                                                                                                                    }
                                                                                                                default:
                                                                                                                    fmt.Println(fi46, v45)
                                                                                                                }
                                                                                                                v64 := GenS1{F0: 177}
                                                                                                                _ = v64
                                                                                                            } else {
                                                                                                                v3[(v15 ^ (v15 | fi34))] = 2147483648
                                                                                                            }
                                                                                                        }
                                                                                                        fmt.Println(fi34, v45, v50)
                                                                                                        var v65 float32 = float32(0.0)
                                                                                                        var v66 int64 = 42
                                                                                                        var v67 bool = (666 < 0)
                                                                                                        _ = v58
                                                                                                        _ = v65
                                                                                                        _ = v66
                                                                                                        _ = v67
                                                                                                    }
                                                                                                    var v68 uint8 = v51
                                                                                                    _ = v68
                                                                                                } else {
                                                                                                    v51 = 127
                                                                                                    if (v47 == 0) {
                                                                                                        var v69 int64 = 1
                                                                                                        v1 = v1
                                                                                                        _ = v69
                                                                                                    } else {
                                                                                                        if (!false) {
                                                                                                            switch -1 {
                                                                                                            case 268:
                                                                                                                var v70 int64 = -929
                                                                                                                if len(v2) > 0 {
                                                                                                                    v2[0] = (v1 - v1)
                                                                                                                }
                                                                                                                _ = v70
                                                                                                            case 0:
                                                                                                                v1 -= (v1 * (v1 | v1))
                                                                                                            default:
                                                                                                                fmt.Println(v45, v51, v50)
                                                                                                            }
                                                                                                        }
                                                                                                        fmt.Println(fi46)
                                                                                                    }
                                                                                                    fmt.Println(v15, v51, v15)
                                                                                                    v71 := GenS7{F0: (v1 / (((^(v1))) | 1)), F1: 1}
                                                                                                    if ("hello" != "") {
                                                                                                        v72 := GenS6{F0: -32768, F1: -393, F2: v47}
                                                                                                        _ = v72
                                                                                                    }
                                                                                                    _ = v71
                                                                                                }
                                                                                                v51 += 132
                                                                                                _ = v56
                                                                                            }
                                                                                            if true {
                                                                                                var v73 uint32 = 496
                                                                                                if (false || false) {
                                                                                                    var v74 uint32 = 390
                                                                                                    _ = v74
                                                                                                } else {
                                                                                                    var v75 bool = ((^(v47)) >= v47)
                                                                                                    if (18446744073709551615 < 42) {
                                                                                                        v45 = ((v45 * (v45 / v45)) / (float32(3.14159) + (v45 * (v45 + 0.1))))
                                                                                                        if (((v75 && true) || v75) || false) {
                                                                                                            var v76 int = (v15 & fi46)
                                                                                                            switch fi46 {
                                                                                                            case 2147483647:
                                                                                                                fmt.Println(v76, v15, v1)
                                                                                                            case -9223372036854775808:
                                                                                                                fmt.Println(v47, v76)
                                                                                                            }
                                                                                                            _ = v76
                                                                                                        } else {
                                                                                                            if len(v2) > 0 {
                                                                                                                v2[0] = (v73 >> ((1) & 63))
                                                                                                            }
                                                                                                        }
                                                                                                        fmt.Println(v50, v47, v50)
                                                                                                    } else {
                                                                                                        v50 = (-(v50))
                                                                                                    }
                                                                                                    var v77 int = fi46
                                                                                                    if v75 {
                                                                                                        var v78 int16 = 694
                                                                                                        switch (fi46 / (((-(v15))) | 1)) {
                                                                                                        case 9223372036854775807:
                                                                                                            v79 := GenS4{F0: uint16(v1), F1: ("\n\t" > "")}
                                                                                                            _ = len(v2)
                                                                                                            var v80 int32 = (v47 * v47)
                                                                                                            var v81 int32 = 742
                                                                                                            _ = v79
                                                                                                            _ = v80
                                                                                                            _ = v81
                                                                                                        case 32767:
                                                                                                            v75 = (int32(fi34) != -2147483648)
                                                                                                        case -2147483648:
                                                                                                            v82 := GenS8{F0: (^(v1)), F1: (^(v50)), F2: 9223372036854775807}
                                                                                                            _ = v82
                                                                                                        }
                                                                                                        _ = v78
                                                                                                    }
                                                                                                    _ = v75
                                                                                                    _ = v77
                                                                                                }
                                                                                                fmt.Println(v45, v1)
                                                                                                _ = v73
                                                                                            } else {
                                                                                                switch (fi46 * (v15 | -1)) {
                                                                                                case 42:
                                                                                                    switch (v15 * (fi46 - 42)) {
                                                                                                    case 42:
                                                                                                        var v83 uint8 = (v51 >> ((65535) & 63))
                                                                                                        _ = v83
                                                                                                    case 857:
                                                                                                        var v84 uint8 = 1
                                                                                                        _ = v84
                                                                                                    case -804:
                                                                                                        v51 += (v51 & 85)
                                                                                                    default:
                                                                                                        v85 := GenS4{F0: 65535, F1: (!true)}
                                                                                                        _ = v85
                                                                                                    }
                                                                                                default:
                                                                                                    v86 := GenS1{F0: 65535}
                                                                                                    _ = v86
                                                                                                }
                                                                                                switch ("\n\t" + ("\n\t" + ("foo bar" + "hello"))) {
                                                                                                case "hello":
                                                                                                    _ = len(v3)
                                                                                                case "":
                                                                                                    if true {
                                                                                                        var v87 uint = 65535
                                                                                                        var v88 int = fi46
                                                                                                        _ = v87
                                                                                                        _ = v88
                                                                                                    } else {
                                                                                                        v15 -= (fi34 ^ (fi46 >> ((1) & 63)))
                                                                                                        v89 := GenS9{F0: -32768, F1: -93, F2: -32768}
                                                                                                        if (!(1.5 == (v45 + 1e10))) {
                                                                                                            var v90 string = ("" + "")
                                                                                                            _ = v90
                                                                                                        }
                                                                                                        if false {
                                                                                                            fmt.Println(v15)
                                                                                                            if len(v2) > 0 {
                                                                                                                v2[0] = (v1 + v1)
                                                                                                            }
                                                                                                        } else {
                                                                                                            v45 -= 1.5
                                                                                                            v50 = v50
                                                                                                        }
                                                                                                        v50 += (v50 ^ genFn1(-1, ("hello" + "a"), ((!(1 > -946)) || true)))
                                                                                                        _ = v89
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        _ = v51
                                                                                    } else {
                                                                                        v91 := GenS2{F0: 56, F1: v1, F2: -1}
                                                                                        if ((0 == -1) || false) {
                                                                                            v92 := GenS10{F0: ((v45 != 0.1) || true), F1: (v47 / ((-1) | 1))}
                                                                                            v45 -= 0.0
                                                                                            switch 127 {
                                                                                            case 42:
                                                                                                v15 += fi34
                                                                                            default:
                                                                                                v93 := GenS11{F0: 1000, F1: fi34, F2: (v50 / (((v50 * int8(v47))) | 1))}
                                                                                                _ = v93
                                                                                            }
                                                                                            _ = v92
                                                                                        } else {
                                                                                            v94 := v3[(^(fi46))]
                                                                                            v15 -= fi46
                                                                                            v45 -= v45
                                                                                            _ = v94
                                                                                        }
                                                                                        var v95 int16 = -32768
                                                                                        _ = v91
                                                                                        _ = v95
                                                                                    }
                                                                                    _ = v50
                                                                                } else {
                                                                                    delete(v3, (v15 % (((fi46 - 32768)) | 1)))
                                                                                    v47 -= (v47 << ((316) & 63))
                                                                                }
                                                                            }
                                                                        case "\n\t":
                                                                            v15 -= fi34
                                                                        default:
                                                                            v15 += -656
                                                                        }
                                                                    } else {
                                                                        var v96 int8 = genFn1(-772, ("" + "hello"), true)
                                                                        _ = v96
                                                                    }
                                                                    delete(v3, fi46)
                                                                } else {
                                                                    if (genFn1(1, "hello", true) <= -1) {
                                                                        var v97 uint32 = v1
                                                                        v1 += (v1 >> ((987) & 63))
                                                                        _ = v97
                                                                    }
                                                                    v98 := GenS2{F0: 188, F1: v1, F2: -1}
                                                                    _ = v98
                                                                }
                                                            }
                                                        case "foo bar":
                                                            _ = len(v2)
                                                        default:
                                                            v1 -= v1
                                                        }
                                                        var v99 int16 = -1
                                                        _ = fi46
                                                        _ = v47
                                                        _ = v99
                                                    }
                                                case "a":
                                                    var v100 uint32 = (v1 | 0)
                                                    _ = v100
                                                }
                                                _ = len(v3)
                                                _ = v45
                                            }
                                        case "\n\t":
                                            _ = len(v2)
                                        default:
                                            v1 = (v1 & (v1 | (v1 / ((v1) | 1))))
                                        }
                                    }
                                    var v101 float64 = -2.25
                                    _ = v101
                                }
                                _ = len(v2)
                                _ = v42
                            }
                        case 0:
                            var v102 uint = 1
                            _ = v102
                        case -432:
                            var v103 uint64 = 18446744073709551615
                            _ = v103
                        default:
                            _ = len(v2)
                        }
                        delete(v3, -9223372036854775808)
                    } else {
                        var v104 int = 1
                        _ = v104
                    }
                case -967:
                    var v105 uint32 = 225
                    _ = v105
                }
            }
        } else {
            if true {
                if false {
                    var v106 int32 = 670
                    var v107 float64 = -2.25
                    _ = v106
                    _ = v107
                }
            }
        }
        v1 = 836
        _ = fi34
    }
    v108 := GenS7{F0: (v1 >> ((0) & 63)), F1: (v15 % (((v15 + 645)) | 1))}
    fmt.Println(v1, v1)
    if (true || ((3.14159 - (1.5 - 1e10)) == 3.14159)) {
        switch "foo bar" {
        case "\n\t":
            for fi109 := 0; fi109 < 6; fi109++ {
                v1 -= (^(v1))
                _ = fi109
            }
        default:
            var v110 uint64 = 2
            _ = v110
        }
        switch "\n\t" {
        case "foo bar":
            if (!(3.14159 > ((1.5 * ((0.1 + (3.14159 - (3.14159 - 1.5))) * 3.14159)) * -2.25))) {
                switch (v15 / ((v15) | 1)) {
                case 127:
                    v111 := GenS12{F0: 2147483647}
                    _ = v111
                case 284:
                    v15 = -8
                case -389:
                    var v112 uint32 = 2147483648
                    _ = v112
                }
                v113 := GenS13{F0: (0 != 2147483647), F1: (("hello" + "a") + "hello")}
                var v114 int16 = -808
                _ = v113
                _ = v114
            } else {
                v115 := GenS12{F0: -696}
                _ = v115
            }
        case "\n\t":
            if ((((8 != -1) || (312 <= 1)) || false) && (432 != 18446744073709551615)) {
                v15 += v15
                var v116 float64 = float64(1e10)
                for fi117 := 0; fi117 < 15; fi117++ {
                    v116 = 0.1
                    _ = fi117
                }
                _ = v116
            }
        }
        v118 := GenS7{F0: (v1 % ((v1) | 1)), F1: (^(v15))}
        _ = v118
    }
    var v119 int = 32767
    fmt.Println(v15)
    fmt.Println("v1:", v1)
    fmt.Println("v119:", v119)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v108
    _ = v119
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
