package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 string; F1 int64; F2 bool }
type GenS4 struct { F0 uint32; F1 string; F2 uint8 }
type GenS5 struct { F0 uint; F1 uint8 }
type GenS6 struct { F0 bool; F1 uint8 }
type GenS7 struct { F0 int32; F1 uint32; F2 string }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int16; F1 uint64 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 uint32; F1 bool; F2 int16 }
type GenS15 struct { F0 uint }
func genFn0() uint16 {
    switch "" {
    case "a":
        v5 := []int{42, -2147483648}
        _ = v5
    case "foo bar":
        var v6 int32 = -1
        _ = v6
    case "hello":
        if true {
            for fi7 := 0; fi7 < 11; fi7++ {
                v8 := []int{fi7, fi7, 0, (fi7 << ((0) & 63))}
                var v9 int = fi7
                v8 = append(v8, v9)
                _ = fi7
                _ = v8
                _ = v9
            }
            var v10 uint8 = 146
            _ = v10
        }
    }
    v11 := GenS1{F0: 4294967295}
    var v12 int8 = 42
    _ = v11
    _ = v12
    return 837
}
func genFn1(p0 int8, p1 string, p2 string) int32 {
    if true {
        switch "hello" {
        case "hello":
            fmt.Println(p1, p2)
        case "a":
            v13 := GenS2{F0: 42, F1: true}
            _ = v13
        case "foo bar":
            v14 := []int{4294967295, 32767, 127}
            _ = v14
        default:
            v15 := GenS3{F0: "\n\t", F1: 1, F2: false}
            _ = v15
        }
    } else {
        p1 = p1
    }
    v16 := GenS2{F0: 65535, F1: ((-1 >= 1) || (!(!(4294967295 <= 2))))}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 2147483647
}

