package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 int64; F2 uint64 }
type GenS5 struct { F0 bool; F1 uint64; F2 int8 }
type GenS6 struct { F0 uint64; F1 int8; F2 int }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int; F1 uint16 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int; F1 uint8; F2 uint8 }
type GenS11 struct { F0 uint64; F1 int8 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 string; F1 int16; F2 int64 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int32; F1 uint32 }
type GenS16 struct { F0 uint; F1 int32 }
type GenS17 struct { F0 uint16; F1 uint16 }
func genFn0() string {
    v7 := GenS1{F0: 255, F1: 0, F2: 2147483647}
    var v8 int16 = 1
    fmt.Println(v8)
    if (false || (!true)) {
        for fi9 := 0; fi9 < 2; fi9++ {
            switch fi9 {
            case 127:
                v10 := []int{fi9, (fi9 - 2147483647)}
                _ = v10
            case 992:
                v11 := GenS2{F0: 42, F1: 0}
                _ = v11
            default:
                for fi12 := 0; fi12 < 4; fi12++ {
                    v8 = v8
                    v8 = (^(v8))
                    _ = fi12
                }
            }
            for fi13 := 0; fi13 < 6; fi13++ {
                v8 = v8
                switch ("foo bar" + "\n\t") {
                case "foo bar":
                    v8 = v8
                }
                _ = fi13
            }
            _ = fi9
        }
        if false {
            v14 := []int{2147483647, 9223372036854775807, 4294967295, 672}
            var v15 float64 = 1e10
            _ = v14
            _ = v15
        }
        switch -536 {
        case 42:
            v16 := []int{9223372036854775807}
            _ = v16
        case -1:
            fmt.Println(v8)
        case 4294967295:
            var v17 uint8 = 255
            _ = v17
        default:
            for fi18 := 0; fi18 < 14; fi18++ {
                switch (fi18 >> ((18446744073709551615) & 63)) {
                case 42:
                    if (false || false) {
                        v19 := GenS3{F0: (("hello" + ("hello" + ((("\n\t" + "a") + "hello") + ((("foo bar" + "\n\t") + "a") + "foo bar")))) + "")}
                        var v20 uint8 = 216
                        _ = v19
                        _ = v20
                    } else {
                        v8 = (v8 ^ v8)
                    }
                }
                _ = fi18
            }
        }
    }
    if (false || (-635 == -74)) {
        var v21 int32 = -2147483648
        v8 = 42
        _ = v21
    }
    _ = v7
    _ = v8
    return "a"
}

