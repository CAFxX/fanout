package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS2 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 uint; F1 bool; F2 int16 }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS8 struct { F0 int8; F1 string; F2 uint32 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 string; F1 bool }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int64; F1 uint8; F2 bool }
type GenS13 struct { F0 int16; F1 uint8 }
type GenS14 struct { F0 int64; F1 int }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 uint64; F1 uint32 }
type GenS18 struct { F0 int8; F1 bool }
type GenS19 struct { F0 int16 }
type GenS20 struct { F0 int }
type GenS21 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS22 struct { F0 uint8; F1 int8; F2 int64 }
type GenS23 struct { F0 string; F1 uint16 }
type GenS24 struct { F0 uint }

func main() {
    var v1 int = 4294967295
    var v2 int64 = 42
    v3 := make(map[uint32]int)
    v4 := v3[4294967295]
    v5 := GenS1{F0: 18446744073709551615, F1: -128, F2: 32767}
    fmt.Println(v4, v4, v4)
    v6 := GenS2{F0: -67, F1: 255, F2: 255}
    if (255 <= 255) {
        v4 = v4
        switch "\n\t" {
        case "":
            if false {
                v7 := GenS3{F0: -2147483648, F1: 2}
                var v8 int8 = 100
                for fi9 := 0; fi9 < 16; fi9++ {
                    var v10 uint8 = 1
                    v2 = 349
                    v8 -= int8(v4)
                    _ = fi9
                    _ = v10
                }
                _ = v7
                _ = v8
            }
        case "foo bar":
            for fi11 := 0; fi11 < 11; fi11++ {
                v2 += v2
                v4 = v4
                _ = fi11
            }
        }
        if true {
            var v12 int16 = 1000
            var v13 uint16 = 32767
            var v14 uint8 = 165
            _ = v12
            _ = v13
            _ = v14
        } else {
            v15 := []int{}
            _ = v15
        }
    } else {
        v16 := GenS1{F0: 1, F1: -2147483648, F2: 32767}
        if false {
            v1 += v4
            switch -720 {
            case 840:
                for fi17 := 0; fi17 < 6; fi17++ {
                    if (2147483648 <= 2147483647) {
                        for fi18 := 0; fi18 < 12; fi18++ {
                            switch fi17 {
                            case -709:
                                v1 = 128
                            case 42:
                                switch "\n\t" {
                                case "hello":
                                    v19 := []int{(v1 + v1), fi17, fi18, fi17}
                                    _ = v19
                                case "a":
                                    v20 := []int{v4}
                                    _ = v20
                                }
                            case -850:
                                _ = len(v3)
                            }
                            if true {
                                v21 := GenS1{F0: 67, F1: -2147483648, F2: 37}
                                if false {
                                    v1 = v4
                                    v2 = v2
                                    var v22 float64 = 0.0
                                    v23 := GenS1{F0: 1, F1: 0, F2: 65535}
                                    _ = v22
                                    _ = v23
                                }
                                _ = v21
                            }
                            _ = fi18
                        }
                        v4 -= v4
                    } else {
                        var v24 uint = 42
                        _ = v24
                    }
                    var v25 uint32 = 754
                    _ = len(v3)
                    _ = fi17
                    _ = v25
                }
            case 32768:
                if (0.1 > (((((1e10 / 1.5) * -2.25) * (1.5 + -2.25)) + 1e10) - float32(3.14159))) {
                    var v26 int16 = -831
                    v2 += v2
                    var v27 uint = 1
                    _ = v26
                    _ = v27
                } else {
                    v28 := GenS1{F0: 948, F1: -373, F2: 81}
                    v29 := GenS3{F0: 2147483647, F1: 0}
                    for fi30 := 0; fi30 < 11; fi30++ {
                        var v31 uint32 = 834
                        switch "hello" {
                        case "":
                            v3[979] = (^(v1))
                        case "\n\t":
                            var v32 int64 = -1
                            _ = v32
                        case "hello":
                            v33 := GenS4{F0: 1, F1: (true || (float64(1.5) > float64((0.1 * 3.14159)))), F2: 0}
                            _ = v33
                        default:
                            var v34 uint = 42
                            _ = v34
                        }
                        for fi35 := 0; fi35 < 4; fi35++ {
                            v4 = v1
                            v36 := GenS3{F0: 2147483647, F1: 1}
                            if (-9223372036854775808 != v2) {
                                v2 = v2
                                switch fi30 {
                                case -69:
                                    switch "a" {
                                    case "a":
                                        v37 := GenS5{F0: 739, F1: v2}
                                        _ = v37
                                    case "\n\t":
                                        v38 := GenS6{F0: ("foo bar" + "a")}
                                        _ = v38
                                    case "foo bar":
                                        if false {
                                            if false {
                                                var v39 int64 = (v2 + v2)
                                                v1 += (^(fi30))
                                                var v40 string = (("hello" + "hello") + "")
                                                _ = v39
                                                _ = v40
                                            } else {
                                                v4 -= v1
                                                if (((424 >= 0) || (!(!(false || ((v31 % ((v31) | 1)) > 65535))))) && false) {
                                                    v1 += 9223372036854775807
                                                    var v41 float32 = (1.5 + 1.5)
                                                    if true {
                                                        v42 := []int{}
                                                        if false {
                                                            v31 += v31
                                                        } else {
                                                            v1 = (v4 / ((0) | 1))
                                                        }
                                                        if len(v42) > 0 {
                                                            v42[0] = v4
                                                        }
                                                        _ = v42
                                                    } else {
                                                        _ = len(v3)
                                                    }
                                                    _ = v41
                                                }
                                            }
                                            var v43 uint = 18446744073709551615
                                            _ = v43
                                        }
                                    }
                                case -2147483648:
                                    switch 1 {
                                    case -696:
                                        fmt.Println(v2, v4, v31)
                                    case 2147483647:
                                        v1 -= -1
                                    }
                                case 42:
                                    v44 := []int{32767}
                                    _ = v44
                                }
                                v45 := []int{(-(fi30)), (fi30 >> ((4) & 63))}
                                v46 := GenS4{F0: 2, F1: (42 <= uint(v31)), F2: 1000}
                                _ = v45
                                _ = v46
                            }
                            _ = fi35
                            _ = v36
                        }
                        _ = fi30
                        _ = v31
                    }
                    _ = v28
                    _ = v29
                }
            case 9223372036854775807:
                if (1 <= 2147483647) {
                    v47 := []int{v1, 1}
                    var v48 float32 = 3.14159
                    _ = v47
                    _ = v48
                }
            }
            v2 -= (v2 >> ((351) & 63))
        }
        _ = v16
    }
    v4 -= int(v2)
    v2 += v2
    if true {
        v49 := GenS4{F0: 255, F1: false, F2: 32767}
        if false {
            fmt.Println(v1, v1, v1)
            if true {
                switch (v1 * (v1 / ((v1) | 1))) {
                case 1:
                    var v50 float32 = 3.14159
                    _ = v50
                case 299:
                    fmt.Println(v2, v2)
                }
                var v51 float32 = ((1.5 / 3.14159) + (float32(float64(1.5)) * 1e10))
                _ = v51
            }
            for fi52 := 0; fi52 < 18; fi52++ {
                if (false && ((150 > 42) || false)) {
                    v53 := GenS7{F0: 0, F1: 255, F2: 42}
                    switch (fi52 | (^(v1))) {
                    case 0:
                        switch ("hello" + "\n\t") {
                        case "a":
                            v3[214] = (v4 >> ((255) & 63))
                        case "":
                            var v54 uint64 = 811
                            _ = v54
                        default:
                            var v55 bool = true
                            _ = v55
                        }
                    default:
                        for fi56 := 0; fi56 < 9; fi56++ {
                            var v57 uint = 65535
                            delete(v3, 2147483648)
                            _ = fi56
                            _ = v57
                        }
                    }
                    _ = v53
                }
                _ = fi52
            }
        } else {
            v4 -= (v4 << ((2) & 63))
        }
        if false {
            if ((false && (1000 < -899)) && true) {
                for fi58 := 0; fi58 < 2; fi58++ {
                    for fi59 := 0; fi59 < 11; fi59++ {
                        v60 := v3[875]
                        _ = fi59
                        _ = v60
                    }
                    for fi61 := 0; fi61 < 6; fi61++ {
                        var v62 int32 = -2147483648
                        v63 := []int{}
                        fmt.Println(v2)
                        if (true || (false || false)) {
                            if len(v63) > 0 {
                                v63[0] = (-(v1))
                            }
                        } else {
                            v1 = 418
                            v64 := GenS4{F0: 255, F1: (32767 <= 32767), F2: 42}
                            v62 = (-(v62))
                            _ = v64
                        }
                        v2 = v2
                        var v65 bool = (190 != 661)
                        _ = fi61
                        _ = v62
                        _ = v63
                        _ = v65
                    }
                    var v66 uint32 = 1
                    _ = fi58
                    _ = v66
                }
            } else {
                fmt.Println(v4, v2, v2)
            }
            var v67 uint64 = 2
            _ = len(v3)
            _ = v67
        } else {
            var v68 float64 = 0.1
            fmt.Println(v2, v2)
            _ = v68
        }
        _ = v49
    }
    _ = len(v3)
    var v69 int32 = 42
    v70 := []int{int(v2), v4, v4}
    if false {
        if false {
            v71 := GenS8{F0: 1, F1: ("\n\t" + "a"), F2: 509}
            fmt.Println(v4)
            var v72 uint32 = 2
            _ = v71
            _ = v72
        }
        for fi73 := 0; fi73 < 18; fi73++ {
            for fi74 := 0; fi74 < 11; fi74++ {
                if false {
                    var v75 int16 = -112
                    if true {
                        v76 := GenS6{F0: ("foo bar" + ("\n\t" + "a"))}
                        v77 := GenS9{F0: (true || (357 >= 255))}
                        _ = v76
                        _ = v77
                    }
                    _ = v75
                }
                v2 = v2
                if true {
                    v4 = (v4 | fi73)
                    v69 = v69
                    v2 -= (v2 - v2)
                } else {
                    var v78 uint16 = 2
                    if (false && true) {
                        if false {
                            if false {
                                v1 = 4294967295
                                var v79 uint8 = 42
                                v78 -= (v78 + (v78 | v78))
                                _ = v79
                            }
                            fmt.Println(v2)
                            if ((2 < 2147483648) && false) {
                                v1 += fi73
                                v69 -= (^(v69))
                                var v80 int32 = v69
                                _ = v80
                            }
                        }
                        var v81 int32 = -1
                        if true {
                            switch (^(v1)) {
                            case 9223372036854775807:
                                if true {
                                    v1 = 56
                                    var v82 float32 = 0.1
                                    if (4294967295 <= 992) {
                                        v81 = v81
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v83 string = "hello"
                                            _ = v83
                                        case "":
                                            var v84 int64 = (-(v2))
                                            _ = v84
                                        case "a":
                                            fmt.Println(fi74, v82, fi74)
                                        }
                                    }
                                    _ = v82
                                }
                            case 4294967295:
                                switch int(v2) {
                                case -591:
                                    if true {
                                        v85 := GenS8{F0: 1, F1: "hello", F2: 2147483647}
                                        _ = v85
                                    }
                                case -373:
                                    var v86 int16 = -32768
                                    _ = v86
                                }
                            }
                            switch (v4 & 2147483647) {
                            case 32768:
                                switch (fi74 << ((65535) & 63)) {
                                case -32:
                                    var v87 uint8 = 127
                                    v78 -= (v78 ^ v78)
                                    v88 := GenS6{F0: "\n\t"}
                                    _ = v87
                                    _ = v88
                                default:
                                    switch v4 {
                                    case 32767:
                                        v4 = v1
                                    case -525:
                                        if (true && (true || false)) {
                                            var v89 int16 = -1
                                            _ = v89
                                        }
                                    default:
                                        var v90 int32 = (v81 * v81)
                                        if false {
                                            if len(v70) > 0 {
                                                v70[0] = (fi74 - (fi74 >> ((42) & 63)))
                                            }
                                        }
                                        if true {
                                            v4 -= -9223372036854775808
                                            fmt.Println(v78, v2)
                                            var v91 uint = 584
                                            _ = v91
                                        }
                                        _ = v90
                                    }
                                }
                            default:
                                v2 -= (^(v2))
                            }
                        } else {
                            switch "a" {
                            case "foo bar":
                                switch "foo bar" {
                                case "hello":
                                    _ = len(v70)
                                case "foo bar":
                                    var v92 int = 1
                                    _ = v92
                                default:
                                    var v93 uint8 = 128
                                    _ = v93
                                }
                            case "\n\t":
                                var v94 int64 = 0
                                _ = v94
                            case "hello":
                                if ((794 < 42) || (1 <= -32768)) {
                                    v95 := GenS9{F0: true}
                                    var v96 uint64 = 857
                                    var v97 bool = (4294967295 <= 2)
                                    _ = v95
                                    _ = v96
                                    _ = v97
                                } else {
                                    var v98 int = 4294967295
                                    _ = v98
                                }
                            default:
                                _ = len(v3)
                            }
                            var v99 string = (("a" + "foo bar") + "hello")
                            _ = v99
                        }
                        _ = v81
                    } else {
                        v100 := GenS10{F0: "a", F1: (!true)}
                        if (v69 <= (v69 / ((v69) | 1))) {
                            if (!false) {
                                v69 += 555
                                v78 = 42
                                switch fi73 {
                                case 42:
                                    _ = len(v3)
                                case 733:
                                    switch -941 {
                                    case 0:
                                        switch (fi74 ^ (v1 * v4)) {
                                        case -1:
                                            if len(v70) >= 2 {
                                                v70 = v70[1:len(v70)]
                                            }
                                        }
                                    case -2147483648:
                                        var v101 int = 2147483647
                                        _ = v101
                                    case 897:
                                        if (668 < 42) {
                                            v102 := GenS8{F0: 100, F1: ("hello" + "hello"), F2: 397}
                                            switch (("" + "\n\t") + "") {
                                            case "foo bar":
                                                if false {
                                                    var v103 uint32 = 482
                                                    if len(v70) >= 2 {
                                                        v70 = v70[1:len(v70)]
                                                    }
                                                    if false {
                                                        v103 -= v103
                                                        v2 = (-(v2))
                                                        var v104 uint32 = (^(v103))
                                                        _ = v104
                                                    } else {
                                                        if (1.5 >= (float64(3.14159) * -2.25)) {
                                                            var v105 int16 = 1
                                                            _ = v105
                                                        } else {
                                                            v1 = fi74
                                                        }
                                                        v1 += v4
                                                    }
                                                    _ = v103
                                                } else {
                                                    v70 = append(v70, (v4 >> ((495) & 63)))
                                                    if true {
                                                        switch fi73 {
                                                        case 1:
                                                            var v106 uint16 = 281
                                                            _ = v106
                                                        default:
                                                            var v107 uint = 255
                                                            _ = v107
                                                        }
                                                        if ((false && ((true || (true && (80 <= 100))) || true)) || (-231 == v69)) {
                                                            v108 := GenS11{F0: 85}
                                                            if ((1.5 / ((1.5 - 1.5) - (-2.25 - 0.1))) >= float32(0.1)) {
                                                                v109 := GenS11{F0: 42}
                                                                v110 := v3[431]
                                                                var v111 bool = false
                                                                _ = v109
                                                                _ = v110
                                                                _ = v111
                                                            }
                                                            _ = len(v3)
                                                            var v112 float64 = (1.5 - -2.25)
                                                            if (42 == 479) {
                                                                delete(v3, 2147483647)
                                                                switch (("" + "hello") + "hello") {
                                                                case "\n\t":
                                                                    var v113 bool = false
                                                                    _ = v113
                                                                }
                                                                fmt.Println(fi73, v78)
                                                            }
                                                            v69 += v69
                                                            switch v1 {
                                                            case -541:
                                                                var v114 uint16 = v78
                                                                _ = v114
                                                            default:
                                                                if true {
                                                                    var v115 int = (fi74 % ((v1) | 1))
                                                                    if (978 < 0) {
                                                                        v1 += fi73
                                                                        switch (("foo bar" + "hello") + "hello") {
                                                                        case "hello":
                                                                            switch fi74 {
                                                                            case -1:
                                                                                var v116 float64 = v112
                                                                                _ = v116
                                                                            case 2147483647:
                                                                                var v117 int = v115
                                                                                _ = v117
                                                                            default:
                                                                                fmt.Println(v4)
                                                                            }
                                                                        default:
                                                                            if (1000 > -156) {
                                                                                v69 = (v69 * (v69 & -2147483648))
                                                                                if false {
                                                                                    if len(v70) > 0 {
                                                                                        v70[0] = 32767
                                                                                    }
                                                                                    if false {
                                                                                        _ = len(v3)
                                                                                    } else {
                                                                                        var v118 uint64 = 1
                                                                                        if len(v70) >= 2 {
                                                                                            v70 = v70[1:len(v70)]
                                                                                        }
                                                                                        _ = v118
                                                                                    }
                                                                                    v70 = append(v70, (fi73 * fi73))
                                                                                    v70 = append(v70, (fi74 ^ (fi73 + fi73)))
                                                                                }
                                                                                v119 := GenS12{F0: (v2 >> ((2) & 63)), F1: 128, F2: (!(2 <= 596))}
                                                                                _ = v119
                                                                            } else {
                                                                                if len(v70) >= 2 {
                                                                                    v70 = v70[1:len(v70)]
                                                                                }
                                                                                fmt.Println(v78, v69, v69)
                                                                                var v120 uint32 = 2
                                                                                _ = v120
                                                                            }
                                                                        }
                                                                        v121 := GenS13{F0: 1, F1: 128}
                                                                        v4 = 2147483647
                                                                        var v122 int64 = (v2 & (v2 >> ((255) & 63)))
                                                                        var v123 int = (v115 | v4)
                                                                        _ = v121
                                                                        _ = v122
                                                                        _ = v123
                                                                    } else {
                                                                        if true {
                                                                            fmt.Println(v69)
                                                                        }
                                                                        v2 -= 9223372036854775807
                                                                        v69 = v69
                                                                    }
                                                                    _ = v115
                                                                } else {
                                                                    switch (fi73 % (((v4 + fi74)) | 1)) {
                                                                    case 0:
                                                                        v78 += 42
                                                                    case 32767:
                                                                        if (!(255 != v78)) {
                                                                            var v124 uint = 2
                                                                            v4 -= (fi74 & (fi74 << (((v124 - (v124 / (((^(v124))) | 1)))) & 63)))
                                                                            var v125 uint64 = 268
                                                                            v3[452] = (-(fi73))
                                                                            v126 := GenS14{F0: int64(fi73), F1: -635}
                                                                            switch "hello" {
                                                                            case "\n\t":
                                                                                var v127 string = "foo bar"
                                                                                _ = v127
                                                                            case "foo bar":
                                                                                var v128 float64 = (((((v112 - 3.14159) / v112) * 1e10) - 3.14159) + v112)
                                                                                _ = v128
                                                                            }
                                                                            v70 = append(v70, (v4 & (fi73 / ((fi73) | 1))))
                                                                            if (fi74 == (v1 % (((-(fi73))) | 1))) {
                                                                                v124 = v124
                                                                                var v129 int16 = int16(v124)
                                                                                v130 := GenS14{F0: (v2 % ((v2) | 1)), F1: v4}
                                                                                _ = v129
                                                                                _ = v130
                                                                            } else {
                                                                                if false {
                                                                                    v125 = (v125 / ((885) | 1))
                                                                                    _ = len(v70)
                                                                                    var v131 bool = true
                                                                                    _ = v131
                                                                                } else {
                                                                                    v132 := GenS15{F0: (v125 >> (((v124 * v124)) & 63))}
                                                                                    v133 := GenS16{F0: -1}
                                                                                    var v134 uint32 = 278
                                                                                    _ = v132
                                                                                    _ = v133
                                                                                    _ = v134
                                                                                }
                                                                            }
                                                                            _ = v124
                                                                            _ = v125
                                                                            _ = v126
                                                                        }
                                                                    default:
                                                                        v69 = (v69 | v69)
                                                                    }
                                                                }
                                                                var v135 uint16 = 2
                                                                var v136 uint = 18446744073709551615
                                                                _ = v135
                                                                _ = v136
                                                            }
                                                            fmt.Println(fi74)
                                                            v137 := v3[57]
                                                            switch (("hello" + ((("" + ("a" + "hello")) + "foo bar") + "a")) + "\n\t") {
                                                            case "\n\t":
                                                                if (42 <= 406) {
                                                                    switch (v137 << ((643) & 63)) {
                                                                    case 9223372036854775807:
                                                                        switch (("\n\t" + "foo bar") + (("a" + ((("foo bar" + "a") + "a") + "foo bar")) + "")) {
                                                                        case "foo bar":
                                                                            if true {
                                                                                v137 = fi73
                                                                                if (!true) {
                                                                                    v78 += 42
                                                                                    var v138 float32 = 1.5
                                                                                    _ = v138
                                                                                } else {
                                                                                    var v139 int16 = 459
                                                                                    v140 := v3[1]
                                                                                    v141 := GenS17{F0: 4294967295, F1: 4294967295}
                                                                                    _ = v139
                                                                                    _ = v140
                                                                                    _ = v141
                                                                                }
                                                                            } else {
                                                                                v142 := GenS16{F0: 32767}
                                                                                var v143 bool = true
                                                                                _ = v142
                                                                                _ = v143
                                                                            }
                                                                        case "":
                                                                            fmt.Println(v69, v137, v4)
                                                                        case "hello":
                                                                            var v144 bool = false
                                                                            _ = v144
                                                                        default:
                                                                            fmt.Println(v78, fi73, fi73)
                                                                        }
                                                                    }
                                                                    switch "" {
                                                                    case "hello":
                                                                        _ = len(v3)
                                                                    case "a":
                                                                        switch "" {
                                                                        case "\n\t":
                                                                            v78 += v78
                                                                        case "foo bar":
                                                                            switch "a" {
                                                                            case "hello":
                                                                                var v145 int64 = -1
                                                                                _ = v145
                                                                            case "\n\t":
                                                                                var v146 int16 = 1
                                                                                _ = v146
                                                                            }
                                                                        case "a":
                                                                            switch "\n\t" {
                                                                            case "foo bar":
                                                                                v69 += (v69 - -2147483648)
                                                                            case "":
                                                                                v147 := GenS18{F0: 104, F1: false}
                                                                                _ = v147
                                                                            case "\n\t":
                                                                                var v148 string = ("hello" + "")
                                                                                _ = v148
                                                                            default:
                                                                                v2 += -1
                                                                            }
                                                                        }
                                                                    case "\n\t":
                                                                        _ = len(v70)
                                                                    }
                                                                }
                                                            case "hello":
                                                                var v149 uint = 2
                                                                _ = v149
                                                            case "a":
                                                                v150 := GenS9{F0: false}
                                                                _ = v150
                                                            default:
                                                                fmt.Println(fi73, v78)
                                                            }
                                                            _ = v108
                                                            _ = v112
                                                            _ = v137
                                                        }
                                                    } else {
                                                        v151 := GenS6{F0: (("hello" + "\n\t") + ((("hello" + "hello") + "hello") + "hello"))}
                                                        if true {
                                                            if (0 != 255) {
                                                                v1 += v4
                                                                v152 := GenS8{F0: -128, F1: (("hello" + "") + "foo bar"), F2: 324}
                                                                var v153 float32 = (((-2.25 * 0.0) / 1.0) + -2.25)
                                                                _ = v152
                                                                _ = v153
                                                            }
                                                            if (!false) {
                                                                if false {
                                                                    _ = len(v70)
                                                                    var v154 uint32 = 42
                                                                    _ = v154
                                                                } else {
                                                                    switch ("hello" + ((("\n\t" + ("\n\t" + "hello")) + ("foo bar" + (("a" + "\n\t") + "a"))) + "foo bar")) {
                                                                    case "":
                                                                        var v155 uint64 = 2
                                                                        _ = v155
                                                                    case "hello":
                                                                        fmt.Println(v2, v69)
                                                                        var v156 uint16 = 65535
                                                                        _ = v156
                                                                    default:
                                                                        if true {
                                                                            if (true || false) {
                                                                                v157 := GenS9{F0: ("" >= (("" + "foo bar") + (("foo bar" + "foo bar") + "")))}
                                                                                switch "hello" {
                                                                                case "a":
                                                                                    var v158 float32 = 3.14159
                                                                                    _ = v158
                                                                                case "hello":
                                                                                    fmt.Println(fi73, fi73)
                                                                                }
                                                                                _ = v157
                                                                            } else {
                                                                                var v159 bool = true
                                                                                var v160 int16 = -32768
                                                                                switch v1 {
                                                                                case -1:
                                                                                    if v159 {
                                                                                        var v161 int64 = (v2 | (v2 | (^(v2))))
                                                                                        _ = v161
                                                                                    }
                                                                                default:
                                                                                    if (int16(v78) >= (v160 ^ (v160 % (((-(v160))) | 1)))) {
                                                                                        v162 := GenS4{F0: 753, F1: v159, F2: (v160 + v160)}
                                                                                        fmt.Println(v69)
                                                                                        _ = v162
                                                                                    }
                                                                                }
                                                                                _ = v159
                                                                                _ = v160
                                                                            }
                                                                            var v163 float64 = -2.25
                                                                            var v164 uint32 = 482
                                                                            _ = v163
                                                                            _ = v164
                                                                        }
                                                                    }
                                                                    switch -655 {
                                                                    case 32767:
                                                                        var v165 uint64 = 877
                                                                        var v166 int = -458
                                                                        _ = v165
                                                                        _ = v166
                                                                    }
                                                                    v78 += (v78 << ((uint(v78)) & 63))
                                                                }
                                                            } else {
                                                                v167 := GenS16{F0: -32768}
                                                                fmt.Println(v69, v69)
                                                                _ = v167
                                                            }
                                                            v2 -= (v2 & v2)
                                                            var v168 int32 = (v69 | v69)
                                                            if ((0 != 2) && false) {
                                                                v78 = v78
                                                                var v169 float32 = ((3.14159 - 1.5) + (float32((0.0 * 1.5)) * ((1.5 - -2.25) + ((-2.25 / (1.5 - 0.1)) / 1e10))))
                                                                fmt.Println(v78, v1)
                                                                _ = v169
                                                            } else {
                                                                fmt.Println(fi74, v168, fi74)
                                                                var v170 int16 = 279
                                                                _ = v170
                                                            }
                                                            _ = v168
                                                        } else {
                                                            var v171 uint32 = 2147483647
                                                            _ = v171
                                                        }
                                                        var v172 int64 = v2
                                                        _ = v151
                                                        _ = v172
                                                    }
                                                }
                                            case "":
                                                v2 -= (v2 % ((42) | 1))
                                            case "\n\t":
                                                v70 = append(v70, 128)
                                            }
                                            var v173 uint64 = 1
                                            var v174 int32 = -1
                                            v2 -= (v2 ^ 884)
                                            _ = v102
                                            _ = v173
                                            _ = v174
                                        }
                                    default:
                                        switch "a" {
                                        case "\n\t":
                                            if false {
                                                switch "" {
                                                case "":
                                                    var v175 string = "foo bar"
                                                    _ = v175
                                                case "foo bar":
                                                    if len(v70) > 0 {
                                                        v70[0] = fi73
                                                    }
                                                default:
                                                    v176 := GenS19{F0: -120}
                                                    _ = v176
                                                }
                                                var v177 int16 = -118
                                                _ = v177
                                            } else {
                                                fmt.Println(v4, v78)
                                                var v178 uint8 = 1
                                                v70 = append(v70, (v1 - v4))
                                                _ = v178
                                            }
                                        case "hello":
                                            switch fi73 {
                                            case 426:
                                                v179 := GenS20{F0: 32767}
                                                _ = v179
                                            default:
                                                v180 := GenS6{F0: ("foo bar" + "a")}
                                                _ = v180
                                            }
                                        }
                                    }
                                default:
                                    var v181 int16 = -1
                                    _ = v181
                                }
                            }
                            var v182 int = int(v78)
                            _ = v182
                        }
                        v2 -= 0
                        _ = v100
                    }
                    fmt.Println(v78, fi74, v69)
                    v183 := GenS12{F0: v2, F1: 127, F2: (false || true)}
                    v184 := GenS21{F0: 246, F1: 833, F2: 2}
                    _ = v183
                    _ = v184
                    _ = v78
                }
                _ = fi74
            }
            fmt.Println(v4, v4)
            _ = fi73
        }
        if ((1.5 * 0.1) >= -2.25) {
            if (true && (!((92 < 42) && true))) {
                if (int32(v2) >= v69) {
                    for fi185 := 0; fi185 < 4; fi185++ {
                        var v186 string = "hello"
                        _ = fi185
                        _ = v186
                    }
                    switch (v4 % (((v4 * v4)) | 1)) {
                    case 27:
                        v187 := GenS22{F0: 255, F1: 91, F2: 1}
                        _ = v187
                    default:
                        var v188 float32 = 1e10
                        _ = v188
                    }
                }
                v189 := GenS13{F0: 0, F1: 116}
                _ = v189
            } else {
                v2 = (v2 / ((521) | 1))
                var v190 string = "foo bar"
                var v191 string = v190
                _ = v190
                _ = v191
            }
        } else {
            if (0 < 65535) {
                var v192 int64 = v2
                v193 := GenS20{F0: 2147483647}
                _ = len(v70)
                _ = v192
                _ = v193
            } else {
                v194 := GenS23{F0: "\n\t", F1: 32767}
                v69 = v69
                fmt.Println(v1, v69, v2)
                _ = v194
            }
            _ = len(v70)
        }
    }
    for fi195 := 0; fi195 < 2; fi195++ {
        v196 := GenS3{F0: 0, F1: 127}
        _ = fi195
        _ = v196
    }
    switch (v4 / ((v1) | 1)) {
    case 128:
        var v197 uint16 = 1000
        _ = v197
    case 212:
        var v198 int8 = -128
        _ = v198
    case 32767:
        for fi199 := 0; fi199 < 2; fi199++ {
            if (v69 != (v69 | (v69 >> ((815) & 63)))) {
                var v200 float64 = 0.1
                v201 := v3[2147483647]
                var v202 int32 = v69
                _ = v200
                _ = v201
                _ = v202
            }
            if true {
                if true {
                    for fi203 := 0; fi203 < 1; fi203++ {
                        var v204 uint = 2
                        _ = fi203
                        _ = v204
                    }
                }
            }
            _ = fi199
        }
    default:
        if (2147483647 < 141) {
            for fi205 := 0; fi205 < 13; fi205++ {
                v69 = (^(v69))
                fmt.Println(v69)
                _ = len(v70)
                _ = fi205
            }
            v206 := GenS24{F0: 18446744073709551615}
            var v207 uint8 = 0
            _ = v206
            _ = v207
        }
    }
    for fi208 := 0; fi208 < 7; fi208++ {
        var v209 string = ("hello" + "\n\t")
        v210 := v3[2147483648]
        _ = fi208
        _ = v209
        _ = v210
    }
    fmt.Println(v1, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v69:", v69)
    fmt.Println("slicelen v70:", len(v70), v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v70
}
