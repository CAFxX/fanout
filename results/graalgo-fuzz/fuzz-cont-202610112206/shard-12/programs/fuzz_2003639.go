package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 int32; F2 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 int8; F2 int16 }
type GenS5 struct { F0 int; F1 string }
type GenS6 struct { F0 int; F1 uint; F2 int }
type GenS7 struct { F0 int64; F1 string; F2 int16 }
type GenS8 struct { F0 int; F1 uint32 }
type GenS9 struct { F0 uint32; F1 int16 }
type GenS10 struct { F0 bool; F1 uint64 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int; F1 uint16; F2 int8 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int16; F1 int16 }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 uint16; F1 uint; F2 int8 }
type GenS18 struct { F0 int16; F1 uint }
type GenS19 struct { F0 int8 }
type GenS20 struct { F0 string }
type GenS21 struct { F0 string; F1 uint }
type GenS22 struct { F0 uint64 }
type GenS23 struct { F0 uint16; F1 int64; F2 int }

func main() {
    var v1 float64 = 1e10
    var v2 int64 = 1
    v3 := make(map[uint32]uint32)
    var v4 int8 = int8(v2)
    v5 := make(map[uint32]bool)
    v6 := make(map[uint32]uint32)
    var v7 string = "\n\t"
    switch "foo bar" {
    case "hello":
        v7 = "a"
    }
    v1 -= v1
    for fi8 := 0; fi8 < 11; fi8++ {
        if true {
            fmt.Println(v2, v2)
            _ = len(v5)
            v6[760] = 1
        } else {
            v7 = v7
        }
        delete(v6, 221)
        v9 := GenS1{F0: 128}
        _ = fi8
        _ = v9
    }
    v2 = v2
    if (v1 < v1) {
        v4 -= 0
        fmt.Println(v2, v2, v4)
        switch v7 {
        case "a":
            if (false && (false && ((!((1 >= 2) || false)) || (32767 >= 42)))) {
                switch 2147483647 {
                case -216:
                    fmt.Println(v2)
                default:
                    v10 := []int{}
                    _ = v10
                }
                for fi11 := 0; fi11 < 1; fi11++ {
                    if true {
                        v12 := []int{fi11, (fi11 >> ((2) & 63)), (^(fi11)), fi11}
                        var v13 uint32 = 2147483648
                        _ = v12
                        _ = v13
                    } else {
                        v1 += v1
                        if false {
                            v14 := []int{}
                            _ = len(v6)
                            _ = v14
                        }
                        fmt.Println(fi11)
                    }
                    _ = fi11
                }
                if (!true) {
                    v15 := []int{128, 128, 128, -963}
                    for fi16 := 0; fi16 < 13; fi16++ {
                        if true {
                            if true {
                                var v17 float32 = 1.5
                                if (!true) {
                                    for fi18 := 0; fi18 < 15; fi18++ {
                                        fmt.Println(v2, v4, fi16)
                                        v19 := v3[65535]
                                        v20 := GenS1{F0: 32767}
                                        _ = fi18
                                        _ = v19
                                        _ = v20
                                    }
                                }
                                if false {
                                    for fi21 := 0; fi21 < 13; fi21++ {
                                        v17 = (v17 * ((v17 * v17) / v17))
                                        _ = fi21
                                    }
                                    v1 -= float64(0.0)
                                } else {
                                    fmt.Println(v17, v1, v7)
                                    var v22 uint16 = 48
                                    if ((false || false) || false) {
                                        v23 := v6[765]
                                        v4 -= v4
                                        _ = v23
                                    } else {
                                        v24 := GenS1{F0: -339}
                                        _ = v24
                                    }
                                    _ = v22
                                }
                                _ = v17
                            }
                        }
                        var v25 uint64 = 145
                        for fi26 := 0; fi26 < 18; fi26++ {
                            switch fi26 {
                            case -2147483648:
                                _ = len(v3)
                            case -9223372036854775808:
                                v3[234] = 587
                            }
                            _ = fi26
                        }
                        _ = fi16
                        _ = v25
                    }
                    fmt.Println(v4, v4, v7)
                    _ = v15
                }
            }
        case "hello":
            v27 := GenS2{F0: 32767, F1: 0, F2: (v7 + v7)}
            _ = v27
        }
    } else {
        var v28 int64 = 9223372036854775807
        switch -2147483648 {
        case 32767:
            switch v7 {
            case "":
                v29 := []int{-2147483648}
                _ = v29
            default:
                var v30 int = -638
                _ = v30
            }
        case 0:
            if ((42 != 1) || false) {
                fmt.Println(v4)
            }
        }
        _ = v28
    }
    var v31 string = v7
    v32 := GenS3{F0: 1}
    v33 := GenS1{F0: -1}
    v1 = 1.5
    v3[1] = 750
    v4 = v4
    if true {
        if (-730 <= 9223372036854775807) {
            for fi34 := 0; fi34 < 1; fi34++ {
                var v35 int8 = -1
                v36 := GenS4{F0: 2, F1: 66, F2: -1}
                _ = fi34
                _ = v35
                _ = v36
            }
            v37 := []int{9223372036854775807, 333}
            for fi38 := 0; fi38 < 11; fi38++ {
                var v39 bool = ((2 > 323) || true)
                for fi40 := 0; fi40 < 11; fi40++ {
                    v41 := GenS4{F0: 1, F1: (v4 | (^(v4))), F2: 42}
                    v7 = (v7 + "a")
                    _ = fi40
                    _ = v41
                }
                for fi42 := 0; fi42 < 7; fi42++ {
                    v3[919] = 0
                    if (true || v39) {
                        var v43 int64 = (v2 & v2)
                        v7 = v7
                        if (!(32767 == -1)) {
                            v43 += 42
                            fmt.Println(v39, v43)
                        } else {
                            fmt.Println(v43)
                            fmt.Println(v2, v39)
                            var v44 uint16 = 32767
                            v43 += (v43 & 812)
                            _ = v44
                        }
                        _ = v43
                    } else {
                        var v45 uint = 4294967295
                        var v46 bool = v39
                        _ = v45
                        _ = v46
                    }
                    var v47 int = 31
                    v48 := GenS5{F0: fi38, F1: ""}
                    _ = fi42
                    _ = v47
                    _ = v48
                }
                _ = fi38
                _ = v39
            }
            _ = v37
        } else {
            for fi49 := 0; fi49 < 3; fi49++ {
                for fi50 := 0; fi50 < 13; fi50++ {
                    var v51 int8 = 93
                    if false {
                        switch fi50 {
                        case 1:
                            if (v51 < (v4 - v51)) {
                                var v52 int32 = 0
                                if (42 >= 32767) {
                                    var v53 string = (v31 + "hello")
                                    _ = v53
                                }
                                if (1 < 42) {
                                    delete(v6, 65535)
                                    switch (fi50 + fi49) {
                                    case 2147483647:
                                        v54 := []int{(fi49 / ((int(v51)) | 1)), fi49}
                                        _ = v54
                                    }
                                    var v55 int = fi50
                                    if ((((v52 > v52) && false) && ((-(fi50)) == 9223372036854775807)) && (true && (v7 == "foo bar"))) {
                                        if (((((!true) || true) && true) || (false || false)) && false) {
                                            switch fi50 {
                                            case 1:
                                                var v56 int16 = 882
                                                _ = v56
                                            case 127:
                                                if true {
                                                    switch (fi49 / ((42) | 1)) {
                                                    case -564:
                                                        fmt.Println(fi49, v2, v1)
                                                        v57 := GenS4{F0: 255, F1: v51, F2: -908}
                                                        var v58 float64 = 0.0
                                                        switch (fi49 + (v55 | fi49)) {
                                                        case 1:
                                                            var v59 int8 = (v51 / ((v4) | 1))
                                                            _ = v59
                                                        default:
                                                            v60 := []int{v55, 32767, v55, 2147483647}
                                                            _ = v60
                                                        }
                                                        _ = v57
                                                        _ = v58
                                                    default:
                                                        v51 += 28
                                                    }
                                                } else {
                                                    v7 = v31
                                                    switch (v55 >> ((18446744073709551615) & 63)) {
                                                    case -284:
                                                        var v61 float64 = (1e10 * -2.25)
                                                        _ = v61
                                                    case 1:
                                                        var v62 int64 = v2
                                                        _ = v62
                                                    }
                                                    var v63 uint8 = 42
                                                    _ = v63
                                                }
                                            case -525:
                                                v2 += v2
                                            default:
                                                if true {
                                                    if false {
                                                        if (698 >= 1) {
                                                            switch ((("a" + "\n\t") + ("a" + ("hello" + (v7 + (v7 + ("a" + "\n\t")))))) + (((v7 + "hello") + v31) + (v31 + v31))) {
                                                            case "foo bar":
                                                                fmt.Println(v7)
                                                            case "a":
                                                                switch 127 {
                                                                case -386:
                                                                    delete(v6, 0)
                                                                }
                                                            }
                                                            v4 = v51
                                                            switch "a" {
                                                            case "hello":
                                                                switch "hello" {
                                                                case "hello":
                                                                    var v64 uint32 = 849
                                                                    _ = v64
                                                                case "":
                                                                    var v65 bool = true
                                                                    _ = v65
                                                                case "foo bar":
                                                                    v31 = (v7 + v31)
                                                                default:
                                                                    var v66 uint8 = 10
                                                                    _ = v66
                                                                }
                                                            case "a":
                                                                if false {
                                                                    v1 -= v1
                                                                }
                                                            case "":
                                                                switch (fi49 & (fi50 & (v55 % ((v55) | 1)))) {
                                                                case 32768:
                                                                    v2 += v2
                                                                case 277:
                                                                    if (false && true) {
                                                                        v67 := []int{}
                                                                        switch v31 {
                                                                        case "hello":
                                                                            switch ("\n\t" + v7) {
                                                                            case "a":
                                                                                v68 := v3[1]
                                                                                if false {
                                                                                    _ = len(v67)
                                                                                    v68 += 2
                                                                                }
                                                                                _ = v68
                                                                            case "\n\t":
                                                                                var v69 int64 = int64(v52)
                                                                                _ = v69
                                                                            }
                                                                        case "\n\t":
                                                                            v31 = ("a" + v31)
                                                                        }
                                                                        _ = v67
                                                                    }
                                                                case -703:
                                                                    if (4294967295 >= 1) {
                                                                        v70 := []int{0, (fi49 >> ((4294967295) & 63))}
                                                                        v71 := GenS5{F0: 716, F1: v31}
                                                                        _ = v70
                                                                        _ = v71
                                                                    } else {
                                                                        var v72 bool = false
                                                                        switch "\n\t" {
                                                                        case "\n\t":
                                                                            v73 := GenS6{F0: 128, F1: 2, F2: 32767}
                                                                            _ = v73
                                                                        case "a":
                                                                            var v74 float32 = 1.5
                                                                            _ = v74
                                                                        }
                                                                        var v75 uint64 = 4294967295
                                                                        if (!(18446744073709551615 == 1)) {
                                                                            switch (-(v55)) {
                                                                            case 128:
                                                                                v75 -= 2
                                                                            case 974:
                                                                                if ((v72 || v72) && (42 >= v4)) {
                                                                                    switch 1 {
                                                                                    case -987:
                                                                                        if v72 {
                                                                                            if true {
                                                                                                v51 = 100
                                                                                            }
                                                                                        } else {
                                                                                            if (933 != 18446744073709551615) {
                                                                                                v76 := GenS6{F0: int(v51), F1: 42, F2: (-(fi50))}
                                                                                                var v77 int64 = (-(v2))
                                                                                                v78 := GenS4{F0: 902, F1: 127, F2: 0}
                                                                                                _ = v76
                                                                                                _ = v77
                                                                                                _ = v78
                                                                                            } else {
                                                                                                v79 := []int{}
                                                                                                var v80 uint = 18446744073709551615
                                                                                                v4 = (v51 % ((v4) | 1))
                                                                                                _ = v79
                                                                                                _ = v80
                                                                                            }
                                                                                            v81 := GenS7{F0: v2, F1: ((v31 + v7) + v7), F2: 0}
                                                                                            _ = v81
                                                                                        }
                                                                                    case 128:
                                                                                        var v82 bool = v72
                                                                                        _ = v82
                                                                                    }
                                                                                    v83 := GenS8{F0: 0, F1: 760}
                                                                                    _ = v83
                                                                                } else {
                                                                                    var v84 uint32 = 65535
                                                                                    switch "\n\t" {
                                                                                    case "foo bar":
                                                                                        if (0 != 599) {
                                                                                            v85 := []int{2147483647, (-(fi50)), fi49}
                                                                                            _ = v85
                                                                                        }
                                                                                    case "hello":
                                                                                        var v86 int16 = 1
                                                                                        _ = v86
                                                                                    case "\n\t":
                                                                                        v4 = -5
                                                                                    default:
                                                                                        v84 += v84
                                                                                        switch (((v31 + ((("a" + "hello") + v31) + (v31 + (v7 + (("hello" + v31) + v31))))) + ("\n\t" + v31)) + "") {
                                                                                        case "":
                                                                                            v6[(v84 + (v84 & v84))] = v84
                                                                                        case "\n\t":
                                                                                            v4 = v51
                                                                                        case "foo bar":
                                                                                            if false {
                                                                                                fmt.Println(v75, v1, v75)
                                                                                                v55 -= (fi49 & fi49)
                                                                                            } else {
                                                                                                fmt.Println(v52)
                                                                                                var v87 uint = 18446744073709551615
                                                                                                _ = v87
                                                                                            }
                                                                                        default:
                                                                                            var v88 uint64 = v75
                                                                                            _ = v88
                                                                                        }
                                                                                        v1 += v1
                                                                                    }
                                                                                    v2 = v2
                                                                                    _ = v84
                                                                                }
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v1, v55)
                                                                            v89 := GenS1{F0: (-(fi49))}
                                                                            v90 := GenS9{F0: 1, F1: 940}
                                                                            _ = v89
                                                                            _ = v90
                                                                        }
                                                                        _ = v72
                                                                        _ = v75
                                                                    }
                                                                    switch -311 {
                                                                    case 2147483647:
                                                                        v4 -= v51
                                                                        v4 -= v51
                                                                    case 32767:
                                                                        if (!(!true)) {
                                                                            v1 += float64(0.1)
                                                                            v91 := GenS10{F0: true, F1: 1}
                                                                            var v92 uint8 = 127
                                                                            v93 := v5[2147483648]
                                                                            _ = v91
                                                                            _ = v92
                                                                            _ = v93
                                                                        } else {
                                                                            switch fi49 {
                                                                            case 4294967295:
                                                                                if false {
                                                                                    if false {
                                                                                        fmt.Println(v2, v1, v2)
                                                                                    } else {
                                                                                        v55 = (fi49 ^ 457)
                                                                                        fmt.Println(v52, v55, fi49)
                                                                                        var v94 uint8 = 1
                                                                                        _ = v94
                                                                                    }
                                                                                    var v95 float64 = float64(-2.25)
                                                                                    var v96 int8 = v51
                                                                                    _ = v95
                                                                                    _ = v96
                                                                                } else {
                                                                                    var v97 int8 = v4
                                                                                    v98 := GenS11{F0: v2}
                                                                                    v99 := GenS12{F0: (fi50 / ((fi50) | 1)), F1: 0, F2: v4}
                                                                                    var v100 string = (v31 + v7)
                                                                                    _ = v100
                                                                                    _ = v97
                                                                                    _ = v98
                                                                                    _ = v99
                                                                                }
                                                                            case 9223372036854775807:
                                                                                v52 += v52
                                                                                fmt.Println(v4)
                                                                                var v101 int8 = (v51 * 18)
                                                                                v102 := GenS13{F0: 236}
                                                                                _ = v101
                                                                                _ = v102
                                                                            case 42:
                                                                                v52 = -2147483648
                                                                            }
                                                                            switch ((v7 + v7) + (v31 + v31)) {
                                                                            case "foo bar":
                                                                                fmt.Println(fi50, fi50)
                                                                            }
                                                                            var v103 uint32 = 65535
                                                                            _ = v103
                                                                        }
                                                                    default:
                                                                        v104 := GenS14{F0: 1000, F1: -32768}
                                                                        _ = v104
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        fmt.Println(v1, v7)
                                                    }
                                                    v105 := GenS7{F0: (^(v2)), F1: v7, F2: 32767}
                                                    _ = v105
                                                } else {
                                                    fmt.Println(fi50)
                                                    switch fi49 {
                                                    case 127:
                                                        v106 := GenS15{F0: 42}
                                                        _ = v106
                                                    default:
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            v107 := GenS16{F0: 1}
                                                            _ = v107
                                                        case "a":
                                                            v4 += v51
                                                        }
                                                    }
                                                    var v108 uint32 = 0
                                                    fmt.Println(v52)
                                                    _ = v108
                                                }
                                            }
                                            v55 = 32768
                                            if false {
                                                if true {
                                                    v109 := GenS12{F0: 9223372036854775807, F1: 255, F2: (v4 + v51)}
                                                    if (true || false) {
                                                        if (!((v51 * (v4 % (((^(v4))) | 1))) > -128)) {
                                                            if true {
                                                                _ = len(v6)
                                                                fmt.Println(v7, v1)
                                                                var v110 int = -116
                                                                _ = v110
                                                            } else {
                                                                var v111 float64 = 1.5
                                                                var v112 float32 = 3.14159
                                                                _ = v111
                                                                _ = v112
                                                            }
                                                            fmt.Println(v52)
                                                            var v113 uint8 = 25
                                                            var v114 int64 = -9223372036854775808
                                                            v114 += 0
                                                            _ = v113
                                                            _ = v114
                                                        }
                                                        if false {
                                                            fmt.Println(v55)
                                                            v115 := GenS17{F0: 255, F1: 0, F2: v51}
                                                            _ = v115
                                                        }
                                                        v52 -= v52
                                                        v51 -= v4
                                                        v2 = 0
                                                        v116 := GenS18{F0: -960, F1: 42}
                                                        _ = v116
                                                    }
                                                    _ = v109
                                                } else {
                                                    if true {
                                                        v5[2147483647] = true
                                                    } else {
                                                        delete(v3, 933)
                                                        if false {
                                                            v4 -= (v51 % ((-8) | 1))
                                                        } else {
                                                            if false {
                                                                v117 := GenS19{F0: -62}
                                                                _ = v117
                                                            } else {
                                                                v1 = 0.1
                                                                v118 := GenS2{F0: 1, F1: (v52 + v52), F2: v7}
                                                                v2 -= (v2 / ((-802) | 1))
                                                                fmt.Println(v4, v31)
                                                                _ = v118
                                                            }
                                                            v55 += 128
                                                        }
                                                        v119 := GenS20{F0: ("\n\t" + v7)}
                                                        var v120 int8 = v4
                                                        fmt.Println(v7, v31)
                                                        switch (v55 << ((11) & 63)) {
                                                        case -1:
                                                            if (65535 > 32767) {
                                                                v1 -= (-2.25 + -2.25)
                                                                var v121 string = v31
                                                                var v122 int8 = 0
                                                                _ = v121
                                                                _ = v122
                                                            } else {
                                                                v123 := GenS21{F0: "", F1: 929}
                                                                if (false || true) {
                                                                    v124 := GenS13{F0: 995}
                                                                    _ = v124
                                                                } else {
                                                                    switch v7 {
                                                                    case "hello":
                                                                        var v125 int16 = 224
                                                                        _ = v125
                                                                    }
                                                                    switch fi49 {
                                                                    case 380:
                                                                        if true {
                                                                            v2 += v2
                                                                            v126 := []int{(^(v55)), (v55 | (fi49 / ((fi50) | 1))), (fi49 >> ((65535) & 63))}
                                                                            switch v7 {
                                                                            case "hello":
                                                                                var v127 int8 = (v120 * v51)
                                                                                _ = v127
                                                                            }
                                                                            switch ("" + v31) {
                                                                            case "a":
                                                                                v128 := GenS8{F0: v55, F1: 441}
                                                                                _ = v128
                                                                            case "":
                                                                                v129 := GenS14{F0: 1, F1: 32767}
                                                                                _ = v129
                                                                            default:
                                                                                fmt.Println(fi50, v2)
                                                                                switch 9223372036854775807 {
                                                                                case 2147483647:
                                                                                    v130 := GenS7{F0: (^(v2)), F1: "\n\t", F2: 1000}
                                                                                    _ = v130
                                                                                case 32767:
                                                                                    var v131 float64 = v1
                                                                                    _ = v131
                                                                                case -1:
                                                                                    v120 += v4
                                                                                }
                                                                                v126 = append(v126, (fi49 & -317))
                                                                                v132 := GenS15{F0: 42}
                                                                                _ = v132
                                                                            }
                                                                            _ = v126
                                                                        } else {
                                                                            v133 := GenS1{F0: 9223372036854775807}
                                                                            _ = v133
                                                                        }
                                                                    case 4294967295:
                                                                        fmt.Println(v120)
                                                                    case 544:
                                                                        if (!false) {
                                                                            fmt.Println(v52)
                                                                            if (255 != 42) {
                                                                                v134 := GenS9{F0: 2147483647, F1: -528}
                                                                                _ = len(v6)
                                                                                if true {
                                                                                    var v135 uint32 = 42
                                                                                    v55 = (fi49 / ((fi49) | 1))
                                                                                    switch -1 {
                                                                                    case 749:
                                                                                        v2 = (v2 % ((v2) | 1))
                                                                                    case 9223372036854775807:
                                                                                        _ = len(v5)
                                                                                        if false {
                                                                                            var v136 float32 = 1.5
                                                                                            if false {
                                                                                                var v137 float64 = (v1 - v1)
                                                                                                if (!(false && true)) {
                                                                                                    switch (fi49 % ((-2147483648) | 1)) {
                                                                                                    case -9223372036854775808:
                                                                                                        v1 = (((float64((v136 * float32(v1))) / (3.14159 / ((v137 * (v1 * v1)) / 1.0))) * (((float64((((((v136 * v136) * float32(v1)) - (3.14159 - float32((v137 / 0.1)))) / (v136 * (-2.25 - v136))) - 3.14159)) - (v1 - 1e10)) * (1e10 + v137)) / 1.0)) * (v1 + v1))
                                                                                                    case 0:
                                                                                                        var v138 uint32 = v135
                                                                                                        _ = v138
                                                                                                    default:
                                                                                                        _ = len(v3)
                                                                                                    }
                                                                                                    if (!false) {
                                                                                                        var v139 uint8 = 2
                                                                                                        fmt.Println(v137, v2)
                                                                                                        v7 = "a"
                                                                                                        _ = v139
                                                                                                    } else {
                                                                                                        v140 := []int{(fi49 ^ (fi49 << ((578) & 63)))}
                                                                                                        v31 = "\n\t"
                                                                                                        _ = v140
                                                                                                    }
                                                                                                    v141 := v6[105]
                                                                                                    _ = v141
                                                                                                }
                                                                                                if false {
                                                                                                    switch v31 {
                                                                                                    case "\n\t":
                                                                                                        var v142 int8 = v51
                                                                                                        _ = v142
                                                                                                    default:
                                                                                                        v143 := GenS17{F0: 175, F1: 255, F2: (v51 & v4)}
                                                                                                        _ = v143
                                                                                                    }
                                                                                                    if false {
                                                                                                        v144 := v5[2]
                                                                                                        v145 := v5[v135]
                                                                                                        var v146 int = fi50
                                                                                                        _ = v144
                                                                                                        _ = v145
                                                                                                        _ = v146
                                                                                                    } else {
                                                                                                        if false {
                                                                                                            var v147 int8 = v120
                                                                                                            if ((255 > 981) || true) {
                                                                                                                fmt.Println(v137, v51, v4)
                                                                                                                var v148 uint16 = 43
                                                                                                                v135 = v135
                                                                                                                var v149 uint16 = 350
                                                                                                                _ = v148
                                                                                                                _ = v149
                                                                                                            } else {
                                                                                                                v51 = int8(v2)
                                                                                                                _ = len(v3)
                                                                                                                _ = len(v3)
                                                                                                            }
                                                                                                            _ = len(v3)
                                                                                                            var v150 uint64 = 4294967295
                                                                                                            if (!((!false) && (!(2 <= 127)))) {
                                                                                                                v151 := v3[v135]
                                                                                                                _ = v151
                                                                                                            } else {
                                                                                                                var v152 uint8 = 146
                                                                                                                if (true && (!true)) {
                                                                                                                    v120 += v120
                                                                                                                    _ = len(v6)
                                                                                                                }
                                                                                                                _ = v152
                                                                                                            }
                                                                                                            _ = v147
                                                                                                            _ = v150
                                                                                                        }
                                                                                                    }
                                                                                                    v4 += (v120 % ((v120) | 1))
                                                                                                    v153 := GenS10{F0: (v31 >= v7), F1: 220}
                                                                                                    var v154 uint64 = 42
                                                                                                    _ = v153
                                                                                                    _ = v154
                                                                                                } else {
                                                                                                    v4 -= (v51 >> ((42) & 63))
                                                                                                }
                                                                                                v155 := []int{fi49, fi50}
                                                                                                v7 = v31
                                                                                                _ = v137
                                                                                                _ = v155
                                                                                            } else {
                                                                                                v4 -= v51
                                                                                                v156 := v6[0]
                                                                                                _ = v156
                                                                                            }
                                                                                            _ = len(v3)
                                                                                            _ = v136
                                                                                        }
                                                                                    case 2147483647:
                                                                                        v31 = v31
                                                                                    default:
                                                                                        var v157 uint = 4294967295
                                                                                        _ = v157
                                                                                    }
                                                                                    delete(v5, v135)
                                                                                    var v158 int8 = (v4 * (v4 & v51))
                                                                                    _ = v135
                                                                                    _ = v158
                                                                                } else {
                                                                                    switch -702 {
                                                                                    case 127:
                                                                                        if (!false) {
                                                                                            switch v7 {
                                                                                            case "":
                                                                                                fmt.Println(v1, v52, v1)
                                                                                            case "foo bar":
                                                                                                fmt.Println(fi50, fi49)
                                                                                            }
                                                                                        } else {
                                                                                            v159 := v3[0]
                                                                                            _ = v159
                                                                                        }
                                                                                    }
                                                                                    v120 = 127
                                                                                    var v160 int64 = -636
                                                                                    _ = v160
                                                                                }
                                                                                _ = v134
                                                                            } else {
                                                                                if false {
                                                                                    switch v31 {
                                                                                    case "":
                                                                                        v55 -= fi49
                                                                                    case "\n\t":
                                                                                        v161 := []int{fi49, (^(fi49)), (v55 - (v55 | fi50))}
                                                                                        _ = v161
                                                                                    case "foo bar":
                                                                                        var v162 float32 = 3.14159
                                                                                        _ = v162
                                                                                    }
                                                                                    if false {
                                                                                        fmt.Println(v120)
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                    default:
                                                                        var v163 float64 = 1e10
                                                                        _ = v163
                                                                    }
                                                                    switch v7 {
                                                                    case "a":
                                                                        switch (v7 + "\n\t") {
                                                                        case "\n\t":
                                                                            v51 = (-(v4))
                                                                        }
                                                                    case "":
                                                                        v1 = (float64(0.0) * 1e10)
                                                                    default:
                                                                        v31 = (v7 + ("foo bar" + ("hello" + ("\n\t" + (v31 + ("a" + ""))))))
                                                                    }
                                                                    var v164 int = (fi49 >> ((255) & 63))
                                                                    _ = v164
                                                                }
                                                                v165 := []int{}
                                                                _ = v123
                                                                _ = v165
                                                            }
                                                        }
                                                        var v166 uint = 45
                                                        v167 := GenS22{F0: 42}
                                                        switch (v55 / ((v55) | 1)) {
                                                        case 32768:
                                                            v55 -= (fi50 & fi49)
                                                        case 127:
                                                            fmt.Println(v120, v31, v31)
                                                        case 335:
                                                            v168 := GenS20{F0: "\n\t"}
                                                            var v169 uint64 = 886
                                                            v55 = fi49
                                                            _ = v168
                                                            _ = v169
                                                        default:
                                                            fmt.Println(v4, v31, v51)
                                                            v170 := []int{v55}
                                                            v171 := GenS16{F0: (v166 | (^(v166)))}
                                                            _ = v170
                                                            _ = v171
                                                        }
                                                        _ = v119
                                                        _ = v120
                                                        _ = v166
                                                        _ = v167
                                                    }
                                                    if (2 < 42) {
                                                        delete(v3, 1)
                                                    } else {
                                                        switch v31 {
                                                        case "":
                                                            var v172 uint16 = 630
                                                            _ = v172
                                                        default:
                                                            var v173 int8 = (v4 | v51)
                                                            _ = v173
                                                        }
                                                    }
                                                }
                                                var v174 bool = true
                                                switch -467 {
                                                case -1:
                                                    switch "\n\t" {
                                                    case "":
                                                        switch (v55 % (((-(fi50))) | 1)) {
                                                        case -559:
                                                            v31 = "\n\t"
                                                        }
                                                    case "a":
                                                        fmt.Println(v1)
                                                    case "hello":
                                                        fmt.Println(fi49, v4, fi50)
                                                    }
                                                default:
                                                    v175 := []int{(fi49 << ((140) & 63)), fi49}
                                                    var v176 int32 = v52
                                                    var v177 float64 = v1
                                                    var v178 float64 = 0.0
                                                    _ = v175
                                                    _ = v176
                                                    _ = v177
                                                    _ = v178
                                                }
                                                _ = v174
                                            } else {
                                                switch v31 {
                                                case "":
                                                    v179 := v5[0]
                                                    _ = v179
                                                case "a":
                                                    v180 := GenS23{F0: 683, F1: v2, F2: 502}
                                                    switch v7 {
                                                    case "\n\t":
                                                        v181 := []int{-2147483648, 32768, v55, (v55 / ((374) | 1))}
                                                        _ = v181
                                                    default:
                                                        var v182 uint16 = 42
                                                        _ = v182
                                                    }
                                                    if (1000 <= 32767) {
                                                        fmt.Println(v2, v4)
                                                        switch (^(v55)) {
                                                        case 32767:
                                                            v183 := []int{fi50, fi50, (fi50 ^ 386), v55}
                                                            _ = v183
                                                        case 0:
                                                            var v184 string = v31
                                                            _ = v184
                                                        }
                                                    }
                                                    _ = v180
                                                case "hello":
                                                    var v185 int64 = 1
                                                    _ = v185
                                                }
                                                v2 -= v2
                                            }
                                        } else {
                                            var v186 uint = 740
                                            v187 := v3[65535]
                                            var v188 uint64 = 2
                                            _ = v186
                                            _ = v187
                                            _ = v188
                                        }
                                        switch (v7 + (v7 + v7)) {
                                        case "a":
                                            delete(v3, 1)
                                        case "":
                                            fmt.Println(v52)
                                        }
                                    } else {
                                        v52 -= v52
                                        fmt.Println(v2)
                                        var v189 float32 = float32(v1)
                                        _ = v189
                                    }
                                    _ = v55
                                }
                                _ = v52
                            }
                        case -2147483648:
                            var v190 int32 = 278
                            _ = v190
                        case 0:
                            v6[42] = 2147483648
                        default:
                            var v191 int64 = v2
                            _ = v191
                        }
                    } else {
                        var v192 uint = 880
                        _ = v192
                    }
                    v2 += -736
                    _ = fi50
                    _ = v51
                }
                delete(v6, 4294967295)
                if true {
                    for fi193 := 0; fi193 < 4; fi193++ {
                        v194 := v5[804]
                        var v195 int8 = v4
                        v196 := GenS2{F0: -1, F1: 333, F2: "hello"}
                        _ = fi193
                        _ = v194
                        _ = v195
                        _ = v196
                    }
                    v197 := GenS9{F0: 42, F1: 1000}
                    _ = v197
                }
                _ = fi49
            }
            v31 = "\n\t"
        }
    } else {
        switch 1 {
        case 612:
            var v198 string = v7
            _ = v198
        case 9223372036854775807:
            switch "\n\t" {
            case "foo bar":
                v199 := GenS1{F0: 32767}
                _ = v199
            default:
                v1 = (v1 - 1.5)
            }
        case -29:
            var v200 int16 = 231
            _ = v200
        default:
            for fi201 := 0; fi201 < 11; fi201++ {
                if false {
                    switch (fi201 / (((fi201 & 128)) | 1)) {
                    case 2147483647:
                        var v202 uint64 = 42
                        _ = v202
                    case -231:
                        var v203 int64 = v2
                        _ = v203
                    default:
                        fmt.Println(v31, v7, v7)
                    }
                    v204 := GenS2{F0: -432, F1: 871, F2: v7}
                    _ = v204
                }
                _ = fi201
            }
        }
        _ = len(v6)
        v205 := []int{743, 0, 4294967295, 9223372036854775807}
        _ = v205
    }
    if (1 <= 664) {
        for fi206 := 0; fi206 < 8; fi206++ {
            fmt.Println(v7, v2, v31)
            for fi207 := 0; fi207 < 18; fi207++ {
                if (128 == 128) {
                    var v208 string = "\n\t"
                    _ = v208
                }
                _ = fi207
            }
            var v209 uint32 = 240
            _ = fi206
            _ = v209
        }
        for fi210 := 0; fi210 < 14; fi210++ {
            var v211 int64 = -1
            v212 := v3[65]
            _ = fi210
            _ = v211
            _ = v212
        }
    }
    if false {
        var v213 uint16 = 42
        _ = v213
    }
    v31 = v7
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
