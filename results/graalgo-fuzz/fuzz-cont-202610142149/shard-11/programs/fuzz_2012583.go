package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int32; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int64; F1 int; F2 string }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS10 struct { F0 int64; F1 int32; F2 int16 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 int16; F1 bool }
type GenS13 struct { F0 int16; F1 int32 }
type GenS14 struct { F0 int16; F1 int32; F2 int16 }
type GenS15 struct { F0 int; F1 int; F2 uint }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 int8; F1 int64 }
type GenS18 struct { F0 string; F1 uint32 }
type GenS19 struct { F0 uint64; F1 string; F2 bool }
type GenS20 struct { F0 int32; F1 int32 }
type GenS21 struct { F0 uint16; F1 int64; F2 bool }
type GenS22 struct { F0 string; F1 uint64 }
type GenS23 struct { F0 uint8; F1 int16 }
type GenS24 struct { F0 bool }
type GenS25 struct { F0 int }
type GenS26 struct { F0 int64; F1 int16 }
type GenS27 struct { F0 bool; F1 int32; F2 int64 }
type GenS28 struct { F0 uint32; F1 uint64 }
type GenS29 struct { F0 int32; F1 bool }
type GenS30 struct { F0 int8 }
type GenS31 struct { F0 int }
type GenS32 struct { F0 uint64; F1 int; F2 int8 }
func genFn0() int64 {
    for fi7 := 0; fi7 < 6; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 uint16 = 907
            var v10 float32 = -2.25
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    if false {
        var v11 uint64 = 962
        _ = v11
    } else {
        var v12 uint8 = 127
        v12 = v12
        _ = v12
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        v14 := GenS1{F0: fi13, F1: 2, F2: 672}
        v15 := GenS1{F0: -318, F1: 191, F2: 164}
        _ = fi13
        _ = v14
        _ = v15
    }
    if true {
        var v16 string = ("" + ("hello" + "hello"))
        switch 0 {
        case 42:
            if false {
                fmt.Println(v16)
                if (-2147483648 == -774) {
                    if true {
                        for fi17 := 0; fi17 < 12; fi17++ {
                            if (true || false) {
                                v16 = "foo bar"
                                for fi18 := 0; fi18 < 14; fi18++ {
                                    if false {
                                        v16 = v16
                                        v19 := []int{}
                                        var v20 int32 = 1
                                        _ = v19
                                        _ = v20
                                    } else {
                                        if (!(668 < 9223372036854775807)) {
                                            var v21 int16 = -647
                                            switch (v16 + "hello") {
                                            case "":
                                                v22 := []int{}
                                                var v23 int8 = -1
                                                _ = v22
                                                _ = v23
                                            case "foo bar":
                                                if (!false) {
                                                    v16 = (v16 + (v16 + v16))
                                                }
                                            }
                                            if (-128 != 100) {
                                                v16 = (v16 + (v16 + ("\n\t" + v16)))
                                                v21 += v21
                                                var v24 uint8 = 0
                                                _ = v24
                                            }
                                            _ = v21
                                        } else {
                                            switch fi17 {
                                            case 127:
                                                var v25 int8 = 100
                                                _ = v25
                                            case 955:
                                                v26 := []int{132, -647, fi17}
                                                _ = v26
                                            case -467:
                                                v27 := []int{fi18, -383}
                                                _ = v27
                                            }
                                            v28 := []int{309, fi17, fi18}
                                            _ = v28
                                        }
                                        switch 4294967295 {
                                        case -1:
                                            fmt.Println(fi18, v16)
                                            if true {
                                                v16 = v16
                                                if (255 > 128) {
                                                    v29 := []int{-145, fi18}
                                                    _ = v29
                                                }
                                                switch v16 {
                                                case "hello":
                                                    if (32767 <= 468) {
                                                        if true {
                                                            var v30 int32 = -145
                                                            _ = v30
                                                        } else {
                                                            v16 = (v16 + (((v16 + (v16 + v16)) + v16) + "hello"))
                                                            fmt.Println(fi18, v16)
                                                            var v31 int16 = 1000
                                                            _ = v31
                                                        }
                                                        var v32 float64 = float64(3.14159)
                                                        _ = v32
                                                    }
                                                case "a":
                                                    v16 = v16
                                                }
                                            } else {
                                                v16 = "\n\t"
                                            }
                                            var v33 int32 = 1
                                            _ = v33
                                        case 93:
                                            var v34 int16 = 1000
                                            _ = v34
                                        default:
                                            switch (fi18 - -371) {
                                            case 4294967295:
                                                if (58 < 18446744073709551615) {
                                                    v16 = v16
                                                    v16 = v16
                                                    if true {
                                                        v16 = "foo bar"
                                                        var v35 int = fi18
                                                        _ = v35
                                                    }
                                                } else {
                                                    switch v16 {
                                                    case "foo bar":
                                                        var v36 uint32 = 42
                                                        v37 := GenS1{F0: (fi18 + -796), F1: 146, F2: 2}
                                                        var v38 float32 = (0.0 * 3.14159)
                                                        _ = v36
                                                        _ = v37
                                                        _ = v38
                                                    case "\n\t":
                                                        fmt.Println(v16, fi17)
                                                    default:
                                                        var v39 bool = true
                                                        _ = v39
                                                    }
                                                    if false {
                                                        var v40 int32 = 0
                                                        _ = v40
                                                    } else {
                                                        var v41 int16 = 1000
                                                        var v42 uint16 = 0
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                    v16 = ((v16 + v16) + "\n\t")
                                                    fmt.Println(fi18, fi18)
                                                }
                                            default:
                                                if false {
                                                    fmt.Println(fi18, fi17, v16)
                                                } else {
                                                    var v43 float32 = (1e10 / 0.1)
                                                    v44 := GenS1{F0: (fi18 ^ fi17), F1: 42, F2: 215}
                                                    v45 := []int{}
                                                    _ = v43
                                                    _ = v44
                                                    _ = v45
                                                }
                                                var v46 uint64 = 1
                                                fmt.Println(fi18, v46, fi17)
                                                _ = v46
                                            }
                                        }
                                        switch "\n\t" {
                                        case "a":
                                            if true {
                                                v47 := []int{(fi17 - fi17), fi17}
                                                var v48 string = ("\n\t" + "a")
                                                v16 = "foo bar"
                                                _ = v47
                                                _ = v48
                                            } else {
                                                var v49 uint = 4294967295
                                                if (!(true && false)) {
                                                    v50 := GenS1{F0: 1, F1: 2, F2: 32767}
                                                    _ = v50
                                                }
                                                var v51 uint32 = 42
                                                _ = v49
                                                _ = v51
                                            }
                                        }
                                        v16 = ("\n\t" + v16)
                                        v16 = v16
                                        switch 524 {
                                        case 4294967295:
                                            v16 = (v16 + v16)
                                            if true {
                                                switch "" {
                                                case "foo bar":
                                                    var v52 int = (-(fi17))
                                                    _ = v52
                                                }
                                            }
                                        case 32767:
                                            var v53 int16 = 255
                                            _ = v53
                                        case 1:
                                            v54 := GenS2{F0: -348, F1: 0, F2: 4294967295}
                                            _ = v54
                                        }
                                    }
                                    _ = fi18
                                }
                                var v55 bool = (2147483648 < 445)
                                _ = v55
                            } else {
                                v16 = v16
                            }
                            _ = fi17
                        }
                        for fi56 := 0; fi56 < 11; fi56++ {
                            if false {
                                v57 := []int{}
                                _ = v57
                            }
                            _ = fi56
                        }
                    } else {
                        for fi58 := 0; fi58 < 7; fi58++ {
                            switch -1 {
                            case 127:
                                fmt.Println(v16, fi58)
                            default:
                                var v59 int64 = 984
                                _ = v59
                            }
                            v16 = v16
                            v60 := []int{(-(fi58))}
                            _ = fi58
                            _ = v60
                        }
                        var v61 uint = 42
                        _ = v61
                    }
                }
                v16 = v16
            }
        case 1:
            var v62 uint16 = 42
            _ = v62
        case 32768:
            v63 := []int{-9223372036854775808, 4294967295, -2147483648}
            _ = v63
        default:
            v64 := []int{}
            _ = v64
        }
        _ = v16
    }
    return -1
}
func genFn1(p0 int8) int {
    fmt.Println(p0)
    p0 = 56
    for fi65 := 0; fi65 < 11; fi65++ {
        if (true && true) {
            var v66 float64 = 0.0
            if (42 == 236) {
                switch ("foo bar" + "") {
                case "hello":
                    v66 = v66
                default:
                    var v67 float32 = 0.0
                    _ = v67
                }
                var v68 int8 = p0
                _ = v68
            }
            switch "foo bar" {
            case "a":
                var v69 uint8 = 128
                _ = v69
            default:
                v66 += 1.5
            }
            _ = v66
        } else {
            for fi70 := 0; fi70 < 4; fi70++ {
                v71 := []int{(fi70 + fi70), (fi70 | fi65)}
                p0 -= p0
                _ = fi70
                _ = v71
            }
            switch "hello" {
            case "foo bar":
                var v72 float32 = (1e10 * ((0.0 - ((((((float32(0.0) / 1.5) - 0.1) - 0.0) / -2.25) / -2.25) - float32((0.0 - 0.1)))) * 1e10))
                _ = v72
            case "hello":
                v73 := GenS1{F0: fi65, F1: 682, F2: 1}
                _ = v73
            }
            v74 := []int{}
            _ = v74
        }
        if (!true) {
            for fi75 := 0; fi75 < 12; fi75++ {
                v76 := GenS3{F0: 840}
                p0 = (p0 << ((1) & 63))
                v77 := GenS4{F0: 18446744073709551615}
                v78 := GenS4{F0: 4294967295}
                v79 := []int{(fi75 % ((fi75) | 1)), (fi75 << ((2) & 63)), fi65}
                _ = fi75
                _ = v76
                _ = v77
                _ = v78
                _ = v79
            }
            p0 = (p0 << ((255) & 63))
            if true {
                switch "a" {
                case "":
                    v80 := GenS1{F0: fi65, F1: 2, F2: 32767}
                    _ = v80
                }
                var v81 float32 = 0.1
                _ = v81
            }
        } else {
            if true {
                v82 := []int{(-(fi65)), fi65}
                _ = v82
            } else {
                v83 := GenS5{F0: 170}
                _ = v83
            }
        }
        _ = fi65
    }
    var v84 float32 = float32(-2.25)
    _ = p0
    _ = v84
    return -63
}

func main() {
    var v1 int32 = 39
    v2 := []uint32{4294967295, 1}
    var v3 uint = 1
    v4 := []int{-280, int(v3), -448}
    var v5 int8 = -128
    v6 := []string{"", "a", "hello", ""}
    if (!(false || (4294967295 <= (v3 ^ (v3 | v3))))) {
        switch "a" {
        case "a":
            for fi85 := 0; fi85 < 9; fi85++ {
                v86 := GenS2{F0: v1, F1: v1, F2: 513}
                _ = fi85
                _ = v86
            }
        }
    } else {
        for fi87 := 0; fi87 < 17; fi87++ {
            fmt.Println(fi87)
            v88 := GenS2{F0: v1, F1: int32(fi87), F2: 0}
            for fi89 := 0; fi89 < 17; fi89++ {
                if (!(0.1 == ((((float32(3.14159) + (((1e10 * 1e10) + 3.14159) / 0.1)) * (-2.25 * ((0.1 + 1e10) * 1e10))) / (1e10 / (1e10 * 3.14159))) - -2.25))) {
                    switch fi89 {
                    case 0:
                        v90 := GenS6{F0: (v5 << ((v3) & 63))}
                        _ = v90
                    case 32767:
                        if false {
                            if (!(v5 != (v5 / ((42) | 1)))) {
                                v6 = append(v6, "foo bar")
                            }
                            v2 = append(v2, 4294967295)
                            v1 += (v1 / (((v1 ^ v1)) | 1))
                            v91 := GenS5{F0: 128}
                            _ = v91
                        }
                    default:
                        if false {
                            var v92 int32 = -541
                            var v93 float32 = ((((0.1 - 3.14159) - 0.1) * (1.5 + (1e10 - 1e10))) + 1e10)
                            if false {
                                v4 = append(v4, fi89)
                                v93 += 0.0
                                switch (fi89 - genFn1(126)) {
                                case 9223372036854775807:
                                    v5 = (-(v5))
                                case -503:
                                    if true {
                                        var v94 int16 = 0
                                        _ = v94
                                    } else {
                                        v92 = -1
                                        fmt.Println(fi87, v93, v1)
                                        fmt.Println(v3, fi87, v3)
                                    }
                                }
                            } else {
                                var v95 int64 = -9223372036854775808
                                if (520 <= 369) {
                                    if false {
                                        if len(v2) > 0 {
                                            v2[0] = 65535
                                        }
                                        switch 1 {
                                        case -2147483648:
                                            v92 = (^(v92))
                                        default:
                                            v96 := GenS6{F0: (v5 << ((v3) & 63))}
                                            v1 -= (v1 | (v1 & v1))
                                            fmt.Println(fi87, v92)
                                            v97 := GenS7{F0: (v95 / ((v95) | 1)), F1: fi89, F2: ""}
                                            var v98 int = fi89
                                            _ = v96
                                            _ = v97
                                            _ = v98
                                        }
                                    }
                                    var v99 uint32 = 614
                                    _ = v99
                                }
                                if len(v4) > 0 {
                                    v4[0] = (fi87 / ((32767) | 1))
                                }
                                _ = v95
                            }
                            _ = v92
                            _ = v93
                        } else {
                            v3 += (v3 << (((v3 / ((v3) | 1))) & 63))
                            switch "" {
                            case "a":
                                v4 = append(v4, (^(fi89)))
                            case "":
                                v100 := GenS8{F0: 4294967295}
                                _ = v100
                            case "\n\t":
                                var v101 int32 = -1
                                _ = v101
                            default:
                                if len(v2) > 0 {
                                    v2[0] = 1
                                }
                            }
                        }
                    }
                    v102 := GenS9{F0: 714, F1: 65535, F2: 42}
                    _ = v102
                } else {
                    if false {
                        v103 := GenS10{F0: genFn0(), F1: v1, F2: -32768}
                        switch 0 {
                        case -1:
                            if false {
                                v3 -= (v3 >> ((v3) & 63))
                                v104 := GenS11{F0: -222}
                                var v105 int = fi87
                                _ = v104
                                _ = v105
                            }
                        case -449:
                            switch "hello" {
                            case "a":
                                switch (("hello" + "a") + "a") {
                                case "":
                                    var v106 uint = 790
                                    _ = v106
                                default:
                                    fmt.Println(fi87, fi89, fi87)
                                }
                            case "hello":
                                var v107 int64 = -824
                                _ = v107
                            }
                        default:
                            v3 = (v3 / ((v3) | 1))
                        }
                        _ = v103
                    }
                    v5 += -113
                }
                v1 = 226
                if len(v2) > 0 {
                    v2[0] = 4294967295
                }
                _ = fi89
            }
            _ = fi87
            _ = v88
        }
        v1 = int32(v5)
    }
    if ("a" <= "foo bar") {
        for fi108 := 0; fi108 < 11; fi108++ {
            switch "foo bar" {
            case "":
                v1 = (v1 % (((v1 + (^(v1)))) | 1))
            case "foo bar":
                if true {
                    if false {
                        for fi109 := 0; fi109 < 20; fi109++ {
                            switch "" {
                            case "":
                                if (1 <= v5) {
                                    if false {
                                        v1 = -2147483648
                                        if true {
                                            fmt.Println(fi108, v3, v1)
                                        }
                                    }
                                    v5 = (v5 ^ (v5 / (((v5 & (^(v5)))) | 1)))
                                }
                            default:
                                if false {
                                    if (v3 != (v3 ^ v3)) {
                                        var v110 uint16 = 8
                                        if false {
                                            fmt.Println(v1)
                                            switch (fi108 + fi109) {
                                            case -285:
                                                var v111 float64 = (((1.5 / (1.5 * 1e10)) + 1e10) - (-2.25 - 1.5))
                                                _ = v111
                                            case 32768:
                                                if (!true) {
                                                    v5 = (v5 ^ (v5 % ((31) | 1)))
                                                } else {
                                                    v112 := GenS4{F0: 42}
                                                    _ = len(v6)
                                                    _ = v112
                                                }
                                            default:
                                                v5 = -1
                                            }
                                            v113 := GenS12{F0: int16(v5), F1: true}
                                            v114 := GenS7{F0: -1, F1: (fi109 << ((65535) & 63)), F2: (("hello" + "foo bar") + "a")}
                                            _ = v113
                                            _ = v114
                                        }
                                        var v115 int16 = 42
                                        _ = v110
                                        _ = v115
                                    } else {
                                        v116 := GenS13{F0: int16(v1), F1: v1}
                                        if (((1e10 / (float32(0.1) - 1.5)) + ((float32((-2.25 - 3.14159)) * (float32(1.5) - -2.25)) / (1.5 + -2.25))) == 0.0) {
                                            var v117 uint32 = 2
                                            _ = v117
                                        }
                                        _ = v116
                                    }
                                    switch fi109 {
                                    case 32767:
                                        var v118 uint32 = 575
                                        _ = v118
                                    case 1:
                                        v119 := GenS2{F0: 926, F1: v1, F2: 42}
                                        _ = v119
                                    }
                                    var v120 uint32 = 2
                                    _ = v120
                                }
                            }
                            if len(v6) > 0 {
                                v6[0] = "hello"
                            }
                            _ = fi109
                        }
                        if (uint32(v3) < 650) {
                            v121 := GenS14{F0: -32768, F1: 1, F2: -109}
                            v1 = v1
                            _ = v121
                        } else {
                            var v122 uint8 = 127
                            if false {
                                var v123 uint64 = 407
                                for fi124 := 0; fi124 < 10; fi124++ {
                                    if true {
                                        switch ("foo bar" + "foo bar") {
                                        case "\n\t":
                                            switch ("foo bar" + (("" + "a") + "hello")) {
                                            case "hello":
                                                var v125 uint8 = (v122 | v122)
                                                v122 = v125
                                                _ = v125
                                            }
                                        default:
                                            fmt.Println(v5, fi108)
                                        }
                                        var v126 int64 = 42
                                        _ = len(v4)
                                        _ = v126
                                    }
                                    var v127 uint16 = 218
                                    _ = fi124
                                    _ = v127
                                }
                                v128 := GenS15{F0: fi108, F1: int(v1), F2: (^(v3))}
                                _ = v123
                                _ = v128
                            } else {
                                var v129 uint16 = 255
                                var v130 int32 = (^(v1))
                                var v131 uint16 = v129
                                _ = v129
                                _ = v130
                                _ = v131
                            }
                            for fi132 := 0; fi132 < 3; fi132++ {
                                v133 := GenS14{F0: -672, F1: v1, F2: 1}
                                var v134 uint8 = (v122 | v122)
                                _ = len(v4)
                                _ = fi132
                                _ = v133
                                _ = v134
                            }
                            _ = v122
                        }
                        switch "\n\t" {
                        case "\n\t":
                            v1 += -960
                        case "a":
                            for fi135 := 0; fi135 < 3; fi135++ {
                                if (-631 >= -1) {
                                    switch ("\n\t" + "") {
                                    case "":
                                        fmt.Println(v5, fi108, v3)
                                    case "a":
                                        var v136 uint32 = 0
                                        _ = v136
                                    default:
                                        if (65535 >= 42) {
                                            var v137 uint32 = 622
                                            fmt.Println(fi135, v1)
                                            v1 += v1
                                            _ = v137
                                        } else {
                                            var v138 uint64 = 709
                                            _ = v138
                                        }
                                    }
                                    var v139 uint64 = 18446744073709551615
                                    v140 := GenS2{F0: 42, F1: v1, F2: (v139 / ((v139) | 1))}
                                    _ = v139
                                    _ = v140
                                }
                                _ = fi135
                            }
                        case "":
                            fmt.Println(fi108)
                        }
                    }
                    if (0.1 > (1e10 - (float32((1.5 + -2.25)) + (float32(1e10) * 1e10)))) {
                        v141 := GenS4{F0: 2}
                        switch genFn1(v5) {
                        case 2147483647:
                            if len(v2) > 0 {
                                v2[0] = 351
                            }
                        case 549:
                            for fi142 := 0; fi142 < 14; fi142++ {
                                var v143 int8 = -1
                                v144 := GenS16{F0: -9223372036854775808}
                                var v145 float64 = float64(0.1)
                                _ = fi142
                                _ = v143
                                _ = v144
                                _ = v145
                            }
                        case -9223372036854775808:
                            for fi146 := 0; fi146 < 15; fi146++ {
                                switch (-(fi146)) {
                                case 42:
                                    v5 += (v5 + -128)
                                default:
                                    if (v3 >= (v3 & v3)) {
                                        _ = len(v2)
                                    } else {
                                        v3 += v3
                                        var v147 int = 128
                                        if false {
                                            fmt.Println(fi146, v5, v1)
                                        } else {
                                            if true {
                                                if ((32767 == (^(fi146))) || (138 == 88)) {
                                                    fmt.Println(v1, fi108, v1)
                                                    var v148 uint8 = 255
                                                    var v149 int = (fi108 + fi146)
                                                    _ = v148
                                                    _ = v149
                                                } else {
                                                    _ = len(v2)
                                                    switch fi146 {
                                                    case -1:
                                                        var v150 float32 = ((1.5 / 3.14159) * (0.0 + 0.0))
                                                        _ = v150
                                                    }
                                                    switch fi146 {
                                                    case 128:
                                                        v3 += (v3 - (^(v3)))
                                                    case 32768:
                                                        switch (v147 - (fi108 >> ((629) & 63))) {
                                                        case 275:
                                                            v3 += v3
                                                        }
                                                    case -71:
                                                        var v151 bool = false
                                                        _ = v151
                                                    default:
                                                        if true {
                                                            v152 := GenS11{F0: 764}
                                                            _ = v152
                                                        } else {
                                                            var v153 int64 = 42
                                                            _ = v153
                                                        }
                                                    }
                                                }
                                                var v154 uint16 = 42
                                                if (1 < 9223372036854775807) {
                                                    v5 = v5
                                                    switch genFn1(v5) {
                                                    case -9223372036854775808:
                                                        if false {
                                                            _ = len(v2)
                                                            v4 = append(v4, 850)
                                                        }
                                                    case 42:
                                                        _ = len(v4)
                                                    }
                                                }
                                                var v155 float64 = (0.0 + 3.14159)
                                                fmt.Println(fi108, fi146, fi108)
                                                var v156 int16 = -32768
                                                v147 = (-(fi146))
                                                _ = v154
                                                _ = v155
                                                _ = v156
                                            }
                                            v2 = append(v2, 2)
                                        }
                                        _ = v147
                                    }
                                }
                                v157 := GenS17{F0: 100, F1: 583}
                                var v158 uint32 = 42
                                _ = fi146
                                _ = v157
                                _ = v158
                            }
                        }
                        if true {
                            fmt.Println(fi108, fi108, v3)
                        }
                        _ = v141
                    } else {
                        var v159 int16 = 1000
                        _ = v159
                    }
                    if false {
                        v5 -= v5
                        for fi160 := 0; fi160 < 16; fi160++ {
                            _ = len(v6)
                            var v161 int = (-(fi108))
                            _ = fi160
                            _ = v161
                        }
                    } else {
                        v5 += v5
                        switch ("foo bar" + "foo bar") {
                        case "foo bar":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case "hello":
                            if (!(false || false)) {
                                switch genFn1((-(v5))) {
                                case -397:
                                    for fi162 := 0; fi162 < 17; fi162++ {
                                        v1 -= v1
                                        var v163 uint16 = 255
                                        _ = fi162
                                        _ = v163
                                    }
                                }
                            }
                        case "":
                            switch "a" {
                            case "":
                                if false {
                                    if len(v2) > 0 {
                                        v2[0] = 0
                                    }
                                    v164 := GenS14{F0: 1, F1: (v1 | 1), F2: 1}
                                    _ = v164
                                } else {
                                    var v165 int = genFn1(v5)
                                    if len(v4) > 0 {
                                        v4[0] = genFn1((v5 << (((v3 >> ((65535) & 63))) & 63)))
                                    }
                                    switch ("a" + "a") {
                                    case "":
                                        v166 := GenS17{F0: v5, F1: -9223372036854775808}
                                        _ = v166
                                    default:
                                        v167 := GenS18{F0: "foo bar", F1: 0}
                                        _ = v167
                                    }
                                    _ = v165
                                }
                            }
                        default:
                            for fi168 := 0; fi168 < 17; fi168++ {
                                v169 := GenS10{F0: 1, F1: v1, F2: int16(v3)}
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                if (233 < 1) {
                                    _ = len(v4)
                                    v3 += v3
                                    var v170 int8 = v5
                                    _ = v170
                                }
                                _ = fi168
                                _ = v169
                            }
                        }
                    }
                }
            default:
                if true {
                    v6 = append(v6, "")
                    for fi171 := 0; fi171 < 6; fi171++ {
                        var v172 uint8 = 0
                        if len(v4) > 0 {
                            v4[0] = 128
                        }
                        v2 = append(v2, 2147483648)
                        _ = fi171
                        _ = v172
                    }
                } else {
                    switch fi108 {
                    case -1:
                        for fi173 := 0; fi173 < 3; fi173++ {
                            _ = len(v4)
                            _ = fi173
                        }
                    case 32767:
                        var v174 uint32 = 1
                        _ = v174
                    }
                }
            }
            v6 = append(v6, "foo bar")
            _ = fi108
        }
        if (42 <= 304) {
            _ = len(v6)
            v5 += (v5 % ((v5) | 1))
            fmt.Println(v5, v5)
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi175 := 0; fi175 < 8; fi175++ {
                v5 += v5
                fmt.Println(v1)
                v176 := GenS7{F0: 1, F1: fi175, F2: ("hello" + "\n\t")}
                _ = fi175
                _ = v176
            }
            v3 = v3
        }
    } else {
        var v177 uint16 = 1
        switch 32767 {
        case -801:
            var v178 uint64 = 0
            _ = v178
        default:
            fmt.Println(v3, v3, v1)
        }
        _ = v177
    }
    if true {
        _ = len(v4)
        fmt.Println(v3)
        v179 := GenS1{F0: -348, F1: 65535, F2: 0}
        _ = v179
    }
    if len(v6) > 0 {
        v6[0] = "hello"
    }
    v5 += 100
    v180 := GenS13{F0: -384, F1: v1}
    v6 = append(v6, "foo bar")
    for fi181 := 0; fi181 < 8; fi181++ {
        switch fi181 {
        case 127:
            v182 := GenS19{F0: 412, F1: "", F2: false}
            _ = v182
        case 353:
            fmt.Println(fi181)
        default:
            if true {
                for fi183 := 0; fi183 < 17; fi183++ {
                    if false {
                        fmt.Println(v5, fi181, v1)
                        switch "" {
                        case "hello":
                            fmt.Println(v3, fi181)
                        case "a":
                            _ = len(v2)
                            var v184 float64 = (((1e10 + (1.5 + ((float64(0.0) * float64(float32((0.1 + float64((-2.25 / 0.1)))))) - 1e10))) - (1e10 - -2.25)) * (float64(3.14159) - ((0.0 * 0.1) + 0.0)))
                            v185 := GenS20{F0: (v1 / ((int32(fi183)) | 1)), F1: int32(fi181)}
                            _ = v184
                            _ = v185
                        case "\n\t":
                            var v186 uint16 = 155
                            _ = v186
                        default:
                            var v187 int = (fi183 % ((-9223372036854775808) | 1))
                            _ = v187
                        }
                        switch fi183 {
                        case 2147483647:
                            var v188 float64 = -2.25
                            _ = v188
                        }
                    }
                    _ = fi183
                }
                v5 -= 127
                v189 := GenS21{F0: 32767, F1: 1, F2: (!true)}
                _ = v189
            }
        }
        v5 = v5
        for fi190 := 0; fi190 < 2; fi190++ {
            if len(v4) > 0 {
                v4[0] = fi181
            }
            if len(v2) > 0 {
                v2[0] = 10
            }
            _ = fi190
        }
        _ = fi181
    }
    v191 := GenS1{F0: 578, F1: 42, F2: 2}
    v192 := GenS13{F0: 0, F1: v1}
    v2 = append(v2, 2147483648)
    v193 := GenS7{F0: -1, F1: -2147483648, F2: "a"}
    v194 := GenS22{F0: "a", F1: 2}
    for fi195 := 0; fi195 < 13; fi195++ {
        switch (-(fi195)) {
        case 128:
            v2 = append(v2, 103)
        case 32768:
            if (int32(v5) < (v1 + v1)) {
                switch ("" + "a") {
                case "hello":
                    var v196 uint = (v3 >> ((v3) & 63))
                    _ = v196
                case "a":
                    for fi197 := 0; fi197 < 4; fi197++ {
                        var v198 string = (("\n\t" + "\n\t") + ("a" + "foo bar"))
                        var v199 int16 = 724
                        _ = fi197
                        _ = v198
                        _ = v199
                    }
                case "\n\t":
                    for fi200 := 0; fi200 < 3; fi200++ {
                        if len(v4) > 0 {
                            v4[0] = (fi195 << ((42) & 63))
                        }
                        _ = fi200
                    }
                default:
                    v201 := GenS23{F0: 128, F1: -32768}
                    _ = v201
                }
                var v202 uint64 = 307
                fmt.Println(fi195)
                _ = v202
            } else {
                if (2 != 65535) {
                    v1 += v1
                    switch genFn1((^(v5))) {
                    case 127:
                        v5 = 52
                    case 32768:
                        var v203 uint32 = 0
                        _ = v203
                    default:
                        fmt.Println(v3)
                    }
                } else {
                    v5 -= (v5 & int8(fi195))
                    v3 = 911
                    if (0 == 42) {
                        if (v3 > v3) {
                            var v204 uint32 = 0
                            v3 -= 65535
                            for fi205 := 0; fi205 < 8; fi205++ {
                                if false {
                                    var v206 int64 = 9223372036854775807
                                    var v207 uint8 = 141
                                    switch (fi195 % ((fi205) | 1)) {
                                    case 1:
                                        v208 := GenS24{F0: (!(v3 != (v3 >> (((v3 + (v3 >> ((1) & 63)))) & 63))))}
                                        v2 = append(v2, 4294967295)
                                        _ = v208
                                    }
                                    _ = v206
                                    _ = v207
                                } else {
                                    fmt.Println(v1)
                                    v204 = (v204 + v204)
                                    v1 += (-(v1))
                                }
                                _ = fi205
                            }
                            _ = v204
                        } else {
                            if false {
                                if (!(true || false)) {
                                    if false {
                                        switch (fi195 * 4294967295) {
                                        case 42:
                                            var v209 int = (-(fi195))
                                            _ = v209
                                        case 127:
                                            v210 := GenS25{F0: (fi195 / ((genFn1(100)) | 1))}
                                            _ = v210
                                        }
                                    } else {
                                        v211 := GenS26{F0: 42, F1: 0}
                                        if true {
                                            switch ("" + ("foo bar" + (("" + (("a" + "") + "\n\t")) + "hello"))) {
                                            case "":
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            default:
                                                var v212 uint64 = 2
                                                _ = v212
                                            }
                                        } else {
                                            if (709 > 4294967295) {
                                                v6 = append(v6, ("a" + "hello"))
                                                v1 += (v1 / ((v1) | 1))
                                            }
                                        }
                                        if len(v6) > 0 {
                                            v6[0] = "hello"
                                        }
                                        _ = v211
                                    }
                                } else {
                                    v213 := GenS10{F0: 768, F1: v1, F2: 1000}
                                    v214 := GenS10{F0: -1, F1: v1, F2: 1000}
                                    if false {
                                        if false {
                                            if ((genFn0() >= 9223372036854775807) || (220 == 42)) {
                                                v1 += v1
                                            }
                                        }
                                        v5 = (v5 * 1)
                                    } else {
                                        fmt.Println(v5, v5, v5)
                                    }
                                    _ = v213
                                    _ = v214
                                }
                            } else {
                                if (1 >= 42) {
                                    var v215 string = "a"
                                    fmt.Println(v5)
                                    _ = v215
                                }
                                if true {
                                    var v216 string = "hello"
                                    v1 -= (v1 % (((v1 / ((0) | 1))) | 1))
                                    v1 = (v1 >> ((v3) & 63))
                                    _ = v216
                                } else {
                                    if ((-1 < -1) || (!false)) {
                                        switch "hello" {
                                        case "":
                                            if len(v2) > 0 {
                                                v2[0] = 1
                                            }
                                        case "hello":
                                            var v217 uint32 = 792
                                            _ = v217
                                        case "foo bar":
                                            for fi218 := 0; fi218 < 14; fi218++ {
                                                var v219 int = 885
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                var v220 int32 = (v1 ^ v1)
                                                if (float32(-2.25) < 3.14159) {
                                                    v221 := GenS17{F0: 127, F1: 1}
                                                    v1 -= -2147483648
                                                    var v222 uint = (v3 & 65535)
                                                    _ = v221
                                                    _ = v222
                                                } else {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    if false {
                                                        v1 += 1
                                                    } else {
                                                        if false {
                                                            v223 := GenS10{F0: int64(v3), F1: (v1 + (v220 - 1)), F2: -733}
                                                            v1 -= (v1 * (v220 / ((v1) | 1)))
                                                            var v224 int16 = 0
                                                            _ = v223
                                                            _ = v224
                                                        }
                                                    }
                                                }
                                                if (true || false) {
                                                    var v225 string = "hello"
                                                    switch ("foo bar" + v225) {
                                                    case "\n\t":
                                                        fmt.Println(v5)
                                                    }
                                                    v1 -= 42
                                                    if false {
                                                        fmt.Println(fi195)
                                                        if (v3 < 18446744073709551615) {
                                                            v226 := GenS27{F0: false, F1: 42, F2: 9223372036854775807}
                                                            v227 := GenS28{F0: 4294967295, F1: 42}
                                                            var v228 uint64 = 1
                                                            _ = v226
                                                            _ = v227
                                                            _ = v228
                                                        } else {
                                                            fmt.Println(fi218)
                                                        }
                                                        switch v225 {
                                                        case "\n\t":
                                                            v220 = v220
                                                        case "foo bar":
                                                            var v229 int32 = v1
                                                            _ = v229
                                                        default:
                                                            v225 = ("a" + v225)
                                                        }
                                                    } else {
                                                        v3 += v3
                                                        v219 += -638
                                                        var v230 uint64 = 18446744073709551615
                                                        _ = v230
                                                    }
                                                    _ = v225
                                                } else {
                                                    var v231 uint8 = 127
                                                    if true {
                                                        if true {
                                                            var v232 uint16 = 65535
                                                            v233 := GenS3{F0: 2}
                                                            _ = v232
                                                            _ = v233
                                                        }
                                                    } else {
                                                        if false {
                                                            v1 += (v1 * (v1 & 0))
                                                            v234 := GenS26{F0: 570, F1: -32768}
                                                            _ = v234
                                                        } else {
                                                            v235 := GenS11{F0: 0}
                                                            v236 := GenS2{F0: int32(v231), F1: v1, F2: 42}
                                                            var v237 float32 = float32(0.0)
                                                            _ = v235
                                                            _ = v236
                                                            _ = v237
                                                        }
                                                        v2 = append(v2, 0)
                                                    }
                                                    if (!true) {
                                                        var v238 uint64 = 0
                                                        if false {
                                                            v239 := GenS19{F0: (v238 ^ 615), F1: ("\n\t" + ("foo bar" + "a")), F2: false}
                                                            if true {
                                                                switch fi218 {
                                                                case 929:
                                                                    v220 = (-(v220))
                                                                case 785:
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        var v240 bool = (v3 < (^(v3)))
                                                                        _ = v240
                                                                    default:
                                                                        var v241 uint = 18446744073709551615
                                                                        _ = v241
                                                                    }
                                                                default:
                                                                    if true {
                                                                        fmt.Println(v3, v5, fi218)
                                                                        if false {
                                                                            v242 := GenS9{F0: 4294967295, F1: 4294967295, F2: -9223372036854775808}
                                                                            v243 := GenS17{F0: 1, F1: genFn0()}
                                                                            _ = v242
                                                                            _ = v243
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            _ = v239
                                                        }
                                                        var v244 int8 = (v5 - v5)
                                                        if ((!true) && true) {
                                                            switch "" {
                                                            case "":
                                                                v238 -= 18446744073709551615
                                                            default:
                                                                if (-737 >= int64(v5)) {
                                                                    if false {
                                                                        var v245 string = ("\n\t" + "\n\t")
                                                                        _ = v245
                                                                    }
                                                                    var v246 bool = true
                                                                    fmt.Println(fi218, v231)
                                                                    fmt.Println(v246)
                                                                    _ = v246
                                                                } else {
                                                                    fmt.Println(v1, v5, v244)
                                                                    var v247 int = v219
                                                                    _ = v247
                                                                }
                                                            }
                                                        } else {
                                                            var v248 int16 = 0
                                                            v4 = append(v4, 1)
                                                            _ = v248
                                                        }
                                                        if true {
                                                            _ = len(v2)
                                                            if (193 < -626) {
                                                                if true {
                                                                    var v249 int64 = 1
                                                                    _ = v249
                                                                } else {
                                                                    if (((876 > 1000) && ((v3 ^ v3) >= (v3 + v3))) && false) {
                                                                        v250 := GenS29{F0: (-(v220)), F1: true}
                                                                        if (v3 < 2) {
                                                                            switch genFn1(v5) {
                                                                            case 32768:
                                                                                switch 4294967295 {
                                                                                case 75:
                                                                                    var v251 string = ("a" + "a")
                                                                                    _ = v251
                                                                                }
                                                                                if (256 > 512) {
                                                                                    v244 = (v244 | 100)
                                                                                    if false {
                                                                                        switch "hello" {
                                                                                        case "hello":
                                                                                            v231 += v231
                                                                                            if (true || (false || false)) {
                                                                                                v252 := GenS13{F0: 0, F1: v220}
                                                                                                v1 = (v1 / (((v1 | (v1 + v220))) | 1))
                                                                                                var v253 int8 = v5
                                                                                                _ = v252
                                                                                                _ = v253
                                                                                            }
                                                                                        case "foo bar":
                                                                                            v220 += (v220 >> ((v3) & 63))
                                                                                        }
                                                                                        _ = len(v6)
                                                                                    }
                                                                                }
                                                                                var v254 string = "a"
                                                                                var v255 bool = false
                                                                                _ = v254
                                                                                _ = v255
                                                                            case -9223372036854775808:
                                                                                if (9223372036854775807 >= 119) {
                                                                                    v256 := GenS2{F0: (-(v1)), F1: v220, F2: v238}
                                                                                    _ = v256
                                                                                }
                                                                            case 184:
                                                                                v257 := GenS22{F0: (("" + "foo bar") + ("a" + "foo bar")), F1: v238}
                                                                                _ = v257
                                                                            }
                                                                            v258 := GenS30{F0: (^(v5))}
                                                                            _ = v258
                                                                        } else {
                                                                            if len(v6) > 0 {
                                                                                v6[0] = "\n\t"
                                                                            }
                                                                        }
                                                                        _ = v250
                                                                    } else {
                                                                        v219 = v219
                                                                        v244 = int8(v238)
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            fmt.Println(v244)
                                                                        }
                                                                    }
                                                                    if (v5 == (v5 - 100)) {
                                                                        switch (^(fi218)) {
                                                                        case 42:
                                                                            fmt.Println(v5)
                                                                        case 667:
                                                                            switch "\n\t" {
                                                                            case "":
                                                                                if len(v6) > 0 {
                                                                                    v6[0] = "\n\t"
                                                                                }
                                                                            case "\n\t":
                                                                                var v259 int64 = int64(v231)
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = 493
                                                                                }
                                                                                _ = v259
                                                                            case "a":
                                                                                var v260 uint = v3
                                                                                _ = v260
                                                                            default:
                                                                                fmt.Println(v244)
                                                                            }
                                                                        default:
                                                                            if false {
                                                                                v220 = (v1 | (^(v1)))
                                                                            } else {
                                                                                if len(v6) >= 2 {
                                                                                    v6 = v6[1:len(v6)]
                                                                                }
                                                                            }
                                                                        }
                                                                        var v261 uint64 = (v238 ^ 2)
                                                                        _ = v261
                                                                    } else {
                                                                        v2 = append(v2, 238)
                                                                    }
                                                                    if (-59 == 42) {
                                                                        v238 = uint64(v231)
                                                                        fmt.Println(v1)
                                                                        if ((v3 | 0) <= v3) {
                                                                            if false {
                                                                                v244 = v244
                                                                            }
                                                                            var v262 float64 = 0.0
                                                                            _ = v262
                                                                        } else {
                                                                            if true {
                                                                                if (3.14159 > 0.1) {
                                                                                    var v263 uint8 = 127
                                                                                    v244 = -30
                                                                                    switch "hello" {
                                                                                    case "":
                                                                                        v4 = append(v4, fi218)
                                                                                    case "a":
                                                                                        v264 := GenS18{F0: "foo bar", F1: 1}
                                                                                        _ = v264
                                                                                    case "hello":
                                                                                        v265 := GenS20{F0: v1, F1: (v220 & (v1 << ((v3) & 63)))}
                                                                                        _ = v265
                                                                                    default:
                                                                                        switch "a" {
                                                                                        case "a":
                                                                                            switch "a" {
                                                                                            case "a":
                                                                                                var v266 string = "hello"
                                                                                                _ = v266
                                                                                            }
                                                                                        default:
                                                                                            v267 := GenS16{F0: -838}
                                                                                            _ = v267
                                                                                        }
                                                                                    }
                                                                                    _ = v263
                                                                                } else {
                                                                                    var v268 float64 = 0.0
                                                                                    _ = v268
                                                                                }
                                                                                var v269 uint8 = (v231 | (v231 - (^(v231))))
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = 42
                                                                                }
                                                                                _ = v269
                                                                            } else {
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = genFn1(int8(v3))
                                                                                }
                                                                                var v270 uint = v3
                                                                                _ = v270
                                                                            }
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            if ((255 < 65535) || false) {
                                                                                var v271 string = "a"
                                                                                _ = len(v2)
                                                                                _ = v271
                                                                            } else {
                                                                                fmt.Println(fi195, v220, v5)
                                                                                v3 += 1
                                                                                v272 := GenS16{F0: -9223372036854775808}
                                                                                _ = v272
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                v273 := GenS31{F0: genFn1((^(v244)))}
                                                                if len(v2) > 0 {
                                                                    v2[0] = 2
                                                                }
                                                                _ = v273
                                                            }
                                                            v3 -= 42
                                                            if false {
                                                                fmt.Println(v238, v1)
                                                                if false {
                                                                    if false {
                                                                        v231 -= v231
                                                                    }
                                                                    v2 = append(v2, 2)
                                                                    v5 -= v244
                                                                    var v274 uint = 739
                                                                    switch (("hello" + "a") + ("a" + "hello")) {
                                                                    case "hello":
                                                                        var v275 uint16 = 65535
                                                                        var v276 uint64 = (v238 << ((v3) & 63))
                                                                        _ = v275
                                                                        _ = v276
                                                                    default:
                                                                        _ = len(v4)
                                                                    }
                                                                    switch "" {
                                                                    case "":
                                                                        if (!true) {
                                                                            if (true || false) {
                                                                                v244 += (v5 + -4)
                                                                                if true {
                                                                                    var v277 int8 = v244
                                                                                    v244 -= (^(v5))
                                                                                    if true {
                                                                                        var v278 int64 = -9223372036854775808
                                                                                        var v279 int64 = (^(v278))
                                                                                        v238 = v238
                                                                                        _ = v278
                                                                                        _ = v279
                                                                                    } else {
                                                                                        var v280 uint8 = v231
                                                                                        switch "a" {
                                                                                        case "a":
                                                                                            fmt.Println(v3, fi218, fi195)
                                                                                        }
                                                                                        switch "\n\t" {
                                                                                        case "a":
                                                                                            v280 -= (v280 * v231)
                                                                                        case "hello":
                                                                                            var v281 float64 = 1.5
                                                                                            _ = v281
                                                                                        case "":
                                                                                            v3 = (v3 / ((65535) | 1))
                                                                                        }
                                                                                        _ = v280
                                                                                    }
                                                                                    _ = v277
                                                                                } else {
                                                                                    var v282 int16 = -199
                                                                                    var v283 uint8 = v231
                                                                                    _ = v282
                                                                                    _ = v283
                                                                                }
                                                                            }
                                                                        }
                                                                        if false {
                                                                            if true {
                                                                                if (0 == -9223372036854775808) {
                                                                                    v1 = v220
                                                                                    var v284 uint64 = (^(v238))
                                                                                    _ = v284
                                                                                }
                                                                                var v285 int = (-(v219))
                                                                                v285 = (fi195 & v219)
                                                                                var v286 bool = ((!false) || true)
                                                                                if len(v6) > 0 {
                                                                                    v6[0] = "foo bar"
                                                                                }
                                                                                _ = v285
                                                                                _ = v286
                                                                            }
                                                                            var v287 bool = (genFn1(1) == int(v274))
                                                                            var v288 string = "foo bar"
                                                                            _ = v287
                                                                            _ = v288
                                                                        }
                                                                        fmt.Println(v1, v219)
                                                                    case "a":
                                                                        v1 += 949
                                                                    case "foo bar":
                                                                        var v289 int = (v219 - 1)
                                                                        _ = v289
                                                                    }
                                                                    _ = v274
                                                                }
                                                            }
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            var v290 uint = (v3 * (v3 & (v3 & v3)))
                                                            _ = v290
                                                        } else {
                                                            var v291 float32 = (float32(float64(-2.25)) - float32(-2.25))
                                                            _ = v291
                                                        }
                                                        if (!(false || false)) {
                                                            var v292 float32 = -2.25
                                                            var v293 uint8 = v231
                                                            _ = v292
                                                            _ = v293
                                                        } else {
                                                            var v294 uint = 2
                                                            var v295 uint = 18446744073709551615
                                                            _ = v294
                                                            _ = v295
                                                        }
                                                        switch fi195 {
                                                        case 127:
                                                            v1 += (^(v1))
                                                        case -533:
                                                            v220 = v1
                                                        }
                                                        _ = v238
                                                        _ = v244
                                                    } else {
                                                        v2 = append(v2, 65535)
                                                        v1 = (v220 ^ (v220 % ((int32(v3)) | 1)))
                                                    }
                                                    _ = v231
                                                }
                                                var v296 string = ("\n\t" + "hello")
                                                var v297 float64 = 0.0
                                                _ = fi218
                                                _ = v219
                                                _ = v220
                                                _ = v296
                                                _ = v297
                                            }
                                        default:
                                            var v298 uint32 = 379
                                            _ = v298
                                        }
                                        if (true && (uint64(v3) < 342)) {
                                            if len(v4) > 0 {
                                                v4[0] = 59
                                            }
                                            v3 = 317
                                        } else {
                                            var v299 uint16 = 255
                                            fmt.Println(v299)
                                            fmt.Println(v5)
                                            _ = v299
                                        }
                                    }
                                    v3 += (v3 ^ (v3 % (((^(v3))) | 1)))
                                }
                                v3 -= v3
                            }
                            for fi300 := 0; fi300 < 7; fi300++ {
                                v301 := GenS21{F0: 853, F1: 9223372036854775807, F2: false}
                                fmt.Println(fi300, v5, v3)
                                v3 = v3
                                _ = fi300
                                _ = v301
                            }
                            v5 = v5
                        }
                        var v302 int64 = -1
                        _ = v302
                    }
                }
                v303 := GenS7{F0: 9223372036854775807, F1: int(v3), F2: "a"}
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = v303
            }
        default:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        if true {
            v304 := GenS32{F0: 947, F1: (fi195 % ((fi195) | 1)), F2: (^(v5))}
            for fi305 := 0; fi305 < 5; fi305++ {
                var v306 uint = (^(v3))
                v1 = (v1 ^ (v1 ^ (v1 ^ v1)))
                _ = len(v6)
                v1 -= v1
                _ = fi305
                _ = v306
            }
            var v307 uint32 = 870
            _ = v304
            _ = v307
        }
        v5 = v5
        _ = fi195
    }
    v3 = (v3 ^ (^(v3)))
    var v308 int = 21
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v308:", v308)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v180
    _ = v191
    _ = v192
    _ = v193
    _ = v194
    _ = v2
    _ = v3
    _ = v308
    _ = v4
    _ = v5
    _ = v6
}
