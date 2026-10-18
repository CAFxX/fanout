package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int; F1 int8; F2 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16; F1 string }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int8; F1 uint32 }
type GenS8 struct { F0 int; F1 uint16; F2 bool }
type GenS9 struct { F0 string; F1 uint32; F2 string }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 int32; F1 uint64 }
type GenS15 struct { F0 uint16; F1 string }
type GenS16 struct { F0 uint32 }
type GenS17 struct { F0 uint64; F1 int }
type GenS18 struct { F0 int8; F1 int }
func genFn0(p0 uint32, p1 int) uint64 {
    for fi7 := 0; fi7 < 6; fi7++ {
        if false {
            v8 := []int{p1, (p1 << ((2) & 63)), (fi7 / ((fi7) | 1)), (p1 - p1)}
            for fi9 := 0; fi9 < 17; fi9++ {
                v8 = append(v8, (^(fi9)))
                fmt.Println(fi9)
                v10 := GenS1{F0: 853}
                _ = fi9
                _ = v10
            }
            fmt.Println(fi7, p1, p1)
            _ = v8
        } else {
            var v11 float64 = 3.14159
            _ = v11
        }
        for fi12 := 0; fi12 < 14; fi12++ {
            var v13 uint16 = 0
            if (false && ((((^(v13)) > (v13 % (((v13 & v13)) | 1))) || (122 < 227)) && ((100 == int8(p1)) || (-2.25 > float32((1e10 / -2.25)))))) {
                var v14 uint32 = p0
                _ = v14
            }
            _ = fi12
            _ = v13
        }
        _ = fi7
    }
    switch "a" {
    case "\n\t":
        if (1 <= 42) {
            v15 := []int{(^(p1)), (p1 & p1), p1, (p1 & p1)}
            var v16 int8 = 42
            _ = v15
            _ = v16
        }
    default:
        switch p1 {
        case -2147483648:
            var v17 float32 = -2.25
            _ = v17
        case 32767:
            var v18 float32 = -2.25
            _ = v18
        case -773:
            var v19 uint16 = 431
            _ = v19
        }
    }
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 uint, p1 uint16, p2 uint16) uint8 {
    v20 := []int{32767, -869, -801}
    p2 -= p2
    p2 += p1
    if true {
        for fi21 := 0; fi21 < 11; fi21++ {
            p2 = (p2 + (p1 + 565))
            _ = fi21
        }
        for fi22 := 0; fi22 < 4; fi22++ {
            for fi23 := 0; fi23 < 9; fi23++ {
                v24 := GenS1{F0: 65535}
                if (-2147483648 >= 821) {
                    p2 = p1
                    var v25 bool = true
                    _ = v25
                } else {
                    var v26 float32 = (float32(3.14159) / 1e10)
                    var v27 int32 = 2147483647
                    switch ("foo bar" + ("hello" + "\n\t")) {
                    case "\n\t":
                        switch fi23 {
                        case 271:
                            if false {
                                if true {
                                    _ = len(v20)
                                }
                                v20 = append(v20, 2147483647)
                                v27 -= v27
                                if len(v20) > 0 {
                                    v20[0] = (fi22 - (fi23 * 958))
                                }
                            }
                        case 128:
                            var v28 uint8 = 185
                            _ = v28
                        case -598:
                            if len(v20) > 0 {
                                v20[0] = fi22
                            }
                        default:
                            var v29 uint64 = 1
                            _ = v29
                        }
                    }
                    _ = v26
                    _ = v27
                }
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
        switch 42 {
        case 2147483647:
            switch "a" {
            case "":
                if true {
                    p2 += p2
                } else {
                    if false {
                        v30 := GenS2{F0: false, F1: 898, F2: p1}
                        for fi31 := 0; fi31 < 5; fi31++ {
                            _ = len(v20)
                            var v32 string = "a"
                            fmt.Println(p2, v32, p2)
                            _ = fi31
                            _ = v32
                        }
                        _ = v30
                    }
                    p1 -= 1
                    if (128 < 174) {
                        switch ("" + "\n\t") {
                        case "\n\t":
                            for fi33 := 0; fi33 < 2; fi33++ {
                                switch -9223372036854775808 {
                                case 128:
                                    switch "" {
                                    case "\n\t":
                                        v34 := GenS1{F0: (^(p0))}
                                        _ = v34
                                    case "":
                                        if true {
                                            var v35 string = ("" + "hello")
                                            if len(v20) > 0 {
                                                v20[0] = fi33
                                            }
                                            for fi36 := 0; fi36 < 4; fi36++ {
                                                fmt.Println(fi33, p1)
                                                p0 += (^(p0))
                                                _ = fi36
                                            }
                                            _ = v35
                                        }
                                    default:
                                        fmt.Println(p0, p1, fi33)
                                    }
                                }
                                p2 -= (p1 << ((65535) & 63))
                                p0 += 2
                                _ = fi33
                            }
                        case "a":
                            p1 = p1
                        case "foo bar":
                            v20 = append(v20, -603)
                        default:
                            p0 += 255
                        }
                        var v37 int64 = 42
                        var v38 float64 = 3.14159
                        _ = v37
                        _ = v38
                    } else {
                        _ = len(v20)
                        for fi39 := 0; fi39 < 11; fi39++ {
                            for fi40 := 0; fi40 < 6; fi40++ {
                                var v41 uint16 = (^(p2))
                                v41 += (p2 << (((p0 | (^(p0)))) & 63))
                                _ = fi40
                                _ = v41
                            }
                            fmt.Println(p2, p2, p0)
                            _ = fi39
                        }
                        v20 = append(v20, -9223372036854775808)
                    }
                }
            case "hello":
                if (p2 == (p1 >> ((p0) & 63))) {
                    p1 += p1
                } else {
                    var v42 uint64 = 4294967295
                    var v43 uint = (^(p0))
                    var v44 int8 = -128
                    _ = v42
                    _ = v43
                    _ = v44
                }
            case "\n\t":
                switch 1 {
                case 42:
                    for fi45 := 0; fi45 < 14; fi45++ {
                        for fi46 := 0; fi46 < 2; fi46++ {
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                            _ = fi46
                        }
                        _ = fi45
                    }
                }
            default:
                if (uint32(p0) > 2) {
                    switch "a" {
                    case "a":
                        _ = len(v20)
                    default:
                        v47 := GenS1{F0: 2}
                        _ = v47
                    }
                    if (!(true || true)) {
                        switch "a" {
                        case "":
                            if (2147483648 <= 65535) {
                                v48 := GenS1{F0: p0}
                                if (p2 != p1) {
                                    fmt.Println(p0, p2)
                                    v49 := GenS2{F0: ((65535 > 131) && (66 != -2147483648)), F1: p2, F2: p2}
                                    if true {
                                        var v50 int16 = 42
                                        if (true || (-9223372036854775808 < 1)) {
                                            for fi51 := 0; fi51 < 3; fi51++ {
                                                p2 -= p1
                                                for fi52 := 0; fi52 < 19; fi52++ {
                                                    var v53 uint64 = 2
                                                    fmt.Println(p2)
                                                    _ = fi52
                                                    _ = v53
                                                }
                                                _ = fi51
                                            }
                                            p0 += 18446744073709551615
                                            if ((float32(((1e10 - 1.5) - -2.25)) + (3.14159 - -2.25)) >= float32((float64(0.0) * ((((0.1 * -2.25) - 0.0) * -2.25) * -2.25)))) {
                                                if ((!true) && ((-128 >= -1) || false)) {
                                                    switch "a" {
                                                    case "hello":
                                                        if false {
                                                            var v54 float64 = 1.5
                                                            _ = v54
                                                        }
                                                    }
                                                    var v55 uint16 = p1
                                                    for fi56 := 0; fi56 < 10; fi56++ {
                                                        fmt.Println(p0)
                                                        _ = fi56
                                                    }
                                                    _ = v55
                                                } else {
                                                    for fi57 := 0; fi57 < 16; fi57++ {
                                                        p1 = (^(p2))
                                                        _ = fi57
                                                    }
                                                    v50 = v50
                                                }
                                            } else {
                                                for fi58 := 0; fi58 < 18; fi58++ {
                                                    var v59 float32 = float32((1e10 - ((-2.25 + 3.14159) * (-2.25 - (0.0 / (1e10 / 1.0))))))
                                                    var v60 string = "foo bar"
                                                    _ = fi58
                                                    _ = v59
                                                    _ = v60
                                                }
                                                var v61 int8 = -113
                                                _ = v61
                                            }
                                        }
                                        _ = v50
                                    } else {
                                        for fi62 := 0; fi62 < 9; fi62++ {
                                            if (!((1 == 100) || false)) {
                                                var v63 uint64 = 2
                                                fmt.Println(v63, p0, p0)
                                                p2 = p2
                                                _ = v63
                                            } else {
                                                v64 := GenS3{F0: 86, F1: 104, F2: "a"}
                                                p1 -= 2
                                                v65 := GenS2{F0: ((637 > 2) && (1 <= 127)), F1: (p2 + 2), F2: p2}
                                                _ = v64
                                                _ = v65
                                            }
                                            var v66 int16 = -741
                                            var v67 uint16 = p2
                                            _ = fi62
                                            _ = v66
                                            _ = v67
                                        }
                                        p0 -= p0
                                    }
                                    _ = v49
                                } else {
                                    if len(v20) > 0 {
                                        v20[0] = -2147483648
                                    }
                                    var v68 uint16 = (p1 + p2)
                                    v69 := GenS4{F0: 496}
                                    _ = v68
                                    _ = v69
                                }
                                var v70 uint16 = p1
                                _ = v48
                                _ = v70
                            }
                        default:
                            for fi71 := 0; fi71 < 1; fi71++ {
                                var v72 int32 = 0
                                if (42 <= (p1 << (((p0 & p0)) & 63))) {
                                    p2 -= p2
                                    p0 -= (p0 & p0)
                                    for fi73 := 0; fi73 < 6; fi73++ {
                                        var v74 string = (("a" + "hello") + "hello")
                                        _ = fi73
                                        _ = v74
                                    }
                                } else {
                                    fmt.Println(fi71, p2, p0)
                                }
                                v20 = append(v20, 1)
                                _ = fi71
                                _ = v72
                            }
                        }
                    } else {
                        var v75 bool = true
                        _ = v75
                    }
                    switch -165 {
                    case 0:
                        p1 -= p1
                    case -815:
                        for fi76 := 0; fi76 < 12; fi76++ {
                            var v77 string = "foo bar"
                            for fi78 := 0; fi78 < 1; fi78++ {
                                switch fi78 {
                                case 2147483647:
                                    v79 := GenS5{F0: 572, F1: (v77 + (v77 + "\n\t"))}
                                    _ = v79
                                case 97:
                                    var v80 int8 = 42
                                    var v81 float64 = ((0.1 * -2.25) - 0.1)
                                    p2 -= p1
                                    _ = v80
                                    _ = v81
                                default:
                                    fmt.Println(v77)
                                }
                                v77 = v77
                                switch (v77 + v77) {
                                case "foo bar":
                                    var v82 int = fi78
                                    _ = v82
                                case "":
                                    switch fi78 {
                                    case 127:
                                        p0 = (^(p0))
                                    case -1:
                                        p0 += (^(p0))
                                    }
                                default:
                                    var v83 int8 = -1
                                    _ = v83
                                }
                                if (p0 >= (p0 % (((p0 & p0)) | 1))) {
                                    p2 = p1
                                    if (false && true) {
                                        if (v77 <= v77) {
                                            var v84 int8 = -38
                                            switch "a" {
                                            case "hello":
                                                var v85 float32 = 3.14159
                                                _ = v85
                                            case "":
                                                if false {
                                                    v86 := GenS4{F0: 9223372036854775807}
                                                    p2 += 255
                                                    fmt.Println(p0, fi76)
                                                    _ = v86
                                                } else {
                                                    v87 := GenS3{F0: -1, F1: (v84 | (v84 + v84)), F2: (("\n\t" + v77) + "")}
                                                    if (((v77 + ("" + (v77 + v77))) + v77) < (v77 + "a")) {
                                                        v84 = v84
                                                    } else {
                                                        var v88 uint = (p0 >> ((uint(p2)) & 63))
                                                        switch -9223372036854775808 {
                                                        case 326:
                                                            if false {
                                                                if false {
                                                                    v84 += (-(v84))
                                                                    v84 += (v84 << ((p0) & 63))
                                                                } else {
                                                                    switch v77 {
                                                                    case "foo bar":
                                                                        v20 = append(v20, 32768)
                                                                    case "":
                                                                        var v89 int8 = v84
                                                                        fmt.Println(fi78)
                                                                        _ = v89
                                                                    case "a":
                                                                        v88 += 65535
                                                                    }
                                                                    fmt.Println(v77)
                                                                    var v90 int64 = 86
                                                                    _ = v90
                                                                }
                                                                v88 -= (p0 >> ((1) & 63))
                                                            } else {
                                                                v91 := GenS5{F0: p2, F1: (v77 + "")}
                                                                p1 = (p2 | (p1 ^ (p2 << ((p0) & 63))))
                                                                _ = v91
                                                            }
                                                            fmt.Println(p1)
                                                        default:
                                                            p1 += p2
                                                        }
                                                        _ = v88
                                                    }
                                                    var v92 float64 = float64(1.5)
                                                    _ = v87
                                                    _ = v92
                                                }
                                            case "foo bar":
                                                if false {
                                                    switch (("\n\t" + "") + (v77 + v77)) {
                                                    case "hello":
                                                        if false {
                                                            p2 -= (p2 + 0)
                                                            var v93 uint32 = 0
                                                            _ = v93
                                                        } else {
                                                            _ = len(v20)
                                                            var v94 bool = (p2 >= 2)
                                                            if (!((fi76 >> ((p0) & 63)) >= fi78)) {
                                                                var v95 bool = (678 <= 0)
                                                                v94 = v95
                                                                _ = v95
                                                            } else {
                                                                v84 -= (v84 - v84)
                                                            }
                                                            _ = v94
                                                        }
                                                    default:
                                                        if (42 == 42) {
                                                            p1 -= p2
                                                            p0 -= p0
                                                            if (-749 == 42) {
                                                                v77 = v77
                                                                p1 = (p1 | p1)
                                                                p2 = p2
                                                            }
                                                        }
                                                        p0 = (p0 + 42)
                                                        if false {
                                                            var v96 uint8 = 0
                                                            v20 = append(v20, -9223372036854775808)
                                                            _ = v96
                                                        } else {
                                                            _ = len(v20)
                                                            p1 = (p2 << ((p0) & 63))
                                                            v97 := GenS1{F0: 0}
                                                            _ = v97
                                                        }
                                                    }
                                                    var v98 int64 = 0
                                                    var v99 int32 = -1
                                                    v77 = (v77 + v77)
                                                    if false {
                                                        switch v77 {
                                                        case "hello":
                                                            switch ((v77 + (("foo bar" + v77) + v77)) + "") {
                                                            case "":
                                                                v20 = append(v20, -705)
                                                            case "\n\t":
                                                                fmt.Println(p0, fi78)
                                                            default:
                                                                switch ("a" + v77) {
                                                                case "a":
                                                                    if len(v20) >= 2 {
                                                                        v20 = v20[1:len(v20)]
                                                                    }
                                                                case "hello":
                                                                    v100 := GenS6{F0: -1}
                                                                    _ = v100
                                                                default:
                                                                    var v101 uint = p0
                                                                    switch fi78 {
                                                                    case -710:
                                                                        var v102 int = fi76
                                                                        _ = v102
                                                                    }
                                                                    v103 := GenS2{F0: true, F1: p2, F2: p1}
                                                                    v99 = (v99 << ((255) & 63))
                                                                    v104 := GenS7{F0: v84, F1: 2147483648}
                                                                    v77 = (v77 + "")
                                                                    v84 += v84
                                                                    _ = v101
                                                                    _ = v103
                                                                    _ = v104
                                                                }
                                                            }
                                                        case "a":
                                                            fmt.Println(p2, v98, p1)
                                                        case "foo bar":
                                                            fmt.Println(fi78, v98, p1)
                                                        default:
                                                            var v105 int32 = v99
                                                            _ = v105
                                                        }
                                                        fmt.Println(v98)
                                                    }
                                                    fmt.Println(v98, v84)
                                                    if (false && (1000 >= 1000)) {
                                                        p2 -= p2
                                                        v77 = (v77 + v77)
                                                        fmt.Println(v84)
                                                    } else {
                                                        var v106 int = (fi76 >> (((p0 ^ p0)) & 63))
                                                        v77 = (v77 + "hello")
                                                        _ = v106
                                                    }
                                                    _ = v98
                                                    _ = v99
                                                }
                                            }
                                            switch v77 {
                                            case "hello":
                                                fmt.Println(v84)
                                            case "\n\t":
                                                p0 = (^(p0))
                                            }
                                            var v107 bool = false
                                            _ = v107
                                            _ = v84
                                        }
                                    }
                                }
                                p1 = (p1 * (p1 & 255))
                                _ = fi78
                            }
                            _ = fi76
                            _ = v77
                        }
                    case -781:
                        fmt.Println(p0)
                    }
                } else {
                    p1 += p2
                    var v108 uint8 = 0
                    _ = v108
                }
            }
        case -13:
            p0 = 80
        }
    }
    if len(v20) > 0 {
        v20[0] = -270
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return 167
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[string]string)
    var v3 uint16 = 2
    v4 := []string{}
    v5 := make(map[string]uint32)
    var v6 uint32 = 1
    v1 = (v1 - float64((-2.25 + (1e10 * float32(v1)))))
    fmt.Println(v3, v1)
    v109 := GenS8{F0: -43, F1: (v3 | (v3 << ((165) & 63))), F2: true}
    fmt.Println(v3, v6)
    v6 = 42
    var v110 uint = 411
    switch 128 {
    case -915:
        v110 = (v110 | 42)
    default:
        v111 := GenS3{F0: 1, F1: -111, F2: "a"}
        _ = v111
    }
    v6 -= (v6 ^ (v6 << ((v110) & 63)))
    var v112 int16 = 0
    v112 -= (v112 % (((v112 | v112)) | 1))
    v1 = 0.0
    v6 -= (v6 | (v6 ^ (v6 | v6)))
    for fi113 := 0; fi113 < 14; fi113++ {
        if false {
            v110 = 255
            for fi114 := 0; fi114 < 9; fi114++ {
                if (("foo bar" + (("foo bar" + ("\n\t" + "\n\t")) + ("hello" + "foo bar"))) >= "\n\t") {
                    v4 = append(v4, "foo bar")
                    switch ("a" + "") {
                    case "a":
                        if true {
                            if true {
                                if len(v4) > 0 {
                                    v4[0] = ""
                                }
                            }
                        } else {
                            var v115 string = "foo bar"
                            v116 := GenS9{F0: "foo bar", F1: (v6 % ((v6) | 1)), F2: v115}
                            _ = v115
                            _ = v116
                        }
                    case "hello":
                        if (!false) {
                            switch fi113 {
                            case 42:
                                v5["a"] = v6
                            default:
                                if false {
                                    switch fi114 {
                                    case 32767:
                                        if (-557 > 42) {
                                            v110 = v110
                                            fmt.Println(v112, v110)
                                        }
                                    default:
                                        fmt.Println(v3)
                                    }
                                }
                            }
                            fmt.Println(v110, v3, v110)
                            v1 -= float64((-2.25 - 0.1))
                        }
                    default:
                        v117 := GenS5{F0: (v3 + 385), F1: "foo bar"}
                        _ = v117
                    }
                } else {
                    if (!(1 != -518)) {
                        switch (-(fi114)) {
                        case 2147483647:
                            v3 += 32767
                        }
                    } else {
                        if (!false) {
                            v118 := v5["foo bar"]
                            _ = v118
                        }
                    }
                    v112 = v112
                }
                if (7 > 2147483647) {
                    fmt.Println(v3)
                } else {
                    v119 := GenS10{F0: 546}
                    if true {
                        var v120 float64 = v1
                        v120 = ((float64(0.0) + v1) + float64(-2.25))
                        _ = v120
                    }
                    _ = v119
                }
                if (false || false) {
                    switch fi114 {
                    case 127:
                        switch (^(fi113)) {
                        case 0:
                            v121 := GenS2{F0: (128 <= (fi113 >> ((473) & 63))), F1: v3, F2: 32767}
                            _ = v121
                        case 127:
                            v110 -= v110
                            v122 := GenS5{F0: v3, F1: ("hello" + "foo bar")}
                            _ = v122
                        default:
                            v1 = v1
                        }
                    case 128:
                        v110 = 654
                    case 0:
                        var v123 uint = v110
                        _ = v123
                    default:
                        if ((v6 ^ (v6 >> ((4294967295) & 63))) > 42) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            switch (("hello" + (("" + "") + "hello")) + ("" + "")) {
                            case "hello":
                                var v124 uint16 = v3
                                _ = v124
                            case "\n\t":
                                var v125 uint32 = v6
                                _ = v125
                            case "":
                                if false {
                                    if (!(!("foo bar" == ""))) {
                                        fmt.Println(v3)
                                    }
                                }
                            }
                            switch "hello" {
                            case "a":
                                switch "" {
                                case "\n\t":
                                    v1 = 3.14159
                                case "hello":
                                    var v126 int8 = 100
                                    _ = v126
                                default:
                                    switch "hello" {
                                    case "":
                                        v127 := GenS4{F0: -9223372036854775808}
                                        _ = v127
                                    }
                                }
                            }
                        }
                    }
                    v112 -= (v112 % ((v112) | 1))
                    v128 := GenS11{F0: true}
                    _ = v128
                }
                v129 := v2["foo bar"]
                if ((v1 <= (-2.25 + v1)) || ((-(fi113)) <= fi113)) {
                    if false {
                        v2[v129] = v129
                        if ((float64(float32((3.14159 + (1.5 - 3.14159)))) / (((v1 * (v1 * (v1 * (1e10 + 0.1)))) - v1) - v1)) >= -2.25) {
                            switch fi113 {
                            case 32768:
                                delete(v5, (v129 + v129))
                            case -540:
                                var v130 int8 = 43
                                _ = v130
                            }
                            v6 -= (v6 | 24)
                        } else {
                            if false {
                                if false {
                                    if true {
                                        switch (fi113 >> ((v110) & 63)) {
                                        case 2147483647:
                                            if false {
                                                _ = len(v5)
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                _ = len(v5)
                                            }
                                        case -1:
                                            var v131 uint32 = v6
                                            _ = v131
                                        case 32767:
                                            v132 := GenS9{F0: "foo bar", F1: v6, F2: "foo bar"}
                                            _ = v132
                                        default:
                                            if len(v4) > 0 {
                                                v4[0] = v129
                                            }
                                        }
                                        v3 += v3
                                        if (true || ((fi113 ^ (fi113 | 136)) < 128)) {
                                            v110 = (v110 * 969)
                                            switch "\n\t" {
                                            case "hello":
                                                v133 := GenS12{F0: (^(v3))}
                                                _ = v133
                                            case "a":
                                                v134 := GenS7{F0: -128, F1: v6}
                                                _ = v134
                                            }
                                            delete(v5, v129)
                                        } else {
                                            if (3.14159 < (float32(((v1 + float64(-2.25)) + v1)) - 0.1)) {
                                                fmt.Println(v1, fi113, v129)
                                                v135 := v2[v129]
                                                v136 := GenS13{F0: (v110 ^ v110)}
                                                _ = v135
                                                _ = v136
                                            }
                                            fmt.Println(v112, v1)
                                            var v137 int32 = -1
                                            _ = v137
                                        }
                                    }
                                    if false {
                                        var v138 int16 = v112
                                        var v139 float64 = v1
                                        fmt.Println(v138, v110, fi114)
                                        _ = v138
                                        _ = v139
                                    } else {
                                        v6 -= v6
                                    }
                                    var v140 uint64 = 1
                                    _ = v140
                                } else {
                                    v141 := GenS5{F0: (v3 & v3), F1: v129}
                                    _ = v141
                                }
                                fmt.Println(fi114, v1, v112)
                                v1 -= (v1 - 0.0)
                                v129 = ""
                                if true {
                                    v112 += (v112 >> ((v110) & 63))
                                    if true {
                                        v129 = ("" + v129)
                                        switch ("" + ((v129 + v129) + "\n\t")) {
                                        case "hello":
                                            var v142 int32 = 760
                                            _ = v142
                                        case "\n\t":
                                            if (-370 < 1) {
                                                if true {
                                                    v143 := GenS2{F0: false, F1: 171, F2: 255}
                                                    v1 = (0.0 * v1)
                                                    _ = v143
                                                } else {
                                                    if (!true) {
                                                        var v144 int16 = 1
                                                        var v145 uint64 = 2
                                                        _ = v144
                                                        _ = v145
                                                    }
                                                    if true {
                                                        if (false || true) {
                                                            v146 := GenS4{F0: 0}
                                                            _ = v146
                                                        }
                                                        v147 := v5[(("foo bar" + v129) + (v129 + v129))]
                                                        v147 += v147
                                                        _ = v147
                                                    } else {
                                                        v4 = append(v4, (("hello" + v129) + "hello"))
                                                        v1 = v1
                                                        if false {
                                                            var v148 uint = (v110 % (((v110 & (^(v110)))) | 1))
                                                            fmt.Println(v148)
                                                            _ = v148
                                                        } else {
                                                            v110 = (^(v110))
                                                            switch fi114 {
                                                            case 632:
                                                                v6 = (v6 << ((18446744073709551615) & 63))
                                                            case -2147483648:
                                                                if true {
                                                                    switch v129 {
                                                                    case "":
                                                                        var v149 int64 = 1
                                                                        _ = v149
                                                                    }
                                                                } else {
                                                                    if (true && (false || (1.5 != float32(1.5)))) {
                                                                        v3 -= v3
                                                                        fmt.Println(v1, v6, v110)
                                                                    } else {
                                                                        if true {
                                                                            var v150 int64 = -9223372036854775808
                                                                            var v151 float64 = -2.25
                                                                            _ = v150
                                                                            _ = v151
                                                                        }
                                                                        fmt.Println(v112)
                                                                        var v152 int8 = 100
                                                                        _ = v152
                                                                    }
                                                                    v112 -= (v112 & v112)
                                                                }
                                                            }
                                                        }
                                                    }
                                                    v6 = 65535
                                                }
                                            }
                                        default:
                                            var v153 bool = (fi114 != 780)
                                            _ = v153
                                        }
                                    } else {
                                        if true {
                                            v154 := v5[v129]
                                            switch v129 {
                                            case "hello":
                                                if (false && (!true)) {
                                                    var v155 int = (fi113 % ((fi113) | 1))
                                                    v3 = v3
                                                    _ = v155
                                                } else {
                                                    v3 -= v3
                                                    v129 = v129
                                                    var v156 uint64 = 42
                                                    _ = v156
                                                }
                                            default:
                                                v6 = 2147483647
                                            }
                                            v154 = (v154 % ((2147483648) | 1))
                                            v6 = v6
                                            _ = v154
                                        }
                                        var v157 int64 = -402
                                        var v158 int = 775
                                        _ = v157
                                        _ = v158
                                    }
                                    delete(v5, "")
                                    v159 := GenS14{F0: 326, F1: 383}
                                    _ = v159
                                }
                                if false {
                                    var v160 int64 = int64(v112)
                                    _ = v160
                                }
                                v4 = append(v4, v129)
                            } else {
                                var v161 int64 = -496
                                switch fi114 {
                                case 753:
                                    v1 -= float64(float32(v1))
                                case 32767:
                                    fmt.Println(v3, fi113)
                                case 426:
                                    v110 += v110
                                default:
                                    var v162 int64 = int64(v110)
                                    _ = v162
                                }
                                _ = v161
                            }
                            switch (^(fi113)) {
                            case -802:
                                v6 = v6
                            case -9223372036854775808:
                                switch v129 {
                                case "foo bar":
                                    var v163 int = fi114
                                    _ = v163
                                case "hello":
                                    v4 = append(v4, v129)
                                case "a":
                                    v164 := GenS13{F0: (v110 & 42)}
                                    _ = v164
                                default:
                                    switch ((v129 + v129) + v129) {
                                    case "":
                                        _ = len(v2)
                                    }
                                }
                                if (-899 <= int64(fi114)) {
                                    switch (fi113 >> ((uint(v6)) & 63)) {
                                    case -801:
                                        var v165 int = fi113
                                        _ = v165
                                    case -872:
                                        if (v129 != "foo bar") {
                                            v4 = append(v4, "hello")
                                            if true {
                                                v166 := GenS1{F0: 4294967295}
                                                _ = v166
                                            } else {
                                                if (4294967295 < fi113) {
                                                    v167 := GenS15{F0: 0, F1: "foo bar"}
                                                    _ = v167
                                                }
                                                v168 := GenS16{F0: (v6 >> ((0) & 63))}
                                                fmt.Println(v110, v6)
                                                _ = v168
                                            }
                                        }
                                    case 32768:
                                        if (true && (false && true)) {
                                            var v169 int32 = 0
                                            var v170 int16 = 1000
                                            _ = v169
                                            _ = v170
                                        }
                                    default:
                                        v171 := GenS6{F0: v112}
                                        _ = v171
                                    }
                                    v110 = v110
                                } else {
                                    v172 := GenS17{F0: 42, F1: 884}
                                    _ = len(v4)
                                    v173 := GenS18{F0: 42, F1: (fi114 >> ((v110) & 63))}
                                    _ = v172
                                    _ = v173
                                }
                                v174 := GenS7{F0: 0, F1: v6}
                                _ = v174
                            case 4294967295:
                                _ = len(v2)
                                if ((((v1 / (v1 / (v1 - (v1 / float64(((0.0 + (((0.1 / 1.0) + float32(float64((3.14159 / 1e10)))) / 0.1)) + 0.1)))))) + float64((0.1 + 3.14159))) / v1) < (v1 * ((0.0 * v1) / (((((-2.25 - ((-2.25 + float64((-2.25 / float32(v1)))) - 0.0)) / (3.14159 / 1.0)) - 0.0) * ((v1 - v1) * (1e10 * v1))) / (((3.14159 / v1) + -2.25) * ((v1 - v1) + float64(((3.14159 - float32(0.1)) + (-2.25 / (1.5 * 1e10)))))))))) {
                                    v175 := GenS17{F0: 2, F1: fi114}
                                    v3 -= 2
                                    switch fi114 {
                                    case -2147483648:
                                        v129 = "foo bar"
                                    }
                                    if (42 >= 366) {
                                        delete(v2, (v129 + ("" + "a")))
                                        if (127 < genFn1((v110 & (v110 & 2)), v3, (^(v3)))) {
                                            var v176 uint64 = 137
                                            delete(v5, v129)
                                            fmt.Println(v129)
                                            _ = v176
                                        } else {
                                            v177 := v2["foo bar"]
                                            var v178 bool = false
                                            _ = v177
                                            _ = v178
                                        }
                                        switch fi114 {
                                        case 128:
                                            if (false && true) {
                                                v179 := GenS13{F0: (v110 - 710)}
                                                var v180 int32 = 43
                                                v129 = (v129 + "hello")
                                                _ = v179
                                                _ = v180
                                            }
                                        case -373:
                                            var v181 string = "hello"
                                            _ = v181
                                        case 2147483647:
                                            if (3.14159 != ((v1 * v1) - -2.25)) {
                                                if len(v4) > 0 {
                                                    v4[0] = (("\n\t" + "foo bar") + v129)
                                                }
                                            }
                                        }
                                    } else {
                                        v5[v129] = (v6 ^ (^(v6)))
                                    }
                                    _ = v175
                                }
                                var v182 float64 = 1e10
                                _ = v182
                            }
                        }
                        v183 := GenS4{F0: -9223372036854775808}
                        _ = v183
                    }
                }
                var v184 int = (fi113 >> (((v110 / ((v110) | 1))) & 63))
                _ = fi114
                _ = v129
                _ = v184
            }
        }
        _ = fi113
    }
    fmt.Println(v110, v110, v3)
    v1 = (0.0 / (float64((3.14159 + (0.0 + (0.1 / 0.1)))) * v1))
    fmt.Println("v1:", v1)
    fmt.Println("v110:", v110)
    fmt.Println("v112:", v112)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v109
    _ = v110
    _ = v112
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
