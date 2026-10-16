package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8; F1 int8; F2 int64 }
type GenS5 struct { F0 uint16; F1 int32; F2 string }
type GenS6 struct { F0 uint16; F1 bool }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 bool; F1 uint16 }
type GenS9 struct { F0 string; F1 uint; F2 int16 }
func genFn0() string {
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 int16 = -197
        _ = fi5
        _ = v6
    }
    v7 := []int{-377, -445}
    for fi8 := 0; fi8 < 15; fi8++ {
        fmt.Println(fi8)
        if false {
            if false {
                var v9 float64 = 3.14159
                v9 += 0.0
                _ = v9
            }
            var v10 uint = 1
            _ = v10
        } else {
            var v11 int64 = 790
            var v12 int16 = -901
            var v13 uint16 = 32767
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi8
    }
    _ = v7
    return "\n\t"
}

func main() {
    var v1 uint16 = 32767
    var v2 uint = 255
    v3 := make(map[int]bool)
    var v4 uint64 = 227
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 uint32 = 0
        _ = fi14
        _ = v15
    }
    v1 = (v1 ^ 42)
    v1 = (^(v1))
    for fi16 := 0; fi16 < 4; fi16++ {
        fmt.Println(v4, v1)
        switch fi16 {
        case 32767:
            for fi17 := 0; fi17 < 4; fi17++ {
                switch fi17 {
                case -9223372036854775808:
                    v18 := GenS1{F0: false, F1: 42}
                    _ = v18
                case 0:
                    if (127 <= int8(fi17)) {
                        var v19 uint = (v2 + 527)
                        _ = v19
                    }
                default:
                    v2 = 1
                }
                _ = fi17
            }
        case 2147483647:
            var v20 int16 = 32767
            _ = v20
        case -348:
            fmt.Println(fi16, v1)
        default:
            var v21 string = ""
            _ = v21
        }
        _ = fi16
    }
    v1 -= uint16(v2)
    v22 := []int{-875, 647, 9223372036854775807, 4294967295}
    if ((^(v1)) == (^(v1))) {
        switch 32768 {
        case 266:
            var v23 int = 9223372036854775807
            _ = v23
        case -269:
            for fi24 := 0; fi24 < 20; fi24++ {
                v25 := GenS1{F0: false, F1: -701}
                v1 -= v1
                _ = fi24
                _ = v25
            }
        default:
            fmt.Println(v1, v1, v1)
        }
        for fi26 := 0; fi26 < 3; fi26++ {
            v27 := GenS1{F0: false, F1: -9223372036854775808}
            _ = fi26
            _ = v27
        }
    } else {
        v28 := GenS2{F0: ("\n\t" + "foo bar")}
        v29 := GenS2{F0: ""}
        for fi30 := 0; fi30 < 2; fi30++ {
            switch "\n\t" {
            case "hello":
                var v31 string = "foo bar"
                _ = v31
            case "\n\t":
                v4 -= 1
            }
            _ = fi30
        }
        _ = v28
        _ = v29
    }
    v32 := v3[2147483647]
    if (42 > 0) {
        delete(v3, 756)
        var v33 uint = v2
        switch "" {
        case "a":
            v34 := GenS1{F0: false, F1: -9223372036854775808}
            _ = v34
        default:
            var v35 float64 = (float64((0.0 * 1e10)) - float64(float32(1e10)))
            _ = v35
        }
        _ = v33
    } else {
        var v36 float64 = 0.0
        for fi37 := 0; fi37 < 1; fi37++ {
            delete(v3, 2147483647)
            if (!v32) {
                v38 := GenS2{F0: "foo bar"}
                v3[713] = v32
                switch fi37 {
                case -1:
                    var v39 int32 = 2147483647
                    _ = v39
                }
                _ = v38
            }
            _ = fi37
        }
        for fi40 := 0; fi40 < 19; fi40++ {
            for fi41 := 0; fi41 < 9; fi41++ {
                switch (fi40 + fi41) {
                case 229:
                    v42 := GenS1{F0: (42 < -54), F1: 0}
                    _ = v42
                }
                v2 -= (^(v2))
                _ = fi41
            }
            v36 = v36
            _ = fi40
        }
        _ = v36
    }
    if v32 {
        fmt.Println(v2)
        _ = len(v3)
        var v43 int8 = 117
        _ = v43
    } else {
        switch 2147483647 {
        case 4294967295:
            if true {
                for fi44 := 0; fi44 < 12; fi44++ {
                    v4 = v4
                    delete(v3, (fi44 & (fi44 + 32768)))
                    _ = fi44
                }
                for fi45 := 0; fi45 < 13; fi45++ {
                    var v46 uint16 = 219
                    fmt.Println(v46, v32, fi45)
                    _ = fi45
                    _ = v46
                }
                if (-40 != int(v2)) {
                    fmt.Println(v32)
                    v1 = 2
                } else {
                    switch "foo bar" {
                    case "foo bar":
                        var v47 uint16 = (v1 & (^(v1)))
                        _ = v47
                    case "a":
                        v32 = v32
                    default:
                        v48 := GenS3{F0: 42}
                        _ = v48
                    }
                    v49 := GenS2{F0: "hello"}
                    fmt.Println(v2, v2)
                    _ = v49
                }
            } else {
                fmt.Println(v2, v1, v4)
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
                v50 := GenS1{F0: (0 <= 485), F1: 42}
                _ = v50
            }
        default:
            _ = len(v3)
        }
    }
    switch (("foo bar" + "") + ("a" + "")) {
    case "a":
        for fi51 := 0; fi51 < 12; fi51++ {
            if ((3.14159 == -2.25) && true) {
                switch (fi51 & fi51) {
                case 32767:
                    switch 883 {
                    case 9223372036854775807:
                        var v52 uint32 = 42
                        _ = v52
                    }
                case 42:
                    fmt.Println(v2, v1, v2)
                }
            }
            _ = fi51
        }
    case "\n\t":
        for fi53 := 0; fi53 < 17; fi53++ {
            if (2147483647 == 42) {
                switch "a" {
                case "a":
                    v54 := v3[fi53]
                    _ = v54
                case "\n\t":
                    var v55 uint8 = 1
                    _ = v55
                default:
                    for fi56 := 0; fi56 < 5; fi56++ {
                        switch fi53 {
                        case 9223372036854775807:
                            v22 = append(v22, (fi53 / ((fi56) | 1)))
                        case -2147483648:
                            var v57 int32 = 1
                            _ = v57
                        default:
                            v58 := v3[fi53]
                            _ = v58
                        }
                        switch "\n\t" {
                        case "foo bar":
                            var v59 uint8 = uint8(v4)
                            _ = v59
                        }
                        _ = fi56
                    }
                }
                var v60 uint16 = v1
                _ = v60
            }
            _ = fi53
        }
    case "hello":
        if (855 < 83) {
            if true {
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
                for fi61 := 0; fi61 < 1; fi61++ {
                    switch (fi61 * fi61) {
                    case -352:
                        v62 := GenS2{F0: ("" + (("foo bar" + "foo bar") + "\n\t"))}
                        _ = v62
                    case -835:
                        for fi63 := 0; fi63 < 1; fi63++ {
                            v64 := GenS2{F0: "hello"}
                            if true {
                                var v65 bool = v32
                                if true {
                                    if true {
                                        if v65 {
                                            _ = len(v3)
                                            var v66 int8 = 100
                                            _ = v66
                                        }
                                    } else {
                                        v1 -= v1
                                        if len(v22) >= 2 {
                                            v22 = v22[1:len(v22)]
                                        }
                                        v67 := GenS4{F0: 54, F1: 100, F2: -182}
                                        _ = v67
                                    }
                                    v1 = 65535
                                } else {
                                    fmt.Println(v2)
                                    switch "a" {
                                    case "a":
                                        if v32 {
                                            fmt.Println(fi61)
                                            v4 = uint64(v2)
                                            v68 := GenS4{F0: 127, F1: 0, F2: 1}
                                            _ = v68
                                        } else {
                                            var v69 uint32 = 65535
                                            var v70 int8 = -128
                                            v2 -= (v2 - v2)
                                            _ = v69
                                            _ = v70
                                        }
                                    case "foo bar":
                                        if ((0 > 1) || (false && ("hello" <= "foo bar"))) {
                                            v71 := GenS2{F0: ""}
                                            _ = v71
                                        } else {
                                            v72 := GenS4{F0: 127, F1: 101, F2: 9223372036854775807}
                                            _ = v72
                                        }
                                    case "\n\t":
                                        v73 := GenS1{F0: v65, F1: 0}
                                        _ = v73
                                    }
                                }
                                _ = len(v22)
                                var v74 string = "hello"
                                _ = v65
                                _ = v74
                            } else {
                                if v32 {
                                    fmt.Println(fi61, v1)
                                    v4 = v4
                                    switch fi63 {
                                    case 195:
                                        var v75 bool = false
                                        _ = v75
                                    case 32767:
                                        v22 = append(v22, (fi61 >> ((v2) & 63)))
                                    case -15:
                                        fmt.Println(v2, fi61, v4)
                                    default:
                                        switch int(v4) {
                                        case 128:
                                            var v76 string = ("\n\t" + "foo bar")
                                            _ = v76
                                        }
                                    }
                                }
                                var v77 int8 = -1
                                var v78 uint = (v2 | v2)
                                _ = len(v3)
                                _ = v77
                                _ = v78
                            }
                            if (0 >= 986) {
                                if false {
                                    v79 := GenS2{F0: (("a" + "\n\t") + "a")}
                                    if len(v22) >= 2 {
                                        v22 = v22[1:len(v22)]
                                    }
                                    if (87 <= -1) {
                                        var v80 int32 = -2147483648
                                        var v81 int64 = 0
                                        fmt.Println(v81)
                                        switch fi63 {
                                        case 9223372036854775807:
                                            switch (fi63 >> ((v2) & 63)) {
                                            case 127:
                                                v1 -= v1
                                            default:
                                                if v32 {
                                                    v82 := GenS4{F0: 115, F1: int8(v4), F2: (v81 / (((v81 % ((v81) | 1))) | 1))}
                                                    var v83 int64 = (v81 ^ v81)
                                                    v84 := GenS2{F0: ((("\n\t" + "\n\t") + "") + "")}
                                                    _ = v82
                                                    _ = v83
                                                    _ = v84
                                                }
                                            }
                                        case 1:
                                            v85 := v3[(fi63 | (fi61 * fi63))]
                                            _ = v85
                                        case -2147483648:
                                            switch (^(fi61)) {
                                            case -685:
                                                var v86 uint = (v2 / ((v2) | 1))
                                                _ = v86
                                            case 233:
                                                v32 = false
                                            }
                                        }
                                        _ = v80
                                        _ = v81
                                    } else {
                                        v4 = v4
                                        var v87 int64 = -9223372036854775808
                                        v2 = (v2 >> (((v2 + v2)) & 63))
                                        _ = v87
                                    }
                                    switch "a" {
                                    case "\n\t":
                                        var v88 string = "foo bar"
                                        _ = v88
                                    case "foo bar":
                                        fmt.Println(v1)
                                    }
                                    _ = v79
                                } else {
                                    var v89 string = "foo bar"
                                    v4 = v4
                                    var v90 uint = (^(v2))
                                    _ = v89
                                    _ = v90
                                }
                                v91 := GenS3{F0: 100}
                                var v92 int8 = 0
                                v93 := GenS1{F0: true, F1: -9223372036854775808}
                                v92 += (v92 ^ v92)
                                _ = v91
                                _ = v92
                                _ = v93
                            }
                            _ = fi63
                            _ = v64
                        }
                    case 127:
                        switch "\n\t" {
                        case "hello":
                            if true {
                                var v94 int16 = -353
                                if len(v22) >= 2 {
                                    v22 = v22[1:len(v22)]
                                }
                                _ = v94
                            }
                        default:
                            if (-32768 == -664) {
                                v2 -= v2
                            } else {
                                if (0 != uint16(v4)) {
                                    var v95 uint32 = 2147483647
                                    fmt.Println(v95, v2, v95)
                                    fmt.Println(v95)
                                    _ = v95
                                } else {
                                    v96 := GenS5{F0: 255, F1: 42, F2: "a"}
                                    _ = v96
                                }
                                var v97 int32 = 0
                                switch "\n\t" {
                                case "hello":
                                    v98 := GenS6{F0: 42, F1: ("a" != "\n\t")}
                                    _ = v98
                                case "foo bar":
                                    delete(v3, fi61)
                                case "":
                                    if (1e10 < float64(3.14159)) {
                                        v99 := GenS5{F0: (v1 % (((v1 % (((v1 << ((v2) & 63))) | 1))) | 1)), F1: (v97 - v97), F2: "a"}
                                        v22 = append(v22, 128)
                                        v100 := v3[(fi61 / ((fi61) | 1))]
                                        _ = v100
                                        _ = v99
                                    } else {
                                        v101 := GenS7{F0: (v4 * v4)}
                                        switch (^(fi61)) {
                                        case 127:
                                            var v102 uint32 = 1
                                            _ = v102
                                        default:
                                            if v32 {
                                                var v103 int64 = -1
                                                v3[fi61] = (!(v97 <= 42))
                                                _ = v103
                                            } else {
                                                v104 := GenS7{F0: 2}
                                                for fi105 := 0; fi105 < 10; fi105++ {
                                                    fmt.Println(fi105)
                                                    var v106 int8 = 127
                                                    fmt.Println(fi105, v2)
                                                    _ = fi105
                                                    _ = v106
                                                }
                                                v107 := GenS5{F0: 32767, F1: -1, F2: "foo bar"}
                                                _ = v104
                                                _ = v107
                                            }
                                        }
                                        switch (fi61 >> ((v2) & 63)) {
                                        case 0:
                                            switch "hello" {
                                            case "":
                                                if len(v22) >= 2 {
                                                    v22 = v22[1:len(v22)]
                                                }
                                            case "a":
                                                switch fi61 {
                                                case -898:
                                                    v2 += (^(v2))
                                                case -2147483648:
                                                    v108 := GenS3{F0: -5}
                                                    _ = v108
                                                default:
                                                    var v109 uint16 = v1
                                                    _ = v109
                                                }
                                            default:
                                                v3[(fi61 + fi61)] = ((((((0.0 / float64(float32(-2.25))) * (1.5 / float64(float32(1.5)))) + ((((1.5 * 1.5) + -2.25) * (1e10 / 3.14159)) + 0.1)) - -2.25) / 0.1) < 1e10)
                                            }
                                        case -2147483648:
                                            if (0 > uint32(v4)) {
                                                switch (^(fi61)) {
                                                case 32767:
                                                    if len(v22) >= 2 {
                                                        v22 = v22[1:len(v22)]
                                                    }
                                                default:
                                                    fmt.Println(v32, v32, v2)
                                                }
                                                switch "" {
                                                case "hello":
                                                    fmt.Println(fi61)
                                                case "foo bar":
                                                    v3[(^(fi61))] = ((!(v4 > 2)) && false)
                                                case "a":
                                                    switch (fi61 * fi61) {
                                                    case 98:
                                                        v110 := GenS5{F0: (v1 / (((v1 / ((42) | 1))) | 1)), F1: 42, F2: "foo bar"}
                                                        _ = v110
                                                    case 913:
                                                        switch "" {
                                                        case "hello":
                                                            var v111 bool = (v1 != (v1 * (v1 ^ v1)))
                                                            _ = v111
                                                        case "a":
                                                            v1 = v1
                                                        }
                                                    default:
                                                        v32 = v32
                                                    }
                                                default:
                                                    v112 := GenS8{F0: true, F1: 353}
                                                    _ = v112
                                                }
                                                v113 := v3[(fi61 / ((32768) | 1))]
                                                _ = v113
                                            }
                                        case 9223372036854775807:
                                            if v32 {
                                                for fi114 := 0; fi114 < 18; fi114++ {
                                                    var v115 int16 = 42
                                                    var v116 uint32 = 2147483648
                                                    _ = len(v3)
                                                    _ = fi114
                                                    _ = v115
                                                    _ = v116
                                                }
                                                for fi117 := 0; fi117 < 14; fi117++ {
                                                    fmt.Println(v2)
                                                    var v118 int8 = 1
                                                    if len(v22) >= 2 {
                                                        v22 = v22[1:len(v22)]
                                                    }
                                                    _ = fi117
                                                    _ = v118
                                                }
                                                fmt.Println(v2)
                                            } else {
                                                for fi119 := 0; fi119 < 1; fi119++ {
                                                    v120 := GenS8{F0: (!false), F1: v1}
                                                    var v121 string = "hello"
                                                    _ = fi119
                                                    _ = v120
                                                    _ = v121
                                                }
                                            }
                                        }
                                        _ = v101
                                    }
                                default:
                                    v97 += (v97 - (v97 - (v97 << (((^(v2))) & 63))))
                                }
                                _ = v97
                            }
                        }
                    default:
                        v122 := GenS9{F0: ("\n\t" + "a"), F1: (^(v2)), F2: -1}
                        _ = v122
                    }
                    _ = fi61
                }
            }
        }
    }
    switch 889 {
    case 127:
        v32 = v32
    }
    for fi123 := 0; fi123 < 2; fi123++ {
        var v124 int8 = 0
        _ = fi123
        _ = v124
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v32
    _ = v4
}
