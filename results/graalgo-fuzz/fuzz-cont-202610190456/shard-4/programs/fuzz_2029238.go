package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int32; F1 uint8; F2 int }
type GenS3 struct { F0 int64; F1 int16; F2 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 uint16; F1 int16; F2 uint }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int32; F1 int8; F2 uint16 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int; F1 string }
type GenS12 struct { F0 int; F1 int64 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 int8; F1 int64; F2 uint }
type GenS16 struct { F0 string }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint8 = 59
    v3 := make(map[string]uint32)
    var v4 int64 = (v1 * (v1 % (((v1 & v1)) | 1)))
    v5 := []uint32{284, 2147483648, 195, 210}
    if (18446744073709551615 <= 0) {
        fmt.Println(v4, v4, v4)
        var v6 uint16 = uint16(v2)
        var v7 uint64 = 2
        _ = v6
        _ = v7
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 int8 = 1
        _ = fi8
        _ = v9
    }
    var v10 int32 = -1
    var v11 float64 = (((-2.25 + 0.1) / -2.25) - (0.0 - ((0.0 - 1.5) + 0.0)))
    if false {
        v12 := GenS1{F0: 344, F1: (3.14159 == (((1e10 + float32(v11)) - (0.0 - -2.25)) + 1e10))}
        _ = v12
    }
    var v13 uint32 = 996
    v4 = v4
    var v14 int8 = 42
    if false {
        v1 = -1
    } else {
        for fi15 := 0; fi15 < 16; fi15++ {
            v13 += 2
            v1 = 1
            _ = fi15
        }
        if true {
            switch -2147483648 {
            case 525:
                v2 = v2
            case 32767:
                var v16 int = -9223372036854775808
                _ = v16
            case 128:
                switch -938 {
                case 4294967295:
                    v4 = 42
                case -572:
                    if false {
                        delete(v3, "foo bar")
                    } else {
                        if ((^(v1)) <= v4) {
                            var v17 int32 = 1
                            _ = len(v3)
                            fmt.Println(v11)
                            _ = v17
                        }
                    }
                default:
                    var v18 int8 = 61
                    _ = v18
                }
            default:
                if ((v13 <= uint32(v2)) && true) {
                    var v19 uint8 = (^(v2))
                    _ = v19
                }
            }
        } else {
            var v20 string = "hello"
            _ = v20
        }
        switch "" {
        case "\n\t":
            if false {
                var v21 int64 = v1
                for fi22 := 0; fi22 < 10; fi22++ {
                    v1 = v1
                    for fi23 := 0; fi23 < 13; fi23++ {
                        var v24 uint64 = uint64(v2)
                        _ = fi23
                        _ = v24
                    }
                    v11 = 1.5
                    _ = fi22
                }
                _ = v21
            } else {
                v25 := GenS1{F0: 0, F1: false}
                var v26 bool = (3.14159 != (1.5 / ((1e10 - 1e10) + 3.14159)))
                v27 := GenS2{F0: (v10 ^ v10), F1: (^(v2)), F2: 428}
                _ = v25
                _ = v26
                _ = v27
            }
        case "hello":
            v2 = (v2 & v2)
        }
    }
    var v28 int8 = int8(v4)
    if true {
        switch "\n\t" {
        case "":
            switch "a" {
            case "foo bar":
                switch "foo bar" {
                case "":
                    v29 := v3[("\n\t" + "foo bar")]
                    _ = v29
                default:
                    fmt.Println(v14, v14)
                }
            case "a":
                var v30 float32 = 3.14159
                _ = v30
            }
        case "a":
            fmt.Println(v2, v28)
        }
    } else {
        var v31 int8 = -128
        for fi32 := 0; fi32 < 15; fi32++ {
            var v33 uint32 = (v13 * v13)
            _ = fi32
            _ = v33
        }
        var v34 string = ("a" + ("hello" + ("hello" + (("hello" + "foo bar") + ("hello" + "a")))))
        _ = v31
        _ = v34
    }
    var v35 int16 = -252
    var v36 uint = 42
    switch "a" {
    case "a":
        if (-2.25 >= v11) {
            v28 += (v14 - v28)
        }
    case "hello":
        if true {
            switch 0 {
            case 2147483647:
                if false {
                    if false {
                        for fi37 := 0; fi37 < 1; fi37++ {
                            var v38 int = -613
                            _ = fi37
                            _ = v38
                        }
                        for fi39 := 0; fi39 < 9; fi39++ {
                            _ = len(v5)
                            v40 := GenS2{F0: 2147483647, F1: v2, F2: (fi39 ^ (fi39 | -1))}
                            _ = fi39
                            _ = v40
                        }
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = len(v5)
                } else {
                    v13 += (v13 % ((42) | 1))
                    v11 += 0.1
                    v11 -= ((v11 * v11) / 1.0)
                }
            case 983:
                for fi41 := 0; fi41 < 7; fi41++ {
                    v42 := GenS3{F0: (v1 ^ v4), F1: 42, F2: 939}
                    _ = fi41
                    _ = v42
                }
            case -1:
                fmt.Println(v2, v1, v35)
            default:
                v43 := v3[""]
                _ = v43
            }
            switch "\n\t" {
            case "\n\t":
                for fi44 := 0; fi44 < 7; fi44++ {
                    v1 += v4
                    var v45 string = "foo bar"
                    for fi46 := 0; fi46 < 12; fi46++ {
                        fmt.Println(v28, v13, v35)
                        fmt.Println(v45, v45, fi44)
                        v47 := GenS1{F0: (v35 * -1), F1: (42 >= 32767)}
                        _ = fi46
                        _ = v47
                    }
                    _ = fi44
                    _ = v45
                }
            case "a":
                for fi48 := 0; fi48 < 20; fi48++ {
                    for fi49 := 0; fi49 < 15; fi49++ {
                        v50 := GenS2{F0: (v10 / (((v10 + int32(v36))) | 1)), F1: (v2 * 127), F2: fi49}
                        if (2147483647 <= 32768) {
                            if false {
                                v10 -= v10
                            }
                            var v51 int16 = (v35 ^ (v35 + v35))
                            _ = v51
                        } else {
                            v5 = append(v5, v13)
                            v14 -= v28
                        }
                        var v52 uint = uint(v13)
                        _ = fi49
                        _ = v50
                        _ = v52
                    }
                    if true {
                        v4 = (-(v1))
                        var v53 bool = false
                        switch (fi48 ^ (fi48 / (((fi48 & int(v10))) | 1))) {
                        case 128:
                            var v54 int64 = 665
                            _ = v54
                        case -549:
                            if (false && (593 < (v10 ^ -2147483648))) {
                                var v55 uint = 2
                                _ = v55
                            } else {
                                switch "" {
                                case "hello":
                                    switch 9223372036854775807 {
                                    case -1:
                                        v3["\n\t"] = 2147483647
                                    }
                                default:
                                    for fi56 := 0; fi56 < 4; fi56++ {
                                        switch 128 {
                                        case 1:
                                            switch ("" + ("" + ("hello" + ""))) {
                                            case "foo bar":
                                                v14 = (v28 ^ (v28 >> ((v36) & 63)))
                                            case "":
                                                var v57 uint32 = (v13 ^ 4294967295)
                                                _ = v57
                                            case "\n\t":
                                                v2 = v2
                                                var v58 uint64 = 40
                                                v36 -= uint(v13)
                                                v5 = append(v5, (v13 << ((42) & 63)))
                                                v14 -= v28
                                                _ = v58
                                            }
                                            switch (fi56 | fi48) {
                                            case -389:
                                                if (true && (("" + (("a" + "a") + ("\n\t" + ""))) <= "foo bar")) {
                                                    var v59 int = fi56
                                                    switch -716 {
                                                    case 32768:
                                                        v11 += (0.0 / 1.0)
                                                    }
                                                    _ = v59
                                                }
                                            case -1:
                                                if (v13 != (v13 % ((2147483647) | 1))) {
                                                    v4 += v1
                                                    if (58 == (v13 % ((v13) | 1))) {
                                                        var v60 int32 = (v10 << (((v36 - v36)) & 63))
                                                        _ = v60
                                                    }
                                                }
                                            case -2147483648:
                                                v14 = (v28 ^ 89)
                                            default:
                                                var v61 bool = ((fi56 % ((fi56) | 1)) == fi48)
                                                _ = v61
                                            }
                                            var v62 int64 = (v1 + 9223372036854775807)
                                            _ = v62
                                        default:
                                            v63 := GenS4{F0: v2}
                                            _ = v63
                                        }
                                        v53 = false
                                        _ = fi56
                                    }
                                }
                                var v64 uint64 = uint64(v2)
                                v65 := GenS2{F0: v10, F1: 2, F2: fi48}
                                _ = v64
                                _ = v65
                            }
                        default:
                            var v66 int = (fi48 << ((v36) & 63))
                            _ = v66
                        }
                        _ = v53
                    } else {
                        fmt.Println(v13, v13)
                        switch fi48 {
                        case 4294967295:
                            fmt.Println(v28)
                        case -479:
                            v67 := GenS1{F0: int16(v4), F1: true}
                            _ = v67
                        case -496:
                            if false {
                                switch fi48 {
                                case 42:
                                    var v68 int = (fi48 | fi48)
                                    _ = v68
                                case -10:
                                    v10 += -1
                                case 32768:
                                    switch 550 {
                                    case -2147483648:
                                        if true {
                                            fmt.Println(v36, v10)
                                            if false {
                                                var v69 uint8 = v2
                                                _ = v69
                                            }
                                        } else {
                                            if true {
                                                switch "" {
                                                case "":
                                                    for fi70 := 0; fi70 < 11; fi70++ {
                                                        if (v36 != v36) {
                                                            if (true && true) {
                                                                v2 += (^(v2))
                                                                fmt.Println(v11, fi48)
                                                                var v71 bool = true
                                                                _ = v71
                                                            }
                                                        }
                                                        _ = fi70
                                                    }
                                                case "foo bar":
                                                    v13 -= v13
                                                default:
                                                    v72 := v3["\n\t"]
                                                    _ = v72
                                                }
                                                v73 := GenS2{F0: (v10 ^ v10), F1: (v2 % ((130) | 1)), F2: fi48}
                                                v28 += (v28 * v14)
                                                _ = v73
                                            }
                                        }
                                    }
                                }
                                for fi74 := 0; fi74 < 8; fi74++ {
                                    var v75 int16 = v35
                                    var v76 float64 = 0.1
                                    v77 := GenS1{F0: -32768, F1: (!true)}
                                    if true {
                                        v78 := GenS5{F0: 509, F1: 84}
                                        _ = v78
                                    } else {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        switch fi48 {
                                        case 1:
                                            var v79 string = "\n\t"
                                            _ = v79
                                        default:
                                            var v80 int64 = (^(v4))
                                            _ = v80
                                        }
                                        v35 = v35
                                    }
                                    v5 = append(v5, v13)
                                    _ = fi74
                                    _ = v75
                                    _ = v76
                                    _ = v77
                                }
                            }
                        }
                        switch (fi48 * (^(fi48))) {
                        case -343:
                            v10 = 2147483647
                        case -2147483648:
                            v36 = uint(v2)
                        default:
                            var v81 float32 = -2.25
                            _ = v81
                        }
                    }
                    for fi82 := 0; fi82 < 9; fi82++ {
                        var v83 uint8 = (v2 - 1)
                        if ((!true) && true) {
                            var v84 uint16 = 255
                            fmt.Println(v28)
                            var v85 float32 = float32(((v11 - (float64(-2.25) / (1e10 - (v11 * (1e10 / 3.14159))))) / v11))
                            switch ("foo bar" + "hello") {
                            case "foo bar":
                                var v86 uint32 = (v13 & v13)
                                _ = v86
                            default:
                                v87 := GenS4{F0: 1}
                                _ = v87
                            }
                            fmt.Println(v36, v4, v1)
                            switch "hello" {
                            case "\n\t":
                                if false {
                                    switch "a" {
                                    case "a":
                                        v5 = append(v5, 852)
                                    case "":
                                        var v88 uint = (v36 | (v36 / (((v36 * (v36 % ((v36) | 1)))) | 1)))
                                        _ = v88
                                    case "foo bar":
                                        var v89 int = 42
                                        _ = v89
                                    default:
                                        var v90 int64 = v1
                                        _ = v90
                                    }
                                }
                            case "":
                                fmt.Println(v13, fi48)
                            case "hello":
                                fmt.Println(v13, v28)
                            }
                            v91 := GenS1{F0: -992, F1: (!true)}
                            var v92 int32 = v10
                            v93 := GenS6{F0: (v84 ^ (v84 * 0)), F1: v35, F2: v36}
                            v35 += int16(v13)
                            var v94 uint32 = v13
                            _ = v84
                            _ = v85
                            _ = v91
                            _ = v92
                            _ = v93
                            _ = v94
                        }
                        _ = fi82
                        _ = v83
                    }
                    _ = fi48
                }
            case "hello":
                fmt.Println(v11, v14)
            }
        } else {
            for fi95 := 0; fi95 < 1; fi95++ {
                switch "\n\t" {
                case "foo bar":
                    for fi96 := 0; fi96 < 17; fi96++ {
                        fmt.Println(v4)
                        _ = fi96
                    }
                case "a":
                    for fi97 := 0; fi97 < 18; fi97++ {
                        if len(v5) > 0 {
                            v5[0] = (v13 & 0)
                        }
                        v10 += 0
                        if false {
                            switch "\n\t" {
                            case "\n\t":
                                switch fi97 {
                                case 128:
                                    if (100 >= 47) {
                                        var v98 int = 721
                                        _ = v98
                                    }
                                default:
                                    v35 = (-(v35))
                                }
                            case "foo bar":
                                switch (fi95 ^ (fi97 - int(v36))) {
                                case -136:
                                    if false {
                                        if false {
                                            if true {
                                                fmt.Println(v2, v14)
                                            }
                                            if ((v35 % (((^(v35))) | 1)) != v35) {
                                                if true {
                                                    switch (-(fi95)) {
                                                    case 975:
                                                        var v99 string = ((("foo bar" + ("" + "hello")) + "\n\t") + "")
                                                        _ = v99
                                                    default:
                                                        switch fi97 {
                                                        case 47:
                                                            var v100 uint64 = 0
                                                            _ = v100
                                                        default:
                                                            v101 := GenS7{F0: -9223372036854775808}
                                                            _ = v101
                                                        }
                                                    }
                                                    switch fi97 {
                                                    case -1:
                                                        var v102 uint32 = v13
                                                        _ = v102
                                                    default:
                                                        v4 -= v4
                                                    }
                                                    v103 := GenS4{F0: (v2 >> ((768) & 63))}
                                                    _ = v103
                                                } else {
                                                    fmt.Println(v14, v11, v13)
                                                    v11 = v11
                                                }
                                                if (v11 < 1.5) {
                                                    fmt.Println(v10)
                                                } else {
                                                    if (!false) {
                                                        var v104 int8 = (v14 + (v28 / ((v14) | 1)))
                                                        if true {
                                                            var v105 uint8 = v2
                                                            _ = v105
                                                        } else {
                                                            switch (("" + "a") + ("" + "foo bar")) {
                                                            case "\n\t":
                                                                v1 += -1
                                                            }
                                                            _ = len(v5)
                                                        }
                                                        var v106 uint = (v36 / (((v36 << (((v36 >> ((v36) & 63))) & 63))) | 1))
                                                        _ = v104
                                                        _ = v106
                                                    }
                                                    switch "" {
                                                    case "\n\t":
                                                        v107 := GenS8{F0: (v10 >> ((228) & 63)), F1: (v14 - (v28 + 42)), F2: 1}
                                                        _ = v107
                                                    case "":
                                                        if (!false) {
                                                            fmt.Println(v11, v28)
                                                        } else {
                                                            v4 = (v4 % (((v4 << ((42) & 63))) | 1))
                                                            _ = len(v3)
                                                            if true {
                                                                v108 := GenS1{F0: int16(v10), F1: true}
                                                                _ = v108
                                                            }
                                                        }
                                                    case "a":
                                                        v109 := GenS7{F0: (v4 * (v1 >> ((2) & 63)))}
                                                        _ = v109
                                                    }
                                                }
                                                var v110 uint16 = 0
                                                switch (fi95 | (fi95 - fi97)) {
                                                case 1:
                                                    v111 := v3[("\n\t" + "")]
                                                    v112 := GenS9{F0: v35}
                                                    v35 = v35
                                                    switch -1 {
                                                    case 42:
                                                        var v113 int64 = 9223372036854775807
                                                        _ = v113
                                                    case 2147483647:
                                                        if (true || true) {
                                                            switch (^(fi97)) {
                                                            case -207:
                                                                var v114 int = 42
                                                                _ = v114
                                                            case 346:
                                                                switch (^(fi97)) {
                                                                case 128:
                                                                    v1 = v4
                                                                case 2147483647:
                                                                    v2 = (v2 + (v2 % ((1) | 1)))
                                                                case 42:
                                                                    var v115 float64 = float64(0.0)
                                                                    _ = v115
                                                                }
                                                            }
                                                            var v116 int8 = (v14 - v28)
                                                            _ = v116
                                                        } else {
                                                            v111 = (v13 ^ (v13 / (((v111 & (v13 ^ v111))) | 1)))
                                                            fmt.Println(v110, v28, v10)
                                                            var v117 int64 = (v4 + -534)
                                                            _ = v117
                                                        }
                                                    }
                                                    v118 := GenS6{F0: 65535, F1: 32767, F2: 18446744073709551615}
                                                    v119 := GenS8{F0: v10, F1: int8(v10), F2: (v110 & (v110 % (((^(v110))) | 1)))}
                                                    v120 := GenS9{F0: v35}
                                                    _ = v111
                                                    _ = v112
                                                    _ = v118
                                                    _ = v119
                                                    _ = v120
                                                case -9223372036854775808:
                                                    if false {
                                                        var v121 uint = 18446744073709551615
                                                        _ = v121
                                                    } else {
                                                        v5 = append(v5, 0)
                                                        switch ("\n\t" + ("foo bar" + "a")) {
                                                        case "foo bar":
                                                            if ((v1 - -279) > (v4 - -877)) {
                                                                v14 += v28
                                                                v2 += v2
                                                            } else {
                                                                var v122 uint32 = (^(v13))
                                                                var v123 float32 = 3.14159
                                                                _ = v122
                                                                _ = v123
                                                            }
                                                        case "":
                                                            var v124 int = fi95
                                                            _ = v124
                                                        case "\n\t":
                                                            v125 := GenS2{F0: v10, F1: (v2 * v2), F2: (-(fi97))}
                                                            _ = v125
                                                        default:
                                                            v126 := GenS10{F0: uint8(v13)}
                                                            _ = v126
                                                        }
                                                    }
                                                default:
                                                    switch "" {
                                                    case "a":
                                                        if len(v5) > 0 {
                                                            v5[0] = v13
                                                        }
                                                    }
                                                }
                                                _ = v110
                                            }
                                            v127 := GenS9{F0: (v35 ^ v35)}
                                            _ = v127
                                        }
                                    } else {
                                        switch "" {
                                        case "\n\t":
                                            var v128 uint64 = 2
                                            var v129 int8 = (v28 ^ (v14 + v28))
                                            _ = v128
                                            _ = v129
                                        case "hello":
                                            var v130 int32 = (v10 * v10)
                                            _ = v130
                                        case "":
                                            switch "hello" {
                                            case "":
                                                if false {
                                                    v131 := GenS4{F0: v2}
                                                    _ = v131
                                                } else {
                                                    switch "a" {
                                                    case "foo bar":
                                                        fmt.Println(v36, v1)
                                                    case "hello":
                                                        switch 1 {
                                                        case 32768:
                                                            if (int8(v2) != (v28 % ((v14) | 1))) {
                                                                var v132 bool = false
                                                                var v133 int = int(v10)
                                                                _ = v132
                                                                _ = v133
                                                            } else {
                                                                var v134 uint16 = uint16(v2)
                                                                v134 = (^(v134))
                                                                _ = v134
                                                            }
                                                        case 32767:
                                                            _ = len(v3)
                                                        case 0:
                                                            v14 = (v28 - (v14 * (^(v14))))
                                                        }
                                                    }
                                                    v135 := v3["foo bar"]
                                                    var v136 float32 = (0.1 + 1.5)
                                                    _ = v135
                                                    _ = v136
                                                }
                                            case "hello":
                                                v10 += v10
                                            default:
                                                v14 = v14
                                            }
                                        }
                                        v4 -= v1
                                        v137 := v3[""]
                                        _ = v137
                                    }
                                case 969:
                                    if len(v5) > 0 {
                                        v5[0] = 42
                                    }
                                default:
                                    var v138 float32 = 1e10
                                    _ = v138
                                }
                            case "":
                                fmt.Println(fi95, v14)
                            default:
                                if (42 > 255) {
                                    var v139 uint8 = (v2 & v2)
                                    if (!true) {
                                        var v140 uint16 = 32767
                                        switch int(v139) {
                                        case -561:
                                            if true {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                var v141 int64 = (v1 / ((-34) | 1))
                                                v10 = v10
                                                v142 := GenS11{F0: 421, F1: "a"}
                                                if true {
                                                    var v143 int16 = v35
                                                    if false {
                                                        v144 := GenS4{F0: v2}
                                                        var v145 int64 = v1
                                                        switch -9223372036854775808 {
                                                        case 286:
                                                            var v146 bool = true
                                                            _ = v146
                                                        case -564:
                                                            v13 = (^(v13))
                                                        case 32767:
                                                            v10 -= v10
                                                        }
                                                        v11 -= (v11 - -2.25)
                                                        _ = v144
                                                        _ = v145
                                                    }
                                                    _ = v143
                                                } else {
                                                    delete(v3, "\n\t")
                                                    v36 -= 4294967295
                                                    if (!((-(v10)) == int32(fi95))) {
                                                        var v147 float32 = 3.14159
                                                        v35 = (v35 & (v35 % ((v35) | 1)))
                                                        v1 -= (v1 * int64(v14))
                                                        _ = v147
                                                    } else {
                                                        if ((v10 & 686) < (v10 >> (((^(v36))) & 63))) {
                                                            if true {
                                                                v36 += 4294967295
                                                                v10 = int32(v28)
                                                                switch ("hello" + "") {
                                                                case "foo bar":
                                                                    v36 -= 18446744073709551615
                                                                case "":
                                                                    v148 := GenS4{F0: 16}
                                                                    var v149 bool = (int8(v36) != (v28 & (v28 & v28)))
                                                                    _ = v148
                                                                    _ = v149
                                                                default:
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                }
                                                            } else {
                                                                switch "hello" {
                                                                case "a":
                                                                    v150 := GenS12{F0: -2147483648, F1: v4}
                                                                    _ = v150
                                                                case "hello":
                                                                    v151 := GenS9{F0: 1}
                                                                    _ = v151
                                                                case "\n\t":
                                                                    var v152 int32 = v10
                                                                    _ = v152
                                                                }
                                                                var v153 int = (fi97 & 32768)
                                                                v139 += uint8(v36)
                                                                _ = v153
                                                            }
                                                            if false {
                                                                if (v2 > v2) {
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                    var v154 uint16 = (^(v140))
                                                                    _ = v154
                                                                } else {
                                                                    switch fi95 {
                                                                    case 127:
                                                                        fmt.Println(fi97)
                                                                    default:
                                                                        fmt.Println(v139, v28, v139)
                                                                        if (!true) {
                                                                            v5 = append(v5, v13)
                                                                            if (!true) {
                                                                                if false {
                                                                                    if (((((!true) && false) || false) || false) && true) {
                                                                                        var v155 int16 = v35
                                                                                        v5 = append(v5, (v13 >> ((v36) & 63)))
                                                                                        var v156 string = "foo bar"
                                                                                        _ = v155
                                                                                        _ = v156
                                                                                    } else {
                                                                                        v35 = v35
                                                                                    }
                                                                                    if false {
                                                                                        var v157 int8 = v28
                                                                                        var v158 uint64 = 1
                                                                                        _ = v157
                                                                                        _ = v158
                                                                                    } else {
                                                                                        var v159 uint16 = (^(v140))
                                                                                        var v160 int64 = (v4 - (v1 - v1))
                                                                                        var v161 int64 = v141
                                                                                        _ = v159
                                                                                        _ = v160
                                                                                        _ = v161
                                                                                    }
                                                                                } else {
                                                                                    v11 = v11
                                                                                    v35 += v35
                                                                                }
                                                                                v162 := v3[("foo bar" + (("\n\t" + ("foo bar" + "foo bar")) + ("a" + ("" + ("hello" + ((("" + "") + "\n\t") + ("hello" + (("a" + "\n\t") + "hello"))))))))]
                                                                                v2 = (^(v139))
                                                                                _ = v162
                                                                            }
                                                                            v5 = append(v5, v13)
                                                                        } else {
                                                                            var v163 float32 = ((1.5 + float32(1.5)) + 0.0)
                                                                            var v164 bool = true
                                                                            var v165 uint32 = (^(v13))
                                                                            _ = v163
                                                                            _ = v164
                                                                            _ = v165
                                                                        }
                                                                        fmt.Println(v141)
                                                                    }
                                                                    switch (fi97 / ((-9223372036854775808) | 1)) {
                                                                    case 4294967295:
                                                                        switch (fi97 * (fi95 >> ((v36) & 63))) {
                                                                        case -996:
                                                                            if false {
                                                                                switch (fi97 % ((fi95) | 1)) {
                                                                                case 0:
                                                                                    v13 -= 4294967295
                                                                                case -353:
                                                                                    var v166 string = "\n\t"
                                                                                    _ = v166
                                                                                }
                                                                                var v167 float64 = 1e10
                                                                                fmt.Println(fi95, v35)
                                                                                _ = v167
                                                                            }
                                                                            var v168 uint8 = (v2 << ((v36) & 63))
                                                                            v169 := v3[((("" + "a") + ("a" + "a")) + "hello")]
                                                                            if false {
                                                                                var v170 uint16 = v140
                                                                                _ = v170
                                                                            }
                                                                            _ = v168
                                                                            _ = v169
                                                                        default:
                                                                            if (2 < 282) {
                                                                                switch (fi97 / (((-(fi97))) | 1)) {
                                                                                case -667:
                                                                                    switch fi95 {
                                                                                    case 42:
                                                                                        if len(v5) >= 2 {
                                                                                            v5 = v5[1:len(v5)]
                                                                                        }
                                                                                    case -646:
                                                                                        fmt.Println(v36, v36, v139)
                                                                                    }
                                                                                case -387:
                                                                                    fmt.Println(fi95, v36, v10)
                                                                                case -891:
                                                                                    var v171 int64 = -1
                                                                                    _ = v171
                                                                                default:
                                                                                    switch ("\n\t" + "a") {
                                                                                    case "a":
                                                                                        v5 = append(v5, 4294967295)
                                                                                    default:
                                                                                        v10 = (-(v10))
                                                                                        if false {
                                                                                            if (566 == 18446744073709551615) {
                                                                                                _ = len(v5)
                                                                                            } else {
                                                                                                v5 = append(v5, v13)
                                                                                            }
                                                                                            var v172 uint32 = v13
                                                                                            fmt.Println(v14)
                                                                                            _ = v172
                                                                                        } else {
                                                                                            fmt.Println(v14, v1, v36)
                                                                                        }
                                                                                        var v173 string = "\n\t"
                                                                                        v173 = (v173 + v173)
                                                                                        v173 = (((("" + ("" + "\n\t")) + v173) + ((v173 + v173) + v173)) + "\n\t")
                                                                                        _ = v173
                                                                                    }
                                                                                }
                                                                                v36 -= v36
                                                                                var v174 uint = 255
                                                                                _ = v174
                                                                            }
                                                                        }
                                                                    default:
                                                                        switch "" {
                                                                        case "\n\t":
                                                                            v175 := GenS3{F0: (-(v1)), F1: -888, F2: (v140 + (v140 / (((v140 + v140)) | 1)))}
                                                                            _ = v175
                                                                        }
                                                                    }
                                                                }
                                                                v11 += v11
                                                                var v176 int16 = int16(v139)
                                                                _ = v176
                                                            } else {
                                                                fmt.Println(v11, v13)
                                                                var v177 uint16 = (v140 - v140)
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                                var v178 uint = 255
                                                                _ = v177
                                                                _ = v178
                                                            }
                                                            var v179 bool = true
                                                            _ = v179
                                                        }
                                                        var v180 string = ""
                                                        var v181 string = ((v180 + (v180 + "a")) + v180)
                                                        _ = v180
                                                        _ = v181
                                                    }
                                                    v182 := GenS13{F0: "hello"}
                                                    _ = v182
                                                }
                                                var v183 int = 4294967295
                                                _ = len(v3)
                                                v183 += (fi97 + fi97)
                                                _ = v141
                                                _ = v142
                                                _ = v183
                                            } else {
                                                var v184 uint32 = v13
                                                v35 += (v35 / (((v35 + (v35 >> (((v36 << (((v36 >> (((^(v36))) & 63))) & 63))) & 63)))) | 1))
                                                var v185 int16 = 1000
                                                v3["foo bar"] = v184
                                                var v186 uint64 = 879
                                                _ = v184
                                                _ = v185
                                                _ = v186
                                            }
                                        }
                                        _ = v140
                                    }
                                    switch "hello" {
                                    case "foo bar":
                                        var v187 uint32 = (^(v13))
                                        _ = v187
                                    case "\n\t":
                                        v11 += v11
                                    case "":
                                        v188 := GenS1{F0: (v35 * (v35 + 1000)), F1: false}
                                        _ = v188
                                    default:
                                        if true {
                                            v189 := GenS14{F0: v36}
                                            _ = v189
                                        }
                                    }
                                    _ = v139
                                } else {
                                    if len(v5) > 0 {
                                        v5[0] = (v13 << (((^(v36))) & 63))
                                    }
                                    var v190 int16 = (v35 >> (((v36 | (v36 & (v36 | 276)))) & 63))
                                    switch ("" + "\n\t") {
                                    case "\n\t":
                                        v11 += ((3.14159 * -2.25) / (float64(float32(1.5)) + float64(0.0)))
                                        var v191 uint32 = (v13 % (((v13 * 34)) | 1))
                                        fmt.Println(fi97, v13, v28)
                                        _ = v191
                                    case "a":
                                        if true {
                                            var v192 float32 = 1.5
                                            v13 = (v13 >> ((v36) & 63))
                                            if (((fi95 << (((v36 | v36)) & 63)) < fi97) || false) {
                                                if (v13 >= v13) {
                                                    var v193 uint8 = (v2 & (v2 | 42))
                                                    var v194 float32 = ((v192 * (v192 * float32((float64(v192) / v11)))) - 0.1)
                                                    v1 += int64(v10)
                                                    if ((("\n\t" + "hello") + "a") <= "hello") {
                                                        v14 = v14
                                                    } else {
                                                        v195 := GenS15{F0: -106, F1: int64(v13), F2: v36}
                                                        var v196 uint64 = 1
                                                        switch (("a" + "\n\t") + "\n\t") {
                                                        case "\n\t":
                                                            switch (("a" + ("" + "hello")) + "a") {
                                                            case "hello":
                                                                if true {
                                                                    v10 = (-(v10))
                                                                    if len(v5) > 0 {
                                                                        v5[0] = v13
                                                                    }
                                                                    var v197 int64 = v4
                                                                    _ = v197
                                                                } else {
                                                                    fmt.Println(v11, v28)
                                                                    fmt.Println(v10)
                                                                }
                                                            case "":
                                                                if false {
                                                                    v196 = 1
                                                                } else {
                                                                    if (v196 != 2) {
                                                                        v196 -= v196
                                                                        v198 := GenS12{F0: (fi97 + (fi95 / ((128) | 1))), F1: v1}
                                                                        if (!false) {
                                                                            delete(v3, ((("foo bar" + "") + "") + ("hello" + "foo bar")))
                                                                        }
                                                                        _ = v198
                                                                    } else {
                                                                        fmt.Println(v10)
                                                                    }
                                                                }
                                                            default:
                                                                v199 := GenS16{F0: "a"}
                                                                _ = v199
                                                            }
                                                        case "hello":
                                                            var v200 uint8 = v2
                                                            v11 = ((v11 * v11) / v11)
                                                            if true {
                                                                fmt.Println(v200)
                                                            }
                                                            _ = v200
                                                        }
                                                        _ = v195
                                                        _ = v196
                                                    }
                                                    v5 = append(v5, v13)
                                                    _ = v193
                                                    _ = v194
                                                }
                                            }
                                            var v201 string = ("a" + ((("a" + "a") + "") + (("hello" + "a") + "a")))
                                            _ = v192
                                            _ = v201
                                        } else {
                                            v11 -= v11
                                        }
                                    }
                                    v1 = v1
                                    _ = v190
                                }
                            }
                        }
                        _ = fi97
                    }
                case "\n\t":
                    v202 := v3[("\n\t" + "a")]
                    _ = v202
                }
                v35 = (v35 | (v35 & (v35 << ((18446744073709551615) & 63))))
                _ = fi95
            }
        }
    case "":
        v1 -= (v1 + 9223372036854775807)
    }
    if true {
        v5 = append(v5, v13)
        if true {
            var v203 uint32 = 2
            v11 = -2.25
            _ = v203
        } else {
            var v204 int32 = 0
            if true {
                fmt.Println(v28, v11, v13)
                v28 = v28
            } else {
                var v205 int16 = (v35 << (((v36 * uint(v13))) & 63))
                if (!true) {
                    for fi206 := 0; fi206 < 14; fi206++ {
                        _ = len(v3)
                        if ((v205 + 1000) != 0) {
                            switch (fi206 ^ (^(fi206))) {
                            case 32768:
                                if false {
                                    var v207 int8 = (v28 & v28)
                                    _ = v207
                                }
                            case -500:
                                var v208 uint8 = v2
                                _ = v208
                            case 127:
                                v205 -= (-(v35))
                            }
                            var v209 int64 = 1
                            _ = v209
                        }
                        _ = fi206
                    }
                    v5 = append(v5, (^(v13)))
                } else {
                    var v210 uint32 = (^(v13))
                    fmt.Println(v11, v210, v4)
                    for fi211 := 0; fi211 < 6; fi211++ {
                        _ = len(v5)
                        _ = fi211
                    }
                    _ = v210
                }
                _ = v205
            }
            _ = v204
        }
        v36 += v36
    }
    v212 := v3[((("\n\t" + "a") + "") + "hello")]
    switch ("hello" + "") {
    case "hello":
        v213 := GenS9{F0: (v35 ^ -32768)}
        _ = v213
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v212:", v212)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v212
    _ = v28
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
