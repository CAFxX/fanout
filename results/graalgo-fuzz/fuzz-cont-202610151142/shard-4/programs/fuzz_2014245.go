package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 uint32; F2 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int16; F1 uint16 }
type GenS9 struct { F0 bool; F1 int; F2 uint }
type GenS10 struct { F0 int32; F1 int16; F2 int }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint64; F1 int16 }
type GenS13 struct { F0 uint32; F1 uint32; F2 string }
type GenS14 struct { F0 int16; F1 uint16 }
type GenS15 struct { F0 bool; F1 uint16 }
type GenS16 struct { F0 int16 }
func genFn0() uint32 {
    var v6 int8 = -128
    var v7 uint = 836
    v6 = (v6 | -1)
    v7 = 18446744073709551615
    for fi8 := 0; fi8 < 5; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            v7 = (v7 ^ (v7 / ((v7) | 1)))
            v10 := GenS1{F0: true}
            v11 := []int{}
            _ = fi9
            _ = v10
            _ = v11
        }
        switch (("a" + ("hello" + "\n\t")) + "") {
        case "\n\t":
            for fi12 := 0; fi12 < 8; fi12++ {
                v7 = v7
                var v13 int64 = 9223372036854775807
                _ = fi12
                _ = v13
            }
        case "foo bar":
            for fi14 := 0; fi14 < 6; fi14++ {
                switch "hello" {
                case "hello":
                    v15 := []int{fi14, 1, (-(fi8))}
                    _ = v15
                case "foo bar":
                    v16 := GenS2{F0: 717}
                    _ = v16
                case "":
                    if false {
                        v17 := []int{(fi14 * (fi14 & (fi8 >> ((510) & 63))))}
                        v18 := GenS1{F0: (0 >= 42)}
                        fmt.Println(fi8, v7)
                        _ = v17
                        _ = v18
                    }
                }
                v7 += (v7 | 634)
                switch "" {
                case "\n\t":
                    fmt.Println(v6)
                default:
                    if (1 != 0) {
                        v19 := GenS3{F0: ("foo bar" + ""), F1: 2147483648, F2: 2147483647}
                        _ = v19
                    }
                }
                _ = fi14
            }
        case "a":
            var v20 bool = false
            _ = v20
        default:
            v6 -= (v6 >> ((2) & 63))
        }
        _ = fi8
    }
    _ = v6
    _ = v7
    return uint32(v7)
}

