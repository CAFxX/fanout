package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 int8 }
type GenS6 struct { F0 uint8; F1 uint; F2 bool }
type GenS7 struct { F0 int32; F1 int32 }
type GenS8 struct { F0 int; F1 uint32 }
type GenS9 struct { F0 string; F1 bool }
type GenS10 struct { F0 int64; F1 int64; F2 int32 }
type GenS11 struct { F0 int8; F1 uint16; F2 int64 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int8; F1 uint; F2 int8 }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 int; F1 int32 }
type GenS18 struct { F0 int }
func genFn0() int {
    var v6 uint32 = 42
    if true {
        v7 := GenS1{F0: 0, F1: (v6 > v6)}
        var v8 int16 = 962
        fmt.Println(v6, v8)
        _ = v7
        _ = v8
    } else {
        if false {
            v9 := GenS2{F0: 896, F1: 4294967295}
            v10 := GenS3{F0: 1}
            v6 -= (v6 + (v6 / (((v6 << ((415) & 63))) | 1)))
            _ = v10
            _ = v9
        } else {
            v6 = v6
            v6 = v6
        }
    }
    v6 += v6
    v6 += v6
    v11 := GenS3{F0: 20}
    _ = v11
    _ = v6
    return -1
}
func genFn1(p0 uint8) int32 {
    v12 := GenS1{F0: 0, F1: true}
    fmt.Println(p0)
    if false {
        var v13 bool = ((127 <= 127) || (-277 != genFn0()))
        var v14 uint = 1
        var v15 uint64 = 18446744073709551615
        _ = v13
        _ = v14
        _ = v15
    } else {
        if true {
            v16 := []int{-1, 42}
            v17 := GenS4{F0: 1}
            switch ("foo bar" + ("\n\t" + "foo bar")) {
            case "\n\t":
                p0 = p0
            default:
                v16 = append(v16, 32767)
            }
            _ = v16
            _ = v17
        } else {
            for fi18 := 0; fi18 < 2; fi18++ {
                p0 = p0
                switch 9223372036854775807 {
                case 2147483647:
                    if (fi18 >= genFn0()) {
                        for fi19 := 0; fi19 < 3; fi19++ {
                            fmt.Println(fi19, fi19, p0)
                            _ = fi19
                        }
                        if true {
                            if (!(-1 > -26)) {
                                var v20 int64 = -961
                                var v21 int16 = 0
                                var v22 int = fi18
                                _ = v20
                                _ = v21
                                _ = v22
                            } else {
                                v23 := GenS5{F0: 65535, F1: 0}
                                for fi24 := 0; fi24 < 2; fi24++ {
                                    v25 := []int{(fi24 << ((4294967295) & 63))}
                                    p0 = 58
                                    v26 := GenS6{F0: p0, F1: 97, F2: false}
                                    p0 = p0
                                    v27 := GenS4{F0: 127}
                                    _ = fi24
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                }
                                if (!(0 >= -128)) {
                                    for fi28 := 0; fi28 < 6; fi28++ {
                                        p0 = (p0 ^ (p0 << ((4294967295) & 63)))
                                        p0 -= p0
                                        switch (fi18 & genFn0()) {
                                        case -2147483648:
                                            if ("a" >= "hello") {
                                                v29 := []int{fi28}
                                                p0 = (^(p0))
                                                if false {
                                                    p0 = (^(p0))
                                                    var v30 int16 = 42
                                                    v31 := GenS1{F0: (fi18 | genFn0()), F1: (true || false)}
                                                    var v32 float64 = 3.14159
                                                    var v33 int16 = (v30 + (v30 * (v30 & v30)))
                                                    _ = v30
                                                    _ = v31
                                                    _ = v32
                                                    _ = v33
                                                } else {
                                                    switch "a" {
                                                    case "":
                                                        var v34 uint64 = 4294967295
                                                        _ = v34
                                                    case "\n\t":
                                                        p0 = 109
                                                        p0 = p0
                                                    default:
                                                        if len(v29) > 0 {
                                                            v29[0] = genFn0()
                                                        }
                                                    }
                                                    if (!(!(true && false))) {
                                                        var v35 bool = (false || false)
                                                        _ = v35
                                                    } else {
                                                        if (42 <= 583) {
                                                            var v36 int16 = -10
                                                            if (958 == 415) {
                                                                v37 := GenS7{F0: -2147483648, F1: 2147483647}
                                                                p0 += (^(p0))
                                                                _ = v37
                                                            } else {
                                                                v38 := GenS8{F0: (fi18 - (^(fi28))), F1: 1}
                                                                fmt.Println(p0)
                                                                _ = v38
                                                            }
                                                            var v39 uint16 = 255
                                                            _ = v36
                                                            _ = v39
                                                        }
                                                    }
                                                }
                                                _ = v29
                                            } else {
                                                v40 := []int{}
                                                if len(v40) >= 2 {
                                                    v40 = v40[1:len(v40)]
                                                }
                                                _ = v40
                                            }
                                            v41 := GenS3{F0: -1}
                                            v42 := GenS6{F0: (p0 << ((6) & 63)), F1: 255, F2: true}
                                            if (42 <= 4294967295) {
                                                v43 := []int{1, fi28, (^(fi28))}
                                                switch 0 {
                                                case 128:
                                                    switch "\n\t" {
                                                    case "hello":
                                                        var v44 uint16 = 65535
                                                        _ = v44
                                                    case "foo bar":
                                                        v45 := GenS9{F0: ("" + ("hello" + "")), F1: (-9223372036854775808 != 516)}
                                                        _ = v45
                                                    case "\n\t":
                                                        p0 -= 1
                                                    default:
                                                        p0 -= p0
                                                    }
                                                }
                                                _ = v43
                                            }
                                            _ = v41
                                            _ = v42
                                        case 32768:
                                            p0 = p0
                                        default:
                                            p0 -= (p0 + 2)
                                        }
                                        var v46 bool = true
                                        var v47 int32 = -218
                                        _ = fi28
                                        _ = v46
                                        _ = v47
                                    }
                                    v48 := []int{}
                                    _ = v48
                                } else {
                                    v49 := GenS2{F0: 207, F1: 65535}
                                    _ = v49
                                }
                                _ = v23
                            }
                            for fi50 := 0; fi50 < 18; fi50++ {
                                if false {
                                    var v51 int16 = -32768
                                    if false {
                                        p0 -= (^(p0))
                                    } else {
                                        if false {
                                            var v52 float64 = float64(((((1e10 - (0.0 - 0.1)) * (1.5 / 1.5)) * 1e10) + 3.14159))
                                            _ = v52
                                        }
                                    }
                                    _ = v51
                                }
                                _ = fi50
                            }
                            switch (-(fi18)) {
                            case -906:
                                p0 = p0
                            case -106:
                                if ((!false) && (-9223372036854775808 <= 0)) {
                                    p0 += (p0 ^ (p0 ^ p0))
                                    p0 += 0
                                    if (!false) {
                                        var v53 uint32 = 2147483648
                                        _ = v53
                                    }
                                }
                            default:
                                var v54 int64 = 42
                                _ = v54
                            }
                        } else {
                            var v55 uint64 = 42
                            _ = v55
                        }
                        if false {
                            if (42 != -753) {
                                switch (fi18 / ((fi18) | 1)) {
                                case 1:
                                    switch "" {
                                    case "\n\t":
                                        if true {
                                            p0 = (p0 << ((846) & 63))
                                        } else {
                                            if true {
                                                v56 := GenS4{F0: -81}
                                                var v57 uint16 = 0
                                                _ = v56
                                                _ = v57
                                            }
                                            fmt.Println(p0)
                                            var v58 int16 = -32768
                                            _ = v58
                                        }
                                    case "hello":
                                        var v59 uint = 255
                                        _ = v59
                                    case "a":
                                        var v60 uint32 = 96
                                        _ = v60
                                    default:
                                        for fi61 := 0; fi61 < 11; fi61++ {
                                            var v62 int8 = int8(p0)
                                            v62 = (v62 & v62)
                                            _ = fi61
                                            _ = v62
                                        }
                                    }
                                case 0:
                                    for fi63 := 0; fi63 < 20; fi63++ {
                                        switch "foo bar" {
                                        case "":
                                            v64 := GenS5{F0: uint32(p0), F1: 126}
                                            _ = v64
                                        case "a":
                                            p0 += p0
                                        default:
                                            p0 -= (^(p0))
                                        }
                                        if (2147483647 > -665) {
                                            var v65 int64 = 0
                                            _ = v65
                                        }
                                        var v66 uint32 = 326
                                        _ = fi63
                                        _ = v66
                                    }
                                default:
                                    var v67 uint = 2
                                    _ = v67
                                }
                                p0 = (p0 - 255)
                                var v68 uint8 = p0
                                _ = v68
                            }
                            switch ((("a" + "foo bar") + "hello") + "a") {
                            case "foo bar":
                                p0 -= p0
                            case "a":
                                var v69 float64 = 0.0
                                _ = v69
                            default:
                                p0 -= 2
                            }
                        } else {
                            v70 := []int{fi18, -415, (fi18 + genFn0()), 270}
                            switch fi18 {
                            case 32767:
                                if false {
                                    for fi71 := 0; fi71 < 6; fi71++ {
                                        fmt.Println(p0, fi18, fi71)
                                        p0 -= 105
                                        _ = fi71
                                    }
                                    var v72 float32 = 1e10
                                    _ = v72
                                } else {
                                    for fi73 := 0; fi73 < 20; fi73++ {
                                        switch -339 {
                                        case 42:
                                            v70 = append(v70, -37)
                                        case 0:
                                            if ((0.0 * (1.5 / 1.5)) < (0.1 / ((0.1 / 0.1) + 1.5))) {
                                                var v74 uint8 = (p0 >> ((65535) & 63))
                                                _ = v74
                                            }
                                        }
                                        if false {
                                            if len(v70) > 0 {
                                                v70[0] = fi18
                                            }
                                            switch (("a" + "foo bar") + "a") {
                                            case "\n\t":
                                                v70 = append(v70, genFn0())
                                            case "a":
                                                if (1 != 699) {
                                                    fmt.Println(fi73, fi73)
                                                    switch fi73 {
                                                    case 1:
                                                        p0 -= (p0 & p0)
                                                    case 9223372036854775807:
                                                        fmt.Println(p0, p0, fi18)
                                                    case 127:
                                                        p0 -= (^(p0))
                                                    default:
                                                        if len(v70) >= 2 {
                                                            v70 = v70[1:len(v70)]
                                                        }
                                                    }
                                                    var v75 uint = 42
                                                    _ = v75
                                                } else {
                                                    _ = len(v70)
                                                    var v76 uint16 = 42
                                                    _ = v76
                                                }
                                            case "foo bar":
                                                var v77 int32 = 42
                                                _ = v77
                                            default:
                                                v78 := GenS6{F0: (p0 ^ (p0 ^ p0)), F1: 65535, F2: false}
                                                _ = v78
                                            }
                                        } else {
                                            switch (fi73 & fi73) {
                                            case 672:
                                                if true {
                                                    v79 := GenS10{F0: -1, F1: -690, F2: 0}
                                                    _ = v79
                                                }
                                            case 1:
                                                if ((p0 + p0) > p0) {
                                                    switch "hello" {
                                                    case "hello":
                                                        if (p0 >= p0) {
                                                            if len(v70) >= 2 {
                                                                v70 = v70[1:len(v70)]
                                                            }
                                                        } else {
                                                            p0 -= p0
                                                        }
                                                    default:
                                                        v80 := GenS9{F0: "", F1: true}
                                                        _ = v80
                                                    }
                                                    switch (^(fi18)) {
                                                    case -2147483648:
                                                        v81 := GenS9{F0: "hello", F1: true}
                                                        var v82 uint32 = 65535
                                                        if false {
                                                            if (((0.1 - (0.1 + 1e10)) == 1.5) && (float64(3.14159) != (1.5 - float64((1.5 + ((float32(1.5) / 1e10) / 0.1)))))) {
                                                                var v83 float32 = ((0.0 * 1.5) * float32((((-2.25 / ((1e10 - ((1.5 - 1e10) - (-2.25 * 3.14159))) - 1.5)) / 0.1) - 1e10)))
                                                                p0 = 0
                                                                v83 = 1.5
                                                                _ = v83
                                                            }
                                                            if false {
                                                                var v84 float32 = (3.14159 + 0.0)
                                                                if true {
                                                                    v82 -= 1
                                                                } else {
                                                                    if false {
                                                                        var v85 int64 = 42
                                                                        _ = v85
                                                                    } else {
                                                                        p0 = p0
                                                                    }
                                                                }
                                                                var v86 bool = (991 != 1)
                                                                _ = v84
                                                                _ = v86
                                                            }
                                                            if false {
                                                                if true {
                                                                    if ((v82 * v82) <= v82) {
                                                                        if len(v70) > 0 {
                                                                            v70[0] = (fi73 >> ((4294967295) & 63))
                                                                        }
                                                                        switch "foo bar" {
                                                                        case "foo bar":
                                                                            switch fi73 {
                                                                            case -9223372036854775808:
                                                                                fmt.Println(fi73)
                                                                            case 2147483647:
                                                                                if true {
                                                                                    v87 := GenS3{F0: 104}
                                                                                    _ = v87
                                                                                }
                                                                            default:
                                                                                v88 := GenS2{F0: 1, F1: 4294967295}
                                                                                _ = v88
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                v89 := GenS11{F0: 100, F1: 803, F2: 1}
                                                                _ = v89
                                                            } else {
                                                                var v90 uint64 = 42
                                                                v82 += v82
                                                                var v91 float32 = float32((0.0 * 0.1))
                                                                _ = v90
                                                                _ = v91
                                                            }
                                                        } else {
                                                            var v92 int32 = 2147483647
                                                            var v93 int16 = 0
                                                            _ = v92
                                                            _ = v93
                                                        }
                                                        _ = v81
                                                        _ = v82
                                                    case -323:
                                                        fmt.Println(fi73)
                                                    case -115:
                                                        p0 = 1
                                                    default:
                                                        p0 = p0
                                                    }
                                                    fmt.Println(fi18, p0, p0)
                                                } else {
                                                    if true {
                                                        if (!(!false)) {
                                                            var v94 uint16 = 371
                                                            switch (fi18 ^ 42) {
                                                            case 135:
                                                                var v95 uint = uint(v94)
                                                                var v96 int32 = 42
                                                                v94 -= (v94 ^ v94)
                                                                _ = v95
                                                                _ = v96
                                                            case 42:
                                                                if (false && true) {
                                                                    v97 := GenS4{F0: -89}
                                                                    _ = v97
                                                                }
                                                                var v98 int16 = -285
                                                                if (v98 == v98) {
                                                                    v99 := GenS8{F0: (fi18 << ((4294967295) & 63)), F1: 32}
                                                                    v100 := GenS11{F0: -1, F1: (v94 % (((v94 | (v94 ^ (v94 | 255)))) | 1)), F2: 498}
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        switch "foo bar" {
                                                                        case "":
                                                                            var v101 uint8 = p0
                                                                            _ = v101
                                                                        default:
                                                                            var v102 int = (fi73 ^ fi18)
                                                                            _ = v102
                                                                        }
                                                                    }
                                                                    _ = v100
                                                                    _ = v99
                                                                }
                                                                _ = v98
                                                            case 9223372036854775807:
                                                                var v103 int64 = 1
                                                                _ = v103
                                                            }
                                                            _ = v94
                                                        }
                                                        if true {
                                                            p0 = p0
                                                            if len(v70) > 0 {
                                                                v70[0] = fi18
                                                            }
                                                            if false {
                                                                p0 = p0
                                                            }
                                                        }
                                                        fmt.Println(p0, fi18)
                                                        if len(v70) > 0 {
                                                            v70[0] = fi18
                                                        }
                                                    } else {
                                                        p0 += p0
                                                    }
                                                    fmt.Println(fi18, fi18, fi73)
                                                    var v104 int16 = 1000
                                                    _ = v104
                                                }
                                            case 32767:
                                                v105 := GenS7{F0: -803, F1: 42}
                                                _ = v105
                                            default:
                                                if (false && false) {
                                                    switch "foo bar" {
                                                    case "":
                                                        var v106 string = "foo bar"
                                                        _ = v106
                                                    case "foo bar":
                                                        var v107 int32 = int32(fi73)
                                                        _ = v107
                                                    case "\n\t":
                                                        if true {
                                                            v108 := GenS9{F0: "hello", F1: (true && (false && ((!false) && true)))}
                                                            p0 = 63
                                                            _ = v108
                                                        }
                                                    }
                                                    if false {
                                                        if (true || ("foo bar" < ("foo bar" + ""))) {
                                                            switch "" {
                                                            case "foo bar":
                                                                switch (fi73 >> ((255) & 63)) {
                                                                case 9223372036854775807:
                                                                    switch (fi18 / (((^(fi18))) | 1)) {
                                                                    case -9223372036854775808:
                                                                        var v109 string = "\n\t"
                                                                        _ = v109
                                                                    case -887:
                                                                        if true {
                                                                            var v110 float64 = 0.1
                                                                            switch -274 {
                                                                            case 32767:
                                                                                if false {
                                                                                    var v111 uint = 18446744073709551615
                                                                                    _ = v111
                                                                                }
                                                                                v110 = 3.14159
                                                                                if false {
                                                                                    fmt.Println(p0, fi73, fi73)
                                                                                } else {
                                                                                    if (!(!true)) {
                                                                                        v110 += ((-2.25 + v110) - 3.14159)
                                                                                        if (!(42 > -1)) {
                                                                                            var v112 int16 = 1
                                                                                            switch fi73 {
                                                                                            case 239:
                                                                                                v113 := GenS4{F0: -128}
                                                                                                _ = v113
                                                                                            case 32768:
                                                                                                var v114 int32 = -841
                                                                                                _ = v114
                                                                                            case 9223372036854775807:
                                                                                                if (1 > -212) {
                                                                                                    var v115 uint64 = 18446744073709551615
                                                                                                    _ = len(v70)
                                                                                                    _ = v115
                                                                                                } else {
                                                                                                    v70 = append(v70, fi73)
                                                                                                    if (true && false) {
                                                                                                        var v116 int16 = v112
                                                                                                        v117 := GenS8{F0: (fi18 / (((fi18 - (fi73 >> ((2) & 63)))) | 1)), F1: 2147483648}
                                                                                                        _ = v116
                                                                                                        _ = v117
                                                                                                    }
                                                                                                    v110 += (float64(1e10) - 0.0)
                                                                                                }
                                                                                            default:
                                                                                                _ = len(v70)
                                                                                            }
                                                                                            v110 = (((float64((0.0 + -2.25)) / 1e10) - 1e10) / v110)
                                                                                            _ = v112
                                                                                        } else {
                                                                                            v110 = v110
                                                                                            if len(v70) >= 2 {
                                                                                                v70 = v70[1:len(v70)]
                                                                                            }
                                                                                        }
                                                                                        p0 += (p0 / ((p0) | 1))
                                                                                    }
                                                                                    fmt.Println(fi73, fi73, p0)
                                                                                }
                                                                            case 572:
                                                                                v118 := GenS9{F0: (("hello" + ("foo bar" + ("a" + (("a" + "hello") + "")))) + "hello"), F1: (false && true)}
                                                                                _ = v118
                                                                            default:
                                                                                var v119 uint8 = (p0 | 255)
                                                                                v120 := GenS8{F0: 128, F1: 1}
                                                                                if len(v70) > 0 {
                                                                                    v70[0] = genFn0()
                                                                                }
                                                                                _ = v119
                                                                                _ = v120
                                                                            }
                                                                            _ = v110
                                                                        } else {
                                                                            p0 = p0
                                                                            var v121 uint32 = 2
                                                                            p0 = 1
                                                                            _ = v121
                                                                        }
                                                                    case 128:
                                                                        fmt.Println(fi73)
                                                                        var v122 int64 = -9223372036854775808
                                                                        _ = v122
                                                                    }
                                                                }
                                                            }
                                                            p0 = p0
                                                            if false {
                                                                if len(v70) >= 2 {
                                                                    v70 = v70[1:len(v70)]
                                                                }
                                                            }
                                                            var v123 int32 = 1
                                                            v124 := GenS6{F0: p0, F1: 4294967295, F2: ((^(p0)) != (^(p0)))}
                                                            _ = v123
                                                            _ = v124
                                                        } else {
                                                            var v125 uint8 = 255
                                                            v125 = (^(v125))
                                                            _ = v125
                                                        }
                                                        switch fi73 {
                                                        case 128:
                                                            if ((772 <= 13) && false) {
                                                                if true {
                                                                    fmt.Println(p0, fi73, p0)
                                                                    p0 += 42
                                                                    if true {
                                                                        v70 = append(v70, (fi73 + (fi73 + fi73)))
                                                                    }
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        fmt.Println(fi73)
                                                                    case "a":
                                                                        var v126 uint32 = 548
                                                                        _ = v126
                                                                    case "":
                                                                        p0 = p0
                                                                    default:
                                                                        switch (^(fi73)) {
                                                                        case 811:
                                                                            var v127 uint64 = 18446744073709551615
                                                                            _ = v127
                                                                        }
                                                                    }
                                                                } else {
                                                                    p0 = p0
                                                                    switch (fi73 * 9223372036854775807) {
                                                                    case 9223372036854775807:
                                                                        var v128 uint64 = 1
                                                                        _ = v128
                                                                    case 627:
                                                                        p0 = p0
                                                                        var v129 bool = (127 != p0)
                                                                        v129 = v129
                                                                        var v130 float32 = ((3.14159 - 0.0) * float32(0.0))
                                                                        _ = v129
                                                                        _ = v130
                                                                    case -625:
                                                                        switch ("hello" + "hello") {
                                                                        case "":
                                                                            p0 += p0
                                                                        }
                                                                    default:
                                                                        var v131 int64 = -9223372036854775808
                                                                        _ = v131
                                                                    }
                                                                }
                                                                var v132 int32 = 42
                                                                _ = v132
                                                            }
                                                        case 32768:
                                                            switch (("a" + "foo bar") + "foo bar") {
                                                            case "hello":
                                                                fmt.Println(fi73)
                                                            }
                                                        }
                                                        fmt.Println(p0, fi18)
                                                    } else {
                                                        p0 = p0
                                                        if len(v70) >= 2 {
                                                            v70 = v70[1:len(v70)]
                                                        }
                                                    }
                                                } else {
                                                    p0 = (p0 / ((p0) | 1))
                                                    fmt.Println(p0)
                                                    if true {
                                                        var v133 int16 = 32767
                                                        var v134 uint64 = 18446744073709551615
                                                        _ = len(v70)
                                                        _ = v133
                                                        _ = v134
                                                    } else {
                                                        p0 = p0
                                                    }
                                                }
                                                p0 -= 36
                                            }
                                            var v135 uint16 = 2
                                            v136 := GenS7{F0: 1, F1: -2147483648}
                                            v70 = append(v70, genFn0())
                                            var v137 int32 = 230
                                            var v138 bool = (false || false)
                                            if len(v70) > 0 {
                                                v70[0] = fi18
                                            }
                                            var v139 uint = 0
                                            v70 = append(v70, fi73)
                                            _ = v135
                                            _ = v136
                                            _ = v137
                                            _ = v138
                                            _ = v139
                                        }
                                        if true {
                                            p0 = 255
                                            if len(v70) > 0 {
                                                v70[0] = fi73
                                            }
                                        } else {
                                            p0 = 1
                                            p0 += p0
                                            var v140 string = ("hello" + "a")
                                            _ = v140
                                        }
                                        _ = fi73
                                    }
                                    p0 += p0
                                    v141 := GenS12{F0: 122}
                                    _ = v141
                                }
                            case 42:
                                for fi142 := 0; fi142 < 16; fi142++ {
                                    v143 := GenS12{F0: 1}
                                    var v144 uint64 = 1
                                    fmt.Println(fi142, fi142)
                                    _ = fi142
                                    _ = v143
                                    _ = v144
                                }
                            }
                            _ = v70
                        }
                    } else {
                        for fi145 := 0; fi145 < 1; fi145++ {
                            if (false && true) {
                                var v146 float64 = (-2.25 - 3.14159)
                                if true {
                                    v147 := GenS13{F0: 7}
                                    _ = v147
                                } else {
                                    v148 := []int{(fi145 - fi145), fi145, fi18}
                                    fmt.Println(fi145, p0)
                                    p0 = (p0 << ((787) & 63))
                                    _ = v148
                                }
                                p0 = p0
                                _ = v146
                            } else {
                                p0 -= p0
                                fmt.Println(fi145, fi145, p0)
                            }
                            p0 = 128
                            if (629 <= 1) {
                                v149 := []int{genFn0()}
                                var v150 float64 = 0.0
                                if len(v149) >= 2 {
                                    v149 = v149[1:len(v149)]
                                }
                                v151 := GenS3{F0: -128}
                                var v152 uint = 4294967295
                                v150 -= v150
                                _ = v149
                                _ = v150
                                _ = v151
                                _ = v152
                            } else {
                                if ((p0 * p0) > p0) {
                                    var v153 uint8 = p0
                                    p0 = 237
                                    if false {
                                        p0 = v153
                                        var v154 int16 = 32767
                                        _ = v154
                                    } else {
                                        if (659 <= 1) {
                                            var v155 uint = 18446744073709551615
                                            var v156 int64 = 42
                                            v157 := GenS14{F0: 32767}
                                            _ = v155
                                            _ = v156
                                            _ = v157
                                        } else {
                                            switch "a" {
                                            case "hello":
                                                if false {
                                                    var v158 bool = false
                                                    var v159 float64 = 0.1
                                                    _ = v158
                                                    _ = v159
                                                } else {
                                                    fmt.Println(v153)
                                                    fmt.Println(fi145, v153)
                                                    v153 -= (p0 / ((1) | 1))
                                                }
                                            case "a":
                                                v160 := GenS9{F0: "a", F1: (-566 <= 1)}
                                                _ = v160
                                            case "\n\t":
                                                if false {
                                                    p0 -= (^(p0))
                                                    var v161 string = "hello"
                                                    _ = v161
                                                } else {
                                                    if true {
                                                        v162 := GenS12{F0: 0}
                                                        v163 := []int{(fi145 << ((255) & 63)), 127, 599, genFn0()}
                                                        _ = v162
                                                        _ = v163
                                                    }
                                                    if (!(0.0 > ((0.0 * 3.14159) / 1e10))) {
                                                        v164 := []int{genFn0(), fi145}
                                                        _ = v164
                                                    } else {
                                                        v165 := []int{(fi18 * (fi145 ^ (fi145 / ((fi18) | 1)))), fi145, (^(fi145))}
                                                        if (!(638 <= 2)) {
                                                            p0 = v153
                                                            switch 127 {
                                                            case -1:
                                                                if (-18 < -30) {
                                                                    var v166 uint32 = 961
                                                                    fmt.Println(p0)
                                                                    _ = v166
                                                                }
                                                            case -569:
                                                                v165 = append(v165, (fi145 + (fi145 * fi18)))
                                                            default:
                                                                v153 = 241
                                                            }
                                                            var v167 bool = true
                                                            _ = v167
                                                        } else {
                                                            p0 = p0
                                                        }
                                                        if ((((float64(-2.25) / 3.14159) - 1.5) - ((3.14159 - ((0.1 + -2.25) - 3.14159)) * float64(0.0))) != 1.5) {
                                                            if false {
                                                                p0 -= (v153 % ((42) | 1))
                                                            }
                                                        } else {
                                                            v168 := GenS3{F0: 42}
                                                            if (32767 < 1000) {
                                                                if len(v165) >= 2 {
                                                                    v165 = v165[1:len(v165)]
                                                                }
                                                                v169 := GenS10{F0: 0, F1: int64(v153), F2: 0}
                                                                switch (("" + "") + "foo bar") {
                                                                case "a":
                                                                    v153 += 2
                                                                case "\n\t":
                                                                    p0 += 255
                                                                }
                                                                _ = v169
                                                            } else {
                                                                v170 := GenS4{F0: 1}
                                                                _ = v170
                                                            }
                                                            _ = v168
                                                        }
                                                        if true {
                                                            _ = len(v165)
                                                            v153 += (p0 + 2)
                                                        } else {
                                                            p0 = v153
                                                            var v171 uint32 = 4294967295
                                                            _ = v171
                                                        }
                                                        var v172 uint8 = p0
                                                        _ = v165
                                                        _ = v172
                                                    }
                                                    v173 := GenS9{F0: "", F1: true}
                                                    _ = v173
                                                }
                                            }
                                        }
                                        v153 = (^(p0))
                                        v153 += (v153 - 1)
                                    }
                                    var v174 int16 = 0
                                    _ = v153
                                    _ = v174
                                } else {
                                    fmt.Println(fi18, p0, fi18)
                                    if (-128 == 1) {
                                        p0 -= p0
                                        var v175 uint16 = uint16(p0)
                                        _ = v175
                                    }
                                    v176 := GenS15{F0: 100, F1: 65535, F2: 42}
                                    fmt.Println(fi145)
                                    if false {
                                        var v177 int8 = -128
                                        _ = v177
                                    }
                                    _ = v176
                                }
                                fmt.Println(fi145)
                                v178 := []int{}
                                _ = v178
                            }
                            p0 = 255
                            _ = fi145
                        }
                        v179 := GenS9{F0: "foo bar", F1: (0 != 229)}
                        if true {
                            switch ("\n\t" + "foo bar") {
                            case "\n\t":
                                var v180 int16 = 0
                                _ = v180
                            case "a":
                                v181 := []int{fi18}
                                _ = v181
                            default:
                                fmt.Println(p0, fi18)
                            }
                        } else {
                            p0 = 1
                        }
                        _ = v179
                    }
                case 32767:
                    var v182 int32 = -2147483648
                    _ = v182
                default:
                    fmt.Println(fi18)
                }
                _ = fi18
            }
            p0 = p0
        }
    }
    _ = p0
    _ = v12
    return -2147483648
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{4294967295, -425}
    v3 := make(map[string]string)
    v4 := make(map[int]uint32)
    var v5 int8 = 85
    if (1 > 255) {
        if false {
            v5 -= v5
        } else {
            v5 -= 100
            for fi183 := 0; fi183 < 10; fi183++ {
                v184 := GenS16{F0: -1}
                for fi185 := 0; fi185 < 5; fi185++ {
                    v2 = append(v2, (fi183 - (fi185 >> ((255) & 63))))
                    fmt.Println(fi183, v5)
                    _ = fi185
                }
                _ = fi183
                _ = v184
            }
            for fi186 := 0; fi186 < 4; fi186++ {
                v187 := v4[(-(fi186))]
                _ = fi186
                _ = v187
            }
        }
    } else {
        v188 := GenS17{F0: 32768, F1: -2147483648}
        _ = v188
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi189 := 0; fi189 < 10; fi189++ {
        _ = len(v2)
        var v190 uint = 42
        v190 += 247
        _ = fi189
        _ = v190
    }
    var v191 int8 = -76
    var v192 bool = (!false)
    var v193 uint16 = 42
    switch "a" {
    case "":
        var v194 uint64 = 840
        _ = v194
    default:
        v5 = 123
    }
    v195 := GenS18{F0: -877}
    var v196 int32 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v191:", v191)
    fmt.Println("v192:", v192)
    fmt.Println("v193:", v193)
    fmt.Println("v196:", v196)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v191
    _ = v192
    _ = v193
    _ = v195
    _ = v196
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
