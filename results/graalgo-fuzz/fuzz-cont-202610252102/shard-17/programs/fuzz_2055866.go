package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int32; F1 string }
type GenS8 struct { F0 uint8; F1 int16 }
type GenS9 struct { F0 int32; F1 bool }
type GenS10 struct { F0 int8; F1 int64 }
type GenS11 struct { F0 uint16; F1 int }
type GenS12 struct { F0 uint32; F1 int32 }
type GenS13 struct { F0 int32; F1 uint; F2 uint64 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 string }
func genFn0(p0 int16, p1 int8, p2 int8) int8 {
    switch "\n\t" {
    case "foo bar":
        p0 -= p0
    case "":
        v6 := GenS1{F0: 42, F1: 42}
        _ = v6
    case "hello":
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := GenS1{F0: 637, F1: 2}
            _ = fi7
            _ = v8
        }
    }
    fmt.Println(p1)
    if true {
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 int = fi9
            var v11 float32 = -2.25
            _ = fi9
            _ = v10
            _ = v11
        }
        v12 := []int{738}
        _ = v12
    }
    v13 := GenS1{F0: 1, F1: 121}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 127
}
func genFn1(p0 int8) uint8 {
    v14 := GenS1{F0: 18446744073709551615, F1: 255}
    v15 := GenS2{F0: -9223372036854775808, F1: 2147483647, F2: 498}
    v16 := GenS1{F0: 42, F1: 255}
    p0 += (p0 / (((p0 * (p0 >> ((1) & 63)))) | 1))
    _ = p0
    _ = v14
    _ = v15
    _ = v16
    return 148
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 1000
    var v3 int16 = (v2 << ((18446744073709551615) & 63))
    var v4 float64 = 1e10
    var v5 int = 32767
    v3 += (v2 / ((42) | 1))
    var v17 int16 = v2
    if false {
        fmt.Println(v2)
        if true {
            for fi18 := 0; fi18 < 5; fi18++ {
                var v19 uint64 = 0
                v20 := GenS1{F0: v19, F1: 128}
                if (v19 >= v19) {
                    for fi21 := 0; fi21 < 11; fi21++ {
                        switch ("a" + "") {
                        case "foo bar":
                            v22 := GenS1{F0: (v19 & 18446744073709551615), F1: 127}
                            _ = v22
                        case "hello":
                            _ = len(v1)
                        case "":
                            var v23 uint = 255
                            _ = v23
                        default:
                            v4 += v4
                        }
                        _ = fi21
                    }
                } else {
                    v2 = v3
                }
                _ = fi18
                _ = v19
                _ = v20
            }
            for fi24 := 0; fi24 < 2; fi24++ {
                var v25 int16 = -1
                _ = fi24
                _ = v25
            }
        }
    }
    var v26 float64 = 1.5
    v26 += v26
    v27 := GenS3{F0: 1, F1: 42}
    switch (-(v5)) {
    case 32767:
        var v28 int8 = 127
        _ = v28
    case 1:
        v3 += (v3 + (v2 % ((32767) | 1)))
    default:
        var v29 uint32 = 0
        _ = v29
    }
    if (v5 > (^(v5))) {
        if ((!((!true) || false)) || (false || true)) {
            switch (v5 - 2147483647) {
            case -573:
                var v30 uint8 = 1
                _ = v30
            case 9223372036854775807:
                if false {
                    if false {
                        for fi31 := 0; fi31 < 2; fi31++ {
                            var v32 float64 = (v4 / (v4 * v4))
                            fmt.Println(fi31, v4, v17)
                            if false {
                                var v33 uint64 = 971
                                for fi34 := 0; fi34 < 3; fi34++ {
                                    v35 := []int{v5, v5, (fi31 | 9223372036854775807)}
                                    var v36 int64 = 1
                                    _ = fi34
                                    _ = v35
                                    _ = v36
                                }
                                _ = v33
                            }
                            _ = fi31
                            _ = v32
                        }
                        v37 := GenS4{F0: 2147483647}
                        _ = v37
                    } else {
                        if (0 > 105) {
                            var v38 float64 = 3.14159
                            _ = v38
                        } else {
                            for fi39 := 0; fi39 < 1; fi39++ {
                                v40 := GenS2{F0: 9223372036854775807, F1: 2147483648, F2: 4294967295}
                                for fi41 := 0; fi41 < 5; fi41++ {
                                    switch (^(fi41)) {
                                    case -2147483648:
                                        if (genFn1(-111) > 0) {
                                            switch (fi39 & 229) {
                                            case 747:
                                                var v42 int = (fi39 + (fi39 << ((4294967295) & 63)))
                                                if (-2.25 != 0.1) {
                                                    v43 := GenS5{F0: 2147483647}
                                                    _ = v43
                                                } else {
                                                    v26 += 1.5
                                                    v26 -= v26
                                                    var v44 float32 = -2.25
                                                    v45 := []int{128, fi41, (v42 & (^(v5)))}
                                                    switch -603 {
                                                    case -1:
                                                        var v46 int16 = -32768
                                                        _ = v46
                                                    case 0:
                                                        if true {
                                                            v17 -= v17
                                                            fmt.Println(v3)
                                                            if ((!false) || ("\n\t" > "foo bar")) {
                                                                var v47 uint32 = 2147483647
                                                                switch ("foo bar" + "a") {
                                                                case "hello":
                                                                    fmt.Println(v17)
                                                                }
                                                                _ = v47
                                                            } else {
                                                                if ((3.14159 == (float64(((v44 / 1.0) - v44)) * 3.14159)) || ((0 >= 65535) && ((42 > 1) || (255 != genFn1(55))))) {
                                                                    v45 = append(v45, (v5 % ((-399) | 1)))
                                                                }
                                                            }
                                                        } else {
                                                            if true {
                                                                var v48 bool = false
                                                                v42 += (fi39 % (((^(fi39))) | 1))
                                                                if false {
                                                                    v2 = 607
                                                                    v2 -= v3
                                                                }
                                                                _ = v48
                                                            } else {
                                                                v42 = (fi39 | fi39)
                                                                var v49 int32 = -1
                                                                v50 := GenS6{F0: v49}
                                                                if (true && false) {
                                                                    fmt.Println(v2)
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        v51 := GenS7{F0: (^(v49)), F1: ("hello" + (("a" + "\n\t") + ("\n\t" + "")))}
                                                                        _ = v51
                                                                    }
                                                                    v42 = (fi41 << ((1) & 63))
                                                                    switch (^(v5)) {
                                                                    case -364:
                                                                        var v52 string = "\n\t"
                                                                        _ = v52
                                                                    case 128:
                                                                        v53 := GenS1{F0: 42, F1: 0}
                                                                        _ = v53
                                                                    case 32767:
                                                                        var v54 uint16 = 255
                                                                        _ = v54
                                                                    }
                                                                    if false {
                                                                        v42 = fi41
                                                                        if len(v45) > 0 {
                                                                            v45[0] = (-(v5))
                                                                        }
                                                                    } else {
                                                                        if len(v45) > 0 {
                                                                            v45[0] = fi41
                                                                        }
                                                                        v17 = 1
                                                                        v55 := GenS7{F0: (-(v49)), F1: ("a" + ("a" + (("" + "foo bar") + ("hello" + "a"))))}
                                                                        var v56 int64 = 1
                                                                        _ = v55
                                                                        _ = v56
                                                                    }
                                                                    if false {
                                                                        var v57 uint64 = 2
                                                                        _ = v57
                                                                    } else {
                                                                        v49 += v49
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            v26 = (0.0 * 0.0)
                                                                        default:
                                                                            var v58 uint16 = 255
                                                                            _ = v58
                                                                        }
                                                                        var v59 string = "foo bar"
                                                                        _ = v59
                                                                    }
                                                                    var v60 bool = true
                                                                    v44 -= v44
                                                                    delete(v1, "foo bar")
                                                                    _ = v60
                                                                } else {
                                                                    _ = len(v1)
                                                                }
                                                                v61 := GenS3{F0: 18446744073709551615, F1: fi41}
                                                                _ = v49
                                                                _ = v50
                                                                _ = v61
                                                            }
                                                            v2 = v2
                                                        }
                                                    case 32767:
                                                        var v62 bool = true
                                                        _ = v62
                                                    }
                                                    _ = v44
                                                    _ = v45
                                                }
                                                if (genFn0(v3, 0, 87) != 127) {
                                                    if (248 <= 5) {
                                                        if (0.0 <= 1e10) {
                                                            v63 := []int{2147483647, v5, v42}
                                                            var v64 uint = 65535
                                                            _ = v63
                                                            _ = v64
                                                        } else {
                                                            v65 := []int{}
                                                            v4 = 0.1
                                                            fmt.Println(v5, v42, fi41)
                                                            if false {
                                                                v66 := GenS3{F0: 2, F1: fi39}
                                                                _ = v66
                                                            }
                                                            _ = v65
                                                        }
                                                        v1["foo bar"] = 1
                                                        switch (fi39 >> ((2) & 63)) {
                                                        case -1:
                                                            fmt.Println(v26, v26, v2)
                                                        case 127:
                                                            v5 -= (-(v42))
                                                            fmt.Println(fi41)
                                                        }
                                                        v67 := v1["a"]
                                                        if true {
                                                            v4 = v4
                                                            v68 := v1[("foo bar" + "a")]
                                                            _ = len(v1)
                                                            fmt.Println(fi39)
                                                            switch v42 {
                                                            case 987:
                                                                if true {
                                                                    if false {
                                                                        v5 += v42
                                                                    } else {
                                                                        var v69 int = v42
                                                                        _ = v69
                                                                    }
                                                                    var v70 uint32 = (v68 << ((393) & 63))
                                                                    _ = v70
                                                                }
                                                            default:
                                                                v67 = 1
                                                            }
                                                            _ = v68
                                                        }
                                                        delete(v1, "foo bar")
                                                        fmt.Println(v26)
                                                        _ = v67
                                                    } else {
                                                        var v71 bool = true
                                                        _ = v71
                                                    }
                                                    v72 := GenS4{F0: 4294967295}
                                                    v1["a"] = 549
                                                    _ = v72
                                                }
                                                v73 := GenS8{F0: 2, F1: -791}
                                                _ = v42
                                                _ = v73
                                            default:
                                                if (4294967295 > 547) {
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        v74 := GenS9{F0: -1, F1: true}
                                                        _ = v74
                                                    case "hello":
                                                        v17 -= (^(v17))
                                                        v75 := []int{fi39, v5}
                                                        _ = v75
                                                    case "\n\t":
                                                        if true {
                                                            var v76 uint64 = 423
                                                            v26 -= 0.1
                                                            v77 := []int{}
                                                            _ = v76
                                                            _ = v77
                                                        }
                                                    default:
                                                        v17 += v3
                                                        v78 := GenS10{F0: 0, F1: 620}
                                                        _ = v78
                                                    }
                                                    v79 := GenS7{F0: 252, F1: "foo bar"}
                                                    v80 := GenS11{F0: 395, F1: fi39}
                                                    _ = v79
                                                    _ = v80
                                                }
                                            }
                                            var v81 uint = 18446744073709551615
                                            var v82 int8 = 100
                                            _ = v81
                                            _ = v82
                                        }
                                    case 4294967295:
                                        v83 := []int{}
                                        _ = v83
                                    default:
                                        var v84 uint32 = 1
                                        _ = v84
                                    }
                                    fmt.Println(fi41, v5, v2)
                                    switch v5 {
                                    case 127:
                                        v85 := GenS2{F0: -9223372036854775808, F1: 2, F2: 1}
                                        _ = v85
                                    case -819:
                                        if false {
                                            if false {
                                                var v86 int8 = -128
                                                var v87 int64 = 1
                                                _ = v86
                                                _ = v87
                                            }
                                            v2 += v17
                                        }
                                    case -9223372036854775808:
                                        var v88 float64 = (v4 * v26)
                                        _ = v88
                                    default:
                                        v89 := GenS12{F0: 2, F1: 2147483647}
                                        if false {
                                            var v90 uint = 348
                                            _ = v90
                                        } else {
                                            v91 := []int{4294967295, 128, 127, 0}
                                            var v92 float64 = (1.5 + v26)
                                            var v93 int64 = 9223372036854775807
                                            _ = v91
                                            _ = v92
                                            _ = v93
                                        }
                                        var v94 float32 = (((1.5 / -2.25) + ((0.1 + ((3.14159 + 1e10) - -2.25)) + float32(float64((-2.25 + (float32(0.0) - -2.25)))))) + 3.14159)
                                        _ = v89
                                        _ = v94
                                    }
                                    _ = fi41
                                }
                                switch "\n\t" {
                                case "":
                                    v95 := []int{-2147483648, 42}
                                    _ = v95
                                case "a":
                                    var v96 uint32 = 1
                                    _ = v96
                                case "\n\t":
                                    switch fi39 {
                                    case -495:
                                        var v97 uint64 = 2
                                        _ = v97
                                    case 2147483647:
                                        var v98 int64 = -572
                                        _ = v98
                                    default:
                                        v26 = (v26 / (v4 + (v4 / float64(0.1))))
                                    }
                                default:
                                    v99 := GenS7{F0: 1, F1: "a"}
                                    _ = v99
                                }
                                _ = fi39
                                _ = v40
                            }
                            var v100 uint8 = 255
                            switch 9223372036854775807 {
                            case 1:
                                v101 := GenS7{F0: -2147483648, F1: "foo bar"}
                                _ = v101
                            case 32768:
                                switch "foo bar" {
                                case "\n\t":
                                    v102 := []int{int(v3), (v5 / (((v5 * v5)) | 1)), (v5 >> ((107) & 63))}
                                    _ = v102
                                default:
                                    switch (-(v5)) {
                                    case 127:
                                        fmt.Println(v100, v26)
                                    case 32767:
                                        for fi103 := 0; fi103 < 17; fi103++ {
                                            fmt.Println(fi103)
                                            v104 := GenS6{F0: -395}
                                            _ = fi103
                                            _ = v104
                                        }
                                    }
                                }
                            default:
                                switch v5 {
                                case 127:
                                    v5 -= -2147483648
                                default:
                                    if false {
                                        _ = len(v1)
                                        v17 = v17
                                    }
                                }
                            }
                            _ = v100
                        }
                    }
                    if (0 < 100) {
                        switch 2147483647 {
                        case 127:
                            var v105 uint32 = 1
                            _ = v105
                        case 1:
                            v5 = v5
                        default:
                            for fi106 := 0; fi106 < 18; fi106++ {
                                if false {
                                    var v107 uint = 0
                                    if (!(v5 >= fi106)) {
                                        if true {
                                            v108 := []int{9223372036854775807, v5}
                                            v2 += v2
                                            _ = v108
                                        }
                                        var v109 int = 514
                                        _ = v109
                                    } else {
                                        v110 := GenS8{F0: 128, F1: v2}
                                        _ = v110
                                    }
                                    _ = v107
                                } else {
                                    if ((v17 << ((0) & 63)) <= -32768) {
                                        v4 -= float64(3.14159)
                                    }
                                }
                                v5 += v5
                                _ = fi106
                            }
                        }
                        if (true || (100 != 0)) {
                            fmt.Println(v26)
                            _ = len(v1)
                            v3 -= (v2 * (-(v2)))
                        } else {
                            var v111 uint = 2
                            var v112 int16 = (v2 << ((v111) & 63))
                            _ = v111
                            _ = v112
                        }
                        for fi113 := 0; fi113 < 4; fi113++ {
                            for fi114 := 0; fi114 < 3; fi114++ {
                                v115 := []int{756}
                                var v116 uint8 = genFn1(-128)
                                _ = len(v1)
                                _ = fi114
                                _ = v115
                                _ = v116
                            }
                            if ((!(v3 < v3)) || (false && true)) {
                                var v117 uint = 351
                                _ = v117
                            }
                            _ = fi113
                        }
                    } else {
                        for fi118 := 0; fi118 < 6; fi118++ {
                            v4 += 1.5
                            v5 -= (fi118 * (v5 & fi118))
                            _ = fi118
                        }
                        var v119 float64 = ((v26 - (1.5 + v4)) * ((v4 - ((v4 - (v4 + (((v26 / (3.14159 / v26)) + v4) * v4))) + v4)) / 0.1))
                        var v120 uint32 = 4294967295
                        _ = v119
                        _ = v120
                    }
                    for fi121 := 0; fi121 < 7; fi121++ {
                        var v122 int8 = 100
                        _ = fi121
                        _ = v122
                    }
                }
            }
            for fi123 := 0; fi123 < 7; fi123++ {
                if (0 <= 4294967295) {
                    v124 := GenS4{F0: 4294967295}
                    v125 := v1["foo bar"]
                    for fi126 := 0; fi126 < 6; fi126++ {
                        var v127 uint32 = v125
                        v17 = 1
                        switch (v5 << ((2) & 63)) {
                        case 32768:
                            v4 -= (v26 + ((v4 * (3.14159 / (3.14159 - v26))) * ((-2.25 + v26) / -2.25)))
                        case 2147483647:
                            if (0 < genFn0((v3 ^ 959), -128, 0)) {
                                switch (v5 >> ((65535) & 63)) {
                                case -387:
                                    v17 = v2
                                case 0:
                                    v127 -= 394
                                default:
                                    v128 := GenS10{F0: 127, F1: 7}
                                    _ = v128
                                }
                            } else {
                                v129 := v1[("a" + ("hello" + "foo bar"))]
                                _ = v129
                            }
                            if (!(32767 > 0)) {
                                v130 := GenS13{F0: 218, F1: 380, F2: 42}
                                _ = v130
                            }
                        }
                        _ = fi126
                        _ = v127
                    }
                    _ = v124
                    _ = v125
                }
                var v131 uint64 = 148
                _ = fi123
                _ = v131
            }
            var v132 int = v5
            _ = v132
        }
        var v133 string = "foo bar"
        var v134 int16 = (-(v17))
        _ = v133
        _ = v134
    } else {
        fmt.Println(v2, v4)
    }
    var v135 int32 = 572
    v136 := []int{v5, v5}
    var v137 int64 = -1
    switch ("a" + "foo bar") {
    case "":
        v138 := GenS4{F0: 2147483647}
        _ = v138
    case "hello":
        if false {
            v17 += v3
        } else {
            if true {
                if (615 < 1) {
                    var v139 float32 = -2.25
                    if (v26 == v26) {
                        v5 = (-(v5))
                        if false {
                            v136 = append(v136, int(v135))
                            switch 32768 {
                            case 32768:
                                for fi140 := 0; fi140 < 18; fi140++ {
                                    v2 += (v17 + v3)
                                    if (334 > 2) {
                                        if len(v136) >= 2 {
                                            v136 = v136[1:len(v136)]
                                        }
                                        var v141 int32 = (v135 / ((v135) | 1))
                                        for fi142 := 0; fi142 < 8; fi142++ {
                                            v4 = v26
                                            var v143 int = 32767
                                            _ = fi142
                                            _ = v143
                                        }
                                        _ = v141
                                    }
                                    _ = fi140
                                }
                            default:
                                for fi144 := 0; fi144 < 7; fi144++ {
                                    v3 += (v2 + v3)
                                    var v145 int8 = -1
                                    _ = fi144
                                    _ = v145
                                }
                            }
                        }
                        var v146 uint8 = 127
                        _ = v146
                    }
                    var v147 bool = true
                    _ = v139
                    _ = v147
                }
                v148 := GenS14{F0: 180}
                _ = v148
            }
            v135 = (v135 + (v135 | (v135 >> ((255) & 63))))
            v5 += (v5 >> ((2) & 63))
        }
    case "foo bar":
        if len(v136) >= 2 {
            v136 = v136[1:len(v136)]
        }
    }
    if false {
        v149 := GenS15{F0: ("" + "\n\t")}
        var v150 int16 = (v2 + (v2 % (((v2 % (((^(v3))) | 1))) | 1)))
        var v151 int16 = v150
        _ = v149
        _ = v150
        _ = v151
    }
    var v152 uint16 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v135:", v135)
    fmt.Println("slicelen v136:", len(v136), v136)
    fmt.Println("v137:", v137)
    fmt.Println("v152:", v152)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v135
    _ = v136
    _ = v137
    _ = v152
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
