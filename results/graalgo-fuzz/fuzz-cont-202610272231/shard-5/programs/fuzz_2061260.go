package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 uint8; F2 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32; F1 int32; F2 uint64 }
type GenS6 struct { F0 int16; F1 int; F2 int16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int64; F1 uint64; F2 string }
type GenS9 struct { F0 uint64; F1 uint64 }
type GenS10 struct { F0 int8; F1 uint32 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int; F1 uint; F2 uint16 }
type GenS13 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS14 struct { F0 uint16; F1 int32 }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 uint32; F1 bool }
type GenS17 struct { F0 uint8 }
type GenS18 struct { F0 int16 }
type GenS19 struct { F0 int }
type GenS20 struct { F0 int32; F1 int64; F2 uint32 }
type GenS21 struct { F0 uint16; F1 int32; F2 uint }
type GenS22 struct { F0 int8 }
type GenS23 struct { F0 int16; F1 string; F2 string }
func genFn0(p0 uint, p1 bool, p2 int16) bool {
    var v7 uint = p0
    var v8 uint32 = 2147483648
    if (p1 && (0 >= (p2 | (p2 ^ (p2 | (p2 * p2)))))) {
        v9 := GenS1{F0: (^(v8))}
        var v10 int = 1
        v11 := []int{(v10 - (v10 / ((v10) | 1))), v10}
        _ = v10
        _ = v11
        _ = v9
    } else {
        if p1 {
            for fi12 := 0; fi12 < 16; fi12++ {
                if p1 {
                    var v13 int = fi12
                    fmt.Println(p1, v7)
                    for fi14 := 0; fi14 < 17; fi14++ {
                        v15 := []int{-105, -1}
                        if (p1 && p1) {
                            if p1 {
                                v15 = append(v15, (-(fi14)))
                                if (true && p1) {
                                    switch "foo bar" {
                                    case "foo bar":
                                        fmt.Println(v8)
                                    case "a":
                                        p1 = (102 < 66)
                                    }
                                    switch ("\n\t" + "\n\t") {
                                    case "a":
                                        var v16 string = ("hello" + "\n\t")
                                        _ = v16
                                    case "":
                                        if len(v15) >= 2 {
                                            v15 = v15[1:len(v15)]
                                        }
                                    }
                                    v15 = append(v15, fi12)
                                    p2 = p2
                                    var v17 int32 = 1
                                    v18 := GenS2{F0: p0, F1: 128, F2: 445}
                                    _ = v17
                                    _ = v18
                                } else {
                                    _ = len(v15)
                                }
                                var v19 bool = p1
                                var v20 int16 = p2
                                _ = v19
                                _ = v20
                            } else {
                                var v21 uint16 = 42
                                _ = v21
                            }
                        } else {
                            switch ("hello" + "\n\t") {
                            case "":
                                if (!p1) {
                                    var v22 bool = true
                                    _ = v22
                                } else {
                                    v23 := GenS2{F0: 2, F1: 1, F2: 0}
                                    fmt.Println(v13, fi12, p2)
                                    p2 -= 885
                                    if len(v15) >= 2 {
                                        v15 = v15[1:len(v15)]
                                    }
                                    if (946 <= 9223372036854775807) {
                                        v24 := GenS3{F0: 1}
                                        var v25 uint64 = 0
                                        if p1 {
                                            v13 = 9223372036854775807
                                        } else {
                                            var v26 int32 = -990
                                            if ((!p1) && p1) {
                                                v25 += v25
                                                if false {
                                                    var v27 uint16 = 1
                                                    var v28 uint64 = 1
                                                    switch "\n\t" {
                                                    case "hello":
                                                        if (v27 <= 255) {
                                                            p0 += p0
                                                            v25 += v28
                                                            v29 := GenS4{F0: 162}
                                                            switch fi14 {
                                                            case 96:
                                                                if p1 {
                                                                    p2 = (p2 << (((^(v7))) & 63))
                                                                    v30 := GenS3{F0: 255}
                                                                    v25 -= v25
                                                                    v15 = append(v15, (-(fi12)))
                                                                    var v31 int64 = -290
                                                                    _ = v30
                                                                    _ = v31
                                                                } else {
                                                                    _ = len(v15)
                                                                    v32 := GenS5{F0: (-(v26)), F1: 1, F2: uint64(v27)}
                                                                    v33 := GenS6{F0: 428, F1: 286, F2: p2}
                                                                    _ = v32
                                                                    _ = v33
                                                                }
                                                            default:
                                                                var v34 string = "\n\t"
                                                                _ = v34
                                                            }
                                                            _ = v29
                                                        } else {
                                                            if false {
                                                                if false {
                                                                    var v35 uint64 = v28
                                                                    _ = v35
                                                                } else {
                                                                    v28 -= (v25 >> ((v7) & 63))
                                                                }
                                                                v36 := GenS2{F0: v7, F1: 2, F2: (v26 | v26)}
                                                                _ = v36
                                                            }
                                                        }
                                                        p0 = 65535
                                                    case "foo bar":
                                                        switch ((("hello" + (((((("foo bar" + "foo bar") + "hello") + "hello") + ((("\n\t" + "a") + "hello") + "hello")) + ("hello" + "foo bar")) + "")) + "hello") + "hello") {
                                                        case "a":
                                                            fmt.Println(p2, p0, v25)
                                                            switch -143 {
                                                            case 4294967295:
                                                                var v37 uint16 = v27
                                                                _ = v37
                                                            case -1:
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    if len(v15) >= 2 {
                                                                        v15 = v15[1:len(v15)]
                                                                    }
                                                                case "":
                                                                    var v38 int8 = int8(v28)
                                                                    _ = v38
                                                                case "hello":
                                                                    fmt.Println(p2)
                                                                default:
                                                                    if len(v15) > 0 {
                                                                        v15[0] = (-(fi14))
                                                                    }
                                                                }
                                                            default:
                                                                var v39 uint16 = 32767
                                                                _ = v39
                                                            }
                                                            var v40 uint64 = (^(v25))
                                                            _ = v40
                                                        }
                                                    }
                                                    _ = v27
                                                    _ = v28
                                                }
                                                var v41 float32 = -2.25
                                                _ = v41
                                            } else {
                                                fmt.Println(v7, fi14)
                                                switch "hello" {
                                                case "":
                                                    if p1 {
                                                        v25 -= 2
                                                        fmt.Println(fi14, fi14)
                                                        v25 -= 871
                                                    } else {
                                                        if len(v15) > 0 {
                                                            v15[0] = fi12
                                                        }
                                                        var v42 bool = true
                                                        v43 := GenS7{F0: 9223372036854775807}
                                                        if ((128 < 86) || (100 != 1)) {
                                                            var v44 int32 = v26
                                                            _ = v44
                                                        }
                                                        switch -164 {
                                                        case 866:
                                                            v25 -= 257
                                                        case 42:
                                                            v45 := GenS1{F0: (v8 | (v8 * uint32(v7)))}
                                                            _ = v45
                                                        case -515:
                                                            if (false || v42) {
                                                                p0 += 4294967295
                                                                if len(v15) > 0 {
                                                                    v15[0] = 42
                                                                }
                                                                fmt.Println(fi14, v8)
                                                            }
                                                        default:
                                                            p1 = true
                                                        }
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                default:
                                                    _ = len(v15)
                                                }
                                                v46 := GenS7{F0: -1}
                                                _ = v46
                                            }
                                            _ = v26
                                        }
                                        _ = v24
                                        _ = v25
                                    } else {
                                        switch (("hello" + "hello") + "") {
                                        case "foo bar":
                                            fmt.Println(fi14, v13)
                                        }
                                    }
                                    v13 = (-(v13))
                                    p0 += 4294967295
                                    var v47 int8 = 100
                                    if p1 {
                                        fmt.Println(v13)
                                        v48 := GenS3{F0: 127}
                                        _ = v48
                                    }
                                    switch ("\n\t" + ("" + "\n\t")) {
                                    case "\n\t":
                                        _ = len(v15)
                                    case "":
                                        switch (fi14 * (fi12 >> ((2) & 63))) {
                                        case 2147483647:
                                            v49 := GenS5{F0: -1, F1: 2147483647, F2: 1}
                                            _ = v49
                                        default:
                                            var v50 int = 2147483647
                                            p2 = (p2 ^ (p2 + p2))
                                            v51 := GenS8{F0: -648, F1: 0, F2: "\n\t"}
                                            if (!(p1 || true)) {
                                                p2 = p2
                                            }
                                            switch (fi12 / ((int(v8)) | 1)) {
                                            case -991:
                                                fmt.Println(v50)
                                            case 42:
                                                if true {
                                                    var v52 bool = (0 != 1)
                                                    var v53 int8 = (v47 + v47)
                                                    fmt.Println(p2)
                                                    _ = v52
                                                    _ = v53
                                                }
                                            }
                                            var v54 float32 = 0.1
                                            fmt.Println(fi14, fi12)
                                            switch ((("\n\t" + ("hello" + (("a" + "a") + ((("\n\t" + "hello") + ("\n\t" + "a")) + "")))) + "a") + "foo bar") {
                                            case "\n\t":
                                                var v55 uint32 = v8
                                                _ = v55
                                            case "a":
                                                v7 = v7
                                            default:
                                                if (true && ((p1 && true) || (0 <= -1))) {
                                                    var v56 uint16 = 0
                                                    fmt.Println(v54, p2, v8)
                                                    _ = v56
                                                } else {
                                                    v57 := GenS9{F0: 0, F1: 570}
                                                    var v58 bool = true
                                                    _ = v57
                                                    _ = v58
                                                }
                                            }
                                            fmt.Println(v47, v47)
                                            var v59 uint = v7
                                            v60 := GenS10{F0: v47, F1: v8}
                                            var v61 int8 = (-(v47))
                                            if p1 {
                                                fmt.Println(v8)
                                                p1 = p1
                                            }
                                            var v62 bool = p1
                                            switch (fi14 | 0) {
                                            case 127:
                                                fmt.Println(p2, v54)
                                                var v63 float32 = (v54 / v54)
                                                _ = v63
                                            }
                                            _ = v50
                                            _ = v51
                                            _ = v54
                                            _ = v59
                                            _ = v60
                                            _ = v61
                                            _ = v62
                                        }
                                    default:
                                        p0 = (^(p0))
                                        fmt.Println(p0, p0)
                                    }
                                    fmt.Println(v47)
                                    var v64 int16 = (-(p2))
                                    if p1 {
                                        if p1 {
                                            if len(v15) > 0 {
                                                v15[0] = (fi12 ^ v13)
                                            }
                                        }
                                        v8 += v8
                                    } else {
                                        if (p1 || (p1 && ((^(v13)) <= fi12))) {
                                            var v65 float64 = 0.1
                                            if len(v15) >= 2 {
                                                v15 = v15[1:len(v15)]
                                            }
                                            if ((v47 / (((v47 / ((-13) | 1))) | 1)) >= (v47 ^ (v47 | v47))) {
                                                v8 += (v8 % ((v8) | 1))
                                                switch "\n\t" {
                                                case "":
                                                    p0 = (v7 / ((65535) | 1))
                                                case "foo bar":
                                                    if (255 <= 1) {
                                                        switch ("\n\t" + "a") {
                                                        case "a":
                                                            switch ("hello" + ("a" + "hello")) {
                                                            case "":
                                                                fmt.Println(v64, v47, v8)
                                                            default:
                                                                v8 += v8
                                                            }
                                                        }
                                                        if len(v15) >= 2 {
                                                            v15 = v15[1:len(v15)]
                                                        }
                                                        switch "a" {
                                                        case "\n\t":
                                                            v65 = ((v65 + v65) - 0.1)
                                                        }
                                                    }
                                                case "hello":
                                                    v66 := GenS11{F0: 1}
                                                    _ = v66
                                                default:
                                                    v67 := GenS5{F0: 697, F1: 42, F2: 0}
                                                    _ = v67
                                                }
                                            } else {
                                                switch 4294967295 {
                                                case 42:
                                                    if (v8 <= v8) {
                                                        p1 = ((^(v7)) == (p0 | (v7 & (^(p0)))))
                                                        var v68 string = "foo bar"
                                                        _ = len(v15)
                                                        if len(v15) > 0 {
                                                            v15[0] = (^(fi12))
                                                        }
                                                        _ = v68
                                                    } else {
                                                        var v69 int16 = (v64 >> (((v7 >> ((v7) & 63))) & 63))
                                                        v15 = append(v15, 42)
                                                        var v70 float32 = 0.1
                                                        v71 := GenS12{F0: fi12, F1: v7, F2: 620}
                                                        _ = v69
                                                        _ = v70
                                                        _ = v71
                                                    }
                                                case 2147483647:
                                                    if len(v15) >= 2 {
                                                        v15 = v15[1:len(v15)]
                                                    }
                                                case -1:
                                                    var v72 float32 = float32(v65)
                                                    _ = v72
                                                default:
                                                    v65 -= (v65 - (v65 / 0.1))
                                                }
                                            }
                                            if (!(((240 <= 696) && (p1 || ((("a" + ("a" + ("a" + ("\n\t" + "")))) + "\n\t") != "\n\t"))) && (42 != 623))) {
                                                fmt.Println(p2, v64, v65)
                                            } else {
                                                var v73 uint32 = v8
                                                _ = v73
                                            }
                                            _ = v65
                                        } else {
                                            var v74 float32 = float32(3.14159)
                                            _ = v74
                                        }
                                        switch "" {
                                        case "":
                                            v64 = (^(p2))
                                        case "hello":
                                            var v75 bool = (("foo bar" + "a") == "\n\t")
                                            _ = v75
                                        case "foo bar":
                                            v76 := GenS11{F0: 42}
                                            _ = v76
                                        default:
                                            var v77 int64 = 325
                                            _ = v77
                                        }
                                        if (v64 == p2) {
                                            fmt.Println(v64)
                                            v8 += 0
                                            var v78 float32 = -2.25
                                            _ = v78
                                        }
                                    }
                                    _ = v23
                                    _ = v47
                                    _ = v64
                                }
                            case "hello":
                                var v79 int32 = -343
                                _ = v79
                            default:
                                var v80 int8 = -128
                                _ = v80
                            }
                            fmt.Println(v8)
                            v8 -= 2147483647
                            if p1 {
                                var v81 int32 = 1
                                fmt.Println(p1, v81)
                                _ = v81
                            } else {
                                if (p1 && (596 < v8)) {
                                    fmt.Println(p1, v7, fi12)
                                    v82 := GenS7{F0: -653}
                                    _ = v82
                                } else {
                                    fmt.Println(v7, v8, fi14)
                                    _ = len(v15)
                                    var v83 bool = ((1e10 + 1e10) < (-2.25 * ((1e10 - 1.5) - 0.1)))
                                    var v84 string = "hello"
                                    var v85 bool = (1e10 != ((0.0 * 1.5) + 0.1))
                                    _ = v83
                                    _ = v84
                                    _ = v85
                                }
                            }
                        }
                        fmt.Println(fi14, v7, v8)
                        _ = fi14
                        _ = v15
                    }
                    _ = v13
                } else {
                    if ((!p1) && (-1 <= 2147483647)) {
                        p0 -= (p0 >> ((p0) & 63))
                        if true {
                            v86 := []int{32767, fi12}
                            _ = len(v86)
                            fmt.Println(p2)
                            _ = v86
                        }
                    }
                    var v87 int32 = 42
                    v88 := []int{fi12, fi12}
                    _ = v87
                    _ = v88
                }
                v89 := GenS13{F0: 147, F1: 255, F2: 65535}
                _ = fi12
                _ = v89
            }
            switch ("hello" + "a") {
            case "\n\t":
                if (!true) {
                    var v90 string = ("" + "hello")
                    for fi91 := 0; fi91 < 17; fi91++ {
                        v92 := GenS14{F0: 42, F1: 2147483647}
                        switch fi91 {
                        case 9223372036854775807:
                            fmt.Println(fi91, v8, v7)
                        case -9223372036854775808:
                            v93 := GenS7{F0: int64(v7)}
                            _ = v93
                        }
                        for fi94 := 0; fi94 < 14; fi94++ {
                            v8 = (v8 >> ((0) & 63))
                            _ = fi94
                        }
                        _ = fi91
                        _ = v92
                    }
                    _ = v90
                }
            case "foo bar":
                v95 := GenS14{F0: 456, F1: 42}
                _ = v95
            case "hello":
                var v96 int16 = p2
                _ = v96
            }
            v97 := GenS15{F0: p1}
            _ = v97
        }
        switch (("" + "a") + "a") {
        case "\n\t":
            if false {
                var v98 int8 = 42
                if (0 < 0) {
                    if (!(42 < -9223372036854775808)) {
                        if false {
                            var v99 int16 = 0
                            _ = v99
                        }
                        for fi100 := 0; fi100 < 11; fi100++ {
                            p2 += 1
                            v101 := GenS3{F0: 1}
                            fmt.Println(v8, v7, v98)
                            _ = fi100
                            _ = v101
                        }
                        v98 = v98
                    } else {
                        v102 := []int{-393, -1, -287, 881}
                        _ = v102
                    }
                } else {
                    v103 := GenS16{F0: v8, F1: (4294967295 >= 0)}
                    var v104 uint16 = 32767
                    var v105 uint32 = v8
                    _ = v103
                    _ = v104
                    _ = v105
                }
                _ = v98
            }
        }
        if p1 {
            if ((2147483647 < int32(v7)) || (int32(p2) != -550)) {
                v106 := GenS17{F0: 0}
                var v107 uint = (v7 - 2)
                _ = v106
                _ = v107
            } else {
                p2 = p2
                v108 := []int{-645, 1, 144, 0}
                v109 := GenS18{F0: p2}
                _ = v108
                _ = v109
            }
            fmt.Println(v8, p0, p2)
        } else {
            for fi110 := 0; fi110 < 5; fi110++ {
                var v111 uint = (p0 >> ((v7) & 63))
                _ = fi110
                _ = v111
            }
            p1 = (uint8(v8) > 1)
            var v112 uint = (v7 ^ (^(p0)))
            _ = v112
        }
    }
    switch 127 {
    case -774:
        if ((p2 << (((p0 << ((v7) & 63))) & 63)) <= p2) {
            for fi113 := 0; fi113 < 2; fi113++ {
                var v114 int = (fi113 >> ((p0) & 63))
                for fi115 := 0; fi115 < 2; fi115++ {
                    v116 := []int{fi113, fi113, v114, (^(fi113))}
                    p1 = (!(((float32((-2.25 * ((float64(1e10) - 3.14159) + (0.0 / 1.0)))) - -2.25) >= 0.1) && p1))
                    var v117 int = (^(fi115))
                    _ = fi115
                    _ = v116
                    _ = v117
                }
                v118 := []int{}
                _ = fi113
                _ = v114
                _ = v118
            }
            v7 -= p0
        }
    case 4294967295:
        if p1 {
            var v119 int64 = 963
            _ = v119
        }
    case 9223372036854775807:
        if p1 {
            if p1 {
                for fi120 := 0; fi120 < 5; fi120++ {
                    var v121 int16 = p2
                    var v122 int = fi120
                    _ = fi120
                    _ = v121
                    _ = v122
                }
                var v123 uint = p0
                var v124 int64 = -9223372036854775808
                _ = v123
                _ = v124
            }
            v125 := []int{0}
            var v126 bool = p1
            _ = v125
            _ = v126
        } else {
            p1 = (244 > 18)
            var v127 int64 = 0
            v128 := GenS3{F0: 196}
            _ = v127
            _ = v128
        }
    default:
        switch ("\n\t" + "foo bar") {
        case "foo bar":
            fmt.Println(p2, v7)
        case "a":
            for fi129 := 0; fi129 < 10; fi129++ {
                fmt.Println(p1)
                fmt.Println(v8, p1, p0)
                _ = fi129
            }
        }
    }
    p2 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p2 > -217)
}

