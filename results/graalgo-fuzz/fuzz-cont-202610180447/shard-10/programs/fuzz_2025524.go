package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint32; F1 int32 }
type GenS7 struct { F0 uint16; F1 uint64 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint8; F1 int8 }
type GenS10 struct { F0 string; F1 int32 }
type GenS11 struct { F0 uint; F1 int32 }
type GenS12 struct { F0 int32; F1 string; F2 uint8 }
type GenS13 struct { F0 int16; F1 int64 }
type GenS14 struct { F0 int16; F1 int8; F2 int32 }
type GenS15 struct { F0 string; F1 int32 }
func genFn0(p0 int64, p1 bool, p2 uint8) uint16 {
    var v4 bool = (!true)
    var v5 uint16 = 0
    v6 := GenS1{F0: v4, F1: (^(p0))}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return v5
}
func genFn1(p0 uint, p1 int) int8 {
    p0 = (p0 << ((652) & 63))
    if false {
        var v7 uint32 = 42
        for fi8 := 0; fi8 < 20; fi8++ {
            switch "a" {
            case "a":
                var v9 bool = false
                _ = v9
            case "\n\t":
                p0 = (p0 * p0)
            case "foo bar":
                if ((0.0 * (1e10 - 0.0)) == 0.0) {
                    for fi10 := 0; fi10 < 14; fi10++ {
                        if false {
                            v7 -= v7
                            var v11 uint64 = 2
                            var v12 string = ("" + "\n\t")
                            _ = v11
                            _ = v12
                        } else {
                            v13 := GenS1{F0: true, F1: -918}
                            var v14 float64 = ((0.0 + (3.14159 - (3.14159 - ((((3.14159 + float64((1e10 - float32(float64(3.14159))))) + ((((float64(1.5) / 1.0) / -2.25) - ((0.0 / -2.25) * 0.1)) - (((0.1 - -2.25) + ((((float64(3.14159) + -2.25) + 0.0) - float64(((1.5 - -2.25) / float32(float64((float32(3.14159) + 3.14159)))))) + ((float64(1e10) - (3.14159 * 0.0)) + 1e10))) * ((float64((float32((float64(0.0) + -2.25)) - 1e10)) + (0.0 * (3.14159 + float64(-2.25)))) * float64(-2.25))))) + 1.5) - ((3.14159 * -2.25) * (float64(1.5) / 1.5)))))) * 1.5)
                            if true {
                                var v15 int32 = 0
                                v16 := []int{fi8}
                                p0 += (p0 + 2)
                                _ = v15
                                _ = v16
                            }
                            v14 = (3.14159 * v14)
                            v17 := GenS1{F0: (!false), F1: 143}
                            _ = v13
                            _ = v14
                            _ = v17
                        }
                        _ = fi10
                    }
                    switch (p1 - (fi8 + p1)) {
                    case 0:
                        v18 := []int{633, p1, p1, 2147483647}
                        _ = v18
                    default:
                        for fi19 := 0; fi19 < 17; fi19++ {
                            v20 := []int{fi8, 9223372036854775807, (p1 << ((p0) & 63)), 409}
                            v21 := GenS1{F0: (-472 > 42), F1: -1}
                            var v22 int32 = -781
                            fmt.Println(v22)
                            v22 += 2147483647
                            _ = fi19
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                    }
                } else {
                    v7 += v7
                    for fi23 := 0; fi23 < 15; fi23++ {
                        v24 := []int{128, (fi8 ^ (fi23 % (((-(fi23))) | 1)))}
                        if len(v24) > 0 {
                            v24[0] = fi8
                        }
                        switch (("foo bar" + (("" + ("foo bar" + (("foo bar" + (("a" + (("a" + "hello") + "hello")) + "foo bar")) + "\n\t"))) + "hello")) + "\n\t") {
                        case "hello":
                            var v25 int8 = int8(v7)
                            _ = v25
                        case "foo bar":
                            fmt.Println(p0, p0)
                        }
                        _ = fi23
                        _ = v24
                    }
                }
            }
            switch fi8 {
            case 2147483647:
                for fi26 := 0; fi26 < 2; fi26++ {
                    if ((false && true) || (((-2.25 - 1.5) * float32(-2.25)) != 1e10)) {
                        v27 := GenS2{F0: (false && (((((1.5 - ((0.0 / -2.25) * 1.5)) * 3.14159) + -2.25) >= -2.25) || true)), F1: ""}
                        var v28 uint32 = (v7 | (v7 & v7))
                        var v29 uint = (p0 & (p0 % ((p0) | 1)))
                        v30 := GenS3{F0: 47, F1: 562}
                        _ = v27
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                    _ = fi26
                }
            }
            for fi31 := 0; fi31 < 16; fi31++ {
                p1 -= (^(p1))
                _ = fi31
            }
            _ = fi8
        }
        switch "\n\t" {
        case "foo bar":
            if false {
                v7 = v7
                for fi32 := 0; fi32 < 18; fi32++ {
                    p1 = (fi32 + (-(fi32)))
                    switch ("foo bar" + ("hello" + "hello")) {
                    case "":
                        fmt.Println(p1, p0)
                    case "hello":
                        v33 := GenS4{F0: (p1 << (((p0 * p0)) & 63)), F1: (!(true || false))}
                        _ = v33
                    default:
                        if true {
                            v34 := GenS1{F0: (0 >= 100), F1: 42}
                            _ = v34
                        } else {
                            for fi35 := 0; fi35 < 8; fi35++ {
                                if true {
                                    v36 := GenS5{F0: 1000}
                                    v37 := GenS3{F0: 42, F1: 106}
                                    _ = v36
                                    _ = v37
                                }
                                fmt.Println(p1)
                                v7 += v7
                                v38 := GenS3{F0: v7, F1: 510}
                                v39 := []int{32767, 239, 273, 201}
                                v7 = (^(v7))
                                fmt.Println(p0)
                                p0 = p0
                                if (2 <= 815) {
                                    switch (fi35 ^ (^(p1))) {
                                    case 127:
                                        v40 := GenS5{F0: 910}
                                        _ = v40
                                    case 32767:
                                        switch fi35 {
                                        case -2147483648:
                                            v7 += (v7 ^ v7)
                                        case 42:
                                            var v41 int64 = 9223372036854775807
                                            _ = v41
                                        }
                                    case 9223372036854775807:
                                        if true {
                                            if len(v39) >= 2 {
                                                v39 = v39[1:len(v39)]
                                            }
                                            v7 = v7
                                            v42 := GenS4{F0: 352, F1: true}
                                            _ = v42
                                        } else {
                                            _ = len(v39)
                                        }
                                    }
                                    v43 := GenS3{F0: (^(v7)), F1: -1}
                                    _ = v43
                                }
                                _ = fi35
                                _ = v38
                                _ = v39
                            }
                        }
                    }
                    _ = fi32
                }
                v7 += v7
            } else {
                switch "\n\t" {
                case "hello":
                    p0 -= (p0 + p0)
                case "":
                    v44 := GenS6{F0: 1, F1: int32(v7)}
                    _ = v44
                case "\n\t":
                    for fi45 := 0; fi45 < 6; fi45++ {
                        v7 -= v7
                        var v46 uint64 = 42
                        var v47 uint64 = 18446744073709551615
                        _ = fi45
                        _ = v46
                        _ = v47
                    }
                }
                fmt.Println(v7, p0, p0)
            }
        }
        _ = v7
    } else {
        if (1 >= -113) {
            if (uint8(p0) == 255) {
                var v48 int64 = 1
                p1 = (p1 | (p1 << ((p0) & 63)))
                _ = v48
            }
            for fi49 := 0; fi49 < 20; fi49++ {
                v50 := []int{9223372036854775807, fi49, (-(p1))}
                _ = fi49
                _ = v50
            }
            p0 -= p0
        }
        fmt.Println(p1, p0)
    }
    _ = p0
    _ = p1
    return -128
}

func main() {
    var v1 uint64 = 0
    var v2 uint8 = 2
    v3 := make(map[int]uint32)
    v51 := []int{128, -958}
    if true {
        for fi52 := 0; fi52 < 11; fi52++ {
            var v53 int8 = 42
            v3[fi52] = 1
            _ = fi52
            _ = v53
        }
        fmt.Println(v1)
    } else {
        v2 = v2
        if (("\n\t" == ("a" + ("" + "hello"))) || false) {
            for fi54 := 0; fi54 < 2; fi54++ {
                var v55 int64 = -1
                _ = fi54
                _ = v55
            }
        }
        var v56 int32 = -2147483648
        _ = v56
    }
    v57 := GenS7{F0: genFn0(0, (605 < 543), v2), F1: (v1 >> ((255) & 63))}
    v1 -= 830
    if (!(uint(v1) == uint(v1))) {
        v51 = append(v51, -114)
        var v58 string = ("\n\t" + "a")
        _ = v58
    } else {
        if len(v51) > 0 {
            v51[0] = 640
        }
        if false {
            if (v1 <= v1) {
                switch "" {
                case "\n\t":
                    var v59 int16 = -1
                    _ = v59
                case "hello":
                    switch 1 {
                    case 42:
                        fmt.Println(v1, v2)
                    }
                case "foo bar":
                    for fi60 := 0; fi60 < 6; fi60++ {
                        var v61 uint8 = v2
                        for fi62 := 0; fi62 < 3; fi62++ {
                            var v63 uint = 0
                            var v64 uint = v63
                            var v65 int64 = 0
                            _ = fi62
                            _ = v63
                            _ = v64
                            _ = v65
                        }
                        var v66 uint = 65535
                        _ = fi60
                        _ = v61
                        _ = v66
                    }
                }
                switch -467 {
                case 42:
                    for fi67 := 0; fi67 < 9; fi67++ {
                        for fi68 := 0; fi68 < 2; fi68++ {
                            switch (fi67 << ((4294967295) & 63)) {
                            case -9223372036854775808:
                                var v69 uint64 = (v1 - 608)
                                _ = v69
                            case 0:
                                var v70 int32 = 42
                                _ = v70
                            }
                            v2 -= (v2 ^ (v2 << ((1) & 63)))
                            v2 += 0
                            _ = fi68
                        }
                        _ = fi67
                    }
                case -2147483648:
                    if len(v51) > 0 {
                        v51[0] = -2147483648
                    }
                }
                v1 += 962
            } else {
                if (2147483648 > 1) {
                    for fi71 := 0; fi71 < 12; fi71++ {
                        fmt.Println(v1, v2, v2)
                        _ = fi71
                    }
                    if false {
                        v2 = 127
                    }
                } else {
                    for fi72 := 0; fi72 < 10; fi72++ {
                        if (!(127 <= -128)) {
                            var v73 int64 = 0
                            if (!true) {
                                if false {
                                    fmt.Println(v2, v1, v2)
                                } else {
                                    v1 += (v1 | (v1 % (((v1 % ((v1) | 1))) | 1)))
                                    for fi74 := 0; fi74 < 6; fi74++ {
                                        if (!true) {
                                            if true {
                                                v1 -= 1
                                                if false {
                                                    fmt.Println(fi72)
                                                    var v75 bool = (0 < -1000)
                                                    _ = v75
                                                }
                                                v51 = append(v51, 32768)
                                            } else {
                                                if (fi74 <= int(v73)) {
                                                    var v76 string = "\n\t"
                                                    _ = v76
                                                }
                                            }
                                        } else {
                                            if true {
                                                var v77 int16 = -32768
                                                fmt.Println(fi74, v77, fi72)
                                                v78 := GenS6{F0: 2147483647, F1: -2147483648}
                                                _ = v77
                                                _ = v78
                                            } else {
                                                if false {
                                                    var v79 string = ""
                                                    var v80 string = ""
                                                    var v81 int64 = 42
                                                    v82 := GenS1{F0: true, F1: 42}
                                                    switch "hello" {
                                                    case "":
                                                        if (v81 < (v73 % ((0) | 1))) {
                                                            if (755 > 32767) {
                                                                v2 = 1
                                                                v83 := GenS1{F0: true, F1: (^(v73))}
                                                                _ = v83
                                                            } else {
                                                                _ = len(v51)
                                                                var v84 int64 = v73
                                                                fmt.Println(fi72, fi72)
                                                                var v85 uint32 = 2
                                                                v86 := GenS6{F0: v85, F1: -1}
                                                                if (v85 <= v85) {
                                                                    var v87 int = fi74
                                                                    _ = v87
                                                                }
                                                                var v88 uint64 = (v1 & 42)
                                                                _ = v84
                                                                _ = v85
                                                                _ = v86
                                                                _ = v88
                                                            }
                                                            v51 = append(v51, (-(fi72)))
                                                            var v89 int16 = 0
                                                            _ = v89
                                                        } else {
                                                            var v90 int32 = 1
                                                            fmt.Println(fi72, v81, v73)
                                                            var v91 uint8 = v2
                                                            fmt.Println(fi74, v2)
                                                            _ = v90
                                                            _ = v91
                                                        }
                                                    case "a":
                                                        if true {
                                                            delete(v3, fi74)
                                                            v92 := GenS8{F0: ((-(v81)) >= 0)}
                                                            _ = v92
                                                        } else {
                                                            var v93 float64 = (3.14159 / 1.5)
                                                            v94 := GenS7{F0: 42, F1: v1}
                                                            _ = v93
                                                            _ = v94
                                                        }
                                                    default:
                                                        var v95 int16 = 42
                                                        _ = v95
                                                    }
                                                    v81 -= (v73 & (^(v73)))
                                                    _ = v79
                                                    _ = v80
                                                    _ = v81
                                                    _ = v82
                                                } else {
                                                    var v96 uint16 = 0
                                                    var v97 int16 = 686
                                                    _ = v96
                                                    _ = v97
                                                }
                                                v1 += 427
                                                v98 := v3[(fi72 & 62)]
                                                switch fi74 {
                                                case 32767:
                                                    _ = len(v51)
                                                case 9223372036854775807:
                                                    v98 = (v98 >> ((uint(v2)) & 63))
                                                default:
                                                    var v99 int64 = int64(v2)
                                                    _ = v99
                                                }
                                                v98 += v98
                                                fmt.Println(fi74)
                                                _ = v98
                                            }
                                        }
                                        _ = fi74
                                    }
                                    switch "foo bar" {
                                    case "hello":
                                        for fi100 := 0; fi100 < 19; fi100++ {
                                            switch (fi72 << ((141) & 63)) {
                                            case 1:
                                                switch fi100 {
                                                case 127:
                                                    var v101 int64 = v73
                                                    _ = v101
                                                case 32767:
                                                    switch fi72 {
                                                    case 409:
                                                        fmt.Println(fi100, v2)
                                                    case -2147483648:
                                                        v102 := GenS9{F0: v2, F1: 48}
                                                        _ = v102
                                                    default:
                                                        if true {
                                                            var v103 uint8 = v2
                                                            _ = v103
                                                        }
                                                    }
                                                case 842:
                                                    var v104 int64 = v73
                                                    _ = v104
                                                default:
                                                    if false {
                                                        var v105 float32 = 0.0
                                                        if (((!true) && true) || (745 < 996)) {
                                                            v106 := GenS8{F0: true}
                                                            delete(v3, 127)
                                                            switch (fi100 << ((2) & 63)) {
                                                            case -706:
                                                                if (!false) {
                                                                    if true {
                                                                        if len(v51) > 0 {
                                                                            v51[0] = fi100
                                                                        }
                                                                    } else {
                                                                        v51 = append(v51, 410)
                                                                    }
                                                                }
                                                            case -555:
                                                                switch "" {
                                                                case "\n\t":
                                                                    fmt.Println(v73, v73)
                                                                case "a":
                                                                    fmt.Println(v73, v1)
                                                                case "hello":
                                                                    var v107 uint32 = 154
                                                                    _ = v107
                                                                }
                                                            default:
                                                                switch (^(fi72)) {
                                                                case -570:
                                                                    var v108 uint16 = 32767
                                                                    _ = v108
                                                                case -9223372036854775808:
                                                                    var v109 float64 = 1e10
                                                                    _ = v109
                                                                }
                                                            }
                                                            var v110 int16 = 1000
                                                            v73 = -1
                                                            fmt.Println(v1)
                                                            _ = v106
                                                            _ = v110
                                                        }
                                                        switch -9223372036854775808 {
                                                        case 32768:
                                                            fmt.Println(v73, fi72, fi72)
                                                        default:
                                                            var v111 string = "hello"
                                                            var v112 float32 = (0.0 / (v105 * (v105 * v105)))
                                                            _ = v111
                                                            _ = v112
                                                        }
                                                        _ = v105
                                                    }
                                                }
                                            default:
                                                if (65535 <= 18446744073709551615) {
                                                    v113 := GenS7{F0: 414, F1: 4294967295}
                                                    v1 = uint64(v2)
                                                    _ = v113
                                                }
                                                var v114 bool = ((3.14159 * 3.14159) != 0.1)
                                                var v115 float64 = 0.1
                                                _ = v114
                                                _ = v115
                                            }
                                            v73 += (v73 >> ((2) & 63))
                                            v116 := v3[(fi72 - fi72)]
                                            _ = fi100
                                            _ = v116
                                        }
                                    }
                                }
                                v1 -= (v1 + (v1 | 4294967295))
                                var v117 uint8 = v2
                                _ = v117
                            } else {
                                fmt.Println(v1)
                                if true {
                                    for fi118 := 0; fi118 < 5; fi118++ {
                                        var v119 uint8 = 255
                                        v119 = v2
                                        if (-2.25 != 0.1) {
                                            v2 = (v2 & v119)
                                            switch "\n\t" {
                                            case "foo bar":
                                                fmt.Println(v2)
                                            case "a":
                                                var v120 uint16 = 650
                                                _ = v120
                                            case "\n\t":
                                                v2 += (v2 & uint8(v1))
                                            default:
                                                _ = len(v3)
                                            }
                                            v51 = append(v51, (fi118 - (fi118 >> ((4294967295) & 63))))
                                        }
                                        _ = fi118
                                        _ = v119
                                    }
                                    var v121 int32 = 42
                                    _ = v121
                                } else {
                                    v1 = 1
                                    v122 := GenS10{F0: "", F1: -762}
                                    _ = v122
                                }
                            }
                            _ = v73
                        } else {
                            switch "\n\t" {
                            case "foo bar":
                                v123 := GenS2{F0: ((((("foo bar" + ("hello" + ("foo bar" + "hello"))) + ("hello" + "\n\t")) + "foo bar") + "\n\t") == "a"), F1: ("a" + "")}
                                _ = v123
                            case "a":
                                for fi124 := 0; fi124 < 2; fi124++ {
                                    v3[fi72] = 65535
                                    if (42 != 65535) {
                                        _ = len(v3)
                                        if len(v51) > 0 {
                                            v51[0] = 0
                                        }
                                        var v125 string = ("" + ("hello" + "a"))
                                        switch ("a" + "") {
                                        case "hello":
                                            v126 := v3[fi72]
                                            delete(v3, (fi72 % (((fi124 % ((fi72) | 1))) | 1)))
                                            _ = v126
                                        case "foo bar":
                                            if (-630 < -534) {
                                                v2 += (^(v2))
                                                switch (^(fi72)) {
                                                case -2147483648:
                                                    if ((0.0 + (3.14159 + 1e10)) <= (1.5 * 0.1)) {
                                                        var v127 uint8 = v2
                                                        _ = v127
                                                    } else {
                                                        v51 = append(v51, (fi72 / (((fi124 - 551)) | 1)))
                                                    }
                                                case 42:
                                                    v2 += v2
                                                    var v128 uint64 = v1
                                                    var v129 int16 = -755
                                                    _ = v128
                                                    _ = v129
                                                }
                                            }
                                        }
                                        if false {
                                            switch "hello" {
                                            case "":
                                                if false {
                                                    v2 = uint8(v1)
                                                    if len(v51) >= 2 {
                                                        v51 = v51[1:len(v51)]
                                                    }
                                                    v125 = "hello"
                                                }
                                            }
                                            var v130 int8 = -74
                                            var v131 int = fi72
                                            _ = v130
                                            _ = v131
                                        } else {
                                            v132 := GenS8{F0: (0 == 2)}
                                            var v133 int16 = 42
                                            v134 := GenS2{F0: false, F1: v125}
                                            _ = v132
                                            _ = v133
                                            _ = v134
                                        }
                                        var v135 uint32 = 2147483647
                                        _ = v125
                                        _ = v135
                                    } else {
                                        switch "hello" {
                                        case "\n\t":
                                            v1 -= v1
                                        default:
                                            fmt.Println(fi124)
                                        }
                                        var v136 int16 = 551
                                        v1 += v1
                                        _ = v136
                                    }
                                    _ = fi124
                                }
                            default:
                                v1 = (^(v1))
                            }
                            v137 := GenS11{F0: 0, F1: -1}
                            _ = v137
                        }
                        _ = fi72
                    }
                    var v138 int16 = 0
                    if true {
                        if (0 != 0) {
                            var v139 uint8 = (v2 * v2)
                            for fi140 := 0; fi140 < 13; fi140++ {
                                for fi141 := 0; fi141 < 15; fi141++ {
                                    if false {
                                        fmt.Println(fi140, v2, v138)
                                        v142 := v3[fi140]
                                        _ = v142
                                    }
                                    _ = fi141
                                }
                                v138 += 1
                                v143 := GenS12{F0: 63, F1: "\n\t", F2: uint8(v1)}
                                _ = fi140
                                _ = v143
                            }
                            v144 := GenS13{F0: (v138 + 1000), F1: -9223372036854775808}
                            _ = v139
                            _ = v144
                        }
                    }
                    _ = v138
                }
                if (true || false) {
                    v2 = v2
                    fmt.Println(v1, v1)
                } else {
                    var v145 uint16 = uint16(v2)
                    if (127 <= 1) {
                        for fi146 := 0; fi146 < 19; fi146++ {
                            v3[fi146] = 727
                            v145 = (^(v145))
                            _ = fi146
                        }
                    }
                    for fi147 := 0; fi147 < 17; fi147++ {
                        v2 = v2
                        v1 = 986
                        var v148 int64 = 9223372036854775807
                        _ = fi147
                        _ = v148
                    }
                    _ = v145
                }
                var v149 uint64 = 18446744073709551615
                _ = v149
            }
            switch 127 {
            case 42:
                fmt.Println(v1)
            }
            v1 = v1
        } else {
            for fi150 := 0; fi150 < 16; fi150++ {
                v2 = (v2 & v2)
                _ = fi150
            }
        }
        if (!true) {
            for fi151 := 0; fi151 < 6; fi151++ {
                var v152 int8 = 34
                v153 := GenS14{F0: 42, F1: -1, F2: -930}
                if false {
                    fmt.Println(v152, fi151)
                }
                _ = fi151
                _ = v152
                _ = v153
            }
            switch 2147483647 {
            case -1:
                if len(v51) >= 2 {
                    v51 = v51[1:len(v51)]
                }
            case -257:
                v1 = v1
            case 395:
                if false {
                    var v154 bool = (!false)
                    if (-9223372036854775808 == -9223372036854775808) {
                        v3[-2147483648] = 4294967295
                    }
                    _ = v154
                }
            }
            if (1e10 >= 3.14159) {
                v2 += v2
                v155 := GenS15{F0: "\n\t", F1: 1}
                _ = v155
            } else {
                if (4294967295 != 341) {
                    fmt.Println(v1)
                    for fi156 := 0; fi156 < 9; fi156++ {
                        var v157 bool = (!true)
                        _ = fi156
                        _ = v157
                    }
                    fmt.Println(v2)
                } else {
                    v51 = append(v51, 494)
                    v2 = (v2 / ((v2) | 1))
                }
            }
        }
    }
    v158 := GenS10{F0: "a", F1: 928}
    for fi159 := 0; fi159 < 19; fi159++ {
        if true {
            v1 = v1
            v2 += v2
            var v160 int = (fi159 | 626)
            _ = v160
        } else {
            if len(v51) > 0 {
                v51[0] = 32768
            }
            fmt.Println(v1, v2)
            fmt.Println(fi159)
        }
        _ = fi159
    }
    var v161 bool = true
    v2 = v2
    var v162 float32 = ((1e10 * 0.1) + 3.14159)
    for fi163 := 0; fi163 < 10; fi163++ {
        fmt.Println(v162, v161)
        _ = fi163
    }
    if (false || true) {
        fmt.Println(v162, v162)
        v2 += (v2 & v2)
        var v164 uint8 = v2
        _ = v164
    } else {
        fmt.Println(v2, v1, v162)
        fmt.Println(v2, v162, v161)
        v162 += -2.25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v161:", v161)
    fmt.Println("v162:", v162)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v51:", len(v51), v51)
    _ = v1
    _ = v158
    _ = v161
    _ = v162
    _ = v2
    _ = v3
    _ = v51
    _ = v57
}
