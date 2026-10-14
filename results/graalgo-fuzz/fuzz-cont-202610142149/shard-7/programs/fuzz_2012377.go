package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 string; F2 uint }
type GenS6 struct { F0 uint16; F1 int16 }
type GenS7 struct { F0 bool; F1 int8; F2 bool }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint32; F1 uint }
type GenS10 struct { F0 int64; F1 int64; F2 uint64 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 string; F1 uint }
type GenS14 struct { F0 int16; F1 bool }
type GenS15 struct { F0 uint32; F1 int8 }
type GenS16 struct { F0 uint8; F1 bool }
type GenS17 struct { F0 uint32; F1 uint32; F2 bool }
type GenS18 struct { F0 uint }
type GenS19 struct { F0 bool; F1 uint64; F2 int }

func main() {
    var v1 uint16 = 11
    var v2 int64 = -1
    var v3 int64 = (v2 ^ (-(v2)))
    var v4 int = 2147483647
    var v5 uint = 227
    v6 := GenS1{F0: 1}
    for fi7 := 0; fi7 < 2; fi7++ {
        v1 -= v1
        v1 -= (^(v1))
        _ = fi7
    }
    v8 := GenS2{F0: v4}
    if ((v5 / (((v5 ^ v5)) | 1)) != 2) {
        v4 -= (v4 - (v4 - 1))
        var v9 uint16 = 2
        _ = v9
    } else {
        v10 := []int{0, v4}
        switch (("" + "a") + "\n\t") {
        case "\n\t":
            _ = len(v10)
        default:
            fmt.Println(v5, v4, v4)
        }
        for fi11 := 0; fi11 < 1; fi11++ {
            if true {
                var v12 float32 = 0.0
                var v13 bool = true
                v3 = v3
                _ = v12
                _ = v13
            } else {
                var v14 int16 = 1000
                v10 = append(v10, (v4 * 1))
                _ = v14
            }
            _ = fi11
        }
        _ = v10
    }
    v15 := GenS1{F0: 54}
    for fi16 := 0; fi16 < 19; fi16++ {
        v17 := []int{v4}
        if ("foo bar" != "hello") {
            v18 := GenS3{F0: 940, F1: ("foo bar" + (("hello" + "a") + "\n\t"))}
            v17 = append(v17, 32767)
            switch "\n\t" {
            case "a":
                if ((^(v1)) != v1) {
                    var v19 float32 = (1.5 - 0.0)
                    v20 := GenS3{F0: 2, F1: ("" + ("a" + "hello"))}
                    _ = v19
                    _ = v20
                } else {
                    if true {
                        v1 += 42
                    } else {
                        v1 -= v1
                    }
                }
            default:
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
            }
            _ = v18
        } else {
            v5 = v5
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
            var v21 bool = (!true)
            _ = v21
        }
        v3 = int64(v5)
        _ = fi16
        _ = v17
    }
    switch "a" {
    case "a":
        v4 = v4
    default:
        for fi22 := 0; fi22 < 16; fi22++ {
            v23 := []int{2147483647}
            for fi24 := 0; fi24 < 8; fi24++ {
                fmt.Println(v5, v3, v4)
                var v25 float64 = 0.0
                if len(v23) > 0 {
                    v23[0] = (fi22 % ((42) | 1))
                }
                _ = fi24
                _ = v25
            }
            v23 = append(v23, fi22)
            _ = fi22
            _ = v23
        }
    }
    v1 = v1
    var v26 uint8 = 0
    v26 -= uint8(v1)
    for fi27 := 0; fi27 < 15; fi27++ {
        var v28 uint32 = 2147483647
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 1; fi29++ {
        for fi30 := 0; fi30 < 12; fi30++ {
            var v31 uint32 = 0
            if (true || ("a" == "")) {
                if false {
                    if ((true || (false && true)) && (1 != (^(v1)))) {
                        v32 := GenS4{F0: v2}
                        var v33 string = "foo bar"
                        var v34 uint64 = 4294967295
                        var v35 int32 = 42
                        var v36 uint32 = v31
                        _ = v32
                        _ = v33
                        _ = v34
                        _ = v35
                        _ = v36
                    } else {
                        var v37 uint8 = uint8(v5)
                        switch fi30 {
                        case 128:
                            switch ((("a" + (("foo bar" + "hello") + ((("a" + ("a" + "a")) + ("foo bar" + "foo bar")) + "a"))) + "hello") + "foo bar") {
                            case "a":
                                var v38 string = ("" + "\n\t")
                                _ = v38
                            case "foo bar":
                                if false {
                                    v5 -= v5
                                }
                            case "\n\t":
                                v5 = 0
                            default:
                                v39 := GenS2{F0: v4}
                                _ = v39
                            }
                        case 32768:
                            v40 := GenS5{F0: -1, F1: "hello", F2: (^(v5))}
                            _ = v40
                        }
                        _ = v37
                    }
                    var v41 int16 = int16(v31)
                    switch fi30 {
                    case 41:
                        if true {
                            if (!((v5 | (v5 * v5)) == (v5 - v5))) {
                                v42 := GenS6{F0: 255, F1: 576}
                                switch ("\n\t" + "hello") {
                                case "a":
                                    v4 -= (fi29 / (((-(fi30))) | 1))
                                case "foo bar":
                                    var v43 bool = (false || (true || false))
                                    _ = v43
                                case "hello":
                                    v44 := []int{(fi30 + v4), v4, (v4 + 655)}
                                    _ = v44
                                }
                                _ = v42
                            }
                            var v45 uint8 = v26
                            _ = v45
                        }
                    case 1:
                        v46 := []int{}
                        _ = v46
                    default:
                        var v47 uint = (v5 | v5)
                        _ = v47
                    }
                    _ = v41
                } else {
                    v48 := []int{}
                    var v49 uint8 = (^(v26))
                    _ = v48
                    _ = v49
                }
            }
            v50 := []int{}
            v4 = fi30
            _ = fi30
            _ = v31
            _ = v50
        }
        for fi51 := 0; fi51 < 6; fi51++ {
            v52 := GenS7{F0: (!(!(4294967295 > 2147483648))), F1: 14, F2: ((-409 != -1) && false)}
            var v53 int = (-(v4))
            switch (-(fi29)) {
            case 4294967295:
                v53 -= (^(v53))
            default:
                v54 := GenS8{F0: 65535}
                _ = v54
            }
            if ((v5 * (v5 >> (((v5 + v5)) & 63))) <= v5) {
                switch (("hello" + "hello") + ("a" + ("\n\t" + ""))) {
                case "\n\t":
                    v55 := []int{(v53 & (fi51 * 32767)), (fi51 + v4), fi29}
                    _ = v55
                case "":
                    v56 := []int{(fi29 / ((fi29) | 1)), int(v3), v53, int(v3)}
                    _ = v56
                }
                if (380 >= 943) {
                    if true {
                        switch "a" {
                        case "\n\t":
                            switch v4 {
                            case 2147483647:
                                var v57 int8 = 114
                                _ = v57
                            case 128:
                                if (((((((float64((((0.1 - (-2.25 / 1.0)) + 0.1) * -2.25)) / (1.5 - 0.0)) * 0.1) * ((3.14159 / ((1e10 / 1e10) + (-2.25 * float64(0.1)))) / 1.0)) - 0.0) - (1.5 / -2.25)) / ((3.14159 * (((0.1 - (1.5 + 1.5)) / 1.0) - 1e10)) - 0.1)) > -2.25) {
                                    if false {
                                        v58 := GenS2{F0: 9223372036854775807}
                                        v59 := []int{(fi51 + v53), (fi29 | v53), v4, v53}
                                        v60 := GenS9{F0: 365, F1: v5}
                                        if true {
                                            switch "" {
                                            case "a":
                                                if false {
                                                    if len(v59) > 0 {
                                                        v59[0] = fi29
                                                    }
                                                    fmt.Println(v1)
                                                }
                                            default:
                                                _ = len(v59)
                                            }
                                            switch 42 {
                                            case 187:
                                                if len(v59) > 0 {
                                                    v59[0] = (fi29 | fi29)
                                                }
                                            case -9223372036854775808:
                                                v4 += fi51
                                            case -594:
                                                var v61 float32 = 0.1
                                                v62 := GenS1{F0: v26}
                                                fmt.Println(v26)
                                                _ = v61
                                                _ = v62
                                            default:
                                                v5 = 0
                                            }
                                            if (!true) {
                                                fmt.Println(v2, v26, v3)
                                                var v63 uint64 = 857
                                                _ = v63
                                            } else {
                                                v3 = (v2 * (^(v2)))
                                                switch "a" {
                                                case "hello":
                                                    v2 -= v3
                                                case "a":
                                                    switch (fi29 | (fi51 & (v4 + 42))) {
                                                    case -9223372036854775808:
                                                        v3 -= -1
                                                    case -1:
                                                        v64 := GenS9{F0: 0, F1: 255}
                                                        _ = v64
                                                    }
                                                }
                                                v59 = append(v59, (v4 - v4))
                                            }
                                        } else {
                                            _ = len(v59)
                                            v65 := GenS9{F0: 42, F1: v5}
                                            switch ("" + (("hello" + (("hello" + "foo bar") + "hello")) + "a")) {
                                            case "hello":
                                                fmt.Println(v53, fi29)
                                            case "foo bar":
                                                fmt.Println(v1, v3)
                                            case "a":
                                                fmt.Println(v4, v4, v5)
                                            }
                                            _ = v65
                                        }
                                        var v66 int32 = -201
                                        _ = v58
                                        _ = v59
                                        _ = v60
                                        _ = v66
                                    }
                                } else {
                                    var v67 int32 = 2147483647
                                    v3 = v2
                                    v26 += v26
                                    _ = v67
                                }
                            }
                        case "a":
                            var v68 int32 = 42
                            _ = v68
                        }
                        v3 = 0
                        if (v26 != (v26 / ((v26) | 1))) {
                            v69 := []int{(v4 * 32767), (fi29 % (((fi51 << (((v5 % (((v5 >> ((v5) & 63))) | 1))) & 63))) | 1))}
                            _ = v69
                        }
                    }
                } else {
                    var v70 uint16 = v1
                    var v71 uint16 = v70
                    v72 := []int{fi29, fi29, v53}
                    _ = v70
                    _ = v71
                    _ = v72
                }
                if true {
                    var v73 string = ("" + "\n\t")
                    var v74 float32 = float32(0.0)
                    v75 := []int{127}
                    v2 = v3
                    _ = v73
                    _ = v74
                    _ = v75
                } else {
                    switch "a" {
                    case "":
                        if (!(v3 < -576)) {
                            if false {
                                v3 = 0
                                v76 := []int{(-(fi51)), v53, (fi29 / ((1) | 1))}
                                var v77 uint64 = 4294967295
                                _ = v76
                                _ = v77
                            } else {
                                var v78 uint32 = 65535
                                if true {
                                    fmt.Println(v3, v4, fi29)
                                    var v79 float64 = 1e10
                                    fmt.Println(v2)
                                    _ = v79
                                } else {
                                    switch "foo bar" {
                                    case "hello":
                                        v80 := GenS8{F0: v1}
                                        _ = v80
                                    case "\n\t":
                                        v81 := GenS8{F0: uint16(v78)}
                                        switch (("" + "a") + "foo bar") {
                                        case "a":
                                            if false {
                                                fmt.Println(v4, v1, v53)
                                                v82 := GenS1{F0: 127}
                                                if true {
                                                    v2 = 10
                                                    v1 = v1
                                                    var v83 int32 = 0
                                                    v83 = v83
                                                    fmt.Println(v5)
                                                    _ = v83
                                                } else {
                                                    if ((v53 - (v4 + (v4 | fi29))) < (v53 << ((793) & 63))) {
                                                        v5 = v5
                                                        v84 := []int{fi29, 32768, -2147483648}
                                                        _ = v84
                                                    } else {
                                                        var v85 int16 = 32767
                                                        var v86 int = int(v85)
                                                        v53 -= fi29
                                                        _ = v85
                                                        _ = v86
                                                    }
                                                    v78 = 2147483647
                                                    v87 := GenS10{F0: v3, F1: 0, F2: uint64(v26)}
                                                    _ = v87
                                                }
                                                _ = v82
                                            }
                                            fmt.Println(v53, v26, v26)
                                        }
                                        v88 := GenS11{F0: (v5 % ((v5) | 1))}
                                        if ((v53 | -304) >= fi29) {
                                            fmt.Println(fi29, v26, v5)
                                            v78 = v78
                                            var v89 float32 = 1e10
                                            _ = v89
                                        }
                                        _ = v81
                                        _ = v88
                                    case "a":
                                        v5 = 203
                                    }
                                    v90 := []int{(v53 ^ (v4 & v53)), fi51}
                                    _ = v90
                                }
                                var v91 uint8 = (^(v26))
                                _ = v78
                                _ = v91
                            }
                        }
                    case "a":
                        if (!true) {
                            if false {
                                switch -90 {
                                case 2147483647:
                                    v26 -= (v26 | v26)
                                }
                            } else {
                                if ((!(false || (-119 < 58))) && true) {
                                    if true {
                                        v26 = (v26 - v26)
                                        v92 := GenS12{F0: 42}
                                        if true {
                                            fmt.Println(v53)
                                            var v93 int8 = 0
                                            fmt.Println(fi51)
                                            _ = v93
                                        } else {
                                            if true {
                                                if false {
                                                    v1 = 1
                                                    if true {
                                                        switch v4 {
                                                        case 128:
                                                            v94 := []int{v4, (v4 ^ v53), fi29}
                                                            _ = v94
                                                        case 269:
                                                            switch (v4 | (v53 >> ((v5) & 63))) {
                                                            case 1:
                                                                if true {
                                                                    v53 += 0
                                                                }
                                                            case 32768:
                                                                if (((141 <= 42) && (!((false && true) || false))) && (!(false && true))) {
                                                                    if (false && true) {
                                                                        var v95 int64 = (v2 + v2)
                                                                        var v96 float64 = 0.0
                                                                        _ = v95
                                                                        _ = v96
                                                                    } else {
                                                                        fmt.Println(v4)
                                                                        v97 := []int{}
                                                                        _ = v97
                                                                    }
                                                                }
                                                            case -901:
                                                                if (42 > 1) {
                                                                    if ((fi51 % ((fi29) | 1)) < (v53 % ((493) | 1))) {
                                                                        var v98 string = ""
                                                                        _ = v98
                                                                    } else {
                                                                        v99 := GenS13{F0: "a", F1: (v5 & (v5 - (v5 * v5)))}
                                                                        v1 -= v1
                                                                        _ = v99
                                                                    }
                                                                    v100 := []int{fi51, (v53 | v53)}
                                                                    _ = v100
                                                                } else {
                                                                    var v101 uint = uint(v26)
                                                                    _ = v101
                                                                }
                                                            }
                                                        case -699:
                                                            v102 := []int{(fi29 % (((-(v53))) | 1)), 32768}
                                                            _ = v102
                                                        default:
                                                            v2 = v3
                                                        }
                                                    }
                                                    var v103 int = v4
                                                    v104 := []int{v4, 0, fi29}
                                                    _ = v103
                                                    _ = v104
                                                }
                                            }
                                            switch "a" {
                                            case "\n\t":
                                                if true {
                                                    var v105 uint16 = (v1 * v1)
                                                    if (((v105 / ((873) | 1)) > v105) && true) {
                                                        v5 -= 0
                                                        var v106 int64 = v2
                                                        if true {
                                                            var v107 uint32 = 314
                                                            v108 := []int{(fi29 << ((v5) & 63))}
                                                            v109 := GenS2{F0: fi29}
                                                            _ = v107
                                                            _ = v108
                                                            _ = v109
                                                        } else {
                                                            fmt.Println(fi29)
                                                        }
                                                        _ = v106
                                                    } else {
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            v1 = (^(v1))
                                                        case "hello":
                                                            v3 += (v3 - 751)
                                                        case "foo bar":
                                                            v2 += (v3 + -9223372036854775808)
                                                        default:
                                                            if ((true && (!true)) && false) {
                                                                switch int(v1) {
                                                                case 894:
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        if true {
                                                                            v2 = 0
                                                                        }
                                                                    default:
                                                                        v110 := []int{}
                                                                        _ = v110
                                                                    }
                                                                case 42:
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        v111 := []int{-421, 858, (v4 ^ v53), (v53 | (-(fi51)))}
                                                                        _ = v111
                                                                    case "a":
                                                                        switch 650 {
                                                                        case 42:
                                                                            fmt.Println(v4)
                                                                        default:
                                                                            fmt.Println(v3, v4, v1)
                                                                        }
                                                                    }
                                                                default:
                                                                    var v112 int = -1
                                                                    _ = v112
                                                                }
                                                            } else {
                                                                if false {
                                                                    v113 := []int{v53}
                                                                    _ = v113
                                                                }
                                                                var v114 float64 = 0.0
                                                                switch "hello" {
                                                                case "\n\t":
                                                                    v5 -= (v5 >> (((v5 + v5)) & 63))
                                                                }
                                                                _ = v114
                                                            }
                                                        }
                                                        switch fi29 {
                                                        case 128:
                                                            v115 := GenS14{F0: -1, F1: false}
                                                            _ = v115
                                                        case 9223372036854775807:
                                                            var v116 int32 = -925
                                                            _ = v116
                                                        case 821:
                                                            v117 := []int{-40, v53, v4}
                                                            _ = v117
                                                        default:
                                                            v118 := GenS6{F0: v105, F1: 330}
                                                            _ = v118
                                                        }
                                                        v4 += (v4 | (v4 - int(v5)))
                                                    }
                                                    _ = v105
                                                }
                                            case "a":
                                                v119 := []int{}
                                                _ = v119
                                            }
                                            v120 := []int{v53, (v53 % ((fi51) | 1)), 32768}
                                            _ = v120
                                        }
                                        var v121 int8 = -1
                                        fmt.Println(v121, v2)
                                        _ = v121
                                        _ = v92
                                    }
                                    switch "" {
                                    case "a":
                                        var v122 int16 = -416
                                        _ = v122
                                    }
                                }
                                fmt.Println(v1)
                                if false {
                                    v123 := GenS10{F0: 9223372036854775807, F1: v3, F2: 4294967295}
                                    var v124 int16 = -561
                                    fmt.Println(v2)
                                    v1 = (v1 / ((587) | 1))
                                    if true {
                                        var v125 int8 = 42
                                        _ = v125
                                    }
                                    _ = v123
                                    _ = v124
                                } else {
                                    v2 = v3
                                    switch fi51 {
                                    case 42:
                                        var v126 uint64 = 93
                                        _ = v126
                                    case 127:
                                        switch ("hello" + "foo bar") {
                                        case "":
                                            var v127 float64 = 0.0
                                            _ = v127
                                        case "a":
                                            v128 := GenS15{F0: 42, F1: int8(v3)}
                                            v129 := GenS16{F0: v26, F1: (2 != 2)}
                                            fmt.Println(v2, v26, fi51)
                                            var v130 bool = false
                                            _ = v128
                                            _ = v129
                                            _ = v130
                                        case "hello":
                                            switch "\n\t" {
                                            case "hello":
                                                v131 := GenS13{F0: "", F1: 221}
                                                _ = v131
                                            default:
                                                v132 := GenS9{F0: 65535, F1: 51}
                                                _ = v132
                                            }
                                        }
                                    case 2147483647:
                                        v133 := []int{128, 32767, fi51}
                                        _ = v133
                                    }
                                    v134 := []int{}
                                    _ = v134
                                }
                            }
                        }
                    case "\n\t":
                        var v135 bool = true
                        _ = v135
                    }
                    v136 := GenS5{F0: 1, F1: "a", F2: v5}
                    _ = v136
                }
            } else {
                v137 := []int{}
                switch "\n\t" {
                case "hello":
                    if len(v137) > 0 {
                        v137[0] = (v53 << (((v5 & v5)) & 63))
                    }
                default:
                    var v138 int16 = 768
                    _ = v138
                }
                v139 := GenS2{F0: 42}
                _ = len(v137)
                _ = v137
                _ = v139
            }
            if false {
                var v140 string = "hello"
                switch v4 {
                case 1:
                    v141 := GenS17{F0: 2147483647, F1: 684, F2: false}
                    _ = v141
                case -2147483648:
                    var v142 uint = (v5 + v5)
                    _ = v142
                case 87:
                    if false {
                        v143 := GenS18{F0: v5}
                        _ = v143
                    } else {
                        v4 += (-(fi51))
                        switch v140 {
                        case "\n\t":
                            var v144 uint = (v5 << (((v5 * v5)) & 63))
                            _ = v144
                        case "foo bar":
                            v140 = (v140 + "\n\t")
                        default:
                            fmt.Println(v4)
                        }
                    }
                }
                _ = v140
            }
            v2 = 0
            _ = fi51
            _ = v52
            _ = v53
        }
        v145 := GenS19{F0: true, F1: 2, F2: 42}
        _ = fi29
        _ = v145
    }
    for fi146 := 0; fi146 < 19; fi146++ {
        switch "" {
        case "":
            for fi147 := 0; fi147 < 3; fi147++ {
                fmt.Println(v26, v4)
                fmt.Println(v3, fi146, v1)
                _ = fi147
            }
        case "a":
            var v148 uint8 = v26
            _ = v148
        case "foo bar":
            var v149 int8 = 1
            _ = v149
        }
        v3 = (v3 | (-(v3)))
        _ = fi146
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