func main() {
    v1 := []uint32{}
    v2 := []int{-617}
    v3 := make(map[uint32]int)
    v4 := make(map[string]bool)
    var v17 bool = false
    var v18 int32 = -1
    if v17 {
        var v19 int = 32767
        v17 = (269 != 2147483648)
        for fi20 := 0; fi20 < 20; fi20++ {
            if len(v2) > 0 {
                v2[0] = v19
            }
            if (v18 <= (v18 << ((1) & 63))) {
                if (!v17) {
                    v17 = v17
                    v18 += genFn1(100, "hello", ("foo bar" + "\n\t"))
                }
                for fi21 := 0; fi21 < 19; fi21++ {
                    var v22 int64 = 42
                    fmt.Println(v17)
                    v23 := GenS4{F0: 2, F1: "foo bar", F2: 222}
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            } else {
                for fi24 := 0; fi24 < 11; fi24++ {
                    var v25 int = fi20
                    var v26 float32 = (0.1 / 0.1)
                    var v27 float32 = v26
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                }
            }
            _ = fi20
        }
        _ = v19
    } else {
        if v17 {
            if (644 != 0) {
                if v17 {
                    switch "" {
                    case "\n\t":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    v28 := GenS1{F0: 42}
                    if len(v1) > 0 {
                        v1[0] = 42
                    }
                    _ = v28
                }
            }
            fmt.Println(v17)
            v29 := GenS5{F0: 2, F1: 0}
            _ = v29
        }
    }
    var v30 uint64 = 42
    v1 = append(v1, 42)
    if (58 == 2147483648) {
        v31 := GenS6{F0: false, F1: 156}
        _ = v31
    } else {
        v32 := GenS7{F0: (v18 / ((v18) | 1)), F1: 211, F2: "hello"}
        _ = v32
    }
    switch "" {
    case "hello":
        v4["foo bar"] = v17
    case "":
        switch 4294967295 {
        case 128:
            for fi33 := 0; fi33 < 2; fi33++ {
                if ((0.0 > 3.14159) || (v17 && (!true))) {
                    if v17 {
                        v1 = append(v1, 4294967295)
                    }
                }
                if (v17 && (true || (v17 || (!(v17 && true))))) {
                    if v17 {
                        var v34 string = "a"
                        v34 = (v34 + (v34 + (v34 + v34)))
                        var v35 uint = 0
                        _ = v34
                        _ = v35
                    }
                }
                _ = fi33
            }
        case -747:
            for fi36 := 0; fi36 < 2; fi36++ {
                v37 := GenS3{F0: "hello", F1: 1, F2: true}
                _ = fi36
                _ = v37
            }
        case 839:
            if (!v17) {
                switch ("foo bar" + ("foo bar" + ("" + ("foo bar" + "foo bar")))) {
                case "hello":
                    v38 := v4["hello"]
                    _ = v38
                case "":
                    if false {
                        if v17 {
                            switch "" {
                            case "foo bar":
                                for fi39 := 0; fi39 < 16; fi39++ {
                                    if true {
                                        fmt.Println(fi39, fi39)
                                        fmt.Println(fi39, v30, v17)
                                    } else {
                                        fmt.Println(fi39)
                                    }
                                    v40 := GenS3{F0: "\n\t", F1: 42, F2: v17}
                                    _ = fi39
                                    _ = v40
                                }
                            }
                            if (v17 && (!v17)) {
                                v18 = (v18 >> ((255) & 63))
                                for fi41 := 0; fi41 < 10; fi41++ {
                                    if v17 {
                                        if (((true || (620 <= 0)) && v17) || ((3.14159 * -2.25) == ((0.0 * -2.25) * -2.25))) {
                                            v17 = (!(255 <= 567))
                                            for fi42 := 0; fi42 < 1; fi42++ {
                                                if (!true) {
                                                    fmt.Println(v30, fi42)
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                } else {
                                                    if v17 {
                                                        var v43 uint8 = 74
                                                        v44 := GenS8{F0: -2147483648}
                                                        if (42 <= -1) {
                                                            var v45 int = fi42
                                                            var v46 int8 = 39
                                                            fmt.Println(v45, v18)
                                                            _ = v45
                                                            _ = v46
                                                        }
                                                        _ = v43
                                                        _ = v44
                                                    }
                                                    if (-1 <= -5) {
                                                        v18 += 1
                                                        var v47 uint16 = 42
                                                        _ = v47
                                                    } else {
                                                        v18 += -2147483648
                                                    }
                                                    var v48 float64 = ((0.0 * (0.0 / 3.14159)) + -2.25)
                                                    _ = v48
                                                }
                                                switch "\n\t" {
                                                case "a":
                                                    v30 = 112
                                                }
                                                _ = fi42
                                            }
                                            if (231 >= 255) {
                                                fmt.Println(v30, v30)
                                                switch (fi41 & fi41) {
                                                case 1:
                                                    if ("" <= "\n\t") {
                                                        v30 -= (v30 / (((v30 & v30)) | 1))
                                                    }
                                                case 500:
                                                    fmt.Println(fi41, v17)
                                                case -808:
                                                    v49 := GenS5{F0: 65535, F1: 0}
                                                    _ = v49
                                                }
                                            } else {
                                                if false {
                                                    var v50 uint64 = (^(v30))
                                                    v30 += v30
                                                    var v51 int16 = 42
                                                    _ = v50
                                                    _ = v51
                                                }
                                            }
                                        }
                                        switch (fi41 / ((int(v30)) | 1)) {
                                        case 9223372036854775807:
                                            v30 -= v30
                                        case -742:
                                            var v52 int = fi41
                                            _ = v52
                                        case 563:
                                            for fi53 := 0; fi53 < 9; fi53++ {
                                                fmt.Println(v30, fi41, v17)
                                                if (0 <= 395) {
                                                    if len(v2) > 0 {
                                                        v2[0] = fi53
                                                    }
                                                }
                                                _ = fi53
                                            }
                                        default:
                                            fmt.Println(fi41, v18)
                                        }
                                    } else {
                                        if (!true) {
                                            switch fi41 {
                                            case 32768:
                                                var v54 bool = (4294967295 >= 406)
                                                _ = v54
                                            }
                                            v17 = v17
                                        }
                                        var v55 float64 = 1.5
                                        _ = v55
                                    }
                                    _ = fi41
                                }
                                var v56 uint = 65535
                                _ = v56
                            } else {
                                for fi57 := 0; fi57 < 6; fi57++ {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    var v58 uint = 592
                                    _ = fi57
                                    _ = v58
                                }
                                if (945 > 18446744073709551615) {
                                    v59 := GenS7{F0: 231, F1: 4294967295, F2: "hello"}
                                    switch 42 {
                                    case 2147483647:
                                        if v17 {
                                            switch 127 {
                                            case 939:
                                                v17 = v17
                                            case 59:
                                                v17 = (!false)
                                            case -557:
                                                v30 -= 42
                                            }
                                            var v60 int = 127
                                            _ = v60
                                        }
                                    default:
                                        var v61 string = "\n\t"
                                        _ = v61
                                    }
                                    _ = v59
                                }
                            }
                            switch 2147483647 {
                            case 4294967295:
                                if len(v2) > 0 {
                                    v2[0] = 32768
                                }
                            case 32767:
                                var v62 int = 421
                                _ = v62
                            case -9223372036854775808:
                                if (2 <= 4294967295) {
                                    if v17 {
                                        fmt.Println(v17, v17)
                                        v63 := GenS9{F0: -1}
                                        switch "\n\t" {
                                        case "foo bar":
                                            switch -62 {
                                            case -281:
                                                var v64 int16 = 0
                                                _ = v64
                                            }
                                        case "a":
                                            var v65 string = "\n\t"
                                            _ = v65
                                        default:
                                            v17 = (v17 || v17)
                                        }
                                        _ = v63
                                    }
                                } else {
                                    v66 := GenS10{F0: 32768}
                                    _ = v66
                                }
                            default:
                                switch ((("" + "a") + ("" + ("foo bar" + "foo bar"))) + "hello") {
                                case "hello":
                                    switch 2147483647 {
                                    case -9223372036854775808:
                                        v3[242] = 9223372036854775807
                                    case 127:
                                        v3[299] = 743
                                    case 0:
                                        v17 = (3.14159 == -2.25)
                                    }
                                case "":
                                    v30 -= 0
                                default:
                                    switch -9223372036854775808 {
                                    case 42:
                                        if v17 {
                                            v18 = v18
                                            v17 = ((0.1 > (((float64(3.14159) / (((0.0 * ((1.5 - (0.1 / float64(0.1))) + -2.25)) + 3.14159) + 1e10)) - 0.1) - (((0.1 + -2.25) - 1e10) * 3.14159))) || v17)
                                            switch 4294967295 {
                                            case 9223372036854775807:
                                                if v17 {
                                                    for fi67 := 0; fi67 < 11; fi67++ {
                                                        v68 := GenS9{F0: 0}
                                                        _ = fi67
                                                        _ = v68
                                                    }
                                                } else {
                                                    var v69 uint32 = 894
                                                    for fi70 := 0; fi70 < 14; fi70++ {
                                                        v18 = (v18 >> ((270) & 63))
                                                        _ = fi70
                                                    }
                                                    _ = v69
                                                }
                                            case 852:
                                                for fi71 := 0; fi71 < 10; fi71++ {
                                                    for fi72 := 0; fi72 < 18; fi72++ {
                                                        switch "" {
                                                        case "a":
                                                            fmt.Println(v17)
                                                        case "foo bar":
                                                            switch (^(fi71)) {
                                                            case 9223372036854775807:
                                                                if v17 {
                                                                    v18 -= v18
                                                                    switch fi71 {
                                                                    case -2147483648:
                                                                        v30 += (^(v30))
                                                                    case 1:
                                                                        var v73 uint64 = (v30 << ((255) & 63))
                                                                        _ = v73
                                                                    default:
                                                                        if (!v17) {
                                                                            v18 += 2147483647
                                                                            if v17 {
                                                                                var v74 int64 = -630
                                                                                _ = v74
                                                                            } else {
                                                                                if v17 {
                                                                                    if true {
                                                                                        fmt.Println(v17, v30, fi72)
                                                                                        var v75 string = ("foo bar" + "a")
                                                                                        _ = v75
                                                                                    }
                                                                                    v18 = (v18 - (v18 >> ((42) & 63)))
                                                                                    if true {
                                                                                        v18 += (v18 | v18)
                                                                                        if len(v2) >= 2 {
                                                                                            v2 = v2[1:len(v2)]
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    switch 0 {
                                                                                    case 128:
                                                                                        v18 = v18
                                                                                    case 4294967295:
                                                                                        v76 := GenS7{F0: 2147483647, F1: 65535, F2: "a"}
                                                                                        _ = v76
                                                                                    case -641:
                                                                                        if v17 {
                                                                                            switch 1 {
                                                                                            case 128:
                                                                                                v77 := v3[4294967295]
                                                                                                _ = v77
                                                                                            case 32767:
                                                                                                _ = len(v2)
                                                                                            }
                                                                                            var v78 uint32 = 2147483648
                                                                                            _ = v78
                                                                                        }
                                                                                    }
                                                                                    v3[65535] = (fi71 + 32768)
                                                                                    var v79 uint64 = 42
                                                                                    _ = v79
                                                                                }
                                                                                if (!v17) {
                                                                                    var v80 int32 = -2147483648
                                                                                    _ = v80
                                                                                } else {
                                                                                    if v17 {
                                                                                        v30 = v30
                                                                                    }
                                                                                }
                                                                            }
                                                                            if len(v2) > 0 {
                                                                                v2[0] = (-(fi71))
                                                                            }
                                                                        } else {
                                                                            v30 += v30
                                                                            v81 := GenS2{F0: 4294967295, F1: ((fi71 / ((fi71) | 1)) != (fi71 >> ((65535) & 63)))}
                                                                            _ = v81
                                                                        }
                                                                    }
                                                                    fmt.Println(fi72, v17, fi71)
                                                                }
                                                            default:
                                                                v82 := GenS3{F0: ("hello" + "\n\t"), F1: -1, F2: (!false)}
                                                                _ = v82
                                                            }
                                                        case "":
                                                            fmt.Println(v18)
                                                        }
                                                        v4[(("\n\t" + ("" + "")) + ("hello" + ("hello" + ("\n\t" + "hello"))))] = v17
                                                        var v83 int32 = v18
                                                        _ = fi72
                                                        _ = v83
                                                    }
                                                    for fi84 := 0; fi84 < 6; fi84++ {
                                                        if (v17 && ((!(!true)) || true)) {
                                                            v17 = (!false)
                                                            switch "hello" {
                                                            case "a":
                                                                if v17 {
                                                                    v18 -= genFn1(100, "", "a")
                                                                    fmt.Println(v17, v18, fi84)
                                                                    if (34 <= 1000) {
                                                                        switch (("foo bar" + "hello") + "hello") {
                                                                        case "\n\t":
                                                                            v85 := GenS3{F0: "a", F1: 9223372036854775807, F2: (2147483647 > 1)}
                                                                            _ = v85
                                                                        case "hello":
                                                                            v86 := v4[("foo bar" + "\n\t")]
                                                                            _ = v86
                                                                        }
                                                                        v87 := GenS8{F0: (fi71 ^ 127)}
                                                                        var v88 uint64 = v30
                                                                        var v89 uint32 = 1
                                                                        _ = v87
                                                                        _ = v88
                                                                        _ = v89
                                                                    } else {
                                                                        var v90 int8 = -1
                                                                        if (v17 || v17) {
                                                                            v91 := GenS11{F0: 1, F1: v30}
                                                                            var v92 float64 = (1e10 - ((0.1 / (-2.25 / 3.14159)) / (1e10 / -2.25)))
                                                                            v30 -= 1
                                                                            v90 = int8(fi84)
                                                                            var v93 bool = false
                                                                            _ = v91
                                                                            _ = v92
                                                                            _ = v93
                                                                        }
                                                                        _ = v90
                                                                    }
                                                                }
                                                            case "\n\t":
                                                                delete(v3, 405)
                                                            }
                                                            _ = len(v1)
                                                            var v94 int = (fi84 - 318)
                                                            if ((!true) || (360 > 4294967295)) {
                                                                v94 = fi84
                                                                fmt.Println(fi71)
                                                                if v17 {
                                                                    if true {
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        fmt.Println(fi71, v30)
                                                                    } else {
                                                                        _ = len(v2)
                                                                    }
                                                                    if (!(((-(fi71)) < fi71) && (4294967295 > 277))) {
                                                                        _ = len(v3)
                                                                        v95 := GenS2{F0: 18446744073709551615, F1: (v17 && ((v17 || v17) && false))}
                                                                        _ = v95
                                                                    }
                                                                    v96 := GenS12{F0: v30}
                                                                    fmt.Println(v17, fi84, v94)
                                                                    var v97 bool = (0.0 < -2.25)
                                                                    _ = v96
                                                                    _ = v97
                                                                }
                                                            }
                                                            switch (v94 * fi71) {
                                                            case 4294967295:
                                                                v98 := GenS13{F0: 0}
                                                                _ = v98
                                                            case -708:
                                                                v99 := GenS13{F0: genFn1(127, "hello", "foo bar")}
                                                                _ = v99
                                                            case 696:
                                                                switch "\n\t" {
                                                                case "":
                                                                    fmt.Println(fi71, fi84)
                                                                case "\n\t":
                                                                    if (!(genFn1(42, (("\n\t" + "") + "hello"), ("" + (("foo bar" + "a") + "foo bar"))) != v18)) {
                                                                        v18 += v18
                                                                        var v100 float32 = float32(-2.25)
                                                                        v101 := GenS14{F0: 1, F1: (false && (!(44 < 100))), F2: 0}
                                                                        _ = v100
                                                                        _ = v101
                                                                    }
                                                                case "a":
                                                                    v17 = v17
                                                                }
                                                            default:
                                                                v102 := GenS13{F0: 1}
                                                                v103 := GenS3{F0: "foo bar", F1: 161, F2: v17}
                                                                v17 = (0.0 < float64(((1.5 / (1e10 / 1.0)) - float32(1e10))))
                                                                _ = v102
                                                                _ = v103
                                                            }
                                                            var v104 string = "a"
                                                            v3[1] = fi84
                                                            _ = v104
                                                            _ = v94
                                                        }
                                                        v105 := v4[("" + "\n\t")]
                                                        _ = fi84
                                                        _ = v105
                                                    }
                                                    for fi106 := 0; fi106 < 13; fi106++ {
                                                        delete(v3, 973)
                                                        _ = fi106
                                                    }
                                                    _ = fi71
                                                }
                                            case -691:
                                                var v107 uint32 = 82
                                                _ = v107
                                            default:
                                                v30 -= v30
                                            }
                                        }
                                    case -527:
                                        v17 = ((2 <= 255) && (!(!v17)))
                                    default:
                                        v18 = v18
                                    }
                                }
                            }
                        }
                        v17 = (0.0 != 1e10)
                        if len(v1) > 0 {
                            v1[0] = 2
                        }
                    }
                default:
                    fmt.Println(v17)
                }
                v108 := GenS11{F0: -32768, F1: v30}
                v109 := GenS2{F0: 2, F1: v17}
                _ = v108
                _ = v109
            } else {
                switch "hello" {
                case "":
                    v110 := GenS1{F0: v30}
                    _ = v110
                default:
                    v30 = (^(v30))
                }
                var v111 uint64 = (^(v30))
                _ = v111
            }
        }
    default:
        v30 += (v30 >> ((37) & 63))
    }
    _ = len(v4)
    for fi112 := 0; fi112 < 4; fi112++ {
        v113 := GenS6{F0: (1 != uint32(v30)), F1: 0}
        _ = fi112
        _ = v113
    }
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        for fi114 := 0; fi114 < 17; fi114++ {
            var v115 uint = 4294967295
            v116 := GenS7{F0: (v18 / (((v18 << ((18446744073709551615) & 63))) | 1)), F1: 2147483648, F2: "hello"}
            fmt.Println(v115, v115, v17)
            _ = fi114
            _ = v115
            _ = v116
        }
    case "a":
        delete(v4, "")
    default:
        var v117 int64 = 1
        _ = v117
    }
    for fi118 := 0; fi118 < 12; fi118++ {
        v119 := GenS15{F0: 65535}
        var v120 uint16 = 2
        if true {
            switch (fi118 << ((1) & 63)) {
            case 127:
                v120 = v120
            case -1:
                v30 -= 4294967295
            }
            v121 := GenS11{F0: 0, F1: (^(v30))}
            _ = v121
        } else {
            for fi122 := 0; fi122 < 1; fi122++ {
                if (0 != -64) {
                    fmt.Println(fi122, v18, v18)
                } else {
                    v18 = (v18 ^ (v18 - v18))
                    var v123 int16 = -180
                    _ = v123
                }
                _ = fi122
            }
            for fi124 := 0; fi124 < 16; fi124++ {
                v17 = v17
                delete(v3, 835)
                var v125 uint8 = 75
                _ = fi124
                _ = v125
            }
        }
        _ = fi118
        _ = v119
        _ = v120
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}
