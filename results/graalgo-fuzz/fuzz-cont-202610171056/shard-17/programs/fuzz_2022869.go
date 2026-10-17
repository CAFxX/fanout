package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 bool }
type GenS2 struct { F0 int; F1 int64; F2 uint8 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 uint; F1 int8; F2 int32 }
type GenS5 struct { F0 string; F1 uint8; F2 uint16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int8; F1 uint64; F2 int64 }
type GenS9 struct { F0 uint; F1 int64 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int16; F1 string; F2 int64 }
type GenS12 struct { F0 uint; F1 uint16 }
type GenS13 struct { F0 uint64 }
func genFn0(p0 uint32) bool {
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := []int{fi6, -564, fi6, 512}
        _ = fi6
        _ = v7
    }
    v8 := []int{544, -1, 317, 128}
    var v9 float64 = 1e10
    var v10 uint64 = 1
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return true
}
func genFn1(p0 uint16, p1 int64) uint {
    var v11 uint8 = 2
    var v12 int8 = 100
    switch "" {
    case "\n\t":
        v13 := []int{}
        _ = v13
    case "foo bar":
        for fi14 := 0; fi14 < 17; fi14++ {
            for fi15 := 0; fi15 < 7; fi15++ {
                v11 += 42
                _ = fi15
            }
            for fi16 := 0; fi16 < 2; fi16++ {
                p1 -= 1
                if false {
                    v17 := GenS1{F0: (fi14 | (-(fi16))), F1: 728, F2: (1.5 <= ((1e10 - ((-2.25 - 0.1) * -2.25)) * (3.14159 / float64(0.1))))}
                    _ = v17
                } else {
                    var v18 bool = (true || false)
                    v19 := GenS1{F0: fi16, F1: -32768, F2: ((255 < 65535) || false)}
                    switch ("hello" + "\n\t") {
                    case "\n\t":
                        v20 := []int{(-(fi16)), -770, fi14, fi16}
                        v20 = append(v20, (-(fi16)))
                        if (false || (!false)) {
                            if v18 {
                                switch ("foo bar" + ("foo bar" + ("foo bar" + "hello"))) {
                                case "hello":
                                    var v21 float32 = 1.5
                                    _ = v21
                                case "":
                                    switch "\n\t" {
                                    case "":
                                        switch fi16 {
                                        case 128:
                                            _ = len(v20)
                                        case -2147483648:
                                            v22 := GenS2{F0: 127, F1: (p1 - int64(v11)), F2: v11}
                                            _ = v22
                                        }
                                    default:
                                        v11 = (^(v11))
                                    }
                                }
                                if v18 {
                                    var v23 int64 = 53
                                    _ = v23
                                }
                                v11 = (v11 << ((255) & 63))
                            }
                            var v24 float64 = 0.1
                            _ = len(v20)
                            _ = v24
                        } else {
                            fmt.Println(p1)
                        }
                        v25 := GenS3{F0: false, F1: (^(v12))}
                        switch "a" {
                        case "a":
                            var v26 int = (fi16 >> ((663) & 63))
                            _ = v26
                        }
                        _ = v20
                        _ = v25
                    case "a":
                        v27 := GenS3{F0: (v18 || (true && (42 > -1))), F1: v12}
                        _ = v27
                    }
                    _ = v18
                    _ = v19
                }
                _ = fi16
            }
            var v28 int64 = (p1 * (^(p1)))
            _ = fi14
            _ = v28
        }
    }
    fmt.Println(v12, p0)
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 4294967295
}

