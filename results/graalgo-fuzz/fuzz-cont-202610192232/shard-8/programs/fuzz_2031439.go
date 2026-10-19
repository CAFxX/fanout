package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 bool; F1 bool; F2 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int32; F1 int16 }
type GenS7 struct { F0 int16; F1 bool; F2 uint64 }
type GenS8 struct { F0 int8; F1 bool }
type GenS9 struct { F0 uint32; F1 bool }
type GenS10 struct { F0 int8; F1 string }
type GenS11 struct { F0 int8; F1 string }
type GenS12 struct { F0 int64; F1 int64; F2 int64 }
type GenS13 struct { F0 int; F1 uint64; F2 int32 }
func genFn0(p0 uint64, p1 string) uint32 {
    v5 := []int{-880, -883, 359, 42}
    for fi6 := 0; fi6 < 19; fi6++ {
        for fi7 := 0; fi7 < 12; fi7++ {
            v8 := GenS1{F0: -9223372036854775808, F1: 2}
            _ = fi7
            _ = v8
        }
        p1 = "hello"
        if true {
            v9 := GenS2{F0: (!(fi6 > (fi6 + (^(fi6)))))}
            if false {
                p1 = "hello"
                if len(v5) > 0 {
                    v5[0] = fi6
                }
            } else {
                p0 = 1
                for fi10 := 0; fi10 < 4; fi10++ {
                    switch p1 {
                    case "a":
                        var v11 string = "a"
                        _ = v11
                    default:
                        var v12 string = p1
                        _ = v12
                    }
                    if len(v5) > 0 {
                        v5[0] = 1
                    }
                    _ = fi10
                }
                p1 = ("hello" + p1)
            }
            _ = v9
        }
        _ = fi6
    }
    _ = len(v5)
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v5
    return 2147483648
}

