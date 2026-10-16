package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int; F1 int64; F2 int16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 bool; F1 uint16; F2 int }
type GenS10 struct { F0 bool; F1 uint64 }
type GenS11 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS12 struct { F0 int64; F1 uint16 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 int }
type GenS15 struct { F0 string; F1 int32; F2 bool }
type GenS16 struct { F0 int64; F1 string }
type GenS17 struct { F0 int32; F1 int8; F2 int32 }
type GenS18 struct { F0 int32 }
type GenS19 struct { F0 uint8 }
type GenS20 struct { F0 int32; F1 int8 }
type GenS21 struct { F0 uint16; F1 int32 }
func genFn0() uint16 {
    if ("" < "") {
        var v4 bool = false
        if true {
            v5 := GenS1{F0: 0, F1: 33}
            v6 := GenS1{F0: 65535, F1: 132}
            for fi7 := 0; fi7 < 11; fi7++ {
                v8 := GenS1{F0: 4294967295, F1: 255}
                _ = fi7
                _ = v8
            }
            _ = v5
            _ = v6
        }
        _ = v4
    } else {
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 int = (fi9 % (((fi9 % ((fi9) | 1))) | 1))
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 11; fi11++ {
            v12 := []int{fi11, fi11, fi11}
            for fi13 := 0; fi13 < 6; fi13++ {
                if false {
                    if true {
                        var v14 int64 = -1
                        v15 := GenS1{F0: 2147483647, F1: 143}
                        _ = v14
                        _ = v15
                    }
                    switch "\n\t" {
                    case "hello":
                        var v16 float64 = -2.25
                        _ = v16
                    case "\n\t":
                        var v17 uint8 = 255
                        _ = v17
                    }
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                }
                _ = fi13
            }
            var v18 bool = true
            _ = fi11
            _ = v12
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        if false {
            fmt.Println(fi19)
            if (!(false && (2 >= 110))) {
                v20 := []int{fi19}
                v21 := GenS2{F0: (("a" + "a") + "\n\t")}
                _ = v20
                _ = v21
            }
            fmt.Println(fi19)
        }
        _ = fi19
    }
    v22 := []int{2147483647, 691, -2147483648}
    switch 73 {
    case 32767:
        if true {
            v23 := GenS1{F0: 2147483647, F1: 128}
            var v24 uint64 = 42
            _ = v23
            _ = v24
        }
    case -799:
        var v25 uint8 = 128
        _ = v25
    case 2147483647:
        var v26 int32 = -2147483648
        _ = v26
    default:
        if false {
            if true {
                var v27 float64 = 3.14159
                if (!(804 > -433)) {
                    var v28 int8 = 100
                    var v29 float64 = v27
                    _ = v28
                    _ = v29
                } else {
                    for fi30 := 0; fi30 < 16; fi30++ {
                        v27 += v27
                        var v31 uint = 867
                        _ = fi30
                        _ = v31
                    }
                    v27 = -2.25
                }
                _ = v27
            } else {
                switch "" {
                case "a":
                    var v32 int16 = 722
                    _ = v32
                default:
                    if (295 == 572) {
                        for fi33 := 0; fi33 < 16; fi33++ {
                            for fi34 := 0; fi34 < 17; fi34++ {
                                fmt.Println(fi34, fi33)
                                _ = fi34
                            }
                            var v35 uint8 = 2
                            _ = fi33
                            _ = v35
                        }
                    } else {
                        for fi36 := 0; fi36 < 5; fi36++ {
                            if len(v22) > 0 {
                                v22[0] = fi36
                            }
                            switch (fi36 - (^(fi36))) {
                            case 32767:
                                if len(v22) > 0 {
                                    v22[0] = (fi36 - (-(fi36)))
                                }
                            case 2147483647:
                                if (((3.14159 - 0.0) == ((1e10 + 0.1) / 1.5)) && (("" + ("a" + "hello")) != ("\n\t" + ("foo bar" + "\n\t")))) {
                                    var v37 float32 = (3.14159 * -2.25)
                                    var v38 uint8 = 255
                                    _ = v37
                                    _ = v38
                                } else {
                                    v39 := GenS1{F0: 589, F1: 2}
                                    _ = v39
                                }
                            case 128:
                                fmt.Println(fi36)
                            }
                            for fi40 := 0; fi40 < 14; fi40++ {
                                var v41 uint64 = 109
                                _ = fi40
                                _ = v41
                            }
                            _ = fi36
                        }
                        _ = len(v22)
                    }
                }
                var v42 uint16 = 65535
                var v43 float32 = 1e10
                _ = v42
                _ = v43
            }
            var v44 uint32 = 4294967295
            var v45 int8 = 127
            _ = v44
            _ = v45
        }
    }
    var v46 uint64 = 42
    _ = v22
    _ = v46
    return 1
}
func genFn1() int8 {
    var v47 int32 = -2147483648
    fmt.Println(v47)
    var v48 int32 = 1
    _ = v47
    _ = v48
    return 1
}

func main() {
    var v1 uint32 = 1
    v2 := make(map[int]bool)
    var v3 string = ""
    v49 := GenS2{F0: v3}
    v50 := GenS2{F0: ""}
    switch v3 {
    case "a":
        var v51 float32 = (1e10 - float32((((1.5 * (3.14159 * 3.14159)) - float64(float32(3.14159))) - 0.1)))
        _ = v51
    case "\n\t":
        var v52 int64 = 1
        _ = v52
    }
    if false {
        v1 = v1
        var v53 bool = true
        if (42 < 564) {
            v53 = (!true)
            var v54 int = 32767
            if (v53 || ((1.5 < ((3.14159 + float64(float32(3.14159))) * 3.14159)) && (42 != 3))) {
                v55 := []int{(v54 + (v54 & v54))}
                _ = v55
            }
            _ = v54
        } else {
            switch "foo bar" {
            case "hello":
                v56 := GenS2{F0: v3}
                _ = v56
            }
        }
        _ = v53
    } else {
        for fi57 := 0; fi57 < 20; fi57++ {
            var v58 uint = 224
            _ = len(v2)
            _ = fi57
            _ = v58
        }
    }
    v1 = v1
    switch 925 {
    case 42:
        var v59 uint16 = 65535
        _ = v59
    default:
        v2[32768] = true
    }
    v3 = v3
    for fi60 := 0; fi60 < 20; fi60++ {
        v61 := GenS3{F0: 0}
        switch v3 {
        case "hello":
            for fi62 := 0; fi62 < 13; fi62++ {
                switch fi60 {
                case 52:
                    if (127 >= 127) {
                        if true {
                            v63 := GenS4{F0: v3, F1: 1}
                            if false {
                                v64 := v2[-1]
                                var v65 uint32 = 1
                                _ = v64
                                _ = v65
                            } else {
                                fmt.Println(v1, fi60, fi62)
                                fmt.Println(v1, v3)
                                switch 32767 {
                                case 9223372036854775807:
                                    var v66 int = 127
                                    _ = v66
                                case 4294967295:
                                    if false {
                                        fmt.Println(fi60)
                                    }
                                case 2147483647:
                                    var v67 bool = false
                                    v67 = true
                                    v68 := GenS1{F0: v1, F1: 2}
                                    _ = v67
                                    _ = v68
                                }
                            }
                            v1 += (v1 / (((v1 / ((2147483648) | 1))) | 1))
                            if false {
                                v1 -= (v1 / (((v1 - (^(v1)))) | 1))
                            } else {
                                var v69 int32 = 42
                                v70 := GenS5{F0: -1}
                                if ((!(1 >= -1)) || (!(-2.25 >= 3.14159))) {
                                    fmt.Println(fi60, v69, fi60)
                                    v2[fi60] = false
                                } else {
                                    v3 = (("" + v3) + "hello")
                                    if ((false && true) || (false && true)) {
                                        v1 += (^(v1))
                                        v71 := GenS5{F0: 1}
                                        _ = v71
                                    }
                                    switch ((("foo bar" + v3) + "\n\t") + ((v3 + ((v3 + "\n\t") + (v3 + v3))) + "hello")) {
                                    case "\n\t":
                                        v72 := GenS4{F0: "a", F1: -463}
                                        _ = v72
                                    case "":
                                        if false {
                                            fmt.Println(v3, v3, fi60)
                                            v1 -= v1
                                            v73 := GenS6{F0: (v69 << ((255) & 63))}
                                            v69 -= v69
                                            _ = v73
                                        }
                                        v3 = v3
                                    default:
                                        var v74 uint = 1
                                        _ = v74
                                    }
                                    var v75 uint64 = 4294967295
                                    v69 = (v69 & v69)
                                    v76 := []int{fi62}
                                    _ = v75
                                    _ = v76
                                }
                                _ = v69
                                _ = v70
                            }
                            if true {
                                v77 := GenS2{F0: v3}
                                delete(v2, (fi60 / (((fi62 ^ fi62)) | 1)))
                                _ = v77
                            } else {
                                v78 := []int{fi60, (-(fi60)), (fi60 - -2147483648)}
                                switch 32768 {
                                case 32768:
                                    if true {
                                        v79 := GenS2{F0: (v3 + (v3 + "foo bar"))}
                                        _ = v79
                                    } else {
                                        v1 = 866
                                    }
                                case -1:
                                    if (1 <= 127) {
                                        v3 = ("a" + v3)
                                        v78 = append(v78, (fi62 ^ (fi60 % ((fi60) | 1))))
                                    }
                                }
                                if (false || (42 > 0)) {
                                    switch 127 {
                                    case 127:
                                        v1 += v1
                                    case 0:
                                        var v80 string = ("foo bar" + v3)
                                        _ = v80
                                    case 128:
                                        _ = len(v78)
                                    default:
                                        if len(v78) > 0 {
                                            v78[0] = fi62
                                        }
                                    }
                                    var v81 uint8 = 185
                                    var v82 string = ((v3 + v3) + v3)
                                    _ = v81
                                    _ = v82
                                } else {
                                    switch v3 {
                                    case "a":
                                        var v83 int64 = 483
                                        _ = v83
                                    default:
                                        if (1 <= 32767) {
                                            _ = len(v2)
                                            if true {
                                                var v84 uint16 = 0
                                                if (-9223372036854775808 > 935) {
                                                    var v85 int = fi62
                                                    var v86 uint32 = (v1 * (^(v1)))
                                                    v84 = (v84 + v84)
                                                    v1 += (v1 - v1)
                                                    v1 -= v1
                                                    _ = v85
                                                    _ = v86
                                                } else {
                                                    var v87 uint16 = genFn0()
                                                    _ = v87
                                                }
                                                if ("" != (v3 + "")) {
                                                    var v88 uint32 = (v1 * (v1 / ((v1) | 1)))
                                                    switch ("foo bar" + "hello") {
                                                    case "":
                                                        _ = len(v78)
                                                        switch "hello" {
                                                        case "hello":
                                                            v3 = ((v3 + (v3 + "a")) + v3)
                                                        default:
                                                            var v89 uint16 = (v84 ^ 2)
                                                            _ = v89
                                                        }
                                                    case "a":
                                                        var v90 uint16 = v84
                                                        _ = v90
                                                    case "\n\t":
                                                        switch ((v3 + "") + "foo bar") {
                                                        case "hello":
                                                            var v91 uint8 = 1
                                                            _ = v91
                                                        case "foo bar":
                                                            var v92 uint32 = (^(v88))
                                                            _ = v92
                                                        case "":
                                                            if len(v78) > 0 {
                                                                v78[0] = -1
                                                            }
                                                        default:
                                                            var v93 float32 = 0.0
                                                            _ = v93
                                                        }
                                                    default:
                                                        v2[fi60] = ((fi62 / ((4294967295) | 1)) > fi60)
                                                        v78 = append(v78, (fi62 - 1))
                                                    }
                                                    v94 := GenS7{F0: fi62, F1: -926, F2: -1}
                                                    v95 := GenS6{F0: 0}
                                                    _ = v88
                                                    _ = v94
                                                    _ = v95
                                                } else {
                                                    _ = len(v2)
                                                }
                                                _ = v84
                                            }
                                        } else {
                                            v1 += 2147483647
                                            var v96 int32 = 2147483647
                                            switch -2147483648 {
                                            case 9223372036854775807:
                                                v3 = v3
                                            case 519:
                                                var v97 uint16 = 255
                                                _ = v97
                                            default:
                                                var v98 bool = false
                                                _ = v98
                                            }
                                            _ = v96
                                        }
                                    }
                                    if false {
                                        if false {
                                            var v99 int = 127
                                            if true {
                                                v3 = ((v3 + v3) + v3)
                                                if len(v78) >= 2 {
                                                    v78 = v78[1:len(v78)]
                                                }
                                            }
                                            v3 = ("foo bar" + "\n\t")
                                            fmt.Println(v99, fi62)
                                            switch (v99 % ((v99) | 1)) {
                                            case 127:
                                                v100 := GenS8{F0: 218}
                                                _ = v100
                                            case 4294967295:
                                                switch (fi62 - (v99 >> ((405) & 63))) {
                                                case -149:
                                                    v99 += fi60
                                                }
                                            case -65:
                                                v1 += (v1 << ((0) & 63))
                                            default:
                                                var v101 uint16 = 0
                                                _ = v101
                                            }
                                            _ = v99
                                        }
                                    }
                                    v1 += 2147483647
                                    var v102 int32 = -2147483648
                                    v103 := GenS9{F0: (false && false), F1: 0, F2: (fi60 | (fi62 + (fi62 * fi62)))}
                                    _ = v102
                                    _ = v103
                                }
                                _ = v78
                            }
                            _ = v63
                        }
                        if (-2.25 >= 3.14159) {
                            v2[(fi62 ^ 9223372036854775807)] = ((65535 != v1) || (false && false))
                        }
                        v2[fi62] = (65535 >= 2)
                    } else {
                        if false {
                            v104 := GenS8{F0: 0}
                            var v105 int32 = -79
                            v106 := []int{(fi62 | fi62)}
                            var v107 int32 = 1
                            _ = v104
                            _ = v105
                            _ = v106
                            _ = v107
                        } else {
                            var v108 int8 = -63
                            if false {
                                v109 := []int{(fi62 | 757), -1}
                                v108 = genFn1()
                                v108 = (v108 << ((42) & 63))
                                _ = v109
                            } else {
                                v3 = "hello"
                                v110 := []int{}
                                if false {
                                    if (false || (2 > 0)) {
                                        if (1 >= -896) {
                                            fmt.Println(v3, v3)
                                        } else {
                                            v108 += v108
                                            v3 = v3
                                        }
                                        fmt.Println(v108, v1)
                                        v3 = "\n\t"
                                    } else {
                                        fmt.Println(fi60, v3, fi62)
                                        v3 = v3
                                        var v111 bool = true
                                        _ = v111
                                    }
                                    fmt.Println(fi62)
                                    v1 -= v1
                                } else {
                                    fmt.Println(fi60, fi60, fi60)
                                    v112 := GenS10{F0: false, F1: 18446744073709551615}
                                    var v113 int16 = 620
                                    _ = v112
                                    _ = v113
                                }
                                if (true && ((!true) && (2 < 0))) {
                                    v108 -= 96
                                    switch ((((v3 + (v3 + v3)) + ("a" + v3)) + (v3 + "")) + v3) {
                                    case "":
                                        v3 = ("a" + "foo bar")
                                    case "hello":
                                        v108 += (v108 / ((v108) | 1))
                                    case "foo bar":
                                        v3 = ""
                                    default:
                                        v114 := GenS4{F0: v3, F1: 976}
                                        _ = v114
                                    }
                                } else {
                                    var v115 int = fi62
                                    if (true && false) {
                                        v115 = 793
                                    } else {
                                        switch v115 {
                                        case -940:
                                            v2[fi60] = (!false)
                                        case 0:
                                            _ = len(v110)
                                        default:
                                            v1 = v1
                                        }
                                    }
                                    if true {
                                        var v116 int32 = 423
                                        switch (v3 + v3) {
                                        case "a":
                                            v117 := GenS11{F0: 696, F1: true, F2: 441}
                                            switch (^(v115)) {
                                            case -2147483648:
                                                if false {
                                                    var v118 bool = true
                                                    switch 1 {
                                                    case 0:
                                                        v116 = v116
                                                    case 128:
                                                        switch ((v3 + "hello") + (v3 + (("\n\t" + "\n\t") + "a"))) {
                                                        case "foo bar":
                                                            v115 = int(v116)
                                                            if len(v110) >= 2 {
                                                                v110 = v110[1:len(v110)]
                                                            }
                                                            v119 := GenS12{F0: 0, F1: 32767}
                                                            _ = v119
                                                        default:
                                                            var v120 bool = (((!(v3 != "a")) || ((v115 << ((508) & 63)) > (v115 / ((fi60) | 1)))) && v118)
                                                            _ = v120
                                                        }
                                                        switch (v115 / ((32768) | 1)) {
                                                        case 1:
                                                            if len(v110) >= 2 {
                                                                v110 = v110[1:len(v110)]
                                                            }
                                                        }
                                                    case 337:
                                                        v110 = append(v110, fi60)
                                                    default:
                                                        if v118 {
                                                            if v118 {
                                                                var v121 int16 = 32767
                                                                v2[fi62] = (!(255 == 18446744073709551615))
                                                                v122 := GenS13{F0: 9223372036854775807}
                                                                _ = v121
                                                                _ = v122
                                                            }
                                                            if ((v115 << ((0) & 63)) > (fi62 / (((v115 << ((0) & 63))) | 1))) {
                                                                v118 = true
                                                            } else {
                                                                if (0 != 1) {
                                                                    switch 1 {
                                                                    case -9223372036854775808:
                                                                        switch v3 {
                                                                        case "foo bar":
                                                                            v1 -= v1
                                                                        case "\n\t":
                                                                            var v123 int32 = 2147483647
                                                                            _ = v123
                                                                        }
                                                                    default:
                                                                        v116 -= v116
                                                                    }
                                                                }
                                                                v124 := GenS7{F0: (^(fi62)), F1: -353, F2: 0}
                                                                var v125 uint16 = 816
                                                                _ = v124
                                                                _ = v125
                                                            }
                                                        }
                                                        if true {
                                                            var v126 uint = 65535
                                                            _ = v126
                                                        } else {
                                                            if v118 {
                                                                if len(v110) >= 2 {
                                                                    v110 = v110[1:len(v110)]
                                                                }
                                                                v116 -= v116
                                                                v108 = int8(fi62)
                                                            } else {
                                                                v3 = ("" + v3)
                                                                fmt.Println(v118)
                                                                v127 := GenS1{F0: v1, F1: 26}
                                                                _ = v127
                                                            }
                                                            var v128 uint8 = 88
                                                            var v129 int8 = 1
                                                            _ = v128
                                                            _ = v129
                                                        }
                                                    }
                                                    fmt.Println(v116, v116)
                                                    var v130 int8 = 126
                                                    var v131 float64 = float64(float32(3.14159))
                                                    _ = v118
                                                    _ = v130
                                                    _ = v131
                                                }
                                            case -9223372036854775808:
                                                v132 := GenS14{F0: fi62}
                                                _ = v132
                                            default:
                                                if true {
                                                    switch fi60 {
                                                    case -9223372036854775808:
                                                        var v133 uint32 = v1
                                                        _ = v133
                                                    case -2147483648:
                                                        var v134 uint64 = 1
                                                        _ = v134
                                                    case 9223372036854775807:
                                                        v1 += 42
                                                    default:
                                                        v3 = "a"
                                                        var v135 int8 = (v108 ^ (v108 / ((genFn1()) | 1)))
                                                        if (fi60 != (fi60 & v115)) {
                                                            var v136 uint32 = v1
                                                            _ = v136
                                                        }
                                                        v3 = ("foo bar" + v3)
                                                        _ = v135
                                                    }
                                                    _ = len(v110)
                                                } else {
                                                    var v137 string = "a"
                                                    _ = v137
                                                }
                                            }
                                            _ = v117
                                        case "foo bar":
                                            v138 := GenS8{F0: (v116 * (v116 ^ (v116 * -128)))}
                                            _ = v138
                                        case "\n\t":
                                            var v139 float64 = 1e10
                                            _ = v139
                                        }
                                        switch v115 {
                                        case 1:
                                            var v140 int64 = -9223372036854775808
                                            _ = v140
                                        }
                                        _ = v116
                                    } else {
                                        var v141 uint32 = (v1 >> ((1) & 63))
                                        _ = v141
                                    }
                                    var v142 uint = 0
                                    _ = v115
                                    _ = v142
                                }
                                var v143 int = (-(fi62))
                                v144 := GenS5{F0: -701}
                                switch (v143 - 9223372036854775807) {
                                case 1:
                                    v3 = ("a" + (v3 + v3))
                                case 128:
                                    v108 = v108
                                case -9223372036854775808:
                                    var v145 uint64 = 42
                                    _ = v145
                                default:
                                    v146 := GenS3{F0: 255}
                                    _ = v146
                                }
                                _ = v110
                                _ = v143
                                _ = v144
                            }
                            _ = v108
                        }
                        v147 := GenS15{F0: "a", F1: 42, F2: (!false)}
                        var v148 bool = (32767 >= -981)
                        _ = v147
                        _ = v148
                    }
                case 4294967295:
                    if (82 <= 0) {
                        if (false && (0 < 42)) {
                            var v149 float32 = 3.14159
                            fmt.Println(v3, fi62, fi60)
                            v149 -= (1e10 - v149)
                            var v150 uint8 = 255
                            v1 = v1
                            if false {
                                var v151 uint16 = 238
                                var v152 uint = 91
                                _ = v151
                                _ = v152
                            } else {
                                if (4294967295 != 65535) {
                                    if true {
                                        v149 = -2.25
                                        if false {
                                            var v153 uint64 = 0
                                            v154 := GenS8{F0: -2147483648}
                                            v155 := []int{-489, (fi60 % (((fi62 / (((fi62 | fi62)) | 1))) | 1)), (fi62 + 4294967295), fi60}
                                            _ = v153
                                            _ = v154
                                            _ = v155
                                        } else {
                                            var v156 uint8 = v150
                                            _ = v156
                                        }
                                    }
                                    var v157 uint16 = 255
                                    v150 = v150
                                    _ = v157
                                }
                                v158 := GenS8{F0: 503}
                                _ = v158
                            }
                            _ = v149
                            _ = v150
                        }
                        var v159 int = (fi60 | fi60)
                        v160 := []int{(fi62 & fi62)}
                        _ = v159
                        _ = v160
                    } else {
                        fmt.Println(v1, v3, fi62)
                        fmt.Println(fi62, v1)
                    }
                default:
                    var v161 bool = false
                    _ = v161
                }
                v162 := GenS12{F0: -9223372036854775808, F1: 42}
                _ = fi62
                _ = v162
            }
        case "foo bar":
            v3 = (("" + (((v3 + "hello") + (v3 + "hello")) + "\n\t")) + "a")
        }
        _ = fi60
        _ = v61
    }
    fmt.Println(v3, v3)
    var v163 bool = false
    var v164 float32 = (3.14159 * 1e10)
    if true {
        for fi165 := 0; fi165 < 13; fi165++ {
            var v166 float32 = 1.5
            var v167 int8 = -1
            switch fi165 {
            case 32768:
                for fi168 := 0; fi168 < 19; fi168++ {
                    if ((fi165 == 4294967295) || (278 != 1)) {
                        v169 := []int{}
                        _ = v169
                    } else {
                        v170 := GenS13{F0: -1}
                        v171 := v2[-858]
                        var v172 uint16 = 32767
                        _ = v170
                        _ = v171
                        _ = v172
                    }
                    v3 = v3
                    v173 := GenS14{F0: (fi165 % ((fi168) | 1))}
                    v3 = (v3 + v3)
                    var v174 bool = v163
                    _ = fi168
                    _ = v173
                    _ = v174
                }
            case -9223372036854775808:
                if (573 <= 0) {
                    v175 := []int{fi165, (^(fi165)), (fi165 << ((uint(v1)) & 63)), (-(fi165))}
                    _ = v175
                } else {
                    var v176 float32 = (v166 + v166)
                    fmt.Println(v176, v166)
                    _ = v176
                }
            case -431:
                v177 := []int{}
                _ = v177
            default:
                v163 = false
            }
            _ = fi165
            _ = v166
            _ = v167
        }
        var v178 uint16 = 1
        v178 = 848
        _ = v178
    } else {
        var v179 float64 = float64((float32(-2.25) / float32((0.1 / -2.25))))
        v179 = v179
        if v163 {
            v1 -= (v1 % ((v1) | 1))
            for fi180 := 0; fi180 < 10; fi180++ {
                var v181 float64 = v179
                v182 := []int{fi180, (fi180 % (((fi180 + fi180)) | 1))}
                if (true || (!(!(255 < 32767)))) {
                    switch fi180 {
                    case 32768:
                        var v183 bool = (422 >= 42)
                        _ = v183
                    case 0:
                        for fi184 := 0; fi184 < 19; fi184++ {
                            var v185 int64 = 486
                            if true {
                                if true {
                                    if true {
                                        if (fi184 < 128) {
                                            var v186 int = fi184
                                            var v187 int = fi180
                                            if true {
                                                var v188 uint32 = (v1 ^ v1)
                                                v189 := GenS16{F0: -9223372036854775808, F1: v3}
                                                _ = v188
                                                _ = v189
                                            } else {
                                                v190 := GenS17{F0: -1, F1: 127, F2: -149}
                                                switch v3 {
                                                case "hello":
                                                    fmt.Println(v3, v164, v163)
                                                case "\n\t":
                                                    var v191 uint = 81
                                                    _ = v191
                                                case "a":
                                                    v192 := GenS11{F0: v1, F1: false, F2: 32767}
                                                    _ = v192
                                                default:
                                                    var v193 uint32 = v1
                                                    _ = v193
                                                }
                                                _ = v190
                                            }
                                            _ = v186
                                            _ = v187
                                        }
                                        v185 = 1
                                    } else {
                                        fmt.Println(v163)
                                        if (!((2147483647 >= 801) || v163)) {
                                            v164 = (v164 / 1.0)
                                            if true {
                                                if true {
                                                    if v163 {
                                                        switch (v3 + v3) {
                                                        case "a":
                                                            fmt.Println(fi180)
                                                        default:
                                                            fmt.Println(v185, v185)
                                                        }
                                                    } else {
                                                        switch (v3 + ((v3 + (v3 + (v3 + v3))) + v3)) {
                                                        case "foo bar":
                                                            if ((-2147483648 == -351) && v163) {
                                                                v194 := GenS18{F0: 0}
                                                                switch v3 {
                                                                case "hello":
                                                                    v185 = -1
                                                                }
                                                                _ = v194
                                                            } else {
                                                                fmt.Println(v185, fi180)
                                                                v195 := GenS3{F0: 1}
                                                                _ = v195
                                                            }
                                                        case "":
                                                            v163 = (1 == -25)
                                                        }
                                                        var v196 string = v3
                                                        if (((((v163 && (!v163)) || v163) && v163) && v163) || v163) {
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                var v197 string = (v3 + "a")
                                                                _ = v197
                                                            case "":
                                                                v3 = v196
                                                            default:
                                                                if v163 {
                                                                    var v198 int64 = -12
                                                                    if (!(100 <= 0)) {
                                                                        fmt.Println(v185)
                                                                    }
                                                                    _ = v198
                                                                }
                                                            }
                                                            _ = len(v2)
                                                            if v163 {
                                                                v182 = append(v182, 435)
                                                                v196 = v196
                                                                var v199 int32 = -1
                                                                _ = v199
                                                            }
                                                        }
                                                        _ = v196
                                                    }
                                                    v185 = int64(v1)
                                                    v200 := v2[(fi184 & fi184)]
                                                    _ = v200
                                                }
                                                if false {
                                                    v201 := GenS19{F0: 129}
                                                    _ = v201
                                                }
                                            } else {
                                                var v202 int64 = v185
                                                switch v3 {
                                                case "hello":
                                                    var v203 bool = (!(v181 != float64(v164)))
                                                    _ = v203
                                                case "\n\t":
                                                    if ((1 <= 0) || (!(!(v1 != (v1 ^ v1))))) {
                                                        v204 := v2[(fi184 * 947)]
                                                        v179 = v179
                                                        v182 = append(v182, (-(fi180)))
                                                        if len(v182) > 0 {
                                                            v182[0] = int(v1)
                                                        }
                                                        if true {
                                                            if (((v3 > v3) || (v204 || true)) && v204) {
                                                                var v205 int32 = -2147483648
                                                                if true {
                                                                    v206 := GenS19{F0: 1}
                                                                    _ = v206
                                                                } else {
                                                                    var v207 uint16 = 32767
                                                                    _ = v207
                                                                }
                                                                _ = v205
                                                            } else {
                                                                if (!false) {
                                                                    var v208 uint = 65535
                                                                    _ = v208
                                                                } else {
                                                                    if (!(!v163)) {
                                                                        v209 := GenS14{F0: (-(fi180))}
                                                                        var v210 uint = 339
                                                                        fmt.Println(v179, v204, v164)
                                                                        _ = v209
                                                                        _ = v210
                                                                    } else {
                                                                        v204 = (float64((v164 - (v164 + v164))) != (v179 + (v179 / 1.0)))
                                                                    }
                                                                }
                                                                fmt.Println(v179)
                                                            }
                                                        } else {
                                                            var v211 uint32 = v1
                                                            if true {
                                                                var v212 int16 = -7
                                                                fmt.Println(v179, v185, v1)
                                                                _ = v212
                                                            } else {
                                                                v1 = v1
                                                                switch "a" {
                                                                case "\n\t":
                                                                    fmt.Println(fi184, v204)
                                                                case "hello":
                                                                    delete(v2, fi180)
                                                                }
                                                                var v213 bool = ((int8(v185) >= -128) || (!(!(true && v204))))
                                                                _ = v213
                                                            }
                                                            if v163 {
                                                                v164 = (0.1 + (v164 * v164))
                                                                v214 := GenS10{F0: (v202 > 573), F1: 2}
                                                                _ = v214
                                                            } else {
                                                                var v215 int16 = -143
                                                                v202 = v185
                                                                v216 := GenS20{F0: -227, F1: -1}
                                                                _ = v215
                                                                _ = v216
                                                            }
                                                            _ = v211
                                                        }
                                                        _ = v204
                                                    } else {
                                                        switch "\n\t" {
                                                        case "a":
                                                            if v163 {
                                                                v217 := GenS18{F0: int32(v1)}
                                                                var v218 uint8 = 255
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    switch 128 {
                                                                    case -1:
                                                                        if len(v182) >= 2 {
                                                                            v182 = v182[1:len(v182)]
                                                                        }
                                                                    }
                                                                }
                                                                v185 = -9223372036854775808
                                                                v181 -= ((v179 - v181) + (v179 + 1e10))
                                                                _ = v217
                                                                _ = v218
                                                            }
                                                        case "\n\t":
                                                            var v219 float64 = ((v179 * (0.1 + v179)) - (float64((-2.25 * (1.5 / v164))) * -2.25))
                                                            _ = v219
                                                        case "hello":
                                                            v1 -= (v1 + v1)
                                                        }
                                                        var v220 int64 = v185
                                                        if v163 {
                                                            var v221 float64 = 1.5
                                                            _ = v221
                                                        } else {
                                                            var v222 int8 = 42
                                                            _ = len(v2)
                                                            v223 := GenS16{F0: v202, F1: v3}
                                                            var v224 int32 = 148
                                                            _ = v222
                                                            _ = v223
                                                            _ = v224
                                                        }
                                                        if len(v182) > 0 {
                                                            v182[0] = 774
                                                        }
                                                        _ = v220
                                                    }
                                                case "foo bar":
                                                    var v225 int64 = (v185 * 1)
                                                    _ = v225
                                                default:
                                                    if v163 {
                                                        if v163 {
                                                            var v226 int8 = 91
                                                            _ = v226
                                                        }
                                                        v3 = v3
                                                    }
                                                }
                                                var v227 uint32 = (^(v1))
                                                if true {
                                                    v185 = v185
                                                    delete(v2, (fi180 % ((fi184) | 1)))
                                                    switch v3 {
                                                    case "":
                                                        v179 -= 1e10
                                                    case "a":
                                                        var v228 string = v3
                                                        _ = v228
                                                    }
                                                }
                                                _ = len(v2)
                                                _ = v202
                                                _ = v227
                                            }
                                            var v229 uint8 = 42
                                            _ = v229
                                        }
                                        var v230 uint32 = v1
                                        _ = v230
                                    }
                                    v164 = 0.0
                                }
                            } else {
                                v179 += -2.25
                                _ = len(v182)
                            }
                            _ = fi184
                            _ = v185
                        }
                    case 4294967295:
                        v164 -= v164
                    }
                    var v231 int8 = -120
                    _ = v231
                }
                _ = fi180
                _ = v181
                _ = v182
            }
            v232 := GenS21{F0: 2, F1: 1}
            _ = v232
        }
        _ = v179
    }
    var v233 float64 = float64(v164)
    var v234 float32 = v164
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v163:", v163)
    fmt.Println("v164:", v164)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v233:", v233)
    fmt.Println("v234:", v234)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v163
    _ = v164
    _ = v2
    _ = v233
    _ = v234
    _ = v3
    _ = v49
    _ = v50
}
