package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int; F1 string }
type GenS6 struct { F0 int64; F1 string }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint; F1 int64; F2 bool }
type GenS11 struct { F0 uint16; F1 int32 }
type GenS12 struct { F0 uint; F1 uint16; F2 uint }
type GenS13 struct { F0 int16; F1 string; F2 uint64 }
type GenS14 struct { F0 int8; F1 uint }
type GenS15 struct { F0 int8; F1 int16; F2 int16 }
func genFn0() int16 {
    v7 := GenS1{F0: 2, F1: 1}
    var v8 float64 = (float64((1.5 / (1.5 - (1e10 * ((0.1 - -2.25) - -2.25))))) / 1.0)
    var v9 bool = false
    _ = v7
    _ = v8
    _ = v9
    return -1
}
func genFn1(p0 int8, p1 int64, p2 int) int16 {
    var v10 uint = 998
    var v11 float64 = ((float64((0.1 * (0.1 * float32((float64((-2.25 * (1.5 + 0.0))) / 1.5))))) * 1e10) - (0.1 / -2.25))
    var v12 string = ""
    for fi13 := 0; fi13 < 5; fi13++ {
        if (204 > -296) {
            for fi14 := 0; fi14 < 12; fi14++ {
                if false {
                    v15 := GenS1{F0: 1, F1: (p0 ^ -1)}
                    if true {
                        v11 -= float64(1.5)
                        var v16 uint32 = 90
                        var v17 uint32 = 2147483647
                        _ = v16
                        _ = v17
                    }
                    if (0 >= 4294967295) {
                        if (255 >= 28) {
                            fmt.Println(v11, v10, fi13)
                            v11 -= v11
                        }
                    } else {
                        v18 := GenS1{F0: 4294967295, F1: 8}
                        switch v12 {
                        case "foo bar":
                            var v19 bool = false
                            _ = v19
                        }
                        _ = v18
                    }
                    _ = v15
                } else {
                    switch p2 {
                    case 127:
                        v20 := GenS2{F0: (p2 << ((42) & 63)), F1: 1}
                        _ = v20
                    case 42:
                        var v21 int8 = (p0 ^ (p0 >> (((v10 - (v10 + v10))) & 63)))
                        _ = v21
                    case -296:
                        switch fi14 {
                        case 0:
                            var v22 int32 = -2147483648
                            _ = v22
                        case -9223372036854775808:
                            fmt.Println(v12, v10)
                        case 42:
                            p2 -= p2
                        }
                    default:
                        switch p2 {
                        case -2147483648:
                            v23 := GenS3{F0: p1, F1: p2}
                            _ = v23
                        case 128:
                            var v24 bool = true
                            _ = v24
                        case 2147483647:
                            switch fi13 {
                            case 32768:
                                if ((fi13 >> (((v10 | 255)) & 63)) != fi14) {
                                    v11 += ((v11 / 1.5) - (0.1 / v11))
                                }
                            case -398:
                                v12 = ("a" + v12)
                            case -451:
                                var v25 string = v12
                                _ = v25
                            default:
                                var v26 int8 = (p0 >> ((v10) & 63))
                                _ = v26
                            }
                        default:
                            p2 = (fi14 * fi14)
                        }
                    }
                    p1 = -1
                    fmt.Println(p1, fi13)
                }
                _ = fi14
            }
        }
        for fi27 := 0; fi27 < 10; fi27++ {
            p1 = (p1 >> ((v10) & 63))
            _ = fi27
        }
        _ = fi13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return -810
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]int)
    var v3 uint16 = 42
    var v4 uint32 = 967
    var v5 int16 = 0
    var v6 uint32 = v4
    if (2 >= 0) {
        if false {
            v28 := GenS2{F0: 0, F1: (v6 ^ (v4 >> ((2) & 63)))}
            _ = v28
        } else {
            v5 = (v5 - (v5 >> ((2) & 63)))
            for fi29 := 0; fi29 < 16; fi29++ {
                if false {
                    v6 += (v4 * 2147483648)
                }
                v2["a"] = (fi29 & 4294967295)
                var v30 uint32 = (v4 % (((v4 | 242)) | 1))
                _ = fi29
                _ = v30
            }
            v6 = (v6 - v6)
        }
        switch (("\n\t" + "a") + "hello") {
        case "foo bar":
            switch "foo bar" {
            case "\n\t":
                v31 := []int{}
                _ = v31
            case "":
                if (true || true) {
                    v3 -= v3
                    if (!false) {
                        for fi32 := 0; fi32 < 18; fi32++ {
                            switch fi32 {
                            case -1:
                                for fi33 := 0; fi33 < 6; fi33++ {
                                    v34 := v1[v6]
                                    _ = fi33
                                    _ = v34
                                }
                            case 526:
                                v2["foo bar"] = fi32
                            case 9223372036854775807:
                                v35 := GenS2{F0: (fi32 - 0), F1: 42}
                                _ = v35
                            default:
                                v36 := GenS4{F0: 42}
                                _ = v36
                            }
                            fmt.Println(v4, v6, v6)
                            _ = fi32
                        }
                        v3 = (v3 / ((v3) | 1))
                        var v37 float32 = ((0.0 + (-2.25 + (0.0 + 0.1))) + 0.0)
                        _ = v37
                    } else {
                        v4 -= v6
                    }
                    v5 = (v5 << ((4294967295) & 63))
                }
            default:
                v38 := GenS1{F0: 0, F1: 0}
                _ = v38
            }
        default:
            v6 = (^(v6))
        }
    }
    if true {
        v39 := GenS5{F0: 0, F1: ("a" + "foo bar")}
        _ = v39
    } else {
        for fi40 := 0; fi40 < 1; fi40++ {
            if true {
                fmt.Println(v6, v6, fi40)
                fmt.Println(v3, v6)
                if false {
                    v41 := []int{}
                    for fi42 := 0; fi42 < 8; fi42++ {
                        v43 := GenS6{F0: -1, F1: ("a" + ("a" + "\n\t"))}
                        v4 -= (v4 | (v6 * (v4 ^ (v6 << ((18446744073709551615) & 63)))))
                        v44 := GenS3{F0: int64(v6), F1: 127}
                        _ = fi42
                        _ = v43
                        _ = v44
                    }
                    _ = v41
                } else {
                    var v45 uint64 = 0
                    _ = v45
                }
            } else {
                v4 += (v4 - v6)
            }
            v4 -= (v4 * (v6 * v6))
            _ = fi40
        }
        v46 := GenS3{F0: -1, F1: 265}
        for fi47 := 0; fi47 < 5; fi47++ {
            switch (fi47 | fi47) {
            case -374:
                var v48 uint64 = 18446744073709551615
                _ = v48
            default:
                var v49 int8 = 127
                _ = v49
            }
            if (!false) {
                if (1 != 2) {
                    v50 := []int{137}
                    switch -9223372036854775808 {
                    case 32768:
                        v51 := GenS7{F0: 127}
                        _ = v51
                    case 2147483647:
                        if false {
                            var v52 string = ""
                            v50 = append(v50, (fi47 + 1))
                            _ = v52
                        }
                    case 4294967295:
                        if false {
                            var v53 string = ("foo bar" + "\n\t")
                            _ = v53
                        } else {
                            switch fi47 {
                            case 4294967295:
                                var v54 uint64 = 289
                                _ = v54
                            case -1:
                                delete(v2, "hello")
                            case 443:
                                for fi55 := 0; fi55 < 16; fi55++ {
                                    v5 = -1
                                    v6 = (v4 ^ v4)
                                    fmt.Println(v3, fi47, v5)
                                    _ = fi55
                                }
                            }
                            v56 := GenS8{F0: (float64(3.14159) >= 3.14159)}
                            _ = v56
                        }
                    }
                    for fi57 := 0; fi57 < 10; fi57++ {
                        v58 := GenS8{F0: (42 > 278)}
                        _ = fi57
                        _ = v58
                    }
                    _ = v50
                } else {
                    for fi59 := 0; fi59 < 10; fi59++ {
                        switch (("foo bar" + ("foo bar" + ("a" + ""))) + "hello") {
                        case "foo bar":
                            v60 := []int{}
                            _ = v60
                        case "a":
                            if true {
                                v61 := GenS4{F0: -1}
                                if ((fi47 | (fi47 * (fi59 + (fi47 - 32768)))) > fi59) {
                                    switch "\n\t" {
                                    case "\n\t":
                                        if (127 != 42) {
                                            if false {
                                                if (uint(v6) != 0) {
                                                    var v62 bool = true
                                                    fmt.Println(v62, v3)
                                                    fmt.Println(v4)
                                                    var v63 string = "a"
                                                    var v64 int16 = v5
                                                    fmt.Println(fi47)
                                                    if ((!v62) && true) {
                                                        v64 -= v64
                                                    } else {
                                                        v3 -= 2
                                                        var v65 uint64 = 2
                                                        var v66 int64 = 0
                                                        _ = v65
                                                        _ = v66
                                                    }
                                                    v67 := GenS3{F0: int64(v4), F1: fi47}
                                                    v68 := GenS6{F0: 592, F1: v63}
                                                    _ = v62
                                                    _ = v63
                                                    _ = v64
                                                    _ = v67
                                                    _ = v68
                                                }
                                                v6 = 2147483647
                                                v4 = 319
                                                var v69 int64 = 9223372036854775807
                                                var v70 int32 = 42
                                                if true {
                                                    var v71 uint = 92
                                                    _ = v71
                                                } else {
                                                    v4 -= 4294967295
                                                    v72 := []int{}
                                                    v69 -= -879
                                                    _ = v72
                                                }
                                                _ = v69
                                                _ = v70
                                            } else {
                                                var v73 int8 = -127
                                                if true {
                                                    switch int(v3) {
                                                    case 32768:
                                                        v73 += (v73 - (v73 % (((v73 & (v73 / ((-1) | 1)))) | 1)))
                                                    case -106:
                                                        var v74 uint64 = 1
                                                        _ = v74
                                                    case -469:
                                                        if (42 > 18446744073709551615) {
                                                            v75 := []int{fi59, fi59, 2147483647, (fi59 - fi59)}
                                                            v4 = v4
                                                            var v76 float32 = -2.25
                                                            if (v76 == v76) {
                                                                v76 += (v76 - -2.25)
                                                            } else {
                                                                if false {
                                                                    v77 := GenS9{F0: v5}
                                                                    _ = v77
                                                                }
                                                                v73 = v73
                                                            }
                                                            _ = v75
                                                            _ = v76
                                                        } else {
                                                            v78 := v2[("" + "foo bar")]
                                                            _ = v78
                                                        }
                                                    default:
                                                        var v79 uint8 = uint8(v3)
                                                        _ = v79
                                                    }
                                                    v2["hello"] = 1
                                                    var v80 uint64 = 236
                                                    fmt.Println(fi47, v3, v6)
                                                    _ = v80
                                                } else {
                                                    var v81 uint64 = 210
                                                    if false {
                                                        var v82 string = "\n\t"
                                                        _ = v82
                                                    } else {
                                                        fmt.Println(v3)
                                                        v5 = (-(v5))
                                                    }
                                                    _ = v81
                                                }
                                                var v83 uint8 = 128
                                                _ = v73
                                                _ = v83
                                            }
                                            v4 += (v6 << ((0) & 63))
                                        } else {
                                            if (1.5 <= float64(float32((-2.25 * float64(0.0))))) {
                                                _ = len(v2)
                                                if (!false) {
                                                    delete(v1, v4)
                                                    v84 := []int{32768, (fi47 % (((fi47 >> ((0) & 63))) | 1))}
                                                    if false {
                                                        _ = len(v1)
                                                        v5 = (v5 & genFn0())
                                                        v5 += v5
                                                        var v85 int8 = -128
                                                        v3 -= v3
                                                        _ = v85
                                                    } else {
                                                        if true {
                                                            var v86 float64 = ((3.14159 * -2.25) * -2.25)
                                                            if true {
                                                                v84 = append(v84, (fi47 & 347))
                                                            } else {
                                                                v6 += (v4 % (((^(v6))) | 1))
                                                                _ = len(v84)
                                                                v6 -= v4
                                                            }
                                                            _ = v86
                                                        }
                                                        v87 := GenS2{F0: fi59, F1: (^(v6))}
                                                        v88 := GenS10{F0: 65535, F1: 9223372036854775807, F2: true}
                                                        _ = v87
                                                        _ = v88
                                                    }
                                                    if (-589 != -1) {
                                                        var v89 uint64 = 42
                                                        switch "a" {
                                                        case "hello":
                                                            var v90 bool = (!false)
                                                            _ = v90
                                                        case "\n\t":
                                                            switch int(v5) {
                                                            case -2147483648:
                                                                v3 += v3
                                                            case -1:
                                                                if false {
                                                                    v4 += (v4 ^ 669)
                                                                } else {
                                                                    v3 += 42
                                                                    var v91 int64 = 1
                                                                    _ = v91
                                                                }
                                                            case 935:
                                                                v89 = v89
                                                            default:
                                                                if true {
                                                                    var v92 float64 = (1e10 * (-2.25 + (-2.25 - float64(float32(0.1)))))
                                                                    v84 = append(v84, -277)
                                                                    v93 := GenS11{F0: (v3 & (v3 | v3)), F1: 104}
                                                                    _ = len(v84)
                                                                    v3 -= (v3 * 32767)
                                                                    _ = v92
                                                                    _ = v93
                                                                } else {
                                                                    if (true && true) {
                                                                        if len(v84) >= 2 {
                                                                            v84 = v84[1:len(v84)]
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v4, v89)
                                                                        fmt.Println(v6, fi59, fi59)
                                                                    }
                                                                }
                                                            }
                                                        default:
                                                            v94 := GenS3{F0: 42, F1: fi47}
                                                            _ = v94
                                                        }
                                                        var v95 int8 = 95
                                                        _ = v89
                                                        _ = v95
                                                    } else {
                                                        v5 = genFn1(-128, 0, fi59)
                                                        v5 += 0
                                                        v2["hello"] = (fi47 << ((42) & 63))
                                                    }
                                                    _ = v84
                                                }
                                            }
                                            var v96 uint32 = 0
                                            v97 := GenS12{F0: 115, F1: v3, F2: 2}
                                            _ = v96
                                            _ = v97
                                        }
                                    case "foo bar":
                                        switch "a" {
                                        case "\n\t":
                                            var v98 bool = true
                                            _ = v98
                                        case "hello":
                                            if true {
                                                v99 := GenS3{F0: 0, F1: fi59}
                                                v6 -= 35
                                                _ = v99
                                            }
                                        }
                                    default:
                                        v4 -= 0
                                    }
                                    v5 = (v5 * v5)
                                    v5 = (v5 >> ((728) & 63))
                                } else {
                                    switch "" {
                                    case "foo bar":
                                        if true {
                                            delete(v1, (v6 | 0))
                                        }
                                    case "\n\t":
                                        var v100 int8 = -1
                                        _ = v100
                                    }
                                    var v101 float64 = (3.14159 - ((1e10 - ((float64(-2.25) * 3.14159) + 3.14159)) * 1e10))
                                    var v102 uint16 = (v3 << ((65535) & 63))
                                    _ = v101
                                    _ = v102
                                }
                                if (!true) {
                                    if (1.5 >= 0.0) {
                                        v4 += v6
                                        var v103 int8 = -1
                                        _ = v103
                                    } else {
                                        v104 := GenS8{F0: (127 == 0)}
                                        fmt.Println(v5)
                                        _ = v104
                                    }
                                }
                                _ = v61
                            } else {
                                if false {
                                    var v105 uint32 = v6
                                    _ = v105
                                } else {
                                    var v106 float32 = 0.1
                                    v3 = v3
                                    _ = v106
                                }
                                fmt.Println(v6, v3)
                                v6 -= v4
                            }
                            var v107 uint32 = 2
                            v3 -= v3
                            _ = v107
                        }
                        _ = fi59
                    }
                    var v108 uint16 = 686
                    _ = v108
                }
            } else {
                v109 := GenS8{F0: (!(((fi47 <= (^(fi47))) || false) || false))}
                v3 += v3
                for fi110 := 0; fi110 < 3; fi110++ {
                    if false {
                        switch (-(fi110)) {
                        case 42:
                            v111 := []int{(fi47 | fi47)}
                            _ = v111
                        case 988:
                            _ = len(v1)
                        }
                        v3 = (v3 << ((18446744073709551615) & 63))
                    } else {
                        var v112 string = "hello"
                        v5 = -157
                        _ = v112
                    }
                    _ = fi110
                }
                _ = v109
            }
            var v113 uint16 = (v3 % (((^(v3))) | 1))
            _ = fi47
            _ = v113
        }
        _ = v46
    }
    switch "foo bar" {
    case "foo bar":
        for fi114 := 0; fi114 < 16; fi114++ {
            v115 := []int{fi114}
            for fi116 := 0; fi116 < 1; fi116++ {
                if true {
                    v115 = append(v115, fi114)
                    var v117 int32 = 0
                    _ = v117
                }
                if false {
                    v5 += (^(v5))
                } else {
                    var v118 uint16 = v3
                    _ = v118
                }
                _ = fi116
            }
            var v119 int64 = -1
            _ = fi114
            _ = v115
            _ = v119
        }
    }
    var v120 uint64 = 4294967295
    v4 += (^(v6))
    v121 := []int{382, -2147483648}
    fmt.Println(v3, v3, v6)
    for fi122 := 0; fi122 < 20; fi122++ {
        _ = len(v121)
        v4 += v6
        delete(v1, (v6 % ((2147483648) | 1)))
        _ = fi122
    }
    if true {
        var v123 int8 = 38
        _ = len(v121)
        v124 := GenS13{F0: 1000, F1: "", F2: (v120 | (v120 & v120))}
        _ = v123
        _ = v124
    } else {
        if true {
            var v125 float64 = 1.5
            if false {
                var v126 float64 = (v125 / float64(float32(v125)))
                _ = v126
            } else {
                v127 := GenS10{F0: 0, F1: 42, F2: true}
                for fi128 := 0; fi128 < 17; fi128++ {
                    switch (fi128 % (((fi128 | -559)) | 1)) {
                    case 0:
                        for fi129 := 0; fi129 < 16; fi129++ {
                            if false {
                                v3 += 1
                                if (!true) {
                                    var v130 uint8 = 255
                                    switch "\n\t" {
                                    case "":
                                        v121 = append(v121, fi128)
                                    case "hello":
                                        _ = len(v1)
                                    }
                                    switch fi128 {
                                    case 199:
                                        v125 -= v125
                                    case -2147483648:
                                        v121 = append(v121, 9223372036854775807)
                                    case 32768:
                                        fmt.Println(v5)
                                    }
                                    _ = v130
                                } else {
                                    if ((fi129 * 551) <= (-(fi128))) {
                                        v131 := GenS10{F0: 4294967295, F1: 1, F2: false}
                                        _ = v131
                                    }
                                    var v132 int16 = (v5 - v5)
                                    _ = v132
                                }
                                fmt.Println(v3, v125)
                                switch "hello" {
                                case "hello":
                                    v125 -= 3.14159
                                case "foo bar":
                                    if (v4 > 2147483648) {
                                        v133 := GenS4{F0: 410}
                                        fmt.Println(v120)
                                        _ = v133
                                    }
                                case "\n\t":
                                    fmt.Println(fi129, fi129)
                                default:
                                    var v134 int = (fi128 << ((0) & 63))
                                    v134 = -952
                                    _ = v134
                                }
                            } else {
                                var v135 int64 = -9223372036854775808
                                _ = v135
                            }
                            v136 := GenS14{F0: -128, F1: 0}
                            var v137 uint64 = 42
                            _ = fi129
                            _ = v136
                            _ = v137
                        }
                    case -664:
                        v5 += 1
                    }
                    if true {
                        v120 = v120
                        var v138 string = ("" + ("hello" + "hello"))
                        v6 = 65535
                        _ = v138
                    }
                    for fi139 := 0; fi139 < 18; fi139++ {
                        var v140 uint8 = 127
                        _ = fi139
                        _ = v140
                    }
                    _ = fi128
                }
                v6 = (v4 << ((255) & 63))
                _ = v127
            }
            _ = len(v121)
            _ = v125
        } else {
            v6 = (^(v4))
            v141 := GenS15{F0: -1, F1: (^(v5)), F2: (v5 / ((676) | 1))}
            switch 128 {
            case 9223372036854775807:
                if true {
                    v6 = v4
                    _ = len(v2)
                }
            case 4294967295:
                var v142 int64 = 9223372036854775807
                _ = v142
            case 345:
                v143 := GenS11{F0: (v3 % ((v3) | 1)), F1: -1}
                _ = v143
            default:
                v144 := GenS7{F0: -64}
                _ = v144
            }
            _ = v141
        }
        v3 = v3
        var v145 int8 = -1
        _ = v145
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v120:", v120)
    fmt.Println("slicelen v121:", len(v121), v121)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v120
    _ = v121
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
