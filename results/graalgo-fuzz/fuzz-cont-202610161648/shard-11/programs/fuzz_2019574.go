package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 uint32; F2 uint }
type GenS4 struct { F0 int; F1 int16; F2 int }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint; F1 string }
type GenS7 struct { F0 int8; F1 string }
type GenS8 struct { F0 uint16; F1 int32; F2 int64 }
type GenS9 struct { F0 bool; F1 int32 }
type GenS10 struct { F0 int32; F1 int16; F2 string }
type GenS11 struct { F0 int32; F1 int8; F2 uint32 }
func genFn0(p0 int32, p1 uint32) uint8 {
    v4 := []int{128, 0}
    var v5 int = 32768
    if false {
        v5 += (v5 % (((v5 + v5)) | 1))
        if (42 != 18446744073709551615) {
            v5 = v5
            var v6 int8 = 42
            p1 -= (p1 | 1)
            _ = v6
        } else {
            _ = len(v4)
        }
        var v7 uint64 = 904
        _ = v7
    } else {
        var v8 int = v5
        _ = v8
    }
    p1 += p1
    fmt.Println(v5)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 0
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = 0
    v3 := make(map[int]int)
    v9 := GenS1{F0: -116, F1: 853, F2: "a"}
    var v10 uint64 = 489
    v11 := []int{-2147483648}
    fmt.Println(v2, v2)
    if true {
        if false {
            v12 := GenS2{F0: 580}
            _ = v12
        }
        var v13 uint8 = genFn0(1, 0)
        _ = v13
    } else {
        fmt.Println(v10, v2)
        v14 := GenS2{F0: 1}
        var v15 uint8 = 228
        _ = v14
        _ = v15
    }
    v2 = v2
    v16 := GenS2{F0: 0}
    for fi17 := 0; fi17 < 19; fi17++ {
        v18 := GenS3{F0: -1, F1: 2147483648, F2: 0}
        _ = fi17
        _ = v18
    }
    v2 += v2
    fmt.Println(v2)
    if (4294967295 > 2147483647) {
        var v19 int = 128
        if true {
            v3[v19] = (v19 * (v19 / ((v19) | 1)))
            for fi20 := 0; fi20 < 13; fi20++ {
                var v21 int16 = (v2 >> ((42) & 63))
                v22 := GenS4{F0: v19, F1: (v21 << ((18446744073709551615) & 63)), F2: v19}
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = len(v1)
        }
        _ = v19
    } else {
        var v23 float64 = 0.1
        if (!(3.14159 >= ((float32((((v23 * float64((0.0 + (((1e10 / ((-2.25 - 0.0) * 1.5)) + -2.25) - 0.0)))) - v23) - (v23 / float64(3.14159)))) / (0.1 / 3.14159)) / ((-2.25 * ((-2.25 + float32(3.14159)) / (3.14159 - ((3.14159 * (((1.5 - 1e10) + float32((v23 - float64((0.0 / 1.0))))) * 1.5)) * (1.5 - (1e10 + (1.5 - 3.14159))))))) * 3.14159)))) {
            var v24 bool = true
            var v25 int32 = -1
            _ = v24
            _ = v25
        }
        _ = v23
    }
    fmt.Println(v10)
    v10 += (v10 / ((v10) | 1))
    v1["a"] = "hello"
    switch ("foo bar" + "foo bar") {
    case "hello":
        fmt.Println(v10, v2)
    }
    _ = len(v3)
    if true {
        if false {
            v26 := GenS5{F0: false}
            for fi27 := 0; fi27 < 3; fi27++ {
                if ((2 > 1) || (!(!false))) {
                    for fi28 := 0; fi28 < 11; fi28++ {
                        _ = len(v1)
                        _ = len(v3)
                        var v29 uint64 = 797
                        _ = fi28
                        _ = v29
                    }
                    switch ("a" + "") {
                    case "hello":
                        switch "\n\t" {
                        case "\n\t":
                            var v30 int = 128
                            _ = v30
                        default:
                            v31 := GenS4{F0: (fi27 << ((601) & 63)), F1: (v2 << ((255) & 63)), F2: fi27}
                            _ = v31
                        }
                    default:
                        var v32 int8 = 127
                        _ = v32
                    }
                }
                switch (("foo bar" + (("hello" + ((("a" + "a") + "foo bar") + "hello")) + "foo bar")) + "a") {
                case "":
                    for fi33 := 0; fi33 < 2; fi33++ {
                        v1[("foo bar" + "hello")] = ("foo bar" + "hello")
                        if false {
                            v34 := GenS4{F0: fi27, F1: (v2 & v2), F2: -849}
                            var v35 int64 = int64(fi33)
                            _ = v34
                            _ = v35
                        } else {
                            v2 = v2
                            fmt.Println(fi27, v2, v10)
                            var v36 uint = 1
                            if ((v2 % (((v2 | 1000)) | 1)) <= (^(v2))) {
                                v2 -= (v2 ^ -32768)
                                if len(v11) > 0 {
                                    v11[0] = (fi33 >> (((v36 * 255)) & 63))
                                }
                            }
                            var v37 int32 = -392
                            if (!false) {
                                fmt.Println(fi27, v37)
                                var v38 int32 = (v37 >> (((v36 ^ 945)) & 63))
                                delete(v3, fi33)
                                _ = v38
                            }
                            v39 := v3[(fi27 << ((396) & 63))]
                            switch (fi33 * (-(fi33))) {
                            case -9223372036854775808:
                                var v40 uint32 = 120
                                _ = v40
                            case -2147483648:
                                v36 -= 42
                            case 261:
                                v10 = 42
                            }
                            _ = v36
                            _ = v37
                            _ = v39
                        }
                        _ = fi33
                    }
                case "\n\t":
                    v2 += (v2 >> ((255) & 63))
                default:
                    v2 -= (v2 << ((65535) & 63))
                }
                for fi41 := 0; fi41 < 10; fi41++ {
                    v10 = (^(v10))
                    _ = fi41
                }
                _ = fi27
            }
            if (!((float32((((1.5 / ((0.1 * 0.1) - ((0.1 - 1.5) / float64(3.14159)))) + 0.0) / -2.25)) - (((0.0 - (3.14159 + 0.0)) * 0.0) - 3.14159)) == (1e10 - (0.1 / (1.5 * float32((0.1 / 0.1))))))) {
                v3[-586] = -2147483648
                for fi42 := 0; fi42 < 8; fi42++ {
                    var v43 uint32 = 1
                    switch (fi42 * fi42) {
                    case 518:
                        var v44 int32 = -1
                        _ = v44
                    case -2147483648:
                        v10 = (v10 ^ v10)
                    }
                    switch (fi42 | (^(fi42))) {
                    case 2147483647:
                        var v45 uint = 65535
                        _ = v45
                    case 9223372036854775807:
                        v46 := GenS3{F0: -65, F1: (v43 << ((uint(v10)) & 63)), F2: 42}
                        _ = v46
                    case 32767:
                        if false {
                            v2 -= v2
                        } else {
                            switch (fi42 + 42) {
                            case 88:
                                fmt.Println(v43)
                            case 0:
                                if (-2147483648 == int32(v2)) {
                                    var v47 float32 = 1e10
                                    _ = v47
                                }
                            default:
                                _ = len(v1)
                            }
                        }
                    default:
                        var v48 string = "\n\t"
                        _ = v48
                    }
                    _ = fi42
                    _ = v43
                }
            } else {
                if false {
                    switch "a" {
                    case "a":
                        v10 -= 515
                    case "hello":
                        for fi49 := 0; fi49 < 13; fi49++ {
                            var v50 int32 = 2147483647
                            _ = fi49
                            _ = v50
                        }
                    default:
                        if (false || (!(!(((-2.25 + ((3.14159 / 1e10) * (1e10 * 1e10))) + 0.0) == (((1.5 * 0.0) * ((-2.25 - 1e10) * 0.0)) - 1e10))))) {
                            if (98 > -105) {
                                v2 = 42
                                var v51 int32 = 420
                                _ = v51
                            }
                        } else {
                            switch -1 {
                            case 307:
                                var v52 bool = (32767 <= 617)
                                _ = v52
                            default:
                                v53 := GenS5{F0: (v2 >= v2)}
                                _ = v53
                            }
                            v54 := GenS3{F0: 100, F1: 170, F2: uint(v10)}
                            if (true && false) {
                                var v55 string = "hello"
                                switch ("\n\t" + "\n\t") {
                                case "hello":
                                    var v56 bool = ((v55 > "foo bar") || false)
                                    _ = v56
                                case "a":
                                    if false {
                                        var v57 uint8 = 1
                                        v55 = v55
                                        for fi58 := 0; fi58 < 9; fi58++ {
                                            v59 := GenS4{F0: 9223372036854775807, F1: v2, F2: (fi58 - -9223372036854775808)}
                                            if ((1e10 == 0.0) && true) {
                                                fmt.Println(v2, v55)
                                            }
                                            switch fi58 {
                                            case -464:
                                                var v60 uint = 255
                                                _ = v60
                                            case 415:
                                                v61 := GenS1{F0: 127, F1: 1, F2: "\n\t"}
                                                _ = v61
                                            }
                                            _ = fi58
                                            _ = v59
                                        }
                                        _ = v57
                                    } else {
                                        fmt.Println(v2)
                                        var v62 float32 = 3.14159
                                        var v63 uint32 = 71
                                        _ = v62
                                        _ = v63
                                    }
                                }
                                v55 = (v55 + ("\n\t" + (v55 + (("" + "") + "\n\t"))))
                                _ = v55
                            } else {
                                v2 -= 1000
                                var v64 float32 = float32((0.1 * (0.0 * 3.14159)))
                                if true {
                                    v65 := GenS6{F0: 255, F1: ("" + "")}
                                    switch ("a" + (("" + "foo bar") + ("" + ""))) {
                                    case "hello":
                                        var v66 float32 = (v64 / float32(3.14159))
                                        _ = v66
                                    case "\n\t":
                                        v2 -= -1
                                    default:
                                        var v67 uint16 = 32767
                                        _ = v67
                                    }
                                    v68 := v3[-848]
                                    _ = v65
                                    _ = v68
                                } else {
                                    v69 := GenS4{F0: 4294967295, F1: v2, F2: -2147483648}
                                    for fi70 := 0; fi70 < 9; fi70++ {
                                        if true {
                                            if false {
                                                var v71 int32 = -553
                                                for fi72 := 0; fi72 < 16; fi72++ {
                                                    var v73 float32 = (v64 / ((0.0 - -2.25) - float32((3.14159 / 1.5))))
                                                    if len(v11) >= 2 {
                                                        v11 = v11[1:len(v11)]
                                                    }
                                                    _ = fi72
                                                    _ = v73
                                                }
                                                _ = v71
                                            }
                                            if len(v11) > 0 {
                                                v11[0] = (^(fi70))
                                            }
                                            for fi74 := 0; fi74 < 1; fi74++ {
                                                var v75 int16 = v2
                                                if ((("" + ((("hello" + "hello") + "foo bar") + "a")) >= "hello") && (("" + "foo bar") != "\n\t")) {
                                                    v75 = -1
                                                    var v76 string = "\n\t"
                                                    var v77 string = (v76 + v76)
                                                    _ = v76
                                                    _ = v77
                                                }
                                                v64 -= v64
                                                _ = fi74
                                                _ = v75
                                            }
                                        } else {
                                            var v78 uint64 = v10
                                            _ = v78
                                        }
                                        switch 127 {
                                        case -9223372036854775808:
                                            if (v64 <= float32(1e10)) {
                                                fmt.Println(v64)
                                                var v79 uint64 = (v10 - (^(v10)))
                                                _ = v79
                                            } else {
                                                switch fi70 {
                                                case 9223372036854775807:
                                                    if len(v11) > 0 {
                                                        v11[0] = (fi70 * fi70)
                                                    }
                                                case -1:
                                                    if len(v11) > 0 {
                                                        v11[0] = -1
                                                    }
                                                default:
                                                    if (v64 > v64) {
                                                        v2 -= v2
                                                        if (1 < 2) {
                                                            if true {
                                                                v10 = v10
                                                            } else {
                                                                var v80 float32 = v64
                                                                var v81 int = -9223372036854775808
                                                                var v82 float64 = ((((-2.25 * (0.0 + 0.0)) * (((0.0 + ((-2.25 / 1e10) / -2.25)) / 1.5) - (0.1 + (0.1 - (1.5 - 0.0))))) - (0.0 / 1e10)) / (float64(v64) * (1.5 / 1.0)))
                                                                _ = v80
                                                                _ = v81
                                                                _ = v82
                                                            }
                                                            fmt.Println(v64, fi70, fi70)
                                                            switch ("foo bar" + ((("a" + "\n\t") + "hello") + "")) {
                                                            case "\n\t":
                                                                fmt.Println(v2, v10, fi70)
                                                            case "foo bar":
                                                                fmt.Println(v64)
                                                            case "":
                                                                if (true && (0.0 != -2.25)) {
                                                                    for fi83 := 0; fi83 < 20; fi83++ {
                                                                        v84 := GenS6{F0: 1, F1: (("hello" + "") + "\n\t")}
                                                                        v1[("\n\t" + ("a" + "a"))] = ("\n\t" + "hello")
                                                                        _ = fi83
                                                                        _ = v84
                                                                    }
                                                                    if false {
                                                                        v85 := GenS7{F0: 78, F1: "foo bar"}
                                                                        _ = v85
                                                                    }
                                                                } else {
                                                                    var v86 int8 = 42
                                                                    for fi87 := 0; fi87 < 4; fi87++ {
                                                                        switch 437 {
                                                                        case 42:
                                                                            if len(v11) > 0 {
                                                                                v11[0] = fi87
                                                                            }
                                                                        case -102:
                                                                            var v88 float32 = (0.0 * v64)
                                                                            _ = v88
                                                                        case 2147483647:
                                                                            var v89 uint16 = 32767
                                                                            _ = v89
                                                                        default:
                                                                            if (-1 < -450) {
                                                                                if len(v11) >= 2 {
                                                                                    v11 = v11[1:len(v11)]
                                                                                }
                                                                            } else {
                                                                                v2 += int16(v10)
                                                                                var v90 uint8 = genFn0(-2147483648, 34)
                                                                                _ = v90
                                                                            }
                                                                        }
                                                                        var v91 uint8 = 207
                                                                        switch "a" {
                                                                        case "hello":
                                                                            var v92 int = fi70
                                                                            _ = v92
                                                                        }
                                                                        _ = fi87
                                                                        _ = v91
                                                                    }
                                                                    _ = v86
                                                                }
                                                            default:
                                                                if (177 > 255) {
                                                                    if (42 != 65535) {
                                                                        v93 := GenS7{F0: -128, F1: ""}
                                                                        v94 := GenS8{F0: 103, F1: -2147483648, F2: -1}
                                                                        _ = v93
                                                                        _ = v94
                                                                    } else {
                                                                        if (true || (127 <= 128)) {
                                                                            v95 := GenS9{F0: (!((v2 > v2) && false)), F1: 2147483647}
                                                                            if true {
                                                                                fmt.Println(v10)
                                                                                if (-2147483648 >= 322) {
                                                                                    if len(v11) >= 2 {
                                                                                        v11 = v11[1:len(v11)]
                                                                                    }
                                                                                    v10 = 0
                                                                                }
                                                                                switch "foo bar" {
                                                                                case "a":
                                                                                    fmt.Println(v2, v64)
                                                                                case "foo bar":
                                                                                    var v96 float32 = (((v64 + float32(0.0)) + v64) * v64)
                                                                                    _ = v96
                                                                                case "\n\t":
                                                                                    var v97 float64 = 0.1
                                                                                    _ = v97
                                                                                }
                                                                            } else {
                                                                                if true {
                                                                                    switch (fi70 / ((-2147483648) | 1)) {
                                                                                    case 1:
                                                                                        v10 -= 4294967295
                                                                                    case 42:
                                                                                        v98 := GenS1{F0: -1, F1: 678, F2: ""}
                                                                                        _ = v98
                                                                                    }
                                                                                }
                                                                            }
                                                                            var v99 string = "\n\t"
                                                                            _ = v95
                                                                            _ = v99
                                                                        } else {
                                                                            switch (("\n\t" + "a") + "\n\t") {
                                                                            case "a":
                                                                                delete(v1, "")
                                                                            case "\n\t":
                                                                                var v100 bool = ((185 >= 127) && (255 > 718))
                                                                                _ = v100
                                                                            }
                                                                            var v101 string = ""
                                                                            _ = v101
                                                                        }
                                                                        for fi102 := 0; fi102 < 17; fi102++ {
                                                                            var v103 bool = true
                                                                            if (202 >= 32767) {
                                                                                if v103 {
                                                                                    v1[""] = "a"
                                                                                    v103 = false
                                                                                }
                                                                            } else {
                                                                                var v104 uint32 = 514
                                                                                var v105 bool = (false && v103)
                                                                                _ = v104
                                                                                _ = v105
                                                                            }
                                                                            v106 := GenS8{F0: 642, F1: 42, F2: 1}
                                                                            _ = fi102
                                                                            _ = v103
                                                                            _ = v106
                                                                        }
                                                                        if true {
                                                                            v64 = float32(-2.25)
                                                                        } else {
                                                                            switch fi70 {
                                                                            case 9223372036854775807:
                                                                                for fi107 := 0; fi107 < 18; fi107++ {
                                                                                    if false {
                                                                                        delete(v1, (("a" + "foo bar") + "a"))
                                                                                        v108 := GenS1{F0: 0, F1: 533, F2: "hello"}
                                                                                        if true {
                                                                                            v64 = v64
                                                                                            var v109 int32 = -1
                                                                                            _ = v109
                                                                                        }
                                                                                        v11 = append(v11, (fi107 / ((4294967295) | 1)))
                                                                                        _ = v108
                                                                                    } else {
                                                                                        v110 := GenS3{F0: -1, F1: 2147483648, F2: 65535}
                                                                                        switch "" {
                                                                                        case "hello":
                                                                                            v111 := GenS6{F0: 255, F1: ""}
                                                                                            _ = v111
                                                                                        }
                                                                                        var v112 uint64 = (v10 ^ v10)
                                                                                        var v113 int32 = -566
                                                                                        _ = v110
                                                                                        _ = v112
                                                                                        _ = v113
                                                                                    }
                                                                                    fmt.Println(v10, v64)
                                                                                    if true {
                                                                                        v114 := GenS8{F0: 1, F1: 570, F2: 9223372036854775807}
                                                                                        _ = v114
                                                                                    }
                                                                                    _ = fi107
                                                                                }
                                                                            case 32768:
                                                                                if (!true) {
                                                                                    v64 = v64
                                                                                    var v115 int16 = 246
                                                                                    v2 = (v2 & v2)
                                                                                    _ = v115
                                                                                } else {
                                                                                    var v116 string = "hello"
                                                                                    for fi117 := 0; fi117 < 1; fi117++ {
                                                                                        var v118 uint8 = 127
                                                                                        v2 = 1000
                                                                                        _ = fi117
                                                                                        _ = v118
                                                                                    }
                                                                                    _ = v116
                                                                                }
                                                                            case 32767:
                                                                                for fi119 := 0; fi119 < 16; fi119++ {
                                                                                    fmt.Println(fi70)
                                                                                    v1["a"] = "a"
                                                                                    fmt.Println(fi70)
                                                                                    _ = fi119
                                                                                }
                                                                            default:
                                                                                var v120 int = (fi70 & (fi70 * 32768))
                                                                                _ = v120
                                                                            }
                                                                            v10 = (v10 << ((65535) & 63))
                                                                        }
                                                                    }
                                                                    v10 = (v10 << ((42) & 63))
                                                                    for fi121 := 0; fi121 < 7; fi121++ {
                                                                        fmt.Println(fi121)
                                                                        v10 += (v10 % ((v10) | 1))
                                                                        var v122 int16 = v2
                                                                        _ = fi121
                                                                        _ = v122
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        v123 := GenS4{F0: 4294967295, F1: v2, F2: (fi70 % (((fi70 + (fi70 | 917))) | 1))}
                                                        _ = v123
                                                    }
                                                }
                                                for fi124 := 0; fi124 < 14; fi124++ {
                                                    var v125 int8 = 100
                                                    v10 = v10
                                                    v2 = (v2 + (v2 ^ 42))
                                                    _ = fi124
                                                    _ = v125
                                                }
                                            }
                                        }
                                        _ = fi70
                                    }
                                    v126 := GenS10{F0: 270, F1: (^(v2)), F2: "a"}
                                    _ = v126
                                    _ = v69
                                }
                                _ = v64
                            }
                            _ = v54
                        }
                    }
                    for fi127 := 0; fi127 < 16; fi127++ {
                        v128 := GenS6{F0: 255, F1: ("\n\t" + ("" + "\n\t"))}
                        _ = fi127
                        _ = v128
                    }
                    v129 := GenS11{F0: -1, F1: 120, F2: 65535}
                    _ = v129
                }
            }
            _ = v26
        } else {
            v130 := GenS1{F0: 0, F1: 65535, F2: "hello"}
            _ = v130
        }
        var v131 string = ""
        v131 = "hello"
        _ = v131
    } else {
        fmt.Println(v2)
    }
    v10 = (v10 + v10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v9
}
