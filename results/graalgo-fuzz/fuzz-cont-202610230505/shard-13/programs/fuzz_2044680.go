package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 uint; F1 int; F2 int32 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int32; F1 uint16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 bool; F1 int64 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint; F1 uint64 }
type GenS12 struct { F0 bool; F1 bool; F2 uint32 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 uint; F1 int }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 int64; F1 int8; F2 uint }
type GenS17 struct { F0 uint8 }
type GenS18 struct { F0 int64; F1 uint; F2 uint }
type GenS19 struct { F0 uint64; F1 int64 }
type GenS20 struct { F0 bool; F1 int16; F2 int16 }
type GenS21 struct { F0 uint8; F1 int32; F2 int64 }
type GenS22 struct { F0 uint16; F1 uint; F2 int32 }
type GenS23 struct { F0 uint32; F1 string; F2 bool }
type GenS24 struct { F0 int; F1 int16 }
type GenS25 struct { F0 uint32; F1 int16; F2 uint64 }
type GenS26 struct { F0 uint16; F1 uint32 }
type GenS27 struct { F0 string }
func genFn0(p0 int64) bool {
    if (((!false) && true) && true) {
        var v7 int = -286
        fmt.Println(v7, p0)
        v8 := GenS1{F0: 1}
        _ = v7
        _ = v8
    } else {
        var v9 uint32 = 0
        v9 = 4294967295
        for fi10 := 0; fi10 < 2; fi10++ {
            for fi11 := 0; fi11 < 9; fi11++ {
                var v12 bool = false
                if (v12 || (!(("foo bar" + ((((("a" + "hello") + "hello") + "foo bar") + "a") + ("\n\t" + "a"))) <= "hello"))) {
                    v13 := GenS2{F0: 18446744073709551615}
                    v14 := []int{}
                    if len(v14) > 0 {
                        v14[0] = fi10
                    }
                    _ = v13
                    _ = v14
                } else {
                    if (false && v12) {
                        fmt.Println(v12, v12)
                    } else {
                        var v15 uint8 = 128
                        switch "foo bar" {
                        case "hello":
                            v16 := []int{-42, fi11, fi11}
                            _ = v16
                        }
                        v15 += 22
                        _ = v15
                    }
                    v17 := []int{fi11, 42, fi11, fi10}
                    switch 32768 {
                    case 834:
                        var v18 int = int(p0)
                        _ = v18
                    case -2147483648:
                        p0 -= (p0 - p0)
                    }
                    switch fi11 {
                    case 32768:
                        _ = len(v17)
                        var v19 uint8 = 63
                        _ = v19
                    }
                    switch "\n\t" {
                    case "hello":
                        switch "hello" {
                        case "a":
                            switch "" {
                            case "\n\t":
                                switch "" {
                                case "":
                                    if (959 != 597) {
                                        p0 = p0
                                    } else {
                                        var v20 int64 = 0
                                        _ = v20
                                    }
                                case "hello":
                                    var v21 float64 = float64((-2.25 + 0.1))
                                    _ = v21
                                case "a":
                                    v9 = (v9 / ((v9) | 1))
                                }
                            }
                        }
                        v9 = v9
                        var v22 float32 = ((0.0 / (((1e10 * (1.5 * -2.25)) * 0.1) - 1e10)) / 1e10)
                        _ = v22
                    }
                    _ = v17
                }
                p0 -= p0
                var v23 uint32 = v9
                _ = fi11
                _ = v12
                _ = v23
            }
            if ("" != "\n\t") {
                var v24 bool = (false || true)
                var v25 int64 = -337
                if (115 == 42) {
                    switch "foo bar" {
                    case "\n\t":
                        fmt.Println(p0, v9, v9)
                    default:
                        v26 := GenS2{F0: 255}
                        _ = v26
                    }
                } else {
                    for fi27 := 0; fi27 < 13; fi27++ {
                        v28 := GenS3{F0: fi10, F1: 1}
                        if false {
                            v29 := []int{fi10, fi27}
                            v9 += v9
                            _ = v29
                        }
                        if (4294967295 == 895) {
                            v30 := GenS2{F0: 255}
                            var v31 float32 = float32((0.1 / 1e10))
                            var v32 int64 = 807
                            if (((0 == 42) || true) && (v31 > v31)) {
                                var v33 int32 = 2147483647
                                v34 := []int{fi27, fi10}
                                v33 = v33
                                switch fi10 {
                                case 567:
                                    _ = len(v34)
                                case 0:
                                    v24 = (false || (v24 && (!(!true))))
                                case -1:
                                    var v35 float64 = 0.0
                                    _ = v35
                                default:
                                    if (!(!(0 <= 494))) {
                                        fmt.Println(v31, v32)
                                        fmt.Println(fi27, p0)
                                        v36 := GenS4{F0: v33, F1: (fi27 | fi27)}
                                        switch (-(fi27)) {
                                        case 42:
                                            if true {
                                                var v37 uint16 = 1
                                                _ = v37
                                            } else {
                                                fmt.Println(v33, fi27, fi27)
                                            }
                                        default:
                                            _ = len(v34)
                                        }
                                        p0 -= (v32 & (v25 << ((720) & 63)))
                                        var v38 int64 = 42
                                        _ = v36
                                        _ = v38
                                    }
                                }
                                v34 = append(v34, (fi10 - 32767))
                                _ = v33
                                _ = v34
                            }
                            _ = v30
                            _ = v31
                            _ = v32
                        } else {
                            fmt.Println(fi10, v25, fi10)
                            switch (fi10 * 4294967295) {
                            case -493:
                                var v39 string = "hello"
                                _ = v39
                            }
                        }
                        var v40 uint16 = 0
                        v41 := GenS5{F0: 4294967295, F1: fi10, F2: 1}
                        _ = fi27
                        _ = v28
                        _ = v40
                        _ = v41
                    }
                    if v24 {
                        v25 -= p0
                    } else {
                        v24 = (725 >= 2)
                        p0 += 1
                    }
                    var v42 uint8 = 36
                    _ = v42
                }
                _ = v24
                _ = v25
            } else {
                v43 := []int{(fi10 & (fi10 >> ((274) & 63))), -994, (fi10 * fi10)}
                p0 += (p0 & (^(p0)))
                if len(v43) >= 2 {
                    v43 = v43[1:len(v43)]
                }
                _ = v43
            }
            _ = fi10
        }
        _ = v9
    }
    v44 := GenS4{F0: int32(p0), F1: 32768}
    v45 := []int{290}
    _ = p0
    _ = v44
    _ = v45
    return (0 >= 629)
}
func genFn1(p0 uint16) uint32 {
    fmt.Println(p0)
    p0 -= (p0 << ((42) & 63))
    v46 := []int{-765}
    for fi47 := 0; fi47 < 20; fi47++ {
        for fi48 := 0; fi48 < 5; fi48++ {
            v46 = append(v46, fi47)
            _ = fi48
        }
        if (100 >= 0) {
            p0 = (^(p0))
            _ = len(v46)
            for fi49 := 0; fi49 < 1; fi49++ {
                v46 = append(v46, 4294967295)
                switch fi49 {
                case -9223372036854775808:
                    var v50 uint16 = 395
                    _ = v50
                case 127:
                    fmt.Println(fi47, fi47, fi47)
                    v51 := GenS6{F0: 0}
                    v46 = append(v46, 9223372036854775807)
                    var v52 uint = 42
                    if true {
                        fmt.Println(fi49)
                    }
                    if true {
                        fmt.Println(fi49, fi49, fi47)
                    }
                    _ = v51
                    _ = v52
                case -2147483648:
                    var v53 float64 = 1.5
                    _ = v53
                default:
                    if true {
                        fmt.Println(fi47)
                        p0 = (p0 * p0)
                    }
                }
                var v54 int8 = 100
                if true {
                    var v55 uint16 = (p0 >> ((0) & 63))
                    var v56 int32 = 1
                    _ = v55
                    _ = v56
                }
                v57 := GenS7{F0: 1, F1: p0}
                _ = fi49
                _ = v54
                _ = v57
            }
        }
        _ = fi47
    }
    _ = p0
    _ = v46
    return 716
}

