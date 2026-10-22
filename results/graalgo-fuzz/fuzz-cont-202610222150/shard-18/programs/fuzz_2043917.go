package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint; F1 uint64; F2 uint }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 uint16; F1 uint }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint8; F1 uint8 }
type GenS8 struct { F0 string; F1 uint8 }
type GenS9 struct { F0 uint64; F1 uint16; F2 int8 }
func genFn0() uint32 {
    v5 := GenS1{F0: -741, F1: -130}
    var v6 uint64 = 69
    var v7 int32 = -1
    fmt.Println(v6, v7)
    _ = v5
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{}
    var v3 int8 = 11
    v4 := []int{}
    if false {
        var v8 int16 = 32767
        v9 := GenS1{F0: int64(v8), F1: -584}
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 42, F1: -2147483648}
    fmt.Println(v3)
    v3 += v3
    v3 += v3
    fmt.Println(v3)
    v11 := GenS2{F0: 255, F1: 42, F2: 255}
    v3 += v3
    var v12 uint = 2
    v12 += v12
    var v13 uint16 = 32767
    var v14 uint64 = 2
    for fi15 := 0; fi15 < 11; fi15++ {
        v16 := GenS2{F0: v12, F1: 4294967295, F2: 42}
        if (0.0 > -2.25) {
            if false {
                if false {
                    fmt.Println(v14, v3)
                    switch ("foo bar" + "a") {
                    case "\n\t":
                        fmt.Println(v3, fi15, v3)
                    default:
                        _ = len(v4)
                    }
                } else {
                    v1[("a" + "hello")] = 2147483647
                }
            }
            var v17 uint = (v12 + 924)
            _ = v17
        }
        _ = fi15
        _ = v16
    }
    if false {
        if (!("a" != "\n\t")) {
            if (!((-2.25 * 1.5) < 0.0)) {
                if false {
                    var v18 string = "foo bar"
                    _ = v18
                } else {
                    var v19 uint64 = v14
                    _ = v19
                }
                if false {
                    var v20 uint8 = 1
                    var v21 uint8 = v20
                    switch 0 {
                    case 1:
                        v1[("" + "\n\t")] = 2147483647
                    case 9223372036854775807:
                        if true {
                            switch 4294967295 {
                            case 0:
                                var v22 float32 = (3.14159 + 0.0)
                                _ = v22
                            case 220:
                                if ("hello" != "foo bar") {
                                    for fi23 := 0; fi23 < 14; fi23++ {
                                        v21 = 178
                                        var v24 int64 = 1
                                        _ = fi23
                                        _ = v24
                                    }
                                    if (!false) {
                                        v4 = append(v4, -921)
                                        if true {
                                            for fi25 := 0; fi25 < 12; fi25++ {
                                                if true {
                                                    if true {
                                                        var v26 int16 = -310
                                                        v20 = 215
                                                        _ = v26
                                                    } else {
                                                        fmt.Println(v13, v13)
                                                    }
                                                    var v27 int64 = 1
                                                    v28 := GenS3{F0: v27, F1: (v12 / ((v12) | 1))}
                                                    _ = v27
                                                    _ = v28
                                                }
                                                _ = fi25
                                            }
                                            if (!false) {
                                                var v29 float64 = -2.25
                                                var v30 float32 = 0.1
                                                v20 += v20
                                                _ = v29
                                                _ = v30
                                            } else {
                                                if false {
                                                    v31 := v1["\n\t"]
                                                    switch 4294967295 {
                                                    case -9223372036854775808:
                                                        var v32 int32 = 170
                                                        _ = v32
                                                    default:
                                                        if true {
                                                            fmt.Println(v31)
                                                        } else {
                                                            v3 = v3
                                                            if true {
                                                                for fi33 := 0; fi33 < 7; fi33++ {
                                                                    for fi34 := 0; fi34 < 2; fi34++ {
                                                                        v20 += v20
                                                                        _ = fi34
                                                                    }
                                                                    fmt.Println(v3, v13, v31)
                                                                    var v35 float64 = 3.14159
                                                                    _ = fi33
                                                                    _ = v35
                                                                }
                                                            }
                                                        }
                                                    }
                                                    v1[""] = uint32(v14)
                                                    _ = v31
                                                }
                                                var v36 float64 = (0.1 * ((0.1 / ((0.1 * 0.1) * ((0.1 * -2.25) / (3.14159 * 1.5)))) / (1.5 * (-2.25 * ((-2.25 * -2.25) * -2.25)))))
                                                _ = v36
                                            }
                                            v2 = append(v2, 128)
                                        } else {
                                            fmt.Println(v12)
                                            var v37 int64 = 0
                                            _ = v37
                                        }
                                    } else {
                                        v38 := GenS3{F0: 42, F1: v12}
                                        for fi39 := 0; fi39 < 13; fi39++ {
                                            switch fi39 {
                                            case 128:
                                                fmt.Println(v14, v14, v3)
                                            case 0:
                                                v14 = (v14 % ((0) | 1))
                                            case 1:
                                                var v40 string = ""
                                                _ = v40
                                            }
                                            v41 := GenS4{F0: v13, F1: (v12 / ((v12) | 1))}
                                            _ = fi39
                                            _ = v41
                                        }
                                        _ = v38
                                    }
                                } else {
                                    v42 := GenS5{F0: true}
                                    _ = v42
                                }
                            }
                            v21 -= (v21 >> ((v12) & 63))
                        } else {
                            var v43 uint32 = 1
                            var v44 uint64 = v14
                            for fi45 := 0; fi45 < 10; fi45++ {
                                fmt.Println(v14, v43, v21)
                                v46 := GenS6{F0: uint8(v14)}
                                if (true && false) {
                                    v43 = (v43 << ((v12) & 63))
                                    if ((542 > -48) || (!true)) {
                                        v47 := GenS5{F0: (!true)}
                                        _ = v47
                                    } else {
                                        v21 = v20
                                        for fi48 := 0; fi48 < 12; fi48++ {
                                            v21 = (v21 ^ v20)
                                            var v49 bool = false
                                            switch (^(fi48)) {
                                            case 42:
                                                fmt.Println(v13, v49)
                                            default:
                                                fmt.Println(v14)
                                            }
                                            _ = fi48
                                            _ = v49
                                        }
                                        v50 := GenS6{F0: v21}
                                        _ = v50
                                    }
                                    v1[("hello" + "foo bar")] = genFn0()
                                }
                                _ = fi45
                                _ = v46
                            }
                            _ = v43
                            _ = v44
                        }
                    case 42:
                        var v51 uint32 = 42
                        _ = v51
                    }
                    _ = v20
                    _ = v21
                } else {
                    if (v3 > v3) {
                        var v52 int = -9223372036854775808
                        var v53 int8 = -73
                        _ = v52
                        _ = v53
                    } else {
                        v54 := GenS1{F0: 100, F1: -1}
                        if false {
                            if ((v14 & v14) < 1) {
                                if true {
                                    v12 -= (v12 ^ 1)
                                    v4 = append(v4, 39)
                                }
                                if true {
                                    for fi55 := 0; fi55 < 1; fi55++ {
                                        v3 -= (v3 ^ (v3 << ((v12) & 63)))
                                        _ = fi55
                                    }
                                    for fi56 := 0; fi56 < 18; fi56++ {
                                        if false {
                                            for fi57 := 0; fi57 < 6; fi57++ {
                                                v12 -= v12
                                                v3 = v3
                                                v58 := GenS3{F0: 0, F1: 2}
                                                _ = fi57
                                                _ = v58
                                            }
                                        }
                                        _ = fi56
                                    }
                                }
                            }
                            fmt.Println(v13)
                            var v59 string = "hello"
                            _ = v59
                        }
                        var v60 int16 = -433
                        _ = v54
                        _ = v60
                    }
                }
            } else {
                switch -435 {
                case 127:
                    if true {
                        if (v13 == 638) {
                            v61 := GenS7{F0: 1, F1: 42}
                            _ = v61
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v62 := GenS1{F0: int64(v3), F1: -426}
                            _ = v62
                        }
                        if true {
                            for fi63 := 0; fi63 < 2; fi63++ {
                                switch (fi63 - (fi63 >> ((v12) & 63))) {
                                case -707:
                                    if true {
                                        switch (("foo bar" + "a") + "hello") {
                                        case "\n\t":
                                            var v64 uint8 = 2
                                            _ = v64
                                        case "foo bar":
                                            v12 = v12
                                        default:
                                            for fi65 := 0; fi65 < 16; fi65++ {
                                                delete(v1, "")
                                                _ = fi65
                                            }
                                        }
                                        if (!((3.14159 + 1e10) == 3.14159)) {
                                            for fi66 := 0; fi66 < 4; fi66++ {
                                                var v67 uint8 = 1
                                                fmt.Println(fi66)
                                                _ = fi66
                                                _ = v67
                                            }
                                            var v68 int = fi63
                                            _ = v68
                                        }
                                    } else {
                                        if true {
                                            v69 := GenS3{F0: 0, F1: v12}
                                            if len(v4) > 0 {
                                                v4[0] = fi63
                                            }
                                            if true {
                                                for fi70 := 0; fi70 < 18; fi70++ {
                                                    if false {
                                                        var v71 float64 = (3.14159 - 0.1)
                                                        v72 := GenS3{F0: -677, F1: (v12 ^ v12)}
                                                        if (221 == v13) {
                                                            var v73 int16 = -32768
                                                            var v74 int8 = v3
                                                            var v75 float32 = -2.25
                                                            _ = v73
                                                            _ = v74
                                                            _ = v75
                                                        }
                                                        _ = v71
                                                        _ = v72
                                                    } else {
                                                        v3 -= v3
                                                        var v76 uint16 = (v13 >> (((v12 & v12)) & 63))
                                                        fmt.Println(v76, v13)
                                                        _ = v76
                                                    }
                                                    if (v14 > 2) {
                                                        v77 := GenS4{F0: (^(v13)), F1: (v12 - (v12 << (((v12 / ((v12) | 1))) & 63)))}
                                                        fmt.Println(v14)
                                                        switch fi70 {
                                                        case 32767:
                                                            if false {
                                                                if ((v14 + (v14 | (^(v14)))) >= 0) {
                                                                    v78 := GenS5{F0: true}
                                                                    _ = v78
                                                                } else {
                                                                    v79 := GenS6{F0: 1}
                                                                    switch fi63 {
                                                                    case 4294967295:
                                                                        v14 += (v14 * v14)
                                                                    case 32768:
                                                                        var v80 uint8 = 255
                                                                        _ = v80
                                                                    case -931:
                                                                        var v81 int64 = -820
                                                                        _ = v81
                                                                    default:
                                                                        v3 = v3
                                                                    }
                                                                    var v82 uint8 = 2
                                                                    _ = v79
                                                                    _ = v82
                                                                }
                                                                if (105 > 2147483648) {
                                                                    switch fi63 {
                                                                    case -9223372036854775808:
                                                                        v3 -= v3
                                                                    case 32767:
                                                                        var v83 string = (("a" + ("foo bar" + "hello")) + "\n\t")
                                                                        _ = v83
                                                                    case 9223372036854775807:
                                                                        v2 = append(v2, 42)
                                                                    }
                                                                }
                                                                var v84 int64 = 1
                                                                _ = v84
                                                            } else {
                                                                var v85 string = "\n\t"
                                                                _ = v85
                                                            }
                                                        case 317:
                                                            v86 := GenS2{F0: (^(v12)), F1: (^(v14)), F2: (v12 >> ((v12) & 63))}
                                                            _ = v86
                                                        case 4294967295:
                                                            var v87 uint32 = 2
                                                            _ = v87
                                                        default:
                                                            switch ("\n\t" + "\n\t") {
                                                            case "\n\t":
                                                                if (!false) {
                                                                    var v88 int = (^(fi63))
                                                                    _ = v88
                                                                }
                                                            case "":
                                                                fmt.Println(fi63)
                                                            default:
                                                                v13 = v13
                                                            }
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                var v89 bool = true
                                                                _ = v89
                                                            default:
                                                                v13 = (^(v13))
                                                            }
                                                            v13 = (v13 * (v13 / ((v13) | 1)))
                                                        }
                                                        _ = v77
                                                    } else {
                                                        v13 = v13
                                                    }
                                                    _ = fi70
                                                }
                                                for fi90 := 0; fi90 < 14; fi90++ {
                                                    var v91 bool = false
                                                    _ = fi90
                                                    _ = v91
                                                }
                                                for fi92 := 0; fi92 < 16; fi92++ {
                                                    fmt.Println(fi63, fi63)
                                                    var v93 uint8 = 183
                                                    v93 += v93
                                                    _ = fi92
                                                    _ = v93
                                                }
                                            } else {
                                                var v94 uint = (v12 ^ v12)
                                                fmt.Println(v13)
                                                if (v12 >= (v94 / (((v12 & (v12 * v12))) | 1))) {
                                                    v95 := GenS5{F0: false}
                                                    _ = v95
                                                } else {
                                                    var v96 string = ("\n\t" + ("\n\t" + ""))
                                                    var v97 string = (v96 + v96)
                                                    fmt.Println(v97, v12, v3)
                                                    _ = v96
                                                    _ = v97
                                                }
                                                _ = v94
                                            }
                                            _ = v69
                                        }
                                        v98 := GenS5{F0: true}
                                        _ = v98
                                    }
                                default:
                                    v14 = v14
                                }
                                var v99 int = fi63
                                _ = fi63
                                _ = v99
                            }
                            v12 -= 65535
                        } else {
                            for fi100 := 0; fi100 < 4; fi100++ {
                                var v101 string = ("" + ("a" + "foo bar"))
                                switch (v101 + v101) {
                                case "a":
                                    fmt.Println(v101, v101, v13)
                                case "":
                                    v102 := GenS7{F0: 48, F1: 0}
                                    _ = v102
                                default:
                                    v101 = "\n\t"
                                }
                                _ = fi100
                                _ = v101
                            }
                            v14 -= (v14 % ((v14) | 1))
                        }
                        var v103 float32 = 0.0
                        _ = v103
                    } else {
                        v3 -= (-(v3))
                        v3 = v3
                    }
                case 2147483647:
                    v104 := v1["a"]
                    _ = v104
                default:
                    for fi105 := 0; fi105 < 15; fi105++ {
                        var v106 int32 = 1
                        _ = fi105
                        _ = v106
                    }
                }
                switch (("" + "foo bar") + "\n\t") {
                case "foo bar":
                    fmt.Println(v12)
                default:
                    var v107 int32 = -1
                    _ = v107
                }
            }
            if false {
                v4 = append(v4, 128)
            }
        }
        fmt.Println(v13, v12)
    }
    if false {
        if (!false) {
            fmt.Println(v13, v13)
            if (0.1 > (((float64(3.14159) - float64((float32(1.5) - 0.1))) * (1e10 / 1.5)) / -2.25)) {
                v108 := GenS8{F0: "hello", F1: 31}
                var v109 int64 = -198
                _ = v108
                _ = v109
            } else {
                if ((v3 % ((v3) | 1)) <= 1) {
                    v12 += 914
                } else {
                    if (!((true && true) && false)) {
                        if false {
                            for fi110 := 0; fi110 < 11; fi110++ {
                                if false {
                                    if (11 < 0) {
                                        var v111 uint32 = 2147483648
                                        _ = v111
                                    } else {
                                        v1["hello"] = 2147483648
                                        var v112 bool = (263 == 521)
                                        var v113 float32 = (0.0 - (float32((-2.25 + 0.0)) - 3.14159))
                                        _ = v112
                                        _ = v113
                                    }
                                    for fi114 := 0; fi114 < 5; fi114++ {
                                        v4 = append(v4, fi110)
                                        fmt.Println(v12, v3)
                                        switch int(v3) {
                                        case 0:
                                            delete(v1, "\n\t")
                                        default:
                                            v3 += int8(fi110)
                                        }
                                        var v115 uint8 = 213
                                        _ = fi114
                                        _ = v115
                                    }
                                    v14 -= 4294967295
                                }
                                _ = fi110
                            }
                            for fi116 := 0; fi116 < 12; fi116++ {
                                v13 = v13
                                switch (fi116 - 0) {
                                case -9223372036854775808:
                                    var v117 int16 = int16(v13)
                                    _ = v117
                                }
                                if (((true && false) && ((fi116 ^ (fi116 | fi116)) > fi116)) || (true && true)) {
                                    for fi118 := 0; fi118 < 15; fi118++ {
                                        switch ("" + "") {
                                        case "hello":
                                            if (-2.25 > float32(0.1)) {
                                                var v119 float64 = 1.5
                                                var v120 int8 = v3
                                                _ = v119
                                                _ = v120
                                            } else {
                                                var v121 int32 = 1
                                                var v122 uint16 = (^(v13))
                                                _ = v121
                                                _ = v122
                                            }
                                        case "foo bar":
                                            v12 -= (v12 & v12)
                                        default:
                                            v2 = append(v2, (fi118 | (fi118 % ((fi118) | 1))))
                                        }
                                        var v123 string = ""
                                        switch "a" {
                                        case "\n\t":
                                            _ = len(v4)
                                        default:
                                            v3 += (v3 >> (((v12 - (v12 << (((v12 | (v12 >> (((v12 >> ((v12) & 63))) & 63)))) & 63)))) & 63))
                                        }
                                        _ = fi118
                                        _ = v123
                                    }
                                    fmt.Println(v3, v12)
                                    var v124 float64 = (0.0 * 0.0)
                                    _ = v124
                                }
                                _ = fi116
                            }
                            switch ("foo bar" + "foo bar") {
                            case "a":
                                if false {
                                    fmt.Println(v3, v3, v12)
                                } else {
                                    var v125 string = ""
                                    _ = v125
                                }
                            default:
                                var v126 bool = (false && ((v3 >> (((v12 & (v12 << ((0) & 63)))) & 63)) <= v3))
                                _ = v126
                            }
                        }
                        var v127 uint32 = 2147483647
                        _ = v127
                    }
                }
            }
        } else {
            fmt.Println(v12, v13)
            var v128 int32 = 1
            if len(v4) > 0 {
                v4[0] = -460
            }
            _ = v128
        }
        if (!(144 > 0)) {
            _ = len(v1)
            _ = len(v2)
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    } else {
        if false {
            for fi129 := 0; fi129 < 10; fi129++ {
                v12 += (v12 ^ 474)
                if false {
                    switch ("foo bar" + ("\n\t" + "a")) {
                    case "hello":
                        v130 := GenS4{F0: v13, F1: v12}
                        _ = v130
                    case "":
                        if len(v2) > 0 {
                            v2[0] = fi129
                        }
                    }
                    for fi131 := 0; fi131 < 13; fi131++ {
                        if (v13 != v13) {
                            delete(v1, ("\n\t" + "hello"))
                        }
                        _ = fi131
                    }
                } else {
                    delete(v1, ("" + "foo bar"))
                }
                _ = fi129
            }
            v132 := GenS1{F0: 42, F1: -2147483648}
            _ = v132
        } else {
            v12 -= v12
            var v133 string = ""
            _ = v133
        }
    }
    if len(v4) > 0 {
        v4[0] = 90
    }
    if (0 >= -86) {
        for fi134 := 0; fi134 < 13; fi134++ {
            var v135 uint32 = uint32(v14)
            _ = fi134
            _ = v135
        }
        var v136 int32 = 42
        if (v12 == (v12 >> ((v12) & 63))) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v137 := GenS9{F0: v14, F1: 32767, F2: (^(v3))}
            _ = v137
        }
        _ = v136
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
