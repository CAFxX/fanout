package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16; F1 uint64; F2 uint }
type GenS6 struct { F0 int16; F1 int8; F2 string }
type GenS7 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS8 struct { F0 string; F1 uint16; F2 int16 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int; F1 string }
type GenS12 struct { F0 uint8; F1 uint8; F2 uint }
type GenS13 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 uint32; F1 int64; F2 uint16 }
func genFn0() bool {
    if false {
        for fi6 := 0; fi6 < 10; fi6++ {
            var v7 uint64 = 645
            for fi8 := 0; fi8 < 7; fi8++ {
                var v9 float64 = 0.1
                switch "a" {
                case "foo bar":
                    var v10 int = (fi8 + fi8)
                    _ = v10
                }
                _ = fi8
                _ = v9
            }
            for fi11 := 0; fi11 < 11; fi11++ {
                v7 = v7
                v12 := []int{fi6}
                v7 = v7
                _ = fi11
                _ = v12
            }
            _ = fi6
            _ = v7
        }
        switch 128 {
        case 161:
            for fi13 := 0; fi13 < 4; fi13++ {
                var v14 int32 = -1
                _ = fi13
                _ = v14
            }
        case 0:
            v15 := []int{127, 1, 2147483647}
            _ = v15
        case 4294967295:
            var v16 uint32 = 0
            _ = v16
        default:
            v17 := GenS1{F0: 1000, F1: 559}
            _ = v17
        }
    }
    v18 := []int{}
    var v19 float64 = (1.5 + 0.0)
    for fi20 := 0; fi20 < 12; fi20++ {
        v19 = v19
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 int16 = 503
            v23 := GenS2{F0: 2, F1: 2}
            var v24 string = "\n\t"
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        v19 = ((0.1 / float64(1e10)) - 3.14159)
        _ = fi20
    }
    _ = v18
    _ = v19
    return false
}
func genFn1(p0 string, p1 bool) uint64 {
    v25 := GenS2{F0: 533, F1: 2}
    var v26 uint = 540
    p0 = p0
    _ = p0
    _ = p1
    _ = v25
    _ = v26
    return 1
}

