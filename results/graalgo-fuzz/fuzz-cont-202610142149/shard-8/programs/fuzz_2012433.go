package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 uint8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int8; F1 int16 }
type GenS9 struct { F0 int64; F1 int32 }
type GenS10 struct { F0 int32; F1 uint32 }
type GenS11 struct { F0 uint16; F1 uint16 }
type GenS12 struct { F0 uint32; F1 uint32 }
type GenS13 struct { F0 string; F1 uint64 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 int64; F1 int16; F2 int8 }
type GenS16 struct { F0 bool; F1 uint16 }
type GenS17 struct { F0 bool; F1 int16 }
type GenS18 struct { F0 int64 }
type GenS19 struct { F0 int64; F1 int64 }
type GenS20 struct { F0 uint16 }
type GenS21 struct { F0 int8 }
type GenS22 struct { F0 int16 }

func main() {
    var v1 int16 = 689
    var v2 int8 = 0
    v3 := []uint32{745, 65535, 375}
    var v4 bool = true
    var v5 uint64 = 0
    v6 := []int{2147483647, -31, 128}
    for fi7 := 0; fi7 < 4; fi7++ {
        v4 = (true || (65535 == 42))
        v8 := GenS1{F0: true, F1: "", F2: 52}
        _ = fi7
        _ = v8
    }
    if true {
        var v9 uint16 = 2
        switch "" {
        case "\n\t":
            var v10 int8 = (v2 | (v2 >> ((255) & 63)))
            _ = v10
        case "a":
            if false {
                for fi11 := 0; fi11 < 15; fi11++ {
                    if ((v2 & v2) >= (v2 >> ((65535) & 63))) {
                        var v12 int8 = 110
                        v12 = v12
                        _ = v12
                    }
                    _ = fi11
                }
                if v4 {
                    for fi13 := 0; fi13 < 13; fi13++ {
                        var v14 int = fi13
                        v14 = 42
                        var v15 int16 = (v1 + v1)
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    for fi16 := 0; fi16 < 8; fi16++ {
                        var v17 uint32 = 42
                        switch 128 {
                        case -536:
                            var v18 int = (fi16 % ((32767) | 1))
                            _ = v18
                        case 279:
                            if ((fi16 + 9223372036854775807) < -1) {
                                if v4 {
                                    fmt.Println(v4, v17)
                                    switch ("\n\t" + ("\n\t" + "a")) {
                                    case "a":
                                        var v19 uint64 = 0
                                        _ = v19
                                    case "\n\t":
                                        var v20 string = "foo bar"
                                        _ = v20
                                    case "":
                                        fmt.Println(v9, v5)
                                    }
                                }
                                fmt.Println(v17, v5, v4)
                                var v21 uint16 = 255
                                _ = v21
                            }
                        case -97:
                            if (!(2 == 1)) {
                                v22 := GenS2{F0: 515}
                                for fi23 := 0; fi23 < 18; fi23++ {
                                    _ = len(v3)
                                    _ = fi23
                                }
                                v5 -= 203
                                _ = v22
                            } else {
                                for fi24 := 0; fi24 < 6; fi24++ {
                                    var v25 int32 = 42
                                    _ = fi24
                                    _ = v25
                                }
                            }
                        default:
                            var v26 float32 = -2.25
                            _ = v26
                        }
                        _ = fi16
                        _ = v17
                    }
                    var v27 int8 = (v2 % ((-128) | 1))
                    _ = v27
                }
                switch 128 {
                case -162:
                    if ("" < "foo bar") {
                        if true {
                            if (("foo bar" == "a") || (!false)) {
                                if len(v6) > 0 {
                                    v6[0] = 0
                                }
                                if (false || false) {
                                    v2 = v2
                                }
                            } else {
                                var v28 uint = 129
                                v29 := GenS3{F0: 241, F1: (v28 << ((v28) & 63))}
                                _ = v28
                                _ = v29
                            }
                            switch 9223372036854775807 {
                            case 0:
                                var v30 int16 = v1
                                _ = v30
                            case 722:
                                if v4 {
                                    fmt.Println(v4, v1, v4)
                                    for fi31 := 0; fi31 < 8; fi31++ {
                                        var v32 uint8 = 51
                                        var v33 int64 = -597
                                        _ = fi31
                                        _ = v32
                                        _ = v33
                                    }
                                    for fi34 := 0; fi34 < 20; fi34++ {
                                        var v35 int8 = 42
                                        if v4 {
                                            var v36 bool = (v4 || v4)
                                            for fi37 := 0; fi37 < 13; fi37++ {
                                                v2 = (v35 << ((105) & 63))
                                                v2 -= (v35 + 0)
                                                fmt.Println(v36)
                                                _ = fi37
                                            }
                                            v38 := GenS4{F0: ("a" + "\n\t"), F1: -1}
                                            _ = v36
                                            _ = v38
                                        }
                                        fmt.Println(v35, v1)
                                        _ = fi34
                                        _ = v35
                                    }
                                } else {
                                    var v39 int16 = 438
                                    _ = v39
                                }
                            default:
                                _ = len(v6)
                            }
                        }
                    }
                case 806:
                    switch -9223372036854775808 {
                    case 42:
                        fmt.Println(v4, v5)
                    case 2147483647:
                        if (true || v4) {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        } else {
                            for fi40 := 0; fi40 < 1; fi40++ {
                                var v41 uint8 = 27
                                _ = fi40
                                _ = v41
                            }
                            for fi42 := 0; fi42 < 20; fi42++ {
                                v9 = v9
                                _ = fi42
                            }
                            switch "" {
                            case "\n\t":
                                fmt.Println(v5, v4, v1)
                            case "foo bar":
                                if false {
                                    if v4 {
                                        v6 = append(v6, -811)
                                    }
                                }
                            case "":
                                for fi43 := 0; fi43 < 7; fi43++ {
                                    var v44 uint64 = 493
                                    v44 = v44
                                    _ = fi43
                                    _ = v44
                                }
                            }
                        }
                    }
                }
            }
        case "":
            for fi45 := 0; fi45 < 5; fi45++ {
                if true {
                    for fi46 := 0; fi46 < 8; fi46++ {
                        if v4 {
                            v47 := GenS1{F0: (v1 <= int16(v5)), F1: "foo bar", F2: v5}
                            if ((v1 - (^(v1))) == v1) {
                                switch (("a" + ("" + "a")) + "hello") {
                                case "a":
                                    if v4 {
                                        if (268 < 2) {
                                            switch "" {
                                            case "\n\t":
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            case "":
                                                if len(v3) > 0 {
                                                    v3[0] = 2147483648
                                                }
                                            case "hello":
                                                v4 = true
                                            }
                                        } else {
                                            v9 = (v9 >> ((1) & 63))
                                            var v48 int32 = 2147483647
                                            var v49 uint = 65535
                                            _ = v48
                                            _ = v49
                                        }
                                    }
                                default:
                                    v6 = append(v6, (fi46 >> ((18446744073709551615) & 63)))
                                }
                                v4 = (uint8(v5) != 0)
                                switch "\n\t" {
                                case "":
                                    switch fi45 {
                                    case 9223372036854775807:
                                        switch (fi46 * 42) {
                                        case -292:
                                            switch (fi46 / ((2147483647) | 1)) {
                                            case -9223372036854775808:
                                                v50 := GenS3{F0: 1, F1: 4294967295}
                                                if len(v3) > 0 {
                                                    v3[0] = 4294967295
                                                }
                                                var v51 uint16 = (v9 & (^(v9)))
                                                _ = v50
                                                _ = v51
                                            }
                                        case 42:
                                            v9 = v9
                                        default:
                                            v5 = (v5 << ((18446744073709551615) & 63))
                                        }
                                    case 4294967295:
                                        fmt.Println(v4)
                                    case 128:
                                        switch int(v1) {
                                        case -975:
                                            v5 += v5
                                        case -775:
                                            if (true && v4) {
                                                switch "foo bar" {
                                                case "foo bar":
                                                    var v52 float32 = 3.14159
                                                    _ = v52
                                                case "a":
                                                    var v53 int16 = v1
                                                    _ = v53
                                                case "":
                                                    fmt.Println(fi46, v1)
                                                }
                                                var v54 int8 = (v2 & v2)
                                                _ = v54
                                            }
                                            v4 = true
                                            v5 = v5
                                        case 2147483647:
                                            if false {
                                                if true {
                                                    var v55 int32 = int32(v2)
                                                    if v4 {
                                                        v4 = v4
                                                        _ = len(v6)
                                                        var v56 float32 = float32(float64(3.14159))
                                                        _ = v56
                                                    } else {
                                                        var v57 uint = 4294967295
                                                        var v58 uint16 = 942
                                                        if (1e10 >= 0.1) {
                                                            v59 := GenS5{F0: true}
                                                            if false {
                                                                if false {
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 65535
                                                                    }
                                                                    var v60 int = (^(fi45))
                                                                    _ = v60
                                                                } else {
                                                                    if (128 != 248) {
                                                                        v1 += (v1 * (v1 / ((v1) | 1)))
                                                                        v61 := GenS4{F0: ("hello" + "hello"), F1: v55}
                                                                        _ = v61
                                                                    }
                                                                }
                                                                fmt.Println(v58)
                                                                if (!(!true)) {
                                                                    switch ("a" + "a") {
                                                                    case "":
                                                                        var v62 uint8 = 186
                                                                        _ = v62
                                                                    case "hello":
                                                                        var v63 int64 = -9223372036854775808
                                                                        _ = v63
                                                                    case "foo bar":
                                                                        fmt.Println(v4)
                                                                    }
                                                                } else {
                                                                    fmt.Println(v55, v4)
                                                                }
                                                            }
                                                            var v64 uint = (^(v57))
                                                            _ = v59
                                                            _ = v64
                                                        } else {
                                                            var v65 int64 = 1
                                                            _ = v65
                                                        }
                                                        if true {
                                                            if (v55 <= 42) {
                                                                v2 += v2
                                                                fmt.Println(v9)
                                                                fmt.Println(v55)
                                                            }
                                                            v66 := GenS4{F0: ("foo bar" + ""), F1: v55}
                                                            _ = v66
                                                        }
                                                        _ = v57
                                                        _ = v58
                                                    }
                                                    _ = v55
                                                } else {
                                                    v67 := GenS4{F0: "hello", F1: -1}
                                                    v3 = append(v3, 0)
                                                    _ = v67
                                                }
                                                var v68 string = "foo bar"
                                                fmt.Println(v2, v68, v2)
                                                var v69 uint = 1
                                                _ = v68
                                                _ = v69
                                            } else {
                                                if ((^(v9)) < 32767) {
                                                    fmt.Println(v9)
                                                } else {
                                                    _ = len(v6)
                                                    if v4 {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    }
                                                    v3 = append(v3, 2147483647)
                                                }
                                            }
                                        }
                                        switch "\n\t" {
                                        case "":
                                            v70 := GenS3{F0: 2, F1: 940}
                                            _ = v70
                                        case "hello":
                                            v9 -= (v9 ^ (v9 + v9))
                                        }
                                        var v71 uint16 = 66
                                        _ = v71
                                    }
                                }
                                switch fi45 {
                                case -9223372036854775808:
                                    v1 += v1
                                case 0:
                                    fmt.Println(v5, v1)
                                    v2 -= v2
                                    switch "\n\t" {
                                    case "hello":
                                        if len(v3) > 0 {
                                            v3[0] = 2
                                        }
                                    }
                                }
                                if true {
                                    var v72 uint = 1
                                    v9 += (v9 << ((v72) & 63))
                                    if len(v6) > 0 {
                                        v6[0] = 128
                                    }
                                    _ = v72
                                }
                                v73 := GenS6{F0: ("foo bar" + "a"), F1: 127}
                                _ = v73
                            } else {
                                fmt.Println(fi46)
                                v3 = append(v3, 4294967295)
                                v5 = (v5 | (v5 ^ v5))
                                fmt.Println(v1)
                            }
                            var v74 int16 = v1
                            _ = v47
                            _ = v74
                        }
                        fmt.Println(v2)
                        _ = fi46
                    }
                }
                fmt.Println(v5)
                var v75 uint16 = (v9 * (v9 * v9))
                _ = fi45
                _ = v75
            }
        default:
            v1 -= (^(v1))
        }
        v76 := GenS3{F0: uint8(v9), F1: 0}
        _ = v76
        _ = v9
    } else {
        switch ("hello" + (("" + "hello") + "a")) {
        case "hello":
            var v77 int = -811
            _ = v77
        }
        _ = len(v6)
        for fi78 := 0; fi78 < 6; fi78++ {
            var v79 uint32 = 165
            v1 += v1
            var v80 int16 = (^(v1))
            _ = fi78
            _ = v79
            _ = v80
        }
    }
    if v4 {
        switch 32767 {
        case 32767:
            for fi81 := 0; fi81 < 3; fi81++ {
                if v4 {
                    var v82 int32 = -2147483648
                    _ = v82
                }
                for fi83 := 0; fi83 < 1; fi83++ {
                    var v84 float64 = 0.1
                    _ = fi83
                    _ = v84
                }
                _ = fi81
            }
        default:
            switch "" {
            case "hello":
                switch "foo bar" {
                case "foo bar":
                    var v85 uint16 = 32767
                    _ = v85
                default:
                    var v86 uint16 = 32767
                    _ = v86
                }
            case "\n\t":
                _ = len(v3)
            case "a":
                v87 := GenS1{F0: (!true), F1: ((("foo bar" + "\n\t") + "") + (("a" + "a") + "")), F2: v5}
                _ = v87
            }
        }
        v5 = v5
    } else {
        fmt.Println(v1)
    }
    if (!true) {
        var v88 int8 = (v2 / (((v2 & v2)) | 1))
        if v4 {
            var v89 uint16 = 32767
            fmt.Println(v5, v89, v4)
            _ = v89
        } else {
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
            for fi90 := 0; fi90 < 15; fi90++ {
                for fi91 := 0; fi91 < 14; fi91++ {
                    if (v4 || ("" < "\n\t")) {
                        if (!v4) {
                            v3 = append(v3, 859)
                            fmt.Println(fi90, v4)
                            var v92 int8 = (v2 << ((0) & 63))
                            _ = v92
                        }
                    }
                    if v4 {
                        v88 = v88
                        var v93 int8 = int8(fi91)
                        fmt.Println(v1)
                        _ = v93
                    } else {
                        var v94 int32 = 42
                        switch "foo bar" {
                        case "\n\t":
                            v88 = v2
                        case "":
                            var v95 int16 = v1
                            _ = v95
                        case "a":
                            var v96 int64 = -1
                            _ = v96
                        }
                        _ = v94
                    }
                    _ = fi91
                }
                var v97 uint16 = 794
                _ = fi90
                _ = v97
            }
            var v98 float32 = 1.5
            _ = v98
        }
        _ = v88
    }
    v1 = v1
    fmt.Println(v5, v2)
    var v99 float64 = 0.0
    fmt.Println(v5, v1)
    if len(v3) > 0 {
        v3[0] = 2147483648
    }
    v4 = v4
    if v4 {
        if v4 {
            if v4 {
                v4 = v4
            } else {
                for fi100 := 0; fi100 < 3; fi100++ {
                    v3 = append(v3, 42)
                    var v101 uint = 18446744073709551615
                    var v102 uint16 = 65535
                    _ = fi100
                    _ = v101
                    _ = v102
                }
                switch -1 {
                case 127:
                    for fi103 := 0; fi103 < 5; fi103++ {
                        var v104 int = fi103
                        for fi105 := 0; fi105 < 17; fi105++ {
                            fmt.Println(v104)
                            switch "\n\t" {
                            case "a":
                                fmt.Println(v5, v104, v4)
                            case "hello":
                                v106 := GenS7{F0: 1}
                                switch 629 {
                                case 327:
                                    fmt.Println(v2, v99, v5)
                                case 9223372036854775807:
                                    v107 := GenS4{F0: "foo bar", F1: 0}
                                    _ = v107
                                case 32768:
                                    var v108 float32 = (float32(v99) + float32(1e10))
                                    _ = v108
                                default:
                                    if v4 {
                                        v109 := GenS8{F0: -57, F1: v1}
                                        v5 = 42
                                        fmt.Println(v99, v104, v99)
                                        var v110 int8 = 127
                                        if (!v4) {
                                            var v111 int64 = -1
                                            v112 := GenS6{F0: "foo bar", F1: 1}
                                            var v113 int64 = v111
                                            _ = v111
                                            _ = v112
                                            _ = v113
                                        } else {
                                            var v114 float32 = 1e10
                                            if ((((1.5 - (1.5 + v114)) * 1.5) - float32(-2.25)) != v114) {
                                                v115 := GenS6{F0: ("" + "\n\t"), F1: 128}
                                                fmt.Println(v5, v4, v104)
                                                v116 := GenS8{F0: v110, F1: (v1 / ((v1) | 1))}
                                                v117 := GenS9{F0: -577, F1: 1}
                                                v118 := GenS10{F0: -667, F1: 4294967295}
                                                v119 := GenS7{F0: 42}
                                                var v120 bool = (708 <= 42)
                                                if len(v6) > 0 {
                                                    v6[0] = 1
                                                }
                                                _ = v115
                                                _ = v116
                                                _ = v117
                                                _ = v118
                                                _ = v119
                                                _ = v120
                                            }
                                            v2 += v2
                                            _ = v114
                                        }
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        _ = v109
                                        _ = v110
                                    } else {
                                        v121 := GenS4{F0: "hello", F1: 2147483647}
                                        if false {
                                            var v122 float32 = -2.25
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            _ = v122
                                        }
                                        fmt.Println(fi103)
                                        _ = v121
                                    }
                                    var v123 float32 = (3.14159 + float32(-2.25))
                                    _ = v123
                                }
                                var v124 int64 = int64(v2)
                                _ = v106
                                _ = v124
                            default:
                                v1 = v1
                            }
                            if (v2 == (v2 ^ (v2 - int8(v1)))) {
                                if (v4 && ("a" != "")) {
                                    v125 := GenS11{F0: 0, F1: 1}
                                    v126 := GenS12{F0: 1, F1: 2147483648}
                                    _ = v125
                                    _ = v126
                                } else {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                var v127 int = (fi105 << ((uint(v5)) & 63))
                                v128 := GenS13{F0: "a", F1: v5}
                                _ = v127
                                _ = v128
                            }
                            _ = fi105
                        }
                        var v129 uint64 = 4294967295
                        _ = fi103
                        _ = v104
                        _ = v129
                    }
                case 1:
                    v1 -= 950
                default:
                    v130 := GenS14{F0: 990}
                    _ = v130
                }
                v131 := GenS9{F0: -1, F1: int32(v1)}
                _ = v131
            }
            var v132 int32 = int32(v1)
            _ = v132
        }
    }
    if (v4 || true) {
        switch 127 {
        case 210:
            if v4 {
                if v4 {
                    switch -9223372036854775808 {
                    case -374:
                        v2 = (v2 / ((42) | 1))
                    case -2147483648:
                        for fi133 := 0; fi133 < 5; fi133++ {
                            var v134 int = (fi133 + (fi133 * 897))
                            var v135 int = (fi133 | fi133)
                            _ = fi133
                            _ = v134
                            _ = v135
                        }
                    }
                    if len(v6) > 0 {
                        v6[0] = int(v1)
                    }
                } else {
                    var v136 string = ("" + "")
                    v137 := GenS13{F0: v136, F1: v5}
                    for fi138 := 0; fi138 < 12; fi138++ {
                        switch 128 {
                        case 321:
                            for fi139 := 0; fi139 < 19; fi139++ {
                                v6 = append(v6, 2147483647)
                                v4 = false
                                if v4 {
                                    v140 := GenS9{F0: -776, F1: -797}
                                    _ = v140
                                }
                                if ((!(((2147483647 >= 1) || v4) && (655 <= 985))) && (v4 || false)) {
                                    if len(v3) > 0 {
                                        v3[0] = 618
                                    }
                                }
                                v2 -= v2
                                v2 = (v2 ^ v2)
                                _ = fi139
                            }
                        default:
                            switch (-(fi138)) {
                            case -304:
                                var v141 float64 = float64(1e10)
                                _ = v141
                            case -2147483648:
                                switch v136 {
                                case "foo bar":
                                    v142 := GenS9{F0: 9223372036854775807, F1: 130}
                                    _ = v142
                                case "hello":
                                    switch (^(fi138)) {
                                    case -9223372036854775808:
                                        var v143 int32 = -2147483648
                                        _ = v143
                                    case 4294967295:
                                        switch "\n\t" {
                                        case "":
                                            fmt.Println(v2, v5)
                                        case "hello":
                                            v144 := GenS15{F0: 384, F1: -32768, F2: -107}
                                            _ = v144
                                        default:
                                            switch fi138 {
                                            case 0:
                                                v1 += v1
                                            case 32768:
                                                if true {
                                                    for fi145 := 0; fi145 < 20; fi145++ {
                                                        v6 = append(v6, (^(fi145)))
                                                        fmt.Println(v1, v1)
                                                        _ = fi145
                                                    }
                                                } else {
                                                    v146 := GenS6{F0: "foo bar", F1: 204}
                                                    if false {
                                                        for fi147 := 0; fi147 < 14; fi147++ {
                                                            _ = len(v3)
                                                            v148 := GenS8{F0: v2, F1: 32767}
                                                            _ = len(v3)
                                                            _ = fi147
                                                            _ = v148
                                                        }
                                                        v4 = false
                                                    }
                                                    if true {
                                                        if false {
                                                            switch ((((v136 + v136) + v136) + (v136 + v136)) + "") {
                                                            case "a":
                                                                if (124 > 9223372036854775807) {
                                                                    if len(v6) >= 2 {
                                                                        v6 = v6[1:len(v6)]
                                                                    }
                                                                    fmt.Println(v2, v4, v2)
                                                                } else {
                                                                    if true {
                                                                        if (!v4) {
                                                                            var v149 int8 = (v2 - (v2 + v2))
                                                                            var v150 uint16 = 1
                                                                            if len(v6) >= 2 {
                                                                                v6 = v6[1:len(v6)]
                                                                            }
                                                                            _ = v149
                                                                            _ = v150
                                                                        }
                                                                        v4 = v4
                                                                        fmt.Println(v5, v5)
                                                                    } else {
                                                                        var v151 uint64 = (^(v5))
                                                                        _ = v151
                                                                    }
                                                                }
                                                            case "foo bar":
                                                                var v152 int8 = v2
                                                                _ = v152
                                                            }
                                                        } else {
                                                            if (!(63 != 77)) {
                                                                if (!v4) {
                                                                    var v153 uint16 = 32767
                                                                    _ = v153
                                                                } else {
                                                                    var v154 float32 = 1.5
                                                                    var v155 int64 = 1
                                                                    _ = v154
                                                                    _ = v155
                                                                }
                                                                var v156 bool = true
                                                                _ = v156
                                                            } else {
                                                                fmt.Println(v5)
                                                                var v157 int32 = 840
                                                                _ = v157
                                                            }
                                                            var v158 float64 = (float64(3.14159) - (v99 / 3.14159))
                                                            switch "hello" {
                                                            case "hello":
                                                                switch "hello" {
                                                                case "a":
                                                                    v4 = false
                                                                default:
                                                                    fmt.Println(v4, fi138, v4)
                                                                }
                                                            default:
                                                                for fi159 := 0; fi159 < 15; fi159++ {
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 709
                                                                    }
                                                                    v1 += -450
                                                                    fmt.Println(v1)
                                                                    _ = fi159
                                                                }
                                                            }
                                                            _ = v158
                                                        }
                                                        v6 = append(v6, -2147483648)
                                                        v136 = "\n\t"
                                                    } else {
                                                        for fi160 := 0; fi160 < 13; fi160++ {
                                                            if ((v4 && (!(0 >= -9223372036854775808))) || (v2 > 100)) {
                                                                var v161 uint8 = 0
                                                                var v162 uint16 = 26
                                                                _ = v161
                                                                _ = v162
                                                            } else {
                                                                if (float64(0.0) == v99) {
                                                                    var v163 int16 = (v1 - (v1 / (((v1 << ((0) & 63))) | 1)))
                                                                    if false {
                                                                        if v4 {
                                                                            v4 = (0 > 2147483648)
                                                                        } else {
                                                                            var v164 uint64 = 18446744073709551615
                                                                            var v165 float32 = (0.1 - 3.14159)
                                                                            v163 = (v1 + (-(v163)))
                                                                            _ = v164
                                                                            _ = v165
                                                                        }
                                                                        if (-1 > -1) {
                                                                            switch v136 {
                                                                            case "a":
                                                                                var v166 int32 = 0
                                                                                _ = v166
                                                                            case "":
                                                                                if ((!((42 < 740) && false)) && (-9223372036854775808 == -9223372036854775808)) {
                                                                                    v4 = v4
                                                                                    var v167 bool = true
                                                                                    _ = v167
                                                                                } else {
                                                                                    switch "\n\t" {
                                                                                    case "\n\t":
                                                                                        if v4 {
                                                                                            v5 -= v5
                                                                                            var v168 uint8 = 128
                                                                                            _ = v168
                                                                                        } else {
                                                                                            if false {
                                                                                                if v4 {
                                                                                                    if len(v6) >= 2 {
                                                                                                        v6 = v6[1:len(v6)]
                                                                                                    }
                                                                                                    if len(v3) >= 2 {
                                                                                                        v3 = v3[1:len(v3)]
                                                                                                    }
                                                                                                } else {
                                                                                                    v169 := GenS11{F0: 46, F1: 255}
                                                                                                    if (32767 == fi138) {
                                                                                                        v99 = v99
                                                                                                    } else {
                                                                                                        v4 = true
                                                                                                        var v170 uint8 = 224
                                                                                                        var v171 int16 = v163
                                                                                                        _ = v170
                                                                                                        _ = v171
                                                                                                    }
                                                                                                    _ = v169
                                                                                                }
                                                                                                fmt.Println(v163, fi160)
                                                                                            }
                                                                                            v1 = -32768
                                                                                            v2 += v2
                                                                                        }
                                                                                    case "a":
                                                                                        v172 := GenS13{F0: "a", F1: v5}
                                                                                        _ = v172
                                                                                    case "hello":
                                                                                        var v173 uint64 = v5
                                                                                        _ = v173
                                                                                    }
                                                                                    v99 = (1.5 - v99)
                                                                                    if v4 {
                                                                                        v5 += 511
                                                                                        fmt.Println(v4, fi160)
                                                                                        if (false || (0 == 1)) {
                                                                                            if (true || ((fi160 >> ((1) & 63)) == fi138)) {
                                                                                                v136 = ("" + ("a" + v136))
                                                                                            }
                                                                                            if false {
                                                                                                var v174 uint8 = 1
                                                                                                _ = v174
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        var v175 string = v136
                                                                                        v1 = -32768
                                                                                        switch 0 {
                                                                                        case 532:
                                                                                            var v176 uint64 = (v5 | 42)
                                                                                            _ = v176
                                                                                        case 32768:
                                                                                            if len(v6) >= 2 {
                                                                                                v6 = v6[1:len(v6)]
                                                                                            }
                                                                                        default:
                                                                                            var v177 uint32 = 511
                                                                                            _ = v177
                                                                                        }
                                                                                        _ = v175
                                                                                    }
                                                                                }
                                                                            case "\n\t":
                                                                                switch ((v136 + v136) + (((v136 + (v136 + v136)) + v136) + v136)) {
                                                                                case "":
                                                                                    _ = len(v6)
                                                                                case "foo bar":
                                                                                    switch v136 {
                                                                                    case "a":
                                                                                        if true {
                                                                                            var v178 bool = v4
                                                                                            fmt.Println(v4, fi160, v5)
                                                                                            var v179 uint8 = 255
                                                                                            _ = v178
                                                                                            _ = v179
                                                                                        }
                                                                                    default:
                                                                                        v5 -= 4294967295
                                                                                    }
                                                                                case "hello":
                                                                                    var v180 uint8 = 0
                                                                                    _ = v180
                                                                                }
                                                                            }
                                                                            if v4 {
                                                                                v181 := GenS4{F0: v136, F1: -2147483648}
                                                                                _ = v181
                                                                            }
                                                                            if false {
                                                                                v99 += v99
                                                                            } else {
                                                                                var v182 uint8 = 2
                                                                                fmt.Println(fi160, v4)
                                                                                v183 := GenS16{F0: (!v4), F1: 42}
                                                                                v184 := GenS13{F0: "\n\t", F1: v5}
                                                                                _ = v182
                                                                                _ = v183
                                                                                _ = v184
                                                                            }
                                                                        }
                                                                        v185 := GenS6{F0: "hello", F1: 0}
                                                                        _ = v185
                                                                    }
                                                                    if false {
                                                                        fmt.Println(fi138, v99)
                                                                        var v186 uint64 = 42
                                                                        v4 = (false || (!v4))
                                                                        _ = v186
                                                                    }
                                                                    _ = v163
                                                                }
                                                                if true {
                                                                    v3 = append(v3, 4294967295)
                                                                    var v187 string = v136
                                                                    var v188 float32 = 1e10
                                                                    _ = v187
                                                                    _ = v188
                                                                } else {
                                                                    v189 := GenS17{F0: (!v4), F1: v1}
                                                                    _ = v189
                                                                }
                                                                v1 -= v1
                                                                v1 += v1
                                                            }
                                                            _ = fi160
                                                        }
                                                        switch v136 {
                                                        case "\n\t":
                                                            var v190 uint32 = 524
                                                            _ = v190
                                                        }
                                                        if true {
                                                            if v4 {
                                                                var v191 float32 = float32(v99)
                                                                _ = v191
                                                            } else {
                                                                switch (fi138 + fi138) {
                                                                case -2147483648:
                                                                    if (false && false) {
                                                                        switch (fi138 + fi138) {
                                                                        case 32768:
                                                                            var v192 int16 = (-(v1))
                                                                            _ = v192
                                                                        }
                                                                        var v193 uint16 = 2
                                                                        var v194 string = "\n\t"
                                                                        _ = v193
                                                                        _ = v194
                                                                    }
                                                                case 2147483647:
                                                                    var v195 uint = 2
                                                                    _ = v195
                                                                case -611:
                                                                    v1 -= v1
                                                                default:
                                                                    for fi196 := 0; fi196 < 9; fi196++ {
                                                                        v136 = (v136 + "a")
                                                                        if len(v6) > 0 {
                                                                            v6[0] = fi138
                                                                        }
                                                                        v5 = 2
                                                                        _ = fi196
                                                                    }
                                                                }
                                                                var v197 int64 = 9223372036854775807
                                                                v197 += (v197 - int64(fi138))
                                                                _ = v197
                                                            }
                                                            if v4 {
                                                                v3 = append(v3, 0)
                                                                v6 = append(v6, (fi138 & -399))
                                                            }
                                                            var v198 uint16 = 783
                                                            _ = v198
                                                        }
                                                    }
                                                    _ = v146
                                                }
                                            default:
                                                v136 = (v136 + (v136 + v136))
                                            }
                                        }
                                    case 338:
                                        for fi199 := 0; fi199 < 3; fi199++ {
                                            if len(v6) > 0 {
                                                v6[0] = fi138
                                            }
                                            var v200 float64 = float64(0.1)
                                            if false {
                                                v3 = append(v3, 42)
                                            }
                                            _ = fi199
                                            _ = v200
                                        }
                                    default:
                                        var v201 uint = 1
                                        _ = v201
                                    }
                                case "\n\t":
                                    for fi202 := 0; fi202 < 6; fi202++ {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = len(v6)
                                        v203 := GenS15{F0: -694, F1: 1000, F2: (v2 - v2)}
                                        var v204 uint8 = 255
                                        _ = fi202
                                        _ = v203
                                        _ = v204
                                    }
                                default:
                                    var v205 bool = ((2 != 0) && v4)
                                    _ = v205
                                }
                            case 0:
                                v2 += (-(v2))
                            default:
                                fmt.Println(v2)
                            }
                        }
                        _ = fi138
                    }
                    _ = v136
                    _ = v137
                }
            }
        case -2147483648:
            if len(v3) > 0 {
                v3[0] = 2147483648
            }
        default:
            v5 -= (^(v5))
        }
        if v4 {
            v5 += (v5 ^ 0)
            if false {
                v2 = v2
                switch 469 {
                case 9223372036854775807:
                    v2 = (v2 % ((1) | 1))
                case 42:
                    for fi206 := 0; fi206 < 19; fi206++ {
                        v99 -= v99
                        switch "hello" {
                        case "\n\t":
                            switch "\n\t" {
                            case "foo bar":
                                for fi207 := 0; fi207 < 4; fi207++ {
                                    v2 += (v2 & -128)
                                    _ = fi207
                                }
                            case "hello":
                                var v208 int16 = v1
                                _ = v208
                            case "a":
                                if v4 {
                                    fmt.Println(fi206)
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                }
                            default:
                                for fi209 := 0; fi209 < 6; fi209++ {
                                    if v4 {
                                        v210 := GenS17{F0: ((4294967295 == 2) && true), F1: -596}
                                        v211 := GenS1{F0: (!((1 < 0) || false)), F1: "a", F2: 42}
                                        _ = v210
                                        _ = v211
                                    } else {
                                        v5 = (v5 & v5)
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = 2147483647
                                    }
                                    _ = fi209
                                }
                            }
                        case "a":
                            switch fi206 {
                            case 574:
                                fmt.Println(v5)
                            case 42:
                                v212 := GenS3{F0: 127, F1: 4294967295}
                                _ = v212
                            }
                        case "foo bar":
                            switch fi206 {
                            case -173:
                                for fi213 := 0; fi213 < 3; fi213++ {
                                    v4 = v4
                                    var v214 float64 = v99
                                    _ = fi213
                                    _ = v214
                                }
                            case 128:
                                v4 = false
                            case -530:
                                if false {
                                    var v215 float64 = float64(((0.0 + float32((v99 * ((v99 * (v99 / v99)) + v99)))) * 1.5))
                                    _ = v215
                                } else {
                                    var v216 int64 = 335
                                    v99 -= v99
                                    switch "" {
                                    case "hello":
                                        for fi217 := 0; fi217 < 11; fi217++ {
                                            var v218 uint32 = 2147483648
                                            if v4 {
                                                v219 := GenS9{F0: v216, F1: 2147483647}
                                                _ = v219
                                            } else {
                                                var v220 uint8 = 255
                                                if v4 {
                                                    fmt.Println(v5)
                                                }
                                                var v221 int64 = -344
                                                _ = v220
                                                _ = v221
                                            }
                                            switch "a" {
                                            case "foo bar":
                                                var v222 bool = v4
                                                _ = v222
                                            case "hello":
                                                v223 := GenS12{F0: v218, F1: 458}
                                                _ = v223
                                            }
                                            _ = fi217
                                            _ = v218
                                        }
                                    case "foo bar":
                                        switch (-(fi206)) {
                                        case 128:
                                            var v224 uint = 4294967295
                                            _ = v224
                                        case -396:
                                            var v225 int = (fi206 | (-(fi206)))
                                            _ = v225
                                        case -234:
                                            v1 -= v1
                                        }
                                    default:
                                        fmt.Println(v5, v2)
                                    }
                                    _ = v216
                                }
                            default:
                                v5 += 732
                            }
                        default:
                            var v226 bool = (2 > 42)
                            _ = v226
                        }
                        for fi227 := 0; fi227 < 7; fi227++ {
                            switch (fi227 & (fi227 / (((fi206 * fi227)) | 1))) {
                            case -810:
                                v5 = v5
                            }
                            if v4 {
                                if v4 {
                                    var v228 float64 = ((((1.5 / (3.14159 / 1.0)) * v99) + 1e10) - (v99 - (((v99 - v99) + 1.5) + v99)))
                                    switch 2147483647 {
                                    case 1:
                                        var v229 uint8 = 128
                                        _ = v229
                                    }
                                    _ = v228
                                }
                                v230 := GenS18{F0: 0}
                                _ = v230
                            } else {
                                var v231 int16 = (v1 % ((42) | 1))
                                _ = v231
                            }
                            _ = fi227
                        }
                        _ = fi206
                    }
                }
            } else {
                for fi232 := 0; fi232 < 14; fi232++ {
                    if true {
                        if (-323 >= -436) {
                            if ((1 <= (v5 & 18446744073709551615)) && false) {
                                v233 := GenS19{F0: 0, F1: 1}
                                _ = v233
                            }
                            if v4 {
                                if ((("a" + "a") + "a") != ((((("foo bar" + "\n\t") + ("hello" + ((("foo bar" + "foo bar") + "\n\t") + ("foo bar" + ("\n\t" + (("a" + "") + "a")))))) + ("" + "\n\t")) + "hello") + "")) {
                                    v5 = 78
                                }
                                v4 = false
                            } else {
                                for fi234 := 0; fi234 < 5; fi234++ {
                                    if (v5 != 2) {
                                        v235 := GenS16{F0: (2 == 2), F1: 457}
                                        v3 = append(v3, 376)
                                        v4 = v4
                                        switch (fi232 % ((int(v2)) | 1)) {
                                        case -9223372036854775808:
                                            var v236 uint16 = 32767
                                            _ = v236
                                        default:
                                            if len(v6) > 0 {
                                                v6[0] = (fi232 - (fi232 - fi232))
                                            }
                                        }
                                        _ = v235
                                    } else {
                                        v237 := GenS20{F0: 565}
                                        var v238 uint16 = 42
                                        _ = v237
                                        _ = v238
                                    }
                                    v239 := GenS21{F0: v2}
                                    v99 += 1.5
                                    _ = fi234
                                    _ = v239
                                }
                                var v240 int16 = -337
                                if v4 {
                                    var v241 string = ""
                                    v6 = append(v6, -1)
                                    v242 := GenS20{F0: 32767}
                                    _ = v241
                                    _ = v242
                                }
                                _ = v240
                            }
                            fmt.Println(v1)
                        } else {
                            if ((v5 >> ((255) & 63)) < 4294967295) {
                                for fi243 := 0; fi243 < 13; fi243++ {
                                    var v244 float64 = (v99 * -2.25)
                                    _ = fi243
                                    _ = v244
                                }
                                v1 += v1
                                for fi245 := 0; fi245 < 15; fi245++ {
                                    var v246 int32 = 2147483647
                                    _ = fi245
                                    _ = v246
                                }
                            } else {
                                for fi247 := 0; fi247 < 9; fi247++ {
                                    v248 := GenS22{F0: (^(v1))}
                                    fmt.Println(v2, v1, fi232)
                                    _ = fi247
                                    _ = v248
                                }
                            }
                            for fi249 := 0; fi249 < 17; fi249++ {
                                switch "\n\t" {
                                case "a":
                                    var v250 int64 = -942
                                    _ = v250
                                default:
                                    if len(v3) > 0 {
                                        v3[0] = 4294967295
                                    }
                                }
                                switch "foo bar" {
                                case "a":
                                    fmt.Println(v2)
                                case "\n\t":
                                    v1 = v1
                                default:
                                    _ = len(v6)
                                }
                                _ = fi249
                            }
                        }
                    }
                    _ = fi232
                }
            }
            var v251 uint8 = 1
            _ = v251
        } else {
            v252 := GenS1{F0: v4, F1: (((("hello" + "") + "") + ((("hello" + ("foo bar" + ("hello" + ("" + "\n\t")))) + "a") + ("" + "foo bar"))) + ("foo bar" + ((("a" + "") + "a") + (("" + "foo bar") + "foo bar")))), F2: 2}
            var v253 int32 = 1
            _ = v252
            _ = v253
        }
        var v254 float32 = 1e10
        _ = v254
    } else {
        var v255 uint16 = 311
        var v256 uint64 = (^(v5))
        _ = v255
        _ = v256
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v99
}
