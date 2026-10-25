package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 bool; F2 uint32 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 uint; F1 uint8; F2 int32 }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS7 struct { F0 int8; F1 int32 }
type GenS8 struct { F0 int64; F1 string; F2 int32 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 string; F1 uint16; F2 uint8 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS13 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 string; F1 int32 }
type GenS16 struct { F0 uint32; F1 uint; F2 int8 }
type GenS17 struct { F0 int16; F1 uint64; F2 int16 }
type GenS18 struct { F0 uint64; F1 uint32 }
type GenS19 struct { F0 uint; F1 int16 }
func genFn0() int64 {
    v7 := []int{-134, -158}
    if true {
        for fi8 := 0; fi8 < 7; fi8++ {
            for fi9 := 0; fi9 < 11; fi9++ {
                var v10 string = ("hello" + "foo bar")
                v10 = v10
                if true {
                    switch v10 {
                    case "hello":
                        fmt.Println(fi9)
                    }
                    var v11 int = fi8
                    _ = v11
                } else {
                    if (-9223372036854775808 != 387) {
                        if (42 <= 0) {
                            v10 = v10
                            v12 := GenS1{F0: (true || (v10 == "foo bar"))}
                            fmt.Println(fi9, fi9)
                            fmt.Println(fi9, fi9)
                            _ = v12
                        }
                        var v13 uint16 = 65535
                        if (4294967295 >= 42) {
                            fmt.Println(v10, fi8)
                            if (true && false) {
                                _ = len(v7)
                            }
                            var v14 int64 = 1
                            v15 := GenS1{F0: (4294967295 >= 1)}
                            if (!(true && false)) {
                                v16 := GenS2{F0: -1, F1: false, F2: 42}
                                var v17 int16 = 1000
                                _ = v16
                                _ = v17
                            } else {
                                if true {
                                    var v18 int32 = -1
                                    switch "\n\t" {
                                    case "hello":
                                        v13 = (v13 % (((v13 << ((1) & 63))) | 1))
                                    case "foo bar":
                                        switch fi9 {
                                        case 9223372036854775807:
                                            switch (fi8 | int(v13)) {
                                            case 4294967295:
                                                _ = len(v7)
                                            }
                                        }
                                    case "a":
                                        var v19 uint16 = v13
                                        v10 = (v10 + "")
                                        var v20 int = (fi8 % (((fi8 / (((fi9 | fi8)) | 1))) | 1))
                                        _ = v19
                                        _ = v20
                                    }
                                    _ = v18
                                }
                                if len(v7) > 0 {
                                    v7[0] = (fi9 - 6)
                                }
                            }
                            var v21 int32 = -20
                            if ((!true) || false) {
                                var v22 string = v10
                                _ = v22
                            }
                            fmt.Println(fi9, v13, v13)
                            var v23 bool = false
                            _ = v14
                            _ = v15
                            _ = v21
                            _ = v23
                        } else {
                            var v24 string = "foo bar"
                            _ = v24
                        }
                        _ = v13
                    } else {
                        if false {
                            var v25 uint = 18446744073709551615
                            var v26 int8 = 107
                            switch (((("hello" + ("" + ("" + "hello"))) + "") + "foo bar") + v10) {
                            case "\n\t":
                                var v27 int64 = 1
                                _ = v27
                            case "hello":
                                var v28 uint32 = uint32(v25)
                                _ = v28
                            }
                            _ = len(v7)
                            v10 = (v10 + v10)
                            var v29 int32 = -1
                            _ = v25
                            _ = v26
                            _ = v29
                        }
                        if false {
                            v10 = ""
                            var v30 int64 = -1
                            v30 = v30
                            var v31 uint32 = 2
                            v30 -= (v30 * -9223372036854775808)
                            _ = v30
                            _ = v31
                        } else {
                            var v32 int8 = 1
                            v32 = v32
                            v33 := GenS3{F0: 1, F1: 1}
                            _ = v32
                            _ = v33
                        }
                    }
                    fmt.Println(fi9, fi8, fi8)
                }
                var v34 float32 = ((0.1 - 0.0) / -2.25)
                v35 := GenS4{F0: 0, F1: 255, F2: 203}
                _ = fi9
                _ = v10
                _ = v34
                _ = v35
            }
            for fi36 := 0; fi36 < 1; fi36++ {
                var v37 uint = 65535
                if len(v7) > 0 {
                    v7[0] = fi8
                }
                _ = fi36
                _ = v37
            }
            fmt.Println(fi8)
            _ = fi8
        }
    }
    _ = v7
    return 516
}
func genFn1(p0 int32, p1 bool, p2 uint16) string {
    fmt.Println(p1)
    var v38 int32 = -1
    var v39 int32 = (p0 >> ((548) & 63))
    if false {
        fmt.Println(p2, v39, v39)
        v40 := []int{-845}
        _ = v40
    } else {
        if false {
            fmt.Println(p2)
            if p1 {
                var v41 int64 = -1
                v39 = 42
                _ = v41
            }
        } else {
            for fi42 := 0; fi42 < 1; fi42++ {
                var v43 uint = 216
                switch "foo bar" {
                case "foo bar":
                    var v44 int = fi42
                    _ = v44
                }
                _ = fi42
                _ = v43
            }
            var v45 uint32 = 2147483648
            p2 -= p2
            _ = v45
        }
        v46 := []int{-963, 770, 128, -903}
        if false {
            v47 := GenS5{F0: (p2 / ((p2) | 1)), F1: 9223372036854775807}
            v48 := GenS6{F0: 919, F1: 2, F2: 497}
            _ = v47
            _ = v48
        }
        _ = v46
    }
    var v49 int = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v38
    _ = v39
    _ = v49
    return (((("a" + "a") + "foo bar") + ((("hello" + (("hello" + "\n\t") + ("hello" + ""))) + ("a" + "foo bar")) + "foo bar")) + "foo bar")
}

