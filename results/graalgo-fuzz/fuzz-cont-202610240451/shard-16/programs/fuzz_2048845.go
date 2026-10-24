package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint16; F2 int }
type GenS3 struct { F0 uint16; F1 uint32; F2 uint64 }
type GenS4 struct { F0 int16; F1 uint32; F2 int }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int8; F1 bool }
type GenS8 struct { F0 uint64; F1 int }
type GenS9 struct { F0 int; F1 uint }
type GenS10 struct { F0 int32; F1 uint8 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 string }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 string }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 int64 }
func genFn0() int32 {
    v6 := []int{473}
    _ = len(v6)
    v7 := GenS1{F0: -32768}
    var v8 bool = ((0.1 == -2.25) || (0.0 > 1.5))
    switch ("foo bar" + "\n\t") {
    case "":
        var v9 int8 = 127
        _ = v9
    }
    _ = v6
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 510
    var v3 int16 = -1
    v4 := make(map[string]string)
    v5 := []int{(v2 & (v2 * 42)), v2}
    fmt.Println(v3, v2)
    if false {
        v2 = v2
        v3 = 1000
        var v10 uint64 = 4294967295
        _ = v10
    } else {
        var v11 float64 = ((1.5 * 1.5) / 3.14159)
        for fi12 := 0; fi12 < 18; fi12++ {
            for fi13 := 0; fi13 < 9; fi13++ {
                if ((true || ((4294967295 > 99) || (!false))) && true) {
                    var v14 uint = 0
                    _ = v14
                } else {
                    v11 -= 3.14159
                }
                switch (fi13 << ((925) & 63)) {
                case 127:
                    switch fi12 {
                    case -130:
                        v15 := GenS2{F0: 1, F1: 516, F2: 213}
                        _ = v15
                    case -68:
                        fmt.Println(v3, fi12)
                    case 32767:
                        v16 := GenS2{F0: 42, F1: 65535, F2: v2}
                        _ = v16
                    default:
                        switch "\n\t" {
                        case "foo bar":
                            var v17 float64 = (((0.1 * v11) - 1e10) + (((v11 / (float64(0.1) / v11)) / 1.0) + (float64(0.0) * (0.1 - (3.14159 / 0.1)))))
                            _ = v17
                        default:
                            v5 = append(v5, (fi12 << ((143) & 63)))
                        }
                    }
                case -977:
                    if ("hello" <= "") {
                        switch fi12 {
                        case -2147483648:
                            var v18 uint64 = 1
                            var v19 int32 = 1
                            var v20 int16 = (v3 - v3)
                            var v21 int = fi12
                            v22 := GenS3{F0: 255, F1: 65535, F2: 0}
                            _ = v18
                            _ = v19
                            _ = v20
                            _ = v21
                            _ = v22
                        case -429:
                            var v23 float64 = (0.0 / 1.0)
                            _ = v23
                        default:
                            v5 = append(v5, (v2 ^ 52))
                        }
                        v11 = v11
                    }
                case -2147483648:
                    if len(v5) > 0 {
                        v5[0] = (fi13 * fi13)
                    }
                default:
                    if false {
                        v24 := v1[4294967295]
                        v24 += 902
                        _ = v24
                    } else {
                        var v25 int = fi12
                        _ = v25
                    }
                }
                _ = fi13
            }
            _ = fi12
        }
        _ = v11
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        v27 := GenS1{F0: v3}
        v28 := GenS4{F0: v3, F1: 1, F2: (v2 | v2)}
        switch v2 {
        case 127:
            for fi29 := 0; fi29 < 17; fi29++ {
                fmt.Println(v2, fi29)
                if false {
                    v30 := GenS1{F0: -1}
                    _ = v30
                }
                _ = fi29
            }
        case 1:
            if (0.0 >= 3.14159) {
                _ = len(v4)
                var v31 string = "foo bar"
                _ = v31
            }
        default:
            delete(v1, 417)
        }
        _ = fi26
        _ = v27
        _ = v28
    }
    var v32 int16 = -973
    delete(v1, 1)
    fmt.Println(v32)
    fmt.Println(v32, v32, v32)
    v33 := GenS3{F0: 604, F1: 2147483647, F2: 2}
    switch (v2 + -355) {
    case 314:
        fmt.Println(v32)
    default:
        for fi34 := 0; fi34 < 16; fi34++ {
            v32 += (v3 >> ((2) & 63))
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            delete(v4, ("" + "\n\t"))
            _ = fi34
        }
    }
    v35 := GenS1{F0: (v3 + 1000)}
    for fi36 := 0; fi36 < 3; fi36++ {
        if (("hello" + ("a" + "")) > "foo bar") {
            if (!false) {
                if false {
                    fmt.Println(fi36, v3)
                } else {
                    switch ("hello" + "\n\t") {
                    case "foo bar":
                        v37 := GenS2{F0: 1, F1: 0, F2: (fi36 * (v2 + 372))}
                        _ = v37
                    default:
                        if (true && (false || true)) {
                            var v38 int = (fi36 | fi36)
                            if false {
                                v39 := GenS3{F0: 237, F1: 2147483647, F2: 18446744073709551615}
                                _ = v39
                            } else {
                                for fi40 := 0; fi40 < 2; fi40++ {
                                    var v41 string = "a"
                                    var v42 int = (fi36 * (fi40 << ((135) & 63)))
                                    var v43 uint64 = 18446744073709551615
                                    v44 := GenS3{F0: 1, F1: 589, F2: 6}
                                    if true {
                                        fmt.Println(v32, v2, fi40)
                                        v5 = append(v5, v42)
                                    }
                                    _ = fi40
                                    _ = v41
                                    _ = v42
                                    _ = v43
                                    _ = v44
                                }
                                fmt.Println(v3, v2)
                                var v45 uint64 = 0
                                _ = v45
                            }
                            if (20 != 32767) {
                                v4["hello"] = ("hello" + "hello")
                                switch (("hello" + "hello") + ("a" + ("foo bar" + "a"))) {
                                case "hello":
                                    if false {
                                        for fi46 := 0; fi46 < 12; fi46++ {
                                            var v47 bool = true
                                            v4[""] = ""
                                            v48 := GenS5{F0: (v32 & (-(v32))), F1: -186}
                                            _ = fi46
                                            _ = v47
                                            _ = v48
                                        }
                                        switch (v2 + (v38 >> ((0) & 63))) {
                                        case 32768:
                                            var v49 uint64 = 2
                                            _ = v49
                                        case -996:
                                            v2 = fi36
                                        case 1:
                                            var v50 int8 = 42
                                            _ = v50
                                        default:
                                            _ = len(v1)
                                        }
                                        v2 = v38
                                    } else {
                                        if len(v5) > 0 {
                                            v5[0] = (v38 / ((128) | 1))
                                        }
                                        v51 := GenS2{F0: 42, F1: 0, F2: v38}
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        _ = v51
                                    }
                                case "":
                                    v52 := v1[1]
                                    _ = v52
                                case "foo bar":
                                    v1[193] = 427
                                }
                            } else {
                                var v53 uint16 = 255
                                v3 = v3
                                var v54 float64 = 3.14159
                                _ = v53
                                _ = v54
                            }
                            _ = v38
                        } else {
                            switch -2147483648 {
                            case 32767:
                                switch ("" + "") {
                                case "foo bar":
                                    for fi55 := 0; fi55 < 2; fi55++ {
                                        fmt.Println(fi55, fi36)
                                        v32 -= (v3 / (((v32 ^ v32)) | 1))
                                        v56 := GenS6{F0: 42}
                                        _ = fi55
                                        _ = v56
                                    }
                                case "hello":
                                    if ((!((1.5 / 1.0) >= 3.14159)) && true) {
                                        fmt.Println(v32, v3)
                                        for fi57 := 0; fi57 < 3; fi57++ {
                                            switch (v2 + (fi57 ^ (fi36 % (((-(fi57))) | 1)))) {
                                            case 127:
                                                v2 -= fi36
                                            case 411:
                                                fmt.Println(fi57, fi36)
                                            }
                                            _ = fi57
                                        }
                                    }
                                }
                            case -675:
                                for fi58 := 0; fi58 < 20; fi58++ {
                                    var v59 float32 = 1.5
                                    _ = fi58
                                    _ = v59
                                }
                            case 128:
                                for fi60 := 0; fi60 < 7; fi60++ {
                                    v32 -= 42
                                    var v61 int64 = -9223372036854775808
                                    var v62 int8 = -1
                                    _ = fi60
                                    _ = v61
                                    _ = v62
                                }
                            default:
                                var v63 uint64 = 18446744073709551615
                                _ = v63
                            }
                            switch "\n\t" {
                            case "\n\t":
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            if (!true) {
                                switch fi36 {
                                case 2147483647:
                                    v64 := GenS7{F0: 1, F1: (1 == 4294967295)}
                                    _ = v64
                                default:
                                    var v65 int = (^(v2))
                                    _ = v65
                                }
                            }
                        }
                    }
                    v32 -= (v32 * v32)
                }
                v4[""] = "a"
            }
        } else {
            fmt.Println(v2)
            if ("\n\t" == (("hello" + ("a" + "foo bar")) + (((((("" + "foo bar") + "foo bar") + "foo bar") + "foo bar") + "foo bar") + "hello"))) {
                v2 = -631
            } else {
                switch v2 {
                case -9223372036854775808:
                    fmt.Println(fi36)
                }
                v66 := GenS8{F0: 98, F1: (v2 & v2)}
                v3 -= (v3 - 32767)
                _ = v66
            }
        }
        delete(v4, ("foo bar" + "foo bar"))
        _ = fi36
    }
    switch ("a" + "a") {
    case "":
        if false {
            switch v2 {
            case 9223372036854775807:
                switch v2 {
                case 28:
                    if true {
                        var v67 string = "foo bar"
                        var v68 float32 = 1.5
                        fmt.Println(v67, v3, v2)
                        _ = v67
                        _ = v68
                    } else {
                        if true {
                            var v69 int32 = -1
                            v70 := GenS4{F0: 940, F1: 2, F2: v2}
                            _ = len(v5)
                            _ = v69
                            _ = v70
                        }
                        for fi71 := 0; fi71 < 13; fi71++ {
                            fmt.Println(fi71)
                            _ = fi71
                        }
                        if false {
                            if (127 <= 1) {
                                var v72 float32 = (float32(3.14159) - 1.5)
                                switch v2 {
                                case -188:
                                    fmt.Println(v3, v3, v72)
                                case -908:
                                    if true {
                                        var v73 int = (v2 | v2)
                                        _ = v73
                                    }
                                }
                                var v74 uint = 383
                                _ = v72
                                _ = v74
                            }
                            v75 := GenS5{F0: v3, F1: -880}
                            _ = v75
                        }
                    }
                default:
                    var v76 float64 = 3.14159
                    _ = v76
                }
            case 4294967295:
                v1[283] = 32767
            default:
                for fi77 := 0; fi77 < 16; fi77++ {
                    switch v2 {
                    case 32768:
                        for fi78 := 0; fi78 < 12; fi78++ {
                            if false {
                                v3 += 730
                                var v79 bool = true
                                _ = v79
                            } else {
                                v80 := GenS9{F0: (fi77 | (fi78 >> ((42) & 63))), F1: 255}
                                switch 1 {
                                case -815:
                                    if (false || ((v2 << ((1) & 63)) == -514)) {
                                        var v81 uint16 = 32767
                                        v3 = (v3 / ((v32) | 1))
                                        var v82 uint = 69
                                        fmt.Println(v2, v3, v82)
                                        _ = v81
                                        _ = v82
                                    } else {
                                        v2 = (v2 & fi78)
                                        var v83 int32 = 305
                                        _ = len(v5)
                                        v5 = append(v5, fi77)
                                        _ = v83
                                    }
                                case -546:
                                    _ = len(v4)
                                case -951:
                                    var v84 uint64 = 1
                                    _ = v84
                                }
                                _ = v80
                            }
                            var v85 int64 = -9223372036854775808
                            _ = fi78
                            _ = v85
                        }
                    case -9223372036854775808:
                        switch v2 {
                        case 356:
                            var v86 int8 = 0
                            _ = v86
                        case 1:
                            _ = len(v5)
                        case -599:
                            v3 = (v3 >> ((0) & 63))
                        default:
                            var v87 int = -1
                            _ = v87
                        }
                    case 0:
                        v88 := GenS10{F0: -356, F1: 2}
                        _ = v88
                    }
                    for fi89 := 0; fi89 < 1; fi89++ {
                        var v90 bool = (41 > 1)
                        if (false && (463 < -2147483648)) {
                            if true {
                                v90 = true
                                var v91 string = "hello"
                                _ = v91
                            } else {
                                var v92 uint = 926
                                v2 += (fi77 ^ (-(fi77)))
                                v93 := GenS11{F0: 845}
                                _ = v92
                                _ = v93
                            }
                        }
                        _ = fi89
                        _ = v90
                    }
                    var v94 float64 = float64(0.1)
                    _ = fi77
                    _ = v94
                }
            }
            var v95 uint16 = 65535
            _ = len(v4)
            _ = v95
        } else {
            for fi96 := 0; fi96 < 20; fi96++ {
                for fi97 := 0; fi97 < 6; fi97++ {
                    v5 = append(v5, v2)
                    fmt.Println(v32)
                    _ = fi97
                }
                if false {
                    v98 := GenS12{F0: (("a" + "a") + "\n\t")}
                    fmt.Println(v32, fi96)
                    v2 = v2
                    _ = v98
                }
                var v99 int64 = -557
                _ = fi96
                _ = v99
            }
        }
    }
    fmt.Println(v32, v32, v3)
    v2 = 42
    var v100 uint = 18446744073709551615
    switch 32767 {
    case -777:
        if (!(-9223372036854775808 != 9223372036854775807)) {
            switch ("\n\t" + ("a" + "\n\t")) {
            case "":
                var v101 float32 = ((float32(float64(0.1)) / 0.1) + 1e10)
                _ = v101
            case "\n\t":
                var v102 int8 = -128
                _ = v102
            case "hello":
                var v103 uint32 = 531
                _ = v103
            default:
                var v104 int64 = 1
                _ = v104
            }
        }
    case 1:
        v4[("foo bar" + (((("foo bar" + ("a" + ("" + ((("\n\t" + ("a" + "")) + ((("foo bar" + "\n\t") + "") + "foo bar")) + "foo bar")))) + "hello") + "\n\t") + ("\n\t" + ("foo bar" + (("hello" + "hello") + "")))))] = "\n\t"
    case -830:
        for fi105 := 0; fi105 < 17; fi105++ {
            if ((0.1 != (0.0 / float64(1.5))) && ((0 == 4294967295) || (561 <= 55))) {
                v106 := v4[""]
                v107 := GenS9{F0: (v2 | (v2 | (v2 ^ (fi105 + 4294967295)))), F1: v100}
                switch (v106 + "hello") {
                case "\n\t":
                    if (255 < 255) {
                        for fi108 := 0; fi108 < 18; fi108++ {
                            v4["a"] = v106
                            _ = fi108
                        }
                        if (v106 != "") {
                            switch v106 {
                            case "\n\t":
                                fmt.Println(v106, v32)
                            default:
                                if (!true) {
                                    for fi109 := 0; fi109 < 9; fi109++ {
                                        if true {
                                            switch v106 {
                                            case "foo bar":
                                                fmt.Println(fi109)
                                            case "\n\t":
                                                var v110 string = ((v106 + "foo bar") + (("\n\t" + (v106 + ("" + v106))) + v106))
                                                _ = v110
                                            case "a":
                                                var v111 uint = v100
                                                _ = v111
                                            }
                                            var v112 float64 = (1e10 - ((-2.25 + 0.0) + (-2.25 / (-2.25 + -2.25))))
                                            _ = len(v4)
                                            _ = v112
                                        }
                                        switch v2 {
                                        case -756:
                                            v3 += v3
                                        case 127:
                                            var v113 int = v2
                                            _ = v113
                                        }
                                        switch 2147483647 {
                                        case 716:
                                            if true {
                                                var v114 string = (v106 + "\n\t")
                                                v100 = (v100 / (((v100 * (v100 << ((4294967295) & 63)))) | 1))
                                                _ = v114
                                            } else {
                                                v3 += 0
                                                v115 := GenS13{F0: (-446 == 32767)}
                                                _ = v115
                                            }
                                        case -92:
                                            v1[0] = (fi105 / (((fi109 - 711)) | 1))
                                        default:
                                            switch -186 {
                                            case 2147483647:
                                                if ((583 >= 42) && true) {
                                                    var v116 int16 = (v32 ^ v32)
                                                    _ = v116
                                                } else {
                                                    v3 -= 1
                                                    var v117 int8 = -128
                                                    if (-999 != 1) {
                                                        switch fi105 {
                                                        case 802:
                                                            v118 := GenS14{F0: v100}
                                                            _ = v118
                                                        case 42:
                                                            v106 = (v106 + "hello")
                                                        case 4294967295:
                                                            _ = len(v1)
                                                        default:
                                                            fmt.Println(v100)
                                                        }
                                                    } else {
                                                        var v119 int16 = v32
                                                        var v120 uint32 = 42
                                                        if true {
                                                            var v121 uint8 = 15
                                                            v122 := v1[(v120 << ((v100) & 63))]
                                                            _ = v121
                                                            _ = v122
                                                        } else {
                                                            if false {
                                                                v119 = v119
                                                                _ = len(v5)
                                                                var v123 int = fi109
                                                                _ = v123
                                                            } else {
                                                                fmt.Println(v32, fi109)
                                                                var v124 bool = true
                                                                _ = v124
                                                            }
                                                        }
                                                        v100 -= (v100 | v100)
                                                        _ = v119
                                                        _ = v120
                                                    }
                                                    _ = v117
                                                }
                                            case 624:
                                                v125 := GenS1{F0: v32}
                                                _ = v125
                                            default:
                                                v106 = v106
                                                fmt.Println(fi105)
                                                fmt.Println(v106, v2, v100)
                                            }
                                        }
                                        fmt.Println(v2)
                                        v1[709] = 32767
                                        _ = fi109
                                    }
                                    v126 := GenS15{F0: ""}
                                    v127 := GenS16{F0: -128}
                                    _ = v126
                                    _ = v127
                                }
                            }
                            var v128 bool = (!false)
                            if (2147483647 != -122) {
                                if len(v5) > 0 {
                                    v5[0] = 873
                                }
                            }
                            _ = v128
                        }
                        if (float64(3.14159) > (1e10 - -2.25)) {
                            var v129 uint64 = 42
                            v130 := GenS2{F0: 159, F1: 65535, F2: v2}
                            _ = v129
                            _ = v130
                        }
                    } else {
                        v2 += fi105
                        var v131 uint64 = 125
                        v132 := GenS3{F0: 440, F1: 2147483648, F2: 42}
                        _ = v131
                        _ = v132
                    }
                case "foo bar":
                    v133 := GenS17{F0: false}
                    _ = v133
                default:
                    switch fi105 {
                    case 1:
                        var v134 uint16 = 2
                        _ = v134
                    case -9223372036854775808:
                        v135 := GenS18{F0: 693}
                        _ = v135
                    case 0:
                        v106 = ((v106 + "a") + ("foo bar" + "hello"))
                    default:
                        fmt.Println(fi105, v32, v3)
                    }
                }
                _ = v106
                _ = v107
            }
            _ = fi105
        }
    default:
        for fi136 := 0; fi136 < 10; fi136++ {
            _ = len(v5)
            switch ("foo bar" + "foo bar") {
            case "a":
                _ = len(v5)
            case "\n\t":
                v3 -= 800
            }
            _ = fi136
        }
    }
    v100 = 669
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v35
    _ = v4
    _ = v5
}
