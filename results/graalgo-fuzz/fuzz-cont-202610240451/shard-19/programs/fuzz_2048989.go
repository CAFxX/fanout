package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 int8; F1 uint8; F2 int8 }
type GenS6 struct { F0 uint8; F1 int8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint; F1 int32; F2 uint64 }
type GenS10 struct { F0 int; F1 bool; F2 int64 }
type GenS11 struct { F0 string; F1 int32; F2 int64 }
type GenS12 struct { F0 int; F1 int32 }
type GenS13 struct { F0 string; F1 int8; F2 int8 }
type GenS14 struct { F0 string; F1 int }
type GenS15 struct { F0 uint8; F1 uint8 }
type GenS16 struct { F0 uint; F1 int64; F2 int16 }
type GenS17 struct { F0 uint16 }
func genFn0(p0 int64, p1 uint16) int64 {
    for fi6 := 0; fi6 < 14; fi6++ {
        for fi7 := 0; fi7 < 20; fi7++ {
            switch "a" {
            case "foo bar":
                fmt.Println(fi7)
            default:
                fmt.Println(fi6)
            }
            var v8 uint64 = 2
            _ = fi7
            _ = v8
        }
        p0 = p0
        var v9 int32 = 0
        _ = fi6
        _ = v9
    }
    switch "" {
    case "\n\t":
        for fi10 := 0; fi10 < 9; fi10++ {
            v11 := GenS1{F0: (^(p1))}
            for fi12 := 0; fi12 < 4; fi12++ {
                var v13 float32 = float32(float64(1e10))
                fmt.Println(fi10, fi12)
                _ = fi12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
    case "foo bar":
        p0 = 1
    case "a":
        v14 := []int{128, -2147483648, 127}
        _ = v14
    }
    p0 += p0
    fmt.Println(p1)
    _ = p0
    _ = p1
    return p0
}
func genFn1(p0 int64, p1 uint, p2 uint16) int8 {
    if (!false) {
        v15 := []int{4294967295}
        v15 = append(v15, 785)
        if false {
            if (42 < -1) {
                var v16 uint16 = 255
                switch "a" {
                case "foo bar":
                    _ = len(v15)
                }
                for fi17 := 0; fi17 < 9; fi17++ {
                    v16 = p2
                    if (!true) {
                        v18 := GenS1{F0: (v16 << ((p1) & 63))}
                        if true {
                            for fi19 := 0; fi19 < 1; fi19++ {
                                if true {
                                    if len(v15) >= 2 {
                                        v15 = v15[1:len(v15)]
                                    }
                                } else {
                                    var v20 float32 = (-2.25 * (1.5 / (0.1 + float32(((0.1 / 1e10) / 1.5)))))
                                    p1 = 1
                                    _ = v20
                                }
                                p0 -= genFn0(genFn0(p0, (v16 / (((v16 & (p2 * p2))) | 1))), 1)
                                _ = fi19
                            }
                            v21 := GenS2{F0: 65535, F1: (^(p0))}
                            var v22 uint8 = 2
                            _ = v21
                            _ = v22
                        } else {
                            if (!(((0.1 - 3.14159) != 0.1) || false)) {
                                if (42 > 127) {
                                    for fi23 := 0; fi23 < 15; fi23++ {
                                        fmt.Println(fi17, v16)
                                        v24 := GenS3{F0: 32767, F1: 1000}
                                        if false {
                                            p1 += p1
                                        } else {
                                            var v25 int16 = 32767
                                            v25 += -286
                                            _ = v25
                                        }
                                        _ = fi23
                                        _ = v24
                                    }
                                }
                                _ = len(v15)
                                v26 := GenS3{F0: 567, F1: 0}
                                _ = v26
                            }
                            switch fi17 {
                            case -1:
                                var v27 int = (fi17 + fi17)
                                _ = v27
                            case -513:
                                _ = len(v15)
                            case 32768:
                                var v28 float32 = 0.1
                                _ = v28
                            default:
                                var v29 int32 = 42
                                _ = v29
                            }
                            v30 := GenS1{F0: v16}
                            _ = v30
                        }
                        _ = v18
                    }
                    if ((0.0 + (float32(1e10) * float32(-2.25))) > ((((0.0 * (0.1 - -2.25)) - 3.14159) / (3.14159 * 1.5)) / 1.0)) {
                        switch (-(fi17)) {
                        case 647:
                            fmt.Println(fi17, p0, fi17)
                        default:
                            for fi31 := 0; fi31 < 3; fi31++ {
                                p2 = (v16 % (((p2 % (((p2 >> (((p1 / ((p1) | 1))) & 63))) | 1))) | 1))
                                if len(v15) >= 2 {
                                    v15 = v15[1:len(v15)]
                                }
                                var v32 int64 = (^(p0))
                                _ = fi31
                                _ = v32
                            }
                        }
                        var v33 uint16 = (v16 % ((v16) | 1))
                        _ = v33
                    } else {
                        if true {
                            for fi34 := 0; fi34 < 17; fi34++ {
                                fmt.Println(p0, v16)
                                if ((float64(0.0) * 1e10) != ((1.5 / (-2.25 + (1.5 + -2.25))) - (1e10 + (((1.5 * ((((0.0 - 1.5) - (1.5 / -2.25)) + float64((((1e10 * float32(0.0)) * (0.0 * (((3.14159 * (-2.25 / 1.5)) / -2.25) / 1.5))) * 0.1))) + -2.25)) * 1e10) + ((1e10 + 0.0) * float64(3.14159)))))) {
                                    p0 = (-(p0))
                                    if (!true) {
                                        var v35 float64 = 3.14159
                                        p0 -= (p0 % (((p0 * p0)) | 1))
                                        v15 = append(v15, (fi17 << ((p1) & 63)))
                                        if true {
                                            v16 = p2
                                            var v36 uint16 = v16
                                            _ = v36
                                        }
                                        v37 := GenS2{F0: 0, F1: (^(p0))}
                                        _ = v35
                                        _ = v37
                                    } else {
                                        p1 = p1
                                        var v38 float64 = -2.25
                                        _ = v38
                                    }
                                }
                                v39 := GenS4{F0: (fi17 >> ((255) & 63)), F1: fi34}
                                _ = fi34
                                _ = v39
                            }
                            var v40 int64 = (p0 & genFn0(829, p2))
                            _ = v40
                        }
                        v41 := GenS4{F0: fi17, F1: (-(fi17))}
                        v16 = v16
                        _ = v41
                    }
                    _ = fi17
                }
                _ = v16
            } else {
                var v42 int64 = (p0 % (((p0 / ((-471) | 1))) | 1))
                if false {
                    switch ("\n\t" + "\n\t") {
                    case "a":
                        _ = len(v15)
                    case "hello":
                        var v43 int8 = 1
                        _ = v43
                    }
                    for fi44 := 0; fi44 < 6; fi44++ {
                        if (18446744073709551615 < 400) {
                            var v45 bool = false
                            _ = v45
                        }
                        var v46 float64 = (float64(3.14159) + 1e10)
                        v47 := GenS4{F0: (fi44 / ((9223372036854775807) | 1)), F1: 838}
                        _ = fi44
                        _ = v46
                        _ = v47
                    }
                }
                _ = v42
            }
            switch "" {
            case "a":
                for fi48 := 0; fi48 < 17; fi48++ {
                    var v49 uint8 = 1
                    if (646 < 4294967295) {
                        v49 += v49
                    }
                    var v50 int32 = 1
                    _ = fi48
                    _ = v49
                    _ = v50
                }
            case "\n\t":
                v51 := GenS1{F0: p2}
                _ = v51
            }
            for fi52 := 0; fi52 < 7; fi52++ {
                switch ("hello" + ("a" + (("a" + ("" + "")) + "a"))) {
                case "":
                    if false {
                        for fi53 := 0; fi53 < 17; fi53++ {
                            if true {
                                p2 -= (p2 << ((uint(p2)) & 63))
                                var v54 uint32 = 42
                                _ = v54
                            }
                            var v55 int16 = 168
                            _ = fi53
                            _ = v55
                        }
                        for fi56 := 0; fi56 < 14; fi56++ {
                            if (!(!false)) {
                                _ = len(v15)
                                var v57 uint32 = 42
                                _ = v57
                            }
                            var v58 int8 = 100
                            p0 = p0
                            _ = fi56
                            _ = v58
                        }
                    } else {
                        if false {
                            for fi59 := 0; fi59 < 20; fi59++ {
                                v60 := GenS5{F0: 44, F1: 42, F2: -85}
                                if true {
                                    var v61 float32 = -2.25
                                    v15 = append(v15, 2147483647)
                                    var v62 int8 = 42
                                    _ = v61
                                    _ = v62
                                }
                                if len(v15) >= 2 {
                                    v15 = v15[1:len(v15)]
                                }
                                _ = fi59
                                _ = v60
                            }
                            v63 := GenS3{F0: 416, F1: 1000}
                            fmt.Println(p2, p1, p0)
                            _ = v63
                        } else {
                            p2 += 411
                            v64 := GenS3{F0: int16(p2), F1: 863}
                            if false {
                                p2 = 25
                                p0 -= p0
                                if false {
                                    switch (fi52 ^ (fi52 % (((^(fi52))) | 1))) {
                                    case -576:
                                        p0 -= (p0 << ((18446744073709551615) & 63))
                                    case 2147483647:
                                        switch fi52 {
                                        case -9223372036854775808:
                                            if true {
                                                if len(v15) >= 2 {
                                                    v15 = v15[1:len(v15)]
                                                }
                                            } else {
                                                p2 = p2
                                                var v65 uint = 18446744073709551615
                                                _ = v65
                                            }
                                        case 0:
                                            if false {
                                                if len(v15) >= 2 {
                                                    v15 = v15[1:len(v15)]
                                                }
                                                v66 := GenS5{F0: 118, F1: 255, F2: 42}
                                                _ = v66
                                            }
                                        case 4294967295:
                                            p1 = (^(p1))
                                        default:
                                            v67 := GenS6{F0: 42, F1: int8(fi52)}
                                            _ = v67
                                        }
                                    case 801:
                                        v68 := GenS6{F0: 1, F1: -24}
                                        _ = v68
                                    }
                                    p1 += p1
                                    fmt.Println(p0)
                                }
                            } else {
                                for fi69 := 0; fi69 < 18; fi69++ {
                                    var v70 string = ("" + ("a" + ("\n\t" + ("" + ("a" + (("foo bar" + "a") + "\n\t"))))))
                                    p2 += p2
                                    _ = fi69
                                    _ = v70
                                }
                                switch "\n\t" {
                                case "a":
                                    switch ("foo bar" + ("a" + "\n\t")) {
                                    case "a":
                                        if false {
                                            if len(v15) > 0 {
                                                v15[0] = (^(fi52))
                                            }
                                            switch fi52 {
                                            case -95:
                                                p1 = p1
                                            case -1:
                                                p1 += (p1 << ((p1) & 63))
                                            }
                                        }
                                    default:
                                        if (false || ((!false) || false)) {
                                            for fi71 := 0; fi71 < 8; fi71++ {
                                                var v72 uint32 = 2
                                                _ = fi71
                                                _ = v72
                                            }
                                        }
                                    }
                                case "\n\t":
                                    p2 = (p2 - (p2 & 1))
                                default:
                                    switch "\n\t" {
                                    case "":
                                        if len(v15) > 0 {
                                            v15[0] = 4294967295
                                        }
                                    case "hello":
                                        if len(v15) > 0 {
                                            v15[0] = (fi52 + fi52)
                                        }
                                    case "a":
                                        fmt.Println(fi52, p0)
                                    default:
                                        var v73 float64 = -2.25
                                        _ = v73
                                    }
                                }
                            }
                            _ = v64
                        }
                        fmt.Println(fi52, p1)
                    }
                case "hello":
                    if (p0 != genFn0(1, p2)) {
                        p1 = (^(p1))
                        for fi74 := 0; fi74 < 2; fi74++ {
                            v15 = append(v15, 1)
                            fmt.Println(p2, fi74, fi74)
                            var v75 uint64 = 743
                            _ = fi74
                            _ = v75
                        }
                    }
                default:
                    _ = len(v15)
                }
                p0 += p0
                v76 := GenS6{F0: 0, F1: 12}
                _ = fi52
                _ = v76
            }
        }
        _ = v15
    } else {
        v77 := []int{}
        if false {
            p0 += (p0 % ((genFn0(p0, p2)) | 1))
            for fi78 := 0; fi78 < 14; fi78++ {
                p2 -= (p2 * (p2 % (((^(p2))) | 1)))
                if (832 == 0) {
                    v79 := GenS7{F0: p0}
                    var v80 int = (fi78 & (fi78 & fi78))
                    p1 = uint(p2)
                    _ = v79
                    _ = v80
                } else {
                    var v81 string = "foo bar"
                    for fi82 := 0; fi82 < 6; fi82++ {
                        if true {
                            if true {
                                var v83 string = ("hello" + "")
                                switch (v81 + v81) {
                                case "foo bar":
                                    var v84 uint8 = 128
                                    _ = v84
                                case "hello":
                                    p0 += p0
                                case "\n\t":
                                    p0 -= p0
                                }
                                _ = len(v77)
                                _ = v83
                            } else {
                                var v85 uint = p1
                                var v86 uint64 = 877
                                var v87 uint64 = v86
                                _ = v85
                                _ = v86
                                _ = v87
                            }
                            if (((v81 + v81) > v81) || (fi82 == (fi82 & 213))) {
                                if len(v77) > 0 {
                                    v77[0] = fi78
                                }
                                _ = len(v77)
                            } else {
                                v88 := GenS5{F0: int8(p2), F1: 1, F2: 0}
                                v81 = ("" + (v81 + v81))
                                switch -49 {
                                case 4294967295:
                                    _ = len(v77)
                                default:
                                    if true {
                                        v89 := GenS4{F0: (fi82 / (((fi78 ^ fi78)) | 1)), F1: 9223372036854775807}
                                        var v90 uint8 = 255
                                        var v91 float32 = (-2.25 / float32(-2.25))
                                        var v92 int = (fi82 % ((fi82) | 1))
                                        _ = v89
                                        _ = v90
                                        _ = v91
                                        _ = v92
                                    }
                                }
                                _ = v88
                            }
                            switch ((("a" + v81) + v81) + "hello") {
                            case "":
                                v93 := GenS1{F0: p2}
                                v94 := GenS7{F0: -1}
                                _ = v93
                                _ = v94
                            case "a":
                                var v95 uint = (p1 * p1)
                                _ = v95
                            case "hello":
                                v81 = "hello"
                            default:
                                var v96 uint = (p1 % ((p1) | 1))
                                _ = v96
                            }
                            p1 = 599
                            if (false && (1.5 >= 1e10)) {
                                var v97 bool = ((-1 >= 42) && (!(!true)))
                                switch 0 {
                                case -466:
                                    if len(v77) > 0 {
                                        v77[0] = fi82
                                    }
                                    p0 = (p0 | p0)
                                case 521:
                                    fmt.Println(fi78, fi82, v97)
                                default:
                                    var v98 int = -131
                                    _ = v98
                                }
                                if (v97 && false) {
                                    v99 := GenS1{F0: p2}
                                    if (299 == (p2 / ((p2) | 1))) {
                                        if ((2147483647 > 1) && (false && (-2 == -937))) {
                                            p1 -= p1
                                            switch (((v81 + "hello") + ("" + v81)) + v81) {
                                            case "foo bar":
                                                fmt.Println(p0)
                                            case "\n\t":
                                                p0 = (^(p0))
                                            }
                                            v81 = v81
                                            var v100 int32 = -757
                                            var v101 uint16 = (p2 ^ 2)
                                            switch (fi82 + -1) {
                                            case 32767:
                                                _ = len(v77)
                                            default:
                                                switch "hello" {
                                                case "a":
                                                    var v102 int8 = -128
                                                    _ = v102
                                                case "foo bar":
                                                    v103 := GenS8{F0: 1}
                                                    v77 = append(v77, (fi82 - (fi78 + fi78)))
                                                    _ = v103
                                                default:
                                                    if len(v77) >= 2 {
                                                        v77 = v77[1:len(v77)]
                                                    }
                                                }
                                            }
                                            var v104 uint = p1
                                            switch "a" {
                                            case "":
                                                var v105 uint8 = 190
                                                _ = v105
                                            case "a":
                                                if v97 {
                                                    v106 := GenS2{F0: 246, F1: (-(p0))}
                                                    fmt.Println(fi82, p0)
                                                    fmt.Println(fi78)
                                                    _ = v106
                                                }
                                            case "foo bar":
                                                switch v81 {
                                                case "a":
                                                    var v107 int8 = -1
                                                    _ = v107
                                                case "foo bar":
                                                    fmt.Println(fi78)
                                                default:
                                                    fmt.Println(v97, fi82)
                                                }
                                                if len(v77) > 0 {
                                                    v77[0] = (fi78 & 32767)
                                                }
                                                var v108 float32 = 0.0
                                                _ = v108
                                            }
                                            _ = v100
                                            _ = v101
                                            _ = v104
                                        }
                                        if len(v77) > 0 {
                                            v77[0] = (fi82 | (-(fi78)))
                                        }
                                        if len(v77) >= 2 {
                                            v77 = v77[1:len(v77)]
                                        }
                                    }
                                    _ = v99
                                }
                                if ((-28 <= 0) || v97) {
                                    v109 := GenS9{F0: (p1 % (((^(p1))) | 1)), F1: -2147483648, F2: 4294967295}
                                    _ = v109
                                }
                                _ = v97
                            } else {
                                fmt.Println(p1, p1, fi82)
                            }
                        }
                        var v110 uint64 = 0
                        _ = fi82
                        _ = v110
                    }
                    _ = len(v77)
                    _ = v81
                }
                _ = fi78
            }
        }
        v77 = append(v77, 32768)
        _ = v77
    }
    v111 := GenS10{F0: 128, F1: true, F2: (p0 * p0)}
    p2 = 571
    p2 -= p2
    _ = p0
    _ = p1
    _ = p2
    _ = v111
    return 6
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []uint32{634, 42, uint32(v1), 426}
    var v3 bool = false
    var v4 uint16 = 255
    var v5 string = (((("" + "hello") + ("" + ("\n\t" + "foo bar"))) + (("" + (("hello" + ("" + "hello")) + "\n\t")) + "")) + (("\n\t" + "") + "hello"))
    var v112 bool = v3
    v5 = v5
    v4 += (v4 % ((32767) | 1))
    for fi113 := 0; fi113 < 12; fi113++ {
        for fi114 := 0; fi114 < 15; fi114++ {
            fmt.Println(v5)
            v5 = "\n\t"
            fmt.Println(v4, v4)
            _ = fi114
        }
        if true {
            for fi115 := 0; fi115 < 15; fi115++ {
                var v116 int16 = 0
                _ = fi115
                _ = v116
            }
            v117 := GenS6{F0: uint8(v4), F1: genFn1(42, 273, v4)}
            fmt.Println(v112, v1)
            _ = v117
        }
        v118 := GenS2{F0: 42, F1: 408}
        _ = fi113
        _ = v118
    }
    if (true && (1 >= genFn0(1, (^(v4))))) {
        switch ("a" + "\n\t") {
        case "foo bar":
            v119 := GenS7{F0: 583}
            _ = v119
        case "":
            switch 511 {
            case 2147483647:
                for fi120 := 0; fi120 < 4; fi120++ {
                    var v121 int16 = -32768
                    _ = fi120
                    _ = v121
                }
            }
        case "a":
            fmt.Println(v3, v112)
        default:
            v1 = v1
        }
    }
    if (118 > (v4 / (((v4 / ((v4) | 1))) | 1))) {
        switch 86 {
        case -855:
            _ = len(v2)
        case -619:
            if true {
                v122 := GenS11{F0: ((v5 + v5) + v5), F1: -2147483648, F2: -1}
                fmt.Println(v4)
                _ = v122
            }
        case -9223372036854775808:
            var v123 int16 = 600
            _ = v123
        default:
            for fi124 := 0; fi124 < 2; fi124++ {
                if (42 == 2147483647) {
                    switch ("" + "") {
                    case "":
                        switch "a" {
                        case "":
                            if ((v1 + 468) != v1) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v125 float32 = (3.14159 / 1.0)
                                _ = v125
                            }
                        default:
                            var v126 float32 = 3.14159
                            _ = v126
                        }
                    default:
                        var v127 uint = 42
                        _ = v127
                    }
                    var v128 int8 = -1
                    switch ("a" + v5) {
                    case "foo bar":
                        if (2147483648 <= 4294967295) {
                            v129 := GenS12{F0: fi124, F1: int32(v4)}
                            v130 := GenS13{F0: v5, F1: -1, F2: -1}
                            var v131 int16 = 556
                            _ = v129
                            _ = v130
                            _ = v131
                        }
                    default:
                        if false {
                            switch 128 {
                            case 42:
                                var v132 uint16 = v4
                                _ = v132
                            case 9223372036854775807:
                                v5 = (v5 + v5)
                            case -2147483648:
                                v1 -= 719
                            default:
                                v5 = v5
                            }
                        }
                    }
                    _ = v128
                }
                v5 = "\n\t"
                _ = fi124
            }
        }
    } else {
        if v3 {
            var v133 uint32 = 2147483648
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v133
        }
    }
    if (!v112) {
        var v134 int64 = -862
        _ = v134
    } else {
        switch "a" {
        case "\n\t":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "":
            if false {
                v4 = 682
                for fi135 := 0; fi135 < 20; fi135++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi135
                }
            }
        }
        for fi136 := 0; fi136 < 17; fi136++ {
            var v137 int32 = -1
            for fi138 := 0; fi138 < 11; fi138++ {
                v1 = (v1 * (v1 - (v1 - 236)))
                if (v137 < v137) {
                    var v139 string = v5
                    _ = v139
                } else {
                    fmt.Println(v3, fi136, fi136)
                    if ((fi136 ^ fi138) != 341) {
                        switch "hello" {
                        case "foo bar":
                            var v140 int8 = 1
                            _ = v140
                        case "a":
                            v141 := GenS10{F0: fi136, F1: (true || false), F2: 1}
                            _ = v141
                        }
                        if (false || v112) {
                            _ = len(v2)
                            var v142 uint8 = 2
                            v3 = v3
                            var v143 uint64 = v1
                            v142 = v142
                            _ = v142
                            _ = v143
                        }
                    } else {
                        v144 := GenS8{F0: -70}
                        var v145 int = fi136
                        v3 = v3
                        v146 := GenS14{F0: ((v5 + v5) + "foo bar"), F1: 4294967295}
                        _ = v144
                        _ = v145
                        _ = v146
                    }
                    v5 = v5
                    var v147 int = (fi138 | (fi138 & fi138))
                    v2 = append(v2, 4294967295)
                    _ = v147
                }
                v5 = (v5 + ("\n\t" + v5))
                fmt.Println(fi138, v5, v4)
                v148 := GenS15{F0: 127, F1: 174}
                v149 := GenS10{F0: (fi136 - fi138), F1: (v112 && v112), F2: 566}
                var v150 string = "foo bar"
                _ = fi138
                _ = v148
                _ = v149
                _ = v150
            }
            v137 = 42
            _ = fi136
            _ = v137
        }
        v151 := GenS16{F0: 888, F1: 9223372036854775807, F2: -32768}
        _ = v151
    }
    var v152 int8 = -128
    v5 = v5
    v112 = true
    fmt.Println(v1, v3, v152)
    v153 := GenS17{F0: v4}
    var v154 uint16 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v112:", v112)
    fmt.Println("v152:", v152)
    fmt.Println("v154:", v154)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v112
    _ = v152
    _ = v153
    _ = v154
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
