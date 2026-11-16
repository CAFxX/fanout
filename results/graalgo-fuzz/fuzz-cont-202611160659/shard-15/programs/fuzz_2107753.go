package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 int; F2 uint64 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS5 struct { F0 uint32; F1 uint; F2 int16 }
type GenS6 struct { F0 int32; F1 int8; F2 uint32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint; F1 int16; F2 string }
type GenS9 struct { F0 uint16; F1 int32 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int16; F1 int; F2 int16 }
type GenS13 struct { F0 int64; F1 uint32 }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 int; F1 uint32 }
type GenS16 struct { F0 int16; F1 string; F2 uint64 }
type GenS17 struct { F0 int32 }
type GenS18 struct { F0 uint16; F1 int32 }
type GenS19 struct { F0 int }
type GenS20 struct { F0 int8 }
type GenS21 struct { F0 string; F1 uint; F2 string }
type GenS22 struct { F0 int32; F1 string; F2 int64 }
type GenS23 struct { F0 int16; F1 bool; F2 uint32 }
type GenS24 struct { F0 string }
func genFn0(p0 int16, p1 uint, p2 string) uint32 {
    var v6 int16 = -149
    switch 42 {
    case 128:
        for fi7 := 0; fi7 < 15; fi7++ {
            fmt.Println(p1)
            switch p2 {
            case "foo bar":
                if true {
                    fmt.Println(v6, p0)
                    switch p2 {
                    case "":
                        v8 := GenS1{F0: (p0 << (((p1 & 18446744073709551615)) & 63))}
                        _ = v8
                    case "foo bar":
                        for fi9 := 0; fi9 < 19; fi9++ {
                            v10 := []int{(fi9 >> (((^(p1))) & 63))}
                            v6 += (v6 * v6)
                            var v11 uint16 = 0
                            _ = fi9
                            _ = v10
                            _ = v11
                        }
                    default:
                        v12 := []int{}
                        _ = v12
                    }
                }
            case "\n\t":
                v13 := []int{fi7, fi7, int(p1), -1}
                _ = v13
            case "":
                switch (fi7 << ((p1) & 63)) {
                case -391:
                    var v14 int32 = 42
                    _ = v14
                case 2147483647:
                    if ((-(fi7)) == fi7) {
                        fmt.Println(p0, p0, v6)
                        p1 = 1
                        v15 := GenS1{F0: p0}
                        _ = v15
                    } else {
                        v16 := GenS1{F0: 32767}
                        p2 = (p2 + p2)
                        var v17 bool = false
                        _ = v16
                        _ = v17
                    }
                default:
                    p0 -= v6
                }
            default:
                for fi18 := 0; fi18 < 17; fi18++ {
                    fmt.Println(fi18, v6, v6)
                    switch ("\n\t" + p2) {
                    case "":
                        if true {
                            p0 = -32768
                            v19 := GenS2{F0: 42, F1: (fi18 & 0), F2: 42}
                            if false {
                                switch (fi7 << ((p1) & 63)) {
                                case 42:
                                    fmt.Println(v6, p0, fi7)
                                case 629:
                                    v20 := []int{}
                                    _ = v20
                                case 32767:
                                    var v21 string = "foo bar"
                                    _ = v21
                                }
                                fmt.Println(p1, v6, v6)
                                var v22 uint32 = 2
                                _ = v22
                            } else {
                                if (p0 > (p0 % ((p0) | 1))) {
                                    switch "" {
                                    case "":
                                        if (p2 == ((p2 + p2) + p2)) {
                                            if (-287 == -588) {
                                                var v23 float32 = float32((1.5 + 0.1))
                                                v24 := GenS2{F0: 2147483648, F1: (fi18 - 74), F2: 42}
                                                p1 += p1
                                                var v25 float64 = 1e10
                                                _ = v23
                                                _ = v24
                                                _ = v25
                                            } else {
                                                var v26 bool = false
                                                _ = v26
                                            }
                                        } else {
                                            v27 := GenS2{F0: 0, F1: (fi18 ^ 42), F2: 1}
                                            _ = v27
                                        }
                                        if false {
                                            switch (p2 + ("foo bar" + ("hello" + (p2 + (p2 + ("a" + p2)))))) {
                                            case "foo bar":
                                                if true {
                                                    var v28 uint32 = 0
                                                    _ = v28
                                                }
                                            case "a":
                                                if (9223372036854775807 != -1) {
                                                    p0 = (v6 ^ v6)
                                                    v29 := GenS1{F0: v6}
                                                    p0 = v6
                                                    _ = v29
                                                }
                                            case "hello":
                                                p0 += -1
                                            }
                                            var v30 float32 = -2.25
                                            p1 = (p1 - 42)
                                            _ = v30
                                        }
                                    case "a":
                                        v31 := []int{(fi7 >> ((p1) & 63)), fi7, (fi7 >> ((42) & 63)), 292}
                                        _ = v31
                                    case "\n\t":
                                        v32 := GenS3{F0: 60, F1: fi18}
                                        _ = v32
                                    default:
                                        if false {
                                            v33 := GenS1{F0: p0}
                                            _ = v33
                                        } else {
                                            v34 := []int{(fi7 - (fi7 >> (((p1 >> (((p1 & p1)) & 63))) & 63)))}
                                            _ = v34
                                        }
                                        switch ((p2 + "hello") + "\n\t") {
                                        case "foo bar":
                                            v35 := GenS2{F0: 306, F1: (fi7 + -1), F2: 18446744073709551615}
                                            _ = v35
                                        case "\n\t":
                                            switch (-(fi7)) {
                                            case -1:
                                                if false {
                                                    if false {
                                                        fmt.Println(p0)
                                                        switch p2 {
                                                        case "":
                                                            if (0 >= 42) {
                                                                if false {
                                                                    v36 := GenS3{F0: 110, F1: fi7}
                                                                    if false {
                                                                        v37 := GenS2{F0: 21, F1: (fi18 & (fi18 / ((fi18) | 1))), F2: 2}
                                                                        _ = v37
                                                                    }
                                                                    p2 = "foo bar"
                                                                    _ = v36
                                                                } else {
                                                                    v38 := GenS2{F0: 932, F1: (fi7 >> ((0) & 63)), F2: 18446744073709551615}
                                                                    _ = v38
                                                                }
                                                            }
                                                        }
                                                        p0 -= -826
                                                        if (42 <= 2147483648) {
                                                            v39 := []int{(fi7 % ((fi18) | 1))}
                                                            v6 = int16(fi7)
                                                            var v40 int = fi18
                                                            _ = v39
                                                            _ = v40
                                                        }
                                                    }
                                                    v41 := []int{(fi7 >> (((p1 % (((p1 << ((0) & 63))) | 1))) & 63)), 9223372036854775807}
                                                    _ = v41
                                                } else {
                                                    fmt.Println(p1)
                                                    if true {
                                                        var v42 uint32 = 42
                                                        _ = v42
                                                    } else {
                                                        v43 := GenS4{F0: 65535, F1: 42, F2: 4294967295}
                                                        switch fi18 {
                                                        case 128:
                                                            if true {
                                                                v44 := GenS4{F0: 304, F1: 682, F2: 115}
                                                                _ = v44
                                                            }
                                                        case -1:
                                                            v45 := []int{}
                                                            _ = v45
                                                        case 127:
                                                            v46 := GenS4{F0: 1, F1: 255, F2: 2147483648}
                                                            _ = v46
                                                        default:
                                                            p0 = -32768
                                                        }
                                                        p0 = -155
                                                        _ = v43
                                                    }
                                                    p1 -= 903
                                                }
                                            case -2147483648:
                                                var v47 uint32 = 740
                                                _ = v47
                                            }
                                        default:
                                            var v48 uint16 = 1
                                            _ = v48
                                        }
                                    }
                                    if (176 <= 143) {
                                        switch (fi18 >> ((p1) & 63)) {
                                        case 2147483647:
                                            v6 = v6
                                        default:
                                            switch fi7 {
                                            case 674:
                                                if false {
                                                    v49 := []int{fi18, int(p1), (^(fi7)), fi7}
                                                    v6 += (v6 ^ p0)
                                                    _ = v49
                                                }
                                            case 0:
                                                v6 -= (v6 ^ p0)
                                            case -9223372036854775808:
                                                v6 = p0
                                            }
                                        }
                                    }
                                }
                                v50 := GenS5{F0: 63, F1: (p1 << ((p1) & 63)), F2: v6}
                                var v51 string = p2
                                _ = v50
                                _ = v51
                            }
                            fmt.Println(p1, p1, p0)
                            if (true && (!true)) {
                                switch p2 {
                                case "\n\t":
                                    v52 := GenS6{F0: 0, F1: 56, F2: 2147483647}
                                    _ = v52
                                case "hello":
                                    switch p2 {
                                    case "":
                                        v6 = v6
                                    case "foo bar":
                                        v53 := GenS7{F0: 42}
                                        p1 += p1
                                        _ = v53
                                    case "\n\t":
                                        var v54 uint8 = 42
                                        _ = v54
                                    default:
                                        switch (fi18 + 32768) {
                                        case 9223372036854775807:
                                            fmt.Println(fi18, p1)
                                        case 42:
                                            if (397 > p0) {
                                                v55 := GenS7{F0: (^(p1))}
                                                _ = v55
                                            }
                                        case 32768:
                                            switch "a" {
                                            case "":
                                                var v56 bool = (42 <= 71)
                                                _ = v56
                                            }
                                        }
                                    }
                                case "a":
                                    if (false || true) {
                                        v57 := GenS2{F0: 4294967295, F1: (fi18 & -1), F2: 18446744073709551615}
                                        _ = v57
                                    } else {
                                        var v58 uint64 = 4294967295
                                        _ = v58
                                    }
                                }
                                fmt.Println(fi18, v6)
                                fmt.Println(p0, v6)
                                var v59 string = ((p2 + p2) + (p2 + (p2 + "hello")))
                                _ = v59
                            }
                            _ = v19
                        } else {
                            v60 := []int{fi18, (fi7 & (fi7 | (fi7 / ((fi18) | 1))))}
                            v61 := GenS2{F0: 1, F1: fi18, F2: 259}
                            var v62 bool = ((65535 != 255) && false)
                            _ = v60
                            _ = v61
                            _ = v62
                        }
                    case "\n\t":
                        fmt.Println(p2, fi18)
                        var v63 int8 = 42
                        fmt.Println(fi7, p0)
                        _ = v63
                    }
                    v64 := []int{fi18, fi7, (fi7 - (fi18 ^ (fi7 / (((-(fi7))) | 1))))}
                    switch "\n\t" {
                    case "foo bar":
                        if len(v64) >= 2 {
                            v64 = v64[1:len(v64)]
                        }
                    default:
                        p1 -= 18446744073709551615
                    }
                    var v65 int8 = 42
                    _ = fi18
                    _ = v64
                    _ = v65
                }
            }
            _ = fi7
        }
    case 32767:
        p2 = "hello"
    }
    var v66 string = p2
    for fi67 := 0; fi67 < 4; fi67++ {
        var v68 uint64 = 42
        for fi69 := 0; fi69 < 7; fi69++ {
            v68 += v68
            _ = fi69
        }
        _ = fi67
        _ = v68
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v66
    return 42
}
func genFn1(p0 bool, p1 uint64, p2 uint16) int32 {
    if ((!false) && (true && p0)) {
        v70 := []int{}
        var v71 int32 = -513
        _ = v70
        _ = v71
    } else {
        p0 = p0
        var v72 int64 = 1
        _ = v72
    }
    var v73 string = "\n\t"
    p0 = true
    _ = p0
    _ = p1
    _ = p2
    _ = v73
    return 364
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]string)
    v3 := []int{-915, 32768, 128}
    var v4 float64 = (-2.25 - 1e10)
    var v5 uint = 255
    switch "foo bar" {
    case "foo bar":
        v5 -= (^(v5))
    case "hello":
        if (!false) {
            switch "a" {
            case "hello":
                var v74 uint32 = 0
                _ = v74
            case "foo bar":
                var v75 bool = (!(-692 < -497))
                _ = v75
            default:
                var v76 int32 = 1
                _ = v76
            }
            var v77 float32 = 0.0
            var v78 bool = ((v4 * (v4 / v4)) > (v4 + 0.1))
            _ = v77
            _ = v78
        } else {
            switch 980 {
            case 42:
                v5 = v5
            default:
                v5 += v5
            }
            v5 += 255
            var v79 float32 = 0.0
            _ = v79
        }
    case "\n\t":
        v80 := GenS8{F0: (v5 - v5), F1: 42, F2: ("a" + "")}
        _ = v80
    default:
        fmt.Println(v5, v4)
    }
    for fi81 := 0; fi81 < 10; fi81++ {
        for fi82 := 0; fi82 < 6; fi82++ {
            if true {
                var v83 int64 = 421
                if false {
                    v4 += float64(((1e10 / 1.0) + 0.1))
                    var v84 uint32 = 65535
                    _ = v84
                } else {
                    v5 = v5
                }
                v85 := GenS9{F0: 404, F1: 1}
                _ = v83
                _ = v85
            } else {
                switch (fi82 * fi81) {
                case 2147483647:
                    v4 = v4
                case -2147483648:
                    if (185 <= (v5 * (v5 | v5))) {
                        v4 = 0.1
                    } else {
                        v2[(fi81 | (fi82 >> ((v5) & 63)))] = "foo bar"
                        switch "a" {
                        case "foo bar":
                            switch fi82 {
                            case 4294967295:
                                if true {
                                    var v86 int32 = 1
                                    v87 := GenS3{F0: 34, F1: fi82}
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = v86
                                    _ = v87
                                } else {
                                    switch ("foo bar" + (("" + "\n\t") + ("hello" + ("hello" + ("hello" + ""))))) {
                                    case "\n\t":
                                        var v88 uint16 = 0
                                        _ = v88
                                    case "":
                                        fmt.Println(fi81, fi82)
                                    case "a":
                                        if len(v3) > 0 {
                                            v3[0] = 0
                                        }
                                    }
                                    v5 = (^(v5))
                                    v4 = 3.14159
                                }
                            case -9223372036854775808:
                                if (738 < 603) {
                                    var v89 uint16 = 42
                                    v5 = (v5 / ((0) | 1))
                                    _ = v89
                                } else {
                                    var v90 float64 = v4
                                    switch "foo bar" {
                                    case "\n\t":
                                        if (float32(v4) != 1e10) {
                                            if (255 <= 86) {
                                                if (!true) {
                                                    if (0.0 <= (0.1 / 1.0)) {
                                                        v90 = v90
                                                        var v91 uint = v5
                                                        _ = v91
                                                    } else {
                                                        _ = len(v1)
                                                    }
                                                    if (!(false || (v5 < v5))) {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    } else {
                                                        var v92 uint64 = 2
                                                        v3 = append(v3, fi82)
                                                        if ((0.0 / float32(1e10)) == 1.5) {
                                                            if true {
                                                                v92 += v92
                                                                v93 := GenS1{F0: 0}
                                                                var v94 uint16 = 255
                                                                v95 := GenS10{F0: 2147483647}
                                                                _ = v93
                                                                _ = v94
                                                                _ = v95
                                                            } else {
                                                                switch "" {
                                                                case "\n\t":
                                                                    switch ("\n\t" + "a") {
                                                                    case "":
                                                                        var v96 uint64 = 2
                                                                        v97 := GenS9{F0: 255, F1: genFn1((2147483647 < (fi81 | 684)), v96, 32767)}
                                                                        _ = v96
                                                                        _ = v97
                                                                    }
                                                                case "":
                                                                    if (!(false && (!(32767 <= 385)))) {
                                                                        v3 = append(v3, (fi81 ^ fi82))
                                                                    }
                                                                    var v98 float64 = v90
                                                                    delete(v1, ("\n\t" + ("\n\t" + "hello")))
                                                                    _ = v98
                                                                default:
                                                                    if ((fi81 + (fi81 / (((fi82 & (fi81 * fi81))) | 1))) > fi81) {
                                                                        var v99 uint = (^(v5))
                                                                        v4 -= v90
                                                                        _ = v99
                                                                    }
                                                                }
                                                                v100 := GenS1{F0: 32767}
                                                                _ = v100
                                                            }
                                                            switch fi82 {
                                                            case -2147483648:
                                                                if true {
                                                                    v101 := GenS9{F0: uint16(v92), F1: -818}
                                                                    _ = v101
                                                                } else {
                                                                    _ = len(v2)
                                                                    v102 := GenS11{F0: 0}
                                                                    _ = v102
                                                                }
                                                            }
                                                            v90 = v90
                                                        } else {
                                                            v92 += 378
                                                        }
                                                        v103 := GenS11{F0: 42}
                                                        _ = v103
                                                        _ = v92
                                                    }
                                                    v2[fi81] = "a"
                                                } else {
                                                    var v104 bool = true
                                                    if (!false) {
                                                        v90 = v90
                                                        var v105 int32 = -588
                                                        _ = v105
                                                    }
                                                    _ = v104
                                                }
                                                if len(v3) > 0 {
                                                    v3[0] = fi81
                                                }
                                                if true {
                                                    var v106 int8 = -128
                                                    v4 += v4
                                                    switch "hello" {
                                                    case "hello":
                                                        v106 = 100
                                                    default:
                                                        v1["foo bar"] = (!(!false))
                                                    }
                                                    _ = v106
                                                } else {
                                                    var v107 int16 = 1000
                                                    var v108 float64 = v4
                                                    var v109 uint16 = 2
                                                    _ = v107
                                                    _ = v108
                                                    _ = v109
                                                }
                                            } else {
                                                var v110 bool = true
                                                v111 := GenS12{F0: -915, F1: (fi82 + (fi81 * fi81)), F2: 42}
                                                v4 -= v90
                                                _ = v110
                                                _ = v111
                                            }
                                        } else {
                                            v112 := GenS3{F0: 127, F1: -1}
                                            var v113 bool = (!false)
                                            v114 := GenS6{F0: 2147483647, F1: -27, F2: 2147483647}
                                            _ = v112
                                            _ = v113
                                            _ = v114
                                        }
                                    }
                                    if true {
                                        v4 = 1e10
                                        v90 -= v90
                                        switch fi82 {
                                        case 0:
                                            var v115 string = "a"
                                            _ = v115
                                        case -811:
                                            var v116 uint32 = 2147483648
                                            _ = v116
                                        default:
                                            var v117 uint32 = 2147483647
                                            _ = v117
                                        }
                                    }
                                    fmt.Println(v90, fi82)
                                    if ((2 < 156) && (!(0.0 == (float32(3.14159) / float32(v4))))) {
                                        if false {
                                            var v118 int32 = -2147483648
                                            v4 += float64((0.0 * (0.1 * float32(((v4 - 0.0) / float64(0.1))))))
                                            _ = v118
                                        } else {
                                            if true {
                                                v4 += float64((float32(v4) - 1.5))
                                                if (!(!(-1 != -13))) {
                                                    var v119 uint8 = 255
                                                    _ = v119
                                                }
                                            }
                                            v120 := GenS5{F0: 796, F1: 0, F2: 42}
                                            fmt.Println(v5)
                                            _ = v120
                                        }
                                        switch 127 {
                                        case -1:
                                            v4 += float64(-2.25)
                                        case -640:
                                            v90 += (v4 * (-2.25 / v90))
                                        case 1:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        var v121 int32 = 0
                                        switch ("hello" + "") {
                                        case "a":
                                            var v122 int32 = v121
                                            _ = v122
                                        case "":
                                            v4 += v90
                                        default:
                                            if false {
                                                if len(v3) > 0 {
                                                    v3[0] = -44
                                                }
                                            }
                                        }
                                        _ = v121
                                    } else {
                                        if true {
                                            var v123 string = ""
                                            if false {
                                                var v124 string = v123
                                                switch 32767 {
                                                case 640:
                                                    fmt.Println(v4, fi81, fi82)
                                                case 302:
                                                    switch v123 {
                                                    case "foo bar":
                                                        v90 = v90
                                                    case "a":
                                                        switch "hello" {
                                                        case "\n\t":
                                                            v5 = v5
                                                        case "hello":
                                                            switch v124 {
                                                            case "foo bar":
                                                                v90 += float64(0.1)
                                                            case "\n\t":
                                                                switch (fi81 + fi82) {
                                                                case 210:
                                                                    v125 := GenS12{F0: -1, F1: fi81, F2: -424}
                                                                    _ = v125
                                                                default:
                                                                    _ = len(v3)
                                                                }
                                                            default:
                                                                var v126 float32 = ((float32(((-2.25 + (((((v4 - v4) * ((v4 - (float64((-2.25 + float32(0.0))) / ((v90 / ((3.14159 / (v90 + ((float64((-2.25 + float32(v4))) - (v90 + (v4 - float64(float32(float64(float32(float64(0.0)))))))) - ((-2.25 / v90) / (v90 * ((v4 / v4) * v4)))))) + v90)) / float64((float32(v4) * 1.5))))) * v90)) / v4) - v4) / v4)) + v90)) - (((1.5 / 1.5) / ((3.14159 - (((((0.1 - 1e10) - -2.25) + (1e10 * float32(0.0))) - 1.5) + 1e10)) + 1e10)) * (3.14159 * 1.5))) / 1e10)
                                                                _ = v126
                                                            }
                                                        }
                                                    }
                                                case 0:
                                                    delete(v2, fi82)
                                                default:
                                                    v90 = (-2.25 + v90)
                                                }
                                                _ = v124
                                            } else {
                                                v90 -= 3.14159
                                                if (!true) {
                                                    if false {
                                                        v127 := GenS6{F0: 1, F1: 1, F2: 4294967295}
                                                        _ = v127
                                                    }
                                                    var v128 float64 = ((float64((float32((0.1 + ((((0.0 / 1.0) - v4) - v90) - float64(((0.1 * 1.5) * (0.1 + 3.14159)))))) / float32(((v90 / (v4 / (3.14159 * v90))) * 0.0)))) + v90) + float64(0.1))
                                                    v129 := GenS13{F0: 42, F1: 4294967295}
                                                    _ = v128
                                                    _ = v129
                                                }
                                            }
                                            v123 = v123
                                            _ = v123
                                        } else {
                                            fmt.Println(fi82, v4)
                                            var v130 int = 4294967295
                                            _ = len(v3)
                                            _ = v130
                                        }
                                    }
                                    _ = v90
                                }
                            case -651:
                                if (true || false) {
                                    if (!((!true) || ((fi81 - (fi82 >> ((v5) & 63))) > fi82))) {
                                        var v131 uint32 = 42
                                        _ = v131
                                    } else {
                                        switch "" {
                                        case "a":
                                            switch "\n\t" {
                                            case "\n\t":
                                                fmt.Println(v4, v4)
                                            default:
                                                if (false || true) {
                                                    v132 := GenS14{F0: 42}
                                                    _ = len(v3)
                                                    fmt.Println(fi82)
                                                    _ = v132
                                                } else {
                                                    v4 -= float64(float32(v4))
                                                }
                                            }
                                        case "foo bar":
                                            v5 = (v5 - (v5 << ((v5) & 63)))
                                        case "":
                                            var v133 int32 = 1
                                            _ = v133
                                        default:
                                            fmt.Println(v4, v4, v4)
                                        }
                                        v5 -= 2
                                    }
                                    var v134 float64 = v4
                                    v4 = v134
                                    _ = v134
                                }
                            default:
                                var v135 int8 = 76
                                _ = v135
                            }
                        case "hello":
                            switch "hello" {
                            case "\n\t":
                                fmt.Println(fi82, v5)
                            case "":
                                v136 := GenS13{F0: 1, F1: 0}
                                _ = v136
                            case "hello":
                                if true {
                                    fmt.Println(v5, v4, v5)
                                }
                            }
                        case "a":
                            v5 = v5
                        default:
                            if true {
                                v4 = 0.0
                                switch fi81 {
                                case 32768:
                                    fmt.Println(fi81, fi81, fi82)
                                case -2147483648:
                                    v3 = append(v3, (-(fi81)))
                                }
                                if true {
                                    v137 := GenS11{F0: -128}
                                    v5 += (v5 ^ 18446744073709551615)
                                    var v138 int16 = -1
                                    _ = v137
                                    _ = v138
                                }
                            } else {
                                var v139 uint = 889
                                _ = v139
                            }
                        }
                        v140 := GenS15{F0: (fi81 * (fi82 >> ((135) & 63))), F1: 1}
                        if true {
                            fmt.Println(fi82, v5, v5)
                            v141 := GenS16{F0: 500, F1: "", F2: 4294967295}
                            _ = v141
                        } else {
                            fmt.Println(fi81, fi82)
                            fmt.Println(fi81)
                        }
                        v142 := GenS13{F0: 42, F1: 1}
                        _ = v140
                        _ = v142
                    }
                default:
                    v4 = (v4 * (v4 * v4))
                }
            }
            v4 -= 3.14159
            v143 := GenS15{F0: (fi81 % (((fi81 | fi82)) | 1)), F1: genFn0(int16(v5), v5, "a")}
            _ = fi82
            _ = v143
        }
        switch "" {
        case "a":
            v4 = v4
        case "foo bar":
            for fi144 := 0; fi144 < 20; fi144++ {
                var v145 int16 = -660
                v146 := GenS3{F0: -12, F1: 2147483647}
                _ = fi144
                _ = v145
                _ = v146
            }
        case "":
            for fi147 := 0; fi147 < 3; fi147++ {
                v148 := GenS17{F0: 2147483647}
                v4 = 0.0
                _ = fi147
                _ = v148
            }
        default:
            var v149 bool = (2147483648 <= 2)
            _ = v149
        }
        _ = fi81
    }
    var v150 uint64 = 1
    switch "" {
    case "foo bar":
        if (true || true) {
            v150 = (v150 % (((v150 % (((v150 & v150)) | 1))) | 1))
            if true {
                var v151 float32 = 3.14159
                for fi152 := 0; fi152 < 7; fi152++ {
                    switch "\n\t" {
                    case "a":
                        if (!true) {
                            var v153 string = ("foo bar" + ("a" + (("" + "foo bar") + (("hello" + "foo bar") + "foo bar"))))
                            _ = len(v2)
                            var v154 int16 = -32768
                            _ = v153
                            _ = v154
                        }
                    case "\n\t":
                        for fi155 := 0; fi155 < 6; fi155++ {
                            if false {
                                if false {
                                    if (genFn0(1000, v5, ("" + "\n\t")) > 1) {
                                        v2[9223372036854775807] = (("\n\t" + "\n\t") + "foo bar")
                                    }
                                    if true {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    } else {
                                        var v156 int32 = 1
                                        var v157 string = ""
                                        if true {
                                            var v158 uint64 = v150
                                            _ = len(v1)
                                            _ = v158
                                        }
                                        _ = v156
                                        _ = v157
                                    }
                                    switch (fi152 >> ((v5) & 63)) {
                                    case 9223372036854775807:
                                        if (10 == 137) {
                                            var v159 uint8 = 1
                                            v150 += v150
                                            v160 := GenS8{F0: (v5 & v5), F1: -1, F2: "\n\t"}
                                            _ = v159
                                            _ = v160
                                        }
                                    case -9223372036854775808:
                                        switch "" {
                                        case "hello":
                                            v5 = 65535
                                        case "a":
                                            v161 := v1["foo bar"]
                                            _ = v161
                                        case "":
                                            var v162 int8 = -1
                                            _ = v162
                                        }
                                        _ = len(v1)
                                        v163 := GenS18{F0: 199, F1: -2147483648}
                                        var v164 int64 = 0
                                        _ = v163
                                        _ = v164
                                    case 42:
                                        fmt.Println(v151, v4, v150)
                                    default:
                                        delete(v2, (fi152 | -9223372036854775808))
                                    }
                                }
                                v151 = (v151 + float32((1e10 / 1.5)))
                                var v165 uint32 = 951
                                _ = v165
                            } else {
                                var v166 bool = (1 != 42)
                                v167 := GenS11{F0: -128}
                                _ = len(v1)
                                fmt.Println(fi155, fi152)
                                _ = v166
                                _ = v167
                            }
                            v168 := v1["hello"]
                            v169 := GenS19{F0: (fi155 << ((709) & 63))}
                            v170 := GenS3{F0: 100, F1: fi155}
                            _ = fi155
                            _ = v168
                            _ = v169
                            _ = v170
                        }
                    case "foo bar":
                        var v171 uint = (v5 / ((v5) | 1))
                        _ = v171
                    default:
                        for fi172 := 0; fi172 < 18; fi172++ {
                            var v173 bool = false
                            _ = fi172
                            _ = v173
                        }
                    }
                    _ = fi152
                }
                _ = v151
            }
            v174 := GenS5{F0: 429, F1: v5, F2: -682}
            _ = v174
        }
    default:
        var v175 int8 = -70
        _ = v175
    }
    if true {
        for fi176 := 0; fi176 < 12; fi176++ {
            for fi177 := 0; fi177 < 3; fi177++ {
                var v178 bool = (uint8(v5) > 0)
                delete(v2, -9223372036854775808)
                v179 := GenS20{F0: 42}
                fmt.Println(v150)
                if ((!v178) || (!true)) {
                    var v180 bool = true
                    _ = v180
                } else {
                    v181 := GenS3{F0: 42, F1: 127}
                    _ = v181
                }
                _ = fi177
                _ = v178
                _ = v179
            }
            v182 := v2[int(v5)]
            var v183 int64 = -9223372036854775808
            _ = fi176
            _ = v182
            _ = v183
        }
        switch 626 {
        case 1:
            v150 = 18446744073709551615
        case -2147483648:
            switch "a" {
            case "foo bar":
                if (100 != 123) {
                    v184 := GenS18{F0: 65535, F1: 66}
                    v4 = v4
                    _ = v184
                }
            default:
                var v185 uint16 = 0
                _ = v185
            }
        case -1:
            var v186 int32 = 0
            _ = v186
        }
    } else {
        _ = len(v3)
    }
    switch "a" {
    case "foo bar":
        for fi187 := 0; fi187 < 15; fi187++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v188 bool = true
            _ = fi187
            _ = v188
        }
    case "a":
        switch ((("" + ("a" + "foo bar")) + "hello") + "") {
        case "\n\t":
            if true {
                fmt.Println(v150)
                var v189 uint16 = 124
                switch 32767 {
                case -9223372036854775808:
                    for fi190 := 0; fi190 < 10; fi190++ {
                        var v191 uint32 = 65535
                        _ = fi190
                        _ = v191
                    }
                case 0:
                    _ = len(v3)
                }
                _ = v189
            }
        default:
            var v192 int16 = 42
            _ = v192
        }
    case "hello":
        var v193 uint64 = v150
        _ = v193
    default:
        var v194 bool = false
        _ = v194
    }
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        var v195 float32 = ((-2.25 + 0.1) / 0.1)
        _ = v195
    case "hello":
        if (false || (-2147483648 >= -490)) {
            switch 42 {
            case 0:
                var v196 int = 128
                _ = v196
            case -934:
                var v197 int64 = -9223372036854775808
                _ = v197
            default:
                if (-90 >= -128) {
                    v5 -= v5
                    for fi198 := 0; fi198 < 16; fi198++ {
                        v199 := GenS7{F0: v5}
                        v5 = v5
                        if true {
                            if (-522 != 524) {
                                switch -9223372036854775808 {
                                case 915:
                                    if false {
                                        if (false && true) {
                                            switch 32768 {
                                            case -2147483648:
                                                v150 -= (v150 << (((v5 / ((v5) | 1))) & 63))
                                            case 42:
                                                v200 := GenS10{F0: genFn1(false, 0, 1)}
                                                _ = v200
                                            }
                                            fmt.Println(v150, v150, v4)
                                            switch "foo bar" {
                                            case "a":
                                                v201 := GenS21{F0: ("foo bar" + (("foo bar" + ("foo bar" + "hello")) + "foo bar")), F1: (v5 / ((v5) | 1)), F2: "hello"}
                                                _ = v201
                                            case "":
                                                v5 = 903
                                            }
                                        }
                                        if false {
                                            fmt.Println(v5)
                                            v202 := GenS22{F0: genFn1(true, 569, 42), F1: "hello", F2: 0}
                                            _ = v202
                                        }
                                    } else {
                                        fmt.Println(v150)
                                        var v203 int = fi198
                                        v5 -= v5
                                        _ = v203
                                    }
                                case 621:
                                    var v204 uint64 = v150
                                    _ = v204
                                default:
                                    if false {
                                        if (!(1 > 193)) {
                                            if false {
                                                if false {
                                                    v150 = (v150 >> (((v5 & (^(v5)))) & 63))
                                                    v205 := GenS23{F0: 155, F1: false, F2: 1}
                                                    _ = v205
                                                } else {
                                                    v206 := GenS24{F0: "hello"}
                                                    var v207 uint64 = (v150 << ((v5) & 63))
                                                    fmt.Println(v150, v207)
                                                    _ = v206
                                                    _ = v207
                                                }
                                                v150 += v150
                                            }
                                            for fi208 := 0; fi208 < 4; fi208++ {
                                                v4 = 0.1
                                                delete(v1, "foo bar")
                                                _ = len(v3)
                                                _ = fi208
                                            }
                                            var v209 uint64 = 155
                                            _ = v209
                                        } else {
                                            if false {
                                                var v210 uint = 4294967295
                                                if true {
                                                    fmt.Println(fi198, fi198, fi198)
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    switch fi198 {
                                                    case -969:
                                                        var v211 int = fi198
                                                        _ = v211
                                                    }
                                                }
                                                var v212 bool = (-9223372036854775808 < -9223372036854775808)
                                                _ = v210
                                                _ = v212
                                            }
                                            v4 -= ((v4 / v4) - (float64(1.5) + ((v4 - (((v4 - 0.0) + (1e10 / (3.14159 / (v4 - ((-2.25 * -2.25) - v4))))) / (1e10 - (v4 - v4)))) / ((v4 + 1e10) + v4))))
                                        }
                                        if true {
                                            for fi213 := 0; fi213 < 3; fi213++ {
                                                v5 += v5
                                                _ = fi213
                                            }
                                            v4 += ((float64(3.14159) * v4) * (v4 - 1.5))
                                            var v214 uint16 = 1
                                            _ = v214
                                        }
                                        v4 += 0.1
                                    } else {
                                        v150 = 0
                                    }
                                }
                                v215 := GenS23{F0: -194, F1: (!("a" < "")), F2: 2147483647}
                                _ = v215
                            }
                            fmt.Println(v5, v5)
                            v5 += (v5 / (((v5 - v5)) | 1))
                        }
                        _ = fi198
                        _ = v199
                    }
                    var v216 float32 = 0.1
                    _ = v216
                } else {
                    for fi217 := 0; fi217 < 1; fi217++ {
                        var v218 uint = (v5 * (v5 / (((^(v5))) | 1)))
                        switch (-(fi217)) {
                        case 42:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case -87:
                            if ((!(2 == (v150 + v150))) || (false && true)) {
                                var v219 uint32 = 4294967295
                                var v220 uint = (v218 / ((v218) | 1))
                                _ = v219
                                _ = v220
                            } else {
                                for fi221 := 0; fi221 < 15; fi221++ {
                                    v3 = append(v3, (fi221 + fi217))
                                    var v222 int16 = -810
                                    _ = fi221
                                    _ = v222
                                }
                                if false {
                                    for fi223 := 0; fi223 < 2; fi223++ {
                                        fmt.Println(v5)
                                        var v224 int16 = 0
                                        v225 := GenS15{F0: fi217, F1: 2147483647}
                                        _ = fi223
                                        _ = v224
                                        _ = v225
                                    }
                                    v226 := GenS20{F0: 127}
                                    _ = v226
                                }
                            }
                        case 579:
                            for fi227 := 0; fi227 < 8; fi227++ {
                                _ = len(v1)
                                v228 := GenS22{F0: 889, F1: "hello", F2: 1}
                                v229 := GenS6{F0: -73, F1: -128, F2: 2}
                                _ = fi227
                                _ = v228
                                _ = v229
                            }
                        default:
                            v218 = v218
                        }
                        _ = fi217
                        _ = v218
                    }
                    var v230 uint = 65535
                    _ = v230
                }
            }
            v231 := GenS11{F0: -123}
            _ = v231
        }
    }
    var v232 float64 = v4
    if len(v3) > 0 {
        v3[0] = 42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v150:", v150)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v232:", v232)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v150
    _ = v2
    _ = v232
    _ = v3
    _ = v4
    _ = v5
}
