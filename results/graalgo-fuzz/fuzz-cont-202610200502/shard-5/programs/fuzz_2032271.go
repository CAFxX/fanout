package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int64 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 int64; F1 uint16; F2 int16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint16; F1 uint64; F2 int64 }
type GenS9 struct { F0 uint; F1 uint }
type GenS10 struct { F0 uint; F1 int8; F2 int32 }
type GenS11 struct { F0 bool; F1 int8 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 uint16; F1 int32; F2 bool }
type GenS14 struct { F0 uint64; F1 string }
func genFn0() int32 {
    switch 9223372036854775807 {
    case 32768:
        var v7 uint64 = 1
        _ = v7
    case 42:
        if false {
            v8 := GenS1{F0: -710, F1: 114, F2: 1}
            var v9 int32 = -1
            _ = v8
            _ = v9
        }
    case -9223372036854775808:
        switch 127 {
        case 926:
            var v10 int32 = -2147483648
            _ = v10
        case 4294967295:
            var v11 int64 = 42
            _ = v11
        default:
            var v12 bool = ((1 == -32768) && true)
            _ = v12
        }
    default:
        var v13 float64 = 0.1
        _ = v13
    }
    var v14 int32 = -274
    if (42 > 128) {
        v15 := GenS1{F0: 1, F1: 128, F2: 42}
        _ = v15
    } else {
        switch "" {
        case "hello":
            v16 := []int{-882, 42, -26}
            _ = v16
        case "foo bar":
            for fi17 := 0; fi17 < 10; fi17++ {
                v18 := GenS1{F0: -1, F1: 42, F2: 299}
                var v19 int64 = -196
                for fi20 := 0; fi20 < 18; fi20++ {
                    fmt.Println(v19, fi17, v19)
                    _ = fi20
                }
                _ = fi17
                _ = v18
                _ = v19
            }
        case "":
            if true {
                v14 += (v14 / (((v14 % (((v14 << ((503) & 63))) | 1))) | 1))
                var v21 uint8 = 128
                for fi22 := 0; fi22 < 12; fi22++ {
                    for fi23 := 0; fi23 < 10; fi23++ {
                        var v24 uint32 = 671
                        if true {
                            if false {
                                v25 := GenS1{F0: 0, F1: v21, F2: 42}
                                _ = v25
                            } else {
                                v26 := GenS1{F0: -456, F1: v21, F2: 122}
                                _ = v26
                            }
                            var v27 float64 = ((3.14159 * 0.0) - 0.1)
                            v28 := []int{}
                            _ = v27
                            _ = v28
                        }
                        v29 := []int{(fi22 & (fi23 >> ((42) & 63))), fi23, (fi22 + (fi22 ^ fi23))}
                        _ = fi23
                        _ = v24
                        _ = v29
                    }
                    _ = fi22
                }
                _ = v21
            }
        default:
            var v30 uint8 = 2
            _ = v30
        }
        switch "\n\t" {
        case "a":
            var v31 string = "foo bar"
            _ = v31
        case "hello":
            if false {
                v14 = -375
            }
        }
        fmt.Println(v14)
    }
    switch int(v14) {
    case -806:
        for fi32 := 0; fi32 < 9; fi32++ {
            if ((!false) || (true || false)) {
                if (!(!false)) {
                    v14 = v14
                    if false {
                        var v33 uint32 = 0
                        v33 -= v33
                        _ = v33
                    }
                    fmt.Println(fi32)
                }
                v34 := []int{fi32}
                _ = v34
            } else {
                fmt.Println(fi32, v14)
            }
            switch fi32 {
            case -2147483648:
                v14 += (v14 * v14)
            }
            var v35 float64 = -2.25
            _ = fi32
            _ = v35
        }
    default:
        v14 += 2147483647
    }
    _ = v14
    return v14
}
func genFn1() bool {
    if true {
        if true {
            v36 := []int{485, -395}
            if false {
                v37 := GenS1{F0: -1, F1: 197, F2: 0}
                v38 := GenS2{F0: false, F1: -494}
                _ = v37
                _ = v38
            } else {
                v36 = append(v36, -1)
                v36 = append(v36, -2147483648)
            }
            _ = v36
        }
        if (-14 > 122) {
            var v39 float64 = (0.0 * 0.1)
            for fi40 := 0; fi40 < 7; fi40++ {
                for fi41 := 0; fi41 < 3; fi41++ {
                    v42 := []int{fi41, (fi41 * -9223372036854775808), fi40}
                    _ = fi41
                    _ = v42
                }
                var v43 uint8 = 73
                _ = fi40
                _ = v43
            }
            if true {
                v44 := GenS3{F0: "hello"}
                v45 := GenS1{F0: -9223372036854775808, F1: 2, F2: 339}
                var v46 uint8 = 205
                _ = v44
                _ = v45
                _ = v46
            } else {
                v39 = 1.5
                var v47 uint = 4294967295
                _ = v47
            }
            _ = v39
        }
        for fi48 := 0; fi48 < 4; fi48++ {
            for fi49 := 0; fi49 < 7; fi49++ {
                var v50 uint64 = 42
                v50 = (v50 - (v50 / (((v50 - (v50 / ((v50) | 1)))) | 1)))
                v50 = 4294967295
                var v51 int64 = -1
                _ = fi49
                _ = v50
                _ = v51
            }
            var v52 int64 = 642
            _ = fi48
            _ = v52
        }
    }
    var v53 uint16 = 550
    v53 -= (v53 ^ (v53 << ((1) & 63)))
    _ = v53
    return true
}

