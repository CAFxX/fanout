package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 uint16; F2 int32 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 uint; F1 uint16 }
type GenS6 struct { F0 int64; F1 uint; F2 int16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int8; F1 int8 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int64; F1 uint64; F2 string }
type GenS11 struct { F0 uint; F1 uint16 }
func genFn0() uint {
    v4 := []int{4294967295, 32768}
    var v5 int32 = 2147483647
    if len(v4) > 0 {
        v4[0] = 127
    }
    for fi6 := 0; fi6 < 3; fi6++ {
        if len(v4) > 0 {
            v4[0] = (fi6 & 9223372036854775807)
        }
        _ = fi6
    }
    if (true || (42 <= 899)) {
        for fi7 := 0; fi7 < 9; fi7++ {
            for fi8 := 0; fi8 < 11; fi8++ {
                if true {
                    v5 = (v5 << ((255) & 63))
                }
                _ = fi8
            }
            v5 -= (^(v5))
            _ = fi7
        }
        var v9 int = -2147483648
        _ = v9
    }
    _ = v4
    _ = v5
    return 17
}
func genFn1(p0 uint8, p1 uint32) int16 {
    switch 127 {
    case 2147483647:
        for fi10 := 0; fi10 < 17; fi10++ {
            fmt.Println(p1, p0)
            p1 += p1
            switch "" {
            case "\n\t":
                if (-2147483648 >= 907) {
                    v11 := GenS1{F0: 1}
                    v12 := []int{fi10}
                    _ = v11
                    _ = v12
                } else {
                    if ((1e10 < (((float32(float64(0.0)) * (0.0 + (0.1 * ((((float32(0.1) / -2.25) * (1e10 + float32(0.0))) / (float32(1e10) + 1.5)) / 3.14159)))) + -2.25) / 1e10)) || true) {
                        for fi13 := 0; fi13 < 7; fi13++ {
                            v14 := []int{(fi10 - fi13)}
                            v15 := GenS2{F0: 42, F1: 2, F2: 42}
                            var v16 int64 = -9223372036854775808
                            _ = fi13
                            _ = v14
                            _ = v15
                            _ = v16
                        }
                    } else {
                        switch fi10 {
                        case -36:
                            switch (fi10 >> ((516) & 63)) {
                            case -9223372036854775808:
                                switch fi10 {
                                case -1:
                                    for fi17 := 0; fi17 < 8; fi17++ {
                                        v18 := GenS3{F0: true, F1: -706}
                                        v19 := GenS1{F0: -1}
                                        v20 := []int{}
                                        _ = fi17
                                        _ = v18
                                        _ = v19
                                        _ = v20
                                    }
                                case -2147483648:
                                    p0 -= 1
                                case 32767:
                                    switch (("\n\t" + "a") + "a") {
                                    case "\n\t":
                                        var v21 int = -2147483648
                                        _ = v21
                                    default:
                                        v22 := []int{fi10, -453}
                                        _ = v22
                                    }
                                default:
                                    if false {
                                        if (false && (255 < 1)) {
                                            fmt.Println(fi10, p0, p0)
                                        }
                                        var v23 float64 = float64(1.5)
                                        for fi24 := 0; fi24 < 3; fi24++ {
                                            var v25 int32 = 0
                                            if ((p0 - 255) == 128) {
                                                if (false && true) {
                                                    switch "\n\t" {
                                                    case "hello":
                                                        v26 := []int{fi24}
                                                        _ = v26
                                                    default:
                                                        switch "a" {
                                                        case "hello":
                                                            var v27 uint64 = uint64(p1)
                                                            _ = v27
                                                        case "foo bar":
                                                            fmt.Println(v23)
                                                            var v28 float64 = ((v23 + v23) / 0.1)
                                                            switch 32 {
                                                            case 4294967295:
                                                                if (264 >= fi10) {
                                                                    var v29 float32 = 1e10
                                                                    switch fi24 {
                                                                    case -9223372036854775808:
                                                                        p1 -= p1
                                                                    case 128:
                                                                        v30 := GenS3{F0: true, F1: v25}
                                                                        _ = v30
                                                                    }
                                                                    _ = v29
                                                                } else {
                                                                    fmt.Println(v23, p1)
                                                                }
                                                            case 651:
                                                                var v31 uint32 = p1
                                                                _ = v31
                                                            case -158:
                                                                p0 = 1
                                                            }
                                                            _ = v28
                                                        case "":
                                                            var v32 string = ("hello" + "foo bar")
                                                            _ = v32
                                                        }
                                                    }
                                                } else {
                                                    switch "a" {
                                                    case "\n\t":
                                                        v33 := []int{774, (fi24 | int(p0)), (fi24 | 127), (fi24 * fi24)}
                                                        _ = v33
                                                    }
                                                    fmt.Println(fi10)
                                                    v34 := []int{-2147483648, (fi10 - (^(fi24)))}
                                                    _ = v34
                                                }
                                                v35 := GenS4{F0: 42, F1: 42}
                                                var v36 string = "foo bar"
                                                _ = v35
                                                _ = v36
                                            }
                                            if (!false) {
                                                var v37 uint = uint(p1)
                                                var v38 uint64 = 0
                                                switch "a" {
                                                case "hello":
                                                    v39 := []int{(fi10 - 128), (fi24 >> ((0) & 63)), (-(fi24))}
                                                    _ = v39
                                                default:
                                                    v40 := []int{(fi24 + 0), (^(fi10)), (fi24 | int(v25)), 32767}
                                                    _ = v40
                                                }
                                                _ = v37
                                                _ = v38
                                            } else {
                                                p0 += 2
                                                if (1 != 4294967295) {
                                                    v41 := []int{fi10, fi24}
                                                    var v42 uint32 = p1
                                                    _ = v41
                                                    _ = v42
                                                } else {
                                                    if false {
                                                        fmt.Println(fi10)
                                                        switch "a" {
                                                        case "\n\t":
                                                            var v43 uint32 = p1
                                                            _ = v43
                                                        case "a":
                                                            var v44 int32 = (v25 * v25)
                                                            _ = v44
                                                        }
                                                    }
                                                    v45 := GenS5{F0: 431, F1: uint16(p0)}
                                                    _ = v45
                                                }
                                                v46 := []int{(^(fi24))}
                                                _ = v46
                                            }
                                            _ = fi24
                                            _ = v25
                                        }
                                        _ = v23
                                    }
                                }
                            case 1:
                                for fi47 := 0; fi47 < 1; fi47++ {
                                    if ((42 >= -994) || true) {
                                        v48 := []int{(-(fi47)), 832, fi47}
                                        _ = v48
                                    }
                                    _ = fi47
                                }
                            case -2147483648:
                                if (9223372036854775807 >= -1) {
                                    fmt.Println(p0, p0)
                                    var v49 string = (("a" + ("hello" + "")) + "")
                                    _ = v49
                                }
                            }
                        case 9223372036854775807:
                            fmt.Println(fi10, p1)
                        case -938:
                            if (-128 != 1) {
                                var v50 int16 = -32768
                                fmt.Println(p1, fi10, p0)
                                v51 := []int{0, fi10}
                                _ = v50
                                _ = v51
                            }
                        }
                    }
                }
            default:
                switch (fi10 % ((fi10) | 1)) {
                case 32767:
                    switch fi10 {
                    case -9223372036854775808:
                        p0 += (p0 ^ (p0 & 1))
                    case 127:
                        if true {
                            p1 += (p1 | (^(p1)))
                            p0 -= (p0 / ((p0) | 1))
                            if (!true) {
                                v52 := []int{fi10, (^(fi10)), 0}
                                _ = v52
                            } else {
                                if false {
                                    var v53 bool = false
                                    var v54 float32 = 0.1
                                    _ = v53
                                    _ = v54
                                } else {
                                    v55 := []int{}
                                    _ = v55
                                }
                                v56 := GenS4{F0: 2, F1: 1}
                                for fi57 := 0; fi57 < 10; fi57++ {
                                    var v58 float32 = 1.5
                                    v58 = v58
                                    _ = fi57
                                    _ = v58
                                }
                                _ = v56
                            }
                        } else {
                            fmt.Println(p1, fi10, p1)
                            if (976 > 42) {
                                var v59 uint = 18446744073709551615
                                var v60 float32 = 1.5
                                if false {
                                    for fi61 := 0; fi61 < 14; fi61++ {
                                        v60 -= 3.14159
                                        p0 = p0
                                        var v62 uint = v59
                                        _ = fi61
                                        _ = v62
                                    }
                                } else {
                                    v63 := []int{(fi10 & fi10), fi10, fi10}
                                    _ = v63
                                }
                                _ = v59
                                _ = v60
                            } else {
                                p1 = (p1 >> ((uint(p1)) & 63))
                                if false {
                                    switch "hello" {
                                    case "hello":
                                        v64 := GenS4{F0: 18446744073709551615, F1: 1}
                                        _ = v64
                                    }
                                    p1 = (p1 >> ((18446744073709551615) & 63))
                                } else {
                                    v65 := GenS4{F0: 42, F1: -9223372036854775808}
                                    switch ("foo bar" + "") {
                                    case "hello":
                                        p0 = p0
                                    }
                                    _ = v65
                                }
                            }
                        }
                    case -870:
                        if (4294967295 != 611) {
                            if (915 <= 255) {
                                switch fi10 {
                                case -1:
                                    if ((1e10 / ((((1.5 / 1.0) - 0.1) * 1.5) / 3.14159)) < 0.0) {
                                        switch -400 {
                                        case -9223372036854775808:
                                            v66 := GenS6{F0: 42, F1: 0, F2: 0}
                                            _ = v66
                                        case 127:
                                            var v67 int16 = int16(p1)
                                            _ = v67
                                        }
                                        p1 = (p1 - (p1 / ((p1) | 1)))
                                        v68 := GenS6{F0: 0, F1: 2, F2: 95}
                                        _ = v68
                                    }
                                case -387:
                                    var v69 uint32 = (p1 * p1)
                                    _ = v69
                                }
                                p1 += 65535
                            }
                            var v70 uint32 = 42
                            _ = v70
                        } else {
                            fmt.Println(p0, p0)
                            var v71 float32 = float32(((0.0 + 1e10) - (1e10 + (((0.0 * (-2.25 / 1e10)) * ((1e10 + 3.14159) * (0.0 + (0.0 * (0.0 - 3.14159))))) / 1.5))))
                            v72 := GenS2{F0: 4294967295, F1: 1, F2: 0}
                            _ = v71
                            _ = v72
                        }
                    default:
                        v73 := GenS7{F0: 2147483647}
                        _ = v73
                    }
                case 189:
                    p0 = 2
                }
            }
            _ = fi10
        }
    case -953:
        for fi74 := 0; fi74 < 20; fi74++ {
            v75 := []int{}
            for fi76 := 0; fi76 < 2; fi76++ {
                v77 := GenS6{F0: -856, F1: 65535, F2: 0}
                if len(v75) > 0 {
                    v75[0] = (fi74 - (fi74 % (((fi76 | fi76)) | 1)))
                }
                v78 := GenS3{F0: true, F1: 1}
                _ = fi76
                _ = v77
                _ = v78
            }
            v79 := GenS8{F0: 45, F1: 7}
            _ = fi74
            _ = v75
            _ = v79
        }
    case 1:
        for fi80 := 0; fi80 < 3; fi80++ {
            for fi81 := 0; fi81 < 5; fi81++ {
                p0 += 42
                _ = fi81
            }
            _ = fi80
        }
    default:
        for fi82 := 0; fi82 < 14; fi82++ {
            switch "\n\t" {
            case "":
                if (!true) {
                    var v83 int16 = -1
                    _ = v83
                } else {
                    var v84 int = (fi82 % ((fi82) | 1))
                    var v85 uint8 = 128
                    var v86 int16 = 32767
                    _ = v84
                    _ = v85
                    _ = v86
                }
            case "\n\t":
                switch (fi82 % ((0) | 1)) {
                case -9223372036854775808:
                    p1 += p1
                case 0:
                    var v87 string = ("foo bar" + "\n\t")
                    _ = v87
                default:
                    switch 0 {
                    case 9223372036854775807:
                        var v88 int8 = 1
                        _ = v88
                    case 4294967295:
                        v89 := []int{(-(fi82)), fi82}
                        _ = v89
                    default:
                        var v90 string = ""
                        _ = v90
                    }
                }
            }
            _ = fi82
        }
    }
    switch 42 {
    case 419:
        p1 = p1
    case 2147483647:
        v91 := []int{32768, -181}
        _ = v91
    default:
        for fi92 := 0; fi92 < 13; fi92++ {
            fmt.Println(p1, fi92)
            _ = fi92
        }
    }
    v93 := GenS9{F0: 0}
    _ = p0
    _ = p1
    _ = v93
    return 0
}

