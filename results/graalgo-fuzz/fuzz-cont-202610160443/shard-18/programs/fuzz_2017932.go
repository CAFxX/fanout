package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 int64; F1 int8; F2 uint }
type GenS3 struct { F0 int64; F1 int16; F2 int8 }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 string; F1 int; F2 uint32 }
type GenS7 struct { F0 int64; F1 int8 }
type GenS8 struct { F0 int32; F1 uint8; F2 int8 }
type GenS9 struct { F0 uint32; F1 uint8 }
type GenS10 struct { F0 int32; F1 int64 }
type GenS11 struct { F0 string; F1 uint32 }
type GenS12 struct { F0 uint32; F1 int16; F2 uint32 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 string; F1 int8 }
type GenS15 struct { F0 string; F1 int }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 uint64 }
func genFn0(p0 int, p1 string) uint8 {
    for fi7 := 0; fi7 < 13; fi7++ {
        p1 = (("" + (p1 + "a")) + (p1 + ("hello" + (p1 + (p1 + (p1 + ("\n\t" + (("a" + "\n\t") + p1))))))))
        if (!(!(false && (!(false && true))))) {
            v8 := GenS1{F0: 65535, F1: 128}
            _ = v8
        } else {
            var v9 int32 = 2147483647
            v10 := GenS1{F0: 550, F1: 2}
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    if false {
        for fi11 := 0; fi11 < 6; fi11++ {
            if false {
                var v12 uint64 = 267
                _ = v12
            } else {
                var v13 float64 = 0.1
                p0 = (p0 % ((128) | 1))
                _ = v13
            }
            p0 -= (p0 | 42)
            v14 := []int{}
            _ = fi11
            _ = v14
        }
    } else {
        switch p1 {
        case "":
            v15 := []int{(p0 ^ p0), (p0 * p0)}
            _ = v15
        }
        p1 = (p1 + p1)
        v16 := []int{(p0 >> ((761) & 63))}
        _ = v16
    }
    var v17 uint16 = 2
    var v18 bool = (!false)
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 251
}
func genFn1(p0 int8) uint16 {
    v19 := []int{}
    var v20 uint8 = 2
    var v21 int8 = (p0 - (p0 | (p0 % (((p0 * (p0 << ((645) & 63)))) | 1))))
    _ = p0
    _ = v19
    _ = v20
    _ = v21
    return 0
}

func main() {
    var v1 int = -1
    var v2 int8 = 92
    var v3 uint32 = 2
    var v4 uint32 = 42
    v5 := make(map[int]uint32)
    var v6 uint32 = (v3 * 4294967295)
    var v22 bool = true
    if v22 {
        fmt.Println(v6)
        fmt.Println(v22)
        for fi23 := 0; fi23 < 3; fi23++ {
            for fi24 := 0; fi24 < 5; fi24++ {
                switch fi23 {
                case 9223372036854775807:
                    v2 += v2
                    var v25 float32 = ((1e10 / -2.25) / 1.5)
                    _ = v25
                case -1:
                    var v26 uint8 = 2
                    _ = v26
                default:
                    switch "hello" {
                    case "":
                        v27 := GenS1{F0: 0, F1: genFn0(fi23, "hello")}
                        if v22 {
                            v3 += v4
                            if false {
                                if true {
                                    _ = len(v5)
                                    v2 += -1
                                } else {
                                    v1 -= fi23
                                }
                                _ = len(v5)
                                v28 := GenS1{F0: 980, F1: 0}
                                _ = v28
                            }
                            v4 -= 0
                            v29 := GenS1{F0: 4294967295, F1: 255}
                            _ = v29
                        } else {
                            var v30 int8 = v2
                            var v31 uint16 = 606
                            v32 := GenS2{F0: -9223372036854775808, F1: (v2 | v30), F2: 963}
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                        _ = v27
                    case "\n\t":
                        if true {
                            if v22 {
                                fmt.Println(fi23, fi23, v6)
                                if true {
                                    var v33 int16 = 620
                                    v1 -= (fi23 * (fi23 + fi24))
                                    _ = v33
                                } else {
                                    var v34 uint32 = (v3 & (^(v3)))
                                    _ = v34
                                }
                            }
                            switch "a" {
                            case "a":
                                v6 = (^(v4))
                            case "\n\t":
                                fmt.Println(v1)
                            }
                            if (0.0 == (3.14159 * 1.5)) {
                                if ((65535 <= 42) || v22) {
                                    var v35 uint = 11
                                    _ = v35
                                }
                                var v36 bool = false
                                v37 := v5[(v1 + fi23)]
                                _ = v36
                                _ = v37
                            } else {
                                v38 := GenS2{F0: 1, F1: v2, F2: 44}
                                v6 = v3
                                v39 := GenS3{F0: 917, F1: -1, F2: (v2 / ((v2) | 1))}
                                _ = v38
                                _ = v39
                            }
                        } else {
                            fmt.Println(v3, v2)
                            if (127 == 42) {
                                var v40 int = fi24
                                v41 := []int{v1, -938}
                                _ = v40
                                _ = v41
                            }
                        }
                    case "a":
                        v5[(fi24 + (v1 >> ((376) & 63)))] = 42
                        v42 := GenS2{F0: 42, F1: v2, F2: 1}
                        v22 = (false && true)
                        _ = v42
                    default:
                        v6 = (v3 >> ((18446744073709551615) & 63))
                    }
                }
                var v43 int8 = (v2 >> ((4294967295) & 63))
                v2 = 1
                var v44 int32 = 1
                _ = fi24
                _ = v43
                _ = v44
            }
            var v45 float64 = 0.0
            v6 -= (v6 >> ((65535) & 63))
            _ = fi23
            _ = v45
        }
    }
    if false {
        var v46 uint64 = 0
        _ = v46
    }
    for fi47 := 0; fi47 < 17; fi47++ {
        v4 += 800
        if v22 {
            var v48 float32 = 0.1
            _ = v48
        } else {
            fmt.Println(v22)
            fmt.Println(v22, v4)
            var v49 uint8 = 2
            _ = v49
        }
        if (0.0 == -2.25) {
            if (((610 == 309) && (2 >= 2)) || false) {
                switch ("a" + "a") {
                case "foo bar":
                    for fi50 := 0; fi50 < 12; fi50++ {
                        switch (v1 >> ((1) & 63)) {
                        case 127:
                            v51 := []int{-391}
                            _ = len(v5)
                            var v52 int64 = -9223372036854775808
                            _ = v51
                            _ = v52
                        }
                        v2 = (v2 & v2)
                        _ = fi50
                    }
                case "a":
                    if (936 <= 65535) {
                        _ = len(v5)
                        if v22 {
                            if false {
                                fmt.Println(v6, v2, v6)
                                v5[(v1 / ((v1) | 1))] = (v3 - v4)
                                if ((v1 / ((-774) | 1)) == fi47) {
                                    v1 = (fi47 + 127)
                                } else {
                                    v3 -= (v6 >> ((0) & 63))
                                    switch "a" {
                                    case "a":
                                        for fi53 := 0; fi53 < 2; fi53++ {
                                            switch (^(fi47)) {
                                            case 0:
                                                v5[2147483647] = v4
                                                v54 := []int{(v1 * (fi47 >> ((756) & 63)))}
                                                if (42 > 981) {
                                                    v55 := GenS4{F0: (^(v1)), F1: 32768}
                                                    if (1 != -9223372036854775808) {
                                                        v2 = (v2 * (v2 >> ((65535) & 63)))
                                                        if (0 >= 127) {
                                                            v56 := GenS5{F0: -321}
                                                            var v57 int64 = -1
                                                            var v58 string = (("" + "hello") + ("\n\t" + ("" + "")))
                                                            if v22 {
                                                                v59 := GenS3{F0: v57, F1: 32767, F2: (^(v2))}
                                                                v22 = v22
                                                                switch v58 {
                                                                case "\n\t":
                                                                    var v60 uint64 = 0
                                                                    _ = v60
                                                                case "foo bar":
                                                                    var v61 uint8 = 127
                                                                    _ = v61
                                                                case "hello":
                                                                    var v62 uint16 = 547
                                                                    _ = v62
                                                                }
                                                                var v63 int32 = -376
                                                                _ = v59
                                                                _ = v63
                                                            } else {
                                                                var v64 uint8 = genFn0((fi53 << ((614) & 63)), "hello")
                                                                v6 = v3
                                                                fmt.Println(v22, v58)
                                                                if (genFn1(-4) == 255) {
                                                                    if true {
                                                                        var v65 uint8 = v64
                                                                        _ = v65
                                                                    }
                                                                    switch "a" {
                                                                    case "a":
                                                                        v58 = ((("hello" + (((v58 + v58) + (v58 + (v58 + "a"))) + ((v58 + ("a" + "foo bar")) + v58))) + "hello") + "foo bar")
                                                                    case "foo bar":
                                                                        v66 := GenS2{F0: (v57 ^ (v57 << ((440) & 63))), F1: v2, F2: 4294967295}
                                                                        _ = v66
                                                                    case "\n\t":
                                                                        v64 -= 127
                                                                    default:
                                                                        if ((v2 + v2) <= (v2 | (v2 + v2))) {
                                                                            v67 := GenS2{F0: v57, F1: v2, F2: 4294967295}
                                                                            v4 -= (v6 % (((v3 | (v4 / (((^(v4))) | 1)))) | 1))
                                                                            var v68 uint64 = 4294967295
                                                                            _ = v67
                                                                            _ = v68
                                                                        }
                                                                    }
                                                                    if len(v54) > 0 {
                                                                        v54[0] = fi53
                                                                    }
                                                                    v69 := GenS3{F0: v57, F1: 1, F2: (v2 ^ v2)}
                                                                    fmt.Println(fi47)
                                                                    _ = v69
                                                                } else {
                                                                    var v70 float64 = 0.1
                                                                    if (!v22) {
                                                                        v64 = v64
                                                                    } else {
                                                                        if (((float32(v70) != 3.14159) && (18446744073709551615 >= 215)) || ((128 > v64) || v22)) {
                                                                            fmt.Println(v1, v6)
                                                                            var v71 uint16 = 1
                                                                            _ = v71
                                                                        }
                                                                        var v72 uint8 = 26
                                                                        _ = v72
                                                                    }
                                                                    v70 = v70
                                                                    _ = v70
                                                                }
                                                                v22 = (false && (1 >= 42))
                                                                _ = v64
                                                            }
                                                            var v73 bool = ((fi47 * (-(fi53))) == fi53)
                                                            _ = v56
                                                            _ = v57
                                                            _ = v58
                                                            _ = v73
                                                        } else {
                                                            v6 += v3
                                                            _ = len(v5)
                                                            var v74 int8 = v2
                                                            _ = v74
                                                        }
                                                        if len(v54) >= 2 {
                                                            v54 = v54[1:len(v54)]
                                                        }
                                                    }
                                                    _ = v55
                                                } else {
                                                    var v75 bool = (-2.25 == 1.5)
                                                    var v76 int64 = 0
                                                    var v77 uint16 = genFn1((^(v2)))
                                                    _ = v75
                                                    _ = v76
                                                    _ = v77
                                                }
                                                _ = v54
                                            default:
                                                v3 += 0
                                                var v78 uint = 255
                                                v3 = v4
                                                if v22 {
                                                    var v79 float32 = float32(-2.25)
                                                    _ = v79
                                                } else {
                                                    switch "" {
                                                    case "a":
                                                        var v80 bool = ((-(v2)) > v2)
                                                        if false {
                                                            fmt.Println(v22)
                                                            if v22 {
                                                                v4 += (v6 % ((1) | 1))
                                                            } else {
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    var v81 int = fi53
                                                                    _ = v81
                                                                case "\n\t":
                                                                    v1 = fi47
                                                                    var v82 uint = uint(v4)
                                                                    var v83 int64 = 746
                                                                    _ = v82
                                                                    _ = v83
                                                                case "hello":
                                                                    var v84 uint8 = genFn0(9223372036854775807, "a")
                                                                    _ = v84
                                                                }
                                                            }
                                                            v3 = v6
                                                            v85 := GenS1{F0: uint(v6), F1: 127}
                                                            if v80 {
                                                                var v86 uint32 = (v6 * (v4 >> ((v78) & 63)))
                                                                _ = v86
                                                            }
                                                            _ = v85
                                                        }
                                                        _ = v80
                                                    case "foo bar":
                                                        v78 -= 65535
                                                    case "":
                                                        var v87 int16 = 0
                                                        _ = v87
                                                    }
                                                    switch ("\n\t" + "foo bar") {
                                                    case "foo bar":
                                                        v88 := []int{(fi47 / ((fi53) | 1)), (fi47 / (((fi47 << ((364) & 63))) | 1)), v1, v1}
                                                        _ = v88
                                                    default:
                                                        v89 := GenS6{F0: ("hello" + "foo bar"), F1: -9223372036854775808, F2: (v3 | 2147483647)}
                                                        _ = v89
                                                    }
                                                }
                                                var v90 uint32 = v3
                                                var v91 uint = v78
                                                _ = v78
                                                _ = v90
                                                _ = v91
                                            }
                                            var v92 uint64 = 1
                                            _ = fi53
                                            _ = v92
                                        }
                                    case "hello":
                                        v3 -= (v3 ^ 1)
                                    default:
                                        v1 = -2147483648
                                    }
                                    for fi93 := 0; fi93 < 20; fi93++ {
                                        fmt.Println(v6, v2, fi93)
                                        fmt.Println(v2, fi47)
                                        _ = fi93
                                    }
                                }
                            }
                            if false {
                                v6 = v3
                            }
                            var v94 int64 = 9223372036854775807
                            _ = v94
                        } else {
                            if v22 {
                                if ((-2147483648 == -1) && (18446744073709551615 <= 1)) {
                                    fmt.Println(v1, v2)
                                }
                            } else {
                                v95 := []int{fi47}
                                v3 = (v3 % (((v6 << ((18446744073709551615) & 63))) | 1))
                                if (v4 != (v3 << ((18446744073709551615) & 63))) {
                                    for fi96 := 0; fi96 < 5; fi96++ {
                                        var v97 bool = (float64(1e10) <= ((3.14159 * -2.25) / 3.14159))
                                        if (v1 > fi96) {
                                            var v98 float64 = (0.0 * 1e10)
                                            _ = v98
                                        } else {
                                            v22 = (v22 && (1 >= 1))
                                            if (!((v2 != (v2 & (v2 + v2))) && (!v97))) {
                                                if (47 >= 2) {
                                                    switch "" {
                                                    case "foo bar":
                                                        fmt.Println(v2)
                                                    }
                                                    if (!v97) {
                                                        fmt.Println(v97)
                                                        if v97 {
                                                            v4 = v3
                                                            if (!(!(42 <= 42))) {
                                                                if v97 {
                                                                    var v99 uint = 542
                                                                    _ = v99
                                                                }
                                                                v95 = append(v95, 2147483647)
                                                            } else {
                                                                if v22 {
                                                                    var v100 int64 = 283
                                                                    _ = v100
                                                                } else {
                                                                    v22 = ((!true) || ((fi47 >> ((4294967295) & 63)) > v1))
                                                                }
                                                                if v97 {
                                                                    v101 := GenS3{F0: -1, F1: 0, F2: v2}
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        v5[int(v6)] = (v6 >> ((759) & 63))
                                                                    case "\n\t":
                                                                        v22 = (v3 == v6)
                                                                    case "foo bar":
                                                                        v3 = v6
                                                                    default:
                                                                        var v102 float32 = -2.25
                                                                        _ = v102
                                                                    }
                                                                    _ = v101
                                                                } else {
                                                                    if true {
                                                                        v103 := GenS7{F0: -1, F1: v2}
                                                                        switch v1 {
                                                                        case 2147483647:
                                                                            var v104 int16 = -32768
                                                                            _ = v104
                                                                        default:
                                                                            switch ("" + "\n\t") {
                                                                            case "hello":
                                                                                v1 += -508
                                                                            default:
                                                                                var v105 int = (^(v1))
                                                                                _ = v105
                                                                            }
                                                                        }
                                                                        var v106 uint64 = 4294967295
                                                                        var v107 int64 = 42
                                                                        _ = v103
                                                                        _ = v106
                                                                        _ = v107
                                                                    }
                                                                    v1 = v1
                                                                }
                                                            }
                                                            var v108 uint32 = v4
                                                            _ = v108
                                                        } else {
                                                            v109 := GenS8{F0: 42, F1: 42, F2: v2}
                                                            if (1 < 32767) {
                                                                var v110 float32 = 0.1
                                                                v111 := GenS8{F0: -2147483648, F1: 71, F2: (v2 % ((v2) | 1))}
                                                                fmt.Println(v3, v1, v22)
                                                                var v112 int64 = 9223372036854775807
                                                                var v113 int32 = 42
                                                                fmt.Println(v22, fi47)
                                                                _ = v110
                                                                _ = v111
                                                                _ = v112
                                                                _ = v113
                                                            } else {
                                                                v95 = append(v95, fi47)
                                                                v114 := GenS6{F0: "a", F1: (v1 * -442), F2: v6}
                                                                v5[fi47] = 2147483648
                                                                _ = v114
                                                            }
                                                            if v22 {
                                                                if false {
                                                                    fmt.Println(v97)
                                                                    if (v22 || (!v97)) {
                                                                        v115 := GenS2{F0: -1, F1: (v2 % ((-24) | 1)), F2: 201}
                                                                        if v22 {
                                                                            var v116 int64 = 1
                                                                            var v117 int64 = (^(v116))
                                                                            _ = v116
                                                                            _ = v117
                                                                        } else {
                                                                            switch fi47 {
                                                                            case -9223372036854775808:
                                                                                v118 := v5[(fi47 >> ((42) & 63))]
                                                                                _ = v118
                                                                            case 32768:
                                                                                v97 = v22
                                                                            case 0:
                                                                                if v22 {
                                                                                    v2 = v2
                                                                                    if len(v95) >= 2 {
                                                                                        v95 = v95[1:len(v95)]
                                                                                    }
                                                                                    var v119 uint32 = v6
                                                                                    var v120 uint16 = 0
                                                                                    v4 = v119
                                                                                    _ = v119
                                                                                    _ = v120
                                                                                }
                                                                                if ((3.14159 * 3.14159) > float32(0.0)) {
                                                                                    v4 = 4294967295
                                                                                    var v121 int = -2147483648
                                                                                    var v122 float32 = (3.14159 + float32(float64(0.1)))
                                                                                    v123 := v5[(v121 - 1)]
                                                                                    var v124 int64 = -1
                                                                                    _ = v121
                                                                                    _ = v122
                                                                                    _ = v123
                                                                                    _ = v124
                                                                                }
                                                                                v97 = false
                                                                            default:
                                                                                fmt.Println(v97, v1)
                                                                            }
                                                                            var v125 int8 = v2
                                                                            var v126 int32 = 1
                                                                            _ = v125
                                                                            _ = v126
                                                                        }
                                                                        var v127 int16 = 42
                                                                        if v22 {
                                                                            v3 = (v3 << ((2) & 63))
                                                                            var v128 int16 = (v127 / ((0) | 1))
                                                                            v129 := GenS5{F0: 1}
                                                                            var v130 int32 = 0
                                                                            switch ("foo bar" + ("foo bar" + "hello")) {
                                                                            case "foo bar":
                                                                                v22 = (v97 || (v97 && true))
                                                                                v2 = (^(v2))
                                                                            case "\n\t":
                                                                                fmt.Println(v1, v128)
                                                                            case "a":
                                                                                var v131 float32 = -2.25
                                                                                v22 = (v127 <= 1)
                                                                                _ = v131
                                                                            default:
                                                                                fmt.Println(v127, v6, v127)
                                                                            }
                                                                            _ = v128
                                                                            _ = v129
                                                                            _ = v130
                                                                        }
                                                                        fmt.Println(v6, v6, v127)
                                                                        _ = v115
                                                                        _ = v127
                                                                    }
                                                                    var v132 float64 = ((0.1 / float64(float32(1.5))) + float64(1e10))
                                                                    v22 = (true && (18446744073709551615 != 1))
                                                                    _ = v132
                                                                }
                                                                if (42 != -1) {
                                                                    delete(v5, (v1 % ((fi96) | 1)))
                                                                    fmt.Println(v22, fi96)
                                                                    delete(v5, v1)
                                                                    var v133 uint32 = (v3 - v6)
                                                                    _ = v133
                                                                }
                                                            }
                                                            fmt.Println(fi47, v1)
                                                            _ = v109
                                                        }
                                                    } else {
                                                        v97 = (1e10 != 0.1)
                                                        var v134 string = (("\n\t" + "foo bar") + "foo bar")
                                                        var v135 string = v134
                                                        v136 := GenS9{F0: (v6 * 42), F1: 219}
                                                        switch ((((((v134 + v134) + "\n\t") + "") + ((v134 + (((v135 + v134) + v134) + "a")) + ((v135 + ((v134 + (("hello" + v134) + "\n\t")) + v134)) + "\n\t"))) + v134) + v134) {
                                                        case "\n\t":
                                                            fmt.Println(v135, v1)
                                                        default:
                                                            var v137 float64 = 1.5
                                                            _ = v137
                                                        }
                                                        var v138 uint64 = 1
                                                        _ = v134
                                                        _ = v135
                                                        _ = v136
                                                        _ = v138
                                                    }
                                                    v5[v1] = v6
                                                } else {
                                                    v97 = (65535 < 65535)
                                                    v1 = (^(v1))
                                                    if (v2 > (v2 >> ((1) & 63))) {
                                                        fmt.Println(v22)
                                                        var v139 uint64 = 51
                                                        v140 := GenS2{F0: 864, F1: v2, F2: 2}
                                                        _ = v139
                                                        _ = v140
                                                    } else {
                                                        var v141 uint64 = 42
                                                        fmt.Println(v97, fi96)
                                                        fmt.Println(v22, v3)
                                                        _ = v141
                                                    }
                                                    fmt.Println(v2, fi47, v1)
                                                }
                                            } else {
                                                var v142 uint16 = 255
                                                var v143 uint8 = 220
                                                _ = v142
                                                _ = v143
                                            }
                                            fmt.Println(v6)
                                            if len(v95) >= 2 {
                                                v95 = v95[1:len(v95)]
                                            }
                                            fmt.Println(fi96, v4)
                                        }
                                        v4 = v4
                                        v6 = 2
                                        _ = fi96
                                        _ = v97
                                    }
                                    var v144 uint32 = 42
                                    var v145 uint = 65535
                                    _ = v144
                                    _ = v145
                                } else {
                                    v4 = 524
                                    for fi146 := 0; fi146 < 13; fi146++ {
                                        v147 := v5[-781]
                                        var v148 uint32 = 338
                                        switch "foo bar" {
                                        case "hello":
                                            switch "foo bar" {
                                            case "hello":
                                                var v149 bool = (475 < 1)
                                                _ = v149
                                            case "":
                                                v150 := GenS3{F0: -605, F1: 14, F2: (v2 * (v2 / (((v2 | v2)) | 1)))}
                                                _ = v150
                                            }
                                        }
                                        if len(v95) >= 2 {
                                            v95 = v95[1:len(v95)]
                                        }
                                        _ = fi146
                                        _ = v147
                                        _ = v148
                                    }
                                }
                                _ = v95
                            }
                            var v151 uint32 = (v6 ^ (v3 * 2147483648))
                            _ = v151
                        }
                    } else {
                        if (((86 >= 128) && (179 >= 4294967295)) || v22) {
                            var v152 int8 = v2
                            _ = len(v5)
                            _ = v152
                        } else {
                            var v153 int16 = 1
                            switch v1 {
                            case -890:
                                for fi154 := 0; fi154 < 13; fi154++ {
                                    v155 := GenS10{F0: 464, F1: -1}
                                    switch "hello" {
                                    case "":
                                        v156 := GenS3{F0: 0, F1: (v153 | 1000), F2: v2}
                                        _ = v156
                                    }
                                    fmt.Println(fi154)
                                    _ = fi154
                                    _ = v155
                                }
                            default:
                                switch ("\n\t" + "") {
                                case "hello":
                                    for fi157 := 0; fi157 < 3; fi157++ {
                                        var v158 float64 = -2.25
                                        v159 := GenS3{F0: 0, F1: 1, F2: (v2 + (v2 & int8(fi157)))}
                                        _ = fi157
                                        _ = v158
                                        _ = v159
                                    }
                                case "\n\t":
                                    if v22 {
                                        fmt.Println(v1)
                                        var v160 string = ""
                                        switch "\n\t" {
                                        case "\n\t":
                                            for fi161 := 0; fi161 < 4; fi161++ {
                                                v162 := GenS1{F0: 0, F1: 243}
                                                switch (fi161 | 67) {
                                                case -9223372036854775808:
                                                    v163 := GenS6{F0: v160, F1: (v1 | 1), F2: 420}
                                                    _ = v163
                                                case -2147483648:
                                                    v5[fi47] = 0
                                                    if false {
                                                        var v164 uint32 = (v6 / ((2147483647) | 1))
                                                        var v165 int8 = v2
                                                        v164 = (^(v4))
                                                        _ = v164
                                                        _ = v165
                                                    } else {
                                                        var v166 uint16 = 2
                                                        var v167 int32 = 0
                                                        v168 := GenS6{F0: v160, F1: (-(fi161)), F2: 1}
                                                        var v169 int16 = (v153 * (v153 / ((v153) | 1)))
                                                        var v170 bool = v22
                                                        _ = v166
                                                        _ = v167
                                                        _ = v168
                                                        _ = v169
                                                        _ = v170
                                                    }
                                                }
                                                v171 := GenS11{F0: v160, F1: v4}
                                                switch -585 {
                                                case 698:
                                                    v172 := GenS12{F0: (v4 + (v3 * 2)), F1: v153, F2: (v3 ^ v6)}
                                                    _ = v172
                                                default:
                                                    v160 = (((("a" + (("" + v160) + (((v160 + v160) + v160) + "a"))) + v160) + v160) + ("foo bar" + v160))
                                                }
                                                v3 = (v6 & (v4 | 1))
                                                v22 = (1 <= 131)
                                                _ = fi161
                                                _ = v162
                                                _ = v171
                                            }
                                        case "":
                                            v153 = (v153 ^ 253)
                                        }
                                        _ = v160
                                    }
                                case "foo bar":
                                    var v173 uint8 = 0
                                    _ = v173
                                }
                            }
                            _ = v153
                        }
                        for fi174 := 0; fi174 < 11; fi174++ {
                            fmt.Println(v3, v3)
                            fmt.Println(v4, v22, v1)
                            v175 := GenS13{F0: -1}
                            _ = fi174
                            _ = v175
                        }
                    }
                default:
                    switch ("hello" + "hello") {
                    case "hello":
                        for fi176 := 0; fi176 < 17; fi176++ {
                            var v177 int32 = -1
                            var v178 int8 = (v2 & v2)
                            switch (fi47 & 1) {
                            case 2147483647:
                                fmt.Println(v3, v178)
                            case 9223372036854775807:
                                v3 = (^(v4))
                            case -225:
                                switch (v1 & (v1 - -1)) {
                                case 586:
                                    v179 := []int{fi176, fi176, v1, -505}
                                    _ = v179
                                }
                            default:
                                var v180 string = (("a" + "hello") + ("a" + "a"))
                                _ = v180
                            }
                            _ = fi176
                            _ = v177
                            _ = v178
                        }
                    default:
                        var v181 uint32 = v3
                        _ = v181
                    }
                }
                delete(v5, fi47)
                if v22 {
                    var v182 uint64 = 867
                    fmt.Println(v3, v3)
                    var v183 uint = 218
                    _ = v182
                    _ = v183
                }
            }
            for fi184 := 0; fi184 < 3; fi184++ {
                switch (fi47 ^ (fi184 - fi47)) {
                case -583:
                    fmt.Println(v6)
                case 606:
                    fmt.Println(v4)
                }
                v185 := GenS14{F0: "", F1: 60}
                _ = fi184
                _ = v185
            }
        }
        _ = fi47
    }
    var v186 int8 = v2
    if v22 {
        v3 -= v4
    }
    var v187 uint = 318
    var v188 uint32 = v6
    switch -572 {
    case -111:
        for fi189 := 0; fi189 < 2; fi189++ {
            switch "foo bar" {
            case "hello":
                fmt.Println(v187, v1)
            case "a":
                fmt.Println(v187, v188)
            default:
                v190 := GenS15{F0: "", F1: (v1 ^ fi189)}
                _ = v190
            }
            fmt.Println(v4, v3, v2)
            _ = fi189
        }
    case 238:
        v22 = (18446744073709551615 <= uint64(v187))
    }
    if v22 {
        v191 := []int{v1, v1, v1, v1}
        var v192 uint = (v187 % ((v187) | 1))
        fmt.Println(v4, v22)
        _ = v191
        _ = v192
    } else {
        if v22 {
            v193 := []int{(v1 ^ (^(v1))), v1, (v1 ^ int(v186)), 127}
            v194 := GenS16{F0: (false || v22)}
            v1 = v1
            _ = v193
            _ = v194
        } else {
            v3 = 42
            v195 := GenS17{F0: 1}
            _ = v195
        }
    }
    var v196 int = (v1 * (-(v1)))
    v5[v1] = (v188 + v6)
    fmt.Println("v1:", v1)
    fmt.Println("v186:", v186)
    fmt.Println("v187:", v187)
    fmt.Println("v188:", v188)
    fmt.Println("v196:", v196)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v186
    _ = v187
    _ = v188
    _ = v196
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
