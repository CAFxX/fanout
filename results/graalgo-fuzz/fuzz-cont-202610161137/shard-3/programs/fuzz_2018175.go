package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 int8; F2 int }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 int16; F1 uint64; F2 string }
type GenS7 struct { F0 int16; F1 uint8 }
type GenS8 struct { F0 uint32; F1 int; F2 int32 }
type GenS9 struct { F0 uint32; F1 int }
type GenS10 struct { F0 uint8; F1 uint64 }
type GenS11 struct { F0 int }
func genFn0(p0 int8) uint {
    var v7 float32 = -2.25
    switch 444 {
    case -9223372036854775808:
        for fi8 := 0; fi8 < 3; fi8++ {
            switch (fi8 - 2147483647) {
            case -204:
                var v9 int8 = int8(fi8)
                _ = v9
            case -1:
                for fi10 := 0; fi10 < 7; fi10++ {
                    fmt.Println(p0)
                    _ = fi10
                }
            case -912:
                p0 = (p0 << ((18446744073709551615) & 63))
            default:
                p0 = (p0 >> ((4294967295) & 63))
            }
            var v11 uint64 = 18446744073709551615
            v11 = 591
            _ = fi8
            _ = v11
        }
    case 127:
        if true {
            v12 := []int{14}
            var v13 uint32 = 2
            _ = v12
            _ = v13
        } else {
            if false {
                for fi14 := 0; fi14 < 8; fi14++ {
                    v15 := []int{}
                    p0 -= 42
                    p0 -= (p0 % ((p0) | 1))
                    _ = fi14
                    _ = v15
                }
            }
            for fi16 := 0; fi16 < 17; fi16++ {
                v17 := GenS1{F0: 491}
                v18 := GenS1{F0: 65535}
                _ = fi16
                _ = v17
                _ = v18
            }
            var v19 int32 = 0
            _ = v19
        }
    case 4294967295:
        if (!true) {
            switch "\n\t" {
            case "foo bar":
                if true {
                    switch "" {
                    case "\n\t":
                        fmt.Println(p0)
                    default:
                        v7 = 1.5
                    }
                    v7 -= 1e10
                    v7 -= float32(0.1)
                } else {
                    for fi20 := 0; fi20 < 3; fi20++ {
                        switch "\n\t" {
                        case "foo bar":
                            var v21 string = ("hello" + "\n\t")
                            _ = v21
                        }
                        p0 = p0
                        _ = fi20
                    }
                }
            case "hello":
                v7 += ((v7 / v7) - float32(3.14159))
            case "":
                var v22 uint64 = 42
                _ = v22
            default:
                for fi23 := 0; fi23 < 9; fi23++ {
                    var v24 string = ("" + "hello")
                    v24 = v24
                    _ = fi23
                    _ = v24
                }
            }
            var v25 uint64 = 884
            switch "foo bar" {
            case "a":
                for fi26 := 0; fi26 < 6; fi26++ {
                    v27 := GenS2{F0: 138, F1: (-(p0)), F2: (fi26 + fi26)}
                    for fi28 := 0; fi28 < 4; fi28++ {
                        v29 := GenS2{F0: 1, F1: 127, F2: 1}
                        v7 = (float32(3.14159) + (v7 - ((v7 - (float32(3.14159) / v7)) - ((1.5 - float32(float64((v7 + v7)))) / float32(3.14159)))))
                        v30 := []int{(-(fi26)), fi26, fi26}
                        switch ("\n\t" + ("hello" + "a")) {
                        case "foo bar":
                            v25 -= (v25 / (((v25 + v25)) | 1))
                        case "a":
                            var v31 uint32 = 1
                            _ = v31
                        default:
                            var v32 int = fi28
                            _ = v32
                        }
                        v25 -= v25
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                    var v33 uint64 = (v25 & 0)
                    _ = fi26
                    _ = v27
                    _ = v33
                }
            case "":
                v34 := []int{}
                _ = v34
            case "foo bar":
                v35 := GenS3{F0: 18446744073709551615, F1: 2}
                _ = v35
            }
            _ = v25
        } else {
            switch 128 {
            case 2147483647:
                v36 := GenS2{F0: 150, F1: p0, F2: 0}
                _ = v36
            case 0:
                v37 := GenS2{F0: 127, F1: (p0 - -104), F2: 814}
                _ = v37
            case 356:
                var v38 int8 = (p0 >> ((1) & 63))
                _ = v38
            default:
                p0 = p0
            }
        }
    default:
        v7 += ((v7 * (float32((0.1 - float64(v7))) - v7)) + ((float32(-2.25) - float32(((3.14159 * (1e10 + (-2.25 * (1.5 / float64(v7))))) - (float64(v7) * 3.14159)))) * float32(float64(v7))))
    }
    _ = p0
    _ = v7
    return 18446744073709551615
}
func genFn1() string {
    var v39 int = 127
    v40 := GenS4{F0: 32767, F1: 42}
    _ = v39
    _ = v40
    return (("hello" + "\n\t") + "\n\t")
}

func main() {
    var v1 int16 = 1000
    v2 := make(map[uint32]string)
    var v3 float64 = (1e10 * 3.14159)
    v4 := []string{"foo bar"}
    var v5 int8 = 31
    var v6 int32 = 0
    v41 := GenS1{F0: 454}
    for fi42 := 0; fi42 < 3; fi42++ {
        v1 += 1
        for fi43 := 0; fi43 < 15; fi43++ {
            switch ("\n\t" + ("" + "hello")) {
            case "hello":
                var v44 int32 = v6
                _ = v44
            default:
                switch "" {
                case "":
                    if true {
                        if ("a" == ("" + "")) {
                            switch "foo bar" {
                            case "a":
                                if ((fi42 > fi43) || ((v1 & v1) == 624)) {
                                    fmt.Println(fi43)
                                    switch int(v5) {
                                    case 1:
                                        delete(v2, 65535)
                                    case 980:
                                        var v45 uint = 42
                                        _ = v45
                                    }
                                }
                            default:
                                var v46 float32 = (1e10 - (1.5 + 0.0))
                                _ = v46
                            }
                        }
                        v47 := GenS5{F0: 841, F1: -955}
                        _ = v47
                    }
                case "foo bar":
                    var v48 uint32 = 2
                    _ = v48
                }
            }
            var v49 float32 = ((((1.5 - 3.14159) * 1.5) - 3.14159) / 1e10)
            _ = fi43
            _ = v49
        }
        v6 = v6
        _ = fi42
    }
    var v50 string = (("" + "") + "\n\t")
    v50 = v50
    for fi51 := 0; fi51 < 19; fi51++ {
        fmt.Println(fi51)
        if len(v4) > 0 {
            v4[0] = v50
        }
        if true {
            var v52 int64 = 42
            v50 = ("foo bar" + "a")
            if false {
                v52 = v52
                fmt.Println(fi51, v6, v6)
                v5 -= -1
            } else {
                v6 = (v6 * v6)
                if (((!((v52 % (((v52 / ((9223372036854775807) | 1))) | 1)) < (v52 + 882))) && (4294967295 < 65535)) && ((^(fi51)) < (^(fi51)))) {
                    v5 += 71
                    _ = len(v2)
                } else {
                    v52 = 0
                }
                var v53 float64 = (v3 + float64(3.14159))
                _ = v53
            }
            _ = v52
        }
        _ = fi51
    }
    var v54 uint64 = 693
    v54 = v54
    switch 328 {
    case -618:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case 1:
        v55 := GenS1{F0: 17}
        _ = v55
    }
    fmt.Println(v6)
    var v56 int64 = 0
    for fi57 := 0; fi57 < 14; fi57++ {
        v58 := v2[2147483647]
        _ = fi57
        _ = v58
    }
    fmt.Println(v54, v54)
    var v59 string = v50
    v60 := v2[42]
    v61 := GenS2{F0: 2, F1: (v5 / (((v5 | v5)) | 1)), F2: 78}
    for fi62 := 0; fi62 < 16; fi62++ {
        for fi63 := 0; fi63 < 16; fi63++ {
            _ = len(v4)
            _ = fi63
        }
        _ = fi62
    }
    v5 -= 50
    if true {
        for fi64 := 0; fi64 < 8; fi64++ {
            if ((-1 != 944) || (42 >= 255)) {
                v54 = (v54 & v54)
            } else {
                v59 = "\n\t"
                v65 := GenS2{F0: 255, F1: v5, F2: (^(fi64))}
                if true {
                    var v66 int = (fi64 << ((0) & 63))
                    _ = v66
                } else {
                    v67 := GenS5{F0: 0, F1: v56}
                    _ = len(v4)
                    if false {
                        v5 = (v5 >> ((606) & 63))
                    } else {
                        for fi68 := 0; fi68 < 11; fi68++ {
                            var v69 string = (v60 + "hello")
                            var v70 int16 = (-(v1))
                            _ = len(v4)
                            switch (("hello" + "hello") + "hello") {
                            case "\n\t":
                                switch 32768 {
                                case 852:
                                    fmt.Println(v59, v50)
                                case 128:
                                    v71 := v2[4294967295]
                                    if ((fi68 / ((fi68) | 1)) == fi64) {
                                        if len(v4) > 0 {
                                            v4[0] = (v71 + v50)
                                        }
                                        if true {
                                            var v72 uint8 = 2
                                            if (32768 <= (-(fi64))) {
                                                switch (-(fi64)) {
                                                case 9223372036854775807:
                                                    v50 = "foo bar"
                                                }
                                            } else {
                                                var v73 float32 = ((1e10 + 1.5) / ((1e10 + 1e10) - 1e10))
                                                _ = v73
                                            }
                                            var v74 bool = false
                                            _ = v72
                                            _ = v74
                                        }
                                    } else {
                                        var v75 int = (fi64 >> ((127) & 63))
                                        _ = v75
                                    }
                                    if false {
                                        var v76 int32 = (-(v6))
                                        v77 := v2[0]
                                        fmt.Println(v60, fi64)
                                        _ = v76
                                        _ = v77
                                    }
                                    var v78 int16 = (-(v70))
                                    _ = v71
                                    _ = v78
                                case -273:
                                    switch 32767 {
                                    case -232:
                                        if (true && true) {
                                            fmt.Println(v60, v50)
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                    default:
                                        v5 = (v5 & -51)
                                    }
                                default:
                                    if (((v70 >> ((1) & 63)) > v1) && (uint16(v54) < 1)) {
                                        var v79 uint64 = 0
                                        _ = v79
                                    }
                                }
                            case "a":
                                var v80 uint64 = v54
                                _ = v80
                            case "hello":
                                v81 := v2[1]
                                _ = v81
                            default:
                                var v82 uint = 641
                                _ = v82
                            }
                            fmt.Println(v70)
                            var v83 int64 = v56
                            _ = fi68
                            _ = v69
                            _ = v70
                            _ = v83
                        }
                        var v84 int32 = (v6 / ((v6) | 1))
                        var v85 int = 1
                        _ = v84
                        _ = v85
                    }
                    _ = v67
                }
                _ = v65
            }
            for fi86 := 0; fi86 < 9; fi86++ {
                if true {
                    if false {
                        var v87 int32 = (v6 * v6)
                        var v88 uint16 = 1
                        _ = v87
                        _ = v88
                    } else {
                        v59 = ((v59 + (v50 + v59)) + v60)
                    }
                    var v89 string = "hello"
                    _ = v89
                } else {
                    if (2 > 1) {
                        var v90 uint8 = uint8(v54)
                        fmt.Println(v50, v54)
                        _ = v90
                    } else {
                        v56 -= 0
                        switch "" {
                        case "a":
                            v2[2147483648] = v50
                        case "hello":
                            var v91 uint16 = 44
                            _ = v91
                        default:
                            if false {
                                switch ("\n\t" + v50) {
                                case "":
                                    var v92 int32 = v6
                                    _ = v92
                                case "\n\t":
                                    v1 = (^(v1))
                                default:
                                    fmt.Println(v5)
                                }
                                v60 = "a"
                            } else {
                                if false {
                                    v54 = (v54 - v54)
                                    if true {
                                        var v93 float32 = ((float32(v3) + 1.5) / -2.25)
                                        _ = v93
                                    }
                                    var v94 float64 = ((float64(1e10) / (float64(-2.25) - 0.0)) - 0.1)
                                    if false {
                                        v59 = "a"
                                        var v95 int16 = v1
                                        if (!(!(42 < 65535))) {
                                            var v96 string = v50
                                            _ = v96
                                        } else {
                                            var v97 float64 = -2.25
                                            if true {
                                                if (1 > 127) {
                                                    v98 := GenS5{F0: 42, F1: (v56 * (v56 & (v56 << ((2) & 63))))}
                                                    var v99 int64 = v56
                                                    v50 = ""
                                                    _ = v98
                                                    _ = v99
                                                } else {
                                                    v4 = append(v4, v50)
                                                    _ = len(v2)
                                                }
                                                var v100 float64 = v97
                                                _ = v100
                                            }
                                            v101 := v2[0]
                                            _ = v101
                                            _ = v97
                                        }
                                        var v102 bool = (255 >= 2)
                                        _ = v102
                                        _ = v95
                                    }
                                    v59 = v50
                                    var v103 string = "hello"
                                    var v104 bool = (v1 < 32767)
                                    _ = v103
                                    _ = v104
                                    _ = v94
                                }
                            }
                            if true {
                                v5 = -128
                                v105 := GenS5{F0: 622, F1: v56}
                                _ = v105
                            } else {
                                var v106 uint64 = v54
                                var v107 int8 = v5
                                var v108 uint64 = v106
                                if true {
                                    v1 = (v1 >> ((410) & 63))
                                }
                                v109 := GenS6{F0: (v1 ^ v1), F1: (v106 & v106), F2: ""}
                                _ = v106
                                _ = v107
                                _ = v108
                                _ = v109
                            }
                            switch -2147483648 {
                            case 128:
                                if true {
                                    fmt.Println(fi64, v3)
                                } else {
                                    fmt.Println(v3, v54, v59)
                                }
                            case -2147483648:
                                if (false || false) {
                                    var v110 float32 = float32(v3)
                                    v5 -= v5
                                    fmt.Println(v6, v3)
                                    _ = v110
                                } else {
                                    var v111 bool = false
                                    _ = v111
                                }
                            case 230:
                                if false {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    var v112 int8 = v5
                                    v1 = (v1 ^ (v1 - v1))
                                    _ = v112
                                } else {
                                    v56 -= (v56 % ((42) | 1))
                                }
                            }
                        }
                        v113 := GenS7{F0: (v1 * (v1 - v1)), F1: 1}
                        _ = v113
                    }
                    var v114 uint32 = 0
                    var v115 uint64 = (v54 | (v54 % ((v54) | 1)))
                    _ = v114
                    _ = v115
                }
                if true {
                    if false {
                        switch (((v60 + ((v60 + "foo bar") + "a")) + v50) + v59) {
                        case "a":
                            delete(v2, 65535)
                        default:
                            var v116 uint64 = (v54 % ((v54) | 1))
                            _ = v116
                        }
                        var v117 int16 = -765
                        v6 = (v6 ^ (^(v6)))
                        _ = v117
                    }
                } else {
                    fmt.Println(v50, v50)
                }
                switch "foo bar" {
                case "":
                    var v118 float64 = ((v3 / v3) + v3)
                    _ = v118
                case "foo bar":
                    v54 = 4294967295
                case "hello":
                    var v119 uint32 = 2147483648
                    _ = v119
                default:
                    if true {
                        var v120 uint8 = 255
                        v54 = (v54 / (((^(v54))) | 1))
                        if (!false) {
                            if false {
                                var v121 uint8 = (v120 % ((42) | 1))
                                _ = v121
                            } else {
                                var v122 float32 = 0.1
                                switch (("\n\t" + v50) + (v59 + v60)) {
                                case "hello":
                                    switch 4294967295 {
                                    case 4294967295:
                                        v123 := GenS8{F0: 2147483647, F1: -665, F2: (v6 << ((4294967295) & 63))}
                                        _ = v123
                                    case 128:
                                        v124 := GenS8{F0: 42, F1: (fi86 & (fi64 + fi86)), F2: (^(v6))}
                                        _ = v124
                                    case 127:
                                        var v125 int16 = int16(v120)
                                        _ = v125
                                    }
                                }
                                var v126 int16 = v1
                                _ = v122
                                _ = v126
                            }
                            fmt.Println(v6, v50, fi64)
                            v127 := v2[471]
                            _ = v127
                        } else {
                            v120 = (v120 & 128)
                            if true {
                                if (fi86 <= (^(fi86))) {
                                    v3 += -2.25
                                    if len(v4) > 0 {
                                        v4[0] = v59
                                    }
                                }
                                v120 = v120
                            } else {
                                if (true && (v54 >= v54)) {
                                    v128 := GenS2{F0: 0, F1: 100, F2: (fi86 * -895)}
                                    v60 = "foo bar"
                                    _ = v128
                                } else {
                                    var v129 bool = (-2147483648 <= (v6 ^ (v6 & (v6 << ((1) & 63)))))
                                    if v129 {
                                        var v130 float32 = 0.1
                                        _ = v130
                                    }
                                    switch "foo bar" {
                                    case "hello":
                                        v131 := GenS9{F0: 2147483647, F1: (fi86 - (fi86 + 9223372036854775807))}
                                        _ = v131
                                    }
                                    if v129 {
                                        var v132 int8 = (v5 ^ (^(v5)))
                                        _ = v132
                                    }
                                    var v133 string = ""
                                    var v134 int = 42
                                    v59 = "foo bar"
                                    _ = v129
                                    _ = v133
                                    _ = v134
                                }
                                v54 = 1
                                v120 -= v120
                                fmt.Println(v56, fi64, v50)
                            }
                            v2[4294967295] = v59
                            var v135 int32 = (v6 >> ((0) & 63))
                            _ = v135
                        }
                        _ = v120
                    }
                }
                _ = fi86
            }
            _ = fi64
        }
        switch -2147483648 {
        case -1:
            var v136 float32 = ((1.5 * float32(v3)) * 3.14159)
            _ = v136
        case -423:
            for fi137 := 0; fi137 < 8; fi137++ {
                for fi138 := 0; fi138 < 12; fi138++ {
                    if (v56 > v56) {
                        var v139 bool = true
                        fmt.Println(v6)
                        switch fi138 {
                        case 42:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case -2147483648:
                            v59 = v60
                        case 886:
                            if len(v4) > 0 {
                                v4[0] = v60
                            }
                        }
                        _ = v139
                    }
                    _ = fi138
                }
                _ = fi137
            }
        case 32768:
            v56 = -9
        default:
            if true {
                for fi140 := 0; fi140 < 9; fi140++ {
                    v50 = (v50 + "foo bar")
                    switch -2147483648 {
                    case 0:
                        v56 = 1
                    case 369:
                        for fi141 := 0; fi141 < 13; fi141++ {
                            switch "\n\t" {
                            case "a":
                                var v142 uint = 42
                                _ = v142
                            default:
                                v143 := GenS10{F0: 0, F1: (^(v54))}
                                _ = v143
                            }
                            v54 = v54
                            if false {
                                v5 = v5
                                switch v59 {
                                case "foo bar":
                                    var v144 int16 = v1
                                    _ = v144
                                default:
                                    fmt.Println(fi141, v56, v60)
                                }
                            }
                            _ = fi141
                        }
                    default:
                        var v145 float64 = 3.14159
                        _ = v145
                    }
                    _ = fi140
                }
            } else {
                for fi146 := 0; fi146 < 9; fi146++ {
                    v4 = append(v4, v59)
                    var v147 int16 = (v1 + (v1 - -32768))
                    v148 := GenS11{F0: fi146}
                    _ = fi146
                    _ = v147
                    _ = v148
                }
                var v149 int64 = (v56 & int64(v1))
                _ = v149
            }
        }
    } else {
        if (!false) {
            v6 -= (v6 ^ v6)
            fmt.Println(v5)
            var v150 int8 = v5
            _ = v150
        }
        v60 = (v60 + v60)
        if (v1 == (v1 * int16(v56))) {
            v6 = v6
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v54:", v54)
    fmt.Println("v56:", v56)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v50
    _ = v54
    _ = v56
    _ = v59
    _ = v6
    _ = v60
    _ = v61
}
