package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int; F1 int64 }
type GenS7 struct { F0 uint16; F1 int64; F2 int }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int32; F1 uint64 }
type GenS12 struct { F0 int; F1 uint16 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 int }
type GenS15 struct { F0 string; F1 bool }
func genFn0(p0 uint16, p1 int) int {
    if ((1 > 2) && false) {
        v6 := GenS1{F0: 1000, F1: "hello"}
        _ = v6
    } else {
        if false {
            var v7 int8 = 0
            v7 = v7
            if true {
                v7 -= v7
            } else {
                for fi8 := 0; fi8 < 7; fi8++ {
                    p0 = 2
                    var v9 float64 = (float64(((1.5 * float32((1.5 / 3.14159))) * 0.0)) + float64(3.14159))
                    v10 := GenS2{F0: 6, F1: -9223372036854775808, F2: 42}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            }
            _ = v7
        } else {
            for fi11 := 0; fi11 < 16; fi11++ {
                v12 := []int{-650, p1}
                p0 = (p0 / ((p0) | 1))
                _ = fi11
                _ = v12
            }
            var v13 uint32 = 65535
            _ = v13
        }
        var v14 uint32 = 2
        p0 -= p0
        _ = v14
    }
    switch "" {
    case "a":
        p0 -= 2
    case "\n\t":
        switch p1 {
        case -443:
            switch ("a" + "") {
            case "foo bar":
                v15 := GenS1{F0: 1, F1: "a"}
                _ = v15
            }
        case 128:
            v16 := []int{946, 2147483647, (p1 >> ((594) & 63))}
            _ = v16
        }
    case "foo bar":
        if (9223372036854775807 > 0) {
            switch (p1 + (p1 * p1)) {
            case 42:
                var v17 uint64 = 42
                _ = v17
            case 1:
                fmt.Println(p0)
            }
        } else {
            p1 = p1
            p0 -= p0
            switch 99 {
            case -803:
                v18 := GenS3{F0: 296}
                _ = v18
            case -885:
                if ((true || ((309 > -1) && ((0 != -1) || true))) || true) {
                    if false {
                        if true {
                            fmt.Println(p1)
                            v19 := GenS4{F0: 0, F1: 790}
                            _ = v19
                        }
                        v20 := GenS2{F0: 255, F1: 0, F2: 55}
                        _ = v20
                    }
                    var v21 int16 = 1
                    var v22 int = p1
                    _ = v21
                    _ = v22
                } else {
                    var v23 bool = false
                    var v24 uint64 = 4294967295
                    _ = v23
                    _ = v24
                }
            }
        }
    }
    var v25 uint8 = 223
    switch p1 {
    case 127:
        v26 := GenS4{F0: 1000, F1: 0}
        _ = v26
    default:
        if ((-128 <= -1) && true) {
            switch p1 {
            case 32768:
                v25 -= 1
            case 4294967295:
                var v27 float32 = float32(0.1)
                _ = v27
            case 9223372036854775807:
                v28 := GenS4{F0: 1, F1: 1}
                _ = v28
            }
            var v29 float64 = (0.1 + 1.5)
            switch 4294967295 {
            case -9223372036854775808:
                if true {
                    var v30 uint32 = 65535
                    var v31 bool = true
                    var v32 string = ("\n\t" + "\n\t")
                    _ = v30
                    _ = v31
                    _ = v32
                }
            case 0:
                for fi33 := 0; fi33 < 20; fi33++ {
                    var v34 int8 = 127
                    _ = fi33
                    _ = v34
                }
            default:
                var v35 uint64 = 42
                _ = v35
            }
            _ = v29
        } else {
            if true {
                p0 = (p0 << ((18446744073709551615) & 63))
                var v36 bool = true
                _ = v36
            } else {
                v37 := GenS2{F0: v25, F1: -1, F2: v25}
                if false {
                    v25 = v25
                } else {
                    v25 += v25
                    v38 := []int{}
                    if true {
                        p1 -= int(p0)
                    }
                    _ = v38
                }
                p1 = p1
                _ = v37
            }
            switch "foo bar" {
            case "":
                var v39 int8 = 42
                _ = v39
            case "a":
                var v40 int32 = 2147483647
                _ = v40
            default:
                fmt.Println(v25, v25, p1)
            }
            fmt.Println(v25, p0, p1)
        }
    }
    var v41 uint8 = v25
    _ = p0
    _ = p1
    _ = v25
    _ = v41
    return p1
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[int]uint32)
    var v3 int8 = 1
    v4 := []uint32{}
    var v5 int = 42
    switch v5 {
    case 128:
        for fi42 := 0; fi42 < 4; fi42++ {
            v5 += (v5 << ((899) & 63))
            var v43 int = 32768
            if true {
                v3 = (v3 + (v3 * (v3 / (((v3 % ((v3) | 1))) | 1))))
                if false {
                    fmt.Println(v5)
                } else {
                    var v44 string = "\n\t"
                    var v45 int16 = -32768
                    _ = v44
                    _ = v45
                }
                v46 := v2[genFn0(2, genFn0(2, genFn0(42, (fi42 ^ -882))))]
                _ = v46
            }
            _ = fi42
            _ = v43
        }
    case 414:
        if (-336 != 0) {
            var v47 bool = false
            v48 := GenS5{F0: 1}
            v49 := GenS4{F0: 172, F1: -418}
            _ = v47
            _ = v48
            _ = v49
        } else {
            switch ("hello" + ((("foo bar" + "a") + "foo bar") + "\n\t")) {
            case "a":
                if true {
                    v50 := GenS3{F0: 255}
                    var v51 int8 = v3
                    _ = v50
                    _ = v51
                } else {
                    fmt.Println(v5, v1)
                    v5 = v5
                }
            default:
                v52 := GenS4{F0: -32768, F1: 0}
                _ = v52
            }
        }
    }
    v4 = append(v4, 311)
    delete(v2, v5)
    switch ("foo bar" + "foo bar") {
    case "hello":
        v53 := GenS6{F0: (v5 * (v5 << ((2) & 63))), F1: -9223372036854775808}
        _ = v53
    }
    var v54 int16 = 42
    for fi55 := 0; fi55 < 15; fi55++ {
        if (255 > 573) {
            v5 -= fi55
            if true {
                v56 := GenS5{F0: 951}
                _ = v56
            } else {
                v57 := GenS2{F0: 128, F1: 1, F2: 0}
                v58 := GenS7{F0: 2, F1: int64(v5), F2: (v5 >> ((65535) & 63))}
                _ = v57
                _ = v58
            }
            v59 := GenS8{F0: 1}
            _ = v59
        } else {
            fmt.Println(v5, v54, fi55)
        }
        switch 2147483647 {
        case -9223372036854775808:
            v1 = float32(1.5)
        case 32767:
            if (((int16(v3) == v54) || false) || (0.1 <= ((float32(1.5) / (3.14159 / v1)) - ((v1 / float32(3.14159)) + float32(-2.25))))) {
                var v60 uint16 = 2
                _ = v60
            }
        case -173:
            if true {
                _ = len(v4)
            } else {
                for fi61 := 0; fi61 < 18; fi61++ {
                    v1 -= float32(-2.25)
                    var v62 string = ("hello" + "\n\t")
                    _ = fi61
                    _ = v62
                }
                v63 := GenS9{F0: 1, F1: -1, F2: 0}
                switch "" {
                case "hello":
                    if false {
                        fmt.Println(v5, fi55, v5)
                        var v64 int8 = -128
                        _ = v64
                    }
                case "foo bar":
                    if (!false) {
                        v5 += v5
                        if (((554 != 1) || false) || true) {
                            var v65 string = "\n\t"
                            v66 := v2[4294967295]
                            v54 -= (v54 >> ((4294967295) & 63))
                            _ = v65
                            _ = v66
                        } else {
                            var v67 int32 = -1
                            var v68 int = genFn0(0, (fi55 + (fi55 | (fi55 % ((fi55) | 1)))))
                            var v69 int8 = 1
                            _ = v67
                            _ = v68
                            _ = v69
                        }
                        v3 = v3
                    } else {
                        v70 := GenS7{F0: 42, F1: 823, F2: v5}
                        v71 := GenS1{F0: (v54 >> ((2) & 63)), F1: "a"}
                        v1 -= v1
                        _ = v70
                        _ = v71
                    }
                case "a":
                    for fi72 := 0; fi72 < 4; fi72++ {
                        v73 := GenS6{F0: (-(fi72)), F1: 0}
                        if true {
                            v1 += (((1e10 * ((3.14159 * v1) * ((float32((((3.14159 / float64(float32(1e10))) * 1.5) * (-2.25 / 1e10))) / v1) * 1e10))) + (v1 * (v1 * (v1 - v1)))) / (-2.25 / 0.1))
                            var v74 int64 = -567
                            var v75 uint = 234
                            v76 := GenS6{F0: (^(fi72)), F1: (^(v74))}
                            _ = v74
                            _ = v75
                            _ = v76
                        } else {
                            if false {
                                var v77 int16 = v54
                                v77 = v77
                                _ = v77
                            }
                            v78 := GenS5{F0: 4294967295}
                            if false {
                                v79 := GenS5{F0: 0}
                                _ = v79
                            }
                            _ = v78
                        }
                        var v80 uint64 = 0
                        if false {
                            var v81 bool = (!(!false))
                            if true {
                                switch "" {
                                case "":
                                    v2[v5] = 2
                                    var v82 int64 = -424
                                    _ = v82
                                case "hello":
                                    if v81 {
                                        v3 = v3
                                        var v83 uint64 = 286
                                        fmt.Println(v5, v80)
                                        _ = v83
                                    } else {
                                        switch "" {
                                        case "foo bar":
                                            v54 -= v54
                                        case "\n\t":
                                            if (!v81) {
                                                v84 := GenS1{F0: (-(v54)), F1: ""}
                                                if (v1 >= (((v1 - 1e10) - v1) * v1)) {
                                                    v85 := GenS8{F0: 110}
                                                    v86 := GenS9{F0: 0, F1: (v54 % (((v54 / ((v54) | 1))) | 1)), F2: 1}
                                                    if v81 {
                                                        v87 := GenS6{F0: (v5 | (fi72 ^ v5)), F1: 0}
                                                        switch (-(fi72)) {
                                                        case 0:
                                                            v54 = (v54 & v54)
                                                        default:
                                                            _ = len(v2)
                                                        }
                                                        _ = v87
                                                    } else {
                                                        var v88 float64 = 1e10
                                                        if (0 < 87) {
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                var v89 bool = (("\n\t" + "\n\t") < "hello")
                                                                _ = v89
                                                            case "a":
                                                                var v90 uint32 = 65535
                                                                _ = v90
                                                            default:
                                                                v91 := v2[v5]
                                                                v92 := GenS2{F0: uint8(v80), F1: 499, F2: 48}
                                                                _ = v91
                                                                _ = v92
                                                            }
                                                            switch fi55 {
                                                            case 0:
                                                                v1 = (v1 * float32(0.1))
                                                            }
                                                        } else {
                                                            v81 = v81
                                                        }
                                                        _ = v88
                                                    }
                                                    _ = v85
                                                    _ = v86
                                                }
                                                v2[v5] = 65535
                                                v93 := GenS10{F0: 65535}
                                                _ = v84
                                                _ = v93
                                            } else {
                                                var v94 int64 = -401
                                                if (42 <= 0) {
                                                    _ = len(v2)
                                                }
                                                v95 := GenS4{F0: (-(v54)), F1: v94}
                                                _ = v94
                                                _ = v95
                                            }
                                        case "hello":
                                            fmt.Println(v3, v80)
                                        default:
                                            if false {
                                                if v81 {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    if true {
                                                        v2[genFn0(295, -9223372036854775808)] = 239
                                                        var v96 int = -9223372036854775808
                                                        v97 := GenS11{F0: -1, F1: (v80 * v80)}
                                                        _ = v96
                                                        _ = v97
                                                    }
                                                    var v98 int32 = 42
                                                    _ = v98
                                                } else {
                                                    var v99 string = ""
                                                    var v100 string = (v99 + v99)
                                                    _ = v100
                                                    _ = v99
                                                }
                                                var v101 float64 = 1.5
                                                switch "a" {
                                                case "":
                                                    var v102 int32 = -757
                                                    _ = v102
                                                case "a":
                                                    v101 -= 1.5
                                                case "hello":
                                                    fmt.Println(v3)
                                                default:
                                                    if (-2147483648 > -688) {
                                                        var v103 int64 = 9223372036854775807
                                                        _ = v103
                                                    }
                                                }
                                                _ = v101
                                            }
                                        }
                                        if true {
                                            if v81 {
                                                v4 = append(v4, 375)
                                                v104 := GenS3{F0: 4294967295}
                                                var v105 string = "\n\t"
                                                var v106 bool = v81
                                                _ = v104
                                                _ = v105
                                                _ = v106
                                            } else {
                                                v107 := GenS12{F0: genFn0(473, v5), F1: 872}
                                                fmt.Println(v3)
                                                _ = v107
                                            }
                                        }
                                        v108 := GenS12{F0: genFn0(723, fi55), F1: 2}
                                        _ = v108
                                    }
                                default:
                                    if v81 {
                                        v54 += 1000
                                    }
                                }
                                var v109 int32 = 2147483647
                                var v110 int8 = v3
                                if false {
                                    switch (v5 + genFn0(153, int(v80))) {
                                    case 2147483647:
                                        fmt.Println(fi55, fi55)
                                    default:
                                        if v81 {
                                            if true {
                                                v5 += fi55
                                                if len(v4) > 0 {
                                                    v4[0] = 902
                                                }
                                            } else {
                                                v110 -= 127
                                                var v111 float32 = v1
                                                v112 := GenS13{F0: (v110 % ((int8(fi72)) | 1))}
                                                _ = v111
                                                _ = v112
                                            }
                                        } else {
                                            v4 = append(v4, 2)
                                            v4 = append(v4, uint32(v80))
                                            var v113 string = "\n\t"
                                            _ = v113
                                        }
                                    }
                                    if (2147483648 > 572) {
                                        fmt.Println(v3, fi72, v81)
                                    }
                                } else {
                                    v114 := GenS14{F0: (fi72 - fi72)}
                                    v115 := GenS15{F0: ("" + ("" + "a")), F1: true}
                                    _ = v114
                                    _ = v115
                                }
                                _ = v109
                                _ = v110
                            } else {
                                fmt.Println(v80, v80)
                                switch (-(fi72)) {
                                case 127:
                                    v116 := GenS10{F0: 2}
                                    _ = v116
                                case 32768:
                                    switch "foo bar" {
                                    case "\n\t":
                                        fmt.Println(v54, v3)
                                    case "hello":
                                        v54 += (^(v54))
                                    case "foo bar":
                                        if (65535 == 992) {
                                            var v117 int8 = (-(v3))
                                            fmt.Println(v117, v117)
                                            _ = v117
                                        } else {
                                            fmt.Println(v80)
                                            v54 -= int16(v80)
                                        }
                                    }
                                default:
                                    switch (fi55 ^ genFn0(0, (-(fi72)))) {
                                    case -2147483648:
                                        v1 = v1
                                    case 1:
                                        if ((v3 < v3) || ((v54 * (v54 % ((int16(v3)) | 1))) >= (v54 & (v54 | v54)))) {
                                            fmt.Println(v81, fi72, v3)
                                            if (!v81) {
                                                if (("foo bar" + "foo bar") >= "foo bar") {
                                                    var v118 uint8 = 164
                                                    var v119 int16 = 437
                                                    _ = v118
                                                    _ = v119
                                                }
                                                if ((v3 < (v3 & (v3 << ((4294967295) & 63)))) || (int8(v80) >= v3)) {
                                                    v5 = genFn0(1, v5)
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                } else {
                                                    v54 = v54
                                                }
                                            }
                                            var v120 bool = ((198 == 0) && v81)
                                            _ = v120
                                        }
                                    }
                                }
                            }
                            var v121 int32 = 603
                            _ = v121
                            _ = v81
                        } else {
                            switch 42 {
                            case 1:
                                var v122 int = 32768
                                _ = v122
                            default:
                                v3 += v3
                                var v123 uint = 1
                                _ = v123
                            }
                            if (!("a" < "hello")) {
                                var v124 uint8 = 1
                                fmt.Println(v5, v124)
                                _ = v124
                            } else {
                                v54 = 1000
                            }
                            var v125 uint64 = v80
                            _ = v125
                        }
                        v80 -= v80
                        _ = fi72
                        _ = v73
                        _ = v80
                    }
                default:
                    var v126 float64 = (0.1 + 0.0)
                    _ = v126
                }
                _ = v63
            }
        }
        _ = fi55
    }
    var v127 float64 = 0.1
    var v128 uint = 2
    if true {
        if ((81 >= 42) || (0 > 2)) {
            var v129 int8 = v3
            v2[(v5 & genFn0(0, (v5 >> (((v128 * (v128 / ((624) | 1)))) & 63))))] = 0
            var v130 uint16 = 2
            _ = v129
            _ = v130
        } else {
            for fi131 := 0; fi131 < 18; fi131++ {
                if true {
                    var v132 int16 = (v54 | v54)
                    _ = v132
                } else {
                    for fi133 := 0; fi133 < 4; fi133++ {
                        fmt.Println(fi133)
                        v134 := v2[(-(v5))]
                        var v135 bool = true
                        _ = fi133
                        _ = v134
                        _ = v135
                    }
                    fmt.Println(v127, v127, v128)
                }
                fmt.Println(v1)
                v128 = 2
                _ = fi131
            }
        }
    }
    v127 = v127
    fmt.Println(v128, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v127:", v127)
    fmt.Println("v128:", v128)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v127
    _ = v128
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
}
