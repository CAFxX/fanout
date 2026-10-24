package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint16; F1 uint16; F2 int }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint; F1 string }
type GenS7 struct { F0 uint8; F1 int; F2 uint64 }
type GenS8 struct { F0 int32; F1 uint16; F2 int16 }
type GenS9 struct { F0 bool; F1 uint8; F2 uint }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint16; F1 bool; F2 int32 }
type GenS13 struct { F0 bool; F1 int16; F2 int16 }
type GenS14 struct { F0 uint8; F1 int64; F2 string }
type GenS15 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS16 struct { F0 uint32; F1 uint; F2 uint16 }
func genFn0() int8 {
    v6 := GenS1{F0: 523}
    var v7 uint8 = 1
    _ = v6
    _ = v7
    return 100
}
func genFn1(p0 string, p1 int8) int {
    var v8 float64 = 0.0
    v9 := GenS1{F0: -253}
    var v10 float32 = 0.1
    p1 = (^(p1))
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 int32 = -2147483648
    var v2 uint = 255
    var v3 uint = v2
    var v4 int8 = -128
    var v5 int8 = v4
    v11 := GenS2{F0: "", F1: -32768}
    var v12 int = -2147483648
    v13 := GenS1{F0: int(v3)}
    if ((65535 <= 2147483647) && true) {
        for fi14 := 0; fi14 < 1; fi14++ {
            if false {
                var v15 int64 = 0
                _ = v15
            }
            v16 := []int{1}
            var v17 uint16 = 39
            _ = fi14
            _ = v16
            _ = v17
        }
        switch -331 {
        case 42:
            switch (^(v12)) {
            case 32767:
                if false {
                    v18 := GenS3{F0: 32767, F1: 65535, F2: genFn1("a", 87)}
                    _ = v18
                } else {
                    switch v12 {
                    case 4294967295:
                        var v19 uint16 = 65535
                        _ = v19
                    case -2147483648:
                        if (((v12 * (^(v12))) >= (-(v12))) || true) {
                            v20 := GenS3{F0: 65535, F1: 2, F2: (v12 - v12)}
                            fmt.Println(v3, v2)
                            for fi21 := 0; fi21 < 1; fi21++ {
                                if ((!false) || ((906 < 434) && ((42 == 1) || (false && (32767 != -32768))))) {
                                    v22 := GenS4{F0: 42}
                                    _ = v22
                                } else {
                                    switch ("a" + "") {
                                    case "hello":
                                        for fi23 := 0; fi23 < 11; fi23++ {
                                            var v24 float64 = 1e10
                                            if (v3 > 0) {
                                                switch fi21 {
                                                case -9223372036854775808:
                                                    var v25 uint8 = 1
                                                    _ = v25
                                                default:
                                                    v5 += genFn0()
                                                }
                                            }
                                            fmt.Println(fi21, v12, v4)
                                            v26 := GenS3{F0: 972, F1: 2, F2: (-(fi21))}
                                            var v27 int16 = -1
                                            _ = fi23
                                            _ = v24
                                            _ = v26
                                            _ = v27
                                        }
                                    case "foo bar":
                                        switch v12 {
                                        case 4294967295:
                                            var v28 float64 = ((0.0 * (3.14159 / (0.1 * 3.14159))) / ((3.14159 * float64(3.14159)) - 0.1))
                                            _ = v28
                                        case 0:
                                            v29 := []int{(v12 | (v12 >> (((v2 + (^(v2)))) & 63)))}
                                            _ = v29
                                        case -447:
                                            var v30 float64 = (0.1 / 1.0)
                                            _ = v30
                                        default:
                                            v1 -= (-(v1))
                                        }
                                    default:
                                        switch "a" {
                                        case "a":
                                            v31 := []int{128, (fi21 >> (((v3 - v3)) & 63))}
                                            _ = v31
                                        default:
                                            for fi32 := 0; fi32 < 20; fi32++ {
                                                var v33 int = 42
                                                switch (v12 + (v12 & genFn1("\n\t", (v5 + genFn0())))) {
                                                case 127:
                                                    switch (-(v33)) {
                                                    case 1:
                                                        v34 := []int{fi32, 1, (-(fi21))}
                                                        _ = v34
                                                    default:
                                                        v35 := GenS1{F0: (fi21 / (((v33 & (v12 << ((390) & 63)))) | 1))}
                                                        _ = v35
                                                    }
                                                }
                                                var v36 float64 = (0.0 - 0.1)
                                                fmt.Println(v1, v36)
                                                var v37 uint32 = 65535
                                                _ = fi32
                                                _ = v33
                                                _ = v36
                                                _ = v37
                                            }
                                        }
                                    }
                                    if true {
                                        v3 = (v3 - (v3 >> (((^(v3))) & 63)))
                                        switch "hello" {
                                        case "hello":
                                            v38 := GenS5{F0: -9223372036854775808}
                                            _ = v38
                                        case "\n\t":
                                            v39 := []int{(-(v12)), int(v1), (^(fi21))}
                                            _ = v39
                                        }
                                    } else {
                                        for fi40 := 0; fi40 < 14; fi40++ {
                                            v1 = v1
                                            _ = fi40
                                        }
                                        switch "foo bar" {
                                        case "foo bar":
                                            if false {
                                                var v41 uint16 = 255
                                                v42 := []int{fi21}
                                                v2 += 660
                                                _ = v41
                                                _ = v42
                                            } else {
                                                var v43 float32 = 0.0
                                                switch int(v1) {
                                                case 621:
                                                    if (65535 <= 0) {
                                                        for fi44 := 0; fi44 < 8; fi44++ {
                                                            v45 := GenS6{F0: 896, F1: ""}
                                                            v43 += float32(1.5)
                                                            _ = fi44
                                                            _ = v45
                                                        }
                                                        switch 793 {
                                                        case -1:
                                                            var v46 int = v12
                                                            _ = v46
                                                        }
                                                        for fi47 := 0; fi47 < 13; fi47++ {
                                                            if true {
                                                                v48 := []int{32768, 128, genFn1("", (^(v4))), (fi21 % ((v12) | 1))}
                                                                _ = v48
                                                            } else {
                                                                v49 := GenS7{F0: 2, F1: fi21, F2: 1}
                                                                var v50 uint64 = 565
                                                                _ = v49
                                                                _ = v50
                                                            }
                                                            v51 := GenS8{F0: (-(v1)), F1: 2, F2: -868}
                                                            _ = fi47
                                                            _ = v51
                                                        }
                                                    } else {
                                                        var v52 float64 = -2.25
                                                        for fi53 := 0; fi53 < 9; fi53++ {
                                                            var v54 int = genFn1("hello", v5)
                                                            if true {
                                                                v55 := GenS6{F0: 65535, F1: (((("" + "\n\t") + "") + "foo bar") + ((("\n\t" + "\n\t") + ("" + "foo bar")) + "hello"))}
                                                                if false {
                                                                    fmt.Println(fi53)
                                                                    var v56 int8 = (v5 * v5)
                                                                    var v57 string = "foo bar"
                                                                    v57 = v57
                                                                    _ = v56
                                                                    _ = v57
                                                                } else {
                                                                    var v58 bool = (519 <= fi53)
                                                                    var v59 uint64 = uint64(v2)
                                                                    _ = v58
                                                                    _ = v59
                                                                }
                                                                if true {
                                                                    switch (fi53 + v54) {
                                                                    case 1:
                                                                        v1 += v1
                                                                    default:
                                                                        if true {
                                                                            var v60 uint = 896
                                                                            _ = v60
                                                                        } else {
                                                                            v2 = (v2 * v3)
                                                                        }
                                                                    }
                                                                    v5 += (v4 % (((v4 % ((100) | 1))) | 1))
                                                                    var v61 uint = (v3 - (^(v2)))
                                                                    _ = v61
                                                                } else {
                                                                    if true {
                                                                        var v62 uint16 = 255
                                                                        var v63 int16 = 42
                                                                        var v64 float32 = v43
                                                                        v65 := []int{v54, (fi21 % ((-1) | 1)), (^(fi53))}
                                                                        switch "" {
                                                                        case "":
                                                                            _ = len(v65)
                                                                            if len(v65) > 0 {
                                                                                v65[0] = (fi53 + -655)
                                                                            }
                                                                        case "hello":
                                                                            var v66 uint32 = 2147483647
                                                                            _ = v66
                                                                        case "foo bar":
                                                                            v65 = append(v65, (v12 << (((v3 << (((v3 >> ((v2) & 63))) & 63))) & 63)))
                                                                            if true {
                                                                                var v67 int = -287
                                                                                _ = v67
                                                                            } else {
                                                                                if false {
                                                                                    fmt.Println(v1, v5)
                                                                                    switch "hello" {
                                                                                    case "a":
                                                                                        if ((2 == 0) || false) {
                                                                                            v68 := GenS5{F0: 228}
                                                                                            fmt.Println(v62)
                                                                                            _ = v68
                                                                                        }
                                                                                    case "\n\t":
                                                                                        if ((!false) || (2 != 4294967295)) {
                                                                                            var v69 int8 = v5
                                                                                            _ = v69
                                                                                        } else {
                                                                                            v70 := GenS9{F0: (((v52 + (((v52 * -2.25) * ((((v52 / v52) / (float64(((v43 * float32(((v52 * 0.0) + 0.0))) * v64)) + v52)) / 0.1) - 3.14159)) * (v52 * 1.5))) + 0.0) == v52), F1: 42, F2: (^(v3))}
                                                                                            var v71 bool = true
                                                                                            var v72 uint16 = v62
                                                                                            v73 := GenS8{F0: (v1 << (((v3 & 358)) & 63)), F1: v72, F2: (v63 * v63)}
                                                                                            v52 -= v52
                                                                                            v74 := GenS6{F0: (^(v3)), F1: ("" + "\n\t")}
                                                                                            _ = v70
                                                                                            _ = v71
                                                                                            _ = v72
                                                                                            _ = v73
                                                                                            _ = v74
                                                                                        }
                                                                                    default:
                                                                                        var v75 int = genFn1(("" + "foo bar"), int8(v63))
                                                                                        _ = v75
                                                                                    }
                                                                                    var v76 float32 = (v64 + float32(-2.25))
                                                                                    _ = v76
                                                                                } else {
                                                                                    var v77 uint8 = uint8(v62)
                                                                                    if ((v52 * (1.5 / 0.1)) == ((v52 / v52) * v52)) {
                                                                                        fmt.Println(v43, v12, v3)
                                                                                    } else {
                                                                                        v2 -= 2
                                                                                        if len(v65) >= 2 {
                                                                                            v65 = v65[1:len(v65)]
                                                                                        }
                                                                                        v62 += v62
                                                                                    }
                                                                                    fmt.Println(fi21)
                                                                                    _ = v77
                                                                                }
                                                                                var v78 float32 = float32(v52)
                                                                                if len(v65) > 0 {
                                                                                    v65[0] = fi53
                                                                                }
                                                                                fmt.Println(v12, v54, fi53)
                                                                                var v79 int8 = (v4 * (^(v4)))
                                                                                var v80 uint = (v2 | (v3 - (v3 / ((v3) | 1))))
                                                                                var v81 int = fi53
                                                                                _ = v78
                                                                                _ = v79
                                                                                _ = v80
                                                                                _ = v81
                                                                            }
                                                                            fmt.Println(v3)
                                                                        default:
                                                                            v1 = (-(v1))
                                                                        }
                                                                        var v82 uint8 = 255
                                                                        v4 += (v5 - (v5 ^ -110))
                                                                        _ = v62
                                                                        _ = v63
                                                                        _ = v64
                                                                        _ = v65
                                                                        _ = v82
                                                                    }
                                                                }
                                                                _ = v55
                                                            }
                                                            _ = fi53
                                                            _ = v54
                                                        }
                                                        _ = v52
                                                    }
                                                case -931:
                                                    switch v12 {
                                                    case -429:
                                                        var v83 uint32 = 4294967295
                                                        _ = v83
                                                    case 32767:
                                                        v5 -= v4
                                                    default:
                                                        v43 = float32((-2.25 - 0.1))
                                                    }
                                                case 42:
                                                    v1 -= (-(v1))
                                                default:
                                                    if (!true) {
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            var v84 int64 = -803
                                                            _ = v84
                                                        case "":
                                                            v85 := GenS6{F0: v2, F1: "a"}
                                                            _ = v85
                                                        case "a":
                                                            if (!false) {
                                                                v4 -= (v5 & (-(v4)))
                                                            }
                                                        default:
                                                            v12 += (fi21 << (((^(v3))) & 63))
                                                        }
                                                    }
                                                }
                                                _ = v43
                                            }
                                        case "a":
                                            v3 = (v2 - (v2 / (((v2 | v2)) | 1)))
                                        }
                                    }
                                }
                                _ = fi21
                            }
                            _ = v20
                        } else {
                            for fi86 := 0; fi86 < 16; fi86++ {
                                v12 -= genFn1(((("\n\t" + ("foo bar" + ("a" + "foo bar"))) + "hello") + "a"), (v4 + v5))
                                _ = fi86
                            }
                        }
                    }
                }
            case 573:
                var v87 float32 = 0.0
                _ = v87
            default:
                v88 := GenS10{F0: 4294967295}
                _ = v88
            }
        default:
            var v89 float64 = 1e10
            _ = v89
        }
    } else {
        v12 += v12
        if true {
            v90 := []int{}
            if len(v90) > 0 {
                v90[0] = -477
            }
            switch genFn1("", (^(v4))) {
            case 783:
                if false {
                    switch genFn1("", v4) {
                    case 24:
                        switch "" {
                        case "":
                            for fi91 := 0; fi91 < 11; fi91++ {
                                v1 -= (v1 * (v1 * (v1 + v1)))
                                _ = fi91
                            }
                        case "hello":
                            v1 += 0
                        case "\n\t":
                            for fi92 := 0; fi92 < 2; fi92++ {
                                fmt.Println(v2, v1, v2)
                                _ = fi92
                            }
                        }
                    case -39:
                        v93 := GenS3{F0: 1, F1: 65535, F2: 127}
                        _ = v93
                    default:
                        var v94 uint64 = 1
                        _ = v94
                    }
                    v95 := GenS2{F0: ("a" + ""), F1: -1}
                    v1 = (v1 ^ -2147483648)
                    _ = v95
                } else {
                    v2 -= 0
                    for fi96 := 0; fi96 < 3; fi96++ {
                        if true {
                            _ = len(v90)
                        }
                        if false {
                            if len(v90) >= 2 {
                                v90 = v90[1:len(v90)]
                            }
                            switch fi96 {
                            case 1:
                                for fi97 := 0; fi97 < 11; fi97++ {
                                    switch genFn1("\n\t", -128) {
                                    case -436:
                                        if false {
                                            if true {
                                                var v98 uint = v3
                                                _ = v98
                                            } else {
                                                if (float64(0.1) < -2.25) {
                                                    if (("foo bar" + "") == "") {
                                                        var v99 uint64 = 71
                                                        if ((false && (4294967295 == 1)) || true) {
                                                            var v100 uint = (v3 & 467)
                                                            if (786 > 65535) {
                                                                if len(v90) > 0 {
                                                                    v90[0] = genFn1(("a" + ""), (v4 - (-(v5))))
                                                                }
                                                                var v101 float64 = 1.5
                                                                _ = v101
                                                            } else {
                                                                if (18446744073709551615 > (v99 / ((278) | 1))) {
                                                                    switch (("" + ("hello" + "a")) + "foo bar") {
                                                                    case "":
                                                                        v2 = 18446744073709551615
                                                                    case "\n\t":
                                                                        switch (fi96 % ((genFn1("foo bar", v5)) | 1)) {
                                                                        case 1:
                                                                            if false {
                                                                                v99 += (^(v99))
                                                                                v102 := GenS11{F0: fi96}
                                                                                if len(v90) > 0 {
                                                                                    v90[0] = 1
                                                                                }
                                                                                v100 = v2
                                                                                v90 = append(v90, fi96)
                                                                                _ = v102
                                                                            }
                                                                        case 4294967295:
                                                                            var v103 uint64 = (v99 - (v99 ^ (^(v99))))
                                                                            _ = v103
                                                                        default:
                                                                            var v104 uint32 = 297
                                                                            var v105 int8 = v4
                                                                            _ = v104
                                                                            _ = v105
                                                                        }
                                                                    default:
                                                                        v106 := GenS1{F0: (fi97 % (((v12 + genFn1("hello", 42))) | 1))}
                                                                        _ = v106
                                                                    }
                                                                    v107 := GenS3{F0: uint16(v2), F1: 42, F2: fi96}
                                                                    if true {
                                                                        switch (("" + ("hello" + "\n\t")) + "") {
                                                                        case "foo bar":
                                                                            v108 := GenS11{F0: fi97}
                                                                            _ = v108
                                                                        case "hello":
                                                                            switch "\n\t" {
                                                                            case "a":
                                                                                var v109 int64 = 43
                                                                                _ = v109
                                                                            }
                                                                        case "a":
                                                                            v1 = (v1 << (((v3 ^ (v100 % ((v3) | 1)))) & 63))
                                                                        default:
                                                                            v110 := GenS12{F0: 32767, F1: (false && false), F2: (v1 * 42)}
                                                                            _ = v110
                                                                        }
                                                                        if len(v90) >= 2 {
                                                                            v90 = v90[1:len(v90)]
                                                                        }
                                                                    } else {
                                                                        v90 = append(v90, genFn1("foo bar", v5))
                                                                        v12 -= v12
                                                                        var v111 float32 = -2.25
                                                                        _ = v111
                                                                    }
                                                                    _ = v107
                                                                } else {
                                                                    v90 = append(v90, (fi97 * genFn1("", genFn0())))
                                                                    _ = len(v90)
                                                                }
                                                                if len(v90) > 0 {
                                                                    v90[0] = v12
                                                                }
                                                                v5 += v4
                                                            }
                                                            _ = v100
                                                        }
                                                        _ = v99
                                                    }
                                                    switch "foo bar" {
                                                    case "a":
                                                        var v112 uint16 = 32767
                                                        _ = v112
                                                    default:
                                                        var v113 int16 = 32767
                                                        _ = v113
                                                    }
                                                    var v114 float32 = ((3.14159 / 1.0) + -2.25)
                                                    _ = v114
                                                }
                                                _ = len(v90)
                                                v4 += v4
                                            }
                                            var v115 uint64 = 2
                                            _ = v115
                                        }
                                    }
                                    if (0 >= 65535) {
                                        v116 := GenS13{F0: false, F1: -350, F2: -1}
                                        v12 = (fi96 | v12)
                                        v3 = (v2 | (v3 >> (((v2 / ((739) | 1))) & 63)))
                                        _ = v116
                                    }
                                    var v117 uint8 = 42
                                    _ = fi97
                                    _ = v117
                                }
                            }
                        }
                        _ = fi96
                    }
                    v90 = append(v90, v12)
                }
            default:
                if false {
                    v118 := GenS3{F0: 1, F1: 635, F2: v12}
                    if (((((float64(float32(-2.25)) - 1e10) + float64(-2.25)) - ((1.5 * -2.25) + float64(0.0))) + -2.25) > 3.14159) {
                        if ((false && false) && true) {
                            switch (v12 >> ((409) & 63)) {
                            case -305:
                                v2 = v2
                            case 9223372036854775807:
                                if false {
                                    for fi119 := 0; fi119 < 12; fi119++ {
                                        v90 = append(v90, (fi119 % (((v12 << (((^(v3))) & 63))) | 1)))
                                        var v120 uint64 = 4294967295
                                        v121 := GenS14{F0: 197, F1: 390, F2: "hello"}
                                        _ = fi119
                                        _ = v120
                                        _ = v121
                                    }
                                    if false {
                                        for fi122 := 0; fi122 < 5; fi122++ {
                                            var v123 uint32 = 42
                                            _ = fi122
                                            _ = v123
                                        }
                                        var v124 uint16 = 0
                                        var v125 int32 = v1
                                        _ = v124
                                        _ = v125
                                    } else {
                                        v5 -= v4
                                        v5 += 0
                                    }
                                }
                            case 4294967295:
                                var v126 uint16 = 945
                                _ = v126
                            }
                            if (1 != 32767) {
                                var v127 uint32 = 0
                                if len(v90) >= 2 {
                                    v90 = v90[1:len(v90)]
                                }
                                _ = v127
                            } else {
                                _ = len(v90)
                                for fi128 := 0; fi128 < 15; fi128++ {
                                    for fi129 := 0; fi129 < 8; fi129++ {
                                        v130 := GenS12{F0: 782, F1: true, F2: 2147483647}
                                        _ = fi129
                                        _ = v130
                                    }
                                    v3 = v3
                                    _ = fi128
                                }
                            }
                        }
                    }
                    _ = v118
                }
            }
            _ = v90
        } else {
            v131 := []int{}
            _ = v131
        }
    }
    for fi132 := 0; fi132 < 12; fi132++ {
        switch (fi132 << ((65535) & 63)) {
        case -2147483648:
            var v133 int8 = v4
            _ = v133
        case 4294967295:
            v134 := GenS6{F0: 42, F1: ("a" + "")}
            _ = v134
        case 2147483647:
            var v135 int16 = 1000
            _ = v135
        }
        _ = fi132
    }
    var v136 string = ("\n\t" + "a")
    v2 = (v2 - (^(v3)))
    v137 := []int{}
    fmt.Println(v136)
    switch genFn1("a", v4) {
    case 2147483647:
        v138 := GenS15{F0: 42, F1: 128, F2: -219}
        _ = v138
    case 9223372036854775807:
        switch v136 {
        case "a":
            v136 = "\n\t"
        case "hello":
            fmt.Println(v1, v136, v12)
        case "foo bar":
            v139 := GenS16{F0: 0, F1: (v2 | v3), F2: 65535}
            _ = v139
        }
    case 1:
        v1 -= v1
    }
    var v140 int16 = 382
    var v141 float64 = 0.1
    switch (^(v12)) {
    case 32767:
        v136 = (v136 + v136)
    }
    fmt.Println(v4, v2)
    var v142 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v136:", v136)
    fmt.Println("slicelen v137:", len(v137), v137)
    fmt.Println("v140:", v140)
    fmt.Println("v141:", v141)
    fmt.Println("v142:", v142)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v136
    _ = v137
    _ = v140
    _ = v141
    _ = v142
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
