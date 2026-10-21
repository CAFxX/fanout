package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 int8; F2 uint }
type GenS4 struct { F0 int; F1 int16; F2 uint64 }
type GenS5 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS6 struct { F0 int8; F1 int32; F2 uint8 }
type GenS7 struct { F0 uint16; F1 uint }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 string; F1 uint16; F2 uint64 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 string; F1 int; F2 uint16 }
type GenS13 struct { F0 uint64 }
func genFn0(p0 int8, p1 int16, p2 uint16) uint {
    v7 := GenS1{F0: (p1 >> ((65535) & 63))}
    if (!(true && true)) {
        if false {
            fmt.Println(p0, p1, p1)
            if (9223372036854775807 < -26) {
                if true {
                    var v8 uint16 = p2
                    var v9 int8 = p0
                    fmt.Println(v9)
                    _ = v8
                    _ = v9
                } else {
                    var v10 int16 = (^(p1))
                    var v11 uint8 = 255
                    switch -318 {
                    case 4294967295:
                        switch -2147483648 {
                        case 2147483647:
                            var v12 int = -425
                            _ = v12
                        case 42:
                            v13 := []int{2147483647}
                            _ = v13
                        }
                    case -1:
                        fmt.Println(v10, p0, v10)
                    case -37:
                        fmt.Println(p0)
                    }
                    _ = v10
                    _ = v11
                }
            } else {
                var v14 int16 = p1
                _ = v14
            }
        } else {
            v15 := []int{32767}
            if len(v15) > 0 {
                v15[0] = 2147483647
            }
            fmt.Println(p0)
            _ = v15
        }
    } else {
        v16 := []int{}
        var v17 int8 = (-(p0))
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        p2 += (p2 % ((0) | 1))
        for fi19 := 0; fi19 < 12; fi19++ {
            v20 := []int{27}
            var v21 string = "foo bar"
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 330
}
func genFn1(p0 int32, p1 int8, p2 int16) int {
    fmt.Println(p0)
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 int = fi22
        _ = fi22
        _ = v23
    }
    if ((int16(p0) > 32767) && false) {
        switch 32768 {
        case 32768:
            v24 := []int{127, -838, -247, 32768}
            _ = v24
        default:
            p0 = (^(p0))
        }
        p1 = 1
        v25 := GenS1{F0: 1}
        _ = v25
    } else {
        for fi26 := 0; fi26 < 6; fi26++ {
            var v27 uint32 = 6
            _ = fi26
            _ = v27
        }
        p2 += (p2 - -32768)
        var v28 float64 = 3.14159
        _ = v28
    }
    p2 = 1
    _ = p0
    _ = p1
    _ = p2
    return 32768
}

func main() {
    var v1 uint16 = 2
    var v2 int16 = 1000
    v3 := make(map[int]bool)
    v4 := make(map[string]int)
    var v5 uint8 = 127
    var v6 float32 = (0.1 - (((0.0 * float32((((3.14159 - 1.5) / 1.5) / -2.25))) - 0.0) * 0.1))
    for fi29 := 0; fi29 < 14; fi29++ {
        v30 := GenS1{F0: v2}
        if true {
            if (0 == 0) {
                v31 := GenS1{F0: v2}
                fmt.Println(v6, v2)
                v32 := []int{fi29, (fi29 % ((genFn1(-1, -128, v2)) | 1))}
                _ = v31
                _ = v32
            } else {
                for fi33 := 0; fi33 < 11; fi33++ {
                    var v34 uint16 = 65535
                    v5 = v5
                    _ = fi33
                    _ = v34
                }
                var v35 bool = (v5 > v5)
                v36 := []int{}
                _ = v35
                _ = v36
            }
            if true {
                if false {
                    if ((((v1 / ((565) | 1)) > 0) && (((384 > 42) || true) || ((1 >= 42) && true))) || (!true)) {
                        for fi37 := 0; fi37 < 3; fi37++ {
                            var v38 uint = 4294967295
                            _ = fi37
                            _ = v38
                        }
                        for fi39 := 0; fi39 < 5; fi39++ {
                            v2 += v2
                            var v40 int = (fi39 >> ((uint(v5)) & 63))
                            var v41 uint16 = 262
                            _ = fi39
                            _ = v40
                            _ = v41
                        }
                        var v42 uint16 = v1
                        _ = v42
                    }
                    if (0 == 0) {
                        var v43 bool = (!true)
                        _ = v43
                    }
                    if (761 < 4294967295) {
                        _ = len(v3)
                        if true {
                            v3[(fi29 & fi29)] = (false && false)
                        } else {
                            fmt.Println(v5)
                        }
                        for fi44 := 0; fi44 < 7; fi44++ {
                            v6 = (v6 + 3.14159)
                            _ = fi44
                        }
                    } else {
                        v45 := []int{fi29}
                        _ = v45
                    }
                }
                var v46 uint16 = (v1 / ((v1) | 1))
                _ = v46
            }
            var v47 uint16 = v1
            _ = v47
        } else {
            v6 -= float32(1e10)
            if true {
                if true {
                    for fi48 := 0; fi48 < 20; fi48++ {
                        if false {
                            fmt.Println(v2)
                        } else {
                            v49 := GenS2{F0: fi29}
                            v50 := []int{(fi29 | (fi29 | 128))}
                            _ = len(v50)
                            _ = v49
                            _ = v50
                        }
                        _ = fi48
                    }
                    if (!(!(int32(v2) >= 0))) {
                        v51 := GenS3{F0: v1, F1: 127, F2: 598}
                        for fi52 := 0; fi52 < 20; fi52++ {
                            var v53 int32 = 0
                            var v54 string = "hello"
                            _ = fi52
                            _ = v53
                            _ = v54
                        }
                        var v55 uint64 = 1
                        _ = v51
                        _ = v55
                    } else {
                        var v56 int8 = -1
                        _ = v56
                    }
                }
                fmt.Println(v2)
                var v57 uint64 = 212
                _ = v57
            }
        }
        _ = fi29
        _ = v30
    }
    if (!(140 != 42)) {
        v58 := GenS4{F0: -103, F1: v2, F2: 2}
        fmt.Println(v6)
        _ = v58
    } else {
        for fi59 := 0; fi59 < 2; fi59++ {
            v60 := v4["\n\t"]
            _ = fi59
            _ = v60
        }
        v2 -= 1000
        if (42 > 0) {
            switch 482 {
            case -538:
                var v61 float64 = 3.14159
                _ = v61
            }
            if (-2147483648 <= 1) {
                v3[127] = (1 > 2147483648)
                v62 := []int{genFn1(-2147483648, 0, (v2 & (v2 | 1)))}
                _ = v62
            } else {
                fmt.Println(v1)
                for fi63 := 0; fi63 < 19; fi63++ {
                    v64 := GenS2{F0: genFn1(-612, -128, (v2 + (v2 & v2)))}
                    _ = fi63
                    _ = v64
                }
            }
        }
    }
    if (4294967295 != 2) {
        v65 := v4[("hello" + "\n\t")]
        v6 = 3.14159
        v4[""] = v65
        _ = v65
    } else {
        v66 := []int{-9223372036854775808, 0, -2147483648}
        v4[""] = genFn1(2147483647, 18, v2)
        v2 = (-(v2))
        _ = v66
    }
    var v67 uint16 = v1
    var v68 int16 = (-(v2))
    for fi69 := 0; fi69 < 3; fi69++ {
        v70 := GenS5{F0: v5, F1: 118, F2: 2147483647}
        v67 += (^(v67))
        _ = fi69
        _ = v70
    }
    fmt.Println(v1, v6, v67)
    var v71 int64 = 0
    var v72 int = -271
    if (-295 > -202) {
        if false {
            v73 := []int{genFn1(935, 127, v68), (v72 + v72)}
            for fi74 := 0; fi74 < 5; fi74++ {
                switch "" {
                case "":
                    v6 = ((-2.25 + -2.25) / v6)
                default:
                    v75 := GenS2{F0: fi74}
                    _ = v75
                }
                _ = fi74
            }
            for fi76 := 0; fi76 < 2; fi76++ {
                switch genFn1(1, 42, 845) {
                case 32768:
                    delete(v4, "foo bar")
                default:
                    if (((float64(1e10) <= (-2.25 / -2.25)) || ((3.14159 + 1.5) > float64(v6))) || ((0.1 / float64(v6)) != 0.1)) {
                        switch fi76 {
                        case -519:
                            v77 := GenS3{F0: (^(v1)), F1: -4, F2: 42}
                            _ = v77
                        case 42:
                            v78 := GenS3{F0: (v1 / (((v67 & 42)) | 1)), F1: 127, F2: 1}
                            _ = v78
                        default:
                            var v79 uint32 = uint32(v5)
                            _ = v79
                        }
                        v5 += 169
                    } else {
                        if false {
                            for fi80 := 0; fi80 < 4; fi80++ {
                                if true {
                                    if false {
                                        var v81 int8 = -33
                                        _ = v81
                                    }
                                    var v82 int = (fi76 % ((fi80) | 1))
                                    var v83 int8 = 0
                                    _ = v82
                                    _ = v83
                                }
                                v6 = (v6 * -2.25)
                                fmt.Println(fi80)
                                var v84 uint32 = 65535
                                v71 += v71
                                v6 += -2.25
                                var v85 float32 = 0.0
                                if (!true) {
                                    v68 = int16(v5)
                                }
                                if (!true) {
                                    var v86 uint32 = 0
                                    v84 -= v86
                                    var v87 string = "\n\t"
                                    _ = v86
                                    _ = v87
                                } else {
                                    var v88 int32 = 42
                                    var v89 int = v72
                                    _ = v88
                                    _ = v89
                                }
                                _ = fi80
                                _ = v84
                                _ = v85
                            }
                        } else {
                            if ((true || false) || (uint(v5) != 4294967295)) {
                                if true {
                                    var v90 string = (("hello" + "hello") + "foo bar")
                                    var v91 uint64 = 42
                                    _ = v90
                                    _ = v91
                                } else {
                                    v72 -= (v72 / ((genFn1(-2147483648, 127, (v2 ^ v68))) | 1))
                                    v92 := GenS6{F0: 127, F1: -1, F2: v5}
                                    _ = v92
                                }
                                v93 := GenS7{F0: 816, F1: genFn0(-43, 32767, uint16(v5))}
                                v5 += v5
                                _ = v93
                            }
                        }
                        var v94 int = v72
                        _ = v94
                    }
                }
                switch "hello" {
                case "\n\t":
                    var v95 uint32 = 2
                    _ = v95
                default:
                    for fi96 := 0; fi96 < 19; fi96++ {
                        if (18446744073709551615 <= 4294967295) {
                            _ = len(v4)
                            v97 := GenS4{F0: 128, F1: v2, F2: 102}
                            _ = v97
                        } else {
                            v1 = (v1 & 931)
                            var v98 float64 = 0.0
                            v68 -= (v68 >> ((0) & 63))
                            v98 -= (float64(float32(v98)) + 0.1)
                            var v99 float64 = -2.25
                            if false {
                                var v100 uint16 = v67
                                var v101 int32 = 172
                                v102 := GenS3{F0: uint16(v5), F1: int8(v68), F2: 65535}
                                _ = v100
                                _ = v101
                                _ = v102
                            } else {
                                if true {
                                    var v103 int32 = 42
                                    switch ("a" + "\n\t") {
                                    case "hello":
                                        fmt.Println(fi96, v99, v5)
                                    case "\n\t":
                                        v104 := GenS4{F0: (^(fi76)), F1: (^(v68)), F2: 18446744073709551615}
                                        if len(v73) > 0 {
                                            v73[0] = (-(fi96))
                                        }
                                        var v105 uint = 65535
                                        _ = v104
                                        _ = v105
                                    case "a":
                                        _ = len(v73)
                                    default:
                                        var v106 int16 = 0
                                        _ = v106
                                    }
                                    _ = v103
                                } else {
                                    var v107 bool = false
                                    if (-128 != 42) {
                                        fmt.Println(fi76, v5, fi96)
                                        v99 = ((v99 - float64(v6)) / 1.0)
                                    } else {
                                        v108 := GenS8{F0: ((("" + "") + "hello") + "foo bar")}
                                        delete(v4, ((("\n\t" + "\n\t") + ("a" + "hello")) + "a"))
                                        var v109 uint16 = v67
                                        _ = v108
                                        _ = v109
                                    }
                                    _ = v107
                                }
                                fmt.Println(v99, v68)
                                if (v71 > v71) {
                                    v5 -= 146
                                    if (v5 > v5) {
                                        v1 = (v67 | (v67 - v67))
                                        v110 := GenS9{F0: 0}
                                        delete(v4, ("\n\t" + ((("" + ("\n\t" + "")) + "") + ("\n\t" + ""))))
                                        if len(v73) >= 2 {
                                            v73 = v73[1:len(v73)]
                                        }
                                        delete(v4, "")
                                        switch fi76 {
                                        case 0:
                                            if len(v73) >= 2 {
                                                v73 = v73[1:len(v73)]
                                            }
                                        case 2147483647:
                                            v111 := GenS1{F0: v68}
                                            _ = v111
                                        default:
                                            v99 = ((float64(float32((-2.25 * 3.14159))) - v98) / v98)
                                        }
                                        v73 = append(v73, fi76)
                                        _ = v110
                                    } else {
                                        switch "foo bar" {
                                        case "":
                                            v72 += fi76
                                        }
                                    }
                                    if true {
                                        var v112 uint = genFn0(0, v68, (v1 & (v67 * 42)))
                                        if (0 == (v112 ^ genFn0(127, (v68 * (v68 << ((v112) & 63))), uint16(v112)))) {
                                            if len(v73) >= 2 {
                                                v73 = v73[1:len(v73)]
                                            }
                                            var v113 int32 = 2147483647
                                            if true {
                                                v1 = v67
                                            } else {
                                                v114 := GenS10{F0: "hello", F1: v1, F2: 2}
                                                _ = v114
                                            }
                                            _ = v113
                                        }
                                        switch genFn1(-739, int8(fi76), (^(v2))) {
                                        case 42:
                                            if (fi76 >= (-(v72))) {
                                                fmt.Println(v67, v1, v71)
                                                v72 += (-(fi96))
                                            } else {
                                                if (!((^(v1)) > 42)) {
                                                    var v115 uint32 = 2
                                                    _ = v115
                                                } else {
                                                    fmt.Println(fi76, v6, v67)
                                                    var v116 float64 = 1.5
                                                    var v117 uint8 = v5
                                                    _ = v116
                                                    _ = v117
                                                }
                                                var v118 float64 = v99
                                                _ = v118
                                            }
                                        }
                                        _ = v112
                                    } else {
                                        if (4294967295 == 4294967295) {
                                            v119 := GenS11{F0: 32767}
                                            switch "foo bar" {
                                            case "\n\t":
                                                v1 = (v67 & v67)
                                            case "hello":
                                                var v120 float32 = v6
                                                _ = v120
                                            case "":
                                                var v121 int32 = 226
                                                v2 = v2
                                                _ = v121
                                            }
                                            v122 := GenS12{F0: "a", F1: 870, F2: (v1 & 435)}
                                            if len(v73) > 0 {
                                                v73[0] = fi76
                                            }
                                            _ = v119
                                            _ = v122
                                        }
                                        var v123 int64 = 52
                                        _ = v123
                                    }
                                    v124 := GenS6{F0: 1, F1: int32(v2), F2: (v5 | 217)}
                                    _ = v124
                                } else {
                                    v125 := GenS13{F0: 280}
                                    _ = v125
                                }
                            }
                            var v126 int16 = v2
                            var v127 int32 = 433
                            _ = v126
                            _ = v127
                            _ = v98
                            _ = v99
                        }
                        if ((-(fi76)) < 0) {
                            delete(v3, (fi76 << ((2) & 63)))
                            v68 = 1
                            fmt.Println(v1)
                        } else {
                            v68 += int16(v5)
                            if len(v73) >= 2 {
                                v73 = v73[1:len(v73)]
                            }
                            v2 = v68
                        }
                        _ = fi96
                    }
                }
                v72 = fi76
                _ = fi76
            }
            _ = v73
        } else {
            var v128 uint64 = 1
            _ = v128
        }
        v71 = (^(v71))
        var v129 uint64 = 22
        _ = v129
    }
    v130 := v3[(v72 >> ((544) & 63))]
    v131 := GenS1{F0: v2}
    for fi132 := 0; fi132 < 7; fi132++ {
        for fi133 := 0; fi133 < 2; fi133++ {
            var v134 uint32 = 65535
            v1 -= (v67 - (v67 << ((genFn0(86, (v68 & (v2 - 1)), 42)) & 63)))
            var v135 int8 = -128
            _ = fi133
            _ = v134
            _ = v135
        }
        var v136 string = (((("hello" + "") + ((("hello" + "") + ("foo bar" + "foo bar")) + ("foo bar" + "\n\t"))) + "foo bar") + ("\n\t" + ("" + "\n\t")))
        if false {
            if v130 {
                if ((v71 <= int64(v2)) || ((!(!true)) && (v72 < (v72 << ((36) & 63))))) {
                    v1 -= v1
                    v130 = v130
                }
                fmt.Println(v1, v1)
                delete(v4, "foo bar")
            } else {
                fmt.Println(v2, v72, v71)
                var v137 int32 = -531
                fmt.Println(v72)
                _ = v137
            }
        } else {
            switch int(v68) {
            case 42:
                v138 := v4["\n\t"]
                _ = v138
            case 4294967295:
                switch v136 {
                case "foo bar":
                    switch (fi132 * genFn1(-399, 1, v68)) {
                    case 127:
                        v71 = (-(v71))
                    case -120:
                        v139 := GenS3{F0: 65535, F1: -37, F2: genFn0(0, v68, v67)}
                        _ = v139
                    case 806:
                        fmt.Println(v72, v67)
                    }
                case "":
                    v140 := []int{}
                    _ = v140
                case "\n\t":
                    fmt.Println(v71, v6)
                }
            case 891:
                var v141 int16 = (v2 << ((0) & 63))
                _ = v141
            }
            v68 = (v2 >> ((0) & 63))
            fmt.Println(v67)
        }
        _ = fi132
        _ = v136
    }
    for fi142 := 0; fi142 < 12; fi142++ {
        v71 += (^(v71))
        var v143 bool = ((^(v5)) > (^(v5)))
        _ = fi142
        _ = v143
    }
    v144 := GenS5{F0: 0, F1: 0, F2: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v130:", v130)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v130
    _ = v131
    _ = v144
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v67
    _ = v68
    _ = v71
    _ = v72
}