func main() {
    var v1 string = "foo bar"
    var v2 string = ""
    var v3 bool = ((-2.25 - 1.5) > 1.5)
    v4 := make(map[string]int)
    v5 := []uint32{2147483648, 985}
    var v6 float32 = ((float32(1e10) - ((1.5 + 3.14159) - (1.5 + 0.0))) / (((1.5 * 1e10) / 1.0) - 1.5))
    v2 = v2
    for fi54 := 0; fi54 < 11; fi54++ {
        v55 := GenS2{F0: v3, F1: int64(fi54)}
        v6 = float32(((-2.25 + (1.5 + float64(((v6 / v6) * v6)))) / (0.1 * 0.1)))
        if (391 > 4294967295) {
            for fi56 := 0; fi56 < 6; fi56++ {
                fmt.Println(fi56, fi54)
                _ = fi56
            }
            for fi57 := 0; fi57 < 13; fi57++ {
                v58 := GenS1{F0: 120, F1: 151, F2: -9223372036854775808}
                var v59 uint16 = 599
                _ = fi57
                _ = v58
                _ = v59
            }
            for fi60 := 0; fi60 < 1; fi60++ {
                fmt.Println(fi60, fi54)
                _ = fi60
            }
        }
        _ = fi54
        _ = v55
    }
    var v61 int32 = -1
    var v62 uint = 26
    if (2 != 42) {
        switch "hello" {
        case "":
            for fi63 := 0; fi63 < 12; fi63++ {
                for fi64 := 0; fi64 < 14; fi64++ {
                    v6 = v6
                    _ = fi64
                }
                for fi65 := 0; fi65 < 13; fi65++ {
                    v66 := GenS4{F0: 827, F1: v2}
                    var v67 uint64 = 0
                    _ = fi65
                    _ = v66
                    _ = v67
                }
                _ = fi63
            }
        case "\n\t":
            v6 = v6
        }
        v1 = "hello"
        switch ("" + (v1 + v1)) {
        case "":
            if (71 != 103) {
                v6 = v6
                v5 = append(v5, 1)
                for fi68 := 0; fi68 < 5; fi68++ {
                    switch (fi68 & (fi68 ^ (^(fi68)))) {
                    case 4294967295:
                        for fi69 := 0; fi69 < 2; fi69++ {
                            v70 := GenS5{F0: (v61 + (v61 / (((v61 + (v61 / ((int32(fi68)) | 1)))) | 1))), F1: (v61 + (v61 / (((v61 + (v61 * 1))) | 1)))}
                            v1 = "\n\t"
                            if (((4294967295 < 441) || (255 == 127)) && false) {
                                var v71 int64 = -306
                                if ((2147483647 > 55) || v3) {
                                    switch v1 {
                                    case "a":
                                        v61 = v61
                                    default:
                                        fmt.Println(v3)
                                    }
                                } else {
                                    v72 := v4[v1]
                                    var v73 float32 = (1.5 / (v6 - (((v6 * v6) + (float32(1.5) * 0.0)) + v6)))
                                    _ = v72
                                    _ = v73
                                }
                                _ = v71
                            }
                            _ = fi69
                            _ = v70
                        }
                    case 492:
                        delete(v4, "")
                    case 42:
                        var v74 uint8 = 2
                        _ = v74
                    default:
                        v6 = ((0.0 * 0.0) - (v6 - v6))
                    }
                    var v75 uint16 = 32767
                    _ = fi68
                    _ = v75
                }
            } else {
                var v76 int64 = -39
                var v77 bool = v3
                _ = v76
                _ = v77
            }
        case "a":
            for fi78 := 0; fi78 < 11; fi78++ {
                var v79 int32 = (^(v61))
                if ((v6 >= (((v6 - v6) - 1.5) / -2.25)) && ((255 > 2) && false)) {
                    delete(v4, v2)
                    v4[v2] = fi78
                    delete(v4, (v1 + "a"))
                } else {
                    v2 = "foo bar"
                    delete(v4, (v2 + v2))
                    if (false || (42 < 429)) {
                        v61 = genFn0()
                        v80 := GenS2{F0: v3, F1: 42}
                        if (!v3) {
                            for fi81 := 0; fi81 < 14; fi81++ {
                                delete(v4, ("\n\t" + (v1 + "foo bar")))
                                _ = fi81
                            }
                        }
                        _ = v80
                    } else {
                        for fi82 := 0; fi82 < 19; fi82++ {
                            v83 := GenS6{F0: 42, F1: 1, F2: 0}
                            _ = fi82
                            _ = v83
                        }
                        v79 = genFn0()
                    }
                }
                _ = fi78
                _ = v79
            }
        }
    }
    for fi84 := 0; fi84 < 16; fi84++ {
        for fi85 := 0; fi85 < 5; fi85++ {
            v3 = (0 != 0)
            _ = fi85
        }
        v86 := GenS5{F0: int32(v62), F1: (v61 & v61)}
        var v87 uint8 = 38
        _ = fi84
        _ = v86
        _ = v87
    }
    var v88 int8 = -77
    for fi89 := 0; fi89 < 20; fi89++ {
        if (!false) {
            if ((v88 == (-(v88))) && v3) {
                var v90 uint64 = 4294967295
                var v91 uint32 = 0
                v90 = (v90 >> ((0) & 63))
                _ = v90
                _ = v91
            }
            v5 = append(v5, 929)
        } else {
            if false {
                for fi92 := 0; fi92 < 20; fi92++ {
                    if false {
                        var v93 float64 = 1.5
                        var v94 int = (fi92 % (((fi92 >> (((v62 * 18446744073709551615)) & 63))) | 1))
                        _ = v93
                        _ = v94
                    } else {
                        _ = len(v4)
                    }
                    var v95 uint32 = 2
                    if v3 {
                        var v96 uint = v62
                        var v97 int16 = 1
                        _ = v96
                        _ = v97
                    }
                    fmt.Println(v88, fi89)
                    switch fi89 {
                    case 0:
                        switch (fi92 + (fi92 % ((32768) | 1))) {
                        case -49:
                            if (v3 && (0 != 165)) {
                                var v98 string = ("" + v1)
                                _ = v98
                            } else {
                                var v99 int16 = -419
                                _ = v99
                            }
                        default:
                            if ((((((1.5 / 1.5) / float64(float32(1e10))) + ((-2.25 / 1e10) + 0.1)) + 0.1) - (-2.25 / float64(v6))) < 1e10) {
                                v3 = true
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                switch ("" + (v2 + (v2 + v2))) {
                                case "foo bar":
                                    fmt.Println(v2, fi92)
                                case "a":
                                    if true {
                                        fmt.Println(v6)
                                        v2 = ""
                                    }
                                default:
                                    if ((v62 % ((1) | 1)) < v62) {
                                        v61 += v61
                                        var v100 int16 = 0
                                        switch "a" {
                                        case "\n\t":
                                            v62 = 255
                                            fmt.Println(v61, v2, v95)
                                            if v3 {
                                                v2 = v1
                                                fmt.Println(v1)
                                                delete(v4, v1)
                                                var v101 string = "\n\t"
                                                if (true || (-2.25 != (1e10 * float64(((v6 / (((v6 + (v6 + float32((-2.25 - -2.25)))) + v6) + 1e10)) * v6))))) {
                                                    var v102 int32 = genFn0()
                                                    _ = v102
                                                }
                                                _ = v101
                                            }
                                        case "foo bar":
                                            switch v2 {
                                            case "":
                                                if (-622 < 1) {
                                                    v88 = (v88 | int8(v95))
                                                }
                                            }
                                        case "hello":
                                            v100 = (^(v100))
                                        }
                                        _ = v100
                                    }
                                    var v103 int16 = -21
                                    if v3 {
                                        var v104 uint = v62
                                        _ = v104
                                    }
                                    _ = v103
                                }
                            } else {
                                v1 = v1
                                var v105 float64 = (1.5 + 3.14159)
                                _ = v105
                            }
                        }
                    }
                    _ = fi92
                    _ = v95
                }
                fmt.Println(v6)
            } else {
                fmt.Println(v61)
                v106 := GenS2{F0: (18446744073709551615 > 91), F1: -485}
                _ = v106
            }
        }
        switch (^(fi89)) {
        case 42:
            var v107 int = fi89
            _ = v107
        default:
            if (0 != 128) {
                switch (^(fi89)) {
                case -2147483648:
                    var v108 bool = false
                    _ = v108
                case 0:
                    switch fi89 {
                    case 32767:
                        v1 = "foo bar"
                    case 111:
                        for fi109 := 0; fi109 < 15; fi109++ {
                            v110 := GenS5{F0: int32(fi109), F1: v61}
                            switch fi109 {
                            case 9223372036854775807:
                                v1 = ((v1 + (v2 + "")) + (v1 + (v1 + (v2 + (v2 + "hello")))))
                            case -353:
                                switch v2 {
                                case "":
                                    var v111 uint16 = 1
                                    _ = v111
                                }
                            case -9223372036854775808:
                                v112 := GenS6{F0: -930, F1: 2, F2: -32768}
                                _ = v112
                            default:
                                if len(v5) > 0 {
                                    v5[0] = 1
                                }
                            }
                            _ = fi109
                            _ = v110
                        }
                    }
                case 32768:
                    var v113 uint32 = 2147483647
                    _ = v113
                }
                for fi114 := 0; fi114 < 19; fi114++ {
                    v115 := GenS7{F0: 769}
                    v116 := GenS8{F0: 32767, F1: 0, F2: 1}
                    v3 = (!false)
                    _ = fi114
                    _ = v115
                    _ = v116
                }
            }
        }
        for fi117 := 0; fi117 < 4; fi117++ {
            fmt.Println(fi117)
            v118 := GenS9{F0: v62, F1: 2}
            _ = fi117
            _ = v118
        }
        _ = fi89
    }
    v88 += 100
    v119 := GenS6{F0: -1, F1: 482, F2: 32767}
    switch v2 {
    case "hello":
        switch "foo bar" {
        case "a":
            v120 := GenS5{F0: v61, F1: 2147483647}
            _ = v120
        default:
            v61 = (v61 / ((34) | 1))
        }
    case "a":
        fmt.Println(v2, v62)
    }
    var v121 float64 = float64((0.0 + 1e10))
    v122 := GenS10{F0: v62, F1: 123, F2: -673}
    fmt.Println(v121)
    var v123 uint16 = 2
    if (v3 && (!(v2 > v2))) {
        switch v2 {
        case "hello":
            switch 204 {
            case -97:
                v124 := GenS3{F0: "foo bar"}
                _ = v124
            default:
                if (!(131 < 154)) {
                    for fi125 := 0; fi125 < 3; fi125++ {
                        if (0 <= 1) {
                            _ = len(v4)
                            if ((float32((float64(0.0) - ((v121 + (v121 / v121)) - v121))) / v6) <= float32(((v121 * (((((0.1 - (-2.25 / (((v121 + (-2.25 / (-2.25 - 1e10))) + float64(((0.1 * ((v6 - 1.5) - (v6 + (v6 * (v6 + 0.0))))) / (v6 / -2.25)))) / v121))) * v121) - float64(v6)) - (float64((v6 * v6)) - float64(v6))) - 0.1)) / v121))) {
                                var v126 float64 = (v121 + 0.1)
                                _ = v126
                            } else {
                                v121 = v121
                                var v127 float32 = v6
                                fmt.Println(v62, v62)
                                _ = v127
                            }
                        } else {
                            if false {
                                switch "\n\t" {
                                case "foo bar":
                                    v88 += int8(v61)
                                case "hello":
                                    v128 := GenS11{F0: ((v62 >= v62) || (3.14159 == ((((v6 * v6) / 1.5) - (v6 + v6)) + v6))), F1: v88}
                                    _ = v128
                                case "\n\t":
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                            } else {
                                if false {
                                    var v129 uint16 = v123
                                    fmt.Println(v61, v62, v129)
                                    _ = v129
                                } else {
                                    fmt.Println(v6, v121, v3)
                                    if ((v121 + (v121 / v121)) < v121) {
                                        v4[v2] = fi125
                                        switch (fi125 - (fi125 + 387)) {
                                        case 128:
                                            v130 := GenS12{F0: v61}
                                            _ = v130
                                        default:
                                            if len(v5) > 0 {
                                                v5[0] = 886
                                            }
                                        }
                                    }
                                    v6 -= v6
                                }
                            }
                        }
                        fmt.Println(v1)
                        _ = fi125
                    }
                    for fi131 := 0; fi131 < 10; fi131++ {
                        switch (fi131 / (((fi131 << (((v62 << ((v62) & 63))) & 63))) | 1)) {
                        case -9223372036854775808:
                            if v3 {
                                if v3 {
                                    v88 -= (-(v88))
                                    var v132 float32 = v6
                                    var v133 int8 = (-(v88))
                                    _ = v132
                                    _ = v133
                                }
                            } else {
                                if false {
                                    var v134 int16 = -562
                                    _ = v134
                                }
                            }
                        case 0:
                            if v3 {
                                v135 := GenS1{F0: -632, F1: 255, F2: 9223372036854775807}
                                _ = v135
                            } else {
                                v121 += -2.25
                                v62 -= (^(v62))
                            }
                        default:
                            var v136 float32 = -2.25
                            _ = v136
                        }
                        v137 := GenS5{F0: (v61 << ((v62) & 63)), F1: genFn0()}
                        v138 := GenS4{F0: v123, F1: v2}
                        _ = fi131
                        _ = v137
                        _ = v138
                    }
                } else {
                    for fi139 := 0; fi139 < 5; fi139++ {
                        if (v3 && v3) {
                            if len(v5) > 0 {
                                v5[0] = 548
                            }
                        } else {
                            var v140 float32 = 3.14159
                            switch "" {
                            case "":
                                if v3 {
                                    switch (-(fi139)) {
                                    case -9223372036854775808:
                                        if false {
                                            switch "a" {
                                            case "foo bar":
                                                fmt.Println(fi139, v2)
                                            case "a":
                                                v4[("" + "")] = fi139
                                            }
                                            switch 9223372036854775807 {
                                            case -2147483648:
                                                for fi141 := 0; fi141 < 1; fi141++ {
                                                    var v142 uint = (v62 ^ v62)
                                                    var v143 uint16 = v123
                                                    var v144 float32 = 0.0
                                                    _ = fi141
                                                    _ = v142
                                                    _ = v143
                                                    _ = v144
                                                }
                                            case 577:
                                                switch (fi139 >> ((v62) & 63)) {
                                                case 1:
                                                    for fi145 := 0; fi145 < 14; fi145++ {
                                                        switch fi145 {
                                                        case 0:
                                                            var v146 bool = (!v3)
                                                            _ = v146
                                                        case -9223372036854775808:
                                                            switch ("\n\t" + ("\n\t" + v2)) {
                                                            case "a":
                                                                var v147 uint8 = 0
                                                                _ = v147
                                                            }
                                                        }
                                                        v121 -= (v121 - v121)
                                                        v140 = -2.25
                                                        _ = fi145
                                                    }
                                                case -1:
                                                    switch fi139 {
                                                    case 42:
                                                        if (false || (0 != 1)) {
                                                            fmt.Println(v140, v2, fi139)
                                                        }
                                                    default:
                                                        var v148 int32 = v61
                                                        _ = v148
                                                    }
                                                case -684:
                                                    var v149 int = -1
                                                    _ = v149
                                                }
                                            default:
                                                if false {
                                                    v5 = append(v5, 1)
                                                } else {
                                                    v88 = (v88 ^ v88)
                                                }
                                            }
                                            v3 = v3
                                        } else {
                                            v61 = (v61 | v61)
                                            fmt.Println(v61)
                                        }
                                    case -520:
                                        fmt.Println(v88, v123, v3)
                                    case 127:
                                        switch (((v2 + "\n\t") + v2) + "foo bar") {
                                        case "":
                                            if v3 {
                                                switch v1 {
                                                case "hello":
                                                    v121 += -2.25
                                                case "":
                                                    var v150 bool = v3
                                                    _ = v150
                                                }
                                            }
                                        case "a":
                                            v6 -= 0.1
                                        default:
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        }
                                    default:
                                        v151 := v4[((v1 + v1) + v2)]
                                        _ = v151
                                    }
                                } else {
                                    var v152 uint64 = 1
                                    _ = v152
                                }
                            }
                            fmt.Println(v123, v62, v3)
                            _ = v140
                        }
                        _ = fi139
                    }
                    for fi153 := 0; fi153 < 8; fi153++ {
                        fmt.Println(v61, v1, fi153)
                        _ = fi153
                    }
                    fmt.Println(v2, v6)
                }
            }
        case "a":
            fmt.Println(v61, v62, v62)
        case "":
            for fi154 := 0; fi154 < 12; fi154++ {
                switch 32768 {
                case 32767:
                    switch -508 {
                    case -2147483648:
                        var v155 int16 = -396
                        _ = v155
                    }
                case -9223372036854775808:
                    var v156 uint8 = 255
                    _ = v156
                }
                fmt.Println(v6, v88)
                if false {
                    if v3 {
                        switch 9223372036854775807 {
                        case -343:
                            switch ("" + v1) {
                            case "hello":
                                v157 := GenS1{F0: -590, F1: 2, F2: 9223372036854775807}
                                _ = v157
                            case "a":
                                switch v1 {
                                case "":
                                    for fi158 := 0; fi158 < 1; fi158++ {
                                        var v159 int = 128
                                        if len(v5) > 0 {
                                            v5[0] = 65535
                                        }
                                        fmt.Println(v88, v3, v123)
                                        var v160 int8 = v88
                                        fmt.Println(v1)
                                        _ = fi158
                                        _ = v159
                                        _ = v160
                                    }
                                default:
                                    switch ("\n\t" + (v2 + v1)) {
                                    case "a":
                                        for fi161 := 0; fi161 < 2; fi161++ {
                                            var v162 bool = ((v121 * v121) <= (float64((0.0 / float32(0.1))) * (((v121 - 0.1) / v121) - ((((1.5 + ((v121 - v121) / ((v121 + ((v121 * v121) * v121)) * (3.14159 + 1e10)))) / (float64(((float32((v121 / ((v121 - float64(v6)) * float64(v6)))) * 3.14159) + v6)) - v121)) - -2.25) - v121))))
                                            v4["\n\t"] = (^(fi154))
                                            v163 := GenS2{F0: v3, F1: 818}
                                            var v164 int8 = 123
                                            v165 := GenS1{F0: 0, F1: 2, F2: 472}
                                            if v3 {
                                                if ((float64((1.5 / 1.0)) - 0.0) >= v121) {
                                                    if true {
                                                        v123 -= (v123 ^ v123)
                                                        var v166 float64 = 1.5
                                                        _ = len(v5)
                                                        _ = v166
                                                    }
                                                    var v167 float32 = v6
                                                    _ = v167
                                                } else {
                                                    if (!v162) {
                                                        if (v3 || (v121 != (float64(v6) + float64((((((-2.25 + (v6 / float32(float64(3.14159)))) - 1.5) * ((((float32(v121) + (v6 + v6)) + (-2.25 + (v6 * (0.0 * 1.5)))) - 0.1) / v6)) * v6) - (((v6 + 0.1) * float32(-2.25)) - -2.25)))))) {
                                                            v168 := GenS13{F0: v123, F1: (v61 >> (((v62 ^ v62)) & 63)), F2: (v123 == 42)}
                                                            var v169 int16 = 32767
                                                            v2 = v2
                                                            var v170 string = (v2 + ((v2 + v2) + v1))
                                                            var v171 int16 = (v169 - (^(v169)))
                                                            _ = v168
                                                            _ = v169
                                                            _ = v170
                                                            _ = v171
                                                        } else {
                                                            v2 = v1
                                                        }
                                                        if v162 {
                                                            switch v2 {
                                                            case "":
                                                                v123 += v123
                                                            default:
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                                if false {
                                                                    v88 -= (v88 | -19)
                                                                    var v172 uint32 = uint32(v62)
                                                                    switch "a" {
                                                                    case "a":
                                                                        v173 := GenS9{F0: v62, F1: 255}
                                                                        _ = v173
                                                                    default:
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                    }
                                                                    _ = v172
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        if (-1 <= 1) {
                                                            var v174 uint = (^(v62))
                                                            switch (^(fi161)) {
                                                            case -9223372036854775808:
                                                                var v175 uint = v62
                                                                _ = v175
                                                            case 4294967295:
                                                                v61 += -2147483648
                                                                if (v1 > (v2 + v2)) {
                                                                    var v176 bool = v162
                                                                    _ = v176
                                                                }
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                            default:
                                                                switch v1 {
                                                                case "\n\t":
                                                                    var v177 uint32 = 2147483647
                                                                    _ = v177
                                                                case "":
                                                                    v178 := GenS9{F0: v62, F1: 65535}
                                                                    _ = v178
                                                                default:
                                                                    v179 := GenS1{F0: 0, F1: 128, F2: 42}
                                                                    _ = v179
                                                                }
                                                            }
                                                            v2 = "a"
                                                            fmt.Println(v62, v121)
                                                            _ = v174
                                                        } else {
                                                            v3 = v3
                                                            var v180 uint = 18446744073709551615
                                                            _ = v180
                                                        }
                                                        if false {
                                                            v1 = v2
                                                        }
                                                        v62 -= uint(v123)
                                                        fmt.Println(v121)
                                                        switch (v1 + v2) {
                                                        case "foo bar":
                                                            v121 += -2.25
                                                        default:
                                                            fmt.Println(fi154, v3, v1)
                                                        }
                                                        _ = len(v4)
                                                        v123 += 846
                                                        var v181 string = (v1 + v2)
                                                        v182 := GenS2{F0: v162, F1: 9223372036854775807}
                                                        _ = v181
                                                        _ = v182
                                                    }
                                                    fmt.Println(v164, v88)
                                                    var v183 int32 = 0
                                                    _ = v183
                                                }
                                                v184 := GenS14{F0: 607, F1: "\n\t"}
                                                _ = v184
                                            } else {
                                                var v185 uint8 = 128
                                                var v186 uint64 = 656
                                                var v187 float32 = ((0.0 * 0.0) - v6)
                                                _ = v185
                                                _ = v186
                                                _ = v187
                                            }
                                            _ = fi161
                                            _ = v162
                                            _ = v163
                                            _ = v164
                                            _ = v165
                                        }
                                    case "\n\t":
                                        _ = len(v5)
                                    default:
                                        _ = len(v4)
                                    }
                                }
                            case "foo bar":
                                _ = len(v4)
                            default:
                                v4[v1] = (fi154 ^ fi154)
                            }
                        }
                    } else {
                        v188 := GenS1{F0: 42, F1: 1, F2: 0}
                        v189 := GenS6{F0: -9223372036854775808, F1: (v123 & v123), F2: int16(fi154)}
                        _ = v188
                        _ = v189
                    }
                } else {
                    v62 += v62
                }
                _ = fi154
            }
        default:
            v190 := GenS11{F0: (true && ((v3 && (int64(v123) == 980)) || (v88 >= v88))), F1: 42}
            _ = v190
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v123 += (v123 * (^(v123)))
    fmt.Println("v1:", v1)
    fmt.Println("v121:", v121)
    fmt.Println("v123:", v123)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v119
    _ = v121
    _ = v122
    _ = v123
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v62
    _ = v88
}
