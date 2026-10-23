package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 int64; F2 int8 }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 int64; F1 int }
type GenS7 struct { F0 string; F1 uint32; F2 string }
type GenS8 struct { F0 int16; F1 int16; F2 uint }
type GenS9 struct { F0 uint8; F1 uint }
type GenS10 struct { F0 uint32; F1 uint }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint32; F1 int16 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 uint8; F1 int32 }
type GenS15 struct { F0 string; F1 uint64 }
type GenS16 struct { F0 int16; F1 uint16 }
type GenS17 struct { F0 int32; F1 uint16 }
func genFn0() bool {
    var v5 uint16 = 18
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 uint16 = (v5 & (v5 * (^(v5))))
        if (-1 == 0) {
            if true {
                if false {
                    v7 += v7
                    v8 := []int{fi6, (fi6 | (fi6 & fi6))}
                    var v9 uint = 4294967295
                    _ = v8
                    _ = v9
                } else {
                    if false {
                        v10 := []int{fi6, fi6, 2147483647}
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        v10 = append(v10, (fi6 * (fi6 & fi6)))
                        _ = v10
                    } else {
                        v11 := GenS1{F0: 42, F1: 2}
                        for fi12 := 0; fi12 < 2; fi12++ {
                            fmt.Println(fi12)
                            v5 += v5
                            _ = fi12
                        }
                        v13 := GenS1{F0: int8(fi6), F1: 42}
                        _ = v11
                        _ = v13
                    }
                    for fi14 := 0; fi14 < 20; fi14++ {
                        var v15 string = ("\n\t" + "\n\t")
                        _ = fi14
                        _ = v15
                    }
                    for fi16 := 0; fi16 < 7; fi16++ {
                        v7 -= v5
                        _ = fi16
                    }
                }
                v5 = (v5 >> ((1) & 63))
            } else {
                v5 = (v7 >> ((1) & 63))
                v7 += 0
                fmt.Println(v7, v5, v7)
            }
        } else {
            v7 = (v7 << ((2) & 63))
            v7 = (v7 & v7)
        }
        _ = fi6
        _ = v7
    }
    v17 := GenS2{F0: 42, F1: 127}
    v18 := []int{960, -18}
    if true {
        switch 32767 {
        case -871:
            v18 = append(v18, -575)
        }
        var v19 float64 = 0.1
        _ = v19
    } else {
        v5 -= v5
        v5 = v5
    }
    _ = v17
    _ = v18
    _ = v5
    return (v5 >= v5)
}
func genFn1(p0 uint16) int {
    if false {
        v20 := []int{447, -2147483648, 128, 127}
        p0 += (p0 % ((32767) | 1))
        for fi21 := 0; fi21 < 18; fi21++ {
            if (!true) {
                if (false && true) {
                    v20 = append(v20, fi21)
                } else {
                    var v22 float64 = ((((0.0 - (3.14159 * 3.14159)) - 1e10) - (1.5 + (0.1 - float64((-2.25 + 0.1))))) - 1e10)
                    p0 += p0
                    _ = v22
                }
            } else {
                p0 = (p0 >> ((255) & 63))
                var v23 bool = (9223372036854775807 == -9223372036854775808)
                for fi24 := 0; fi24 < 10; fi24++ {
                    p0 += (p0 ^ 42)
                    _ = fi24
                }
                _ = v23
            }
            _ = fi21
        }
        _ = v20
    }
    v25 := []int{}
    p0 += p0
    v26 := GenS3{F0: 255}
    _ = p0
    _ = v25
    _ = v26
    return 9223372036854775807
}

