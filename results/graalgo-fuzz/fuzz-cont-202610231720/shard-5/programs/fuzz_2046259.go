package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 int16; F1 int8; F2 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint8; F1 uint16; F2 int8 }
type GenS7 struct { F0 int; F1 int64; F2 uint32 }
type GenS8 struct { F0 uint32; F1 uint; F2 int }
type GenS9 struct { F0 uint16; F1 int16 }
type GenS10 struct { F0 uint8; F1 int16 }
type GenS11 struct { F0 int64; F1 uint8 }
type GenS12 struct { F0 uint; F1 int16; F2 uint64 }
type GenS13 struct { F0 int32; F1 int; F2 int32 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 int32; F1 uint8 }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 int16 }
type GenS18 struct { F0 int64; F1 bool; F2 uint16 }
type GenS19 struct { F0 uint64 }
type GenS20 struct { F0 int8 }
type GenS21 struct { F0 int32; F1 uint; F2 int32 }
type GenS22 struct { F0 uint64; F1 uint8; F2 bool }
type GenS23 struct { F0 int; F1 bool; F2 uint32 }
func genFn0(p0 int64, p1 uint16, p2 int8) int16 {
    p1 = 904
    v7 := GenS1{F0: 122, F1: 0}
    var v8 uint = uint(p1)
    switch -9223372036854775808 {
    case 9223372036854775807:
        v9 := GenS1{F0: p0, F1: 128}
        _ = v9
    case 32768:
        if false {
            p2 = p2
            v10 := GenS2{F0: 225}
            var v11 bool = ((3.14159 / 0.1) == 3.14159)
            _ = v10
            _ = v11
        } else {
            var v12 int32 = -939
            if false {
                v12 -= 607
            } else {
                switch (((("hello" + "foo bar") + "hello") + "") + "a") {
                case "a":
                    var v13 int32 = v12
                    _ = v13
                case "foo bar":
                    if (false || (((^(v8)) < v8) || (!((false && true) || true)))) {
                        v14 := GenS2{F0: 28}
                        if false {
                            switch (("\n\t" + "\n\t") + "") {
                            case "a":
                                for fi15 := 0; fi15 < 3; fi15++ {
                                    if false {
                                        v12 = (-(v12))
                                        for fi16 := 0; fi16 < 17; fi16++ {
                                            var v17 int32 = v12
                                            switch (("" + "\n\t") + "") {
                                            case "\n\t":
                                                switch ("\n\t" + "hello") {
                                                case "foo bar":
                                                    var v18 int8 = 42
                                                    _ = v18
                                                }
                                            case "a":
                                                switch (fi15 & (-(fi16))) {
                                                case -1:
                                                    var v19 uint8 = uint8(v8)
                                                    _ = v19
                                                case 0:
                                                    var v20 int32 = v17
                                                    _ = v20
                                                default:
                                                    var v21 bool = true
                                                    _ = v21
                                                }
                                            }
                                            _ = fi16
                                            _ = v17
                                        }
                                        v22 := GenS1{F0: p0, F1: 255}
                                        _ = v22
                                    } else {
                                        v23 := []int{fi15, (fi15 & fi15)}
                                        if true {
                                            for fi24 := 0; fi24 < 4; fi24++ {
                                                var v25 float64 = 0.0
                                                _ = fi24
                                                _ = v25
                                            }
                                            v12 += -43
                                        } else {
                                            if ((p0 <= (p0 - p0)) && ((p0 + p0) != (p0 << (((v8 | 1)) & 63)))) {
                                                v26 := GenS3{F0: 2147483647, F1: p2}
                                                _ = v26
                                            }
                                        }
                                        v12 -= 2147483647
                                        _ = v23
                                    }
                                    v8 = 1
                                    _ = fi15
                                }
                            case "":
                                v27 := []int{}
                                _ = v27
                            }
                            for fi28 := 0; fi28 < 1; fi28++ {
                                p2 -= (-(p2))
                                p0 = p0
                                _ = fi28
                            }
                            switch "a" {
                            case "foo bar":
                                var v29 uint8 = 0
                                _ = v29
                            default:
                                var v30 int64 = p0
                                _ = v30
                            }
                        }
                        _ = v14
                    }
                case "":
                    if (128 != 42) {
                        v31 := []int{127, -354, -928}
                        switch 1 {
                        case 42:
                            fmt.Println(p0)
                        case -2147483648:
                            v32 := GenS4{F0: 0, F1: p2, F2: 42}
                            _ = v32
                        default:
                            var v33 int16 = 727
                            _ = v33
                        }
                        _ = v31
                    }
                default:
                    var v34 string = ((("\n\t" + "") + "hello") + (("a" + "\n\t") + ("\n\t" + "foo bar")))
                    _ = v34
                }
                p0 = p0
                p0 = p0
            }
            _ = v12
        }
    }
    v35 := GenS5{F0: ("hello" + "a")}
    _ = p0
    _ = p1
    _ = p2
    _ = v35
    _ = v7
    _ = v8
    return -723
}