func main() {
    var v1 float64 = float64(1e10)
    v2 := []uint32{2147483648, 65535, 2147483648, 2147483647}
    var v3 uint64 = 18446744073709551615
    fmt.Println(v3, v3)
    fmt.Println(v1, v3)
    fmt.Println(v3, v1)
    for fi94 := 0; fi94 < 13; fi94++ {
        for fi95 := 0; fi95 < 9; fi95++ {
            var v96 int64 = 9223372036854775807
            _ = fi95
            _ = v96
        }
        _ = fi94
    }
    for fi97 := 0; fi97 < 20; fi97++ {
        switch fi97 {
        case 9223372036854775807:
            _ = len(v2)
        case 2147483647:
            v98 := GenS10{F0: -898, F1: (v3 ^ v3), F2: (("a" + ("hello" + "hello")) + "a")}
            _ = v98
        }
        var v99 uint16 = 255
        _ = fi97
        _ = v99
    }
    v100 := GenS3{F0: (!(v1 < (((v1 * (float64(1.5) - 3.14159)) - float64((-2.25 / (3.14159 + ((float32((v1 + (v1 / v1))) * float32((v1 / ((((3.14159 + 1e10) - v1) - (v1 + (v1 + 3.14159))) + float64(3.14159))))) / 0.1))))) * -2.25))), F1: -803}
    if len(v2) > 0 {
        v2[0] = 493
    }
    switch 9223372036854775807 {
    case -996:
        v101 := GenS6{F0: 9223372036854775807, F1: 0, F2: 328}
        _ = v101
    }
    v102 := GenS6{F0: 0, F1: 2, F2: 0}
    fmt.Println(v3, v1)
    var v103 string = (("foo bar" + ((("\n\t" + ("a" + (("hello" + "hello") + (("" + "\n\t") + ("a" + "\n\t"))))) + "a") + ("\n\t" + "hello"))) + "foo bar")
    var v104 uint8 = 15
    var v105 float32 = (-2.25 - (1.5 - 3.14159))
    fmt.Println(v105)
    for fi106 := 0; fi106 < 17; fi106++ {
        if (0 > 65535) {
            var v107 string = "a"
            if false {
                v108 := GenS9{F0: 1}
                _ = len(v2)
                _ = len(v2)
                _ = v108
            } else {
                v104 += v104
            }
            _ = v107
        } else {
            v105 -= float32(-2.25)
        }
        v3 -= (v3 / (((v3 - v3)) | 1))
        if false {
            for fi109 := 0; fi109 < 19; fi109++ {
                var v110 int = fi109
                _ = fi109
                _ = v110
            }
            switch (v103 + (v103 + ("a" + ("" + (("a" + (v103 + v103)) + ""))))) {
            case "foo bar":
                v104 = v104
            }
            switch fi106 {
            case 9223372036854775807:
                if (true && (!false)) {
                    if (3.14159 < (v1 / 1.5)) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        v104 = v104
                        var v111 string = v103
                        for fi112 := 0; fi112 < 2; fi112++ {
                            var v113 int16 = -32768
                            v113 += v113
                            _ = fi112
                            _ = v113
                        }
                        _ = v111
                    }
                    for fi114 := 0; fi114 < 6; fi114++ {
                        _ = len(v2)
                        if true {
                            fmt.Println(v1, v104)
                        } else {
                            var v115 int8 = 72
                            _ = v115
                        }
                        v104 = (v104 << ((2) & 63))
                        v105 += ((float32((v1 * v1)) / 1.0) - 1e10)
                        _ = fi114
                    }
                } else {
                    switch v103 {
                    case "a":
                        v1 = ((float64(v105) / v1) / v1)
                    case "":
                        for fi116 := 0; fi116 < 13; fi116++ {
                            var v117 float32 = ((v105 / 1.5) * v105)
                            _ = fi116
                            _ = v117
                        }
                    }
                    for fi118 := 0; fi118 < 20; fi118++ {
                        v105 = v105
                        _ = fi118
                    }
                }
            case 2147483647:
                if true {
                    for fi119 := 0; fi119 < 4; fi119++ {
                        var v120 string = v103
                        if true {
                            fmt.Println(v104, fi119, v1)
                        } else {
                            fmt.Println(v103)
                            v120 = v103
                            v121 := GenS11{F0: 393, F1: 42}
                            _ = v121
                        }
                        var v122 string = v103
                        _ = fi119
                        _ = v120
                        _ = v122
                    }
                }
            case 4294967295:
                _ = len(v2)
            }
        } else {
            if true {
                switch v103 {
                case "\n\t":
                    switch ("" + ("hello" + "a")) {
                    case "hello":
                        fmt.Println(v3, v104)
                    }
                case "hello":
                    for fi123 := 0; fi123 < 14; fi123++ {
                        var v124 string = "\n\t"
                        var v125 int32 = -1
                        _ = fi123
                        _ = v124
                        _ = v125
                    }
                case "a":
                    var v126 uint = 975
                    _ = v126
                }
                fmt.Println(v103, fi106, v105)
            }
        }
        _ = fi106
    }
    for fi127 := 0; fi127 < 13; fi127++ {
        switch v103 {
        case "a":
            v2 = append(v2, 42)
        case "hello":
            for fi128 := 0; fi128 < 10; fi128++ {
                v1 = float64(v105)
                v129 := GenS7{F0: 42}
                fmt.Println(v103)
                _ = fi128
                _ = v129
            }
        default:
            var v130 float64 = 0.0
            _ = v130
        }
        var v131 uint = 953
        if (!(877 != 336)) {
            for fi132 := 0; fi132 < 18; fi132++ {
                v131 -= v131
                var v133 int32 = 2147483647
                v105 = v105
                _ = fi132
                _ = v133
            }
        }
        _ = fi127
        _ = v131
    }
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v100
    _ = v102
    _ = v103
    _ = v104
    _ = v105
    _ = v2
    _ = v3
}
