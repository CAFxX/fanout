package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int8 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int64; F1 bool; F2 uint16 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint8; F1 uint8; F2 int16 }
type GenS10 struct { F0 string; F1 int64; F2 uint64 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int; F1 int64; F2 int8 }
type GenS13 struct { F0 uint16; F1 int16; F2 int64 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 int; F1 int16 }
type GenS16 struct { F0 uint32; F1 int64; F2 int64 }
type GenS17 struct { F0 uint; F1 uint8 }
func genFn0(p0 string, p1 int) int {
    v6 := []int{p1}
    p0 = "foo bar"
    if false {
        for fi7 := 0; fi7 < 17; fi7++ {
            switch ((p0 + (p0 + p0)) + "foo bar") {
            case "hello":
                fmt.Println(p0, fi7, p0)
            default:
                fmt.Println(p1, p1)
            }
            fmt.Println(fi7, p0, p1)
            for fi8 := 0; fi8 < 5; fi8++ {
                p0 = (("hello" + p0) + "a")
                p1 -= 21
                _ = fi8
            }
            _ = fi7
        }
        v9 := GenS1{F0: 18446744073709551615, F1: ("hello" + p0), F2: 0}
        _ = v9
    } else {
        var v10 uint16 = 42
        var v11 bool = true
        var v12 int64 = 17
        _ = v10
        _ = v11
        _ = v12
    }
    var v13 string = p0
    var v14 uint16 = 14
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v6
    return (p1 & p1)
}
func genFn1(p0 uint8, p1 uint) uint16 {
    if false {
        var v15 string = "hello"
        var v16 float64 = float64(float32(3.14159))
        _ = v15
        _ = v16
    }
    fmt.Println(p0)
    if false {
        v17 := []int{-9223372036854775808, 762}
        _ = v17
    } else {
        v18 := []int{150, 127, 32768, -437}
        fmt.Println(p0)
        for fi19 := 0; fi19 < 13; fi19++ {
            switch (("\n\t" + ((("\n\t" + "foo bar") + "a") + "hello")) + "") {
            case "\n\t":
                for fi20 := 0; fi20 < 3; fi20++ {
                    var v21 uint32 = 2147483647
                    v22 := GenS1{F0: 614, F1: "a", F2: -36}
                    if len(v18) > 0 {
                        v18[0] = (fi20 - (fi19 - 1))
                    }
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            }
            p1 -= (p1 * (p1 >> ((p1) & 63)))
            _ = fi19
        }
        _ = v18
    }
    var v23 uint = p1
    if false {
        for fi24 := 0; fi24 < 17; fi24++ {
            switch fi24 {
            case 2147483647:
                var v25 int64 = 9223372036854775807
                _ = v25
            case 128:
                fmt.Println(p0)
            default:
                p1 -= (^(v23))
            }
            switch fi24 {
            case -1:
                if true {
                    switch ("foo bar" + "\n\t") {
                    case "":
                        fmt.Println(p1, p0, fi24)
                    case "\n\t":
                        if (int32(p0) != 42) {
                            var v26 int8 = 1
                            _ = v26
                        }
                    case "foo bar":
                        switch ("foo bar" + "\n\t") {
                        case "hello":
                            p1 -= (v23 - v23)
                        default:
                            v27 := []int{fi24, (fi24 * genFn0(("foo bar" + ("a" + ("" + ("a" + ("hello" + "foo bar"))))), fi24))}
                            _ = v27
                        }
                    }
                    v28 := []int{fi24}
                    for fi29 := 0; fi29 < 2; fi29++ {
                        var v30 bool = false
                        if false {
                            if v30 {
                                v23 += (p1 ^ (p1 | (^(p1))))
                                v31 := GenS1{F0: 4294967295, F1: "\n\t", F2: 1}
                                if v30 {
                                    v32 := GenS1{F0: 312, F1: ("" + (("foo bar" + "a") + "a")), F2: -128}
                                    p1 -= v23
                                    _ = v32
                                }
                                if (!false) {
                                    if (1 == -128) {
                                        var v33 int16 = 886
                                        fmt.Println(v23)
                                        _ = v33
                                    }
                                } else {
                                    if (("a" + ("foo bar" + "foo bar")) < ("" + "foo bar")) {
                                        var v34 string = ("foo bar" + "foo bar")
                                        var v35 int8 = 1
                                        _ = v34
                                        _ = v35
                                    } else {
                                        v23 = 42
                                        var v36 uint16 = 2
                                        fmt.Println(v30, p0)
                                        _ = v36
                                    }
                                }
                                _ = v31
                            }
                            fmt.Println(fi24)
                            switch fi24 {
                            case 128:
                                var v37 uint16 = 1
                                _ = v37
                            default:
                                fmt.Println(p1)
                            }
                        }
                        var v38 float64 = float64((0.1 + 0.0))
                        _ = fi29
                        _ = v30
                        _ = v38
                    }
                    _ = v28
                } else {
                    var v39 int8 = 127
                    var v40 uint = (p1 << (((p1 | 1)) & 63))
                    v39 = (^(v39))
                    _ = v39
                    _ = v40
                }
            }
            if (true || (!(42 == 42))) {
                var v41 bool = false
                _ = v41
            } else {
                if false {
                    v23 -= (^(p1))
                }
            }
            _ = fi24
        }
        if false {
            switch 0 {
            case 346:
                var v42 uint = p1
                _ = v42
            case 220:
                var v43 int8 = 7
                _ = v43
            case 9223372036854775807:
                if true {
                    switch "a" {
                    case "\n\t":
                        var v44 uint = 42
                        _ = v44
                    default:
                        if (2147483647 > -2147483648) {
                            var v45 bool = true
                            _ = v45
                        }
                    }
                    var v46 int32 = 1
                    fmt.Println(v23, p0, v23)
                    _ = v46
                } else {
                    if false {
                        if (!false) {
                            fmt.Println(v23, p0, v23)
                            var v47 bool = false
                            p1 -= (v23 & (v23 >> ((p1) & 63)))
                            _ = v47
                        }
                        for fi48 := 0; fi48 < 8; fi48++ {
                            v49 := GenS2{F0: -1, F1: 32767}
                            _ = fi48
                            _ = v49
                        }
                        if false {
                            v50 := GenS1{F0: 18446744073709551615, F1: "\n\t", F2: -1}
                            var v51 int8 = 42
                            _ = v50
                            _ = v51
                        }
                    }
                }
            }
            var v52 string = "foo bar"
            v53 := GenS2{F0: -2147483648, F1: 42}
            _ = v52
            _ = v53
        } else {
            if (true || (1000 > 42)) {
                p0 = p0
                var v54 int64 = 0
                v23 = (v23 >> ((771) & 63))
                _ = v54
            } else {
                var v55 int16 = 32767
                fmt.Println(p1)
                v56 := GenS2{F0: -295, F1: 972}
                _ = v55
                _ = v56
            }
            for fi57 := 0; fi57 < 16; fi57++ {
                for fi58 := 0; fi58 < 17; fi58++ {
                    var v59 uint8 = 6
                    v60 := []int{(fi57 & fi57), (fi57 ^ fi57)}
                    fmt.Println(fi58, v59)
                    _ = fi58
                    _ = v59
                    _ = v60
                }
                if (float32(((0.0 + (1.5 * 0.1)) / 3.14159)) == 0.0) {
                    for fi61 := 0; fi61 < 12; fi61++ {
                        var v62 int32 = 1
                        fmt.Println(fi57, p0)
                        if false {
                            p0 = p0
                            p0 += 255
                        } else {
                            if (fi57 <= 298) {
                                var v63 uint8 = 209
                                if ("a" == ((("hello" + ("hello" + "foo bar")) + ("a" + "\n\t")) + "foo bar")) {
                                    v64 := GenS2{F0: int32(fi61), F1: 1}
                                    var v65 int8 = 105
                                    v65 = v65
                                    var v66 int64 = 42
                                    if false {
                                        v65 -= 127
                                        p1 = 823
                                        if (42 != 2) {
                                            v67 := []int{int(p0)}
                                            _ = v67
                                        }
                                    }
                                    _ = v64
                                    _ = v65
                                    _ = v66
                                }
                                _ = v63
                            } else {
                                switch (-(fi61)) {
                                case 0:
                                    v68 := GenS1{F0: 4294967295, F1: "", F2: 127}
                                    _ = v68
                                case 4294967295:
                                    if (18446744073709551615 >= 2) {
                                        switch ("a" + ("hello" + ("hello" + ""))) {
                                        case "":
                                            if false {
                                                if (false || ((fi61 ^ (fi61 << ((1) & 63))) == -421)) {
                                                    if (!(!true)) {
                                                        var v69 float64 = (1e10 / float64(0.1))
                                                        p1 -= 0
                                                        _ = v69
                                                    }
                                                    p0 -= 0
                                                    v62 = v62
                                                }
                                                v70 := GenS1{F0: 42, F1: ("a" + "a"), F2: 127}
                                                var v71 float64 = 1e10
                                                switch (fi61 / ((32767) | 1)) {
                                                case -426:
                                                    if false {
                                                        switch fi57 {
                                                        case 4294967295:
                                                            switch ("\n\t" + "foo bar") {
                                                            case "hello":
                                                                if true {
                                                                    if true {
                                                                        p0 = (^(p0))
                                                                        if false {
                                                                            if false {
                                                                                if true {
                                                                                    p1 += 2
                                                                                }
                                                                            } else {
                                                                                var v72 uint32 = 4294967295
                                                                                _ = v72
                                                                            }
                                                                        }
                                                                        v71 -= v71
                                                                        fmt.Println(v71, fi57)
                                                                    } else {
                                                                        v73 := GenS2{F0: (^(v62)), F1: 255}
                                                                        fmt.Println(v23, fi57, v23)
                                                                        _ = v73
                                                                    }
                                                                    v71 = 3.14159
                                                                    var v74 uint8 = (p0 % (((p0 + p0)) | 1))
                                                                    if true {
                                                                        switch fi61 {
                                                                        case 127:
                                                                            if (!false) {
                                                                                var v75 float32 = 1e10
                                                                                _ = v75
                                                                            } else {
                                                                                if (32767 == -287) {
                                                                                    var v76 uint16 = 164
                                                                                    _ = v76
                                                                                } else {
                                                                                    if (v71 != (v71 * v71)) {
                                                                                        if false {
                                                                                            v77 := GenS2{F0: v62, F1: uint16(v23)}
                                                                                            if (!(fi57 < fi57)) {
                                                                                                v78 := GenS3{F0: 360, F1: 286, F2: 2}
                                                                                                _ = v78
                                                                                            } else {
                                                                                                v74 -= p0
                                                                                                switch ("a" + ("hello" + "\n\t")) {
                                                                                                case "":
                                                                                                    v79 := GenS1{F0: 230, F1: ((("hello" + ((("\n\t" + "\n\t") + ("\n\t" + ("hello" + (("foo bar" + "") + "hello")))) + (("foo bar" + (("" + "hello") + "foo bar")) + ("a" + "a")))) + "hello") + ("hello" + (("\n\t" + "") + "hello"))), F2: -5}
                                                                                                    _ = v79
                                                                                                case "hello":
                                                                                                    var v80 int32 = v62
                                                                                                    if false {
                                                                                                        v81 := []int{fi57}
                                                                                                        var v82 uint = (p1 - v23)
                                                                                                        if false {
                                                                                                            v81 = append(v81, int(v82))
                                                                                                            v71 -= -2.25
                                                                                                        } else {
                                                                                                            var v83 string = ""
                                                                                                            if true {
                                                                                                                v84 := GenS4{F0: v74, F1: 0}
                                                                                                                if (507 >= uint16(v82)) {
                                                                                                                    if (float32((v71 / (0.1 / v71))) >= float32(((0.1 - float64(((1e10 * 1.5) / 1.5))) / v71))) {
                                                                                                                        v85 := GenS5{F0: (("" + v83) != v83), F1: 42}
                                                                                                                        if len(v81) >= 2 {
                                                                                                                            v81 = v81[1:len(v81)]
                                                                                                                        }
                                                                                                                        var v86 uint32 = 0
                                                                                                                        _ = v85
                                                                                                                        _ = v86
                                                                                                                    }
                                                                                                                    fmt.Println(v74, v62)
                                                                                                                    if false {
                                                                                                                        v71 -= v71
                                                                                                                        if true {
                                                                                                                            v81 = append(v81, genFn0(("hello" + v83), 127))
                                                                                                                            var v87 int16 = -1
                                                                                                                            _ = v87
                                                                                                                        }
                                                                                                                        var v88 bool = false
                                                                                                                        _ = v88
                                                                                                                    } else {
                                                                                                                        fmt.Println(v82, v74)
                                                                                                                    }
                                                                                                                }
                                                                                                                v89 := GenS2{F0: v62, F1: 418}
                                                                                                                _ = v84
                                                                                                                _ = v89
                                                                                                            }
                                                                                                            _ = v83
                                                                                                        }
                                                                                                        fmt.Println(v62)
                                                                                                        var v90 int64 = 42
                                                                                                        _ = v81
                                                                                                        _ = v82
                                                                                                        _ = v90
                                                                                                    } else {
                                                                                                        v91 := GenS6{F0: (fi57 & (fi57 & fi57))}
                                                                                                        var v92 uint8 = (p0 << ((v23) & 63))
                                                                                                        fmt.Println(v62, v74, v62)
                                                                                                        _ = v91
                                                                                                        _ = v92
                                                                                                    }
                                                                                                    var v93 string = (("\n\t" + "") + (("hello" + ("hello" + "a")) + "a"))
                                                                                                    _ = v80
                                                                                                    _ = v93
                                                                                                }
                                                                                                switch ("" + "\n\t") {
                                                                                                case "a":
                                                                                                    var v94 int16 = 321
                                                                                                    _ = v94
                                                                                                case "foo bar":
                                                                                                    switch "hello" {
                                                                                                    case "\n\t":
                                                                                                        fmt.Println(p0)
                                                                                                    default:
                                                                                                        var v95 int = genFn0("", -2147483648)
                                                                                                        _ = v95
                                                                                                    }
                                                                                                case "hello":
                                                                                                    v96 := GenS7{F0: 9223372036854775807, F1: (!true), F2: 369}
                                                                                                    if true {
                                                                                                        p0 = p0
                                                                                                        v97 := GenS4{F0: 38, F1: 42}
                                                                                                        switch (fi57 % (((^(fi57))) | 1)) {
                                                                                                        case -9223372036854775808:
                                                                                                            var v98 uint64 = 4294967295
                                                                                                            _ = v98
                                                                                                        case -1:
                                                                                                            fmt.Println(fi57, p0, fi61)
                                                                                                        case 964:
                                                                                                            v99 := GenS6{F0: fi61}
                                                                                                            _ = v99
                                                                                                        default:
                                                                                                            var v100 uint16 = 32767
                                                                                                            _ = v100
                                                                                                        }
                                                                                                        fmt.Println(v62)
                                                                                                        var v101 bool = false
                                                                                                        fmt.Println(fi57)
                                                                                                        if v101 {
                                                                                                            p0 = 42
                                                                                                            var v102 int8 = -128
                                                                                                            if v101 {
                                                                                                                fmt.Println(v102)
                                                                                                            }
                                                                                                            _ = v102
                                                                                                        } else {
                                                                                                            v103 := []int{(fi61 % (((fi61 * (fi61 * genFn0("hello", (fi57 + genFn0("hello", fi57)))))) | 1)), (fi57 + fi61), fi61}
                                                                                                            _ = v103
                                                                                                        }
                                                                                                        var v104 uint = v23
                                                                                                        _ = v101
                                                                                                        _ = v104
                                                                                                        _ = v97
                                                                                                    }
                                                                                                    _ = v96
                                                                                                }
                                                                                            }
                                                                                            switch "" {
                                                                                            case "a":
                                                                                                v23 = p1
                                                                                            default:
                                                                                                var v105 uint8 = v74
                                                                                                _ = v105
                                                                                            }
                                                                                            v106 := GenS8{F0: (v74 % ((127) | 1))}
                                                                                            switch "\n\t" {
                                                                                            case "\n\t":
                                                                                                if (0.0 != (((v71 / ((v71 / v71) / v71)) / (v71 / 0.1)) + v71)) {
                                                                                                    fmt.Println(v62, v62, v62)
                                                                                                    if (3.14159 == 0.1) {
                                                                                                        v71 += v71
                                                                                                    } else {
                                                                                                        var v107 uint32 = 1
                                                                                                        _ = v107
                                                                                                    }
                                                                                                    v23 -= (v23 | 255)
                                                                                                } else {
                                                                                                    v108 := []int{-9223372036854775808, 4294967295}
                                                                                                    switch ("\n\t" + ("\n\t" + "\n\t")) {
                                                                                                    case "hello":
                                                                                                        v109 := GenS6{F0: (-(fi57))}
                                                                                                        _ = v109
                                                                                                    case "foo bar":
                                                                                                        _ = len(v108)
                                                                                                    }
                                                                                                    _ = v108
                                                                                                }
                                                                                                v74 += (v74 | (p0 * (^(p0))))
                                                                                                switch (fi57 / ((fi57) | 1)) {
                                                                                                case 631:
                                                                                                    v110 := []int{}
                                                                                                    _ = v110
                                                                                                case 128:
                                                                                                    v111 := []int{(fi61 - fi57), fi57}
                                                                                                    _ = v111
                                                                                                default:
                                                                                                    p0 -= v74
                                                                                                }
                                                                                            case "":
                                                                                                v62 -= -1
                                                                                            }
                                                                                            _ = v106
                                                                                            _ = v77
                                                                                        }
                                                                                        v74 -= v74
                                                                                    } else {
                                                                                        switch ("" + "\n\t") {
                                                                                        case "a":
                                                                                            fmt.Println(fi57)
                                                                                        case "foo bar":
                                                                                            v23 = p1
                                                                                        }
                                                                                        var v112 uint16 = 0
                                                                                        switch "foo bar" {
                                                                                        case "hello":
                                                                                            v113 := GenS5{F0: false, F1: 42}
                                                                                            _ = v113
                                                                                        case "\n\t":
                                                                                            v71 = (3.14159 / 0.1)
                                                                                        case "foo bar":
                                                                                            v112 = 42
                                                                                        default:
                                                                                            var v114 uint = v23
                                                                                            _ = v114
                                                                                        }
                                                                                        switch (fi57 >> ((p1) & 63)) {
                                                                                        case 1:
                                                                                            v115 := GenS9{F0: (v74 % ((v74) | 1)), F1: (^(v74)), F2: 1}
                                                                                            switch (fi61 >> (((p1 - p1)) & 63)) {
                                                                                            case 128:
                                                                                                fmt.Println(p0, fi57)
                                                                                            case 32768:
                                                                                                v23 += 18446744073709551615
                                                                                            }
                                                                                            v116 := []int{int(v74), fi57, (fi57 % (((fi61 % ((fi57) | 1))) | 1))}
                                                                                            _ = v115
                                                                                            _ = v116
                                                                                        case 32767:
                                                                                            var v117 int32 = -1
                                                                                            _ = v117
                                                                                        }
                                                                                        _ = v112
                                                                                    }
                                                                                }
                                                                                if false {
                                                                                    v71 += v71
                                                                                    v62 -= (v62 / ((int32(fi61)) | 1))
                                                                                    var v118 uint32 = 0
                                                                                    _ = v118
                                                                                } else {
                                                                                    fmt.Println(v74, fi57)
                                                                                    var v119 uint16 = 284
                                                                                    switch "foo bar" {
                                                                                    case "hello":
                                                                                        v120 := []int{856, (-(fi57)), -2147483648}
                                                                                        _ = v120
                                                                                    }
                                                                                    _ = v119
                                                                                }
                                                                                if true {
                                                                                    p0 -= p0
                                                                                } else {
                                                                                    var v121 uint8 = 42
                                                                                    _ = v121
                                                                                }
                                                                                switch (^(fi57)) {
                                                                                case -22:
                                                                                    if false {
                                                                                        fmt.Println(v71, p1)
                                                                                        if (1 > 18446744073709551615) {
                                                                                            v122 := GenS10{F0: "", F1: 42, F2: 18446744073709551615}
                                                                                            if (true || true) {
                                                                                                if ("foo bar" >= ("\n\t" + "")) {
                                                                                                    p1 = (^(p1))
                                                                                                }
                                                                                                var v123 uint64 = 1
                                                                                                if (386 <= 42) {
                                                                                                    switch ("hello" + "foo bar") {
                                                                                                    case "":
                                                                                                        v124 := GenS11{F0: (v62 >> ((v23) & 63))}
                                                                                                        _ = v124
                                                                                                    }
                                                                                                }
                                                                                                _ = v123
                                                                                            }
                                                                                            _ = v122
                                                                                        } else {
                                                                                            switch fi61 {
                                                                                            case 1:
                                                                                                switch (fi61 - genFn0(("" + "\n\t"), int(v62))) {
                                                                                                case 948:
                                                                                                    if (!((1 <= -231) || (!(((((3.14159 + v71) * v71) * 1.5) - ((v71 + ((float64(((((((1e10 + 3.14159) - 3.14159) - -2.25) * 3.14159) - 0.0) * (1e10 + 1e10))) - 0.1) + 1.5)) + float64(3.14159))) >= 1e10)))) {
                                                                                                        fmt.Println(fi57)
                                                                                                    } else {
                                                                                                        if true {
                                                                                                            v125 := GenS3{F0: 4294967295, F1: 427, F2: 255}
                                                                                                            _ = v125
                                                                                                        }
                                                                                                        v126 := []int{}
                                                                                                        v127 := GenS12{F0: 4294967295, F1: 497, F2: 100}
                                                                                                        _ = v126
                                                                                                        _ = v127
                                                                                                    }
                                                                                                    if false {
                                                                                                        var v128 uint16 = uint16(v23)
                                                                                                        if (4294967295 != uint32(v74)) {
                                                                                                            var v129 int32 = v62
                                                                                                            v130 := GenS8{F0: (^(v74))}
                                                                                                            v74 = p0
                                                                                                            _ = v129
                                                                                                            _ = v130
                                                                                                        } else {
                                                                                                            if (v74 >= p0) {
                                                                                                                v131 := []int{int(v74)}
                                                                                                                _ = v131
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v128
                                                                                                    } else {
                                                                                                        p1 -= v23
                                                                                                        v74 = 62
                                                                                                        if false {
                                                                                                            v71 = v71
                                                                                                            switch (fi61 | fi61) {
                                                                                                            case 0:
                                                                                                                v132 := GenS3{F0: 0, F1: 324, F2: 2}
                                                                                                                _ = v132
                                                                                                            case 42:
                                                                                                                var v133 string = "hello"
                                                                                                                _ = v133
                                                                                                            default:
                                                                                                                switch (fi57 / (((fi61 >> ((v23) & 63))) | 1)) {
                                                                                                                case 2147483647:
                                                                                                                    switch "hello" {
                                                                                                                    case "hello":
                                                                                                                        var v134 uint64 = 733
                                                                                                                        _ = v134
                                                                                                                    default:
                                                                                                                        if true {
                                                                                                                            if false {
                                                                                                                                v135 := GenS9{F0: 255, F1: (p0 / (((p0 ^ 1)) | 1)), F2: 1}
                                                                                                                                if true {
                                                                                                                                    if false {
                                                                                                                                        var v136 int8 = -1
                                                                                                                                        var v137 string = "foo bar"
                                                                                                                                        v138 := []int{-997, (fi57 + (fi57 | int(v136)))}
                                                                                                                                        _ = v136
                                                                                                                                        _ = v137
                                                                                                                                        _ = v138
                                                                                                                                    } else {
                                                                                                                                        fmt.Println(v23, p1, p1)
                                                                                                                                    }
                                                                                                                                    switch genFn0("", -9223372036854775808) {
                                                                                                                                    case -9223372036854775808:
                                                                                                                                        v62 = (v62 >> ((4294967295) & 63))
                                                                                                                                    case 32768:
                                                                                                                                        p1 = v23
                                                                                                                                    case 32767:
                                                                                                                                        var v139 uint16 = 852
                                                                                                                                        _ = v139
                                                                                                                                    default:
                                                                                                                                        p1 = p1
                                                                                                                                    }
                                                                                                                                } else {
                                                                                                                                    switch "\n\t" {
                                                                                                                                    case "\n\t":
                                                                                                                                        var v140 bool = ((false || true) && (p1 != (v23 >> ((255) & 63))))
                                                                                                                                        _ = v140
                                                                                                                                    default:
                                                                                                                                        if true {
                                                                                                                                            v71 = (((v71 / v71) - float64(3.14159)) + (v71 / ((float64((0.0 * (0.0 / (-2.25 / 1.0)))) + 1.5) * (1.5 + 1e10))))
                                                                                                                                            v23 -= 255
                                                                                                                                            var v141 float64 = v71
                                                                                                                                            if (!true) {
                                                                                                                                                v142 := GenS5{F0: true, F1: -450}
                                                                                                                                                _ = v142
                                                                                                                                            } else {
                                                                                                                                                v143 := []int{(fi57 * fi61), genFn0((("" + "a") + "hello"), 0)}
                                                                                                                                                var v144 int = (fi61 | genFn0(("hello" + "hello"), (fi61 + fi61)))
                                                                                                                                                fmt.Println(p1, v71)
                                                                                                                                                var v145 int64 = -388
                                                                                                                                                _ = v143
                                                                                                                                                _ = v144
                                                                                                                                                _ = v145
                                                                                                                                            }
                                                                                                                                            v74 = p0
                                                                                                                                            switch genFn0("foo bar", fi61) {
                                                                                                                                            case 0:
                                                                                                                                                v146 := []int{}
                                                                                                                                                _ = v146
                                                                                                                                            case -1:
                                                                                                                                                v147 := GenS13{F0: 907, F1: 1000, F2: 974}
                                                                                                                                                _ = v147
                                                                                                                                            case 2147483647:
                                                                                                                                                var v148 int64 = 1
                                                                                                                                                _ = v148
                                                                                                                                            }
                                                                                                                                            _ = v141
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    v149 := GenS14{F0: p1}
                                                                                                                                    _ = v149
                                                                                                                                }
                                                                                                                                var v150 string = (("" + "hello") + ("\n\t" + "foo bar"))
                                                                                                                                _ = v135
                                                                                                                                _ = v150
                                                                                                                            }
                                                                                                                            var v151 uint8 = p0
                                                                                                                            _ = v151
                                                                                                                        }
                                                                                                                    }
                                                                                                                default:
                                                                                                                    v152 := GenS13{F0: 784, F1: 21, F2: 141}
                                                                                                                    _ = v152
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        v153 := []int{fi57, fi61}
                                                                                                        v23 += v23
                                                                                                        _ = v153
                                                                                                    }
                                                                                                    v23 += p1
                                                                                                    var v154 bool = (710 == 2)
                                                                                                    _ = v154
                                                                                                case 4294967295:
                                                                                                    v74 -= 142
                                                                                                case -9223372036854775808:
                                                                                                    v155 := GenS9{F0: (p0 << ((v23) & 63)), F1: 255, F2: 32767}
                                                                                                    _ = v155
                                                                                                }
                                                                                            case 56:
                                                                                                v156 := []int{fi57}
                                                                                                _ = v156
                                                                                            case 42:
                                                                                                var v157 string = (("foo bar" + ("hello" + "hello")) + ("foo bar" + "\n\t"))
                                                                                                _ = v157
                                                                                            }
                                                                                            v158 := []int{42, int(v62), (-(fi57))}
                                                                                            if len(v158) >= 2 {
                                                                                                v158 = v158[1:len(v158)]
                                                                                            }
                                                                                            p0 = p0
                                                                                            _ = v158
                                                                                        }
                                                                                        var v159 int32 = (v62 * v62)
                                                                                        _ = v159
                                                                                    }
                                                                                }
                                                                                var v160 string = "hello"
                                                                                _ = v160
                                                                            }
                                                                        default:
                                                                            v161 := GenS1{F0: 339, F1: ("a" + "a"), F2: 127}
                                                                            _ = v161
                                                                        }
                                                                        fmt.Println(v74)
                                                                        v162 := GenS8{F0: v74}
                                                                        if (0 < 1000) {
                                                                            v74 -= p0
                                                                            switch (((("" + ("" + ("hello" + "foo bar"))) + "\n\t") + (("hello" + "hello") + "\n\t")) + "\n\t") {
                                                                            case "foo bar":
                                                                                if false {
                                                                                    if ((!true) || false) {
                                                                                        var v163 uint64 = 29
                                                                                        if (!false) {
                                                                                            v164 := []int{genFn0("hello", -9223372036854775808), (fi57 | fi57)}
                                                                                            _ = v164
                                                                                        } else {
                                                                                            fmt.Println(v163, fi57)
                                                                                            v165 := []int{(^(fi61))}
                                                                                            var v166 uint = v23
                                                                                            _ = v165
                                                                                            _ = v166
                                                                                        }
                                                                                        _ = v163
                                                                                    } else {
                                                                                        v167 := []int{}
                                                                                        _ = v167
                                                                                    }
                                                                                } else {
                                                                                    var v168 uint = p1
                                                                                    v169 := []int{}
                                                                                    p0 = 255
                                                                                    _ = len(v169)
                                                                                    if false {
                                                                                        if len(v169) > 0 {
                                                                                            v169[0] = (fi57 - fi61)
                                                                                        }
                                                                                        v62 = v62
                                                                                    } else {
                                                                                        _ = len(v169)
                                                                                    }
                                                                                    _ = v168
                                                                                    _ = v169
                                                                                }
                                                                            default:
                                                                                var v170 string = "foo bar"
                                                                                fmt.Println(p0, p1, fi61)
                                                                                _ = v170
                                                                            }
                                                                        } else {
                                                                            if (((fi61 / ((fi57) | 1)) != (^(fi57))) || true) {
                                                                                switch 127 {
                                                                                case 198:
                                                                                    var v171 uint16 = 32767
                                                                                    _ = v171
                                                                                }
                                                                                v172 := []int{}
                                                                                _ = v172
                                                                            } else {
                                                                                p0 -= v74
                                                                                var v173 float32 = ((((float32(1.5) * 1e10) * (((0.1 - (1e10 * -2.25)) / ((3.14159 * (1.5 / (0.1 - 1.5))) * ((((0.1 / 1e10) * -2.25) - 1e10) - ((3.14159 * 0.0) - -2.25)))) + float32(v71))) * 0.1) * 1.5)
                                                                                _ = v173
                                                                            }
                                                                        }
                                                                        fmt.Println(v23, v74)
                                                                        _ = v162
                                                                    } else {
                                                                        v174 := GenS14{F0: v23}
                                                                        fmt.Println(p1, v74, p1)
                                                                        switch "hello" {
                                                                        case "\n\t":
                                                                            fmt.Println(v74)
                                                                        case "":
                                                                            var v175 float64 = v71
                                                                            _ = v175
                                                                        default:
                                                                            var v176 bool = (2 > 183)
                                                                            _ = v176
                                                                        }
                                                                        _ = v174
                                                                    }
                                                                    var v177 uint32 = 0
                                                                    _ = v177
                                                                    _ = v74
                                                                } else {
                                                                    v178 := GenS14{F0: 2}
                                                                    v179 := GenS12{F0: -2147483648, F1: -1, F2: -16}
                                                                    switch 1 {
                                                                    case 440:
                                                                        switch genFn0(("" + "foo bar"), (fi57 << ((42) & 63))) {
                                                                        case 2147483647:
                                                                            v180 := []int{460, fi61}
                                                                            _ = v180
                                                                        case 127:
                                                                            v181 := []int{517, fi57}
                                                                            _ = v181
                                                                        }
                                                                    case -2147483648:
                                                                        if false {
                                                                            v71 += ((v71 + (((-2.25 + 3.14159) + (v71 - v71)) * v71)) / v71)
                                                                            v71 += (v71 - float64(1.5))
                                                                        } else {
                                                                            var v182 uint32 = 705
                                                                            _ = v182
                                                                        }
                                                                    case 42:
                                                                        fmt.Println(v23)
                                                                    }
                                                                    _ = v178
                                                                    _ = v179
                                                                }
                                                            default:
                                                                v183 := []int{(fi57 ^ (^(fi57))), 128, fi57, (fi57 / (((fi57 >> ((p1) & 63))) | 1))}
                                                                _ = v183
                                                            }
                                                        case -9223372036854775808:
                                                            if true {
                                                                v23 = p1
                                                                fmt.Println(v23, v71)
                                                                p1 -= p1
                                                            } else {
                                                                p1 = v23
                                                                if true {
                                                                    p0 += (^(p0))
                                                                }
                                                                v184 := GenS14{F0: p1}
                                                                if false {
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        var v185 string = (("" + "\n\t") + "a")
                                                                        _ = v185
                                                                    case "hello":
                                                                        v186 := GenS4{F0: 2, F1: -9223372036854775808}
                                                                        _ = v186
                                                                    }
                                                                }
                                                                if (!(v23 < (p1 / (((p1 << ((4294967295) & 63))) | 1)))) {
                                                                    var v187 float64 = ((v71 / (-2.25 - float64(float32((v71 / (v71 / v71)))))) / v71)
                                                                    _ = v187
                                                                } else {
                                                                    if true {
                                                                        v188 := []int{fi57, fi61, fi61}
                                                                        var v189 uint16 = 1
                                                                        v190 := GenS15{F0: fi61, F1: 42}
                                                                        _ = v188
                                                                        _ = v189
                                                                        _ = v190
                                                                    } else {
                                                                        var v191 int = fi61
                                                                        if false {
                                                                            var v192 int8 = 127
                                                                            if ((^(v62)) == (v62 + v62)) {
                                                                                var v193 uint = 328
                                                                                _ = v193
                                                                            } else {
                                                                                v71 = 0.0
                                                                                var v194 float32 = (0.1 - 0.0)
                                                                                v191 -= (-(v191))
                                                                                p0 -= 255
                                                                                v191 -= 712
                                                                                _ = v194
                                                                            }
                                                                            p1 -= p1
                                                                            _ = v192
                                                                        } else {
                                                                            var v195 int16 = 1000
                                                                            _ = v195
                                                                        }
                                                                        fmt.Println(fi61, v191, v191)
                                                                        _ = v191
                                                                    }
                                                                    var v196 bool = true
                                                                    _ = v196
                                                                }
                                                                var v197 string = ("" + ("" + "foo bar"))
                                                                _ = v184
                                                                _ = v197
                                                            }
                                                        case 128:
                                                            v198 := []int{}
                                                            _ = v198
                                                        }
                                                        switch (fi61 / ((genFn0("", (fi61 * fi61))) | 1)) {
                                                        case 9223372036854775807:
                                                            p1 = (v23 + (^(p1)))
                                                        case 1:
                                                            v23 += (p1 - 756)
                                                        case 4294967295:
                                                            if false {
                                                                var v199 float64 = (v71 * (3.14159 + v71))
                                                                fmt.Println(v71)
                                                                _ = v199
                                                            }
                                                        }
                                                        var v200 uint8 = (p0 % ((p0) | 1))
                                                        if true {
                                                            v71 += ((v71 - 0.1) / v71)
                                                            fmt.Println(v62)
                                                            fmt.Println(v62)
                                                        }
                                                        if (!(true && (1000 < 32767))) {
                                                            var v201 uint = 4294967295
                                                            var v202 int8 = -1
                                                            _ = v201
                                                            _ = v202
                                                        }
                                                        _ = v200
                                                    } else {
                                                        var v203 int32 = (v62 << ((v23) & 63))
                                                        _ = v203
                                                    }
                                                    var v204 uint32 = 540
                                                    v205 := GenS16{F0: 2147483647, F1: -634, F2: 9223372036854775807}
                                                    _ = v204
                                                    _ = v205
                                                case 32768:
                                                    switch "a" {
                                                    case "\n\t":
                                                        var v206 int64 = 65
                                                        _ = v206
                                                    case "":
                                                        var v207 int16 = 42
                                                        _ = v207
                                                    }
                                                default:
                                                    var v208 int = fi57
                                                    _ = v208
                                                }
                                                _ = v70
                                                _ = v71
                                            } else {
                                                var v209 int16 = -1
                                                _ = v209
                                            }
                                        case "a":
                                            var v210 uint64 = 2
                                            _ = v210
                                        }
                                        p1 = p1
                                    }
                                }
                            }
                            v211 := []int{(fi61 | 514), (-(fi57)), fi61, fi57}
                            _ = v211
                        }
                        fmt.Println(v62, fi61, p0)
                        var v212 uint16 = 1
                        _ = fi61
                        _ = v212
                        _ = v62
                    }
                    if (-128 > 93) {
                        fmt.Println(v23, p1)
                        var v213 uint32 = 572
                        v214 := []int{fi57, fi57}
                        _ = v213
                        _ = v214
                    }
                }
                _ = fi57
            }
        }
        var v215 int64 = -1
        _ = v215
    } else {
        var v216 int16 = -32768
        _ = v216
    }
    _ = p0
    _ = p1
    _ = v23
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 42
    v3 := make(map[int]bool)
    var v4 float64 = 3.14159
    v5 := make(map[int]bool)
    v217 := GenS17{F0: 65535, F1: 195}
    if true {
        switch 768 {
        case 1:
            var v218 string = "foo bar"
            _ = v218
        case 419:
            v2 = (v2 % ((2147483647) | 1))
        }
    } else {
        v2 -= 2147483648
        switch "" {
        case "\n\t":
            v4 += (v4 - (((((v4 - (v4 - v4)) / v4) / (float64(0.1) + 1e10)) * (0.0 + 0.1)) * 0.0))
        case "hello":
            for fi219 := 0; fi219 < 19; fi219++ {
                fmt.Println(fi219)
                _ = fi219
            }
        case "foo bar":
            switch 2147483647 {
            case 32768:
                v2 = v2
            default:
                fmt.Println(v4, v2)
            }
        }
        var v220 int64 = 9223372036854775807
        _ = v220
    }
    v2 -= v2
    var v221 int64 = 9223372036854775807
    for fi222 := 0; fi222 < 14; fi222++ {
        for fi223 := 0; fi223 < 7; fi223++ {
            v224 := GenS5{F0: true, F1: (v221 | v221)}
            var v225 uint8 = 127
            var v226 uint = 42
            _ = fi223
            _ = v224
            _ = v225
            _ = v226
        }
        for fi227 := 0; fi227 < 3; fi227++ {
            v221 = (v221 >> ((42) & 63))
            var v228 float32 = -2.25
            _ = fi227
            _ = v228
        }
        var v229 int = genFn0("\n\t", fi222)
        _ = fi222
        _ = v229
    }
    fmt.Println(v221, v2, v221)
    v230 := []int{}
    v4 += (-2.25 * (v4 + 0.0))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v221:", v221)
    fmt.Println("slicelen v230:", len(v230), v230)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v217
    _ = v221
    _ = v230
    _ = v3
    _ = v4
    _ = v5
}
