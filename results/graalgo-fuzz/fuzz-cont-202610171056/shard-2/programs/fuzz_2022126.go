package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int16 }
type GenS2 struct { F0 int32; F1 string; F2 int32 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 uint32 }
type GenS6 struct { F0 uint32; F1 int }
type GenS7 struct { F0 uint64; F1 int32; F2 int }
type GenS8 struct { F0 int; F1 uint16 }
type GenS9 struct { F0 int; F1 uint }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int16; F1 uint; F2 int8 }
type GenS12 struct { F0 uint; F1 int16; F2 uint16 }
type GenS13 struct { F0 uint8; F1 int16; F2 int64 }
type GenS14 struct { F0 int32; F1 int64 }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 int8; F1 bool }
type GenS17 struct { F0 uint16; F1 bool; F2 string }
type GenS18 struct { F0 uint16; F1 int64 }
type GenS19 struct { F0 int64; F1 uint }
type GenS20 struct { F0 int32; F1 int8 }
type GenS21 struct { F0 int16 }
type GenS22 struct { F0 bool; F1 uint32; F2 uint }
type GenS23 struct { F0 int }
type GenS24 struct { F0 uint32; F1 uint64 }
func genFn0(p0 uint8, p1 int64) int16 {
    v5 := GenS1{F0: -1, F1: -956, F2: 21}
    if true {
        v6 := GenS1{F0: 100, F1: 0, F2: 1000}
        fmt.Println(p1, p1)
        _ = v6
    }
    v7 := []int{}
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := GenS1{F0: 1, F1: 0, F2: 330}
        if ((!false) || true) {
            switch "foo bar" {
            case "":
                var v10 bool = (((3.14159 + 0.1) * -2.25) < 1e10)
                _ = v10
            default:
                fmt.Println(p0)
            }
            v7 = append(v7, 2147483647)
        }
        _ = fi8
        _ = v9
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v5
    _ = v7
    return 42
}
func genFn1(p0 uint, p1 int16, p2 uint8) uint {
    switch ("hello" + (("" + "a") + "a")) {
    case "":
        for fi11 := 0; fi11 < 2; fi11++ {
            p1 = genFn0(p2, -403)
            _ = fi11
        }
    }
    if (!(0 >= (p1 % ((p1) | 1)))) {
        if true {
            for fi12 := 0; fi12 < 8; fi12++ {
                if ((float64(3.14159) * float64(3.14159)) <= ((0.0 + 1.5) * (0.1 / 1.0))) {
                    fmt.Println(p2, p0)
                    switch ((((("" + "hello") + "") + "hello") + (("foo bar" + "foo bar") + ("foo bar" + "hello"))) + ("\n\t" + (("\n\t" + "\n\t") + "\n\t"))) {
                    case "foo bar":
                        p1 += p1
                    case "\n\t":
                        if (2147483647 > 2147483647) {
                            if false {
                                v13 := []int{}
                                switch fi12 {
                                case 128:
                                    p1 -= p1
                                case 4294967295:
                                    for fi14 := 0; fi14 < 10; fi14++ {
                                        if false {
                                            var v15 uint16 = 32767
                                            v13 = append(v13, int(v15))
                                            _ = v15
                                        } else {
                                            fmt.Println(p0)
                                            v16 := GenS1{F0: -128, F1: -283, F2: 142}
                                            if len(v13) > 0 {
                                                v13[0] = fi12
                                            }
                                            if true {
                                                if true {
                                                    var v17 float32 = 1.5
                                                    var v18 uint32 = 719
                                                    p0 = (p0 >> ((p0) & 63))
                                                    _ = v17
                                                    _ = v18
                                                }
                                                v19 := GenS1{F0: 100, F1: 0, F2: (p1 ^ (p1 >> (((p0 >> (((p0 ^ p0)) & 63))) & 63)))}
                                                var v20 int64 = 421
                                                _ = v19
                                                _ = v20
                                            } else {
                                                var v21 int32 = -899
                                                _ = v21
                                            }
                                            var v22 int8 = 100
                                            _ = v16
                                            _ = v22
                                        }
                                        p1 += (p1 + (p1 << ((p0) & 63)))
                                        var v23 int16 = 1000
                                        _ = fi14
                                        _ = v23
                                    }
                                case 789:
                                    for fi24 := 0; fi24 < 13; fi24++ {
                                        var v25 int64 = 1
                                        p1 = p1
                                        p2 += (p2 / ((42) | 1))
                                        _ = fi24
                                        _ = v25
                                    }
                                default:
                                    if false {
                                        v26 := GenS1{F0: -128, F1: -2147483648, F2: (^(p1))}
                                        switch 2147483647 {
                                        case -9223372036854775808:
                                            fmt.Println(p2, p1)
                                        case 789:
                                            switch "" {
                                            case "a":
                                                if (-2.25 >= 3.14159) {
                                                    v27 := GenS1{F0: -128, F1: -1, F2: (p1 - p1)}
                                                    _ = v27
                                                }
                                            case "":
                                                if len(v13) >= 2 {
                                                    v13 = v13[1:len(v13)]
                                                }
                                            }
                                        case 42:
                                            _ = len(v13)
                                        }
                                        _ = v26
                                    }
                                }
                                _ = v13
                            }
                        } else {
                            for fi28 := 0; fi28 < 13; fi28++ {
                                if (0 <= 564) {
                                    switch "hello" {
                                    case "\n\t":
                                        p0 = p0
                                    case "":
                                        fmt.Println(p0, p1)
                                    case "a":
                                        var v29 float32 = (0.0 + float32(1.5))
                                        _ = v29
                                    default:
                                        p1 -= p1
                                    }
                                    v30 := GenS2{F0: 245, F1: "\n\t", F2: -2147483648}
                                    _ = v30
                                }
                                v31 := GenS3{F0: (fi28 >> ((2) & 63)), F1: false}
                                v32 := GenS4{F0: 4294967295}
                                p2 = p2
                                var v33 uint32 = 1
                                switch fi12 {
                                case 2147483647:
                                    var v34 int8 = -1
                                    _ = v34
                                case 32767:
                                    fmt.Println(v33)
                                case 1:
                                    fmt.Println(fi12)
                                }
                                _ = fi28
                                _ = v31
                                _ = v32
                                _ = v33
                            }
                        }
                    case "a":
                        v35 := []int{(fi12 + fi12), (fi12 + (fi12 << ((p0) & 63))), -276}
                        _ = v35
                    }
                }
                v36 := []int{808, fi12, fi12}
                for fi37 := 0; fi37 < 11; fi37++ {
                    switch fi12 {
                    case 0:
                        p2 += (p2 ^ p2)
                    }
                    var v38 uint16 = 996
                    _ = fi37
                    _ = v38
                }
                _ = fi12
                _ = v36
            }
            for fi39 := 0; fi39 < 5; fi39++ {
                for fi40 := 0; fi40 < 10; fi40++ {
                    if false {
                        p2 = 182
                        p1 += p1
                        var v41 int32 = 888
                        _ = v41
                    }
                    _ = fi40
                }
                _ = fi39
            }
        }
    } else {
        var v42 uint64 = 0
        if true {
            v43 := []int{127, 32768, 58}
            if (-104 < -9223372036854775808) {
                if false {
                    fmt.Println(v42)
                    fmt.Println(p0, v42)
                    var v44 int64 = -57
                    _ = v44
                }
            }
            if len(v43) >= 2 {
                v43 = v43[1:len(v43)]
            }
            _ = v43
        } else {
            var v45 int64 = 1
            for fi46 := 0; fi46 < 3; fi46++ {
                if false {
                    var v47 uint32 = 4294967295
                    for fi48 := 0; fi48 < 8; fi48++ {
                        v49 := GenS2{F0: 0, F1: "hello", F2: 2147483647}
                        fmt.Println(v42, v45)
                        v47 = 2
                        _ = fi48
                        _ = v49
                    }
                    _ = v47
                } else {
                    v50 := GenS3{F0: 592, F1: (!(!(v45 >= 42)))}
                    _ = v50
                }
                v51 := GenS5{F0: uint64(p2), F1: 42}
                _ = fi46
                _ = v51
            }
            v52 := GenS2{F0: -300, F1: ("hello" + ""), F2: 913}
            _ = v45
            _ = v52
        }
        var v53 bool = true
        _ = v42
        _ = v53
    }
    v54 := GenS1{F0: 0, F1: -2147483648, F2: (p1 + (^(p1)))}
    switch 2147483647 {
    case -1:
        v55 := []int{32768}
        _ = v55
    case 42:
        p2 -= 128
    case 172:
        v56 := GenS6{F0: 233, F1: 2147483647}
        _ = v56
    default:
        var v57 uint = 147
        _ = v57
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v54
    return (^(p0))
}

func main() {
    var v1 int64 = 1
    var v2 int64 = 0
    var v3 float32 = (0.0 / (0.1 / ((0.1 + (1e10 * (-2.25 * -2.25))) + -2.25)))
    var v4 bool = (4294967295 > 2147483648)
    v58 := []int{-1, 142}
    if false {
        _ = len(v58)
    } else {
        switch (("" + "") + "a") {
        case "":
            if true {
                for fi59 := 0; fi59 < 4; fi59++ {
                    fmt.Println(v3, v3, fi59)
                    fmt.Println(fi59, v1, v3)
                    var v60 float32 = 0.1
                    _ = fi59
                    _ = v60
                }
            } else {
                var v61 uint16 = 425
                _ = v61
            }
        }
        v1 -= v1
    }
    v3 = float32(3.14159)
    var v62 int8 = -9
    for fi63 := 0; fi63 < 12; fi63++ {
        switch ("hello" + "\n\t") {
        case "hello":
            if true {
                var v64 int32 = 2147483647
                v58 = append(v58, (fi63 + (fi63 + fi63)))
                switch ("foo bar" + "hello") {
                case "hello":
                    if (v64 <= (v64 ^ (v64 ^ v64))) {
                        v64 += v64
                        var v65 uint32 = 0
                        v66 := GenS2{F0: (v64 << ((genFn1(311, -1, 255)) & 63)), F1: ("foo bar" + ("hello" + "")), F2: v64}
                        _ = v65
                        _ = v66
                    }
                case "foo bar":
                    if true {
                        for fi67 := 0; fi67 < 3; fi67++ {
                            v4 = (!(1 == 0))
                            v4 = v4
                            switch fi67 {
                            case 9223372036854775807:
                                v68 := GenS1{F0: (v62 + (v62 | v62)), F1: (v64 % (((v64 + (v64 - (v64 % (((v64 ^ v64)) | 1))))) | 1)), F2: 1000}
                                _ = v68
                            case 876:
                                fmt.Println(v2)
                            case 257:
                                v1 = 42
                            }
                            var v69 int16 = genFn0(255, v2)
                            _ = fi67
                            _ = v69
                        }
                    } else {
                        if (v4 || v4) {
                            for fi70 := 0; fi70 < 3; fi70++ {
                                fmt.Println(v62, v4)
                                v62 += v62
                                if v4 {
                                    fmt.Println(v64)
                                    v71 := GenS7{F0: 4294967295, F1: (v64 >> ((854) & 63)), F2: fi63}
                                    switch fi63 {
                                    case -1:
                                        if len(v58) >= 2 {
                                            v58 = v58[1:len(v58)]
                                        }
                                    case 127:
                                        v58 = append(v58, 9223372036854775807)
                                    }
                                    _ = v71
                                } else {
                                    if (((v3 > v3) || (!((-(fi63)) == fi70))) || (v4 && v4)) {
                                        var v72 uint8 = 127
                                        if len(v58) >= 2 {
                                            v58 = v58[1:len(v58)]
                                        }
                                        v4 = false
                                        if (!(((0.0 + v3) - (1.5 - v3)) < v3)) {
                                            v73 := GenS8{F0: 9223372036854775807, F1: 65535}
                                            var v74 uint = genFn1(genFn1(26, 1, (v72 >> ((255) & 63))), 1000, (v72 << ((630) & 63)))
                                            _ = v73
                                            _ = v74
                                        }
                                        var v75 int32 = 987
                                        _ = v72
                                        _ = v75
                                    }
                                }
                                fmt.Println(fi70, v2)
                                fmt.Println(v2, v64)
                                _ = fi70
                            }
                            if (42 >= 1) {
                                if (fi63 < -1) {
                                    var v76 int = (fi63 & fi63)
                                    v76 += (fi63 >> ((65535) & 63))
                                    fmt.Println(v3)
                                    _ = v76
                                }
                                if len(v58) > 0 {
                                    v58[0] = fi63
                                }
                                if ("hello" > ((((("foo bar" + "hello") + "a") + "a") + "a") + (("foo bar" + "a") + "a"))) {
                                    v62 += (v62 % (((v62 % (((^(v62))) | 1))) | 1))
                                } else {
                                    v4 = false
                                    v1 += -9223372036854775808
                                    v4 = v4
                                }
                            } else {
                                var v77 float32 = ((1.5 + (v3 / float32(1e10))) + -2.25)
                                v4 = false
                                _ = v77
                            }
                            if v4 {
                                v3 = 0.1
                                for fi78 := 0; fi78 < 20; fi78++ {
                                    fmt.Println(v62, v64, fi78)
                                    _ = fi78
                                }
                                if (!((919 != 0) || v4)) {
                                    fmt.Println(v64, v64, v62)
                                    v79 := GenS9{F0: fi63, F1: 0}
                                    _ = v79
                                }
                            }
                        } else {
                            if v4 {
                                if v4 {
                                    v3 = float32(float64((0.0 - -2.25)))
                                    for fi80 := 0; fi80 < 7; fi80++ {
                                        v81 := GenS10{F0: 0}
                                        _ = fi80
                                        _ = v81
                                    }
                                    fmt.Println(v3, v3)
                                } else {
                                    if false {
                                        for fi82 := 0; fi82 < 3; fi82++ {
                                            switch ((("hello" + "a") + ((("a" + ("\n\t" + "hello")) + (("foo bar" + "\n\t") + "a")) + ("a" + "foo bar"))) + "") {
                                            case "foo bar":
                                                if len(v58) > 0 {
                                                    v58[0] = fi63
                                                }
                                                var v83 int = fi82
                                                var v84 uint8 = 2
                                                fmt.Println(v62)
                                                v85 := GenS8{F0: fi63, F1: 0}
                                                if v4 {
                                                    v83 = -162
                                                } else {
                                                    if v4 {
                                                        if v4 {
                                                            v86 := GenS1{F0: 1, F1: -1, F2: 1000}
                                                            _ = v86
                                                        } else {
                                                            var v87 bool = (!v4)
                                                            v2 = v2
                                                            fmt.Println(fi63, v2, v3)
                                                            v88 := GenS7{F0: 0, F1: v64, F2: (-(v83))}
                                                            var v89 float32 = 0.1
                                                            _ = v87
                                                            _ = v88
                                                            _ = v89
                                                        }
                                                        var v90 string = "a"
                                                        _ = v90
                                                    }
                                                    v4 = false
                                                    if len(v58) >= 2 {
                                                        v58 = v58[1:len(v58)]
                                                    }
                                                }
                                                if (true || v4) {
                                                    if len(v58) > 0 {
                                                        v58[0] = int(v1)
                                                    }
                                                    if ((34 > 42) && true) {
                                                        v58 = append(v58, fi82)
                                                    }
                                                }
                                                _ = v83
                                                _ = v84
                                                _ = v85
                                            default:
                                                var v91 uint64 = 0
                                                _ = v91
                                            }
                                            _ = fi82
                                        }
                                        v64 = v64
                                    }
                                    for fi92 := 0; fi92 < 8; fi92++ {
                                        v4 = ("hello" > ("\n\t" + "\n\t"))
                                        var v93 uint16 = 42
                                        _ = fi92
                                        _ = v93
                                    }
                                    for fi94 := 0; fi94 < 8; fi94++ {
                                        v4 = false
                                        switch 9223372036854775807 {
                                        case 603:
                                            switch ("a" + "hello") {
                                            case "hello":
                                                var v95 int8 = v62
                                                _ = v95
                                            case "a":
                                                switch "" {
                                                case "":
                                                    v4 = ((!((32767 > 42) && true)) || v4)
                                                }
                                            case "foo bar":
                                                if (1 <= 128) {
                                                    var v96 uint = 539
                                                    _ = v96
                                                } else {
                                                    v1 += (v1 & (v2 / ((v1) | 1)))
                                                    if len(v58) >= 2 {
                                                        v58 = v58[1:len(v58)]
                                                    }
                                                    switch "" {
                                                    case "foo bar":
                                                        if v4 {
                                                            v64 = v64
                                                            var v97 uint16 = 32767
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                if (-1 != int64(fi94)) {
                                                                    v98 := GenS11{F0: 42, F1: 42, F2: int8(v64)}
                                                                    v99 := GenS5{F0: 18446744073709551615, F1: 65535}
                                                                    v64 -= 42
                                                                    _ = v98
                                                                    _ = v99
                                                                } else {
                                                                    v97 = v97
                                                                    var v100 int16 = int16(v64)
                                                                    _ = v100
                                                                }
                                                            case "hello":
                                                                var v101 uint32 = 2
                                                                _ = v101
                                                            case "":
                                                                v2 = v1
                                                            default:
                                                                if v4 {
                                                                    switch fi94 {
                                                                    case 32768:
                                                                        var v102 uint8 = 0
                                                                        _ = v102
                                                                    case -468:
                                                                        fmt.Println(v64, v97, fi63)
                                                                    case 489:
                                                                        v97 += (v97 << ((616) & 63))
                                                                        var v103 uint16 = v97
                                                                        _ = v103
                                                                    default:
                                                                        switch 32767 {
                                                                        case 128:
                                                                            var v104 int8 = (v62 << ((442) & 63))
                                                                            _ = v104
                                                                        case 1:
                                                                            _ = len(v58)
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            _ = v97
                                                        } else {
                                                            switch fi63 {
                                                            case -218:
                                                                if ((v4 || (314 == genFn1(355, 32767, 74))) || false) {
                                                                    v105 := GenS12{F0: 1, F1: 42, F2: 226}
                                                                    _ = v105
                                                                } else {
                                                                    var v106 int8 = (v62 & 0)
                                                                    switch 42 {
                                                                    case -1:
                                                                        var v107 uint32 = 1
                                                                        _ = v107
                                                                    case 4294967295:
                                                                        fmt.Println(v4, v106)
                                                                    default:
                                                                        if len(v58) > 0 {
                                                                            v58[0] = fi63
                                                                        }
                                                                    }
                                                                    _ = v106
                                                                }
                                                            case 32768:
                                                                switch fi94 {
                                                                case 520:
                                                                    v3 = float32((0.1 / -2.25))
                                                                case 689:
                                                                    v108 := GenS13{F0: 127, F1: -1, F2: v1}
                                                                    _ = v108
                                                                case 4294967295:
                                                                    if v4 {
                                                                        var v109 uint = genFn1(42, 42, 1)
                                                                        _ = v109
                                                                    } else {
                                                                        switch fi94 {
                                                                        case -9223372036854775808:
                                                                            if len(v58) >= 2 {
                                                                                v58 = v58[1:len(v58)]
                                                                            }
                                                                        case -60:
                                                                            v110 := GenS14{F0: (v64 ^ 42), F1: (-(v1))}
                                                                            _ = v110
                                                                        case -1:
                                                                            var v111 int64 = -3
                                                                            _ = v111
                                                                        default:
                                                                            if len(v58) > 0 {
                                                                                v58[0] = (^(fi63))
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            case -1:
                                                                fmt.Println(v1, fi63, v3)
                                                            }
                                                            _ = len(v58)
                                                            if len(v58) >= 2 {
                                                                v58 = v58[1:len(v58)]
                                                            }
                                                        }
                                                    case "hello":
                                                        if ((v64 | v64) == (v64 ^ (v64 % ((v64) | 1)))) {
                                                            v112 := GenS8{F0: fi94, F1: 32767}
                                                            if true {
                                                                _ = len(v58)
                                                            }
                                                            _ = v112
                                                        } else {
                                                            var v113 uint = genFn1(18446744073709551615, 874, 255)
                                                            _ = v113
                                                        }
                                                    case "\n\t":
                                                        v2 = v2
                                                    default:
                                                        fmt.Println(fi63, fi94)
                                                    }
                                                }
                                            }
                                        case 9223372036854775807:
                                            fmt.Println(v3, v1, fi94)
                                        }
                                        var v114 int8 = int8(v1)
                                        v115 := GenS15{F0: 642}
                                        _ = fi94
                                        _ = v114
                                        _ = v115
                                    }
                                }
                                for fi116 := 0; fi116 < 9; fi116++ {
                                    var v117 uint16 = 32767
                                    switch "" {
                                    case "a":
                                        if v4 {
                                            v118 := GenS2{F0: -1, F1: "a", F2: (v64 % ((1) | 1))}
                                            var v119 int8 = 127
                                            var v120 int8 = (v62 * (v119 & (^(v62))))
                                            _ = v118
                                            _ = v119
                                            _ = v120
                                        }
                                    default:
                                        var v121 uint = 4294967295
                                        _ = v121
                                    }
                                    var v122 uint8 = 128
                                    if v4 {
                                        var v123 int = (fi63 & (fi116 - (fi63 - (fi63 + (fi116 + (-(fi116)))))))
                                        var v124 int = (fi63 / ((fi116) | 1))
                                        v122 = v122
                                        var v125 bool = (!v4)
                                        if false {
                                            _ = len(v58)
                                            v126 := GenS10{F0: 815}
                                            _ = v126
                                        } else {
                                            v64 = (v64 ^ v64)
                                            var v127 uint16 = v117
                                            _ = v127
                                        }
                                        fmt.Println(v62, v4)
                                        _ = v123
                                        _ = v124
                                        _ = v125
                                    }
                                    _ = fi116
                                    _ = v117
                                    _ = v122
                                }
                            } else {
                                var v128 int64 = int64(fi63)
                                _ = v128
                            }
                        }
                        switch fi63 {
                        case 937:
                            v3 = ((0.1 - (v3 * 1.5)) + (((v3 / 1.5) - 0.0) * ((0.0 - v3) * v3)))
                        default:
                            switch (("\n\t" + "a") + "\n\t") {
                            case "":
                                if false {
                                    v129 := GenS16{F0: -36, F1: v4}
                                    v130 := GenS1{F0: (v62 | -1), F1: (v64 ^ 654), F2: -1}
                                    _ = v129
                                    _ = v130
                                } else {
                                    switch fi63 {
                                    case -9223372036854775808:
                                        for fi131 := 0; fi131 < 16; fi131++ {
                                            switch ("foo bar" + "a") {
                                            case "foo bar":
                                                v3 = float32(1e10)
                                            case "\n\t":
                                                var v132 uint16 = 42
                                                _ = v132
                                            }
                                            v3 = float32(float64(float32(1.5)))
                                            _ = fi131
                                        }
                                    }
                                }
                            }
                        }
                    }
                case "a":
                    switch ("hello" + (("a" + "\n\t") + "\n\t")) {
                    case "foo bar":
                        v133 := GenS10{F0: 4294967295}
                        _ = v133
                    case "":
                        for fi134 := 0; fi134 < 13; fi134++ {
                            if v4 {
                                var v135 uint = 65535
                                v4 = ((1e10 < -2.25) || v4)
                                _ = v135
                            } else {
                                if false {
                                    v136 := GenS17{F0: 920, F1: (1.5 < 1e10), F2: (("foo bar" + ("\n\t" + (("" + "") + ""))) + (("\n\t" + ("hello" + "\n\t")) + ("" + ("\n\t" + ("a" + "hello")))))}
                                    if (127 < 0) {
                                        fmt.Println(fi63, v4, v64)
                                    } else {
                                        if true {
                                            if v4 {
                                                if len(v58) >= 2 {
                                                    v58 = v58[1:len(v58)]
                                                }
                                                v64 -= 1
                                                v137 := GenS18{F0: 2, F1: (v1 % ((v2) | 1))}
                                                _ = v137
                                            } else {
                                                v2 = (v1 % ((0) | 1))
                                                v2 -= int64(fi63)
                                                var v138 int = fi134
                                                _ = v138
                                            }
                                        } else {
                                            v2 -= v1
                                        }
                                    }
                                    var v139 int16 = -1
                                    _ = v136
                                    _ = v139
                                } else {
                                    if v4 {
                                        fmt.Println(v1, fi134, v3)
                                    }
                                    fmt.Println(fi63)
                                    if v4 {
                                        fmt.Println(fi63, v3)
                                        var v140 int8 = v62
                                        var v141 int16 = -111
                                        _ = v140
                                        _ = v141
                                    } else {
                                        v1 = v1
                                        v4 = false
                                    }
                                    if (((-1 <= 0) || ("\n\t" > (("foo bar" + ((("hello" + "foo bar") + ("foo bar" + "\n\t")) + "a")) + (("foo bar" + "hello") + "hello")))) && v4) {
                                        if false {
                                            if len(v58) > 0 {
                                                v58[0] = -2147483648
                                            }
                                            if (42 != 32767) {
                                                var v142 uint64 = 4294967295
                                                _ = v142
                                            }
                                            if ((!true) || false) {
                                                v3 += 3.14159
                                                var v143 uint8 = 1
                                                _ = v143
                                            } else {
                                                if true {
                                                    if v4 {
                                                        var v144 int64 = v1
                                                        v145 := GenS19{F0: v144, F1: 845}
                                                        _ = v144
                                                        _ = v145
                                                    }
                                                    v64 -= (^(v64))
                                                }
                                                if len(v58) >= 2 {
                                                    v58 = v58[1:len(v58)]
                                                }
                                                if len(v58) >= 2 {
                                                    v58 = v58[1:len(v58)]
                                                }
                                            }
                                        } else {
                                            fmt.Println(v1, v1, v64)
                                            switch fi134 {
                                            case -1:
                                                v4 = (!v4)
                                            case 1:
                                                switch "a" {
                                                case "a":
                                                    v146 := GenS20{F0: 2147483647, F1: v62}
                                                    _ = v146
                                                case "":
                                                    var v147 uint = 1
                                                    if v4 {
                                                        if true {
                                                            v148 := GenS21{F0: 976}
                                                            _ = v148
                                                        }
                                                    } else {
                                                        v3 -= (v3 * v3)
                                                        switch fi63 {
                                                        case 42:
                                                            if len(v58) >= 2 {
                                                                v58 = v58[1:len(v58)]
                                                            }
                                                        }
                                                        v149 := GenS22{F0: v4, F1: 0, F2: 255}
                                                        var v150 uint16 = 65535
                                                        fmt.Println(v2, fi134)
                                                        _ = v149
                                                        _ = v150
                                                    }
                                                    switch 32768 {
                                                    case -2147483648:
                                                        v2 -= -1
                                                    case 94:
                                                        var v151 uint8 = 225
                                                        _ = v151
                                                    case 127:
                                                        switch 990 {
                                                        case 1:
                                                            switch (fi134 + -282) {
                                                            case 128:
                                                                _ = len(v58)
                                                            }
                                                        case 9223372036854775807:
                                                            fmt.Println(v3, v3, v147)
                                                        case 423:
                                                            var v152 int64 = int64(v62)
                                                            _ = v152
                                                        }
                                                    }
                                                    var v153 uint = (^(v147))
                                                    _ = v147
                                                    _ = v153
                                                default:
                                                    v58 = append(v58, fi63)
                                                }
                                            case 128:
                                                if (false && ((-(fi134)) >= fi63)) {
                                                    switch int(v62) {
                                                    case 127:
                                                        switch 128 {
                                                        case -1:
                                                            v154 := GenS12{F0: 0, F1: 42, F2: 32767}
                                                            _ = v154
                                                        case 4294967295:
                                                            v58 = append(v58, -1)
                                                        case 9223372036854775807:
                                                            var v155 uint32 = 2147483647
                                                            _ = v155
                                                        default:
                                                            var v156 uint64 = 18446744073709551615
                                                            _ = v156
                                                        }
                                                    case -2147483648:
                                                        switch fi63 {
                                                        case -1:
                                                            var v157 uint64 = 952
                                                            v158 := GenS23{F0: fi134}
                                                            _ = v157
                                                            _ = v158
                                                        case 601:
                                                            if (1 != 206) {
                                                                if (!v4) {
                                                                    switch (("\n\t" + "a") + ("" + "hello")) {
                                                                    case "a":
                                                                        v2 -= v1
                                                                    case "\n\t":
                                                                        var v159 uint16 = 65535
                                                                        _ = v159
                                                                    default:
                                                                        var v160 int = (^(fi63))
                                                                        _ = v160
                                                                    }
                                                                    v4 = true
                                                                    v2 -= -895
                                                                } else {
                                                                    var v161 int32 = -629
                                                                    var v162 float32 = (v3 / v3)
                                                                    if v4 {
                                                                        v58 = append(v58, fi134)
                                                                        if v4 {
                                                                            var v163 uint16 = 2
                                                                            _ = v163
                                                                        } else {
                                                                            var v164 uint = 65535
                                                                            v64 = (v64 / ((v161) | 1))
                                                                            var v165 int64 = 42
                                                                            _ = v164
                                                                            _ = v165
                                                                        }
                                                                    } else {
                                                                        v1 = (v2 + v2)
                                                                        _ = len(v58)
                                                                        fmt.Println(v1)
                                                                    }
                                                                    _ = v161
                                                                    _ = v162
                                                                }
                                                            } else {
                                                                v166 := GenS18{F0: 65535, F1: v2}
                                                                switch "\n\t" {
                                                                case "a":
                                                                    var v167 bool = v4
                                                                    _ = v167
                                                                case "hello":
                                                                    v2 = -1
                                                                case "":
                                                                    v62 = (v62 | (-(v62)))
                                                                }
                                                                v1 = 0
                                                                _ = v166
                                                            }
                                                        }
                                                    default:
                                                        v1 = (v1 | v1)
                                                    }
                                                }
                                            }
                                            v1 = v2
                                        }
                                        v62 = v62
                                        v168 := GenS13{F0: 128, F1: 1, F2: v1}
                                        fmt.Println(v3, v62, v2)
                                        var v169 uint32 = 557
                                        _ = v168
                                        _ = v169
                                    } else {
                                        if len(v58) > 0 {
                                            v58[0] = fi63
                                        }
                                        if v4 {
                                            switch fi63 {
                                            case 32767:
                                                v4 = true
                                            case 32768:
                                                v2 = v2
                                            }
                                            switch fi134 {
                                            case -9223372036854775808:
                                                v62 -= 42
                                            case 42:
                                                var v170 bool = v4
                                                _ = v170
                                            }
                                        }
                                        var v171 int32 = (-(v64))
                                        _ = v171
                                    }
                                }
                                var v172 uint64 = 1
                                v64 = 0
                                fmt.Println(v1)
                                _ = v172
                            }
                            var v173 uint64 = 42
                            v62 = (v62 / ((v62) | 1))
                            switch -305 {
                            case -1:
                                v64 += (v64 >> ((65535) & 63))
                            case 128:
                                v2 -= 42
                                switch (fi134 - (fi134 / (((fi134 & (fi134 % ((int(v62)) | 1)))) | 1))) {
                                case 0:
                                    v174 := GenS14{F0: (-(v64)), F1: 1}
                                    _ = v174
                                default:
                                    if len(v58) >= 2 {
                                        v58 = v58[1:len(v58)]
                                    }
                                }
                                var v175 float32 = -2.25
                                _ = v175
                            default:
                                var v176 uint16 = 292
                                _ = v176
                            }
                            _ = fi134
                            _ = v173
                        }
                    case "hello":
                        var v177 bool = false
                        _ = v177
                    default:
                        fmt.Println(v64, fi63, fi63)
                    }
                default:
                    v2 = 9223372036854775807
                }
                _ = v64
            } else {
                if (2147483647 != -1) {
                    v3 -= (float32((1.5 - float64(0.1))) / v3)
                    var v178 int32 = 659
                    var v179 int = int(v178)
                    _ = v178
                    _ = v179
                }
            }
        case "":
            v62 += v62
        case "foo bar":
            v2 += v1
        default:
            switch "" {
            case "\n\t":
                for fi180 := 0; fi180 < 17; fi180++ {
                    if len(v58) > 0 {
                        v58[0] = fi63
                    }
                    var v181 uint64 = 1
                    var v182 string = ("" + "a")
                    _ = fi180
                    _ = v181
                    _ = v182
                }
            case "":
                v183 := GenS24{F0: 1, F1: 479}
                _ = v183
            default:
                var v184 string = ("foo bar" + "a")
                _ = v184
            }
        }
        _ = fi63
    }
    var v185 uint32 = 498
    var v186 uint8 = 0
    var v187 float64 = ((0.0 / 3.14159) / (0.1 / -2.25))
    v62 = (v62 / ((v62) | 1))
    var v188 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v185:", v185)
    fmt.Println("v186:", v186)
    fmt.Println("v187:", v187)
    fmt.Println("v188:", v188)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v58:", len(v58), v58)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v185
    _ = v186
    _ = v187
    _ = v188
    _ = v2
    _ = v3
    _ = v4
    _ = v58
    _ = v62
}