func main() {
    var v1 uint32 = 4294967295
    var v2 bool = false
    var v3 bool = (!v2)
    var v4 uint = 1
    v5 := make(map[string]bool)
    v6 := make(map[uint32]bool)
    for fi50 := 0; fi50 < 14; fi50++ {
        v1 -= (^(v1))
        v51 := GenS7{F0: -1, F1: 42}
        switch 42 {
        case 85:
            if (!((v1 >> (((^(v4))) & 63)) <= (^(v1)))) {
                var v52 uint16 = 1
                _ = v52
            } else {
                for fi53 := 0; fi53 < 8; fi53++ {
                    if ((0 == 100) && (1 >= 2)) {
                        switch fi50 {
                        case 32768:
                            fmt.Println(v1)
                        case 32767:
                            var v54 int = (fi50 + fi50)
                            _ = v54
                        default:
                            if v3 {
                                fmt.Println(fi50, v3, fi50)
                            }
                        }
                        v55 := GenS4{F0: (v4 - (^(v4))), F1: 101, F2: 42}
                        if false {
                            v4 = (v4 * v4)
                            if v2 {
                                var v56 float32 = 1e10
                                v2 = (v3 || false)
                                var v57 bool = true
                                _ = v56
                                _ = v57
                            }
                        } else {
                            var v58 bool = true
                            _ = v58
                        }
                        v59 := GenS5{F0: 366, F1: (fi53 * -2147483648)}
                        v60 := GenS2{F0: 744, F1: v3, F2: (v1 >> ((v4) & 63))}
                        _ = v55
                        _ = v59
                        _ = v60
                    } else {
                        fmt.Println(fi50, v4, v2)
                        if v3 {
                            if (((909 != 0) && v2) && (42 < 18446744073709551615)) {
                                _ = len(v5)
                                v61 := []int{370}
                                if v2 {
                                    switch 32767 {
                                    case 0:
                                        _ = len(v61)
                                    case 600:
                                        if true {
                                            var v62 int = 128
                                            var v63 string = "foo bar"
                                            v64 := GenS8{F0: 9223372036854775807, F1: "\n\t", F2: -1}
                                            switch v63 {
                                            case "a":
                                                v63 = (("a" + ("a" + (v63 + v63))) + v63)
                                            }
                                            _ = v62
                                            _ = v63
                                            _ = v64
                                        } else {
                                            v65 := GenS9{F0: (^(fi50))}
                                            if (true && v2) {
                                                if v2 {
                                                    v2 = (!v3)
                                                } else {
                                                    v66 := GenS2{F0: -32768, F1: v2, F2: v1}
                                                    v3 = (1 == (v4 | v4))
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        v67 := GenS1{F0: (18446744073709551615 < 0)}
                                                        _ = v67
                                                    case "\n\t":
                                                        var v68 int64 = 9223372036854775807
                                                        _ = v68
                                                    default:
                                                        if false {
                                                            delete(v6, 65535)
                                                            var v69 bool = v2
                                                            _ = v69
                                                        }
                                                    }
                                                    _ = v66
                                                }
                                            }
                                            var v70 bool = v3
                                            fmt.Println(fi50)
                                            if (int16(v1) < -529) {
                                                switch (-(fi53)) {
                                                case 3:
                                                    v4 += 190
                                                case 842:
                                                    if v3 {
                                                        switch fi53 {
                                                        case 32767:
                                                            if ("" >= "hello") {
                                                                v70 = ("hello" >= (("a" + ("" + ((("a" + ("" + "a")) + "hello") + "hello"))) + ""))
                                                            } else {
                                                                if v2 {
                                                                    v2 = true
                                                                    v4 -= v4
                                                                    v71 := GenS3{F0: 131, F1: 613}
                                                                    _ = len(v5)
                                                                    v72 := v6[(v1 << (((v4 % (((v4 >> (((v4 - 2)) & 63))) | 1))) & 63))]
                                                                    _ = v71
                                                                    _ = v72
                                                                }
                                                            }
                                                        case 2147483647:
                                                            v70 = v2
                                                        case 0:
                                                            v2 = (0 >= 0)
                                                        default:
                                                            v73 := GenS3{F0: -32768, F1: 380}
                                                            _ = v73
                                                        }
                                                        fmt.Println(v70, v4)
                                                    } else {
                                                        v2 = (!v2)
                                                        var v74 int16 = 0
                                                        if (0 >= 0) {
                                                            switch "" {
                                                            case "a":
                                                                var v75 uint16 = 1
                                                                _ = v75
                                                            case "\n\t":
                                                                switch "a" {
                                                                case "a":
                                                                    if ((0 <= 32767) || v3) {
                                                                        var v76 string = "a"
                                                                        v77 := v6[(v1 | 2147483647)]
                                                                        v78 := GenS10{F0: "", F1: uint16(v4), F2: 0}
                                                                        _ = v76
                                                                        _ = v77
                                                                        _ = v78
                                                                    }
                                                                default:
                                                                    switch (("hello" + "") + "hello") {
                                                                    case "\n\t":
                                                                        fmt.Println(v2, v1)
                                                                    case "a":
                                                                        v79 := GenS11{F0: -921}
                                                                        _ = v79
                                                                    case "":
                                                                        var v80 int16 = (v74 ^ v74)
                                                                        _ = v80
                                                                    default:
                                                                        if (798 < -2147483648) {
                                                                            var v81 int64 = 42
                                                                            _ = v81
                                                                        }
                                                                    }
                                                                }
                                                            default:
                                                                _ = len(v5)
                                                            }
                                                            var v82 uint16 = 482
                                                            if v70 {
                                                                if v70 {
                                                                    var v83 uint64 = uint64(v82)
                                                                    _ = v83
                                                                } else {
                                                                    v70 = (("" + "hello") > "\n\t")
                                                                    fmt.Println(fi50, fi50, v74)
                                                                    var v84 bool = (0 <= -648)
                                                                    _ = v84
                                                                }
                                                            } else {
                                                                v85 := v6[4294967295]
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        v6[42] = v2
                                                                    case "":
                                                                        var v86 int = (fi50 % (((fi53 * -703)) | 1))
                                                                        _ = v86
                                                                    case "hello":
                                                                        var v87 int = 4294967295
                                                                        _ = v87
                                                                    }
                                                                    v82 += v82
                                                                    var v88 bool = (!false)
                                                                    delete(v5, ("foo bar" + ""))
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        var v89 int32 = 0
                                                                        _ = v89
                                                                    case "":
                                                                        var v90 bool = (1 > 0)
                                                                        _ = v90
                                                                    }
                                                                    _ = v88
                                                                case "\n\t":
                                                                    if ((true || false) && v3) {
                                                                        v91 := v5["foo bar"]
                                                                        var v92 uint64 = 2
                                                                        switch 4294967295 {
                                                                        case 2147483647:
                                                                            fmt.Println(v2, v3, v92)
                                                                        case -9223372036854775808:
                                                                            v92 = 1
                                                                        case 4294967295:
                                                                            v6[v1] = v70
                                                                        }
                                                                        switch ("" + ((("hello" + "hello") + ("" + ("a" + "\n\t"))) + "foo bar")) {
                                                                        case "foo bar":
                                                                            var v93 float32 = ((0.1 + ((((1e10 / (1.5 + float32(3.14159))) * (1e10 / 0.1)) - (-2.25 - (1e10 * (-2.25 + 1.5)))) * 3.14159)) * 1e10)
                                                                            v94 := GenS2{F0: -582, F1: false, F2: (v1 ^ (v1 * 1))}
                                                                            if true {
                                                                                var v95 uint8 = 174
                                                                                v4 = v4
                                                                                v96 := GenS2{F0: 1, F1: v2, F2: (v1 | v1)}
                                                                                _ = v95
                                                                                _ = v96
                                                                            } else {
                                                                                var v97 uint64 = v92
                                                                                v98 := GenS3{F0: v74, F1: (v97 % ((v92) | 1))}
                                                                                _ = v97
                                                                                _ = v98
                                                                            }
                                                                            var v99 int = 2147483647
                                                                            var v100 uint = (v4 % ((v4) | 1))
                                                                            _ = v100
                                                                            _ = v93
                                                                            _ = v94
                                                                            _ = v99
                                                                        }
                                                                        _ = v91
                                                                        _ = v92
                                                                    }
                                                                }
                                                                var v101 uint64 = 1
                                                                var v102 int8 = 0
                                                                _ = v101
                                                                _ = v102
                                                                _ = v85
                                                            }
                                                            _ = v82
                                                        } else {
                                                            if ((v74 & 42) < (v74 - v74)) {
                                                                v103 := GenS7{F0: 1, F1: -1}
                                                                _ = v103
                                                            } else {
                                                                var v104 string = (("foo bar" + "\n\t") + "foo bar")
                                                                fmt.Println(v3, v2)
                                                                switch "\n\t" {
                                                                case "a":
                                                                    var v105 uint8 = 255
                                                                    _ = v105
                                                                case "foo bar":
                                                                    switch (fi53 * (fi53 % ((fi50) | 1))) {
                                                                    case -792:
                                                                        if (!(v104 == ("hello" + v104))) {
                                                                            fmt.Println(fi53, v1, v3)
                                                                        }
                                                                    case 129:
                                                                        switch (fi50 >> ((v4) & 63)) {
                                                                        case 128:
                                                                            if len(v61) > 0 {
                                                                                v61[0] = 4294967295
                                                                            }
                                                                        }
                                                                    case 32768:
                                                                        v106 := GenS9{F0: fi50}
                                                                        _ = v106
                                                                    default:
                                                                        var v107 float32 = 3.14159
                                                                        _ = v107
                                                                    }
                                                                }
                                                                _ = v104
                                                            }
                                                        }
                                                        _ = v74
                                                    }
                                                }
                                                switch (fi53 & 32767) {
                                                case -817:
                                                    var v108 float32 = 3.14159
                                                    _ = v108
                                                case 4294967295:
                                                    fmt.Println(v4, v70)
                                                case 32767:
                                                    var v109 int = 211
                                                    _ = v109
                                                default:
                                                    if false {
                                                        var v110 int32 = -1
                                                        delete(v5, "hello")
                                                        _ = v110
                                                    } else {
                                                        var v111 int16 = 223
                                                        _ = v111
                                                    }
                                                }
                                            }
                                            _ = v65
                                            _ = v70
                                        }
                                    }
                                    switch (fi53 | fi50) {
                                    case 415:
                                        var v112 uint16 = 32767
                                        _ = v112
                                    case 754:
                                        v3 = ((((3.14159 - 0.0) / (float64(1.5) + 1e10)) * 0.1) >= 1e10)
                                    case 4294967295:
                                        v4 = (^(v4))
                                        v113 := GenS6{F0: 0, F1: 65535, F2: 18446744073709551615}
                                        if (402 <= (fi53 / (((fi50 >> ((uint(v1)) & 63))) | 1))) {
                                            _ = len(v5)
                                        } else {
                                            v3 = v2
                                        }
                                        v1 -= v1
                                        _ = v113
                                    }
                                    v114 := GenS12{F0: 1, F1: 1, F2: 2147483647}
                                    _ = v114
                                } else {
                                    _ = len(v61)
                                }
                                _ = v61
                            }
                        } else {
                            v115 := GenS12{F0: 465, F1: 1, F2: (^(v1))}
                            v116 := GenS13{F0: 4294967295, F1: 65535, F2: -32768}
                            v1 = (v1 | v1)
                            _ = v115
                            _ = v116
                        }
                    }
                    v117 := []int{(fi53 - fi50), (fi50 - fi53), (fi50 - 128), -2147483648}
                    v118 := GenS14{F0: 42}
                    _ = fi53
                    _ = v117
                    _ = v118
                }
                for fi119 := 0; fi119 < 15; fi119++ {
                    switch (fi50 << (((^(v4))) & 63)) {
                    case 32768:
                        v120 := []int{(fi50 + 761), (fi50 / ((fi50) | 1))}
                        _ = v120
                    default:
                        v121 := GenS15{F0: ("foo bar" + "a"), F1: -2147483648}
                        _ = v121
                    }
                    v1 = (^(v1))
                    _ = fi119
                }
            }
        }
        _ = fi50
        _ = v51
    }
    for fi122 := 0; fi122 < 8; fi122++ {
        switch ("foo bar" + "hello") {
        case "foo bar":
            if false {
                for fi123 := 0; fi123 < 7; fi123++ {
                    if (!true) {
                        var v124 int32 = -2147483648
                        var v125 int = (fi122 | 9223372036854775807)
                        _ = v124
                        _ = v125
                    } else {
                        var v126 uint = (v4 * (v4 | v4))
                        _ = v126
                    }
                    v127 := []int{(fi122 >> ((v4) & 63)), (-(fi122)), (fi123 | fi122)}
                    v128 := GenS16{F0: (v1 | uint32(v4)), F1: v4, F2: 42}
                    _ = fi123
                    _ = v127
                    _ = v128
                }
                if false {
                    v3 = true
                    delete(v6, (v1 >> (((v4 / ((v4) | 1))) & 63)))
                    if true {
                        if (42 < 42) {
                            var v129 int8 = 0
                            _ = v129
                        } else {
                            var v130 uint32 = (v1 / (((v1 / ((v1) | 1))) | 1))
                            var v131 int = (fi122 << ((v4) & 63))
                            _ = v130
                            _ = v131
                        }
                    } else {
                        v132 := GenS13{F0: 4294967295, F1: 1, F2: -96}
                        switch fi122 {
                        case 42:
                            v4 -= v4
                        case 32767:
                            if ((v4 ^ v4) < v4) {
                                var v133 uint16 = 65535
                                for fi134 := 0; fi134 < 20; fi134++ {
                                    if v2 {
                                        fmt.Println(fi122, fi134, fi122)
                                        var v135 float64 = 1.5
                                        if v2 {
                                            v2 = v2
                                            v133 = 42
                                            switch fi122 {
                                            case 186:
                                                var v136 bool = v2
                                                _ = v136
                                            default:
                                                switch "" {
                                                case "foo bar":
                                                    v133 = (v133 % ((v133) | 1))
                                                case "\n\t":
                                                    if (((519 == -656) && v3) && (-1 < 42)) {
                                                        switch "" {
                                                        case "\n\t":
                                                            v4 = v4
                                                        }
                                                    } else {
                                                        var v137 int64 = 0
                                                        v138 := v5[((((("hello" + "\n\t") + "a") + "\n\t") + "") + "foo bar")]
                                                        _ = v137
                                                        _ = v138
                                                    }
                                                }
                                            }
                                            v135 = ((v135 / (-2.25 * v135)) / ((0.0 + v135) - v135))
                                            if v2 {
                                                v3 = v2
                                                var v139 int64 = 446
                                                _ = v139
                                            }
                                            v140 := []int{9223372036854775807, (fi134 + fi134), 0}
                                            _ = v140
                                        } else {
                                            v1 = v1
                                        }
                                        _ = v135
                                    } else {
                                        if ((!(v2 || true)) && (!(127 >= -49))) {
                                            v141 := []int{}
                                            _ = v141
                                        }
                                        if ((1 < 0) || v3) {
                                            v5[("foo bar" + "")] = (("hello" + ("a" + "foo bar")) < "")
                                            var v142 uint8 = 255
                                            v142 -= 255
                                            _ = v142
                                        }
                                        v143 := GenS8{F0: -646, F1: "", F2: int32(v1)}
                                        _ = v143
                                    }
                                    var v144 uint16 = (v133 * (v133 >> (((v4 ^ v4)) & 63)))
                                    _ = fi134
                                    _ = v144
                                }
                                _ = v133
                            }
                        }
                        _ = v132
                    }
                } else {
                    if ((0.1 <= -2.25) && v2) {
                        var v145 int32 = -732
                        _ = v145
                    } else {
                        for fi146 := 0; fi146 < 16; fi146++ {
                            var v147 uint32 = v1
                            v148 := []int{int(v4), fi122, -9223372036854775808}
                            if ((v3 && (42 >= 0)) || (69 != -698)) {
                                if len(v148) > 0 {
                                    v148[0] = int(v4)
                                }
                            }
                            _ = fi146
                            _ = v147
                            _ = v148
                        }
                        switch (fi122 / ((int(v1)) | 1)) {
                        case 9223372036854775807:
                            if v3 {
                                v149 := GenS17{F0: 32767, F1: 1, F2: -1}
                                v150 := GenS13{F0: 2, F1: uint16(v4), F2: 32767}
                                v6[v1] = (4294967295 < 4294967295)
                                _ = v149
                                _ = v150
                            } else {
                                v151 := []int{127, 2147483647, fi122, -1}
                                v152 := GenS3{F0: 0, F1: 18446744073709551615}
                                if true {
                                    var v153 string = ("a" + "a")
                                    var v154 float32 = -2.25
                                    _ = v153
                                    _ = v154
                                } else {
                                    var v155 int32 = -1
                                    if false {
                                        var v156 uint = (^(v4))
                                        switch "foo bar" {
                                        case "foo bar":
                                            v157 := GenS17{F0: -1, F1: 18446744073709551615, F2: 501}
                                            _ = v157
                                        case "a":
                                            if (v3 && v2) {
                                                for fi158 := 0; fi158 < 4; fi158++ {
                                                    var v159 int32 = 163
                                                    _ = fi158
                                                    _ = v159
                                                }
                                                fmt.Println(v4)
                                                switch "hello" {
                                                case "":
                                                    if len(v151) > 0 {
                                                        v151[0] = fi122
                                                    }
                                                }
                                            } else {
                                                if v3 {
                                                    v3 = (!v2)
                                                    var v160 int32 = v155
                                                    _ = v160
                                                } else {
                                                    for fi161 := 0; fi161 < 16; fi161++ {
                                                        switch 4294967295 {
                                                        case -701:
                                                            switch ("a" + ("\n\t" + ("foo bar" + "foo bar"))) {
                                                            case "foo bar":
                                                                var v162 int64 = 799
                                                                _ = v162
                                                            case "a":
                                                                v163 := GenS10{F0: ("hello" + ""), F1: 32767, F2: 127}
                                                                _ = v163
                                                            case "":
                                                                v164 := GenS7{F0: 100, F1: -1}
                                                                _ = v164
                                                            default:
                                                                var v165 uint8 = 107
                                                                _ = v165
                                                            }
                                                        case -9223372036854775808:
                                                            delete(v6, v1)
                                                        case 9223372036854775807:
                                                            switch "\n\t" {
                                                            case "hello":
                                                                v166 := GenS8{F0: genFn0(), F1: "", F2: (v155 << ((v4) & 63))}
                                                                _ = v166
                                                            case "":
                                                                fmt.Println(v4, v3, v1)
                                                            }
                                                        default:
                                                            delete(v5, "hello")
                                                        }
                                                        if v3 {
                                                            var v167 int32 = (v155 | v155)
                                                            v4 += 454
                                                            _ = v167
                                                        }
                                                        _ = fi161
                                                    }
                                                    v168 := GenS18{F0: 4294967295, F1: v1}
                                                    v169 := GenS5{F0: 926, F1: -1}
                                                    _ = v168
                                                    _ = v169
                                                }
                                                var v170 bool = (fi122 <= 32)
                                                v155 = 0
                                                _ = v170
                                            }
                                        case "\n\t":
                                            var v171 int64 = -955
                                            _ = v171
                                        default:
                                            var v172 int = fi122
                                            _ = v172
                                        }
                                        _ = v156
                                    } else {
                                        for fi173 := 0; fi173 < 18; fi173++ {
                                            var v174 uint16 = 386
                                            var v175 int16 = 42
                                            var v176 uint16 = (v174 & v174)
                                            _ = fi173
                                            _ = v174
                                            _ = v175
                                            _ = v176
                                        }
                                        v4 = v4
                                    }
                                    _ = v155
                                }
                                _ = v151
                                _ = v152
                            }
                        case 128:
                            for fi177 := 0; fi177 < 13; fi177++ {
                                v2 = ((255 >= 0) && (v2 || v2))
                                _ = len(v6)
                                _ = fi177
                            }
                        case 4294967295:
                            v178 := v6[(v1 ^ v1)]
                            _ = v178
                        }
                    }
                }
                if ("hello" < "") {
                    for fi179 := 0; fi179 < 12; fi179++ {
                        fmt.Println(v4)
                        switch "a" {
                        case "hello":
                            var v180 uint64 = 2
                            _ = v180
                        case "a":
                            if (127 > 128) {
                                v5["hello"] = v2
                                v181 := GenS1{F0: true}
                                _ = v181
                            }
                        default:
                            v182 := v6[2]
                            fmt.Println(v3, v4, v4)
                            _ = v182
                        }
                        _ = fi179
                    }
                    fmt.Println(fi122, v1, v1)
                }
            }
        }
        _ = fi122
    }
    v3 = true
    for fi183 := 0; fi183 < 5; fi183++ {
        fmt.Println(v4)
        var v184 int64 = -512
        _ = fi183
        _ = v184
    }
    switch "hello" {
    case "foo bar":
        if ((0.1 != -2.25) && (!(!("hello" >= "hello")))) {
            var v185 int64 = 0
            _ = v185
        }
    case "":
        v1 = (v1 << ((v4) & 63))
    case "hello":
        v186 := []int{32767, 32768}
        _ = v186
    default:
        v3 = ("foo bar" != ("" + ("a" + "hello")))
    }
    v187 := GenS6{F0: 172, F1: 42, F2: 857}
    v1 = (^(v1))
    for fi188 := 0; fi188 < 9; fi188++ {
        if false {
            v4 -= v4
        } else {
            fmt.Println(v2, fi188)
            v2 = v2
        }
        var v189 uint16 = 42
        var v190 int64 = -1
        _ = fi188
        _ = v189
        _ = v190
    }
    switch -732 {
    case 9223372036854775807:
        v3 = (!(65535 > 42))
    case 32767:
        var v191 int32 = 0
        _ = v191
    }
    var v192 uint = v4
    switch "\n\t" {
    case "foo bar":
        var v193 uint8 = 247
        _ = v193
    case "a":
        v3 = (42 != 1)
    default:
        switch "foo bar" {
        case "foo bar":
            for fi194 := 0; fi194 < 11; fi194++ {
                var v195 int32 = 1
                switch (-(fi194)) {
                case 32768:
                    var v196 uint16 = 42
                    _ = v196
                default:
                    v197 := GenS17{F0: 0, F1: 4294967295, F2: 32767}
                    _ = v197
                }
                var v198 int = fi194
                _ = fi194
                _ = v195
                _ = v198
            }
        case "hello":
            v199 := GenS1{F0: (32767 < 383)}
            _ = v199
        default:
            for fi200 := 0; fi200 < 2; fi200++ {
                switch 128 {
                case -9223372036854775808:
                    fmt.Println(v3)
                case -2147483648:
                    var v201 uint64 = 808
                    _ = v201
                }
                v202 := GenS19{F0: (^(v4)), F1: 32767}
                _ = fi200
                _ = v202
            }
        }
    }
    fmt.Println(v1, v2)
    v192 = v192
    for fi203 := 0; fi203 < 15; fi203++ {
        var v204 string = "\n\t"
        _ = fi203
        _ = v204
    }
    if v2 {
        var v205 uint32 = v1
        var v206 int = 2147483647
        _ = v205
        _ = v206
    }
    v207 := GenS4{F0: v192, F1: 226, F2: -729}
    fmt.Println("v1:", v1)
    fmt.Println("v192:", v192)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v187
    _ = v192
    _ = v2
    _ = v207
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