func main() {
    v1 := make(map[int]int)
    var v2 float64 = float64(3.14159)
    var v3 float64 = (float64(-2.25) * float64((0.0 - (1.5 * float32(v2)))))
    var v4 string = "\n\t"
    var v5 int16 = 1000
    var v6 float64 = (v2 * -2.25)
    for fi58 := 0; fi58 < 3; fi58++ {
        switch ((v4 + v4) + "foo bar") {
        case "foo bar":
            if false {
                if (!true) {
                    var v59 bool = false
                    v2 = (v6 + float64(0.1))
                    if (v59 || ((65535 <= 573) || v59)) {
                        v60 := GenS3{F0: 128, F1: 103}
                        for fi61 := 0; fi61 < 11; fi61++ {
                            v62 := []int{}
                            _ = fi61
                            _ = v62
                        }
                        _ = v60
                    } else {
                        var v63 int16 = v5
                        switch 127 {
                        case -356:
                            switch (fi58 >> ((891) & 63)) {
                            case -9223372036854775808:
                                switch (fi58 - fi58) {
                                case 32768:
                                    var v64 uint64 = 0
                                    _ = v64
                                case -963:
                                    for fi65 := 0; fi65 < 1; fi65++ {
                                        var v66 uint8 = 255
                                        v67 := []int{int(v66)}
                                        fmt.Println(v4)
                                        if (0 <= 244) {
                                            v4 = v4
                                            if false {
                                                var v68 uint32 = 26
                                                fmt.Println(v66)
                                                var v69 int32 = 1
                                                _ = v68
                                                _ = v69
                                            } else {
                                                var v70 string = "a"
                                                _ = v70
                                            }
                                        }
                                        v67 = append(v67, fi65)
                                        v71 := v1[(fi65 >> ((4294967295) & 63))]
                                        _ = fi65
                                        _ = v66
                                        _ = v67
                                        _ = v71
                                    }
                                case 128:
                                    for fi72 := 0; fi72 < 18; fi72++ {
                                        if (!(!v59)) {
                                            switch "\n\t" {
                                            case "a":
                                                delete(v1, 131)
                                            default:
                                                v73 := GenS7{F0: 2147483647, F1: 0}
                                                _ = v73
                                            }
                                            var v74 uint32 = genFn1(2)
                                            v75 := []int{(fi58 ^ (fi72 * 389)), (fi72 >> ((255) & 63))}
                                            _ = v74
                                            _ = v75
                                        } else {
                                            var v76 uint8 = 152
                                            if (0 != 9223372036854775807) {
                                                switch (v4 + "\n\t") {
                                                case "foo bar":
                                                    var v77 uint8 = v76
                                                    _ = v77
                                                case "\n\t":
                                                    v78 := GenS8{F0: 255}
                                                    _ = v78
                                                case "":
                                                    if (-841 == (fi58 % ((fi58) | 1))) {
                                                        if v59 {
                                                            var v79 uint16 = 886
                                                            v80 := []int{fi58, fi72}
                                                            if len(v80) > 0 {
                                                                v80[0] = 1
                                                            }
                                                            v76 -= (v76 & (v76 & (v76 ^ (v76 & (v76 & (v76 << ((65535) & 63)))))))
                                                            _ = v79
                                                            _ = v80
                                                        }
                                                        if ((90 != 42) && v59) {
                                                            var v81 uint32 = 42
                                                            _ = v81
                                                        }
                                                        fmt.Println(v6, v63, v2)
                                                        v4 = v4
                                                    } else {
                                                        var v82 int = 42
                                                        var v83 string = ((((("\n\t" + v4) + ("\n\t" + "hello")) + v4) + v4) + v4)
                                                        var v84 bool = v59
                                                        fmt.Println(v76)
                                                        _ = v82
                                                        _ = v83
                                                        _ = v84
                                                    }
                                                default:
                                                    v85 := GenS9{F0: v59, F1: 1}
                                                    _ = v85
                                                }
                                                v86 := GenS10{F0: -236}
                                                switch (^(fi72)) {
                                                case 128:
                                                    var v87 int8 = -1
                                                    _ = v87
                                                case -9223372036854775808:
                                                    var v88 int64 = -621
                                                    _ = v88
                                                }
                                                _ = v86
                                            }
                                            _ = v76
                                        }
                                        fmt.Println(v3, v5)
                                        fmt.Println(v59)
                                        var v89 bool = v59
                                        _ = fi72
                                        _ = v89
                                    }
                                }
                            }
                        case 4294967295:
                            var v90 uint64 = 220
                            _ = v90
                        default:
                            v91 := v1[(fi58 & fi58)]
                            _ = v91
                        }
                        v59 = v59
                        _ = v63
                    }
                    _ = v59
                }
                var v92 int = (-(fi58))
                _ = v92
            }
        case "a":
            _ = len(v1)
        default:
            for fi93 := 0; fi93 < 15; fi93++ {
                v94 := GenS8{F0: 4294967295}
                v95 := []int{fi93, (fi58 | fi58)}
                _ = fi93
                _ = v94
                _ = v95
            }
        }
        _ = fi58
    }
    v96 := v1[32767]
    v4 = (v4 + v4)
    var v97 bool = (65535 == 0)
    if v97 {
        fmt.Println(v3)
        switch v96 {
        case -589:
            v3 -= v2
        case -581:
            v6 += 0.0
        case 4294967295:
            var v98 uint32 = 4294967295
            _ = v98
        default:
            var v99 uint32 = 262
            _ = v99
        }
    }
    for fi100 := 0; fi100 < 11; fi100++ {
        v101 := []int{}
        v102 := GenS3{F0: fi100, F1: 100}
        for fi103 := 0; fi103 < 4; fi103++ {
            switch (^(v96)) {
            case 32768:
                if (1 == 0) {
                    if (42 >= 85) {
                        if v97 {
                            v104 := GenS5{F0: 283, F1: v96, F2: 1}
                            v105 := GenS3{F0: fi103, F1: 1}
                            var v106 uint8 = 27
                            _ = v104
                            _ = v105
                            _ = v106
                        } else {
                            switch (v96 | -9223372036854775808) {
                            case -2147483648:
                                if v97 {
                                    switch -9223372036854775808 {
                                    case -943:
                                        fmt.Println(v96, fi103, v4)
                                    }
                                }
                            case 2147483647:
                                switch (-(v96)) {
                                case 128:
                                    switch v4 {
                                    case "\n\t":
                                        fmt.Println(v3)
                                    }
                                case 332:
                                    if v97 {
                                        v1[(-(fi103))] = (v96 + v96)
                                        switch "hello" {
                                        case "":
                                            v2 = v3
                                        case "a":
                                            if len(v101) >= 2 {
                                                v101 = v101[1:len(v101)]
                                            }
                                        case "\n\t":
                                            if len(v101) >= 2 {
                                                v101 = v101[1:len(v101)]
                                            }
                                        }
                                        var v107 int64 = 1
                                        _ = v107
                                    } else {
                                        if true {
                                            v97 = v97
                                            switch fi103 {
                                            case 997:
                                                delete(v1, 32767)
                                            }
                                            switch v4 {
                                            case "":
                                                fmt.Println(v3, v6, v96)
                                            case "foo bar":
                                                v108 := GenS2{F0: 127}
                                                _ = v108
                                            }
                                        }
                                    }
                                case -632:
                                    if (255 >= 1) {
                                        if len(v101) >= 2 {
                                            v101 = v101[1:len(v101)]
                                        }
                                    }
                                    fmt.Println(v5)
                                }
                            default:
                                var v109 uint64 = 18446744073709551615
                                _ = v109
                            }
                            v110 := GenS11{F0: 2, F1: 4294967295}
                            _ = v110
                        }
                        v4 = v4
                        var v111 uint64 = 4294967295
                        _ = v111
                    }
                    if (!v97) {
                        var v112 bool = (0 > 127)
                        var v113 uint8 = 170
                        v114 := GenS12{F0: true, F1: v97, F2: 2147483648}
                        _ = v112
                        _ = v113
                        _ = v114
                    } else {
                        if len(v101) >= 2 {
                            v101 = v101[1:len(v101)]
                        }
                        fmt.Println(v6)
                        var v115 float32 = 1e10
                        _ = v115
                    }
                }
            case 2147483647:
                fmt.Println(fi103, v97)
            case 42:
                if v97 {
                    var v116 uint8 = 242
                    if v97 {
                        v117 := GenS1{F0: 127}
                        if true {
                            var v118 float64 = 3.14159
                            fmt.Println(fi103)
                            _ = v118
                        } else {
                            v6 = v6
                            v3 = ((v6 - v6) - v6)
                        }
                        _ = v117
                    }
                    if len(v101) >= 2 {
                        v101 = v101[1:len(v101)]
                    }
                    fmt.Println(v3)
                    _ = v116
                } else {
                    v97 = v97
                }
            default:
                switch ((v4 + "foo bar") + ("hello" + (v4 + "hello"))) {
                case "hello":
                    v1[v96] = fi100
                case "\n\t":
                    var v119 int8 = -91
                    _ = v119
                }
            }
            if (42 > 977) {
                var v120 uint64 = 42
                switch v4 {
                case "foo bar":
                    var v121 int = v96
                    _ = v121
                }
                _ = v120
            }
            var v122 float32 = 0.0
            var v123 uint64 = 7
            delete(v1, (fi103 % ((fi103) | 1)))
            _ = fi103
            _ = v122
            _ = v123
        }
        _ = fi100
        _ = v101
        _ = v102
    }
    v96 += (-(v96))
    var v124 int32 = -274
    fmt.Println(v2, v4, v3)
    var v125 int32 = 2147483647
    if (v97 || false) {
        for fi126 := 0; fi126 < 17; fi126++ {
            var v127 bool = (128 <= 190)
            v125 -= v124
            v128 := GenS4{F0: v124, F1: (v96 & (fi126 ^ 1))}
            _ = fi126
            _ = v127
            _ = v128
        }
        var v129 uint32 = 1
        _ = v129
    } else {
        v130 := []int{v96}
        _ = v130
    }
    v97 = v97
    if v97 {
        v131 := GenS13{F0: v124}
        var v132 int64 = 781
        _ = v131
        _ = v132
    } else {
        v133 := GenS14{F0: 18446744073709551615, F1: v96}
        if (v97 && v97) {
            if v97 {
                for fi134 := 0; fi134 < 10; fi134++ {
                    delete(v1, -2147483648)
                    if ((v125 % ((v124) | 1)) <= 365) {
                        for fi135 := 0; fi135 < 14; fi135++ {
                            var v136 int64 = -1
                            if (v136 <= (v136 / ((v136) | 1))) {
                                v136 = (v136 & (v136 << ((2) & 63)))
                                if (((fi135 % ((fi134) | 1)) >= 9223372036854775807) && (true || (1e10 > (0.0 + 1.5)))) {
                                    switch fi135 {
                                    case 128:
                                        switch (fi135 + 4294967295) {
                                        case -676:
                                            v96 -= (fi134 * fi134)
                                        }
                                    case 199:
                                        if false {
                                            if (42 > 2147483648) {
                                                v97 = true
                                                if (632 != v96) {
                                                    fmt.Println(fi135, v6, v6)
                                                    v137 := GenS15{F0: 715}
                                                    _ = len(v1)
                                                    _ = v137
                                                } else {
                                                    v138 := GenS1{F0: 128}
                                                    var v139 string = (v4 + v4)
                                                    fmt.Println(v139, fi134)
                                                    _ = v138
                                                    _ = v139
                                                }
                                                v5 += v5
                                            } else {
                                                v2 -= float64(1.5)
                                            }
                                        } else {
                                            if v97 {
                                                if (0 != 949) {
                                                    var v140 uint64 = 1
                                                    if true {
                                                        if false {
                                                            fmt.Println(v125, v97, v97)
                                                        } else {
                                                            fmt.Println(v97, v125)
                                                        }
                                                    } else {
                                                        v141 := v1[1]
                                                        delete(v1, (-(v96)))
                                                        _ = v141
                                                    }
                                                    switch fi135 {
                                                    case -9223372036854775808:
                                                        if (1 > -1) {
                                                            _ = len(v1)
                                                            v6 = (0.1 / v3)
                                                            v5 = v5
                                                        }
                                                    case 128:
                                                        v142 := GenS16{F0: 1, F1: 127, F2: 1}
                                                        _ = v142
                                                    default:
                                                        if false {
                                                            if (v97 || (!(!false))) {
                                                                if false {
                                                                    var v143 int8 = -1
                                                                    _ = v143
                                                                } else {
                                                                    v144 := GenS2{F0: 42}
                                                                    _ = v144
                                                                }
                                                                v140 += 42
                                                                switch ((v4 + v4) + v4) {
                                                                case "foo bar":
                                                                    fmt.Println(v6, v2, v97)
                                                                case "":
                                                                    fmt.Println(v6, fi135, v3)
                                                                }
                                                                if false {
                                                                    v140 = (v140 - 131)
                                                                }
                                                            } else {
                                                                v140 -= (v140 - (v140 | v140))
                                                            }
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                var v145 uint8 = 128
                                                                _ = v145
                                                            default:
                                                                v146 := GenS17{F0: 42}
                                                                _ = v146
                                                            }
                                                            fmt.Println(v6)
                                                            var v147 string = v4
                                                            v148 := GenS10{F0: -542}
                                                            _ = v147
                                                            _ = v148
                                                        } else {
                                                            v3 = 1.5
                                                            switch fi135 {
                                                            case -2147483648:
                                                                v149 := GenS5{F0: 18446744073709551615, F1: v96, F2: v125}
                                                                _ = v149
                                                            case 127:
                                                                fmt.Println(v4)
                                                            case 128:
                                                                var v150 int8 = 127
                                                                _ = v150
                                                            }
                                                            var v151 int32 = 218
                                                            _ = v151
                                                        }
                                                    }
                                                    _ = v140
                                                }
                                                var v152 uint32 = 603
                                                v153 := []int{fi134, (fi134 << ((255) & 63))}
                                                if v97 {
                                                    var v154 int16 = (-(v5))
                                                    if v97 {
                                                        if (4294967295 < 791) {
                                                            v97 = ((v125 & -735) == v124)
                                                            fmt.Println(v6)
                                                        } else {
                                                            var v155 uint32 = v152
                                                            v2 = -2.25
                                                            var v156 uint64 = 42
                                                            _ = v155
                                                            _ = v156
                                                        }
                                                    }
                                                    _ = v154
                                                }
                                                var v157 uint16 = 255
                                                var v158 bool = false
                                                var v159 int = fi134
                                                _ = v152
                                                _ = v153
                                                _ = v157
                                                _ = v158
                                                _ = v159
                                            } else {
                                                v1[(-(fi135))] = 450
                                                var v160 bool = (3 > 1)
                                                _ = v160
                                            }
                                            switch v4 {
                                            case "foo bar":
                                                if ("" < v4) {
                                                    var v161 uint16 = 65535
                                                    v162 := []int{}
                                                    var v163 uint = 65535
                                                    _ = v161
                                                    _ = v162
                                                    _ = v163
                                                } else {
                                                    v164 := GenS4{F0: (v124 ^ (v124 << ((255) & 63))), F1: (^(fi134))}
                                                    _ = v164
                                                }
                                            default:
                                                switch -494 {
                                                case 128:
                                                    fmt.Println(v125, v136)
                                                case -1:
                                                    var v165 uint8 = 128
                                                    _ = v165
                                                case -290:
                                                    fmt.Println(v125, fi134)
                                                }
                                            }
                                        }
                                        switch v4 {
                                        case "\n\t":
                                            v3 = 1e10
                                        case "a":
                                            v166 := GenS18{F0: 8, F1: 255, F2: 87}
                                            _ = v166
                                        case "":
                                            v2 -= 1.5
                                            fmt.Println(v3)
                                            v167 := GenS19{F0: 18446744073709551615, F1: -1}
                                            _ = v167
                                        }
                                    case 9223372036854775807:
                                        v168 := v1[v96]
                                        fmt.Println(fi135, v6, v4)
                                        v169 := GenS20{F0: (1 == 42), F1: (v5 & 0), F2: (-(v5))}
                                        v170 := []int{(v96 ^ v168), (v168 + (v168 & v96)), fi134}
                                        _ = v168
                                        _ = v169
                                        _ = v170
                                    }
                                } else {
                                    v171 := GenS11{F0: 322, F1: 18446744073709551615}
                                    _ = v171
                                }
                                var v172 float32 = 1e10
                                _ = v172
                            } else {
                                if false {
                                    if (0 > 926) {
                                        if (0 != (v136 - 42)) {
                                            var v173 float32 = 1e10
                                            if false {
                                                if v97 {
                                                    v136 += (-(v136))
                                                    var v174 int = v96
                                                    if v97 {
                                                        fmt.Println(fi134, v97, v124)
                                                        var v175 uint8 = 2
                                                        _ = v175
                                                    } else {
                                                        if (2147483648 >= 285) {
                                                            fmt.Println(v173)
                                                            var v176 bool = (771 != (v96 | 128))
                                                            _ = v176
                                                        } else {
                                                            v96 += (fi134 >> ((255) & 63))
                                                        }
                                                        var v177 uint64 = 1
                                                        _ = v177
                                                    }
                                                    _ = v174
                                                }
                                                fmt.Println(v2, v136)
                                                if (genFn1(1) == 2) {
                                                    v178 := []int{42, 42, 661, (fi134 + v96)}
                                                    if (!v97) {
                                                        var v179 float64 = v6
                                                        fmt.Println(v6)
                                                        switch (v4 + v4) {
                                                        case "\n\t":
                                                            v173 += v173
                                                        default:
                                                            if (!(192 <= v5)) {
                                                                var v180 uint32 = 2
                                                                v1[fi134] = -372
                                                                _ = len(v1)
                                                                _ = v180
                                                            } else {
                                                                v3 = v179
                                                                var v181 int8 = int8(v124)
                                                                _ = v181
                                                            }
                                                        }
                                                        _ = v179
                                                    }
                                                    v182 := GenS21{F0: 0, F1: v124, F2: v136}
                                                    _ = v178
                                                    _ = v182
                                                } else {
                                                    v136 -= (v136 - v136)
                                                    fmt.Println(v6)
                                                }
                                                switch (v96 ^ v96) {
                                                case 4294967295:
                                                    v183 := GenS22{F0: 8, F1: 65535, F2: (v124 + 1)}
                                                    _ = v183
                                                }
                                                v184 := v1[v96]
                                                var v185 int = v184
                                                if ((true || (-52 > 0)) || v97) {
                                                    var v186 uint8 = 255
                                                    fmt.Println(v2)
                                                    fmt.Println(fi135, v97, fi134)
                                                    _ = v186
                                                } else {
                                                    if (2 == 32767) {
                                                        if ((true || v97) || false) {
                                                            var v187 uint16 = 32767
                                                            switch (v185 + 128) {
                                                            case 2147483647:
                                                                fmt.Println(v2)
                                                            case 42:
                                                                var v188 uint32 = genFn1(v187)
                                                                _ = v188
                                                            case 128:
                                                                fmt.Println(v97, v96, v2)
                                                            default:
                                                                v4 = "foo bar"
                                                            }
                                                            v189 := GenS23{F0: 256, F1: v4, F2: (!false)}
                                                            _ = v187
                                                            _ = v189
                                                        } else {
                                                            if v97 {
                                                                var v190 bool = (("\n\t" + (("foo bar" + v4) + (v4 + v4))) < ((("hello" + "hello") + "") + v4))
                                                                v191 := []int{v184, (v96 - (^(fi134)))}
                                                                v192 := v1[(v96 / ((int(v136)) | 1))]
                                                                _ = v190
                                                                _ = v191
                                                                _ = v192
                                                            } else {
                                                                v3 = v3
                                                                switch -621 {
                                                                case 127:
                                                                    if false {
                                                                        v193 := GenS24{F0: v96, F1: (v5 & v5)}
                                                                        var v194 bool = v97
                                                                        v2 += float64(v173)
                                                                        _ = v193
                                                                        _ = v194
                                                                    } else {
                                                                        switch (v4 + "a") {
                                                                        case "":
                                                                            v184 += fi135
                                                                        case "a":
                                                                            v5 = -167
                                                                        case "foo bar":
                                                                            var v195 int64 = v136
                                                                            _ = v195
                                                                        }
                                                                        v97 = v97
                                                                    }
                                                                case 32768:
                                                                    var v196 uint32 = 2147483647
                                                                    _ = v196
                                                                default:
                                                                    var v197 float32 = float32((1.5 - (float64(1e10) + v3)))
                                                                    _ = v197
                                                                }
                                                            }
                                                            if v97 {
                                                                var v198 uint8 = 128
                                                                _ = v198
                                                            } else {
                                                                v199 := GenS25{F0: 1, F1: v5, F2: 151}
                                                                v200 := []int{}
                                                                _ = v199
                                                                _ = v200
                                                            }
                                                            var v201 uint = 1
                                                            _ = v201
                                                        }
                                                        var v202 int16 = (v5 - (v5 >> ((603) & 63)))
                                                        v124 -= 2147483647
                                                        _ = v202
                                                    } else {
                                                        v125 -= (^(v125))
                                                        v136 = v136
                                                        if true {
                                                            v1[v184] = 32768
                                                            v2 += 3.14159
                                                        } else {
                                                            v96 = fi134
                                                            v5 = -1
                                                        }
                                                    }
                                                    v203 := GenS7{F0: (v124 >> ((18446744073709551615) & 63)), F1: 2}
                                                    if (875 == 190) {
                                                        var v204 uint8 = 42
                                                        var v205 uint = 42
                                                        v97 = (v6 == (1.5 / (float64((v173 - v173)) - v3)))
                                                        switch v4 {
                                                        case "\n\t":
                                                            v206 := GenS19{F0: 1, F1: v136}
                                                            _ = v206
                                                        case "a":
                                                            if v97 {
                                                                v207 := []int{42}
                                                                if (v97 || v97) {
                                                                    var v208 int8 = 127
                                                                    var v209 float64 = -2.25
                                                                    fmt.Println(v3, v6)
                                                                    _ = v208
                                                                    _ = v209
                                                                } else {
                                                                    var v210 int32 = v125
                                                                    v205 = v205
                                                                    v1[v185] = (v185 % (((v185 * v185)) | 1))
                                                                    _ = v210
                                                                }
                                                                if (v97 && ((!((v205 ^ v205) < (v205 - (v205 >> (((^(v205))) & 63))))) && (42 == 4294967295))) {
                                                                    v97 = v97
                                                                }
                                                                _ = v207
                                                            }
                                                        case "foo bar":
                                                            v205 += (v205 % ((v205) | 1))
                                                        default:
                                                            if v97 {
                                                                fmt.Println(v97, v96)
                                                                v211 := []int{v184, 9223372036854775807, (fi134 ^ (^(fi134)))}
                                                                fmt.Println(v4, v204, v4)
                                                                _ = v211
                                                            }
                                                        }
                                                        _ = v204
                                                        _ = v205
                                                    }
                                                    _ = v203
                                                }
                                                v212 := GenS9{F0: true, F1: v136}
                                                if ((v136 >> ((4294967295) & 63)) > 809) {
                                                    var v213 float32 = (v173 + 1e10)
                                                    _ = v213
                                                } else {
                                                    v214 := GenS11{F0: 940, F1: 42}
                                                    v215 := GenS4{F0: v125, F1: (-(fi135))}
                                                    _ = v214
                                                    _ = v215
                                                }
                                                _ = v184
                                                _ = v185
                                                _ = v212
                                            }
                                            _ = v173
                                        } else {
                                            if (true && v97) {
                                                v4 = v4
                                            }
                                        }
                                        v136 -= v136
                                    }
                                }
                            }
                            v216 := GenS26{F0: 42, F1: 42}
                            _ = fi135
                            _ = v136
                            _ = v216
                        }
                    } else {
                        fmt.Println(v125)
                    }
                    if false {
                        v96 = v96
                    }
                    _ = fi134
                }
                for fi217 := 0; fi217 < 3; fi217++ {
                    for fi218 := 0; fi218 < 14; fi218++ {
                        v6 = v3
                        fmt.Println(v5, v3)
                        var v219 uint32 = 2147483647
                        _ = fi218
                        _ = v219
                    }
                    v220 := GenS27{F0: v4}
                    _ = fi217
                    _ = v220
                }
                switch ("\n\t" + "foo bar") {
                case "hello":
                    var v221 float64 = v3
                    _ = v221
                case "":
                    v4 = ("a" + v4)
                default:
                    switch v4 {
                    case "hello":
                        switch (v96 * (v96 & -2147483648)) {
                        case 1:
                            var v222 uint16 = 693
                            _ = v222
                        }
                    case "foo bar":
                        switch ((v4 + v4) + "hello") {
                        case "foo bar":
                            fmt.Println(v96)
                        case "\n\t":
                            var v223 uint = 255
                            _ = v223
                        }
                    case "":
                        fmt.Println(v124, v125)
                    }
                }
            } else {
                v224 := v1[(v96 & 42)]
                for fi225 := 0; fi225 < 17; fi225++ {
                    delete(v1, fi225)
                    _ = fi225
                }
                for fi226 := 0; fi226 < 8; fi226++ {
                    v125 += v124
                    v6 += -2.25
                    v227 := GenS6{F0: (-(v5))}
                    _ = fi226
                    _ = v227
                }
                _ = v224
            }
            var v228 uint16 = 42
            _ = v228
        }
        var v229 int = 42
        _ = v133
        _ = v229
    }
    v4 = (v4 + v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v124:", v124)
    fmt.Println("v125:", v125)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v124
    _ = v125
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v96
    _ = v97
}
