package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 uint; F2 uint }
type GenS5 struct { F0 uint64; F1 uint }
type GenS6 struct { F0 int; F1 string; F2 uint }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int64; F1 uint }
type GenS10 struct { F0 int8; F1 bool; F2 uint8 }
func genFn0(p0 bool, p1 uint16) int64 {
    v5 := GenS1{F0: -348, F1: (^(p1)), F2: 0}
    if true {
        fmt.Println(p0, p1)
        if (1 >= 1) {
            p1 = (p1 - (^(p1)))
            p0 = true
            for fi6 := 0; fi6 < 2; fi6++ {
                fmt.Println(p0, fi6)
                _ = fi6
            }
        } else {
            if false {
                for fi7 := 0; fi7 < 13; fi7++ {
                    var v8 int8 = 100
                    if p0 {
                        var v9 float64 = 0.0
                        v9 = v9
                        v10 := []int{fi7, (fi7 >> ((4294967295) & 63)), 32768}
                        _ = v10
                        _ = v9
                    }
                    _ = fi7
                    _ = v8
                }
            }
        }
    } else {
        fmt.Println(p1)
        var v11 uint8 = 1
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v5
    return -9223372036854775808
}
func genFn1() uint32 {
    if (((true && (18446744073709551615 == 42)) || ((0.0 + -2.25) >= (float32(3.14159) + -2.25))) && (42 != 72)) {
        if (false && (2 < 42)) {
            var v12 uint64 = 18446744073709551615
            var v13 float64 = ((1.5 * ((1.5 + (1e10 + 3.14159)) / (-2.25 - 0.1))) - (3.14159 + 3.14159))
            _ = v12
            _ = v13
        } else {
            v14 := GenS2{F0: -32768, F1: 669}
            var v15 int16 = 348
            for fi16 := 0; fi16 < 20; fi16++ {
                v15 = (^(v15))
                v15 -= -572
                _ = fi16
            }
            _ = v14
            _ = v15
        }
        switch ("" + "foo bar") {
        case "\n\t":
            var v17 string = ""
            _ = v17
        }
    } else {
        var v18 uint = 647
        v19 := []int{-9223372036854775808, 2147483647}
        _ = v18
        _ = v19
    }
    var v20 uint8 = 42
    v21 := GenS2{F0: -1, F1: 1}
    _ = v20
    _ = v21
    return 826
}

func main() {
    var v1 float32 = ((-2.25 + float32((float64(0.0) - (3.14159 / (-2.25 * ((3.14159 - ((-2.25 + ((1e10 + 1.5) * (0.1 / (1.5 + ((-2.25 - 0.1) / 0.1))))) * (0.1 + 1.5))) - ((1e10 + ((0.1 * 1.5) / 3.14159)) - -2.25))))))) / 1.0)
    var v2 int16 = 0
    v3 := make(map[string]string)
    v4 := make(map[int]string)
    if false {
        for fi22 := 0; fi22 < 19; fi22++ {
            fmt.Println(fi22, v2, v2)
            v23 := []int{int(v2), fi22, fi22}
            _ = fi22
            _ = v23
        }
        if (false && (42 >= (^(v2)))) {
            var v24 uint = 255
            var v25 string = "a"
            if false {
                v24 = v24
                v26 := []int{166, 156, 4294967295}
                v27 := GenS1{F0: (v2 | v2), F1: 32767, F2: (v24 ^ v24)}
                _ = v26
                _ = v27
            } else {
                if (!true) {
                    if (v24 >= v24) {
                        if false {
                            v28 := GenS3{F0: "foo bar"}
                            var v29 uint16 = 2
                            v1 -= 0.0
                            _ = v28
                            _ = v29
                        }
                    } else {
                        for fi30 := 0; fi30 < 11; fi30++ {
                            v24 = v24
                            _ = fi30
                        }
                        var v31 string = v25
                        v2 -= -32768
                        _ = v31
                    }
                }
                var v32 float64 = (1e10 - 1.5)
                if false {
                    if false {
                        v25 = v25
                        _ = len(v3)
                        for fi33 := 0; fi33 < 2; fi33++ {
                            v24 = v24
                            if (true || true) {
                                if true {
                                    v25 = v25
                                    v34 := GenS1{F0: 1000, F1: 65535, F2: (^(v24))}
                                    _ = v34
                                } else {
                                    var v35 int16 = v2
                                    _ = v35
                                }
                            } else {
                                v36 := []int{(fi33 ^ (fi33 >> (((v24 << ((v24) & 63))) & 63))), fi33}
                                _ = v36
                            }
                            for fi37 := 0; fi37 < 13; fi37++ {
                                v24 += v24
                                v38 := GenS3{F0: (("a" + "a") + v25)}
                                var v39 float64 = v32
                                _ = fi37
                                _ = v38
                                _ = v39
                            }
                            _ = fi33
                        }
                    }
                } else {
                    var v40 uint64 = 739
                    v41 := []int{42}
                    _ = v40
                    _ = v41
                }
                _ = v32
            }
            _ = v24
            _ = v25
        } else {
            v1 += 1e10
            v42 := GenS3{F0: "a"}
            _ = v42
        }
    }
    switch "foo bar" {
    case "hello":
        v43 := GenS4{F0: 2, F1: 2, F2: 18446744073709551615}
        _ = v43
    case "a":
        fmt.Println(v2)
    case "foo bar":
        var v44 uint32 = 0
        _ = v44
    default:
        var v45 int8 = -89
        _ = v45
    }
    v1 += float32(3.14159)
    v1 -= v1
    _ = len(v3)
    var v46 uint32 = genFn1()
    for fi47 := 0; fi47 < 1; fi47++ {
        v46 -= genFn1()
        fmt.Println(v1, v1, fi47)
        if (!false) {
            for fi48 := 0; fi48 < 9; fi48++ {
                var v49 uint8 = 42
                _ = fi48
                _ = v49
            }
            if false {
                var v50 uint64 = 906
                var v51 int16 = v2
                if (false && true) {
                    fmt.Println(v1, v2)
                }
                _ = v50
                _ = v51
            } else {
                var v52 int = (^(fi47))
                switch "foo bar" {
                case "foo bar":
                    for fi53 := 0; fi53 < 15; fi53++ {
                        v54 := []int{(fi47 & fi53), (-(v52)), 42}
                        _ = fi53
                        _ = v54
                    }
                default:
                    if (!(false || (!true))) {
                        v55 := []int{(v52 * -747), (fi47 - (fi47 * (fi47 | 127))), 127}
                        for fi56 := 0; fi56 < 10; fi56++ {
                            v1 += (v1 * 1e10)
                            if (!(float64(v1) == 1e10)) {
                                var v57 bool = false
                                v46 -= 4294967295
                                _ = v57
                            } else {
                                v58 := GenS1{F0: v2, F1: 32767, F2: 4294967295}
                                fmt.Println(v52, v46)
                                if false {
                                    v52 = 230
                                    fmt.Println(v46)
                                    switch fi56 {
                                    case -2147483648:
                                        switch (v52 + fi56) {
                                        case 222:
                                            var v59 uint64 = 2
                                            _ = v59
                                        case -329:
                                            _ = len(v55)
                                        }
                                    case 32768:
                                        var v60 uint32 = v46
                                        _ = v60
                                    case 835:
                                        var v61 float32 = v1
                                        _ = v61
                                    }
                                    var v62 int16 = v2
                                    if (false && true) {
                                        if false {
                                            v2 += (v62 | v62)
                                            v63 := GenS4{F0: 1, F1: 18446744073709551615, F2: 42}
                                            _ = v63
                                        }
                                    }
                                    _ = v62
                                }
                                _ = v58
                            }
                            fmt.Println(v46)
                            _ = fi56
                        }
                        _ = v55
                    }
                }
                var v64 int8 = 127
                _ = v52
                _ = v64
            }
            switch fi47 {
            case 1:
                switch "\n\t" {
                case "\n\t":
                    var v65 uint32 = 2147483648
                    _ = v65
                }
            case 127:
                v46 -= v46
            case -1:
                v2 += (v2 ^ (v2 | -111))
            }
        } else {
            _ = len(v3)
            var v66 float64 = (float64(0.0) + 1e10)
            v67 := GenS5{F0: 1, F1: 1}
            _ = v66
            _ = v67
        }
        _ = fi47
    }
    if true {
        v68 := GenS6{F0: 4294967295, F1: "a", F2: 42}
        _ = v68
    } else {
        v69 := GenS3{F0: "\n\t"}
        v70 := v3["foo bar"]
        for fi71 := 0; fi71 < 8; fi71++ {
            if (2 < 42) {
                var v72 uint64 = 18446744073709551615
                if true {
                    for fi73 := 0; fi73 < 15; fi73++ {
                        v72 = (v72 * 18446744073709551615)
                        delete(v4, -2147483648)
                        _ = fi73
                    }
                } else {
                    v72 -= v72
                    fmt.Println(v70)
                }
                _ = v72
            } else {
                for fi74 := 0; fi74 < 10; fi74++ {
                    var v75 int16 = v2
                    _ = fi74
                    _ = v75
                }
            }
            if true {
                v1 = (float32(0.0) / v1)
                if ((v2 << ((374) & 63)) == v2) {
                    v70 = v70
                    v76 := GenS7{F0: 1}
                    _ = v76
                } else {
                    v2 -= (v2 - (v2 | (v2 ^ v2)))
                    v77 := GenS6{F0: (fi71 % (((^(fi71))) | 1)), F1: "a", F2: 18446744073709551615}
                    if true {
                        v46 -= (v46 ^ genFn1())
                        if true {
                            switch fi71 {
                            case 489:
                                v70 = ((v70 + (v70 + v70)) + v70)
                            case 0:
                                v46 = 2
                            case 128:
                                var v78 float64 = (((0.0 + 1.5) * ((1e10 / (0.1 * 3.14159)) + 1.5)) * 1.5)
                                _ = v78
                            default:
                                if (-1 == -735) {
                                    var v79 int = (fi71 << ((255) & 63))
                                    v79 = (fi71 - 393)
                                    _ = v79
                                } else {
                                    fmt.Println(v2, fi71, v1)
                                    v2 = v2
                                }
                            }
                            switch v70 {
                            case "\n\t":
                                v80 := []int{fi71, 2147483647, fi71}
                                _ = v80
                            case "a":
                                v81 := GenS3{F0: ((v70 + v70) + v70)}
                                _ = v81
                            default:
                                for fi82 := 0; fi82 < 16; fi82++ {
                                    v83 := GenS6{F0: fi71, F1: "a", F2: 600}
                                    v46 = v46
                                    _ = fi82
                                    _ = v83
                                }
                            }
                            switch fi71 {
                            case -853:
                                v84 := GenS6{F0: fi71, F1: v70, F2: 255}
                                _ = v84
                            case 127:
                                switch fi71 {
                                case 0:
                                    fmt.Println(v1)
                                case 127:
                                    var v85 int = (fi71 / ((fi71) | 1))
                                    _ = v85
                                case 1:
                                    _ = len(v4)
                                }
                            case 4294967295:
                                _ = len(v3)
                            default:
                                if true {
                                    if ((v46 - v46) >= v46) {
                                        fmt.Println(v46, v2)
                                        var v86 uint8 = 1
                                        _ = v86
                                    } else {
                                        v46 = 42
                                        if ((!true) && (((float32(0.0) > float32(0.0)) && false) || false)) {
                                            for fi87 := 0; fi87 < 16; fi87++ {
                                                v88 := []int{(fi87 & fi87), (fi87 * (fi87 + (fi71 & -782))), (fi71 & fi71)}
                                                _ = fi87
                                                _ = v88
                                            }
                                            if (!(0 >= 42)) {
                                                v1 = v1
                                                v46 = (v46 - genFn1())
                                            } else {
                                                if false {
                                                    fmt.Println(v70, fi71)
                                                    switch v70 {
                                                    case "":
                                                        var v89 uint16 = 255
                                                        _ = v89
                                                    case "a":
                                                        switch v70 {
                                                        case "\n\t":
                                                            v46 -= v46
                                                        case "foo bar":
                                                            v1 = ((1e10 - float32(1e10)) + float32(((-2.25 - (float64(-2.25) + ((-2.25 / 1e10) / (float64(float32(((1e10 - 1e10) - (0.0 - float64(v1))))) * ((1e10 - 1e10) * (float64(0.1) / 1.5)))))) * 3.14159)))
                                                        }
                                                    default:
                                                        v46 = genFn1()
                                                    }
                                                }
                                            }
                                            var v90 int64 = -1
                                            _ = v90
                                        } else {
                                            switch fi71 {
                                            case 1:
                                                var v91 uint32 = 1
                                                _ = v91
                                            default:
                                                var v92 int32 = -216
                                                _ = v92
                                            }
                                            fmt.Println(fi71, v1)
                                        }
                                        var v93 uint64 = 2
                                        _ = v93
                                    }
                                    var v94 string = v70
                                    v95 := GenS2{F0: -872, F1: 79}
                                    _ = v94
                                    _ = v95
                                } else {
                                    if false {
                                        v46 = v46
                                        for fi96 := 0; fi96 < 1; fi96++ {
                                            v97 := v4[(fi96 / ((867) | 1))]
                                            switch fi96 {
                                            case 1:
                                                v1 += v1
                                            case 2147483647:
                                                var v98 string = ""
                                                _ = v98
                                            }
                                            var v99 uint16 = 255
                                            _ = fi96
                                            _ = v97
                                            _ = v99
                                        }
                                    }
                                }
                            }
                        } else {
                            v100 := GenS7{F0: 270}
                            v101 := GenS5{F0: 2, F1: 118}
                            _ = v100
                            _ = v101
                        }
                    } else {
                        if (456 > 42) {
                            v1 += v1
                            if (-2.25 == float32(3.14159)) {
                                v102 := []int{(fi71 % ((fi71) | 1)), fi71}
                                if false {
                                    var v103 float64 = 1.5
                                    delete(v4, fi71)
                                    for fi104 := 0; fi104 < 7; fi104++ {
                                        var v105 uint = 2
                                        if (((("a" + v70) + "\n\t") == v70) || true) {
                                            v106 := v4[fi104]
                                            switch fi104 {
                                            case -1:
                                                var v107 uint16 = 414
                                                _ = v107
                                            case 42:
                                                var v108 int = fi71
                                                _ = v108
                                            case 2147483647:
                                                v109 := GenS3{F0: v106}
                                                _ = v109
                                            default:
                                                if false {
                                                    v110 := GenS5{F0: 4294967295, F1: v105}
                                                    var v111 int64 = -9223372036854775808
                                                    _ = v110
                                                    _ = v111
                                                }
                                                v70 = v106
                                            }
                                            var v112 float64 = v103
                                            _ = v106
                                            _ = v112
                                        } else {
                                            if false {
                                                switch "hello" {
                                                case "\n\t":
                                                    switch (fi104 ^ int(v2)) {
                                                    case 4294967295:
                                                        var v113 int16 = v2
                                                        _ = v113
                                                    default:
                                                        delete(v4, (fi104 ^ (^(fi104))))
                                                        if false {
                                                            var v114 int16 = v2
                                                            var v115 uint32 = (v46 - (v46 & (^(v46))))
                                                            _ = v114
                                                            _ = v115
                                                        } else {
                                                            if len(v102) > 0 {
                                                                v102[0] = fi71
                                                            }
                                                            var v116 uint32 = 712
                                                            if len(v102) >= 2 {
                                                                v102 = v102[1:len(v102)]
                                                            }
                                                            _ = v116
                                                        }
                                                    }
                                                case "foo bar":
                                                    v105 = (v105 % ((0) | 1))
                                                case "a":
                                                    v70 = v70
                                                default:
                                                    if (!false) {
                                                        var v117 string = "\n\t"
                                                        var v118 uint32 = v46
                                                        var v119 int32 = -1
                                                        v120 := GenS2{F0: (-(v2)), F1: 2}
                                                        if true {
                                                            var v121 int32 = -1
                                                            fmt.Println(v46, v105, fi104)
                                                            _ = v121
                                                        }
                                                        _ = v117
                                                        _ = v118
                                                        _ = v119
                                                        _ = v120
                                                    } else {
                                                        var v122 uint64 = 881
                                                        if true {
                                                            delete(v4, fi104)
                                                        }
                                                        switch -535 {
                                                        case -2147483648:
                                                            if ((v70 + ((v70 + "hello") + "hello")) < "foo bar") {
                                                                v123 := GenS4{F0: (v122 / ((42) | 1)), F1: (v105 % (((^(v105))) | 1)), F2: v105}
                                                                switch v70 {
                                                                case "a":
                                                                    _ = len(v102)
                                                                case "":
                                                                    var v124 int = (fi71 ^ (fi104 & fi104))
                                                                    _ = v124
                                                                case "foo bar":
                                                                    var v125 float32 = (v1 * v1)
                                                                    _ = v125
                                                                }
                                                                v102 = append(v102, fi71)
                                                                var v126 int32 = 42
                                                                _ = v123
                                                                _ = v126
                                                            } else {
                                                                var v127 int8 = 0
                                                                _ = v127
                                                            }
                                                        case 0:
                                                            fmt.Println(v103)
                                                        case 42:
                                                            switch v70 {
                                                            case "\n\t":
                                                                fmt.Println(fi104, fi104, fi71)
                                                            }
                                                        default:
                                                            v128 := GenS5{F0: (v122 | (^(v122))), F1: (v105 + v105)}
                                                            _ = v128
                                                        }
                                                        _ = v122
                                                    }
                                                }
                                                v102 = append(v102, 0)
                                            } else {
                                                var v129 int = fi71
                                                v1 = (((v1 + v1) / float32(v103)) + v1)
                                                if true {
                                                    fmt.Println(v1, v1)
                                                    fmt.Println(v2, v1, v1)
                                                    var v130 uint8 = 128
                                                    _ = v130
                                                }
                                                _ = v129
                                            }
                                            if ("hello" <= "a") {
                                                v131 := GenS2{F0: 1, F1: 1}
                                                var v132 bool = (!(-985 < (v2 >> ((v105) & 63))))
                                                _ = v131
                                                _ = v132
                                            }
                                            v46 -= (v46 / ((v46) | 1))
                                        }
                                        _ = fi104
                                        _ = v105
                                    }
                                    _ = v103
                                } else {
                                    fmt.Println(v2, v70)
                                    var v133 bool = false
                                    _ = v133
                                }
                                _ = v102
                            } else {
                                v134 := GenS1{F0: -952, F1: 481, F2: 18446744073709551615}
                                for fi135 := 0; fi135 < 8; fi135++ {
                                    var v136 uint64 = 0
                                    v137 := GenS1{F0: -438, F1: 695, F2: 0}
                                    var v138 string = v70
                                    _ = fi135
                                    _ = v136
                                    _ = v137
                                    _ = v138
                                }
                                _ = v134
                            }
                        }
                        fmt.Println(fi71, v2)
                        v139 := GenS6{F0: fi71, F1: (((v70 + (("hello" + (v70 + (v70 + v70))) + "hello")) + ((v70 + v70) + v70)) + v70), F2: 0}
                        _ = v139
                    }
                    _ = v77
                }
            }
            _ = fi71
        }
        _ = v69
        _ = v70
    }
    v140 := []int{}
    v141 := v4[127]
    v142 := GenS4{F0: 1, F1: 18446744073709551615, F2: 255}
    if (!true) {
        v143 := GenS2{F0: v2, F1: 42}
        _ = v143
    } else {
        v144 := GenS8{F0: 42}
        var v145 bool = false
        switch "" {
        case "hello":
            for fi146 := 0; fi146 < 16; fi146++ {
                fmt.Println(v46, v145)
                if (v2 < v2) {
                    v147 := GenS9{F0: 9223372036854775807, F1: 42}
                    var v148 int = fi146
                    _ = v147
                    _ = v148
                } else {
                    if (!v145) {
                        v140 = append(v140, fi146)
                    } else {
                        for fi149 := 0; fi149 < 4; fi149++ {
                            _ = len(v140)
                            v145 = v145
                            if v145 {
                                v150 := GenS8{F0: v2}
                                if (uint8(v46) <= 0) {
                                    switch (fi146 & fi146) {
                                    case 9223372036854775807:
                                        if ((v46 < genFn1()) || false) {
                                            var v151 float32 = (v1 * (v1 * (v1 * 1.5)))
                                            if (!(false || v145)) {
                                                var v152 int = 2147483647
                                                _ = v152
                                            }
                                            v140 = append(v140, (fi146 - (fi146 ^ (-(fi149)))))
                                            _ = v151
                                        }
                                    case -2147483648:
                                        var v153 int64 = -1
                                        _ = v153
                                    case 800:
                                        v46 = (v46 | (^(v46)))
                                    }
                                } else {
                                    var v154 uint16 = 1
                                    _ = v154
                                }
                                fmt.Println(v1)
                                _ = v150
                            } else {
                                v2 += int16(v46)
                            }
                            _ = fi149
                        }
                        for fi155 := 0; fi155 < 9; fi155++ {
                            v2 = -1
                            v156 := GenS10{F0: -1, F1: false, F2: 255}
                            _ = fi155
                            _ = v156
                        }
                    }
                }
                _ = fi146
            }
        }
        _ = v144
        _ = v145
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v140:", len(v140), v140)
    fmt.Println("v141:", v141)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    _ = v1
    _ = v140
    _ = v141
    _ = v142
    _ = v2
    _ = v3
    _ = v4
    _ = v46
}
