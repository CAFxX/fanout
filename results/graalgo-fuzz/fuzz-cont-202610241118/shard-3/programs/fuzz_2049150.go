package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS2 struct { F0 int32; F1 int8; F2 int }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int16; F1 string; F2 uint64 }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 int64; F1 int32; F2 bool }
type GenS7 struct { F0 int; F1 uint32; F2 uint16 }
type GenS8 struct { F0 uint32; F1 uint32 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int16; F1 bool; F2 uint32 }
type GenS11 struct { F0 string; F1 uint32; F2 uint8 }
type GenS12 struct { F0 bool; F1 uint; F2 int16 }
type GenS13 struct { F0 int; F1 bool; F2 int }
type GenS14 struct { F0 uint16; F1 int32 }
type GenS15 struct { F0 bool; F1 int }
type GenS16 struct { F0 uint }

func main() {
    var v1 uint32 = 0
    v2 := []string{"foo bar"}
    var v3 float32 = 0.0
    v4 := make(map[uint32]int)
    fmt.Println(v1, v1)
    var v5 bool = (0 != 2)
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := GenS1{F0: 128, F1: 64, F2: -128}
        _ = fi6
        _ = v7
    }
    if true {
        if (!v5) {
            if false {
                v5 = true
                fmt.Println(v5)
            } else {
                switch -751 {
                case 576:
                    v3 -= v3
                case 32767:
                    v8 := GenS2{F0: 394, F1: 104, F2: 1}
                    _ = v8
                }
                v9 := GenS2{F0: -1, F1: 127, F2: 50}
                _ = v9
            }
            for fi10 := 0; fi10 < 9; fi10++ {
                v1 = (^(v1))
                var v11 int32 = 1
                switch fi10 {
                case 9223372036854775807:
                    for fi12 := 0; fi12 < 12; fi12++ {
                        fmt.Println(v11, fi12)
                        var v13 uint16 = 42
                        if (!(!(363 != 110))) {
                            v1 = (^(v1))
                            switch "\n\t" {
                            case "a":
                                var v14 uint32 = v1
                                _ = v14
                            case "hello":
                                var v15 int64 = 206
                                _ = v15
                            case "\n\t":
                                v1 = 4294967295
                            }
                            var v16 uint16 = 65535
                            switch ("\n\t" + (("hello" + "") + ("\n\t" + ("foo bar" + ("hello" + ("\n\t" + "hello")))))) {
                            case "":
                                switch -1 {
                                case 127:
                                    v17 := GenS1{F0: 2, F1: uint8(v13), F2: 4}
                                    _ = v17
                                case 32768:
                                    v3 = v3
                                    fmt.Println(v13)
                                default:
                                    v1 = 1
                                }
                            default:
                                v16 += v13
                            }
                            v16 -= v16
                            _ = v16
                        }
                        if (100 != 100) {
                            v11 = v11
                            var v18 uint32 = v1
                            var v19 uint16 = (^(v13))
                            var v20 int16 = int16(v1)
                            if v5 {
                                if false {
                                    v18 = v18
                                    v4[(^(v1))] = (-(fi10))
                                    var v21 int8 = 100
                                    _ = v21
                                } else {
                                    var v22 uint16 = (^(v13))
                                    v23 := GenS3{F0: 65535, F1: v20}
                                    _ = v22
                                    _ = v23
                                }
                            }
                            _ = v18
                            _ = v19
                            _ = v20
                        } else {
                            switch 897 {
                            case -73:
                                if false {
                                    v3 = ((v3 / float32((0.0 + 1.5))) / (v3 * v3))
                                    var v24 uint16 = v13
                                    _ = v24
                                } else {
                                    fmt.Println(fi12, v5)
                                    v11 -= (v11 % (((v11 % ((int32(v13)) | 1))) | 1))
                                }
                            }
                            var v25 uint64 = 284
                            _ = v25
                        }
                        var v26 uint16 = v13
                        _ = fi12
                        _ = v13
                        _ = v26
                    }
                case 667:
                    switch (fi10 & (fi10 >> ((18446744073709551615) & 63))) {
                    case 1:
                        if v5 {
                            var v27 float32 = v3
                            if true {
                                v1 = (v1 ^ 42)
                                var v28 uint = 0
                                switch (fi10 * 42) {
                                case 42:
                                    v27 = -2.25
                                case 2147483647:
                                    v27 += 1.5
                                }
                                _ = v28
                            }
                            fmt.Println(v1, v27)
                            _ = v27
                        } else {
                            fmt.Println(fi10, fi10)
                            var v29 uint32 = 2147483648
                            for fi30 := 0; fi30 < 8; fi30++ {
                                switch "" {
                                case "foo bar":
                                    if v5 {
                                        _ = len(v2)
                                        switch (^(fi30)) {
                                        case 4294967295:
                                            switch "" {
                                            case "foo bar":
                                                fmt.Println(v29, v11)
                                            default:
                                                fmt.Println(v5)
                                            }
                                        case 127:
                                            var v31 uint = 42
                                            _ = v31
                                        }
                                    }
                                case "a":
                                    v5 = false
                                    var v32 uint8 = 74
                                    switch "a" {
                                    case "foo bar":
                                        v32 = uint8(v1)
                                    case "a":
                                        var v33 uint = 1
                                        _ = v33
                                    case "\n\t":
                                        var v34 uint16 = 255
                                        _ = v34
                                    default:
                                        fmt.Println(v3, v29)
                                        v29 -= v29
                                        fmt.Println(v11, v1)
                                        var v35 int64 = -1
                                        fmt.Println(v5, v3, v35)
                                        _ = v35
                                    }
                                    _ = v32
                                }
                                _ = fi30
                            }
                            _ = v29
                        }
                    case -468:
                        if (42 < 42) {
                            for fi36 := 0; fi36 < 17; fi36++ {
                                var v37 int32 = 0
                                _ = fi36
                                _ = v37
                            }
                        } else {
                            v1 -= 0
                            if v5 {
                                var v38 float64 = 3.14159
                                _ = len(v2)
                                _ = v38
                            } else {
                                v1 = 22
                            }
                            fmt.Println(fi10)
                        }
                    default:
                        v11 -= v11
                    }
                default:
                    var v39 int8 = -1
                    _ = v39
                }
                _ = fi10
                _ = v11
            }
            switch "hello" {
            case "a":
                for fi40 := 0; fi40 < 6; fi40++ {
                    var v41 float64 = 0.1
                    fmt.Println(v1, v41)
                    _ = fi40
                    _ = v41
                }
            case "foo bar":
                switch 42 {
                case 453:
                    delete(v4, (^(v1)))
                case -203:
                    for fi42 := 0; fi42 < 3; fi42++ {
                        switch (("a" + "a") + "\n\t") {
                        case "foo bar":
                            var v43 uint64 = 2
                            _ = v43
                        case "":
                            var v44 uint64 = 4294967295
                            _ = v44
                        }
                        if v5 {
                            for fi45 := 0; fi45 < 14; fi45++ {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                fmt.Println(fi42, v3, v1)
                                v1 += 2
                                v1 -= (^(v1))
                                var v46 uint = 1
                                delete(v4, 2147483647)
                                _ = fi45
                                _ = v46
                            }
                            var v47 string = "hello"
                            var v48 uint = 255
                            _ = v47
                            _ = v48
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            switch 4294967295 {
                            case -512:
                                var v49 uint16 = 65535
                                _ = v49
                            case -489:
                                v3 = v3
                            default:
                                var v50 uint32 = (v1 / ((v1) | 1))
                                _ = v50
                            }
                        }
                        _ = fi42
                    }
                }
            case "":
                if ("foo bar" > "hello") {
                    if (1 != 42) {
                        v51 := GenS4{F0: 32767, F1: "a", F2: 42}
                        _ = v51
                    } else {
                        switch 32767 {
                        case 449:
                            v1 = (v1 << ((2) & 63))
                        case -9223372036854775808:
                            v3 -= float32((3.14159 - 1.5))
                        case -117:
                            switch "\n\t" {
                            case "a":
                                for fi52 := 0; fi52 < 13; fi52++ {
                                    switch (("\n\t" + "") + "") {
                                    case "":
                                        for fi53 := 0; fi53 < 13; fi53++ {
                                            v5 = true
                                            v54 := GenS1{F0: 0, F1: 42, F2: -122}
                                            v55 := GenS4{F0: 415, F1: "\n\t", F2: 42}
                                            _ = fi53
                                            _ = v54
                                            _ = v55
                                        }
                                    case "\n\t":
                                        if v5 {
                                            v5 = (v5 && (!v5))
                                        } else {
                                            if (2 >= 662) {
                                                var v56 uint64 = 42
                                                v1 = (v1 * v1)
                                                v1 = 65535
                                                _ = v56
                                            }
                                        }
                                    case "foo bar":
                                        for fi57 := 0; fi57 < 4; fi57++ {
                                            var v58 float32 = v3
                                            if v5 {
                                                var v59 uint32 = (v1 ^ (v1 & v1))
                                                fmt.Println(fi52, v58, v1)
                                                fmt.Println(v58, v58)
                                                _ = v59
                                            } else {
                                                var v60 int8 = -1
                                                if v5 {
                                                    if len(v2) > 0 {
                                                        v2[0] = "\n\t"
                                                    }
                                                    var v61 int32 = -2147483648
                                                    var v62 bool = (!((float64(1.5) * float64(v58)) != float64(v3)))
                                                    _ = v61
                                                    _ = v62
                                                }
                                                var v63 int16 = -34
                                                _ = v60
                                                _ = v63
                                            }
                                            v64 := GenS4{F0: 0, F1: ("\n\t" + "a"), F2: 42}
                                            var v65 int8 = 12
                                            _ = fi57
                                            _ = v58
                                            _ = v64
                                            _ = v65
                                        }
                                    default:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    for fi66 := 0; fi66 < 10; fi66++ {
                                        var v67 uint16 = 65535
                                        v68 := GenS2{F0: 1, F1: -97, F2: (fi66 + 128)}
                                        if (v5 || v5) {
                                            switch (fi66 & fi66) {
                                            case 32768:
                                                if true {
                                                    var v69 int8 = 127
                                                    v5 = false
                                                    _ = v69
                                                } else {
                                                    var v70 int8 = 42
                                                    if (v70 > -48) {
                                                        v70 += 42
                                                    } else {
                                                        v1 -= (v1 & 2147483648)
                                                        var v71 int16 = -796
                                                        switch "foo bar" {
                                                        case "hello":
                                                            var v72 int = fi52
                                                            _ = v72
                                                        case "":
                                                            fmt.Println(fi52)
                                                        case "a":
                                                            if (-9223372036854775808 == 0) {
                                                                var v73 uint16 = 42
                                                                v70 -= v70
                                                                fmt.Println(v73)
                                                                fmt.Println(fi52, fi66)
                                                                if (false || v5) {
                                                                    v74 := GenS1{F0: 0, F1: 2, F2: (v70 << ((0) & 63))}
                                                                    v1 = 1
                                                                    v75 := v4[v1]
                                                                    v5 = v5
                                                                    v76 := GenS5{F0: v71, F1: int64(v75)}
                                                                    _ = v74
                                                                    _ = v75
                                                                    _ = v76
                                                                } else {
                                                                    v67 += (v73 % (((v67 - v73)) | 1))
                                                                }
                                                                switch "a" {
                                                                case "":
                                                                    fmt.Println(v73)
                                                                case "foo bar":
                                                                    v70 -= (v70 + (v70 | -1))
                                                                case "hello":
                                                                    switch fi52 {
                                                                    case 182:
                                                                        if v5 {
                                                                            v73 += v67
                                                                            var v77 uint = 566
                                                                            _ = v77
                                                                        }
                                                                    case 32767:
                                                                        var v78 uint = 0
                                                                        _ = v78
                                                                    }
                                                                }
                                                                if (18446744073709551615 == 118) {
                                                                    v71 += (v71 % ((v71) | 1))
                                                                }
                                                                var v79 float32 = v3
                                                                _ = v73
                                                                _ = v79
                                                            } else {
                                                                if ((false && v5) || v5) {
                                                                    var v80 string = ((("" + "a") + ("hello" + "foo bar")) + ("\n\t" + ""))
                                                                    v67 = (v67 | (v67 + v67))
                                                                    _ = v80
                                                                } else {
                                                                    var v81 string = ("a" + ("\n\t" + ("a" + "foo bar")))
                                                                    _ = v81
                                                                }
                                                            }
                                                        }
                                                        _ = v71
                                                    }
                                                    v5 = v5
                                                    _ = v70
                                                }
                                            case 127:
                                                if ((-1 > 0) && (v5 || v5)) {
                                                    v5 = (!v5)
                                                    v5 = v5
                                                    v5 = true
                                                    var v82 float64 = (float64(1.5) + 1.5)
                                                    _ = v82
                                                } else {
                                                    v67 -= v67
                                                    v83 := GenS6{F0: 9223372036854775807, F1: -1, F2: v5}
                                                    _ = v83
                                                }
                                                fmt.Println(v67)
                                            default:
                                                v3 += (v3 * float32(0.0))
                                            }
                                            if v5 {
                                                var v84 uint64 = 1
                                                if len(v2) > 0 {
                                                    v2[0] = "\n\t"
                                                }
                                                var v85 bool = (!(!((122 > 0) || true)))
                                                var v86 int64 = 533
                                                _ = v84
                                                _ = v85
                                                _ = v86
                                            }
                                            v4[(v1 / (((v1 | v1)) | 1))] = fi52
                                        } else {
                                            v1 = 2147483647
                                            fmt.Println(v1, v1)
                                            var v87 int64 = -9223372036854775808
                                            _ = v87
                                        }
                                        v88 := GenS7{F0: 32767, F1: v1, F2: v67}
                                        var v89 int16 = -1
                                        _ = fi66
                                        _ = v67
                                        _ = v68
                                        _ = v88
                                        _ = v89
                                    }
                                    v90 := GenS8{F0: 153, F1: v1}
                                    _ = fi52
                                    _ = v90
                                }
                            default:
                                if true {
                                    if false {
                                        fmt.Println(v5, v5, v5)
                                        v91 := GenS7{F0: -188, F1: v1, F2: 32767}
                                        _ = v91
                                    }
                                    if false {
                                        if (v5 && (true && (-45 <= 42))) {
                                            if ((v5 || v5) && v5) {
                                                var v92 bool = (1e10 == 3.14159)
                                                v3 += v3
                                                _ = v92
                                            }
                                        } else {
                                            v93 := GenS9{F0: -146}
                                            for fi94 := 0; fi94 < 15; fi94++ {
                                                var v95 int32 = 2147483647
                                                for fi96 := 0; fi96 < 4; fi96++ {
                                                    v97 := GenS10{F0: -1, F1: v5, F2: v1}
                                                    _ = fi96
                                                    _ = v97
                                                }
                                                if (115 != 1) {
                                                    var v98 uint64 = 1
                                                    v98 = 1
                                                    _ = v98
                                                }
                                                _ = fi94
                                                _ = v95
                                            }
                                            v99 := GenS9{F0: 633}
                                            _ = v93
                                            _ = v99
                                        }
                                        for fi100 := 0; fi100 < 7; fi100++ {
                                            var v101 int8 = -35
                                            v102 := GenS11{F0: "a", F1: (^(v1)), F2: 255}
                                            _ = fi100
                                            _ = v101
                                            _ = v102
                                        }
                                        fmt.Println(v1, v5)
                                    }
                                } else {
                                    for fi103 := 0; fi103 < 1; fi103++ {
                                        var v104 uint32 = (v1 >> ((586) & 63))
                                        switch "hello" {
                                        case "\n\t":
                                            fmt.Println(v104, v5, fi103)
                                        case "hello":
                                            if v5 {
                                                fmt.Println(fi103, v5)
                                            }
                                        case "":
                                            v104 = v104
                                        default:
                                            _ = len(v4)
                                        }
                                        _ = fi103
                                        _ = v104
                                    }
                                    var v105 int16 = -613
                                    _ = v105
                                }
                            }
                        default:
                            switch "" {
                            case "":
                                var v106 float64 = 0.0
                                _ = v106
                            case "hello":
                                switch "" {
                                case "":
                                    v5 = ((0.0 == (1.5 - float64((1.5 * float32(0.0))))) && v5)
                                }
                            case "foo bar":
                                v1 += v1
                            }
                        }
                        v3 = float32((1.5 * (0.1 + 1e10)))
                        var v107 int16 = -934
                        _ = v107
                    }
                    for fi108 := 0; fi108 < 19; fi108++ {
                        v109 := GenS12{F0: true, F1: 1, F2: 42}
                        _ = fi108
                        _ = v109
                    }
                    fmt.Println(v5, v3)
                }
            }
        } else {
            delete(v4, (^(v1)))
            var v110 uint64 = 425
            for fi111 := 0; fi111 < 15; fi111++ {
                var v112 int32 = -676
                if len(v2) > 0 {
                    v2[0] = ("a" + ("a" + ""))
                }
                for fi113 := 0; fi113 < 20; fi113++ {
                    var v114 uint = 1
                    var v115 uint8 = 244
                    _ = fi113
                    _ = v114
                    _ = v115
                }
                _ = fi111
                _ = v112
            }
            _ = v110
        }
        for fi116 := 0; fi116 < 16; fi116++ {
            fmt.Println(v5)
            if (!(!v5)) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            switch (fi116 << ((481) & 63)) {
            case 1:
                v117 := GenS13{F0: (fi116 - 4294967295), F1: v5, F2: (-(fi116))}
                _ = v117
            case 128:
                _ = len(v4)
            default:
                var v118 int8 = 100
                _ = v118
            }
            _ = fi116
        }
        if v5 {
            for fi119 := 0; fi119 < 4; fi119++ {
                var v120 uint16 = 32767
                fmt.Println(v1, v5, fi119)
                _ = fi119
                _ = v120
            }
        } else {
            for fi121 := 0; fi121 < 11; fi121++ {
                delete(v4, (v1 & (^(v1))))
                _ = fi121
            }
            v1 = (v1 % (((v1 ^ v1)) | 1))
            switch "" {
            case "\n\t":
                var v122 uint16 = 1
                _ = v122
            }
        }
    } else {
        switch 128 {
        case 2147483647:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 32768:
            v123 := GenS9{F0: 42}
            _ = v123
        }
        if (1 != 42) {
            switch -2147483648 {
            case 32767:
                fmt.Println(v5, v1, v5)
            case -719:
                for fi124 := 0; fi124 < 11; fi124++ {
                    var v125 float32 = v3
                    _ = fi124
                    _ = v125
                }
            }
            v5 = v5
        } else {
            for fi126 := 0; fi126 < 12; fi126++ {
                switch ("foo bar" + (("hello" + "foo bar") + ("\n\t" + "a"))) {
                case "a":
                    switch "foo bar" {
                    case "\n\t":
                        if false {
                            fmt.Println(v3, v1, v1)
                            if ((!true) && v5) {
                                v3 = v3
                            }
                            v5 = (v5 || ((("hello" < "hello") && (((!v5) && (1 != 42)) || false)) || (v3 <= (v3 + 1.5))))
                        }
                    }
                }
                _ = fi126
            }
        }
        fmt.Println(v1)
    }
    v3 = v3
    for fi127 := 0; fi127 < 3; fi127++ {
        for fi128 := 0; fi128 < 12; fi128++ {
            v3 -= ((v3 + (1e10 / float32((1.5 - 0.1)))) - ((1.5 + (v3 - (v3 - v3))) + v3))
            switch ((((("\n\t" + "\n\t") + "\n\t") + "") + ("hello" + "")) + "foo bar") {
            case "\n\t":
                fmt.Println(v1)
            }
            _ = fi128
        }
        _ = fi127
    }
    var v129 uint = 738
    var v130 bool = false
    var v131 float64 = (float64(v3) / 1.5)
    if (v5 && v130) {
        v131 = (v131 * ((3.14159 * v131) * v131))
        v132 := v4[(v1 & v1)]
        _ = v132
    } else {
        v5 = v130
        v131 -= 1e10
    }
    fmt.Println(v5, v129)
    if true {
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v133 float32 = (float32(1e10) + v3)
            _ = v133
        } else {
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v131)
            case "foo bar":
                var v134 uint32 = (v1 | 0)
                _ = v134
            case "":
                var v135 string = ((("hello" + ((("a" + ("\n\t" + "foo bar")) + ("hello" + "hello")) + "foo bar")) + (("a" + ("hello" + "hello")) + "")) + ("a" + "\n\t"))
                _ = v135
            default:
                for fi136 := 0; fi136 < 18; fi136++ {
                    var v137 uint = (v129 & (v129 + 18446744073709551615))
                    _ = fi136
                    _ = v137
                }
            }
            fmt.Println(v129)
            switch ("" + ("" + ("\n\t" + "\n\t"))) {
            case "\n\t":
                var v138 uint64 = 18446744073709551615
                _ = v138
            }
        }
    } else {
        v4[(v1 & 2)] = 9223372036854775807
    }
    v131 = 0.1
    if (!false) {
        if v130 {
            fmt.Println(v130, v130, v1)
        }
    }
    if v5 {
        if (42 != 542) {
            v139 := GenS7{F0: 32767, F1: (v1 - v1), F2: 42}
            _ = v139
        } else {
            v5 = true
        }
        v140 := GenS4{F0: 1000, F1: "foo bar", F2: 42}
        var v141 int16 = 552
        _ = v140
        _ = v141
    } else {
        if (!v130) {
            v1 = v1
        } else {
            if v130 {
                v142 := GenS14{F0: 627, F1: 0}
                v143 := GenS5{F0: 0, F1: 42}
                _ = v142
                _ = v143
            }
            if (-2.25 != 3.14159) {
                v130 = (-2147483648 < 42)
                if true {
                    v144 := GenS15{F0: (32767 > 328), F1: 127}
                    for fi145 := 0; fi145 < 20; fi145++ {
                        v131 -= (((((0.0 + ((-2.25 / 1.0) / 1.0)) - (1.5 / (((v131 - v131) - 1e10) * 0.0))) - (v131 / v131)) - (((-2.25 - v131) * (float64(3.14159) * v131)) - (v131 / float64((((0.1 * v3) / (((((v3 - (float32(float64((1.5 * -2.25))) / v3)) + 3.14159) / (v3 - v3)) - (v3 - 3.14159)) / v3)) / 1e10))))) / 1.0)
                        _ = fi145
                    }
                    if (!v5) {
                        var v146 uint = (v129 | v129)
                        _ = v146
                    } else {
                        if (2147483648 < (^(v1))) {
                            var v147 int8 = -26
                            _ = v147
                        }
                        var v148 uint64 = 393
                        _ = v148
                    }
                    _ = v144
                }
                var v149 bool = v130
                _ = v149
            } else {
                for fi150 := 0; fi150 < 5; fi150++ {
                    switch "\n\t" {
                    case "\n\t":
                        if (!((-499 > 2147483647) && (!(v129 > (^(v129)))))) {
                            v3 = (v3 * (v3 - 1e10))
                        } else {
                            var v151 uint16 = 2
                            for fi152 := 0; fi152 < 11; fi152++ {
                                v1 = uint32(v129)
                                v131 = v131
                                fmt.Println(v130, v1)
                                _ = fi152
                            }
                            _ = v151
                        }
                    case "foo bar":
                        if false {
                            switch fi150 {
                            case 2147483647:
                                if len(v2) > 0 {
                                    v2[0] = "foo bar"
                                }
                            }
                            for fi153 := 0; fi153 < 15; fi153++ {
                                v1 += (^(v1))
                                fmt.Println(v130, v131, v1)
                                _ = fi153
                            }
                        }
                    default:
                        var v154 int8 = 100
                        _ = v154
                    }
                    v1 = v1
                    _ = fi150
                }
                v155 := GenS16{F0: v129}
                v4[v1] = 0
                _ = v155
            }
        }
        for fi156 := 0; fi156 < 13; fi156++ {
            for fi157 := 0; fi157 < 3; fi157++ {
                v131 = v131
                v3 = (((v3 - v3) - (v3 - (v3 - v3))) - v3)
                v158 := v4[(v1 ^ (v1 * v1))]
                var v159 int64 = -1
                var v160 uint16 = 0
                _ = fi157
                _ = v158
                _ = v159
                _ = v160
            }
            if v130 {
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
                v161 := GenS12{F0: false, F1: 0, F2: -99}
                fmt.Println(v3, v129)
                _ = v161
            }
            v162 := GenS2{F0: 560, F1: -69, F2: fi156}
            _ = fi156
            _ = v162
        }
    }
    v5 = (!(1 != 1))
    fmt.Println("v1:", v1)
    fmt.Println("v129:", v129)
    fmt.Println("v130:", v130)
    fmt.Println("v131:", v131)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v129
    _ = v130
    _ = v131
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