func main() {
    var v1 int8 = 42
    var v2 uint32 = 0
    v3 := make(map[int]string)
    v4 := []int{744}
    var v27 uint = 557
    var v28 float64 = float64(0.0)
    v29 := GenS1{F0: (v1 | v1), F1: v2}
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v2, v1, v1)
    case "\n\t":
        var v30 int8 = (v1 & (v1 - v1))
        _ = v30
    case "hello":
        fmt.Println(v27)
    }
    v31 := GenS3{F0: 65}
    v32 := GenS3{F0: 176}
    if false {
        switch ("\n\t" + "a") {
        case "a":
            var v33 uint64 = 629
            _ = v33
        default:
            for fi34 := 0; fi34 < 12; fi34++ {
                v28 = float64(1.5)
                _ = fi34
            }
        }
        if ("foo bar" >= "hello") {
            if true {
                v4 = append(v4, 715)
            }
            v1 -= 0
        } else {
            var v35 int = -706
            for fi36 := 0; fi36 < 8; fi36++ {
                fmt.Println(fi36, v2, v28)
                _ = fi36
            }
            v37 := GenS4{F0: v1, F1: -9223372036854775808, F2: v1}
            _ = v35
            _ = v37
        }
        var v38 uint8 = 42
        _ = v38
    }
    v39 := v3[751]
    v40 := GenS2{F0: 252, F1: (v1 | (v1 << (((v27 ^ v27)) & 63)))}
    if len(v4) > 0 {
        v4[0] = genFn1(65535)
    }
    if (1 < 182) {
        var v41 string = ("a" + v39)
        v42 := GenS5{F0: (v27 - v27), F1: v27}
        var v43 float32 = float32(v28)
        _ = v41
        _ = v42
        _ = v43
    }
    if false {
        var v44 int16 = 1000
        v44 += (^(v44))
        _ = v44
    } else {
        for fi45 := 0; fi45 < 14; fi45++ {
            v3[(fi45 << (((^(v27))) & 63))] = v39
            _ = fi45
        }
        if false {
            var v46 float64 = float64(0.1)
            for fi47 := 0; fi47 < 9; fi47++ {
                v28 -= (float64((1e10 * 0.0)) + v28)
                var v48 int64 = 0
                _ = fi47
                _ = v48
            }
            _ = v46
        } else {
            for fi49 := 0; fi49 < 17; fi49++ {
                switch fi49 {
                case -2147483648:
                    switch (-(fi49)) {
                    case 721:
                        v28 = ((0.1 - float64(0.0)) + v28)
                    case 2147483647:
                        v39 = v39
                    }
                case 0:
                    var v50 uint64 = 897
                    _ = v50
                case -1:
                    v2 = (^(v2))
                default:
                    if true {
                        fmt.Println(v28)
                        v51 := GenS3{F0: 42}
                        for fi52 := 0; fi52 < 11; fi52++ {
                            v27 += v27
                            _ = fi52
                        }
                        _ = v51
                    }
                }
                _ = len(v3)
                _ = fi49
            }
            for fi53 := 0; fi53 < 1; fi53++ {
                for fi54 := 0; fi54 < 17; fi54++ {
                    if (199 != 1) {
                        v28 += v28
                    }
                    _ = fi54
                }
                var v55 string = ("" + v39)
                _ = fi53
                _ = v55
            }
        }
        if false {
            if true {
                v56 := GenS5{F0: v27, F1: (v27 >> (((v27 >> (((v27 ^ (v27 >> (((^(v27))) & 63)))) & 63))) & 63))}
                var v57 bool = (false && (2 != 300))
                fmt.Println(v27)
                _ = v56
                _ = v57
            }
            for fi58 := 0; fi58 < 3; fi58++ {
                v59 := GenS5{F0: v27, F1: v27}
                v1 += -128
                for fi60 := 0; fi60 < 15; fi60++ {
                    v61 := GenS5{F0: (^(v27)), F1: (^(v27))}
                    switch "hello" {
                    case "":
                        var v62 float64 = 3.14159
                        _ = v62
                    case "a":
                        var v63 bool = true
                        _ = v63
                    }
                    _ = fi60
                    _ = v61
                }
                _ = fi58
                _ = v59
            }
            _ = len(v3)
        } else {
            for fi64 := 0; fi64 < 1; fi64++ {
                switch (fi64 ^ fi64) {
                case 1:
                    fmt.Println(v28, v39, v27)
                case -973:
                    v27 += 18446744073709551615
                case -532:
                    v1 -= (v1 >> ((v27) & 63))
                }
                switch "" {
                case "":
                    for fi65 := 0; fi65 < 10; fi65++ {
                        if false {
                            if (true || (47 != 27)) {
                                var v66 int64 = 42
                                var v67 uint64 = 1
                                switch genFn1(0) {
                                case 1:
                                    if (0.0 > (float32(v28) * 1e10)) {
                                        switch fi65 {
                                        case -2147483648:
                                            var v68 int8 = v1
                                            _ = v68
                                        case 42:
                                            var v69 bool = (v28 > (0.1 + 0.1))
                                            _ = v69
                                        case 9223372036854775807:
                                            v2 += v2
                                        default:
                                            switch v39 {
                                            case "hello":
                                                v39 = v39
                                            case "\n\t":
                                                v27 += (v27 - 18446744073709551615)
                                            case "":
                                                if (true && ((!true) && true)) {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                            }
                                        }
                                        v28 = v28
                                        _ = len(v3)
                                        v3[fi64] = ("\n\t" + (v39 + v39))
                                    } else {
                                        if false {
                                            _ = len(v4)
                                            var v70 int64 = 0
                                            _ = v70
                                        } else {
                                            v66 -= 130
                                            var v71 int8 = (v1 ^ (v1 - (v1 ^ (v1 + 100))))
                                            v67 = v67
                                            _ = v71
                                        }
                                    }
                                case 127:
                                    if true {
                                        v72 := GenS6{F0: (v66 << (((v27 >> ((1) & 63))) & 63)), F1: 1}
                                        v28 = v28
                                        fmt.Println(v2, v1)
                                        _ = v72
                                    } else {
                                        v73 := GenS6{F0: v66, F1: (fi65 >> ((v27) & 63))}
                                        switch ((v39 + v39) + "foo bar") {
                                        case "\n\t":
                                            v74 := GenS3{F0: 255}
                                            _ = v74
                                        case "a":
                                            switch ("\n\t" + "foo bar") {
                                            case "":
                                                v75 := GenS7{F0: "hello", F1: (v2 - v2), F2: "a"}
                                                _ = v75
                                            }
                                        case "foo bar":
                                            v76 := GenS4{F0: int8(v27), F1: v66, F2: v1}
                                            _ = v76
                                        }
                                        _ = v73
                                    }
                                default:
                                    delete(v3, (fi64 - fi65))
                                }
                                _ = v66
                                _ = v67
                            } else {
                                v77 := GenS7{F0: v39, F1: (v2 / (((v2 ^ 0)) | 1)), F2: (v39 + (("hello" + "foo bar") + "hello"))}
                                _ = v77
                            }
                            if (uint16(v27) == 280) {
                                if false {
                                    switch ((("hello" + v39) + "foo bar") + "") {
                                    case "foo bar":
                                        var v78 uint64 = 18446744073709551615
                                        _ = v78
                                    default:
                                        v1 = 42
                                        if (true && true) {
                                            var v79 uint32 = v2
                                            _ = v79
                                        } else {
                                            var v80 string = "foo bar"
                                            fmt.Println(v28, v39)
                                            v81 := v3[int(v2)]
                                            _ = v80
                                            _ = v81
                                        }
                                    }
                                    switch "a" {
                                    case "\n\t":
                                        v39 = "\n\t"
                                    case "foo bar":
                                        v2 -= 2
                                    default:
                                        var v82 int32 = 42
                                        _ = v82
                                    }
                                }
                                var v83 float64 = ((1e10 + 1e10) + (1.5 * v28))
                                v28 = v83
                                var v84 uint32 = (v2 % (((^(v2))) | 1))
                                var v85 int16 = 1
                                _ = v83
                                _ = v84
                                _ = v85
                            }
                            if false {
                                switch genFn1(450) {
                                case 1:
                                    switch genFn1(0) {
                                    case -194:
                                        fmt.Println(v28, v1, v2)
                                    case -839:
                                        v1 += 42
                                    case 128:
                                        v28 += 0.1
                                    default:
                                        fmt.Println(v1)
                                    }
                                case 224:
                                    switch 1 {
                                    case 2147483647:
                                        v3[fi65] = ("hello" + "\n\t")
                                    case -2147483648:
                                        v86 := GenS8{F0: 32767, F1: 1, F2: (^(v27))}
                                        _ = v86
                                    }
                                default:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                switch fi64 {
                                case -2147483648:
                                    var v87 uint16 = 0
                                    _ = v87
                                case -137:
                                    if (2 > 92) {
                                        var v88 float64 = v28
                                        v27 -= v27
                                        _ = v88
                                    } else {
                                        v89 := v3[-47]
                                        if true {
                                            fmt.Println(v2, v2)
                                            var v90 int32 = 42
                                            _ = v90
                                        }
                                        fmt.Println(v89, v2)
                                        _ = v89
                                    }
                                case 127:
                                    v91 := v3[(fi64 | genFn1(793))]
                                    _ = v91
                                default:
                                    if true {
                                        var v92 uint16 = 1
                                        _ = v92
                                    }
                                }
                                fmt.Println(v2, v39)
                            } else {
                                fmt.Println(v39, v2)
                                if false {
                                    if (!(-1 >= 0)) {
                                        v1 -= v1
                                        if len(v4) > 0 {
                                            v4[0] = (-(fi65))
                                        }
                                        switch v39 {
                                        case "hello":
                                            if true {
                                                if ((v1 & (v1 / ((-1) | 1))) == v1) {
                                                    var v93 int8 = v1
                                                    v93 += -1
                                                    _ = v93
                                                }
                                                if false {
                                                    v94 := GenS2{F0: 305, F1: v1}
                                                    v95 := GenS7{F0: v39, F1: 792, F2: (v39 + ((v39 + v39) + v39))}
                                                    if true {
                                                        _ = len(v4)
                                                        if false {
                                                            var v96 uint = v27
                                                            if true {
                                                                v97 := GenS9{F0: 2, F1: v96}
                                                                _ = v97
                                                            } else {
                                                                v39 = "foo bar"
                                                                if (((3.14159 * (((1e10 / 0.1) * -2.25) / (1e10 / -2.25))) < 3.14159) && false) {
                                                                    v39 = v39
                                                                    v98 := GenS3{F0: 77}
                                                                    var v99 uint8 = 2
                                                                    _ = v98
                                                                    _ = v99
                                                                } else {
                                                                    v96 += (v27 - (v96 % ((v96) | 1)))
                                                                    switch v39 {
                                                                    case "\n\t":
                                                                        v2 = v2
                                                                    case "":
                                                                        switch ("\n\t" + v39) {
                                                                        case "\n\t":
                                                                            v100 := GenS10{F0: v2, F1: 0}
                                                                            _ = v100
                                                                        }
                                                                    case "a":
                                                                        var v101 uint = (^(v27))
                                                                        _ = v101
                                                                    }
                                                                }
                                                                v96 = (v96 | 96)
                                                            }
                                                            _ = v96
                                                        }
                                                    } else {
                                                        var v102 float64 = v28
                                                        v39 = (v39 + (("foo bar" + v39) + v39))
                                                        _ = v102
                                                    }
                                                    _ = v94
                                                    _ = v95
                                                } else {
                                                    switch (fi64 << (((v27 - (v27 | 2))) & 63)) {
                                                    case -723:
                                                        if (true || false) {
                                                            var v103 uint64 = 1
                                                            _ = v103
                                                        }
                                                    default:
                                                        fmt.Println(fi65, v28)
                                                    }
                                                }
                                            } else {
                                                v27 = (^(v27))
                                                fmt.Println(v39, v28, fi64)
                                                v2 = v2
                                            }
                                            var v104 uint = 2
                                            fmt.Println(fi65)
                                            _ = v104
                                        case "foo bar":
                                            var v105 int64 = -9223372036854775808
                                            switch fi65 {
                                            case 32768:
                                                if (v39 != (v39 + v39)) {
                                                    var v106 int8 = v1
                                                    switch fi64 {
                                                    case 4294967295:
                                                        v1 += (v1 >> ((v27) & 63))
                                                    case 571:
                                                        fmt.Println(v106, v27, v1)
                                                    case -2147483648:
                                                        switch (fi65 / ((fi64) | 1)) {
                                                        case 225:
                                                            var v107 uint64 = 314
                                                            _ = v107
                                                        case 2147483647:
                                                            var v108 uint = 65535
                                                            _ = v108
                                                        case 127:
                                                            v1 = v1
                                                        default:
                                                            var v109 uint16 = 0
                                                            _ = v109
                                                        }
                                                    default:
                                                        var v110 uint8 = 42
                                                        _ = v110
                                                    }
                                                    switch ((("a" + v39) + ("foo bar" + ("" + v39))) + "") {
                                                    case "a":
                                                        if false {
                                                            v106 = 6
                                                            fmt.Println(v105, v28)
                                                        }
                                                    default:
                                                        var v111 int32 = 1
                                                        _ = v111
                                                    }
                                                    var v112 float64 = v28
                                                    _ = v106
                                                    _ = v112
                                                } else {
                                                    v1 = 100
                                                }
                                            case -2147483648:
                                                switch v39 {
                                                case "a":
                                                    v28 -= v28
                                                case "hello":
                                                    v113 := GenS7{F0: ("a" + ("" + v39)), F1: 327, F2: v39}
                                                    _ = v113
                                                }
                                            case -1:
                                                v114 := GenS11{F0: 2147483647}
                                                _ = v114
                                            default:
                                                v115 := GenS12{F0: 4294967295, F1: 1}
                                                _ = v115
                                            }
                                            _ = v105
                                        default:
                                            if (((v1 - v1) > v1) && true) {
                                                if ((((v28 + (((v28 - v28) / float64(3.14159)) + (((v28 * 1e10) * ((float64((1e10 - float32(((v28 - v28) * 1.5)))) - v28) / 3.14159)) - 1e10))) * v28) * -2.25) == 1e10) {
                                                    v116 := GenS13{F0: 2}
                                                    _ = v116
                                                } else {
                                                    v27 -= v27
                                                    fmt.Println(v1, v39, v1)
                                                    v4 = append(v4, 127)
                                                    var v117 bool = (127 == 32)
                                                    _ = v117
                                                }
                                                v39 = ((v39 + v39) + (v39 + v39))
                                                if false {
                                                    v3[fi65] = (((((v39 + ("" + v39)) + ("hello" + (v39 + v39))) + "\n\t") + v39) + "a")
                                                }
                                                v27 = 44
                                                var v118 uint8 = 127
                                                var v119 uint16 = uint16(v27)
                                                fmt.Println(v1, v28, v118)
                                                _ = len(v4)
                                                v39 = v39
                                                v28 = (v28 + (3.14159 + (float64(1e10) - 0.1)))
                                                _ = v118
                                                _ = v119
                                            } else {
                                                switch 2147483647 {
                                                case 9223372036854775807:
                                                    _ = len(v4)
                                                default:
                                                    if (1 < v2) {
                                                        v3[fi65] = v39
                                                        v1 = 122
                                                    } else {
                                                        fmt.Println(v28, fi64)
                                                    }
                                                }
                                            }
                                        }
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                    if false {
                                        if (false && true) {
                                            v120 := GenS14{F0: 2, F1: 2147483647}
                                            if true {
                                                var v121 uint16 = 911
                                                var v122 float32 = 3.14159
                                                _ = v121
                                                _ = v122
                                            }
                                            _ = v120
                                        }
                                        switch v39 {
                                        case "":
                                            var v123 int8 = v1
                                            _ = v123
                                        }
                                        var v124 uint32 = v2
                                        _ = v124
                                    }
                                    v2 += (v2 / (((v2 - (v2 / ((65535) | 1)))) | 1))
                                    var v125 uint8 = 0
                                    var v126 uint64 = uint64(v27)
                                    var v127 uint8 = (v125 >> ((133) & 63))
                                    if (!false) {
                                        var v128 uint64 = (v126 % (((v126 * v126)) | 1))
                                        _ = v128
                                    }
                                    _ = v125
                                    _ = v126
                                    _ = v127
                                } else {
                                    switch (fi64 + fi64) {
                                    case 2147483647:
                                        var v129 uint64 = 42
                                        _ = v129
                                    case 32767:
                                        delete(v3, (-(fi64)))
                                    case 1:
                                        v130 := GenS15{F0: v39, F1: 42}
                                        _ = v130
                                    default:
                                        if false {
                                            if true {
                                                var v131 int32 = 771
                                                fmt.Println(v39, v2, v28)
                                                v27 += (v27 << ((v27) & 63))
                                                v132 := v3[(fi64 - 128)]
                                                _ = v131
                                                _ = v132
                                            }
                                            v133 := v3[(fi64 / ((fi64) | 1))]
                                            v133 = v39
                                            v2 -= (v2 - (v2 ^ 42))
                                            if true {
                                                fmt.Println(v2, fi65, v2)
                                            } else {
                                                if len(v4) > 0 {
                                                    v4[0] = (-(fi64))
                                                }
                                                v28 = v28
                                                var v134 int = genFn1(559)
                                                _ = v134
                                            }
                                            _ = v133
                                        }
                                    }
                                    var v135 int = (fi65 % ((fi64) | 1))
                                    v136 := GenS1{F0: (v1 / ((v1) | 1)), F1: (v2 * v2)}
                                    _ = v135
                                    _ = v136
                                }
                            }
                            v28 = ((v28 * v28) - (v28 * 1e10))
                        }
                        var v137 uint = (v27 >> ((v27) & 63))
                        switch fi64 {
                        case 313:
                            delete(v3, 128)
                        case 0:
                            v138 := GenS16{F0: 1, F1: 2}
                            _ = v138
                        }
                        _ = fi65
                        _ = v137
                    }
                case "a":
                    v39 = v39
                }
                _ = fi64
            }
        }
    }
    var v139 float32 = (3.14159 * 3.14159)
    if false {
        var v140 uint = v27
        _ = v140
    }
    if false {
        v141 := GenS17{F0: 1, F1: 42}
        for fi142 := 0; fi142 < 1; fi142++ {
            var v143 uint16 = 65535
            v144 := GenS10{F0: v2, F1: (v27 + uint(v143))}
            _ = fi142
            _ = v143
            _ = v144
        }
        v145 := v3[0]
        _ = v141
        _ = v145
    } else {
        v28 -= v28
        _ = len(v4)
        var v146 bool = (0 <= 42)
        _ = v146
    }
    v28 = v28
    var v147 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v139:", v139)
    fmt.Println("v147:", v147)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v139
    _ = v147
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v39
    _ = v4
    _ = v40
}
