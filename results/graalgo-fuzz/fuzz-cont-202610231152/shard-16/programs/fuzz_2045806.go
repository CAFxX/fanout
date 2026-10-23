package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int64; F1 uint64; F2 bool }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint; F1 uint16 }
type GenS9 struct { F0 int32; F1 uint; F2 int16 }
type GenS10 struct { F0 uint16; F1 string; F2 int32 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 string; F1 bool }
type GenS13 struct { F0 string }

func main() {
    var v1 string = "hello"
    v2 := []int{469, 1}
    v3 := []int{128, 127}
    var v4 int32 = 2147483647
    v5 := []int{}
    v6 := make(map[uint32]uint32)
    if true {
        switch v1 {
        case "a":
            v7 := GenS1{F0: -1, F1: 177}
            _ = v7
        case "hello":
            v4 -= v4
        case "":
            var v8 bool = false
            _ = v8
        default:
            if ((!false) && (1 <= 0)) {
                if len(v5) > 0 {
                    v5[0] = -1
                }
                fmt.Println(v4, v1)
                delete(v6, 2147483648)
            } else {
                var v9 float32 = -2.25
                _ = v9
            }
        }
        if (-206 > 723) {
            var v10 int64 = 0
            v1 = v1
            var v11 uint16 = 2
            _ = v10
            _ = v11
        } else {
            if false {
                v4 = v4
                v6[2147483647] = 191
                fmt.Println(v4, v4)
            }
            for fi12 := 0; fi12 < 2; fi12++ {
                switch "a" {
                case "":
                    for fi13 := 0; fi13 < 17; fi13++ {
                        v6[2] = 2
                        if false {
                            var v14 uint8 = 128
                            _ = v14
                        } else {
                            switch (fi12 / ((fi13) | 1)) {
                            case 32768:
                                if (!(2147483647 != 1)) {
                                    _ = len(v6)
                                    v15 := GenS2{F0: 242}
                                    _ = v15
                                }
                            }
                        }
                        switch 2147483647 {
                        case 9223372036854775807:
                            if (1 >= 114) {
                                fmt.Println(v1, fi13, v1)
                                var v16 uint16 = 32767
                                _ = v16
                            } else {
                                var v17 uint = 981
                                v4 += (v4 * int32(v17))
                                switch (("\n\t" + (v1 + v1)) + (v1 + v1)) {
                                case "\n\t":
                                    switch 0 {
                                    case 128:
                                        var v18 bool = true
                                        v6[415] = 401
                                        _ = v18
                                    }
                                case "foo bar":
                                    var v19 bool = false
                                    _ = v19
                                default:
                                    var v20 uint = (v17 << (((v17 - (v17 + v17))) & 63))
                                    _ = v20
                                }
                                _ = v17
                            }
                        case -1:
                            v6[4294967295] = 2147483648
                            fmt.Println(v4, fi13)
                            var v21 string = (v1 + v1)
                            _ = v21
                        case 60:
                            if (127 <= 1) {
                                if false {
                                    fmt.Println(v4, fi13, fi12)
                                    v6[35] = 65535
                                    v4 += (-(v4))
                                }
                            }
                        }
                        _ = fi13
                    }
                default:
                    if true {
                        v1 = v1
                    }
                }
                _ = fi12
            }
            v22 := GenS1{F0: 32767, F1: 203}
            _ = v22
        }
        v23 := GenS3{F0: 9, F1: "hello"}
        _ = v23
    } else {
        v1 = v1
        fmt.Println(v4, v1)
        for fi24 := 0; fi24 < 1; fi24++ {
            if true {
                var v25 uint32 = 4294967295
                v1 = ((v1 + v1) + "\n\t")
                fmt.Println(v4, v1)
                _ = v25
            } else {
                v26 := GenS1{F0: -1, F1: 2147483648}
                switch v1 {
                case "\n\t":
                    fmt.Println(v1)
                case "":
                    v4 = v4
                case "hello":
                    fmt.Println(fi24)
                }
                for fi27 := 0; fi27 < 9; fi27++ {
                    v1 = v1
                    v1 = "foo bar"
                    v4 = (-(v4))
                    _ = fi27
                }
                _ = v26
            }
            v28 := GenS4{F0: 397, F1: 854, F2: (v4 > (v4 >> ((42) & 63)))}
            switch int(v4) {
            case -2147483648:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case 32767:
                switch (("foo bar" + "foo bar") + v1) {
                case "\n\t":
                    _ = len(v6)
                default:
                    if true {
                        switch 1 {
                        case 2147483647:
                            v4 = v4
                        case -2147483648:
                            if (!(586 < 4294967295)) {
                                var v29 int32 = 2147483647
                                _ = v29
                            }
                        }
                        for fi30 := 0; fi30 < 2; fi30++ {
                            _ = len(v6)
                            v2 = append(v2, fi24)
                            var v31 int8 = 100
                            _ = fi30
                            _ = v31
                        }
                        var v32 uint8 = 1
                        _ = v32
                    } else {
                        var v33 uint16 = 2
                        _ = v33
                    }
                }
            default:
                v1 = (v1 + v1)
            }
            _ = fi24
            _ = v28
        }
    }
    var v34 uint8 = 128
    if false {
        var v35 uint8 = v34
        v36 := GenS5{F0: 653}
        var v37 int = 32767
        _ = v35
        _ = v36
        _ = v37
    } else {
        var v38 string = v1
        v39 := GenS6{F0: 114}
        _ = v38
        _ = v39
    }
    if true {
        var v40 float64 = 0.0
        v41 := GenS7{F0: 1}
        _ = v40
        _ = v41
    }
    var v42 uint8 = v34
    var v43 uint32 = 1
    v6[325] = 4294967295
    for fi44 := 0; fi44 < 10; fi44++ {
        switch (fi44 - fi44) {
        case 2147483647:
            var v45 uint32 = v43
            _ = v45
        case 247:
            switch ("" + "") {
            case "":
                v1 = ""
            case "foo bar":
                for fi46 := 0; fi46 < 16; fi46++ {
                    v47 := GenS8{F0: 218, F1: 255}
                    _ = fi46
                    _ = v47
                }
            }
        default:
            switch fi44 {
            case 32767:
                if (230 == 840) {
                    _ = len(v2)
                }
            case 9223372036854775807:
                if (true || (false || true)) {
                    var v48 uint32 = (v43 << ((269) & 63))
                    for fi49 := 0; fi49 < 13; fi49++ {
                        var v50 uint16 = 42
                        fmt.Println(v42, v4, v48)
                        var v51 float32 = 0.0
                        _ = fi49
                        _ = v50
                        _ = v51
                    }
                    v34 += (v34 - 255)
                    _ = v48
                } else {
                    fmt.Println(v4, v43)
                    if ((18446744073709551615 >= 42) && true) {
                        if (!(false && (!true))) {
                            var v52 int8 = 100
                            _ = v52
                        } else {
                            var v53 float32 = 3.14159
                            _ = v53
                        }
                        switch (v1 + ("\n\t" + "")) {
                        case "":
                            switch 9223372036854775807 {
                            case 4294967295:
                                for fi54 := 0; fi54 < 12; fi54++ {
                                    switch fi44 {
                                    case 4294967295:
                                        switch (fi44 ^ -9223372036854775808) {
                                        case 498:
                                            if (!false) {
                                                v55 := GenS5{F0: 1}
                                                v43 += v43
                                                _ = v55
                                            } else {
                                                var v56 uint64 = 0
                                                fmt.Println(v43)
                                                switch "hello" {
                                                case "":
                                                    v43 += (v43 / (((v43 ^ (v43 / (((v43 * 0)) | 1)))) | 1))
                                                case "\n\t":
                                                    v1 = ((v1 + ("hello" + v1)) + ((v1 + v1) + "\n\t"))
                                                case "hello":
                                                    v42 = (^(v34))
                                                }
                                                v34 = v34
                                                _ = v56
                                            }
                                        default:
                                            var v57 float64 = ((3.14159 - 3.14159) + (3.14159 / 1e10))
                                            _ = v57
                                        }
                                    case -2147483648:
                                        v43 -= (^(v43))
                                    }
                                    var v58 float64 = float64(0.1)
                                    v43 = v43
                                    _ = fi54
                                    _ = v58
                                }
                            case 951:
                                for fi59 := 0; fi59 < 9; fi59++ {
                                    if (!true) {
                                        v60 := GenS4{F0: 962, F1: 4294967295, F2: false}
                                        v34 = 56
                                        _ = v60
                                    } else {
                                        v1 = v1
                                        fmt.Println(v1)
                                        var v61 uint64 = 521
                                        var v62 uint16 = 1
                                        if (v4 != -21) {
                                            if ((^(v4)) == (-(v4))) {
                                                var v63 uint16 = v62
                                                var v64 uint32 = 2147483647
                                                _ = v63
                                                _ = v64
                                            }
                                            v65 := GenS4{F0: 437, F1: v61, F2: false}
                                            _ = v65
                                        } else {
                                            _ = len(v2)
                                            if false {
                                                var v66 float32 = 1.5
                                                if (v62 >= v62) {
                                                    var v67 uint16 = 730
                                                    var v68 float32 = v66
                                                    var v69 float32 = -2.25
                                                    _ = v67
                                                    _ = v68
                                                    _ = v69
                                                } else {
                                                    v42 = (v42 >> ((18446744073709551615) & 63))
                                                    if false {
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    }
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    var v70 int = 42
                                                    _ = v70
                                                }
                                                _ = v66
                                            } else {
                                                v34 -= 127
                                                if (299 != uint32(v34)) {
                                                    var v71 uint16 = (v62 / ((v62) | 1))
                                                    fmt.Println(v61)
                                                    _ = v71
                                                }
                                            }
                                            delete(v6, (v43 % ((998) | 1)))
                                            v72 := GenS7{F0: v43}
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            _ = v72
                                        }
                                        _ = v61
                                        _ = v62
                                    }
                                    _ = fi59
                                }
                            case -9223372036854775808:
                                if (12 == 42) {
                                    var v73 float32 = ((1.5 * ((1.5 + (float32(0.1) - 1e10)) - (3.14159 - float32(-2.25)))) + 1.5)
                                    _ = v73
                                } else {
                                    v43 = (^(v43))
                                }
                            default:
                                if ((fi44 ^ fi44) >= -2147483648) {
                                    switch (fi44 - -2147483648) {
                                    case 32768:
                                        v1 = ((v1 + (("hello" + v1) + v1)) + ((v1 + v1) + "a"))
                                    }
                                    switch 42 {
                                    case -1:
                                        for fi74 := 0; fi74 < 4; fi74++ {
                                            var v75 int = (fi74 | -9223372036854775808)
                                            v76 := v6[(v43 / ((v43) | 1))]
                                            var v77 string = v1
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            _ = fi74
                                            _ = v75
                                            _ = v76
                                            _ = v77
                                        }
                                    case -915:
                                        var v78 uint = 18446744073709551615
                                        _ = v78
                                    default:
                                        if true {
                                            for fi79 := 0; fi79 < 19; fi79++ {
                                                v80 := GenS1{F0: -308, F1: v43}
                                                v1 = v1
                                                v43 = v43
                                                v43 = v43
                                                _ = fi79
                                                _ = v80
                                            }
                                        }
                                    }
                                } else {
                                    if (true && ((false && (!false)) && (false || (-1 >= 9223372036854775807)))) {
                                        if (3.14159 != -2.25) {
                                            var v81 uint32 = 2147483648
                                            for fi82 := 0; fi82 < 3; fi82++ {
                                                v83 := GenS5{F0: 1}
                                                _ = fi82
                                                _ = v83
                                            }
                                            v4 -= 2147483647
                                            _ = v81
                                        }
                                        for fi84 := 0; fi84 < 20; fi84++ {
                                            v1 = v1
                                            fmt.Println(v42)
                                            _ = fi84
                                        }
                                    } else {
                                        var v85 int64 = -669
                                        fmt.Println(v43, v43)
                                        _ = v85
                                    }
                                    fmt.Println(v1, v1)
                                }
                            }
                        case "hello":
                            switch "" {
                            case "hello":
                                switch "" {
                                case "\n\t":
                                    v86 := GenS9{F0: (v4 * 43), F1: 255, F2: -450}
                                    _ = v86
                                case "hello":
                                    if ((v34 % ((v34) | 1)) < (v34 % (((v34 ^ 255)) | 1))) {
                                        if (1.5 == -2.25) {
                                            switch (fi44 << ((42) & 63)) {
                                            case 128:
                                                if len(v2) > 0 {
                                                    v2[0] = -93
                                                }
                                            case 1:
                                                if false {
                                                    for fi87 := 0; fi87 < 13; fi87++ {
                                                        v34 = (v34 | (v42 * v42))
                                                        var v88 float64 = (1e10 + 1e10)
                                                        v42 = 1
                                                        v89 := GenS1{F0: -494, F1: (^(v43))}
                                                        delete(v6, 65535)
                                                        if true {
                                                            var v90 uint32 = 4294967295
                                                            if true {
                                                                switch "a" {
                                                                case "hello":
                                                                    v90 = (v43 - 369)
                                                                case "foo bar":
                                                                    v91 := GenS6{F0: 127}
                                                                    _ = v91
                                                                case "\n\t":
                                                                    var v92 int16 = 98
                                                                    _ = v92
                                                                }
                                                                if ((false || (42 >= 18446744073709551615)) && false) {
                                                                    fmt.Println(v4, v1)
                                                                    delete(v6, (v43 * (v43 & v43)))
                                                                    v88 = ((0.1 / (v88 / 1.5)) - (float64((0.1 - 0.0)) - v88))
                                                                }
                                                            } else {
                                                                v4 = -175
                                                                if (!(!(-2147483648 == v4))) {
                                                                    var v93 int = (fi44 / ((fi87) | 1))
                                                                    _ = v93
                                                                }
                                                                v94 := GenS10{F0: 280, F1: "a", F2: v4}
                                                                _ = v94
                                                            }
                                                            _ = v90
                                                        }
                                                        v34 += v34
                                                        _ = fi87
                                                        _ = v88
                                                        _ = v89
                                                    }
                                                }
                                            case 467:
                                                fmt.Println(v43, v42, v34)
                                            }
                                        } else {
                                            fmt.Println(v34, v1, v4)
                                            if ((false && (2 == 1)) && (3.14159 != 0.1)) {
                                                if (1e10 <= 0.0) {
                                                    v95 := v6[4294967295]
                                                    var v96 string = v1
                                                    _ = v95
                                                    _ = v96
                                                } else {
                                                    switch (fi44 ^ fi44) {
                                                    case -1:
                                                        v97 := GenS11{F0: v4}
                                                        _ = v97
                                                    case -9223372036854775808:
                                                        if true {
                                                            if (false || (!(!true))) {
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                v98 := GenS4{F0: 1, F1: 0, F2: (!(32767 <= 1))}
                                                                _ = v98
                                                            } else {
                                                                switch fi44 {
                                                                case 42:
                                                                    switch fi44 {
                                                                    case 9223372036854775807:
                                                                        if len(v3) > 0 {
                                                                            v3[0] = 42
                                                                        }
                                                                    case 731:
                                                                        switch (-(fi44)) {
                                                                        case 208:
                                                                            var v99 int64 = 0
                                                                            _ = v99
                                                                        case 127:
                                                                            var v100 int8 = -128
                                                                            _ = v100
                                                                        }
                                                                    }
                                                                case 4294967295:
                                                                    for fi101 := 0; fi101 < 20; fi101++ {
                                                                        v102 := GenS12{F0: v1, F1: false}
                                                                        _ = fi101
                                                                        _ = v102
                                                                    }
                                                                case -2147483648:
                                                                    delete(v6, (^(v43)))
                                                                default:
                                                                    for fi103 := 0; fi103 < 13; fi103++ {
                                                                        if (!true) {
                                                                            v34 += v34
                                                                            v4 = 2147483647
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v4)
                                                                        }
                                                                        fmt.Println(fi103)
                                                                        v43 += v43
                                                                        _ = fi103
                                                                    }
                                                                }
                                                                v34 = 42
                                                            }
                                                            var v104 string = v1
                                                            _ = v104
                                                        }
                                                    default:
                                                        fmt.Println(v1, v4)
                                                    }
                                                    v42 = (v42 * (^(v42)))
                                                }
                                                v34 = (^(v42))
                                                var v105 uint32 = 2
                                                _ = v105
                                            }
                                            var v106 int64 = -1
                                            _ = v106
                                        }
                                        v107 := GenS9{F0: -935, F1: 4294967295, F2: -1}
                                        v43 = v43
                                        _ = v107
                                    } else {
                                        var v108 int16 = 1
                                        v109 := GenS11{F0: int32(v43)}
                                        _ = v108
                                        _ = v109
                                    }
                                }
                            case "":
                                var v110 int = (fi44 ^ (fi44 | (fi44 % ((fi44) | 1))))
                                _ = v110
                            default:
                                _ = len(v6)
                            }
                        case "\n\t":
                            var v111 float32 = float32((0.0 / float64(0.1)))
                            _ = v111
                        }
                    }
                    fmt.Println(fi44, v42)
                }
            default:
                var v112 uint8 = (v34 % ((128) | 1))
                _ = v112
            }
        }
        _ = fi44
    }
    switch -728 {
    case -2147483648:
        v42 -= 42
    case 265:
        for fi113 := 0; fi113 < 4; fi113++ {
            if len(v2) > 0 {
                v2[0] = int(v42)
            }
            var v114 int32 = (v4 << ((613) & 63))
            v114 += (v114 & 2147483647)
            _ = fi113
            _ = v114
        }
    case -9223372036854775808:
        var v115 float32 = 3.14159
        _ = v115
    default:
        v2 = append(v2, 32768)
    }
    switch -2147483648 {
    case 4294967295:
        v1 = ""
    }
    v116 := GenS11{F0: v4}
    if (false && true) {
        v5 = append(v5, -9223372036854775808)
        if (9223372036854775807 == int(v43)) {
            var v117 float64 = (((float64(((1e10 - 3.14159) / (float32((float64((-2.25 + float32(-2.25))) - 1e10)) - 0.1))) / -2.25) - float64(0.1)) - 0.0)
            switch ((v1 + "a") + v1) {
            case "hello":
                if true {
                    fmt.Println(v117, v117)
                    fmt.Println(v117, v42, v117)
                    var v118 int16 = 0
                    _ = v118
                } else {
                    fmt.Println(v117)
                    v119 := v6[630]
                    _ = v119
                }
            }
            _ = v117
        }
        if len(v5) > 0 {
            v5[0] = 42
        }
    } else {
        for fi120 := 0; fi120 < 17; fi120++ {
            _ = len(v6)
            var v121 int = (fi120 * fi120)
            _ = fi120
            _ = v121
        }
        switch (v1 + v1) {
        case "\n\t":
            v42 = (v42 << ((2) & 63))
        case "foo bar":
            if true {
                if (0.0 < 1e10) {
                    fmt.Println(v1, v1)
                    var v122 string = ("" + v1)
                    _ = v122
                } else {
                    var v123 int16 = 157
                    _ = v123
                }
                v43 = (v43 % (((v43 * v43)) | 1))
                if (42 < 255) {
                    for fi124 := 0; fi124 < 11; fi124++ {
                        v4 = -715
                        _ = fi124
                    }
                    switch ((("" + "hello") + v1) + "") {
                    case "hello":
                        v43 = (v43 >> ((65535) & 63))
                    case "":
                        v125 := GenS2{F0: 965}
                        _ = v125
                    default:
                        if false {
                            v42 += v34
                        } else {
                            var v126 uint32 = 42
                            var v127 uint = 664
                            _ = v126
                            _ = v127
                        }
                    }
                }
            }
        case "":
            for fi128 := 0; fi128 < 4; fi128++ {
                var v129 int16 = 231
                v130 := GenS2{F0: 42}
                delete(v6, 1)
                _ = fi128
                _ = v129
                _ = v130
            }
        }
        for fi131 := 0; fi131 < 11; fi131++ {
            fmt.Println(v34, v4)
            fmt.Println(v43, v42)
            if (true && false) {
                var v132 int8 = -1
                _ = v132
            } else {
                v34 += (v42 - (^(v42)))
                for fi133 := 0; fi133 < 1; fi133++ {
                    v134 := v6[v43]
                    _ = fi133
                    _ = v134
                }
                var v135 int8 = 97
                _ = v135
            }
            _ = fi131
        }
    }
    for fi136 := 0; fi136 < 9; fi136++ {
        var v137 uint8 = 2
        _ = fi136
        _ = v137
    }
    switch v1 {
    case "":
        v1 = v1
    case "a":
        if (!(!((v1 + v1) <= v1))) {
            var v138 string = v1
            v6[(^(v43))] = v43
            _ = v138
        } else {
            var v139 int32 = 519
            v34 = v42
            if true {
                v140 := GenS13{F0: v1}
                if ((v4 / ((v139) | 1)) <= (v4 - v139)) {
                    var v141 uint32 = uint32(v34)
                    _ = v141
                } else {
                    v42 = v34
                    var v142 float64 = 0.0
                    var v143 uint16 = 27
                    _ = v142
                    _ = v143
                }
                fmt.Println(v43)
                _ = v140
            }
            _ = v139
        }
    }
    switch -2147483648 {
    case 91:
        if false {
            var v144 int16 = int16(v34)
            var v145 uint16 = 439
            v42 = (^(v34))
            _ = v144
            _ = v145
        }
    case 0:
        var v146 int32 = v4
        _ = v146
    default:
        fmt.Println(v43, v4, v34)
    }
    if (1 > 0) {
        switch -714 {
        case -2147483648:
            if (!true) {
                var v147 uint8 = 42
                for fi148 := 0; fi148 < 17; fi148++ {
                    switch (fi148 - (fi148 & fi148)) {
                    case 960:
                        v147 += v42
                    case -1:
                        var v149 int8 = 100
                        _ = v149
                    }
                    if ((v34 | v34) >= 128) {
                        delete(v6, v43)
                    }
                    v150 := GenS3{F0: 53, F1: v1}
                    _ = fi148
                    _ = v150
                }
                _ = v147
            } else {
                fmt.Println(v43, v1)
                v42 += v34
                switch (v1 + ((((v1 + v1) + v1) + "hello") + "")) {
                case "\n\t":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "a":
                    fmt.Println(v34)
                case "":
                    var v151 uint64 = 762
                    _ = v151
                }
            }
        }
        var v152 int = 0
        delete(v6, 42)
        _ = v152
    }
    v43 += (v43 * (v43 - (^(v43))))
    var v153 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v153:", v153)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v116
    _ = v153
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
