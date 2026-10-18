package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 int32; F1 int; F2 string }
type GenS4 struct { F0 bool; F1 uint16; F2 int8 }
type GenS5 struct { F0 int; F1 int64; F2 uint8 }
type GenS6 struct { F0 int64; F1 int16; F2 int }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint8; F1 int; F2 uint32 }
type GenS9 struct { F0 string; F1 int32 }
type GenS10 struct { F0 int8; F1 int16; F2 int8 }
type GenS11 struct { F0 uint16; F1 int; F2 uint16 }
type GenS12 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 string; F1 uint16; F2 int64 }
type GenS15 struct { F0 int; F1 int8 }
type GenS16 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS17 struct { F0 int16; F1 bool }
type GenS18 struct { F0 int }
type GenS19 struct { F0 int8; F1 string; F2 uint8 }
type GenS20 struct { F0 bool; F1 uint16 }
type GenS21 struct { F0 string }
type GenS22 struct { F0 uint32 }
type GenS23 struct { F0 string; F1 int32 }
type GenS24 struct { F0 string; F1 int64 }
type GenS25 struct { F0 int32; F1 bool; F2 int16 }
func genFn0(p0 int32) uint16 {
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := []int{(fi5 + (fi5 & 32)), 348, fi5, (fi5 % ((fi5) | 1))}
        v7 := GenS1{F0: (false || false)}
        _ = fi5
        _ = v6
        _ = v7
    }
    switch ((("a" + "foo bar") + "a") + "a") {
    case "a":
        var v8 uint = 0
        _ = v8
    case "hello":
        if (!((0 == 911) || (("foo bar" + "a") >= "hello"))) {
            v9 := []int{-55, -86}
            for fi10 := 0; fi10 < 8; fi10++ {
                switch (fi10 - (-(fi10))) {
                case 32768:
                    v11 := GenS1{F0: (p0 <= (p0 % ((p0) | 1)))}
                    _ = v11
                case 9223372036854775807:
                    p0 -= -2147483648
                case 127:
                    p0 -= (-(p0))
                }
                _ = fi10
            }
            var v12 int32 = p0
            _ = v12
            _ = v9
        }
    }
    if false {
        p0 = 1
        p0 -= (p0 >> ((18446744073709551615) & 63))
        if (636 >= -629) {
            v13 := GenS1{F0: (65535 > 18446744073709551615)}
            v14 := []int{-1, 128, 128}
            for fi15 := 0; fi15 < 9; fi15++ {
                p0 = p0
                v16 := GenS1{F0: true}
                v17 := GenS1{F0: (float64((float32(1e10) * (float32(3.14159) * -2.25))) != 1.5)}
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = v13
            _ = v14
        }
    } else {
        p0 -= p0
        var v18 float32 = float32(-2.25)
        _ = v18
    }
    var v19 float64 = 1.5
    v20 := GenS1{F0: (false && true)}
    _ = p0
    _ = v19
    _ = v20
    return 2
}
func genFn1(p0 int64) string {
    for fi21 := 0; fi21 < 13; fi21++ {
        v22 := GenS1{F0: false}
        v23 := GenS1{F0: true}
        _ = fi21
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        fmt.Println(fi24)
        if (!false) {
            if (true && (367 >= 9223372036854775807)) {
                for fi25 := 0; fi25 < 1; fi25++ {
                    var v26 uint16 = 739
                    _ = fi25
                    _ = v26
                }
            }
            p0 = p0
        }
        _ = fi24
    }
    v27 := GenS2{F0: 104, F1: 42}
    for fi28 := 0; fi28 < 5; fi28++ {
        p0 -= p0
        switch "a" {
        case "":
            var v29 int32 = -840
            _ = v29
        case "hello":
            v30 := GenS3{F0: 0, F1: 2147483647, F2: "\n\t"}
            _ = v30
        default:
            p0 = (p0 | p0)
        }
        _ = fi28
    }
    var v31 float64 = 0.0
    _ = p0
    _ = v27
    _ = v31
    return ""
}