func main() {
    v1 := []string{"hello"}
    var v2 bool = false
    var v3 int32 = 2147483647
    v4 := make(map[int]uint32)
    var v5 float64 = 1.5
    var v6 float32 = (((-2.25 * (3.14159 + (1.5 / 3.14159))) / float32(0.1)) - float32(float64(0.1)))
    var v22 uint8 = 108
    if (0 <= 1) {
        var v23 int16 = 88
        _ = v23
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        if (!true) {
            v25 := GenS4{F0: "a", F1: -9223372036854775808, F2: 42}
            _ = v25
        } else {
            if v2 {
                for fi26 := 0; fi26 < 7; fi26++ {
                    if (v6 <= 0.1) {
                        if (1 != 2147483648) {
                            if len(v1) > 0 {
                                v1[0] = ("\n\t" + "\n\t")
                            }
                            v27 := GenS5{F0: (1 != 0), F1: 0, F2: 127}
                            v6 -= (v6 - v6)
                            _ = v27
                        }
                        switch fi26 {
                        case 637:
                            if false {
                                switch "a" {
                                case "":
                                    var v28 float32 = v6
                                    _ = v28
                                case "\n\t":
                                    fmt.Println(v22)
                                default:
                                    if (!v2) {
                                        v29 := v4[(-(fi26))]
                                        var v30 int = (fi26 | fi24)
                                        _ = v29
                                        _ = v30
                                    }
                                }
                                var v31 uint32 = 1
                                _ = v31
                            } else {
                                if (!((v2 && true) && v2)) {
                                    if (false && true) {
                                        if v2 {
                                            v1 = append(v1, "foo bar")
                                        } else {
                                            switch fi26 {
                                            case 2147483647:
                                                fmt.Println(v5)
                                            case 9223372036854775807:
                                                var v32 string = "foo bar"
                                                _ = v32
                                            default:
                                                v33 := GenS3{F0: "foo bar"}
                                                _ = v33
                                            }
                                            fmt.Println(v5)
                                        }
                                        v22 = (v22 ^ v22)
                                        v2 = (v6 < (v6 + v6))
                                    } else {
                                        v34 := GenS1{F0: v22, F1: 1, F2: -698}
                                        fmt.Println(fi26)
                                        switch (^(fi26)) {
                                        case 42:
                                            v3 += v3
                                        case -9223372036854775808:
                                            var v35 uint16 = 2
                                            _ = v35
                                        case -925:
                                            var v36 string = "a"
                                            _ = v36
                                        default:
                                            _ = len(v1)
                                        }
                                        _ = v34
                                    }
                                }
                            }
                        default:
                            var v37 float32 = 1.5
                            _ = v37
                        }
                        v38 := GenS1{F0: (v22 - v22), F1: -9223372036854775808, F2: (fi26 >> ((18446744073709551615) & 63))}
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if (fi26 > fi24) {
                            var v39 int8 = 53
                            v40 := GenS3{F0: ""}
                            if (v2 && (!v2)) {
                                fmt.Println(v3, fi26)
                                var v41 string = ""
                                v42 := v4[9223372036854775807]
                                _ = v41
                                _ = v42
                            }
                            _ = v39
                            _ = v40
                        }
                        var v43 float64 = (1e10 - 1e10)
                        _ = v38
                        _ = v43
                    }
                    v3 = v3
                    var v44 int16 = 0
                    _ = fi26
                    _ = v44
                }
                fmt.Println(fi24)
            }
            var v45 float64 = (v5 * v5)
            if (!true) {
                v46 := GenS1{F0: v22, F1: 723, F2: 4294967295}
                _ = v46
            }
            _ = v45
        }
        _ = fi24
    }
    for fi47 := 0; fi47 < 17; fi47++ {
        var v48 uint32 = 778
        switch "foo bar" {
        case "hello":
            if v2 {
                _ = len(v1)
            }
        case "\n\t":
            if v2 {
                var v49 int64 = 1
                _ = v49
            }
        case "a":
            v22 = (v22 / (((^(v22))) | 1))
        default:
            v2 = v2
        }
        _ = fi47
        _ = v48
    }
    switch "a" {
    case "hello":
        v22 = v22
    }
    if false {
        for fi50 := 0; fi50 < 12; fi50++ {
            fmt.Println(v6)
            v22 = 127
            _ = fi50
        }
        for fi51 := 0; fi51 < 13; fi51++ {
            for fi52 := 0; fi52 < 6; fi52++ {
                var v53 float64 = v5
                if (0 == 42) {
                    var v54 uint64 = 1
                    v3 = (v3 / ((v3) | 1))
                    switch ("hello" + "hello") {
                    case "":
                        if v2 {
                            if (2 <= 18446744073709551615) {
                                if (v22 == (v22 % ((v22) | 1))) {
                                    v55 := GenS3{F0: "\n\t"}
                                    switch "foo bar" {
                                    case "a":
                                        switch "foo bar" {
                                        case "":
                                            var v56 uint32 = 4294967295
                                            _ = v56
                                        case "foo bar":
                                            v57 := GenS6{F0: (v54 >> ((339) & 63)), F1: 100, F2: fi52}
                                            _ = v57
                                        default:
                                            _ = len(v1)
                                        }
                                    case "hello":
                                        switch "foo bar" {
                                        case "a":
                                            v58 := GenS7{F0: 1}
                                            _ = v58
                                        default:
                                            if (-30 == -128) {
                                                v59 := GenS8{F0: (fi52 * fi51), F1: 2}
                                                if (42 != -32768) {
                                                    if true {
                                                        var v60 uint = 4294967295
                                                        v3 = (v3 + (v3 ^ v3))
                                                        var v61 uint = v60
                                                        if true {
                                                            switch fi52 {
                                                            case -2147483648:
                                                                v1 = append(v1, "")
                                                            case 900:
                                                                var v62 bool = v2
                                                                _ = v62
                                                            case 381:
                                                                v61 = (v60 - (v60 % ((414) | 1)))
                                                            default:
                                                                v63 := GenS9{F0: false}
                                                                _ = v63
                                                            }
                                                        } else {
                                                            var v64 uint16 = 255
                                                            if v2 {
                                                                var v65 uint = 4294967295
                                                                var v66 float32 = 1e10
                                                                _ = v65
                                                                _ = v66
                                                            }
                                                            _ = v64
                                                        }
                                                        v67 := GenS2{F0: 42, F1: (fi51 % (((fi52 % ((fi52) | 1))) | 1))}
                                                        if (!v2) {
                                                            if len(v1) > 0 {
                                                                v1[0] = "hello"
                                                            }
                                                            var v68 bool = ("a" >= "\n\t")
                                                            switch "hello" {
                                                            case "hello":
                                                                switch (^(fi51)) {
                                                                case 1:
                                                                    v54 += v54
                                                                    var v69 float32 = v6
                                                                    _ = v69
                                                                case 0:
                                                                    if true {
                                                                        v70 := GenS4{F0: "", F1: 42, F2: (v54 >> ((v61) & 63))}
                                                                        _ = v70
                                                                    } else {
                                                                        if false {
                                                                            fmt.Println(v60)
                                                                            var v71 int8 = 127
                                                                            v61 = 0
                                                                            _ = v71
                                                                        } else {
                                                                            if false {
                                                                                var v72 int32 = v3
                                                                                switch (("\n\t" + "foo bar") + "a") {
                                                                                case "a":
                                                                                    v4[(fi52 << ((642) & 63))] = 0
                                                                                    switch fi52 {
                                                                                    case 2147483647:
                                                                                        var v73 int16 = -303
                                                                                        _ = v73
                                                                                    default:
                                                                                        if ((2147483648 <= 2) && v68) {
                                                                                            var v74 int8 = 1
                                                                                            if v68 {
                                                                                                switch (fi51 * (fi52 << (((^(v60))) & 63))) {
                                                                                                case 32767:
                                                                                                    var v75 int8 = -56
                                                                                                    _ = v75
                                                                                                case 561:
                                                                                                    if v68 {
                                                                                                        v53 -= float64((((1.5 + (v6 / 1e10)) - (-2.25 + v6)) / (((3.14159 / v6) * v6) / 1.0)))
                                                                                                    }
                                                                                                }
                                                                                                if false {
                                                                                                    v4[323] = 2
                                                                                                }
                                                                                                if true {
                                                                                                    v76 := GenS1{F0: v22, F1: 9223372036854775807, F2: fi52}
                                                                                                    var v77 string = "hello"
                                                                                                    if v2 {
                                                                                                        v2 = ((v77 + v77) == v77)
                                                                                                        if false {
                                                                                                            v6 = v6
                                                                                                            switch fi52 {
                                                                                                            case 930:
                                                                                                                v53 -= 3.14159
                                                                                                            case 1:
                                                                                                                var v78 float64 = 0.0
                                                                                                                _ = v78
                                                                                                            case -1:
                                                                                                                v79 := GenS10{F0: 32767, F1: (^(v22)), F2: 176}
                                                                                                                _ = v79
                                                                                                            default:
                                                                                                                v53 = v5
                                                                                                            }
                                                                                                        }
                                                                                                        fmt.Println(v72, v61, v68)
                                                                                                        fmt.Println(v2, v61)
                                                                                                        v53 += (float64(v6) + v5)
                                                                                                        delete(v4, fi52)
                                                                                                        v4[fi51] = 2147483647
                                                                                                    } else {
                                                                                                        if (!v68) {
                                                                                                            var v80 uint = v60
                                                                                                            var v81 uint8 = (v22 << ((18446744073709551615) & 63))
                                                                                                            _ = v80
                                                                                                            _ = v81
                                                                                                        }
                                                                                                        if (v6 >= 1e10) {
                                                                                                            v4[0] = 662
                                                                                                        } else {
                                                                                                            v3 = (v3 << ((0) & 63))
                                                                                                            var v82 uint32 = 629
                                                                                                            v83 := GenS2{F0: v82, F1: (fi51 >> (((v60 + 42)) & 63))}
                                                                                                            _ = v82
                                                                                                            _ = v83
                                                                                                        }
                                                                                                        v3 = -1
                                                                                                    }
                                                                                                    _ = v76
                                                                                                    _ = v77
                                                                                                } else {
                                                                                                    var v84 int64 = -810
                                                                                                    _ = v84
                                                                                                }
                                                                                            }
                                                                                            var v85 float32 = v6
                                                                                            v68 = (!true)
                                                                                            _ = len(v1)
                                                                                            _ = v74
                                                                                            _ = v85
                                                                                        }
                                                                                    }
                                                                                    switch ("a" + ((("a" + "\n\t") + "\n\t") + "\n\t")) {
                                                                                    case "foo bar":
                                                                                        v3 = v3
                                                                                        v54 += (v54 ^ v54)
                                                                                        v86 := GenS9{F0: false}
                                                                                        _ = v86
                                                                                    case "a":
                                                                                        var v87 float32 = v6
                                                                                        _ = v87
                                                                                    case "hello":
                                                                                        var v88 bool = false
                                                                                        _ = v88
                                                                                    default:
                                                                                        _ = len(v1)
                                                                                    }
                                                                                case "hello":
                                                                                    var v89 uint64 = (^(v54))
                                                                                    _ = v89
                                                                                case "foo bar":
                                                                                    var v90 int16 = 42
                                                                                    _ = v90
                                                                                default:
                                                                                    v91 := GenS7{F0: (^(v61))}
                                                                                    _ = v91
                                                                                }
                                                                                v60 -= 4294967295
                                                                                _ = v72
                                                                            }
                                                                        }
                                                                        if v68 {
                                                                            fmt.Println(v22, v22)
                                                                            var v92 int8 = 42
                                                                            _ = v92
                                                                        }
                                                                        if len(v1) > 0 {
                                                                            v1[0] = "\n\t"
                                                                        }
                                                                    }
                                                                default:
                                                                    var v93 string = ""
                                                                    _ = v93
                                                                }
                                                                var v94 uint32 = 65535
                                                                _ = v94
                                                            case "a":
                                                                v95 := GenS3{F0: (("foo bar" + "foo bar") + ("foo bar" + "a"))}
                                                                _ = v95
                                                            }
                                                            fmt.Println(v6, v61, v54)
                                                            _ = v68
                                                        } else {
                                                            delete(v4, 9223372036854775807)
                                                        }
                                                        _ = v60
                                                        _ = v61
                                                        _ = v67
                                                    }
                                                }
                                                var v96 int32 = (v3 + int32(v22))
                                                _ = v59
                                                _ = v96
                                            } else {
                                                var v97 uint = 65535
                                                fmt.Println(v97, v3, v2)
                                                var v98 int32 = (v3 << (((v97 + 825)) & 63))
                                                _ = v97
                                                _ = v98
                                            }
                                        }
                                    }
                                    _ = v55
                                } else {
                                    fmt.Println(v22, v3)
                                    var v99 uint32 = 0
                                    _ = v99
                                }
                                v54 += (v54 >> ((287) & 63))
                                fmt.Println(v5)
                            } else {
                                switch fi51 {
                                case 42:
                                    v6 -= ((1.5 - 3.14159) * 0.1)
                                default:
                                    if (!v2) {
                                        switch "\n\t" {
                                        case "a":
                                            v100 := GenS4{F0: "foo bar", F1: -14, F2: v54}
                                            _ = v100
                                        case "":
                                            var v101 bool = (true && (uint(v54) == 4294967295))
                                            _ = v101
                                        }
                                        switch fi52 {
                                        case 80:
                                            var v102 uint = 65535
                                            _ = v102
                                        case -211:
                                            v103 := GenS7{F0: 18446744073709551615}
                                            _ = v103
                                        case 128:
                                            v104 := GenS4{F0: ("foo bar" + ("hello" + ("\n\t" + ("\n\t" + "\n\t")))), F1: int64(v22), F2: 42}
                                            _ = v104
                                        default:
                                            v105 := GenS2{F0: 4294967295, F1: int(v54)}
                                            _ = v105
                                        }
                                        v106 := GenS11{F0: 1, F1: 11}
                                        _ = v106
                                    }
                                }
                                v107 := GenS12{F0: v54}
                                switch fi51 {
                                case 2147483647:
                                    delete(v4, (fi51 << ((480) & 63)))
                                case 128:
                                    fmt.Println(v22, v53, v54)
                                }
                                _ = v107
                            }
                            v5 = 3.14159
                        } else {
                            var v108 int32 = -2147483648
                            switch (fi52 >> ((42) & 63)) {
                            case -9223372036854775808:
                                var v109 int64 = 0
                                _ = v109
                            case 9223372036854775807:
                                if v2 {
                                    if ("hello" <= ("a" + "foo bar")) {
                                        v110 := GenS13{F0: ("" + "\n\t"), F1: 32767, F2: -9223372036854775808}
                                        var v111 int16 = 857
                                        var v112 uint16 = 337
                                        _ = v110
                                        _ = v111
                                        _ = v112
                                    }
                                } else {
                                    var v113 float32 = v6
                                    if (!(!(-1 == v3))) {
                                        v114 := GenS12{F0: v54}
                                        v115 := GenS13{F0: "a", F1: 1, F2: 0}
                                        v2 = v2
                                        _ = v114
                                        _ = v115
                                    } else {
                                        var v116 uint32 = 1
                                        v117 := GenS5{F0: (v2 && true), F1: v54, F2: 1}
                                        _ = v116
                                        _ = v117
                                    }
                                    v118 := GenS3{F0: "\n\t"}
                                    _ = v113
                                    _ = v118
                                }
                            case 127:
                                v119 := GenS7{F0: 4294967295}
                                _ = v119
                            default:
                                var v120 uint64 = uint64(v22)
                                _ = v120
                            }
                            fmt.Println(v6, v53)
                            _ = v108
                        }
                    default:
                        var v121 int = fi51
                        _ = v121
                    }
                    v1 = append(v1, "a")
                    _ = v54
                }
                _ = len(v4)
                v1 = append(v1, "a")
                v122 := GenS14{F0: "hello"}
                switch "foo bar" {
                case "foo bar":
                    v123 := v4[(fi52 & (fi51 - (-(fi51))))]
                    _ = v123
                case "\n\t":
                    switch "a" {
                    case "\n\t":
                        v3 += v3
                        v3 = 42
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v124 uint8 = (v22 + 0)
                        _ = v124
                    case "a":
                        if true {
                            if v2 {
                                if v2 {
                                    delete(v4, -574)
                                }
                            }
                            var v125 uint32 = 605
                            _ = v125
                        }
                    default:
                        v126 := GenS15{F0: (v3 - v3), F1: 65535}
                        _ = v126
                    }
                case "":
                    v127 := GenS8{F0: 9223372036854775807, F1: 1}
                    _ = v127
                default:
                    v53 = 1e10
                }
                _ = fi52
                _ = v122
                _ = v53
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v128 int = 4294967295
            _ = fi51
            _ = v128
        }
        v3 = -895
    } else {
        v129 := GenS16{F0: 255, F1: -610}
        for fi130 := 0; fi130 < 1; fi130++ {
            var v131 uint8 = v22
            _ = fi130
            _ = v131
        }
        _ = v129
    }
    fmt.Println(v6)
    for fi132 := 0; fi132 < 5; fi132++ {
        for fi133 := 0; fi133 < 9; fi133++ {
            v6 -= 0.0
            fmt.Println(v2, v5, v6)
            _ = fi133
        }
        _ = fi132
    }
    v134 := GenS2{F0: 65535, F1: 127}
    for fi135 := 0; fi135 < 11; fi135++ {
        for fi136 := 0; fi136 < 3; fi136++ {
            var v137 int32 = (v3 + (v3 % ((-1) | 1)))
            fmt.Println(v6, v2, fi135)
            var v138 uint8 = v22
            _ = fi136
            _ = v137
            _ = v138
        }
        for fi139 := 0; fi139 < 1; fi139++ {
            v140 := v4[fi135]
            switch fi139 {
            case -197:
                v5 = -2.25
            case 721:
                v4[(fi139 | fi139)] = (v140 << ((255) & 63))
            case 2147483647:
                v6 = v6
            }
            var v141 int64 = 924
            _ = fi139
            _ = v140
            _ = v141
        }
        _ = fi135
    }
    v142 := GenS5{F0: (962 <= 2), F1: 396, F2: -14}
    v143 := GenS17{F0: 1, F1: 0}
    v144 := GenS3{F0: "foo bar"}
    v5 = 3.14159
    v6 -= ((1.5 + (v6 / v6)) * (v6 + 0.0))
    var v145 int32 = (v3 | v3)
    v146 := GenS14{F0: "foo bar"}
    for fi147 := 0; fi147 < 10; fi147++ {
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
        if false {
            v4[fi147] = 2147483647
        } else {
            v22 = (v22 << ((255) & 63))
            _ = len(v4)
            var v148 int = fi147
            _ = v148
        }
        _ = len(v1)
        _ = fi147
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v145:", v145)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v134
    _ = v142
    _ = v143
    _ = v144
    _ = v145
    _ = v146
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
