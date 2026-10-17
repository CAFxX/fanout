package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint8 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 uint8; F1 uint16; F2 uint }
type GenS5 struct { F0 uint8; F1 int64; F2 int32 }
type GenS6 struct { F0 int8; F1 uint; F2 string }
type GenS7 struct { F0 uint8; F1 uint; F2 uint }
type GenS8 struct { F0 int64; F1 int8; F2 uint64 }
type GenS9 struct { F0 bool; F1 int32 }
type GenS10 struct { F0 string; F1 uint64 }
type GenS11 struct { F0 uint32; F1 uint; F2 string }
type GenS12 struct { F0 uint32; F1 uint32 }
func genFn0(p0 int32) uint8 {
    v4 := []int{0, 32767, 0}
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := GenS1{F0: (int16(p0) < 0), F1: fi5, F2: 222}
        if len(v4) > 0 {
            v4[0] = fi5
        }
        _ = len(v4)
        _ = fi5
        _ = v6
    }
    var v7 string = "a"
    if (!true) {
        if false {
            for fi8 := 0; fi8 < 8; fi8++ {
                v9 := GenS1{F0: true, F1: fi8, F2: 0}
                v7 = v7
                _ = fi8
                _ = v9
            }
            var v10 uint8 = 0
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v10
        }
        v11 := GenS1{F0: (42 > -37), F1: 2147483647, F2: 255}
        _ = v11
    }
    switch v7 {
    case "foo bar":
        if len(v4) > 0 {
            v4[0] = 1
        }
    case "\n\t":
        v12 := GenS1{F0: (1000 >= 42), F1: -9223372036854775808, F2: 1}
        _ = v12
    case "a":
        switch "\n\t" {
        case "a":
            for fi13 := 0; fi13 < 18; fi13++ {
                switch (fi13 >> ((18446744073709551615) & 63)) {
                case 1:
                    if (19 > -1) {
                        if (724 == 254) {
                            var v14 float32 = (0.1 * 1e10)
                            _ = v14
                        } else {
                            for fi15 := 0; fi15 < 6; fi15++ {
                                v7 = v7
                                _ = fi15
                            }
                            switch fi13 {
                            case 2147483647:
                                p0 -= 0
                            default:
                                fmt.Println(p0, fi13, p0)
                            }
                            v7 = v7
                        }
                        p0 = (-(p0))
                        for fi16 := 0; fi16 < 19; fi16++ {
                            v7 = ""
                            v17 := GenS2{F0: ("foo bar" + ("hello" + "foo bar")), F1: 2}
                            v7 = ("foo bar" + v7)
                            _ = fi16
                            _ = v17
                        }
                    } else {
                        v18 := GenS2{F0: "hello", F1: 0}
                        _ = v18
                    }
                case 605:
                    v7 = "foo bar"
                case -1:
                    var v19 int8 = 1
                    _ = v19
                }
                var v20 int16 = -32768
                v21 := GenS3{F0: 2, F1: (v20 - v20)}
                _ = fi13
                _ = v20
                _ = v21
            }
        case "hello":
            if len(v4) > 0 {
                v4[0] = 0
            }
        case "\n\t":
            if len(v4) > 0 {
                v4[0] = -1
            }
        default:
            v7 = v7
        }
    }
    _ = p0
    _ = v4
    _ = v7
    return 42
}
func genFn1(p0 uint64, p1 int, p2 uint) uint8 {
    p2 = (^(p2))
    for fi22 := 0; fi22 < 20; fi22++ {
        switch 32768 {
        case 736:
            for fi23 := 0; fi23 < 4; fi23++ {
                switch "a" {
                case "foo bar":
                    switch "a" {
                    case "":
                        if true {
                            if false {
                                fmt.Println(fi23)
                            } else {
                                p2 = (p2 + (p2 << ((0) & 63)))
                                v24 := GenS4{F0: 203, F1: 1, F2: (^(p2))}
                                p0 -= 4294967295
                                _ = v24
                            }
                            v25 := []int{}
                            _ = v25
                        }
                    case "hello":
                        p1 += fi22
                    case "foo bar":
                        switch ("\n\t" + ((((("" + ("foo bar" + "a")) + "") + ("a" + (("hello" + "hello") + ("\n\t" + "hello")))) + "\n\t") + ("hello" + ("" + "\n\t")))) {
                        case "foo bar":
                            v26 := []int{(^(fi22)), 128, fi22, fi23}
                            _ = v26
                        }
                    default:
                        v27 := GenS1{F0: false, F1: (fi23 * p1), F2: 0}
                        _ = v27
                    }
                case "":
                    var v28 float64 = 0.0
                    _ = v28
                case "hello":
                    switch ("" + (("foo bar" + "foo bar") + "hello")) {
                    case "":
                        p2 = 292
                    case "foo bar":
                        if (65535 != 0) {
                            v29 := GenS5{F0: 255, F1: 9223372036854775807, F2: 2147483647}
                            var v30 float32 = 3.14159
                            _ = v29
                            _ = v30
                        } else {
                            p0 += 825
                            p2 += p2
                            fmt.Println(fi22)
                            v31 := []int{}
                            p0 = p0
                            var v32 int = (fi22 >> ((4294967295) & 63))
                            _ = v31
                            _ = v32
                        }
                    case "\n\t":
                        v33 := []int{(fi23 | p1), fi23, 128}
                        _ = v33
                    }
                }
                p0 = (p0 * (p0 >> (((p2 * 381)) & 63)))
                var v34 int = (fi23 % ((333) | 1))
                _ = fi23
                _ = v34
            }
        case 42:
            if true {
                for fi35 := 0; fi35 < 12; fi35++ {
                    v36 := GenS5{F0: 0, F1: -159, F2: -2147483648}
                    p2 = 547
                    _ = fi35
                    _ = v36
                }
                fmt.Println(p1)
                var v37 uint16 = 42
                _ = v37
            }
        default:
            var v38 uint32 = 0
            _ = v38
        }
        _ = fi22
    }
    switch "foo bar" {
    case "hello":
        if (!false) {
            p2 += p2
            p2 -= p2
            var v39 string = ("a" + ("foo bar" + "foo bar"))
            _ = v39
        } else {
            if (-89 < 42) {
                v40 := []int{(p1 * (p1 ^ p1))}
                p2 -= p2
                _ = v40
            }
        }
    case "\n\t":
        if (-128 <= 0) {
            if true {
                fmt.Println(p2, p0, p0)
                v41 := GenS6{F0: 42, F1: p2, F2: "foo bar"}
                _ = v41
            } else {
                v42 := []int{-9223372036854775808}
                _ = v42
            }
            for fi43 := 0; fi43 < 15; fi43++ {
                switch 32767 {
                case -445:
                    var v44 uint8 = 214
                    _ = v44
                case -12:
                    switch (p1 / ((fi43) | 1)) {
                    case 677:
                        if true {
                            if false {
                                var v45 int64 = 1
                                _ = v45
                            } else {
                                switch 0 {
                                case -1:
                                    var v46 uint64 = 690
                                    _ = v46
                                case 128:
                                    p0 += 208
                                case -9223372036854775808:
                                    v47 := GenS7{F0: 2, F1: 884, F2: p2}
                                    _ = v47
                                default:
                                    switch ("" + ("hello" + "a")) {
                                    case "foo bar":
                                        p2 = 65535
                                    }
                                }
                                p1 += 128
                                var v48 int = 602
                                _ = v48
                            }
                            v49 := []int{fi43, -89, fi43}
                            _ = v49
                        } else {
                            for fi50 := 0; fi50 < 4; fi50++ {
                                v51 := GenS1{F0: (9223372036854775807 >= -9223372036854775808), F1: (-(fi50)), F2: 127}
                                if (false && (!(1 == 2))) {
                                    var v52 string = "hello"
                                    var v53 int32 = 42
                                    if (0 != 100) {
                                        v54 := GenS7{F0: 110, F1: 944, F2: (p2 >> ((p2) & 63))}
                                        if false {
                                            var v55 uint8 = 0
                                            _ = v55
                                        }
                                        _ = v54
                                    }
                                    v52 = ("" + "\n\t")
                                    _ = v52
                                    _ = v53
                                } else {
                                    v56 := GenS8{F0: -9223372036854775808, F1: 0, F2: p0}
                                    var v57 uint = (p2 >> (((p2 * p2)) & 63))
                                    _ = v56
                                    _ = v57
                                }
                                _ = fi50
                                _ = v51
                            }
                        }
                    case 32768:
                        for fi58 := 0; fi58 < 18; fi58++ {
                            p2 -= p2
                            _ = fi58
                        }
                    case -9223372036854775808:
                        var v59 int8 = 42
                        _ = v59
                    }
                case -2147483648:
                    var v60 uint64 = (^(p0))
                    _ = v60
                default:
                    p2 += p2
                }
                var v61 bool = false
                _ = fi43
                _ = v61
            }
        } else {
            if (1 < genFn0(657)) {
                if (446 >= 65535) {
                    v62 := []int{(p1 - -2147483648), p1}
                    var v63 float64 = -2.25
                    for fi64 := 0; fi64 < 11; fi64++ {
                        var v65 uint64 = p0
                        if (false && (-797 > 2147483647)) {
                            v62 = append(v62, (p1 >> (((p2 + (p2 & p2))) & 63)))
                            if len(v62) >= 2 {
                                v62 = v62[1:len(v62)]
                            }
                            var v66 uint8 = 127
                            _ = v66
                        }
                        for fi67 := 0; fi67 < 11; fi67++ {
                            _ = len(v62)
                            if (float32((((float64(0.0) + (((1.5 + (((v63 - float64(((-2.25 * ((1e10 - float32(v63)) + (3.14159 - 0.0))) - ((3.14159 / ((3.14159 + 0.1) * (0.0 + float32((0.1 * -2.25))))) + (-2.25 / 1.0))))) * float64((float32(v63) - ((-2.25 + -2.25) + 3.14159)))) / v63)) + float64(1e10)) + 1.5)) + v63) - v63)) != float32(v63)) {
                                if (4294967295 > 269) {
                                    fmt.Println(p0)
                                    v68 := GenS2{F0: "hello", F1: 128}
                                    var v69 int8 = -1
                                    _ = v68
                                    _ = v69
                                } else {
                                    var v70 int = fi64
                                    var v71 string = ""
                                    if (2147483647 != 2) {
                                        if false {
                                            p0 = (v65 | p0)
                                            v72 := GenS3{F0: 75, F1: 807}
                                            v73 := GenS2{F0: v71, F1: 64}
                                            _ = v72
                                            _ = v73
                                        } else {
                                            if false {
                                                p0 += (v65 << ((p2) & 63))
                                            }
                                            var v74 uint8 = 228
                                            fmt.Println(v71)
                                            _ = v74
                                        }
                                        var v75 int = fi67
                                        _ = v75
                                    } else {
                                        if false {
                                            fmt.Println(v70, fi64, p0)
                                        } else {
                                            var v76 int8 = 127
                                            if len(v62) > 0 {
                                                v62[0] = 128
                                            }
                                            _ = v76
                                        }
                                        if (!true) {
                                            if (289 == 372) {
                                                var v77 bool = (255 != 2)
                                                fmt.Println(v77, v63, p2)
                                                if len(v62) >= 2 {
                                                    v62 = v62[1:len(v62)]
                                                }
                                                _ = v77
                                            }
                                        } else {
                                            fmt.Println(fi67)
                                            if len(v62) >= 2 {
                                                v62 = v62[1:len(v62)]
                                            }
                                        }
                                    }
                                    v62 = append(v62, -294)
                                    if true {
                                        switch 897 {
                                        case 9223372036854775807:
                                            if (374 <= (p0 ^ p0)) {
                                                v71 = (((((("\n\t" + "hello") + (v71 + "\n\t")) + "a") + "foo bar") + "a") + (v71 + v71))
                                            } else {
                                                if true {
                                                    if (4294967295 == p0) {
                                                        v70 += (fi64 >> ((0) & 63))
                                                        switch ((("" + v71) + "") + (v71 + v71)) {
                                                        case "":
                                                            v78 := GenS4{F0: genFn0(-1), F1: 2, F2: uint(p0)}
                                                            if len(v62) >= 2 {
                                                                v62 = v62[1:len(v62)]
                                                            }
                                                            var v79 uint8 = 66
                                                            _ = v78
                                                            _ = v79
                                                        case "\n\t":
                                                            v80 := GenS4{F0: 123, F1: 65535, F2: (p2 * p2)}
                                                            _ = v80
                                                        case "foo bar":
                                                            var v81 float64 = float64((float32(v63) / -2.25))
                                                            _ = v81
                                                        default:
                                                            var v82 uint16 = 928
                                                            _ = v82
                                                        }
                                                    }
                                                    _ = len(v62)
                                                    var v83 string = ("\n\t" + ("hello" + v71))
                                                    _ = v83
                                                }
                                            }
                                        default:
                                            var v84 int64 = 9223372036854775807
                                            _ = v84
                                        }
                                    }
                                    _ = v70
                                    _ = v71
                                }
                                if (!true) {
                                    v85 := GenS7{F0: 183, F1: (^(p2)), F2: p2}
                                    switch (fi67 * 2147483647) {
                                    case 9223372036854775807:
                                        var v86 int16 = 140
                                        _ = v86
                                    case 971:
                                        fmt.Println(v65)
                                    }
                                    if true {
                                        v63 = float64(0.0)
                                        if (fi64 == fi67) {
                                            var v87 uint = (^(p2))
                                            p0 = 18446744073709551615
                                            var v88 uint32 = 42
                                            _ = v87
                                            _ = v88
                                        } else {
                                            v89 := GenS3{F0: 0, F1: -5}
                                            if len(v62) >= 2 {
                                                v62 = v62[1:len(v62)]
                                            }
                                            _ = v89
                                        }
                                    }
                                    fmt.Println(v65)
                                    _ = v85
                                } else {
                                    v90 := GenS4{F0: 2, F1: 0, F2: p2}
                                    if (!(0 > 42)) {
                                        v91 := GenS8{F0: -745, F1: -92, F2: (v65 << ((p2) & 63))}
                                        _ = v91
                                    } else {
                                        v92 := GenS7{F0: genFn0(-1), F1: p2, F2: (p2 & p2)}
                                        switch "hello" {
                                        case "hello":
                                            if (fi64 == (fi64 & (fi64 & (-(fi67))))) {
                                                var v93 uint = (p2 - p2)
                                                _ = v93
                                            }
                                        case "a":
                                            p2 = p2
                                        default:
                                            if (-1 < -1) {
                                                var v94 int16 = 1
                                                var v95 int16 = (v94 << ((42) & 63))
                                                _ = v94
                                                _ = v95
                                            } else {
                                                if (!false) {
                                                    if len(v62) >= 2 {
                                                        v62 = v62[1:len(v62)]
                                                    }
                                                    switch fi67 {
                                                    case 2147483647:
                                                        switch "hello" {
                                                        case "a":
                                                            v96 := GenS9{F0: (false && (273 == 608)), F1: 0}
                                                            _ = v96
                                                        }
                                                    case 473:
                                                        p1 += fi64
                                                    case 728:
                                                        var v97 bool = (!false)
                                                        _ = v97
                                                    }
                                                    var v98 int8 = 1
                                                    _ = v98
                                                } else {
                                                    fmt.Println(p2, p2, v65)
                                                    v62 = append(v62, (fi64 | fi64))
                                                    if true {
                                                        switch -2147483648 {
                                                        case 32768:
                                                            p0 = (^(p0))
                                                        case 127:
                                                            v62 = append(v62, -884)
                                                        }
                                                        v99 := GenS6{F0: 26, F1: p2, F2: ""}
                                                        if (-32 <= 42) {
                                                            p2 = p2
                                                        } else {
                                                            var v100 uint16 = 1
                                                            switch (p1 - (fi64 / ((-2147483648) | 1))) {
                                                            case 385:
                                                                v62 = append(v62, fi64)
                                                            case 4294967295:
                                                                var v101 int64 = 42
                                                                _ = v101
                                                            }
                                                            _ = v100
                                                        }
                                                        _ = v99
                                                    } else {
                                                        fmt.Println(p1, p2, v63)
                                                        _ = len(v62)
                                                    }
                                                }
                                                fmt.Println(v63, fi67)
                                            }
                                        }
                                        fmt.Println(p0, p0)
                                        _ = v92
                                    }
                                    _ = v90
                                }
                                if true {
                                    var v102 bool = (!true)
                                    fmt.Println(v63)
                                    if len(v62) >= 2 {
                                        v62 = v62[1:len(v62)]
                                    }
                                    _ = v102
                                }
                            } else {
                                v103 := GenS10{F0: "", F1: p0}
                                switch (fi64 % ((fi64) | 1)) {
                                case -817:
                                    fmt.Println(v63, v65)
                                case 127:
                                    switch "" {
                                    case "":
                                        v62 = append(v62, (fi64 << ((p2) & 63)))
                                    case "\n\t":
                                        var v104 int32 = -2147483648
                                        _ = v104
                                    case "a":
                                        switch "" {
                                        case "a":
                                            var v105 string = ("hello" + "a")
                                            _ = v105
                                        default:
                                            if true {
                                                if false {
                                                    fmt.Println(p1)
                                                    if false {
                                                        v106 := GenS5{F0: 2, F1: -756, F2: -1}
                                                        _ = v106
                                                    }
                                                } else {
                                                    p2 -= (^(p2))
                                                }
                                                var v107 float32 = (3.14159 / 3.14159)
                                                fmt.Println(p1)
                                                _ = len(v62)
                                                _ = v107
                                            }
                                        }
                                    }
                                }
                                _ = v103
                            }
                            _ = fi67
                        }
                        _ = fi64
                        _ = v65
                    }
                    _ = v62
                    _ = v63
                } else {
                    v108 := []int{p1, (p1 + p1)}
                    _ = v108
                }
            }
            p1 += (-(p1))
            fmt.Println(p1, p2, p1)
        }
    case "foo bar":
        switch p1 {
        case -9223372036854775808:
            if false {
                for fi109 := 0; fi109 < 7; fi109++ {
                    for fi110 := 0; fi110 < 3; fi110++ {
                        if (-348 >= fi109) {
                            v111 := []int{-185, 32768, p1, fi109}
                            var v112 int16 = 32767
                            switch "hello" {
                            case "":
                                fmt.Println(fi110)
                            case "\n\t":
                                switch fi109 {
                                case -274:
                                    if (false || false) {
                                        var v113 int = (fi109 % ((9223372036854775807) | 1))
                                        var v114 float64 = 1e10
                                        if (255 >= 1) {
                                            switch ("\n\t" + "a") {
                                            case "foo bar":
                                                switch p1 {
                                                case -513:
                                                    if false {
                                                        v115 := GenS11{F0: 811, F1: p2, F2: "\n\t"}
                                                        _ = v115
                                                    } else {
                                                        var v116 string = ((("hello" + "foo bar") + "\n\t") + "\n\t")
                                                        v117 := GenS5{F0: 0, F1: -9223372036854775808, F2: 2147483647}
                                                        var v118 uint16 = 2
                                                        _ = len(v111)
                                                        _ = v116
                                                        _ = v117
                                                        _ = v118
                                                    }
                                                case 1:
                                                    _ = len(v111)
                                                default:
                                                    if false {
                                                        p1 = 0
                                                    }
                                                }
                                                if false {
                                                    var v119 uint = p2
                                                    var v120 bool = false
                                                    _ = v119
                                                    _ = v120
                                                }
                                            }
                                            switch ("foo bar" + ("hello" + "foo bar")) {
                                            case "hello":
                                                p1 -= v113
                                            default:
                                                p2 = (p2 >> ((255) & 63))
                                            }
                                            p1 = fi109
                                            var v121 string = ("\n\t" + "hello")
                                            if len(v111) > 0 {
                                                v111[0] = (v113 / ((fi110) | 1))
                                            }
                                            _ = v121
                                        } else {
                                            p2 += p2
                                            switch 407 {
                                            case 32767:
                                                v114 += 1e10
                                            case -212:
                                                v114 += ((v114 + float64(3.14159)) + (v114 + ((v114 * (3.14159 - v114)) - (0.1 + -2.25))))
                                            default:
                                                if ((0 < -1) || (!(!false))) {
                                                    var v122 float64 = (v114 * v114)
                                                    if len(v111) > 0 {
                                                        v111[0] = 1
                                                    }
                                                    _ = v122
                                                } else {
                                                    _ = len(v111)
                                                }
                                            }
                                        }
                                        _ = v113
                                        _ = v114
                                    } else {
                                        switch -1000 {
                                        case 128:
                                            v123 := GenS8{F0: 9223372036854775807, F1: 1, F2: (p0 ^ uint64(p2))}
                                            _ = v123
                                        case -2147483648:
                                            v124 := GenS4{F0: uint8(p0), F1: 1, F2: 155}
                                            _ = v124
                                        case 42:
                                            var v125 int8 = 100
                                            _ = v125
                                        default:
                                            var v126 int8 = -1
                                            _ = v126
                                        }
                                    }
                                case 42:
                                    fmt.Println(fi110, fi109)
                                default:
                                    if (v112 >= (v112 >> ((2) & 63))) {
                                        var v127 bool = (!false)
                                        _ = len(v111)
                                        _ = v127
                                    }
                                }
                            }
                            _ = v111
                            _ = v112
                        }
                        _ = fi110
                    }
                    _ = fi109
                }
                v128 := GenS8{F0: -158, F1: -78, F2: 4294967295}
                _ = v128
            } else {
                p2 -= p2
                for fi129 := 0; fi129 < 16; fi129++ {
                    fmt.Println(fi129, p2)
                    _ = fi129
                }
                var v130 uint16 = 65535
                _ = v130
            }
        }
    default:
        v131 := GenS4{F0: 246, F1: 42, F2: p2}
        _ = v131
    }
    var v132 int8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v132
    return 1
}

