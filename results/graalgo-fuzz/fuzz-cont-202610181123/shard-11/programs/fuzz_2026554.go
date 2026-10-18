package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 string; F2 int }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int32; F1 bool }
type GenS7 struct { F0 string; F1 uint64 }
type GenS8 struct { F0 uint8; F1 bool }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int; F1 int32; F2 bool }
type GenS11 struct { F0 int; F1 uint64 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 int }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 uint64; F1 string; F2 int8 }
type GenS18 struct { F0 int; F1 int8; F2 uint }
type GenS19 struct { F0 uint64; F1 string; F2 string }
func genFn0(p0 uint32, p1 int8) int64 {
    for fi5 := 0; fi5 < 20; fi5++ {
        if false {
            v6 := GenS1{F0: 0}
            switch (fi5 + 545) {
            case 128:
                var v7 string = "a"
                _ = v7
            case -1:
                var v8 string = "a"
                _ = v8
            }
            _ = v6
        }
        _ = fi5
    }
    v9 := GenS1{F0: -1}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v9
    return 383
}
func genFn1() int64 {
    var v10 uint32 = 447
    v10 = v10
    fmt.Println(v10)
    switch 4294967295 {
    case 4294967295:
        var v11 uint32 = 903
        _ = v11
    case 42:
        fmt.Println(v10)
    case -2147483648:
        v12 := []int{127, 581}
        _ = v12
    }
    var v13 float64 = float64(1.5)
    _ = v10
    _ = v13
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 42
    var v3 uint32 = (^(v2))
    var v4 string = "\n\t"
    for fi14 := 0; fi14 < 18; fi14++ {
        if (1 == -128) {
            v2 += 3
            for fi15 := 0; fi15 < 19; fi15++ {
                if false {
                    v16 := v1[fi14]
                    switch v4 {
                    case "hello":
                        var v17 string = (v4 + (((v4 + ((v4 + v4) + "a")) + v4) + ((v4 + "foo bar") + "foo bar")))
                        _ = v17
                    default:
                        v18 := GenS2{F0: v16, F1: 2}
                        v19 := GenS3{F0: true}
                        v16 = fi15
                        v20 := GenS2{F0: fi14, F1: 4294967295}
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    _ = v16
                } else {
                    if true {
                        var v21 int8 = -96
                        v3 = (v3 >> ((706) & 63))
                        _ = v21
                    } else {
                        var v22 bool = (false || ((155 < 1) && (false && false)))
                        if (2 == 140) {
                            var v23 int8 = -74
                            v24 := []int{(fi14 >> ((673) & 63)), (fi14 * 1), (fi14 - (fi15 >> ((65535) & 63)))}
                            v25 := GenS4{F0: v22, F1: v4, F2: fi15}
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                        if (-2147483648 > int32(v2)) {
                            var v26 float64 = 1e10
                            _ = v26
                        } else {
                            v27 := GenS2{F0: (fi14 + fi14), F1: 184}
                            if ((fi14 << ((1) & 63)) < -2147483648) {
                                var v28 float64 = 1e10
                                var v29 string = "hello"
                                v29 = ("hello" + v4)
                                _ = v28
                                _ = v29
                            }
                            v30 := GenS2{F0: fi15, F1: 2}
                            _ = v27
                            _ = v30
                        }
                        if v22 {
                            v3 += v2
                            v31 := GenS4{F0: (!v22), F1: (v4 + v4), F2: -131}
                            _ = v31
                        } else {
                            var v32 uint8 = 152
                            if false {
                                var v33 bool = (!v22)
                                _ = v33
                            }
                            _ = v32
                        }
                        fmt.Println(v4, fi15, v3)
                        _ = v22
                    }
                    v4 = v4
                }
                fmt.Println(v4, v4, fi15)
                _ = fi15
            }
            fmt.Println(v3, v4)
        } else {
            v4 = v4
        }
        var v34 uint32 = v3
        _ = fi14
        _ = v34
    }
    v2 = (^(v3))
    v3 -= 0
    var v35 float64 = 0.0
    v36 := GenS5{F0: v3}
    fmt.Println(v35)
    switch 9223372036854775807 {
    case 127:
        fmt.Println(v35, v35)
    default:
        var v37 int16 = -921
        _ = v37
    }
    if (!true) {
        if true {
            var v38 uint64 = 42
            _ = v38
        } else {
            if false {
                var v39 uint = 0
                for fi40 := 0; fi40 < 17; fi40++ {
                    for fi41 := 0; fi41 < 6; fi41++ {
                        v42 := GenS6{F0: 42, F1: (-298 >= genFn0(42, 34))}
                        v35 = v35
                        switch fi41 {
                        case 127:
                            v3 = v2
                        case 285:
                            fmt.Println(fi41, fi41)
                        default:
                            fmt.Println(v35, v39)
                        }
                        _ = fi41
                        _ = v42
                    }
                    v43 := GenS7{F0: ("" + v4), F1: 1}
                    v3 += v3
                    _ = fi40
                    _ = v43
                }
                _ = v39
            }
            if true {
                if false {
                    if true {
                        switch 128 {
                        case 368:
                            v44 := v1[2147483647]
                            _ = v44
                        }
                        v35 = 1.5
                        _ = len(v1)
                    } else {
                        switch v4 {
                        case "":
                            v45 := []int{32767}
                            _ = v45
                        }
                        v4 = (((v4 + v4) + ((("\n\t" + "hello") + ("\n\t" + "foo bar")) + v4)) + "a")
                        switch "hello" {
                        case "hello":
                            if false {
                                switch 42 {
                                case 500:
                                    var v46 int32 = 2147483647
                                    _ = v46
                                }
                                var v47 int32 = 42
                                v3 = 120
                                _ = v47
                            } else {
                                for fi48 := 0; fi48 < 1; fi48++ {
                                    _ = len(v1)
                                    for fi49 := 0; fi49 < 5; fi49++ {
                                        if true {
                                            var v50 uint16 = 65535
                                            _ = v50
                                        } else {
                                            switch (fi49 & fi48) {
                                            case 127:
                                                delete(v1, (fi49 / ((fi49) | 1)))
                                            case 1:
                                                var v51 int64 = -9223372036854775808
                                                _ = v51
                                            case 2147483647:
                                                var v52 int64 = 9223372036854775807
                                                var v53 int8 = -128
                                                _ = v52
                                                _ = v53
                                            default:
                                                if (4294967295 >= (-(fi48))) {
                                                    v2 = v2
                                                    var v54 uint = 65535
                                                    v55 := []int{-322}
                                                    _ = v54
                                                    _ = v55
                                                } else {
                                                    if true {
                                                        var v56 int8 = 100
                                                        v35 = v35
                                                        fmt.Println(v35, fi49)
                                                        _ = v56
                                                    }
                                                }
                                            }
                                            v4 = v4
                                        }
                                        if false {
                                            var v57 string = v4
                                            var v58 uint16 = 255
                                            _ = v57
                                            _ = v58
                                        }
                                        var v59 int8 = -53
                                        _ = fi49
                                        _ = v59
                                    }
                                    v35 -= (float64(0.1) - v35)
                                    _ = fi48
                                }
                                var v60 int32 = 42
                                _ = v60
                            }
                        case "a":
                            var v61 uint64 = 1
                            _ = v61
                        }
                    }
                    v62 := GenS8{F0: 2, F1: false}
                    var v63 bool = (65535 < 1)
                    _ = v62
                    _ = v63
                }
                v1[32768] = 32768
            } else {
                v64 := []int{42}
                v3 -= (^(v3))
                v2 = (^(v3))
                _ = v64
            }
        }
        for fi65 := 0; fi65 < 19; fi65++ {
            var v66 uint64 = 1
            if (0 == 253) {
                var v67 uint32 = v2
                if true {
                    v68 := GenS2{F0: fi65, F1: 7}
                    _ = v68
                }
                _ = v67
            } else {
                if (404 < 1) {
                    if ((3.14159 - (3.14159 * 1.5)) == 1.5) {
                        if false {
                            var v69 int64 = 1
                            _ = v69
                        } else {
                            v35 = 0.0
                            v70 := []int{4294967295}
                            switch v4 {
                            case "\n\t":
                                if (false && false) {
                                    if (-9223372036854775808 != 1) {
                                        for fi71 := 0; fi71 < 15; fi71++ {
                                            v66 += v66
                                            if false {
                                                v72 := GenS9{F0: 42}
                                                _ = v72
                                            } else {
                                                switch "" {
                                                case "":
                                                    switch ((v4 + v4) + ((v4 + v4) + "foo bar")) {
                                                    case "\n\t":
                                                        if true {
                                                            var v73 int32 = -2147483648
                                                            v73 = (^(v73))
                                                            v66 = v66
                                                            _ = v73
                                                        } else {
                                                            fmt.Println(v3, fi71)
                                                        }
                                                    case "hello":
                                                        if false {
                                                            switch v4 {
                                                            case "hello":
                                                                fmt.Println(fi71)
                                                            case "":
                                                                var v74 int16 = 991
                                                                _ = v74
                                                            default:
                                                                var v75 int8 = 24
                                                                _ = v75
                                                            }
                                                            fmt.Println(v2, v4, v2)
                                                        }
                                                    case "a":
                                                        var v76 uint = 2
                                                        v1[fi71] = (-(fi65))
                                                        v4 = "\n\t"
                                                        _ = v76
                                                    default:
                                                        v2 -= 65535
                                                    }
                                                case "a":
                                                    if true {
                                                        var v77 int64 = 1
                                                        _ = v77
                                                    }
                                                default:
                                                    v78 := GenS10{F0: fi71, F1: 1, F2: false}
                                                    _ = v78
                                                }
                                                var v79 uint = 2
                                                _ = v79
                                            }
                                            if true {
                                                v66 = (^(v66))
                                                v4 = v4
                                                var v80 float32 = 0.1
                                                var v81 int16 = 112
                                                var v82 float32 = 0.0
                                                _ = v80
                                                _ = v81
                                                _ = v82
                                            } else {
                                                v3 += (v2 % ((v3) | 1))
                                                if (v66 > 1) {
                                                    var v83 uint64 = (^(v66))
                                                    v84 := GenS11{F0: (fi65 >> ((18446744073709551615) & 63)), F1: (v83 % (((^(v66))) | 1))}
                                                    v1[-580] = 914
                                                    if true {
                                                        v35 = (-2.25 / v35)
                                                        if len(v70) > 0 {
                                                            v70[0] = (fi71 | (fi65 >> ((uint(v66)) & 63)))
                                                        }
                                                        fmt.Println(fi71)
                                                    }
                                                    fmt.Println(v4, v66)
                                                    _ = v83
                                                    _ = v84
                                                }
                                                v85 := GenS12{F0: (v66 / ((v66) | 1))}
                                                _ = v85
                                            }
                                            _ = fi71
                                        }
                                        for fi86 := 0; fi86 < 14; fi86++ {
                                            v87 := GenS10{F0: (fi65 | (fi65 & fi86)), F1: 1, F2: false}
                                            if false {
                                                var v88 int8 = -23
                                                var v89 uint8 = 1
                                                _ = v88
                                                _ = v89
                                            } else {
                                                if (false && true) {
                                                    if false {
                                                        var v90 string = ("\n\t" + ("\n\t" + (v4 + "")))
                                                        var v91 int8 = int8(v2)
                                                        if true {
                                                            v92 := GenS2{F0: (fi86 ^ fi86), F1: 399}
                                                            fmt.Println(v4, v2)
                                                            v93 := GenS13{F0: v91}
                                                            _ = v92
                                                            _ = v93
                                                        } else {
                                                            v90 = v4
                                                        }
                                                        _ = v90
                                                        _ = v91
                                                    } else {
                                                        switch ((v4 + "\n\t") + (v4 + "a")) {
                                                        case "\n\t":
                                                            var v94 uint8 = 127
                                                            _ = v94
                                                        case "foo bar":
                                                            if (!(true || true)) {
                                                                var v95 int16 = 0
                                                                if false {
                                                                    if ((v3 + v3) >= v3) {
                                                                        v96 := GenS14{F0: (!true)}
                                                                        v95 = (v95 % ((32767) | 1))
                                                                        delete(v1, fi86)
                                                                        _ = v96
                                                                    } else {
                                                                        v3 -= 42
                                                                    }
                                                                    if len(v70) >= 2 {
                                                                        v70 = v70[1:len(v70)]
                                                                    }
                                                                } else {
                                                                    if false {
                                                                        v4 = v4
                                                                        switch fi86 {
                                                                        case 1:
                                                                            var v97 int32 = -2147483648
                                                                            _ = v97
                                                                        case 42:
                                                                            v98 := v1[fi86]
                                                                            _ = v98
                                                                        }
                                                                    }
                                                                }
                                                                v99 := GenS3{F0: true}
                                                                var v100 int64 = genFn1()
                                                                _ = v100
                                                                _ = v95
                                                                _ = v99
                                                            } else {
                                                                switch -463 {
                                                                case 42:
                                                                    if len(v70) >= 2 {
                                                                        v70 = v70[1:len(v70)]
                                                                    }
                                                                case 4294967295:
                                                                    var v101 uint16 = 909
                                                                    _ = v101
                                                                case 32767:
                                                                    if (!false) {
                                                                        var v102 string = (v4 + "a")
                                                                        _ = v102
                                                                    }
                                                                default:
                                                                    v3 = v2
                                                                }
                                                                v70 = append(v70, (fi65 >> ((0) & 63)))
                                                            }
                                                        case "":
                                                            v66 -= v66
                                                            fmt.Println(v2)
                                                            switch -115 {
                                                            case 42:
                                                                delete(v1, fi86)
                                                            case 128:
                                                                var v103 uint16 = 0
                                                                _ = v103
                                                            case 32768:
                                                                v104 := v1[fi86]
                                                                _ = v104
                                                            }
                                                            if false {
                                                                v70 = append(v70, -620)
                                                                var v105 int = (fi65 * 984)
                                                                fmt.Println(v66, v3)
                                                                _ = v105
                                                            } else {
                                                                switch -2147483648 {
                                                                case -310:
                                                                    var v106 int64 = 872
                                                                    _ = v106
                                                                }
                                                            }
                                                            v2 = (v3 + 120)
                                                        }
                                                        v4 = v4
                                                        if ((!false) && false) {
                                                            if (true && false) {
                                                                v2 = 0
                                                                v2 += v3
                                                                var v107 uint8 = 255
                                                                _ = v107
                                                            }
                                                            if true {
                                                                v108 := GenS9{F0: (v3 - (^(v3)))}
                                                                _ = v108
                                                            } else {
                                                                v109 := GenS2{F0: (fi86 + 2147483647), F1: 42}
                                                                v110 := GenS2{F0: (fi65 % ((0) | 1)), F1: 18446744073709551615}
                                                                _ = v109
                                                                _ = v110
                                                            }
                                                            if (100 <= 1) {
                                                                if true {
                                                                    if (-738 < 0) {
                                                                        var v111 int64 = -866
                                                                        if false {
                                                                            if ((v66 <= 42) || false) {
                                                                                var v112 uint8 = 79
                                                                                v113 := GenS9{F0: (v2 / ((597) | 1))}
                                                                                _ = v112
                                                                                _ = v113
                                                                            } else {
                                                                                switch v4 {
                                                                                case "hello":
                                                                                    var v114 uint16 = 255
                                                                                    _ = v114
                                                                                }
                                                                                v66 -= v66
                                                                                var v115 float32 = -2.25
                                                                                _ = v115
                                                                            }
                                                                        }
                                                                        _ = v111
                                                                    } else {
                                                                        var v116 float64 = float64(-2.25)
                                                                        var v117 float32 = 1e10
                                                                        var v118 string = ("a" + "foo bar")
                                                                        _ = v116
                                                                        _ = v117
                                                                        _ = v118
                                                                    }
                                                                } else {
                                                                    switch v4 {
                                                                    case "hello":
                                                                        if true {
                                                                            if true {
                                                                                fmt.Println(v4, v2)
                                                                                if true {
                                                                                    v119 := v1[(fi65 % ((fi65) | 1))]
                                                                                    if len(v70) >= 2 {
                                                                                        v70 = v70[1:len(v70)]
                                                                                    }
                                                                                    v120 := GenS15{F0: (fi86 ^ v119)}
                                                                                    _ = v119
                                                                                    _ = v120
                                                                                }
                                                                                fmt.Println(v66, v35)
                                                                            }
                                                                            var v121 float64 = v35
                                                                            _ = v121
                                                                        }
                                                                        v70 = append(v70, (fi65 + fi65))
                                                                        fmt.Println(v2)
                                                                        var v122 uint32 = v3
                                                                        _ = v122
                                                                    default:
                                                                        var v123 string = "hello"
                                                                        _ = v123
                                                                    }
                                                                    switch "a" {
                                                                    case "a":
                                                                        v2 = v3
                                                                    default:
                                                                        switch ("" + "hello") {
                                                                        case "hello":
                                                                            var v124 float32 = (((0.1 + 3.14159) / 1e10) + 3.14159)
                                                                            _ = v124
                                                                        case "":
                                                                            switch v4 {
                                                                            case "":
                                                                                var v125 uint = 842
                                                                                _ = v125
                                                                            default:
                                                                                v4 = ""
                                                                            }
                                                                        }
                                                                    }
                                                                    v4 = ("foo bar" + v4)
                                                                }
                                                            }
                                                        } else {
                                                            v126 := v1[127]
                                                            _ = v126
                                                        }
                                                    }
                                                    if len(v70) > 0 {
                                                        v70[0] = (fi65 * fi86)
                                                    }
                                                    v127 := GenS16{F0: 65535}
                                                    _ = v127
                                                } else {
                                                    v128 := GenS17{F0: 2, F1: v4, F2: -97}
                                                    fmt.Println(v2)
                                                    var v129 uint64 = 2
                                                    _ = v128
                                                    _ = v129
                                                }
                                                if len(v70) > 0 {
                                                    v70[0] = -9223372036854775808
                                                }
                                                v4 = (v4 + (((("a" + v4) + ("hello" + v4)) + v4) + v4))
                                            }
                                            switch fi86 {
                                            case -250:
                                                v4 = v4
                                            case -9223372036854775808:
                                                if len(v70) > 0 {
                                                    v70[0] = (fi86 ^ (fi65 & int(v2)))
                                                }
                                            }
                                            _ = fi86
                                            _ = v87
                                        }
                                    }
                                    var v130 bool = false
                                    _ = v130
                                } else {
                                    v4 = (v4 + (("\n\t" + (v4 + v4)) + v4))
                                    v3 = (v3 ^ 2147483647)
                                }
                            case "hello":
                                switch (v4 + v4) {
                                case "foo bar":
                                    var v131 bool = true
                                    _ = v131
                                case "\n\t":
                                    for fi132 := 0; fi132 < 8; fi132++ {
                                        fmt.Println(fi65, v2)
                                        v2 = v2
                                        var v133 uint8 = 1
                                        _ = fi132
                                        _ = v133
                                    }
                                case "hello":
                                    v35 += 3.14159
                                default:
                                    switch (fi65 - fi65) {
                                    case 1:
                                        v134 := GenS2{F0: 127, F1: v66}
                                        _ = v134
                                    default:
                                        fmt.Println(v35, v66)
                                    }
                                }
                            default:
                                if false {
                                    for fi135 := 0; fi135 < 4; fi135++ {
                                        v35 += v35
                                        v4 = ("" + "")
                                        _ = fi135
                                    }
                                    fmt.Println(v35)
                                } else {
                                    for fi136 := 0; fi136 < 2; fi136++ {
                                        switch fi136 {
                                        case 0:
                                            v2 += (v2 + v3)
                                        default:
                                            var v137 int64 = -673
                                            _ = v137
                                        }
                                        _ = fi136
                                    }
                                }
                            }
                            _ = v70
                        }
                        var v138 uint64 = v66
                        v139 := GenS18{F0: (fi65 | fi65), F1: 0, F2: 0}
                        _ = v138
                        _ = v139
                    } else {
                        var v140 int16 = 1
                        _ = v140
                    }
                    for fi141 := 0; fi141 < 11; fi141++ {
                        var v142 int64 = 9223372036854775807
                        _ = fi141
                        _ = v142
                    }
                }
            }
            if false {
                for fi143 := 0; fi143 < 20; fi143++ {
                    v3 = 2147483648
                    _ = fi143
                }
                v3 += 2147483647
                v144 := GenS19{F0: 4294967295, F1: v4, F2: v4}
                _ = v144
            }
            _ = fi65
            _ = v66
        }
    } else {
        if true {
            v145 := []int{-2147483648, 802, 477}
            v4 = v4
            _ = v145
        } else {
            var v146 uint8 = 180
            v147 := []int{42}
            fmt.Println(v146, v3)
            _ = v146
            _ = v147
        }
    }
    var v148 int32 = -2147483648
    v149 := v1[42]
    var v150 int8 = -17
    v151 := []int{v149, v149, 32768}
    switch 9223372036854775807 {
    case -332:
        v2 += 2147483648
    case -1:
        switch "foo bar" {
        case "\n\t":
            for fi152 := 0; fi152 < 6; fi152++ {
                v149 += (v149 / ((-881) | 1))
                var v153 bool = true
                _ = fi152
                _ = v153
            }
        case "a":
            var v154 int64 = 1
            _ = v154
        case "hello":
            var v155 int16 = 0
            _ = v155
        }
    default:
        delete(v1, (-(v149)))
    }
    switch v149 {
    case 9223372036854775807:
        var v156 float32 = 3.14159
        _ = v156
    }
    fmt.Println(v3)
    for fi157 := 0; fi157 < 13; fi157++ {
        var v158 int = v149
        _ = fi157
        _ = v158
    }
    v149 += v149
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v148:", v148)
    fmt.Println("v149:", v149)
    fmt.Println("v150:", v150)
    fmt.Println("slicelen v151:", len(v151), v151)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v148
    _ = v149
    _ = v150
    _ = v151
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
}
