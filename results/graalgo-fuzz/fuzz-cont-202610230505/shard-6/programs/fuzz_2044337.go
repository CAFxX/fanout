package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 string; F1 uint64; F2 string }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 uint64; F1 int; F2 int }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint16; F1 uint64; F2 uint }
type GenS9 struct { F0 int; F1 string }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 int32; F1 uint64; F2 int }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 bool; F1 uint64 }
type GenS16 struct { F0 bool; F1 int }
type GenS17 struct { F0 int32; F1 int16 }
func genFn0(p0 uint64, p1 int16) int32 {
    switch "" {
    case "\n\t":
        fmt.Println(p1)
    case "a":
        p1 += 716
    case "":
        switch 32767 {
        case 32767:
            p1 += (p1 % ((p1) | 1))
        case -1:
            var v7 float32 = 0.0
            _ = v7
        case 640:
            p0 = 2
        default:
            var v8 int16 = p1
            _ = v8
        }
    default:
        p0 = 210
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 string, p1 uint64, p2 bool) uint {
    var v9 uint8 = 42
    switch -897 {
    case 128:
        v9 -= (v9 * (v9 + v9))
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        v11 := []int{(-(fi10)), 42, (fi10 / ((2147483647) | 1)), (fi10 | (fi10 / ((-2147483648) | 1)))}
        v12 := GenS1{F0: -345}
        var v13 string = ("hello" + "hello")
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 38
}

func main() {
    var v1 int8 = 42
    var v2 string = "\n\t"
    v3 := []string{v2, v2, v2}
    v4 := make(map[int]bool)
    v5 := make(map[uint32]int)
    var v6 int8 = (v1 - (v1 + -20))
    v6 += v6
    for fi14 := 0; fi14 < 9; fi14++ {
        v15 := GenS2{F0: 255}
        var v16 int64 = 42
        v17 := GenS1{F0: (v16 | 1)}
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        fmt.Println(v1, fi18)
        _ = fi18
    }
    v4[127] = false
    v6 = (v1 + (-(v6)))
    if (!true) {
        if true {
            var v19 uint8 = 255
            _ = v19
        }
    } else {
        var v20 uint8 = 2
        _ = v20
    }
    var v21 uint = 2
    v22 := GenS3{F0: 845, F1: -2147483648}
    fmt.Println(v21, v6)
    var v23 int64 = -245
    fmt.Println(v1)
    fmt.Println(v6, v6)
    for fi24 := 0; fi24 < 14; fi24++ {
        if false {
            if (-2.25 > -2.25) {
                v6 = int8(fi24)
                var v25 int16 = 58
                _ = v25
            } else {
                v26 := GenS1{F0: (v23 ^ int64(fi24))}
                _ = v26
            }
            var v27 float32 = (3.14159 + (3.14159 * 3.14159))
            var v28 string = "\n\t"
            _ = v27
            _ = v28
        } else {
            v29 := GenS1{F0: (v23 % (((v23 | (v23 ^ (v23 * (v23 | 9223372036854775807))))) | 1))}
            switch 32767 {
            case -9223372036854775808:
                for fi30 := 0; fi30 < 13; fi30++ {
                    var v31 int8 = 42
                    var v32 int8 = v6
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            case 128:
                for fi33 := 0; fi33 < 7; fi33++ {
                    v2 = v2
                    if (0.1 >= float64((3.14159 + 0.0))) {
                        v3 = append(v3, v2)
                        v34 := GenS1{F0: -1}
                        var v35 int32 = -1
                        _ = v34
                        _ = v35
                    }
                    if false {
                        delete(v5, 1)
                        var v36 uint = genFn1("", 4294967295, (158 != 1))
                        var v37 int = fi33
                        var v38 int16 = -32768
                        if false {
                            v37 = (-(v37))
                            switch (fi24 ^ (-(v37))) {
                            case -898:
                                switch fi33 {
                                case 1:
                                    v4[v37] = (false && (v38 >= (^(v38))))
                                case -858:
                                    var v39 uint32 = 65535
                                    switch (v37 % ((335) | 1)) {
                                    case 2147483647:
                                        if (!true) {
                                            if (v2 != "a") {
                                                var v40 int8 = 42
                                                _ = v40
                                            } else {
                                                v23 -= -1
                                                var v41 uint64 = 200
                                                _ = v41
                                            }
                                            if (!(0 == 2)) {
                                                var v42 int64 = int64(v38)
                                                if len(v3) > 0 {
                                                    v3[0] = ("" + ("a" + (("hello" + (("" + (v2 + v2)) + v2)) + (v2 + v2))))
                                                }
                                                v21 -= v21
                                                _ = v42
                                            } else {
                                                var v43 string = (v2 + v2)
                                                _ = v43
                                            }
                                            var v44 string = v2
                                            _ = v44
                                        } else {
                                            var v45 uint32 = (v39 * (v39 & 2147483648))
                                            _ = v45
                                        }
                                    case -9223372036854775808:
                                        v36 = v21
                                    case 42:
                                        if ((v21 % ((v21) | 1)) >= 18446744073709551615) {
                                            v46 := GenS2{F0: 42}
                                            _ = v46
                                        } else {
                                            if (-9223372036854775808 >= -1) {
                                                fmt.Println(fi24, v6)
                                                var v47 uint64 = 2
                                                _ = v47
                                            }
                                        }
                                    }
                                    _ = v39
                                default:
                                    v2 = v2
                                }
                            case -33:
                                if len(v3) > 0 {
                                    v3[0] = (v2 + v2)
                                }
                                if (!true) {
                                    v1 += (v1 / ((-128) | 1))
                                    if len(v3) > 0 {
                                        v3[0] = v2
                                    }
                                }
                            default:
                                if false {
                                    v37 = fi33
                                }
                                delete(v4, (v37 | fi24))
                                v48 := GenS3{F0: v38, F1: -197}
                                _ = v48
                            }
                        }
                        v49 := GenS4{F0: "\n\t", F1: 0, F2: v2}
                        v37 = 9223372036854775807
                        _ = v36
                        _ = v37
                        _ = v38
                        _ = v49
                    } else {
                        v23 = (v23 & 1)
                    }
                    _ = fi33
                }
            default:
                v2 = "hello"
            }
            _ = v29
        }
        _ = fi24
    }
    v21 = (v21 ^ v21)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    delete(v4, -816)
    if false {
        for fi50 := 0; fi50 < 14; fi50++ {
            var v51 int64 = v23
            v52 := GenS5{F0: (v2 + ("" + "hello")), F1: v23}
            if false {
                switch fi50 {
                case 128:
                    var v53 uint8 = 128
                    _ = v53
                case -2147483648:
                    var v54 int = 4
                    _ = v54
                case 1:
                    var v55 uint = 18446744073709551615
                    _ = v55
                }
                var v56 string = v2
                v1 += v1
                _ = v56
            }
            _ = fi50
            _ = v51
            _ = v52
        }
        if false {
            v57 := GenS3{F0: 0, F1: 1}
            var v58 int64 = int64(v21)
            switch 4294967295 {
            case 522:
                if false {
                    v58 = -1
                    v6 += v6
                    v59 := GenS3{F0: 1000, F1: -957}
                    _ = v59
                }
            case 1:
                var v60 uint = (v21 / ((v21) | 1))
                _ = v60
            case -516:
                for fi61 := 0; fi61 < 13; fi61++ {
                    switch 53 {
                    case -983:
                        for fi62 := 0; fi62 < 19; fi62++ {
                            switch 730 {
                            case -1:
                                if true {
                                    v58 = v23
                                    fmt.Println(v1, v58)
                                    fmt.Println(v1)
                                } else {
                                    fmt.Println(v6)
                                    v21 += (^(v21))
                                    if true {
                                        var v63 int16 = 0
                                        if true {
                                            switch -520 {
                                            case -1:
                                                var v64 int16 = 0
                                                _ = v64
                                            case -2147483648:
                                                v6 = v6
                                            case 4294967295:
                                                v21 = (v21 >> (((v21 & (^(v21)))) & 63))
                                            default:
                                                if false {
                                                    var v65 int16 = v63
                                                    if true {
                                                        v6 = 42
                                                    }
                                                    _ = v65
                                                } else {
                                                    switch v2 {
                                                    case "":
                                                        var v66 float64 = 0.0
                                                        _ = v66
                                                    case "hello":
                                                        fmt.Println(v1, v1)
                                                    case "a":
                                                        v1 += -101
                                                    default:
                                                        v21 = v21
                                                    }
                                                    fmt.Println(v2, fi61)
                                                    fmt.Println(v1, fi61)
                                                }
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = v2
                                            }
                                            if false {
                                                if true {
                                                    switch (("" + v2) + "foo bar") {
                                                    case "hello":
                                                        var v67 int32 = 2147483647
                                                        _ = v67
                                                    case "":
                                                        v3 = append(v3, "\n\t")
                                                    case "foo bar":
                                                        v68 := GenS2{F0: 2}
                                                        _ = v68
                                                    default:
                                                        if true {
                                                            if true {
                                                                v1 -= int8(v63)
                                                            }
                                                            switch v2 {
                                                            case "":
                                                                switch v2 {
                                                                case "foo bar":
                                                                    var v69 uint64 = 18446744073709551615
                                                                    var v70 uint32 = 2
                                                                    _ = v69
                                                                    _ = v70
                                                                case "":
                                                                    v3 = append(v3, "foo bar")
                                                                }
                                                            case "foo bar":
                                                                switch (-(fi61)) {
                                                                case 9223372036854775807:
                                                                    if (v1 == (v6 % ((v6) | 1))) {
                                                                        fmt.Println(fi62)
                                                                        var v71 float32 = -2.25
                                                                        var v72 uint64 = 2
                                                                        _ = v71
                                                                        _ = v72
                                                                    } else {
                                                                        if false {
                                                                            v3 = append(v3, "")
                                                                        }
                                                                        if true {
                                                                            switch v2 {
                                                                            case "foo bar":
                                                                                switch (fi62 << (((^(v21))) & 63)) {
                                                                                case -9223372036854775808:
                                                                                    if (genFn0(0, v63) == 2147483647) {
                                                                                        v73 := GenS4{F0: "a", F1: 869, F2: (v2 + ((v2 + (((("\n\t" + (((v2 + ("foo bar" + (v2 + (("a" + v2) + "a")))) + v2) + "foo bar")) + v2) + v2) + (v2 + (((v2 + v2) + (v2 + "foo bar")) + v2)))) + (v2 + "foo bar")))}
                                                                                        v74 := GenS6{F0: 0, F1: (^(fi62)), F2: (fi62 * fi62)}
                                                                                        v6 = v1
                                                                                        _ = v73
                                                                                        _ = v74
                                                                                    } else {
                                                                                        if (!(v23 >= (v23 | v23))) {
                                                                                            var v75 float64 = ((1e10 - (((1.5 / 1e10) * (float64(1.5) - (-2.25 / 1.0))) - 1e10)) * (0.0 + 1.5))
                                                                                            _ = v75
                                                                                        }
                                                                                        var v76 string = v2
                                                                                        _ = v76
                                                                                    }
                                                                                default:
                                                                                    _ = len(v3)
                                                                                }
                                                                            default:
                                                                                v77 := GenS4{F0: "\n\t", F1: 4294967295, F2: "foo bar"}
                                                                                _ = v77
                                                                            }
                                                                            v21 = 69
                                                                            var v78 uint8 = 0
                                                                            _ = v78
                                                                        }
                                                                    }
                                                                default:
                                                                    if true {
                                                                        var v79 int8 = 127
                                                                        if false {
                                                                            v23 -= (v23 % (((v23 ^ v58)) | 1))
                                                                            if ((v23 + (v23 % (((v58 / ((v23) | 1))) | 1))) == v23) {
                                                                                v6 += (v1 % (((^(v6))) | 1))
                                                                                v80 := GenS7{F0: (v23 << (((v21 & (^(v21)))) & 63))}
                                                                                v6 = v79
                                                                                var v81 float64 = float64(float32(((1e10 / (1.5 - -2.25)) - float64(1.5))))
                                                                                _ = v80
                                                                                _ = v81
                                                                            } else {
                                                                                if true {
                                                                                    var v82 float32 = float32(0.0)
                                                                                    _ = v82
                                                                                } else {
                                                                                    v21 = (v21 - v21)
                                                                                    fmt.Println(v63, v6)
                                                                                }
                                                                                v83 := GenS8{F0: 1, F1: uint64(v21), F2: (v21 % ((genFn1((v2 + ((("foo bar" + "foo bar") + v2) + v2)), 18446744073709551615, (2 <= 42))) | 1))}
                                                                                if (!false) {
                                                                                    var v84 int64 = (v23 | (v58 >> ((v21) & 63)))
                                                                                    var v85 float64 = -2.25
                                                                                    _ = v84
                                                                                    _ = v85
                                                                                }
                                                                                _ = v83
                                                                            }
                                                                        } else {
                                                                            v23 -= (v23 ^ (^(v23)))
                                                                            v1 = -79
                                                                            var v86 string = (v2 + "hello")
                                                                            v23 = v58
                                                                            var v87 int = (fi62 << ((v21) & 63))
                                                                            _ = v86
                                                                            _ = v87
                                                                        }
                                                                        v2 = (((v2 + "") + v2) + v2)
                                                                        _ = v79
                                                                    }
                                                                }
                                                            case "a":
                                                                var v88 int = fi61
                                                                _ = v88
                                                            }
                                                            v5[0] = fi62
                                                            v89 := GenS9{F0: fi62, F1: (v2 + "")}
                                                            _ = v89
                                                        } else {
                                                            v4[(fi62 << ((0) & 63))] = true
                                                        }
                                                    }
                                                } else {
                                                    if (!(!true)) {
                                                        v90 := GenS4{F0: "a", F1: 18446744073709551615, F2: v2}
                                                        switch (^(fi62)) {
                                                        case 1:
                                                            delete(v4, (-(fi62)))
                                                        }
                                                        if false {
                                                            if true {
                                                                _ = len(v3)
                                                                var v91 uint = genFn1("\n\t", 4294967295, false)
                                                                _ = v91
                                                            } else {
                                                                v23 = v23
                                                                v92 := GenS2{F0: 0}
                                                                _ = v92
                                                            }
                                                            var v93 uint64 = 0
                                                            _ = v93
                                                        }
                                                        var v94 uint16 = 65535
                                                        v95 := GenS1{F0: v23}
                                                        if ((^(v63)) <= v63) {
                                                            if true {
                                                                fmt.Println(v58)
                                                            }
                                                            delete(v5, 42)
                                                            var v96 bool = ((!(-2147483648 != -2147483648)) && (42 <= 485))
                                                            fmt.Println(v1)
                                                            var v97 uint64 = 2
                                                            if false {
                                                                switch (("a" + v2) + v2) {
                                                                case "\n\t":
                                                                    var v98 uint16 = v94
                                                                    _ = v98
                                                                default:
                                                                    v99 := GenS1{F0: v23}
                                                                    _ = v99
                                                                }
                                                                v94 = (v94 - v94)
                                                                v100 := GenS10{F0: 4294967295}
                                                                _ = v100
                                                            }
                                                            _ = v96
                                                            _ = v97
                                                        } else {
                                                            v4[fi62] = true
                                                        }
                                                        v21 += 18446744073709551615
                                                        _ = v90
                                                        _ = v94
                                                        _ = v95
                                                    }
                                                }
                                                var v101 uint16 = 65535
                                                v58 = 9223372036854775807
                                                if false {
                                                    switch fi61 {
                                                    case 128:
                                                        if ((v23 << ((v21) & 63)) != 0) {
                                                            var v102 int32 = -1
                                                            _ = v102
                                                        }
                                                    case 127:
                                                        if (795 <= -1) {
                                                            var v103 uint8 = 42
                                                            switch (fi62 % ((42) | 1)) {
                                                            case -9223372036854775808:
                                                                v104 := GenS9{F0: 128, F1: ("foo bar" + ("foo bar" + ("a" + "")))}
                                                                _ = v104
                                                            default:
                                                                v23 += v58
                                                                if (!false) {
                                                                    var v105 float32 = (((1.5 - 0.1) * (0.0 * 0.0)) + 1.5)
                                                                    switch (-(fi62)) {
                                                                    case -1:
                                                                        if false {
                                                                            v2 = "foo bar"
                                                                        }
                                                                    case 965:
                                                                        fmt.Println(fi61, v101, v58)
                                                                    }
                                                                    _ = v105
                                                                } else {
                                                                    if ("a" < ("foo bar" + "hello")) {
                                                                        if len(v3) > 0 {
                                                                            v3[0] = "foo bar"
                                                                        }
                                                                        v2 = (v2 + (v2 + (v2 + (v2 + v2))))
                                                                    }
                                                                    _ = len(v5)
                                                                    v106 := GenS11{F0: ((v101 != (v101 / ((225) | 1))) && false)}
                                                                    _ = v106
                                                                }
                                                                v107 := GenS12{F0: genFn0(4294967295, v63), F1: 0, F2: -798}
                                                                var v108 uint16 = (v101 ^ (v101 | v101))
                                                                var v109 int16 = 32767
                                                                _ = v107
                                                                _ = v108
                                                                _ = v109
                                                            }
                                                            _ = v103
                                                        } else {
                                                            if false {
                                                                fmt.Println(v21, v2)
                                                                v1 -= int8(fi62)
                                                                v110 := GenS5{F0: ((v2 + v2) + v2), F1: (v23 & (-(v58)))}
                                                                if ((v63 >= (v63 << (((^(v21))) & 63))) || (((true && false) || (!false)) && (v101 >= (v101 & (^(v101)))))) {
                                                                    var v111 bool = ((v1 == (-(v6))) || (!true))
                                                                    _ = v111
                                                                }
                                                                _ = v110
                                                            }
                                                            if false {
                                                                v2 = v2
                                                                _ = len(v4)
                                                            }
                                                            switch (((v2 + v2) + v2) + ("foo bar" + ((v2 + v2) + "a"))) {
                                                            case "hello":
                                                                var v112 uint64 = 948
                                                                _ = v112
                                                            case "":
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                            case "\n\t":
                                                                v113 := GenS3{F0: v63, F1: -2147483648}
                                                                _ = v113
                                                            }
                                                        }
                                                    case 762:
                                                        switch fi61 {
                                                        case 722:
                                                            if len(v3) > 0 {
                                                                v3[0] = v2
                                                            }
                                                        case 1:
                                                            var v114 float64 = -2.25
                                                            _ = v114
                                                        }
                                                    default:
                                                        v115 := GenS6{F0: 2, F1: (fi61 ^ (fi62 / (((fi62 & int(v21))) | 1))), F2: fi62}
                                                        _ = v115
                                                    }
                                                    var v116 uint32 = 2147483647
                                                    if false {
                                                        v116 = 2147483647
                                                    } else {
                                                        switch -431 {
                                                        case 32768:
                                                            var v117 int8 = v6
                                                            _ = v117
                                                        case 32767:
                                                            if true {
                                                                switch "" {
                                                                case "\n\t":
                                                                    switch -2147483648 {
                                                                    case -2147483648:
                                                                        var v118 uint32 = (v116 & v116)
                                                                        _ = v118
                                                                    default:
                                                                        var v119 int = fi61
                                                                        _ = v119
                                                                    }
                                                                case "":
                                                                    v116 = v116
                                                                }
                                                            } else {
                                                                if (!true) {
                                                                    v58 += v58
                                                                    var v120 uint = 2
                                                                    var v121 bool = false
                                                                    _ = v120
                                                                    _ = v121
                                                                } else {
                                                                    v122 := GenS13{F0: 255}
                                                                    v123 := GenS4{F0: v2, F1: 42, F2: (v2 + v2)}
                                                                    v124 := GenS14{F0: v21}
                                                                    _ = v122
                                                                    _ = v123
                                                                    _ = v124
                                                                }
                                                            }
                                                        }
                                                    }
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    if true {
                                                        v2 = "a"
                                                        var v125 uint8 = uint8(v116)
                                                        delete(v4, fi62)
                                                        _ = v125
                                                    } else {
                                                        v63 -= 566
                                                    }
                                                    _ = v116
                                                }
                                                if true {
                                                    if ((("foo bar" + v2) + v2) < v2) {
                                                        v126 := GenS11{F0: ((-(fi62)) >= (fi61 | (fi61 & -230)))}
                                                        v5[803] = 4294967295
                                                        var v127 float64 = 1.5
                                                        fmt.Println(v6, v6)
                                                        v128 := GenS1{F0: v23}
                                                        v129 := GenS15{F0: (!false), F1: uint64(v21)}
                                                        _ = v126
                                                        _ = v127
                                                        _ = v128
                                                        _ = v129
                                                    } else {
                                                        v130 := GenS6{F0: 793, F1: (fi62 + fi61), F2: (fi62 / (((fi61 % ((fi62) | 1))) | 1))}
                                                        v58 += -779
                                                        var v131 uint8 = 127
                                                        _ = v130
                                                        _ = v131
                                                    }
                                                } else {
                                                    var v132 uint16 = (v101 ^ v101)
                                                    v23 = (v58 % ((v58) | 1))
                                                    _ = v132
                                                }
                                                _ = v101
                                            }
                                        } else {
                                            v133 := GenS3{F0: v63, F1: 0}
                                            var v134 float32 = (float32(-2.25) + 0.1)
                                            v135 := GenS8{F0: 33, F1: 753, F2: 18446744073709551615}
                                            _ = v133
                                            _ = v134
                                            _ = v135
                                        }
                                        v58 = (v58 >> (((^(v21))) & 63))
                                        _ = len(v5)
                                        fmt.Println(v23, v58)
                                        _ = v63
                                    }
                                }
                            default:
                                v58 = (v23 % ((-149) | 1))
                            }
                            switch (("a" + "a") + (((v2 + v2) + "foo bar") + ("hello" + v2))) {
                            case "foo bar":
                                v1 = -56
                            case "hello":
                                if (42 != genFn0(18446744073709551615, 32767)) {
                                    v58 += (v58 * (v23 | (v23 % ((v23) | 1))))
                                    var v136 int8 = int8(v23)
                                    _ = v136
                                }
                            }
                            switch 4294967295 {
                            case -967:
                                if false {
                                    switch (fi61 << ((genFn1(v2, 1, true)) & 63)) {
                                    case 128:
                                        v23 = 9223372036854775807
                                    case 998:
                                        var v137 int8 = 0
                                        v138 := GenS5{F0: "hello", F1: -9223372036854775808}
                                        _ = v137
                                        _ = v138
                                    default:
                                        v3 = append(v3, ("foo bar" + ((v2 + v2) + v2)))
                                    }
                                    var v139 uint32 = 275
                                    _ = v139
                                }
                            default:
                                v1 -= v6
                                switch v2 {
                                case "\n\t":
                                    v140 := GenS16{F0: ((((v2 + v2) + v2) + ("hello" + (v2 + v2))) != v2), F1: 32767}
                                    var v141 int8 = 36
                                    var v142 uint8 = 255
                                    _ = v140
                                    _ = v141
                                    _ = v142
                                case "hello":
                                    var v143 float32 = 3.14159
                                    _ = v143
                                default:
                                    v144 := GenS7{F0: 9223372036854775807}
                                    _ = v144
                                }
                            }
                            v6 += v6
                            v145 := GenS12{F0: 42, F1: 173, F2: 128}
                            if false {
                                var v146 int = (fi62 + 0)
                                v58 -= (v58 >> ((genFn1("foo bar", 1, false)) & 63))
                                var v147 uint64 = 0
                                fmt.Println(fi62, fi61)
                                _ = v146
                                _ = v147
                            } else {
                                if true {
                                    v148 := GenS12{F0: 1, F1: 546, F2: fi62}
                                    _ = v148
                                }
                                switch v2 {
                                case "a":
                                    v23 += (v23 + -9223372036854775808)
                                case "hello":
                                    var v149 bool = (127 <= 0)
                                    _ = v149
                                }
                                var v150 int16 = 62
                                _ = v150
                            }
                            _ = fi62
                            _ = v145
                        }
                    }
                    v151 := GenS16{F0: (300 > 2), F1: fi61}
                    for fi152 := 0; fi152 < 8; fi152++ {
                        v1 -= v6
                        if true {
                            delete(v4, fi61)
                            switch fi152 {
                            case 4294967295:
                                v153 := GenS17{F0: -1, F1: 42}
                                _ = v153
                            case 32767:
                                delete(v5, 2147483648)
                                var v154 int16 = -1
                                var v155 int8 = (^(v1))
                                _ = v154
                                _ = v155
                            case 42:
                                v156 := GenS2{F0: 65535}
                                _ = v156
                            default:
                                v157 := GenS14{F0: v21}
                                _ = v157
                            }
                        }
                        _ = fi152
                    }
                    _ = fi61
                    _ = v151
                }
            }
            _ = v57
            _ = v58
        } else {
            var v158 uint8 = 2
            for fi159 := 0; fi159 < 10; fi159++ {
                for fi160 := 0; fi160 < 9; fi160++ {
                    v5[65535] = (-(fi160))
                    v161 := GenS13{F0: 127}
                    _ = fi160
                    _ = v161
                }
                fmt.Println(v2)
                _ = fi159
            }
            var v162 float32 = -2.25
            _ = v158
            _ = v162
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
