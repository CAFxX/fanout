package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint; F1 uint32; F2 int16 }
type GenS6 struct { F0 uint64; F1 int }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint16; F1 int64 }
type GenS10 struct { F0 uint32; F1 bool }
type GenS11 struct { F0 int64; F1 uint64 }
type GenS12 struct { F0 int; F1 bool; F2 int }
type GenS13 struct { F0 uint64; F1 int }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 uint64; F1 int8 }
type GenS18 struct { F0 uint64 }
type GenS19 struct { F0 int64; F1 int32; F2 uint8 }
type GenS20 struct { F0 int; F1 uint32 }

func main() {
    var v1 uint8 = 232
    v2 := []int{127, -331}
    v3 := make(map[string]bool)
    var v4 uint64 = 18446744073709551615
    var v5 int32 = 2147483647
    if true {
        for fi6 := 0; fi6 < 9; fi6++ {
            v1 += 2
            v1 += (v1 & (v1 >> ((42) & 63)))
            _ = fi6
        }
        if (18446744073709551615 <= 18446744073709551615) {
            for fi7 := 0; fi7 < 19; fi7++ {
                v1 = v1
                _ = fi7
            }
        } else {
            v5 -= v5
            fmt.Println(v1, v4, v5)
            fmt.Println(v5, v1, v1)
        }
    } else {
        v4 += v4
        if false {
            var v8 uint8 = v1
            _ = v8
        }
    }
    if false {
        v1 -= 170
    }
    switch 127 {
    case -831:
        switch "" {
        case "foo bar":
            v5 = (v5 >> ((0) & 63))
        }
    case 127:
        for fi9 := 0; fi9 < 12; fi9++ {
            v10 := GenS1{F0: -128, F1: (fi9 >> ((65535) & 63))}
            _ = fi9
            _ = v10
        }
    }
    v11 := GenS2{F0: 1, F1: (v1 >> ((275) & 63))}
    for fi12 := 0; fi12 < 12; fi12++ {
        switch -5 {
        case 2147483647:
            for fi13 := 0; fi13 < 15; fi13++ {
                if true {
                    v14 := GenS1{F0: -55, F1: (fi12 >> ((65535) & 63))}
                    _ = v14
                }
                _ = fi13
            }
        default:
            fmt.Println(v5)
        }
        for fi15 := 0; fi15 < 16; fi15++ {
            var v16 uint16 = 0
            _ = fi15
            _ = v16
        }
        _ = fi12
    }
    switch "a" {
    case "":
        var v17 int = 0
        _ = v17
    case "hello":
        v18 := GenS1{F0: 1, F1: 104}
        _ = v18
    }
    switch -2147483648 {
    case -405:
        v19 := GenS2{F0: (v5 ^ (v5 | 42)), F1: (^(v1))}
        _ = v19
    case 2147483647:
        if (2147483648 != 210) {
            fmt.Println(v4, v5, v5)
            v4 = (v4 % ((v4) | 1))
            delete(v3, "foo bar")
        } else {
            v1 = v1
            if (512 >= 42) {
                switch 32767 {
                case -9223372036854775808:
                    var v20 float32 = (0.1 - (3.14159 * (0.1 * 3.14159)))
                    _ = v20
                case 853:
                    var v21 float32 = 0.0
                    _ = v21
                }
                var v22 uint64 = (v4 >> ((42) & 63))
                _ = v22
            }
            var v23 uint = 0
            _ = v23
        }
    case 9223372036854775807:
        var v24 float32 = (1e10 / float32(-2.25))
        _ = v24
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println(v4, v4)
    for fi25 := 0; fi25 < 9; fi25++ {
        switch (fi25 & (^(fi25))) {
        case 36:
            fmt.Println(v5, v5)
        case 1:
            switch -9223372036854775808 {
            case 1:
                for fi26 := 0; fi26 < 10; fi26++ {
                    switch "\n\t" {
                    case "a":
                        if (592 < 255) {
                            v27 := GenS3{F0: 42, F1: 42}
                            fmt.Println(fi25, fi25)
                            _ = v27
                        } else {
                            var v28 uint8 = v1
                            switch fi26 {
                            case -819:
                                _ = len(v2)
                            }
                            fmt.Println(v4, v4, v28)
                            _ = v28
                        }
                    default:
                        v5 += v5
                    }
                    switch "\n\t" {
                    case "\n\t":
                        if false {
                            v4 -= 411
                            if true {
                                switch (("foo bar" + "hello") + "\n\t") {
                                case "":
                                    switch (-(fi26)) {
                                    case -9223372036854775808:
                                        var v29 uint64 = 42
                                        _ = v29
                                    }
                                case "\n\t":
                                    var v30 uint64 = v4
                                    _ = v30
                                default:
                                    v31 := GenS2{F0: v5, F1: v1}
                                    _ = v31
                                }
                                v4 -= 754
                            } else {
                                v4 = 510
                            }
                        }
                    case "foo bar":
                        v32 := v3["hello"]
                        _ = v32
                    default:
                        var v33 int16 = 0
                        _ = v33
                    }
                    v4 = (v4 * 4294967295)
                    fmt.Println(v1)
                    v5 = (v5 / (((v5 << ((796) & 63))) | 1))
                    fmt.Println(fi25, fi25)
                    _ = fi26
                }
            case -9223372036854775808:
                for fi34 := 0; fi34 < 4; fi34++ {
                    var v35 float64 = ((-2.25 - ((float64(float32(0.1)) * (0.0 * 1.5)) - 0.1)) - 1e10)
                    switch ((("\n\t" + (("a" + "hello") + ("foo bar" + "a"))) + "") + ("" + ("" + "foo bar"))) {
                    case "\n\t":
                        var v36 uint16 = 255
                        _ = v36
                    case "a":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    v35 = -2.25
                    _ = fi34
                    _ = v35
                }
            default:
                v1 = (^(v1))
            }
        case 0:
            v37 := v3["foo bar"]
            _ = v37
        }
        if false {
            var v38 bool = (v5 >= -396)
            if (1 < v1) {
                v4 += 665
                v39 := GenS4{F0: 0}
                _ = v39
            }
            for fi40 := 0; fi40 < 13; fi40++ {
                if v38 {
                    fmt.Println(fi25, fi25, v38)
                    fmt.Println(fi25, v1, v1)
                } else {
                    var v41 float32 = -2.25
                    v42 := GenS1{F0: 0, F1: 2147483647}
                    _ = v41
                    _ = v42
                }
                fmt.Println(v4)
                v43 := GenS3{F0: 1, F1: int64(fi40)}
                _ = fi40
                _ = v43
            }
            _ = v38
        } else {
            v44 := GenS3{F0: -19, F1: -1}
            for fi45 := 0; fi45 < 14; fi45++ {
                v5 = 198
                _ = fi45
            }
            _ = v44
        }
        _ = len(v3)
        _ = fi25
    }
    var v46 uint32 = 418
    v46 += 2147483648
    switch -9223372036854775808 {
    case 32768:
        v4 += v4
    case 167:
        fmt.Println(v5, v46)
    }
    if true {
        if false {
            v47 := v3["foo bar"]
            var v48 uint16 = 65535
            v4 = (v4 - 0)
            _ = v47
            _ = v48
        }
        v49 := GenS5{F0: 4294967295, F1: v46, F2: int16(v4)}
        switch "hello" {
        case "a":
            var v50 uint64 = v4
            _ = v50
        }
        _ = v49
    }
    switch 0 {
    case -985:
        if false {
            var v51 uint64 = (v4 - 99)
            v52 := GenS1{F0: int8(v51), F1: 4294967295}
            for fi53 := 0; fi53 < 6; fi53++ {
                for fi54 := 0; fi54 < 5; fi54++ {
                    if (823 > 2) {
                        v3["\n\t"] = false
                        _ = len(v2)
                        v55 := GenS3{F0: 122, F1: 0}
                        _ = v55
                    }
                    if false {
                        if false {
                            var v56 uint64 = 193
                            v57 := GenS3{F0: 0, F1: 230}
                            _ = v56
                            _ = v57
                        } else {
                            if true {
                                v3["foo bar"] = (0 <= v4)
                            }
                            var v58 int8 = -128
                            if false {
                                if true {
                                    v59 := GenS4{F0: -128}
                                    var v60 int64 = 0
                                    v5 = (v5 >> ((255) & 63))
                                    var v61 uint8 = (v1 & (v1 | uint8(v46)))
                                    _ = v59
                                    _ = v60
                                    _ = v61
                                }
                                if (65535 < 32767) {
                                    if (!true) {
                                        var v62 uint32 = (v46 % (((^(v46))) | 1))
                                        _ = v62
                                    }
                                    fmt.Println(v51)
                                    var v63 uint16 = 1
                                    _ = v63
                                }
                                v4 += (^(v51))
                            }
                            var v64 int32 = (v5 ^ -1)
                            _ = v58
                            _ = v64
                        }
                        _ = len(v2)
                        if true {
                            switch (fi53 >> ((65535) & 63)) {
                            case 832:
                                v65 := GenS6{F0: v51, F1: -2147483648}
                                fmt.Println(v4, fi54, v1)
                                _ = v65
                            case 0:
                                if true {
                                    fmt.Println(v46, v5)
                                    var v66 int = 32767
                                    switch (fi53 + v66) {
                                    case 9223372036854775807:
                                        v1 += v1
                                    case 1:
                                        if len(v2) > 0 {
                                            v2[0] = 680
                                        }
                                    case -699:
                                        v67 := GenS5{F0: 255, F1: (^(v46)), F2: 42}
                                        _ = v67
                                    }
                                    _ = v66
                                } else {
                                    v68 := GenS2{F0: v5, F1: 42}
                                    if (true && true) {
                                        if (true && (1e10 > 1.5)) {
                                            v69 := GenS7{F0: 915}
                                            var v70 int = fi54
                                            _ = v69
                                            _ = v70
                                        }
                                    }
                                    _ = v68
                                }
                            }
                            if true {
                                v1 = (v1 % ((153) | 1))
                            } else {
                                fmt.Println(v51, v4, v51)
                                if (1 != v46) {
                                    var v71 uint16 = 65535
                                    var v72 int16 = -1
                                    var v73 int = 4294967295
                                    _ = v71
                                    _ = v72
                                    _ = v73
                                }
                            }
                            var v74 int64 = 42
                            _ = v74
                        } else {
                            var v75 int16 = 0
                            _ = v75
                        }
                        v76 := GenS8{F0: 127}
                        _ = v76
                    } else {
                        v77 := GenS3{F0: -128, F1: -268}
                        var v78 uint16 = 384
                        var v79 int64 = -225
                        v5 -= -2147483648
                        _ = v77
                        _ = v78
                        _ = v79
                    }
                    switch 134 {
                    case -24:
                        v80 := GenS9{F0: uint16(v4), F1: -27}
                        v81 := GenS3{F0: -1, F1: 939}
                        v82 := GenS3{F0: 127, F1: 952}
                        var v83 float64 = 1e10
                        _ = v80
                        _ = v81
                        _ = v82
                        _ = v83
                    default:
                        switch fi53 {
                        case 2147483647:
                            switch -2147483648 {
                            case 32767:
                                switch fi54 {
                                case -118:
                                    v84 := GenS10{F0: (v46 << ((2) & 63)), F1: true}
                                    _ = v84
                                case 2147483647:
                                    switch ("\n\t" + "foo bar") {
                                    case "a":
                                        v85 := GenS1{F0: -87, F1: fi54}
                                        _ = v85
                                    case "\n\t":
                                        v86 := v3["foo bar"]
                                        _ = v86
                                    }
                                case 42:
                                    if (!true) {
                                        if ((4294967295 < 1) && true) {
                                            if ((v46 <= (^(v46))) && (false && true)) {
                                                v51 = v4
                                                if true {
                                                    if (!false) {
                                                        v46 = 2147483648
                                                    }
                                                } else {
                                                    switch fi53 {
                                                    case 32:
                                                        if ((^(v1)) < 127) {
                                                            var v87 string = "a"
                                                            switch fi54 {
                                                            case -1:
                                                                fmt.Println(v1)
                                                            case 4294967295:
                                                                switch (fi54 << ((4294967295) & 63)) {
                                                                case 32767:
                                                                    if true {
                                                                        if ((v5 % ((-218) | 1)) == (v5 ^ -656)) {
                                                                            if len(v2) > 0 {
                                                                                v2[0] = -2147483648
                                                                            }
                                                                            _ = len(v3)
                                                                            v46 = v46
                                                                        }
                                                                    }
                                                                case 127:
                                                                    v1 = uint8(v46)
                                                                default:
                                                                    v88 := v3[(((v87 + "a") + (("foo bar" + ("\n\t" + ((v87 + v87) + ""))) + "a")) + "\n\t")]
                                                                    _ = v88
                                                                }
                                                            case 9223372036854775807:
                                                                v46 -= (v46 >> ((255) & 63))
                                                            default:
                                                                switch 32767 {
                                                                case -2147483648:
                                                                    v87 = ((v87 + v87) + "foo bar")
                                                                default:
                                                                    var v89 bool = false
                                                                    _ = v89
                                                                }
                                                                v2 = append(v2, -9223372036854775808)
                                                                v90 := GenS5{F0: 1, F1: uint32(v4), F2: 0}
                                                                _ = v90
                                                            }
                                                            _ = v87
                                                        } else {
                                                            switch fi53 {
                                                            case -972:
                                                                var v91 uint32 = (v46 & v46)
                                                                _ = v91
                                                            }
                                                            var v92 bool = ((v46 * 2147483648) == (v46 >> ((0) & 63)))
                                                            _ = v92
                                                        }
                                                    }
                                                }
                                            }
                                            _ = len(v2)
                                            v5 += v5
                                            v3["hello"] = true
                                            if false {
                                                fmt.Println(v46, v1, v46)
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                        }
                                    }
                                default:
                                    fmt.Println(fi54, v1)
                                }
                            }
                        case -1:
                            switch 128 {
                            case 1:
                                var v93 string = "\n\t"
                                _ = v93
                            default:
                                v94 := GenS8{F0: int8(v4)}
                                _ = v94
                            }
                        case 32767:
                            var v95 float64 = float64(-2.25)
                            _ = v95
                        default:
                            if false {
                                var v96 int16 = 377
                                _ = len(v3)
                                var v97 uint64 = 39
                                _ = v96
                                _ = v97
                            } else {
                                var v98 string = "a"
                                if (!true) {
                                    v4 = v51
                                }
                                _ = v98
                            }
                        }
                    }
                    _ = fi54
                }
                var v99 uint8 = (v1 * 255)
                if (4294967295 == 757) {
                    fmt.Println(v46)
                    var v100 string = ""
                    _ = v100
                }
                _ = fi53
                _ = v99
            }
            _ = v51
            _ = v52
        }
    case -9223372036854775808:
        v101 := GenS11{F0: -798, F1: 42}
        _ = v101
    case 2147483647:
        fmt.Println(v5, v46)
    default:
        for fi102 := 0; fi102 < 12; fi102++ {
            var v103 string = ("hello" + ("hello" + "hello"))
            _ = fi102
            _ = v103
        }
    }
    for fi104 := 0; fi104 < 14; fi104++ {
        for fi105 := 0; fi105 < 4; fi105++ {
            if true {
                v106 := GenS4{F0: -109}
                var v107 float64 = ((0.0 / (-2.25 - 0.0)) + float64((float32(0.1) * 1e10)))
                _ = v106
                _ = v107
            } else {
                switch "foo bar" {
                case "hello":
                    v1 -= v1
                    v108 := GenS1{F0: -51, F1: int(v5)}
                    v109 := GenS11{F0: 1, F1: v4}
                    _ = v108
                    _ = v109
                case "\n\t":
                    _ = len(v3)
                case "":
                    _ = len(v2)
                }
                v110 := GenS11{F0: 1, F1: (v4 % ((v4) | 1))}
                _ = v110
            }
            if false {
                v3[(("\n\t" + ("foo bar" + ("foo bar" + "foo bar"))) + (("a" + "") + "hello"))] = ("foo bar" < "\n\t")
                if (42 < -1) {
                    var v111 string = "hello"
                    var v112 int8 = int8(fi105)
                    var v113 int8 = v112
                    v114 := GenS12{F0: 42, F1: true, F2: (fi105 - 128)}
                    v115 := GenS13{F0: v4, F1: fi104}
                    _ = v111
                    _ = v112
                    _ = v113
                    _ = v114
                    _ = v115
                } else {
                    if (0 < 42) {
                        v116 := GenS5{F0: 65535, F1: uint32(v4), F2: -1}
                        _ = len(v3)
                        _ = v116
                    } else {
                        v117 := GenS14{F0: 32767}
                        v5 = v5
                        v4 = v4
                        _ = v117
                    }
                    switch int(v46) {
                    case 2147483647:
                        var v118 int32 = v5
                        _ = v118
                    }
                    if false {
                        delete(v3, "a")
                    } else {
                        var v119 uint8 = (v1 ^ (v1 + (v1 << ((42) & 63))))
                        switch -9223372036854775808 {
                        case 32768:
                            switch "" {
                            case "":
                                v120 := GenS8{F0: -1}
                                if false {
                                    if true {
                                        v46 = 42
                                        v46 = (v46 / (((v46 << ((uint(v46)) & 63))) | 1))
                                        var v121 int16 = 1
                                        _ = v121
                                    } else {
                                        if (42 <= 140) {
                                            fmt.Println(v4)
                                        } else {
                                            if false {
                                                var v122 int64 = int64(v4)
                                                var v123 uint64 = v4
                                                var v124 int16 = -32768
                                                if false {
                                                    v125 := GenS15{F0: 0}
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    v122 += v122
                                                    _ = v125
                                                } else {
                                                    if len(v2) > 0 {
                                                        v2[0] = (fi104 >> ((42) & 63))
                                                    }
                                                    var v126 int = (^(fi105))
                                                    var v127 int16 = (v124 & v124)
                                                    _ = v126
                                                    _ = v127
                                                }
                                                _ = v122
                                                _ = v123
                                                _ = v124
                                            } else {
                                                v5 += v5
                                                v3["foo bar"] = true
                                                v4 = 1
                                            }
                                            if (!false) {
                                                switch fi105 {
                                                case 128:
                                                    var v128 int16 = 0
                                                    _ = v128
                                                }
                                                v129 := GenS11{F0: -1, F1: v4}
                                                _ = v129
                                            }
                                        }
                                        var v130 int8 = -128
                                        _ = v130
                                    }
                                }
                                _ = v120
                            default:
                                switch (fi104 - 1) {
                                case -335:
                                    v4 -= uint64(v1)
                                }
                            }
                        case 0:
                            if false {
                                v131 := GenS16{F0: 1}
                                _ = v131
                            } else {
                                var v132 int = int(v5)
                                v133 := GenS11{F0: -414, F1: v4}
                                _ = len(v3)
                                _ = v132
                                _ = v133
                            }
                            switch "foo bar" {
                            case "foo bar":
                                fmt.Println(fi104, v119)
                                switch "foo bar" {
                                case "hello":
                                    v1 = uint8(v46)
                                default:
                                    var v134 int8 = 97
                                    _ = v134
                                }
                                var v135 uint32 = (v46 << ((738) & 63))
                                _ = v135
                            case "a":
                                v119 += v119
                            default:
                                var v136 float32 = (1e10 / 1.0)
                                _ = v136
                            }
                        }
                        _ = v119
                    }
                    var v137 uint64 = (v4 ^ (v4 >> ((42) & 63)))
                    var v138 float32 = float32((float64(1e10) * (float64(-2.25) / (-2.25 - 3.14159))))
                    _ = v137
                    _ = v138
                }
                v139 := GenS16{F0: 0}
                _ = v139
            } else {
                v46 = (v46 << ((2) & 63))
                if (true || true) {
                    if true {
                        if (3.14159 == 1e10) {
                            v5 = (v5 % ((v5) | 1))
                            switch (fi105 ^ fi104) {
                            case 0:
                                if (!(721 == 42)) {
                                    var v140 int = (fi105 + fi104)
                                    if false {
                                        if (false && true) {
                                            var v141 uint = 773
                                            v1 -= 42
                                            _ = v141
                                        } else {
                                            _ = len(v3)
                                        }
                                    } else {
                                        var v142 uint = 400
                                        _ = v142
                                    }
                                    switch 2147483647 {
                                    case 2147483647:
                                        if len(v2) > 0 {
                                            v2[0] = fi105
                                        }
                                    case 42:
                                        v143 := GenS12{F0: 591, F1: (uint16(v1) >= 32767), F2: (^(fi105))}
                                        _ = v143
                                    default:
                                        v144 := GenS9{F0: 32767, F1: 9223372036854775807}
                                        _ = v144
                                    }
                                    _ = v140
                                }
                            case -2147483648:
                                fmt.Println(fi104, fi105, v5)
                            case 128:
                                var v145 int16 = -1
                                _ = v145
                            }
                        } else {
                            v3[("\n\t" + "a")] = (!false)
                            switch "foo bar" {
                            case "\n\t":
                                v46 = v46
                            default:
                                v5 += v5
                            }
                        }
                        var v146 float32 = 1.5
                        if (uint32(v4) < v46) {
                            var v147 uint16 = 482
                            var v148 uint64 = (v4 + (v4 % ((763) | 1)))
                            _ = v147
                            _ = v148
                        } else {
                            v146 += (0.0 * (3.14159 / -2.25))
                            switch "hello" {
                            case "\n\t":
                                v149 := GenS3{F0: -34, F1: -591}
                                _ = v149
                            default:
                                v5 = (v5 % ((-545) | 1))
                            }
                        }
                        _ = v146
                    }
                } else {
                    switch ("foo bar" + ("a" + "\n\t")) {
                    case "foo bar":
                        v1 -= v1
                    case "":
                        v4 = (v4 & v4)
                    case "a":
                        if len(v2) > 0 {
                            v2[0] = (fi105 / ((32767) | 1))
                        }
                        if (!(true || ((v1 & (v1 << ((631) & 63))) >= v1))) {
                            if true {
                                v150 := GenS17{F0: (v4 | (v4 >> ((42) & 63))), F1: int8(v4)}
                                switch fi105 {
                                case 4294967295:
                                    if (!false) {
                                        v2 = append(v2, (fi104 >> ((906) & 63)))
                                        fmt.Println(fi104, v5, v5)
                                        var v151 int32 = int32(v4)
                                        _ = v151
                                    }
                                case 127:
                                    v3[""] = false
                                case 32767:
                                    v152 := GenS2{F0: v5, F1: (v1 + v1)}
                                    _ = v152
                                }
                                _ = v150
                            } else {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                            switch "hello" {
                            case "foo bar":
                                switch "" {
                                case "\n\t":
                                    var v153 int = 42
                                    _ = v153
                                }
                            case "\n\t":
                                var v154 int16 = 808
                                _ = v154
                            case "hello":
                                fmt.Println(fi104, fi105, v1)
                                v155 := GenS18{F0: uint64(v46)}
                                _ = v155
                            }
                            v46 += uint32(v4)
                        } else {
                            var v156 int = fi105
                            if false {
                                if false {
                                    if true {
                                        v46 = 0
                                        switch 32767 {
                                        case -958:
                                            fmt.Println(v4)
                                        case 475:
                                            fmt.Println(v4)
                                        default:
                                            switch (fi105 ^ int(v1)) {
                                            case 42:
                                                if (((float64(-2.25) + 0.1) <= ((1e10 / 1.0) * 3.14159)) || true) {
                                                    v46 = uint32(v1)
                                                    v157 := GenS5{F0: 1, F1: v46, F2: 358}
                                                    _ = v157
                                                }
                                            case 0:
                                                switch ("foo bar" + "hello") {
                                                case "\n\t":
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                case "a":
                                                    if true {
                                                        var v158 uint = 0
                                                        _ = v158
                                                    } else {
                                                        fmt.Println(v46, v156)
                                                        switch 128 {
                                                        case -9223372036854775808:
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        case 874:
                                                            if (!false) {
                                                                if (true && true) {
                                                                    v4 -= v4
                                                                    var v159 float32 = -2.25
                                                                    _ = v159
                                                                }
                                                            } else {
                                                                v3[""] = false
                                                            }
                                                        default:
                                                            var v160 float32 = 1.5
                                                            _ = v160
                                                        }
                                                        v3["a"] = (!(false && true))
                                                    }
                                                case "":
                                                    if (1 >= 448) {
                                                        var v161 uint64 = v4
                                                        _ = v161
                                                    } else {
                                                        v4 += uint64(v1)
                                                    }
                                                }
                                            case -837:
                                                var v162 uint = 2
                                                _ = v162
                                            }
                                        }
                                    } else {
                                        v163 := GenS19{F0: -165, F1: (v5 ^ -1), F2: (^(v1))}
                                        var v164 int32 = 2147483647
                                        var v165 int64 = 42
                                        var v166 uint32 = (^(v46))
                                        if len(v2) > 0 {
                                            v2[0] = v156
                                        }
                                        _ = v163
                                        _ = v164
                                        _ = v165
                                        _ = v166
                                    }
                                }
                                v167 := GenS20{F0: -1, F1: 42}
                                v2 = append(v2, (v156 * 127))
                                v156 = v156
                                _ = v167
                            } else {
                                v4 += (v4 >> ((4294967295) & 63))
                                fmt.Println(fi104, fi105, fi105)
                                if false {
                                    v3[("\n\t" + "foo bar")] = false
                                }
                                v168 := GenS3{F0: 92, F1: 9223372036854775807}
                                _ = v168
                            }
                            _ = v156
                        }
                    default:
                        var v169 uint8 = uint8(v4)
                        _ = v169
                    }
                    v46 = (v46 % ((2) | 1))
                    v1 += v1
                }
                var v170 int = (fi104 >> ((18446744073709551615) & 63))
                v171 := GenS8{F0: -1}
                _ = v170
                _ = v171
            }
            _ = fi105
        }
        for fi172 := 0; fi172 < 8; fi172++ {
            fmt.Println(v1)
            if ((!true) || true) {
                var v173 int64 = -9223372036854775808
                _ = v173
            } else {
                var v174 bool = (((int(v1) < 9223372036854775807) && ((v1 >> ((0) & 63)) <= v1)) || false)
                _ = v174
            }
            _ = fi172
        }
        var v175 int64 = 9223372036854775807
        _ = fi104
        _ = v175
    }
    var v176 uint64 = (v4 >> ((342) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v176:", v176)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v176
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
}
