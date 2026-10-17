package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 uint64; F1 uint64; F2 uint }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 string; F1 int8; F2 int64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 bool; F1 uint32; F2 bool }
type GenS9 struct { F0 uint8; F1 string }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint; F1 int32 }
type GenS12 struct { F0 uint16; F1 int64 }
type GenS13 struct { F0 int32; F1 uint16; F2 int32 }
type GenS14 struct { F0 uint64; F1 uint; F2 int16 }
type GenS15 struct { F0 uint64; F1 string }
type GenS16 struct { F0 int16; F1 string }
type GenS17 struct { F0 uint64 }
type GenS18 struct { F0 int64 }
func genFn0(p0 uint32, p1 int32) int8 {
    p0 = (^(p0))
    p0 += (^(p0))
    p1 += p1
    p1 = (p1 * p1)
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 uint32 = 2
    v2 := []string{"", "\n\t", "\n\t", "foo bar"}
    v3 := []uint32{146, v1, v1}
    v4 := make(map[string]string)
    v5 := []uint32{}
    v6 := GenS1{F0: 145, F1: 1000, F2: 4294967295}
    switch "" {
    case "\n\t":
        v1 -= v1
    }
    if (("foo bar" != "hello") || true) {
        v1 -= v1
    } else {
        var v7 int32 = 2147483647
        var v8 int16 = 32767
        switch "" {
        case "foo bar":
            fmt.Println(v1, v8, v1)
        case "":
            for fi9 := 0; fi9 < 13; fi9++ {
                v10 := GenS2{F0: v1, F1: 928}
                fmt.Println(fi9, v1, fi9)
                switch (fi9 - (-(fi9))) {
                case 0:
                    v1 = (v1 * (v1 & v1))
                case 127:
                    v7 = (v7 / ((v7) | 1))
                case 42:
                    v7 -= v7
                default:
                    v11 := GenS3{F0: 42, F1: 2, F2: 65535}
                    _ = v11
                }
                _ = fi9
                _ = v10
            }
        case "\n\t":
            if (2 == 2) {
                fmt.Println(v1)
                for fi12 := 0; fi12 < 14; fi12++ {
                    v8 = v8
                    _ = fi12
                }
                if true {
                    var v13 uint = 0
                    for fi14 := 0; fi14 < 15; fi14++ {
                        for fi15 := 0; fi15 < 10; fi15++ {
                            var v16 uint = (v13 % (((v13 >> (((v13 >> (((v13 ^ (v13 << ((v13) & 63)))) & 63))) & 63))) | 1))
                            var v17 bool = (!true)
                            if v17 {
                                switch "foo bar" {
                                case "\n\t":
                                    v18 := GenS2{F0: (v1 / (((v1 ^ 2147483647)) | 1)), F1: v13}
                                    _ = v18
                                case "hello":
                                    v4[("hello" + "hello")] = "a"
                                    fmt.Println(fi14, v7, v8)
                                case "foo bar":
                                    if ((42 < 147) && false) {
                                        v19 := GenS1{F0: 2, F1: v8, F2: 4294967295}
                                        switch "hello" {
                                        case "foo bar":
                                            if (!true) {
                                                v17 = (42 < 1)
                                            } else {
                                                if false {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    fmt.Println(v16, fi14, v17)
                                                    v20 := GenS4{F0: (v1 % ((v1) | 1)), F1: -1}
                                                    _ = v20
                                                }
                                                var v21 bool = (!v17)
                                                v1 -= v1
                                                _ = v21
                                            }
                                        case "hello":
                                            switch -259 {
                                            case -355:
                                                if (v17 && true) {
                                                    if (true && (-35 != 50)) {
                                                        v22 := GenS4{F0: (^(v1)), F1: 42}
                                                        _ = v22
                                                    } else {
                                                        var v23 int16 = (v8 << ((v13) & 63))
                                                        if (v17 && v17) {
                                                            var v24 bool = (v17 || (590 != 0))
                                                            _ = v24
                                                        } else {
                                                            v25 := GenS2{F0: 42, F1: (v13 & (v13 >> (((v16 << (((v13 / ((v16) | 1))) & 63))) & 63)))}
                                                            switch "foo bar" {
                                                            case "":
                                                                v23 -= v8
                                                                var v26 uint16 = uint16(v16)
                                                                v27 := GenS1{F0: 148, F1: v8, F2: 397}
                                                                fmt.Println(fi14, fi14, fi14)
                                                                _ = v26
                                                                _ = v27
                                                            default:
                                                                if v17 {
                                                                    fmt.Println(v23, v13, v7)
                                                                } else {
                                                                    fmt.Println(fi14, v23, v13)
                                                                }
                                                            }
                                                            _ = v25
                                                        }
                                                        var v28 int32 = int32(fi15)
                                                        if v17 {
                                                            v29 := GenS5{F0: (("" + (("" + ("\n\t" + "hello")) + "\n\t")) + ((("a" + (("a" + "") + "a")) + "\n\t") + "hello")), F1: -123, F2: 0}
                                                            _ = v29
                                                        } else {
                                                            var v30 bool = (fi15 != (-(fi14)))
                                                            var v31 uint = v16
                                                            v8 = (v8 << (((v16 * v31)) & 63))
                                                            fmt.Println(v30)
                                                            v23 = v23
                                                            _ = v30
                                                            _ = v31
                                                        }
                                                        _ = v23
                                                        _ = v28
                                                    }
                                                } else {
                                                    v1 -= v1
                                                    fmt.Println(v17, fi14, v16)
                                                    v7 -= v7
                                                }
                                            }
                                        }
                                        _ = v19
                                    } else {
                                        v32 := GenS3{F0: 1, F1: 1, F2: (v16 ^ 42)}
                                        _ = v32
                                    }
                                }
                                v7 = (v7 >> ((v13) & 63))
                                switch fi14 {
                                case 32767:
                                    var v33 int8 = 42
                                    _ = v33
                                case -958:
                                    var v34 float32 = (float32(3.14159) * ((0.0 * (3.14159 + (1e10 / 3.14159))) * (0.0 - (1e10 / -2.25))))
                                    _ = v34
                                default:
                                    if (v17 || v17) {
                                        v35 := GenS6{F0: 42}
                                        var v36 string = "a"
                                        switch 0 {
                                        case -641:
                                            switch v36 {
                                            case "foo bar":
                                                v37 := GenS7{F0: v17}
                                                _ = v37
                                            }
                                            v16 -= v16
                                            v38 := GenS7{F0: v17}
                                            _ = v38
                                        default:
                                            v8 -= 1000
                                        }
                                        _ = v35
                                        _ = v36
                                    }
                                }
                            } else {
                                v39 := GenS8{F0: v17, F1: v1, F2: v17}
                                _ = v39
                            }
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                        _ = fi14
                    }
                    _ = v13
                } else {
                    if (true && (100 < -128)) {
                        for fi40 := 0; fi40 < 20; fi40++ {
                            if false {
                                switch (fi40 - (fi40 % ((fi40) | 1))) {
                                case 0:
                                    v41 := v4["a"]
                                    _ = v41
                                case 127:
                                    v42 := GenS4{F0: (^(v1)), F1: 9223372036854775807}
                                    _ = v42
                                case 2147483647:
                                    v43 := v4[""]
                                    _ = v43
                                }
                                var v44 float64 = 1e10
                                _ = v44
                            }
                            fmt.Println(v7, fi40, v1)
                            var v45 bool = true
                            _ = fi40
                            _ = v45
                        }
                        if (true || false) {
                            v7 += v7
                        }
                    } else {
                        var v46 int32 = 2147483647
                        if true {
                            v47 := GenS2{F0: (v1 << ((65535) & 63)), F1: 42}
                            for fi48 := 0; fi48 < 9; fi48++ {
                                var v49 float64 = -2.25
                                _ = fi48
                                _ = v49
                            }
                            _ = v47
                        } else {
                            switch int(v1) {
                            case -129:
                                var v50 uint64 = 42
                                _ = v50
                            }
                            switch "hello" {
                            case "a":
                                fmt.Println(v8, v8)
                            case "":
                                for fi51 := 0; fi51 < 11; fi51++ {
                                    v52 := GenS9{F0: 121, F1: "foo bar"}
                                    var v53 float64 = 0.0
                                    if (!true) {
                                        var v54 uint32 = 0
                                        _ = v54
                                    } else {
                                        v8 += (v8 - (v8 ^ (v8 | v8)))
                                        for fi55 := 0; fi55 < 3; fi55++ {
                                            var v56 uint64 = 1
                                            switch fi51 {
                                            case -1:
                                                if true {
                                                    v57 := GenS10{F0: true}
                                                    _ = v57
                                                } else {
                                                    fmt.Println(v46, v46, v8)
                                                    _ = len(v3)
                                                    var v58 float32 = 0.0
                                                    _ = v58
                                                }
                                            }
                                            _ = fi55
                                            _ = v56
                                        }
                                    }
                                    _ = fi51
                                    _ = v52
                                    _ = v53
                                }
                            }
                            v59 := GenS11{F0: 65535, F1: 0}
                            _ = v59
                        }
                        var v60 float64 = -2.25
                        _ = v46
                        _ = v60
                    }
                    if false {
                        if true {
                            v7 += (^(v7))
                        } else {
                            if false {
                                if ("" == ("hello" + "hello")) {
                                    if false {
                                        var v61 uint = 65535
                                        for fi62 := 0; fi62 < 6; fi62++ {
                                            v1 = (^(v1))
                                            fmt.Println(v8, v7)
                                            _ = fi62
                                        }
                                        _ = v61
                                    }
                                    if false {
                                        fmt.Println(v1, v7, v7)
                                    } else {
                                        v8 = v8
                                        if false {
                                            v7 += 2147483647
                                            v63 := GenS10{F0: false}
                                            _ = v63
                                        }
                                        v64 := GenS5{F0: "a", F1: 100, F2: -9223372036854775808}
                                        _ = v64
                                    }
                                    v4[""] = "a"
                                } else {
                                    fmt.Println(v1, v8)
                                }
                            }
                        }
                        if (-9223372036854775808 < 42) {
                            for fi65 := 0; fi65 < 12; fi65++ {
                                v66 := GenS12{F0: 392, F1: -1}
                                switch -2147483648 {
                                case 32768:
                                    v1 = v1
                                }
                                v5 = append(v5, v1)
                                _ = fi65
                                _ = v66
                            }
                            for fi67 := 0; fi67 < 2; fi67++ {
                                for fi68 := 0; fi68 < 9; fi68++ {
                                    v69 := GenS13{F0: (v7 + (^(v7))), F1: 65535, F2: (v7 << ((483) & 63))}
                                    var v70 string = "\n\t"
                                    if (0.0 != 1.5) {
                                        if true {
                                            v71 := GenS11{F0: 0, F1: (v7 / (((^(v7))) | 1))}
                                            v70 = "\n\t"
                                            _ = v71
                                        }
                                    } else {
                                        var v72 float32 = 1.5
                                        v1 = v1
                                        fmt.Println(v72, fi67, v72)
                                        _ = v72
                                    }
                                    _ = fi68
                                    _ = v69
                                    _ = v70
                                }
                                _ = fi67
                            }
                            var v73 int32 = 2147483647
                            _ = v73
                        }
                    }
                }
            }
        default:
            for fi74 := 0; fi74 < 6; fi74++ {
                _ = len(v4)
                v7 += v7
                _ = fi74
            }
        }
        _ = v7
        _ = v8
    }
    if (2147483647 != -9223372036854775808) {
        fmt.Println(v1)
        v75 := GenS3{F0: 42, F1: 1, F2: 60}
        for fi76 := 0; fi76 < 5; fi76++ {
            for fi77 := 0; fi77 < 7; fi77++ {
                if len(v3) > 0 {
                    v3[0] = (v1 | (v1 ^ (v1 ^ (v1 % ((v1) | 1)))))
                }
                v78 := GenS3{F0: 810, F1: 4294967295, F2: 51}
                v1 -= 4294967295
                _ = fi77
                _ = v78
            }
            _ = fi76
        }
        _ = v75
    }
    var v79 int = -9223372036854775808
    for fi80 := 0; fi80 < 20; fi80++ {
        v81 := GenS2{F0: (v1 | v1), F1: 18446744073709551615}
        _ = fi80
        _ = v81
    }
    for fi82 := 0; fi82 < 12; fi82++ {
        var v83 string = "hello"
        _ = fi82
        _ = v83
    }
    switch (v79 - (-(v79))) {
    case -583:
        if false {
            var v84 float64 = (((-2.25 - 1e10) - ((0.1 + ((3.14159 + 3.14159) - 1e10)) / (((((1e10 - ((3.14159 * 1e10) - 1.5)) * 0.1) + (1.5 - 1e10)) + ((1e10 + -2.25) / (1e10 - float64(0.0)))) * (float64((1e10 + -2.25)) * 0.1)))) / 3.14159)
            for fi85 := 0; fi85 < 16; fi85++ {
                var v86 int64 = 9223372036854775807
                for fi87 := 0; fi87 < 13; fi87++ {
                    v84 = v84
                    _ = fi87
                }
                for fi88 := 0; fi88 < 17; fi88++ {
                    fmt.Println(v86, v1, v79)
                    v89 := v4["a"]
                    fmt.Println(v84, v89, v79)
                    _ = fi88
                    _ = v89
                }
                _ = fi85
                _ = v86
            }
            _ = v84
        } else {
            fmt.Println(v79, v79)
            v90 := GenS5{F0: "\n\t", F1: -12, F2: -713}
            _ = v90
        }
    case 42:
        switch (("" + (((("" + "\n\t") + "a") + "") + "a")) + "hello") {
        case "hello":
            v91 := GenS2{F0: (v1 | 1), F1: 772}
            _ = v91
        case "\n\t":
            if true {
                v2 = append(v2, "a")
                v92 := GenS4{F0: (^(v1)), F1: 9223372036854775807}
                v93 := GenS14{F0: 18446744073709551615, F1: 1, F2: 779}
                _ = v92
                _ = v93
            }
        case "a":
            fmt.Println(v1, v1)
        default:
            fmt.Println(v1, v79)
        }
    case -433:
        if (!(-9223372036854775808 >= 229)) {
            v94 := GenS15{F0: 4294967295, F1: ("\n\t" + "hello")}
            _ = len(v4)
            _ = v94
        } else {
            if true {
                v95 := GenS12{F0: 65535, F1: -1}
                switch "a" {
                case "foo bar":
                    v96 := GenS8{F0: true, F1: v1, F2: false}
                    _ = v96
                }
                v79 = v79
                _ = v95
            }
            if false {
                v1 += v1
            } else {
                if (597 != 66) {
                    for fi97 := 0; fi97 < 20; fi97++ {
                        fmt.Println(v79, fi97, fi97)
                        var v98 uint = 324
                        _ = fi97
                        _ = v98
                    }
                    v1 = 65535
                    fmt.Println(v1, v79)
                } else {
                    if (0 < 579) {
                        switch v79 {
                        case -9223372036854775808:
                            if len(v3) > 0 {
                                v3[0] = 65535
                            }
                        case -554:
                            var v99 string = "hello"
                            _ = v99
                        default:
                            var v100 int16 = 42
                            _ = v100
                        }
                        v79 += v79
                        if false {
                            switch 645 {
                            case -641:
                                _ = len(v5)
                            case 9223372036854775807:
                                for fi101 := 0; fi101 < 12; fi101++ {
                                    for fi102 := 0; fi102 < 5; fi102++ {
                                        if ((2 != 128) || true) {
                                            v4["\n\t"] = ("a" + "a")
                                            switch "" {
                                            case "a":
                                                if (82 != 2) {
                                                    v103 := v4["\n\t"]
                                                    switch v103 {
                                                    case "\n\t":
                                                        v79 -= 128
                                                    case "a":
                                                        var v104 int64 = -1
                                                        _ = v104
                                                    case "foo bar":
                                                        v105 := GenS11{F0: 417, F1: 0}
                                                        _ = v105
                                                    }
                                                    v106 := v4[((v103 + (v103 + ("hello" + "foo bar"))) + v103)]
                                                    var v107 uint16 = 255
                                                    _ = v103
                                                    _ = v106
                                                    _ = v107
                                                }
                                            default:
                                                fmt.Println(fi102, v1)
                                            }
                                            v79 = -292
                                        } else {
                                            switch (fi101 >> ((0) & 63)) {
                                            case 32768:
                                                switch "foo bar" {
                                                case "\n\t":
                                                    switch (v79 - (-(v79))) {
                                                    case 128:
                                                        var v108 string = "a"
                                                        _ = v108
                                                    }
                                                }
                                            case 128:
                                                v109 := GenS2{F0: (^(v1)), F1: 2}
                                                _ = v109
                                            case 42:
                                                var v110 int = (fi101 | fi101)
                                                _ = v110
                                            }
                                            fmt.Println(fi102)
                                            if ((v1 << ((2) & 63)) <= v1) {
                                                v111 := GenS3{F0: 2, F1: 123, F2: 1}
                                                if (!((!(981 >= 18446744073709551615)) || false)) {
                                                    v112 := GenS8{F0: false, F1: (v1 | 65535), F2: false}
                                                    fmt.Println(fi102)
                                                    switch (v79 - (v79 ^ v79)) {
                                                    case 127:
                                                        fmt.Println(fi102, fi101, v79)
                                                    case -361:
                                                        if true {
                                                            if true {
                                                                var v113 int = (fi102 % ((32767) | 1))
                                                                fmt.Println(v1, fi102)
                                                                _ = v113
                                                            }
                                                            if false {
                                                                fmt.Println(fi101, fi102)
                                                            } else {
                                                                v1 = (v1 + v1)
                                                                v114 := GenS5{F0: ("foo bar" + ""), F1: int8(v1), F2: 42}
                                                                switch -934 {
                                                                case 1:
                                                                    var v115 int64 = 9223372036854775807
                                                                    _ = v115
                                                                }
                                                                _ = v114
                                                            }
                                                            v116 := GenS11{F0: 255, F1: -2147483648}
                                                            _ = v116
                                                        }
                                                        var v117 uint64 = 382
                                                        _ = v117
                                                    }
                                                    _ = v112
                                                }
                                                if (2 >= 127) {
                                                    v118 := GenS8{F0: (!false), F1: v1, F2: false}
                                                    _ = v118
                                                } else {
                                                    v79 = fi101
                                                    _ = len(v4)
                                                }
                                                _ = v111
                                            }
                                        }
                                        _ = fi102
                                    }
                                    _ = fi101
                                }
                            case 128:
                                v79 -= (v79 & (v79 >> ((65535) & 63)))
                            }
                        }
                    }
                }
                for fi119 := 0; fi119 < 19; fi119++ {
                    var v120 int = v79
                    if (1 >= 127) {
                        var v121 bool = (!(!(18446744073709551615 < 683)))
                        _ = v121
                    } else {
                        v4[""] = "\n\t"
                        v79 -= (v79 + 224)
                    }
                    v122 := GenS16{F0: 979, F1: "a"}
                    _ = fi119
                    _ = v120
                    _ = v122
                }
                switch v79 {
                case 128:
                    if ((!false) && false) {
                        if ((!(255 == 128)) && ((^(v1)) > (v1 | v1))) {
                            for fi123 := 0; fi123 < 9; fi123++ {
                                for fi124 := 0; fi124 < 2; fi124++ {
                                    v2 = append(v2, "hello")
                                    _ = fi124
                                }
                                if (4294967295 == 4294967295) {
                                    fmt.Println(fi123, fi123, fi123)
                                } else {
                                    fmt.Println(v79, fi123, v1)
                                    v79 -= v79
                                    v125 := GenS17{F0: 152}
                                    _ = v125
                                }
                                _ = fi123
                            }
                        }
                        if ((!(((v1 & (v1 ^ v1)) != v1) || false)) || (942 < 844)) {
                            switch ("a" + "") {
                            case "a":
                                fmt.Println(v1, v1)
                            case "foo bar":
                                var v126 int16 = 1
                                _ = v126
                            default:
                                v79 -= (v79 & 32768)
                            }
                            if (!true) {
                                _ = len(v5)
                                v79 -= (v79 - v79)
                            } else {
                                v79 = v79
                                fmt.Println(v79, v79)
                                v127 := GenS9{F0: 255, F1: "\n\t"}
                                _ = v127
                            }
                            var v128 float64 = -2.25
                            _ = v128
                        }
                    }
                case 9223372036854775807:
                    v5 = append(v5, 1)
                case 0:
                    v129 := GenS4{F0: 0, F1: 0}
                    _ = v129
                default:
                    var v130 uint8 = 45
                    _ = v130
                }
            }
        }
    default:
        if ((v79 - (v79 + (v79 | v79))) < v79) {
            if ((!(false && false)) && false) {
                for fi131 := 0; fi131 < 5; fi131++ {
                    v132 := GenS18{F0: -1}
                    for fi133 := 0; fi133 < 7; fi133++ {
                        delete(v4, "")
                        _ = len(v4)
                        _ = fi133
                    }
                    fmt.Println(v1, fi131)
                    _ = fi131
                    _ = v132
                }
            }
            v79 = v79
            var v134 float64 = (-2.25 + -2.25)
            _ = v134
        } else {
            for fi135 := 0; fi135 < 14; fi135++ {
                if true {
                    var v136 uint = 1
                    _ = v136
                } else {
                    delete(v4, ("hello" + "hello"))
                    var v137 float32 = 0.1
                    v1 += v1
                    _ = v137
                }
                _ = fi135
            }
            var v138 uint32 = 119
            fmt.Println(v79)
            _ = v138
        }
    }
    v79 -= v79
    v1 = (v1 ^ 4294967295)
    if true {
        v79 -= (v79 << ((842) & 63))
    } else {
        for fi139 := 0; fi139 < 16; fi139++ {
            if len(v5) > 0 {
                v5[0] = (v1 * 42)
            }
            var v140 float32 = 0.0
            _ = fi139
            _ = v140
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v79
}
