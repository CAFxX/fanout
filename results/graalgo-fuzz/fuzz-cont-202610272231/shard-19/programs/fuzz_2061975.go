package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 int8; F2 int64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int32; F1 bool; F2 int64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint; F1 string; F2 uint8 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint8; F1 bool }
type GenS11 struct { F0 int; F1 uint }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS14 struct { F0 int8; F1 uint16 }
type GenS15 struct { F0 uint16; F1 int16 }
type GenS16 struct { F0 int64; F1 int64 }
type GenS17 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS18 struct { F0 int8; F1 bool }

func main() {
    var v1 int32 = 42
    var v2 int64 = 9223372036854775807
    var v3 int = 128
    v4 := []string{((("\n\t" + (("\n\t" + "a") + ("a" + "a"))) + "foo bar") + ""), "hello", "hello"}
    var v5 int = (v3 % (((v3 | -1)) | 1))
    var v6 bool = false
    if (!((4294967295 > 2147483648) && (2 > 1))) {
        v5 = (-(v5))
        fmt.Println(v6)
    } else {
        var v7 bool = v6
        if v7 {
            if v6 {
                if (2147483647 >= 2147483648) {
                    v8 := GenS1{F0: 2, F1: 2}
                    v9 := GenS1{F0: 255, F1: 0}
                    var v10 int8 = 100
                    _ = v10
                    _ = v8
                    _ = v9
                }
                if (v6 || (!v7)) {
                    var v11 uint = 172
                    _ = v11
                }
                v12 := GenS1{F0: 42, F1: 1}
                _ = v12
            }
            var v13 bool = v7
            for fi14 := 0; fi14 < 10; fi14++ {
                var v15 int8 = 100
                var v16 uint8 = 249
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = v13
        }
        if false {
            switch (-(v3)) {
            case -2147483648:
                v1 = 42
            case -700:
                for fi17 := 0; fi17 < 13; fi17++ {
                    for fi18 := 0; fi18 < 16; fi18++ {
                        var v19 int32 = -2147483648
                        _ = fi18
                        _ = v19
                    }
                    _ = fi17
                }
            default:
                var v20 uint32 = 42
                _ = v20
            }
        }
        _ = v7
    }
    v3 = -1
    for fi21 := 0; fi21 < 17; fi21++ {
        v22 := GenS2{F0: 255}
        switch ("hello" + ("\n\t" + "a")) {
        case "hello":
            var v23 float64 = -2.25
            _ = v23
        case "":
            v4 = append(v4, "")
        case "a":
            for fi24 := 0; fi24 < 14; fi24++ {
                if true {
                    switch (v3 - -2147483648) {
                    case 4294967295:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case 128:
                        var v25 uint64 = 2
                        _ = v25
                    default:
                        if len(v4) > 0 {
                            v4[0] = ""
                        }
                    }
                    var v26 float32 = 1.5
                    v4 = append(v4, ("" + "foo bar"))
                    var v27 string = ("foo bar" + ("\n\t" + "foo bar"))
                    _ = v26
                    _ = v27
                } else {
                    if (false && (-1 > v1)) {
                        if len(v4) > 0 {
                            v4[0] = "\n\t"
                        }
                        fmt.Println(v5, v1)
                        if false {
                            if v6 {
                                if ((v6 && true) && v6) {
                                    v28 := GenS3{F0: -1, F1: 42, F2: -1}
                                    fmt.Println(v6)
                                    v29 := GenS4{F0: 1000}
                                    _ = v28
                                    _ = v29
                                } else {
                                    if (!true) {
                                        if len(v4) > 0 {
                                            v4[0] = ("foo bar" + ((("" + "") + ((("\n\t" + "") + "foo bar") + ("\n\t" + ("hello" + ("\n\t" + "a"))))) + "a"))
                                        }
                                    } else {
                                        var v30 uint16 = 1
                                        if (0 == 127) {
                                            v31 := GenS3{F0: v1, F1: 127, F2: (v2 * v2)}
                                            _ = v31
                                        } else {
                                            if len(v4) > 0 {
                                                v4[0] = ("a" + ((("hello" + "a") + ("hello" + "\n\t")) + ""))
                                            }
                                        }
                                        _ = v30
                                    }
                                    switch fi24 {
                                    case -9223372036854775808:
                                        var v32 uint16 = 42
                                        _ = v32
                                    case 592:
                                        var v33 uint16 = 2
                                        _ = v33
                                    case 9223372036854775807:
                                        v34 := GenS5{F0: 476}
                                        _ = v34
                                    default:
                                        var v35 uint32 = 1
                                        _ = v35
                                    }
                                }
                            }
                        }
                        if (("foo bar" > "hello") || v6) {
                            switch v3 {
                            case 2147483647:
                                var v36 uint64 = 0
                                _ = v36
                            case 295:
                                var v37 float32 = (float32(3.14159) * (3.14159 - 0.1))
                                _ = v37
                            case -562:
                                if len(v4) > 0 {
                                    v4[0] = ("hello" + (("\n\t" + "hello") + "\n\t"))
                                }
                                v1 += (v1 << ((18446744073709551615) & 63))
                                var v38 uint64 = 915
                                _ = v38
                            }
                            var v39 uint8 = 42
                            if (65535 > 404) {
                                v40 := GenS5{F0: 407}
                                _ = v40
                            }
                            _ = v39
                        }
                        fmt.Println(fi24, v3, v2)
                    } else {
                        var v41 uint8 = 216
                        switch (fi24 % (((v5 | v3)) | 1)) {
                        case 175:
                            switch ("a" + "") {
                            case "\n\t":
                                var v42 int64 = v2
                                _ = v42
                            case "foo bar":
                                v43 := GenS1{F0: 42, F1: 486}
                                _ = v43
                            default:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        case 504:
                            v41 = (v41 ^ (v41 / ((v41) | 1)))
                        default:
                            switch (fi21 / ((42) | 1)) {
                            case -382:
                                fmt.Println(v41, v6)
                            default:
                                fmt.Println(v6)
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        }
                        _ = v41
                    }
                    var v44 uint = 4294967295
                    _ = v44
                }
                if v6 {
                    if v6 {
                        var v45 uint16 = 1
                        _ = v45
                    }
                    v5 = fi21
                    var v46 uint32 = 565
                    v2 += int64(v46)
                    _ = v46
                } else {
                    v1 = v1
                    v6 = (v6 || (317 > 757))
                    var v47 bool = (1.5 != (0.0 / -2.25))
                    switch fi24 {
                    case 1:
                        _ = len(v4)
                        var v48 uint32 = 503
                        var v49 int = (v3 - fi24)
                        _ = v48
                        _ = v49
                    case 32768:
                        var v50 uint32 = 1
                        switch "hello" {
                        case "\n\t":
                            switch (-(fi24)) {
                            case 32767:
                                v3 = fi21
                            case 0:
                                switch ("foo bar" + (("\n\t" + "foo bar") + "hello")) {
                                case "":
                                    v51 := GenS6{F0: 705, F1: v47, F2: (v2 + v2)}
                                    _ = v51
                                }
                            case 694:
                                v47 = ((1 != 0) || v6)
                            default:
                                switch ("hello" + ("" + "foo bar")) {
                                case "\n\t":
                                    if v6 {
                                        switch "" {
                                        case "":
                                            v2 = 287
                                        case "a":
                                            fmt.Println(v5, v2, v5)
                                        case "hello":
                                            var v52 int32 = int32(v2)
                                            _ = v52
                                        default:
                                            v2 += (v2 << ((276) & 63))
                                        }
                                    } else {
                                        v53 := GenS7{F0: v47}
                                        var v54 int16 = 63
                                        _ = v53
                                        _ = v54
                                    }
                                }
                            }
                        default:
                            var v55 uint8 = 255
                            _ = v55
                        }
                        _ = v50
                    }
                    switch ("a" + ("hello" + ((("" + ("a" + (("\n\t" + ("\n\t" + "foo bar")) + ("a" + (((("foo bar" + "\n\t") + (("hello" + "") + "\n\t")) + "") + (("hello" + ("hello" + "hello")) + ("\n\t" + "foo bar"))))))) + (("hello" + "") + "hello")) + ("a" + "\n\t")))) {
                    case "\n\t":
                        v47 = v6
                    }
                    if v47 {
                        var v56 float32 = 1.5
                        _ = v56
                    } else {
                        if v47 {
                            switch (fi21 ^ (v5 + (-(v5)))) {
                            case 2147483647:
                                fmt.Println(v6)
                            case -2147483648:
                                switch "hello" {
                                case "\n\t":
                                    v4 = append(v4, "hello")
                                case "a":
                                    v57 := GenS2{F0: 4294967295}
                                    _ = v57
                                case "":
                                    fmt.Println(fi21, fi24, v3)
                                }
                            case -78:
                                v58 := GenS8{F0: 1, F1: "a", F2: 103}
                                _ = v58
                            }
                            v6 = false
                        }
                        if true {
                            v59 := GenS5{F0: 18446744073709551615}
                            fmt.Println(v1, fi24)
                            fmt.Println(v5, v3, v2)
                            _ = v59
                        } else {
                            v5 -= (fi21 + (fi24 - v5))
                            var v60 uint32 = 4294967295
                            var v61 uint = 138
                            switch fi24 {
                            case -426:
                                var v62 uint64 = 18446744073709551615
                                v62 += v62
                                _ = v62
                            case 32768:
                                v63 := GenS9{F0: "\n\t"}
                                _ = v63
                            default:
                                if v47 {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v2 -= (v2 << ((1) & 63))
                                    v4 = append(v4, "\n\t")
                                }
                            }
                            _ = v60
                            _ = v61
                        }
                    }
                    var v64 string = "\n\t"
                    _ = v47
                    _ = v64
                }
                _ = fi24
            }
        }
        var v65 uint32 = 2
        _ = fi21
        _ = v22
        _ = v65
    }
    var v66 int64 = v2
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (("" + ("hello" + ("foo bar" + ("hello" + (("foo bar" + (("hello" + "\n\t") + "a")) + "hello"))))) > "foo bar") {
        var v67 float32 = (float32((float64(1e10) + 1e10)) * 3.14159)
        var v68 float32 = float32(-2.25)
        v69 := GenS10{F0: 238, F1: v6}
        _ = v67
        _ = v68
        _ = v69
    }
    fmt.Println(v5, v66)
    if v6 {
        v1 = v1
        if v6 {
            fmt.Println(v66, v2, v1)
            for fi70 := 0; fi70 < 1; fi70++ {
                v6 = v6
                var v71 int32 = (v1 | 2147483647)
                var v72 int = (fi70 ^ v3)
                _ = fi70
                _ = v71
                _ = v72
            }
            if (255 > 0) {
                v73 := GenS11{F0: (v3 / ((32767) | 1)), F1: 1}
                if false {
                    fmt.Println(v6, v2, v3)
                    for fi74 := 0; fi74 < 6; fi74++ {
                        fmt.Println(v1, v66)
                        v6 = (true || v6)
                        _ = fi74
                    }
                    for fi75 := 0; fi75 < 2; fi75++ {
                        var v76 float64 = 3.14159
                        var v77 uint = 18446744073709551615
                        if (42 > 671) {
                            if (0.0 < 1e10) {
                                if false {
                                    for fi78 := 0; fi78 < 18; fi78++ {
                                        if len(v4) > 0 {
                                            v4[0] = ""
                                        }
                                        _ = fi78
                                    }
                                } else {
                                    if (65535 > 1) {
                                        v4 = append(v4, "hello")
                                        if ((65535 == 42) && (!v6)) {
                                            if (1 == 42) {
                                                if ((fi75 >> ((533) & 63)) == (v5 + (^(fi75)))) {
                                                    var v79 bool = false
                                                    _ = v79
                                                }
                                                v4 = append(v4, ("foo bar" + "foo bar"))
                                            }
                                            fmt.Println(v2)
                                            var v80 uint32 = 2
                                            _ = v80
                                        } else {
                                            if (1 <= -70) {
                                                if ((v77 & 352) <= v77) {
                                                    if ((v2 << ((1) & 63)) <= v2) {
                                                        fmt.Println(v66)
                                                        if v6 {
                                                            v76 = (1e10 * v76)
                                                            var v81 float32 = (0.0 + -2.25)
                                                            _ = v81
                                                        }
                                                        v1 -= v1
                                                    }
                                                    if false {
                                                        for fi82 := 0; fi82 < 15; fi82++ {
                                                            switch "hello" {
                                                            case "a":
                                                                var v83 uint = 138
                                                                _ = v83
                                                            case "":
                                                                fmt.Println(v2)
                                                            case "hello":
                                                                if ((-32768 >= 1000) || (!v6)) {
                                                                    if ((v6 && (48 <= -118)) && false) {
                                                                        v2 += v2
                                                                        if (-32768 >= -686) {
                                                                            v3 = fi75
                                                                            if v6 {
                                                                                fmt.Println(v6, v77, v76)
                                                                                v3 = (fi82 | 102)
                                                                                v66 -= v66
                                                                                v4 = append(v4, ("a" + ""))
                                                                            }
                                                                            v84 := GenS5{F0: 18446744073709551615}
                                                                            _ = v84
                                                                        } else {
                                                                            var v85 int32 = (-(v1))
                                                                            v76 = v76
                                                                            _ = v85
                                                                        }
                                                                    }
                                                                    if len(v4) > 0 {
                                                                        v4[0] = "hello"
                                                                    }
                                                                    var v86 int64 = 42
                                                                    _ = v86
                                                                } else {
                                                                    switch -1 {
                                                                    case -9223372036854775808:
                                                                        v87 := GenS10{F0: 92, F1: (1e10 <= 3.14159)}
                                                                        _ = v87
                                                                    }
                                                                }
                                                            default:
                                                                v77 += v77
                                                            }
                                                            _ = fi82
                                                        }
                                                        if v6 {
                                                            fmt.Println(v76, v2, v77)
                                                            v88 := GenS3{F0: (v1 % ((2147483647) | 1)), F1: 1, F2: (v2 / ((-879) | 1))}
                                                            _ = v88
                                                        } else {
                                                            fmt.Println(v1, v1)
                                                            switch "foo bar" {
                                                            case "":
                                                                var v89 float32 = float32((-2.25 / 1e10))
                                                                _ = v89
                                                            case "\n\t":
                                                                var v90 bool = (false || (85 <= 42))
                                                                _ = v90
                                                            }
                                                        }
                                                        if v6 {
                                                            v4 = append(v4, "hello")
                                                            if len(v4) > 0 {
                                                                v4[0] = (((("foo bar" + "") + "foo bar") + "hello") + "")
                                                            }
                                                        }
                                                    }
                                                    switch "a" {
                                                    case "a":
                                                        for fi91 := 0; fi91 < 20; fi91++ {
                                                            v66 = v2
                                                            _ = fi91
                                                        }
                                                    }
                                                } else {
                                                    v3 = (v5 | fi75)
                                                }
                                            } else {
                                                var v92 int64 = (v66 % (((v2 * int64(v77))) | 1))
                                                _ = v92
                                            }
                                            for fi93 := 0; fi93 < 14; fi93++ {
                                                var v94 string = "a"
                                                _ = fi93
                                                _ = v94
                                            }
                                            for fi95 := 0; fi95 < 1; fi95++ {
                                                if (-2.25 > (((v76 + v76) * float64(float32((v76 / v76)))) - 0.0)) {
                                                    v96 := GenS12{F0: v77}
                                                    _ = v96
                                                }
                                                if v6 {
                                                    v5 = v5
                                                    switch "foo bar" {
                                                    case "a":
                                                        if ((float64(0.1) / (v76 * v76)) > (v76 / v76)) {
                                                            var v97 bool = v6
                                                            v4 = append(v4, "")
                                                            _ = v97
                                                        }
                                                    }
                                                } else {
                                                    v98 := GenS3{F0: 554, F1: 3, F2: (v66 % (((^(v66))) | 1))}
                                                    _ = v98
                                                }
                                                var v99 bool = v6
                                                _ = fi95
                                                _ = v99
                                            }
                                        }
                                    }
                                }
                                switch ("\n\t" + ("foo bar" + "foo bar")) {
                                case "a":
                                    v76 = v76
                                case "hello":
                                    v3 -= int(v66)
                                default:
                                    if (!(794 == 65535)) {
                                        v100 := GenS11{F0: v5, F1: (v77 >> (((v77 << ((1) & 63))) & 63))}
                                        for fi101 := 0; fi101 < 1; fi101++ {
                                            var v102 bool = (!v6)
                                            v103 := GenS5{F0: 4294967295}
                                            var v104 int = (^(v5))
                                            _ = fi101
                                            _ = v102
                                            _ = v103
                                            _ = v104
                                        }
                                        _ = v100
                                    }
                                }
                                var v105 int32 = 853
                                _ = v105
                            } else {
                                v77 -= (v77 / ((v77) | 1))
                                switch ("foo bar" + "foo bar") {
                                case "a":
                                    v106 := GenS13{F0: 32767, F1: (v1 % ((v1) | 1)), F2: 2}
                                    _ = v106
                                case "hello":
                                    fmt.Println(v5, v3, v2)
                                case "\n\t":
                                    _ = len(v4)
                                default:
                                    if (1 > 42) {
                                        if (2 <= 99) {
                                            var v107 int16 = 42
                                            _ = v107
                                        }
                                        if len(v4) > 0 {
                                            v4[0] = ("hello" + "\n\t")
                                        }
                                    } else {
                                        v108 := GenS11{F0: 9223372036854775807, F1: v77}
                                        _ = v108
                                    }
                                }
                            }
                            if false {
                                switch ("foo bar" + "a") {
                                case "a":
                                    var v109 float32 = ((float32(1.5) / 1.0) * 3.14159)
                                    _ = v109
                                case "\n\t":
                                    v110 := GenS14{F0: -128, F1: 919}
                                    _ = v110
                                case "hello":
                                    v1 -= -1
                                default:
                                    v1 += v1
                                }
                            } else {
                                v111 := GenS8{F0: (v77 * v77), F1: "a", F2: 1}
                                fmt.Println(v6, fi75)
                                _ = v111
                            }
                        }
                        _ = fi75
                        _ = v76
                        _ = v77
                    }
                } else {
                    _ = len(v4)
                    for fi112 := 0; fi112 < 8; fi112++ {
                        fmt.Println(v6)
                        if ("" < "") {
                            if v6 {
                                if v6 {
                                    var v113 float64 = -2.25
                                    if v6 {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v66 = v2
                                        fmt.Println(v1, v66)
                                    }
                                    var v114 bool = v6
                                    _ = v113
                                    _ = v114
                                } else {
                                    var v115 uint8 = 0
                                    var v116 int16 = -427
                                    _ = v115
                                    _ = v116
                                }
                                v2 += (v66 % ((9223372036854775807) | 1))
                            } else {
                                var v117 string = "\n\t"
                                _ = v117
                            }
                            v1 -= v1
                        } else {
                            v118 := GenS15{F0: 32767, F1: -746}
                            if len(v4) > 0 {
                                v4[0] = "a"
                            }
                            v66 = -9223372036854775808
                            _ = v118
                        }
                        _ = fi112
                    }
                    for fi119 := 0; fi119 < 18; fi119++ {
                        v66 = v2
                        _ = fi119
                    }
                }
                var v120 float32 = -2.25
                _ = v120
                _ = v73
            }
        } else {
            if (v2 >= int64(v1)) {
                fmt.Println(v3, v3)
            }
            var v121 int8 = 53
            v3 = (v3 / ((128) | 1))
            _ = v121
        }
    } else {
        switch (v5 ^ (v5 % ((948) | 1))) {
        case -9223372036854775808:
            v3 = (v3 | v5)
        case 42:
            v3 += v3
        case -2147483648:
            switch (v3 ^ v3) {
            case -200:
                switch (^(v3)) {
                case 905:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case -1:
                    for fi122 := 0; fi122 < 8; fi122++ {
                        for fi123 := 0; fi123 < 18; fi123++ {
                            var v124 uint32 = 0
                            _ = fi123
                            _ = v124
                        }
                        for fi125 := 0; fi125 < 1; fi125++ {
                            v126 := GenS16{F0: (v2 * (^(v66))), F1: 0}
                            _ = fi125
                            _ = v126
                        }
                        for fi127 := 0; fi127 < 11; fi127++ {
                            v4 = append(v4, (("a" + "foo bar") + "a"))
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            var v128 string = "a"
                            _ = fi127
                            _ = v128
                        }
                        _ = fi122
                    }
                case 127:
                    v3 = (v3 - 534)
                default:
                    switch "" {
                    case "\n\t":
                        var v129 uint64 = 4294967295
                        _ = v129
                    case "hello":
                        var v130 uint = 18446744073709551615
                        _ = v130
                    }
                }
            case 4294967295:
                var v131 int64 = (v2 + -9223372036854775808)
                _ = v131
            default:
                if (82 != 35) {
                    var v132 uint16 = 65535
                    if (!(-55 != 42)) {
                        for fi133 := 0; fi133 < 12; fi133++ {
                            if false {
                                if len(v4) > 0 {
                                    v4[0] = ""
                                }
                            }
                            _ = fi133
                        }
                        fmt.Println(v3, v132)
                        var v134 float32 = 0.0
                        _ = v134
                    }
                    v135 := GenS16{F0: v66, F1: (v2 | v2)}
                    _ = v132
                    _ = v135
                }
            }
        }
        if v6 {
            var v136 int8 = 100
            v1 -= v1
            v1 += v1
            _ = v136
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    fmt.Println(v2, v66)
    var v137 uint8 = 162
    var v138 int8 = 46
    v139 := GenS7{F0: (!false)}
    for fi140 := 0; fi140 < 18; fi140++ {
        v141 := GenS17{F0: 2, F1: (v137 / (((v137 - v137)) | 1)), F2: (v66 % ((v66) | 1))}
        if false {
            for fi142 := 0; fi142 < 3; fi142++ {
                var v143 int64 = v66
                if (-2147483648 <= v1) {
                    v144 := GenS5{F0: 1}
                    var v145 uint32 = 0
                    var v146 float64 = (((((1e10 + (0.0 / ((((0.0 + 0.0) - 0.1) / -2.25) - 3.14159))) * float64(-2.25)) * 3.14159) + 0.0) * (1e10 * (-2.25 / 0.1)))
                    v138 = v138
                    if false {
                        switch "hello" {
                        case "foo bar":
                            var v147 uint8 = (v137 * 104)
                            _ = v147
                        case "a":
                            switch (v3 / (((v5 >> ((4294967295) & 63))) | 1)) {
                            case -342:
                                var v148 int = (v3 / ((-2147483648) | 1))
                                _ = v148
                            case 32768:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case -1:
                                v1 = (v1 ^ (v1 ^ v1))
                            }
                        }
                        fmt.Println(v66, v146, v3)
                        var v149 uint8 = (v137 & v137)
                        v4 = append(v4, ("a" + "a"))
                        _ = v149
                    } else {
                        fmt.Println(v145, v5, v3)
                    }
                    _ = v144
                    _ = v145
                    _ = v146
                }
                var v150 float32 = ((3.14159 - (1e10 - ((((1.5 / (-2.25 + (((1.5 * (3.14159 + float32(3.14159))) * (1.5 / (0.1 + ((0.1 / (3.14159 + 1e10)) * 1.5)))) / 3.14159))) - 0.1) + (-2.25 + 0.0)) / -2.25))) * (-2.25 * -2.25))
                _ = fi142
                _ = v143
                _ = v150
            }
        } else {
            fmt.Println(v1)
            var v151 bool = (-906 != 0)
            _ = v151
        }
        _ = fi140
        _ = v141
    }
    switch v5 {
    case -386:
        var v152 bool = v6
        _ = v152
    default:
        switch "hello" {
        case "foo bar":
            var v153 uint8 = v137
            _ = v153
        }
    }
    v154 := GenS18{F0: -128, F1: false}
    if v6 {
        var v155 uint32 = 0
        fmt.Println(v155)
        _ = v155
    }
    var v156 int64 = -369
    fmt.Println("v1:", v1)
    fmt.Println("v137:", v137)
    fmt.Println("v138:", v138)
    fmt.Println("v156:", v156)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v137
    _ = v138
    _ = v139
    _ = v154
    _ = v156
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
}
