package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 uint32; F1 int; F2 int32 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int16; F1 int16; F2 uint32 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int; F1 uint }
type GenS12 struct { F0 int32; F1 bool; F2 uint64 }
type GenS13 struct { F0 uint64; F1 int16; F2 int }
type GenS14 struct { F0 string; F1 uint8 }
type GenS15 struct { F0 uint16; F1 uint16; F2 string }
type GenS16 struct { F0 int8; F1 uint32; F2 int32 }
func genFn0(p0 int64) uint16 {
    fmt.Println(p0)
    v7 := []int{32768, -1}
    _ = p0
    _ = v7
    return 1000
}
func genFn1(p0 uint32, p1 bool) int8 {
    var v8 uint16 = 255
    if (p1 && ((18446744073709551615 == 1) && true)) {
        v9 := GenS1{F0: 0, F1: (v8 & v8)}
        switch 1 {
        case -2147483648:
            switch 9223372036854775807 {
            case 32768:
                switch 9223372036854775807 {
                case 32768:
                    for fi10 := 0; fi10 < 17; fi10++ {
                        if p1 {
                            p0 = (p0 / ((2147483647) | 1))
                            if ((p0 > 152) || p1) {
                                fmt.Println(p0)
                            }
                        } else {
                            if p1 {
                                var v11 bool = (42 > 294)
                                var v12 int = fi10
                                switch "a" {
                                case "a":
                                    v11 = v11
                                case "hello":
                                    for fi13 := 0; fi13 < 17; fi13++ {
                                        var v14 uint16 = 901
                                        _ = fi13
                                        _ = v14
                                    }
                                case "\n\t":
                                    if (4294967295 < 935) {
                                        v11 = (-2.25 <= 0.1)
                                        var v15 int8 = 42
                                        _ = v15
                                    } else {
                                        for fi16 := 0; fi16 < 19; fi16++ {
                                            if v11 {
                                                fmt.Println(v8, v12)
                                            } else {
                                                v11 = (!p1)
                                            }
                                            _ = fi16
                                        }
                                    }
                                }
                                _ = v11
                                _ = v12
                            }
                            switch "foo bar" {
                            case "\n\t":
                                var v17 uint8 = 42
                                _ = v17
                            default:
                                if true {
                                    v18 := GenS1{F0: 0, F1: 1}
                                    v19 := GenS2{F0: 127, F1: 2}
                                    for fi20 := 0; fi20 < 7; fi20++ {
                                        p1 = (!(-2.25 <= 3.14159))
                                        fmt.Println(fi20, p0)
                                        v8 -= (v8 >> ((4294967295) & 63))
                                        _ = fi20
                                    }
                                    _ = v18
                                    _ = v19
                                } else {
                                    v21 := GenS1{F0: 1, F1: (v8 & v8)}
                                    _ = v21
                                }
                            }
                            fmt.Println(p1, p1, p1)
                        }
                        for fi22 := 0; fi22 < 16; fi22++ {
                            fmt.Println(v8, v8)
                            _ = fi22
                        }
                        _ = fi10
                    }
                case 32767:
                    fmt.Println(p1, p1)
                default:
                    v23 := []int{-2147483648, -9223372036854775808}
                    _ = v23
                }
            }
        case 127:
            if ("\n\t" == "\n\t") {
                fmt.Println(p1, v8)
            }
        case -137:
            var v24 float64 = ((1.5 - 0.1) * (-2.25 + (3.14159 * -2.25)))
            _ = v24
        default:
            v25 := GenS1{F0: 18446744073709551615, F1: (v8 - (v8 & 889))}
            _ = v25
        }
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v8
    return 42
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int64 = 1
    v3 := []int{859, -2147483648, 342, 4294967295}
    var v4 uint16 = 2
    v5 := []int{42}
    var v6 bool = true
    for fi26 := 0; fi26 < 7; fi26++ {
        v2 = -9223372036854775808
        var v27 uint16 = (v4 << ((797) & 63))
        v28 := GenS1{F0: v1, F1: 65535}
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println(v1)
    for fi29 := 0; fi29 < 6; fi29++ {
        var v30 uint = 293
        switch 4294967295 {
        case 4294967295:
            var v31 float64 = float64(0.1)
            _ = v31
        }
        fmt.Println(fi29, v1, v1)
        _ = fi29
        _ = v30
    }
    for fi32 := 0; fi32 < 14; fi32++ {
        switch fi32 {
        case -698:
            for fi33 := 0; fi33 < 5; fi33++ {
                switch (((((("foo bar" + "\n\t") + "foo bar") + "") + "\n\t") + "a") + (("" + ("" + ("hello" + (("\n\t" + "foo bar") + "hello")))) + "")) {
                case "\n\t":
                    var v34 int32 = -289
                    _ = v34
                }
                _ = fi33
            }
        case 128:
            _ = len(v3)
        case 1:
            v35 := GenS3{F0: v2, F1: 32767}
            _ = v35
        }
        switch (fi32 >> ((369) & 63)) {
        case 267:
            if v6 {
                var v36 uint8 = 205
                v3 = append(v3, (fi32 ^ fi32))
                _ = v36
            } else {
                if (((v4 << ((582) & 63)) > v4) && (v6 && v6)) {
                    switch fi32 {
                    case 2147483647:
                        if (v6 && v6) {
                            _ = len(v3)
                            for fi37 := 0; fi37 < 6; fi37++ {
                                switch "" {
                                case "\n\t":
                                    v2 += (v2 + v2)
                                default:
                                    if (!((v1 % (((v1 << ((2) & 63))) | 1)) != (v1 | v1))) {
                                        var v38 float64 = float64((1.5 - -2.25))
                                        if true {
                                            if ((((((("hello" + "foo bar") + "foo bar") + "foo bar") + ("" + "foo bar")) + ("\n\t" + "foo bar")) + ("\n\t" + (("foo bar" + "foo bar") + ""))) != ("" + "foo bar")) {
                                                switch 32768 {
                                                case -603:
                                                    _ = len(v5)
                                                case 20:
                                                    v38 -= v38
                                                }
                                                v5 = append(v5, (fi37 >> ((666) & 63)))
                                                v39 := GenS2{F0: 100, F1: v1}
                                                _ = v39
                                            }
                                            v38 = v38
                                            switch 4294967295 {
                                            case -842:
                                                if (738 <= 42) {
                                                    fmt.Println(fi37, v2, v1)
                                                    v40 := GenS3{F0: 891, F1: 849}
                                                    v5 = append(v5, (fi32 % (((^(fi37))) | 1)))
                                                    _ = v40
                                                }
                                            case 555:
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            }
                                        }
                                        _ = v38
                                    }
                                    switch fi32 {
                                    case -214:
                                        if (1e10 < 0.1) {
                                            v6 = (!(!v6))
                                            v41 := GenS4{F0: 4294967295}
                                            _ = v41
                                        } else {
                                            var v42 int = 788
                                            if len(v3) > 0 {
                                                v3[0] = v42
                                            }
                                            _ = v42
                                        }
                                    case 248:
                                        fmt.Println(fi32, v6, fi37)
                                    case 32768:
                                        switch (^(fi32)) {
                                        case 382:
                                            switch (fi37 << ((18446744073709551615) & 63)) {
                                            case 127:
                                                var v43 float64 = -2.25
                                                _ = v43
                                            case -9223372036854775808:
                                                v3 = append(v3, fi32)
                                            }
                                        case 127:
                                            var v44 float64 = 0.1
                                            _ = v44
                                        }
                                    default:
                                        fmt.Println(v4)
                                    }
                                    var v45 string = ("a" + "foo bar")
                                    _ = v45
                                }
                                if v6 {
                                    v46 := GenS4{F0: 255}
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = v46
                                } else {
                                    if v6 {
                                        var v47 int = 2147483647
                                        var v48 int32 = 1
                                        if v6 {
                                            var v49 string = "foo bar"
                                            if v6 {
                                                if false {
                                                    fmt.Println(v1)
                                                    if ((v48 * v48) == v48) {
                                                        v47 = v47
                                                        v50 := GenS4{F0: 255}
                                                        v2 += v2
                                                        _ = v50
                                                    } else {
                                                        if len(v5) > 0 {
                                                            v5[0] = (v47 >> ((0) & 63))
                                                        }
                                                    }
                                                    v48 = v48
                                                }
                                                var v51 float32 = (3.14159 * 0.0)
                                                fmt.Println(v2, v2)
                                                _ = v51
                                            }
                                            if (2 > 42) {
                                                switch (-(fi37)) {
                                                case -585:
                                                    v48 += (v48 ^ (v48 / ((v48) | 1)))
                                                case -622:
                                                    var v52 uint8 = 42
                                                    _ = v52
                                                case 32767:
                                                    fmt.Println(fi32)
                                                }
                                                v2 -= v2
                                            } else {
                                                var v53 float32 = -2.25
                                                if v6 {
                                                    if v6 {
                                                        v54 := GenS5{F0: 42, F1: 65535}
                                                        _ = v54
                                                    }
                                                } else {
                                                    v55 := GenS6{F0: 512, F1: 42, F2: (v48 | (v48 / ((1) | 1)))}
                                                    var v56 int16 = -32768
                                                    _ = v55
                                                    _ = v56
                                                }
                                                _ = v53
                                            }
                                            _ = v49
                                        }
                                        _ = v47
                                        _ = v48
                                    }
                                    v57 := GenS6{F0: 65535, F1: fi37, F2: -890}
                                    _ = v57
                                }
                                _ = fi37
                            }
                        } else {
                            v58 := GenS3{F0: v2, F1: 1}
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = v58
                        }
                    case 127:
                        for fi59 := 0; fi59 < 11; fi59++ {
                            v60 := GenS2{F0: 0, F1: v1}
                            _ = fi59
                            _ = v60
                        }
                    }
                }
                v2 = 1
                var v61 uint = 500
                _ = v61
            }
        case -1:
            switch (fi32 + fi32) {
            case 128:
                switch "\n\t" {
                case "foo bar":
                    switch "a" {
                    case "":
                        for fi62 := 0; fi62 < 18; fi62++ {
                            _ = len(v3)
                            var v63 int32 = int32(v2)
                            if v6 {
                                var v64 int16 = -32768
                                switch ((("\n\t" + "foo bar") + ("" + "a")) + ("hello" + ("a" + "\n\t"))) {
                                case "\n\t":
                                    var v65 uint = 400
                                    _ = v65
                                case "foo bar":
                                    v6 = v6
                                }
                                switch "foo bar" {
                                case "hello":
                                    v66 := GenS6{F0: 65535, F1: 32768, F2: (v63 - v63)}
                                    _ = v66
                                case "a":
                                    v2 += 1
                                default:
                                    fmt.Println(v4, fi32)
                                }
                                _ = v64
                            }
                            if (!(2 == 1)) {
                                if ((true && v6) && v6) {
                                    v1 += 569
                                }
                                fmt.Println(v2, v6, v4)
                                v1 = (^(v1))
                            } else {
                                var v67 int32 = -1
                                v68 := GenS4{F0: 517}
                                _ = v67
                                _ = v68
                            }
                            if ((!false) || (((0.0 + 0.1) == (1.5 + (1.5 / 1e10))) || (false && (!v6)))) {
                                if v6 {
                                    fmt.Println(fi62)
                                    var v69 int16 = 1
                                    _ = len(v5)
                                    _ = v69
                                }
                                switch "" {
                                case "foo bar":
                                    var v70 int16 = 42
                                    _ = v70
                                default:
                                    v1 += v1
                                }
                                var v71 int8 = 1
                                _ = v71
                            }
                            _ = fi62
                            _ = v63
                        }
                    case "a":
                        fmt.Println(v4)
                    default:
                        if (2 > 2) {
                            for fi72 := 0; fi72 < 9; fi72++ {
                                v4 += (v4 - genFn0((v2 >> ((172) & 63))))
                                if len(v5) > 0 {
                                    v5[0] = 42
                                }
                                var v73 int64 = v2
                                _ = fi72
                                _ = v73
                            }
                            for fi74 := 0; fi74 < 3; fi74++ {
                                fmt.Println(fi74, v2, fi32)
                                _ = len(v3)
                                _ = fi74
                            }
                            var v75 uint64 = v1
                            _ = v75
                        }
                    }
                case "\n\t":
                    var v76 uint64 = v1
                    _ = v76
                }
            case 127:
                var v77 bool = false
                _ = v77
            case 32767:
                v1 = 4294967295
            }
        case 32768:
            if len(v3) > 0 {
                v3[0] = -9223372036854775808
            }
        default:
            v5 = append(v5, (fi32 & (fi32 + 32767)))
        }
        switch (fi32 << ((735) & 63)) {
        case -759:
            fmt.Println(v4)
        case 32767:
            for fi78 := 0; fi78 < 9; fi78++ {
                var v79 int16 = 1
                v80 := GenS3{F0: -107, F1: v79}
                if (-2147483648 <= 42) {
                    var v81 uint = 0
                    v1 += (v1 >> (((^(v81))) & 63))
                    v82 := GenS2{F0: 127, F1: v1}
                    _ = v81
                    _ = v82
                }
                _ = fi78
                _ = v79
                _ = v80
            }
        case 1000:
            var v83 bool = (!(genFn0(-1) > v4))
            _ = v83
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        _ = fi32
    }
    switch 128 {
    case 127:
        for fi84 := 0; fi84 < 8; fi84++ {
            var v85 int32 = 0
            _ = fi84
            _ = v85
        }
    case -2147483648:
        var v86 uint32 = 4294967295
        _ = v86
    case 2147483647:
        switch "hello" {
        case "":
            v4 = genFn0((v2 / ((1) | 1)))
        case "foo bar":
            v3 = append(v3, 128)
        default:
            v87 := GenS5{F0: 78, F1: 465}
            _ = v87
        }
    }
    v6 = (!(!v6))
    v4 -= 42
    if (3.14159 == (0.1 + -2.25)) {
        switch 573 {
        case 2147483647:
            var v88 uint64 = (v1 >> ((1) & 63))
            _ = v88
        case -2147483648:
            if v6 {
                if v6 {
                    fmt.Println(v2, v4)
                    v4 -= genFn0(v2)
                }
                v89 := GenS2{F0: -128, F1: (v1 + v1)}
                if v6 {
                    if (-698 >= -9223372036854775808) {
                        for fi90 := 0; fi90 < 8; fi90++ {
                            for fi91 := 0; fi91 < 15; fi91++ {
                                var v92 uint32 = 2
                                v93 := GenS1{F0: 18446744073709551615, F1: genFn0((v2 << ((518) & 63)))}
                                var v94 uint8 = 2
                                _ = fi91
                                _ = v92
                                _ = v93
                                _ = v94
                            }
                            var v95 bool = ((!(v1 <= (v1 | 392))) && ((("hello" < (("\n\t" + "foo bar") + "")) || v6) && (!(0 != 127))))
                            _ = fi90
                            _ = v95
                        }
                        v96 := GenS7{F0: 1000}
                        if false {
                            var v97 bool = ((v6 && false) || (!v6))
                            v98 := GenS8{F0: 42}
                            var v99 float32 = -2.25
                            _ = v97
                            _ = v98
                            _ = v99
                        } else {
                            var v100 float64 = (float64(float32(0.0)) / 1.0)
                            _ = v100
                        }
                        _ = v96
                    } else {
                        v101 := GenS8{F0: 42}
                        _ = v101
                    }
                    for fi102 := 0; fi102 < 1; fi102++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi102
                    }
                }
                _ = v89
            } else {
                for fi103 := 0; fi103 < 7; fi103++ {
                    fmt.Println(v1, fi103)
                    var v104 float32 = -2.25
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi103
                    _ = v104
                }
                v105 := GenS9{F0: 1000, F1: 523, F2: 65535}
                _ = v105
            }
        case 42:
            var v106 uint8 = 53
            _ = v106
        default:
            v107 := GenS9{F0: -177, F1: 526, F2: 102}
            _ = v107
        }
    } else {
        fmt.Println(v1, v4, v4)
        if v6 {
            if v6 {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch -801 {
                case 32767:
                    var v108 uint8 = 42
                    _ = v108
                case 42:
                    for fi109 := 0; fi109 < 16; fi109++ {
                        v1 -= (v1 ^ 2)
                        if true {
                            var v110 uint = 65535
                            if (32767 == 1) {
                                switch 4294967295 {
                                case -822:
                                    for fi111 := 0; fi111 < 1; fi111++ {
                                        fmt.Println(v1)
                                        v1 = v1
                                        _ = fi111
                                    }
                                case 128:
                                    if v6 {
                                        v2 = (v2 | v2)
                                    } else {
                                        v1 += 18446744073709551615
                                        var v112 uint64 = 2
                                        _ = v112
                                    }
                                case -401:
                                    if v6 {
                                        if (!(!true)) {
                                            _ = len(v5)
                                        }
                                        fmt.Println(v2, fi109)
                                        var v113 bool = (52 != (v110 % ((v110) | 1)))
                                        _ = v113
                                    } else {
                                        var v114 string = ""
                                        fmt.Println(fi109)
                                        _ = v114
                                    }
                                default:
                                    v115 := GenS10{F0: v6}
                                    _ = v115
                                }
                            } else {
                                v6 = true
                            }
                            _ = len(v5)
                            _ = v110
                        }
                        v6 = v6
                        _ = fi109
                    }
                default:
                    for fi116 := 0; fi116 < 1; fi116++ {
                        switch fi116 {
                        case 2147483647:
                            var v117 float64 = float64((1.5 - -2.25))
                            _ = v117
                        case 128:
                            var v118 float32 = ((1e10 + 1e10) * 1e10)
                            _ = v118
                        case 32767:
                            if len(v5) > 0 {
                                v5[0] = (fi116 % ((fi116) | 1))
                            }
                        default:
                            switch (fi116 - fi116) {
                            case 128:
                                switch (fi116 % ((fi116) | 1)) {
                                case 971:
                                    v4 += (^(v4))
                                default:
                                    v6 = v6
                                }
                            case -2147483648:
                                v3 = append(v3, (fi116 ^ fi116))
                            default:
                                for fi119 := 0; fi119 < 1; fi119++ {
                                    switch (fi116 % (((^(fi116))) | 1)) {
                                    case 4294967295:
                                        v6 = v6
                                    default:
                                        v120 := GenS11{F0: fi116, F1: 119}
                                        _ = v120
                                    }
                                    _ = fi119
                                }
                            }
                        }
                        if (!(42 <= 58)) {
                            var v121 int32 = -734
                            _ = v121
                        } else {
                            v4 = v4
                            switch (fi116 ^ 1) {
                            case 557:
                                v1 = 1
                            case -1:
                                v122 := GenS12{F0: 0, F1: (100 > -1), F2: (v1 / ((v1) | 1))}
                                _ = v122
                            default:
                                fmt.Println(fi116, fi116)
                            }
                        }
                        _ = fi116
                    }
                }
                v3 = append(v3, 1)
            }
            var v123 float64 = float64(0.1)
            if ((!false) && (!v6)) {
                switch 556 {
                case 4294967295:
                    fmt.Println(v123)
                case 0:
                    fmt.Println(v2)
                default:
                    for fi124 := 0; fi124 < 2; fi124++ {
                        if (float32(((v123 + (v123 * (v123 / v123))) / (0.1 + ((v123 * (1.5 + (1.5 * ((v123 / (((0.1 / (v123 - (float64(0.0) / v123))) - ((float64(1.5) - v123) / (v123 / v123))) + (1e10 / ((0.1 - float64((1.5 * -2.25))) + v123)))) + (float64(1e10) / v123))))) + float64(((((1.5 * ((float32((v123 + (((0.1 / (v123 / 1e10)) + v123) + ((0.1 / 0.1) - v123)))) / float32(-2.25)) - 3.14159)) / 3.14159) + 1e10) / 1.5)))))) == 3.14159) {
                            for fi125 := 0; fi125 < 18; fi125++ {
                                var v126 float64 = v123
                                v127 := GenS12{F0: 2147483647, F1: false, F2: (^(v1))}
                                _ = fi125
                                _ = v126
                                _ = v127
                            }
                        }
                        v128 := GenS4{F0: 65535}
                        v129 := GenS13{F0: v1, F1: -19, F2: (fi124 ^ (fi124 & (fi124 - 4294967295)))}
                        _ = fi124
                        _ = v128
                        _ = v129
                    }
                }
            } else {
                fmt.Println(v2, v123)
                v130 := GenS7{F0: 806}
                v1 += 2
                _ = v130
            }
            _ = v123
        } else {
            switch "hello" {
            case "hello":
                var v131 uint64 = v1
                _ = v131
            default:
                v4 = 1
            }
            v5 = append(v5, -656)
        }
        switch "" {
        case "":
            for fi132 := 0; fi132 < 2; fi132++ {
                if (9223372036854775807 > v2) {
                    v6 = (42 >= -2)
                    v133 := GenS14{F0: "foo bar", F1: 128}
                    _ = v133
                } else {
                    v4 -= v4
                }
                _ = fi132
            }
        default:
            v134 := GenS15{F0: (^(v4)), F1: 255, F2: "a"}
            _ = v134
        }
    }
    if (4294967295 == 42) {
        fmt.Println(v1, v1, v1)
    }
    var v135 uint64 = (v1 >> ((42) & 63))
    _ = len(v5)
    var v136 int32 = -156
    for fi137 := 0; fi137 < 8; fi137++ {
        for fi138 := 0; fi138 < 4; fi138++ {
            var v139 float64 = float64((3.14159 / 1.0))
            if v6 {
                var v140 int64 = v2
                _ = v140
            } else {
                var v141 bool = true
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch (fi138 - (-(fi138))) {
                case -9223372036854775808:
                    if true {
                        v142 := GenS7{F0: -38}
                        if true {
                            var v143 string = "\n\t"
                            v6 = ((v2 & (v2 | -9223372036854775808)) < 1)
                            _ = len(v5)
                            _ = v143
                        } else {
                            v139 -= v139
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                        if (0 == (v1 ^ (^(v1)))) {
                            if (((100 == 127) && (-1 >= genFn1(2147483647, (874 > 4294967295)))) || v6) {
                                v6 = v6
                                fmt.Println(v139, v6, v136)
                                v135 = 2
                            }
                        }
                        v144 := GenS16{F0: -1, F1: 825, F2: (^(v136))}
                        _ = v142
                        _ = v144
                    }
                }
                _ = v141
            }
            var v145 int32 = v136
            _ = fi138
            _ = v139
            _ = v145
        }
        if ((!v6) && (100 != 100)) {
            v136 += int32(v135)
            if v6 {
                var v146 string = ""
                var v147 uint32 = 2147483647
                _ = v146
                _ = v147
            } else {
                fmt.Println(v1, v4, v136)
            }
            v148 := GenS12{F0: 2147483647, F1: v6, F2: (^(v135))}
            _ = v148
        } else {
            v136 = v136
            if len(v3) > 0 {
                v3[0] = (fi137 + (fi137 >> ((42) & 63)))
            }
            v135 = (^(v135))
        }
        var v149 uint = 255
        _ = fi137
        _ = v149
    }
    fmt.Println("v1:", v1)
    fmt.Println("v135:", v135)
    fmt.Println("v136:", v136)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v135
    _ = v136
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