func main() {
    var v1 int16 = 790
    v2 := []uint32{}
    var v3 uint16 = 1
    v4 := make(map[int]string)
    var v5 uint8 = 42
    if (false || (1 == 0)) {
        v3 += v3
        fmt.Println(v5, v5, v1)
        var v21 int32 = 0
        _ = v21
    }
    var v22 int32 = 0
    for fi23 := 0; fi23 < 14; fi23++ {
        v24 := v4[(fi23 * (fi23 << ((65535) & 63)))]
        var v25 int8 = -1
        for fi26 := 0; fi26 < 8; fi26++ {
            v27 := GenS2{F0: fi23}
            _ = fi26
            _ = v27
        }
        _ = fi23
        _ = v24
        _ = v25
    }
    fmt.Println(v5, v22)
    fmt.Println(v3, v3)
    v28 := GenS4{F0: v5}
    switch "foo bar" {
    case "":
        var v29 uint8 = v5
        _ = v29
    case "foo bar":
        for fi30 := 0; fi30 < 8; fi30++ {
            switch (((("a" + "\n\t") + "hello") + "a") + "\n\t") {
            case "\n\t":
                v31 := GenS5{F0: -571, F1: 643}
                _ = v31
            default:
                var v32 int = fi30
                _ = v32
            }
            var v33 uint64 = 42
            if len(v2) > 0 {
                v2[0] = 595
            }
            _ = fi30
            _ = v33
        }
    case "\n\t":
        var v34 float64 = 0.0
        _ = v34
    default:
        fmt.Println(v5, v3, v1)
    }
    fmt.Println(v3)
    if (0 == 128) {
        for fi35 := 0; fi35 < 12; fi35++ {
            for fi36 := 0; fi36 < 12; fi36++ {
                var v37 int64 = -1
                fmt.Println(v5)
                _ = fi36
                _ = v37
            }
            _ = fi35
        }
        if true {
            v1 = v1
            v38 := GenS6{F0: (v5 + v5)}
            _ = v38
        }
    } else {
        var v39 string = "\n\t"
        var v40 string = (v39 + v39)
        delete(v4, 128)
        _ = v39
        _ = v40
    }
    _ = len(v4)
    for fi41 := 0; fi41 < 2; fi41++ {
        if false {
            for fi42 := 0; fi42 < 5; fi42++ {
                v3 += 1
                switch fi41 {
                case 9223372036854775807:
                    v43 := GenS4{F0: v5}
                    _ = v43
                case -1:
                    v44 := GenS7{F0: v3}
                    _ = v44
                case 32768:
                    fmt.Println(fi42, fi41, v22)
                }
                if false {
                    if (((float32(-2.25) / ((1e10 - 0.1) + (1.5 - (0.0 + (0.0 + ((3.14159 / (0.0 + 1e10)) - (0.0 + 3.14159))))))) / 0.1) > (1e10 + 0.1)) {
                        fmt.Println(v1, v22, v1)
                    } else {
                        switch ("\n\t" + "\n\t") {
                        case "foo bar":
                            var v45 int16 = v1
                            _ = v45
                        case "a":
                            fmt.Println(v5)
                        case "hello":
                            var v46 uint64 = 966
                            _ = v46
                        default:
                            if (42 < fi41) {
                                switch "\n\t" {
                                case "foo bar":
                                    if (!(((!(0.0 >= 0.1)) && true) && false)) {
                                        fmt.Println(fi42, v3)
                                    } else {
                                        if true {
                                            if ((v22 >= -814) || ((!false) || (!true))) {
                                                var v47 int16 = -32768
                                                var v48 string = "foo bar"
                                                _ = v47
                                                _ = v48
                                            }
                                        }
                                        var v49 string = ""
                                        _ = v49
                                    }
                                case "hello":
                                    var v50 uint8 = 42
                                    var v51 int8 = -81
                                    if (4294967295 <= 1) {
                                        var v52 int32 = (^(v22))
                                        if (719 >= int64(v3)) {
                                            if ((fi41 >= (fi41 ^ fi41)) || (fi41 <= (fi42 & -847))) {
                                                v2 = append(v2, 2)
                                            }
                                            if ((false || false) && (v3 == v3)) {
                                                fmt.Println(v5, v5, v52)
                                                v53 := v4[-695]
                                                _ = v53
                                            }
                                            v1 = (v1 - v1)
                                        } else {
                                            switch 9223372036854775807 {
                                            case 32767:
                                                switch fi41 {
                                                case -737:
                                                    v1 -= v1
                                                case 127:
                                                    v4[fi42] = ((("a" + "\n\t") + (("\n\t" + ("hello" + "a")) + "")) + "hello")
                                                }
                                                _ = len(v4)
                                                switch "a" {
                                                case "hello":
                                                    if (!false) {
                                                        v1 += (-(v1))
                                                        if true {
                                                            v54 := GenS5{F0: -9223372036854775808, F1: 18446744073709551615}
                                                            _ = v54
                                                        } else {
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                var v55 int32 = -85
                                                                _ = v55
                                                            default:
                                                                if true {
                                                                    if false {
                                                                        v5 -= 30
                                                                    } else {
                                                                        fmt.Println(v52)
                                                                    }
                                                                }
                                                            }
                                                            switch "foo bar" {
                                                            case "a":
                                                                switch 42 {
                                                                case 127:
                                                                    var v56 int32 = v52
                                                                    _ = v56
                                                                case 356:
                                                                    var v57 float64 = 1e10
                                                                    _ = v57
                                                                case 4294967295:
                                                                    v3 += (v3 * v3)
                                                                default:
                                                                    if true {
                                                                        v52 = v22
                                                                        if (v22 >= 856) {
                                                                            switch (fi42 % (((fi41 + (fi42 >> ((18446744073709551615) & 63)))) | 1)) {
                                                                            case 32767:
                                                                                v22 -= -30
                                                                            case 0:
                                                                                switch (fi41 << ((18446744073709551615) & 63)) {
                                                                                case 492:
                                                                                    if (-998 <= 0) {
                                                                                        switch -826 {
                                                                                        case -914:
                                                                                            var v58 int8 = (-(v51))
                                                                                            _ = v58
                                                                                        case -9223372036854775808:
                                                                                            fmt.Println(v51)
                                                                                        default:
                                                                                            v3 = v3
                                                                                        }
                                                                                    } else {
                                                                                        var v59 int16 = int16(v50)
                                                                                        var v60 string = ""
                                                                                        switch ("\n\t" + (v60 + "a")) {
                                                                                        case "hello":
                                                                                            var v61 string = (v60 + ((v60 + v60) + ((("\n\t" + v60) + (v60 + "foo bar")) + "foo bar")))
                                                                                            var v62 uint16 = 255
                                                                                            if false {
                                                                                                fmt.Println(fi42, v3, v61)
                                                                                                v59 += (v59 % ((v1) | 1))
                                                                                                var v63 uint8 = 127
                                                                                                v3 += v62
                                                                                                _ = v63
                                                                                            } else {
                                                                                                v64 := GenS8{F0: v1, F1: (v62 - (v62 + (v3 | v3)))}
                                                                                                v65 := GenS2{F0: (^(fi41))}
                                                                                                if false {
                                                                                                    var v66 uint = 468
                                                                                                    _ = v66
                                                                                                }
                                                                                                _ = v64
                                                                                                _ = v65
                                                                                            }
                                                                                            v67 := GenS1{F0: true}
                                                                                            fmt.Println(v5, v62)
                                                                                            _ = v61
                                                                                            _ = v62
                                                                                            _ = v67
                                                                                        default:
                                                                                            v5 -= v50
                                                                                        }
                                                                                        _ = v59
                                                                                        _ = v60
                                                                                    }
                                                                                }
                                                                            }
                                                                            if (true && (18446744073709551615 >= 4294967295)) {
                                                                                if true {
                                                                                    fmt.Println(fi42, fi42, fi41)
                                                                                    v52 = (v22 << ((42) & 63))
                                                                                } else {
                                                                                    var v68 uint16 = (v3 >> ((1) & 63))
                                                                                    var v69 float64 = (1.5 / -2.25)
                                                                                    _ = v68
                                                                                    _ = v69
                                                                                }
                                                                                v5 = v50
                                                                                delete(v4, fi41)
                                                                            }
                                                                        } else {
                                                                            v1 += v1
                                                                            v50 += (v50 - (v50 + (^(v5))))
                                                                            fmt.Println(fi41, v52, v22)
                                                                        }
                                                                    }
                                                                }
                                                            case "hello":
                                                                switch 42 {
                                                                case 4294967295:
                                                                    v52 -= 2147483647
                                                                case 127:
                                                                    switch fi42 {
                                                                    case -1:
                                                                        fmt.Println(v5, v5, v22)
                                                                    case 4294967295:
                                                                        v5 = 2
                                                                    case -126:
                                                                        var v70 bool = (!false)
                                                                        _ = v70
                                                                    default:
                                                                        if true {
                                                                            switch ("\n\t" + (("hello" + ("hello" + "")) + "a")) {
                                                                            case "a":
                                                                                var v71 uint = 522
                                                                                _ = v71
                                                                            }
                                                                            v72 := GenS9{F0: false, F1: (fi42 % ((fi41) | 1)), F2: 18446744073709551615}
                                                                            _ = v72
                                                                        }
                                                                    }
                                                                default:
                                                                    var v73 uint32 = uint32(v3)
                                                                    _ = v73
                                                                }
                                                            case "foo bar":
                                                                v1 -= v1
                                                            default:
                                                                _ = len(v4)
                                                            }
                                                            if (1 < (v3 * v3)) {
                                                                if false {
                                                                    v74 := GenS6{F0: v5}
                                                                    var v75 uint = 623
                                                                    _ = v74
                                                                    _ = v75
                                                                } else {
                                                                    v76 := v4[fi42]
                                                                    _ = v76
                                                                }
                                                                v51 += v51
                                                            }
                                                        }
                                                    }
                                                default:
                                                    if len(v2) > 0 {
                                                        v2[0] = 296
                                                    }
                                                }
                                            }
                                            switch 0 {
                                            case -211:
                                                fmt.Println(fi41, v52)
                                            case 4294967295:
                                                v77 := GenS6{F0: v5}
                                                _ = v77
                                            default:
                                                switch fi41 {
                                                case 128:
                                                    var v78 int = (fi41 % ((-1) | 1))
                                                    _ = v78
                                                }
                                            }
                                            var v79 float32 = -2.25
                                            _ = v79
                                        }
                                        _ = v52
                                    }
                                    _ = v50
                                    _ = v51
                                default:
                                    v22 = -1
                                }
                                var v80 uint16 = uint16(v5)
                                _ = v80
                            }
                        }
                        v3 += 65535
                        if false {
                            switch (-(fi41)) {
                            case -1:
                                var v81 uint = 255
                                _ = v81
                            }
                        }
                    }
                    var v82 string = ""
                    var v83 float32 = 0.0
                    _ = v82
                    _ = v83
                }
                _ = fi42
            }
            if false {
                v3 = 2
                var v84 float64 = ((0.1 - 1.5) / 1.0)
                _ = v84
            }
            v85 := GenS9{F0: (!false), F1: fi41, F2: 4294967295}
            _ = v85
        } else {
            for fi86 := 0; fi86 < 13; fi86++ {
                if true {
                    if true {
                        v87 := v4[(fi86 << ((255) & 63))]
                        v3 += v3
                        var v88 float64 = 0.1
                        _ = v87
                        _ = v88
                    } else {
                        v89 := GenS1{F0: true}
                        _ = v89
                    }
                    v1 += v1
                    if true {
                        v3 += (v3 | 2)
                        if (!(!false)) {
                            delete(v4, (-(fi41)))
                        } else {
                            fmt.Println(v22, v22, v3)
                        }
                        v3 = (v3 / (((v3 * 0)) | 1))
                        var v90 float64 = float64(0.0)
                        v22 += (v22 % ((v22) | 1))
                        _ = v90
                    } else {
                        v91 := GenS10{F0: (v22 >> ((4294967295) & 63)), F1: v1, F2: (fi86 ^ (fi86 ^ (fi86 % ((fi41) | 1))))}
                        v1 = (v1 + (v1 >> ((810) & 63)))
                        _ = v91
                    }
                    v92 := GenS8{F0: (^(v1)), F1: v3}
                    if (int32(fi86) == (v22 + 2147483647)) {
                        if (0.1 <= 1.5) {
                            if true {
                                switch fi41 {
                                case 0:
                                    if true {
                                        var v93 uint32 = 581
                                        var v94 float64 = float64(0.1)
                                        if true {
                                            v95 := GenS2{F0: -957}
                                            _ = v95
                                        } else {
                                            fmt.Println(v1, v1, v94)
                                        }
                                        v5 -= 42
                                        switch (fi41 % ((32768) | 1)) {
                                        case 42:
                                            v1 = (-(v1))
                                        default:
                                            delete(v4, fi41)
                                        }
                                        _ = v93
                                        _ = v94
                                    }
                                default:
                                    v3 = (v3 | v3)
                                }
                            }
                            switch (fi41 / ((fi86) | 1)) {
                            case -251:
                                fmt.Println(fi86)
                                v1 += (^(v1))
                            case 127:
                                delete(v4, (fi41 - fi41))
                            case 32767:
                                switch (("" + "") + (("hello" + "foo bar") + ("foo bar" + "foo bar"))) {
                                case "":
                                    v5 += 42
                                case "a":
                                    fmt.Println(v5)
                                case "hello":
                                    v1 = 1
                                }
                            default:
                                v5 = (v5 & v5)
                            }
                            v5 = 1
                            if ((-737 != -1) || (false || (!(v3 <= (v3 << ((18446744073709551615) & 63)))))) {
                                var v96 uint32 = 2147483648
                                fmt.Println(v5)
                                v4[2147483647] = "a"
                                var v97 int16 = int16(v22)
                                var v98 int16 = int16(v22)
                                var v99 int64 = int64(v1)
                                _ = v96
                                _ = v97
                                _ = v98
                                _ = v99
                            }
                            fmt.Println(v3, v1)
                            v2 = append(v2, 556)
                        } else {
                            if (!(float64(1e10) <= -2.25)) {
                                v22 = v22
                                if false {
                                    if false {
                                        var v100 uint16 = (v3 << ((1) & 63))
                                        v101 := GenS11{F0: (962 >= 939)}
                                        if (255 >= (v5 ^ v5)) {
                                            switch "hello" {
                                            case "a":
                                                fmt.Println(v3, fi41, v5)
                                            case "foo bar":
                                                v3 += (v100 + (v3 / ((322) | 1)))
                                            }
                                        }
                                        _ = v100
                                        _ = v101
                                    } else {
                                        fmt.Println(v3, v22)
                                    }
                                    _ = len(v2)
                                } else {
                                    v102 := GenS5{F0: -9223372036854775808, F1: 1}
                                    _ = v102
                                }
                                var v103 uint64 = 18446744073709551615
                                _ = v103
                            }
                        }
                        var v104 uint32 = genFn0()
                        var v105 int64 = 214
                        _ = v104
                        _ = v105
                    } else {
                        var v106 float32 = 0.1
                        _ = v106
                    }
                    _ = v92
                } else {
                    v22 = (v22 - (-(v22)))
                    fmt.Println(v5, fi41)
                    var v107 uint32 = 65535
                    _ = v107
                }
                _ = fi86
            }
        }
        fmt.Println(v1)
        for fi108 := 0; fi108 < 20; fi108++ {
            if (1.5 > 0.0) {
                v5 += 128
                var v109 int = fi41
                _ = v109
            }
            var v110 uint64 = 662
            _ = fi108
            _ = v110
        }
        _ = fi41
    }
    for fi111 := 0; fi111 < 11; fi111++ {
        if len(v2) > 0 {
            v2[0] = uint32(v5)
        }
        _ = fi111
    }
    for fi112 := 0; fi112 < 19; fi112++ {
        fmt.Println(v1, v22, v1)
        var v113 int = (fi112 | (-(fi112)))
        if ("foo bar" >= "\n\t") {
            switch -39 {
            case -1:
                v114 := GenS2{F0: (v113 / (((fi112 * 32768)) | 1))}
                _ = v114
            }
        } else {
            if ((v3 << ((1) & 63)) <= v3) {
                fmt.Println(v22)
                v115 := v4[(v113 / (((-(v113))) | 1))]
                if (!true) {
                    if (false && true) {
                        var v116 float64 = 1e10
                        var v117 int16 = (v1 >> ((42) & 63))
                        _ = v116
                        _ = v117
                    } else {
                        v118 := v4[0]
                        if len(v2) > 0 {
                            v2[0] = genFn0()
                        }
                        v22 += v22
                        _ = v118
                    }
                    for fi119 := 0; fi119 < 18; fi119++ {
                        if false {
                            v113 -= fi112
                            v120 := GenS12{F0: 850, F1: 1000}
                            switch 42 {
                            case -9223372036854775808:
                                fmt.Println(v3, v3, v22)
                            case 2147483647:
                                switch (-(fi112)) {
                                case -2147483648:
                                    switch ((v115 + v115) + v115) {
                                    case "a":
                                        switch "" {
                                        case "a":
                                            switch "a" {
                                            case "foo bar":
                                                v5 = v5
                                            case "\n\t":
                                                var v121 uint8 = uint8(v3)
                                                var v122 uint32 = 242
                                                _ = v121
                                                _ = v122
                                            }
                                        case "hello":
                                            if (900 >= 42) {
                                                v113 -= v113
                                                var v123 uint16 = 1
                                                v124 := GenS13{F0: 0, F1: 2, F2: "\n\t"}
                                                v115 = v115
                                                _ = v123
                                                _ = v124
                                            }
                                        case "":
                                            v113 = fi112
                                        }
                                    case "\n\t":
                                        var v125 uint32 = 799
                                        _ = v125
                                    case "":
                                        fmt.Println(v3, v113)
                                    default:
                                        v5 = 0
                                    }
                                }
                            case 32767:
                                v5 += v5
                            default:
                                v126 := GenS14{F0: (v1 + (v1 | (v1 << ((uint(v3)) & 63)))), F1: v3}
                                _ = v126
                            }
                            _ = v120
                        }
                        if (!(v22 != -1)) {
                            v1 -= -956
                            var v127 int16 = int16(v3)
                            if false {
                                switch 32767 {
                                case 32768:
                                    v128 := GenS8{F0: 1, F1: (v3 + 65535)}
                                    _ = v128
                                case 4294967295:
                                    v22 = (^(v22))
                                default:
                                    v5 = v5
                                }
                            } else {
                                fmt.Println(fi119)
                                switch fi119 {
                                case -1:
                                    if true {
                                        switch fi119 {
                                        case -2147483648:
                                            var v129 int8 = 127
                                            _ = v129
                                        default:
                                            v22 += v22
                                        }
                                        var v130 uint8 = 255
                                        _ = v130
                                    }
                                case 800:
                                    var v131 bool = false
                                    _ = v131
                                case 32767:
                                    v22 = (v22 | v22)
                                default:
                                    fmt.Println(fi119, v1, v1)
                                }
                            }
                            if len(v2) > 0 {
                                v2[0] = 4294967295
                            }
                            _ = v127
                        }
                        v132 := GenS15{F0: true, F1: (v3 << ((65535) & 63))}
                        _ = fi119
                        _ = v132
                    }
                }
                _ = v115
            } else {
                v22 += v22
                for fi133 := 0; fi133 < 1; fi133++ {
                    v134 := GenS13{F0: 0, F1: 413, F2: "\n\t"}
                    var v135 int = (v113 * fi112)
                    _ = fi133
                    _ = v134
                    _ = v135
                }
                for fi136 := 0; fi136 < 16; fi136++ {
                    v137 := GenS16{F0: 1000}
                    fmt.Println(fi112, v113, v1)
                    v138 := GenS11{F0: (!(42 > 0))}
                    _ = fi136
                    _ = v137
                    _ = v138
                }
            }
        }
        _ = fi112
        _ = v113
    }
    v22 = (-(v22))
    for fi139 := 0; fi139 < 16; fi139++ {
        _ = len(v4)
        _ = fi139
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
