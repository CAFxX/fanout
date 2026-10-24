package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 uint; F2 uint32 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint16; F1 int32 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int32; F1 uint8 }
type GenS12 struct { F0 int8; F1 string }
type GenS13 struct { F0 uint16; F1 int8; F2 string }
type GenS14 struct { F0 uint64; F1 string; F2 uint16 }
type GenS15 struct { F0 uint8; F1 int; F2 int }
type GenS16 struct { F0 uint8; F1 int; F2 bool }
func genFn0(p0 bool) int32 {
    p0 = true
    var v4 uint16 = 1
    var v5 uint = 255
    _ = p0
    _ = v4
    _ = v5
    return -1
}
func genFn1(p0 uint8) bool {
    var v6 uint64 = 42
    var v7 int64 = -1
    v6 = (v6 & (v6 - v6))
    fmt.Println(v7, v6)
    if ((v6 / ((v6) | 1)) < 4294967295) {
        for fi8 := 0; fi8 < 7; fi8++ {
            for fi9 := 0; fi9 < 12; fi9++ {
                fmt.Println(fi8)
                _ = fi9
            }
            _ = fi8
        }
    } else {
        var v10 uint8 = 1
        v11 := []int{-67}
        for fi12 := 0; fi12 < 11; fi12++ {
            fmt.Println(p0, p0, v10)
            fmt.Println(fi12, v6)
            if ((v7 | (v7 & (v7 & v7))) >= (v7 - -1)) {
                v13 := GenS1{F0: 1000}
                var v14 bool = false
                _ = v13
                _ = v14
            }
            _ = fi12
        }
        _ = v10
        _ = v11
    }
    _ = p0
    _ = v6
    _ = v7
    return false
}

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 731
    v3 := make(map[uint32]bool)
    v2 += 897
    if (((0.0 + (1.5 * ((-2.25 * float32((float64((1.5 + (0.0 + -2.25))) + 1e10))) / 0.1))) > 1.5) && true) {
        for fi15 := 0; fi15 < 5; fi15++ {
            var v16 uint16 = 0
            _ = fi15
            _ = v16
        }
    } else {
        fmt.Println(v2)
        switch "foo bar" {
        case "":
            fmt.Println(v2)
        case "\n\t":
            if ((false && true) || true) {
                for fi17 := 0; fi17 < 6; fi17++ {
                    fmt.Println(fi17, fi17)
                    fmt.Println(v2, v2)
                    switch "" {
                    case "":
                        switch fi17 {
                        case 4294967295:
                            var v18 int = (^(fi17))
                            _ = v18
                        case -817:
                            v2 -= (v2 + v2)
                        }
                    default:
                        switch 128 {
                        case 1:
                            v19 := GenS2{F0: 2147483648, F1: 829, F2: uint32(v2)}
                            _ = v19
                        case 32768:
                            var v20 uint64 = v2
                            _ = v20
                        case 42:
                            v2 = (v2 * (^(v2)))
                        default:
                            var v21 string = ("foo bar" + "hello")
                            _ = v21
                        }
                    }
                    _ = fi17
                }
                if (!false) {
                    if (!false) {
                        v2 = 745
                    }
                } else {
                    var v22 string = "\n\t"
                    if false {
                        if (321 != 1) {
                            var v23 int64 = 106
                            for fi24 := 0; fi24 < 18; fi24++ {
                                if ((42 >= 2147483647) && (-1 == -159)) {
                                    v25 := GenS1{F0: 0}
                                    _ = v25
                                } else {
                                    if true {
                                        var v26 float64 = 1e10
                                        _ = v26
                                    } else {
                                        delete(v1, v22)
                                        v23 += 0
                                    }
                                    v27 := GenS3{F0: 323, F1: 127}
                                    _ = v27
                                }
                                v2 -= (v2 % ((v2) | 1))
                                _ = fi24
                            }
                            var v28 uint64 = 18446744073709551615
                            _ = v23
                            _ = v28
                        } else {
                            switch v22 {
                            case "foo bar":
                                var v29 int8 = 1
                                _ = v29
                            case "\n\t":
                                v30 := GenS4{F0: 100}
                                _ = v30
                            }
                            v22 = "hello"
                            for fi31 := 0; fi31 < 15; fi31++ {
                                var v32 uint64 = (v2 / ((2) | 1))
                                fmt.Println(v2, v2, fi31)
                                switch (v22 + "hello") {
                                case "hello":
                                    _ = len(v1)
                                case "a":
                                    switch v22 {
                                    case "a":
                                        switch 1 {
                                        case 42:
                                            switch fi31 {
                                            case -479:
                                                var v33 uint = 2
                                                _ = v33
                                            case -636:
                                                v2 -= 181
                                            default:
                                                v34 := GenS1{F0: 18446744073709551615}
                                                _ = v34
                                            }
                                        }
                                    case "":
                                        v32 -= v32
                                    }
                                case "\n\t":
                                    fmt.Println(v2, v22, fi31)
                                default:
                                    switch (fi31 ^ (fi31 / (((-(fi31))) | 1))) {
                                    case 304:
                                        fmt.Println(v22)
                                    case 127:
                                        switch (((((("" + v22) + v22) + "hello") + v22) + v22) + v22) {
                                        case "":
                                            fmt.Println(fi31)
                                        case "a":
                                            for fi35 := 0; fi35 < 3; fi35++ {
                                                var v36 uint16 = 42
                                                _ = fi35
                                                _ = v36
                                            }
                                        case "\n\t":
                                            for fi37 := 0; fi37 < 6; fi37++ {
                                                fmt.Println(fi31)
                                                fmt.Println(fi31, fi31)
                                                var v38 int32 = genFn0(false)
                                                _ = fi37
                                                _ = v38
                                            }
                                        default:
                                            v39 := GenS1{F0: 150}
                                            _ = v39
                                        }
                                    case 42:
                                        v40 := GenS5{F0: v22, F1: -1}
                                        _ = v40
                                    default:
                                        var v41 uint64 = 2
                                        _ = v41
                                    }
                                }
                                _ = fi31
                                _ = v32
                            }
                        }
                        v2 = (v2 | (v2 ^ v2))
                        var v42 int64 = 0
                        _ = v42
                    }
                    _ = v22
                }
                v2 += 1
            }
        case "foo bar":
            v2 += (^(v2))
        default:
            var v43 uint = 158
            _ = v43
        }
    }
    if false {
        switch 4294967295 {
        case 0:
            v2 += (v2 | (v2 + v2))
        case 32768:
            fmt.Println(v2)
        default:
            for fi44 := 0; fi44 < 7; fi44++ {
                for fi45 := 0; fi45 < 17; fi45++ {
                    v46 := GenS6{F0: 143}
                    var v47 uint32 = uint32(v2)
                    _ = fi45
                    _ = v46
                    _ = v47
                }
                v48 := []int{(fi44 + (fi44 & fi44)), (fi44 * fi44), fi44}
                _ = fi44
                _ = v48
            }
        }
        v49 := GenS5{F0: "hello", F1: 0}
        _ = v49
    } else {
        if true {
            v50 := GenS7{F0: 601}
            _ = v50
        }
        v3[2] = false
    }
    if false {
        v2 = 0
        if (2 >= 1) {
            if (-1 < -1) {
                fmt.Println(v2)
            } else {
                if false {
                    var v51 int = -55
                    fmt.Println(v2, v2)
                    _ = v51
                } else {
                    v52 := v3[845]
                    for fi53 := 0; fi53 < 3; fi53++ {
                        v52 = (1 <= -439)
                        if v52 {
                            switch fi53 {
                            case 32768:
                                if (2147483647 <= 1) {
                                    var v54 uint = 346
                                    _ = v54
                                } else {
                                    fmt.Println(v2)
                                    switch (fi53 | 127) {
                                    case -158:
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v55 uint64 = (v2 + (v2 % (((v2 - 1)) | 1)))
                                            _ = v55
                                        }
                                    case 4294967295:
                                        var v56 uint8 = 1
                                        _ = v56
                                    default:
                                        v52 = v52
                                    }
                                }
                            case -517:
                                var v57 uint64 = (v2 ^ (v2 - 616))
                                _ = v57
                            default:
                                if v52 {
                                    var v58 int64 = 42
                                    if (!(v52 && (!false))) {
                                        if v52 {
                                            for fi59 := 0; fi59 < 18; fi59++ {
                                                v60 := GenS8{F0: 8}
                                                var v61 string = "foo bar"
                                                v62 := GenS9{F0: 482, F1: 2147483647}
                                                v63 := []int{(fi53 >> ((42) & 63)), (fi53 >> ((18446744073709551615) & 63))}
                                                v2 += (v2 - 0)
                                                var v64 string = v61
                                                _ = fi59
                                                _ = v60
                                                _ = v61
                                                _ = v62
                                                _ = v63
                                                _ = v64
                                            }
                                            v52 = (1 <= (v58 ^ v58))
                                            for fi65 := 0; fi65 < 9; fi65++ {
                                                var v66 int16 = 424
                                                switch ("hello" + (("hello" + ((("" + "") + "\n\t") + "")) + "a")) {
                                                case "foo bar":
                                                    v67 := GenS1{F0: 361}
                                                    _ = v67
                                                case "a":
                                                    v2 = (v2 % ((42) | 1))
                                                case "hello":
                                                    var v68 int16 = v66
                                                    var v69 uint8 = 1
                                                    if (v52 || true) {
                                                        var v70 int16 = v68
                                                        fmt.Println(v58, v70)
                                                        if v52 {
                                                            v71 := []int{(^(fi53)), fi65, (fi53 ^ (-(fi53))), (fi65 % ((-9223372036854775808) | 1))}
                                                            _ = v71
                                                        } else {
                                                            v58 -= v58
                                                        }
                                                        var v72 int64 = (v58 ^ v58)
                                                        switch ("\n\t" + "\n\t") {
                                                        case "foo bar":
                                                            var v73 float64 = (1.5 / float64((0.0 + 1e10)))
                                                            _ = v73
                                                        case "hello":
                                                            switch fi65 {
                                                            case -258:
                                                                var v74 float32 = 0.1
                                                                _ = v74
                                                            case 32767:
                                                                var v75 int32 = -2147483648
                                                                _ = v75
                                                            default:
                                                                v2 = 419
                                                            }
                                                        case "a":
                                                            v76 := []int{(fi53 >> ((0) & 63)), fi53, fi65}
                                                            v77 := GenS3{F0: 2147483648, F1: -1}
                                                            _ = v76
                                                            _ = v77
                                                        }
                                                        if (!v52) {
                                                            v72 = (-(v58))
                                                            v78 := []int{fi53, (fi53 * (fi65 - (fi53 >> ((65535) & 63)))), fi53, (fi53 >> ((772) & 63))}
                                                            v79 := GenS10{F0: 42}
                                                            if v52 {
                                                                var v80 int16 = v68
                                                                var v81 int64 = v72
                                                                if true {
                                                                    _ = len(v3)
                                                                }
                                                                _ = v80
                                                                _ = v81
                                                            }
                                                            var v82 int8 = 127
                                                            _ = v78
                                                            _ = v79
                                                            _ = v82
                                                        } else {
                                                            var v83 float64 = ((1.5 - (1e10 / -2.25)) * 0.1)
                                                            _ = v83
                                                        }
                                                        _ = v70
                                                        _ = v72
                                                    }
                                                    _ = v68
                                                    _ = v69
                                                }
                                                if (!false) {
                                                    v58 -= (-(v58))
                                                }
                                                var v84 int32 = -162
                                                v52 = ((0 <= 173) || v52)
                                                _ = fi65
                                                _ = v66
                                                _ = v84
                                            }
                                        }
                                    }
                                    _ = v58
                                }
                            }
                            v85 := GenS3{F0: 0, F1: -125}
                            delete(v3, 92)
                            _ = v85
                        }
                        v2 = (v2 + (v2 | 18446744073709551615))
                        _ = fi53
                    }
                    _ = v52
                }
            }
            fmt.Println(v2)
            fmt.Println(v2)
        }
    }
    var v86 int32 = 381
    var v87 int64 = 0
    var v88 int16 = 42
    var v89 uint16 = 2
    v90 := GenS11{F0: 2147483647, F1: 128}
    v91 := GenS12{F0: 1, F1: "foo bar"}
    switch -1 {
    case 4294967295:
        v92 := GenS13{F0: 42, F1: -1, F2: (("hello" + "a") + ((("\n\t" + (("a" + "\n\t") + "\n\t")) + "") + "a"))}
        _ = v92
    case -2147483648:
        v88 = v88
    }
    v89 = 0
    v86 = v86
    var v93 int16 = v88
    fmt.Println(v89)
    if ((v2 > v2) || false) {
        for fi94 := 0; fi94 < 4; fi94++ {
            var v95 bool = true
            if (-28 == 1) {
                var v96 uint = 2
                if (!(((-1 < (v87 - v87)) || v95) && v95)) {
                    if v95 {
                        for fi97 := 0; fi97 < 3; fi97++ {
                            var v98 int64 = (v87 % ((v87) | 1))
                            v99 := GenS14{F0: (^(v2)), F1: ("foo bar" + (("hello" + "a") + "a")), F2: (v89 & uint16(v96))}
                            _ = fi97
                            _ = v98
                            _ = v99
                        }
                    }
                } else {
                    fmt.Println(v96, v86, v95)
                    v2 = v2
                    for fi100 := 0; fi100 < 18; fi100++ {
                        var v101 string = ("foo bar" + "a")
                        _ = fi100
                        _ = v101
                    }
                }
                for fi102 := 0; fi102 < 12; fi102++ {
                    var v103 uint32 = 42
                    _ = fi102
                    _ = v103
                }
                _ = v96
            } else {
                var v104 uint16 = 42
                _ = v104
            }
            fmt.Println(fi94)
            _ = fi94
            _ = v95
        }
        for fi105 := 0; fi105 < 6; fi105++ {
            for fi106 := 0; fi106 < 17; fi106++ {
                fmt.Println(v2, v86)
                if ((!false) && true) {
                    v93 = -222
                } else {
                    v107 := GenS15{F0: 42, F1: 32768, F2: -402}
                    _ = v107
                }
                v87 -= (-(v87))
                switch fi106 {
                case 127:
                    var v108 uint = 55
                    _ = v108
                case 0:
                    if true {
                        v86 = genFn0(("a" >= ""))
                    }
                case -7:
                    var v109 uint64 = (v2 / ((uint64(v89)) | 1))
                    _ = v109
                default:
                    if (!true) {
                        var v110 int8 = -128
                        _ = v110
                    }
                }
                _ = fi106
            }
            for fi111 := 0; fi111 < 15; fi111++ {
                if (!(42 > int64(v2))) {
                    if false {
                        if true {
                            var v112 int64 = (v87 % ((v87) | 1))
                            v113 := []int{(fi105 * (fi111 * 32767)), fi111, -2147483648}
                            var v114 int64 = v87
                            _ = v112
                            _ = v113
                            _ = v114
                        }
                    } else {
                        var v115 int64 = int64(v2)
                        var v116 int32 = v86
                        _ = v115
                        _ = v116
                    }
                    fmt.Println(v86, v86, v88)
                    switch (fi111 + int(v87)) {
                    case -2147483648:
                        switch (fi111 & 1) {
                        case 127:
                            switch fi105 {
                            case 32767:
                                var v117 int64 = (v87 - (v87 + (v87 >> ((398) & 63))))
                                _ = v117
                            case 9223372036854775807:
                                v118 := []int{}
                                _ = v118
                            }
                        }
                    case 0:
                        v3[4294967295] = true
                    case 127:
                        v87 = v87
                    default:
                        if (-128 >= 0) {
                            var v119 int = (fi111 % ((-1) | 1))
                            v2 += v2
                            v88 = int16(v89)
                            _ = v119
                        } else {
                            var v120 int = fi111
                            v2 -= 598
                            _ = v120
                        }
                    }
                    v121 := GenS9{F0: v89, F1: (v86 << ((1) & 63))}
                    if false {
                        if false {
                            fmt.Println(v87, v89)
                            v122 := []int{}
                            v2 = (^(v2))
                            _ = v122
                        } else {
                            if (("a" + "foo bar") != ("\n\t" + "")) {
                                v86 += (^(v86))
                            }
                            var v123 uint64 = 4294967295
                            _ = v123
                        }
                    }
                    _ = v121
                } else {
                    var v124 bool = true
                    v125 := GenS16{F0: 2, F1: (fi105 - fi105), F2: ((int(v86) > (fi105 ^ fi105)) && ((!v124) && v124))}
                    v126 := v3[863]
                    _ = v124
                    _ = v125
                    _ = v126
                }
                v89 = v89
                if false {
                    v86 = genFn0(((v86 & (^(v86))) >= genFn0(true)))
                    var v127 uint8 = 30
                    delete(v3, 36)
                    _ = v127
                }
                _ = fi111
            }
            _ = fi105
        }
        switch 32768 {
        case 1:
            v86 = (v86 + (v86 >> ((808) & 63)))
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v86
    _ = v87
    _ = v88
    _ = v89
    _ = v90
    _ = v91
    _ = v93
}
