package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int8; F1 string; F2 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int8; F1 int64 }
type GenS8 struct { F0 int16; F1 int16; F2 uint8 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint32; F1 int32; F2 int }
type GenS11 struct { F0 string; F1 bool }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint32; F1 uint32; F2 int64 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 int64; F1 int32 }
type GenS16 struct { F0 uint8; F1 int8; F2 uint }
type GenS17 struct { F0 int16; F1 uint32; F2 int }
type GenS18 struct { F0 int }
type GenS19 struct { F0 uint32; F1 bool }
type GenS20 struct { F0 uint; F1 int; F2 string }
type GenS21 struct { F0 int8; F1 string; F2 uint32 }
type GenS22 struct { F0 int32; F1 string }
type GenS23 struct { F0 int }
type GenS24 struct { F0 uint16 }
type GenS25 struct { F0 uint32; F1 uint8; F2 uint32 }
type GenS26 struct { F0 uint16; F1 string; F2 bool }
func genFn0() uint16 {
    v4 := GenS1{F0: 72}
    var v5 int8 = 8
    v6 := GenS2{F0: 921, F1: -1}
    for fi7 := 0; fi7 < 5; fi7++ {
        fmt.Println(v5, fi7)
        _ = fi7
    }
    v8 := GenS3{F0: v5, F1: ("hello" + ("hello" + "hello")), F2: 1000}
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    return 2
}
func genFn1(p0 uint16, p1 int16) uint {
    var v9 uint = 4294967295
    p0 += (p0 * (^(p0)))
    var v10 uint16 = p0
    p1 = p1
    switch 316 {
    case 127:
        if false {
            if (((p1 << (((v9 >> (((v9 * 1)) & 63))) & 63)) <= 0) || (!false)) {
                switch 128 {
                case 9223372036854775807:
                    if false {
                        var v11 float64 = 3.14159
                        _ = v11
                    }
                case 128:
                    v9 += v9
                case -839:
                    if false {
                        v12 := GenS4{F0: 419}
                        if (127 < 1) {
                            p1 += (^(p1))
                            v9 -= (^(v9))
                        }
                        if true {
                            var v13 uint32 = 0
                            for fi14 := 0; fi14 < 14; fi14++ {
                                switch (fi14 ^ (fi14 & (fi14 + (fi14 + -1)))) {
                                case 4294967295:
                                    v10 = (^(v10))
                                case -1:
                                    var v15 uint64 = 18446744073709551615
                                    _ = v15
                                default:
                                    for fi16 := 0; fi16 < 19; fi16++ {
                                        var v17 int32 = 42
                                        if ((v13 / ((uint32(v9)) | 1)) == v13) {
                                            v18 := GenS1{F0: 128}
                                            if true {
                                                v19 := GenS4{F0: v17}
                                                v20 := GenS5{F0: 32, F1: 261}
                                                _ = v19
                                                _ = v20
                                            } else {
                                                if (2147483647 > (v17 % (((^(v17))) | 1))) {
                                                    v10 -= (p0 + uint16(v9))
                                                    var v21 int64 = -243
                                                    v22 := GenS2{F0: (v13 + v13), F1: -1}
                                                    _ = v21
                                                    _ = v22
                                                } else {
                                                    switch "" {
                                                    case "":
                                                        var v23 string = "a"
                                                        _ = v23
                                                    case "foo bar":
                                                        if false {
                                                            v24 := GenS4{F0: v17}
                                                            if (!true) {
                                                                v25 := GenS1{F0: 0}
                                                                if (!false) {
                                                                    v17 -= v17
                                                                    var v26 int8 = int8(v13)
                                                                    v17 = v17
                                                                    _ = v26
                                                                }
                                                                v17 -= (v17 & (v17 % ((v17) | 1)))
                                                                if true {
                                                                    if (false && ((p1 * (p1 | p1)) != p1)) {
                                                                        switch ("foo bar" + "hello") {
                                                                        case "hello":
                                                                            v27 := []int{(fi14 & 2147483647)}
                                                                            _ = v27
                                                                        case "a":
                                                                            v10 -= 0
                                                                        default:
                                                                            switch (fi16 / ((fi14) | 1)) {
                                                                            case 896:
                                                                                switch (^(fi14)) {
                                                                                case -476:
                                                                                    v28 := GenS6{F0: (p1 ^ -492)}
                                                                                    _ = v28
                                                                                case 32767:
                                                                                    v29 := []int{42, (fi14 << ((uint(p0)) & 63))}
                                                                                    _ = v29
                                                                                }
                                                                            case 4294967295:
                                                                                v30 := []int{}
                                                                                _ = v30
                                                                            case -9223372036854775808:
                                                                                v31 := GenS4{F0: v17}
                                                                                _ = v31
                                                                            }
                                                                        }
                                                                    } else {
                                                                        switch "hello" {
                                                                        case "foo bar":
                                                                            switch ("foo bar" + "hello") {
                                                                            case "hello":
                                                                                fmt.Println(v17)
                                                                            case "foo bar":
                                                                                v17 = v17
                                                                            case "a":
                                                                                p1 += (p1 & p1)
                                                                            default:
                                                                                v32 := GenS7{F0: 10, F1: -9223372036854775808}
                                                                                _ = v32
                                                                            }
                                                                        case "\n\t":
                                                                            fmt.Println(fi16)
                                                                        case "hello":
                                                                            switch ("" + ("hello" + "a")) {
                                                                            case "hello":
                                                                                switch "hello" {
                                                                                case "\n\t":
                                                                                    var v33 uint16 = genFn0()
                                                                                    _ = v33
                                                                                case "":
                                                                                    v34 := []int{(fi16 - int(v9)), fi16, (^(fi14)), 1}
                                                                                    _ = v34
                                                                                case "hello":
                                                                                    switch "" {
                                                                                    case "a":
                                                                                        var v35 string = "foo bar"
                                                                                        _ = v35
                                                                                    }
                                                                                default:
                                                                                    fmt.Println(p0, v10, fi16)
                                                                                }
                                                                            case "\n\t":
                                                                                var v36 uint32 = (v13 + v13)
                                                                                _ = v36
                                                                            case "a":
                                                                                var v37 string = ("a" + ((("foo bar" + "\n\t") + "foo bar") + "hello"))
                                                                                _ = v37
                                                                            }
                                                                        default:
                                                                            var v38 int8 = 1
                                                                            _ = v38
                                                                        }
                                                                        v39 := []int{(fi16 / (((fi16 * fi16)) | 1)), fi14, (^(fi14))}
                                                                        _ = v39
                                                                    }
                                                                    var v40 int64 = 399
                                                                    _ = v40
                                                                }
                                                                var v41 float32 = -2.25
                                                                _ = v25
                                                                _ = v41
                                                            } else {
                                                                v10 -= genFn0()
                                                            }
                                                            _ = v24
                                                        } else {
                                                            v17 = (v17 - (v17 - v17))
                                                        }
                                                    default:
                                                        var v42 int64 = -1
                                                        _ = v42
                                                    }
                                                    if true {
                                                        var v43 int = fi14
                                                        _ = v43
                                                    }
                                                    var v44 bool = (!(128 < 0))
                                                    _ = v44
                                                }
                                            }
                                            v45 := []int{}
                                            v9 -= v9
                                            var v46 float64 = (1.5 + 1e10)
                                            switch "\n\t" {
                                            case "\n\t":
                                                if (p1 < (p1 / ((p1) | 1))) {
                                                    v47 := GenS2{F0: (v13 - (v13 >> (((v9 + (v9 % ((v9) | 1)))) & 63))), F1: v17}
                                                    var v48 uint8 = 255
                                                    if false {
                                                        if (!true) {
                                                            if false {
                                                                var v49 uint64 = 667
                                                                _ = v49
                                                            } else {
                                                                _ = len(v45)
                                                                var v50 int32 = (v17 & (v17 + v17))
                                                                _ = v50
                                                            }
                                                        }
                                                        v10 = (v10 * v10)
                                                        if true {
                                                            p0 -= p0
                                                            var v51 uint8 = (v48 + (v48 | (v48 - 2)))
                                                            _ = v51
                                                        } else {
                                                            switch -865 {
                                                            case -9223372036854775808:
                                                                var v52 uint = (v9 - 641)
                                                                _ = v52
                                                            case 0:
                                                                var v53 int = fi16
                                                                _ = v53
                                                            case 209:
                                                                if len(v45) >= 2 {
                                                                    v45 = v45[1:len(v45)]
                                                                }
                                                            }
                                                            v54 := GenS4{F0: v17}
                                                            if false {
                                                                v10 = 1
                                                                if false {
                                                                    v17 += (v17 << ((255) & 63))
                                                                } else {
                                                                    var v55 uint32 = (v13 & (v13 | 101))
                                                                    var v56 int64 = 742
                                                                    if (int8(p0) >= -1) {
                                                                        p0 -= (^(p0))
                                                                        switch fi14 {
                                                                        case 42:
                                                                            var v57 int8 = 1
                                                                            _ = v57
                                                                        case 127:
                                                                            if false {
                                                                                v58 := GenS5{F0: (^(v48)), F1: (v10 % (((p0 & v10)) | 1))}
                                                                                p1 = 32767
                                                                                var v59 uint8 = 0
                                                                                _ = v58
                                                                                _ = v59
                                                                            } else {
                                                                                fmt.Println(p1, p1)
                                                                                p1 = 1
                                                                            }
                                                                            switch -1 {
                                                                            case 128:
                                                                                var v60 uint = (v9 + v9)
                                                                                _ = v60
                                                                            }
                                                                        default:
                                                                            var v61 float32 = 0.1
                                                                            _ = v61
                                                                        }
                                                                        if (!(((v13 << (((v9 * (^(v9)))) & 63)) > 195) || (!false))) {
                                                                            p1 -= p1
                                                                            if ((((float32(-2.25) - 3.14159) / 1e10) - -2.25) < -2.25) {
                                                                                fmt.Println(fi16)
                                                                                if len(v45) > 0 {
                                                                                    v45[0] = (fi14 << ((v9) & 63))
                                                                                }
                                                                                switch "\n\t" {
                                                                                case "foo bar":
                                                                                    p0 = p0
                                                                                case "":
                                                                                    v46 -= v46
                                                                                case "a":
                                                                                    _ = len(v45)
                                                                                default:
                                                                                    v62 := GenS8{F0: (p1 ^ -1), F1: p1, F2: uint8(p0)}
                                                                                    _ = v62
                                                                                }
                                                                            }
                                                                            p1 -= 42
                                                                        }
                                                                        if true {
                                                                            v63 := GenS9{F0: fi16}
                                                                            _ = v63
                                                                        }
                                                                        v45 = append(v45, (fi14 % (((fi14 | fi14)) | 1)))
                                                                        switch "hello" {
                                                                        case "\n\t":
                                                                            if (uint64(v48) == 42) {
                                                                                v64 := GenS10{F0: v13, F1: v17, F2: 128}
                                                                                switch "" {
                                                                                case "":
                                                                                    fmt.Println(v56)
                                                                                case "a":
                                                                                    var v65 int = (fi16 >> (((v9 - (v9 + (v9 << ((v9) & 63))))) & 63))
                                                                                    _ = v65
                                                                                case "hello":
                                                                                    var v66 int16 = p1
                                                                                    _ = v66
                                                                                }
                                                                                _ = v64
                                                                            }
                                                                        case "":
                                                                            v9 = v9
                                                                        }
                                                                    }
                                                                    _ = v55
                                                                    _ = v56
                                                                }
                                                            }
                                                            _ = v54
                                                        }
                                                    }
                                                    _ = v47
                                                    _ = v48
                                                } else {
                                                    if (false || (!(("" + "\n\t") <= ("hello" + "\n\t")))) {
                                                        v67 := GenS11{F0: "hello", F1: true}
                                                        _ = v67
                                                    } else {
                                                        v9 = 238
                                                    }
                                                    var v68 string = "\n\t"
                                                    _ = len(v45)
                                                    _ = v68
                                                }
                                                if len(v45) > 0 {
                                                    v45[0] = fi14
                                                }
                                                v45 = append(v45, 32767)
                                                fmt.Println(p0)
                                            case "":
                                                v69 := GenS12{F0: v10}
                                                if true {
                                                    v9 -= 255
                                                    fmt.Println(v10, v46)
                                                    v70 := GenS7{F0: 24, F1: 9223372036854775807}
                                                    _ = v70
                                                }
                                                v17 -= v17
                                                _ = v69
                                            }
                                            p0 = p0
                                            p0 = (v10 % ((p0) | 1))
                                            _ = v18
                                            _ = v45
                                            _ = v46
                                        } else {
                                            v71 := GenS13{F0: (v13 * (v13 | v13)), F1: (^(v13)), F2: 1}
                                            v72 := GenS14{F0: 1}
                                            switch (-(fi14)) {
                                            case 192:
                                                v73 := []int{}
                                                _ = v73
                                            }
                                            _ = v71
                                            _ = v72
                                        }
                                        _ = fi16
                                        _ = v17
                                    }
                                }
                                _ = fi14
                            }
                            if (!(!(true || false))) {
                                for fi74 := 0; fi74 < 13; fi74++ {
                                    v75 := GenS15{F0: 9223372036854775807, F1: -2147483648}
                                    p1 = p1
                                    for fi76 := 0; fi76 < 19; fi76++ {
                                        switch "\n\t" {
                                        case "a":
                                            p1 += p1
                                        }
                                        _ = fi76
                                    }
                                    _ = fi74
                                    _ = v75
                                }
                                v77 := GenS11{F0: "foo bar", F1: (!(false || true))}
                                _ = v77
                            }
                            _ = v13
                        }
                        _ = v12
                    } else {
                        v78 := []int{-271, 32768}
                        if true {
                            v9 = v9
                            for fi79 := 0; fi79 < 19; fi79++ {
                                var v80 string = "\n\t"
                                var v81 uint8 = 127
                                for fi82 := 0; fi82 < 8; fi82++ {
                                    if len(v78) >= 2 {
                                        v78 = v78[1:len(v78)]
                                    }
                                    _ = fi82
                                }
                                _ = fi79
                                _ = v80
                                _ = v81
                            }
                        } else {
                            for fi83 := 0; fi83 < 19; fi83++ {
                                for fi84 := 0; fi84 < 5; fi84++ {
                                    switch (fi83 - 0) {
                                    case 42:
                                        if (0 > 1) {
                                            var v85 int32 = 700
                                            if true {
                                                var v86 uint32 = 31
                                                _ = v86
                                            }
                                            if false {
                                                v85 -= v85
                                                fmt.Println(p1)
                                                var v87 int16 = 1000
                                                _ = v87
                                            } else {
                                                var v88 int16 = p1
                                                v89 := GenS16{F0: 35, F1: 1, F2: v9}
                                                v78 = append(v78, (fi84 * (fi84 * fi84)))
                                                _ = v88
                                                _ = v89
                                            }
                                            _ = v85
                                        }
                                    case 4294967295:
                                        fmt.Println(v10, fi83, p0)
                                    default:
                                        v90 := GenS17{F0: 32767, F1: 1, F2: fi83}
                                        _ = v90
                                    }
                                    var v91 uint64 = 42
                                    _ = fi84
                                    _ = v91
                                }
                                _ = fi83
                            }
                            if false {
                                switch 2147483647 {
                                case 2147483647:
                                    for fi92 := 0; fi92 < 9; fi92++ {
                                        v93 := GenS5{F0: 42, F1: (p0 & genFn0())}
                                        for fi94 := 0; fi94 < 4; fi94++ {
                                            switch fi94 {
                                            case 127:
                                                v9 += 0
                                            case 1:
                                                if len(v78) >= 2 {
                                                    v78 = v78[1:len(v78)]
                                                }
                                            }
                                            _ = fi94
                                        }
                                        _ = fi92
                                        _ = v93
                                    }
                                }
                                switch -346 {
                                case 2147483647:
                                    for fi95 := 0; fi95 < 8; fi95++ {
                                        v96 := GenS12{F0: v10}
                                        for fi97 := 0; fi97 < 5; fi97++ {
                                            switch "hello" {
                                            case "\n\t":
                                                fmt.Println(v10, p0, v9)
                                            default:
                                                _ = len(v78)
                                            }
                                            v98 := GenS16{F0: 1, F1: -128, F2: v9}
                                            _ = fi97
                                            _ = v98
                                        }
                                        _ = fi95
                                        _ = v96
                                    }
                                default:
                                    if len(v78) >= 2 {
                                        v78 = v78[1:len(v78)]
                                    }
                                }
                            } else {
                                var v99 uint16 = p0
                                _ = v99
                            }
                            v100 := GenS16{F0: 127, F1: int8(v9), F2: v9}
                            _ = v100
                        }
                        _ = v78
                    }
                }
                v101 := GenS18{F0: 32767}
                v102 := []int{896}
                _ = v101
                _ = v102
            }
            v103 := GenS19{F0: 65535, F1: true}
            _ = v103
        }
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 1
}

func main() {
    var v1 int8 = 1
    v2 := []uint32{2147483647, 434}
    var v3 bool = true
    if ((("a" + (("" + "") + "a")) + "\n\t") <= ("foo bar" + ("hello" + ("hello" + "hello")))) {
        if v3 {
            var v104 uint32 = 2147483648
            _ = v104
        }
        var v105 int = -732
        fmt.Println(v105)
        _ = v105
    } else {
        v1 = (v1 & -128)
    }
    if false {
        _ = len(v2)
        var v106 int8 = -32
        var v107 int16 = 42
        _ = v106
        _ = v107
    }
    v108 := GenS5{F0: 255, F1: 255}
    v2 = append(v2, 42)
    _ = len(v2)
    var v109 bool = v3
    if v109 {
        if false {
            if true {
                var v110 string = ("hello" + "a")
                var v111 int = -36
                if (!(((0.1 * (0.1 + 0.1)) / 1.5) == 0.1)) {
                    v112 := GenS14{F0: 128}
                    v113 := GenS20{F0: 255, F1: (v111 + v111), F2: v110}
                    var v114 float32 = 1.5
                    _ = v112
                    _ = v113
                    _ = v114
                }
                _ = v110
                _ = v111
            }
            _ = len(v2)
            _ = len(v2)
        }
        v115 := GenS21{F0: (v1 | (v1 / ((127) | 1))), F1: ("a" + "hello"), F2: 2147483647}
        _ = v115
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v116 := GenS6{F0: 51}
        var v117 uint = 4294967295
        _ = v116
        _ = v117
    }
    v3 = (v109 || ((-2.25 / (0.0 - (((0.1 + 3.14159) - 0.0) + 1e10))) >= -2.25))
    v2 = append(v2, 2)
    v118 := GenS22{F0: 42, F1: "foo bar"}
    v119 := GenS6{F0: -32768}
    for fi120 := 0; fi120 < 17; fi120++ {
        var v121 uint64 = 42
        v122 := GenS23{F0: fi120}
        _ = fi120
        _ = v121
        _ = v122
    }
    for fi123 := 0; fi123 < 4; fi123++ {
        for fi124 := 0; fi124 < 20; fi124++ {
            v125 := GenS24{F0: 2}
            var v126 uint64 = 42
            var v127 int = fi124
            _ = fi124
            _ = v125
            _ = v126
            _ = v127
        }
        if ("" == "\n\t") {
            var v128 int32 = 42
            fmt.Println(v128, v128)
            if v109 {
                if false {
                    switch "a" {
                    case "":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "hello":
                        if v3 {
                            switch (("foo bar" + "a") + "\n\t") {
                            case "":
                                for fi129 := 0; fi129 < 5; fi129++ {
                                    fmt.Println(v1, fi129)
                                    _ = fi129
                                }
                            case "a":
                                _ = len(v2)
                            case "foo bar":
                                v130 := GenS24{F0: genFn0()}
                                _ = v130
                            default:
                                var v131 uint = 65535
                                _ = v131
                            }
                            var v132 int16 = 32767
                            switch "hello" {
                            case "a":
                                for fi133 := 0; fi133 < 6; fi133++ {
                                    if v109 {
                                        if len(v2) > 0 {
                                            v2[0] = 247
                                        }
                                    } else {
                                        if ((1e10 + (0.0 * 1e10)) <= float32(float64(1.5))) {
                                            var v134 uint64 = 1
                                            _ = v134
                                        } else {
                                            var v135 bool = true
                                            v132 = (v132 ^ 1000)
                                            var v136 int = 362
                                            _ = v135
                                            _ = v136
                                        }
                                    }
                                    var v137 int8 = 100
                                    _ = fi133
                                    _ = v137
                                }
                            case "hello":
                                var v138 int = (fi123 ^ (fi123 << ((54) & 63)))
                                _ = v138
                            }
                            _ = v132
                        } else {
                            if (fi123 <= fi123) {
                                if (((-2.25 / (0.0 - (float32(0.0) + 1.5))) + -2.25) == 3.14159) {
                                    v3 = v109
                                    var v139 uint16 = 262
                                    _ = v139
                                } else {
                                    for fi140 := 0; fi140 < 19; fi140++ {
                                        if (42 > 625) {
                                            var v141 uint32 = 2
                                            v142 := GenS23{F0: fi123}
                                            _ = v141
                                            _ = v142
                                        }
                                        var v143 int = 164
                                        _ = fi140
                                        _ = v143
                                    }
                                    if (((-2.25 <= 1e10) && (fi123 < -541)) && true) {
                                        v144 := GenS13{F0: 305, F1: 1, F2: 9223372036854775807}
                                        _ = v144
                                    } else {
                                        v128 = v128
                                        v128 -= v128
                                        var v145 uint8 = 0
                                        _ = v145
                                    }
                                    var v146 int8 = (v1 / ((v1) | 1))
                                    _ = v146
                                }
                                switch (fi123 | fi123) {
                                case 42:
                                    switch fi123 {
                                    case -1:
                                        var v147 bool = ((v1 | (v1 << ((255) & 63))) > v1)
                                        _ = v147
                                    case 0:
                                        v109 = (((true && false) || true) && (!v3))
                                    case -9223372036854775808:
                                        var v148 uint32 = 908
                                        _ = v148
                                    }
                                }
                                v109 = v3
                            } else {
                                v109 = false
                                fmt.Println(v3)
                            }
                            var v149 float32 = -2.25
                            if v109 {
                                if (!false) {
                                    v149 += 3.14159
                                }
                                v109 = (!true)
                            } else {
                                v150 := GenS25{F0: 2147483648, F1: 127, F2: 1}
                                v151 := GenS17{F0: 32767, F1: 1, F2: fi123}
                                _ = v150
                                _ = v151
                            }
                            _ = v149
                        }
                    default:
                        if (!(v128 >= 42)) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            for fi152 := 0; fi152 < 19; fi152++ {
                                v1 = v1
                                if (1.5 != 0.1) {
                                    v1 -= 100
                                    var v153 uint32 = 9
                                    _ = v153
                                }
                                v1 += v1
                                _ = fi152
                            }
                        }
                    }
                    for fi154 := 0; fi154 < 13; fi154++ {
                        v3 = v3
                        _ = fi154
                    }
                    fmt.Println(v109, v109)
                } else {
                    v155 := GenS24{F0: 1}
                    v156 := GenS26{F0: 42, F1: (("" + "foo bar") + "a"), F2: ((v1 - (v1 / ((v1) | 1))) != (v1 << ((0) & 63)))}
                    v109 = (65535 < 2147483647)
                    _ = v155
                    _ = v156
                }
            }
            _ = v128
        } else {
            fmt.Println(v1, fi123)
        }
        fmt.Println(v109, v1, v3)
        _ = fi123
    }
    fmt.Println("v1:", v1)
    fmt.Println("v109:", v109)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v108
    _ = v109
    _ = v118
    _ = v119
    _ = v2
    _ = v3
}