func main() {
    var v1 uint64 = 196
    var v2 uint32 = 369
    var v3 bool = (42 == -1)
    v4 := []int{127, 42}
    v13 := GenS1{F0: 22, F1: uint(v2)}
    var v14 float32 = (-2.25 / 3.14159)
    var v15 int32 = -98
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 uint32 = genFn0(v1, "a")
        switch (fi16 >> ((2) & 63)) {
        case -74:
            for fi18 := 0; fi18 < 3; fi18++ {
                var v19 int32 = v15
                if ((v1 / (((v1 + (v1 & 4294967295))) | 1)) > 4294967295) {
                    if (!true) {
                        v20 := GenS1{F0: 9223372036854775807, F1: 42}
                        _ = v20
                    } else {
                        var v21 uint = 116
                        v3 = v3
                        _ = v21
                    }
                    var v22 bool = v3
                    _ = v22
                } else {
                    switch "hello" {
                    case "foo bar":
                        var v23 string = ((("a" + "\n\t") + "a") + "")
                        _ = v23
                    case "hello":
                        var v24 int = fi16
                        _ = v24
                    }
                    v17 = (^(v17))
                    var v25 uint64 = (v1 & v1)
                    _ = v25
                }
                _ = fi18
                _ = v19
            }
        case 9223372036854775807:
            v26 := GenS3{F0: 2, F1: 42}
            _ = v26
        case 2147483647:
            var v27 uint8 = 217
            _ = v27
        default:
            if (v3 && (((^(v17)) >= v17) && (int8(v1) <= 42))) {
                v4 = append(v4, fi16)
                v28 := GenS4{F0: (!false), F1: false, F2: (-1 >= 117)}
                var v29 int16 = 1
                _ = v28
                _ = v29
            }
        }
        _ = fi16
        _ = v17
    }
    v14 = 1e10
    if true {
        switch "\n\t" {
        case "foo bar":
            var v30 int16 = 32767
            _ = v30
        case "a":
            v31 := GenS5{F0: 128}
            _ = v31
        case "":
            var v32 int16 = 1000
            _ = v32
        }
        var v33 string = "\n\t"
        _ = v33
    } else {
        var v34 bool = (42 <= 1000)
        for fi35 := 0; fi35 < 17; fi35++ {
            var v36 uint32 = genFn0(v1, "\n\t")
            v2 -= v36
            for fi37 := 0; fi37 < 7; fi37++ {
                if len(v4) > 0 {
                    v4[0] = fi37
                }
                var v38 int16 = 600
                _ = fi37
                _ = v38
            }
            _ = fi35
            _ = v36
        }
        _ = v34
    }
    switch -253 {
    case 2147483647:
        for fi39 := 0; fi39 < 15; fi39++ {
            switch "foo bar" {
            case "":
                v40 := GenS1{F0: -824, F1: 566}
                _ = v40
            default:
                fmt.Println(fi39)
            }
            var v41 int64 = 982
            _ = fi39
            _ = v41
        }
    case -796:
        for fi42 := 0; fi42 < 17; fi42++ {
            fmt.Println(v1)
            var v43 uint = 4294967295
            fmt.Println(v43)
            _ = fi42
            _ = v43
        }
    case -387:
        if (v3 || v3) {
            for fi44 := 0; fi44 < 6; fi44++ {
                _ = len(v4)
                _ = fi44
            }
            fmt.Println(v3, v3)
            if v3 {
                for fi45 := 0; fi45 < 5; fi45++ {
                    if ((!true) || v3) {
                        var v46 int16 = 0
                        v4 = append(v4, 1)
                        if ((v46 ^ 221) != 1000) {
                            for fi47 := 0; fi47 < 2; fi47++ {
                                v3 = (fi45 != fi45)
                                if (-65 != 1) {
                                    if v3 {
                                        switch fi47 {
                                        case -992:
                                            v48 := GenS6{F0: -1, F1: -990}
                                            _ = v48
                                        case -92:
                                            var v49 uint = 0
                                            var v50 uint8 = 0
                                            v51 := GenS5{F0: v50}
                                            v52 := GenS4{F0: (-1 >= 1), F1: (true || ((v1 % ((42) | 1)) < (v1 + uint64(v50)))), F2: true}
                                            _ = v49
                                            _ = v50
                                            _ = v51
                                            _ = v52
                                        case 32768:
                                            v53 := GenS7{F0: v46, F1: (false && false), F2: 819}
                                            switch (fi45 - -9223372036854775808) {
                                            case 32768:
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v54 uint = 1
                                                    _ = v54
                                                default:
                                                    fmt.Println(v3)
                                                }
                                            case 9223372036854775807:
                                                v14 = v14
                                            default:
                                                v55 := GenS1{F0: 925, F1: 4294967295}
                                                _ = v55
                                            }
                                            var v56 uint32 = v2
                                            _ = v53
                                            _ = v56
                                        }
                                    } else {
                                        switch ("foo bar" + "") {
                                        case "\n\t":
                                            v2 += (v2 - 965)
                                        }
                                    }
                                }
                                if true {
                                    var v57 uint32 = 65535
                                    _ = v57
                                } else {
                                    if v3 {
                                        v3 = v3
                                        var v58 int32 = 2147483647
                                        v59 := GenS8{F0: -128, F1: v3}
                                        if v3 {
                                            if (0.0 == (((1e10 / (float64(v14) + 3.14159)) - 0.0) - -2.25)) {
                                                switch ("foo bar" + "") {
                                                case "hello":
                                                    if ((("a" + "foo bar") + ("hello" + "\n\t")) <= ("a" + ("a" + "hello"))) {
                                                        if len(v4) > 0 {
                                                            v4[0] = 42
                                                        }
                                                        var v60 float32 = v14
                                                        _ = v60
                                                    } else {
                                                        v15 -= v15
                                                        var v61 bool = false
                                                        v62 := GenS9{F0: 2147483648, F1: v61}
                                                        _ = v61
                                                        _ = v62
                                                    }
                                                default:
                                                    v15 += v15
                                                }
                                                var v63 int = (fi45 >> ((48) & 63))
                                                v3 = (!(952 != 32767))
                                                _ = v63
                                            }
                                            v64 := GenS9{F0: genFn0((v1 * v1), "foo bar"), F1: ((1.5 / -2.25) > 3.14159)}
                                            _ = v64
                                        } else {
                                            var v65 uint32 = v2
                                            if false {
                                                switch 4294967295 {
                                                case -76:
                                                    v66 := GenS1{F0: -850, F1: 4294967295}
                                                    _ = v66
                                                case -9223372036854775808:
                                                    var v67 float32 = ((v14 / v14) / 3.14159)
                                                    _ = v67
                                                case -2147483648:
                                                    var v68 uint8 = 1
                                                    _ = v68
                                                default:
                                                    if false {
                                                        var v69 uint16 = 35
                                                        if v3 {
                                                            v65 -= v65
                                                        } else {
                                                            switch fi47 {
                                                            case 42:
                                                                var v70 float32 = (v14 * float32(float64(1e10)))
                                                                _ = v70
                                                            default:
                                                                if (!((^(fi47)) != (fi45 & 0))) {
                                                                    if len(v4) > 0 {
                                                                        v4[0] = -101
                                                                    }
                                                                } else {
                                                                    v71 := GenS5{F0: 127}
                                                                    _ = v71
                                                                }
                                                            }
                                                        }
                                                        v72 := GenS7{F0: (v46 * 0), F1: ((v14 < v14) || true), F2: (v1 & 0)}
                                                        _ = v69
                                                        _ = v72
                                                    } else {
                                                        var v73 int32 = 1
                                                        var v74 uint16 = 1
                                                        v75 := GenS9{F0: (v2 + (v65 * v2)), F1: false}
                                                        var v76 int64 = -1
                                                        fmt.Println(v65)
                                                        v73 = v73
                                                        _ = v73
                                                        _ = v74
                                                        _ = v75
                                                        _ = v76
                                                    }
                                                }
                                                var v77 bool = true
                                                if true {
                                                    v58 += 0
                                                    if len(v4) > 0 {
                                                        v4[0] = (-(fi45))
                                                    }
                                                    var v78 int16 = (v46 | -138)
                                                    var v79 uint16 = 255
                                                    fmt.Println(v58, v58, v15)
                                                    _ = v78
                                                    _ = v79
                                                }
                                                _ = v77
                                            } else {
                                                fmt.Println(v2, fi45, fi45)
                                            }
                                            var v80 float32 = (3.14159 + (v14 * 0.0))
                                            _ = v65
                                            _ = v80
                                        }
                                        if true {
                                            v58 = (v15 | v58)
                                            _ = len(v4)
                                        }
                                        _ = v58
                                        _ = v59
                                    } else {
                                        v81 := GenS4{F0: v3, F1: v3, F2: (!false)}
                                        var v82 int16 = 42
                                        switch (("hello" + (("foo bar" + ("foo bar" + ("\n\t" + "hello"))) + "hello")) + "\n\t") {
                                        case "hello":
                                            v14 -= v14
                                        default:
                                            if (-32768 < v82) {
                                                var v83 int8 = 0
                                                if true {
                                                    fmt.Println(fi45, v83)
                                                    var v84 int64 = 0
                                                    v85 := GenS5{F0: 74}
                                                    _ = v84
                                                    _ = v85
                                                } else {
                                                    var v86 uint32 = v2
                                                    var v87 string = ""
                                                    v14 += v14
                                                    if v3 {
                                                        v4 = append(v4, 1)
                                                    }
                                                    _ = len(v4)
                                                    _ = v86
                                                    _ = v87
                                                }
                                                v88 := GenS4{F0: v3, F1: true, F2: true}
                                                _ = v83
                                                _ = v88
                                            }
                                        }
                                        var v89 uint32 = (v2 >> ((18446744073709551615) & 63))
                                        v90 := GenS2{F0: (v3 && (1 > 9223372036854775807))}
                                        fmt.Println(v82, v46)
                                        _ = v81
                                        _ = v82
                                        _ = v89
                                        _ = v90
                                    }
                                }
                                _ = fi47
                            }
                        }
                        _ = v46
                    }
                    _ = fi45
                }
            } else {
                for fi91 := 0; fi91 < 20; fi91++ {
                    for fi92 := 0; fi92 < 5; fi92++ {
                        switch ("foo bar" + (("hello" + ("foo bar" + ("a" + "\n\t"))) + "\n\t")) {
                        case "":
                            switch (((("foo bar" + "foo bar") + ("\n\t" + "")) + ("\n\t" + "\n\t")) + "a") {
                            case "hello":
                                var v93 uint8 = 127
                                _ = v93
                            case "":
                                switch "foo bar" {
                                case "hello":
                                    v2 += 2
                                default:
                                    fmt.Println(v14, fi91, fi91)
                                }
                                _ = len(v4)
                                var v94 int32 = -2147483648
                                _ = v94
                            case "a":
                                var v95 float32 = (v14 * (((v14 + 1.5) + ((((v14 - -2.25) - 1.5) * 1e10) + float32((1e10 + 3.14159)))) + v14))
                                _ = v95
                            default:
                                switch ((("\n\t" + "foo bar") + ("\n\t" + "hello")) + "hello") {
                                case "foo bar":
                                    v14 -= float32(0.0)
                                case "hello":
                                    v96 := GenS1{F0: -1, F1: 255}
                                    fmt.Println(fi92)
                                    _ = v96
                                case "\n\t":
                                    var v97 uint64 = v1
                                    _ = v97
                                }
                            }
                        case "foo bar":
                            v14 += 3.14159
                        default:
                            switch fi92 {
                            case 572:
                                var v98 string = "a"
                                _ = v98
                            case 127:
                                switch (fi92 + fi91) {
                                case 128:
                                    v14 += 3.14159
                                }
                            case 9223372036854775807:
                                v99 := GenS4{F0: true, F1: (false || (v3 || v3)), F2: false}
                                _ = v99
                            default:
                                v14 = (1.5 - v14)
                            }
                        }
                        v100 := GenS10{F0: 127, F1: "foo bar"}
                        _ = fi92
                        _ = v100
                    }
                    switch (fi91 ^ 2147483647) {
                    case -334:
                        v101 := GenS11{F0: 100, F1: "a"}
                        _ = v101
                    case -366:
                        v2 = v2
                    }
                    var v102 uint8 = 0
                    _ = fi91
                    _ = v102
                }
                v3 = v3
                v4 = append(v4, 32768)
            }
        } else {
            for fi103 := 0; fi103 < 16; fi103++ {
                v104 := GenS11{F0: 1, F1: "hello"}
                if (255 < 72) {
                    if true {
                        v15 = (v15 & v15)
                        var v105 uint64 = v1
                        _ = v105
                    }
                    if v3 {
                        var v106 int16 = -770
                        _ = v106
                    }
                    v15 = v15
                }
                var v107 uint64 = 790
                _ = fi103
                _ = v104
                _ = v107
            }
        }
    }
    switch 264 {
    case 32767:
        var v108 uint16 = 31
        _ = v108
    case 619:
        if v3 {
            if (65535 >= uint16(v2)) {
                v15 = v15
                var v109 float32 = 1.5
                v14 = float32(1e10)
                _ = v109
            } else {
                var v110 float32 = 1e10
                if (4294967295 < v2) {
                    _ = len(v4)
                    for fi111 := 0; fi111 < 4; fi111++ {
                        var v112 bool = (54 <= 100)
                        _ = fi111
                        _ = v112
                    }
                    v4 = append(v4, 250)
                } else {
                    v113 := GenS12{F0: 0, F1: 9223372036854775807, F2: -9223372036854775808}
                    v1 = v1
                    _ = v113
                }
                v3 = (!(42 > -774))
                _ = v110
            }
        } else {
            fmt.Println(v14)
            for fi114 := 0; fi114 < 18; fi114++ {
                var v115 int8 = 127
                _ = fi114
                _ = v115
            }
            var v116 string = "\n\t"
            _ = v116
        }
    case 277:
        for fi117 := 0; fi117 < 8; fi117++ {
            var v118 int = (fi117 * fi117)
            _ = fi117
            _ = v118
        }
    }
    var v119 bool = (v3 && (!v3))
    var v120 int64 = -9223372036854775808
    var v121 bool = v3
    for fi122 := 0; fi122 < 11; fi122++ {
        var v123 int32 = (v15 % (((-(v15))) | 1))
        for fi124 := 0; fi124 < 14; fi124++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi124
        }
        switch fi122 {
        case 128:
            var v125 float64 = (0.1 - -2.25)
            _ = v125
        case -9223372036854775808:
            v2 += v2
        }
        _ = fi122
        _ = v123
    }
    for fi126 := 0; fi126 < 5; fi126++ {
        var v127 float32 = (v14 * (v14 + 3.14159))
        if ((v2 % ((genFn0((v1 & (v1 ^ v1)), (("a" + "hello") + ""))) | 1)) != 2147483648) {
            v128 := GenS13{F0: fi126, F1: 625, F2: v15}
            _ = v128
        } else {
            var v129 float64 = 1.5
            _ = v129
        }
        _ = fi126
        _ = v127
    }
    v4 = append(v4, -2147483648)
    var v130 int8 = -128
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) > 0 {
        v4[0] = 283
    }
    if true {
        var v131 int16 = -93
        var v132 int64 = v120
        var v133 int8 = (^(v130))
        _ = v131
        _ = v132
        _ = v133
    }
    fmt.Println("v1:", v1)
    fmt.Println("v119:", v119)
    fmt.Println("v120:", v120)
    fmt.Println("v121:", v121)
    fmt.Println("v130:", v130)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v119
    _ = v120
    _ = v121
    _ = v13
    _ = v130
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