func main() {
    var v1 string = "a"
    v2 := make(map[int]uint32)
    v3 := make(map[int]int)
    var v4 uint64 = 42
    if ((428 < 456) || true) {
        if true {
            for fi32 := 0; fi32 < 2; fi32++ {
                for fi33 := 0; fi33 < 8; fi33++ {
                    fmt.Println(v1, v1, v4)
                    v4 = (v4 - (v4 * (v4 >> ((1) & 63))))
                    _ = fi33
                }
                _ = fi32
            }
        } else {
            var v34 int64 = int64(v4)
            var v35 uint8 = 255
            _ = v34
            _ = v35
        }
    } else {
        for fi36 := 0; fi36 < 6; fi36++ {
            switch v1 {
            case "a":
                v1 = "a"
            case "foo bar":
                var v37 uint64 = (^(v4))
                _ = v37
            default:
                for fi38 := 0; fi38 < 10; fi38++ {
                    var v39 float32 = (1.5 - (float32(1e10) * 1.5))
                    v39 += v39
                    v1 = v1
                    if false {
                        _ = len(v2)
                        v4 = 4294967295
                        delete(v2, (fi36 % ((fi38) | 1)))
                    }
                    _ = fi38
                    _ = v39
                }
            }
            var v40 int8 = -107
            v4 = 80
            _ = fi36
            _ = v40
        }
        for fi41 := 0; fi41 < 18; fi41++ {
            for fi42 := 0; fi42 < 9; fi42++ {
                var v43 bool = false
                fmt.Println(v1, v1)
                var v44 bool = (!true)
                switch 14 {
                case 1:
                    _ = len(v2)
                case -910:
                    fmt.Println(fi42, v4)
                }
                _ = fi42
                _ = v43
                _ = v44
            }
            v45 := []int{(fi41 | (fi41 + fi41))}
            var v46 int = 127
            _ = fi41
            _ = v45
            _ = v46
        }
    }
    _ = len(v3)
    var v47 bool = false
    switch 32767 {
    case 0:
        fmt.Println(v1)
    case 9223372036854775807:
        v2[-2147483648] = 396
    case 32767:
        if (-918 == 0) {
            for fi48 := 0; fi48 < 7; fi48++ {
                if true {
                    v1 = "hello"
                    delete(v2, (fi48 << ((65535) & 63)))
                    for fi49 := 0; fi49 < 4; fi49++ {
                        v50 := GenS4{F0: ((-(fi48)) > (fi48 & (fi49 >> ((2) & 63)))), F1: 2, F2: 1}
                        v51 := []int{fi49, (fi48 - (fi48 & 250))}
                        var v52 bool = true
                        _ = fi49
                        _ = v50
                        _ = v51
                        _ = v52
                    }
                } else {
                    fmt.Println(fi48, v4)
                }
                var v53 uint16 = 65535
                v4 = v4
                _ = fi48
                _ = v53
            }
        } else {
            for fi54 := 0; fi54 < 12; fi54++ {
                if v47 {
                    fmt.Println(v47)
                }
                if (!((v47 || true) || ((!(((1e10 / (3.14159 / (-2.25 / -2.25))) != float64((0.0 + 0.1))) || false)) || (v47 && (9223372036854775807 >= -192))))) {
                    v3[(fi54 + fi54)] = fi54
                } else {
                    if v47 {
                        var v55 float32 = 0.0
                        _ = v55
                    } else {
                        for fi56 := 0; fi56 < 12; fi56++ {
                            var v57 int32 = 1
                            _ = fi56
                            _ = v57
                        }
                        v47 = true
                    }
                }
                _ = fi54
            }
            var v58 uint32 = 851
            if ((3.14159 * 1e10) >= ((0.0 + -2.25) - 3.14159)) {
                if (!v47) {
                    var v59 uint8 = 54
                    var v60 string = (v1 + (v1 + v1))
                    _ = v59
                    _ = v60
                } else {
                    v4 -= 18446744073709551615
                }
                var v61 string = ((v1 + v1) + ("" + v1))
                _ = v61
            }
            _ = v58
        }
    }
    v47 = v47
    v4 -= (v4 / ((v4) | 1))
    var v62 float64 = (((1e10 + 3.14159) - -2.25) / float64(1e10))
    v4 = 18446744073709551615
    v62 += v62
    var v63 uint32 = 9
    for fi64 := 0; fi64 < 16; fi64++ {
        if (((!((-1 < -1) || (742 <= 2))) && (1 == 4294967295)) && ((!(!(v47 && (0.1 >= 1e10)))) && (!false))) {
            v4 = 0
        }
        _ = fi64
    }
    for fi65 := 0; fi65 < 11; fi65++ {
        switch "\n\t" {
        case "\n\t":
            for fi66 := 0; fi66 < 6; fi66++ {
                v2[(fi65 % ((-432) | 1))] = (v63 / ((v63) | 1))
                _ = fi66
            }
        case "foo bar":
            switch "" {
            case "foo bar":
                if v47 {
                    if (float64((1.5 + 0.1)) <= (v62 - v62)) {
                        switch "" {
                        case "a":
                            v67 := GenS5{F0: fi65, F1: 0, F2: 225}
                            _ = v67
                        case "hello":
                            if false {
                                switch "" {
                                case "\n\t":
                                    var v68 bool = (0 != 0)
                                    _ = v68
                                case "":
                                    for fi69 := 0; fi69 < 4; fi69++ {
                                        v4 -= (^(v4))
                                        fmt.Println(v47, v63)
                                        v70 := GenS4{F0: v47, F1: 2, F2: 127}
                                        _ = fi69
                                        _ = v70
                                    }
                                }
                                v47 = (32767 > -640)
                            } else {
                                if false {
                                    var v71 int = fi65
                                    _ = v71
                                }
                                var v72 uint = 420
                                if v47 {
                                    switch v1 {
                                    case "hello":
                                        for fi73 := 0; fi73 < 16; fi73++ {
                                            if v47 {
                                                if v47 {
                                                    v74 := GenS3{F0: -2147483648, F1: fi65, F2: v1}
                                                    v75 := []int{32767, int(v4), (fi73 >> ((v72) & 63)), 128}
                                                    _ = v74
                                                    _ = v75
                                                }
                                                var v76 uint32 = v63
                                                if true {
                                                    v77 := GenS6{F0: 42, F1: int16(v4), F2: 32767}
                                                    switch -341 {
                                                    case 263:
                                                        var v78 uint64 = 1
                                                        _ = v78
                                                    case -1:
                                                        v79 := GenS7{F0: 127}
                                                        _ = v79
                                                    }
                                                    _ = v77
                                                } else {
                                                    var v80 uint32 = v76
                                                    fmt.Println(v1)
                                                    v81 := GenS6{F0: 21, F1: -474, F2: 127}
                                                    _ = v80
                                                    _ = v81
                                                }
                                                var v82 uint16 = 42
                                                _ = v76
                                                _ = v82
                                            } else {
                                                switch fi65 {
                                                case 128:
                                                    if ("a" > (v1 + (v1 + ("a" + (("a" + v1) + v1))))) {
                                                        v62 += v62
                                                    } else {
                                                        if v47 {
                                                            var v83 int = fi65
                                                            var v84 string = (v1 + v1)
                                                            v85 := v3[(fi73 % ((fi73) | 1))]
                                                            v72 = (v72 % (((v72 << ((0) & 63))) | 1))
                                                            fmt.Println(fi65)
                                                            _ = v83
                                                            _ = v84
                                                            _ = v85
                                                        } else {
                                                            if (v47 && v47) {
                                                                var v86 uint64 = 4294967295
                                                                _ = v86
                                                            } else {
                                                                if false {
                                                                    v72 += 1
                                                                    if false {
                                                                        var v87 float32 = 1e10
                                                                        fmt.Println(v63, v63, v72)
                                                                        v88 := GenS3{F0: -667, F1: (fi73 | 127), F2: v1}
                                                                        _ = v87
                                                                        _ = v88
                                                                    } else {
                                                                        fmt.Println(v72)
                                                                    }
                                                                    var v89 uint8 = 2
                                                                    _ = v89
                                                                }
                                                            }
                                                        }
                                                        var v90 int64 = 42
                                                        _ = v90
                                                    }
                                                case -395:
                                                    if (!(!(!true))) {
                                                        var v91 uint8 = 127
                                                        var v92 uint8 = uint8(v4)
                                                        var v93 uint16 = 942
                                                        v3[(fi73 | (fi65 * (fi73 | fi65)))] = 32768
                                                        if (18446744073709551615 != (v72 + 0)) {
                                                            var v94 float32 = -2.25
                                                            var v95 bool = (v47 && false)
                                                            fmt.Println(v72, v63)
                                                            if (-518 >= 441) {
                                                                if v47 {
                                                                    fmt.Println(v93, v62)
                                                                    switch fi65 {
                                                                    case -2147483648:
                                                                        fmt.Println(fi73, v93)
                                                                    case 127:
                                                                        v96 := GenS7{F0: 0}
                                                                        _ = v96
                                                                    }
                                                                    v97 := []int{(^(fi73)), 9223372036854775807, (-(fi65)), fi65}
                                                                    _ = v97
                                                                } else {
                                                                    v98 := GenS7{F0: int8(fi73)}
                                                                    if true {
                                                                        if v47 {
                                                                            if true {
                                                                                switch v1 {
                                                                                case "\n\t":
                                                                                    fmt.Println(v63)
                                                                                case "a":
                                                                                    v93 = genFn0(664)
                                                                                case "foo bar":
                                                                                    _ = len(v2)
                                                                                }
                                                                                if true {
                                                                                    var v99 string = "hello"
                                                                                    switch v1 {
                                                                                    case "foo bar":
                                                                                        if false {
                                                                                            fmt.Println(v62, v1, v4)
                                                                                            v100 := []int{-928, -5, fi73, fi73}
                                                                                            var v101 uint8 = v91
                                                                                            _ = v100
                                                                                            _ = v101
                                                                                        } else {
                                                                                            v91 += v92
                                                                                            fmt.Println(v93)
                                                                                        }
                                                                                    case "hello":
                                                                                        v72 = v72
                                                                                        v102 := GenS3{F0: -1, F1: 4294967295, F2: v1}
                                                                                        _ = v102
                                                                                    case "\n\t":
                                                                                        var v103 int8 = 127
                                                                                        _ = v103
                                                                                    default:
                                                                                        var v104 uint32 = (v63 | 65535)
                                                                                        _ = v104
                                                                                    }
                                                                                    _ = v99
                                                                                } else {
                                                                                    var v105 float32 = 1e10
                                                                                    v91 = (v91 / (((v92 >> ((uint(v4)) & 63))) | 1))
                                                                                    fmt.Println(v1, v92, fi65)
                                                                                    _ = v105
                                                                                }
                                                                                var v106 uint64 = v4
                                                                                _ = v106
                                                                            } else {
                                                                                v92 -= (v91 >> (((v72 + (^(v72)))) & 63))
                                                                                v107 := GenS2{F0: uint8(v72), F1: (v93 >> ((v72) & 63))}
                                                                                _ = v107
                                                                            }
                                                                            if (v94 >= (v94 / v94)) {
                                                                                v94 -= v94
                                                                            }
                                                                        }
                                                                        v47 = v47
                                                                        var v108 int32 = -1
                                                                        _ = v108
                                                                    } else {
                                                                        switch fi73 {
                                                                        case -2147483648:
                                                                            var v109 int = 1
                                                                            _ = v109
                                                                        case -9223372036854775808:
                                                                            var v110 float32 = 0.1
                                                                            _ = v110
                                                                        default:
                                                                            delete(v3, 1)
                                                                        }
                                                                    }
                                                                    v95 = v95
                                                                    _ = v98
                                                                }
                                                            } else {
                                                                fmt.Println(v47)
                                                            }
                                                            if (-9223372036854775808 <= 0) {
                                                                if (uint64(v92) >= v4) {
                                                                    if true {
                                                                        var v111 int8 = 111
                                                                        _ = v111
                                                                    } else {
                                                                        v112 := []int{(fi65 % ((fi65) | 1)), (fi73 % ((fi65) | 1)), (fi65 + fi65), (fi65 / (((fi73 | fi65)) | 1))}
                                                                        var v113 int32 = 0
                                                                        switch v1 {
                                                                        case "hello":
                                                                            if (v113 >= (v113 | v113)) {
                                                                                v112 = append(v112, fi73)
                                                                                v114 := v3[fi73]
                                                                                _ = v114
                                                                            } else {
                                                                                var v115 uint64 = v4
                                                                                _ = v115
                                                                            }
                                                                            v62 -= 0.1
                                                                            v91 += (v92 << (((^(v72))) & 63))
                                                                        case "\n\t":
                                                                            var v116 uint32 = v63
                                                                            _ = v116
                                                                        case "a":
                                                                            var v117 int = 818
                                                                            _ = v117
                                                                        default:
                                                                            switch (^(fi65)) {
                                                                            case 9223372036854775807:
                                                                                switch (fi73 % (((fi73 / ((2147483647) | 1))) | 1)) {
                                                                                case -726:
                                                                                    switch "" {
                                                                                    case "foo bar":
                                                                                        var v118 uint = v72
                                                                                        _ = v118
                                                                                    }
                                                                                case -447:
                                                                                    fmt.Println(v95, v93)
                                                                                case 32768:
                                                                                    v91 = v92
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v112
                                                                        _ = v113
                                                                    }
                                                                    switch "a" {
                                                                    case "hello":
                                                                        var v119 int8 = 0
                                                                        _ = v119
                                                                    case "":
                                                                        switch (fi65 ^ (fi65 / (((-(fi73))) | 1))) {
                                                                        case 32768:
                                                                            switch v1 {
                                                                            case "foo bar":
                                                                                var v120 float32 = ((float32(((1e10 + ((((v62 - 0.0) - v62) - v62) * ((-2.25 + v62) / v62))) + 0.1)) + (1e10 / float32(((-2.25 / (v62 / -2.25)) * v62)))) / 1.0)
                                                                                _ = v120
                                                                            }
                                                                        default:
                                                                            v121 := v2[fi65]
                                                                            _ = v121
                                                                        }
                                                                    case "foo bar":
                                                                        var v122 int64 = -789
                                                                        _ = v122
                                                                    default:
                                                                        if (fi65 >= (fi73 | fi73)) {
                                                                            v62 += (v62 + v62)
                                                                            v123 := GenS2{F0: (^(v91)), F1: (v93 | 42)}
                                                                            var v124 int16 = 1
                                                                            _ = v123
                                                                            _ = v124
                                                                        } else {
                                                                            v125 := []int{fi73, int(v93)}
                                                                            switch (fi65 << (((v72 | v72)) & 63)) {
                                                                            case -673:
                                                                                switch (("a" + "foo bar") + (v1 + v1)) {
                                                                                case "a":
                                                                                    fmt.Println(v47)
                                                                                case "\n\t":
                                                                                    v95 = false
                                                                                    v126 := GenS2{F0: (^(v92)), F1: (v93 ^ uint16(v4))}
                                                                                    v95 = v95
                                                                                    v95 = ((fi65 ^ (fi65 - (^(fi73)))) == (^(fi65)))
                                                                                    _ = v126
                                                                                default:
                                                                                    v127 := GenS7{F0: 127}
                                                                                    _ = v127
                                                                                }
                                                                            case 124:
                                                                                switch v1 {
                                                                                case "hello":
                                                                                    var v128 bool = v95
                                                                                    _ = v128
                                                                                default:
                                                                                    v93 += (^(v93))
                                                                                }
                                                                            default:
                                                                                v95 = true
                                                                            }
                                                                            var v129 string = ((v1 + (v1 + v1)) + (((v1 + (v1 + v1)) + (v1 + v1)) + (v1 + v1)))
                                                                            _ = v125
                                                                            _ = v129
                                                                        }
                                                                    }
                                                                    switch v1 {
                                                                    case "foo bar":
                                                                        v47 = true
                                                                    case "a":
                                                                        fmt.Println(v93, v62, v47)
                                                                    case "\n\t":
                                                                        if false {
                                                                            if v95 {
                                                                                if ((v62 * 1e10) <= (1.5 + v62)) {
                                                                                    fmt.Println(v91, v72, v4)
                                                                                    var v130 bool = v95
                                                                                    if false {
                                                                                        if v95 {
                                                                                            if (691 != v72) {
                                                                                                switch -985 {
                                                                                                case 4294967295:
                                                                                                    var v131 bool = false
                                                                                                    _ = v131
                                                                                                default:
                                                                                                    v94 -= 1e10
                                                                                                }
                                                                                                v132 := []int{fi65, 127, fi73, 32768}
                                                                                                _ = v132
                                                                                            }
                                                                                            switch (fi65 + (-(fi65))) {
                                                                                            case -2147483648:
                                                                                                if (-9223372036854775808 != -9223372036854775808) {
                                                                                                    switch "" {
                                                                                                    case "foo bar":
                                                                                                        if v47 {
                                                                                                            var v133 float64 = (v62 / float64(v94))
                                                                                                            switch fi65 {
                                                                                                            case -845:
                                                                                                                v134 := []int{fi73, fi73, (fi73 << ((v72) & 63)), fi65}
                                                                                                                _ = v134
                                                                                                            case 859:
                                                                                                                var v135 int = (fi65 / (((fi73 - 9223372036854775807)) | 1))
                                                                                                                _ = v135
                                                                                                            default:
                                                                                                                switch v1 {
                                                                                                                case "\n\t":
                                                                                                                    if (v94 < float32(1.5)) {
                                                                                                                        if false {
                                                                                                                            v94 += 3.14159
                                                                                                                            v93 = 1
                                                                                                                        }
                                                                                                                        if v95 {
                                                                                                                            v133 = v62
                                                                                                                            delete(v2, 2147483647)
                                                                                                                            var v136 string = "\n\t"
                                                                                                                            _ = v136
                                                                                                                        }
                                                                                                                        switch 42 {
                                                                                                                        case 32767:
                                                                                                                            var v137 int64 = 9223372036854775807
                                                                                                                            _ = v137
                                                                                                                        default:
                                                                                                                            v62 = -2.25
                                                                                                                        }
                                                                                                                    }
                                                                                                                default:
                                                                                                                    if (((-(fi65)) <= (^(fi65))) || (!(!(32768 == (fi65 * fi73))))) {
                                                                                                                        fmt.Println(v47, v63)
                                                                                                                        var v138 int8 = 1
                                                                                                                        _ = v138
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                            var v139 uint64 = uint64(v63)
                                                                                                            _ = v133
                                                                                                            _ = v139
                                                                                                        } else {
                                                                                                            v140 := GenS5{F0: 4, F1: -9223372036854775808, F2: 128}
                                                                                                            fmt.Println(v94, v130)
                                                                                                            if true {
                                                                                                                var v141 uint = v72
                                                                                                                var v142 uint32 = (^(v63))
                                                                                                                _ = v141
                                                                                                                _ = v142
                                                                                                            } else {
                                                                                                                var v143 uint32 = (v63 % (((^(v63))) | 1))
                                                                                                                switch v1 {
                                                                                                                case "foo bar":
                                                                                                                    if (127 > (fi65 / ((fi65) | 1))) {
                                                                                                                        switch "" {
                                                                                                                        case "hello":
                                                                                                                            v95 = v47
                                                                                                                            fmt.Println(v72, v72, v93)
                                                                                                                            v144 := GenS5{F0: fi73, F1: -9223372036854775808, F2: v91}
                                                                                                                            _ = v144
                                                                                                                        case "":
                                                                                                                            v2[(^(fi65))] = (^(v143))
                                                                                                                        case "foo bar":
                                                                                                                            fmt.Println(v130)
                                                                                                                        }
                                                                                                                        if (!((42 <= -1) || (0 <= -1))) {
                                                                                                                            v47 = (true && v130)
                                                                                                                        }
                                                                                                                        v145 := GenS2{F0: v92, F1: (v93 | genFn0(232))}
                                                                                                                        _ = v145
                                                                                                                    } else {
                                                                                                                        fmt.Println(v63, v94, v1)
                                                                                                                        fmt.Println(v4)
                                                                                                                        v91 = v92
                                                                                                                    }
                                                                                                                    v146 := GenS3{F0: 42, F1: -1, F2: v1}
                                                                                                                    var v147 string = v1
                                                                                                                    v148 := v2[(fi65 << (((v72 % ((v72) | 1))) & 63))]
                                                                                                                    var v149 float64 = float64((v94 * v94))
                                                                                                                    v150 := GenS8{F0: (v91 & v92), F1: (fi73 % (((fi65 ^ -189)) | 1)), F2: (v143 | 65535)}
                                                                                                                    _ = v146
                                                                                                                    _ = v147
                                                                                                                    _ = v148
                                                                                                                    _ = v149
                                                                                                                    _ = v150
                                                                                                                case "":
                                                                                                                    fmt.Println(v1)
                                                                                                                case "\n\t":
                                                                                                                    fmt.Println(v72, v93)
                                                                                                                }
                                                                                                                if false {
                                                                                                                    var v151 int = (fi73 & (fi73 + fi65))
                                                                                                                    _ = v151
                                                                                                                } else {
                                                                                                                    var v152 int = -551
                                                                                                                    fmt.Println(v4, v130, v62)
                                                                                                                    var v153 int16 = 42
                                                                                                                    var v154 float64 = (v62 * (0.1 * ((((((v62 / 0.1) + (v62 / v62)) - (v62 - (1e10 - v62))) + 3.14159) + v62) * 0.0)))
                                                                                                                    v1 = (v1 + "\n\t")
                                                                                                                    fmt.Println(v63, v92)
                                                                                                                    v155 := GenS6{F0: -1, F1: v153, F2: (v152 + fi73)}
                                                                                                                    fmt.Println(v152)
                                                                                                                    _ = v152
                                                                                                                    _ = v153
                                                                                                                    _ = v154
                                                                                                                    _ = v155
                                                                                                                }
                                                                                                                _ = v143
                                                                                                            }
                                                                                                            _ = v140
                                                                                                        }
                                                                                                    case "":
                                                                                                        if true {
                                                                                                            v72 -= v72
                                                                                                            v91 += 163
                                                                                                        } else {
                                                                                                            var v156 uint16 = (^(v93))
                                                                                                            v130 = false
                                                                                                            v156 -= v156
                                                                                                            switch (fi65 & fi65) {
                                                                                                            case -318:
                                                                                                                if v47 {
                                                                                                                    v157 := GenS9{F0: (v1 + v1), F1: -940}
                                                                                                                    _ = v157
                                                                                                                } else {
                                                                                                                    v63 -= 42
                                                                                                                    if (0 < -402) {
                                                                                                                        var v158 float32 = 0.1
                                                                                                                        var v159 uint8 = (v92 ^ v92)
                                                                                                                        fmt.Println(v63, v93)
                                                                                                                        _ = v158
                                                                                                                        _ = v159
                                                                                                                    } else {
                                                                                                                        v160 := []int{}
                                                                                                                        _ = v160
                                                                                                                    }
                                                                                                                }
                                                                                                            default:
                                                                                                                v4 = (v4 % ((18446744073709551615) | 1))
                                                                                                            }
                                                                                                            _ = v156
                                                                                                        }
                                                                                                    case "hello":
                                                                                                        v94 = (v94 + v94)
                                                                                                    }
                                                                                                } else {
                                                                                                    if false {
                                                                                                        if true {
                                                                                                            fmt.Println(fi65, v47, v4)
                                                                                                            if v47 {
                                                                                                                if false {
                                                                                                                    v47 = (1 < (^(v63)))
                                                                                                                } else {
                                                                                                                    var v161 uint8 = (v91 ^ uint8(v93))
                                                                                                                    var v162 int8 = 100
                                                                                                                    var v163 int8 = 0
                                                                                                                    _ = v161
                                                                                                                    _ = v162
                                                                                                                    _ = v163
                                                                                                                }
                                                                                                                var v164 int = fi73
                                                                                                                _ = v164
                                                                                                            }
                                                                                                        }
                                                                                                    } else {
                                                                                                        var v165 int16 = 0
                                                                                                        switch fi73 {
                                                                                                        case 9223372036854775807:
                                                                                                            v166 := GenS10{F0: -83, F1: 1, F2: 0}
                                                                                                            _ = v166
                                                                                                        case 32767:
                                                                                                            switch (v1 + "hello") {
                                                                                                            case "foo bar":
                                                                                                                var v167 float32 = -2.25
                                                                                                                _ = v167
                                                                                                            default:
                                                                                                                v4 += uint64(v63)
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v165
                                                                                                    }
                                                                                                    if true {
                                                                                                        var v168 int16 = 0
                                                                                                        _ = v168
                                                                                                    } else {
                                                                                                        var v169 uint8 = 1
                                                                                                        _ = v169
                                                                                                    }
                                                                                                }
                                                                                            case 128:
                                                                                                switch (fi65 % ((fi65) | 1)) {
                                                                                                case -2147483648:
                                                                                                    var v170 int16 = 1
                                                                                                    _ = v170
                                                                                                default:
                                                                                                    v171 := v3[(fi73 | 0)]
                                                                                                    _ = v171
                                                                                                }
                                                                                            case -321:
                                                                                                v130 = v47
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    _ = v130
                                                                                } else {
                                                                                    if (true && v95) {
                                                                                        var v172 int8 = 1
                                                                                        if (v63 == uint32(v91)) {
                                                                                            var v173 int64 = 738
                                                                                            fmt.Println(v95, fi65, v63)
                                                                                            var v174 uint64 = (^(v4))
                                                                                            _ = v173
                                                                                            _ = v174
                                                                                        } else {
                                                                                            var v175 int64 = 42
                                                                                            v92 = uint8(v4)
                                                                                            switch ("foo bar" + "foo bar") {
                                                                                            case "hello":
                                                                                                if v47 {
                                                                                                    v176 := GenS5{F0: (-(fi73)), F1: v175, F2: v92}
                                                                                                    _ = v176
                                                                                                }
                                                                                            case "a":
                                                                                                v177 := GenS11{F0: (v93 / (((^(v93))) | 1)), F1: (-(fi73)), F2: 255}
                                                                                                _ = v177
                                                                                            case "\n\t":
                                                                                                fmt.Println(v62, v47, v95)
                                                                                            }
                                                                                            var v178 int8 = 100
                                                                                            var v179 int16 = 1
                                                                                            _ = v175
                                                                                            _ = v178
                                                                                            _ = v179
                                                                                        }
                                                                                        _ = v172
                                                                                    }
                                                                                    v91 -= (v91 >> (((^(v72))) & 63))
                                                                                }
                                                                                var v180 uint64 = v4
                                                                                _ = v180
                                                                            }
                                                                            v181 := GenS7{F0: 42}
                                                                            _ = v181
                                                                        } else {
                                                                            if v47 {
                                                                                v182 := GenS5{F0: fi65, F1: 0, F2: 42}
                                                                                v47 = ((true || v47) || (-358 < -9223372036854775808))
                                                                                _ = v182
                                                                            } else {
                                                                                var v183 uint32 = 905
                                                                                if (!v47) {
                                                                                    v91 += v92
                                                                                    switch "hello" {
                                                                                    case "a":
                                                                                        v184 := GenS12{F0: -1, F1: v183, F2: 127}
                                                                                        _ = v184
                                                                                    case "":
                                                                                        fmt.Println(fi73, v94)
                                                                                    case "foo bar":
                                                                                        v91 = 154
                                                                                        fmt.Println(v92, v95, v4)
                                                                                        v185 := []int{(fi65 * (^(fi65))), fi65}
                                                                                        _ = v185
                                                                                    default:
                                                                                        fmt.Println(fi73, v62, v93)
                                                                                    }
                                                                                    v186 := v2[fi73]
                                                                                    _ = v186
                                                                                } else {
                                                                                    v91 = 1
                                                                                }
                                                                                var v187 uint8 = 127
                                                                                if (!(!false)) {
                                                                                    fmt.Println(fi65, v94, v187)
                                                                                } else {
                                                                                    var v188 uint64 = 4294967295
                                                                                    v189 := GenS13{F0: 0}
                                                                                    if v47 {
                                                                                        fmt.Println(fi65, fi65, v188)
                                                                                        var v190 uint = 65535
                                                                                        _ = v190
                                                                                    } else {
                                                                                        var v191 float64 = 0.1
                                                                                        var v192 int8 = -128
                                                                                        v193 := GenS7{F0: v192}
                                                                                        var v194 int32 = -678
                                                                                        _ = v191
                                                                                        _ = v192
                                                                                        _ = v193
                                                                                        _ = v194
                                                                                    }
                                                                                    v195 := GenS1{F0: (v183 > (v63 << ((v72) & 63)))}
                                                                                    v196 := GenS14{F0: (v1 + v1), F1: uint16(v92), F2: 0}
                                                                                    _ = v188
                                                                                    _ = v189
                                                                                    _ = v195
                                                                                    _ = v196
                                                                                }
                                                                                var v197 float64 = (v62 - ((v62 * v62) / ((v62 / v62) - 1.5)))
                                                                                _ = v183
                                                                                _ = v187
                                                                                _ = v197
                                                                            }
                                                                        }
                                                                    default:
                                                                        var v198 int16 = 32767
                                                                        _ = v198
                                                                    }
                                                                } else {
                                                                    switch ((v1 + v1) + "") {
                                                                    case "hello":
                                                                        switch (fi73 & (fi65 / (((^(fi73))) | 1))) {
                                                                        case 127:
                                                                            if (13 < 0) {
                                                                                v94 -= 1.5
                                                                                v199 := GenS15{F0: (fi73 >> ((255) & 63)), F1: -128}
                                                                                var v200 int16 = 906
                                                                                _ = v199
                                                                                _ = v200
                                                                            }
                                                                            if ((v93 * (v93 / ((v93) | 1))) <= 32767) {
                                                                                if true {
                                                                                    fmt.Println(v93, fi65)
                                                                                    v201 := GenS5{F0: (fi65 % ((-1) | 1)), F1: -1, F2: (v91 + v92)}
                                                                                    if v95 {
                                                                                        switch int(v63) {
                                                                                        case 4294967295:
                                                                                            var v202 uint16 = uint16(v92)
                                                                                            _ = v202
                                                                                        case 32767:
                                                                                            var v203 string = (("" + (v1 + "foo bar")) + "hello")
                                                                                            _ = v203
                                                                                        }
                                                                                    } else {
                                                                                        var v204 int8 = -128
                                                                                        v205 := GenS16{F0: v93, F1: (v72 ^ v72), F2: (v63 % ((881) | 1))}
                                                                                        v1 = v1
                                                                                        _ = v204
                                                                                        _ = v205
                                                                                    }
                                                                                    v92 = (v92 | 202)
                                                                                    v91 = (^(v92))
                                                                                    v47 = (1 <= 100)
                                                                                    v47 = (-9223372036854775808 <= 9223372036854775807)
                                                                                    var v206 uint16 = uint16(v92)
                                                                                    _ = v201
                                                                                    _ = v206
                                                                                } else {
                                                                                    if (false || (!((v72 + v72) <= (v72 & (^(v72)))))) {
                                                                                        var v207 string = ((v1 + (v1 + v1)) + ((((v1 + ("" + v1)) + v1) + v1) + v1))
                                                                                        switch v1 {
                                                                                        case "":
                                                                                            var v208 int = fi65
                                                                                            _ = v208
                                                                                        case "\n\t":
                                                                                            var v209 int32 = -1
                                                                                            var v210 uint = 535
                                                                                            v62 = v62
                                                                                            _ = v209
                                                                                            _ = v210
                                                                                        case "hello":
                                                                                            v91 = v91
                                                                                        default:
                                                                                            fmt.Println(v93, v207)
                                                                                        }
                                                                                        _ = v207
                                                                                    }
                                                                                    v47 = (!(!(-1 < 9223372036854775807)))
                                                                                }
                                                                            }
                                                                            if (!true) {
                                                                                v72 -= v72
                                                                                var v211 float64 = 1e10
                                                                                var v212 uint32 = (v63 | 65535)
                                                                                v93 = genFn0(2147483647)
                                                                                fmt.Println(v92)
                                                                                _ = v211
                                                                                _ = v212
                                                                            }
                                                                        }
                                                                    case "":
                                                                        var v213 int = fi73
                                                                        _ = v213
                                                                    case "a":
                                                                        v4 += v4
                                                                    default:
                                                                        var v214 float64 = float64(float32((-2.25 * (float64(v94) * v62))))
                                                                        _ = v214
                                                                    }
                                                                }
                                                                var v215 int = 127
                                                                if false {
                                                                    var v216 int = 4294967295
                                                                    _ = v216
                                                                }
                                                                _ = v215
                                                            } else {
                                                                fmt.Println(v63, v91)
                                                                var v217 uint32 = uint32(v93)
                                                                var v218 float32 = (0.1 * (v94 - 3.14159))
                                                                _ = v217
                                                                _ = v218
                                                            }
                                                            _ = v94
                                                            _ = v95
                                                        } else {
                                                            v219 := []int{fi73, (fi65 * (^(fi65))), fi73, fi73}
                                                            _ = v219
                                                        }
                                                        fmt.Println(v1, v1, v47)
                                                        fmt.Println(v4, v63, v4)
                                                        var v220 bool = (((0.0 * 1e10) * -2.25) < (((0.1 / 1.0) - float32(v62)) - 0.0))
                                                        var v221 uint8 = v92
                                                        _ = v220
                                                        _ = v221
                                                        _ = v91
                                                        _ = v92
                                                        _ = v93
                                                    }
                                                case 529:
                                                    var v222 uint = v72
                                                    _ = v222
                                                }
                                            }
                                            _ = fi73
                                        }
                                    case "\n\t":
                                        switch v1 {
                                        case "\n\t":
                                            v223 := GenS13{F0: 128}
                                            _ = v223
                                        }
                                    case "foo bar":
                                        if (true || v47) {
                                            if (!false) {
                                                fmt.Println(v1, v72, fi65)
                                            } else {
                                                var v224 uint64 = (v4 >> (((v72 * v72)) & 63))
                                                var v225 int64 = 42
                                                _ = v224
                                                _ = v225
                                            }
                                            v1 = v1
                                        }
                                    }
                                } else {
                                    if ((!(!v47)) && (!(v72 != v72))) {
                                        v226 := GenS17{F0: -32768, F1: (-90 >= 127)}
                                        fmt.Println(v47)
                                        _ = v226
                                    }
                                    for fi227 := 0; fi227 < 13; fi227++ {
                                        var v228 int64 = 270
                                        v229 := GenS4{F0: (v47 && v47), F1: 1, F2: 0}
                                        switch v1 {
                                        case "hello":
                                            v63 -= (^(v63))
                                        case "a":
                                            if (0.1 == float64((0.1 - (0.1 * (0.0 * ((((0.0 - (0.0 - -2.25)) + 1e10) * 0.1) * 0.0)))))) {
                                                if v47 {
                                                    if true {
                                                        fmt.Println(v63)
                                                        v1 = "\n\t"
                                                        v230 := GenS17{F0: 546, F1: (true || v47)}
                                                        v1 = v1
                                                        fmt.Println(v62, v47, v62)
                                                        _ = v230
                                                    }
                                                    v72 += v72
                                                } else {
                                                    v72 = v72
                                                }
                                                if (!((fi65 - (fi65 + 128)) > fi227)) {
                                                    v231 := []int{(fi227 & fi227), (fi65 % ((fi65) | 1)), fi65, (fi227 >> (((v72 ^ v72)) & 63))}
                                                    switch (fi227 & (fi65 & fi227)) {
                                                    case 32767:
                                                        v1 = v1
                                                    case 855:
                                                        v232 := GenS3{F0: 730, F1: fi65, F2: "hello"}
                                                        _ = v232
                                                    case 409:
                                                        v228 -= 9223372036854775807
                                                        fmt.Println(v228)
                                                        v63 -= v63
                                                    default:
                                                        v233 := GenS12{F0: 1, F1: (v63 ^ 483), F2: 255}
                                                        _ = v233
                                                    }
                                                    _ = v231
                                                } else {
                                                    v63 += uint32(v72)
                                                    fmt.Println(fi227, v62)
                                                    v63 = (v63 >> ((v72) & 63))
                                                }
                                                var v234 uint8 = 127
                                                _ = v234
                                            } else {
                                                if (-32768 != 42) {
                                                    if (1 >= v72) {
                                                        v228 -= v228
                                                        _ = len(v3)
                                                        if (!v47) {
                                                            v1 = v1
                                                            var v235 int16 = 42
                                                            _ = v235
                                                        }
                                                    }
                                                } else {
                                                    v236 := []int{(fi65 ^ -373), fi65, fi227, (fi227 * (fi65 * (fi65 ^ fi65)))}
                                                    v62 -= (v62 - (v62 + -2.25))
                                                    switch fi227 {
                                                    case -9223372036854775808:
                                                        v72 += (v72 << (((^(v72))) & 63))
                                                    default:
                                                        v237 := GenS12{F0: 1, F1: v63, F2: 246}
                                                        _ = v237
                                                    }
                                                    _ = v236
                                                }
                                                switch "hello" {
                                                case "hello":
                                                    if ((v72 % ((v72) | 1)) > (v72 % (((v72 | (v72 + (v72 | 65535)))) | 1))) {
                                                        _ = len(v2)
                                                        delete(v2, (fi65 + fi65))
                                                    } else {
                                                        v238 := GenS18{F0: fi227}
                                                        if (v47 && (v47 || v47)) {
                                                            v62 += 0.1
                                                            switch (fi227 + fi65) {
                                                            case 1:
                                                                if (!v47) {
                                                                    var v239 int32 = 403
                                                                    v240 := []int{(fi227 * fi227), -9223372036854775808}
                                                                    v241 := GenS19{F0: -128, F1: v1, F2: 35}
                                                                    _ = v239
                                                                    _ = v240
                                                                    _ = v241
                                                                } else {
                                                                    var v242 int64 = v228
                                                                    delete(v3, (fi227 << ((42) & 63)))
                                                                    var v243 uint8 = 2
                                                                    _ = v242
                                                                    _ = v243
                                                                }
                                                            case 128:
                                                                v244 := v3[fi65]
                                                                _ = v244
                                                            case 0:
                                                                if v47 {
                                                                    var v245 uint16 = 65535
                                                                    v2[-1] = (v63 >> (((v72 * v72)) & 63))
                                                                    _ = v245
                                                                } else {
                                                                    var v246 int64 = (v228 + v228)
                                                                    v247 := []int{(fi65 | (-(fi227))), fi227, 128, 32768}
                                                                    fmt.Println(v47)
                                                                    _ = v246
                                                                    _ = v247
                                                                }
                                                            default:
                                                                _ = len(v3)
                                                            }
                                                        } else {
                                                            fmt.Println(v63, v4, v4)
                                                            v3[992] = 32768
                                                            v248 := GenS20{F0: (-969 <= -2147483648), F1: 2}
                                                            _ = v248
                                                        }
                                                        _ = v238
                                                    }
                                                case "\n\t":
                                                    fmt.Println(fi227, v72)
                                                case "foo bar":
                                                    v63 = v63
                                                default:
                                                    v228 = (v228 ^ v228)
                                                }
                                            }
                                        }
                                        var v249 bool = (v47 || (!v47))
                                        _ = fi227
                                        _ = v228
                                        _ = v229
                                        _ = v249
                                    }
                                }
                                _ = v72
                            }
                        }
                    } else {
                        if true {
                            var v250 uint16 = uint16(v63)
                            for fi251 := 0; fi251 < 3; fi251++ {
                                v47 = (((31 <= 42) || (0 <= (v63 / (((v63 & v63)) | 1)))) || v47)
                                _ = fi251
                            }
                            fmt.Println(v47)
                            _ = v250
                        } else {
                            switch fi65 {
                            case 0:
                                if (!v47) {
                                    fmt.Println(v4, v47)
                                    fmt.Println(v47, v4)
                                    v47 = (0 > 16)
                                }
                            }
                        }
                    }
                }
            case "\n\t":
                if (-2.25 < v62) {
                    if v47 {
                        fmt.Println(v47, v47)
                        if ((v62 - ((v62 / v62) - v62)) >= (v62 + v62)) {
                            v4 = v4
                            if (255 != 1) {
                                for fi252 := 0; fi252 < 4; fi252++ {
                                    var v253 uint8 = 2
                                    var v254 int64 = 509
                                    _ = fi252
                                    _ = v253
                                    _ = v254
                                }
                                var v255 int32 = 511
                                v63 = (v63 << ((132) & 63))
                                _ = v255
                            }
                            fmt.Println(v1, v63, v4)
                        }
                        for fi256 := 0; fi256 < 3; fi256++ {
                            switch fi65 {
                            case -9223372036854775808:
                                v257 := GenS16{F0: 61, F1: 255, F2: v63}
                                _ = v257
                            case 605:
                                v258 := GenS21{F0: ("hello" + v1)}
                                _ = v258
                            }
                            v259 := GenS11{F0: genFn0(42), F1: (fi256 % ((32767) | 1)), F2: 0}
                            v62 -= v62
                            _ = fi256
                            _ = v259
                        }
                    } else {
                        v260 := v2[fi65]
                        if v47 {
                            switch ("\n\t" + v1) {
                            case "":
                                var v261 uint16 = 722
                                _ = v261
                            case "hello":
                                switch -9223372036854775808 {
                                case -877:
                                    v262 := []int{fi65}
                                    _ = v262
                                case 127:
                                    v260 = 1
                                case -636:
                                    v4 -= (v4 * (v4 * 1))
                                default:
                                    var v263 int8 = -1
                                    _ = v263
                                }
                            }
                            v2[fi65] = (v260 ^ v63)
                        }
                        v264 := GenS22{F0: 42}
                        _ = v260
                        _ = v264
                    }
                    switch (fi65 ^ fi65) {
                    case -105:
                        if v47 {
                            if v47 {
                                v47 = v47
                                switch fi65 {
                                case 876:
                                    switch 127 {
                                    case 9223372036854775807:
                                        var v265 uint64 = v4
                                        _ = v265
                                    }
                                }
                            }
                            var v266 int32 = -2147483648
                            v267 := GenS23{F0: (("a" + "a") + (((v1 + "a") + v1) + (v1 + ""))), F1: v266}
                            _ = v266
                            _ = v267
                        }
                    case 4294967295:
                        v268 := []int{fi65, fi65}
                        _ = v268
                    }
                } else {
                    var v269 int64 = 0
                    if ((!true) || ((0.0 * -2.25) < ((1e10 / 1.5) + float32(v62)))) {
                        v270 := []int{}
                        _ = len(v2)
                        for fi271 := 0; fi271 < 2; fi271++ {
                            var v272 uint32 = 440
                            v273 := GenS24{F0: v1, F1: v269}
                            v272 -= v63
                            fmt.Println(v1, fi65)
                            _ = fi271
                            _ = v272
                            _ = v273
                        }
                        _ = v270
                    } else {
                        if (v269 >= (^(v269))) {
                            var v274 float64 = (v62 + ((3.14159 * 0.1) * (v62 + v62)))
                            _ = v274
                        }
                    }
                    _ = v269
                }
            }
        case "":
            for fi275 := 0; fi275 < 11; fi275++ {
                v276 := v3[(fi65 / ((fi65) | 1))]
                v4 += (v4 & v4)
                v276 += fi275
                _ = fi275
                _ = v276
            }
        default:
            for fi277 := 0; fi277 < 9; fi277++ {
                var v278 int16 = -32768
                fmt.Println(v47)
                v4 += (v4 + 69)
                _ = fi277
                _ = v278
            }
        }
        var v279 uint32 = (^(v63))
        _ = fi65
        _ = v279
    }
    v280 := GenS25{F0: 0, F1: (false && ((v4 ^ v4) == 18446744073709551615)), F2: -32768}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v280
    _ = v3
    _ = v4
    _ = v47
    _ = v62
    _ = v63
}