func main() {
    var v1 int64 = 1
    var v2 float64 = (3.14159 - 1e10)
    v3 := make(map[uint32]bool)
    v4 := []uint32{}
    var v5 string = "a"
    var v6 int16 = 1000
    delete(v3, 0)
    v130 := v3[65535]
    for fi131 := 0; fi131 < 9; fi131++ {
        v132 := GenS3{F0: 255}
        _ = fi131
        _ = v132
    }
    for fi133 := 0; fi133 < 18; fi133++ {
        v134 := GenS7{F0: v1}
        if v130 {
            delete(v3, 65535)
            for fi135 := 0; fi135 < 8; fi135++ {
                v136 := GenS19{F0: (fi135 >> ((747) & 63))}
                if (2 > 4294967295) {
                    var v137 float64 = float64(0.1)
                    _ = v137
                } else {
                    v138 := GenS20{F0: -2147483648, F1: 9223372036854775807, F2: 2}
                    v130 = v130
                    var v139 int32 = 42
                    v139 = v139
                    if true {
                        fmt.Println(v130, fi133, v130)
                        if false {
                            v6 -= v6
                            v140 := v3[2147483647]
                            v141 := GenS2{F0: 869, F1: 42, F2: (^(v139))}
                            _ = v140
                            _ = v141
                        } else {
                            v142 := GenS3{F0: 0}
                            switch fi133 {
                            case -934:
                                if len(v4) > 0 {
                                    v4[0] = 4294967295
                                }
                            case -9223372036854775808:
                                if true {
                                    v143 := GenS10{F0: 111, F1: 65535}
                                    v139 = v139
                                    _ = v143
                                } else {
                                    _ = len(v4)
                                    v144 := v3[0]
                                    if (!v130) {
                                        var v145 uint = 245
                                        var v146 int = fi133
                                        _ = v145
                                        _ = v146
                                    } else {
                                        switch -2147483648 {
                                        case 42:
                                            var v147 int32 = v139
                                            _ = v147
                                        case 32767:
                                            var v148 int16 = 0
                                            _ = v148
                                        case -241:
                                            v5 = "a"
                                        }
                                        var v149 float32 = (((((float32(-2.25) / float32(float64(1e10))) / (3.14159 - 1e10)) / 3.14159) + ((1e10 + 3.14159) / 1.5)) - (0.0 + (-2.25 / (float32(v2) - 1.5))))
                                        _ = v149
                                    }
                                    _ = v144
                                }
                            case -451:
                                switch (v5 + ((v5 + v5) + (v5 + "\n\t"))) {
                                case "foo bar":
                                    if (2147483647 < v139) {
                                        v5 = "hello"
                                        var v150 float32 = 1e10
                                        _ = v150
                                    }
                                default:
                                    var v151 uint8 = 0
                                    _ = v151
                                }
                            }
                            _ = v142
                        }
                        v139 = (v139 - v139)
                    } else {
                        fmt.Println(v5, fi133, v2)
                    }
                    if false {
                        v152 := v3[2147483647]
                        v153 := GenS17{F0: 1}
                        _ = v152
                        _ = v153
                    } else {
                        if (1e10 > 1e10) {
                            v139 = (v139 >> ((65535) & 63))
                            if v130 {
                                if (false || true) {
                                    v130 = ((1 < 100) || v130)
                                    if true {
                                        var v154 int8 = 100
                                        v6 += (v6 + v6)
                                        _ = v154
                                    } else {
                                        v155 := GenS21{F0: 311, F1: -77, F2: 4294967295}
                                        _ = v155
                                    }
                                    fmt.Println(v1, v6, v130)
                                    if (!false) {
                                        v139 = (v139 / ((v139) | 1))
                                        if (int8(v139) > 105) {
                                            fmt.Println(v5, v130, fi135)
                                        }
                                    }
                                } else {
                                    if len(v4) > 0 {
                                        v4[0] = 2147483647
                                    }
                                    fmt.Println(fi133, v130)
                                }
                                v139 -= 14
                                if (((894 > 244) || (v130 || false)) || (float64(1.5) != v2)) {
                                    var v156 int16 = (v6 / (((v6 & -1)) | 1))
                                    var v157 uint32 = 448
                                    _ = v156
                                    _ = v157
                                }
                            }
                        } else {
                            if false {
                                var v158 int = 2147483647
                                v3[2] = (((^(v6)) < v6) || ((v130 && (!v130)) || true))
                                _ = v158
                            }
                            v130 = (!false)
                            v130 = false
                        }
                        v6 = -1
                        fmt.Println(fi135, fi133, v139)
                    }
                    v159 := GenS17{F0: 1}
                    _ = v138
                    _ = v139
                    _ = v159
                }
                _ = fi135
                _ = v136
            }
            v1 = v1
        } else {
            fmt.Println(v2, v2)
            _ = len(v3)
        }
        var v160 int64 = v1
        _ = fi133
        _ = v134
        _ = v160
    }
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    v161 := GenS20{F0: 2147483647, F1: v1, F2: 39}
    v162 := GenS11{F0: (v1 * (v1 + (v1 >> ((0) & 63))))}
    if (9223372036854775807 >= 32767) {
        _ = len(v4)
        v5 = "a"
        v2 -= 0.0
    } else {
        for fi163 := 0; fi163 < 13; fi163++ {
            v3[1] = false
            fmt.Println(v5, v1, v5)
            _ = fi163
        }
        _ = len(v4)
        for fi164 := 0; fi164 < 18; fi164++ {
            v6 = v6
            var v165 int32 = -2147483648
            fmt.Println(v1, v130)
            _ = fi164
            _ = v165
        }
    }
    _ = len(v4)
    v130 = (v130 || false)
    fmt.Println(v5)
    if false {
        switch 185 {
        case 1:
            if (!((42 < 1) || v130)) {
                fmt.Println(v1, v6)
            } else {
                switch 0 {
                case 4294967295:
                    var v166 float32 = (0.0 - ((3.14159 * 3.14159) + 1e10))
                    _ = v166
                case 127:
                    var v167 string = v5
                    _ = v167
                default:
                    v168 := GenS22{F0: 0}
                    _ = v168
                }
                var v169 uint = 840
                _ = v169
            }
        case -210:
            for fi170 := 0; fi170 < 19; fi170++ {
                var v171 int16 = (v6 << ((886) & 63))
                v171 += (v6 ^ (v6 << ((350) & 63)))
                var v172 int = (fi170 >> ((630) & 63))
                _ = fi170
                _ = v171
                _ = v172
            }
        case 32767:
            v6 += v6
        default:
            v6 += v6
        }
    }
    v173 := GenS23{F0: v6, F1: ((v5 + "\n\t") + ("hello" + v5)), F2: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("v130:", v130)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v130
    _ = v161
    _ = v162
    _ = v173
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