func main() {
    v1 := make(map[int]string)
    v2 := []string{((("\n\t" + "a") + "foo bar") + ("a" + ("foo bar" + (((("a" + "") + "foo bar") + "foo bar") + "hello"))))}
    var v3 int16 = -32768
    v4 := []uint32{2, 2}
    var v5 uint64 = 42
    var v6 uint = 719
    v1[int(v6)] = "foo bar"
    fmt.Println(v6)
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    var v36 uint16 = 255
    var v37 int8 = -128
    switch ("hello" + "hello") {
    case "foo bar":
        v38 := GenS1{F0: 0, F1: uint8(v5)}
        _ = v38
    case "\n\t":
        v39 := GenS6{F0: 2, F1: 42, F2: -1}
        _ = v39
    case "":
        if false {
            var v40 uint32 = uint32(v5)
            switch ("\n\t" + "hello") {
            case "hello":
                switch "foo bar" {
                case "hello":
                    var v41 uint8 = 128
                    _ = v41
                case "foo bar":
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(v40, v40, v6)
                    }
                case "\n\t":
                    _ = len(v2)
                default:
                    if true {
                        if false {
                            if len(v2) > 0 {
                                v2[0] = ("foo bar" + "\n\t")
                            }
                            if (v6 <= 4294967295) {
                                _ = len(v2)
                            }
                            var v42 int = int(v6)
                            _ = v42
                        }
                        v43 := GenS7{F0: -15, F1: -224, F2: (v40 & 2147483647)}
                        _ = len(v4)
                        _ = v43
                    }
                }
            case "\n\t":
                if true {
                    for fi44 := 0; fi44 < 19; fi44++ {
                        if false {
                            for fi45 := 0; fi45 < 8; fi45++ {
                                var v46 float32 = (1e10 + 1e10)
                                v47 := GenS7{F0: fi45, F1: 0, F2: 65535}
                                _ = fi45
                                _ = v46
                                _ = v47
                            }
                            for fi48 := 0; fi48 < 16; fi48++ {
                                var v49 string = ("hello" + (("\n\t" + "\n\t") + "a"))
                                var v50 int = (fi48 ^ 127)
                                v50 -= (v50 % (((fi48 | (v50 % ((fi44) | 1)))) | 1))
                                _ = fi48
                                _ = v49
                                _ = v50
                            }
                            v51 := GenS1{F0: 1, F1: 2}
                            _ = v51
                        } else {
                            fmt.Println(fi44, v36, v6)
                            switch fi44 {
                            case 9223372036854775807:
                                v52 := GenS7{F0: (-(fi44)), F1: 9223372036854775807, F2: uint32(v6)}
                                _ = v52
                            case 2147483647:
                                v53 := GenS8{F0: uint32(v5), F1: (^(v6)), F2: 9223372036854775807}
                                _ = v53
                            case 32767:
                                v54 := GenS9{F0: 1, F1: int16(v37)}
                                _ = v54
                            default:
                                fmt.Println(v40)
                            }
                            fmt.Println(v3, v36)
                        }
                        v37 = (^(v37))
                        v55 := GenS3{F0: v40, F1: (-(v37))}
                        _ = fi44
                        _ = v55
                    }
                    if true {
                        v36 = v36
                        var v56 int = 128
                        v5 = v5
                        _ = v56
                    }
                } else {
                    if (v3 >= (v3 % (((v3 ^ v3)) | 1))) {
                        v57 := GenS10{F0: 255, F1: v3}
                        v3 = (v3 * (-(v3)))
                        _ = v57
                    } else {
                        v58 := GenS11{F0: -1, F1: 128}
                        _ = v58
                    }
                    var v59 int8 = v37
                    v37 = 1
                    _ = v59
                }
            default:
                var v60 bool = (!false)
                _ = v60
            }
            v37 = (v37 - 100)
            _ = v40
        }
    }
    _ = len(v1)
    switch -9223372036854775808 {
    case 9223372036854775807:
        switch 358 {
        case -510:
            var v61 uint16 = v36
            _ = v61
        case 9223372036854775807:
            var v62 uint = v6
            _ = v62
        case 128:
            v63 := GenS12{F0: (v6 ^ v6), F1: v3, F2: v5}
            _ = v63
        }
    case 2147483647:
        v37 = (v37 & -1)
    case 127:
        if (241 >= 255) {
            for fi64 := 0; fi64 < 11; fi64++ {
                v1[128] = ""
                if (-2147483648 > 974) {
                    _ = len(v4)
                    _ = len(v2)
                    for fi65 := 0; fi65 < 15; fi65++ {
                        switch fi65 {
                        case 32767:
                            if true {
                                var v66 uint8 = 255
                                v67 := GenS13{F0: 548, F1: fi65, F2: -1}
                                _ = v66
                                _ = v67
                            } else {
                                if (!true) {
                                    v68 := GenS2{F0: 1}
                                    if len(v4) > 0 {
                                        v4[0] = uint32(v6)
                                    }
                                    var v69 int = fi65
                                    _ = v68
                                    _ = v69
                                } else {
                                    v5 = v5
                                    fmt.Println(v6, fi64, v5)
                                    var v70 float32 = 0.0
                                    _ = v70
                                }
                                v5 = (v5 / ((42) | 1))
                                v1[(fi64 % ((fi65) | 1))] = ("foo bar" + "hello")
                            }
                        case 2147483647:
                            v3 += (v3 / (((v3 - (^(v3)))) | 1))
                        case 32768:
                            switch ("a" + "foo bar") {
                            case "\n\t":
                                var v71 string = "a"
                                _ = v71
                            case "foo bar":
                                v72 := v1[(fi64 ^ (fi65 ^ fi64))]
                                _ = v72
                            }
                        default:
                            v3 = (v3 / (((v3 % (((v3 - (v3 + v3))) | 1))) | 1))
                        }
                        _ = fi65
                    }
                } else {
                    v6 += v6
                    fmt.Println(v37, v3, v6)
                }
                v37 = -128
                _ = fi64
            }
            if (-2.25 != float32((float64(float32(0.0)) / -2.25))) {
                var v73 int16 = (v3 << ((2) & 63))
                if (true || (0.0 >= 0.1)) {
                    v6 += v6
                } else {
                    var v74 bool = (4294967295 <= 554)
                    v5 += v5
                    var v75 bool = true
                    _ = v74
                    _ = v75
                }
                _ = v73
            }
        } else {
            switch (("a" + "") + "foo bar") {
            case "\n\t":
                if false {
                    var v76 int8 = (v37 % ((v37) | 1))
                    fmt.Println(v36)
                    for fi77 := 0; fi77 < 5; fi77++ {
                        v76 -= (v76 >> (((v6 - uint(v36))) & 63))
                        _ = fi77
                    }
                    _ = v76
                }
            case "hello":
                v78 := GenS14{F0: (v5 - v5)}
                _ = v78
            case "a":
                for fi79 := 0; fi79 < 13; fi79++ {
                    fmt.Println(v3, fi79)
                    switch ("foo bar" + "\n\t") {
                    case "\n\t":
                        fmt.Println(fi79, v36)
                    case "foo bar":
                        var v80 uint64 = (v5 ^ v5)
                        _ = v80
                    default:
                        for fi81 := 0; fi81 < 11; fi81++ {
                            v36 = 934
                            if (65535 == 0) {
                                var v82 uint32 = uint32(v36)
                                switch "" {
                                case "hello":
                                    if (3.14159 > 1e10) {
                                        v6 = v6
                                        if (61 != int8(fi81)) {
                                            var v83 uint16 = v36
                                            v84 := GenS9{F0: uint16(v82), F1: genFn0(9223372036854775807, 0, -1)}
                                            _ = v83
                                            _ = v84
                                        } else {
                                            v5 += v5
                                            if (-2.25 > -2.25) {
                                                var v85 uint = 219
                                                v3 = (v3 ^ v3)
                                                _ = v85
                                            }
                                            fmt.Println(v37)
                                        }
                                        var v86 string = "a"
                                        _ = v86
                                    } else {
                                        v36 = (v36 / ((v36) | 1))
                                        if false {
                                            v36 = (v36 | 0)
                                            v36 -= v36
                                            var v87 float32 = (0.1 - 1e10)
                                            _ = v87
                                        } else {
                                            v36 = uint16(v5)
                                        }
                                    }
                                default:
                                    switch ("\n\t" + "hello") {
                                    case "a":
                                        fmt.Println(v3)
                                    case "":
                                        v36 = v36
                                        if ((v6 & (^(v6))) == v6) {
                                            v3 -= (v3 >> ((uint(v36)) & 63))
                                        }
                                        switch (fi81 / ((fi79) | 1)) {
                                        case 413:
                                            v5 = 18446744073709551615
                                        }
                                    case "hello":
                                        var v88 float32 = float32(3.14159)
                                        _ = v88
                                    default:
                                        v36 = 340
                                        switch (fi79 + (fi81 | (^(fi81)))) {
                                        case 32768:
                                            if false {
                                                var v89 float32 = (((1.5 / 1.5) + float32(0.1)) + 0.1)
                                                if ((v5 <= v5) && true) {
                                                    if true {
                                                        v90 := v1[fi81]
                                                        if (true || true) {
                                                            var v91 int16 = (v3 & v3)
                                                            v92 := GenS15{F0: 0, F1: 0}
                                                            _ = v91
                                                            _ = v92
                                                        }
                                                        if (false && ((-2.25 < v89) || false)) {
                                                            var v93 int64 = -9223372036854775808
                                                            v3 = v3
                                                            _ = v93
                                                        }
                                                        fmt.Println(v89, v3)
                                                        v4 = append(v4, (v82 - (^(v82))))
                                                        _ = v90
                                                    } else {
                                                        v94 := GenS16{F0: true}
                                                        var v95 bool = false
                                                        _ = v94
                                                        _ = v95
                                                    }
                                                    v82 = uint32(v36)
                                                }
                                                var v96 bool = false
                                                _ = v89
                                                _ = v96
                                            }
                                        case -1:
                                            v36 = 255
                                        default:
                                            var v97 int8 = (v37 >> ((18446744073709551615) & 63))
                                            _ = v97
                                        }
                                        v98 := GenS15{F0: 42, F1: 42}
                                        _ = v98
                                    }
                                }
                                v99 := GenS17{F0: (^(v3))}
                                var v100 int8 = (v37 | (v37 | 1))
                                var v101 int8 = v100
                                _ = v100
                                _ = v101
                                _ = v82
                                _ = v99
                            } else {
                                _ = len(v2)
                            }
                            if false {
                                v37 -= v37
                                v102 := GenS5{F0: "hello"}
                                _ = v102
                            }
                            _ = fi81
                        }
                    }
                    _ = fi79
                }
            default:
                v4 = append(v4, 2)
            }
        }
    }
    switch 4294967295 {
    case -222:
        for fi103 := 0; fi103 < 3; fi103++ {
            fmt.Println(v6, v37)
            v37 -= v37
            v104 := v1[fi103]
            _ = fi103
            _ = v104
        }
    }
    if (true || ((v3 != (^(v3))) || true)) {
        v105 := GenS18{F0: 1, F1: (false && false), F2: v36}
        _ = v105
    } else {
        for fi106 := 0; fi106 < 20; fi106++ {
            for fi107 := 0; fi107 < 18; fi107++ {
                v108 := GenS3{F0: 2147483648, F1: (-(v37))}
                if true {
                    var v109 bool = true
                    _ = v109
                } else {
                    var v110 float64 = ((3.14159 + 0.0) * -2.25)
                    v5 = v5
                    var v111 int = -484
                    _ = v110
                    _ = v111
                }
                var v112 int16 = (v3 ^ (^(v3)))
                _ = fi107
                _ = v108
                _ = v112
            }
            v4 = append(v4, 2147483648)
            _ = fi106
        }
        if (0.1 >= ((1.5 * -2.25) / 1.0)) {
            var v113 uint64 = (v5 + uint64(v6))
            v114 := GenS19{F0: (v113 >> ((v6) & 63))}
            _ = v113
            _ = v114
        } else {
            if (!true) {
                for fi115 := 0; fi115 < 20; fi115++ {
                    v116 := GenS10{F0: 127, F1: (^(v3))}
                    if false {
                        v37 = (v37 % ((v37) | 1))
                        if false {
                            v117 := GenS8{F0: 0, F1: 1, F2: fi115}
                            for fi118 := 0; fi118 < 4; fi118++ {
                                var v119 int16 = v3
                                v120 := GenS20{F0: v37}
                                _ = fi118
                                _ = v119
                                _ = v120
                            }
                            for fi121 := 0; fi121 < 9; fi121++ {
                                if true {
                                    v6 += v6
                                }
                                _ = fi121
                            }
                            _ = v117
                        }
                        var v122 int32 = -2147483648
                        _ = v122
                    }
                    _ = fi115
                    _ = v116
                }
                v37 += 100
                var v123 uint8 = 178
                _ = v123
            }
            var v124 int16 = v3
            v125 := GenS21{F0: 892, F1: (v6 | v6), F2: 739}
            _ = v124
            _ = v125
        }
        v1[42] = "a"
    }
    for fi126 := 0; fi126 < 2; fi126++ {
        fmt.Println(fi126, v5)
        var v127 bool = false
        if (!v127) {
            if len(v4) > 0 {
                v4[0] = 1
            }
            v36 = (v36 / (((v36 ^ (^(v36)))) | 1))
            if false {
                fmt.Println(v5)
                v128 := GenS22{F0: v5, F1: 72, F2: v127}
                v129 := v1[fi126]
                _ = v128
                _ = v129
            } else {
                for fi130 := 0; fi130 < 5; fi130++ {
                    if (v36 <= v36) {
                        var v131 float32 = (-2.25 + float32(3.14159))
                        var v132 int64 = 9223372036854775807
                        var v133 int32 = 84
                        _ = v131
                        _ = v132
                        _ = v133
                    } else {
                        v1[(fi130 | fi126)] = ((("\n\t" + "\n\t") + "hello") + "a")
                        fmt.Println(v37, v37, v5)
                    }
                    v37 += (v37 / ((v37) | 1))
                    _ = fi130
                }
                fmt.Println(v127, v6)
                for fi134 := 0; fi134 < 8; fi134++ {
                    switch "foo bar" {
                    case "foo bar":
                        v135 := v1[4294967295]
                        _ = v135
                    case "\n\t":
                        v5 = v5
                    }
                    _ = fi134
                }
            }
        } else {
            fmt.Println(fi126)
            v36 = 255
        }
        _ = fi126
        _ = v127
    }
    switch 9223372036854775807 {
    case -669:
        v136 := GenS12{F0: (^(v6)), F1: (-(v3)), F2: v5}
        _ = v136
    }
    switch "\n\t" {
    case "a":
        var v137 int16 = (v3 % (((v3 * genFn0(564, (v36 % (((v36 >> ((v6) & 63))) | 1)), v37))) | 1))
        _ = v137
    case "foo bar":
        if (!(("hello" + (("\n\t" + "foo bar") + "hello")) != "\n\t")) {
            for fi138 := 0; fi138 < 20; fi138++ {
                var v139 float64 = 1.5
                _ = fi138
                _ = v139
            }
            var v140 uint16 = 215
            _ = v140
        }
    case "":
        for fi141 := 0; fi141 < 19; fi141++ {
            fmt.Println(v5, v36)
            _ = fi141
        }
    default:
        switch -877 {
        case -1:
            if true {
                v5 = (^(v5))
                var v142 int32 = 1
                _ = v142
            } else {
                var v143 uint64 = (v5 * (^(v5)))
                v36 = (v36 % ((v36) | 1))
                var v144 string = "foo bar"
                _ = v143
                _ = v144
            }
        }
    }
    if false {
        v5 = (v5 * (v5 | v5))
        v145 := GenS23{F0: 32768, F1: true, F2: 0}
        v5 -= v5
        _ = v145
    } else {
        v5 += 155
        for fi146 := 0; fi146 < 10; fi146++ {
            for fi147 := 0; fi147 < 10; fi147++ {
                v148 := GenS14{F0: 833}
                v3 += 1
                _ = fi147
                _ = v148
            }
            _ = fi146
        }
        switch "foo bar" {
        case "hello":
            v37 = 100
        default:
            var v149 uint64 = (v5 >> ((v6) & 63))
            _ = v149
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
