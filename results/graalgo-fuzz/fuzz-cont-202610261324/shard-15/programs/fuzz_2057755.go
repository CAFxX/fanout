package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 uint8; F1 uint; F2 int32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int8; F1 uint8 }
type GenS9 struct { F0 int32 }
func genFn0(p0 uint16) int8 {
    var v7 uint = uint(p0)
    var v8 bool = true
    _ = p0
    _ = v7
    _ = v8
    return 1
}
func genFn1() uint32 {
    switch ("foo bar" + ("foo bar" + ("a" + "a"))) {
    case "hello":
        var v9 uint32 = 2
        _ = v9
    case "\n\t":
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 bool = false
            switch fi10 {
            case 4294967295:
                v11 = (fi10 == (fi10 ^ fi10))
            default:
                v12 := GenS1{F0: 128}
                _ = v12
            }
            _ = fi10
            _ = v11
        }
    }
    if (12 > 45) {
        if false {
            var v13 string = "foo bar"
            _ = v13
        }
    }
    if (481 < 0) {
        var v14 float64 = (1.5 * 0.1)
        var v15 uint64 = 0
        var v16 uint8 = 128
        _ = v14
        _ = v15
        _ = v16
    } else {
        var v17 bool = true
        if ((4294967295 < 51) && (false || v17)) {
            if ((1.5 > ((0.1 / (-2.25 - ((0.1 - float64(0.0)) * -2.25))) * 1.5)) || v17) {
                v18 := GenS1{F0: 0}
                _ = v18
            }
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v17)
            case "foo bar":
                v19 := GenS1{F0: 1}
                _ = v19
            }
        }
        _ = v17
    }
    var v20 uint8 = 42
    _ = v20
    return 658
}

