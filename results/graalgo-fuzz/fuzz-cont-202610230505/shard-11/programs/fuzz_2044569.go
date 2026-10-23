package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint16 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS5 struct { F0 int; F1 int16; F2 uint }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int32; F1 uint8 }
type GenS10 struct { F0 uint16; F1 int64 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 int32; F1 string; F2 bool }
type GenS14 struct { F0 uint16; F1 string }
type GenS15 struct { F0 string }
func genFn0() uint16 {
    v6 := []int{676, 32768, -1, 127}
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 float32 = 3.14159
        var v9 bool = true
        if len(v6) > 0 {
            v6[0] = fi7
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = v6
    return 321
}
func genFn1(p0 uint8, p1 int8, p2 int8) int {
    v10 := GenS1{F0: -1, F1: ("hello" + "hello"), F2: 42}
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 uint64 = 18446744073709551615
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        v14 := []int{fi13}
        _ = fi13
        _ = v14
    }
    if false {
        var v15 int64 = 9223372036854775807
        _ = v15
    } else {
        var v16 int8 = p1
        switch "\n\t" {
        case "\n\t":
            var v17 int32 = -2147483648
            _ = v17
        case "foo bar":
            p1 -= 37
        }
        _ = v16
    }
    switch "" {
    case "hello":
        p1 = p2
    case "":
        for fi18 := 0; fi18 < 2; fi18++ {
            if false {
                if true {
                    var v19 uint32 = 413
                    p2 -= (p1 >> ((2) & 63))
                    if false {
                        switch fi18 {
                        case 4294967295:
                            switch "" {
                            case "a":
                                p2 = p2
                            default:
                                v20 := []int{(fi18 + (fi18 * -526)), fi18, fi18, 681}
                                _ = v20
                            }
                        case -2147483648:
                            var v21 bool = false
                            _ = v21
                        case 127:
                            for fi22 := 0; fi22 < 8; fi22++ {
                                v23 := GenS2{F0: 4294967295, F1: 9223372036854775807}
                                _ = fi22
                                _ = v23
                            }
                        }
                        v24 := []int{fi18, (-(fi18)), 9223372036854775807}
                        switch (fi18 - (fi18 + -864)) {
                        case 42:
                            if true {
                                switch ("foo bar" + "hello") {
                                case "a":
                                    for fi25 := 0; fi25 < 20; fi25++ {
                                        if (p0 != (p0 / ((6) | 1))) {
                                            var v26 int64 = 742
                                            _ = v26
                                        }
                                        v27 := GenS3{F0: 59}
                                        var v28 uint64 = 1
                                        var v29 float64 = 1e10
                                        v28 += (^(v28))
                                        _ = fi25
                                        _ = v27
                                        _ = v28
                                        _ = v29
                                    }
                                case "\n\t":
                                    v30 := GenS2{F0: 2, F1: 625}
                                    _ = v30
                                case "foo bar":
                                    var v31 uint8 = 127
                                    _ = v31
                                default:
                                    v32 := GenS2{F0: 18446744073709551615, F1: 0}
                                    _ = v32
                                }
                                v19 = 0
                            } else {
                                fmt.Println(p1, fi18, p2)
                                var v33 string = ""
                                var v34 int64 = 42
                                _ = v33
                                _ = v34
                            }
                        case 4294967295:
                            v19 += v19
                        case -2147483648:
                            _ = len(v24)
                        default:
                            for fi35 := 0; fi35 < 1; fi35++ {
                                if true {
                                    p1 = p1
                                    var v36 float32 = -2.25
                                    _ = v36
                                } else {
                                    if true {
                                        var v37 uint64 = 2
                                        _ = v37
                                    } else {
                                        if ("" >= "\n\t") {
                                            if (0.1 <= 1e10) {
                                                fmt.Println(p1)
                                                var v38 float32 = (0.0 + float32(3.14159))
                                                if (!(-71 >= int32(p0))) {
                                                    v39 := GenS2{F0: 42, F1: -9223372036854775808}
                                                    _ = v39
                                                }
                                                _ = v38
                                            } else {
                                                v40 := GenS4{F0: 152, F1: 42, F2: p0}
                                                var v41 int16 = 0
                                                var v42 bool = false
                                                v19 = (v19 / ((v19) | 1))
                                                if true {
                                                    if len(v24) >= 2 {
                                                        v24 = v24[1:len(v24)]
                                                    }
                                                    v19 = v19
                                                    fmt.Println(v42, p0, fi18)
                                                } else {
                                                    var v43 uint = 553
                                                    v44 := GenS2{F0: 2, F1: 9223372036854775807}
                                                    v45 := GenS5{F0: 42, F1: v41, F2: (v43 >> (((^(v43))) & 63))}
                                                    _ = v43
                                                    _ = v44
                                                    _ = v45
                                                }
                                                _ = v40
                                                _ = v41
                                                _ = v42
                                            }
                                            switch (fi18 | (fi35 ^ (fi18 + -511))) {
                                            case -356:
                                                v46 := GenS6{F0: 42}
                                                _ = v46
                                            default:
                                                p1 = 42
                                            }
                                            var v47 uint32 = v19
                                            _ = v47
                                        }
                                        var v48 int64 = 9223372036854775807
                                        v24 = append(v24, fi35)
                                        _ = v48
                                    }
                                    var v49 uint = 0
                                    _ = v49
                                }
                                _ = fi35
                            }
                        }
                        _ = v24
                    }
                    _ = v19
                } else {
                    for fi50 := 0; fi50 < 2; fi50++ {
                        if false {
                            var v51 uint8 = p0
                            var v52 uint8 = v51
                            _ = v51
                            _ = v52
                        } else {
                            switch -2147483648 {
                            case -1:
                                switch "hello" {
                                case "a":
                                    var v53 int16 = int16(p1)
                                    _ = v53
                                case "\n\t":
                                    v54 := []int{int(p1), (fi18 - fi18)}
                                    _ = v54
                                case "":
                                    if (!true) {
                                        var v55 uint64 = 4294967295
                                        p1 = -36
                                        if ((fi18 / (((^(fi50))) | 1)) > (fi18 % (((-(fi50))) | 1))) {
                                            if false {
                                                switch "hello" {
                                                case "":
                                                    fmt.Println(p0, p1, p1)
                                                case "a":
                                                    v56 := []int{fi50, (fi18 | (fi18 & 869)), (fi18 ^ 956), (fi18 + (fi18 >> ((0) & 63)))}
                                                    _ = v56
                                                case "\n\t":
                                                    if (!true) {
                                                        p0 = p0
                                                        var v57 uint32 = 486
                                                        if (2147483647 < -681) {
                                                            v58 := GenS2{F0: 42, F1: 937}
                                                            v59 := GenS7{F0: -180}
                                                            var v60 uint16 = uint16(v55)
                                                            _ = v58
                                                            _ = v59
                                                            _ = v60
                                                        } else {
                                                            switch fi18 {
                                                            case -311:
                                                                v61 := GenS8{F0: 1}
                                                                _ = v61
                                                            case 127:
                                                                if (42 != 456) {
                                                                    v62 := []int{fi50}
                                                                    _ = v62
                                                                } else {
                                                                    switch -338 {
                                                                    case 32767:
                                                                        var v63 uint64 = v55
                                                                        _ = v63
                                                                    case 9223372036854775807:
                                                                        p0 = 1
                                                                    case 42:
                                                                        if true {
                                                                            v64 := []int{}
                                                                            if len(v64) >= 2 {
                                                                                v64 = v64[1:len(v64)]
                                                                            }
                                                                            _ = v64
                                                                        }
                                                                    default:
                                                                        v65 := []int{fi50, 182, (^(fi18)), fi50}
                                                                        _ = v65
                                                                    }
                                                                }
                                                            case -360:
                                                                var v66 float64 = -2.25
                                                                _ = v66
                                                            }
                                                            var v67 string = ""
                                                            _ = v67
                                                        }
                                                        _ = v57
                                                    }
                                                }
                                                v68 := GenS9{F0: 42, F1: (p0 * p0)}
                                                var v69 int64 = -9223372036854775808
                                                v70 := []int{9223372036854775807}
                                                if ((p1 << ((2) & 63)) == (-(p1))) {
                                                    switch (^(fi18)) {
                                                    case -860:
                                                        p1 = p2
                                                    case 4294967295:
                                                        v69 += (v69 + (^(v69)))
                                                    case 0:
                                                        var v71 uint32 = 65535
                                                        _ = v71
                                                    }
                                                    if (false && (1e10 >= 1.5)) {
                                                        fmt.Println(fi50, p1)
                                                    }
                                                    if (((p1 < p1) || (fi50 > 843)) || false) {
                                                        v72 := GenS3{F0: -895}
                                                        fmt.Println(p1)
                                                        fmt.Println(v55, p0, p1)
                                                        _ = v72
                                                    }
                                                } else {
                                                    p2 = p1
                                                    fmt.Println(p1, fi50, fi18)
                                                    if false {
                                                        v70 = append(v70, 4294967295)
                                                        if len(v70) > 0 {
                                                            v70[0] = fi18
                                                        }
                                                        p0 += p0
                                                    } else {
                                                        if false {
                                                            fmt.Println(p2, p1)
                                                            v73 := GenS10{F0: 1, F1: (v69 ^ 42)}
                                                            _ = v73
                                                        }
                                                        p2 += 127
                                                        var v74 int = -9223372036854775808
                                                        _ = v74
                                                    }
                                                }
                                                _ = v68
                                                _ = v69
                                                _ = v70
                                            } else {
                                                var v75 float32 = (-2.25 - 3.14159)
                                                var v76 int8 = (p1 | (-(p1)))
                                                v77 := GenS7{F0: 1}
                                                v78 := []int{(fi18 >> ((917) & 63)), fi50, (fi18 % ((fi50) | 1)), fi50}
                                                _ = len(v78)
                                                switch (fi18 - fi50) {
                                                case 32768:
                                                    _ = len(v78)
                                                case 2147483647:
                                                    fmt.Println(fi50, p2, v75)
                                                case -763:
                                                    v76 -= (p1 >> ((42) & 63))
                                                default:
                                                    var v79 uint8 = (p0 * 2)
                                                    _ = v79
                                                }
                                                _ = v75
                                                _ = v76
                                                _ = v77
                                                _ = v78
                                            }
                                        } else {
                                            fmt.Println(p1, v55)
                                        }
                                        _ = v55
                                    } else {
                                        p2 = (p1 << ((348) & 63))
                                        var v80 uint16 = 0
                                        if (3.14159 < -2.25) {
                                            v81 := []int{fi18, (fi50 | 1)}
                                            fmt.Println(p0)
                                            fmt.Println(p0)
                                            _ = v81
                                        }
                                        _ = v80
                                    }
                                }
                            case 253:
                                fmt.Println(fi50, p2)
                            }
                        }
                        v82 := []int{-186, (^(fi50))}
                        switch "hello" {
                        case "":
                            p2 -= (p1 % ((-127) | 1))
                        case "foo bar":
                            p0 = 0
                        }
                        var v83 uint8 = p0
                        var v84 float64 = (((1e10 / float64(3.14159)) - (((0.0 / 1.5) - 3.14159) * (1.5 - 1.5))) - -2.25)
                        _ = fi50
                        _ = v82
                        _ = v83
                        _ = v84
                    }
                }
                if ((!(!(genFn0() > 561))) || ((p2 + (p1 * p2)) <= (p1 / (((p2 >> ((255) & 63))) | 1)))) {
                    v85 := GenS1{F0: -32768, F1: "\n\t", F2: 42}
                    fmt.Println(fi18)
                    _ = v85
                }
                p1 += (p1 / ((p2) | 1))
            } else {
                for fi86 := 0; fi86 < 6; fi86++ {
                    p1 = p2
                    _ = fi86
                }
                p1 = p2
                switch fi18 {
                case -2147483648:
                    var v87 float32 = 3.14159
                    _ = v87
                }
            }
            v88 := GenS1{F0: 1000, F1: "hello", F2: 548}
            var v89 int64 = -9223372036854775808
            _ = fi18
            _ = v88
            _ = v89
        }
    case "\n\t":
        v90 := []int{4294967295, 1, -2147483648}
        _ = v90
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 127
}

func main() {
    var v1 int32 = 912
    v2 := []uint32{0, 926, 65535, 0}
    v3 := make(map[int]uint32)
    v4 := make(map[int]string)
    var v5 int = 32768
    if (v1 >= (v1 & (v1 % (((v1 / ((-1) | 1))) | 1)))) {
        v1 = v1
        for fi91 := 0; fi91 < 4; fi91++ {
            if (false || true) {
                for fi92 := 0; fi92 < 7; fi92++ {
                    switch (fi92 >> ((925) & 63)) {
                    case 4294967295:
                        if false {
                            if (!(128 < 148)) {
                                switch (("" + ("foo bar" + (("hello" + "a") + "hello"))) + ("a" + "hello")) {
                                case "hello":
                                    switch "foo bar" {
                                    case "hello":
                                        if false {
                                            var v93 int64 = -113
                                            v93 = (v93 & (v93 / ((1) | 1)))
                                            _ = v93
                                        }
                                    case "foo bar":
                                        var v94 int64 = -879
                                        _ = v94
                                    default:
                                        _ = len(v2)
                                    }
                                    var v95 uint8 = 1
                                    v2 = append(v2, 0)
                                    _ = v95
                                case "\n\t":
                                    var v96 uint16 = 32767
                                    _ = v96
                                case "":
                                    var v97 string = ("\n\t" + "")
                                    _ = v97
                                default:
                                    fmt.Println(v5, fi92)
                                }
                                v1 = (^(v1))
                            }
                        }
                        fmt.Println(v1, fi91, fi91)
                        var v98 uint = 18446744073709551615
                        _ = v98
                    default:
                        v5 = fi91
                    }
                    v99 := GenS3{F0: v1}
                    v3[genFn1(0, 1, 100)] = 12
                    v5 = (fi91 | (fi92 & fi92))
                    switch fi92 {
                    case -9223372036854775808:
                        v1 = 693
                    case -232:
                        if false {
                            v100 := GenS4{F0: 42, F1: -959, F2: 42}
                            fmt.Println(fi91, fi92)
                            _ = v100
                        }
                        fmt.Println(fi92)
                        switch "foo bar" {
                        case "\n\t":
                            fmt.Println(fi91, v5, v1)
                        case "":
                            if (!(65535 != 42)) {
                                var v101 float64 = 0.1
                                v101 = -2.25
                                var v102 string = (("\n\t" + "hello") + "foo bar")
                                v103 := GenS5{F0: (fi91 / ((fi91) | 1)), F1: 0, F2: 255}
                                if true {
                                    v5 -= (fi92 >> ((19) & 63))
                                    if true {
                                        v101 -= v101
                                        v1 = v1
                                    } else {
                                        if false {
                                            switch v102 {
                                            case "a":
                                                v101 = float64((1e10 * (1.5 / (-2.25 - (float32(v101) * (1e10 - 1e10))))))
                                            case "hello":
                                                v104 := GenS8{F0: 160}
                                                _ = v104
                                            case "foo bar":
                                                var v105 uint = 0
                                                _ = v105
                                            }
                                            v5 += (fi91 & genFn1(174, 42, 1))
                                        } else {
                                            v5 -= genFn1(255, -128, 127)
                                            v102 = ("foo bar" + "")
                                            v101 = v101
                                            v106 := GenS4{F0: 2, F1: -9223372036854775808, F2: 255}
                                            _ = v106
                                        }
                                    }
                                    var v107 int64 = 229
                                    _ = v107
                                }
                                if (127 > 1) {
                                    switch 222 {
                                    case -102:
                                        v108 := GenS8{F0: 863}
                                        _ = v108
                                    case -594:
                                        v109 := GenS4{F0: 2, F1: -874, F2: 2}
                                        var v110 uint32 = 0
                                        fmt.Println(v101)
                                        _ = v109
                                        _ = v110
                                    case -2147483648:
                                        v2 = append(v2, 347)
                                    default:
                                        var v111 int8 = 1
                                        _ = v111
                                    }
                                } else {
                                    var v112 int = v5
                                    v101 = v101
                                    _ = v112
                                }
                                _ = v101
                                _ = v102
                                _ = v103
                            }
                        }
                    default:
                        v5 -= fi92
                    }
                    _ = fi92
                    _ = v99
                }
                if false {
                    var v113 uint = 65535
                    if ((0.1 - 0.0) < 3.14159) {
                        v113 = (v113 - 255)
                        v113 += (^(v113))
                    }
                    _ = len(v4)
                    _ = v113
                } else {
                    if true {
                        v5 -= genFn1(160, -23, 127)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        switch "" {
                        case "a":
                            v114 := GenS5{F0: (^(fi91)), F1: 1000, F2: 255}
                            _ = v114
                        case "hello":
                            var v115 uint64 = 4294967295
                            _ = v115
                        }
                    } else {
                        v116 := GenS11{F0: 42}
                        _ = v116
                    }
                    v117 := GenS7{F0: v5}
                    _ = v117
                }
                v1 = -2147483648
            }
            v118 := GenS4{F0: 599, F1: -860, F2: 255}
            for fi119 := 0; fi119 < 15; fi119++ {
                v5 -= v5
                var v120 uint16 = 65535
                v121 := GenS3{F0: v1}
                if ("\n\t" <= ("hello" + "a")) {
                    if (0 >= -128) {
                        var v122 int16 = 828
                        v122 = 0
                        v1 -= (v1 & 556)
                        v120 += 452
                        _ = v122
                    }
                    var v123 string = "a"
                    var v124 int16 = 0
                    _ = v123
                    _ = v124
                } else {
                    v1 = v1
                    v1 += v1
                    switch ("" + "foo bar") {
                    case "foo bar":
                        if (579 < 0) {
                            v125 := GenS12{F0: fi91}
                            _ = v125
                        }
                    }
                }
                v120 = v120
                v1 -= (^(v1))
                var v126 int16 = 42
                v127 := GenS13{F0: 712, F1: "", F2: (-32768 <= 1000)}
                _ = fi119
                _ = v120
                _ = v121
                _ = v126
                _ = v127
            }
            _ = fi91
            _ = v118
        }
    } else {
        for fi128 := 0; fi128 < 15; fi128++ {
            for fi129 := 0; fi129 < 4; fi129++ {
                v1 = 172
                _ = fi129
            }
            v5 += (v5 << ((657) & 63))
            fmt.Println(v1, v5, v5)
            _ = fi128
        }
        v130 := GenS8{F0: 65535}
        fmt.Println(v5, v1)
        _ = v130
    }
    delete(v4, -1000)
    var v131 int16 = 614
    switch "foo bar" {
    case "":
        switch (("foo bar" + ("hello" + "hello")) + ("" + "\n\t")) {
        case "":
            for fi132 := 0; fi132 < 20; fi132++ {
                for fi133 := 0; fi133 < 8; fi133++ {
                    var v134 uint16 = 42
                    _ = fi133
                    _ = v134
                }
                _ = fi132
            }
        case "\n\t":
            var v135 string = ("" + "")
            _ = v135
        default:
            v1 = (v1 - (v1 | v1))
        }
    case "\n\t":
        for fi136 := 0; fi136 < 16; fi136++ {
            var v137 uint64 = 1
            v1 = 312
            switch (v5 - (v5 % (((v5 - 42)) | 1))) {
            case 32768:
                fmt.Println(fi136)
            }
            _ = fi136
            _ = v137
        }
    case "hello":
        v1 += 2147483647
    }
    switch (v5 & (v5 - v5)) {
    case 1:
        v138 := GenS10{F0: 255, F1: -323}
        _ = v138
    case -348:
        switch 101 {
        case 9223372036854775807:
            var v139 int = v5
            _ = v139
        case -2147483648:
            var v140 uint = 18446744073709551615
            _ = v140
        }
    }
    v1 += (-(v1))
    var v141 int = v5
    for fi142 := 0; fi142 < 11; fi142++ {
        v141 += (v141 - -285)
        v143 := GenS6{F0: 2}
        v1 -= (v1 >> ((255) & 63))
        _ = fi142
        _ = v143
    }
    v144 := GenS14{F0: 0, F1: ""}
    switch v5 {
    case 904:
        if true {
            v5 -= v5
        } else {
            v145 := GenS5{F0: int(v1), F1: (v131 % (((^(v131))) | 1)), F2: 250}
            _ = v145
        }
    default:
        switch "hello" {
        case "\n\t":
            if (255 != 18446744073709551615) {
                _ = len(v2)
            }
        case "foo bar":
            v146 := GenS15{F0: "foo bar"}
            _ = v146
        }
    }
    for fi147 := 0; fi147 < 18; fi147++ {
        var v148 int = v5
        v2 = append(v2, 0)
        _ = fi147
        _ = v148
    }
    v2 = append(v2, 495)
    v1 -= 0
    for fi149 := 0; fi149 < 6; fi149++ {
        var v150 uint = 42
        _ = fi149
        _ = v150
    }
    fmt.Println(v131, v5)
    var v151 int32 = 283
    var v152 uint32 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v131:", v131)
    fmt.Println("v141:", v141)
    fmt.Println("v151:", v151)
    fmt.Println("v152:", v152)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v131
    _ = v141
    _ = v144
    _ = v151
    _ = v152
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