func main() {
    var v1 int16 = 1000
    v2 := make(map[int]uint32)
    var v3 uint16 = 455
    v4 := make(map[uint32]uint32)
    var v5 int = 9223372036854775807
    v5 = v5
    if true {
        var v27 float32 = (((((1e10 + 3.14159) + 1e10) * (0.1 / float32(1.5))) - (((1.5 * (0.0 - float32(float64(float32(float64(float32(1e10))))))) - 1.5) + 3.14159)) * 1e10)
        switch v5 {
        case 128:
            if false {
                switch ("\n\t" + "\n\t") {
                case "":
                    for fi28 := 0; fi28 < 9; fi28++ {
                        switch "" {
                        case "\n\t":
                            switch 128 {
                            case 0:
                                if true {
                                    switch "foo bar" {
                                    case "\n\t":
                                        v29 := GenS1{F0: v1, F1: -9223372036854775808}
                                        _ = v29
                                    case "a":
                                        var v30 uint64 = 0
                                        _ = v30
                                    case "foo bar":
                                        for fi31 := 0; fi31 < 8; fi31++ {
                                            switch (v5 & (^(fi28))) {
                                            case 42:
                                                v32 := []int{v5, v5}
                                                _ = v32
                                            case 9223372036854775807:
                                                var v33 string = ("a" + "\n\t")
                                                _ = v33
                                            default:
                                                var v34 uint64 = genFn1((("foo bar" + "foo bar") + "foo bar"), (((-(fi28)) <= fi31) && (!(true && ((false && (true || ((2 == 645) || true))) && true)))))
                                                var v35 int16 = (v1 << ((65535) & 63))
                                                v1 -= 354
                                                if (((670 != 1) || (!(121 > 100))) && false) {
                                                    v36 := GenS2{F0: 1, F1: (^(v34))}
                                                    v2[fi31] = 1
                                                    var v37 float64 = 1.5
                                                    _ = v36
                                                    _ = v37
                                                } else {
                                                    fmt.Println(fi31)
                                                }
                                                _ = v34
                                                _ = v35
                                            }
                                            var v38 int16 = (v1 * v1)
                                            var v39 string = "foo bar"
                                            _ = fi31
                                            _ = v38
                                            _ = v39
                                        }
                                    default:
                                        v40 := []int{fi28}
                                        _ = v40
                                    }
                                }
                            case -9223372036854775808:
                                var v41 bool = (false || false)
                                _ = v41
                            case 515:
                                if true {
                                    v3 += (v3 * 1)
                                }
                            default:
                                switch (fi28 ^ (fi28 / (((v5 << ((18446744073709551615) & 63))) | 1))) {
                                case -112:
                                    v42 := []int{32767, (fi28 + v5), 2147483647, fi28}
                                    _ = v42
                                default:
                                    v43 := []int{fi28, (v5 | (v5 / ((-2147483648) | 1)))}
                                    _ = v43
                                }
                            }
                        }
                        v5 = v5
                        for fi44 := 0; fi44 < 15; fi44++ {
                            v45 := GenS2{F0: 65535, F1: 2}
                            v5 = (fi44 + (fi44 << ((0) & 63)))
                            if false {
                                v1 -= -1
                                _ = len(v2)
                                switch "\n\t" {
                                case "":
                                    v27 = -2.25
                                    var v46 int = fi28
                                    _ = v46
                                case "foo bar":
                                    if (79 != 4294967295) {
                                        v47 := GenS3{F0: false}
                                        if (-124 >= 0) {
                                            var v48 bool = (239 == 153)
                                            switch fi44 {
                                            case 32768:
                                                v49 := GenS3{F0: v48}
                                                _ = v49
                                            case 172:
                                                fmt.Println(fi28, fi44)
                                            case 4294967295:
                                                v50 := GenS2{F0: 0, F1: 1}
                                                _ = v50
                                            }
                                            _ = v48
                                        } else {
                                            switch fi44 {
                                            case 32768:
                                                if (!false) {
                                                    v51 := GenS4{F0: (v1 * 32767)}
                                                    v52 := GenS5{F0: 32767, F1: uint64(v3), F2: 1}
                                                    var v53 int64 = 1
                                                    _ = v51
                                                    _ = v52
                                                    _ = v53
                                                } else {
                                                    if false {
                                                        var v54 uint = 18446744073709551615
                                                        if (42 <= -2147483648) {
                                                            delete(v2, (fi28 * 128))
                                                            v1 = 1
                                                            v5 = fi44
                                                        } else {
                                                            v55 := GenS2{F0: 2147483647, F1: 4294967295}
                                                            var v56 uint16 = v3
                                                            v57 := []int{fi28, (fi44 % ((fi44) | 1)), (fi44 | 42), (fi44 + 4294967295)}
                                                            _ = v55
                                                            _ = v56
                                                            _ = v57
                                                        }
                                                        var v58 uint32 = 65535
                                                        _ = v54
                                                        _ = v58
                                                    }
                                                    var v59 int = (fi44 + -509)
                                                    if (0 == -1) {
                                                        var v60 uint32 = 0
                                                        v61 := GenS6{F0: v1, F1: 42, F2: "a"}
                                                        _ = v60
                                                        _ = v61
                                                    } else {
                                                        fmt.Println(v5, v3)
                                                    }
                                                    var v62 uint = 255
                                                    var v63 string = ""
                                                    switch fi28 {
                                                    case 32767:
                                                        v64 := v4[433]
                                                        _ = v64
                                                    default:
                                                        _ = len(v2)
                                                    }
                                                    _ = v59
                                                    _ = v62
                                                    _ = v63
                                                }
                                            case 2147483647:
                                                v65 := GenS7{F0: 0, F1: 2, F2: 2147483647}
                                                _ = v65
                                            }
                                            if (((813 == 159) || ("foo bar" == "")) || (true && true)) {
                                                if (fi44 >= int(v1)) {
                                                    v66 := GenS8{F0: "hello", F1: 376, F2: (v1 - v1)}
                                                    var v67 string = "\n\t"
                                                    _ = v66
                                                    _ = v67
                                                }
                                                v68 := GenS9{F0: 2}
                                                _ = v68
                                            } else {
                                                if false {
                                                    v5 = fi44
                                                    var v69 float64 = 3.14159
                                                    switch -1 {
                                                    case 449:
                                                        v5 = 42
                                                    case -852:
                                                        if false {
                                                            switch (^(v5)) {
                                                            case 141:
                                                                var v70 uint64 = 404
                                                                _ = v70
                                                            default:
                                                                v71 := GenS3{F0: false}
                                                                _ = v71
                                                            }
                                                            var v72 int8 = 27
                                                            var v73 string = ("\n\t" + "a")
                                                            v5 += (fi44 + int(v72))
                                                            _ = v72
                                                            _ = v73
                                                        }
                                                        var v74 uint16 = 2
                                                        if (false && true) {
                                                            fmt.Println(v1, v1)
                                                            var v75 int32 = -1
                                                            _ = v75
                                                        }
                                                        var v76 int16 = (v1 * -1)
                                                        _ = v74
                                                        _ = v76
                                                    default:
                                                        var v77 float64 = 1.5
                                                        _ = v77
                                                    }
                                                    _ = v69
                                                } else {
                                                    switch "hello" {
                                                    case "\n\t":
                                                        if false {
                                                            v78 := GenS3{F0: false}
                                                            v79 := GenS10{F0: 65535}
                                                            _ = v78
                                                            _ = v79
                                                        } else {
                                                            fmt.Println(v1, v1)
                                                        }
                                                        var v80 float32 = v27
                                                        var v81 int64 = 42
                                                        _ = v80
                                                        _ = v81
                                                    case "a":
                                                        v5 = (fi44 / (((-(fi28))) | 1))
                                                    case "foo bar":
                                                        v27 = (v27 * ((v27 * v27) + (v27 + (1.5 * v27))))
                                                    default:
                                                        delete(v2, v5)
                                                    }
                                                    if false {
                                                        switch (fi44 - (fi44 ^ (fi44 << ((18446744073709551615) & 63)))) {
                                                        case 32767:
                                                            v27 = v27
                                                        default:
                                                            _ = len(v4)
                                                            switch "a" {
                                                            case "\n\t":
                                                                if ((false && (!(true || false))) || (!false)) {
                                                                    var v82 int64 = 1
                                                                    _ = v82
                                                                }
                                                            default:
                                                                v83 := v2[fi44]
                                                                _ = v83
                                                            }
                                                            v84 := GenS10{F0: 4294967295}
                                                            _ = v84
                                                        }
                                                    }
                                                }
                                                var v85 uint32 = 202
                                                _ = v85
                                            }
                                        }
                                        v1 = v1
                                        _ = v47
                                    } else {
                                        v86 := GenS7{F0: 0, F1: 255, F2: 42}
                                        _ = v86
                                    }
                                    v27 = v27
                                default:
                                    if true {
                                        var v87 uint64 = 545
                                        _ = v87
                                    }
                                }
                                v5 = fi44
                                if (725 > 916) {
                                    v88 := v4[65535]
                                    _ = v88
                                }
                            } else {
                                v1 = 1000
                                v89 := []int{v5}
                                _ = v89
                            }
                            _ = fi44
                            _ = v45
                        }
                        _ = fi28
                    }
                case "foo bar":
                    fmt.Println(v1, v1)
                }
                for fi90 := 0; fi90 < 17; fi90++ {
                    fmt.Println(v27, v1)
                    var v91 uint16 = (^(v3))
                    var v92 int32 = -2147483648
                    _ = fi90
                    _ = v91
                    _ = v92
                }
            }
        }
        v3 = v3
        _ = v27
    } else {
        v93 := GenS2{F0: 2147483647, F1: 2}
        var v94 int = (v5 + v5)
        _ = v93
        _ = v94
    }
    var v95 uint8 = 182
    for fi96 := 0; fi96 < 6; fi96++ {
        v97 := []int{v5, (v5 ^ 1)}
        if false {
            var v98 float32 = -2.25
            _ = v98
        } else {
            switch "a" {
            case "foo bar":
                for fi99 := 0; fi99 < 17; fi99++ {
                    var v100 uint16 = (v3 + (v3 + (v3 * uint16(v95))))
                    _ = fi99
                    _ = v100
                }
            default:
                var v101 uint16 = (^(v3))
                _ = v101
            }
            for fi102 := 0; fi102 < 10; fi102++ {
                v1 += v1
                _ = fi102
            }
        }
        _ = fi96
        _ = v97
    }
    v95 += 1
    for fi103 := 0; fi103 < 16; fi103++ {
        var v104 string = (("\n\t" + ("" + "")) + ("a" + ""))
        var v105 float32 = 0.0
        var v106 float32 = -2.25
        _ = fi103
        _ = v104
        _ = v105
        _ = v106
    }
    v1 = 32767
    v4[65535] = uint32(v3)
    v107 := GenS3{F0: false}
    if false {
        fmt.Println(v1, v1)
        v4[42] = 2
    }
    switch (v5 / (((v5 ^ 0)) | 1)) {
    case 42:
        for fi108 := 0; fi108 < 7; fi108++ {
            for fi109 := 0; fi109 < 15; fi109++ {
                var v110 int = (fi109 - -182)
                var v111 float64 = 1.5
                delete(v2, v5)
                _ = fi109
                _ = v110
                _ = v111
            }
            v1 = 1000
            _ = fi108
        }
    case 2147483647:
        if (65535 <= 67) {
            if true {
                if (!false) {
                    fmt.Println(v5, v95, v1)
                    switch 32767 {
                    case -256:
                        v112 := GenS11{F0: (-(v5)), F1: "a"}
                        _ = v112
                    default:
                        v3 += 2
                    }
                    v95 = v95
                }
                v3 -= v3
            } else {
                _ = len(v2)
            }
        } else {
            var v113 int32 = -298
            var v114 uint = 18446744073709551615
            _ = v113
            _ = v114
        }
    default:
        var v115 float32 = 0.0
        _ = v115
    }
    if true {
        if false {
            switch (("hello" + ("" + "a")) + ("foo bar" + "\n\t")) {
            case "\n\t":
                if false {
                    if true {
                        switch "hello" {
                        case "hello":
                            _ = len(v2)
                        case "":
                            if (-9223372036854775808 == -9223372036854775808) {
                                switch ("hello" + "a") {
                                case "":
                                    v116 := []int{v5, v5}
                                    _ = v116
                                case "hello":
                                    var v117 int8 = -128
                                    _ = v117
                                case "foo bar":
                                    if false {
                                        var v118 uint32 = 1
                                        _ = v118
                                    } else {
                                        if true {
                                            fmt.Println(v95, v95)
                                            v3 = 190
                                        }
                                        v95 -= 127
                                        v119 := GenS12{F0: v95, F1: (v95 / ((1) | 1)), F2: 348}
                                        _ = v119
                                    }
                                default:
                                    fmt.Println(v5)
                                }
                                for fi120 := 0; fi120 < 2; fi120++ {
                                    if (!true) {
                                        v5 += v5
                                        v121 := []int{fi120, (-(fi120)), fi120, v5}
                                        v122 := GenS11{F0: 32768, F1: ("\n\t" + ("foo bar" + ""))}
                                        _ = v121
                                        _ = v122
                                    } else {
                                        v1 = (^(v1))
                                        v123 := GenS7{F0: 2147483648, F1: v95, F2: 348}
                                        v1 = (v1 << ((255) & 63))
                                        _ = v123
                                    }
                                    v124 := GenS11{F0: -9223372036854775808, F1: "\n\t"}
                                    _ = fi120
                                    _ = v124
                                }
                            } else {
                                switch (v5 + int(v3)) {
                                case -9223372036854775808:
                                    v125 := []int{(v5 % (((v5 >> ((255) & 63))) | 1))}
                                    _ = v125
                                default:
                                    switch "a" {
                                    case "hello":
                                        switch v5 {
                                        case 9223372036854775807:
                                            for fi126 := 0; fi126 < 14; fi126++ {
                                                fmt.Println(v3)
                                                fmt.Println(v95, v1)
                                                _ = fi126
                                            }
                                        default:
                                            var v127 int8 = 26
                                            _ = v127
                                        }
                                    case "\n\t":
                                        v3 = v3
                                    case "":
                                        fmt.Println(v95)
                                    }
                                }
                                fmt.Println(v3)
                            }
                        case "a":
                            switch 675 {
                            case 127:
                                fmt.Println(v1, v5, v3)
                            case -9223372036854775808:
                                fmt.Println(v5, v5, v1)
                            default:
                                var v128 int64 = -9223372036854775808
                                _ = v128
                            }
                        default:
                            for fi129 := 0; fi129 < 8; fi129++ {
                                v5 = (-(fi129))
                                for fi130 := 0; fi130 < 10; fi130++ {
                                    v1 -= (^(v1))
                                    switch "a" {
                                    case "a":
                                        fmt.Println(fi129, fi130)
                                    default:
                                        if true {
                                            v5 = 0
                                            var v131 string = "foo bar"
                                            _ = v131
                                        }
                                    }
                                    var v132 string = "foo bar"
                                    _ = fi130
                                    _ = v132
                                }
                                if true {
                                    for fi133 := 0; fi133 < 8; fi133++ {
                                        if (((("" + ("hello" + "a")) + "\n\t") + "a") == "") {
                                            if false {
                                                if true {
                                                    fmt.Println(v3, v3, v95)
                                                    if false {
                                                        var v134 int = fi133
                                                        var v135 float32 = 0.1
                                                        v95 += v95
                                                        _ = v134
                                                        _ = v135
                                                    }
                                                }
                                                fmt.Println(fi129, v95, v5)
                                            } else {
                                                fmt.Println(fi129, v5)
                                            }
                                            v136 := v4[1]
                                            _ = v136
                                        } else {
                                            v137 := GenS13{F0: 2, F1: 436, F2: (v3 & (v3 ^ (v3 >> ((2) & 63))))}
                                            _ = v137
                                        }
                                        var v138 uint32 = 2147483648
                                        var v139 int = -470
                                        _ = fi133
                                        _ = v138
                                        _ = v139
                                    }
                                    switch fi129 {
                                    case 244:
                                        switch (v5 << ((18446744073709551615) & 63)) {
                                        case 1:
                                            var v140 int64 = 9223372036854775807
                                            _ = v140
                                        case -9223372036854775808:
                                            if (1e10 >= (3.14159 + 0.0)) {
                                                fmt.Println(v1)
                                                v141 := v2[v5]
                                                _ = v141
                                            }
                                        case 32767:
                                            v3 = v3
                                        default:
                                            if false {
                                                v5 = v5
                                                var v142 uint64 = 622
                                                var v143 int32 = -2147483648
                                                _ = v142
                                                _ = v143
                                            } else {
                                                if (-891 == -1) {
                                                    var v144 int8 = -128
                                                    for fi145 := 0; fi145 < 2; fi145++ {
                                                        if false {
                                                            var v146 uint8 = 1
                                                            v147 := GenS2{F0: 0, F1: 859}
                                                            var v148 int = fi145
                                                            _ = v146
                                                            _ = v147
                                                            _ = v148
                                                        } else {
                                                            switch int(v3) {
                                                            case 127:
                                                                var v149 int = fi129
                                                                _ = v149
                                                            default:
                                                                var v150 string = (("\n\t" + "\n\t") + "\n\t")
                                                                _ = v150
                                                            }
                                                        }
                                                        v151 := []int{v5, fi145, (fi145 % ((fi129) | 1))}
                                                        _ = fi145
                                                        _ = v151
                                                    }
                                                    v3 += (^(v3))
                                                    _ = v144
                                                } else {
                                                    v95 = 0
                                                }
                                                for fi152 := 0; fi152 < 5; fi152++ {
                                                    var v153 float32 = 0.0
                                                    if true {
                                                        v5 -= fi129
                                                        if true {
                                                            var v154 float64 = 0.1
                                                            _ = v154
                                                        }
                                                        var v155 uint16 = v3
                                                        _ = v155
                                                    } else {
                                                        v95 = 71
                                                        v1 += 0
                                                        v156 := GenS8{F0: "a", F1: v3, F2: 42}
                                                        _ = v156
                                                    }
                                                    _ = fi152
                                                    _ = v153
                                                }
                                                if true {
                                                    var v157 uint64 = 0
                                                    fmt.Println(v3, v157)
                                                    fmt.Println(v3, v95)
                                                    _ = v157
                                                } else {
                                                    v158 := GenS10{F0: 0}
                                                    for fi159 := 0; fi159 < 10; fi159++ {
                                                        if true {
                                                            if (0 != 57) {
                                                                v1 += (^(v1))
                                                                var v160 float64 = 1e10
                                                                var v161 int16 = -32768
                                                                _ = v160
                                                                _ = v161
                                                            }
                                                            switch ("\n\t" + "a") {
                                                            case "\n\t":
                                                                if true {
                                                                    v95 = (v95 + v95)
                                                                    v162 := []int{int(v3), -581, 128, 32768}
                                                                    _ = v162
                                                                }
                                                                switch "hello" {
                                                                case "hello":
                                                                    var v163 uint8 = 255
                                                                    _ = v163
                                                                case "\n\t":
                                                                    var v164 float64 = 0.1
                                                                    _ = v164
                                                                default:
                                                                    var v165 string = "hello"
                                                                    _ = v165
                                                                }
                                                            case "foo bar":
                                                                switch (fi129 & 737) {
                                                                case 1:
                                                                    var v166 int = fi129
                                                                    _ = v166
                                                                case 128:
                                                                    var v167 float64 = (float64(float32(1e10)) - float64(3.14159))
                                                                    _ = v167
                                                                case -1:
                                                                    v95 -= 135
                                                                    var v168 string = "a"
                                                                    _ = v168
                                                                default:
                                                                    v3 += v3
                                                                }
                                                                fmt.Println(v5, v5, v1)
                                                                var v169 string = "hello"
                                                                _ = v169
                                                            case "hello":
                                                                var v170 int8 = -1
                                                                _ = v170
                                                            }
                                                            v3 = v3
                                                            var v171 int16 = -1
                                                            v171 = v1
                                                            var v172 uint32 = 2
                                                            var v173 uint64 = 42
                                                            _ = v171
                                                            _ = v172
                                                            _ = v173
                                                        }
                                                        v174 := v4[2]
                                                        _ = fi159
                                                        _ = v174
                                                    }
                                                    _ = v158
                                                }
                                            }
                                        }
                                    case -1:
                                        v175 := []int{(^(v5))}
                                        _ = v175
                                    default:
                                        switch "a" {
                                        case "\n\t":
                                            fmt.Println(v5)
                                        case "foo bar":
                                            var v176 int16 = (-(v1))
                                            _ = v176
                                        default:
                                            if (2 >= 65535) {
                                                v95 = v95
                                            } else {
                                                switch "hello" {
                                                case "\n\t":
                                                    switch (fi129 + (fi129 >> ((0) & 63))) {
                                                    case 0:
                                                        switch (v5 / ((32767) | 1)) {
                                                        case -469:
                                                            fmt.Println(v3, fi129)
                                                        }
                                                    case 127:
                                                        if true {
                                                            if false {
                                                                var v177 float32 = 1e10
                                                                for fi178 := 0; fi178 < 17; fi178++ {
                                                                    var v179 string = "a"
                                                                    _ = fi178
                                                                    _ = v179
                                                                }
                                                                switch (^(fi129)) {
                                                                case 524:
                                                                    v180 := v2[(v5 >> ((0) & 63))]
                                                                    _ = v180
                                                                case 2147483647:
                                                                    v1 -= 32767
                                                                case 9223372036854775807:
                                                                    if true {
                                                                        v1 = -1
                                                                        v95 += v95
                                                                        if ((float32(-2.25) * v177) >= v177) {
                                                                            v181 := GenS13{F0: 2, F1: 1000, F2: (v3 << ((65535) & 63))}
                                                                            v5 = (fi129 << ((255) & 63))
                                                                            _ = v181
                                                                        } else {
                                                                            for fi182 := 0; fi182 < 8; fi182++ {
                                                                                v183 := GenS14{F0: 264}
                                                                                _ = fi182
                                                                                _ = v183
                                                                            }
                                                                            var v184 int = (fi129 ^ v5)
                                                                            switch (v184 * 127) {
                                                                            case 128:
                                                                                v185 := []int{(v5 & (v5 + (v184 << ((380) & 63)))), 4294967295}
                                                                                _ = v185
                                                                            case -1:
                                                                                fmt.Println(v1, v184, v177)
                                                                            case 9223372036854775807:
                                                                                fmt.Println(v177, v177)
                                                                            default:
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    var v186 float64 = (float64(v177) * 0.0)
                                                                                    _ = v186
                                                                                case "":
                                                                                    v3 += v3
                                                                                case "foo bar":
                                                                                    if (v1 == v1) {
                                                                                        v177 -= (v177 * (v177 * 1.5))
                                                                                        v5 += v5
                                                                                    } else {
                                                                                        switch v5 {
                                                                                        case 127:
                                                                                            for fi187 := 0; fi187 < 5; fi187++ {
                                                                                                var v188 int8 = -128
                                                                                                v189 := []int{42, (fi129 - (-(v184)))}
                                                                                                _ = fi187
                                                                                                _ = v188
                                                                                                _ = v189
                                                                                            }
                                                                                        case 4294967295:
                                                                                            switch v5 {
                                                                                            case 9223372036854775807:
                                                                                                var v190 int8 = 127
                                                                                                _ = v190
                                                                                            case -9223372036854775808:
                                                                                                v3 = 0
                                                                                            case 128:
                                                                                                var v191 int16 = -32768
                                                                                                _ = v191
                                                                                            }
                                                                                        default:
                                                                                            v192 := GenS4{F0: v1}
                                                                                            _ = v192
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    v2[v5] = 4294967295
                                                                                }
                                                                            }
                                                                            _ = v184
                                                                        }
                                                                    }
                                                                }
                                                                _ = v177
                                                            } else {
                                                                var v193 bool = ((v1 & 590) <= (^(v1)))
                                                                v194 := []int{}
                                                                v194 = append(v194, (v5 % (((v5 / (((fi129 >> ((255) & 63))) | 1))) | 1)))
                                                                _ = v193
                                                                _ = v194
                                                            }
                                                            v1 -= v1
                                                        }
                                                    }
                                                case "":
                                                    v95 += 0
                                                }
                                                for fi195 := 0; fi195 < 17; fi195++ {
                                                    v196 := GenS11{F0: v5, F1: "foo bar"}
                                                    v5 += (fi195 << ((0) & 63))
                                                    _ = fi195
                                                    _ = v196
                                                }
                                                v5 = (fi129 % ((v5) | 1))
                                            }
                                        }
                                    }
                                    v197 := GenS1{F0: (v1 * v1), F1: 42}
                                    _ = v197
                                } else {
                                    v1 += 42
                                    for fi198 := 0; fi198 < 7; fi198++ {
                                        if (!true) {
                                            switch (v5 << ((65535) & 63)) {
                                            case 127:
                                                var v199 string = "a"
                                                _ = v199
                                            default:
                                                fmt.Println(fi129, v95)
                                            }
                                            var v200 float32 = -2.25
                                            _ = v200
                                        } else {
                                            fmt.Println(v1, v5)
                                            v201 := []int{128}
                                            _ = v201
                                        }
                                        var v202 bool = true
                                        var v203 int16 = (v1 ^ v1)
                                        v204 := []int{fi129}
                                        _ = fi198
                                        _ = v202
                                        _ = v203
                                        _ = v204
                                    }
                                }
                                _ = fi129
                            }
                        }
                        var v205 int = 0
                        _ = v205
                    } else {
                        if (!true) {
                            if ((v3 / (((v3 >> ((1) & 63))) | 1)) == 32767) {
                                if (-2.25 >= float32(((((0.0 / (float64((3.14159 * ((0.1 / (float32((float64(3.14159) + 1.5)) * 1e10)) * 1e10))) - 3.14159)) - -2.25) + 1e10) - (0.0 * 3.14159)))) {
                                    fmt.Println(v3)
                                    v206 := GenS15{F0: 576}
                                    v1 = v1
                                    _ = v206
                                }
                                var v207 uint8 = v95
                                if true {
                                    if true {
                                        for fi208 := 0; fi208 < 10; fi208++ {
                                            if (!((1 <= 500) || true)) {
                                                var v209 uint16 = v3
                                                v3 = (v3 | 176)
                                                var v210 float64 = -2.25
                                                _ = v209
                                                _ = v210
                                            } else {
                                                fmt.Println(v5, v1)
                                                v2[v5] = 4294967295
                                                var v211 float64 = 1e10
                                                _ = v211
                                            }
                                            _ = fi208
                                        }
                                        v212 := GenS6{F0: 32767, F1: 100, F2: "a"}
                                        v213 := GenS16{F0: 2147483647, F1: 0, F2: v3}
                                        _ = v212
                                        _ = v213
                                    }
                                } else {
                                    var v214 uint = 30
                                    for fi215 := 0; fi215 < 14; fi215++ {
                                        v2[(v5 / ((-2147483648) | 1))] = 922
                                        v216 := []int{(v5 % ((v5) | 1)), fi215, (fi215 * (v5 + 127)), (v5 * (-(fi215)))}
                                        _ = fi215
                                        _ = v216
                                    }
                                    v217 := GenS6{F0: 0, F1: 1, F2: "a"}
                                    _ = v214
                                    _ = v217
                                }
                                _ = v207
                            }
                        } else {
                            fmt.Println(v5, v95)
                            var v218 string = "hello"
                            v219 := GenS4{F0: (v1 << ((4294967295) & 63))}
                            _ = v218
                            _ = v219
                        }
                        if ((v3 * (v3 ^ (v3 % ((v3) | 1)))) != 311) {
                            var v220 float64 = 0.0
                            _ = v220
                        }
                    }
                }
            }
        } else {
            var v221 uint64 = genFn1("hello", false)
            _ = v221
        }
    }
    for fi222 := 0; fi222 < 18; fi222++ {
        fmt.Println(v1)
        _ = fi222
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v107
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v95
}
