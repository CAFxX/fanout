package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 int; F2 string }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 int64; F1 int32; F2 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int }
type GenS7 struct { F0 bool; F1 int }
type GenS8 struct { F0 int8; F1 int64 }
type GenS9 struct { F0 bool; F1 uint32 }
type GenS10 struct { F0 int16; F1 uint }
type GenS11 struct { F0 int64; F1 int32 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 uint8; F1 uint16 }
type GenS14 struct { F0 bool; F1 int8; F2 string }
type GenS15 struct { F0 uint; F1 int; F2 int64 }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 uint8 }
type GenS18 struct { F0 int }

func main() {
    var v1 uint8 = 136
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]int)
    var v4 bool = true
    v5 := []uint32{}
    var v6 uint16 = 340
    var v7 uint64 = 42
    if (v4 && (!(false || (!false)))) {
        v1 = 0
        if v4 {
            var v8 bool = v4
            v1 += v1
            var v9 bool = false
            _ = v8
            _ = v9
        } else {
            for fi10 := 0; fi10 < 7; fi10++ {
                for fi11 := 0; fi11 < 10; fi11++ {
                    var v12 uint64 = v7
                    _ = fi11
                    _ = v12
                }
                _ = fi10
            }
            if ("a" > ("\n\t" + (("foo bar" + "hello") + ""))) {
                var v13 uint64 = v7
                switch ("" + "foo bar") {
                case "hello":
                    for fi14 := 0; fi14 < 5; fi14++ {
                        var v15 int = (-(fi14))
                        if ((v15 & (v15 ^ (fi14 | (fi14 | (v15 >> ((393) & 63)))))) <= v15) {
                            v6 -= (v6 >> ((0) & 63))
                        } else {
                            if (42 != -1) {
                                if v4 {
                                    fmt.Println(v6, v7)
                                } else {
                                    for fi16 := 0; fi16 < 19; fi16++ {
                                        switch "\n\t" {
                                        case "a":
                                            var v17 uint16 = uint16(v7)
                                            _ = v17
                                        case "":
                                            var v18 uint = 884
                                            _ = v18
                                        case "\n\t":
                                            switch v15 {
                                            case 32767:
                                                _ = len(v5)
                                            default:
                                                var v19 uint64 = (v7 & 42)
                                                _ = v19
                                            }
                                        default:
                                            v15 += fi14
                                        }
                                        var v20 uint8 = v1
                                        var v21 float32 = (0.1 * -2.25)
                                        fmt.Println(v1)
                                        _ = fi16
                                        _ = v20
                                        _ = v21
                                    }
                                    v15 -= (v15 + v15)
                                }
                                switch fi14 {
                                case 0:
                                    for fi22 := 0; fi22 < 1; fi22++ {
                                        if true {
                                            var v23 string = "\n\t"
                                            _ = v23
                                        } else {
                                            v15 = 4294967295
                                        }
                                        v4 = false
                                        _ = fi22
                                    }
                                case 2147483647:
                                    v15 = (fi14 | (fi14 - 32768))
                                default:
                                    v1 = v1
                                }
                                if (!(2147483648 <= 0)) {
                                    for fi24 := 0; fi24 < 17; fi24++ {
                                        switch ("\n\t" + "a") {
                                        case "hello":
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        case "a":
                                            v13 = (v13 * (v7 << ((0) & 63)))
                                            fmt.Println(v6, v13)
                                            switch 32767 {
                                            case 4294967295:
                                                delete(v3, 2)
                                            default:
                                                v7 = (v7 * 262)
                                            }
                                        }
                                        switch "foo bar" {
                                        case "":
                                            var v25 uint16 = v6
                                            _ = v25
                                        case "a":
                                            var v26 uint32 = 0
                                            _ = v26
                                        case "hello":
                                            if ((v6 - (v6 + v6)) <= 255) {
                                                v27 := GenS1{F0: uint32(v6)}
                                                fmt.Println(fi24, fi24, v1)
                                                _ = v27
                                            }
                                        default:
                                            var v28 uint64 = 289
                                            _ = v28
                                        }
                                        if true {
                                            if v4 {
                                                v5 = append(v5, 65535)
                                            } else {
                                                if (v4 || v4) {
                                                    if true {
                                                        var v29 uint16 = v6
                                                        v15 = (-(fi24))
                                                        if (!false) {
                                                            var v30 float32 = float32(-2.25)
                                                            _ = v30
                                                        } else {
                                                            fmt.Println(v29)
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                _ = len(v2)
                                                            default:
                                                                switch (fi24 >> ((65535) & 63)) {
                                                                case -471:
                                                                    var v31 float32 = (-2.25 + ((-2.25 / ((0.0 - (1e10 * (3.14159 + 0.1))) + 1e10)) + float32(0.1)))
                                                                    _ = v31
                                                                case 9223372036854775807:
                                                                    v1 -= (v1 / (((v1 ^ (v1 % ((v1) | 1)))) | 1))
                                                                default:
                                                                    switch (("a" + "") + "foo bar") {
                                                                    case "":
                                                                        var v32 int64 = 42
                                                                        _ = v32
                                                                    case "hello":
                                                                        switch (v15 - fi24) {
                                                                        case 42:
                                                                            var v33 float32 = float32(((1e10 - 1e10) * ((-2.25 * ((0.1 + (-2.25 / 1.0)) * float64(1e10))) - (float64(0.0) - (-2.25 + (1e10 - (3.14159 + 1e10)))))))
                                                                            _ = v33
                                                                        case -476:
                                                                            v6 = (^(v6))
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                            v13 = v13
                                                                            v29 = (v6 | 2)
                                                                            v6 = v29
                                                                        default:
                                                                            v1 -= v1
                                                                        }
                                                                    case "a":
                                                                        if (!false) {
                                                                            fmt.Println(v1, v15, v15)
                                                                            if v4 {
                                                                                if len(v5) >= 2 {
                                                                                    v5 = v5[1:len(v5)]
                                                                                }
                                                                            } else {
                                                                                v34 := GenS2{F0: 128, F1: fi14, F2: "hello"}
                                                                                _ = v34
                                                                            }
                                                                            if ((v1 & 194) != (v1 - (v1 ^ (v1 >> ((65535) & 63))))) {
                                                                                var v35 float64 = (((float64(3.14159) - -2.25) + ((-2.25 * float64(3.14159)) * float64((1e10 - 0.1)))) - 0.0)
                                                                                _ = v35
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            var v36 uint64 = 2
                                                            _ = v36
                                                        }
                                                        _ = v29
                                                    }
                                                    fmt.Println(v7, fi14)
                                                    if v4 {
                                                        var v37 float32 = float32((-2.25 / ((3.14159 - ((1e10 * float64(1e10)) - float64(1e10))) / 1.0)))
                                                        var v38 int64 = 574
                                                        v39 := GenS3{F0: 42, F1: v4}
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        _ = v37
                                                        _ = v38
                                                        _ = v39
                                                    }
                                                }
                                                var v40 int = (fi14 + (fi24 % ((127) | 1)))
                                                _ = v40
                                            }
                                            switch int(v6) {
                                            case -372:
                                                var v41 string = ""
                                                _ = v41
                                            default:
                                                var v42 float64 = 3.14159
                                                if (v4 || (!(2 <= v1))) {
                                                    if (-404 == 1) {
                                                        if true {
                                                            v43 := GenS1{F0: 2}
                                                            v2[2147483648] = ((0.0 >= 0.0) || false)
                                                            _ = v43
                                                        } else {
                                                            if (!true) {
                                                                v15 += fi14
                                                                switch (-(fi24)) {
                                                                case -156:
                                                                    fmt.Println(fi24)
                                                                }
                                                                if v4 {
                                                                    switch ("" + ("" + "")) {
                                                                    case "\n\t":
                                                                        var v44 uint32 = 4294967295
                                                                        _ = v44
                                                                    case "hello":
                                                                        var v45 int16 = -32768
                                                                        _ = v45
                                                                    case "a":
                                                                        v46 := GenS3{F0: 0, F1: (-1 >= -32768)}
                                                                        _ = v46
                                                                    default:
                                                                        delete(v3, 152)
                                                                    }
                                                                    v42 += v42
                                                                } else {
                                                                    if (v42 == ((v42 - v42) + 0.0)) {
                                                                        fmt.Println(v42, fi14)
                                                                        if (v4 && v4) {
                                                                            if (!("\n\t" >= ("" + ""))) {
                                                                                v4 = v4
                                                                                fmt.Println(v6, v6, fi24)
                                                                            }
                                                                            fmt.Println(v4)
                                                                            v47 := GenS1{F0: 4294967295}
                                                                            _ = v47
                                                                        }
                                                                        var v48 bool = v4
                                                                        _ = v48
                                                                    }
                                                                    switch (("hello" + ("\n\t" + ("\n\t" + "foo bar"))) + "") {
                                                                    case "\n\t":
                                                                        var v49 uint8 = (v1 ^ v1)
                                                                        _ = v49
                                                                    case "a":
                                                                        if (67 > -74) {
                                                                            if len(v5) > 0 {
                                                                                v5[0] = 2
                                                                            }
                                                                            if (-1 != -32768) {
                                                                                var v50 float64 = 0.1
                                                                                var v51 int8 = 0
                                                                                fmt.Println(v4, fi24)
                                                                                _ = v50
                                                                                _ = v51
                                                                            } else {
                                                                                v6 = 2
                                                                                if len(v5) >= 2 {
                                                                                    v5 = v5[1:len(v5)]
                                                                                }
                                                                            }
                                                                            v7 = v13
                                                                        }
                                                                    }
                                                                    fmt.Println(fi14, fi24, fi14)
                                                                }
                                                            }
                                                        }
                                                        var v52 int = -207
                                                        var v53 uint32 = 1
                                                        _ = v52
                                                        _ = v53
                                                    } else {
                                                        if (!(0 >= -4)) {
                                                            var v54 uint16 = (^(v6))
                                                            v55 := GenS4{F0: -9223372036854775808, F1: -689, F2: 42}
                                                            fmt.Println(fi24)
                                                            _ = v54
                                                            _ = v55
                                                        }
                                                    }
                                                    switch 4294967295 {
                                                    case 127:
                                                        var v56 uint64 = v7
                                                        _ = v56
                                                    case -1:
                                                        var v57 int = fi24
                                                        _ = v57
                                                    case -631:
                                                        v5 = append(v5, 617)
                                                    }
                                                } else {
                                                    if v4 {
                                                        v15 = (fi24 + -1)
                                                        v58 := GenS4{F0: 42, F1: -1, F2: 247}
                                                        _ = v58
                                                    } else {
                                                        if v4 {
                                                            _ = len(v3)
                                                            var v59 int8 = -128
                                                            var v60 int32 = 2147483647
                                                            _ = v59
                                                            _ = v60
                                                        }
                                                        fmt.Println(v15)
                                                    }
                                                    v61 := GenS4{F0: 42, F1: 594, F2: -162}
                                                    _ = v61
                                                }
                                                var v62 bool = v4
                                                _ = v42
                                                _ = v62
                                            }
                                        }
                                        _ = fi24
                                    }
                                } else {
                                    v13 = 1
                                }
                            } else {
                                var v63 string = "a"
                                _ = v63
                            }
                        }
                        _ = fi14
                        _ = v15
                    }
                case "a":
                    var v64 uint = 281
                    _ = v64
                default:
                    fmt.Println(v7, v4)
                }
                _ = v13
            } else {
                v4 = v4
                v6 -= 42
            }
            v65 := GenS5{F0: -1}
            _ = v65
        }
        switch 127 {
        case 127:
            for fi66 := 0; fi66 < 11; fi66++ {
                if (!v4) {
                    var v67 int = fi66
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v68 uint = 255
                    _ = v67
                    _ = v68
                } else {
                    fmt.Println(v4, v1)
                    for fi69 := 0; fi69 < 9; fi69++ {
                        v70 := GenS6{F0: -1}
                        _ = len(v2)
                        v71 := GenS1{F0: 898}
                        _ = fi69
                        _ = v70
                        _ = v71
                    }
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi66
            }
        default:
            switch (("a" + "foo bar") + ("foo bar" + ("hello" + ""))) {
            case "\n\t":
                var v72 uint32 = 4294967295
                _ = v72
            case "hello":
                v2[75] = false
            case "foo bar":
                v1 -= (^(v1))
            default:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        }
    }
    fmt.Println(v6, v6, v1)
    if v4 {
        v1 += (v1 >> ((65535) & 63))
        if len(v5) > 0 {
            v5[0] = 2
        }
    } else {
        switch "hello" {
        case "a":
            v7 -= (^(v7))
        default:
            for fi73 := 0; fi73 < 12; fi73++ {
                v74 := GenS5{F0: fi73}
                for fi75 := 0; fi75 < 2; fi75++ {
                    if false {
                        var v76 float32 = 1.5
                        if (((816 == 42) && v4) || v4) {
                            switch ("" + "\n\t") {
                            case "foo bar":
                                if (v76 <= 3.14159) {
                                    var v77 float64 = 1e10
                                    _ = v77
                                } else {
                                    if v4 {
                                        if (631 == 107) {
                                            _ = len(v5)
                                            v7 = (^(v7))
                                            switch fi73 {
                                            case -50:
                                                if ((v76 / (v76 - -2.25)) == v76) {
                                                    if v4 {
                                                        if v4 {
                                                            fmt.Println(v76, v1)
                                                        } else {
                                                            var v78 float32 = (v76 - v76)
                                                            v79 := GenS7{F0: false, F1: fi73}
                                                            fmt.Println(v76)
                                                            _ = v78
                                                            _ = v79
                                                        }
                                                        fmt.Println(fi73, v6)
                                                        if len(v5) > 0 {
                                                            v5[0] = 243
                                                        }
                                                        v1 += (v1 + (v1 % (((^(v1))) | 1)))
                                                    }
                                                    v80 := v3[42]
                                                    fmt.Println(v76)
                                                    _ = v80
                                                }
                                            case 127:
                                                var v81 float32 = float32((-2.25 - 0.1))
                                                _ = v81
                                            case 42:
                                                v82 := GenS8{F0: -128, F1: 1}
                                                _ = v82
                                            }
                                        }
                                        if v4 {
                                            v4 = (v4 || (42 <= (^(v7))))
                                            v76 += v76
                                        } else {
                                            fmt.Println(fi75, v7, fi75)
                                            if false {
                                                fmt.Println(v4)
                                                if (!false) {
                                                    if true {
                                                        if (4294967295 <= 4294967295) {
                                                            v76 += -2.25
                                                        }
                                                    } else {
                                                        switch -2147483648 {
                                                        case -9223372036854775808:
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                        case 127:
                                                            v83 := GenS9{F0: v4, F1: 2147483648}
                                                            _ = v83
                                                        }
                                                    }
                                                    if len(v5) > 0 {
                                                        v5[0] = 2
                                                    }
                                                    v84 := GenS5{F0: (fi75 >> ((198) & 63))}
                                                    _ = v84
                                                } else {
                                                    switch "a" {
                                                    case "\n\t":
                                                        var v85 int = fi73
                                                        _ = v85
                                                    case "":
                                                        fmt.Println(fi75)
                                                    default:
                                                        var v86 float64 = 3.14159
                                                        _ = v86
                                                    }
                                                }
                                            } else {
                                                var v87 int = fi75
                                                _ = v87
                                            }
                                            v6 = v6
                                            fmt.Println(v76)
                                        }
                                        if v4 {
                                            v88 := GenS3{F0: 940, F1: (0 < v6)}
                                            _ = v88
                                        }
                                    } else {
                                        v1 -= (v1 - v1)
                                        switch (("" + "a") + "a") {
                                        case "\n\t":
                                            fmt.Println(v7, v1)
                                        case "foo bar":
                                            var v89 int32 = 1
                                            _ = v89
                                        case "a":
                                            v7 = (v7 | 620)
                                            switch (fi73 << ((224) & 63)) {
                                            case 32768:
                                                v90 := GenS10{F0: -411, F1: 0}
                                                _ = v90
                                            }
                                        default:
                                            var v91 int16 = 32767
                                            _ = v91
                                        }
                                        _ = len(v5)
                                    }
                                }
                                var v92 int = (-(fi73))
                                var v93 float64 = (0.0 * 1.5)
                                _ = v92
                                _ = v93
                            default:
                                v5 = append(v5, 2147483647)
                            }
                            if (v4 && (-1 <= 1)) {
                                v7 += v7
                                if v4 {
                                    v7 = (v7 % ((4294967295) | 1))
                                    switch ((("\n\t" + "foo bar") + ("a" + "hello")) + "") {
                                    case "\n\t":
                                        v94 := GenS9{F0: (100 != 124), F1: 4294967295}
                                        v95 := GenS8{F0: -82, F1: 1}
                                        _ = v94
                                        _ = v95
                                    case "foo bar":
                                        if (!false) {
                                            v96 := GenS11{F0: -731, F1: 1}
                                            _ = v96
                                        } else {
                                            switch (("\n\t" + "hello") + "foo bar") {
                                            case "hello":
                                                if v4 {
                                                    v76 += v76
                                                    v4 = (-128 <= 63)
                                                    var v97 uint64 = v7
                                                    v98 := GenS1{F0: 2147483647}
                                                    _ = v97
                                                    _ = v98
                                                } else {
                                                    v99 := GenS12{F0: 100}
                                                    v100 := GenS12{F0: 1}
                                                    _ = v100
                                                    _ = v99
                                                }
                                            case "":
                                                v101 := GenS6{F0: (fi73 / (((fi75 | 601)) | 1))}
                                                _ = v101
                                            }
                                            v4 = false
                                        }
                                    case "":
                                        v1 = 127
                                    default:
                                        v102 := GenS10{F0: int16(v6), F1: 2}
                                        _ = v102
                                    }
                                    var v103 int16 = 1000
                                    fmt.Println(v6, v4, v76)
                                    _ = v103
                                }
                            } else {
                                var v104 bool = (v4 || (-1 == -6))
                                if v104 {
                                    if true {
                                        var v105 uint64 = v7
                                        switch "foo bar" {
                                        case "hello":
                                            fmt.Println(v7)
                                        case "\n\t":
                                            v3[2147483647] = fi75
                                        }
                                        v105 -= v105
                                        _ = v105
                                    }
                                    var v106 uint8 = v1
                                    var v107 uint8 = v106
                                    var v108 uint8 = (v1 % ((v106) | 1))
                                    if v104 {
                                        if v104 {
                                            fmt.Println(fi75)
                                            var v109 bool = ("\n\t" >= "")
                                            _ = v109
                                        }
                                        v7 = 2
                                    } else {
                                        var v110 int8 = 42
                                        var v111 float32 = ((((1e10 + 0.0) / ((v76 - v76) - v76)) - 1e10) * (v76 / (v76 + v76)))
                                        v4 = v104
                                        _ = v110
                                        _ = v111
                                    }
                                    v112 := v3[0]
                                    _ = v106
                                    _ = v107
                                    _ = v108
                                    _ = v112
                                } else {
                                    var v113 float64 = 0.0
                                    var v114 int8 = -95
                                    _ = len(v3)
                                    _ = v113
                                    _ = v114
                                }
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = v104
                            }
                            delete(v3, 4294967295)
                        }
                        var v115 string = ("hello" + "")
                        if (0 == 126) {
                            if v4 {
                                var v116 int64 = -1
                                _ = v116
                            } else {
                                if true {
                                    var v117 uint32 = 871
                                    _ = v117
                                }
                            }
                            v118 := GenS3{F0: 0, F1: true}
                            _ = v118
                        } else {
                            if v4 {
                                var v119 string = "foo bar"
                                v119 = v119
                                v120 := GenS11{F0: 9223372036854775807, F1: 1}
                                _ = v119
                                _ = v120
                            } else {
                                v1 = 255
                                v121 := GenS13{F0: 146, F1: (v6 | v6)}
                                v122 := GenS14{F0: v4, F1: -128, F2: ""}
                                _ = v121
                                _ = v122
                            }
                            v115 = v115
                            var v123 uint8 = (v1 >> ((42) & 63))
                            _ = v123
                        }
                        switch "foo bar" {
                        case "":
                            var v124 uint = 1
                            _ = v124
                        default:
                            if (int64(fi73) == -1) {
                                if v4 {
                                    var v125 int16 = 32767
                                    _ = v125
                                }
                                var v126 bool = v4
                                switch 128 {
                                case 114:
                                    v4 = (-32768 >= 32767)
                                }
                                _ = v126
                            } else {
                                switch fi73 {
                                case 1:
                                    v76 += v76
                                case 42:
                                    fmt.Println(v4)
                                    var v127 uint64 = v7
                                    if len(v5) > 0 {
                                        v5[0] = 641
                                    }
                                    if v4 {
                                        _ = len(v5)
                                        _ = len(v5)
                                    } else {
                                        var v128 bool = (v4 || v4)
                                        _ = len(v5)
                                        _ = v128
                                    }
                                    _ = v127
                                }
                                switch v115 {
                                case "":
                                    _ = len(v2)
                                case "\n\t":
                                    v1 = (^(v1))
                                }
                            }
                        }
                        _ = v115
                        _ = v76
                    }
                    if v4 {
                        v129 := GenS15{F0: 2, F1: int(v6), F2: 0}
                        _ = v129
                    }
                    v130 := GenS4{F0: 584, F1: 511, F2: 42}
                    _ = fi75
                    _ = v130
                }
                if v4 {
                    for fi131 := 0; fi131 < 6; fi131++ {
                        switch "a" {
                        case "\n\t":
                            var v132 int16 = 1
                            _ = v132
                        case "hello":
                            v133 := GenS16{F0: v7}
                            var v134 uint64 = v7
                            v6 = 32767
                            if v4 {
                                v6 = 1
                            } else {
                                _ = len(v5)
                                fmt.Println(fi131, v7)
                            }
                            fmt.Println(v1)
                            _ = v133
                            _ = v134
                        case "":
                            v7 = 1
                        default:
                            switch fi131 {
                            case 100:
                                switch "\n\t" {
                                case "foo bar":
                                    var v135 bool = (v6 < v6)
                                    var v136 uint64 = (v7 - (v7 - uint64(v6)))
                                    _ = v135
                                    _ = v136
                                case "":
                                    v4 = false
                                }
                            case 407:
                                v3[17] = fi73
                            default:
                                v1 = (v1 % ((v1) | 1))
                                var v137 uint16 = v6
                                switch fi73 {
                                case -2147483648:
                                    var v138 int64 = -1
                                    _ = v138
                                case 4294967295:
                                    v137 = (v6 << ((18446744073709551615) & 63))
                                case 9223372036854775807:
                                    delete(v3, 294)
                                }
                                _ = v137
                            }
                        }
                        _ = fi131
                    }
                } else {
                    v1 += v1
                    v139 := GenS11{F0: -1, F1: 981}
                    _ = v139
                }
                _ = fi73
                _ = v74
            }
        }
    }
    if v4 {
        switch "foo bar" {
        case "\n\t":
            v1 -= v1
        default:
            switch -546 {
            case -341:
                _ = len(v5)
            case -9223372036854775808:
                fmt.Println(v7)
            case -173:
                v140 := GenS4{F0: 42, F1: 0, F2: 566}
                _ = v140
            }
        }
        v141 := GenS17{F0: (v1 | (^(v1)))}
        v142 := GenS18{F0: -9223372036854775808}
        _ = v141
        _ = v142
    }
    v143 := GenS15{F0: 65535, F1: -998, F2: 99}
    _ = len(v3)
    _ = len(v5)
    var v144 uint16 = v6
    delete(v3, 514)
    switch "hello" {
    case "":
        v145 := GenS18{F0: -365}
        _ = v145
    case "foo bar":
        v6 -= v144
    case "a":
        var v146 bool = (("a" + "foo bar") > "")
        _ = v146
    default:
        for fi147 := 0; fi147 < 20; fi147++ {
            var v148 uint16 = uint16(v1)
            for fi149 := 0; fi149 < 18; fi149++ {
                v150 := GenS5{F0: 9223372036854775807}
                if false {
                    var v151 int64 = int64(v7)
                    _ = v151
                } else {
                    switch "foo bar" {
                    case "":
                        fmt.Println(fi149)
                    case "foo bar":
                        if (v4 || v4) {
                            var v152 uint64 = 958
                            v153 := GenS16{F0: (v7 * v152)}
                            v154 := GenS15{F0: 2, F1: (fi147 / (((^(fi149))) | 1)), F2: -9223372036854775808}
                            _ = v152
                            _ = v153
                            _ = v154
                        } else {
                            v7 = 18446744073709551615
                        }
                    case "a":
                        v155 := GenS14{F0: v4, F1: 42, F2: "\n\t"}
                        _ = v155
                    default:
                        var v156 uint8 = v1
                        var v157 int64 = -1
                        _ = v156
                        _ = v157
                    }
                    if v4 {
                        _ = len(v2)
                        var v158 int64 = 42
                        _ = v158
                    } else {
                        v7 -= 1
                    }
                }
                _ = fi149
                _ = v150
            }
            _ = fi147
            _ = v148
        }
    }
    var v159 float32 = float32(-2.25)
    var v160 string = ("\n\t" + ("\n\t" + (("foo bar" + ("" + "\n\t")) + "\n\t")))
    for fi161 := 0; fi161 < 18; fi161++ {
        _ = len(v2)
        if v4 {
            var v162 float32 = v159
            _ = v162
        }
        switch (fi161 >> ((65535) & 63)) {
        case 42:
            v3[849] = fi161
        }
        _ = fi161
    }
    var v163 string = v160
    fmt.Println("v1:", v1)
    fmt.Println("v144:", v144)
    fmt.Println("v159:", v159)
    fmt.Println("v160:", v160)
    fmt.Println("v163:", v163)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v143
    _ = v144
    _ = v159
    _ = v160
    _ = v163
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
