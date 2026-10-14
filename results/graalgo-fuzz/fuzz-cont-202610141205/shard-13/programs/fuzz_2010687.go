package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 int16; F1 uint32; F2 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int; F1 string }
type GenS7 struct { F0 string; F1 string }
type GenS8 struct { F0 uint64; F1 int8; F2 int16 }
type GenS9 struct { F0 int32; F1 int; F2 uint16 }
type GenS10 struct { F0 bool; F1 int16; F2 int64 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int8; F1 int8; F2 int64 }
type GenS13 struct { F0 uint64; F1 uint8 }
type GenS14 struct { F0 uint; F1 string; F2 uint }
type GenS15 struct { F0 int64; F1 bool; F2 int }
type GenS16 struct { F0 int8; F1 uint16; F2 uint16 }
func genFn0(p0 int, p1 int64) int16 {
    if true {
        v7 := []int{(p0 >> ((1) & 63)), 390, (p0 & 1), (p0 - -2147483648)}
        if len(v7) > 0 {
            v7[0] = p0
        }
        _ = v7
    } else {
        p0 += p0
        v8 := []int{(p0 << ((65535) & 63)), 0, 645, (-(p0))}
        fmt.Println(p0)
        _ = v8
    }
    p1 = (p1 & p1)
    p1 -= (^(p1))
    if true {
        var v9 uint = 2
        v10 := GenS1{F0: 255}
        _ = v10
        _ = v9
    } else {
        v11 := GenS1{F0: 0}
        v12 := []int{p0, (p0 >> ((18446744073709551615) & 63)), p0, p0}
        _ = v11
        _ = v12
    }
    switch (p0 ^ (p0 & -1)) {
    case 42:
        fmt.Println(p1, p0)
    }
    _ = p0
    _ = p1
    return -32768
}
func genFn1() uint8 {
    if true {
        var v13 bool = (42 == 32767)
        v14 := GenS2{F0: 323, F1: 149}
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        fmt.Println(fi15)
        var v16 bool = true
        var v17 float32 = 1.5
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 uint16 = 255
    var v19 uint = 392
    _ = v18
    _ = v19
    return 128
}

func main() {
    var v1 uint16 = 42
    var v2 int8 = -89
    var v3 uint16 = (v1 << ((65535) & 63))
    var v4 int32 = 2147483647
    var v5 int8 = 127
    v6 := []uint32{}
    fmt.Println(v1, v4, v5)
    for fi20 := 0; fi20 < 8; fi20++ {
        if true {
            if len(v6) > 0 {
                v6[0] = 4294967295
            }
        } else {
            v3 = v3
            switch fi20 {
            case 127:
                for fi21 := 0; fi21 < 8; fi21++ {
                    switch int(v2) {
                    case -2147483648:
                        v1 = 2
                    default:
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                    if (true || (!false)) {
                        if false {
                            switch "hello" {
                            case "foo bar":
                                var v22 bool = false
                                _ = v22
                            case "":
                                v23 := GenS1{F0: v3}
                                switch "hello" {
                                case "\n\t":
                                    v6 = append(v6, 0)
                                default:
                                    var v24 uint32 = 0
                                    _ = v24
                                }
                                _ = v23
                            case "a":
                                var v25 uint64 = 730
                                _ = v25
                            default:
                                v1 = v1
                            }
                        } else {
                            v26 := GenS3{F0: true, F1: 42}
                            if len(v6) > 0 {
                                v6[0] = 4294967295
                            }
                            var v27 float32 = -2.25
                            _ = v26
                            _ = v27
                        }
                        v28 := GenS3{F0: true, F1: -1}
                        v5 = (v5 * 87)
                        _ = v28
                    }
                    _ = fi21
                }
            }
        }
        if false {
            v2 += 4
            _ = len(v6)
            var v29 float32 = float32(1e10)
            _ = v29
        } else {
            for fi30 := 0; fi30 < 10; fi30++ {
                if ((32767 < 42) && false) {
                    v31 := GenS2{F0: 409, F1: 42}
                    _ = v31
                } else {
                    if (193 > genFn1()) {
                        if ((v3 + (v1 % ((v3) | 1))) != v1) {
                            if (752 < 2147483647) {
                                fmt.Println(fi30)
                                v3 += v1
                            }
                            switch (fi30 / ((9223372036854775807) | 1)) {
                            case 796:
                                var v32 uint8 = 1
                                _ = v32
                            case 32767:
                                switch (^(fi30)) {
                                case 42:
                                    if (((-2.25 * (-2.25 * -2.25)) * 0.1) <= (1.5 - -2.25)) {
                                        v4 = 42
                                        if (false || (!true)) {
                                            v2 = 42
                                            v5 = 106
                                            var v33 int16 = -673
                                            _ = v33
                                        } else {
                                            if (-128 == (v2 << ((0) & 63))) {
                                                fmt.Println(v1, fi20)
                                                v34 := GenS4{F0: -352, F1: 1, F2: (("\n\t" + ("\n\t" + (("foo bar" + ("" + ("foo bar" + (("a" + (("hello" + "\n\t") + ("hello" + "hello"))) + (("hello" + "") + "hello"))))) + "foo bar"))) + "a")}
                                                _ = v34
                                            }
                                            var v35 int8 = (v2 >> ((255) & 63))
                                            var v36 uint8 = 2
                                            _ = v35
                                            _ = v36
                                        }
                                        var v37 uint = 255
                                        _ = v37
                                    }
                                default:
                                    v4 = (v4 / (((v4 - 1)) | 1))
                                }
                            }
                        }
                        v2 += v5
                    }
                    if (false && true) {
                        v5 = v2
                        v6 = append(v6, 0)
                        v38 := GenS5{F0: false}
                        var v39 uint16 = v3
                        fmt.Println(v1)
                        if true {
                            if false {
                                v6 = append(v6, 42)
                                if len(v6) > 0 {
                                    v6[0] = 2
                                }
                                v5 = (-(v5))
                                v40 := GenS1{F0: v1}
                                _ = len(v6)
                                _ = v40
                            } else {
                                var v41 uint32 = 2
                                switch "foo bar" {
                                case "\n\t":
                                    if (!false) {
                                        _ = len(v6)
                                        var v42 int = fi30
                                        fmt.Println(v42, v2)
                                        _ = v42
                                    }
                                case "a":
                                    var v43 string = "foo bar"
                                    _ = v43
                                }
                                if true {
                                    if true {
                                        switch "foo bar" {
                                        case "hello":
                                            if len(v6) > 0 {
                                                v6[0] = v41
                                            }
                                        }
                                        _ = len(v6)
                                    } else {
                                        switch fi20 {
                                        case 4294967295:
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        default:
                                            v5 += -128
                                        }
                                        v44 := GenS3{F0: true, F1: 701}
                                        var v45 int = (fi20 % ((fi20) | 1))
                                        _ = v44
                                        _ = v45
                                    }
                                    v39 = v1
                                    switch "foo bar" {
                                    case "hello":
                                        var v46 int32 = (v4 << ((18446744073709551615) & 63))
                                        _ = v46
                                    case "\n\t":
                                        var v47 int8 = (v2 / (((v2 << ((uint(v3)) & 63))) | 1))
                                        _ = v47
                                    case "":
                                        _ = len(v6)
                                    }
                                    if false {
                                        if false {
                                            fmt.Println(fi20, v39)
                                            if true {
                                                switch "a" {
                                                case "hello":
                                                    if false {
                                                        fmt.Println(v2, v4)
                                                        var v48 uint = 2
                                                        _ = v48
                                                    }
                                                }
                                            } else {
                                                switch "" {
                                                case "hello":
                                                    v2 += (-(v2))
                                                default:
                                                    if (-2.25 > 0.0) {
                                                        if true {
                                                            v49 := GenS2{F0: 1, F1: (^(v41))}
                                                            _ = v49
                                                        }
                                                        v50 := GenS5{F0: false}
                                                        v5 = 0
                                                        _ = v50
                                                    }
                                                }
                                                switch "a" {
                                                case "foo bar":
                                                    switch 918 {
                                                    case 445:
                                                        fmt.Println(v41, v2)
                                                        fmt.Println(v3, v2, v39)
                                                        var v51 int16 = 0
                                                        _ = v51
                                                    case 32767:
                                                        v6 = append(v6, v41)
                                                    }
                                                case "":
                                                    v41 = v41
                                                }
                                                v52 := GenS6{F0: 32768, F1: ("" + ("hello" + "\n\t"))}
                                                _ = v52
                                            }
                                            v4 += (^(v4))
                                        }
                                        var v53 int32 = -1
                                        if false {
                                            v1 += v39
                                            var v54 int64 = 374
                                            _ = v54
                                        } else {
                                            v3 -= (v3 | (v1 * v1))
                                            switch "foo bar" {
                                            case "\n\t":
                                                var v55 int8 = -128
                                                _ = v55
                                            case "hello":
                                                fmt.Println(fi30, v1)
                                            case "foo bar":
                                                v56 := GenS3{F0: false, F1: 0}
                                                _ = v56
                                            default:
                                                if ("" <= "hello") {
                                                    v57 := GenS5{F0: (!(!false))}
                                                    _ = v57
                                                }
                                            }
                                            v58 := GenS7{F0: "a", F1: ((("a" + "") + "a") + "a")}
                                            _ = v58
                                        }
                                        _ = v53
                                    }
                                    fmt.Println(fi20)
                                }
                                _ = v41
                            }
                            var v59 string = ""
                            if true {
                                v1 -= (v39 % ((v39) | 1))
                                v60 := GenS1{F0: (v39 * 477)}
                                v61 := GenS6{F0: (fi20 & 0), F1: (v59 + ((v59 + v59) + v59))}
                                _ = v60
                                _ = v61
                            } else {
                                v6 = append(v6, 2)
                            }
                            _ = v59
                        }
                        _ = v38
                        _ = v39
                    }
                }
                v1 -= (v3 | (v1 ^ (^(v1))))
                _ = fi30
            }
            switch "foo bar" {
            case "hello":
                switch ("foo bar" + "\n\t") {
                case "":
                    v62 := GenS7{F0: "", F1: ""}
                    _ = v62
                case "foo bar":
                    switch ("" + ("\n\t" + ("\n\t" + ""))) {
                    case "\n\t":
                        var v63 float64 = float64(3.14159)
                        _ = v63
                    }
                case "hello":
                    fmt.Println(v1, v3)
                default:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
            case "foo bar":
                for fi64 := 0; fi64 < 6; fi64++ {
                    v65 := GenS2{F0: 141, F1: 0}
                    v2 = v5
                    v6 = append(v6, 0)
                    _ = fi64
                    _ = v65
                }
            }
            var v66 int32 = (^(v4))
            _ = v66
        }
        v5 = (v2 * (v5 * v5))
        _ = fi20
    }
    for fi67 := 0; fi67 < 6; fi67++ {
        v4 += v4
        for fi68 := 0; fi68 < 6; fi68++ {
            switch (fi68 & (fi68 * fi67)) {
            case 32768:
                if len(v6) > 0 {
                    v6[0] = 0
                }
            case 135:
                v1 -= 65535
            case -2147483648:
                fmt.Println(v1, v1)
            default:
                fmt.Println(v3, v2, v2)
            }
            v2 += v5
            _ = fi68
        }
        var v69 uint32 = uint32(v1)
        _ = fi67
        _ = v69
    }
    var v70 uint16 = v3
    var v71 int8 = v2
    v1 = v3
    if true {
        v72 := GenS8{F0: 1, F1: 127, F2: 42}
        for fi73 := 0; fi73 < 11; fi73++ {
            switch ("foo bar" + ("" + ("foo bar" + ("a" + ("\n\t" + "foo bar"))))) {
            case "\n\t":
                v74 := GenS9{F0: (v4 - (v4 & 249)), F1: 32768, F2: 65535}
                _ = v74
            }
            _ = fi73
        }
        v4 = v4
        _ = v72
    } else {
        v2 -= 127
        for fi75 := 0; fi75 < 10; fi75++ {
            switch "foo bar" {
            case "hello":
                v76 := GenS9{F0: -1, F1: fi75, F2: 32767}
                _ = v76
            case "\n\t":
                if false {
                    var v77 int = fi75
                    _ = v77
                } else {
                    if false {
                        v71 = (v2 & (-(v5)))
                        for fi78 := 0; fi78 < 17; fi78++ {
                            switch (-(fi78)) {
                            case 127:
                                switch (fi78 + (fi75 * (fi78 >> ((65535) & 63)))) {
                                case -1:
                                    switch "" {
                                    case "a":
                                        v79 := GenS10{F0: (!(127 < 42)), F1: -560, F2: 42}
                                        _ = v79
                                    case "":
                                        var v80 float64 = 0.0
                                        _ = v80
                                    default:
                                        v81 := GenS8{F0: 539, F1: 0, F2: 32767}
                                        _ = v81
                                    }
                                    fmt.Println(fi78, v71)
                                }
                                if ((int32(v5) > (v4 ^ (v4 % ((v4) | 1)))) && (false && true)) {
                                    var v82 int8 = v71
                                    var v83 int64 = -930
                                    _ = len(v6)
                                    _ = v82
                                    _ = v83
                                }
                            case -870:
                                _ = len(v6)
                            default:
                                if (false || ((128 <= 127) && (v2 > v5))) {
                                    switch "a" {
                                    case "foo bar":
                                        if (1 < 824) {
                                            var v84 uint64 = 42
                                            switch (("" + "hello") + ("hello" + "hello")) {
                                            case "\n\t":
                                                if true {
                                                    v85 := GenS3{F0: (!(true || false)), F1: 16}
                                                    v86 := GenS11{F0: 0}
                                                    v6 = append(v6, 1)
                                                    v4 = (-(v4))
                                                    fmt.Println(v3, v5, v5)
                                                    fmt.Println(v5, fi75, v2)
                                                    v6 = append(v6, 42)
                                                    _ = v85
                                                    _ = v86
                                                } else {
                                                    if ((0 == 128) && (178 == 179)) {
                                                        var v87 string = "\n\t"
                                                        _ = v87
                                                    } else {
                                                        if (-643 < 1) {
                                                            v84 = 42
                                                            v88 := GenS12{F0: v5, F1: -128, F2: 0}
                                                            var v89 uint8 = 136
                                                            _ = v88
                                                            _ = v89
                                                        }
                                                        v3 -= (v3 - (v1 >> ((42) & 63)))
                                                    }
                                                    if (true || (1 >= v84)) {
                                                        var v90 string = ("\n\t" + "hello")
                                                        if len(v6) > 0 {
                                                            v6[0] = 2147483648
                                                        }
                                                        _ = v90
                                                    }
                                                }
                                            default:
                                                v91 := GenS13{F0: v84, F1: genFn1()}
                                                _ = v91
                                            }
                                            switch fi75 {
                                            case 423:
                                                var v92 string = "foo bar"
                                                _ = v92
                                            case 1:
                                                var v93 uint64 = (v84 | v84)
                                                _ = v93
                                            case 32768:
                                                if (v4 >= (-(v4))) {
                                                    switch fi78 {
                                                    case 127:
                                                        var v94 uint16 = v3
                                                        _ = v94
                                                    default:
                                                        if (false || (!true)) {
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                var v95 uint32 = 997
                                                                _ = v95
                                                            default:
                                                                v96 := GenS14{F0: 80, F1: "", F2: uint(v84)}
                                                                _ = v96
                                                            }
                                                            var v97 uint16 = 91
                                                            _ = v97
                                                        }
                                                    }
                                                    if (0.0 < (0.1 * 1.5)) {
                                                        var v98 float32 = float32((1e10 - (float64(0.1) * -2.25)))
                                                        switch "\n\t" {
                                                        case "hello":
                                                            v4 -= (v4 & (v4 << ((65535) & 63)))
                                                        case "a":
                                                            v84 += 18446744073709551615
                                                        case "":
                                                            v99 := GenS2{F0: 2147483648, F1: 4294967295}
                                                            _ = v99
                                                        default:
                                                            v70 += 2
                                                        }
                                                        var v100 uint8 = 223
                                                        _ = v100
                                                        _ = v98
                                                    } else {
                                                        var v101 uint32 = 1
                                                        _ = v101
                                                    }
                                                    if false {
                                                        var v102 int16 = -337
                                                        v70 = v3
                                                        if true {
                                                            v71 = (v2 << ((726) & 63))
                                                        } else {
                                                            v6 = append(v6, 432)
                                                        }
                                                        v4 = v4
                                                        _ = v102
                                                    } else {
                                                        v6 = append(v6, 0)
                                                        v103 := GenS13{F0: v84, F1: genFn1()}
                                                        if true {
                                                            if true {
                                                                if true {
                                                                    var v104 int8 = 107
                                                                    var v105 int16 = 476
                                                                    _ = v104
                                                                    _ = v105
                                                                }
                                                            } else {
                                                                v4 = v4
                                                            }
                                                        } else {
                                                            fmt.Println(v70, v4)
                                                        }
                                                        _ = v103
                                                    }
                                                }
                                            default:
                                                v106 := GenS13{F0: v84, F1: 127}
                                                _ = v106
                                            }
                                            _ = v84
                                        } else {
                                            fmt.Println(v5, v71)
                                            switch "" {
                                            case "a":
                                                v3 -= v70
                                            default:
                                                var v107 bool = (246 >= 0)
                                                v108 := GenS13{F0: 18446744073709551615, F1: 99}
                                                fmt.Println(fi75)
                                                _ = v107
                                                _ = v108
                                            }
                                            v109 := GenS15{F0: -1, F1: (4294967295 <= (fi75 + 0)), F2: (fi75 / ((fi75) | 1))}
                                            _ = v109
                                        }
                                    case "a":
                                        var v110 int8 = -52
                                        _ = v110
                                    case "hello":
                                        var v111 int8 = v2
                                        _ = v111
                                    }
                                } else {
                                    v6 = append(v6, 0)
                                }
                            }
                            _ = len(v6)
                            var v112 int8 = (^(v5))
                            _ = fi78
                            _ = v112
                        }
                        var v113 int32 = (v4 << ((255) & 63))
                        _ = v113
                    }
                }
            default:
                v6 = append(v6, 4294967295)
            }
            _ = fi75
        }
    }
    for fi114 := 0; fi114 < 10; fi114++ {
        for fi115 := 0; fi115 < 2; fi115++ {
            if false {
                fmt.Println(v70, v71)
                v6 = append(v6, 481)
                if (!true) {
                    v2 += (v2 << ((2) & 63))
                    var v116 int64 = -420
                    var v117 uint8 = 1
                    _ = v116
                    _ = v117
                } else {
                    if (false || true) {
                        v2 -= v5
                        var v118 int8 = v71
                        v1 = (v70 / ((255) | 1))
                        v4 += v4
                        if ((!false) || (476 > 4294967295)) {
                            v4 = (v4 ^ v4)
                            if false {
                                if ((!false) || false) {
                                    v119 := GenS6{F0: 127, F1: "hello"}
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    var v120 int32 = v4
                                    _ = v119
                                    _ = v120
                                }
                            } else {
                                var v121 string = "\n\t"
                                v3 += (^(v3))
                                var v122 uint16 = v3
                                _ = v121
                                _ = v122
                            }
                        } else {
                            fmt.Println(v118, v5, v5)
                            v123 := GenS6{F0: (fi115 * 703), F1: ("\n\t" + "")}
                            fmt.Println(fi115, fi114)
                            if len(v6) > 0 {
                                v6[0] = 2
                            }
                            var v124 uint = 255
                            _ = v123
                            _ = v124
                        }
                        var v125 uint8 = 67
                        var v126 string = ""
                        var v127 uint16 = (^(v3))
                        v6 = append(v6, 0)
                        v128 := GenS16{F0: int8(v4), F1: v70, F2: 0}
                        _ = v118
                        _ = v125
                        _ = v126
                        _ = v127
                        _ = v128
                    } else {
                        if (int64(v70) < 0) {
                            v5 -= (v2 ^ 0)
                            var v129 int16 = int16(fi114)
                            _ = len(v6)
                            _ = v129
                        }
                        var v130 int16 = -32768
                        _ = v130
                    }
                    var v131 int16 = -110
                    _ = v131
                }
            }
            _ = len(v6)
            _ = fi115
        }
        if false {
            for fi132 := 0; fi132 < 4; fi132++ {
                var v133 uint64 = 0
                _ = fi132
                _ = v133
            }
            var v134 float32 = 0.0
            v70 += v3
            _ = v134
        } else {
            v71 = 121
        }
        if true {
            v4 = (v4 % ((v4) | 1))
        }
        _ = fi114
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if true {
        v5 -= (v71 - v5)
        v6 = append(v6, 4294967295)
    }
    var v135 float64 = float64((1.5 - (0.1 - -2.25)))
    if false {
        var v136 uint = 18446744073709551615
        v137 := GenS9{F0: v4, F1: 128, F2: (v70 ^ v3)}
        if false {
            if (v1 != 255) {
                v70 += v1
            }
            var v138 uint64 = 42
            for fi139 := 0; fi139 < 11; fi139++ {
                v140 := GenS11{F0: 1}
                _ = fi139
                _ = v140
            }
            _ = v138
        }
        _ = v136
        _ = v137
    } else {
        var v141 string = ""
        var v142 float64 = (-2.25 - v135)
        _ = v141
        _ = v142
    }
    var v143 float32 = 1e10
    switch -392 {
    case 2147483647:
        if (v143 <= (v143 + float32(v135))) {
            var v144 int32 = (v4 % ((v4) | 1))
            var v145 uint = 18446744073709551615
            _ = v144
            _ = v145
        } else {
            fmt.Println(v4)
        }
    case 32767:
        var v146 uint8 = 46
        _ = v146
    case 4294967295:
        var v147 int16 = -1
        _ = v147
    }
    var v148 uint16 = (v70 & (v3 | (v3 | (v3 << ((1) & 63)))))
    _ = len(v6)
    fmt.Println(v3, v4)
    v4 -= (v4 % ((v4) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v135:", v135)
    fmt.Println("v143:", v143)
    fmt.Println("v148:", v148)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v135
    _ = v143
    _ = v148
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v70
    _ = v71
}