func main() {
    v1 := []uint32{}
    v2 := []int{-1, -9223372036854775808}
    var v3 int64 = -990
    var v133 uint64 = 18446744073709551615
    for fi134 := 0; fi134 < 7; fi134++ {
        var v135 string = "foo bar"
        v3 = v3
        for fi136 := 0; fi136 < 20; fi136++ {
            fmt.Println(fi134, fi136)
            var v137 float64 = 0.1
            var v138 int16 = -863
            _ = fi136
            _ = v137
            _ = v138
        }
        _ = fi134
        _ = v135
    }
    if (!true) {
        v3 -= v3
        switch -894 {
        case 32767:
            var v139 int16 = 0
            _ = v139
        default:
            v140 := GenS8{F0: v3, F1: 1, F2: (v133 / ((917) | 1))}
            _ = v140
        }
    }
    v141 := GenS9{F0: true, F1: -1}
    var v142 uint16 = 1
    var v143 float32 = 3.14159
    v133 = v133
    v142 += 32767
    fmt.Println(v3)
    v144 := GenS7{F0: 157, F1: 0, F2: 316}
    v142 -= 255
    switch "\n\t" {
    case "":
        if false {
            for fi145 := 0; fi145 < 16; fi145++ {
                if true {
                    v133 = 551
                } else {
                    v3 += v3
                    v146 := GenS7{F0: 2, F1: 1, F2: 900}
                    var v147 float32 = (0.0 - ((float32((((float64(v143) + 1e10) - 1.5) / 1.0)) / (v143 * 1e10)) + -2.25))
                    _ = v146
                    _ = v147
                }
                fmt.Println(v142, v3)
                _ = fi145
            }
            v148 := GenS2{F0: "\n\t", F1: 128}
            _ = v148
        }
    default:
        if true {
            for fi149 := 0; fi149 < 2; fi149++ {
                v142 -= (v142 / (((v142 << ((4294967295) & 63))) | 1))
                if false {
                    for fi150 := 0; fi150 < 13; fi150++ {
                        var v151 int = (fi150 + (fi150 & 9223372036854775807))
                        v142 -= v142
                        _ = fi150
                        _ = v151
                    }
                    for fi152 := 0; fi152 < 5; fi152++ {
                        v143 -= ((((0.0 + float32(0.1)) - v143) - (v143 / 1.5)) - 1.5)
                        v1 = append(v1, 653)
                        _ = fi152
                    }
                    fmt.Println(fi149)
                } else {
                    if false {
                        v153 := GenS12{F0: 42, F1: 2147483648}
                        _ = v153
                    } else {
                        _ = len(v2)
                        var v154 uint = 18446744073709551615
                        v3 -= v3
                        _ = v154
                    }
                    v155 := GenS9{F0: true, F1: -43}
                    _ = v155
                }
                _ = fi149
            }
        } else {
            if len(v1) > 0 {
                v1[0] = 65535
            }
            var v156 string = "hello"
            _ = v156
        }
    }
    v142 = (v142 >> ((65535) & 63))
    v142 += v142
    for fi157 := 0; fi157 < 17; fi157++ {
        _ = len(v1)
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v143, v3)
        case "":
            var v158 uint = 790
            _ = v158
        }
        _ = fi157
    }
    switch ("hello" + (("\n\t" + ("" + "a")) + ("a" + ("hello" + "hello")))) {
    case "a":
        if true {
            var v159 uint8 = 232
            for fi160 := 0; fi160 < 17; fi160++ {
                for fi161 := 0; fi161 < 7; fi161++ {
                    var v162 int32 = 1
                    var v163 int = fi160
                    v2 = append(v2, v163)
                    _ = fi161
                    _ = v162
                    _ = v163
                }
                v143 -= ((0.0 * (v143 - 3.14159)) / (float32(0.1) * v143))
                if len(v1) > 0 {
                    v1[0] = 0
                }
                _ = fi160
            }
            _ = v159
        } else {
            switch "" {
            case "a":
                if false {
                    switch -812 {
                    case -9223372036854775808:
                        for fi164 := 0; fi164 < 16; fi164++ {
                            v165 := GenS11{F0: 0, F1: 42, F2: ("\n\t" + "foo bar")}
                            _ = fi164
                            _ = v165
                        }
                    }
                    var v166 float64 = (3.14159 * (1.5 * ((-2.25 * (1e10 - 3.14159)) - 1.5)))
                    var v167 int = 0
                    _ = v166
                    _ = v167
                }
            case "":
                var v168 int = -1
                _ = v168
            case "foo bar":
                var v169 int64 = v3
                _ = v169
            }
            v170 := GenS4{F0: 32, F1: v142, F2: 4294967295}
            _ = v170
        }
    case "":
        if true {
            fmt.Println(v142)
            v171 := GenS10{F0: "foo bar", F1: v133}
            _ = v171
        } else {
            if true {
                if (474 != v142) {
                    if true {
                        v133 -= v133
                        v133 = v133
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    var v172 uint8 = 60
                    _ = v172
                }
                v173 := GenS8{F0: v3, F1: 0, F2: v133}
                v142 += (v142 ^ v142)
                _ = v173
            }
        }
    case "\n\t":
        v1 = append(v1, 881)
    }
    v142 -= v142
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v133:", v133)
    fmt.Println("v142:", v142)
    fmt.Println("v143:", v143)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v133
    _ = v141
    _ = v142
    _ = v143
    _ = v144
    _ = v2
    _ = v3
}
