package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 bool; F1 string; F2 int }
type GenS3 struct { F0 uint32; F1 int64; F2 uint }
type GenS4 struct { F0 uint; F1 int8; F2 int16 }
type GenS5 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS6 struct { F0 int; F1 int16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 string; F1 uint8 }
type GenS9 struct { F0 string; F1 uint8 }
type GenS10 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS11 struct { F0 int64; F1 string }
type GenS12 struct { F0 uint32; F1 uint32 }
type GenS13 struct { F0 int16; F1 uint64; F2 int16 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 uint32; F1 int8; F2 uint }
func genFn0() uint {
    switch -246 {
    case -9223372036854775808:
        for fi4 := 0; fi4 < 17; fi4++ {
            var v5 uint8 = 42
            var v6 uint16 = 1
            v7 := []int{fi4, fi4, (fi4 & fi4), (fi4 / ((159) | 1))}
            _ = fi4
            _ = v5
            _ = v6
            _ = v7
        }
    default:
        v8 := GenS1{F0: 2, F1: 400}
        _ = v8
    }
    v9 := []int{1000}
    v10 := GenS2{F0: (!false), F1: ("" + "a"), F2: 32767}
    v9 = append(v9, 32768)
    switch -179 {
    case 32767:
        v11 := GenS3{F0: 65535, F1: -1, F2: 65535}
        _ = v11
    default:
        var v12 uint = 833
        _ = v12
    }
    _ = v10
    _ = v9
    return 1
}
func genFn1(p0 int16, p1 uint64, p2 int8) int16 {
    p2 = (p2 + (-(p2)))
    v13 := []int{-1, -2147483648}
    p0 = (p0 * (p0 << ((4294967295) & 63)))
    if (2147483648 > 2147483647) {
        switch 42 {
        case 2147483647:
            v14 := GenS3{F0: 2147483648, F1: 1, F2: 2}
            _ = v14
        }
        v15 := GenS3{F0: 2, F1: int64(p0), F2: 2}
        fmt.Println(p2, p2, p0)
        _ = v15
    } else {
        for fi16 := 0; fi16 < 13; fi16++ {
            v17 := GenS4{F0: uint(p1), F1: (p2 ^ p2), F2: -951}
            _ = fi16
            _ = v17
        }
        var v18 uint16 = 486
        _ = v18
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 uint64 = (^(p1))
        _ = fi19
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return p0
}

func main() {
    var v1 uint16 = 959
    var v2 int32 = 390
    v3 := []int{0, 0}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if ((float64(0.1) / 3.14159) <= 3.14159) {
        var v21 int16 = 1
        _ = v21
    }
    switch 32767 {
    case 4294967295:
        fmt.Println(v2, v1)
    case 128:
        if (2147483647 < 0) {
            var v22 uint64 = 1
            _ = v22
        } else {
            if ("foo bar" != "\n\t") {
                if true {
                    if (1e10 != (1e10 / 1.5)) {
                        var v23 string = ("hello" + "hello")
                        _ = v23
                    } else {
                        var v24 float32 = (0.1 * (((float32(0.0) * (((float32(1e10) * float32(-2.25)) - float32(-2.25)) / 1.0)) * (1e10 * 1.5)) * 0.0))
                        var v25 float64 = (-2.25 - (-2.25 * 1.5))
                        v24 += (v24 - v24)
                        _ = v24
                        _ = v25
                    }
                    v26 := GenS4{F0: 18446744073709551615, F1: -1, F2: 42}
                    _ = v26
                }
            }
            v1 += (v1 - 2)
            v1 = (v1 << ((88) & 63))
        }
    case -1:
        for fi27 := 0; fi27 < 13; fi27++ {
            v2 += (v2 & -2147483648)
            switch fi27 {
            case 0:
                switch (("\n\t" + "\n\t") + (("" + "\n\t") + "hello")) {
                case "":
                    switch fi27 {
                    case 4294967295:
                        var v28 bool = (float64((float32((0.0 * ((0.0 * float64(0.0)) + (-2.25 + ((0.0 - 1e10) / (0.1 + float64(3.14159))))))) * 1.5)) >= ((float64(float32(1e10)) + 0.0) + float64(-2.25)))
                        _ = v28
                    case -2147483648:
                        var v29 int32 = (v2 ^ v2)
                        _ = v29
                    case -1:
                        switch "foo bar" {
                        case "":
                            if (128 < 42) {
                                if false {
                                    v2 += (v2 + (-(v2)))
                                    for fi30 := 0; fi30 < 1; fi30++ {
                                        if (0 < 189) {
                                            switch "" {
                                            case "a":
                                                v31 := GenS4{F0: 42, F1: 122, F2: 0}
                                                _ = v31
                                            case "\n\t":
                                                v32 := GenS4{F0: genFn0(), F1: 42, F2: 1}
                                                _ = v32
                                            case "hello":
                                                if true {
                                                    var v33 uint8 = 68
                                                    fmt.Println(fi30)
                                                    fmt.Println(v33, v2)
                                                    v33 = v33
                                                    _ = v33
                                                }
                                            default:
                                                fmt.Println(v2, fi27)
                                            }
                                            fmt.Println(v2, v2)
                                        }
                                        var v34 int16 = -805
                                        v34 = 854
                                        var v35 int = 32767
                                        _ = fi30
                                        _ = v34
                                        _ = v35
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                } else {
                                    var v36 int = 1
                                    for fi37 := 0; fi37 < 5; fi37++ {
                                        if len(v3) > 0 {
                                            v3[0] = (fi27 << ((1) & 63))
                                        }
                                        var v38 uint64 = 42
                                        fmt.Println(fi27)
                                        _ = fi37
                                        _ = v38
                                    }
                                    for fi39 := 0; fi39 < 3; fi39++ {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        var v40 float32 = 1e10
                                        _ = fi39
                                        _ = v40
                                    }
                                    _ = v36
                                }
                            } else {
                                if (0 <= (fi27 ^ (fi27 / ((fi27) | 1)))) {
                                    if false {
                                        if false {
                                            for fi41 := 0; fi41 < 11; fi41++ {
                                                if false {
                                                    _ = len(v3)
                                                    switch 931 {
                                                    case -389:
                                                        _ = len(v3)
                                                    case -1:
                                                        var v42 uint8 = 98
                                                        _ = v42
                                                    case -2147483648:
                                                        v43 := GenS4{F0: 2, F1: 1, F2: -32768}
                                                        if len(v3) > 0 {
                                                            v3[0] = 9223372036854775807
                                                        }
                                                        _ = v43
                                                    }
                                                    var v44 float64 = (0.0 + 1e10)
                                                    _ = v44
                                                }
                                                v2 = (-(v2))
                                                v2 = v2
                                                _ = fi41
                                            }
                                            v45 := GenS1{F0: 2, F1: (v1 / (((v1 % (((^(v1))) | 1))) | 1))}
                                            v1 += (^(v1))
                                            _ = v45
                                        }
                                        if false {
                                            switch "hello" {
                                            case "hello":
                                                if (255 < uint8(v1)) {
                                                    v46 := GenS2{F0: (v1 < (^(v1))), F1: "\n\t", F2: (fi27 * (fi27 / ((-725) | 1)))}
                                                    if len(v3) > 0 {
                                                        v3[0] = (^(fi27))
                                                    }
                                                    v1 = (v1 | 0)
                                                    _ = v46
                                                }
                                            case "":
                                                fmt.Println(v1, v1)
                                            default:
                                                switch "" {
                                                case "hello":
                                                    fmt.Println(v2, v1, v2)
                                                case "":
                                                    var v47 uint64 = 886
                                                    _ = v47
                                                default:
                                                    if ((65535 > 65535) || (18446744073709551615 == 18446744073709551615)) {
                                                        switch 4294967295 {
                                                        case 492:
                                                            var v48 float32 = 0.0
                                                            _ = v48
                                                        case 127:
                                                            if false {
                                                                if (729 == 703) {
                                                                    v49 := GenS4{F0: 1, F1: -1, F2: 32767}
                                                                    v3 = append(v3, (fi27 * fi27))
                                                                    _ = v49
                                                                } else {
                                                                    switch (fi27 + 32768) {
                                                                    case 128:
                                                                        fmt.Println(fi27, v2)
                                                                    case 32767:
                                                                        v1 = v1
                                                                    case -2147483648:
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                    default:
                                                                        var v50 string = "hello"
                                                                        _ = v50
                                                                    }
                                                                    fmt.Println(v2)
                                                                }
                                                            } else {
                                                                fmt.Println(v2)
                                                                if ((-2.25 <= float64(1e10)) && (false || true)) {
                                                                    fmt.Println(v2, v1, v2)
                                                                    _ = len(v3)
                                                                } else {
                                                                    v2 -= -2147483648
                                                                    var v51 int16 = genFn1(-586, 0, 42)
                                                                    if ((!(562 != (v2 / ((559) | 1)))) && (42 == (v2 - -1))) {
                                                                        v2 = (v2 % (((v2 / ((v2) | 1))) | 1))
                                                                        if (!true) {
                                                                            v2 = -2147483648
                                                                            switch fi27 {
                                                                            case 32768:
                                                                                switch ("\n\t" + ("" + ("hello" + "foo bar"))) {
                                                                                case "foo bar":
                                                                                    fmt.Println(v1, v2)
                                                                                }
                                                                            case 2147483647:
                                                                                fmt.Println(v1, v51, v1)
                                                                            case 705:
                                                                                if ((-253 >= 9223372036854775807) || true) {
                                                                                    for fi52 := 0; fi52 < 14; fi52++ {
                                                                                        if len(v3) >= 2 {
                                                                                            v3 = v3[1:len(v3)]
                                                                                        }
                                                                                        _ = fi52
                                                                                    }
                                                                                }
                                                                            default:
                                                                                var v53 uint16 = (v1 >> ((437) & 63))
                                                                                _ = v53
                                                                            }
                                                                        } else {
                                                                            if (1 < -1) {
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = 720
                                                                                }
                                                                            } else {
                                                                                for fi54 := 0; fi54 < 5; fi54++ {
                                                                                    v51 = (v51 - 32767)
                                                                                    _ = fi54
                                                                                }
                                                                            }
                                                                            var v55 float64 = (float64(1.5) / 3.14159)
                                                                            v56 := GenS4{F0: genFn0(), F1: -75, F2: 1}
                                                                            _ = v55
                                                                            _ = v56
                                                                        }
                                                                        _ = len(v3)
                                                                    }
                                                                    _ = v51
                                                                }
                                                                v2 = v2
                                                            }
                                                        default:
                                                            v3 = append(v3, int(v2))
                                                        }
                                                        for fi57 := 0; fi57 < 10; fi57++ {
                                                            var v58 int16 = 32767
                                                            v59 := GenS4{F0: 4294967295, F1: 105, F2: 1000}
                                                            _ = fi57
                                                            _ = v58
                                                            _ = v59
                                                        }
                                                    } else {
                                                        v60 := GenS4{F0: 17, F1: -128, F2: -32768}
                                                        _ = v60
                                                    }
                                                }
                                            }
                                            var v61 float32 = 0.1
                                            var v62 float32 = (v61 + 1e10)
                                            _ = v61
                                            _ = v62
                                        } else {
                                            if false {
                                                if (!false) {
                                                    if false {
                                                        var v63 int8 = 1
                                                        switch "\n\t" {
                                                        case "":
                                                            fmt.Println(v63, v2, v2)
                                                        case "\n\t":
                                                            v64 := GenS4{F0: 18446744073709551615, F1: (v63 & 0), F2: 42}
                                                            _ = v64
                                                        }
                                                        v65 := GenS5{F0: 2147483647, F1: 1, F2: v63}
                                                        _ = v63
                                                        _ = v65
                                                    }
                                                    fmt.Println(fi27, fi27)
                                                    switch fi27 {
                                                    case 9223372036854775807:
                                                        v1 -= v1
                                                    }
                                                } else {
                                                    v1 -= (v1 << ((2) & 63))
                                                    v2 = (v2 << ((507) & 63))
                                                    v2 = (v2 >> ((genFn0()) & 63))
                                                }
                                            } else {
                                                for fi66 := 0; fi66 < 2; fi66++ {
                                                    fmt.Println(v1, v2)
                                                    if (true && false) {
                                                        v3 = append(v3, fi27)
                                                        v1 += v1
                                                    }
                                                    _ = fi66
                                                }
                                            }
                                            for fi67 := 0; fi67 < 9; fi67++ {
                                                v1 += (^(v1))
                                                v2 = v2
                                                v2 = -317
                                                _ = fi67
                                            }
                                        }
                                    }
                                    var v68 uint32 = 2
                                    _ = v68
                                } else {
                                    switch "a" {
                                    case "foo bar":
                                        var v69 bool = (!false)
                                        _ = v69
                                    case "a":
                                        _ = len(v3)
                                    case "":
                                        if false {
                                            if (!((4294967295 >= genFn0()) && false)) {
                                                var v70 int8 = 127
                                                v70 = v70
                                                v71 := GenS6{F0: fi27, F1: -653}
                                                _ = v70
                                                _ = v71
                                            }
                                            v1 += 480
                                        }
                                    default:
                                        for fi72 := 0; fi72 < 15; fi72++ {
                                            v2 += 1
                                            _ = len(v3)
                                            var v73 int8 = 1
                                            _ = fi72
                                            _ = v73
                                        }
                                    }
                                }
                            }
                        case "hello":
                            v1 += 65535
                        }
                    default:
                        switch fi27 {
                        case 32767:
                            for fi74 := 0; fi74 < 9; fi74++ {
                                v75 := GenS7{F0: ""}
                                if (42 >= -128) {
                                    var v76 string = ""
                                    if false {
                                        v77 := GenS8{F0: "hello", F1: 127}
                                        var v78 uint16 = v1
                                        _ = v77
                                        _ = v78
                                    }
                                    v79 := GenS6{F0: fi27, F1: 407}
                                    _ = v76
                                    _ = v79
                                }
                                _ = fi74
                                _ = v75
                            }
                        }
                    }
                case "a":
                    _ = len(v3)
                }
            case 4294967295:
                var v80 int = (^(fi27))
                _ = v80
            }
            v1 += v1
            _ = fi27
        }
    }
    switch "foo bar" {
    case "a":
        var v81 uint32 = 65535
        _ = v81
    default:
        switch 1 {
        case -1:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 128:
            v82 := GenS9{F0: "hello", F1: 0}
            _ = v82
        default:
            switch 127 {
            case 0:
                for fi83 := 0; fi83 < 19; fi83++ {
                    for fi84 := 0; fi84 < 14; fi84++ {
                        v2 = 1
                        var v85 string = ((("\n\t" + "foo bar") + "\n\t") + ("hello" + ""))
                        if true {
                            if (0.0 < (float64(((1e10 + 1.5) - 3.14159)) - float64(1e10))) {
                                v1 = 32767
                                v2 = (^(v2))
                            } else {
                                if len(v3) > 0 {
                                    v3[0] = fi83
                                }
                                if false {
                                    if true {
                                        if (127 < 42) {
                                            var v86 uint64 = 845
                                            _ = v86
                                        } else {
                                            v87 := GenS6{F0: int(v2), F1: -405}
                                            switch fi83 {
                                            case 9223372036854775807:
                                                v85 = v85
                                            case -9223372036854775808:
                                                _ = len(v3)
                                            case 361:
                                                v2 += v2
                                            }
                                            _ = v87
                                        }
                                    }
                                }
                                var v88 int32 = (v2 & (v2 << ((2) & 63)))
                                _ = v88
                            }
                            var v89 uint32 = 0
                            _ = v89
                        } else {
                            _ = len(v3)
                            _ = len(v3)
                        }
                        _ = fi84
                        _ = v85
                    }
                    var v90 uint = genFn0()
                    if true {
                        for fi91 := 0; fi91 < 17; fi91++ {
                            v3 = append(v3, int(v1))
                            _ = fi91
                        }
                    } else {
                        v92 := GenS7{F0: "\n\t"}
                        if (!(!(920 <= 18446744073709551615))) {
                            fmt.Println(v1, v90, v90)
                        } else {
                            if (!false) {
                                v93 := GenS10{F0: v1, F1: 42, F2: 1000}
                                var v94 float32 = float32(float64(1e10))
                                _ = v93
                                _ = v94
                            }
                            v95 := GenS10{F0: (v1 % ((v1) | 1)), F1: 1, F2: 0}
                            v90 -= genFn0()
                            _ = v95
                        }
                        for fi96 := 0; fi96 < 12; fi96++ {
                            v3 = append(v3, (fi83 / (((-(fi96))) | 1)))
                            _ = fi96
                        }
                        _ = v92
                    }
                    _ = fi83
                    _ = v90
                }
            case 128:
                switch "foo bar" {
                case "foo bar":
                    if true {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    } else {
                        var v97 uint64 = 1
                        v98 := GenS4{F0: 255, F1: 42, F2: -1}
                        _ = v97
                        _ = v98
                    }
                case "a":
                    v1 -= v1
                case "":
                    var v99 uint = 1
                    _ = v99
                }
            case 32768:
                var v100 float32 = ((float32(float64(1e10)) - 1e10) + 0.1)
                _ = v100
            }
        }
    }
    for fi101 := 0; fi101 < 9; fi101++ {
        for fi102 := 0; fi102 < 6; fi102++ {
            v103 := GenS11{F0: 0, F1: "a"}
            if (!("foo bar" >= "\n\t")) {
                v1 -= 255
                var v104 float32 = (((0.0 * (3.14159 / 1.5)) - -2.25) - -2.25)
                if (v1 < v1) {
                    v105 := GenS12{F0: 65535, F1: 4294967295}
                    if (true || (!true)) {
                        var v106 float64 = -2.25
                        _ = len(v3)
                        _ = v106
                    }
                    v107 := GenS12{F0: 42, F1: 4294967295}
                    _ = v105
                    _ = v107
                }
                _ = v104
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi102
            _ = v103
        }
        if len(v3) > 0 {
            v3[0] = (fi101 >> ((2) & 63))
        }
        for fi108 := 0; fi108 < 13; fi108++ {
            v109 := GenS13{F0: genFn1(819, 2, -128), F1: 18446744073709551615, F2: -1}
            switch "a" {
            case "hello":
                v1 = v1
            case "a":
                if (4294967295 > 65535) {
                    var v110 int8 = -105
                    _ = v110
                } else {
                    var v111 int = fi108
                    v112 := GenS14{F0: "\n\t"}
                    _ = v111
                    _ = v112
                }
            }
            _ = fi108
            _ = v109
        }
        _ = fi101
    }
    switch "\n\t" {
    case "a":
        var v113 float32 = 3.14159
        _ = v113
    default:
        var v114 int64 = 1
        _ = v114
    }
    for fi115 := 0; fi115 < 9; fi115++ {
        v116 := GenS15{F0: 42, F1: 115, F2: 929}
        fmt.Println(v2, v1, v2)
        _ = fi115
        _ = v116
    }
    var v117 bool = true
    for fi118 := 0; fi118 < 15; fi118++ {
        v1 = v1
        v3 = append(v3, fi118)
        for fi119 := 0; fi119 < 2; fi119++ {
            v1 -= v1
            v3 = append(v3, -667)
            fmt.Println(fi118)
            _ = fi119
        }
        _ = fi118
    }
    v1 = (v1 % ((v1) | 1))
    if (3.14159 != 0.1) {
        var v120 int32 = 1
        for fi121 := 0; fi121 < 7; fi121++ {
            if (255 != (v1 & v1)) {
                var v122 int32 = (v2 - v120)
                v123 := GenS9{F0: "a", F1: 128}
                _ = v122
                _ = v123
            }
            for fi124 := 0; fi124 < 2; fi124++ {
                v3 = append(v3, fi124)
                var v125 int = fi124
                var v126 uint64 = 519
                _ = fi124
                _ = v125
                _ = v126
            }
            _ = fi121
        }
        _ = v120
    }
    for fi127 := 0; fi127 < 10; fi127++ {
        for fi128 := 0; fi128 < 17; fi128++ {
            var v129 uint8 = 127
            _ = fi128
            _ = v129
        }
        var v130 uint64 = 591
        _ = fi127
        _ = v130
    }
    fmt.Println("v1:", v1)
    fmt.Println("v117:", v117)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v117
    _ = v2
    _ = v3
}
