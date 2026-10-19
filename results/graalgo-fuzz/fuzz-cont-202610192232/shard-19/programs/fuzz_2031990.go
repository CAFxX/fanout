package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 string; F2 string }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint32; F1 int32 }
type GenS7 struct { F0 int64; F1 uint32 }
type GenS8 struct { F0 int16; F1 int16; F2 uint }
type GenS9 struct { F0 int32; F1 uint32; F2 uint }
type GenS10 struct { F0 uint32; F1 bool; F2 int }
type GenS11 struct { F0 uint; F1 int32; F2 uint64 }
func genFn0(p0 string, p1 uint32) uint {
    p0 = p0
    if (127 < -1) {
        v6 := GenS1{F0: (p1 * p1), F1: 32767, F2: 127}
        v7 := GenS2{F0: 42}
        _ = v6
        _ = v7
    } else {
        if true {
            switch 0 {
            case -9223372036854775808:
                fmt.Println(p1, p0)
            case -224:
                v8 := GenS3{F0: 32767, F1: ("a" + (p0 + p0)), F2: ((p0 + p0) + (p0 + "foo bar"))}
                _ = v8
            case 9223372036854775807:
                v9 := GenS4{F0: p1, F1: 23}
                _ = v9
            default:
                fmt.Println(p1, p1)
            }
            for fi10 := 0; fi10 < 3; fi10++ {
                var v11 uint8 = 2
                fmt.Println(p0, p1)
                _ = fi10
                _ = v11
            }
            if (-1 != -1) {
                if false {
                    p0 = p0
                    v12 := []int{-527, 128}
                    _ = v12
                }
                v13 := []int{665, 32767}
                _ = v13
            } else {
                var v14 int8 = -1
                switch 128 {
                case 128:
                    if (p1 < (^(p1))) {
                        var v15 string = (p0 + "hello")
                        switch 32768 {
                        case 740:
                            for fi16 := 0; fi16 < 14; fi16++ {
                                switch 997 {
                                case 42:
                                    var v17 float64 = float64(-2.25)
                                    _ = v17
                                case 881:
                                    p1 = (p1 << ((2) & 63))
                                case -9223372036854775808:
                                    p0 = ((v15 + v15) + "hello")
                                default:
                                    for fi18 := 0; fi18 < 7; fi18++ {
                                        v14 -= v14
                                        fmt.Println(v15, v14)
                                        switch (fi16 % (((fi16 ^ fi18)) | 1)) {
                                        case 32767:
                                            var v19 int8 = v14
                                            _ = v19
                                        case 1:
                                            var v20 int16 = -671
                                            _ = v20
                                        case 603:
                                            v21 := []int{fi16, fi16, -650}
                                            if len(v21) > 0 {
                                                v21[0] = -673
                                            }
                                            _ = v21
                                        default:
                                            v14 = v14
                                        }
                                        _ = fi18
                                    }
                                }
                                _ = fi16
                            }
                        case 1:
                            var v22 uint64 = 18446744073709551615
                            _ = v22
                        case -9223372036854775808:
                            v23 := GenS1{F0: p1, F1: 1000, F2: 127}
                            _ = v23
                        default:
                            for fi24 := 0; fi24 < 2; fi24++ {
                                switch fi24 {
                                case 128:
                                    switch v15 {
                                    case "":
                                        v15 = ((p0 + p0) + "foo bar")
                                    case "hello":
                                        var v25 int = fi24
                                        _ = v25
                                    }
                                }
                                p1 -= (p1 >> ((721) & 63))
                                _ = fi24
                            }
                        }
                        for fi26 := 0; fi26 < 3; fi26++ {
                            var v27 int = (fi26 * (fi26 ^ fi26))
                            v27 = (fi26 / (((fi26 * v27)) | 1))
                            fmt.Println(v14)
                            _ = fi26
                            _ = v27
                        }
                        _ = v15
                    } else {
                        if false {
                            v28 := GenS5{F0: (v14 ^ (v14 / ((v14) | 1)))}
                            fmt.Println(v14, v14)
                            if (0 >= 2147483647) {
                                fmt.Println(p0, p1, p1)
                                v29 := GenS6{F0: (p1 >> ((2) & 63)), F1: -5}
                                _ = v29
                            }
                            _ = v28
                        } else {
                            v30 := GenS7{F0: -236, F1: 65535}
                            for fi31 := 0; fi31 < 12; fi31++ {
                                if (921 != (p1 >> ((4294967295) & 63))) {
                                    var v32 float64 = 1.5
                                    var v33 int = fi31
                                    switch 42 {
                                    case 127:
                                        v32 -= v32
                                    case 49:
                                        v33 += v33
                                    case -700:
                                        switch p0 {
                                        case "foo bar":
                                            fmt.Println(v32, fi31, fi31)
                                        case "a":
                                            var v34 int64 = 1
                                            _ = v34
                                        case "\n\t":
                                            for fi35 := 0; fi35 < 13; fi35++ {
                                                v36 := GenS1{F0: 277, F1: 1000, F2: 193}
                                                _ = fi35
                                                _ = v36
                                            }
                                        default:
                                            if false {
                                                v37 := []int{(^(v33)), 128, (^(fi31)), (v33 + (v33 / ((fi31) | 1)))}
                                                _ = v37
                                            } else {
                                                for fi38 := 0; fi38 < 10; fi38++ {
                                                    v39 := []int{(fi31 | v33), 128, -816, v33}
                                                    _ = fi38
                                                    _ = v39
                                                }
                                                for fi40 := 0; fi40 < 1; fi40++ {
                                                    v32 = v32
                                                    var v41 int32 = 678
                                                    _ = fi40
                                                    _ = v41
                                                }
                                                switch p0 {
                                                case "hello":
                                                    v42 := GenS8{F0: -767, F1: 1000, F2: 508}
                                                    _ = v42
                                                }
                                            }
                                        }
                                    default:
                                        v43 := []int{(fi31 + fi31), fi31}
                                        _ = v43
                                    }
                                    _ = v32
                                    _ = v33
                                } else {
                                    var v44 uint64 = 1
                                    var v45 uint64 = (^(v44))
                                    p0 = "a"
                                    _ = v44
                                    _ = v45
                                }
                                _ = fi31
                            }
                            if ((true || (1 <= 373)) || false) {
                                switch 0 {
                                case 9223372036854775807:
                                    v46 := GenS2{F0: 224}
                                    _ = v46
                                case -2147483648:
                                    var v47 float32 = 3.14159
                                    _ = v47
                                case 42:
                                    if (((-2.25 * 0.0) + 1.5) != 3.14159) {
                                        for fi48 := 0; fi48 < 6; fi48++ {
                                            p0 = "a"
                                            if false {
                                                var v49 int64 = -291
                                                _ = v49
                                            }
                                            for fi50 := 0; fi50 < 4; fi50++ {
                                                fmt.Println(fi48, p0)
                                                if false {
                                                    switch 128 {
                                                    case 2147483647:
                                                        if (false && (4294967295 > 901)) {
                                                            if true {
                                                                var v51 int8 = 100
                                                                v52 := GenS1{F0: p1, F1: -539, F2: 1}
                                                                _ = v51
                                                                _ = v52
                                                            }
                                                        } else {
                                                            var v53 string = (("a" + p0) + "hello")
                                                            v14 = (^(v14))
                                                            _ = v53
                                                        }
                                                    }
                                                    var v54 float64 = -2.25
                                                    fmt.Println(v14, fi50, fi48)
                                                    _ = v54
                                                }
                                                _ = fi50
                                            }
                                            _ = fi48
                                        }
                                        if false {
                                            if false {
                                                fmt.Println(p1, p1)
                                                fmt.Println(p0, p0)
                                                if true {
                                                    p1 = 2147483647
                                                    var v55 int16 = -1
                                                    switch 203 {
                                                    case -2147483648:
                                                        v14 = (^(v14))
                                                    default:
                                                        var v56 bool = true
                                                        _ = v56
                                                    }
                                                    _ = v55
                                                } else {
                                                    var v57 uint8 = 128
                                                    v57 += (v57 & 165)
                                                    v58 := GenS9{F0: -2147483648, F1: 1, F2: 691}
                                                    _ = v57
                                                    _ = v58
                                                }
                                            }
                                            var v59 uint64 = 931
                                            _ = v59
                                        } else {
                                            switch p0 {
                                            case "\n\t":
                                                v14 = 118
                                            }
                                            var v60 string = ((p0 + "\n\t") + (("a" + p0) + "hello"))
                                            for fi61 := 0; fi61 < 4; fi61++ {
                                                for fi62 := 0; fi62 < 12; fi62++ {
                                                    v63 := []int{fi62}
                                                    _ = fi62
                                                    _ = v63
                                                }
                                                v64 := GenS9{F0: 0, F1: 101, F2: 65535}
                                                if false {
                                                    var v65 uint32 = p1
                                                    var v66 int32 = -1
                                                    _ = v65
                                                    _ = v66
                                                } else {
                                                    for fi67 := 0; fi67 < 8; fi67++ {
                                                        fmt.Println(v14)
                                                        _ = fi67
                                                    }
                                                    v68 := []int{(fi61 & fi61)}
                                                    _ = v68
                                                }
                                                _ = fi61
                                                _ = v64
                                            }
                                            _ = v60
                                        }
                                    }
                                default:
                                    v14 = v14
                                }
                                if ((!(false || true)) || false) {
                                    v14 -= -70
                                    switch 128 {
                                    case 127:
                                        var v69 string = (p0 + (p0 + p0))
                                        _ = v69
                                    default:
                                        v70 := []int{-2147483648}
                                        _ = v70
                                    }
                                } else {
                                    if false {
                                        v71 := []int{4294967295, 1, 0, 0}
                                        _ = v71
                                    } else {
                                        v72 := []int{350}
                                        _ = v72
                                    }
                                    var v73 int32 = 42
                                    _ = v73
                                }
                            } else {
                                fmt.Println(p1, v14)
                            }
                            _ = v30
                        }
                        var v74 int = 4294967295
                        switch p0 {
                        case "foo bar":
                            for fi75 := 0; fi75 < 11; fi75++ {
                                v76 := []int{fi75, fi75, 295, int(v14)}
                                _ = fi75
                                _ = v76
                            }
                        case "\n\t":
                            for fi77 := 0; fi77 < 7; fi77++ {
                                v78 := GenS10{F0: (^(p1)), F1: ((!false) && ((p0 + (((((p0 + "\n\t") + (((("foo bar" + ((((p0 + (((p0 + p0) + "foo bar") + "foo bar")) + p0) + "foo bar") + "\n\t")) + "") + "\n\t") + "hello")) + "foo bar") + p0) + p0)) > "a")), F2: v74}
                                var v79 int64 = 588
                                switch "foo bar" {
                                case "hello":
                                    if true {
                                        switch v74 {
                                        case 1:
                                            fmt.Println(v14)
                                        case 127:
                                            if false {
                                                for fi80 := 0; fi80 < 2; fi80++ {
                                                    var v81 uint64 = 18446744073709551615
                                                    _ = fi80
                                                    _ = v81
                                                }
                                            }
                                        case -2147483648:
                                            var v82 float32 = ((((0.0 / -2.25) + 1.5) + 0.1) - 3.14159)
                                            _ = v82
                                        default:
                                            if (!true) {
                                                for fi83 := 0; fi83 < 3; fi83++ {
                                                    p0 = p0
                                                    _ = fi83
                                                }
                                                for fi84 := 0; fi84 < 13; fi84++ {
                                                    v85 := []int{v74, fi84, fi77, 0}
                                                    _ = fi84
                                                    _ = v85
                                                }
                                            }
                                        }
                                        for fi86 := 0; fi86 < 14; fi86++ {
                                            var v87 float64 = 3.14159
                                            var v88 uint16 = 2
                                            if (!true) {
                                                var v89 int64 = v79
                                                fmt.Println(v14, v88, v79)
                                                _ = v89
                                            }
                                            _ = fi86
                                            _ = v87
                                            _ = v88
                                        }
                                        v74 = fi77
                                    }
                                default:
                                    for fi90 := 0; fi90 < 2; fi90++ {
                                        v91 := []int{fi77, fi90}
                                        _ = fi90
                                        _ = v91
                                    }
                                }
                                _ = fi77
                                _ = v78
                                _ = v79
                            }
                        case "":
                            var v92 float32 = 1e10
                            _ = v92
                        }
                        _ = v74
                    }
                case 127:
                    v93 := []int{899, 124, -95}
                    _ = v93
                }
                var v94 int8 = (v14 & v14)
                _ = v14
                _ = v94
            }
        }
        var v95 uint32 = p1
        _ = v95
    }
    _ = p0
    _ = p1
    return 373
}

