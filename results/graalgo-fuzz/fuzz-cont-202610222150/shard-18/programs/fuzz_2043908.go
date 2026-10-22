package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 int8; F1 int32; F2 uint32 }
type GenS6 struct { F0 int8; F1 bool; F2 uint16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 bool; F1 int8; F2 int8 }
type GenS9 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint32; F1 uint32 }
type GenS12 struct { F0 uint; F1 uint32; F2 uint16 }
type GenS13 struct { F0 int16; F1 uint32 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS16 struct { F0 string; F1 int16; F2 uint32 }
type GenS17 struct { F0 int64 }
type GenS18 struct { F0 bool; F1 uint32; F2 string }
type GenS19 struct { F0 uint32; F1 int32 }
type GenS20 struct { F0 uint8 }
type GenS21 struct { F0 int64; F1 uint8; F2 int }
type GenS22 struct { F0 uint32; F1 int64; F2 int8 }

func main() {
    var v1 uint = 391
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 42
    v4 := GenS1{F0: 255, F1: ("\n\t" + "hello"), F2: 0}
    fmt.Println(v3, v1)
    v5 := GenS1{F0: 0, F1: "foo bar", F2: 32767}
    v3 += v3
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := GenS1{F0: 1, F1: "", F2: 430}
        if true {
            if (!false) {
                switch "\n\t" {
                case "foo bar":
                    v8 := GenS1{F0: 463, F1: (("\n\t" + ("\n\t" + ("foo bar" + "a"))) + "a"), F2: 65535}
                    _ = v8
                }
                if false {
                    v1 = v1
                    var v9 uint32 = 362
                    _ = v9
                }
            } else {
                delete(v2, 65535)
                v10 := GenS2{F0: v1}
                delete(v2, 189)
                _ = v10
            }
            if true {
                switch fi6 {
                case 0:
                    var v11 uint = 617
                    _ = v11
                case 2147483647:
                    if false {
                        if true {
                            for fi12 := 0; fi12 < 11; fi12++ {
                                v1 = uint(v3)
                                var v13 int = fi12
                                if (float32(3.14159) == (-2.25 * 1e10)) {
                                    v1 = (^(v1))
                                    fmt.Println(v1, v1, v1)
                                }
                                v14 := GenS3{F0: 65535}
                                _ = fi12
                                _ = v13
                                _ = v14
                            }
                            _ = len(v2)
                            for fi15 := 0; fi15 < 1; fi15++ {
                                v2[2147483647] = 42
                                fmt.Println(v1, v1)
                                fmt.Println(v1)
                                _ = fi15
                            }
                        }
                    } else {
                        var v16 uint64 = 306
                        for fi17 := 0; fi17 < 2; fi17++ {
                            v18 := v2[4294967295]
                            v3 -= 2
                            v2[0] = v18
                            _ = fi17
                            _ = v18
                        }
                        v19 := []int{}
                        _ = v16
                        _ = v19
                    }
                case -738:
                    for fi20 := 0; fi20 < 17; fi20++ {
                        if ("" > "foo bar") {
                            v21 := GenS1{F0: 255, F1: ("a" + ("hello" + "")), F2: 2}
                            fmt.Println(v1, v1, fi6)
                            v3 += v3
                            _ = v21
                        } else {
                            var v22 uint = 42
                            _ = v22
                        }
                        var v23 int16 = 1000
                        var v24 int16 = v23
                        if (0 < -128) {
                            fmt.Println(fi6, v1)
                            if true {
                                var v25 uint16 = 2
                                var v26 int16 = 42
                                var v27 int8 = 0
                                switch fi6 {
                                case 42:
                                    v24 -= v23
                                    v28 := []int{fi6}
                                    _ = v28
                                default:
                                    if true {
                                        v24 -= int16(v1)
                                        switch "" {
                                        case "\n\t":
                                            if ((((v3 ^ v3) >= (v3 + uint8(v1))) && false) && (!((!false) || true))) {
                                                v26 = (v24 + v23)
                                            } else {
                                                v25 = (^(v25))
                                                if true {
                                                    if true {
                                                        var v29 int64 = -456
                                                        fmt.Println(fi6, v29)
                                                        _ = v29
                                                    }
                                                    v30 := []int{fi20}
                                                    if ((fi20 * fi6) > (fi6 ^ (-(fi20)))) {
                                                        var v31 float32 = 3.14159
                                                        v25 = v25
                                                        v32 := GenS1{F0: (v25 << ((v1) & 63)), F1: ("" + "a"), F2: v25}
                                                        _ = v31
                                                        _ = v32
                                                    } else {
                                                        var v33 int16 = (v24 & 32767)
                                                        var v34 int32 = -156
                                                        _ = v33
                                                        _ = v34
                                                    }
                                                    _ = v30
                                                }
                                            }
                                        case "":
                                            fmt.Println(fi6, v27, v27)
                                        }
                                        v35 := GenS2{F0: (v1 >> ((v1) & 63))}
                                        v24 = 528
                                        if true {
                                            if (65535 > 2) {
                                                v36 := GenS4{F0: 65535, F1: fi20}
                                                _ = len(v2)
                                                switch (fi20 % ((fi6) | 1)) {
                                                case -110:
                                                    switch 42 {
                                                    case 127:
                                                        if true {
                                                            var v37 float64 = (((-2.25 / 1.0) * 3.14159) / 0.1)
                                                            _ = v37
                                                        } else {
                                                            var v38 string = "foo bar"
                                                            _ = v38
                                                        }
                                                    }
                                                case -2147483648:
                                                    if (false && (false || true)) {
                                                        var v39 int = (fi20 % (((fi6 >> ((v1) & 63))) | 1))
                                                        if true {
                                                            var v40 int32 = -1
                                                            _ = v40
                                                        }
                                                        v41 := GenS2{F0: v1}
                                                        _ = v39
                                                        _ = v41
                                                    } else {
                                                        v42 := []int{}
                                                        fmt.Println(v27)
                                                        _ = v42
                                                    }
                                                case 42:
                                                    fmt.Println(v24)
                                                }
                                                _ = v36
                                            } else {
                                                if (true || (!((v27 >> ((uint(v25)) & 63)) < (v27 / ((v27) | 1))))) {
                                                    var v43 float64 = (float64(3.14159) * -2.25)
                                                    _ = v43
                                                }
                                                v44 := GenS5{F0: int8(v1), F1: -1, F2: 0}
                                                _ = v44
                                            }
                                            var v45 float32 = float32(3.14159)
                                            v46 := GenS6{F0: v27, F1: (true && (v25 >= 774)), F2: (v25 << ((v1) & 63))}
                                            _ = v45
                                            _ = v46
                                        } else {
                                            v47 := []int{(-(fi20))}
                                            var v48 int = (fi20 >> ((v1) & 63))
                                            switch "" {
                                            case "\n\t":
                                                switch ("foo bar" + ("" + "")) {
                                                case "a":
                                                    switch fi6 {
                                                    case -2147483648:
                                                        var v49 float32 = 0.1
                                                        _ = v49
                                                    }
                                                }
                                            case "":
                                                fmt.Println(fi20)
                                            default:
                                                var v50 uint8 = v3
                                                _ = v50
                                            }
                                            _ = v47
                                            _ = v48
                                        }
                                        var v51 float64 = ((1e10 * -2.25) - float64((0.0 / ((float32(1e10) - -2.25) - 0.0))))
                                        _ = v35
                                        _ = v51
                                    } else {
                                        v23 -= (v26 | v23)
                                        switch (fi6 % (((^(fi20))) | 1)) {
                                        case -9223372036854775808:
                                            var v52 int64 = 1
                                            _ = v52
                                        }
                                    }
                                }
                                if (-2.25 == ((float32(float64(-2.25)) + float32(float64(0.1))) - ((-2.25 * float32((0.1 - (float64(((float32(3.14159) + (3.14159 - 0.1)) / ((0.0 * float32((-2.25 - 1e10))) + 1e10))) - (((-2.25 - 0.1) / 3.14159) * 0.0))))) * 3.14159))) {
                                    switch -864 {
                                    case 32767:
                                        if false {
                                            var v53 float32 = 3.14159
                                            var v54 float32 = float32(0.0)
                                            _ = v53
                                            _ = v54
                                        } else {
                                            var v55 int64 = 42
                                            if (-981 > -2147483648) {
                                                switch "hello" {
                                                case "a":
                                                    v56 := GenS4{F0: 2147483648, F1: fi20}
                                                    _ = v56
                                                case "":
                                                    v25 = (v25 + (^(v25)))
                                                case "\n\t":
                                                    delete(v2, uint32(v3))
                                                default:
                                                    switch "" {
                                                    case "hello":
                                                        var v57 int = fi6
                                                        _ = v57
                                                    case "foo bar":
                                                        var v58 bool = false
                                                        _ = v58
                                                    case "a":
                                                        v23 = (^(v24))
                                                    }
                                                }
                                                var v59 uint8 = (^(v3))
                                                _ = v59
                                            }
                                            switch (((("a" + "hello") + "foo bar") + "\n\t") + ("foo bar" + "hello")) {
                                            case "":
                                                if ("\n\t" >= ("" + "")) {
                                                    fmt.Println(v23)
                                                } else {
                                                    v25 -= 65535
                                                }
                                            case "foo bar":
                                                fmt.Println(v26)
                                            case "hello":
                                                v26 += v26
                                            }
                                            v60 := GenS3{F0: 65535}
                                            _ = v55
                                            _ = v60
                                        }
                                    case 542:
                                        if true {
                                            v1 = 2
                                        } else {
                                            var v61 uint8 = v3
                                            v62 := []int{(fi6 >> (((v1 / ((v1) | 1))) & 63))}
                                            var v63 uint32 = 2147483647
                                            _ = v61
                                            _ = v62
                                            _ = v63
                                        }
                                    case 817:
                                        if false {
                                            v26 = -1
                                            fmt.Println(fi20)
                                        } else {
                                            v23 += -351
                                            v26 = v23
                                        }
                                        if false {
                                            fmt.Println(v25, v27)
                                        } else {
                                            if false {
                                                var v64 uint16 = v25
                                                var v65 string = (("hello" + ((("" + "") + "hello") + "a")) + "a")
                                                switch ((v65 + v65) + "hello") {
                                                case "hello":
                                                    v27 += v27
                                                    v66 := GenS2{F0: (v1 >> ((v1) & 63))}
                                                    if (((("foo bar" + "") + (v65 + v65)) + "a") > (v65 + ("hello" + v65))) {
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            v26 = (v26 - (v26 | -32768))
                                                        }
                                                        v24 = int16(v1)
                                                        switch v65 {
                                                        case "foo bar":
                                                            if false {
                                                                v65 = (v65 + v65)
                                                                if false {
                                                                    var v67 int64 = 9223372036854775807
                                                                    if (float32(1e10) >= (1.5 - 1e10)) {
                                                                        v68 := []int{(fi20 * (^(fi6))), 9223372036854775807, int(v26)}
                                                                        switch v65 {
                                                                        case "a":
                                                                            if ((!(false && (true && false))) && true) {
                                                                                v69 := v2[2]
                                                                                v64 -= 312
                                                                                v24 = v26
                                                                                var v70 uint32 = 4294967295
                                                                                var v71 uint = 18446744073709551615
                                                                                _ = v69
                                                                                _ = v70
                                                                                _ = v71
                                                                            }
                                                                        case "hello":
                                                                            v72 := GenS2{F0: v1}
                                                                            _ = v72
                                                                        }
                                                                        var v73 int16 = (v26 & -24)
                                                                        _ = v68
                                                                        _ = v73
                                                                    }
                                                                    v74 := []int{fi6, (^(fi20)), fi20, -43}
                                                                    _ = v67
                                                                    _ = v74
                                                                }
                                                            }
                                                        case "a":
                                                            var v75 uint = (v1 % (((v1 | (^(v1)))) | 1))
                                                            _ = v75
                                                        default:
                                                            var v76 int8 = v27
                                                            _ = v76
                                                        }
                                                    }
                                                    _ = v66
                                                }
                                                _ = v64
                                                _ = v65
                                            }
                                            var v77 uint64 = 18446744073709551615
                                            v78 := GenS5{F0: (v27 | v27), F1: 1, F2: 4294967295}
                                            _ = v77
                                            _ = v78
                                        }
                                        if (0 == 9223372036854775807) {
                                            var v79 int16 = 42
                                            _ = v79
                                        } else {
                                            v80 := []int{(fi20 + -224), (fi20 & (fi20 * int(v23))), fi6}
                                            _ = v80
                                        }
                                    default:
                                        if true {
                                            v81 := GenS5{F0: (v27 / (((v27 - (v27 << ((4294967295) & 63)))) | 1)), F1: 2147483647, F2: 121}
                                            v82 := GenS6{F0: v27, F1: (-1 != 42), F2: (v25 * (v25 << (((v1 % ((v1) | 1))) & 63)))}
                                            if true {
                                                v83 := GenS3{F0: 4}
                                                _ = v83
                                            }
                                            _ = v81
                                            _ = v82
                                        } else {
                                            var v84 uint8 = (^(v3))
                                            v85 := v2[65535]
                                            v86 := GenS1{F0: (^(v25)), F1: "\n\t", F2: 2}
                                            if (false || (false || true)) {
                                                if (true && false) {
                                                    v85 = (v85 >> ((4294967295) & 63))
                                                    v87 := GenS7{F0: (uint(v85) == 65535)}
                                                    _ = v87
                                                }
                                                var v88 uint8 = 2
                                                _ = v88
                                            } else {
                                                v84 = v3
                                                v23 -= v26
                                                if (v3 == v3) {
                                                    var v89 int = (-(fi6))
                                                    var v90 int8 = (-(v27))
                                                    _ = v89
                                                    _ = v90
                                                }
                                            }
                                            var v91 float64 = 1.5
                                            _ = v84
                                            _ = v85
                                            _ = v86
                                            _ = v91
                                        }
                                    }
                                    if (-792 > -487) {
                                        v92 := []int{9223372036854775807}
                                        _ = v92
                                    }
                                } else {
                                    if false {
                                        switch ("\n\t" + ("\n\t" + ("a" + (("\n\t" + "\n\t") + "foo bar")))) {
                                        case "":
                                            var v93 float32 = -2.25
                                            _ = v93
                                        case "\n\t":
                                            v94 := []int{}
                                            _ = v94
                                        case "hello":
                                            v26 -= (v23 + (v23 & (v26 >> (((v1 << (((^(v1))) & 63))) & 63))))
                                        default:
                                            fmt.Println(v24, fi6, v27)
                                        }
                                    } else {
                                        switch fi6 {
                                        case 767:
                                            var v95 int32 = 2147483647
                                            _ = v95
                                        case 128:
                                            var v96 bool = (-1 == 0)
                                            _ = v96
                                        case 32768:
                                            switch ("foo bar" + "hello") {
                                            case "\n\t":
                                                v24 -= v24
                                            case "foo bar":
                                                if (false || true) {
                                                    var v97 float64 = 0.1
                                                    var v98 uint16 = (v25 | 1)
                                                    switch fi6 {
                                                    case 138:
                                                        if false {
                                                            switch "hello" {
                                                            case "":
                                                                v27 = (-(v27))
                                                            case "hello":
                                                                v97 += (v97 - ((1.5 + v97) + v97))
                                                            default:
                                                                v99 := GenS7{F0: false}
                                                                _ = v99
                                                            }
                                                        }
                                                    }
                                                    _ = v97
                                                    _ = v98
                                                } else {
                                                    if false {
                                                        v100 := GenS8{F0: (("foo bar" + ("\n\t" + ("a" + "\n\t"))) == ("hello" + ("foo bar" + ("a" + "hello")))), F1: -1, F2: 42}
                                                        if false {
                                                            v101 := v2[uint32(v25)]
                                                            _ = v101
                                                        }
                                                        _ = v100
                                                    }
                                                    _ = len(v2)
                                                    v102 := GenS9{F0: v3, F1: 0, F2: (v26 | v26)}
                                                    var v103 int16 = (v24 * (-(v23)))
                                                    fmt.Println(v24, v23)
                                                    _ = v102
                                                    _ = v103
                                                }
                                                v104 := []int{(fi6 ^ -612), fi6}
                                                v27 = v27
                                                if len(v104) >= 2 {
                                                    v104 = v104[1:len(v104)]
                                                }
                                                var v105 int64 = -1
                                                _ = v104
                                                _ = v105
                                            case "hello":
                                                v25 += (v25 & v25)
                                            }
                                        }
                                        var v106 bool = (false || true)
                                        v107 := GenS10{F0: 9223372036854775807}
                                        if v106 {
                                            var v108 int = fi6
                                            _ = v108
                                        } else {
                                            switch fi6 {
                                            case -9223372036854775808:
                                                v26 -= (v23 << ((v1) & 63))
                                            default:
                                                v26 -= (v26 / (((v26 >> ((255) & 63))) | 1))
                                                if (-32768 > (v26 * v23)) {
                                                    v109 := GenS3{F0: 4294967295}
                                                    _ = v109
                                                } else {
                                                    var v110 float64 = 0.1
                                                    _ = v110
                                                }
                                            }
                                            fmt.Println(fi6, v3)
                                            var v111 int8 = (^(v27))
                                            _ = v111
                                        }
                                        v25 += (v25 | (v25 * (^(v25))))
                                        _ = v106
                                        _ = v107
                                    }
                                    switch 128 {
                                    case 2147483647:
                                        if (!(v26 <= (v23 * 1000))) {
                                            if (0 == -925) {
                                                var v112 int = (-(fi6))
                                                _ = v112
                                            } else {
                                                if true {
                                                    if false {
                                                        if true {
                                                            v113 := []int{127, fi20, fi6}
                                                            _ = v113
                                                        } else {
                                                            if true {
                                                                fmt.Println(fi6)
                                                                if false {
                                                                    v114 := GenS5{F0: (v27 - v27), F1: 0, F2: 0}
                                                                    _ = len(v2)
                                                                    var v115 bool = false
                                                                    _ = v114
                                                                    _ = v115
                                                                }
                                                            }
                                                            v1 = 18446744073709551615
                                                            v116 := GenS11{F0: 2, F1: 65535}
                                                            _ = v116
                                                        }
                                                    }
                                                }
                                                fmt.Println(v27)
                                                switch (-(fi20)) {
                                                case 0:
                                                    var v117 int16 = v24
                                                    _ = v117
                                                case -1:
                                                    v25 = v25
                                                }
                                            }
                                        }
                                    case -2147483648:
                                        var v118 int64 = -9223372036854775808
                                        _ = v118
                                    case -991:
                                        switch (fi6 * int(v27)) {
                                        case 180:
                                            if false {
                                                var v119 uint64 = 2
                                                v120 := []int{(fi6 * fi6), (^(fi20))}
                                                _ = v119
                                                _ = v120
                                            }
                                        case 2147483647:
                                            var v121 int64 = 1
                                            _ = v121
                                        case -1:
                                            if (!true) {
                                                fmt.Println(v26, v26)
                                                switch fi20 {
                                                case 32768:
                                                    var v122 int64 = 584
                                                    _ = v122
                                                }
                                                v25 = (v25 >> ((v1) & 63))
                                            } else {
                                                v26 = v24
                                                switch (("hello" + "hello") + "\n\t") {
                                                case "a":
                                                    if (v27 <= 100) {
                                                        var v123 bool = false
                                                        fmt.Println(v24, v26)
                                                        v123 = (!(-1 >= 0))
                                                        v25 = (v25 * (v25 | (v25 | (v25 << (((v1 & 2)) & 63)))))
                                                        _ = v123
                                                    }
                                                default:
                                                    v124 := GenS12{F0: v1, F1: 42, F2: 0}
                                                    _ = v124
                                                }
                                            }
                                        default:
                                            var v125 float32 = 1e10
                                            _ = v125
                                        }
                                    default:
                                        fmt.Println(v1, v1)
                                    }
                                }
                                _ = v25
                                _ = v26
                                _ = v27
                            }
                        } else {
                            var v126 int8 = -128
                            _ = v126
                        }
                        var v127 string = ""
                        _ = fi20
                        _ = v127
                        _ = v23
                        _ = v24
                    }
                }
            }
        } else {
            switch (fi6 % ((fi6) | 1)) {
            case 32767:
                fmt.Println(v1)
            case 9223372036854775807:
                for fi128 := 0; fi128 < 3; fi128++ {
                    v1 = v1
                    v129 := GenS4{F0: 880, F1: (fi6 | fi128)}
                    v1 -= v1
                    switch "" {
                    case "":
                        delete(v2, 2147483647)
                    case "a":
                        v2[0] = 909
                    }
                    _ = fi128
                    _ = v129
                }
            }
            switch ("a" + "a") {
            case "hello":
                v3 = v3
            case "":
                v3 = v3
            }
            fmt.Println(v3)
        }
        switch "foo bar" {
        case "foo bar":
            var v130 float64 = (float64(1e10) / 1e10)
            _ = v130
        case "a":
            for fi131 := 0; fi131 < 20; fi131++ {
                var v132 float64 = -2.25
                v133 := GenS13{F0: 42, F1: 4294967295}
                _ = fi131
                _ = v132
                _ = v133
            }
        }
        _ = fi6
        _ = v7
    }
    switch ("hello" + ("\n\t" + ("\n\t" + "a"))) {
    case "foo bar":
        for fi134 := 0; fi134 < 20; fi134++ {
            var v135 bool = (!true)
            var v136 string = ("" + "hello")
            _ = fi134
            _ = v135
            _ = v136
        }
    default:
        for fi137 := 0; fi137 < 13; fi137++ {
            var v138 string = "\n\t"
            _ = fi137
            _ = v138
        }
    }
    if (!(!(0 == v1))) {
        var v139 int8 = -117
        fmt.Println(v139, v139)
        v140 := GenS1{F0: 575, F1: "foo bar", F2: 0}
        _ = v139
        _ = v140
    } else {
        var v141 bool = true
        _ = v141
    }
    switch "\n\t" {
    case "foo bar":
        if true {
            for fi142 := 0; fi142 < 6; fi142++ {
                var v143 uint8 = uint8(v1)
                v143 -= v143
                _ = fi142
                _ = v143
            }
        }
    case "a":
        v144 := GenS14{F0: 42}
        _ = v144
    case "hello":
        v145 := GenS15{F0: (v3 % (((v3 << (((v1 | v1)) & 63))) | 1)), F1: 1, F2: 32767}
        _ = v145
    }
    for fi146 := 0; fi146 < 19; fi146++ {
        if true {
            switch (fi146 - (-(fi146))) {
            case -513:
                if (!(0.0 >= 1e10)) {
                    delete(v2, 0)
                } else {
                    for fi147 := 0; fi147 < 18; fi147++ {
                        switch fi147 {
                        case -860:
                            v148 := GenS11{F0: 2147483647, F1: 330}
                            _ = v148
                        case 32767:
                            v149 := GenS12{F0: 42, F1: 0, F2: 255}
                            _ = v149
                        case -9223372036854775808:
                            if (true && (-2.25 != (0.0 * 1e10))) {
                                if true {
                                    switch "hello" {
                                    case "\n\t":
                                        v150 := GenS16{F0: ("" + ("" + ("\n\t" + "hello"))), F1: -828, F2: 384}
                                        _ = v150
                                    case "":
                                        var v151 string = "a"
                                        _ = v151
                                    case "hello":
                                        if false {
                                            if (!true) {
                                                v3 -= v3
                                                v1 -= (v1 & v1)
                                                v3 = v3
                                            }
                                            v1 += v1
                                            var v152 uint8 = (v3 - 42)
                                            _ = v152
                                        }
                                    }
                                    var v153 float64 = 3.14159
                                    var v154 float32 = 1.5
                                    fmt.Println(fi147, v3, v153)
                                    _ = v153
                                    _ = v154
                                }
                                if true {
                                    if false {
                                        if (-34 < 73) {
                                            var v155 uint16 = 65535
                                            _ = v155
                                        } else {
                                            v156 := GenS10{F0: -9223372036854775808}
                                            v157 := GenS17{F0: 831}
                                            _ = v156
                                            _ = v157
                                        }
                                    } else {
                                        if false {
                                            fmt.Println(fi147)
                                            v3 = (^(v3))
                                        }
                                        v158 := GenS18{F0: true, F1: 132, F2: "foo bar"}
                                        var v159 int16 = 1000
                                        _ = v158
                                        _ = v159
                                    }
                                }
                                switch ("hello" + "hello") {
                                case "hello":
                                    var v160 string = ""
                                    _ = v160
                                case "":
                                    fmt.Println(v1, v3, v1)
                                case "\n\t":
                                    var v161 int = (fi146 << (((v1 ^ v1)) & 63))
                                    _ = v161
                                }
                            }
                        }
                        var v162 int64 = 1
                        _ = fi147
                        _ = v162
                    }
                    switch "hello" {
                    case "foo bar":
                        if (1 >= 42) {
                            _ = len(v2)
                            if false {
                                if false {
                                    fmt.Println(fi146)
                                } else {
                                    for fi163 := 0; fi163 < 1; fi163++ {
                                        if (4294967295 == uint64(v3)) {
                                            var v164 uint = 65535
                                            switch (fi146 * -102) {
                                            case 560:
                                                var v165 uint = (v164 << ((v164) & 63))
                                                _ = v165
                                            case 42:
                                                v166 := v2[1]
                                                _ = v166
                                            }
                                            switch (fi146 ^ 828) {
                                            case -2147483648:
                                                v167 := []int{}
                                                _ = v167
                                            case 9223372036854775807:
                                                fmt.Println(fi163)
                                            case 1:
                                                v168 := GenS19{F0: 2147483648, F1: 42}
                                                _ = v168
                                            default:
                                                v3 = v3
                                            }
                                            v169 := GenS8{F0: (-573 <= -1), F1: -1, F2: 127}
                                            _ = v164
                                            _ = v169
                                        }
                                        _ = fi163
                                    }
                                    v1 -= (v1 & (^(v1)))
                                    if (!(((0.0 / 0.1) / (1.5 / (1.5 + 1e10))) > (3.14159 / 1.5))) {
                                        if (v3 != (v3 - 3)) {
                                            var v170 float64 = ((1e10 + (((0.0 - 1e10) * float64(0.1)) * 1.5)) * ((3.14159 * float64(3.14159)) / 3.14159))
                                            _ = v170
                                        }
                                    }
                                }
                            }
                            v3 = (v3 << ((v1) & 63))
                        } else {
                            if true {
                                v171 := []int{fi146, (fi146 + (fi146 - int(v1))), fi146}
                                _ = v171
                            }
                        }
                    default:
                        v2[1] = 4294967295
                    }
                    v3 = 2
                }
            case 42:
                v3 += (^(v3))
            case -9223372036854775808:
                switch fi146 {
                case 0:
                    fmt.Println(v1)
                case 453:
                    if (!true) {
                        v172 := []int{fi146, 2147483647}
                        var v173 int8 = 0
                        _ = v172
                        _ = v173
                    }
                case -368:
                    v1 = (v1 - v1)
                }
            }
            var v174 float32 = 1e10
            v175 := GenS20{F0: v3}
            _ = v174
            _ = v175
        } else {
            v176 := GenS1{F0: 859, F1: "foo bar", F2: 141}
            if (9223372036854775807 <= -1) {
                if false {
                    v1 += (v1 | (^(v1)))
                    v3 -= 42
                }
                var v177 uint = 2
                _ = v177
            } else {
                if ((897 >= 497) && ((false && false) || (false || (2147483647 != 42)))) {
                    switch "foo bar" {
                    case "a":
                        fmt.Println(fi146, fi146)
                    default:
                        v178 := GenS21{F0: -1, F1: (v3 % ((v3) | 1)), F2: 1}
                        _ = v178
                    }
                    v1 -= (^(v1))
                    switch "foo bar" {
                    case "":
                        var v179 uint = 0
                        _ = v179
                    case "a":
                        var v180 string = "foo bar"
                        _ = v180
                    case "hello":
                        v181 := GenS22{F0: 469, F1: -1, F2: -128}
                        _ = v181
                    }
                }
            }
            _ = v176
        }
        fmt.Println(fi146, v1, v1)
        _ = fi146
    }
    v182 := GenS12{F0: v1, F1: 65535, F2: 32767}
    for fi183 := 0; fi183 < 1; fi183++ {
        v184 := v2[2147483647]
        var v185 int8 = 98
        for fi186 := 0; fi186 < 5; fi186++ {
            fmt.Println(v185)
            _ = fi186
        }
        _ = fi183
        _ = v184
        _ = v185
    }
    for fi187 := 0; fi187 < 2; fi187++ {
        var v188 uint8 = (^(v3))
        _ = fi187
        _ = v188
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v182
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
