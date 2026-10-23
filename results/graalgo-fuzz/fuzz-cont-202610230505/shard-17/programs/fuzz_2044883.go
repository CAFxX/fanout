package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 string; F2 string }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS7 struct { F0 int; F1 string; F2 uint16 }
type GenS8 struct { F0 uint16; F1 uint64; F2 int8 }
type GenS9 struct { F0 int16; F1 int32 }
type GenS10 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS11 struct { F0 uint; F1 uint }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 int32; F1 int32 }
type GenS15 struct { F0 uint32; F1 int64 }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 uint; F1 uint64; F2 int16 }
type GenS18 struct { F0 string }
func genFn0() string {
    v6 := GenS1{F0: 0}
    var v7 int8 = -128
    v8 := GenS1{F0: -9223372036854775808}
    if true {
        v9 := GenS2{F0: 1, F1: ("foo bar" + (("" + "\n\t") + "foo bar")), F2: "foo bar"}
        var v10 int8 = (v7 + v7)
        v10 += (v10 | v10)
        _ = v10
        _ = v9
    }
    v11 := []int{-2147483648, 2147483647, 917, 127}
    _ = v11
    _ = v6
    _ = v7
    _ = v8
    return "hello"
}
func genFn1(p0 int, p1 int8, p2 uint16) uint32 {
    v12 := []int{(p0 % ((p0) | 1)), p0, 9223372036854775807, p0}
    var v13 uint64 = 0
    for fi14 := 0; fi14 < 3; fi14++ {
        v15 := GenS3{F0: 1, F1: 127}
        fmt.Println(p0)
        if (!false) {
            if false {
                var v16 bool = (!true)
                for fi17 := 0; fi17 < 14; fi17++ {
                    if v16 {
                        if v16 {
                            v18 := GenS4{F0: 1}
                            switch "\n\t" {
                            case "\n\t":
                                var v19 string = (("a" + "\n\t") + ("foo bar" + "a"))
                                var v20 uint16 = 32767
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                                _ = v19
                                _ = v20
                            case "foo bar":
                                switch "foo bar" {
                                case "a":
                                    v21 := GenS5{F0: (^(p2))}
                                    _ = v21
                                case "hello":
                                    fmt.Println(v16, fi14)
                                case "foo bar":
                                    if (!((fi14 - (fi17 ^ 4294967295)) < (fi14 << ((454) & 63)))) {
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                        switch ("" + "hello") {
                                        case "":
                                            v22 := GenS2{F0: 506, F1: (((("hello" + "\n\t") + ("a" + "a")) + "hello") + ("" + ("foo bar" + "a"))), F2: ("\n\t" + "")}
                                            _ = v22
                                        case "hello":
                                            switch "foo bar" {
                                            case "hello":
                                                var v23 float32 = 1.5
                                                _ = v23
                                            case "foo bar":
                                                v12 = append(v12, (^(fi14)))
                                            case "\n\t":
                                                if v16 {
                                                    switch fi14 {
                                                    case 9223372036854775807:
                                                        if false {
                                                            var v24 float64 = 1e10
                                                            var v25 uint = 4294967295
                                                            _ = v24
                                                            _ = v25
                                                        }
                                                    case -2147483648:
                                                        if v16 {
                                                            v26 := GenS2{F0: (p2 >> ((42) & 63)), F1: "foo bar", F2: "hello"}
                                                            _ = v26
                                                        } else {
                                                            if v16 {
                                                                fmt.Println(fi14, v16, p1)
                                                            } else {
                                                                v27 := GenS5{F0: p2}
                                                                v16 = (0 > 127)
                                                                v28 := GenS5{F0: p2}
                                                                var v29 int16 = -1
                                                                _ = v27
                                                                _ = v28
                                                                _ = v29
                                                            }
                                                        }
                                                    case 4294967295:
                                                        v30 := GenS2{F0: (p2 - (p2 + 65535)), F1: "", F2: "foo bar"}
                                                        _ = v30
                                                    }
                                                } else {
                                                    v13 -= (v13 << ((850) & 63))
                                                    v31 := GenS1{F0: p0}
                                                    fmt.Println(v13)
                                                    if (42 != 20) {
                                                        if true {
                                                            p1 += p1
                                                        }
                                                    }
                                                    p2 = p2
                                                    _ = v31
                                                }
                                            default:
                                                if false {
                                                    v32 := GenS6{F0: v13, F1: 4294967295, F2: (p2 | (p2 ^ 0))}
                                                    var v33 uint32 = 961
                                                    v13 += (v13 >> ((924) & 63))
                                                    _ = v32
                                                    _ = v33
                                                }
                                            }
                                            v12 = append(v12, p0)
                                            var v34 string = ("" + "hello")
                                            _ = v34
                                        case "foo bar":
                                            v35 := GenS7{F0: (-(p0)), F1: "\n\t", F2: (p2 >> ((207) & 63))}
                                            _ = v35
                                        }
                                        var v36 uint32 = 455
                                        _ = v36
                                    } else {
                                        var v37 float32 = 0.0
                                        _ = v37
                                    }
                                default:
                                    var v38 uint8 = 255
                                    _ = v38
                                }
                            }
                            _ = len(v12)
                            _ = v18
                        } else {
                            var v39 int = fi14
                            var v40 uint64 = 2
                            switch 32768 {
                            case 514:
                                var v41 int64 = -495
                                _ = v41
                            case -2147483648:
                                p0 += (p0 + p0)
                            }
                            _ = v39
                            _ = v40
                        }
                    }
                    v42 := GenS6{F0: (v13 - v13), F1: 0, F2: (p2 - p2)}
                    if (!false) {
                        v16 = true
                        if v16 {
                            var v43 uint8 = 42
                            p1 = (-(p1))
                            _ = v43
                        }
                        fmt.Println(v13)
                    }
                    v44 := GenS8{F0: p2, F1: v13, F2: p1}
                    _ = fi17
                    _ = v42
                    _ = v44
                }
                var v45 int16 = -1
                _ = v16
                _ = v45
            }
        } else {
            var v46 float64 = float64(float32(-2.25))
            _ = v46
        }
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return uint32(p2)
}

func main() {
    v1 := []uint32{}
    var v2 uint64 = 2
    var v3 float64 = 3.14159
    var v4 float32 = (-2.25 - ((1e10 + (float32(float64((-2.25 * 0.0))) + 1e10)) * ((3.14159 / -2.25) / 3.14159)))
    v5 := make(map[uint32]string)
    for fi47 := 0; fi47 < 18; fi47++ {
        switch fi47 {
        case -2147483648:
            switch fi47 {
            case 4294967295:
                var v48 uint64 = (v2 >> ((722) & 63))
                _ = v48
            case -9223372036854775808:
                var v49 uint32 = 435
                _ = v49
            case 934:
                if true {
                    var v50 int32 = 0
                    switch (^(fi47)) {
                    case -9223372036854775808:
                        v50 -= 930
                    }
                    var v51 float32 = v4
                    _ = v50
                    _ = v51
                }
            }
        case -46:
            fmt.Println(v2, v4)
        case 9223372036854775807:
            v52 := GenS6{F0: v2, F1: 4294967295, F2: 42}
            _ = v52
        default:
            if (-2.25 >= 1.5) {
                v53 := v5[2]
                var v54 int64 = 9223372036854775807
                var v55 uint32 = 512
                _ = v53
                _ = v54
                _ = v55
            } else {
                var v56 int = 1
                if (uint(v2) != 1) {
                    switch (v56 * fi47) {
                    case -2147483648:
                        var v57 uint16 = 2
                        _ = v57
                    case -9223372036854775808:
                        var v58 int64 = 0
                        _ = v58
                    case -497:
                        v59 := GenS1{F0: 127}
                        _ = v59
                    }
                    for fi60 := 0; fi60 < 13; fi60++ {
                        if ("" != ((("\n\t" + "foo bar") + "") + "foo bar")) {
                            var v61 int8 = 100
                            fmt.Println(fi47)
                            _ = v61
                        }
                        v56 -= fi47
                        if (true || true) {
                            v62 := GenS8{F0: 658, F1: (v2 - v2), F2: 100}
                            var v63 string = "foo bar"
                            _ = v62
                            _ = v63
                        } else {
                            if (0 <= 1) {
                                var v64 bool = false
                                if v64 {
                                    var v65 float32 = (0.1 / v4)
                                    v1 = append(v1, 1)
                                    _ = v65
                                }
                                switch ("foo bar" + ("foo bar" + "")) {
                                case "a":
                                    var v66 uint = 739
                                    _ = v66
                                case "foo bar":
                                    v2 += 4294967295
                                default:
                                    var v67 int64 = 42
                                    _ = v67
                                }
                                _ = v64
                            } else {
                                var v68 uint16 = 42
                                var v69 int32 = 2147483647
                                var v70 uint16 = (v68 << ((475) & 63))
                                if true {
                                    switch (("a" + "\n\t") + "hello") {
                                    case "a":
                                        v3 = float64((v4 / v4))
                                    case "\n\t":
                                        v71 := GenS6{F0: (v2 + 0), F1: 561, F2: uint16(v2)}
                                        _ = v71
                                    case "":
                                        fmt.Println(v3, v3, v70)
                                    }
                                    if true {
                                        fmt.Println(v4, v68)
                                        v4 = 3.14159
                                        v72 := GenS6{F0: v2, F1: 4294967295, F2: v68}
                                        _ = v72
                                    }
                                    if false {
                                        v73 := GenS4{F0: 127}
                                        if (!(127 >= 2)) {
                                            var v74 int32 = v69
                                            v56 = (^(v56))
                                            _ = v74
                                        } else {
                                            _ = len(v5)
                                            switch fi47 {
                                            case 225:
                                                var v75 int64 = -592
                                                _ = v75
                                            default:
                                                var v76 uint8 = 124
                                                _ = v76
                                            }
                                            v77 := GenS9{F0: -32768, F1: (v69 ^ v69)}
                                            _ = v77
                                        }
                                        _ = v73
                                    } else {
                                        v5[391] = "a"
                                        v78 := v5[65535]
                                        v3 -= 1e10
                                        if (842 >= 0) {
                                            v79 := GenS10{F0: (v2 + (v2 ^ v2)), F1: v2, F2: 67}
                                            _ = v79
                                        }
                                        _ = v78
                                    }
                                } else {
                                    switch (("" + "hello") + "hello") {
                                    case "foo bar":
                                        fmt.Println(v3, v2)
                                    case "a":
                                        var v80 float64 = 3.14159
                                        _ = v80
                                    case "\n\t":
                                        switch (-(v56)) {
                                        case 42:
                                            var v81 uint16 = v70
                                            _ = v81
                                        default:
                                            var v82 int32 = 836
                                            _ = v82
                                        }
                                        if ((0.1 >= (float64(((3.14159 / (0.1 - ((1.5 + 0.1) - (0.0 * float32(3.14159))))) * 3.14159)) * (v3 / v3))) || true) {
                                            if false {
                                                v83 := GenS11{F0: 42, F1: 0}
                                                _ = v83
                                            } else {
                                                v84 := GenS12{F0: 42}
                                                if ((("\n\t" <= "") && (-636 <= 556)) || false) {
                                                    fmt.Println(v3)
                                                    var v85 uint64 = v2
                                                    _ = v85
                                                } else {
                                                    var v86 float64 = 0.0
                                                    if (!(!(-128 == -46))) {
                                                        var v87 uint32 = 513
                                                        v88 := GenS1{F0: 32767}
                                                        _ = v87
                                                        _ = v88
                                                    }
                                                    if true {
                                                        if len(v1) > 0 {
                                                            v1[0] = 983
                                                        }
                                                        v89 := v5[genFn1((fi60 >> ((42) & 63)), 84, 255)]
                                                        var v90 uint16 = v68
                                                        var v91 int8 = 64
                                                        var v92 float64 = ((v3 - v86) - ((v3 - 3.14159) - 0.1))
                                                        v5[891] = v89
                                                        var v93 float32 = (v4 / v4)
                                                        if (v93 < v93) {
                                                            v91 = 0
                                                        } else {
                                                            _ = len(v1)
                                                            v70 -= (v90 & (v70 % (((^(v70))) | 1)))
                                                        }
                                                        _ = v89
                                                        _ = v90
                                                        _ = v91
                                                        _ = v92
                                                        _ = v93
                                                    } else {
                                                        v94 := GenS13{F0: 0}
                                                        var v95 uint64 = (v2 ^ 647)
                                                        _ = v94
                                                        _ = v95
                                                    }
                                                    _ = v86
                                                }
                                                if false {
                                                    v3 = (3.14159 / v3)
                                                    var v96 int64 = 0
                                                    if (("" < "\n\t") && false) {
                                                        fmt.Println(v2, v69)
                                                        v68 = v70
                                                        if true {
                                                            v70 -= (v70 % ((v68) | 1))
                                                        } else {
                                                            if ((false && true) || (fi47 == (^(fi60)))) {
                                                                v68 -= v68
                                                                fmt.Println(v68, fi60)
                                                            }
                                                            _ = len(v1)
                                                        }
                                                    } else {
                                                        switch v56 {
                                                        case 319:
                                                            var v97 int16 = 32767
                                                            _ = v97
                                                        case -9223372036854775808:
                                                            fmt.Println(v68, fi47, v69)
                                                        case 476:
                                                            _ = len(v5)
                                                        default:
                                                            v98 := GenS14{F0: v69, F1: v69}
                                                            _ = v98
                                                        }
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        fmt.Println(v4)
                                                    }
                                                    _ = v96
                                                }
                                                _ = v84
                                            }
                                            delete(v5, 886)
                                            v99 := GenS15{F0: 369, F1: -9223372036854775808}
                                            _ = v99
                                        } else {
                                            var v100 int16 = 147
                                            fmt.Println(v70, v3)
                                            v101 := GenS16{F0: 77}
                                            _ = v100
                                            _ = v101
                                        }
                                        var v102 int8 = 0
                                        fmt.Println(v68)
                                        var v103 uint32 = 42
                                        _ = v102
                                        _ = v103
                                    }
                                }
                                _ = v68
                                _ = v69
                                _ = v70
                            }
                            v104 := v5[2147483647]
                            _ = v104
                        }
                        _ = fi60
                    }
                    var v105 int = 9223372036854775807
                    _ = v105
                } else {
                    v3 += 3.14159
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v56
            }
        }
        _ = fi47
    }
    _ = len(v5)
    switch (("" + ("foo bar" + "a")) + ("\n\t" + "\n\t")) {
    case "":
        fmt.Println(v3, v4)
    case "a":
        var v106 int32 = 0
        _ = v106
    case "hello":
        if (9223372036854775807 >= 9223372036854775807) {
            var v107 int = 32768
            switch "\n\t" {
            case "a":
                var v108 uint32 = 4294967295
                _ = v108
            case "foo bar":
                fmt.Println(v2)
            case "\n\t":
                v5[42] = "hello"
            default:
                v4 = v4
            }
            _ = v107
        }
    }
    if (210 >= 229) {
        switch 4294967295 {
        case -708:
            var v109 uint16 = 187
            _ = v109
        case -633:
            v110 := GenS5{F0: 42}
            _ = v110
        default:
            v111 := GenS17{F0: 636, F1: (^(v2)), F2: 42}
            _ = v111
        }
        if false {
            v2 = 2
            v112 := GenS10{F0: v2, F1: (v2 + 4294967295), F2: -1}
            _ = v112
        }
        for fi113 := 0; fi113 < 18; fi113++ {
            delete(v5, 1)
            for fi114 := 0; fi114 < 7; fi114++ {
                v2 = v2
                if len(v1) > 0 {
                    v1[0] = 65535
                }
                _ = fi114
            }
            _ = fi113
        }
    }
    switch -9223372036854775808 {
    case -1:
        v115 := GenS18{F0: "hello"}
        _ = v115
    case 32768:
        v2 = v2
    }
    fmt.Println(v3)
    if len(v1) > 0 {
        v1[0] = 344
    }
    for fi116 := 0; fi116 < 18; fi116++ {
        var v117 int = (fi116 << ((672) & 63))
        var v118 int32 = 39
        _ = fi116
        _ = v117
        _ = v118
    }
    if false {
        fmt.Println(v4, v3)
        var v119 uint32 = 1
        v3 = v3
        _ = v119
    }
    _ = len(v5)
    v120 := GenS5{F0: 32767}
    switch 0 {
    case 32767:
        fmt.Println(v2, v4, v3)
    }
    if true {
        var v121 int32 = 0
        v121 -= v121
        _ = v121
    }
    delete(v5, 65535)
    var v122 uint8 = 56
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v122:", v122)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v120
    _ = v122
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