func main() {
    var v1 float64 = -2.25
    v2 := make(map[uint32]int)
    v3 := make(map[int]bool)
    v4 := make(map[int]uint32)
    var v5 float32 = (-2.25 - (1e10 / 1.5))
    v1 -= float64(0.1)
    switch -722 {
    case 32768:
        var v96 uint = 838
        _ = v96
    case 127:
        var v97 bool = (4294967295 > 175)
        _ = v97
    case 42:
        if (false && false) {
            v5 -= ((v5 / (v5 / 1.0)) / (v5 + -2.25))
        }
    default:
        var v98 int16 = 32767
        _ = v98
    }
    v99 := GenS8{F0: 42, F1: -1, F2: 897}
    var v100 uint32 = 810
    var v101 string = "a"
    if false {
        var v102 uint8 = 42
        _ = v102
    } else {
        fmt.Println(v101, v100)
        for fi103 := 0; fi103 < 19; fi103++ {
            for fi104 := 0; fi104 < 7; fi104++ {
                var v105 int16 = 0
                fmt.Println(v5, v1)
                var v106 uint8 = 127
                var v107 int8 = 100
                var v108 uint64 = 2
                if true {
                    v109 := v2[42]
                    var v110 int = v109
                    delete(v2, 71)
                    v100 = (v100 + 926)
                    if (!(0 > 32767)) {
                        if ((v108 != v108) && (!("\n\t" > (v101 + "a")))) {
                            if true {
                                switch fi104 {
                                case 42:
                                    var v111 int = v110
                                    fmt.Println(v101)
                                    _ = v111
                                }
                                v5 = float32(1.5)
                                v112 := GenS11{F0: 257, F1: -147, F2: 18446744073709551615}
                                v105 = -32768
                                _ = v112
                            }
                        } else {
                            fmt.Println(v5, fi104)
                        }
                        switch "hello" {
                        case "\n\t":
                            var v113 bool = false
                            _ = v113
                        case "":
                            if (!false) {
                                var v114 int16 = (v105 / ((42) | 1))
                                v1 = v1
                                _ = v114
                            } else {
                                v115 := GenS9{F0: -1, F1: 2147483647, F2: 1}
                                var v116 int = v109
                                switch (v109 ^ fi104) {
                                case -1:
                                    switch (v109 << ((42) & 63)) {
                                    case 32768:
                                        v116 = (v109 - 2147483647)
                                    }
                                }
                                if ((fi103 / (((fi103 | v109)) | 1)) > (fi104 / ((-2147483648) | 1))) {
                                    var v117 uint64 = (v108 - (v108 >> ((255) & 63)))
                                    v118 := GenS6{F0: (v100 - (v100 + v100)), F1: -543}
                                    _ = v117
                                    _ = v118
                                }
                                switch int(v106) {
                                case 4294967295:
                                    delete(v3, (^(v109)))
                                case -2147483648:
                                    fmt.Println(v105, v116, v101)
                                case -9223372036854775808:
                                    var v119 int16 = -646
                                    _ = v119
                                }
                                _ = v115
                                _ = v116
                            }
                        case "hello":
                            if ((((fi103 * v110) < fi104) && false) && false) {
                                v120 := []int{-590}
                                var v121 uint = 429
                                var v122 int64 = 1
                                _ = v120
                                _ = v121
                                _ = v122
                            } else {
                                v100 += 732
                            }
                        }
                    }
                    _ = v109
                    _ = v110
                } else {
                    fmt.Println(v1)
                    v123 := []int{fi104, fi104, (fi103 % (((fi104 / ((fi104) | 1))) | 1))}
                    if len(v123) >= 2 {
                        v123 = v123[1:len(v123)]
                    }
                    _ = v123
                }
                _ = fi104
                _ = v105
                _ = v106
                _ = v107
                _ = v108
            }
            _ = fi103
        }
    }
    v124 := v2[(v100 % ((v100) | 1))]
    if false {
        switch (("\n\t" + v101) + "hello") {
        case "foo bar":
            var v125 int64 = -474
            _ = v125
        }
        v126 := []int{v124, 42, 32767, v124}
        if false {
            for fi127 := 0; fi127 < 16; fi127++ {
                fmt.Println(v1)
                for fi128 := 0; fi128 < 15; fi128++ {
                    fmt.Println(fi127, v1)
                    fmt.Println(fi127, v101)
                    if len(v126) >= 2 {
                        v126 = v126[1:len(v126)]
                    }
                    _ = fi128
                }
                if ((!false) || true) {
                    v1 = (0.0 + v1)
                    var v129 float32 = (float32((v1 / v1)) + v5)
                    _ = v129
                }
                _ = fi127
            }
        }
        _ = v126
    } else {
        v101 = "\n\t"
        var v130 int16 = 1
        var v131 float32 = ((v5 + (v5 / v5)) + 0.0)
        _ = v130
        _ = v131
    }
    switch v101 {
    case "hello":
        for fi132 := 0; fi132 < 17; fi132++ {
            v133 := GenS3{F0: 0, F1: "\n\t", F2: v101}
            _ = fi132
            _ = v133
        }
    case "foo bar":
        switch (^(v124)) {
        case -2147483648:
            switch 32767 {
            case 571:
                var v134 int8 = 0
                _ = v134
            case 1:
                fmt.Println(v124, v5)
            case 128:
                var v135 string = "\n\t"
                _ = v135
            }
        case 32768:
            if (true && false) {
                for fi136 := 0; fi136 < 12; fi136++ {
                    var v137 int = 42
                    var v138 float64 = ((0.0 / 0.1) + (1.5 - float64(float32(v1))))
                    _ = fi136
                    _ = v137
                    _ = v138
                }
            }
        case 484:
            if true {
                for fi139 := 0; fi139 < 18; fi139++ {
                    fmt.Println(v101, v5, v124)
                    _ = fi139
                }
            }
        default:
            fmt.Println(v1, v124)
        }
    }
    switch (v124 - v124) {
    case 32767:
        v140 := GenS5{F0: 1}
        _ = v140
    case 204:
        v141 := GenS6{F0: (v100 << ((844) & 63)), F1: 1}
        _ = v141
    case -2147483648:
        for fi142 := 0; fi142 < 12; fi142++ {
            var v143 uint32 = (v100 & v100)
            _ = fi142
            _ = v143
        }
    default:
        v5 += ((float32((1.5 - (-2.25 - (0.0 / v1)))) / float32(v1)) * 1e10)
    }
    v5 -= v5
    var v144 int = v124
    v145 := GenS8{F0: -510, F1: 1, F2: 42}
    if (!true) {
        switch v101 {
        case "hello":
            if (!true) {
                v124 += v124
            }
        case "a":
            v146 := v4[v124]
            _ = v146
        }
    }
    var v147 uint16 = 0
    v148 := GenS8{F0: -1, F1: -32768, F2: 65535}
    if (1 <= 0) {
        v149 := []int{v144}
        _ = v149
    } else {
        v124 -= 351
        v150 := GenS6{F0: uint32(v147), F1: 42}
        delete(v3, 4294967295)
        _ = v150
    }
    v151 := []int{(v144 << ((uint(v147)) & 63)), v124, v124}
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v101:", v101)
    fmt.Println("v124:", v124)
    fmt.Println("v144:", v144)
    fmt.Println("v147:", v147)
    fmt.Println("slicelen v151:", len(v151), v151)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v100
    _ = v101
    _ = v124
    _ = v144
    _ = v145
    _ = v147
    _ = v148
    _ = v151
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v99
}