func main() {
    var v1 string = ""
    var v2 uint32 = 2147483647
    v3 := make(map[int]int)
    v4 := []uint32{(v2 & v2), (v2 << ((uint(v2)) & 63))}
    var v5 int32 = 42
    for fi29 := 0; fi29 < 8; fi29++ {
        v1 = "foo bar"
        _ = fi29
    }
    var v30 float32 = 0.1
    fmt.Println(v2)
    if false {
        if len(v4) > 0 {
            v4[0] = (v2 | v2)
        }
        for fi31 := 0; fi31 < 11; fi31++ {
            if (1 > 301) {
                var v32 uint32 = v2
                for fi33 := 0; fi33 < 18; fi33++ {
                    v34 := GenS3{F0: true, F1: -128}
                    v35 := GenS3{F0: true, F1: -1}
                    switch ((("hello" + v1) + ("" + ("" + v1))) + "\n\t") {
                    case "hello":
                        var v36 int16 = 1000
                        _ = v36
                    case "":
                        var v37 float32 = v30
                        _ = v37
                    case "foo bar":
                        if (true || ((false && (!(int16(fi31) != 42))) && (true || true))) {
                            var v38 uint8 = 121
                            _ = v38
                        }
                    }
                    _ = fi33
                    _ = v34
                    _ = v35
                }
                if false {
                    if (v30 >= (3.14159 / 1.5)) {
                        v2 = v2
                        if true {
                            for fi39 := 0; fi39 < 10; fi39++ {
                                v40 := GenS4{F0: 42, F1: -107, F2: v5}
                                _ = fi39
                                _ = v40
                            }
                        } else {
                            v30 -= v30
                            v41 := GenS2{F0: fi31, F1: 42, F2: 255}
                            v3[534] = (fi31 * (fi31 >> ((2) & 63)))
                            _ = v41
                        }
                        var v42 uint8 = 63
                        _ = v42
                    }
                    switch v1 {
                    case "\n\t":
                        if (int64(v32) > 9223372036854775807) {
                            if len(v4) > 0 {
                                v4[0] = 365
                            }
                        } else {
                            if true {
                                var v43 float64 = 1.5
                                if ((65535 <= 673) || (!true)) {
                                    v32 = v2
                                } else {
                                    switch fi31 {
                                    case -1:
                                        var v44 int64 = 9223372036854775807
                                        _ = v44
                                    case 42:
                                        v2 = v2
                                    default:
                                        if false {
                                            fmt.Println(v32)
                                        }
                                    }
                                }
                                _ = v43
                            } else {
                                fmt.Println(v30, v30)
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                var v45 uint64 = 4294967295
                                _ = v45
                            }
                            v46 := GenS4{F0: 255, F1: -128, F2: (v5 / ((v5) | 1))}
                            _ = v46
                        }
                    }
                    if (42 <= 42) {
                        v5 = v5
                    } else {
                        var v47 int16 = 1
                        _ = v47
                    }
                } else {
                    var v48 int16 = -1
                    _ = v48
                }
                _ = v32
            }
            _ = fi31
        }
        var v49 int64 = 42
        _ = v49
    } else {
        for fi50 := 0; fi50 < 19; fi50++ {
            for fi51 := 0; fi51 < 1; fi51++ {
                if (("\n\t" + v1) <= (v1 + v1)) {
                    fmt.Println(v2, v30, v2)
                    if len(v4) > 0 {
                        v4[0] = (v2 * v2)
                    }
                    if (18446744073709551615 <= 2) {
                        v52 := GenS5{F0: "", F1: 97, F2: 255}
                        _ = v52
                    }
                    v53 := GenS3{F0: (true || (2 >= v2)), F1: -41}
                    _ = v53
                }
                _ = fi51
            }
            _ = fi50
        }
        if true {
            if (!true) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v5 -= v5
            }
        }
    }
    v30 = ((1e10 * -2.25) * 1.5)
    fmt.Println(v5, v2)
    var v54 uint8 = 124
    switch v1 {
    case "":
        switch 2147483647 {
        case 32767:
            switch ((v1 + v1) + "\n\t") {
            case "a":
                var v55 int32 = (v5 % ((v5) | 1))
                _ = v55
            case "\n\t":
                _ = len(v3)
            }
        case 397:
            v54 = (v54 % ((2) | 1))
        case 9223372036854775807:
            v5 = 0
        }
    default:
        v56 := GenS6{F0: v1}
        _ = v56
    }
    v30 = 1e10
    v2 = v2
    for fi57 := 0; fi57 < 2; fi57++ {
        fmt.Println(v54, v1)
        switch (-(fi57)) {
        case 128:
            v54 -= v54
        case 32768:
            var v58 uint = 42
            _ = v58
        }
        var v59 int32 = (v5 * -2147483648)
        _ = fi57
        _ = v59
    }
    v5 = (^(v5))
    fmt.Println(v5)
    switch v1 {
    case "foo bar":
        v60 := GenS7{F0: int64(v54)}
        _ = v60
    }
    if ((true || false) && (("\n\t" + "hello") != v1)) {
        switch -2147483648 {
        case 127:
            for fi61 := 0; fi61 < 9; fi61++ {
                switch (fi61 | (^(fi61))) {
                case 32768:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                if true {
                    switch (-(fi61)) {
                    case -2147483648:
                        _ = len(v4)
                    case 127:
                        fmt.Println(v54, v1)
                    default:
                        for fi62 := 0; fi62 < 4; fi62++ {
                            v4 = append(v4, 42)
                            fmt.Println(v2)
                            var v63 int = (fi61 % (((fi62 | fi62)) | 1))
                            if true {
                                _ = len(v4)
                                if (false || (0 != (^(v2)))) {
                                    if (!(65535 <= (v2 / ((4294967295) | 1)))) {
                                        v1 = ""
                                        fmt.Println(v5)
                                    } else {
                                        var v64 int8 = -114
                                        fmt.Println(v1, v1)
                                        v5 += (v5 >> ((65535) & 63))
                                        var v65 bool = (0.1 != 0.0)
                                        var v66 uint64 = 0
                                        _ = v64
                                        _ = v65
                                        _ = v66
                                    }
                                } else {
                                    fmt.Println(v2, fi61, fi61)
                                }
                                v67 := GenS1{F0: (fi61 & (fi61 << ((2) & 63))), F1: -32768, F2: false}
                                _ = v67
                            } else {
                                var v68 uint = genFn1(1, 0)
                                fmt.Println(fi61, v68)
                                if ("foo bar" < ("\n\t" + (v1 + (((v1 + v1) + "a") + v1)))) {
                                    var v69 uint64 = 42
                                    if false {
                                        var v70 uint = v68
                                        v71 := GenS8{F0: 100, F1: 881, F2: -9223372036854775808}
                                        var v72 int64 = -676
                                        _ = v70
                                        _ = v71
                                        _ = v72
                                    }
                                    _ = v69
                                } else {
                                    var v73 int32 = (v5 % (((v5 ^ 2147483647)) | 1))
                                    if false {
                                        v5 += 1
                                        v54 += 102
                                        if ((710 <= 484) || (uint16(v2) >= 57)) {
                                            if true {
                                                _ = len(v3)
                                                var v74 int8 = -29
                                                _ = v74
                                            }
                                            v2 -= v2
                                        } else {
                                            v75 := v3[(fi62 >> ((4294967295) & 63))]
                                            if (1.5 > 0.1) {
                                                var v76 uint32 = (v2 | 307)
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                switch v75 {
                                                case 4294967295:
                                                    v54 = (v54 / ((uint8(v2)) | 1))
                                                default:
                                                    if ((!(243 > 4294967295)) && true) {
                                                        if (-1 > -442) {
                                                            v73 = v73
                                                            v75 -= v63
                                                            var v77 uint16 = 255
                                                            _ = v77
                                                        } else {
                                                            if len(v4) > 0 {
                                                                v4[0] = (v76 % ((v76) | 1))
                                                            }
                                                            switch 136 {
                                                            case 811:
                                                                if (1e10 < 1.5) {
                                                                    var v78 int32 = v73
                                                                    if true {
                                                                        var v79 int = (^(v63))
                                                                        _ = v79
                                                                    } else {
                                                                        if ("hello" < (("\n\t" + v1) + v1)) {
                                                                            if true {
                                                                                v2 -= v2
                                                                                v75 += 0
                                                                            } else {
                                                                                fmt.Println(fi61, v5)
                                                                                v5 += (v78 * v5)
                                                                                var v80 uint = v68
                                                                                _ = v80
                                                                            }
                                                                            v81 := GenS1{F0: 127, F1: 0, F2: (v76 == v76)}
                                                                            _ = v81
                                                                        } else {
                                                                            var v82 string = (v1 + (("a" + "\n\t") + "foo bar"))
                                                                            _ = v82
                                                                        }
                                                                        v68 = (^(v68))
                                                                    }
                                                                    _ = v78
                                                                }
                                                            default:
                                                                if (false || true) {
                                                                    var v83 int8 = 0
                                                                    if true {
                                                                        v68 = (^(v68))
                                                                        switch (v75 / (((fi62 & (fi62 ^ v75))) | 1)) {
                                                                        case 42:
                                                                            v4 = append(v4, 1)
                                                                        case 69:
                                                                            if false {
                                                                                fmt.Println(v1)
                                                                            } else {
                                                                                if true {
                                                                                    v5 += (v5 * (v73 - (v73 / (((v5 + v5)) | 1))))
                                                                                    if true {
                                                                                        fmt.Println(v73, v30)
                                                                                    }
                                                                                } else {
                                                                                    var v84 string = v1
                                                                                    switch (v75 & fi62) {
                                                                                    case 42:
                                                                                        var v85 bool = ((v68 / (((v68 | (v68 & (v68 % ((152) | 1))))) | 1)) < genFn1(42, -671))
                                                                                        _ = v85
                                                                                    case 0:
                                                                                        if true {
                                                                                            if (false || false) {
                                                                                                var v86 float64 = 1e10
                                                                                                _ = v86
                                                                                            } else {
                                                                                                v54 -= (^(v54))
                                                                                                if true {
                                                                                                    var v87 uint32 = 42
                                                                                                    var v88 float32 = v30
                                                                                                    var v89 int64 = 42
                                                                                                    if false {
                                                                                                        if (0 == v2) {
                                                                                                            var v90 uint8 = v54
                                                                                                            if len(v4) > 0 {
                                                                                                                v4[0] = (v2 ^ 1)
                                                                                                            }
                                                                                                            _ = v90
                                                                                                        }
                                                                                                        switch "" {
                                                                                                        case "hello":
                                                                                                            v88 = 3.14159
                                                                                                        case "foo bar":
                                                                                                            v91 := GenS9{F0: v68, F1: (v89 + 9223372036854775807)}
                                                                                                            _ = v91
                                                                                                        default:
                                                                                                            v92 := GenS10{F0: 32767}
                                                                                                            _ = v92
                                                                                                        }
                                                                                                        v83 += 42
                                                                                                        switch v63 {
                                                                                                        case 0:
                                                                                                            v75 += 1
                                                                                                        case -392:
                                                                                                            if ((!(int16(v75) < 32767)) || false) {
                                                                                                                v75 = (^(v75))
                                                                                                                v5 = v5
                                                                                                                v1 = ((v1 + "") + v84)
                                                                                                                var v93 int16 = 0
                                                                                                                _ = v93
                                                                                                            }
                                                                                                        case 4294967295:
                                                                                                            switch (v63 % (((fi62 % ((int(v83)) | 1))) | 1)) {
                                                                                                            case -2147483648:
                                                                                                                if false {
                                                                                                                    switch (^(v63)) {
                                                                                                                    case 97:
                                                                                                                        if ((v68 * genFn1(42, (-(v89)))) >= genFn1(42, 0)) {
                                                                                                                            switch (v63 % ((9223372036854775807) | 1)) {
                                                                                                                            case -821:
                                                                                                                                var v94 uint8 = (v54 % ((v54) | 1))
                                                                                                                                _ = v94
                                                                                                                            case 127:
                                                                                                                                var v95 bool = (2 != 0)
                                                                                                                                _ = v95
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            v84 = (((((("" + v1) + "foo bar") + v84) + v1) + v84) + (v1 + v1))
                                                                                                                            if (false && true) {
                                                                                                                                var v96 uint64 = 42
                                                                                                                                _ = v96
                                                                                                                            }
                                                                                                                            _ = len(v4)
                                                                                                                        }
                                                                                                                    case 762:
                                                                                                                        var v97 int32 = (v73 * v73)
                                                                                                                        _ = v97
                                                                                                                    case 128:
                                                                                                                        if (0 >= 701) {
                                                                                                                            _ = len(v4)
                                                                                                                        } else {
                                                                                                                            if (((v5 <= (v5 << (((^(v68))) & 63))) && true) || false) {
                                                                                                                                var v98 int8 = (-(v83))
                                                                                                                                v3[fi62] = v75
                                                                                                                                _ = v98
                                                                                                                            }
                                                                                                                            if ((((v83 - -1) < v83) && (255 <= (v68 - v68))) || true) {
                                                                                                                                var v99 int8 = v83
                                                                                                                                if (false && ((!(true || true)) || false)) {
                                                                                                                                    v100 := GenS11{F0: 1000, F1: ("foo bar" + ("" + ((v1 + (v1 + v1)) + v84))), F2: -1}
                                                                                                                                    if (true || false) {
                                                                                                                                        v87 = uint32(v68)
                                                                                                                                        fmt.Println(v68)
                                                                                                                                        v4 = append(v4, v76)
                                                                                                                                    } else {
                                                                                                                                        var v101 float32 = (float32(0.1) * v88)
                                                                                                                                        var v102 uint = v68
                                                                                                                                        _ = v101
                                                                                                                                        _ = v102
                                                                                                                                    }
                                                                                                                                    if (((v87 | v2) > (v76 & (^(v76)))) || false) {
                                                                                                                                        v103 := GenS1{F0: (v75 / ((int(v2)) | 1)), F1: 725, F2: (v87 < 729)}
                                                                                                                                        var v104 uint8 = v54
                                                                                                                                        _ = v103
                                                                                                                                        _ = v104
                                                                                                                                    }
                                                                                                                                    _ = v100
                                                                                                                                } else {
                                                                                                                                    v1 = v1
                                                                                                                                    var v105 uint64 = 435
                                                                                                                                    if len(v4) > 0 {
                                                                                                                                        v4[0] = uint32(v68)
                                                                                                                                    }
                                                                                                                                    var v106 string = v1
                                                                                                                                    _ = v105
                                                                                                                                    _ = v106
                                                                                                                                }
                                                                                                                                v5 = int32(v99)
                                                                                                                                var v107 uint8 = 0
                                                                                                                                _ = v107
                                                                                                                                _ = v99
                                                                                                                            }
                                                                                                                            if false {
                                                                                                                                var v108 uint16 = 699
                                                                                                                                _ = v108
                                                                                                                            } else {
                                                                                                                                v109 := v3[int(v89)]
                                                                                                                                var v110 int = v63
                                                                                                                                if (float32(0.1) <= v30) {
                                                                                                                                    _ = len(v3)
                                                                                                                                    if false {
                                                                                                                                        var v111 float32 = (float32(0.0) / 1e10)
                                                                                                                                        v112 := GenS3{F0: true, F1: v83}
                                                                                                                                        _ = v111
                                                                                                                                        _ = v112
                                                                                                                                    }
                                                                                                                                } else {
                                                                                                                                    if false {
                                                                                                                                        var v113 int64 = v89
                                                                                                                                        _ = len(v3)
                                                                                                                                        v114 := GenS8{F0: -1, F1: 42, F2: (v89 & v89)}
                                                                                                                                        _ = v113
                                                                                                                                        _ = v114
                                                                                                                                    } else {
                                                                                                                                        if (false || (462 < 342)) {
                                                                                                                                            fmt.Println(v30)
                                                                                                                                            v115 := GenS4{F0: v68, F1: v83, F2: -2147483648}
                                                                                                                                            _ = v115
                                                                                                                                        }
                                                                                                                                        v116 := GenS7{F0: 0}
                                                                                                                                        _ = v116
                                                                                                                                    }
                                                                                                                                    v87 = v87
                                                                                                                                    var v117 int32 = v5
                                                                                                                                    _ = v117
                                                                                                                                }
                                                                                                                                _ = v109
                                                                                                                                _ = v110
                                                                                                                            }
                                                                                                                            v75 = v63
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        if (fi61 >= v75) {
                                                                                                                            v118 := GenS12{F0: uint(v76), F1: 65535}
                                                                                                                            _ = len(v4)
                                                                                                                            var v119 int = 127
                                                                                                                            _ = v118
                                                                                                                            _ = v119
                                                                                                                        }
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    var v120 int64 = -9223372036854775808
                                                                                                                    _ = v120
                                                                                                                }
                                                                                                            }
                                                                                                        default:
                                                                                                            var v121 uint64 = 4294967295
                                                                                                            _ = v121
                                                                                                        }
                                                                                                        var v122 uint = genFn1(838, v89)
                                                                                                        _ = v122
                                                                                                    } else {
                                                                                                        var v123 uint8 = (v54 / ((2) | 1))
                                                                                                        _ = v123
                                                                                                    }
                                                                                                    fmt.Println(v76, v87)
                                                                                                    _ = v87
                                                                                                    _ = v88
                                                                                                    _ = v89
                                                                                                }
                                                                                            }
                                                                                            v84 = "\n\t"
                                                                                        }
                                                                                    case 32768:
                                                                                        v30 -= v30
                                                                                    }
                                                                                    _ = v84
                                                                                }
                                                                            }
                                                                        default:
                                                                            var v124 uint16 = 243
                                                                            _ = v124
                                                                        }
                                                                        switch "" {
                                                                        case "foo bar":
                                                                            var v125 uint = v68
                                                                            _ = v125
                                                                        }
                                                                        v126 := GenS1{F0: fi62, F1: 1000, F2: false}
                                                                        _ = v126
                                                                    } else {
                                                                        switch (v1 + (v1 + "hello")) {
                                                                        case "a":
                                                                            switch "hello" {
                                                                            case "":
                                                                                var v127 int32 = (^(v73))
                                                                                _ = v127
                                                                            case "foo bar":
                                                                                if false {
                                                                                    var v128 int = fi62
                                                                                    var v129 float64 = 3.14159
                                                                                    v130 := GenS1{F0: fi61, F1: 42, F2: ((v54 >> ((847) & 63)) == v54)}
                                                                                    fmt.Println(v2, v75)
                                                                                    v131 := GenS13{F0: 218}
                                                                                    _ = v128
                                                                                    _ = v129
                                                                                    _ = v130
                                                                                    _ = v131
                                                                                } else {
                                                                                    v132 := GenS11{F0: -717, F1: (v1 + (v1 + (v1 + "\n\t"))), F2: -296}
                                                                                    _ = v132
                                                                                }
                                                                            case "hello":
                                                                                v1 = v1
                                                                                v63 = (-(v63))
                                                                            }
                                                                        default:
                                                                            var v133 string = v1
                                                                            _ = v133
                                                                        }
                                                                        _ = len(v4)
                                                                        var v134 uint64 = 722
                                                                        v75 = int(v134)
                                                                        v68 = (^(v68))
                                                                        _ = v134
                                                                    }
                                                                    _ = v83
                                                                }
                                                            }
                                                        }
                                                        v63 -= v75
                                                    } else {
                                                        v63 = v63
                                                        switch (v63 * (fi61 & fi61)) {
                                                        case 32768:
                                                            var v135 uint16 = 1
                                                            _ = v135
                                                        case 42:
                                                            var v136 float32 = v30
                                                            _ = v136
                                                        default:
                                                            if true {
                                                                var v137 int16 = -164
                                                                if (false || ((false && true) || true)) {
                                                                    var v138 uint16 = 816
                                                                    _ = v138
                                                                } else {
                                                                    fmt.Println(v2, v54)
                                                                }
                                                                _ = v137
                                                            }
                                                        }
                                                        switch fi61 {
                                                        case 42:
                                                            var v139 float32 = 0.1
                                                            if (uint(v54) >= v68) {
                                                                var v140 int32 = (v5 + 182)
                                                                fmt.Println(v68, v5)
                                                                _ = v140
                                                            }
                                                            _ = v139
                                                        case 908:
                                                            v141 := GenS6{F0: "\n\t"}
                                                            _ = v141
                                                        }
                                                        v142 := GenS6{F0: v1}
                                                        _ = v142
                                                    }
                                                }
                                                _ = v76
                                            } else {
                                                fmt.Println(v63, v73, v1)
                                                fmt.Println(v54, v63)
                                            }
                                            var v143 uint8 = (^(v54))
                                            _ = v143
                                            _ = v75
                                        }
                                    } else {
                                        v144 := GenS4{F0: 685, F1: 20, F2: v5}
                                        v145 := GenS7{F0: 333}
                                        if len(v4) > 0 {
                                            v4[0] = (v2 & v2)
                                        }
                                        _ = v144
                                        _ = v145
                                    }
                                    var v146 int = 0
                                    _ = v146
                                    _ = v73
                                }
                                _ = v68
                            }
                            var v147 int = (fi61 >> ((399) & 63))
                            _ = fi62
                            _ = v147
                            _ = v63
                        }
                    }
                } else {
                    v5 += (v5 | int32(v2))
                    v2 = 4294967295
                    v30 += (v30 / 1.5)
                }
                _ = fi61
            }
        case -685:
            var v148 string = v1
            _ = v148
        case -1:
            if true {
                if false {
                    for fi149 := 0; fi149 < 17; fi149++ {
                        v5 -= (v5 << ((18446744073709551615) & 63))
                        _ = fi149
                    }
                    v150 := GenS1{F0: -2147483648, F1: -643, F2: true}
                    _ = v150
                }
                v151 := GenS5{F0: "", F1: (v54 & 128), F2: 809}
                if (-57 != -1) {
                    v152 := GenS10{F0: 869}
                    _ = v152
                }
                _ = v151
            }
        }
        fmt.Println(v2)
    }
    var v153 uint8 = (v54 << ((0) & 63))
    fmt.Println(v153)
    fmt.Println("v1:", v1)
    fmt.Println("v153:", v153)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v153
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v54
}
