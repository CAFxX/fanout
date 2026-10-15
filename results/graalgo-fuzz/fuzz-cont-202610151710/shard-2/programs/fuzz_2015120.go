package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 uint; F2 uint32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 bool; F1 int16 }
type GenS7 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS8 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 string; F1 int8; F2 int32 }
type GenS11 struct { F0 uint; F1 uint64; F2 int }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int16; F1 int64 }
type GenS14 struct { F0 int32; F1 int }
type GenS15 struct { F0 int8; F1 uint16; F2 bool }
type GenS16 struct { F0 int; F1 int; F2 uint16 }
type GenS17 struct { F0 int64; F1 uint64; F2 int64 }
type GenS18 struct { F0 uint16; F1 int64 }
type GenS19 struct { F0 int; F1 uint64; F2 uint }
func genFn0(p0 uint8, p1 int32, p2 uint64) int8 {
    p0 += (p0 << ((4294967295) & 63))
    var v4 uint16 = 42
    fmt.Println(p0, p0, p1)
    for fi5 := 0; fi5 < 15; fi5++ {
        for fi6 := 0; fi6 < 15; fi6++ {
            switch fi5 {
            case 42:
                if (true || (943 == (v4 << ((732) & 63)))) {
                    if true {
                        p2 -= 1
                        v4 = v4
                    } else {
                        var v7 int = fi5
                        p0 = p0
                        var v8 uint32 = 2
                        _ = v7
                        _ = v8
                    }
                    var v9 int8 = 0
                    _ = v9
                } else {
                    fmt.Println(p0)
                    var v10 uint64 = (p2 * p2)
                    var v11 float32 = ((0.1 + (3.14159 + 0.1)) - -2.25)
                    v10 = 586
                    p2 += v10
                    v4 = 2
                    v12 := GenS1{F0: uint32(v10), F1: v4}
                    fmt.Println(v11)
                    p0 += (p0 - (^(p0)))
                    v11 = v11
                    _ = v10
                    _ = v11
                    _ = v12
                }
            case 4294967295:
                if (!(42 == -912)) {
                    v13 := []int{(fi6 * (fi6 % (((^(fi5))) | 1))), (fi5 / (((fi5 << ((42) & 63))) | 1)), 42}
                    if (42 >= 100) {
                        if (false || true) {
                            v14 := GenS2{F0: (fi5 & (fi6 / (((fi6 ^ (^(fi6)))) | 1)))}
                            _ = v14
                        } else {
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                            if (2147483648 >= 17) {
                                p0 = (^(p0))
                                v13 = append(v13, 1)
                            } else {
                                if ((1e10 + (3.14159 - 1e10)) == float64(1e10)) {
                                    _ = len(v13)
                                    if len(v13) > 0 {
                                        v13[0] = (-(fi5))
                                    }
                                }
                                var v15 int = 4294967295
                                p0 = p0
                                _ = v15
                            }
                        }
                        var v16 int16 = -32768
                        fmt.Println(p0, p0)
                        _ = v16
                    }
                    switch ("\n\t" + "\n\t") {
                    case "foo bar":
                        p1 = (p1 * p1)
                    case "\n\t":
                        v4 = (^(v4))
                    }
                    _ = v13
                }
            default:
                if (false || false) {
                    var v17 uint8 = 0
                    var v18 int8 = 0
                    if (2 >= (^(p2))) {
                        p1 = p1
                    }
                    _ = v17
                    _ = v18
                }
            }
            var v19 float64 = -2.25
            _ = fi6
            _ = v19
        }
        var v20 uint64 = (p2 ^ 0)
        _ = fi5
        _ = v20
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        fmt.Println(fi21, p2)
        _ = fi21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 127
}
func genFn1(p0 int64, p1 uint, p2 uint64) string {
    v22 := GenS3{F0: -418}
    var v23 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v23
    return ("a" + ("foo bar" + ("\n\t" + "foo bar")))
}

func main() {
    var v1 int16 = -1
    v2 := make(map[uint32]int)
    var v3 uint32 = 4294967295
    if false {
        for fi24 := 0; fi24 < 8; fi24++ {
            v2[1] = fi24
            _ = fi24
        }
    }
    v25 := []int{-2147483648}
    var v26 int64 = 42
    switch -9223372036854775808 {
    case 1:
        var v27 int32 = 969
        _ = v27
    case -2147483648:
        v1 = int16(v3)
    case 4294967295:
        v2[v3] = -9223372036854775808
    }
    var v28 int = 32768
    for fi29 := 0; fi29 < 7; fi29++ {
        if false {
            if false {
                switch -9223372036854775808 {
                case 127:
                    var v30 string = "foo bar"
                    _ = v30
                default:
                    for fi31 := 0; fi31 < 5; fi31++ {
                        v28 = (^(v28))
                        _ = fi31
                    }
                }
                if len(v25) > 0 {
                    v25[0] = v28
                }
            }
        } else {
            for fi32 := 0; fi32 < 8; fi32++ {
                v33 := GenS4{F0: v1, F1: 443, F2: (v3 >> ((0) & 63))}
                var v34 uint64 = 986
                if ((1e10 + -2.25) <= 1e10) {
                    if ((false && ((0.0 * -2.25) < -2.25)) && true) {
                        v2[v3] = (-(v28))
                        v35 := GenS5{F0: 42}
                        switch (-(fi29)) {
                        case 4294967295:
                            if ((65535 >= 42) && (2147483647 < 1)) {
                                v3 += (^(v3))
                            } else {
                                v28 -= fi29
                                v2[v3] = fi29
                                v28 = (^(fi29))
                            }
                        case 32768:
                            var v36 uint = 2
                            _ = v36
                        case 0:
                            var v37 uint16 = 65535
                            _ = v37
                        }
                        v38 := GenS6{F0: true, F1: 0}
                        var v39 int8 = -128
                        _ = v35
                        _ = v38
                        _ = v39
                    } else {
                        var v40 bool = (true && (4294967295 == (^(v3))))
                        var v41 int8 = -1
                        fmt.Println(fi32, v1, v1)
                        _ = v40
                        _ = v41
                    }
                    v42 := GenS7{F0: 349, F1: (true && ("" != "foo bar")), F2: 0}
                    _ = v42
                }
                v25 = append(v25, fi29)
                var v43 int64 = (v26 << ((0) & 63))
                _ = fi32
                _ = v33
                _ = v34
                _ = v43
            }
            v28 += (v28 + (^(v28)))
        }
        _ = fi29
    }
    switch (v28 % (((v28 & (-(v28)))) | 1)) {
    case -1:
        for fi44 := 0; fi44 < 3; fi44++ {
            v45 := v2[2]
            if false {
                v46 := GenS8{F0: 2, F1: 42, F2: 218}
                _ = v46
            } else {
                var v47 float64 = 1e10
                for fi48 := 0; fi48 < 18; fi48++ {
                    fmt.Println(v28)
                    v49 := GenS6{F0: ((v3 + v3) < v3), F1: 0}
                    _ = fi48
                    _ = v49
                }
                if (v1 >= (v1 / (((v1 + (v1 - v1))) | 1))) {
                    var v50 uint16 = 52
                    _ = v50
                }
                _ = v47
            }
            if (true || true) {
                for fi51 := 0; fi51 < 2; fi51++ {
                    v52 := GenS9{F0: 841}
                    if len(v25) > 0 {
                        v25[0] = v28
                    }
                    if len(v25) >= 2 {
                        v25 = v25[1:len(v25)]
                    }
                    _ = fi51
                    _ = v52
                }
                fmt.Println(v26, fi44, v26)
            }
            _ = fi44
            _ = v45
        }
    case 0:
        var v53 uint16 = 255
        _ = v53
    }
    v28 = (v28 >> ((18446744073709551615) & 63))
    for fi54 := 0; fi54 < 5; fi54++ {
        for fi55 := 0; fi55 < 12; fi55++ {
            if len(v25) >= 2 {
                v25 = v25[1:len(v25)]
            }
            var v56 int = 32767
            var v57 uint = 2
            _ = fi55
            _ = v56
            _ = v57
        }
        for fi58 := 0; fi58 < 12; fi58++ {
            fmt.Println(v28, fi54, v26)
            _ = fi58
        }
        _ = fi54
    }
    v59 := GenS2{F0: v28}
    for fi60 := 0; fi60 < 2; fi60++ {
        switch (("\n\t" + ("" + ("foo bar" + ((("foo bar" + "\n\t") + "hello") + "foo bar")))) + "hello") {
        case "a":
            v61 := GenS6{F0: (64 >= 0), F1: v1}
            _ = v61
        case "hello":
            if (v1 <= (v1 % ((v1) | 1))) {
                var v62 bool = false
                v1 = v1
                if v62 {
                    fmt.Println(v1)
                    if (true || false) {
                        v62 = (!v62)
                        var v63 uint = 42
                        var v64 int16 = v1
                        _ = v63
                        _ = v64
                    }
                }
                _ = v62
            } else {
                _ = len(v25)
                v3 = (v3 * v3)
                var v65 int16 = (^(v1))
                _ = v65
            }
        case "":
            v26 += v26
        default:
            v66 := GenS8{F0: 128, F1: v26, F2: 128}
            _ = v66
        }
        if ((fi60 % ((v28) | 1)) >= 297) {
            for fi67 := 0; fi67 < 13; fi67++ {
                var v68 uint16 = 2
                v25 = append(v25, fi60)
                _ = fi67
                _ = v68
            }
            if true {
                if false {
                    v69 := GenS8{F0: 99, F1: v26, F2: 0}
                    v70 := GenS10{F0: ("hello" + ("hello" + ("" + "a"))), F1: int8(fi60), F2: 0}
                    _ = v69
                    _ = v70
                } else {
                    v71 := GenS4{F0: v1, F1: 65535, F2: v3}
                    var v72 uint64 = 2
                    var v73 uint = 623
                    _ = v71
                    _ = v72
                    _ = v73
                }
                v74 := GenS11{F0: 42, F1: 244, F2: (v28 | 32767)}
                _ = v74
            }
        } else {
            var v75 uint8 = 255
            var v76 int64 = 9223372036854775807
            if true {
                for fi77 := 0; fi77 < 4; fi77++ {
                    v78 := GenS12{F0: 1}
                    if false {
                        var v79 uint = 1
                        _ = v79
                    }
                    switch fi60 {
                    case -9223372036854775808:
                        if ((!(!(752 > -748))) || true) {
                            if true {
                                v3 = v3
                                switch "" {
                                case "hello":
                                    switch "hello" {
                                    case "\n\t":
                                        if (!((!(false && false)) && (259 >= 4294967295))) {
                                            v2[2147483647] = (v28 + 32767)
                                            if true {
                                                fmt.Println(fi60)
                                            }
                                            fmt.Println(v75, v1)
                                        }
                                        v75 += v75
                                    }
                                default:
                                    fmt.Println(v75, v28)
                                }
                            } else {
                                switch -769 {
                                case -2147483648:
                                    v80 := v2[65535]
                                    if (true || false) {
                                        switch "foo bar" {
                                        case "":
                                            if ("hello" < (("" + ("" + "a")) + ("\n\t" + (("\n\t" + ("foo bar" + (("\n\t" + "\n\t") + "a"))) + ("" + ("hello" + "")))))) {
                                                var v81 int32 = -2147483648
                                                _ = v81
                                            }
                                        case "a":
                                            var v82 int32 = -996
                                            _ = v82
                                        }
                                        var v83 float64 = 1.5
                                        _ = v83
                                    }
                                    _ = v80
                                case 758:
                                    v84 := GenS13{F0: v1, F1: (v76 >> ((255) & 63))}
                                    if (-128 <= 100) {
                                        if false {
                                            if false {
                                                if true {
                                                    if true {
                                                        v1 -= v1
                                                    }
                                                    switch 128 {
                                                    case -1:
                                                        if false {
                                                            v26 -= v76
                                                            fmt.Println(v76, v76)
                                                        } else {
                                                            var v85 string = ""
                                                            v86 := GenS14{F0: -707, F1: (^(fi60))}
                                                            _ = v85
                                                            _ = v86
                                                        }
                                                    case -878:
                                                        var v87 uint = 4294967295
                                                        _ = v87
                                                    case 32767:
                                                        if (4294967295 < v3) {
                                                            var v88 int64 = 447
                                                            var v89 int = fi60
                                                            _ = v88
                                                            _ = v89
                                                        } else {
                                                            if false {
                                                                v90 := GenS6{F0: true, F1: 481}
                                                                fmt.Println(v3, v75, v3)
                                                                _ = v90
                                                            } else {
                                                                v1 -= (^(v1))
                                                                if (v3 == (v3 & (^(v3)))) {
                                                                    if false {
                                                                        v76 += v76
                                                                        if (v75 > (v75 * 127)) {
                                                                            switch "hello" {
                                                                            case "foo bar":
                                                                                v91 := v2[(v3 | uint32(v75))]
                                                                                _ = v91
                                                                            default:
                                                                                var v92 uint16 = uint16(v3)
                                                                                _ = v92
                                                                            }
                                                                        }
                                                                    }
                                                                    var v93 int64 = v76
                                                                    _ = v93
                                                                } else {
                                                                    v2[(v3 * (v3 + 2))] = -679
                                                                    v94 := GenS15{F0: genFn0(v75, 42, 1), F1: 692, F2: ((false && false) && true)}
                                                                    _ = v94
                                                                }
                                                            }
                                                            if true {
                                                                var v95 float64 = (-2.25 / 1.0)
                                                                if false {
                                                                    v1 = v1
                                                                    switch ("\n\t" + (("" + "hello") + "a")) {
                                                                    case "a":
                                                                        v28 = fi60
                                                                        if true {
                                                                            delete(v2, v3)
                                                                            var v96 float32 = (1.5 * (-2.25 / 1.0))
                                                                            _ = v96
                                                                        } else {
                                                                            v97 := v2[(^(v3))]
                                                                            v2[2147483648] = fi77
                                                                            _ = v97
                                                                        }
                                                                    case "\n\t":
                                                                        if (uint16(v75) < 570) {
                                                                            var v98 uint = 18446744073709551615
                                                                            if (true || false) {
                                                                                _ = len(v25)
                                                                                var v99 int16 = v1
                                                                                switch (^(fi77)) {
                                                                                case 127:
                                                                                    var v100 uint64 = 42
                                                                                    _ = v100
                                                                                case 1:
                                                                                    v95 = v95
                                                                                case 42:
                                                                                    v1 += 32767
                                                                                    v99 += v99
                                                                                    if false {
                                                                                        var v101 uint64 = 1
                                                                                        _ = v101
                                                                                    } else {
                                                                                        var v102 uint64 = 786
                                                                                        _ = v102
                                                                                    }
                                                                                    switch "\n\t" {
                                                                                    case "":
                                                                                        if (32767 > v99) {
                                                                                            v103 := GenS13{F0: v99, F1: v26}
                                                                                            v25 = append(v25, int(v1))
                                                                                            var v104 uint64 = 42
                                                                                            _ = v103
                                                                                            _ = v104
                                                                                        } else {
                                                                                            var v105 string = ""
                                                                                            _ = v105
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    v99 += v1
                                                                                }
                                                                                var v106 uint8 = 1
                                                                                _ = v106
                                                                                _ = v99
                                                                            }
                                                                            v26 += v76
                                                                            v28 = (fi77 - v28)
                                                                            _ = v98
                                                                        }
                                                                    case "":
                                                                        v75 -= 255
                                                                    default:
                                                                        if true {
                                                                            var v107 float64 = v95
                                                                            v1 = 1000
                                                                            fmt.Println(fi77)
                                                                            _ = v107
                                                                        }
                                                                    }
                                                                } else {
                                                                    if true {
                                                                        if (307 > (-(fi60))) {
                                                                            if (!true) {
                                                                                v75 = (v75 * 2)
                                                                            } else {
                                                                                if false {
                                                                                    v108 := GenS16{F0: v28, F1: fi77, F2: 255}
                                                                                    _ = v108
                                                                                }
                                                                                var v109 uint32 = 2
                                                                                _ = v109
                                                                            }
                                                                            v28 = fi77
                                                                            fmt.Println(v75, fi60, v3)
                                                                        } else {
                                                                            var v110 uint = 4294967295
                                                                            _ = v110
                                                                        }
                                                                    }
                                                                    v2[v3] = 0
                                                                }
                                                                _ = v95
                                                            } else {
                                                                v111 := GenS3{F0: (^(v1))}
                                                                _ = v111
                                                            }
                                                        }
                                                    default:
                                                        if false {
                                                            fmt.Println(v75)
                                                        } else {
                                                            if true {
                                                                v26 = (v76 / ((v76) | 1))
                                                            } else {
                                                                v112 := GenS13{F0: v1, F1: (v26 % ((v26) | 1))}
                                                                fmt.Println(v28, fi60, fi77)
                                                                v1 -= (v1 % ((1) | 1))
                                                                _ = v112
                                                            }
                                                            v1 = (^(v1))
                                                            v76 = (v76 & -1)
                                                            var v113 int = (fi60 / ((-381) | 1))
                                                            _ = v113
                                                        }
                                                    }
                                                    var v114 int = (-(v28))
                                                    _ = v114
                                                } else {
                                                    fmt.Println(fi60, v1, v3)
                                                }
                                            }
                                            var v115 bool = false
                                            _ = v115
                                        }
                                        var v116 bool = true
                                        v117 := GenS17{F0: (v26 ^ (v76 | v26)), F1: 4294967295, F2: (v76 >> ((1) & 63))}
                                        _ = v116
                                        _ = v117
                                    }
                                    _ = v84
                                case -317:
                                    v25 = append(v25, v28)
                                default:
                                    var v118 int32 = -2147483648
                                    _ = v118
                                }
                                if (167 < 1) {
                                    v1 = (-(v1))
                                    var v119 float64 = (1e10 + 0.1)
                                    v76 += 42
                                    _ = v119
                                } else {
                                    delete(v2, v3)
                                    fmt.Println(v75, v3)
                                    var v120 uint = 65535
                                    _ = v120
                                }
                            }
                            v76 = 0
                            if (42 < 65535) {
                                var v121 float64 = ((((0.1 - (1e10 - 3.14159)) - (0.1 - ((-2.25 * 0.1) * ((0.1 + 3.14159) + -2.25)))) + 1e10) / (0.1 - 3.14159))
                                _ = v121
                            }
                            var v122 uint32 = v3
                            v3 = v3
                            _ = v122
                        } else {
                            v75 = v75
                            delete(v2, v3)
                            var v123 uint32 = v3
                            _ = v123
                        }
                    case 4294967295:
                        v3 += (v3 & 940)
                    case 32767:
                        if true {
                            _ = len(v2)
                        }
                    }
                    _ = fi77
                    _ = v78
                }
                v2[v3] = (fi60 / (((v28 >> ((uint(v3)) & 63))) | 1))
                var v124 uint8 = (v75 * 42)
                _ = v124
            } else {
                var v125 float64 = 0.1
                _ = v125
            }
            _ = v75
            _ = v76
        }
        fmt.Println(v28, v1, v26)
        _ = fi60
    }
    v126 := GenS18{F0: 255, F1: v26}
    var v127 int64 = v26
    var v128 int = 32768
    for fi129 := 0; fi129 < 1; fi129++ {
        if false {
            v2[808] = (fi129 | (v28 << ((18446744073709551615) & 63)))
            var v130 string = ""
            _ = v130
        }
        v127 = v127
        _ = fi129
    }
    var v131 bool = false
    var v132 int8 = genFn0(128, 42, 42)
    v133 := GenS19{F0: v28, F1: 484, F2: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v127:", v127)
    fmt.Println("v128:", v128)
    fmt.Println("v131:", v131)
    fmt.Println("v132:", v132)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v126
    _ = v127
    _ = v128
    _ = v131
    _ = v132
    _ = v133
    _ = v2
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v59
}