func main() {
    var v1 bool = true
    var v2 int = -9223372036854775808
    v3 := make(map[uint32]bool)
    var v4 uint = 242
    v5 := make(map[uint32]uint32)
    var v6 int = (v2 / ((-121) | 1))
    for fi21 := 0; fi21 < 7; fi21++ {
        v2 -= (fi21 - fi21)
        var v22 bool = v1
        v22 = v1
        _ = fi21
        _ = v22
    }
    fmt.Println(v2, v6, v6)
    if v1 {
        switch 1 {
        case -9223372036854775808:
            v1 = (-128 > -7)
        case 42:
            if (!(!(896 > 0))) {
                var v23 int32 = 98
                v24 := GenS1{F0: 128}
                _ = v23
                _ = v24
            }
        case 239:
            if false {
                v2 = v6
            }
        }
    }
    fmt.Println(v1, v1, v4)
    v6 += v6
    for fi25 := 0; fi25 < 15; fi25++ {
        if true {
            var v26 string = "foo bar"
            _ = v26
        } else {
            switch "" {
            case "a":
                if v1 {
                    v1 = (4294967295 == v6)
                    v4 = 255
                    v2 = (v6 ^ fi25)
                }
            case "hello":
                if (0 <= 681) {
                    fmt.Println(v6, v6)
                    if v1 {
                        for fi27 := 0; fi27 < 18; fi27++ {
                            v6 = (v2 >> (((v4 << ((v4) & 63))) & 63))
                            _ = fi27
                        }
                        for fi28 := 0; fi28 < 17; fi28++ {
                            var v29 int8 = -1
                            _ = fi28
                            _ = v29
                        }
                    } else {
                        for fi30 := 0; fi30 < 18; fi30++ {
                            v4 = v4
                            v6 = (^(v6))
                            v31 := []int{fi30, 127}
                            _ = fi30
                            _ = v31
                        }
                        switch v6 {
                        case 2147483647:
                            if ((("foo bar" + "foo bar") + ("a" + "\n\t")) > ("\n\t" + (("foo bar" + "a") + ("hello" + "foo bar")))) {
                                v4 = v4
                                v32 := GenS1{F0: 2}
                                _ = v32
                            }
                        case 32767:
                            if v1 {
                                v2 = -743
                                if (v4 == v4) {
                                    if ((v2 < (fi25 >> ((v4) & 63))) && v1) {
                                        switch "foo bar" {
                                        case "a":
                                            v33 := GenS2{F0: 1}
                                            _ = v33
                                        case "":
                                            var v34 float32 = (0.1 * ((float32(((0.1 / (3.14159 + 3.14159)) / 3.14159)) + 3.14159) / 1e10))
                                            _ = v34
                                        }
                                        v6 += v2
                                    }
                                } else {
                                    fmt.Println(v4, fi25)
                                }
                            } else {
                                if true {
                                    v35 := GenS2{F0: 1000}
                                    _ = v35
                                } else {
                                    var v36 int64 = 42
                                    var v37 uint8 = 86
                                    _ = v36
                                    _ = v37
                                }
                                var v38 int32 = -2147483648
                                switch "\n\t" {
                                case "\n\t":
                                    fmt.Println(fi25)
                                case "hello":
                                    fmt.Println(v2, v4, v2)
                                case "":
                                    for fi39 := 0; fi39 < 19; fi39++ {
                                        v40 := []int{v6, fi39, (^(fi25)), v6}
                                        fmt.Println(fi39, v1, v6)
                                        var v41 int8 = 107
                                        _ = fi39
                                        _ = v40
                                        _ = v41
                                    }
                                }
                                _ = v38
                            }
                        case 32768:
                            if ((!((v1 && false) || v1)) || v1) {
                                var v42 uint32 = 469
                                if v1 {
                                    var v43 float64 = 1.5
                                    var v44 int64 = -9223372036854775808
                                    _ = v43
                                    _ = v44
                                }
                                if false {
                                    var v45 int16 = int16(v42)
                                    if (v1 && v1) {
                                        if (!("\n\t" == ("" + ""))) {
                                            v46 := GenS2{F0: v45}
                                            _ = v46
                                        }
                                        var v47 int16 = v45
                                        v48 := GenS1{F0: 1}
                                        _ = v47
                                        _ = v48
                                    } else {
                                        switch "\n\t" {
                                        case "\n\t":
                                            v4 -= v4
                                        case "a":
                                            for fi49 := 0; fi49 < 16; fi49++ {
                                                var v50 uint16 = 0
                                                switch ("hello" + "") {
                                                case "hello":
                                                    if true {
                                                        fmt.Println(fi49)
                                                        if v1 {
                                                            var v51 uint16 = (v50 & (v50 * (v50 | v50)))
                                                            v3[uint32(v50)] = v1
                                                            v1 = v1
                                                            _ = v51
                                                        } else {
                                                            v45 -= v45
                                                        }
                                                        fmt.Println(v1, v6)
                                                    }
                                                case "\n\t":
                                                    if (v1 || (true && v1)) {
                                                        switch ("a" + "a") {
                                                        case "foo bar":
                                                            var v52 bool = ((uint16(v42) >= v50) && v1)
                                                            _ = v52
                                                        case "hello":
                                                            v42 -= uint32(v4)
                                                            v53 := GenS2{F0: (v45 * int16(fi25))}
                                                            switch ("a" + ("hello" + ("foo bar" + "\n\t"))) {
                                                            case "":
                                                                var v54 uint16 = v50
                                                                _ = v54
                                                            case "\n\t":
                                                                if false {
                                                                    v55 := []int{(v6 & (v2 + (fi25 >> (((v4 / ((v4) | 1))) & 63))))}
                                                                    v42 += v42
                                                                    _ = v55
                                                                } else {
                                                                    fmt.Println(v2, v4)
                                                                    if false {
                                                                        var v56 int8 = genFn0(v50)
                                                                        _ = v56
                                                                    } else {
                                                                        v50 += v50
                                                                        v6 += (fi25 ^ fi49)
                                                                        var v57 string = (("a" + ("hello" + "a")) + "")
                                                                        var v58 uint8 = 137
                                                                        _ = v57
                                                                        _ = v58
                                                                    }
                                                                    var v59 int16 = v45
                                                                    _ = v59
                                                                }
                                                            default:
                                                                _ = len(v3)
                                                            }
                                                            _ = v53
                                                        case "\n\t":
                                                            switch "hello" {
                                                            case "foo bar":
                                                                var v60 int8 = 1
                                                                _ = v60
                                                            default:
                                                                v50 -= 209
                                                            }
                                                        }
                                                    } else {
                                                        v6 -= v6
                                                        if ("foo bar" > "foo bar") {
                                                            v50 += v50
                                                            var v61 uint = v4
                                                            fmt.Println(fi25, v6, v6)
                                                            _ = v61
                                                        }
                                                        v50 = uint16(v4)
                                                    }
                                                }
                                                var v62 float32 = 1e10
                                                _ = fi49
                                                _ = v50
                                                _ = v62
                                            }
                                        case "foo bar":
                                            if (!true) {
                                                v63 := v5[v42]
                                                _ = v63
                                            } else {
                                                v45 += 0
                                                var v64 float32 = (-2.25 / (0.1 * 0.1))
                                                v65 := GenS1{F0: 255}
                                                _ = v64
                                                _ = v65
                                            }
                                        default:
                                            var v66 int = v2
                                            _ = v66
                                        }
                                        var v67 int16 = int16(v42)
                                        v68 := GenS2{F0: (v67 << (((v4 & (v4 - (v4 << (((v4 + v4)) & 63))))) & 63))}
                                        _ = v67
                                        _ = v68
                                    }
                                    _ = v45
                                }
                                _ = v42
                            }
                        default:
                            if (("a" + (("hello" + ("foo bar" + "foo bar")) + "")) >= ("hello" + "hello")) {
                                v4 += v4
                                fmt.Println(v1, fi25, v2)
                                v69 := GenS3{F0: 652, F1: -1, F2: 1}
                                _ = v69
                            } else {
                                v70 := GenS3{F0: 42, F1: 0, F2: 2147483647}
                                _ = v70
                            }
                        }
                    }
                    var v71 string = "\n\t"
                    _ = v71
                } else {
                    v6 = v2
                    v72 := GenS1{F0: 127}
                    _ = v72
                }
            case "":
                for fi73 := 0; fi73 < 3; fi73++ {
                    if false {
                        var v74 string = "a"
                        v1 = v1
                        _ = v74
                    } else {
                        var v75 string = ""
                        _ = v75
                    }
                    switch (v6 ^ 128) {
                    case 2147483647:
                        if v1 {
                            var v76 int32 = int32(v4)
                            var v77 uint8 = 0
                            fmt.Println(v1)
                            _ = v76
                            _ = v77
                        } else {
                            if (((0.0 + (((1.5 * 1.5) / 1.0) + (1.5 * 1e10))) + 0.1) >= 0.0) {
                                fmt.Println(v4)
                                v78 := GenS2{F0: -1}
                                _ = v78
                            } else {
                                v6 = v6
                            }
                        }
                    case 32767:
                        v79 := GenS1{F0: 39}
                        _ = v79
                    default:
                        if v1 {
                            v80 := GenS1{F0: 149}
                            _ = len(v5)
                            if false {
                                var v81 bool = v1
                                _ = v81
                            }
                            _ = v80
                        } else {
                            v2 = (-(v2))
                            var v82 uint8 = 2
                            if true {
                                v1 = ((0.1 - 1e10) != ((0.0 - -2.25) * (1e10 - -2.25)))
                                var v83 int32 = 140
                                _ = v83
                            } else {
                                var v84 uint64 = 1
                                _ = len(v3)
                                if v1 {
                                    if (v2 > (v6 | (v6 + -85))) {
                                        if v1 {
                                            var v85 int64 = 9223372036854775807
                                            var v86 int16 = -32768
                                            _ = v85
                                            _ = v86
                                        }
                                        v87 := []int{fi73, v2, 544}
                                        _ = v87
                                    } else {
                                        if true {
                                            if v1 {
                                                var v88 float32 = 1.5
                                                var v89 float64 = (0.1 + (0.0 + (1.5 + (float64(v88) / 1.5))))
                                                switch "foo bar" {
                                                case "\n\t":
                                                    fmt.Println(v82, v82)
                                                }
                                                _ = v88
                                                _ = v89
                                            }
                                        }
                                        fmt.Println(v82, fi25)
                                        v90 := GenS3{F0: 32767, F1: 100, F2: 2147483648}
                                        v4 = v4
                                        _ = v90
                                    }
                                } else {
                                    fmt.Println(v4, fi25, v6)
                                }
                                if (9223372036854775807 <= 104) {
                                    fmt.Println(fi25, v4)
                                    var v91 int8 = 117
                                    _ = v91
                                } else {
                                    fmt.Println(v6, fi73)
                                    v92 := []int{(fi25 | (v6 % ((fi25) | 1))), v2, v2, (fi25 ^ (fi25 * fi25))}
                                    _ = v92
                                }
                                _ = v84
                            }
                            _ = v82
                        }
                        v6 -= (-(fi25))
                        v4 = v4
                    }
                    var v93 int = (fi25 % ((-9223372036854775808) | 1))
                    fmt.Println(fi25, v93)
                    fmt.Println(v2, v93, fi25)
                    if (!((1e10 / (((float32(1e10) * (-2.25 / 3.14159)) + 1e10) * 1.5)) >= float32((-2.25 - (0.1 - float64((0.0 - 1e10))))))) {
                        if ((!(1 < 13)) && (!v1)) {
                            fmt.Println(v2, v1)
                            if false {
                                v4 = (v4 / (((v4 * v4)) | 1))
                            } else {
                                switch (fi73 << ((v4) & 63)) {
                                case 127:
                                    v94 := GenS3{F0: 42, F1: genFn0(653), F2: 737}
                                    _ = v94
                                }
                            }
                            fmt.Println(v6, fi25)
                        } else {
                            switch "" {
                            case "\n\t":
                                if true {
                                    v95 := GenS4{F0: v4, F1: -1}
                                    var v96 int16 = 1
                                    v4 += 18446744073709551615
                                    _ = v95
                                    _ = v96
                                }
                            case "foo bar":
                                var v97 int16 = 32767
                                _ = v97
                            default:
                                v98 := []int{v93, (v6 | (v93 - (fi25 << ((v4) & 63))))}
                                _ = v98
                            }
                            v2 = (fi73 % ((fi73) | 1))
                            v99 := GenS5{F0: 1, F1: 42}
                            switch v6 {
                            case 32768:
                                v93 = fi73
                            case 4294967295:
                                v100 := GenS2{F0: int16(v93)}
                                _ = v100
                            case 938:
                                var v101 string = "\n\t"
                                _ = v101
                            default:
                                v4 = 0
                                v93 += 127
                                v102 := GenS1{F0: 189}
                                v103 := []int{}
                                switch "a" {
                                case "foo bar":
                                    v1 = false
                                    switch "foo bar" {
                                    case "a":
                                        var v104 uint64 = 861
                                        _ = v104
                                    case "foo bar":
                                        fmt.Println(v2)
                                    default:
                                        v105 := GenS4{F0: (v4 % ((v4) | 1)), F1: -128}
                                        _ = v105
                                    }
                                    if false {
                                        switch "foo bar" {
                                        case "a":
                                            var v106 int16 = 1000
                                            _ = v106
                                        case "foo bar":
                                            var v107 float32 = -2.25
                                            _ = v107
                                        default:
                                            if true {
                                                if (595 >= 0) {
                                                    v108 := GenS6{F0: 228, F1: (v4 & v4), F2: -1}
                                                    v6 -= (v6 % ((0) | 1))
                                                    var v109 uint = (v4 ^ (^(v4)))
                                                    _ = v108
                                                    _ = v109
                                                }
                                                v1 = v1
                                            }
                                        }
                                        delete(v5, 0)
                                        var v110 uint16 = 255
                                        _ = v110
                                    }
                                case "":
                                    v93 -= (^(v6))
                                default:
                                    var v111 uint = v4
                                    _ = v111
                                }
                                var v112 uint8 = 128
                                _ = v102
                                _ = v103
                                _ = v112
                            }
                            _ = v99
                        }
                        if v1 {
                            v113 := []int{v93, fi25}
                            fmt.Println(v6, v1)
                            v6 -= -9223372036854775808
                            _ = v113
                        } else {
                            v2 += (v93 << ((0) & 63))
                            v93 += (fi73 - fi25)
                        }
                    }
                    _ = fi73
                    _ = v93
                }
            default:
                v4 = (v4 / ((23) | 1))
            }
        }
        var v114 int32 = 0
        _ = fi25
        _ = v114
    }
    v4 += (v4 - 255)
    if ((!(((^(v4)) == 65535) && (-1 < 42))) && ((42 != 292) && (241 < 42))) {
        var v115 uint16 = 2
        v5[263] = 42
        v116 := GenS7{F0: (("hello" + "\n\t") + "hello")}
        _ = v115
        _ = v116
    }
    for fi117 := 0; fi117 < 18; fi117++ {
        if (1.5 != 1e10) {
            for fi118 := 0; fi118 < 15; fi118++ {
                v119 := []int{}
                _ = fi118
                _ = v119
            }
            v120 := GenS1{F0: 0}
            var v121 float64 = -2.25
            _ = v120
            _ = v121
        }
        v122 := GenS3{F0: 0, F1: 100, F2: 227}
        _ = fi117
        _ = v122
    }
    switch (v2 * -421) {
    case 1:
        if (-1 > -9223372036854775808) {
            if v1 {
                v123 := GenS1{F0: 127}
                _ = v123
            } else {
                var v124 int64 = 727
                _ = v124
            }
            for fi125 := 0; fi125 < 15; fi125++ {
                v126 := GenS8{F0: -128, F1: 140}
                _ = fi125
                _ = v126
            }
            for fi127 := 0; fi127 < 16; fi127++ {
                for fi128 := 0; fi128 < 19; fi128++ {
                    fmt.Println(v6)
                    _ = fi128
                }
                fmt.Println(v1, v4)
                var v129 int32 = -1
                _ = fi127
                _ = v129
            }
        } else {
            v2 += (v2 % ((-495) | 1))
            if (0 > 2) {
                v1 = (v1 && (!(0.0 <= 1.5)))
                for fi130 := 0; fi130 < 10; fi130++ {
                    v131 := []int{fi130, (^(v2))}
                    for fi132 := 0; fi132 < 10; fi132++ {
                        var v133 int16 = 419
                        _ = fi132
                        _ = v133
                    }
                    v6 = 204
                    _ = fi130
                    _ = v131
                }
                var v134 int8 = 100
                _ = v134
            } else {
                switch -385 {
                case 128:
                    fmt.Println(v6, v4)
                default:
                    switch v2 {
                    case -474:
                        v135 := GenS9{F0: 606}
                        _ = v135
                    }
                }
            }
            if ((0.1 * 3.14159) > -2.25) {
                fmt.Println(v1)
                switch -1 {
                case -1:
                    for fi136 := 0; fi136 < 1; fi136++ {
                        var v137 bool = v1
                        _ = fi136
                        _ = v137
                    }
                case 2147483647:
                    var v138 string = ""
                    _ = v138
                default:
                    for fi139 := 0; fi139 < 4; fi139++ {
                        var v140 uint8 = 2
                        fmt.Println(v2, fi139)
                        var v141 uint = 457
                        _ = fi139
                        _ = v140
                        _ = v141
                    }
                }
                v1 = false
            } else {
                if false {
                    for fi142 := 0; fi142 < 1; fi142++ {
                        if (!v1) {
                            fmt.Println(v2, v6)
                            switch "" {
                            case "hello":
                                v143 := []int{(fi142 + 128), fi142, (fi142 & (^(v2)))}
                                _ = v143
                            case "a":
                                switch -1 {
                                case 127:
                                    var v144 uint8 = 1
                                    _ = v144
                                case 32768:
                                    for fi145 := 0; fi145 < 20; fi145++ {
                                        if (true && v1) {
                                            var v146 uint = v4
                                            v146 = v4
                                            switch "\n\t" {
                                            case "hello":
                                                switch ("a" + "hello") {
                                                case "foo bar":
                                                    if v1 {
                                                        if false {
                                                            switch (("a" + "a") + "foo bar") {
                                                            case "a":
                                                                fmt.Println(v1)
                                                            }
                                                            v6 = v2
                                                        }
                                                        var v147 float32 = 3.14159
                                                        _ = v147
                                                    }
                                                case "\n\t":
                                                    v148 := GenS8{F0: -1, F1: 0}
                                                    _ = v148
                                                default:
                                                    v6 = fi145
                                                }
                                                switch ("\n\t" + ("hello" + "hello")) {
                                                case "foo bar":
                                                    v149 := GenS1{F0: 128}
                                                    v5[1] = 442
                                                    fmt.Println(v2, v6)
                                                    _ = v149
                                                }
                                            case "a":
                                                switch "" {
                                                case "\n\t":
                                                    var v150 int16 = 1000
                                                    _ = v150
                                                }
                                            }
                                            _ = v146
                                        }
                                        _ = fi145
                                    }
                                default:
                                    v151 := []int{1, v6, (v2 << ((42) & 63)), fi142}
                                    _ = v151
                                }
                            }
                            v1 = v1
                        } else {
                            var v152 uint32 = 795
                            _ = v152
                        }
                        for fi153 := 0; fi153 < 8; fi153++ {
                            var v154 int8 = -1
                            v155 := v5[42]
                            _ = fi153
                            _ = v154
                            _ = v155
                        }
                        var v156 int = (v2 * 888)
                        _ = fi142
                        _ = v156
                    }
                    if (!(v1 && ((32767 == 42) || true))) {
                        if false {
                            for fi157 := 0; fi157 < 3; fi157++ {
                                if false {
                                    if v1 {
                                        switch "foo bar" {
                                        case "a":
                                            for fi158 := 0; fi158 < 7; fi158++ {
                                                if (!((v4 << ((255) & 63)) == 558)) {
                                                    if v1 {
                                                        var v159 uint8 = 2
                                                        _ = v159
                                                    }
                                                    switch ("a" + "\n\t") {
                                                    case "a":
                                                        fmt.Println(fi158, v4, v4)
                                                    case "foo bar":
                                                        if true {
                                                            var v160 uint8 = 127
                                                            var v161 int = v6
                                                            _ = v160
                                                            _ = v161
                                                        }
                                                    case "\n\t":
                                                        var v162 uint16 = uint16(v4)
                                                        _ = v162
                                                    default:
                                                        fmt.Println(v4, fi157, v2)
                                                    }
                                                }
                                                fmt.Println(fi158, v1)
                                                _ = fi158
                                            }
                                        case "hello":
                                            delete(v5, 2147483647)
                                        case "foo bar":
                                            for fi163 := 0; fi163 < 12; fi163++ {
                                                v164 := GenS9{F0: 42}
                                                v4 = v4
                                                if ((!(!v1)) || v1) {
                                                    v1 = (v1 || v1)
                                                    switch "" {
                                                    case "\n\t":
                                                        v2 += (-(fi157))
                                                    case "a":
                                                        v4 += 4294967295
                                                    default:
                                                        v1 = v1
                                                    }
                                                    var v165 uint16 = 2
                                                    _ = v165
                                                } else {
                                                    if (v1 || true) {
                                                        if (537 >= 42) {
                                                            var v166 uint8 = 255
                                                            if (1.5 > 3.14159) {
                                                                if v1 {
                                                                    v167 := []int{}
                                                                    v168 := GenS6{F0: v166, F1: (^(v4)), F2: -452}
                                                                    var v169 int8 = -1
                                                                    fmt.Println(v6, fi163, v4)
                                                                    _ = v167
                                                                    _ = v168
                                                                    _ = v169
                                                                }
                                                                var v170 float64 = (1e10 / (0.1 + 1.5))
                                                                if (214 == 4294967295) {
                                                                    var v171 uint16 = 878
                                                                    _ = v171
                                                                }
                                                                _ = v170
                                                            }
                                                            _ = v166
                                                        }
                                                        v172 := GenS1{F0: 176}
                                                        _ = v172
                                                    }
                                                    v1 = (-2147483648 > -1)
                                                    if (1 < 0) {
                                                        fmt.Println(fi163, v6, fi157)
                                                        if (42 > 2) {
                                                            var v173 uint32 = 0
                                                            var v174 uint8 = 127
                                                            var v175 uint64 = 1
                                                            _ = v173
                                                            _ = v174
                                                            _ = v175
                                                        } else {
                                                            v4 = 2
                                                        }
                                                    } else {
                                                        fmt.Println(fi157)
                                                        fmt.Println(fi163, fi163)
                                                    }
                                                }
                                                _ = fi163
                                                _ = v164
                                            }
                                        default:
                                            switch v2 {
                                            case 4294967295:
                                                v176 := []int{fi157, 701}
                                                _ = v176
                                            case -133:
                                                _ = len(v5)
                                            }
                                        }
                                    }
                                    v177 := GenS3{F0: 2, F1: -128, F2: 4294967295}
                                    _ = v177
                                }
                                _ = fi157
                            }
                            fmt.Println(v1, v4, v1)
                            v178 := []int{}
                            _ = v178
                        }
                    }
                    var v179 bool = v1
                    _ = v179
                }
            }
        }
    case 0:
        fmt.Println(v4, v6)
    default:
        var v180 float32 = (float32(0.1) / -2.25)
        _ = v180
    }
    v1 = true
    for fi181 := 0; fi181 < 11; fi181++ {
        v182 := GenS8{F0: -93, F1: 255}
        var v183 uint8 = 1
        _ = fi181
        _ = v182
        _ = v183
    }
    switch "hello" {
    case "\n\t":
        switch v2 {
        case 42:
            if v1 {
                v184 := GenS6{F0: 36, F1: v4, F2: 504}
                _ = v184
            }
        default:
            fmt.Println(v6, v6)
        }
    case "hello":
        switch v6 {
        case 4294967295:
            for fi185 := 0; fi185 < 19; fi185++ {
                for fi186 := 0; fi186 < 18; fi186++ {
                    v187 := GenS4{F0: 0, F1: 42}
                    fmt.Println(v4, v1)
                    v4 = 245
                    _ = fi186
                    _ = v187
                }
                v188 := GenS3{F0: 0, F1: 100, F2: genFn1()}
                _ = fi185
                _ = v188
            }
        case 1:
            delete(v5, genFn1())
        }
    default:
        v3[4294967295] = v1
    }
    v6 -= (v2 ^ -290)
    v6 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
